package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	scf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

func main() {
	// 从环境变量获取SID和SKEY实例化认证对象
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	// 实例化客户端配置对象
	cpf := profile.NewClientProfile()
	// 实例化SCF的client对象
	client, _ := scf.NewClient(credential, regions.Guangzhou, cpf)

	// API -- ListFunctions:获取函数列表
	listFunctionsRequest := scf.NewListFunctionsRequest()
	listFunctionsResponse, err := client.ListFunctions(listFunctionsRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listFunctionsResponse.ToJsonString())
}
