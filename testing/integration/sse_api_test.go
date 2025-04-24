// Copyright 2018-2025 Tencent Ltd.
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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
	"os"
	"testing"
)

func TestChatCompletionsStream(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	request := hunyuan.NewChatCompletionsRequest()

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request.Messages = append(request.Messages, msg)

	request.Stream = common.BoolPtr(true)
	request.Model = common.StringPtr("hunyuan-standard")

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	for event := range response.Events {
		t.Log(string(event.Data))
	}
}

func TestChatCompletionsNonStream(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	request := hunyuan.NewChatCompletionsRequest()

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request.Messages = append(request.Messages, msg)

	request.Stream = common.BoolPtr(false)
	request.Model = common.StringPtr("hunyuan-standard")

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(response.ToJsonString())
}
