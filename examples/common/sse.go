package main

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"log"
	"os"
	"time"
)

type MyStreamApiResponse struct {
	tchttp.BaseSSEResponse
}

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	// 流式请求要设置 ReqTimeout=0, 表示没有超时, 否则在超时时间到达后会中断http流
	cpf.HttpProfile.ReqTimeout = 0

	client := common.NewCommonClient(credential, regions.Guangzhou, cpf).WithLogger(log.Default())

	request := tchttp.NewCommonRequest("apitest", "2017-03-12", "StreamCall")
	response := &MyStreamApiResponse{}

	request.SetHeader(map[string]string{
		"Content-Type": "application/octet-stream",
	})

	// 如果客户端想要主动关闭数据流, 使用 Context
	clientClose := false
	if clientClose {
		ctx, cancel := context.WithCancel(context.Background())
		request.SetContext(ctx)
		go func() {
			time.Sleep(time.Second * 3)
			println("cancel early")
			cancel()
		}()
	}

	err := client.Send(request, response)
	if err != nil {
		log.Printf("fail to invoke api: %v \n", err)
		return
	}

	// 在 SSE流中断或者Context到期 后channel会被关闭
	for e := range response.Events {
		if e.Err != nil {
			log.Println(e.Err.Error())
			break
		}

		log.Printf("id=%s, event=%s, retry=%d, data=%s", e.Id, e.Event, e.Retry, e.Data)
	}
}
