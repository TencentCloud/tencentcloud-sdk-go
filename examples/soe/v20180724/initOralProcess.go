package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	soe "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe/v20180724"
)

func main() {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		// os.Getenv("TENCENTCLOUD_SECRET_ID"),
		// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"...",
		"...",
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	cpf.HttpProfile.ReqMethod = "POST"
	// SDK有默认的超时时间，非必要请不要进行调整。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 30
	// SDK会自动指定域名。通常是不需要特地指定域名的
	cpf.HttpProfile.Endpoint = "soe.tencentcloudapi.com"

	// 实例化要请求产品的client对象
	// 第二个参数是地域信息
	client, _ := soe.NewClient(credential, "ap-guangzhou", cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定InitOralProcessRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := soe.NewInitOralProcessRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.WorkMode = common.Int64Ptr(1)
	request.EvalMode = common.Int64Ptr(0)
	request.ScoreCoeff = common.Float64Ptr(2.0)
	request.SessionId = common.StringPtr("1")
	request.RefText = common.StringPtr("since")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.InitOralProcess(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
