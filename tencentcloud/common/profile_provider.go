package common

import (
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"runtime"
)

const (
	EnvCredentialFile = "TENCENTCLOUD_CREDENTIALS_FILE"
)

type ProfileProvider struct {
	//path []string

}

//new?

// DefaultProfileProvider return a default Profile  provider
// profile path :
// linux: ~/.tencentcloud/credentials
// windows: \c:\User\NAME\.tencentcloud\credentials
func DefaultProfileProvider() *ProfileProvider {
	return &ProfileProvider{}
}

// GetHomePath return home directory according to the system.
// if the environmental variables does not exist, it will return empty string
func GetHomePath() string {
	// Windows
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	// *nix
	return os.Getenv("HOME")
}

func GetCredentialsFilePath() string {
	return filepath.Join(GetHomePath(), ".tencentcloud", "credentials")
}

func checkDefaultFile() (path string, err error) {
	path = GetCredentialsFilePath()
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return path, nil
}

func (p *ProfileProvider) GetCredential() (CredentialIface, error) {
	path, ok := os.LookupEnv(EnvCredentialFile)
	if !ok {
		var err error
		path, err = checkDefaultFile()
		if err != nil {
			return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialNotFound", "Failed to find profile file,"+err.Error(), "")
		}
		if path == "" {
			return nil, FILEDOSENOTEXIST
		}
	} else if path == "" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Environment variable '"+EnvCredentialFile+"' cannot be empty", "")
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Failed to parse profile file,"+err.Error(), "")
	}

	sId := cfg.Section("default").Key("secret_id").String()
	sKey := cfg.Section("default").Key("secret_key").String()

	if sId == "" || sKey == "" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Failed to parse profile file,please confirm whether it contains secret_id and secret_key in section: default ", "")
	}
	return &Credential{
		SecretId:  sId,
		SecretKey: sKey,
		Token:     "",
	}, nil
}
