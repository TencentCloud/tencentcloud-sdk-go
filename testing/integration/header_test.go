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
	"context"
	"net/http/httptrace"
	"os"
	"strings"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestSetHeader(t *testing.T) {
	signMethods := []string{"TC3-HMAC-SHA256", "HmacSHA256", "HmacSHA1"}
	reqMethods := []string{"GET", "POST"}

	endpoint := "vpc.tencentcloudapi.com"
	signHost := "cvm.tencentcloudapi.com"

	for _, signMethod := range signMethods {
		for _, reqMethod := range reqMethods {
			credential := common.NewCredential(
				os.Getenv("TENCENTCLOUD_SECRET_ID"),
				os.Getenv("TENCENTCLOUD_SECRET_KEY"),
			)
			cpf := profile.NewClientProfile()
			cpf.HttpProfile.Endpoint = endpoint
			cpf.HttpProfile.ReqMethod = reqMethod
			cpf.SignMethod = signMethod
			client, err := cvm.NewClient(credential, "ap-guangzhou", cpf)
			if err != nil {
				t.Fatal(err)
			}

			request := cvm.NewDescribeZonesRequest()
			request.SetHeader(map[string]string{
				"TestKey": "TestValue",
				"Host":    signHost,
			})

			var hostChecked, testKeyChecked bool
			trace := &httptrace.ClientTrace{
				DNSStart: func(info httptrace.DNSStartInfo) {
					if info.Host != endpoint {
						t.Fatalf("invalid req endpoint: signMethod=%s reqMethod=%s host=%s", signMethod, reqMethod, info.Host)
					}
				},
				WroteHeaderField: func(key string, value []string) {
					if strings.EqualFold(key, "Host") {
						hostChecked = true
						if value[0] != signHost {
							t.Fatalf("invalid Host: %s", value[0])
						}
					}

					if strings.EqualFold(key, "TestKey") {
						testKeyChecked = true
						if value[0] != "TestValue" {
							t.Fatalf("invalid Header: %s", value[0])
						}
					}
				},
			}
			ctx := httptrace.WithClientTrace(context.Background(), trace)
			_, err = client.DescribeZonesWithContext(ctx, request)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if !hostChecked || !testKeyChecked {
				t.Fatalf("header not sent: %v %v", hostChecked, testKeyChecked)
			}
		}
	}
}
