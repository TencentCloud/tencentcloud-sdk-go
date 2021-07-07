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
	roleName string
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
	return &CvmRoleProvider{roleName: roleName}
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
	if r.roleName != "" {
		return r.roleName, nil
	}
	rn, err := get(RoleUrl)
	if err == nil {
		r.roleName = string(rn)
		return r.roleName, nil
	}
	return string(rn), err
}

func (r *CvmRoleProvider) GetCredential() (CredentialIface, error) {
	roleName, err := r.getRoleName()
	if err != nil {
		if errors.Is(err, RoleNotBound) {
			return nil, CVMNOROLE
		}
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", err.Error(), "")
	}
	//通过元数据接口获取临时凭证 https://cloud.tencent.com/document/product/213/4934

	body, err := get(RoleUrl + roleName)

	if err != nil {
		return nil, err
	}
	rspSt := new(roleRsp)
	if err = json.Unmarshal(body, rspSt); err != nil {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", err.Error(), "")
	}
	if rspSt.Code != "Success" {
		return nil, tcerr.NewTencentCloudSDKError("ClientError.CredentialError", "Get credential from metadata server by role name "+roleName+" failed, code="+rspSt.Code, "")
	}
	cre := &CvmRoleCredential{
		tmpSecretId:  rspSt.TmpSecretId,
		tmpSecretKey: rspSt.TmpSecretKey,
		token:        rspSt.Token,
		roleName:     roleName,
		expiredTime:  rspSt.ExpiredTime,
		source:       r,
	}
	return cre, nil
}
