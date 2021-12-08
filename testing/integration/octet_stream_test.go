package integration

import (
	"encoding/json"
	"fmt"
	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"io/ioutil"
	"testing"
)

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

	data, _ := ioutil.ReadFile("./binary.data")
	response, err := client.UploadLog(request, data)
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
