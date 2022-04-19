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

func TestRequestPublicParams(t *testing.T) {
	region := "ap-guangzhou"
	language := "en-US"

	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.Language = language
	client, _ := cvm.NewClient(credential, region, cpf)

	request := cvm.NewDescribeZonesRequest()

	response, err := client.DescribeZones(request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *response.Response.TotalCount > 0 && strings.Index(response.ToJsonString(), "guangzhou") == -1 {
		t.Fatalf("The en-US setting does not take effect!")
	}
}
