package common

import (
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	ini "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/ini"
	"os"
	"path/filepath"
	"runtime"
)

const (
	EnvCredentialFile = "TENCENTCLOUD_CREDENTIALS_FILE"
)

type ProfileProvider struct {
}

// DefaultProfileProvider return a default Profile  provider
// profile path :
// The value of the environment variable TENCENTCLOUD_CREDENTIALS_FILE
// linux: ~/.tencentcloud/credentials
// windows: \c:\Users\NAME\.tencentcloud\credentials
func DefaultProfileProvider() *ProfileProvider {
	return &ProfileProvider{}
}

// getHomePath return home directory according to the system.
// if the environmental variables does not exist, it will return empty string
func getHomePath() string {
	// Windows
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	// *nix
	return os.Getenv("HOME")
}

func getCredentialsFilePath() string {
	return filepath.Join(getHomePath(), ".tencentcloud", "credentials")
}

func checkDefaultFile() (path string, err error) {
	path = getCredentialsFilePath()
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
			return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Failed to find profile file,"+err.Error(), "")
		}
		if path == "" {
			return nil, fileDoseNotExist
		}
	} else if path == "" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Environment variable '"+EnvCredentialFile+"' cannot be empty", "")
	}

	cfg, err := ini.Parse(path)
	if err != nil {
		return nil, err
	}

	sId := cfg.Section("default").Key("secret_id").String()
	sKey := cfg.Section("default").Key("secret_key").String()

	if sId == "" || sKey == "" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Failed to parse profile file,please confirm whether it contains \"secret_id\" and \"secret_key\" in section: \"default\" ", "")
	}
	return &Credential{
		SecretId:  sId,
		SecretKey: sKey,
		Token:     "",
	}, nil
}
