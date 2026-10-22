package common

import (
	"io/ioutil"
	"os"
	"path/filepath"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

const (
	EnvTCCLIProfile = "TCCLI_PROFILE"
)

type TccliProfileProvider struct {
	profile string
}

// DefaultTccliProfileProvider returns a default TCCLI profile provider.
// profile path :
//
//	Linux: ~/.tccli/{profile}.credential
//	Windows: C:\Users\{UserName}\.tccli\{profile}.credential
//
//	{profile} is determined by environment variable TCCLI_PROFILE, falling back to "default".
func DefaultTccliProfileProvider() *TccliProfileProvider {
	profile, ok := os.LookupEnv(EnvTCCLIProfile)
	if !ok {
		profile = "default"
	}

	return &TccliProfileProvider{profile: profile}
}

// NewTccliProfileProvider returns a TCCLI profile provider with specified profile.
func NewTccliProfileProvider(profile string) *TccliProfileProvider {
	return &TccliProfileProvider{profile: profile}
}

func (p *TccliProfileProvider) getCredentialFilePath() string {
	homePath := getHomePath()
	if homePath == "" {
		return ""
	}
	return filepath.Join(homePath, ".tccli", p.profile+".credential")
}

func (p *TccliProfileProvider) checkCredentialFile() (string, error) {
	path := p.getCredentialFilePath()
	if path == "" {
		return "", nil
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return path, nil
}

func (p *TccliProfileProvider) GetCredential() (CredentialIface, error) {
	path, err := p.checkCredentialFile()
	if err != nil {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "Failed to find TCCLI profile file, "+err.Error(), "")
	}
	if path == "" {
		return nil, fileDoseNotExist
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tccliCred := struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
		Token     string `json:"token"`
	}{}
	if err = json.Unmarshal(data, &tccliCred); err != nil {
		return nil, err
	}

	if tccliCred.SecretId == "" || tccliCred.SecretKey == "" {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "Failed to parse TCCLI profile file, please confirm whether it contains \"secretId\" and \"secretKey\"", "")
	}
	return &Credential{
		SecretId:  tccliCred.SecretId,
		SecretKey: tccliCred.SecretKey,
		Token:     tccliCred.Token,
	}, nil
}
