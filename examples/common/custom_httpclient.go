package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"

	// 如果需要自定义http客户端，可以在这里设置
	transport := http.DefaultTransport
	if _, ok := transport.(*http.Transport); ok {
		// go1.12版本之后支持 http.Transport.Clone
		// clonedTransport := transport.(*http.Transport).Clone()
		// clonedTransport.IdleConnTimeout = 30 * time.Second
		// transport = clonedTransport

		if cloneMethod, hasClone := reflect.TypeOf(transport).MethodByName("Clone"); hasClone {
			cloned := cloneMethod.Func.Call([]reflect.Value{reflect.ValueOf(transport)})[0].Interface().(http.RoundTripper)
			if clonedTransport, ok := cloned.(*http.Transport); ok {
				// 如果需要修改Transport的配置，可以在这里进行修改
				// 例如修改HTTP连接在空闲状态下的最大keep-alive时间
				clonedTransport.IdleConnTimeout = 30 * time.Second
				transport = clonedTransport
			}
		}
	}

	common.DefaultHttpClient = &http.Client{Transport: transport}
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf).WithLogger(log.Default())
	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")
	body := map[string]interface{}{}
	request.SetHeader(map[string]string{
		"X-TC-TraceId": "ffe0c072-8a5d-4e17-8887-a8a60252abca",
	})
	err := request.SetActionParameters(body)
	if err != nil {
		return
	}
	response := tchttp.NewCommonResponse()
	err = client.Send(request, response)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}
	fmt.Println(string(response.GetBody()))
}
