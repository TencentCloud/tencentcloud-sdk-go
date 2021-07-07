package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestEnvProvider(t *testing.T) {
	p := common.DefaultEnvProvider()
	_, err := p.GetCredential()
	if err != nil {
		t.Error(err)
	}
}
