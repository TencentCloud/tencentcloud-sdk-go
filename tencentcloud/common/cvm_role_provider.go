package common

import (
	"encoding/json"
	"errors"
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	MetaUrl = "http://metadata.tencentyun.com/latest/meta-data/"
	RoleUrl = MetaUrl + "cam/security-credentials/"
)

var RoleNotBound = errors.New("get cvm role name failed, Please confirm whether the role is bound")

type CvmRoleProvider struct {
	RoleName string
}

type roleRsp struct {
	TmpSecretId  string    `json:"TmpSecretId"`
	TmpSecretKey string    `json:"TmpSecretKey"`
	ExpiredTime  int64     `json:"ExpiredTime"`
	Expiration   time.Time `json:"Expiration"`
	Token        string    `json:"Token"`
	Code         string    `json:"Code"`
}

func NewCvmRoleProvider(roleName string) *CvmRoleProvider {
	return &CvmRoleProvider{RoleName: roleName}
}

func DefaultCvmRoleProvider() *CvmRoleProvider {
	return NewCvmRoleProvider("")
}

func get(url string) ([]byte, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode == http.StatusNotFound {
		return nil, RoleNotBound
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *CvmRoleProvider) getRoleName() (string, error) {
	if r.RoleName != "" {
		return r.RoleName, nil
	}
	rn, err := get(RoleUrl)
	if err == nil {
		r.RoleName = string(rn)
		return r.RoleName, nil
	}
	return string(rn), err
}

func (r *CvmRoleProvider) GetCredential() (CredentialIface, error) {
	roleName, err := r.getRoleName()
	if err != nil {
		//如果没有就返回特点的err,否则会使 DefaultProviderChain 终止
		if errors.Is(err, RoleNotBound) {
			return nil, CVMNOROLE
		}
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", err.Error(), "")
	}

	body, err := get(RoleUrl + roleName)

	if err != nil {
		return nil, err
	}
	rspSt := new(roleRsp)
	if err = json.Unmarshal(body, rspSt); err != nil {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", err.Error(), "")
	}
	if rspSt.Code != "Success" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "get credential from cvm role name failed,code="+rspSt.Code, "")
	}
	cre := &CvmRoleCredential{
		tmpSecretId:  rspSt.TmpSecretId,
		tmpSecretKey: rspSt.TmpSecretKey,
		token:        rspSt.Token,
		RoleName:     roleName,
		ExpiredTime:  rspSt.ExpiredTime,
		source:       r,
	}
	return cre, nil
}
