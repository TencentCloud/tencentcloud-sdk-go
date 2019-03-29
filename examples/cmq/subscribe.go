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
	client, _ := cmq.NewClient(credential, cpf)
	request := cmq.NewSubscribeRequest()
	request.TopicName = common.StringPtr("sdk_create_topic_post") // tag filter
	// request.TopicName = common.StringPtr("sdk_create_topic") // binding key
	request.SubscriptionName = common.StringPtr("sdk_create_sub")
	request.Protocol = common.StringPtr(cmq.SubscribeProtocolHTTP)
	request.Endpoint = common.StringPtr("http://cmq.abc.com/sdk")
	request.NotifyStrategy = common.StringPtr(cmq.NotifyStrategyBackoffRetry)
	request.NotifyContentFormat = common.StringPtr(cmq.NotifyContentFormatJSON)
	request.FilterTag = common.StringPtrs([]string{
		"tag1",
		"tag2",
	})
	request.BindingKey = common.StringPtrs([]string{
		"sdk.one",
		"sdk.two",
	})
	response, err := client.Subscribe(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("error : %v", err)
		panic(err)
	}
	fmt.Printf("response: %v\n", common.ToJsonString(response))
}
