package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
	"os"
)

func main() {
	// hunyuan ChatStd/ChatPro 同时支持 stream 和非 stream 的情况

	// stream 示例
	stream()

	// 非 stream 示例
	// 通过 Stream=false 参数来指定非 stream 协议, 一次性拿到结果
	nonStream()
}

func stream() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	request := hunyuan.NewChatStdRequest()

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("你好, 可以讲个笑话吗"),
	}
	request.Messages = append(request.Messages, msg)

	response, err := client.ChatStd(request)

	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	for event := range response.Events {
		fmt.Println(string(event.Data))
	}
}

func nonStream() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	request := hunyuan.NewChatStdRequest()

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("你好, 可以讲个笑话吗"),
	}
	request.Messages = append(request.Messages, msg)

	request.Stream = common.BoolPtr(false)

	response, err := client.ChatStd(request)

	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	fmt.Println(response.ToJsonString())
}
