package common

import tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"

var (
	ENVNOTSET        = tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "could not find environmental variable", "")
	FILEDOSENOTEXIST = tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "could not find config file", "")
	CVMNOROLE        = tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "get cvm role name failed, Please confirm whether the role is bound", "")
)

// Provider interface
type Provider interface {
	GetCredential() (CredentialIface, error)
}
