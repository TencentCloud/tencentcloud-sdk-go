package main

import (
	"fmt"

	"github.com/teamlint/tencentcloud-sdk-go/examples/cmq/config"
	cmq "github.com/teamlint/tencentcloud-sdk-go/tencentcloud/cmq/v2"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	credential := common.NewCredential(config.SecretID, config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.TopicEndpoint
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.HttpProfile.Protocol = "http"
	client, _ := cmq.NewClient(credential, cpf)

	request := cmq.NewBatchPublishMessageRequest()
	// 标签过滤
	// request.TopicName = common.StringPtr("demo_topic_tag")
	// request.MsgBody = common.StringPtr("sdk发送消息内容2")
	// request.MsgTag = common.StringPtrs([]string{
	// 	"none",
	// })

	// 路由过滤
	request.TopicName = common.StringPtr("demo_topic")
	msgs := []string{
		"sdk批量发送路由消息3",
		"sdk批量发送路由消息4",
	}
	request.MsgBody = common.StringPtrs(msgs)
	request.RoutingKey = common.StringPtr("order.abc")

	response, err := client.BatchPublishMessage(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("error : %v", err)
		panic(err)
	}
	json := common.ToJsonString(response)
	fmt.Printf("response: %+v", json)
	resp := &cmq.PublishMessageResponse{}
	common.FromJsonString(resp, json)
	fmt.Printf("response: %+v", resp)
}
