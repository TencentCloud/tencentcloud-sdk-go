package main

import (
	"encoding/json"
	"fmt"
	"os"

	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 300 // 流式接口可能耗时较长，因此将请求超时时间设置为 300 秒。
	// 通过 SSE 流式调用此接口时，请务必设置请求域名（Endpoint）为 cls.ai.tencentcloudapi.com。
	cpf.HttpProfile.Endpoint = "cls.ai.tencentcloudapi.com"

	client, _ := cls.NewClient(credential, regions.Guangzhou, cpf)

	request := cls.NewChatCompletionsRequest()

	// 指定要使用的功能，例如 "text2sql"。
	request.Model = common.StringPtr("text2sql")

	// 设置对话消息，角色为 "user"，内容为用户的提问。
	request.Messages = []*cls.Message{
		{
			Role:    common.StringPtr("user"),
			Content: common.StringPtr("状态码200, 统计日志条数"),
		},
	}

	// 设置元数据，传递日志主题地域和日志主题 ID。
	request.Metadata = []*cls.MetadataItem{
		{
			Key:   common.StringPtr("topic_region"),
			Value: common.StringPtr(regions.Guangzhou),
		},
		{
			Key:   common.StringPtr("topic_id"),
			Value: common.StringPtr("xxxxxxxx-xxxx"),
		},
	}

	// CLS ChatCompletions 同时支持流式和非流式两种模式。
	// 将 Stream 参数设置为 true，表示以流式方式获取响应。
	request.Stream = common.BoolPtr(true)

	response, err := client.ChatCompletions(request)

	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	// 非 SDK 异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	if *request.Stream {
		// 流式调用示例
		for event := range response.Events {
			if event.Err != nil {
				fmt.Printf("\n[流式] stream error: %v\n", event.Err)
				break
			}

			// 跳过心跳等 data 为空的事件（如服务端每隔一段时间发送的 :heartbeat 注释）
			if len(event.Data) == 0 {
				continue
			}

			// 服务端通过 event: error 返回应用层错误
			if event.Event == "error" {
				var errResp tchttp.ErrorResponse
				if err := json.Unmarshal(event.Data, &errResp); err != nil {
					fmt.Printf("\n[流式] json unmarshal error event failed: %v\n", err)
				} else {
					fmt.Printf("\n[流式] error code: %s, message: %s\n",
						errResp.Response.Error.Code,
						errResp.Response.Error.Message)
				}
				continue
			}

			var data cls.ChatCompletionsResponseParams
			if err := json.Unmarshal(event.Data, &data); err != nil {
				fmt.Printf("\n[流式] json unmarshal error: %v\n", err)
				continue
			}

			for _, choice := range data.Choices {
				// 实时输出思考过程（ReasoningContent）
				if choice.Delta.ReasoningContent != nil && *choice.Delta.ReasoningContent != "" {
					fmt.Print(*choice.Delta.ReasoningContent)
				}
				// 实时输出回复内容（Content）
				if choice.Delta.Content != nil && *choice.Delta.Content != "" {
					fmt.Print(*choice.Delta.Content)
				}
			}
		}
		fmt.Println()
		fmt.Println("[流式] 完成，开始非流式调用...")
	}

	// 非流式调用示例
	// 通过 Stream=false 参数来指定非 stream 协议，一次性拿到结果。
	request.Stream = common.BoolPtr(false)
	response2, err := client.ChatCompletions(request)

	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}

	if response2.Response != nil && len(response2.Response.Choices) > 0 {
		msg := response2.Response.Choices[0].Message
		if msg.ReasoningContent != nil && *msg.ReasoningContent != "" {
			fmt.Println("[非流式] 思考过程：", *msg.ReasoningContent)
		}
		fmt.Println("[非流式] 回复内容：\n" + *msg.Content)
	}
}
