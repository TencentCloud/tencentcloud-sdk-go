package v2

import (
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/teamlint/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init("").
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func newCmqBaseRequest(action string) *tchttp.BaseRequest {
	req := &tchttp.BaseRequest{}
	req.Init().WithApiInfo("", "", action)
	req.SetPath(V2Path)
	return req
}
func newCmqBaseResponse() *tchttp.BaseResponse {
	return &tchttp.BaseResponse{}

}
