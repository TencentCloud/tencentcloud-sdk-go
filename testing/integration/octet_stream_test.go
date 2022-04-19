package integration

import (
	"io/ioutil"
	"testing"

	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func TestOctetStreamAction(t *testing.T) {
	cr := getCredential(t)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := cls.NewClient(cr, regions.Guangzhou, cpf)
	request := cls.NewUploadLogRequest()
	request.TopicId = common.StringPtr("f6c4fa6f-367a-4f14-8289-1ff6f77ed975")
	request.HashKey = common.StringPtr("0fffffffffffffffffffffffffffffff")
	request.CompressType = common.StringPtr("")

	data, _ := ioutil.ReadFile("./binary.data")
	_, err := client.UploadLog(request, data)
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		if terr.GetCode() == "OperationDenied" || terr.GetCode() == "ResourceNotFound.TopicNotExist" {
			return
		}
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
