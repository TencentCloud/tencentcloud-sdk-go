package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"os"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.Debug = true
	// 如果你希望在请求中携带空数组, 需要开启 OmitNil 功能
	cpf.OmitNil = true

	client, err := cvm.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		panic(err)
	}

	req := cvm.NewDeleteDisasterRecoverGroupsRequest()
	req.DisasterRecoverGroupIds = []*string{}
	response, err := client.DeleteDisasterRecoverGroups(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.ToJsonString())
}
