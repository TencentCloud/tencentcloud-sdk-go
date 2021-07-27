package integration

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"os"
	"testing"
)

func TestRoleArnProvider(t *testing.T) {
	sKey := os.Getenv("TENCENTCLOUD_SECRET_ID")
	sId := os.Getenv("TENCENTCLOUD_SECRET_KEY")
	roleArn := "qcs::cam::uin/123456:roleName/test"

	pc := common.DefaultRoleArnProvider(sKey, sId, roleArn)
	_, e := pc.GetCredential()
	if te, ok := e.(*tcerr.TencentCloudSDKError); ok {
		if te.GetCode() != "InternalError.GetRoleError" {
			t.Error(e)
		}
	}
}
