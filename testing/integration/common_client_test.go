package integration

import (
	"io/ioutil"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func getCredential(t *testing.T) common.CredentialIface {
	pr := common.DefaultProviderChain()
	cr, err := pr.GetCredential()
	if err != nil {
		t.Fatal(err)
	}
	return cr
}

// TestCommonRequest
// Currently only supports signature v3+POST
func TestCommonRequest(t *testing.T) {
	cr := getCredential(t)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client := common.NewCommonClient(cr, regions.Guangzhou, cpf)

	// create common request
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")
	// To customize request parameters, the SetActionParameters function supports input parameters of three data types
	// 1. map[string]interface{}
	// body:=map[string]interface{}{
	//	"InstanceId":"crs-xxx",
	//	"SpanType":2,
	// }

	// 2. string
	bodyStr := `{}`

	// 3. []byte
	bodyBytes := []byte(bodyStr)

	// Set the request data required by the action
	err := request.SetActionParameters(bodyBytes)
	if err != nil {
		return
	}

	// create common response
	response := tchttp.NewCommonResponse()

	// send request
	err = client.Send(request, response)
	if err != nil {
		t.Errorf("fail to invoke api: %v", err)
	}
}

func TestClient_SendOctetStream(t *testing.T) {
	cr := getCredential(t)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	// create common client
	client := common.NewCommonClient(cr, regions.Guangzhou, cpf)
	// create common request
	request := tchttp.NewCommonRequest("cls", "2020-10-16", "UploadLog")
	headers := map[string]string{
		"X-CLS-TopicId":      "f6c4fa6f-367a-4f14-8289-1ff6f77ed975",
		"X-CLS-HashKey":      "0fffffffffffffffffffffffffffffff",
		"X-CLS-CompressType": "",
	}
	body, _ := ioutil.ReadFile("./binary.data")

	request.SetOctetStreamParameters(headers, body)
	//  create common response
	response := tchttp.NewCommonResponse()

	// send request
	err := client.SendOctetStream(request, response)
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		if terr.GetCode() == "OperationDenied" || terr.GetCode() == "ResourceNotFound.TopicNotExist" {
			return
		}
	}
	if err != nil {
		t.Errorf("fail to invoke api: %v", err)
	}
}
