package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	trtc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc/v20190722"
	"os"
)

func main() {
	//SDK 使用 `omitempty` 标签来序列化你的 request 对象, 因为这样可以避免上报空数组/对象.
	//但对有的接口而言, 长度为0的数组 和 nil数组 是有区别的, 如果你希望在请求中携带空数组, 需要使用Common Client 来发送请求.

	// 错误的做法
	wrongWay()

	// 正确的做法
	rightWay()
}

func wrongWay() {
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

func rightWay() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.Debug = true

	client := common.NewCommonClient(credential, regions.Guangzhou, cpf)
	request := tchttp.NewCommonRequest("trtc", "2019-07-22", "UpdatePublishCdnStream")
	req := map[string]interface{}{
		"PublishCdnParams": []*trtc.McuPublishCdnParam{},
	}

	// 发送的请求body为 `{"PublishCdnParams":[]}`
	err := request.SetActionParameters(req)
	if err != nil {
		panic(err)
	}

	response := tchttp.NewCommonResponse()

	err = client.Send(request, response)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}

	fmt.Println(string(response.GetBody()))
}
