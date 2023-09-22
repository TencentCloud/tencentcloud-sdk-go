package common

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"io"
	"io/ioutil"
	"log"
	//"log"
	"net/http"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type StreamResponse interface {
	streamResponseEmbed()
}

type SSEvent struct {
	Event string
	Data  []byte
	Id    string
	Retry int
}

type SSEResponse interface {
	Response
	setEventsChannel(ch chan SSEvent)
}

type BaseSSEResponse struct {
	BaseResponse
	Events chan SSEvent
}

func (r *BaseSSEResponse) setEventsChannel(ch chan SSEvent) {
	r.Events = ch
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func TryReadErr(resp *http.Response) (err error) {
	switch resp.Header.Get("Content-Type") {
	case "text/event-stream", "application/octet-stream":
		return nil
	}

	buf := bytes.NewBuffer(nil)
	tee := io.TeeReader(resp.Body, buf)

	errResp := &ErrorResponse{}
	err = json.NewDecoder(tee).Decode(&errResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", buf.String(), err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if errResp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(errResp.Response.Error.Code, errResp.Response.Error.Message, errResp.Response.RequestId)
	}

	depResp := &DeprecatedAPIErrorResponse{}
	err = json.NewDecoder(buf).Decode(&depResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", buf.String(), err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if depResp.Code != 0 {
		return errors.NewTencentCloudSDKError(depResp.CodeDesc, depResp.Message, "")
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, resp Response) error {
	switch hr.Header.Get("Content-Type") {
	case "text/event-stream":
		return parseFromSSE(hr, resp)
	default:
		return parseFromJson(hr, resp)
	}
}

func parseFromJson(hr *http.Response, resp Response) error {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewTencentCloudSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewTencentCloudSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	err = resp.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	return nil
}

func parseFromSSE(hr *http.Response, resp Response) error {
	reqId := hr.Header.Get("X-TC-RequestId")
	r, ok := resp.(SSEResponse)
	if !ok {
		return errors.NewTencentCloudSDKError("ClientError.TypeError",
			"Response type does not implement SSEResponse", reqId)
	}

	ch := make(chan SSEvent)
	r.setEventsChannel(ch)

	// parser
	go func() {
		defer hr.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(hr.Body)
		scanner.Split(bufio.ScanLines)

		event := SSEvent{}
		for scanner.Scan() {
			line := scanner.Bytes()

			// SSE use empty line to indicate message end
			if len(line) == 0 {
				ch <- event
				event = SSEvent{}
				continue
			}

			idx := bytes.IndexByte(line, ':')
			if idx == -1 {
				log.Printf("ServerError: SSE received invalid line,%s", line)
				break
			}

			key := string(line[:idx])
			val := line[idx+1:]
			switch key {
			case "event":
				event.Event = string(val)
			case "data":
				event.Data = append(event.Data, val...)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("ClientError: " + err.Error())
		}
	}()

	return nil
}
