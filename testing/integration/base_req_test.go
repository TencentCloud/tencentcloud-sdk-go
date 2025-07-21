package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
	"os"
	"testing"
)

func TestInitBaseRequest(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := &cvm.DescribeInstancesRequest{
		Limit: common.Int64Ptr(1),
	}

	_, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitBaseRequest2(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := &cvm.DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Limit:       common.Int64Ptr(1),
	}

	_, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitBaseRequest3(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := &cvm.DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Limit:       common.Int64Ptr(1),
	}
	request.BaseRequest.Init()

	_, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitBaseRequest4(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := &cvm.DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Limit:       common.Int64Ptr(1),
	}
	request.BaseRequest.Init().WithApiInfo("cvm", cvm.APIVersion, "DescribeInstances")

	_, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitBaseRequestSSE(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request := &hunyuan.ChatCompletionsRequest{
		Messages: []*hunyuan.Message{msg},
		Stream:   common.BoolPtr(true),
		Model:    common.StringPtr("hunyuan-standard"),
	}

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	for range response.Events {
	}
}

func TestInitBaseRequestSSE2(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request := &hunyuan.ChatCompletionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Messages:    []*hunyuan.Message{msg},
		Stream:      common.BoolPtr(true),
		Model:       common.StringPtr("hunyuan-standard"),
	}

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	for range response.Events {
	}
}

func TestInitBaseRequestSSE3(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request := &hunyuan.ChatCompletionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Messages:    []*hunyuan.Message{msg},
		Stream:      common.BoolPtr(true),
		Model:       common.StringPtr("hunyuan-standard"),
	}
	request.BaseRequest.Init()

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	for range response.Events {
	}
}

func TestInitBaseRequestSSE4(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 400

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("hi, tell me a joke"),
	}
	request := &hunyuan.ChatCompletionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
		Messages:    []*hunyuan.Message{msg},
		Stream:      common.BoolPtr(true),
		Model:       common.StringPtr("hunyuan-standard"),
	}
	request.Init().WithApiInfo("hunyuan", hunyuan.APIVersion, "ChatCompletions")

	response, err := client.ChatCompletions(request)
	if err != nil {
		t.Fatal(err)
	}

	for range response.Events {
	}
}
