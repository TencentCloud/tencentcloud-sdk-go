package common

import (
	"os"
	"testing"
)

func TestUserRoleTokenCredential(t *testing.T) {
	credential := NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))

	//options := []UserRoleTokenOptions {
	//	WithExtra(""),
	//	WithTargetAction([]string{}),
	//}
	provider := NewUserRoleTokenProvider(credential, 100010801389, 100010801389)
	userTokenCredential := &UserRoleTokenCredential{
		source: provider,
	}
	tmpSecretId := userTokenCredential.GetSecretId()
	tmpSecretKey := userTokenCredential.GetSecretKey()
	token := userTokenCredential.GetToken()
	t.Logf("tmpSecretId: %s", tmpSecretId)
	t.Logf("tmpSecretKey: %s", tmpSecretKey)
	t.Logf("token: %s", token)

}
