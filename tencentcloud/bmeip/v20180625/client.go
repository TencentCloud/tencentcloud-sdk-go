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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// BindEipAcls
// 此接口用于为某个 EIP 关联 ACL。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindEipAcls(request *BindEipAclsRequest) (response *BindEipAclsResponse, err error) {
    return c.BindEipAclsWithContext(context.Background(), request)
}

// BindEipAcls
// 此接口用于为某个 EIP 关联 ACL。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindEipAclsWithContext(ctx context.Context, request *BindEipAclsRequest) (response *BindEipAclsResponse, err error) {
    if request == nil {
        request = NewBindEipAclsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindEipAcls require credential")
    }

    request.SetContext(ctx)
    
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

// BindHosted
// BindHosted接口用于绑定黑石弹性公网IP到黑石托管机器上
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindHosted(request *BindHostedRequest) (response *BindHostedResponse, err error) {
    return c.BindHostedWithContext(context.Background(), request)
}

// BindHosted
// BindHosted接口用于绑定黑石弹性公网IP到黑石托管机器上
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindHostedWithContext(ctx context.Context, request *BindHostedRequest) (response *BindHostedResponse, err error) {
    if request == nil {
        request = NewBindHostedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindHosted require credential")
    }

    request.SetContext(ctx)
    
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

// BindRs
// 绑定黑石EIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindRs(request *BindRsRequest) (response *BindRsResponse, err error) {
    return c.BindRsWithContext(context.Background(), request)
}

// BindRs
// 绑定黑石EIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindRsWithContext(ctx context.Context, request *BindRsRequest) (response *BindRsResponse, err error) {
    if request == nil {
        request = NewBindRsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindRs require credential")
    }

    request.SetContext(ctx)
    
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

// BindVpcIp
// 黑石EIP绑定VPC IP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindVpcIp(request *BindVpcIpRequest) (response *BindVpcIpResponse, err error) {
    return c.BindVpcIpWithContext(context.Background(), request)
}

// BindVpcIp
// 黑石EIP绑定VPC IP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindVpcIpWithContext(ctx context.Context, request *BindVpcIpRequest) (response *BindVpcIpResponse, err error) {
    if request == nil {
        request = NewBindVpcIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindVpcIp require credential")
    }

    request.SetContext(ctx)
    
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

// CreateEip
// 创建黑石弹性公网IP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEip(request *CreateEipRequest) (response *CreateEipResponse, err error) {
    return c.CreateEipWithContext(context.Background(), request)
}

// CreateEip
// 创建黑石弹性公网IP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEipWithContext(ctx context.Context, request *CreateEipRequest) (response *CreateEipResponse, err error) {
    if request == nil {
        request = NewCreateEipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEip require credential")
    }

    request.SetContext(ctx)
    
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

// CreateEipAcl
// 创建黑石弹性公网 EIP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEipAcl(request *CreateEipAclRequest) (response *CreateEipAclResponse, err error) {
    return c.CreateEipAclWithContext(context.Background(), request)
}

// CreateEipAcl
// 创建黑石弹性公网 EIP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateEipAclWithContext(ctx context.Context, request *CreateEipAclRequest) (response *CreateEipAclResponse, err error) {
    if request == nil {
        request = NewCreateEipAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEipAcl require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteEip
// 释放黑石弹性公网IP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteEip(request *DeleteEipRequest) (response *DeleteEipResponse, err error) {
    return c.DeleteEipWithContext(context.Background(), request)
}

// DeleteEip
// 释放黑石弹性公网IP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteEipWithContext(ctx context.Context, request *DeleteEipRequest) (response *DeleteEipResponse, err error) {
    if request == nil {
        request = NewDeleteEipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEip require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteEipAcl
// 删除弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteEipAcl(request *DeleteEipAclRequest) (response *DeleteEipAclResponse, err error) {
    return c.DeleteEipAclWithContext(context.Background(), request)
}

// DeleteEipAcl
// 删除弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteEipAclWithContext(ctx context.Context, request *DeleteEipAclRequest) (response *DeleteEipAclResponse, err error) {
    if request == nil {
        request = NewDeleteEipAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEipAcl require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEipAcls
// 查询弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEipAcls(request *DescribeEipAclsRequest) (response *DescribeEipAclsResponse, err error) {
    return c.DescribeEipAclsWithContext(context.Background(), request)
}

// DescribeEipAcls
// 查询弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEipAclsWithContext(ctx context.Context, request *DescribeEipAclsRequest) (response *DescribeEipAclsResponse, err error) {
    if request == nil {
        request = NewDescribeEipAclsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEipAcls require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEipQuota
// 查询黑石EIP 限额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEipQuota(request *DescribeEipQuotaRequest) (response *DescribeEipQuotaResponse, err error) {
    return c.DescribeEipQuotaWithContext(context.Background(), request)
}

// DescribeEipQuota
// 查询黑石EIP 限额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEipQuotaWithContext(ctx context.Context, request *DescribeEipQuotaRequest) (response *DescribeEipQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEipQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEipQuota require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEipTask
// 黑石EIP查询任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEipTask(request *DescribeEipTaskRequest) (response *DescribeEipTaskResponse, err error) {
    return c.DescribeEipTaskWithContext(context.Background(), request)
}

// DescribeEipTask
// 黑石EIP查询任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEipTaskWithContext(ctx context.Context, request *DescribeEipTaskRequest) (response *DescribeEipTaskResponse, err error) {
    if request == nil {
        request = NewDescribeEipTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEipTask require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeEips
// 黑石EIP查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEips(request *DescribeEipsRequest) (response *DescribeEipsResponse, err error) {
    return c.DescribeEipsWithContext(context.Background(), request)
}

// DescribeEips
// 黑石EIP查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEipsWithContext(ctx context.Context, request *DescribeEipsRequest) (response *DescribeEipsResponse, err error) {
    if request == nil {
        request = NewDescribeEipsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEips require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyEipAcl
// 修改弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipAcl(request *ModifyEipAclRequest) (response *ModifyEipAclResponse, err error) {
    return c.ModifyEipAclWithContext(context.Background(), request)
}

// ModifyEipAcl
// 修改弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipAclWithContext(ctx context.Context, request *ModifyEipAclRequest) (response *ModifyEipAclResponse, err error) {
    if request == nil {
        request = NewModifyEipAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEipAcl require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyEipCharge
// 黑石EIP修改计费方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipCharge(request *ModifyEipChargeRequest) (response *ModifyEipChargeResponse, err error) {
    return c.ModifyEipChargeWithContext(context.Background(), request)
}

// ModifyEipCharge
// 黑石EIP修改计费方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipChargeWithContext(ctx context.Context, request *ModifyEipChargeRequest) (response *ModifyEipChargeResponse, err error) {
    if request == nil {
        request = NewModifyEipChargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEipCharge require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyEipName
// 更新黑石EIP名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipName(request *ModifyEipNameRequest) (response *ModifyEipNameResponse, err error) {
    return c.ModifyEipNameWithContext(context.Background(), request)
}

// ModifyEipName
// 更新黑石EIP名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyEipNameWithContext(ctx context.Context, request *ModifyEipNameRequest) (response *ModifyEipNameResponse, err error) {
    if request == nil {
        request = NewModifyEipNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEipName require credential")
    }

    request.SetContext(ctx)
    
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

// UnbindEipAcls
// 解绑弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindEipAcls(request *UnbindEipAclsRequest) (response *UnbindEipAclsResponse, err error) {
    return c.UnbindEipAclsWithContext(context.Background(), request)
}

// UnbindEipAcls
// 解绑弹性公网IP ACL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindEipAclsWithContext(ctx context.Context, request *UnbindEipAclsRequest) (response *UnbindEipAclsResponse, err error) {
    if request == nil {
        request = NewUnbindEipAclsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindEipAcls require credential")
    }

    request.SetContext(ctx)
    
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

// UnbindHosted
// UnbindHosted接口用于解绑托管机器上的EIP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindHosted(request *UnbindHostedRequest) (response *UnbindHostedResponse, err error) {
    return c.UnbindHostedWithContext(context.Background(), request)
}

// UnbindHosted
// UnbindHosted接口用于解绑托管机器上的EIP
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindHostedWithContext(ctx context.Context, request *UnbindHostedRequest) (response *UnbindHostedResponse, err error) {
    if request == nil {
        request = NewUnbindHostedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindHosted require credential")
    }

    request.SetContext(ctx)
    
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

// UnbindRs
// 解绑黑石EIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindRs(request *UnbindRsRequest) (response *UnbindRsResponse, err error) {
    return c.UnbindRsWithContext(context.Background(), request)
}

// UnbindRs
// 解绑黑石EIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindRsWithContext(ctx context.Context, request *UnbindRsRequest) (response *UnbindRsResponse, err error) {
    if request == nil {
        request = NewUnbindRsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindRs require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindRsResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindRsListRequest() (request *UnbindRsListRequest) {
    request = &UnbindRsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bmeip", APIVersion, "UnbindRsList")
    
    
    return
}

func NewUnbindRsListResponse() (response *UnbindRsListResponse) {
    response = &UnbindRsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindRsList
// 批量解绑物理机弹性公网IP接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UnbindRsList(request *UnbindRsListRequest) (response *UnbindRsListResponse, err error) {
    return c.UnbindRsListWithContext(context.Background(), request)
}

// UnbindRsList
// 批量解绑物理机弹性公网IP接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UnbindRsListWithContext(ctx context.Context, request *UnbindRsListRequest) (response *UnbindRsListResponse, err error) {
    if request == nil {
        request = NewUnbindRsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindRsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindRsListResponse()
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

// UnbindVpcIp
// 黑石EIP解绑VPCIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindVpcIp(request *UnbindVpcIpRequest) (response *UnbindVpcIpResponse, err error) {
    return c.UnbindVpcIpWithContext(context.Background(), request)
}

// UnbindVpcIp
// 黑石EIP解绑VPCIP
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnbindVpcIpWithContext(ctx context.Context, request *UnbindVpcIpRequest) (response *UnbindVpcIpResponse, err error) {
    if request == nil {
        request = NewUnbindVpcIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindVpcIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindVpcIpResponse()
    err = c.Send(request, response)
    return
}
