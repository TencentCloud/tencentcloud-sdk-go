// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180625

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-25"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBindEipAclsRequest() (request *BindEipAclsRequest) {
    request = &BindEipAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "BindEipAcls")
    return
}

func NewBindEipAclsResponse() (response *BindEipAclsResponse) {
    response = &BindEipAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于为某个 EIP 关联 ACL。
func (c *Client) BindEipAcls(request *BindEipAclsRequest) (response *BindEipAclsResponse, err error) {
    if request == nil {
        request = NewBindEipAclsRequest()
    }
    response = NewBindEipAclsResponse()
    err = c.Send(request, response)
    return
}

func NewBindHostedRequest() (request *BindHostedRequest) {
    request = &BindHostedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "BindHosted")
    return
}

func NewBindHostedResponse() (response *BindHostedResponse) {
    response = &BindHostedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindHosted接口用于绑定黑石弹性公网IP到黑石托管机器上
func (c *Client) BindHosted(request *BindHostedRequest) (response *BindHostedResponse, err error) {
    if request == nil {
        request = NewBindHostedRequest()
    }
    response = NewBindHostedResponse()
    err = c.Send(request, response)
    return
}

func NewBindRsRequest() (request *BindRsRequest) {
    request = &BindRsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "BindRs")
    return
}

func NewBindRsResponse() (response *BindRsResponse) {
    response = &BindRsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 绑定黑石EIP
func (c *Client) BindRs(request *BindRsRequest) (response *BindRsResponse, err error) {
    if request == nil {
        request = NewBindRsRequest()
    }
    response = NewBindRsResponse()
    err = c.Send(request, response)
    return
}

func NewBindVpcIpRequest() (request *BindVpcIpRequest) {
    request = &BindVpcIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "BindVpcIp")
    return
}

func NewBindVpcIpResponse() (response *BindVpcIpResponse) {
    response = &BindVpcIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑石EIP绑定VPCIP
func (c *Client) BindVpcIp(request *BindVpcIpRequest) (response *BindVpcIpResponse, err error) {
    if request == nil {
        request = NewBindVpcIpRequest()
    }
    response = NewBindVpcIpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEipRequest() (request *CreateEipRequest) {
    request = &CreateEipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "CreateEip")
    return
}

func NewCreateEipResponse() (response *CreateEipResponse) {
    response = &CreateEipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石弹性公网IP
func (c *Client) CreateEip(request *CreateEipRequest) (response *CreateEipResponse, err error) {
    if request == nil {
        request = NewCreateEipRequest()
    }
    response = NewCreateEipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEipAclRequest() (request *CreateEipAclRequest) {
    request = &CreateEipAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "CreateEipAcl")
    return
}

func NewCreateEipAclResponse() (response *CreateEipAclResponse) {
    response = &CreateEipAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石弹性公网 EIPACL
func (c *Client) CreateEipAcl(request *CreateEipAclRequest) (response *CreateEipAclResponse, err error) {
    if request == nil {
        request = NewCreateEipAclRequest()
    }
    response = NewCreateEipAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEipRequest() (request *DeleteEipRequest) {
    request = &DeleteEipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DeleteEip")
    return
}

func NewDeleteEipResponse() (response *DeleteEipResponse) {
    response = &DeleteEipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 释放黑石弹性公网IP
func (c *Client) DeleteEip(request *DeleteEipRequest) (response *DeleteEipResponse, err error) {
    if request == nil {
        request = NewDeleteEipRequest()
    }
    response = NewDeleteEipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEipAclRequest() (request *DeleteEipAclRequest) {
    request = &DeleteEipAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DeleteEipAcl")
    return
}

func NewDeleteEipAclResponse() (response *DeleteEipAclResponse) {
    response = &DeleteEipAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除弹性公网IP ACL
func (c *Client) DeleteEipAcl(request *DeleteEipAclRequest) (response *DeleteEipAclResponse, err error) {
    if request == nil {
        request = NewDeleteEipAclRequest()
    }
    response = NewDeleteEipAclResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEipAclsRequest() (request *DescribeEipAclsRequest) {
    request = &DescribeEipAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DescribeEipAcls")
    return
}

func NewDescribeEipAclsResponse() (response *DescribeEipAclsResponse) {
    response = &DescribeEipAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询弹性公网IP ACL
func (c *Client) DescribeEipAcls(request *DescribeEipAclsRequest) (response *DescribeEipAclsResponse, err error) {
    if request == nil {
        request = NewDescribeEipAclsRequest()
    }
    response = NewDescribeEipAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEipQuotaRequest() (request *DescribeEipQuotaRequest) {
    request = &DescribeEipQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DescribeEipQuota")
    return
}

func NewDescribeEipQuotaResponse() (response *DescribeEipQuotaResponse) {
    response = &DescribeEipQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询黑石EIP 限额
func (c *Client) DescribeEipQuota(request *DescribeEipQuotaRequest) (response *DescribeEipQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEipQuotaRequest()
    }
    response = NewDescribeEipQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEipTaskRequest() (request *DescribeEipTaskRequest) {
    request = &DescribeEipTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DescribeEipTask")
    return
}

func NewDescribeEipTaskResponse() (response *DescribeEipTaskResponse) {
    response = &DescribeEipTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑石EIP查询任务状态
func (c *Client) DescribeEipTask(request *DescribeEipTaskRequest) (response *DescribeEipTaskResponse, err error) {
    if request == nil {
        request = NewDescribeEipTaskRequest()
    }
    response = NewDescribeEipTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEipsRequest() (request *DescribeEipsRequest) {
    request = &DescribeEipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "DescribeEips")
    return
}

func NewDescribeEipsResponse() (response *DescribeEipsResponse) {
    response = &DescribeEipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑石EIP查询接口
func (c *Client) DescribeEips(request *DescribeEipsRequest) (response *DescribeEipsResponse, err error) {
    if request == nil {
        request = NewDescribeEipsRequest()
    }
    response = NewDescribeEipsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEipAclRequest() (request *ModifyEipAclRequest) {
    request = &ModifyEipAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "ModifyEipAcl")
    return
}

func NewModifyEipAclResponse() (response *ModifyEipAclResponse) {
    response = &ModifyEipAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改弹性公网IP ACL
func (c *Client) ModifyEipAcl(request *ModifyEipAclRequest) (response *ModifyEipAclResponse, err error) {
    if request == nil {
        request = NewModifyEipAclRequest()
    }
    response = NewModifyEipAclResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEipChargeRequest() (request *ModifyEipChargeRequest) {
    request = &ModifyEipChargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "ModifyEipCharge")
    return
}

func NewModifyEipChargeResponse() (response *ModifyEipChargeResponse) {
    response = &ModifyEipChargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑石EIP修改计费方式
func (c *Client) ModifyEipCharge(request *ModifyEipChargeRequest) (response *ModifyEipChargeResponse, err error) {
    if request == nil {
        request = NewModifyEipChargeRequest()
    }
    response = NewModifyEipChargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEipNameRequest() (request *ModifyEipNameRequest) {
    request = &ModifyEipNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "ModifyEipName")
    return
}

func NewModifyEipNameResponse() (response *ModifyEipNameResponse) {
    response = &ModifyEipNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新黑石EIP名称
func (c *Client) ModifyEipName(request *ModifyEipNameRequest) (response *ModifyEipNameResponse, err error) {
    if request == nil {
        request = NewModifyEipNameRequest()
    }
    response = NewModifyEipNameResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindEipAclsRequest() (request *UnbindEipAclsRequest) {
    request = &UnbindEipAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "UnbindEipAcls")
    return
}

func NewUnbindEipAclsResponse() (response *UnbindEipAclsResponse) {
    response = &UnbindEipAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑弹性公网IP ACL
func (c *Client) UnbindEipAcls(request *UnbindEipAclsRequest) (response *UnbindEipAclsResponse, err error) {
    if request == nil {
        request = NewUnbindEipAclsRequest()
    }
    response = NewUnbindEipAclsResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindHostedRequest() (request *UnbindHostedRequest) {
    request = &UnbindHostedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "UnbindHosted")
    return
}

func NewUnbindHostedResponse() (response *UnbindHostedResponse) {
    response = &UnbindHostedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindHosted接口用于解绑托管机器上的EIP
func (c *Client) UnbindHosted(request *UnbindHostedRequest) (response *UnbindHostedResponse, err error) {
    if request == nil {
        request = NewUnbindHostedRequest()
    }
    response = NewUnbindHostedResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindRsRequest() (request *UnbindRsRequest) {
    request = &UnbindRsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "UnbindRs")
    return
}

func NewUnbindRsResponse() (response *UnbindRsResponse) {
    response = &UnbindRsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑黑石EIP
func (c *Client) UnbindRs(request *UnbindRsRequest) (response *UnbindRsResponse, err error) {
    if request == nil {
        request = NewUnbindRsRequest()
    }
    response = NewUnbindRsResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindVpcIpRequest() (request *UnbindVpcIpRequest) {
    request = &UnbindVpcIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bmeip", APIVersion, "UnbindVpcIp")
    return
}

func NewUnbindVpcIpResponse() (response *UnbindVpcIpResponse) {
    response = &UnbindVpcIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 黑石EIP解绑VPCIP
func (c *Client) UnbindVpcIp(request *UnbindVpcIpRequest) (response *UnbindVpcIpResponse, err error) {
    if request == nil {
        request = NewUnbindVpcIpRequest()
    }
    response = NewUnbindVpcIpResponse()
    err = c.Send(request, response)
    return
}
