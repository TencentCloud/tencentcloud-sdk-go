package main

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
	/*
		本示例演示如何通过 **nginx 反向代理** 请求腾讯云 API。
		场景：某些网络环境不能直接访问腾讯云域名(*.tencentcloudapi.com)，只能通过代理服务器。

		========================
		1. nginx 配置示例：
		========================
		server {
		    listen 80;
		    # 指定 DNS 服务器（可以根据自己网络环境进行替换）
		    resolver 114.114.114.114;

		    # 可以自定义请求路径
		    location /tc_api {
		        # http_host 后必须以 / 结尾
		        proxy_pass https://$http_host/;
		    }
		}
	*/
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	// 2. 将 Scheme 设置为 http
	cpf.HttpProfile.Scheme = "http"

	// 3. 替换 1.2.3.4 为真实的 nginx 地址, /tc_api 可以自定义
	nginx := "9.134.89.153/tc_api"
	cpf.HttpProfile.Endpoint = nginx

	client, err := cvm.NewClient(credential, regions.Guangzhou, cpf)
	if err != nil {
		panic(err)
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Limit = common.Int64Ptr(10)
	// 4. 设置 header 为 {服务名}.tencentcloudapi.com
	request.SetHeader(map[string]string{
		"Host": "cvm.tencentcloudapi.com",
	})

	response, err := client.DescribeInstances(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.ToJsonString())
}
