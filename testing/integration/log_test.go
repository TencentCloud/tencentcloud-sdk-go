package integration

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
	"testing"
)

type testLogger struct {
	logged bool
}

func (l *testLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	l.logged = true
}

func TestLogger(t *testing.T) {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	logger := &testLogger{}
	cpf := profile.NewClientProfile()
	cpf.Debug = true // debug will cause logging
	client := common.NewCommonClient(credential, regions.Guangzhou, cpf).WithLogger(logger)

	request := tchttp.NewCommonRequest("cvm", "2017-03-12", "DescribeZones")
	response := tchttp.NewCommonResponse()
	err := client.Send(request, response)
	if err != nil {
		t.Fatal(err)
	}

	if !logger.logged {
		t.Fatal("logger failed")
	}
}
