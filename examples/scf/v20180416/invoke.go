package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	scf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

func main() {
	// 从环境变量获取SID和SKEY实例化认证对象
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	// 实例化客户端配置对象
	cpf := profile.NewClientProfile()
	// 实例化SCF的client对象
	client, _ := scf.NewClient(credential, regions.Guangzhou, cpf)

	// API -- Invoke:运行SCF上的函数
	invokeRequest := scf.NewInvokeRequest()
	// 函数名称
	invokeRequest.FunctionName = common.StringPtr("sum")
	// RequestResponse(同步) 和 Event(异步)，默认为同步
	invokeRequest.InvocationType = common.StringPtr("RequestResponse")
	// 运行函数时的参数，以json格式传入，最大支持的参数长度是 1M
	invokeRequest.ClientContext = common.StringPtr(`{"key1": 1, "key2": 1}`)
	// 同步调用时指定该字段，返回值会包含4K的日志，可选值为None和Tail，默认值为None。当该值为Tail时，返回参数中的logMsg字段会包含对应的函数执行日志
	invokeRequest.LogType = common.StringPtr("Tail")
	invokeResponse, err := client.Invoke(invokeRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(invokeResponse.ToJsonString())
}
