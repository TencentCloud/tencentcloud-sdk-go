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
	"testing"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
)

func getBigString(l int) string {
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, 'a')
	}
	return string(result)
}

func TestUnsignedPayload(t *testing.T) {
	status := [2]bool{true, false}
	var useTime [2]float64

	for i := 0; i < 2; i++ {
		credential := common.NewCredential(
			os.Getenv("TENCENTCLOUD_SECRET_ID"),
			os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
		cpf.UnsignedPayload = status[i]
		client, _ := iai.NewClient(credential, "ap-guangzhou", cpf)

		request := iai.NewCompareFaceRequest()

		request.ImageA = common.StringPtr(getBigString(3145))
		request.ImageB = common.StringPtr(getBigString(2097))
		// request.UrlA = common.StringPtr("https://cloudapi-test-1254240205.cos.ap-nanjing.myqcloud.com/9d1f1393e5fc5ca13032e40e3bf9e882.jpeg")
		// request.UrlB = common.StringPtr("https://cloudapi-test-1254240205.cos.ap-nanjing.myqcloud.com/9d1f1393e5fc5ca13032e40e3bf9e882.jpeg")

		start := time.Now()
		_, err := client.CompareFace(request)
		if _, ok := err.(*errors.TencentCloudSDKError); !ok {
			t.Errorf("request failed with unexpected error.")
		}
		elapsed := time.Since(start)
		useTime[i] = float64(elapsed)
	}

	if useTime[0] > useTime[1] {
		t.Logf("The request time of setting unsigned-payload %f > ordinary %f", useTime[0], useTime[1])
	}

}
