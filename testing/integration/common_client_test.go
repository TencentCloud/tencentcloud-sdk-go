package integration

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
	"testing"
)

// TestCommonRequest
//  目前只支持 签名v3+POST
func TestCommonRequest(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf)

	// 创建common request
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")
	//自定义请求参数,SetActionParameters 函数支持三种数据类型的入参
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
	t.Log(string(response.GetBody()))
}
