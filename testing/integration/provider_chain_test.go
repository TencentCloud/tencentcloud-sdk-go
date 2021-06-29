package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestProviderChain(t *testing.T) {
	rc := common.DefaultProviderChain()
	cre, err := rc.GetCredential()
	if err != nil {
		t.Error(err)
	}
	t.Log("id:" + cre.GetSecretId() + "\n" + "key:" + cre.GetSecretKey() + "\n" + "token:" + cre.GetToken())
}
