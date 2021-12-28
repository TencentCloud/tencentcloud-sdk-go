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
	"strings"
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
		t.Errorf(fmt.Sprintf("fail to init client: %v", err))
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("zone"),
			Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
		},
	}
	resp, err := client.DescribeInstances(request)
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
	}
	fmt.Printf("%s\n", resp.ToJsonString())
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
		t.Errorf(fmt.Sprintf("fail to init client: %v", err))
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("zone"),
			Values: common.StringPtrs([]string{""}),
		},
	}
	resp, err := client.DescribeInstances(request)
	if terr, ok := err.(*errors.TencentCloudSDKError); ok {
		code := terr.GetCode()
		if code == cvm.INVALIDZONE_MISMATCHREGION {
			return
		}
	}
	fmt.Printf("%s\n", resp.ToJsonString())
	t.Errorf(fmt.Sprintf("not expected error: %v", err))
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
		t.Errorf(fmt.Sprintf("fail to init client: %v", err))
	}

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to invoke api: %v", err))
	}
	if *response.Response.TotalCount != (int64)(len(response.Response.InstanceSet)) {
		t.Errorf("response item count inconsisitent!")
	}
}

// Payload limit
// GET: 32 KB
// POST + HmacSHA1: 1 MB
// POST + HmacSHA256: 1 MB
// POST + TC3-HMAC-SHA256: 10 MB
func TestPayloadLimit(t *testing.T) {
	const (
		KB = 1024
		MB = 1024 * KB
	)
	testCases := []struct {
		reqMethod  string
		signMethod string
		payload    int
		pass       bool
	}{
		{reqMethod: "GET", signMethod: "HmacSHA256", payload: 32 * KB, pass: true},
		{reqMethod: "GET", signMethod: "HmacSHA256", payload: 33 * KB, pass: false},
		{reqMethod: "POST", signMethod: "HmacSHA1", payload: 10 * MB, pass: true},
		{reqMethod: "POST", signMethod: "HmacSHA1", payload: 1*MB + 10*KB, pass: false},
		{reqMethod: "POST", signMethod: "HmacSHA256", payload: 1 * MB, pass: true},
		{reqMethod: "POST", signMethod: "HmacSHA256", payload: 1*MB + 10*KB, pass: false},
		{reqMethod: "POST", signMethod: "TC3-HMAC-SHA256", payload: 9 * MB, pass: true},
		{reqMethod: "POST", signMethod: "TC3-HMAC-SHA256", payload: 11 * MB, pass: false},
	}

	for _, tc := range testCases {
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.ReqMethod = tc.reqMethod
		cpf.SignMethod = tc.signMethod
		client, err := cvm.NewClient(getCredential(), regions.Guangzhou, cpf)
		if err != nil {
			t.Fatalf("fail to init client: %v", err)
		}

		req := cvm.NewDescribeInstancesRequest()
		req.InstanceIds = []*string{
			common.StringPtr(strings.Repeat("A", tc.payload)),
		}
		_, err = client.DescribeInstances(req)

		if err == nil {
			t.Fatalf("expect return error ")
		}

		tcErr, ok := err.(*errors.TencentCloudSDKError)
		if !ok {
			t.Fatalf("expect return TencentCloudSDKError, got %t", err)
		}

		isLimitError := tcErr.Code == "AuthFailure.SignatureFailure" ||
			tcErr.Code == "RequestSizeLimitExceeded" ||
			tcErr.Code == "ClientError.ParseJsonError"

		if !tc.pass && !isLimitError {
			t.Fatalf("test fail, reqMethod: %s, signMethod: %s, payload: %d, err: %s",
				tc.reqMethod, tc.signMethod, tc.payload, err)
		}
	}
}
