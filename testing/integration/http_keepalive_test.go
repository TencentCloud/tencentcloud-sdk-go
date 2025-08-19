package integration

import (
	"context"
	"net/http/httptrace"
	"os"
	"testing"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func TestHttpClientIdleConnTimeout(t *testing.T) {
	var firstAddr string

	firstFlag := true

	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			realAddr := info.Conn.LocalAddr().String()
			if firstFlag {
				firstFlag = false
				firstAddr = realAddr
			} else {
				if firstAddr == realAddr {
					t.Fatalf("tcp connection should not be reused!")
				}
			}
		},
	}

	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	client, err := cvm.NewClient(
		credential,
		regions.Guangzhou,
		profile.NewClientProfile())
	if err != nil {
		t.Fatalf("fail to init client: %v", err)
	}

	const (
		serverKeepalive  = time.Second * 60
		clientKeepalive  = time.Second * 30
		timeToDisconnect = (serverKeepalive + clientKeepalive) / 2
	)

	ctx := httptrace.WithClientTrace(context.Background(), trace)

	for i := 0; i < 2; i++ {
		request := cvm.NewDescribeInstancesRequest()
		response, err := client.DescribeInstancesWithContext(ctx, request)
		if err != nil {
			t.Fatalf("fail to invoke api: %v", err)
		}
		if *response.Response.TotalCount != (int64)(len(response.Response.InstanceSet)) {
			t.Fatalf("response item count inconsisitent!")
		}
		if i == 0 {
			time.Sleep(timeToDisconnect)
		}
	}
}
