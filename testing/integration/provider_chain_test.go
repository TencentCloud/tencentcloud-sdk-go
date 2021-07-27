package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestProviderChain(t *testing.T) {
	rc := common.DefaultProviderChain()
	_, err := rc.GetCredential()
	if err != nil {
		t.Error(err)
	}
}
