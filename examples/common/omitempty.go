package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	trtc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc/v20190722"
	"os"
)

func main() {
	// SDK 使用 `omitnil` 标签来序列化你的 request 对象, 因为这样可以区分 `nil数组` 和 `长度为0的数组`

	// SDK 默认会发送 `长度为0的数组` 而忽略 `nil数组`
	sendJsonRequest()

	// 当你不希望发送一个 `长度为0的数组` 时, 可以通过 json.OmitBehaviour 来关闭此特性
	json.OmitBehaviour = json.OmitEmpty
	sendJsonRequest()
}

func sendJsonRequest() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.Debug = true

	client, err := trtc.NewClient(credential, regions.Guangzhou, cpf)

	req := trtc.NewUpdatePublishCdnStreamRequest()

	// 因为 omitempty 的关系, 请求中不会包含 PublishCdnParams 字段
	// 发送的请求body为 `{}`
	req.PublishCdnParams = []*trtc.McuPublishCdnParam{}

	response, err := client.UpdatePublishCdnStream(req)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}

	fmt.Println(response.ToJsonString())
}
