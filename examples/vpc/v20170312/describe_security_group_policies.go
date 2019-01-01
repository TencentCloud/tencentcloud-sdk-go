package main

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	client, _ := vpc.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	sgreq := vpc.NewDescribeSecurityGroupsRequest()
	sgresp, err := client.DescribeSecurityGroups(sgreq)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if *sgresp.Response.TotalCount < 1 {
		fmt.Printf("No security group found.")
		return
	}
	request := vpc.NewDescribeSecurityGroupPoliciesRequest()
	// we don't care which one, hence with no filter and just use the first one
	request.SecurityGroupId = sgresp.Response.SecurityGroupSet[0].SecurityGroupId
	response, err := client.DescribeSecurityGroupPolicies(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
