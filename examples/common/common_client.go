package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
)

// 目前只支持 签名v3+POST
func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf)

	// 创建common request，依次传入产品名、产品版本、接口名称
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")

	// 自定义请求参数:
	// SetActionParameters 函数支持三种数据类型的入参：
	// 1. map[string]interface{}
	// body:=map[string]interface{}{
	// "InstanceId":"crs-xxx",
	// "SpanType":2,
	// }
	body := map[string]interface{}{}

	// // 2. string
	// bodyStr := `{}`

	// // 3. []byte
	// bodyBytes := []byte(bodyStr)

	// set custom headers
	request.SetHeader(map[string]string{
		"X-TC-TraceId": "ffe0c072-8a5d-4e17-8887-a8a60252abca",
	})

	// 设置action所需的请求数据
	err := request.SetActionParameters(body)
	if err != nil {
		return
	}

	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err = client.Send(request, response)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}

	// 获取响应结果
	fmt.Println(string(response.GetBody()))
}
