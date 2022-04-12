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

	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := cvm.NewDescribeInstancesRequest()

	request.Limit = common.Int64Ptr(-1)

	response, err := client.DescribeInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "length less than 1 or greater than 100") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}
}

func TestUintAndStringAndNestType(t *testing.T) {
	//Set the interface parameters to uInt, String, multi-layer nested complex objects
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cbs.tencentcloudapi.com"
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

	response, err := client.DescribeSnapshots(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "SnapshotId  is invalid") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}

}

func TestBoolAndDatetimeIsoType(t *testing.T) {
	//Set the interface parameters to Bool, Datetime_Iso type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cbs.tencentcloudapi.com"
	client, _ := cbs.NewClient(credential, "ap-guangzhou", cpf)

	request := cbs.NewModifySnapshotAttributeRequest()

	request.SnapshotId = common.StringPtr("")
	request.SnapshotName = common.StringPtr("")
	request.IsPermanent = common.BoolPtr(false)
	request.Deadline = common.StringPtr("2021-05-22T09:00:00Z")

	response, err := client.ModifySnapshotAttribute(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "SnapshotId:  is invalid") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}

}

func TestDateType(t *testing.T) {
	//Set the interface parameter to Date type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cbs.tencentcloudapi.com"
	client, _ := cbs.NewClient(credential, "ap-guangzhou", cpf)

	request := cbs.NewDescribeSnapshotOperationLogsRequest()

	request.Filters = []*cbs.Filter{
		{
			Name:   common.StringPtr("snapshot-id"),
			Values: common.StringPtrs([]string{""}),
		},
	}
	request.BeginTime = common.StringPtr("2021-05-21")
	request.EndTime = common.StringPtr("2021-05-21")

	response, err := client.DescribeSnapshotOperationLogs(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "SnapshotId  is invalid") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}
}

func TestDatetimeType(t *testing.T) {
	//Set the interface parameter to Datetime type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "scf.tencentcloudapi.com"
	client, _ := scf.NewClient(credential, "ap-guangzhou", cpf)

	request := scf.NewGetFunctionLogsRequest()

	request.StartTime = common.StringPtr("2021-05-21 11:00:00")
	request.EndTime = common.StringPtr("2021-05-21 11:00:00")

	response, err := client.GetFunctionLogs(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}
	if err != nil {
		t.Errorf(fmt.Sprintf("fail to init client: %v", err))
	}
	fmt.Printf("%s\n", response.ToJsonString())
}

func TestFloatType(t *testing.T) {
	//Set the interface parameter to float type
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
	client, _ := iai.NewClient(credential, "ap-guangzhou", cpf)

	request := iai.NewSearchFacesRequest()

	request.GroupIds = common.StringPtrs([]string{"test"})
	request.Url = common.StringPtr("https://www.test.com")
	request.FaceMatchThreshold = common.Float64Ptr(0.2)

	response, err := client.SearchFaces(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "人员库ID不存在") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}
}

func TestComplexType(t *testing.T) {
	//Set interface parameters to complex types
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	client, _ := faceid.NewClient(credential, "ap-guangzhou", cpf)

	request := faceid.NewMobileStatusRequest()

	request.Mobile = common.StringPtr("null")
	request.Encryption = &faceid.Encryption{
		CiphertextBlob: common.StringPtr("null"),
		EncryptList:    common.StringPtrs([]string{"null", "null"}),
		Iv:             common.StringPtr("null"),
	}

	response, err := client.MobileStatus(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		if strings.Index(err.Error(), "Service not activated") == -1 {
			t.Errorf(fmt.Sprintf("The error is not as expected"))
		}
	} else {
		if err != nil {
			t.Errorf(fmt.Sprintf("fail to init client: %v", err))
		}
		fmt.Printf("%s\n", response.ToJsonString())
		t.Errorf(fmt.Sprintf("The error is not as expected"))
	}
}
