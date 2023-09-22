package main

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"log"
	"os"
)

type MyStreamApiResponse struct {
	tchttp.BaseSSEResponse
}

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))

	cpf := profile.NewClientProfile()
	// 流式请求不要设置 ReqTimeout, 否则在超时时间到达后会中断http流
	//cpf.HttpProfile.ReqTimeout = 1
	cpf.Debug = true

	client := common.NewCommonClient(credential, regions.Guangzhou, cpf).WithLogger(log.Default())

	request := tchttp.NewCommonRequest("apitest", "2017-03-12", "StreamCall")
	response := &MyStreamApiResponse{}

	err := client.Send(request, response)
	if err != nil {
		log.Printf("fail to invoke api: %v \n", err)
		return
	}

	// 在 SSE 流中断后 channel 会自动关闭
	for e := range response.Events {
		log.Printf("id=%s, event=%s, retry=%d, data=%s", e.Id, e.Event, e.Retry, e.Data)
	}
}
