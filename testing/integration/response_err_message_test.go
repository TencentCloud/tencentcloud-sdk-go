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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestResponseMessage(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := cvm.NewDescribeInstancesOperationLimitRequest()
	request.InstanceIds = common.StringPtrs([]string{""})
	request.Operation = common.StringPtr("INSTANCE_DEGRADE")

	_, err := client.DescribeInstancesOperationLimit(request)
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if strings.Index(err.Error(), "InvalidParameterValue.InstanceIdMalformed") == -1 {
		t.Fatalf("The error code %s is not expected value", err.Error())
	}
}
