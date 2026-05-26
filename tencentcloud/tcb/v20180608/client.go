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

package v20180608

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

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


func NewAddProviderRequest() (request *AddProviderRequest) {
    request = &AddProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "AddProvider")
    
    
    return
}

func NewAddProviderResponse() (response *AddProviderResponse) {
    response = &AddProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddProvider
// 添加身份认证源。在指定云开发环境下创建一个新的身份认证源，支持 OAuth 2.0、OIDC、SAML 2.0 等标准协议，以及自定义登录和邮箱登录等多种认证方式。
//
// 创建时需指定身份源协议类型（ProviderType）并配置对应的协议连接参数（Config）。若身份源 ID 已存在将返回错误。
//
// 限制：一个环境最大可允许加入20个认证源。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddProvider(request *AddProviderRequest) (response *AddProviderResponse, err error) {
    return c.AddProviderWithContext(context.Background(), request)
}

// AddProvider
// 添加身份认证源。在指定云开发环境下创建一个新的身份认证源，支持 OAuth 2.0、OIDC、SAML 2.0 等标准协议，以及自定义登录和邮箱登录等多种认证方式。
//
// 创建时需指定身份源协议类型（ProviderType）并配置对应的协议连接参数（Config）。若身份源 ID 已存在将返回错误。
//
// 限制：一个环境最大可允许加入20个认证源。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddProviderWithContext(ctx context.Context, request *AddProviderRequest) (response *AddProviderResponse, err error) {
    if request == nil {
        request = NewAddProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "AddProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddProviderResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateEnvRequest() (request *AllocateEnvRequest) {
    request = &AllocateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "AllocateEnv")
    
    
    return
}

func NewAllocateEnvResponse() (response *AllocateEnvResponse) {
    response = &AllocateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateEnv
// 从环境池里立即取出1个环境
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALLOCATEIDRELEASED = "FailedOperation.AllocateIdReleased"
//  INVALIDPARAMETER_CUSTOMERNOTCONFIGURED = "InvalidParameter.CustomerNotConfigured"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_ENVPOOLEMPTY = "ResourceInsufficient.EnvPoolEmpty"
func (c *Client) AllocateEnv(request *AllocateEnvRequest) (response *AllocateEnvResponse, err error) {
    return c.AllocateEnvWithContext(context.Background(), request)
}

// AllocateEnv
// 从环境池里立即取出1个环境
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALLOCATEIDRELEASED = "FailedOperation.AllocateIdReleased"
//  INVALIDPARAMETER_CUSTOMERNOTCONFIGURED = "InvalidParameter.CustomerNotConfigured"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_ENVPOOLEMPTY = "ResourceInsufficient.EnvPoolEmpty"
func (c *Client) AllocateEnvWithContext(ctx context.Context, request *AllocateEnvRequest) (response *AllocateEnvResponse, err error) {
    if request == nil {
        request = NewAllocateEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "AllocateEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateEnvResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleForAllocatedEnvRequest() (request *AssumeRoleForAllocatedEnvRequest) {
    request = &AssumeRoleForAllocatedEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "AssumeRoleForAllocatedEnv")
    
    
    return
}

func NewAssumeRoleForAllocatedEnvResponse() (response *AssumeRoleForAllocatedEnvResponse) {
    response = &AssumeRoleForAllocatedEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssumeRoleForAllocatedEnv
// 白名单接口，申请Tcb角色临时凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssumeRoleForAllocatedEnv(request *AssumeRoleForAllocatedEnvRequest) (response *AssumeRoleForAllocatedEnvResponse, err error) {
    return c.AssumeRoleForAllocatedEnvWithContext(context.Background(), request)
}

// AssumeRoleForAllocatedEnv
// 白名单接口，申请Tcb角色临时凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssumeRoleForAllocatedEnvWithContext(ctx context.Context, request *AssumeRoleForAllocatedEnvRequest) (response *AssumeRoleForAllocatedEnvResponse, err error) {
    if request == nil {
        request = NewAssumeRoleForAllocatedEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "AssumeRoleForAllocatedEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssumeRoleForAllocatedEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssumeRoleForAllocatedEnvResponse()
    err = c.Send(request, response)
    return
}

func NewBindStorageSourceRequest() (request *BindStorageSourceRequest) {
    request = &BindStorageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "BindStorageSource")
    
    
    return
}

func NewBindStorageSourceResponse() (response *BindStorageSourceResponse) {
    response = &BindStorageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindStorageSource
// 为云存储绑定外部云存储源。
//
// 将一个用户自有的 COS桶 作为外部存储源绑定到指定云开发环境的云存储。绑定后，该环境的云存储文件操作将指向此桶，通过 BasePath 路径前缀实现与其他环境的数据隔离。
//
// 每个环境仅允许绑定 1 个外部云存储源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) BindStorageSource(request *BindStorageSourceRequest) (response *BindStorageSourceResponse, err error) {
    return c.BindStorageSourceWithContext(context.Background(), request)
}

// BindStorageSource
// 为云存储绑定外部云存储源。
//
// 将一个用户自有的 COS桶 作为外部存储源绑定到指定云开发环境的云存储。绑定后，该环境的云存储文件操作将指向此桶，通过 BasePath 路径前缀实现与其他环境的数据隔离。
//
// 每个环境仅允许绑定 1 个外部云存储源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) BindStorageSourceWithContext(ctx context.Context, request *BindStorageSourceRequest) (response *BindStorageSourceResponse, err error) {
    if request == nil {
        request = NewBindStorageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "BindStorageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindStorageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindStorageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckTcbServiceRequest() (request *CheckTcbServiceRequest) {
    request = &CheckTcbServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CheckTcbService")
    
    
    return
}

func NewCheckTcbServiceResponse() (response *CheckTcbServiceResponse) {
    response = &CheckTcbServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbService(request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    return c.CheckTcbServiceWithContext(context.Background(), request)
}

// CheckTcbService
// 检查是否开通Tcb服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) CheckTcbServiceWithContext(ctx context.Context, request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    if request == nil {
        request = NewCheckTcbServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CheckTcbService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckTcbService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckTcbServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIModelRequest() (request *CreateAIModelRequest) {
    request = &CreateAIModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAIModel")
    
    
    return
}

func NewCreateAIModelResponse() (response *CreateAIModelResponse) {
    response = &CreateAIModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAIModel
// 创建 AI 模型配置分组。云开发标准版及以上套餐才可使用
//
// 
//
// 支持自定义类型（custom）：用户自行提供模型服务地址（baseUrl）和访问密钥，分组名可自由命名，适用于接入第三方或自建模型服务。
//
// 
//
// 注意：内置类型（builtin）分组由云开发平台统一创建和管理，不支持通过此接口创建。如需修改内置分组的模型列表或启用状态，请使用 [UpdateAIModel](https://cloud.tencent.com/document/product/876/131316) 接口。
//
// 
//
// 创建成功后，可通过 [DescribeAIModels](https://cloud.tencent.com/document/product/876/131318) 接口查询分组信息，并在云开发 AI+ 功能中使用所配置的模型。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) CreateAIModel(request *CreateAIModelRequest) (response *CreateAIModelResponse, err error) {
    return c.CreateAIModelWithContext(context.Background(), request)
}

// CreateAIModel
// 创建 AI 模型配置分组。云开发标准版及以上套餐才可使用
//
// 
//
// 支持自定义类型（custom）：用户自行提供模型服务地址（baseUrl）和访问密钥，分组名可自由命名，适用于接入第三方或自建模型服务。
//
// 
//
// 注意：内置类型（builtin）分组由云开发平台统一创建和管理，不支持通过此接口创建。如需修改内置分组的模型列表或启用状态，请使用 [UpdateAIModel](https://cloud.tencent.com/document/product/876/131316) 接口。
//
// 
//
// 创建成功后，可通过 [DescribeAIModels](https://cloud.tencent.com/document/product/876/131318) 接口查询分组信息，并在云开发 AI+ 功能中使用所配置的模型。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) CreateAIModelWithContext(ctx context.Context, request *CreateAIModelRequest) (response *CreateAIModelResponse, err error) {
    if request == nil {
        request = NewCreateAIModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateAIModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateApiKey")
    
    
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiKey
// 创建云开发平台的API Key。在指定云开发环境下创建一个 API Key 访问凭证。支持两种类型：api_key（服务端管理员访问凭证，以管理员身份签发，可设置有效期，不设置有效期则永不过期，单个环境最多创建 5 个）和 publish_key（前端匿名访问凭证，固定有效期，每个环境仅保留一个）。创建成功后将返回 API Key 明文 Token，该值仅在创建时返回一次，请妥善保存。需要管理员权限。
//
// 权限范围与注意事项：
//
//  - api_key 是以"管理员身份"签发的访问凭证，持有该 Token 即拥有腾讯云云开发数据流（含数据库、云函数、云存储、托管等）资源的完整访问与操作权限。该凭证不区分子资源粒度，一旦泄露等同于环境管理员权限被泄露，请按"高敏感凭证"管理： 仅用于服务端调用、严禁下发到前端/客户端、严禁写入代码仓库或日志、定期轮换并回收闲置凭证。
//
//  - 在为子账号 / 协作者 / RAM 子用户授权该接口权限时，须先评估该子账号是否有资格获得该环境的管理员级数据流权限，能创建 api_key ≈ 能拥有环境管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    return c.CreateApiKeyWithContext(context.Background(), request)
}

// CreateApiKey
// 创建云开发平台的API Key。在指定云开发环境下创建一个 API Key 访问凭证。支持两种类型：api_key（服务端管理员访问凭证，以管理员身份签发，可设置有效期，不设置有效期则永不过期，单个环境最多创建 5 个）和 publish_key（前端匿名访问凭证，固定有效期，每个环境仅保留一个）。创建成功后将返回 API Key 明文 Token，该值仅在创建时返回一次，请妥善保存。需要管理员权限。
//
// 权限范围与注意事项：
//
//  - api_key 是以"管理员身份"签发的访问凭证，持有该 Token 即拥有腾讯云云开发数据流（含数据库、云函数、云存储、托管等）资源的完整访问与操作权限。该凭证不区分子资源粒度，一旦泄露等同于环境管理员权限被泄露，请按"高敏感凭证"管理： 仅用于服务端调用、严禁下发到前端/客户端、严禁写入代码仓库或日志、定期轮换并回收闲置凭证。
//
//  - 在为子账号 / 协作者 / RAM 子用户授权该接口权限时，须先评估该子账号是否有资格获得该环境的管理员级数据流权限，能创建 api_key ≈ 能拥有环境管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthDomainRequest() (request *CreateAuthDomainRequest) {
    request = &CreateAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAuthDomain")
    
    
    return
}

func NewCreateAuthDomainResponse() (response *CreateAuthDomainResponse) {
    response = &CreateAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuthDomain
// 增加安全域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
//   安全域名绑定成功之后，需要几分钟时间逐步生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomain(request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    return c.CreateAuthDomainWithContext(context.Background(), request)
}

// CreateAuthDomain
// 增加安全域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名。
//
// 
//
// 注意⚠️
//
//   安全域名绑定成功之后，需要几分钟时间逐步生效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuthDomainWithContext(ctx context.Context, request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateAuthDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuthDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBillDealRequest() (request *CreateBillDealRequest) {
    request = &CreateBillDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateBillDeal")
    
    
    return
}

func NewCreateBillDealResponse() (response *CreateBillDealResponse) {
    response = &CreateBillDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBillDeal
// 创建云开发产品计费订单，用于以下几种场景：
//
// 1. 购买云开发环境
//
// 2. 续费云开发环境
//
// 3. 变更云开发环境套餐
//
// 4. 购买云开发资源包
//
// 5. 购买云开发大促包
//
// 
//
// 该接口支持下单并支付(CreateAndPay=true时)，此时会自动在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateBillDeal(request *CreateBillDealRequest) (response *CreateBillDealResponse, err error) {
    return c.CreateBillDealWithContext(context.Background(), request)
}

// CreateBillDeal
// 创建云开发产品计费订单，用于以下几种场景：
//
// 1. 购买云开发环境
//
// 2. 续费云开发环境
//
// 3. 变更云开发环境套餐
//
// 4. 购买云开发资源包
//
// 5. 购买云开发大促包
//
// 
//
// 该接口支持下单并支付(CreateAndPay=true时)，此时会自动在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_FREEPACKAGEDENIED = "OperationDenied.FreePackageDenied"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateBillDealWithContext(ctx context.Context, request *CreateBillDealRequest) (response *CreateBillDealResponse, err error) {
    if request == nil {
        request = NewCreateBillDealRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateBillDeal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBillDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBillDealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomLoginKeyRequest() (request *CreateCustomLoginKeyRequest) {
    request = &CreateCustomLoginKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateCustomLoginKey")
    
    
    return
}

func NewCreateCustomLoginKeyResponse() (response *CreateCustomLoginKeyResponse) {
    response = &CreateCustomLoginKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomLoginKey
// 创建自定义登录密钥。在指定云开发环境下生成一对 RSA 1024 位非对称密钥对，系统仅存储公钥，私钥仅在创建时返回一次且不可恢复，请妥善保存。创建新密钥后，该环境下原有未设置过期时间的旧密钥将被自动标记为 2 小时后过期，请确保客户端及时更新密钥配置。
//
// 返回的 KeyID 和 PrivateKey 需与环境 ID 一起组装为 JSON 配置文件，供客户端 Admin SDK 初始化时使用，文件格式如下：
//
// {
//
//   "private_key_id": "<返回的 KeyID>",
//
//   "private_key": "<返回的 PrivateKey>",
//
//   "env_id": "<请求时传入的 EnvId>"
//
// }
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCustomLoginKey(request *CreateCustomLoginKeyRequest) (response *CreateCustomLoginKeyResponse, err error) {
    return c.CreateCustomLoginKeyWithContext(context.Background(), request)
}

// CreateCustomLoginKey
// 创建自定义登录密钥。在指定云开发环境下生成一对 RSA 1024 位非对称密钥对，系统仅存储公钥，私钥仅在创建时返回一次且不可恢复，请妥善保存。创建新密钥后，该环境下原有未设置过期时间的旧密钥将被自动标记为 2 小时后过期，请确保客户端及时更新密钥配置。
//
// 返回的 KeyID 和 PrivateKey 需与环境 ID 一起组装为 JSON 配置文件，供客户端 Admin SDK 初始化时使用，文件格式如下：
//
// {
//
//   "private_key_id": "<返回的 KeyID>",
//
//   "private_key": "<返回的 PrivateKey>",
//
//   "env_id": "<请求时传入的 EnvId>"
//
// }
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASE = "InternalError.Database"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCustomLoginKeyWithContext(ctx context.Context, request *CreateCustomLoginKeyRequest) (response *CreateCustomLoginKeyResponse, err error) {
    if request == nil {
        request = NewCreateCustomLoginKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateCustomLoginKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomLoginKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomLoginKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvRequest() (request *CreateEnvRequest) {
    request = &CreateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateEnv")
    
    
    return
}

func NewCreateEnvResponse() (response *CreateEnvResponse) {
    response = &CreateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnv
// 本接口用于购买云开发环境。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 环境下单成功之后会返回EnvId。EnvId是全局唯一表示。
//
// 环境发货是异步行为，后续可以通过接口 [DescribeEnvs ](https://cloud.tencent.com/document/product/876/34820) 查询环境状态和各项资源信息；通过 [DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境套餐信息，包括 到期时间、当前套餐等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnv(request *CreateEnvRequest) (response *CreateEnvResponse, err error) {
    return c.CreateEnvWithContext(context.Background(), request)
}

// CreateEnv
// 本接口用于购买云开发环境。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 环境下单成功之后会返回EnvId。EnvId是全局唯一表示。
//
// 环境发货是异步行为，后续可以通过接口 [DescribeEnvs ](https://cloud.tencent.com/document/product/876/34820) 查询环境状态和各项资源信息；通过 [DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境套餐信息，包括 到期时间、当前套餐等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnvWithContext(ctx context.Context, request *CreateEnvRequest) (response *CreateEnvResponse, err error) {
    if request == nil {
        request = NewCreateEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvResourceRequest() (request *CreateEnvResourceRequest) {
    request = &CreateEnvResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateEnvResource")
    
    
    return
}

func NewCreateEnvResourceResponse() (response *CreateEnvResourceResponse) {
    response = &CreateEnvResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvResource
// **创建环境日志资源**
//
// 
//
// 环境开通后，若用户需要开启检索日志功能，需调用此接口进行日志资源的开通。
//
// 
//
// > **注意**：日志资源的开通为**异步操作**，接口调用成功后并不代表日志资源已立即可用。
//
// 
//
// **如何确认日志开通状态：**
//
// 
//
// 可通过 [DescribeEnvs](https://cloud.tencent.com/document/product/876/34820) 接口轮询查询日志开通结果，建议每 5 秒轮询一次，最长等待 5 分钟。在返回的数据结构 [EnvInfo](https://cloud.tencent.com/document/api/876/34822#EnvInfo) 中，检查 `LogServices` 字段下 `LogServiceInfo` 是否包含有效的日志主题（Topic）等相关信息，以此判断日志资源是否已成功开通：
//
// 
//
// - **已开通**：`LogServiceInfo` 中存在日志主题 ID 等有效信息
//
// - **未开通 / 开通中**：`LogServiceInfo` 为空或相关字段缺失
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_LOGEXIST = "ResourceInUse.LogExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateEnvResource(request *CreateEnvResourceRequest) (response *CreateEnvResourceResponse, err error) {
    return c.CreateEnvResourceWithContext(context.Background(), request)
}

// CreateEnvResource
// **创建环境日志资源**
//
// 
//
// 环境开通后，若用户需要开启检索日志功能，需调用此接口进行日志资源的开通。
//
// 
//
// > **注意**：日志资源的开通为**异步操作**，接口调用成功后并不代表日志资源已立即可用。
//
// 
//
// **如何确认日志开通状态：**
//
// 
//
// 可通过 [DescribeEnvs](https://cloud.tencent.com/document/product/876/34820) 接口轮询查询日志开通结果，建议每 5 秒轮询一次，最长等待 5 分钟。在返回的数据结构 [EnvInfo](https://cloud.tencent.com/document/api/876/34822#EnvInfo) 中，检查 `LogServices` 字段下 `LogServiceInfo` 是否包含有效的日志主题（Topic）等相关信息，以此判断日志资源是否已成功开通：
//
// 
//
// - **已开通**：`LogServiceInfo` 中存在日志主题 ID 等有效信息
//
// - **未开通 / 开通中**：`LogServiceInfo` 为空或相关字段缺失
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_LOGEXIST = "ResourceInUse.LogExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateEnvResourceWithContext(ctx context.Context, request *CreateEnvResourceRequest) (response *CreateEnvResourceResponse, err error) {
    if request == nil {
        request = NewCreateEnvResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateEnvResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHTTPServiceRouteRequest() (request *CreateHTTPServiceRouteRequest) {
    request = &CreateHTTPServiceRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateHTTPServiceRoute")
    
    
    return
}

func NewCreateHTTPServiceRouteResponse() (response *CreateHTTPServiceRouteResponse) {
    response = &CreateHTTPServiceRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHTTPServiceRoute
// 本接口CreateHTTPServiceRoute用于创建HTTP访问服务路由。如果不传Domain.Routes，仅创建域名信息。首次创建域名后需要调用DescribeHTTPServiceRoute查询域名状态，如果状态是PROCESSING，需要轮询查询域名状态直到SUCCESS或者FAIL。如果状态是FAIL，可以删除后重新创建。创建成功后域名可能无法访问，原因是异步下发的路由，可通过http或者https探测路由是否下发，如果http访问返回404或者https访问握手失败，可等待一会再试，直到访问正常。此外HTTP访问服务提供了默认域名，通过DescribeHTTPServiceRoute接口可直接获取默认域名。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTVERIFYFAILED = "InvalidParameter.CertVerifyFailed"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  INVALIDPARAMETER_HTTPSERVICEDOMAINNOTICP = "InvalidParameter.HTTPServiceDomainNotICP"
//  INVALIDPARAMETER_HTTPSERVICEDOMAINVERIFYFAILED = "InvalidParameter.HTTPServiceDomainVerifyFailed"
//  LIMITEXCEEDED_HTTPSERVICEDOMAIN = "LimitExceeded.HTTPServiceDomain"
//  LIMITEXCEEDED_HTTPSERVICEROUTE = "LimitExceeded.HTTPServiceRoute"
//  OPERATIONDENIED_HTTPSERVICEDOMAININBLACKLIST = "OperationDenied.HTTPServiceDomainInBlacklist"
//  OPERATIONDENIED_NONINTERNALACCOUNT = "OperationDenied.NonInternalAccount"
//  RESOURCEINUSE_HTTPSERVICEDOMAIN = "ResourceInUse.HTTPServiceDomain"
//  RESOURCEINUSE_HTTPSERVICEROUTE = "ResourceInUse.HTTPServiceRoute"
//  RESOURCENOTFOUND_HTTPSERVICEDOMAIN = "ResourceNotFound.HTTPServiceDomain"
func (c *Client) CreateHTTPServiceRoute(request *CreateHTTPServiceRouteRequest) (response *CreateHTTPServiceRouteResponse, err error) {
    return c.CreateHTTPServiceRouteWithContext(context.Background(), request)
}

// CreateHTTPServiceRoute
// 本接口CreateHTTPServiceRoute用于创建HTTP访问服务路由。如果不传Domain.Routes，仅创建域名信息。首次创建域名后需要调用DescribeHTTPServiceRoute查询域名状态，如果状态是PROCESSING，需要轮询查询域名状态直到SUCCESS或者FAIL。如果状态是FAIL，可以删除后重新创建。创建成功后域名可能无法访问，原因是异步下发的路由，可通过http或者https探测路由是否下发，如果http访问返回404或者https访问握手失败，可等待一会再试，直到访问正常。此外HTTP访问服务提供了默认域名，通过DescribeHTTPServiceRoute接口可直接获取默认域名。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTVERIFYFAILED = "InvalidParameter.CertVerifyFailed"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  INVALIDPARAMETER_HTTPSERVICEDOMAINNOTICP = "InvalidParameter.HTTPServiceDomainNotICP"
//  INVALIDPARAMETER_HTTPSERVICEDOMAINVERIFYFAILED = "InvalidParameter.HTTPServiceDomainVerifyFailed"
//  LIMITEXCEEDED_HTTPSERVICEDOMAIN = "LimitExceeded.HTTPServiceDomain"
//  LIMITEXCEEDED_HTTPSERVICEROUTE = "LimitExceeded.HTTPServiceRoute"
//  OPERATIONDENIED_HTTPSERVICEDOMAININBLACKLIST = "OperationDenied.HTTPServiceDomainInBlacklist"
//  OPERATIONDENIED_NONINTERNALACCOUNT = "OperationDenied.NonInternalAccount"
//  RESOURCEINUSE_HTTPSERVICEDOMAIN = "ResourceInUse.HTTPServiceDomain"
//  RESOURCEINUSE_HTTPSERVICEROUTE = "ResourceInUse.HTTPServiceRoute"
//  RESOURCENOTFOUND_HTTPSERVICEDOMAIN = "ResourceNotFound.HTTPServiceDomain"
func (c *Client) CreateHTTPServiceRouteWithContext(ctx context.Context, request *CreateHTTPServiceRouteRequest) (response *CreateHTTPServiceRouteResponse, err error) {
    if request == nil {
        request = NewCreateHTTPServiceRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateHTTPServiceRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHTTPServiceRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHTTPServiceRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostingDomainRequest() (request *CreateHostingDomainRequest) {
    request = &CreateHostingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateHostingDomain")
    
    
    return
}

func NewCreateHostingDomainResponse() (response *CreateHostingDomainResponse) {
    response = &CreateHostingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomain(request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    return c.CreateHostingDomainWithContext(context.Background(), request)
}

// CreateHostingDomain
// 创建托管域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHostingDomainWithContext(ctx context.Context, request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    if request == nil {
        request = NewCreateHostingDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateHostingDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHostingDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMySQLRequest() (request *CreateMySQLRequest) {
    request = &CreateMySQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateMySQL")
    
    
    return
}

func NewCreateMySQLResponse() (response *CreateMySQLResponse) {
    response = &CreateMySQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMySQL
// 本接口（CreateMySQL）用于开通Mysql型数据库。
//
// 
//
// 开通后，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询开通结果，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令。
//
// 可能返回的错误码:
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
func (c *Client) CreateMySQL(request *CreateMySQLRequest) (response *CreateMySQLResponse, err error) {
    return c.CreateMySQLWithContext(context.Background(), request)
}

// CreateMySQL
// 本接口（CreateMySQL）用于开通Mysql型数据库。
//
// 
//
// 开通后，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询开通结果，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 sql 命令。
//
// 可能返回的错误码:
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
func (c *Client) CreateMySQLWithContext(ctx context.Context, request *CreateMySQLRequest) (response *CreateMySQLResponse, err error) {
    if request == nil {
        request = NewCreateMySQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateMySQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMySQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMySQLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaticStoreRequest() (request *CreateStaticStoreRequest) {
    request = &CreateStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateStaticStore")
    
    
    return
}

func NewCreateStaticStoreResponse() (response *CreateStaticStoreResponse) {
    response = &CreateStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStore(request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    return c.CreateStaticStoreWithContext(context.Background(), request)
}

// CreateStaticStore
// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateStaticStoreWithContext(ctx context.Context, request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    if request == nil {
        request = NewCreateStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableRequest() (request *CreateTableRequest) {
    request = &CreateTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateTable")
    
    
    return
}

func NewCreateTableResponse() (response *CreateTableResponse) {
    response = &CreateTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTable
// 本接口(CreateTable)用于创建文档型数据库表，支持创建capped类型集合，暂时不支持分片表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETABLE = "FailedOperation.CreateTable"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFTABLEQUOTA = "LimitExceeded.OutOfTableQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEEXIST = "ResourceUnavailable.ResourceExist"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTable(request *CreateTableRequest) (response *CreateTableResponse, err error) {
    return c.CreateTableWithContext(context.Background(), request)
}

// CreateTable
// 本接口(CreateTable)用于创建文档型数据库表，支持创建capped类型集合，暂时不支持分片表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETABLE = "FailedOperation.CreateTable"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFTABLEQUOTA = "LimitExceeded.OutOfTableQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEEXIST = "ResourceUnavailable.ResourceExist"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTableWithContext(ctx context.Context, request *CreateTableRequest) (response *CreateTableResponse, err error) {
    if request == nil {
        request = NewCreateTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 创建tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 创建tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVmInstanceRequest() (request *CreateVmInstanceRequest) {
    request = &CreateVmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "CreateVmInstance")
    
    
    return
}

func NewCreateVmInstanceResponse() (response *CreateVmInstanceResponse) {
    response = &CreateVmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVmInstance
// 创建虚拟服务器
//
// 创建流程为先调用[DescribeVmSpec](https://cloud.tencent.com/document/product/876/129360)获取可购买的规格，同时调用[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)拉取镜像列表，选中一个规格和一个镜像后，调用[InquireVmPrice](https://cloud.tencent.com/document/product/876/129759)询价，如果价格可接受，调用此接口创建实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVmInstance(request *CreateVmInstanceRequest) (response *CreateVmInstanceResponse, err error) {
    return c.CreateVmInstanceWithContext(context.Background(), request)
}

// CreateVmInstance
// 创建虚拟服务器
//
// 创建流程为先调用[DescribeVmSpec](https://cloud.tencent.com/document/product/876/129360)获取可购买的规格，同时调用[DescribeBlueprints](https://cloud.tencent.com/document/product/1207/47689)拉取镜像列表，选中一个规格和一个镜像后，调用[InquireVmPrice](https://cloud.tencent.com/document/product/876/129759)询价，如果价格可接受，调用此接口创建实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVmInstanceWithContext(ctx context.Context, request *CreateVmInstanceRequest) (response *CreateVmInstanceResponse, err error) {
    if request == nil {
        request = NewCreateVmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "CreateVmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAIModelRequest() (request *DeleteAIModelRequest) {
    request = &DeleteAIModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteAIModel")
    
    
    return
}

func NewDeleteAIModelResponse() (response *DeleteAIModelResponse) {
    response = &DeleteAIModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAIModel
// 删除 AI 模型配置分组，支持批量删除。内置分组无法删除。分组删除后，该分组下的所有模型配置将同步移除，针对该分组模型的请求将会失败，请在删除前确认业务侧已停止对该分组的调用。
//
// 
//
// 注意：
//
// 此操作不可逆，删除后数据无法恢复，请谨慎操作。
//
// 
//
// 删除前建议通过 [DescribeAIModelList]() 接口确认分组当前状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteAIModel(request *DeleteAIModelRequest) (response *DeleteAIModelResponse, err error) {
    return c.DeleteAIModelWithContext(context.Background(), request)
}

// DeleteAIModel
// 删除 AI 模型配置分组，支持批量删除。内置分组无法删除。分组删除后，该分组下的所有模型配置将同步移除，针对该分组模型的请求将会失败，请在删除前确认业务侧已停止对该分组的调用。
//
// 
//
// 注意：
//
// 此操作不可逆，删除后数据无法恢复，请谨慎操作。
//
// 
//
// 删除前建议通过 [DescribeAIModelList]() 接口确认分组当前状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteAIModelWithContext(ctx context.Context, request *DeleteAIModelRequest) (response *DeleteAIModelResponse, err error) {
    if request == nil {
        request = NewDeleteAIModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteAIModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAIModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAIModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteApiKey")
    
    
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiKey
// 删除指定云开发环境下的某个 API Key 服务端访问凭证。删除后，该 API Key 对应的 Token 将被吊销，已使用该 Key 发起的请求将失败。该操作具有幂等性，若指定的 API Key 不存在则直接返回成功。需要管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    return c.DeleteApiKeyWithContext(context.Background(), request)
}

// DeleteApiKey
// 删除指定云开发环境下的某个 API Key 服务端访问凭证。删除后，该 API Key 对应的 Token 将被吊销，已使用该 Key 发起的请求将失败。该操作具有幂等性，若指定的 API Key 不存在则直接返回成功。需要管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthDomainRequest() (request *DeleteAuthDomainRequest) {
    request = &DeleteAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteAuthDomain")
    
    
    return
}

func NewDeleteAuthDomainResponse() (response *DeleteAuthDomainResponse) {
    response = &DeleteAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthDomain
// 删除合法域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名，将对应安全域名的id填入Domainlds中
//
// 
//
// 注意⚠️
//
// 安全域名被删除之后，可能会引起跨域问题，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuthDomain(request *DeleteAuthDomainRequest) (response *DeleteAuthDomainResponse, err error) {
    return c.DeleteAuthDomainWithContext(context.Background(), request)
}

// DeleteAuthDomain
// 删除合法域名。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [DescribeAuthDomains](https://cloud.tencent.com/document/product/876/42151) 获取当前已绑定生效的安全域名，将对应安全域名的id填入Domainlds中
//
// 
//
// 注意⚠️
//
// 安全域名被删除之后，可能会引起跨域问题，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  OPERATIONDENIED_RESOURCEFROZEN = "OperationDenied.ResourceFrozen"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuthDomainWithContext(ctx context.Context, request *DeleteAuthDomainRequest) (response *DeleteAuthDomainResponse, err error) {
    if request == nil {
        request = NewDeleteAuthDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteAuthDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHTTPServiceRouteRequest() (request *DeleteHTTPServiceRouteRequest) {
    request = &DeleteHTTPServiceRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteHTTPServiceRoute")
    
    
    return
}

func NewDeleteHTTPServiceRouteResponse() (response *DeleteHTTPServiceRouteResponse) {
    response = &DeleteHTTPServiceRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHTTPServiceRoute
// 本接口DeleteHTTPServiceRoute用于删除HTTP访问服务域名或者路由。可批量删除多条path路由、删除域名及所有path路由，如果Paths字段为空则删除域名及所有path路由，如果Paths不为空则仅删除path路由。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteHTTPServiceRoute(request *DeleteHTTPServiceRouteRequest) (response *DeleteHTTPServiceRouteResponse, err error) {
    return c.DeleteHTTPServiceRouteWithContext(context.Background(), request)
}

// DeleteHTTPServiceRoute
// 本接口DeleteHTTPServiceRoute用于删除HTTP访问服务域名或者路由。可批量删除多条path路由、删除域名及所有path路由，如果Paths字段为空则删除域名及所有path路由，如果Paths不为空则仅删除path路由。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteHTTPServiceRouteWithContext(ctx context.Context, request *DeleteHTTPServiceRouteRequest) (response *DeleteHTTPServiceRouteResponse, err error) {
    if request == nil {
        request = NewDeleteHTTPServiceRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteHTTPServiceRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHTTPServiceRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHTTPServiceRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProviderRequest() (request *DeleteProviderRequest) {
    request = &DeleteProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteProvider")
    
    
    return
}

func NewDeleteProviderResponse() (response *DeleteProviderResponse) {
    response = &DeleteProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProvider
// 删除认证源
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteProvider(request *DeleteProviderRequest) (response *DeleteProviderResponse, err error) {
    return c.DeleteProviderWithContext(context.Background(), request)
}

// DeleteProvider
// 删除认证源
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DeleteProviderWithContext(ctx context.Context, request *DeleteProviderRequest) (response *DeleteProviderResponse, err error) {
    if request == nil {
        request = NewDeleteProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableRequest() (request *DeleteTableRequest) {
    request = &DeleteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteTable")
    
    
    return
}

func NewDeleteTableResponse() (response *DeleteTableResponse) {
    response = &DeleteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTable
// 本接口(DeleteTable)用于删除文档型数据库表，删除表后表中数据将会被删除且无法恢复，请谨慎操作。
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTable(request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    return c.DeleteTableWithContext(context.Background(), request)
}

// DeleteTable
// 本接口(DeleteTable)用于删除文档型数据库表，删除表后表中数据将会被删除且无法恢复，请谨慎操作。
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTableWithContext(ctx context.Context, request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    if request == nil {
        request = NewDeleteTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// 删除tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 删除tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVmInstanceRequest() (request *DeleteVmInstanceRequest) {
    request = &DeleteVmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteVmInstance")
    
    
    return
}

func NewDeleteVmInstanceResponse() (response *DeleteVmInstanceResponse) {
    response = &DeleteVmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVmInstance
// 销毁云服务器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVmInstance(request *DeleteVmInstanceRequest) (response *DeleteVmInstanceResponse, err error) {
    return c.DeleteVmInstanceWithContext(context.Background(), request)
}

// DeleteVmInstance
// 销毁云服务器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteVmInstanceWithContext(ctx context.Context, request *DeleteVmInstanceRequest) (response *DeleteVmInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteVmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DeleteVmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIModelsRequest() (request *DescribeAIModelsRequest) {
    request = &DescribeAIModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeAIModels")
    
    
    return
}

func NewDescribeAIModelsResponse() (response *DescribeAIModelsResponse) {
    response = &DescribeAIModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIModels
// 查询指定云开发环境下已配置的 AI 模型分组列表。返回结果包含该环境下所有类型的模型分组（自定义类型 custom、内置类型 builtin），以及各分组下的模型列表、服务地址、启用状态等配置信息。
//
// 
//
// 通常在以下场景中使用：
//
// 
//
// 控制台展示：在云开发控制台 AI+ 功能 → 模型管理页面，展示当前环境已配置的模型分组列表。
//
// 
//
// 模型管理：在调用 [UpdateAIModel]() 或 [DeleteAIModel]() 接口前，通过本接口查询当前分组配置。
//
// 
//
// 业务集成：开发者查询可用模型列表，用于选择合适的模型接入 AI+ 功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeAIModels(request *DescribeAIModelsRequest) (response *DescribeAIModelsResponse, err error) {
    return c.DescribeAIModelsWithContext(context.Background(), request)
}

// DescribeAIModels
// 查询指定云开发环境下已配置的 AI 模型分组列表。返回结果包含该环境下所有类型的模型分组（自定义类型 custom、内置类型 builtin），以及各分组下的模型列表、服务地址、启用状态等配置信息。
//
// 
//
// 通常在以下场景中使用：
//
// 
//
// 控制台展示：在云开发控制台 AI+ 功能 → 模型管理页面，展示当前环境已配置的模型分组列表。
//
// 
//
// 模型管理：在调用 [UpdateAIModel]() 或 [DeleteAIModel]() 接口前，通过本接口查询当前分组配置。
//
// 
//
// 业务集成：开发者查询可用模型列表，用于选择合适的模型接入 AI+ 功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeAIModelsWithContext(ctx context.Context, request *DescribeAIModelsRequest) (response *DescribeAIModelsResponse, err error) {
    if request == nil {
        request = NewDescribeAIModelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeAIModels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIModelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyListRequest() (request *DescribeApiKeyListRequest) {
    request = &DescribeApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeApiKeyList")
    
    
    return
}

func NewDescribeApiKeyListResponse() (response *DescribeApiKeyListResponse) {
    response = &DescribeApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKeyList
// 查询 API Key 列表。分页查询指定云开发环境下的 API Key 访问凭证列表。支持按类型过滤（api_key 或 publish_key）。未指定类型时，默认仅返回 api_key 类型的记录。列表查询中 api_key 类型的令牌值将进行脱敏处理（仅保留前后各 6 位字符）；publish_key 类型始终返回完整明文。接口需要管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeApiKeyList(request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    return c.DescribeApiKeyListWithContext(context.Background(), request)
}

// DescribeApiKeyList
// 查询 API Key 列表。分页查询指定云开发环境下的 API Key 访问凭证列表。支持按类型过滤（api_key 或 publish_key）。未指定类型时，默认仅返回 api_key 类型的记录。列表查询中 api_key 类型的令牌值将进行脱敏处理（仅保留前后各 6 位字符）；publish_key 类型始终返回完整明文。接口需要管理员权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeApiKeyListWithContext(ctx context.Context, request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthDomainsRequest() (request *DescribeAuthDomainsRequest) {
    request = &DescribeAuthDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeAuthDomains")
    
    
    return
}

func NewDescribeAuthDomainsResponse() (response *DescribeAuthDomainsResponse) {
    response = &DescribeAuthDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthDomains
// 本接口用于获取当前环境的安全域名列表。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [CreateAuthDomain](https://cloud.tencent.com/document/product/876/42764) 增加安全域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomains(request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    return c.DescribeAuthDomainsWithContext(context.Background(), request)
}

// DescribeAuthDomains
// 本接口用于获取当前环境的安全域名列表。
//
// 云开发会校验网页应用请求的来源域名，您需要将来源域名加入到WEB安全域名列表中。
//
// 可以通过接口 [CreateAuthDomain](https://cloud.tencent.com/document/product/876/42764) 增加安全域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeAuthDomainsWithContext(ctx context.Context, request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeAuthDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaasPackageListRequest() (request *DescribeBaasPackageListRequest) {
    request = &DescribeBaasPackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeBaasPackageList")
    
    
    return
}

func NewDescribeBaasPackageListResponse() (response *DescribeBaasPackageListResponse) {
    response = &DescribeBaasPackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaasPackageList
// 获取新套餐列表，含详情，如果传了PackageId，则只获取指定套餐详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeBaasPackageList(request *DescribeBaasPackageListRequest) (response *DescribeBaasPackageListResponse, err error) {
    return c.DescribeBaasPackageListWithContext(context.Background(), request)
}

// DescribeBaasPackageList
// 获取新套餐列表，含详情，如果传了PackageId，则只获取指定套餐详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeBaasPackageListWithContext(ctx context.Context, request *DescribeBaasPackageListRequest) (response *DescribeBaasPackageListResponse, err error) {
    if request == nil {
        request = NewDescribeBaasPackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeBaasPackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaasPackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaasPackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingInfoRequest() (request *DescribeBillingInfoRequest) {
    request = &DescribeBillingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeBillingInfo")
    
    
    return
}

func NewDescribeBillingInfoResponse() (response *DescribeBillingInfoResponse) {
    response = &DescribeBillingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingInfo
// 获取计费相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBillingInfo(request *DescribeBillingInfoRequest) (response *DescribeBillingInfoResponse, err error) {
    return c.DescribeBillingInfoWithContext(context.Background(), request)
}

// DescribeBillingInfo
// 获取计费相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBillingInfoWithContext(ctx context.Context, request *DescribeBillingInfoRequest) (response *DescribeBillingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBillingInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeBillingInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientRequest() (request *DescribeClientRequest) {
    request = &DescribeClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeClient")
    
    
    return
}

func NewDescribeClientResponse() (response *DescribeClientResponse) {
    response = &DescribeClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClient
// 查询客户端详情。获取指定云开发环境下某个客户端的配置信息，包括客户端基本信息（名称、图标、描述）、OAuth 凭证（ClientId、ClientSecret）、安全域名、允许的 Scope 列表、Token 有效期、会话控制策略等。当客户端 ID 等于环境 ID 时，返回该环境的默认客户端配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClient(request *DescribeClientRequest) (response *DescribeClientResponse, err error) {
    return c.DescribeClientWithContext(context.Background(), request)
}

// DescribeClient
// 查询客户端详情。获取指定云开发环境下某个客户端的配置信息，包括客户端基本信息（名称、图标、描述）、OAuth 凭证（ClientId、ClientSecret）、安全域名、允许的 Scope 列表、Token 有效期、会话控制策略等。当客户端 ID 等于环境 ID 时，返回该环境的默认客户端配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClientWithContext(ctx context.Context, request *DescribeClientRequest) (response *DescribeClientResponse, err error) {
    if request == nil {
        request = NewDescribeClientRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeClient")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseBuildServiceRequest() (request *DescribeCloudBaseBuildServiceRequest) {
    request = &DescribeCloudBaseBuildServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseBuildService")
    
    
    return
}

func NewDescribeCloudBaseBuildServiceResponse() (response *DescribeCloudBaseBuildServiceResponse) {
    response = &DescribeCloudBaseBuildServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseBuildService(request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    return c.DescribeCloudBaseBuildServiceWithContext(context.Background(), request)
}

// DescribeCloudBaseBuildService
// 获取云托管代码上传url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseBuildServiceWithContext(ctx context.Context, request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseBuildServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCloudBaseBuildService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudBaseBuildService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseBuildServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunServerVersionRequest() (request *DescribeCloudBaseRunServerVersionRequest) {
    request = &DescribeCloudBaseRunServerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunServerVersion")
    
    
    return
}

func NewDescribeCloudBaseRunServerVersionResponse() (response *DescribeCloudBaseRunServerVersionResponse) {
    response = &DescribeCloudBaseRunServerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudBaseRunServerVersion
// 查询服务版本的详情，CPU和MEM  请使用CPUSize和MemSize
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseRunServerVersion(request *DescribeCloudBaseRunServerVersionRequest) (response *DescribeCloudBaseRunServerVersionResponse, err error) {
    return c.DescribeCloudBaseRunServerVersionWithContext(context.Background(), request)
}

// DescribeCloudBaseRunServerVersion
// 查询服务版本的详情，CPU和MEM  请使用CPUSize和MemSize
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudBaseRunServerVersionWithContext(ctx context.Context, request *DescribeCloudBaseRunServerVersionRequest) (response *DescribeCloudBaseRunServerVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunServerVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCloudBaseRunServerVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudBaseRunServerVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudBaseRunServerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreateMySQLResultRequest() (request *DescribeCreateMySQLResultRequest) {
    request = &DescribeCreateMySQLResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCreateMySQLResult")
    
    
    return
}

func NewDescribeCreateMySQLResultResponse() (response *DescribeCreateMySQLResultResponse) {
    response = &DescribeCreateMySQLResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCreateMySQLResult
// 本接口（DescribeCreateMySQLResult）用于查询开通Mysql结果。
//
// 
//
// `Response.Data.Status = "notexist"` 表示未开通，如果未开通，可以调用 [CreateMySQL](https://cloud.tencent.com/document/api/876/128186) 来开通
//
//  `Response.Data. Status = "success"` 表示开通成功，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 MySql 命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCreateMySQLResult(request *DescribeCreateMySQLResultRequest) (response *DescribeCreateMySQLResultResponse, err error) {
    return c.DescribeCreateMySQLResultWithContext(context.Background(), request)
}

// DescribeCreateMySQLResult
// 本接口（DescribeCreateMySQLResult）用于查询开通Mysql结果。
//
// 
//
// `Response.Data.Status = "notexist"` 表示未开通，如果未开通，可以调用 [CreateMySQL](https://cloud.tencent.com/document/api/876/128186) 来开通
//
//  `Response.Data. Status = "success"` 表示开通成功，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 sql 命令，比如创建表格、插入数据、删除表格等 MySql 命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SERVICENOTEXIST = "InvalidParameter.ServiceNotExist"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCreateMySQLResultWithContext(ctx context.Context, request *DescribeCreateMySQLResultRequest) (response *DescribeCreateMySQLResultResponse, err error) {
    if request == nil {
        request = NewDescribeCreateMySQLResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCreateMySQLResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCreateMySQLResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCreateMySQLResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurveDataRequest() (request *DescribeCurveDataRequest) {
    request = &DescribeCurveDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCurveData")
    
    
    return
}

func NewDescribeCurveDataResponse() (response *DescribeCurveDataResponse) {
    response = &DescribeCurveDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCurveData
// 根据指定指标名称，查询某环境在指定时间范围内的监控数据，返回按统计粒度聚合后的时序数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCurveData(request *DescribeCurveDataRequest) (response *DescribeCurveDataResponse, err error) {
    return c.DescribeCurveDataWithContext(context.Background(), request)
}

// DescribeCurveData
// 根据指定指标名称，查询某环境在指定时间范围内的监控数据，返回按统计粒度聚合后的时序数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCurveDataWithContext(ctx context.Context, request *DescribeCurveDataRequest) (response *DescribeCurveDataResponse, err error) {
    if request == nil {
        request = NewDescribeCurveDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeCurveData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurveData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurveDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseACLRequest() (request *DescribeDatabaseACLRequest) {
    request = &DescribeDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDatabaseACL")
    
    
    return
}

func NewDescribeDatabaseACLResponse() (response *DescribeDatabaseACLResponse) {
    response = &DescribeDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseACL
// 本接口（DescribeDatabaseACL）获取文档型数据库权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACL(request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    return c.DescribeDatabaseACLWithContext(context.Background(), request)
}

// DescribeDatabaseACL
// 本接口（DescribeDatabaseACL）获取文档型数据库权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeDatabaseACLWithContext(ctx context.Context, request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeDatabaseACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvAccountCircleRequest() (request *DescribeEnvAccountCircleRequest) {
    request = &DescribeEnvAccountCircleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvAccountCircle")
    
    
    return
}

func NewDescribeEnvAccountCircleResponse() (response *DescribeEnvAccountCircleResponse) {
    response = &DescribeEnvAccountCircleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvAccountCircle
// 查询环境计费周期。
//
// 云开发环境的资源点都是按月结算的，每个月都有一定的抵扣额度。
//
// 
//
// 例如：
//
//   某个环境在 2026-01-05 购买了3个月个人版(到期时间: 2026-04-05)，则他可以在以下3个周期内，分别享有40000资源点的额度：
//
//   1. 2026-01-05 ~ 2026-02-05 23:59:59
//
//   2. 2026-02-06 ~ 2026-03-05 23:59:59
//
//   3. 2026-03-06 ~ 2026-04-05 23:59:59
//
// 
//
// 本接口，用于获取环境当前属于哪个计费周期内。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE_RESOURCEEXPIRED = "ResourceUnavailable.ResourceExpired"
func (c *Client) DescribeEnvAccountCircle(request *DescribeEnvAccountCircleRequest) (response *DescribeEnvAccountCircleResponse, err error) {
    return c.DescribeEnvAccountCircleWithContext(context.Background(), request)
}

// DescribeEnvAccountCircle
// 查询环境计费周期。
//
// 云开发环境的资源点都是按月结算的，每个月都有一定的抵扣额度。
//
// 
//
// 例如：
//
//   某个环境在 2026-01-05 购买了3个月个人版(到期时间: 2026-04-05)，则他可以在以下3个周期内，分别享有40000资源点的额度：
//
//   1. 2026-01-05 ~ 2026-02-05 23:59:59
//
//   2. 2026-02-06 ~ 2026-03-05 23:59:59
//
//   3. 2026-03-06 ~ 2026-04-05 23:59:59
//
// 
//
// 本接口，用于获取环境当前属于哪个计费周期内。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE_RESOURCEEXPIRED = "ResourceUnavailable.ResourceExpired"
func (c *Client) DescribeEnvAccountCircleWithContext(ctx context.Context, request *DescribeEnvAccountCircleRequest) (response *DescribeEnvAccountCircleResponse, err error) {
    if request == nil {
        request = NewDescribeEnvAccountCircleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvAccountCircle")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvAccountCircle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvAccountCircleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvLimitRequest() (request *DescribeEnvLimitRequest) {
    request = &DescribeEnvLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvLimit")
    
    
    return
}

func NewDescribeEnvLimitResponse() (response *DescribeEnvLimitResponse) {
    response = &DescribeEnvLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimit(request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    return c.DescribeEnvLimitWithContext(context.Background(), request)
}

// DescribeEnvLimit
// 查询环境个数上限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeEnvLimitWithContext(ctx context.Context, request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    if request == nil {
        request = NewDescribeEnvLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvsRequest() (request *DescribeEnvsRequest) {
    request = &DescribeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvs")
    
    
    return
}

func NewDescribeEnvsResponse() (response *DescribeEnvsResponse) {
    response = &DescribeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvs(request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    return c.DescribeEnvsWithContext(context.Background(), request)
}

// DescribeEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTION = "InvalidParameter.Action"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) DescribeEnvsWithContext(ctx context.Context, request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayVersionsRequest() (request *DescribeGatewayVersionsRequest) {
    request = &DescribeGatewayVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeGatewayVersions")
    
    
    return
}

func NewDescribeGatewayVersionsResponse() (response *DescribeGatewayVersionsResponse) {
    response = &DescribeGatewayVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayVersions
// 查询网关版本信息
//
// 暂不鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeGatewayVersions(request *DescribeGatewayVersionsRequest) (response *DescribeGatewayVersionsResponse, err error) {
    return c.DescribeGatewayVersionsWithContext(context.Background(), request)
}

// DescribeGatewayVersions
// 查询网关版本信息
//
// 暂不鉴权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeGatewayVersionsWithContext(ctx context.Context, request *DescribeGatewayVersionsRequest) (response *DescribeGatewayVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeGatewayVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHTTPServiceRouteRequest() (request *DescribeHTTPServiceRouteRequest) {
    request = &DescribeHTTPServiceRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeHTTPServiceRoute")
    
    
    return
}

func NewDescribeHTTPServiceRouteResponse() (response *DescribeHTTPServiceRouteResponse) {
    response = &DescribeHTTPServiceRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHTTPServiceRoute
// 本接口DescribeHTTPServiceRoute用于查询环境下HTTP访问服务路由信息。可通过Filters过滤。如果不存在不会返回错误。HTTP访问服务提供了默认域名，通过本接口可直接获取默认域名。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeHTTPServiceRoute(request *DescribeHTTPServiceRouteRequest) (response *DescribeHTTPServiceRouteResponse, err error) {
    return c.DescribeHTTPServiceRouteWithContext(context.Background(), request)
}

// DescribeHTTPServiceRoute
// 本接口DescribeHTTPServiceRoute用于查询环境下HTTP访问服务路由信息。可通过Filters过滤。如果不存在不会返回错误。HTTP访问服务提供了默认域名，通过本接口可直接获取默认域名。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeHTTPServiceRouteWithContext(ctx context.Context, request *DescribeHTTPServiceRouteRequest) (response *DescribeHTTPServiceRouteResponse, err error) {
    if request == nil {
        request = NewDescribeHTTPServiceRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeHTTPServiceRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHTTPServiceRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHTTPServiceRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostingDomainTaskRequest() (request *DescribeHostingDomainTaskRequest) {
    request = &DescribeHostingDomainTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeHostingDomainTask")
    
    
    return
}

func NewDescribeHostingDomainTaskResponse() (response *DescribeHostingDomainTaskResponse) {
    response = &DescribeHostingDomainTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTask(request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    return c.DescribeHostingDomainTaskWithContext(context.Background(), request)
}

// DescribeHostingDomainTask
// 查询静态托管域名任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeHostingDomainTaskWithContext(ctx context.Context, request *DescribeHostingDomainTaskRequest) (response *DescribeHostingDomainTaskResponse, err error) {
    if request == nil {
        request = NewDescribeHostingDomainTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeHostingDomainTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostingDomainTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostingDomainTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginConfigRequest() (request *DescribeLoginConfigRequest) {
    request = &DescribeLoginConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeLoginConfig")
    
    
    return
}

func NewDescribeLoginConfigResponse() (response *DescribeLoginConfigResponse) {
    response = &DescribeLoginConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoginConfig
// 查询指定云开发环境的登录策略配置。包括手机号短信登录、邮箱登录、用户名密码登录和匿名登录方式的开启状态，同时包含短信验证码发送通道、MFA 多因子认证和密码的更新策略。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoginConfig(request *DescribeLoginConfigRequest) (response *DescribeLoginConfigResponse, err error) {
    return c.DescribeLoginConfigWithContext(context.Background(), request)
}

// DescribeLoginConfig
// 查询指定云开发环境的登录策略配置。包括手机号短信登录、邮箱登录、用户名密码登录和匿名登录方式的开启状态，同时包含短信验证码发送通道、MFA 多因子认证和密码的更新策略。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoginConfigWithContext(ctx context.Context, request *DescribeLoginConfigRequest) (response *DescribeLoginConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLoginConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeLoginConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagedAIModelListRequest() (request *DescribeManagedAIModelListRequest) {
    request = &DescribeManagedAIModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeManagedAIModelList")
    
    
    return
}

func NewDescribeManagedAIModelListResponse() (response *DescribeManagedAIModelListResponse) {
    response = &DescribeManagedAIModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeManagedAIModelList
// 查询云开发平台支持的托管类型 AI 模型列表。
//
// 
//
// 托管类型模型由云开发平台统一接入和管理，用户无需自行配置模型服务地址和访问密钥，开通后即可直接使用。返回结果按模型分组（Group）组织，包含各模型的规格参数（ModelSpec）和计费信息（ModelChargingInfo）。
//
// 
//
// 通常在以下场景中使用：
//
// 
//
// 开通托管模型前：通过本接口查询平台支持的托管模型及其规格，结合 [UpdateAIModel](https://cloud.tencent.com/document/product/876/131316) 接口完成模型配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeManagedAIModelList(request *DescribeManagedAIModelListRequest) (response *DescribeManagedAIModelListResponse, err error) {
    return c.DescribeManagedAIModelListWithContext(context.Background(), request)
}

// DescribeManagedAIModelList
// 查询云开发平台支持的托管类型 AI 模型列表。
//
// 
//
// 托管类型模型由云开发平台统一接入和管理，用户无需自行配置模型服务地址和访问密钥，开通后即可直接使用。返回结果按模型分组（Group）组织，包含各模型的规格参数（ModelSpec）和计费信息（ModelChargingInfo）。
//
// 
//
// 通常在以下场景中使用：
//
// 
//
// 开通托管模型前：通过本接口查询平台支持的托管模型及其规格，结合 [UpdateAIModel](https://cloud.tencent.com/document/product/876/131316) 接口完成模型配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) DescribeManagedAIModelListWithContext(ctx context.Context, request *DescribeManagedAIModelListRequest) (response *DescribeManagedAIModelListResponse, err error) {
    if request == nil {
        request = NewDescribeManagedAIModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeManagedAIModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManagedAIModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagedAIModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMySQLClusterDetailRequest() (request *DescribeMySQLClusterDetailRequest) {
    request = &DescribeMySQLClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeMySQLClusterDetail")
    
    
    return
}

func NewDescribeMySQLClusterDetailResponse() (response *DescribeMySQLClusterDetailResponse) {
    response = &DescribeMySQLClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMySQLClusterDetail
// 本接口（DescribeMySQLClusterDetail）查询Mysql集群信息。
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有已开通的才能查到集群信息，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 MySql 命令，比如创建表格、插入数据、删除表格等 MySql 命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLClusterDetail(request *DescribeMySQLClusterDetailRequest) (response *DescribeMySQLClusterDetailResponse, err error) {
    return c.DescribeMySQLClusterDetailWithContext(context.Background(), request)
}

// DescribeMySQLClusterDetail
// 本接口（DescribeMySQLClusterDetail）查询Mysql集群信息。
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有已开通的才能查到集群信息，Mysql开通成功后，可通过接口设置数据库账号相关功能包括但不限于【创建账号、删除账号、查询可授权权限列表、查询账号已有权限、修改主机、修改配置、修改账号库表权限】、集群操作相关【查询集群参数、修改集群参数】，连接设置相关【关闭外网、开通外网、查询集群信息】，备份回档相关【创建手动回档、删除手动回档、修改自动备份配置信息、查询备份文件列表、集群回档、查询任务列表、获取table列表、获取集群数据库列表、查询备份下载地址】，相关功能接口文档：[TDSQL-C MySQL API文档](https://cloud.tencent.com/document/product/1003/48106)，可以通过 [RunSql](https://cloud.tencent.com/document/api/876/127880) 接口来执行 MySql 命令，比如创建表格、插入数据、删除表格等 MySql 命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLClusterDetailWithContext(ctx context.Context, request *DescribeMySQLClusterDetailRequest) (response *DescribeMySQLClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMySQLClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeMySQLClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMySQLClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMySQLClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMySQLTaskStatusRequest() (request *DescribeMySQLTaskStatusRequest) {
    request = &DescribeMySQLTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeMySQLTaskStatus")
    
    
    return
}

func NewDescribeMySQLTaskStatusResponse() (response *DescribeMySQLTaskStatusResponse) {
    response = &DescribeMySQLTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMySQLTaskStatus
// 本接口（DescribeMySQLTaskStatus）用于查询Mysql任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLTaskStatus(request *DescribeMySQLTaskStatusRequest) (response *DescribeMySQLTaskStatusResponse, err error) {
    return c.DescribeMySQLTaskStatusWithContext(context.Background(), request)
}

// DescribeMySQLTaskStatus
// 本接口（DescribeMySQLTaskStatus）用于查询Mysql任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DescribeMySQLTaskStatusWithContext(ctx context.Context, request *DescribeMySQLTaskStatusRequest) (response *DescribeMySQLTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeMySQLTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeMySQLTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMySQLTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMySQLTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaDataRequest() (request *DescribeQuotaDataRequest) {
    request = &DescribeQuotaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeQuotaData")
    
    
    return
}

func NewDescribeQuotaDataResponse() (response *DescribeQuotaDataResponse) {
    response = &DescribeQuotaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaData(request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    return c.DescribeQuotaDataWithContext(context.Background(), request)
}

// DescribeQuotaData
// 查询指定指标的配额使用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeQuotaDataWithContext(ctx context.Context, request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeQuotaData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeRuleRequest() (request *DescribeSafeRuleRequest) {
    request = &DescribeSafeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeSafeRule")
    
    
    return
}

func NewDescribeSafeRuleResponse() (response *DescribeSafeRuleResponse) {
    response = &DescribeSafeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSafeRule
// 查询数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍](https://cloud.tencent.com/document/product/876/123478) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSafeRule(request *DescribeSafeRuleRequest) (response *DescribeSafeRuleResponse, err error) {
    return c.DescribeSafeRuleWithContext(context.Background(), request)
}

// DescribeSafeRule
// 查询数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍](https://cloud.tencent.com/document/product/876/123478) 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) DescribeSafeRuleWithContext(ctx context.Context, request *DescribeSafeRuleRequest) (response *DescribeSafeRuleResponse, err error) {
    if request == nil {
        request = NewDescribeSafeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeSafeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaticStoreRequest() (request *DescribeStaticStoreRequest) {
    request = &DescribeStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeStaticStore")
    
    
    return
}

func NewDescribeStaticStoreResponse() (response *DescribeStaticStoreResponse) {
    response = &DescribeStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStaticStore
// 查看当前环境下静态托管资源信息，根据返回结果判断静态资源的状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStaticStore(request *DescribeStaticStoreRequest) (response *DescribeStaticStoreResponse, err error) {
    return c.DescribeStaticStoreWithContext(context.Background(), request)
}

// DescribeStaticStore
// 查看当前环境下静态托管资源信息，根据返回结果判断静态资源的状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStaticStoreWithContext(ctx context.Context, request *DescribeStaticStoreRequest) (response *DescribeStaticStoreResponse, err error) {
    if request == nil {
        request = NewDescribeStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableRequest() (request *DescribeTableRequest) {
    request = &DescribeTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeTable")
    
    
    return
}

func NewDescribeTableResponse() (response *DescribeTableResponse) {
    response = &DescribeTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTable
// 本接口（DescribeTable）用于查询文档型数据库表的相关信息，包括索引等信息。
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTable(request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    return c.DescribeTableWithContext(context.Background(), request)
}

// DescribeTable
// 本接口（DescribeTable）用于查询文档型数据库表的相关信息，包括索引等信息。
//
// 
//
// 接口入参中的 Tag 为文档型数据库的实例 Id，可以通过 [DescribeEnvs](https://cloud.tencent.com/document/api/876/34820) 接口返回的 EnvList[0].Databases[0].InstanceId 获取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableWithContext(ctx context.Context, request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    if request == nil {
        request = NewDescribeTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTables
// 本接口(DescribeTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// 本接口(DescribeTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// 查询tcb用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// 查询tcb用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeUserList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVmInstancesRequest() (request *DescribeVmInstancesRequest) {
    request = &DescribeVmInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeVmInstances")
    
    
    return
}

func NewDescribeVmInstancesResponse() (response *DescribeVmInstancesResponse) {
    response = &DescribeVmInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVmInstances
// 查询环境下的云服务器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVmInstances(request *DescribeVmInstancesRequest) (response *DescribeVmInstancesResponse, err error) {
    return c.DescribeVmInstancesWithContext(context.Background(), request)
}

// DescribeVmInstances
// 查询环境下的云服务器列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVmInstancesWithContext(ctx context.Context, request *DescribeVmInstancesRequest) (response *DescribeVmInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeVmInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeVmInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVmInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVmInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVmSpecRequest() (request *DescribeVmSpecRequest) {
    request = &DescribeVmSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeVmSpec")
    
    
    return
}

func NewDescribeVmSpecResponse() (response *DescribeVmSpecResponse) {
    response = &DescribeVmSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVmSpec
// 云服务器规格list
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVmSpec(request *DescribeVmSpecRequest) (response *DescribeVmSpecResponse, err error) {
    return c.DescribeVmSpecWithContext(context.Background(), request)
}

// DescribeVmSpec
// 云服务器规格list
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVmSpecWithContext(ctx context.Context, request *DescribeVmSpecRequest) (response *DescribeVmSpecResponse, err error) {
    if request == nil {
        request = NewDescribeVmSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DescribeVmSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVmSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVmSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyEnvRequest() (request *DestroyEnvRequest) {
    request = &DestroyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyEnv")
    
    
    return
}

func NewDestroyEnvResponse() (response *DestroyEnvResponse) {
    response = &DestroyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyEnv
// 本接口用于销毁云开发环境。
//
// 云开发环境遵循腾讯云包年包月预付费产品生命周期，因此环境销毁需要分两步：
//
// 1. 资源退费。此时会根据当前环境剩余有效期，自动退还相关费用(代金券不退)。退款后，环境进入隔离期。
//
// 2. 环境删除。环境在进入隔离期后15天会自动删除。也可以通过本接口，指定 IsForce=true 来强制删除隔离期环境。
//
// 
//
// **注意**⚠️
//
//   1. 环境退费后进入隔离期，则所有资源均无法访问，控制台无法操作和管理。
//
//   2. 环境被彻底删除后，所有数据均无法找回。请谨慎操作。
//
// 
//
// 可以通过接口 [tcb:DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境计费状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnv(request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    return c.DestroyEnvWithContext(context.Background(), request)
}

// DestroyEnv
// 本接口用于销毁云开发环境。
//
// 云开发环境遵循腾讯云包年包月预付费产品生命周期，因此环境销毁需要分两步：
//
// 1. 资源退费。此时会根据当前环境剩余有效期，自动退还相关费用(代金券不退)。退款后，环境进入隔离期。
//
// 2. 环境删除。环境在进入隔离期后15天会自动删除。也可以通过本接口，指定 IsForce=true 来强制删除隔离期环境。
//
// 
//
// **注意**⚠️
//
//   1. 环境退费后进入隔离期，则所有资源均无法访问，控制台无法操作和管理。
//
//   2. 环境被彻底删除后，所有数据均无法找回。请谨慎操作。
//
// 
//
// 可以通过接口 [tcb:DescribeBillingInfo](https://cloud.tencent.com/document/product/876/94390) 查询环境计费状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INVOICEAMOUNTLACK = "ResourceUnavailable.InvoiceAmountLack"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyEnvWithContext(ctx context.Context, request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    if request == nil {
        request = NewDestroyEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyMySQLRequest() (request *DestroyMySQLRequest) {
    request = &DestroyMySQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyMySQL")
    
    
    return
}

func NewDestroyMySQLResponse() (response *DestroyMySQLResponse) {
    response = &DestroyMySQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyMySQL
// 本接口（DestroyMySQL）用于销毁Mysql。
//
// 
//
// 销毁后可以通过 [DescribeMySQLTaskStatus](https://cloud.tencent.com/document/api/876/128183) 接口查询销毁结果，如果 `Response.Data. Status = FAILED ` 表示销毁失败，可以重新调用销毁接口重试。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DestroyMySQL(request *DestroyMySQLRequest) (response *DestroyMySQLResponse, err error) {
    return c.DestroyMySQLWithContext(context.Background(), request)
}

// DestroyMySQL
// 本接口（DestroyMySQL）用于销毁Mysql。
//
// 
//
// 销毁后可以通过 [DescribeMySQLTaskStatus](https://cloud.tencent.com/document/api/876/128183) 接口查询销毁结果，如果 `Response.Data. Status = FAILED ` 表示销毁失败，可以重新调用销毁接口重试。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCENOTEXIST = "FailedOperation.DataSourceNotExist"
func (c *Client) DestroyMySQLWithContext(ctx context.Context, request *DestroyMySQLRequest) (response *DestroyMySQLResponse, err error) {
    if request == nil {
        request = NewDestroyMySQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyMySQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyMySQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyMySQLResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStaticStoreRequest() (request *DestroyStaticStoreRequest) {
    request = &DestroyStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyStaticStore")
    
    
    return
}

func NewDestroyStaticStoreResponse() (response *DestroyStaticStoreResponse) {
    response = &DestroyStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStore(request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    return c.DestroyStaticStoreWithContext(context.Background(), request)
}

// DestroyStaticStore
// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DestroyStaticStoreWithContext(ctx context.Context, request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    if request == nil {
        request = NewDestroyStaticStoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "DestroyStaticStore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyStaticStore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewExecutePGSqlRequest() (request *ExecutePGSqlRequest) {
    request = &ExecutePGSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ExecutePGSql")
    
    
    return
}

func NewExecutePGSqlResponse() (response *ExecutePGSqlResponse) {
    response = &ExecutePGSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecutePGSql
// 在Postgres数据库上执行SQL
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCESTATUSCONFLICT = "FailedOperation.InstanceStatusConflict"
//  FAILEDOPERATION_PGCONNECTERROR = "FailedOperation.PGConnectError"
//  FAILEDOPERATION_PGEXECUTESQLERROR = "FailedOperation.PGExecuteSqlError"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
func (c *Client) ExecutePGSql(request *ExecutePGSqlRequest) (response *ExecutePGSqlResponse, err error) {
    return c.ExecutePGSqlWithContext(context.Background(), request)
}

// ExecutePGSql
// 在Postgres数据库上执行SQL
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCESTATUSCONFLICT = "FailedOperation.InstanceStatusConflict"
//  FAILEDOPERATION_PGCONNECTERROR = "FailedOperation.PGConnectError"
//  FAILEDOPERATION_PGEXECUTESQLERROR = "FailedOperation.PGExecuteSqlError"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
func (c *Client) ExecutePGSqlWithContext(ctx context.Context, request *ExecutePGSqlRequest) (response *ExecutePGSqlResponse, err error) {
    if request == nil {
        request = NewExecutePGSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ExecutePGSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecutePGSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecutePGSqlResponse()
    err = c.Send(request, response)
    return
}

func NewGetProvidersRequest() (request *GetProvidersRequest) {
    request = &GetProvidersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "GetProviders")
    
    
    return
}

func NewGetProvidersResponse() (response *GetProvidersResponse) {
    response = &GetProvidersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetProviders
// 查询指定云开发环境下的身份认证源列表。返回该环境已配置的所有身份认证源信息，包括第三方登录（OAuth、OIDC、SAML）、微信小程序登录、自定义登录和邮箱登录等。返回结果包含认证源基本信息、关联应用、配置状态及启用情况。若自定义登录或邮箱登录的身份源尚未创建，接口会自动追加一个默认关闭状态的身份源记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetProviders(request *GetProvidersRequest) (response *GetProvidersResponse, err error) {
    return c.GetProvidersWithContext(context.Background(), request)
}

// GetProviders
// 查询指定云开发环境下的身份认证源列表。返回该环境已配置的所有身份认证源信息，包括第三方登录（OAuth、OIDC、SAML）、微信小程序登录、自定义登录和邮箱登录等。返回结果包含认证源基本信息、关联应用、配置状态及启用情况。若自定义登录或邮箱登录的身份源尚未创建，接口会自动追加一个默认关闭状态的身份源记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetProvidersWithContext(ctx context.Context, request *GetProvidersRequest) (response *GetProvidersResponse, err error) {
    if request == nil {
        request = NewGetProvidersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "GetProviders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProviders require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProvidersResponse()
    err = c.Send(request, response)
    return
}

func NewInquireVmPriceRequest() (request *InquireVmPriceRequest) {
    request = &InquireVmPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "InquireVmPrice")
    
    
    return
}

func NewInquireVmPriceResponse() (response *InquireVmPriceResponse) {
    response = &InquireVmPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquireVmPrice
// 查询服务器价格
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquireVmPrice(request *InquireVmPriceRequest) (response *InquireVmPriceResponse, err error) {
    return c.InquireVmPriceWithContext(context.Background(), request)
}

// InquireVmPrice
// 查询服务器价格
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquireVmPriceWithContext(ctx context.Context, request *InquireVmPriceRequest) (response *InquireVmPriceResponse, err error) {
    if request == nil {
        request = NewInquireVmPriceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "InquireVmPrice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireVmPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireVmPriceResponse()
    err = c.Send(request, response)
    return
}

func NewListTablesRequest() (request *ListTablesRequest) {
    request = &ListTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ListTables")
    
    
    return
}

func NewListTablesResponse() (response *ListTablesResponse) {
    response = &ListTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTables
// 本接口(ListTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等。
//
// 
//
// 该接口跟 [DescribeTables](https://cloud.tencent.com/document/api/876/127962) 接口功能一致，后续该接口可能会下线，请使用 [DescribeTable](https://cloud.tencent.com/document/api/876/127962)接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListTables(request *ListTablesRequest) (response *ListTablesResponse, err error) {
    return c.ListTablesWithContext(context.Background(), request)
}

// ListTables
// 本接口(ListTables)用于查询文档型数据库所有表信息，包括表名、表中数据条数、表中数据量、索引个数及索引的大小等。
//
// 
//
// 该接口跟 [DescribeTables](https://cloud.tencent.com/document/api/876/127962) 接口功能一致，后续该接口可能会下线，请使用 [DescribeTable](https://cloud.tencent.com/document/api/876/127962)接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFREADREQUESTQUOTA = "LimitExceeded.OutOfReadRequestQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListTablesWithContext(ctx context.Context, request *ListTablesRequest) (response *ListTablesResponse, err error) {
    if request == nil {
        request = NewListTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ListTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTablesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClientRequest() (request *ModifyClientRequest) {
    request = &ModifyClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyClient")
    
    
    return
}

func NewModifyClientResponse() (response *ModifyClientResponse) {
    response = &ModifyClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClient
// 修改客户端配置。采用增量更新策略，仅更新请求中传入的非空字段，未传入的字段保持原值不变。支持修改客户端基本信息（名称、图标、描述、主页地址）、安全域名、允许的 Scope 列表、Token 有效期、会话控制策略及启用状态等配置。
//
// Id、Secret、CreatedAt、Meta 等字段在该接口中不可修改，当客户端 ID 等于环境 ID 且客户端尚未创建时，将自动创建默认客户端配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClient(request *ModifyClientRequest) (response *ModifyClientResponse, err error) {
    return c.ModifyClientWithContext(context.Background(), request)
}

// ModifyClient
// 修改客户端配置。采用增量更新策略，仅更新请求中传入的非空字段，未传入的字段保持原值不变。支持修改客户端基本信息（名称、图标、描述、主页地址）、安全域名、允许的 Scope 列表、Token 有效期、会话控制策略及启用状态等配置。
//
// Id、Secret、CreatedAt、Meta 等字段在该接口中不可修改，当客户端 ID 等于环境 ID 且客户端尚未创建时，将自动创建默认客户端配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyClientWithContext(ctx context.Context, request *ModifyClientRequest) (response *ModifyClientResponse, err error) {
    if request == nil {
        request = NewModifyClientRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyClient")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClientResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClsTopicRequest() (request *ModifyClsTopicRequest) {
    request = &ModifyClsTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyClsTopic")
    
    
    return
}

func NewModifyClsTopicResponse() (response *ModifyClsTopicResponse) {
    response = &ModifyClsTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClsTopic
// 修改日志主题
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClsTopic(request *ModifyClsTopicRequest) (response *ModifyClsTopicResponse, err error) {
    return c.ModifyClsTopicWithContext(context.Background(), request)
}

// ModifyClsTopic
// 修改日志主题
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENT = "LimitExceeded.Concurrent"
//  LIMITEXCEEDED_REQUEST = "LimitExceeded.Request"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyClsTopicWithContext(ctx context.Context, request *ModifyClsTopicRequest) (response *ModifyClsTopicResponse, err error) {
    if request == nil {
        request = NewModifyClsTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyClsTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClsTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClsTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseACLRequest() (request *ModifyDatabaseACLRequest) {
    request = &ModifyDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyDatabaseACL")
    
    
    return
}

func NewModifyDatabaseACLResponse() (response *ModifyDatabaseACLResponse) {
    response = &ModifyDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatabaseACL
// 本接口（ModifyDatabaseACL）用于修改文档型数据库权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"
func (c *Client) ModifyDatabaseACL(request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    return c.ModifyDatabaseACLWithContext(context.Background(), request)
}

// ModifyDatabaseACL
// 本接口（ModifyDatabaseACL）用于修改文档型数据库权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
//  RESOURCEINUSE_FSACLJOBUNDONE = "ResourceInUse.FsACLJobUnDone"
func (c *Client) ModifyDatabaseACLWithContext(ctx context.Context, request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyDatabaseACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvRequest() (request *ModifyEnvRequest) {
    request = &ModifyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnv")
    
    
    return
}

func NewModifyEnvResponse() (response *ModifyEnvResponse) {
    response = &ModifyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnv(request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    return c.ModifyEnvWithContext(context.Background(), request)
}

// ModifyEnv
// 更新环境信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifyEnvWithContext(ctx context.Context, request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvPlanRequest() (request *ModifyEnvPlanRequest) {
    request = &ModifyEnvPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnvPlan")
    
    
    return
}

func NewModifyEnvPlanResponse() (response *ModifyEnvPlanResponse) {
    response = &ModifyEnvPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnvPlan
// 本接口用于变更云开发环境套餐。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) ModifyEnvPlan(request *ModifyEnvPlanRequest) (response *ModifyEnvPlanResponse, err error) {
    return c.ModifyEnvPlanWithContext(context.Background(), request)
}

// ModifyEnvPlan
// 本接口用于变更云开发环境套餐。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) ModifyEnvPlanWithContext(ctx context.Context, request *ModifyEnvPlanRequest) (response *ModifyEnvPlanResponse, err error) {
    if request == nil {
        request = NewModifyEnvPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyEnvPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHTTPServiceRouteRequest() (request *ModifyHTTPServiceRouteRequest) {
    request = &ModifyHTTPServiceRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyHTTPServiceRoute")
    
    
    return
}

func NewModifyHTTPServiceRouteResponse() (response *ModifyHTTPServiceRouteResponse) {
    response = &ModifyHTTPServiceRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHTTPServiceRoute
// 本接口ModifyHTTPServiceRoute用于修改HTTP访问服务路由。支持增量修改，对应字段不传参数则不修改
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTVERIFYFAILED = "InvalidParameter.CertVerifyFailed"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED_HTTPSERVICEROUTE = "LimitExceeded.HTTPServiceRoute"
//  OPERATIONDENIED_HTTPSERVICEDOMAINPROCESSING = "OperationDenied.HTTPServiceDomainProcessing"
func (c *Client) ModifyHTTPServiceRoute(request *ModifyHTTPServiceRouteRequest) (response *ModifyHTTPServiceRouteResponse, err error) {
    return c.ModifyHTTPServiceRouteWithContext(context.Background(), request)
}

// ModifyHTTPServiceRoute
// 本接口ModifyHTTPServiceRoute用于修改HTTP访问服务路由。支持增量修改，对应字段不传参数则不修改
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTVERIFYFAILED = "InvalidParameter.CertVerifyFailed"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  LIMITEXCEEDED_HTTPSERVICEROUTE = "LimitExceeded.HTTPServiceRoute"
//  OPERATIONDENIED_HTTPSERVICEDOMAINPROCESSING = "OperationDenied.HTTPServiceDomainProcessing"
func (c *Client) ModifyHTTPServiceRouteWithContext(ctx context.Context, request *ModifyHTTPServiceRouteRequest) (response *ModifyHTTPServiceRouteResponse, err error) {
    if request == nil {
        request = NewModifyHTTPServiceRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyHTTPServiceRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHTTPServiceRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHTTPServiceRouteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginConfigRequest() (request *ModifyLoginConfigRequest) {
    request = &ModifyLoginConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyLoginConfig")
    
    
    return
}

func NewModifyLoginConfigResponse() (response *ModifyLoginConfigResponse) {
    response = &ModifyLoginConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoginConfig
// 修改指定云开发环境的登录策略配置。支持开启或关闭手机号短信登录、邮箱登录、用户名密码登录和匿名登录，同时可配置短信验证码发送通道、MFA 多因子认证和密码更新策略。
//
// 修改后立即生效，影响该环境下所有终端用户的登录行为。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoginConfig(request *ModifyLoginConfigRequest) (response *ModifyLoginConfigResponse, err error) {
    return c.ModifyLoginConfigWithContext(context.Background(), request)
}

// ModifyLoginConfig
// 修改指定云开发环境的登录策略配置。支持开启或关闭手机号短信登录、邮箱登录、用户名密码登录和匿名登录，同时可配置短信验证码发送通道、MFA 多因子认证和密码更新策略。
//
// 修改后立即生效，影响该环境下所有终端用户的登录行为。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoginConfigWithContext(ctx context.Context, request *ModifyLoginConfigRequest) (response *ModifyLoginConfigResponse, err error) {
    if request == nil {
        request = NewModifyLoginConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyLoginConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoginConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoginConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProviderRequest() (request *ModifyProviderRequest) {
    request = &ModifyProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyProvider")
    
    
    return
}

func NewModifyProviderResponse() (response *ModifyProviderResponse) {
    response = &ModifyProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProvider
// 修改身份认证源。更新指定云开发环境下已有身份认证源的配置信息，支持修改基本信息（名称、图标、描述）、协议连接配置（ClientId、ClientSecret、端点地址等）、登录行为控制（透传模式、自动注册、邮箱/手机号自动关联）以及启用状态。
//
// 对于 OIDC 类型身份源，修改 Issuer 后将自动通过 OpenID Connect Discovery 重新获取端点配置。
//
// 若自定义登录（CUSTOM）或邮箱登录（EMAIL）身份源尚不存在，调用该接口时将自动创建。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyProvider(request *ModifyProviderRequest) (response *ModifyProviderResponse, err error) {
    return c.ModifyProviderWithContext(context.Background(), request)
}

// ModifyProvider
// 修改身份认证源。更新指定云开发环境下已有身份认证源的配置信息，支持修改基本信息（名称、图标、描述）、协议连接配置（ClientId、ClientSecret、端点地址等）、登录行为控制（透传模式、自动注册、邮箱/手机号自动关联）以及启用状态。
//
// 对于 OIDC 类型身份源，修改 Issuer 后将自动通过 OpenID Connect Discovery 重新获取端点配置。
//
// 若自定义登录（CUSTOM）或邮箱登录（EMAIL）身份源尚不存在，调用该接口时将自动创建。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyProviderWithContext(ctx context.Context, request *ModifyProviderRequest) (response *ModifyProviderResponse, err error) {
    if request == nil {
        request = NewModifyProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProviderResponse()
    err = c.Send(request, response)
    return
}

func NewModifySafeRuleRequest() (request *ModifySafeRuleRequest) {
    request = &ModifySafeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifySafeRule")
    
    
    return
}

func NewModifySafeRuleResponse() (response *ModifySafeRuleResponse) {
    response = &ModifySafeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySafeRule
// 设置数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍 ](https://cloud.tencent.com/document/product/876/123478)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifySafeRule(request *ModifySafeRuleRequest) (response *ModifySafeRuleResponse, err error) {
    return c.ModifySafeRuleWithContext(context.Background(), request)
}

// ModifySafeRule
// 设置数据库安全规则。
//
// 安全规则，用于控制C端用户的访问权限。详见 [安全规则介绍 ](https://cloud.tencent.com/document/product/876/123478)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_PARAM = "MissingParameter.Param"
func (c *Client) ModifySafeRuleWithContext(ctx context.Context, request *ModifySafeRuleRequest) (response *ModifySafeRuleResponse, err error) {
    if request == nil {
        request = NewModifySafeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifySafeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySafeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySafeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStorageSourceRequest() (request *ModifyStorageSourceRequest) {
    request = &ModifyStorageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyStorageSource")
    
    
    return
}

func NewModifyStorageSourceResponse() (response *ModifyStorageSourceResponse) {
    response = &ModifyStorageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStorageSource
// 修改指定云开发环境已绑定的外部云存储源配置。
//
// 修改之后，大约3~5分钟生效。
//
// 
//
// 注意⚠️
//
// 本接口仅更新存储源绑定关系，不会迁移您的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) ModifyStorageSource(request *ModifyStorageSourceRequest) (response *ModifyStorageSourceResponse, err error) {
    return c.ModifyStorageSourceWithContext(context.Background(), request)
}

// ModifyStorageSource
// 修改指定云开发环境已绑定的外部云存储源配置。
//
// 修改之后，大约3~5分钟生效。
//
// 
//
// 注意⚠️
//
// 本接口仅更新存储源绑定关系，不会迁移您的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) ModifyStorageSourceWithContext(ctx context.Context, request *ModifyStorageSourceRequest) (response *ModifyStorageSourceResponse, err error) {
    if request == nil {
        request = NewModifyStorageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyStorageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStorageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStorageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// 修改tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 修改tcb用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDDATA = "FailedOperation.DuplicatedData"
//  FAILEDOPERATION_FLEXDBRESOURCEOVERDUE = "FailedOperation.FlexdbResourceOverdue"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXISTS = "ResourceNotFound.UserNotExists"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ModifyUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseEnvRequest() (request *ReleaseEnvRequest) {
    request = &ReleaseEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "ReleaseEnv")
    
    
    return
}

func NewReleaseEnvResponse() (response *ReleaseEnvResponse) {
    response = &ReleaseEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseEnv
// 从环境池里立即取出1个环境
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGIDENTIFIER = "InvalidParameter.MissingIdentifier"
//  RESOURCENOTFOUND_ENVNOTEXIST = "ResourceNotFound.EnvNotExist"
func (c *Client) ReleaseEnv(request *ReleaseEnvRequest) (response *ReleaseEnvResponse, err error) {
    return c.ReleaseEnvWithContext(context.Background(), request)
}

// ReleaseEnv
// 从环境池里立即取出1个环境
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGIDENTIFIER = "InvalidParameter.MissingIdentifier"
//  RESOURCENOTFOUND_ENVNOTEXIST = "ResourceNotFound.EnvNotExist"
func (c *Client) ReleaseEnvWithContext(ctx context.Context, request *ReleaseEnvRequest) (response *ReleaseEnvResponse, err error) {
    if request == nil {
        request = NewReleaseEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "ReleaseEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseEnvResponse()
    err = c.Send(request, response)
    return
}

func NewRenewEnvRequest() (request *RenewEnvRequest) {
    request = &RenewEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RenewEnv")
    
    
    return
}

func NewRenewEnvResponse() (response *RenewEnvResponse) {
    response = &RenewEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewEnv
// 本接口用于云开发环境套餐续费。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) RenewEnv(request *RenewEnvRequest) (response *RenewEnvResponse, err error) {
    return c.RenewEnvWithContext(context.Background(), request)
}

// RenewEnv
// 本接口用于云开发环境套餐续费。
//
// 该接口会自动下单并支付，会在腾讯云账户中扣除余额（余额不足会下单失败）。
//
// 该接口支持自动扣除代金券（AutoVoucher=true时），符合条件的代金券会被自动扣除。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
func (c *Client) RenewEnvWithContext(ctx context.Context, request *RenewEnvRequest) (response *RenewEnvResponse, err error) {
    if request == nil {
        request = NewRenewEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RenewEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewEnvResponse()
    err = c.Send(request, response)
    return
}

func NewRunCommandsRequest() (request *RunCommandsRequest) {
    request = &RunCommandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RunCommands")
    
    
    return
}

func NewRunCommandsResponse() (response *RunCommandsResponse) {
    response = &RunCommandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunCommands
// 本接口（RunCommands）用于执行文档型数据库命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOC = "InvalidParameterValue.InvalidDoc"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFRESULTSIZELIMIT = "LimitExceeded.OutOfResultSizeLimit"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
func (c *Client) RunCommands(request *RunCommandsRequest) (response *RunCommandsResponse, err error) {
    return c.RunCommandsWithContext(context.Background(), request)
}

// RunCommands
// 本接口（RunCommands）用于执行文档型数据库命令。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOC = "InvalidParameterValue.InvalidDoc"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFRESULTSIZELIMIT = "LimitExceeded.OutOfResultSizeLimit"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
func (c *Client) RunCommandsWithContext(ctx context.Context, request *RunCommandsRequest) (response *RunCommandsResponse, err error) {
    if request == nil {
        request = NewRunCommandsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RunCommands")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunCommands require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunCommandsResponse()
    err = c.Send(request, response)
    return
}

func NewRunSqlRequest() (request *RunSqlRequest) {
    request = &RunSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "RunSql")
    
    
    return
}

func NewRunSqlResponse() (response *RunSqlResponse) {
    response = &RunSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunSql
// 本接口（RunSql）用于执行MySQL语句。
//
// 
//
// 该接口用来执行 MySql 语句，比如创建表格、插入数据、修改数据、删除字段、添加索引等可以通过sql 语句实现的都可以使用该接口。
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有开通成功才能操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"
//  FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"
//  FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"
//  FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"
//  FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
//  UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
func (c *Client) RunSql(request *RunSqlRequest) (response *RunSqlResponse, err error) {
    return c.RunSqlWithContext(context.Background(), request)
}

// RunSql
// 本接口（RunSql）用于执行MySQL语句。
//
// 
//
// 该接口用来执行 MySql 语句，比如创建表格、插入数据、修改数据、删除字段、添加索引等可以通过sql 语句实现的都可以使用该接口。
//
// 
//
// 调用该接口前需要先查询Mysql是否开通，可通过 [DescribeCreateMySQLResult ](https://cloud.tencent.com/document/api/876/128185) 查询，只有开通成功才能操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASECONNECTERROR = "FailedOperation.DatabaseConnectError"
//  FAILEDOPERATION_DATABASEEXECSQLERROR = "FailedOperation.DatabaseExecSqlError"
//  FAILEDOPERATION_DATABASESCHEMAERROR = "FailedOperation.DatabaseSchemaError"
//  FAILEDOPERATION_EMPTYDATABASEENDPOINT = "FailedOperation.EmptyDatabaseEndpoint"
//  FAILEDOPERATION_TDSQLPAUSED = "FailedOperation.TdsqlPaused"
//  INTERNALERROR_SYS_ERR = "InternalError.SYS_ERR"
//  INVALIDPARAMETER_INVALID_PARAM = "InvalidParameter.INVALID_PARAM"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
//  UNSUPPORTEDOPERATION_TOOMANYTABLES = "UnsupportedOperation.TooManyTables"
func (c *Client) RunSqlWithContext(ctx context.Context, request *RunSqlRequest) (response *RunSqlResponse, err error) {
    if request == nil {
        request = NewRunSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "RunSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunSqlResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClsLog
// 搜索用户调用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    return c.SearchClsLogWithContext(context.Background(), request)
}

// SearchClsLog
// 搜索用户调用日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "SearchClsLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindStorageSourceRequest() (request *UnbindStorageSourceRequest) {
    request = &UnbindStorageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "UnbindStorageSource")
    
    
    return
}

func NewUnbindStorageSourceResponse() (response *UnbindStorageSourceResponse) {
    response = &UnbindStorageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindStorageSource
// 从指定云开发环境中解绑已绑定的外部云存储源。解绑后，该环境将不再关联外部 存储源，云存储功能恢复为未绑定状态。
//
// 解绑操作仅移除 CloudBase 侧的绑定关系，不会删除桶本身及桶内数据，桶仍由用户自行管理。
//
// 
//
// 注意⚠️
//
// 解绑之后，会导致云存储不可用，请谨慎操作。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) UnbindStorageSource(request *UnbindStorageSourceRequest) (response *UnbindStorageSourceResponse, err error) {
    return c.UnbindStorageSourceWithContext(context.Background(), request)
}

// UnbindStorageSource
// 从指定云开发环境中解绑已绑定的外部云存储源。解绑后，该环境将不再关联外部 存储源，云存储功能恢复为未绑定状态。
//
// 解绑操作仅移除 CloudBase 侧的绑定关系，不会删除桶本身及桶内数据，桶仍由用户自行管理。
//
// 
//
// 注意⚠️
//
// 解绑之后，会导致云存储不可用，请谨慎操作。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) UnbindStorageSourceWithContext(ctx context.Context, request *UnbindStorageSourceRequest) (response *UnbindStorageSourceResponse, err error) {
    if request == nil {
        request = NewUnbindStorageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "UnbindStorageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindStorageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindStorageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIModelRequest() (request *UpdateAIModelRequest) {
    request = &UpdateAIModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "UpdateAIModel")
    
    
    return
}

func NewUpdateAIModelResponse() (response *UpdateAIModelResponse) {
    response = &UpdateAIModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAIModel
// 更新 AI 模型配置分组。支持修改分组的模型列表、服务地址、访问密钥、备注及启用状态。
//
// 
//
// 不同分组类型支持的更新操作如下：
//
// 自定义类型（custom）：可更新模型服务地址（BaseUrl）、访问密钥（Secret）、模型列表及分组备注。
//
// 内置类型（builtin）：默认由云开发平台统一管理服务地址和密钥。若同时传入自定义 BaseURL 和 Secret，该分组将自动转换为自定义类型（custom），后续使用用户自行提供的服务地址和访问密钥，平台不再托管其凭证。
//
// 
//
// 分组类型转换说明:
//
// 调用本接口时，可通过传入 BaseURL 与 Secret 参数触发分组类型的自动转换，规则如下：
//
// (1) builtin → custom（内置转自定义）
//
// 当分组当前类型（Type）为 builtin 时，若请求中同时传入有效的 BaseURL（非 http://default.tcb）和 Secret，系统将自动将该分组转换为自定义类型（Type = custom）。转换后，平台不再托管该分组的服务地址和访问密钥，后续将使用用户自行提供的 BaseUrl 与 Secret 对模型服务发起请求。
//
// 
//
// (2) custom → builtin（自定义恢复内置托管）
//
// 仅当分组的原始类型（OriginType）为 builtin 时，支持将分组恢复为内置托管类型。将 BaseUrl 传入固定值 http://default.tcb，且不传入 Secret，系统将自动将该分组转换回内置托管类型（Type = builtin），平台重新接管其服务地址和访问密钥。
//
// 若 OriginType 为 CUSTOM（即用户通过 [CreateAIModel](https://cloud.tencent.com/document/product/876/131320) 接口自行创建的自定义分组），不支持恢复为内置托管类型。
//
// 
//
// 更新成功后，可通过 [DescribeAIModels](https://cloud.tencent.com/document/product/876/131318) 接口查询最新分组配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) UpdateAIModel(request *UpdateAIModelRequest) (response *UpdateAIModelResponse, err error) {
    return c.UpdateAIModelWithContext(context.Background(), request)
}

// UpdateAIModel
// 更新 AI 模型配置分组。支持修改分组的模型列表、服务地址、访问密钥、备注及启用状态。
//
// 
//
// 不同分组类型支持的更新操作如下：
//
// 自定义类型（custom）：可更新模型服务地址（BaseUrl）、访问密钥（Secret）、模型列表及分组备注。
//
// 内置类型（builtin）：默认由云开发平台统一管理服务地址和密钥。若同时传入自定义 BaseURL 和 Secret，该分组将自动转换为自定义类型（custom），后续使用用户自行提供的服务地址和访问密钥，平台不再托管其凭证。
//
// 
//
// 分组类型转换说明:
//
// 调用本接口时，可通过传入 BaseURL 与 Secret 参数触发分组类型的自动转换，规则如下：
//
// (1) builtin → custom（内置转自定义）
//
// 当分组当前类型（Type）为 builtin 时，若请求中同时传入有效的 BaseURL（非 http://default.tcb）和 Secret，系统将自动将该分组转换为自定义类型（Type = custom）。转换后，平台不再托管该分组的服务地址和访问密钥，后续将使用用户自行提供的 BaseUrl 与 Secret 对模型服务发起请求。
//
// 
//
// (2) custom → builtin（自定义恢复内置托管）
//
// 仅当分组的原始类型（OriginType）为 builtin 时，支持将分组恢复为内置托管类型。将 BaseUrl 传入固定值 http://default.tcb，且不传入 Secret，系统将自动将该分组转换回内置托管类型（Type = builtin），平台重新接管其服务地址和访问密钥。
//
// 若 OriginType 为 CUSTOM（即用户通过 [CreateAIModel](https://cloud.tencent.com/document/product/876/131320) 接口自行创建的自定义分组），不支持恢复为内置托管类型。
//
// 
//
// 更新成功后，可通过 [DescribeAIModels](https://cloud.tencent.com/document/product/876/131318) 接口查询最新分组配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENVID = "InvalidParameter.EnvId"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) UpdateAIModelWithContext(ctx context.Context, request *UpdateAIModelRequest) (response *UpdateAIModelResponse, err error) {
    if request == nil {
        request = NewUpdateAIModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "UpdateAIModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIModelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTableRequest() (request *UpdateTableRequest) {
    request = &UpdateTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcb", APIVersion, "UpdateTable")
    
    
    return
}

func NewUpdateTableResponse() (response *UpdateTableResponse) {
    response = &UpdateTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTable
// 本接口(UpdateTable)用于修改文档型数据库表信息，当前可以支持创建和删除索引。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFINDEXQUOTA = "LimitExceeded.OutOfIndexQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_INDEXCREATING = "ResourceInUse.IndexCreating"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INDEXOPTIONSCONFLICT = "ResourceUnavailable.IndexOptionsConflict"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateTable(request *UpdateTableRequest) (response *UpdateTableResponse, err error) {
    return c.UpdateTableWithContext(context.Background(), request)
}

// UpdateTable
// 本接口(UpdateTable)用于修改文档型数据库表信息，当前可以支持创建和删除索引。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LISTTABLE = "FailedOperation.ListTable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NOVALIDCONNECTION = "LimitExceeded.NoValidConnection"
//  LIMITEXCEEDED_OUTOFINDEXQUOTA = "LimitExceeded.OutOfIndexQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_INDEXCREATING = "ResourceInUse.IndexCreating"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONNECTOR = "ResourceNotFound.Connector"
//  RESOURCENOTFOUND_TABLE = "ResourceNotFound.Table"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_INDEXOPTIONSCONFLICT = "ResourceUnavailable.IndexOptionsConflict"
//  RESOURCEUNAVAILABLE_MONGOISOLATED = "ResourceUnavailable.MongoIsolated"
//  RESOURCEUNAVAILABLE_RESOURCEOVERDUE = "ResourceUnavailable.ResourceOverdue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateTableWithContext(ctx context.Context, request *UpdateTableRequest) (response *UpdateTableResponse, err error) {
    if request == nil {
        request = NewUpdateTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcb", APIVersion, "UpdateTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTableResponse()
    err = c.Send(request, response)
    return
}
