package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 10

	client, err := vpc.NewClient(credential, regions.Guangzhou, cpf)
	if err != nil {
		log.Fatalln(err)
	}

	id := common.StringPtr("rtb-your-routetable-id")
	request := vpc.NewCreateRoutesRequest()
	request.RouteTableId = id
	r1 := `{"DestinationCidrBlock":"10.244.0.0/24","GatewayType":"NORMAL_CVM","GatewayId":"172.16.0.17","RouteDescription":"test flanneld podcidr"}`
	r2 := `{"DestinationCidrBlock":"10.244.1.0/24","GatewayType":"NORMAL_CVM","GatewayId":"172.16.0.17","RouteDescription":"test flanneld podcidr"}`
	r3 := `{"DestinationCidrBlock":"10.244.22.0/24","GatewayType":"NORMAL_CVM","GatewayId":"172.16.0.17","RouteDescription":"test flanneld podcidr"}`

	err = request.FromJsonString(fmt.Sprintf(`{"Routes":[%s,%s,%s]}`, r1, r2, r3))
	if err != nil {
		log.Fatalln("json", err)
	}

	response, err := client.CreateRoutes(request)
	handleRespErr(err)

	fmt.Println("create route:", response.ToJsonString())

}

func handleRespErr(err error) {
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			log.Println("An API error has returned:")
		}
		log.Fatalln(err)
	}
}
