package common

import (
	"fmt"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
	"testing"
)

// Copyright 1999-2021 Tencent Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TestCommonRequest
//  目前只支持 签名v3+POST
func TestCommonRequest(t *testing.T) {
	credential := NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "redis.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	//创建common client
	client := NewCommonClient(credential, regions.Guangzhou, cpf)

	// 创建common request
	request := tchttp.NewCommonRequest("redis", "2018-04-12", "DescribeInstanceMonitorTopNCmdTook")
	//自定义请求参数,支持三种数据类型的入参
	// 1. map[string]interface{}
	//body:=map[string]interface{}{
	//	"InstanceId":"crs-69nvlon9",
	//	"SpanType":2,
	//}

	// 2. string
	bodyStr := `{
	"InstanceId":"crs-xxx",
	"SpanType":2
}`

	// 3. []byte
	bodyBytes := []byte(bodyStr)

	// 设置action所需的请求数据
	err := request.SetActionRequest(bodyBytes)
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
	t.Log(response.GetBody().Raw())
	// 返回[]byte
	// t.Log(response.actionRequest.Raw())
}
