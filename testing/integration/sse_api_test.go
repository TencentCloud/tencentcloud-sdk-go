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
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	es "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20250101"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
	"net"
	"net/http/httptrace"
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

func TestUseAIDomainByDefault(t *testing.T) {
	const expectHost = "es.ai.tencentcloudapi.com"
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	client, _ := es.NewClient(credential, regions.Guangzhou, cpf)

	request := es.NewParseDocumentRequest()

	var requestedHost string
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			var err error
			requestedHost, _, err = net.SplitHostPort(hostPort)
			if err != nil {
				panic(err)
			}
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	client.ParseDocumentWithContext(ctx, request)

	if requestedHost != expectHost {
		t.Fatalf("expected host: %s, got: %s", expectHost, requestedHost)
	}
}

func TestAIDomainOverrideByDomain(t *testing.T) {
	const expectHost = "custom-domain.com"
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	client, _ := es.NewClient(credential, regions.Guangzhou, cpf)

	request := es.NewParseDocumentRequest()
	request.SetDomain(expectHost)

	var requestedHost string
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			var err error
			requestedHost, _, err = net.SplitHostPort(hostPort)
			if err != nil {
				panic(err)
			}
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	client.ParseDocumentWithContext(ctx, request)

	if requestedHost != expectHost {
		t.Fatalf("expected host: %s, got: %s", expectHost, requestedHost)
	}
}

func TestAIDomainOverrideByProfile(t *testing.T) {
	const expectHost = "custom-domain.com"
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = expectHost
	client, _ := es.NewClient(credential, regions.Guangzhou, cpf)

	request := es.NewParseDocumentRequest()

	var requestedHost string
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			var err error
			requestedHost, _, err = net.SplitHostPort(hostPort)
			if err != nil {
				panic(err)
			}
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	client.ParseDocumentWithContext(ctx, request)

	if requestedHost != expectHost {
		t.Fatalf("expected host: %s, got: %s", expectHost, requestedHost)
	}
}

func TestAIDomainOverrideByRootDomain(t *testing.T) {
	const (
		rootDomain = "custom-domain.com"
		expectHost = "es." + rootDomain
	)
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	client, _ := es.NewClient(credential, regions.Guangzhou, cpf)

	request := es.NewParseDocumentRequest()
	request.SetRootDomain(rootDomain)

	var requestedHost string
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			var err error
			requestedHost, _, err = net.SplitHostPort(hostPort)
			if err != nil {
				panic(err)
			}
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	client.ParseDocumentWithContext(ctx, request)

	if requestedHost != expectHost {
		t.Fatalf("expected host: %s, got: %s", expectHost, requestedHost)
	}
}
