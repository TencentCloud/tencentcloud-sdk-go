package main

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
    // 仅填写域名，不包含http(s)前缀，和URL Path
    cpf.HttpProfile.ApigwEndpoint = "service-1q2w3e4r-12345678.gz.apigw.tencentcs.com"
    // 需要和apigw后端配置中指定的保持一致
    cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"

	// 实例化要请求产品(以cvm为例)的client对象
	client, _ := cvm.NewClient(credential, "ap-beijing", cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := cvm.NewDescribeZonesRequest()
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeZones(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Println(response.ToJsonString())
}
