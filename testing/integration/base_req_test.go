package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
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
	request.BaseRequest.Init().WithApiInfo("cvm", "2017-03-12", "DescribeInstances")

	_, err := client.DescribeInstances(request)
	if err != nil {
		t.Fatal(err)
	}
}
