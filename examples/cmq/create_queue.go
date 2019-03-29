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
	cpf.HttpProfile.Endpoint = config.QueueEndpoint
	client, _ := cmq.NewClient(credential, cpf)
	request := cmq.NewCreateQueueRequest()
	request.QueueName = common.StringPtr("sdk_create_queue_t1")
	request.MaxMsgHeapNum = common.IntPtr(1100000)
	request.PollingWaitSeconds = common.IntPtr(8)
	request.VisibilityTimeout = common.IntPtr(9)
	request.MaxMsgSize = common.IntPtr(3072)
	request.MsgRetentionSeconds = common.IntPtr(4096)
	// 最大值不超过MsgRetentionSeconds
	request.RewindSeconds = common.IntPtr(4095)
	response, err := client.CreateQueue(request)
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
