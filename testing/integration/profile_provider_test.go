package integration

import (
	common "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestProfileProvider(t *testing.T) {
	p:=common.DefaultProfileProvider()
	c,err:=p.GetCredential()
	if err!=nil{
		t.Error(err)
	}
	t.Log(c)
}