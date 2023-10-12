package common

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

const (
	codeLimitExceeded = "RequestLimitExceeded"
	tplRateLimitRetry = "[WARN] rate limit exceeded, retrying (%d/%d) in %f seconds: %s"
)

func (c *Client) sendWithRateLimitRetry(req *http.Request, retryable bool) (resp *http.Response, err error) {
	// make sure maxRetries is more than 0
	maxRetries := maxInt(c.profile.RateLimitExceededMaxRetries, 0)
	durationFunc := safeDurationFunc(c.profile.RateLimitExceededRetryDuration)

	for idx := 0; idx <= maxRetries; idx++ {
		resp, err = c.sendWithNetworkFailureRetry(req, retryable)
		if err != nil {
			return
		}

		err = decompressBodyReader(resp)
		if err != nil {
			return resp, err
		}

		err = tchttp.TryReadErr(resp)
		// should not sleep on last request
		if err, ok := err.(*errors.TencentCloudSDKError); ok && err.Code == codeLimitExceeded && idx < maxRetries {
			duration := durationFunc(idx)
			if c.debug {
				c.logger.Printf(tplRateLimitRetry, idx, maxRetries, duration.Seconds(), err.Error())
			}

			time.Sleep(duration)
			continue
		}

		return resp, err
	}

	return resp, err
}

func decompressBodyReader(resp *http.Response) error {
	var readCloser io.ReadCloser
	var err error

	enc := resp.Header.Get("Content-Encoding")
	switch enc {
	case "":
		readCloser = resp.Body
	case "deflate":
		readCloser = flate.NewReader(resp.Body)
	case "gzip":
		readCloser, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Content-Encoding not support: %s", enc)
	}

	resp.Body = readCloser
	// delete the header in case the caller mistake the body being encoded
	delete(resp.Header, "Content-Encoding")
	return nil
}
