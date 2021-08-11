package main

import (
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// 此函数应该会在前5次失败，在后5次成功
func main() {
	pr := common.DefaultProfileProvider()
	credential, _ := pr.GetCredential()
	log.Println(credential.GetSecretId(), credential.GetSecretKey())
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 2
	// 错误的域名，请求一定失败，但是会在第5次之后成功
	cpf.HttpProfile.Endpoint = "cvm.ap-shanghai1.tencentcloudapi.com"
	cpf.BackupEndPoint = "ap-beijing.tencentcloudapi.com"

	// 开启熔断器
	cpf.DisableRegionBreaker = false
	client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
	for i := 0; i < 10; i++ {
		request := cvm.NewDescribeZonesRequest()
		response, err := client.DescribeZones(request)
		log.Println(i+1, "/", 10)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			log.Printf("An API error has returned: %s", err)
			continue
		}
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("%s\n", response.ToJsonString())
	}
}
