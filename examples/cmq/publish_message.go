package main

import (
	"fmt"

	"github.com/teamlint/tencentcloud-sdk-go/examples/cmq/config"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/cmq"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	credential := common.NewCredential(config.SecretID, config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.TopicEndpoint
	client, _ := cmq.NewClient(credential, cpf)

	request := cmq.NewPublishMessageRequest()
	// 标签过滤
	// request.TopicName = "demo_topic_tag"
	// request.MsgBody = "sdk发送消息内容1"
	// request.MsgTag = []string{
	// 	"pc",
	// 	"wx",
	// }
	// 路由过滤
	request.TopicName = "demo_topic"
	request.MsgBody = "sdk发送路由消息"
	request.RoutingKey = "order.confirm1"

	response, err := client.PublishMessage(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("error : %v", err)
		panic(err)
	}
	fmt.Printf("response: %+v", response)
}
