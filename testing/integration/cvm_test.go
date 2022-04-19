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
	"fmt"
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestDescribeInstancesSignV1Get(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.SignMethod = "HmacSHA1"
	client, err := cvm.NewClient(
		credential,
		regions.Guangzhou,
		cpf)
	if err != nil {
		t.Fatalf(fmt.Sprintf("fail to init client: %v", err))
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("zone"),
			Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
		},
	}
	_, err = client.DescribeInstances(request)
	if err != nil {
		t.Fatalf("fail to invoke api: %v", err)
	}
}

func TestEmptyStringGetSignV1HmacSHA1(t *testing.T) {
	testEmptyStringGet(t, "HmacSHA1")
}

func TestEmptyStringGetSignV1HmacSHA256(t *testing.T) {
	testEmptyStringGet(t, "HmacSHA256")
}

func TestEmptyStringGetSignV3(t *testing.T) {
	testEmptyStringGet(t, "TC3-HMAC-SHA256")
}

func testEmptyStringGet(t *testing.T, method string) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.SignMethod = method
	client, err := cvm.NewClient(
		credential,
		regions.Guangzhou,
		cpf)
	if err != nil {
		t.Fatalf("fail to init client: %v", err)
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("zone"),
			Values: common.StringPtrs([]string{""}),
		},
	}
	_, err = client.DescribeInstances(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		if terr.GetCode() == cvm.INVALIDZONE_MISMATCHREGION {
			return
		}
	}
	t.Fatalf("not expected error: %v", err)
}

func TestDescribeInstances(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	client, err := cvm.NewClient(
		credential,
		regions.Guangzhou,
		profile.NewClientProfile())
	if err != nil {
		t.Fatalf("fail to init client: %v", err)
	}

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatalf("fail to invoke api: %v", err)
	}
	if *response.Response.TotalCount != (int64)(len(response.Response.InstanceSet)) {
		t.Fatalf("response item count inconsisitent!")
	}
}
