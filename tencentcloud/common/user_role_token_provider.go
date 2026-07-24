package common

import (
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"log"
)

const (
	accessTokenService  = "sts"
	accessTokenVersion  = "2018-08-13"
	accessTokenAction   = "GetAccessToken"
	accessTokenEndpoint = "sts.internal.tencentcloudapi.com"
	accessTokenRegion   = regions.Guangzhou
)

type UserRoleTokenProvider struct {
	Credential *Credential

	TargetOwnerUin  uint64
	TargetUin       uint64
	DurationSeconds uint64
	Extra           string
	TargetAction    []string
}

type tmpCredential struct {
	tmpSecretId  string
	tmpSecretKey string
	token        string
}

type userRoleTokenRsp struct {
	credentials tmpCredential
	expiredTime uint64
	expiration  string
}

type UserRoleTokenOptions func(*UserRoleTokenProvider)

func WithDurationSeconds(durationSeconds uint64) UserRoleTokenOptions {
	return func(p *UserRoleTokenProvider) {
		p.DurationSeconds = durationSeconds
	}
}

func WithExtra(extra string) UserRoleTokenOptions {
	return func(p *UserRoleTokenProvider) {
		p.Extra = extra
	}
}

func WithTargetAction(targetAction []string) UserRoleTokenOptions {
	return func(p *UserRoleTokenProvider) {
		p.TargetAction = targetAction
	}
}

func NewUserRoleTokenProvider(credential *Credential, OwnerUin, Uin uint64, options ...UserRoleTokenOptions) *UserRoleTokenProvider {
	p := &UserRoleTokenProvider{
		Credential: credential,

		TargetOwnerUin: OwnerUin,
		TargetUin:      Uin,
	}
	for _, op := range options {
		op(p)
	}
	log.Println(p)
	return p
}

// GetCredential 服务账号换取用户身份临时密钥
func (p *UserRoleTokenProvider) GetCredential() (CredentialIface, error) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = accessTokenEndpoint
	client := NewCommonClient(p.Credential, accessTokenRegion, cpf)
	request := tchttp.NewCommonRequest(accessTokenService, accessTokenVersion, accessTokenAction)
	body := map[string]interface{}{
		"TargetOwnerUin": p.TargetOwnerUin,
		"TargetUin":      p.TargetUin,
	}

	if p.DurationSeconds != 0 {
		body["DurationSeconds"] = p.DurationSeconds
	}
	if p.Extra != "" {
		body["Extra"] = p.Extra
	}
	if p.TargetAction != nil {
		body["TargetAction"] = p.TargetAction
	}

	err := request.SetActionParameters(body)
	if err != nil {
		return nil, err
	}

	response := tchttp.NewCommonResponse()
	err = client.Send(request, response)
	if err != nil {
		log.Printf("fail to invoke api: %v \n", err)
		return nil, err
	}
	rsp := userRoleTokenRsp{}
	content := response.GetBody()
	log.Println(string(content))
	_ = json.Unmarshal(content, &rsp)

	cre := &UserRoleTokenCredential{
		tmpSecretId:  rsp.credentials.tmpSecretId,
		tmpSecretKey: rsp.credentials.tmpSecretKey,
		token:        rsp.credentials.token,
		expiredTime:  rsp.expiredTime,
		expiration:   rsp.expiration,
		source:       p,
	}
	return cre, nil
}
