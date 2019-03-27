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
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	credential := common.NewCredential(config.SecretID, config.SecretKey)

	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// cpf.HttpProfile.ReqMethod = "GET"
	// cpf.HttpProfile.ReqTimeout = 10
	// v2 版本必须设置
	cpf.HttpProfile.Endpoint = config.TopicEndpoint
	// cpf.SignMethod = "HmacSHA1"

	// 实例化要请求产品的client对象
	client, _ := cmq.NewClient(credential, cpf)

	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := cmq.NewCreateTopicRequest()
	request.TopicName = "sdk_create_topic_post"
	request.MaxMsgSize = 4096
	// request.FilterType = cmq.FilterTypeRouting

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.CreateTopic(request)
	// // 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		fmt.Printf("API error : %v", err)
		panic(err)
	}
	// 打印返回结果
	// fmt.Printf("%s", response.ToJsonString())
	fmt.Printf("response: %+v", response)
}
