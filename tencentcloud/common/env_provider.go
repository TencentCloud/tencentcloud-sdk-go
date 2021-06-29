package common

import (
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"os"
)

type EnvProvider struct {
	secretIdENV  string
	secretKeyENV string
}

// DefaultEnvProvider return a default provider
func DefaultEnvProvider() *EnvProvider {
	return &EnvProvider{
		secretIdENV:  "TENCENTCLOUD_SECRET_ID",
		secretKeyENV: "TENCENTCLOUD_SECRET_KEY",
	}
}

func NewEnvProvider(idEnv, keyEnv string) *EnvProvider {
	return &EnvProvider{
		secretIdENV:  idEnv,
		secretKeyENV: keyEnv,
	}
}

func (p *EnvProvider) GetCredential() (CredentialIface, error) {
	secretId, ok1 := os.LookupEnv(p.secretIdENV)
	secretKey, ok2 := os.LookupEnv(p.secretKeyENV)
	if !ok1 || !ok2 {
		return nil, ENVNOTSET
	}
	if secretId == "" || secretKey == "" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialNotFound", "Environmental variable ("+p.secretIdENV+" or "+p.secretKeyENV+") is empty", "")
	}
	return NewCredential(secretId, secretKey), nil
}
