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

package integration

import (
	"os"
	"strings"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

func toGetCredential(t *testing.T) (string, string, string) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := sts.NewClient(credential, "ap-guangzhou", cpf)

	request := sts.NewAssumeRoleRequest()
	request.RoleArn = common.StringPtr(os.Getenv("TENCENTCLOUD_ROLE_ARN"))
	if *request.RoleArn == "" {
		t.Skipf("TENCENTCLOUD_ROLE_ARN env var not set")
	}
	request.RoleSessionName = common.StringPtr("cloudapi-test")

	response, err := client.AssumeRole(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		t.Fatalf("unexpected error has returned: %s", err)
	}
	if err != nil {
		t.Fatalf("fail to init client: %v", err)
	}

	tmpId := *response.Response.Credentials.TmpSecretId
	tmpKey := *response.Response.Credentials.TmpSecretKey
	tmpToken := *response.Response.Credentials.Token
	return tmpId, tmpKey, tmpToken
}

func testCredRequestSuccess(t *testing.T, reqMethodList []string) {
	for _, reqMethod := range reqMethodList {
		id, key, token := toGetCredential(t)
		credential := common.NewTokenCredential(id, key, token)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.ReqMethod = reqMethod
		client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

		request := cvm.NewDescribeZonesRequest()

		_, err := client.DescribeZones(request)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

}

func testCredRequestFail(t *testing.T, reMethodList []string) {
	for _, reqMethod := range reMethodList {
		id, key, token := toGetCredential(t)
		credential := common.NewTokenCredential(id, key, token+"error")
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.ReqMethod = reqMethod
		client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

		request := cvm.NewDescribeInstancesOperationLimitRequest()
		request.InstanceIds = common.StringPtrs([]string{" "})
		request.Operation = common.StringPtr("INSTANCE_DEGRADE")

		_, err := client.DescribeInstancesOperationLimit(request)
		if err == nil {
			t.Fatalf("unexpected success")
		}
		if strings.Index(err.Error(), "AuthFailure.TokenFailure") == -1 {
			t.Fatalf("The error code is not as expected!")
		}
	}
}

func TestTempCredentials(t *testing.T) {
	reqMethodList := []string{"GET", "POST"}
	testCredRequestFail(t, reqMethodList)
	testCredRequestSuccess(t, reqMethodList)
}
