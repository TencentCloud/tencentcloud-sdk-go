package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func main() {
	pc := common.DefaultProviderChain()
	cre, err := pc.GetCredential()
	if err != nil {
		fmt.Println("get cre failed!", err)
		return
	}
	fmt.Println("get cre success!")
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(cre, regions.Guangzhou, cpf)
	// 创建common request
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")

	bodyStr := `{}`

	// 设置action所需的请求数据
	err = request.SetActionParameters(bodyStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err = client.Send(request, response)
	if err != nil {
		fmt.Println("fail to invoke api: ", err)
		return
	}
	fmt.Println(string(response.GetBody()))
}
