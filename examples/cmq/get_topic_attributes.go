package main

import (
	"fmt"
	"time"

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
	request := cmq.NewGetTopicAttributesRequest()
	request.TopicName = common.StringPtr("sdk_create_topic")
	response, err := client.GetTopicAttributes(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("error : %v", err)
		panic(err)
	}
	fmt.Printf("response: %v\n", common.ToJsonString(response))
	fmt.Printf("createTime: %v, LastModifyTime: %v", time.Unix(
		int64(*response.CreateTime), 0), time.Unix(int64(*response.LastModifyTime), 0))
}
