package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf)
	// 创建common request
	headers := map[string]string{
		"X-CLS-TopicId":      "cde920aa-61f2-4569-a2c0-d01ecef65c55",
		"X-CLS-HashKey":      "0fffffffffffffffffffffffffffffff",
		"X-CLS-CompressType": "",
	}
	dir, _ := os.Getwd()
	body, err := ioutil.ReadFile(path.Join(dir, "examples", "cls", "binary.data"))
	if err != nil {
		fmt.Println(err)
		return
	}
	request := tchttp.NewCommonRequest("cls", "2020-10-16", "UploadLog")
	request.SetOctetStreamParameters(headers, body)
	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err = client.SendOctetStream(request, response)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail to invoke api: %v", err))
		return
	}
	fmt.Println(string(response.GetBody()))

}
