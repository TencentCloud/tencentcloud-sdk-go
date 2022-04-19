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

	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	faceid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
	scf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

func TestIntType(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := cvm.NewDescribeInstancesRequest()
	request.Limit = common.Int64Ptr(-1)

	_, err := client.DescribeInstances(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if strings.Index(err.Error(), "InvalidParameterValue") == -1 {
		t.Fatalf("The error is not expected: %s", err.Error())
	}
}

func TestUintAndStringAndNestType(t *testing.T) {
	// Set the interface parameters to uInt, String, multi-layer nested objects
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cbs.NewClient(credential, "ap-guangzhou", cpf)

	request := cbs.NewDescribeSnapshotsRequest()
	request.Filters = []*cbs.Filter{
		{
			Name:   common.StringPtr("snapshot-name"),
			Values: common.StringPtrs([]string{""}),
		},
		{
			Name:   common.StringPtr("snapshot-id"),
			Values: common.StringPtrs([]string{""}),
		},
	}
	request.Limit = common.Uint64Ptr(1)

	_, err := client.DescribeSnapshots(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if strings.Index(err.Error(), "InvalidParameter") == -1 {
		t.Fatalf("The error is not expected: %s", err.Error())
	}
}

func TestBoolAndDatetimeIsoType(t *testing.T) {
	// Set the interface parameters to Bool, Datetime_Iso type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cbs.NewClient(credential, "ap-guangzhou", cpf)

	request := cbs.NewModifySnapshotAttributeRequest()
	request.SnapshotId = common.StringPtr("")
	request.SnapshotName = common.StringPtr("")
	request.IsPermanent = common.BoolPtr(false)
	request.Deadline = common.StringPtr("2021-05-22T09:00:00Z")

	_, err := client.ModifySnapshotAttribute(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if strings.Index(err.Error(), "InvalidParameter") == -1 {
		t.Fatalf("The error is not expected: %s", err.Error())
	}
}

func TestDateType(t *testing.T) {
	// Set the interface parameter to Date type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cbs.NewClient(credential, "ap-guangzhou", cpf)

	request := cbs.NewDescribeSnapshotOperationLogsRequest()
	request.Filters = []*cbs.Filter{
		{
			Name:   common.StringPtr("snapshot-id"),
			Values: common.StringPtrs([]string{"invalid-snp-id"}),
		},
	}
	request.BeginTime = common.StringPtr("2021-05-21")
	request.EndTime = common.StringPtr("2021-05-21")

	// not a good api design, the api responses successully with wrong input
	_, err := client.DescribeSnapshotOperationLogs(request)
	if err == nil {
		t.Logf("unexpected success")
	}
}

func TestDatetimeType(t *testing.T) {
	// Set the interface parameter to Datetime type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := scf.NewClient(credential, "ap-guangzhou", cpf)

	request := scf.NewGetFunctionLogsRequest()
	request.StartTime = common.StringPtr("2021-05-21 11:00:00")
	request.EndTime = common.StringPtr("2021-05-21 11:00:00")

	_, err := client.GetFunctionLogs(request)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestFloatType(t *testing.T) {
	// Set the interface parameter to float type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := iai.NewClient(credential, "ap-guangzhou", cpf)

	request := iai.NewSearchFacesRequest()
	request.GroupIds = common.StringPtrs([]string{"test"})
	request.Url = common.StringPtr("https://www.test.com")
	request.FaceMatchThreshold = common.Float64Ptr(0.2)

	_, err := client.SearchFaces(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if strings.Index(err.Error(), "InvalidParameter") == -1 {
		t.Fatalf("The error is not expected: %s", err.Error())
	}
}

func TestComplexType(t *testing.T) {
	// Set interface parameters to complex types
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.Language = "en-US"
	client, _ := faceid.NewClient(credential, "ap-guangzhou", cpf)

	request := faceid.NewMobileStatusRequest()
	request.Mobile = common.StringPtr("null")
	request.Encryption = &faceid.Encryption{
		CiphertextBlob: common.StringPtr("null"),
		EncryptList:    common.StringPtrs([]string{"null", "null"}),
		Iv:             common.StringPtr("null"),
	}

	_, err := client.MobileStatus(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	// not a good test case, we should use cvm describe-instances
	if strings.Index(err.Error(), "UnauthorizedOperation.Nonactivated") == -1 {
		t.Fatalf("The error is not expected: %s", err.Error())
	}
}
