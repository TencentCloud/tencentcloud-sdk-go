package integration

import (
	"encoding/json"
	"fmt"
	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	"io/ioutil"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func getCredential() common.CredentialIface {
	pr := common.DefaultProviderChain()
	cr, err := pr.GetCredential()
	if err != nil {
		panic(err)
	}
	return cr
}

// TestCommonRequest
// 目前只支持 签名v3+POST
func TestCommonRequest(t *testing.T) {
	cr := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(cr, regions.Guangzhou, cpf)

	// 创建common request
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")
	// 自定义请求参数,SetActionParameters 函数支持三种数据类型的入参
	// 1. map[string]interface{}
	//body:=map[string]interface{}{
	//	"InstanceId":"crs-xxx",
	//	"SpanType":2,
	//}

	// 2. string
	bodyStr := `{}`

	// 3. []byte
	bodyBytes := []byte(bodyStr)

	// 设置action所需的请求数据
	err := request.SetActionParameters(bodyBytes)
	if err != nil {
		return
	}

	//创建common response
	response := tchttp.NewCommonResponse()

	//发送请求
	err = client.Send(request, response)
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
	}
}

func TestClient_SendOctetStream(t *testing.T) {
	cr := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	// 创建common client
	client := common.NewCommonClient(cr, regions.Guangzhou, cpf)
	// 创建common request
	request := tchttp.NewCommonRequest("cls", "2020-10-16", "UploadLog")
	headers := map[string]string{
		"X-CLS-TopicId":      "f6c4fa6f-367a-4f14-8289-1ff6f77ed975",
		"X-CLS-HashKey":      "0fffffffffffffffffffffffffffffff",
		"X-CLS-CompressType": "",
	}
	body, _ := ioutil.ReadFile("./binary.data")

	request.SetOctetStreamParameters(headers, body)
	// 创建common response
	response := tchttp.NewCommonResponse()

	// 发送请求
	err := client.SendOctetStream(request, response)
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		if terr.GetCode() == "OperationDenied" || terr.GetCode() == "ResourceNotFound.TopicNotExist" {
			return
		} else {
			t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))

		}
	}
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
	}
	t.Log(string(response.GetBody()))
}

func TestOctetStreamAction(t *testing.T) {
	cr := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := cls.NewClient(cr, regions.Guangzhou, cpf)
	request := cls.NewUploadLogRequest()
	request.TopicId = common.StringPtr("f6c4fa6f-367a-4f14-8289-1ff6f77ed975")
	request.HashKey = common.StringPtr("0fffffffffffffffffffffffffffffff")
	request.CompressType = common.StringPtr("")

	body, _ := ioutil.ReadFile("./binary.data")
	request.SetOctetStreamBody(body)

	response, err := client.UploadLog(request)
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		if terr.GetCode() == "OperationDenied" || terr.GetCode() == "ResourceNotFound.TopicNotExist" {
			return
		} else {
			t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
		}
	}
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
	}
	b, _ := json.Marshal(response.Response)
	t.Log(b)
}
