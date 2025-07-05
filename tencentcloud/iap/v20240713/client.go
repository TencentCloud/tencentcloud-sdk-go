// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20240713

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-07-13"

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


func NewCreateIAPUserOIDCConfigRequest() (request *CreateIAPUserOIDCConfigRequest) {
    request = &CreateIAPUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "CreateIAPUserOIDCConfig")
    
    
    return
}

func NewCreateIAPUserOIDCConfigResponse() (response *CreateIAPUserOIDCConfigResponse) {
    response = &CreateIAPUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIAPUserOIDCConfig
// 创建用户OIDC配置。只能创建一个用户OIDC身份提供商，并且创建用户OIDC配置之后会自动关闭用户SAML SSO身份提供商。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateIAPUserOIDCConfig(request *CreateIAPUserOIDCConfigRequest) (response *CreateIAPUserOIDCConfigResponse, err error) {
    return c.CreateIAPUserOIDCConfigWithContext(context.Background(), request)
}

// CreateIAPUserOIDCConfig
// 创建用户OIDC配置。只能创建一个用户OIDC身份提供商，并且创建用户OIDC配置之后会自动关闭用户SAML SSO身份提供商。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateIAPUserOIDCConfigWithContext(ctx context.Context, request *CreateIAPUserOIDCConfigRequest) (response *CreateIAPUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewCreateIAPUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIAPUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIAPUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIAPLoginSessionDurationRequest() (request *DescribeIAPLoginSessionDurationRequest) {
    request = &DescribeIAPLoginSessionDurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "DescribeIAPLoginSessionDuration")
    
    
    return
}

func NewDescribeIAPLoginSessionDurationResponse() (response *DescribeIAPLoginSessionDurationResponse) {
    response = &DescribeIAPLoginSessionDurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIAPLoginSessionDuration
// 查询登录会话时长
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_RECORDNOTEXISTS = "ResourceNotFound.RecordNotExists"
func (c *Client) DescribeIAPLoginSessionDuration(request *DescribeIAPLoginSessionDurationRequest) (response *DescribeIAPLoginSessionDurationResponse, err error) {
    return c.DescribeIAPLoginSessionDurationWithContext(context.Background(), request)
}

// DescribeIAPLoginSessionDuration
// 查询登录会话时长
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_RECORDNOTEXISTS = "ResourceNotFound.RecordNotExists"
func (c *Client) DescribeIAPLoginSessionDurationWithContext(ctx context.Context, request *DescribeIAPLoginSessionDurationRequest) (response *DescribeIAPLoginSessionDurationResponse, err error) {
    if request == nil {
        request = NewDescribeIAPLoginSessionDurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIAPLoginSessionDuration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIAPLoginSessionDurationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIAPUserOIDCConfigRequest() (request *DescribeIAPUserOIDCConfigRequest) {
    request = &DescribeIAPUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "DescribeIAPUserOIDCConfig")
    
    
    return
}

func NewDescribeIAPUserOIDCConfigResponse() (response *DescribeIAPUserOIDCConfigResponse) {
    response = &DescribeIAPUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIAPUserOIDCConfig
// 查询用户OIDC配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeIAPUserOIDCConfig(request *DescribeIAPUserOIDCConfigRequest) (response *DescribeIAPUserOIDCConfigResponse, err error) {
    return c.DescribeIAPUserOIDCConfigWithContext(context.Background(), request)
}

// DescribeIAPUserOIDCConfig
// 查询用户OIDC配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeIAPUserOIDCConfigWithContext(ctx context.Context, request *DescribeIAPUserOIDCConfigRequest) (response *DescribeIAPUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewDescribeIAPUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIAPUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIAPUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisableIAPUserSSORequest() (request *DisableIAPUserSSORequest) {
    request = &DisableIAPUserSSORequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "DisableIAPUserSSO")
    
    
    return
}

func NewDisableIAPUserSSOResponse() (response *DisableIAPUserSSOResponse) {
    response = &DisableIAPUserSSOResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableIAPUserSSO
// 禁用用户SSO
//
// 可能返回的错误码:
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
func (c *Client) DisableIAPUserSSO(request *DisableIAPUserSSORequest) (response *DisableIAPUserSSOResponse, err error) {
    return c.DisableIAPUserSSOWithContext(context.Background(), request)
}

// DisableIAPUserSSO
// 禁用用户SSO
//
// 可能返回的错误码:
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
func (c *Client) DisableIAPUserSSOWithContext(ctx context.Context, request *DisableIAPUserSSORequest) (response *DisableIAPUserSSOResponse, err error) {
    if request == nil {
        request = NewDisableIAPUserSSORequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableIAPUserSSO require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableIAPUserSSOResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIAPLoginSessionDurationRequest() (request *ModifyIAPLoginSessionDurationRequest) {
    request = &ModifyIAPLoginSessionDurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "ModifyIAPLoginSessionDuration")
    
    
    return
}

func NewModifyIAPLoginSessionDurationResponse() (response *ModifyIAPLoginSessionDurationResponse) {
    response = &ModifyIAPLoginSessionDurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIAPLoginSessionDuration
// 修改登录会话时长
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyIAPLoginSessionDuration(request *ModifyIAPLoginSessionDurationRequest) (response *ModifyIAPLoginSessionDurationResponse, err error) {
    return c.ModifyIAPLoginSessionDurationWithContext(context.Background(), request)
}

// ModifyIAPLoginSessionDuration
// 修改登录会话时长
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyIAPLoginSessionDurationWithContext(ctx context.Context, request *ModifyIAPLoginSessionDurationRequest) (response *ModifyIAPLoginSessionDurationResponse, err error) {
    if request == nil {
        request = NewModifyIAPLoginSessionDurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIAPLoginSessionDuration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIAPLoginSessionDurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateIAPUserOIDCConfigRequest() (request *UpdateIAPUserOIDCConfigRequest) {
    request = &UpdateIAPUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iap", APIVersion, "UpdateIAPUserOIDCConfig")
    
    
    return
}

func NewUpdateIAPUserOIDCConfigResponse() (response *UpdateIAPUserOIDCConfigResponse) {
    response = &UpdateIAPUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateIAPUserOIDCConfig
// 修改用户OIDC配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateIAPUserOIDCConfig(request *UpdateIAPUserOIDCConfigRequest) (response *UpdateIAPUserOIDCConfigResponse, err error) {
    return c.UpdateIAPUserOIDCConfigWithContext(context.Background(), request)
}

// UpdateIAPUserOIDCConfig
// 修改用户OIDC配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateIAPUserOIDCConfigWithContext(ctx context.Context, request *UpdateIAPUserOIDCConfigRequest) (response *UpdateIAPUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewUpdateIAPUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateIAPUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateIAPUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}
