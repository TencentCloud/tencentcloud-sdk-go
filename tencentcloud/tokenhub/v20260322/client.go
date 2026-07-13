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

package v20260322

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-03-22"

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


func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateApiKey")
    
    
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiKey
// 创建 API 密钥。
//
// 
//
// 创建一个新的 API 密钥，创建成功后返回 API 密钥 ID。需指定平台类型、绑定方式和初始状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    return c.CreateApiKeyWithContext(context.Background(), request)
}

// CreateApiKey
// 创建 API 密钥。
//
// 
//
// 创建一个新的 API 密钥，创建成功后返回 API 密钥 ID。需指定平台类型、绑定方式和初始状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndpointRequest() (request *CreateEndpointRequest) {
    request = &CreateEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateEndpoint")
    
    
    return
}

func NewCreateEndpointResponse() (response *CreateEndpointResponse) {
    response = &CreateEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEndpoint
// 创建推理服务。
//
// 
//
// 创建一个在线推理服务，创建成功后返回推理服务 ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEINSTANCEFAILED = "FailedOperation.CreateInstanceFailed"
//  FAILEDOPERATION_ENABLEPOSTPAIDFAILED = "FailedOperation.EnablePostPaidFailed"
//  FAILEDOPERATION_ENDPOINTALREADYEXISTS = "FailedOperation.EndpointAlreadyExists"
//  FAILEDOPERATION_FREEQUOTAEXHAUSTED = "FailedOperation.FreeQuotaExhausted"
//  FAILEDOPERATION_PACKAGEQUERYFAILED = "FailedOperation.PackageQueryFailed"
//  FAILEDOPERATION_PURCHASETPMFAILED = "FailedOperation.PurchaseTpmFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_MODELIDNOTFOUND = "InvalidParameter.ModelIdNotFound"
//  INVALIDPARAMETER_QPMLIMITEXCEEDED = "InvalidParameter.QPMLimitExceeded"
//  INVALIDPARAMETER_TPMBELOWQUOTA = "InvalidParameter.TPMBelowQuota"
//  INVALIDPARAMETER_TPMLIMITEXCEEDED = "InvalidParameter.TPMLimitExceeded"
//  INVALIDPARAMETERVALUE_ENDPOINTNAMETOOLONG = "InvalidParameterValue.EndpointNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_ENDPOINTQUOTA = "LimitExceeded.EndpointQuota"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateEndpoint(request *CreateEndpointRequest) (response *CreateEndpointResponse, err error) {
    return c.CreateEndpointWithContext(context.Background(), request)
}

// CreateEndpoint
// 创建推理服务。
//
// 
//
// 创建一个在线推理服务，创建成功后返回推理服务 ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEINSTANCEFAILED = "FailedOperation.CreateInstanceFailed"
//  FAILEDOPERATION_ENABLEPOSTPAIDFAILED = "FailedOperation.EnablePostPaidFailed"
//  FAILEDOPERATION_ENDPOINTALREADYEXISTS = "FailedOperation.EndpointAlreadyExists"
//  FAILEDOPERATION_FREEQUOTAEXHAUSTED = "FailedOperation.FreeQuotaExhausted"
//  FAILEDOPERATION_PACKAGEQUERYFAILED = "FailedOperation.PackageQueryFailed"
//  FAILEDOPERATION_PURCHASETPMFAILED = "FailedOperation.PurchaseTpmFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_MODELIDNOTFOUND = "InvalidParameter.ModelIdNotFound"
//  INVALIDPARAMETER_QPMLIMITEXCEEDED = "InvalidParameter.QPMLimitExceeded"
//  INVALIDPARAMETER_TPMBELOWQUOTA = "InvalidParameter.TPMBelowQuota"
//  INVALIDPARAMETER_TPMLIMITEXCEEDED = "InvalidParameter.TPMLimitExceeded"
//  INVALIDPARAMETERVALUE_ENDPOINTNAMETOOLONG = "InvalidParameterValue.EndpointNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_ENDPOINTQUOTA = "LimitExceeded.EndpointQuota"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateEndpointWithContext(ctx context.Context, request *CreateEndpointRequest) (response *CreateEndpointResponse, err error) {
    if request == nil {
        request = NewCreateEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlossaryRequest() (request *CreateGlossaryRequest) {
    request = &CreateGlossaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateGlossary")
    
    
    return
}

func NewCreateGlossaryResponse() (response *CreateGlossaryResponse) {
    response = &CreateGlossaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlossary
// 创建术语库。
//
// 
//
// 在当前应用下创建一个新的翻译术语库，用于自定义源语言到目标语言的术语映射。创建成功后返回术语库 ID，可通过该 ID 进一步管理术语条目。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossary(request *CreateGlossaryRequest) (response *CreateGlossaryResponse, err error) {
    return c.CreateGlossaryWithContext(context.Background(), request)
}

// CreateGlossary
// 创建术语库。
//
// 
//
// 在当前应用下创建一个新的翻译术语库，用于自定义源语言到目标语言的术语映射。创建成功后返回术语库 ID，可通过该 ID 进一步管理术语条目。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryWithContext(ctx context.Context, request *CreateGlossaryRequest) (response *CreateGlossaryResponse, err error) {
    if request == nil {
        request = NewCreateGlossaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateGlossary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlossary require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlossaryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlossaryEntriesRequest() (request *CreateGlossaryEntriesRequest) {
    request = &CreateGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateGlossaryEntries")
    
    
    return
}

func NewCreateGlossaryEntriesResponse() (response *CreateGlossaryEntriesResponse) {
    response = &CreateGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlossaryEntries
// 批量创建术语条目。
//
// 
//
// 在指定术语库下批量创建术语条目。单次最多创建 100 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryEntries(request *CreateGlossaryEntriesRequest) (response *CreateGlossaryEntriesResponse, err error) {
    return c.CreateGlossaryEntriesWithContext(context.Background(), request)
}

// CreateGlossaryEntries
// 批量创建术语条目。
//
// 
//
// 在指定术语库下批量创建术语条目。单次最多创建 100 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryEntriesWithContext(ctx context.Context, request *CreateGlossaryEntriesRequest) (response *CreateGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewCreateGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTokenPlanApiKeysRequest() (request *CreateTokenPlanApiKeysRequest) {
    request = &CreateTokenPlanApiKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateTokenPlanApiKeys")
    
    
    return
}

func NewCreateTokenPlanApiKeysResponse() (response *CreateTokenPlanApiKeysResponse) {
    response = &CreateTokenPlanApiKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTokenPlanApiKeys
// 批量创建 TokenPlan APIKey。
//
// 
//
// 传入名称前缀和数量，自动按 {ApiKeyName}-{序号} 格式生成名称（如 aaa-1, aaa-2）。允许同名。支持部分成功，最多 100 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanApiKeys(request *CreateTokenPlanApiKeysRequest) (response *CreateTokenPlanApiKeysResponse, err error) {
    return c.CreateTokenPlanApiKeysWithContext(context.Background(), request)
}

// CreateTokenPlanApiKeys
// 批量创建 TokenPlan APIKey。
//
// 
//
// 传入名称前缀和数量，自动按 {ApiKeyName}-{序号} 格式生成名称（如 aaa-1, aaa-2）。允许同名。支持部分成功，最多 100 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanApiKeysWithContext(ctx context.Context, request *CreateTokenPlanApiKeysRequest) (response *CreateTokenPlanApiKeysResponse, err error) {
    if request == nil {
        request = NewCreateTokenPlanApiKeysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateTokenPlanApiKeys")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTokenPlanApiKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTokenPlanApiKeysResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTokenPlanTeamOrderAndBuyRequest() (request *CreateTokenPlanTeamOrderAndBuyRequest) {
    request = &CreateTokenPlanTeamOrderAndBuyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateTokenPlanTeamOrderAndBuy")
    
    
    return
}

func NewCreateTokenPlanTeamOrderAndBuyResponse() (response *CreateTokenPlanTeamOrderAndBuyResponse) {
    response = &CreateTokenPlanTeamOrderAndBuyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTokenPlanTeamOrderAndBuy
// 购买套餐。
//
// 
//
// 发起 TokenPlan 套餐下单并完成支付，成功后返回大订单 ID 及关联的子订单、资源信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanTeamOrderAndBuy(request *CreateTokenPlanTeamOrderAndBuyRequest) (response *CreateTokenPlanTeamOrderAndBuyResponse, err error) {
    return c.CreateTokenPlanTeamOrderAndBuyWithContext(context.Background(), request)
}

// CreateTokenPlanTeamOrderAndBuy
// 购买套餐。
//
// 
//
// 发起 TokenPlan 套餐下单并完成支付，成功后返回大订单 ID 及关联的子订单、资源信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanTeamOrderAndBuyWithContext(ctx context.Context, request *CreateTokenPlanTeamOrderAndBuyRequest) (response *CreateTokenPlanTeamOrderAndBuyResponse, err error) {
    if request == nil {
        request = NewCreateTokenPlanTeamOrderAndBuyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateTokenPlanTeamOrderAndBuy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTokenPlanTeamOrderAndBuy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTokenPlanTeamOrderAndBuyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteApiKey")
    
    
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiKey
// 删除指定的 API 密钥，同时清理关联的模型绑定关系。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    return c.DeleteApiKeyWithContext(context.Background(), request)
}

// DeleteApiKey
// 删除指定的 API 密钥，同时清理关联的模型绑定关系。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndpointRequest() (request *DeleteEndpointRequest) {
    request = &DeleteEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteEndpoint")
    
    
    return
}

func NewDeleteEndpointResponse() (response *DeleteEndpointResponse) {
    response = &DeleteEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEndpoint
// 删除推理服务。
//
// 
//
// 删除指定的推理服务端点，操作不可逆。调用接口后，若通过 DescribeEndpoint 接口查询不到对应的端点，则表示删除成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteEndpoint(request *DeleteEndpointRequest) (response *DeleteEndpointResponse, err error) {
    return c.DeleteEndpointWithContext(context.Background(), request)
}

// DeleteEndpoint
// 删除推理服务。
//
// 
//
// 删除指定的推理服务端点，操作不可逆。调用接口后，若通过 DescribeEndpoint 接口查询不到对应的端点，则表示删除成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteEndpointWithContext(ctx context.Context, request *DeleteEndpointRequest) (response *DeleteEndpointResponse, err error) {
    if request == nil {
        request = NewDeleteEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlossaryRequest() (request *DeleteGlossaryRequest) {
    request = &DeleteGlossaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteGlossary")
    
    
    return
}

func NewDeleteGlossaryResponse() (response *DeleteGlossaryResponse) {
    response = &DeleteGlossaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlossary
// 删除术语库。
//
// 
//
// 删除指定的术语库及其下所有术语条目。删除操作幂等，对不存在的术语库返回成功。调用接口后，若通过 DescribeGlossaries 接口查询不到对应术语库，则表示删除成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossary(request *DeleteGlossaryRequest) (response *DeleteGlossaryResponse, err error) {
    return c.DeleteGlossaryWithContext(context.Background(), request)
}

// DeleteGlossary
// 删除术语库。
//
// 
//
// 删除指定的术语库及其下所有术语条目。删除操作幂等，对不存在的术语库返回成功。调用接口后，若通过 DescribeGlossaries 接口查询不到对应术语库，则表示删除成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryWithContext(ctx context.Context, request *DeleteGlossaryRequest) (response *DeleteGlossaryResponse, err error) {
    if request == nil {
        request = NewDeleteGlossaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteGlossary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlossary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlossaryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlossaryEntriesRequest() (request *DeleteGlossaryEntriesRequest) {
    request = &DeleteGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteGlossaryEntries")
    
    
    return
}

func NewDeleteGlossaryEntriesResponse() (response *DeleteGlossaryEntriesResponse) {
    response = &DeleteGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlossaryEntries
// 批量删除术语条目。
//
// 
//
// 在指定术语库下批量删除术语条目。单次最多删除 200 条。若术语库不存在或不属于当前应用，返回 ResourceNotFound 错误。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryEntries(request *DeleteGlossaryEntriesRequest) (response *DeleteGlossaryEntriesResponse, err error) {
    return c.DeleteGlossaryEntriesWithContext(context.Background(), request)
}

// DeleteGlossaryEntries
// 批量删除术语条目。
//
// 
//
// 在指定术语库下批量删除术语条目。单次最多删除 200 条。若术语库不存在或不属于当前应用，返回 ResourceNotFound 错误。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryEntriesWithContext(ctx context.Context, request *DeleteGlossaryEntriesRequest) (response *DeleteGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewDeleteGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTokenPlanApiKeyRequest() (request *DeleteTokenPlanApiKeyRequest) {
    request = &DeleteTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteTokenPlanApiKey")
    
    
    return
}

func NewDeleteTokenPlanApiKeyResponse() (response *DeleteTokenPlanApiKeyResponse) {
    response = &DeleteTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTokenPlanApiKey
// 删除 TokenPlan APIKey。
//
// 
//
// 同时删除额度中心子额度包并通知网关清除缓存。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteTokenPlanApiKey(request *DeleteTokenPlanApiKeyRequest) (response *DeleteTokenPlanApiKeyResponse, err error) {
    return c.DeleteTokenPlanApiKeyWithContext(context.Background(), request)
}

// DeleteTokenPlanApiKey
// 删除 TokenPlan APIKey。
//
// 
//
// 同时删除额度中心子额度包并通知网关清除缓存。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteTokenPlanApiKeyWithContext(ctx context.Context, request *DeleteTokenPlanApiKeyRequest) (response *DeleteTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeApiKey")
    
    
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKey
// 根据 API 密钥 ID 或密钥值查询 API 密钥详情，返回明文密钥。ApiKeyId 和 ApiKey 至少需传入其一，优先使用 ApiKeyId。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    return c.DescribeApiKeyWithContext(context.Background(), request)
}

// DescribeApiKey
// 根据 API 密钥 ID 或密钥值查询 API 密钥详情，返回明文密钥。ApiKeyId 和 ApiKey 至少需传入其一，优先使用 ApiKeyId。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyWithContext(ctx context.Context, request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyListRequest() (request *DescribeApiKeyListRequest) {
    request = &DescribeApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeApiKeyList")
    
    
    return
}

func NewDescribeApiKeyListResponse() (response *DescribeApiKeyListResponse) {
    response = &DescribeApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKeyList
// 查询 API 密钥列表。
//
// 
//
// 查询当前用户的 API 密钥列表，密钥值脱敏展示。支持分页、过滤和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyList(request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    return c.DescribeApiKeyListWithContext(context.Background(), request)
}

// DescribeApiKeyList
// 查询 API 密钥列表。
//
// 
//
// 查询当前用户的 API 密钥列表，密钥值脱敏展示。支持分页、过滤和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyListWithContext(ctx context.Context, request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndpointRequest() (request *DescribeEndpointRequest) {
    request = &DescribeEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeEndpoint")
    
    
    return
}

func NewDescribeEndpointResponse() (response *DescribeEndpointResponse) {
    response = &DescribeEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEndpoint
// 查询推理服务详情。
//
// 
//
// 根据推理服务 ID 查询推理服务的详细信息，包括计费信息、免费额度、API 调用地址等。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEndpoint(request *DescribeEndpointRequest) (response *DescribeEndpointResponse, err error) {
    return c.DescribeEndpointWithContext(context.Background(), request)
}

// DescribeEndpoint
// 查询推理服务详情。
//
// 
//
// 根据推理服务 ID 查询推理服务的详细信息，包括计费信息、免费额度、API 调用地址等。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEndpointWithContext(ctx context.Context, request *DescribeEndpointRequest) (response *DescribeEndpointResponse, err error) {
    if request == nil {
        request = NewDescribeEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlossariesRequest() (request *DescribeGlossariesRequest) {
    request = &DescribeGlossariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeGlossaries")
    
    
    return
}

func NewDescribeGlossariesResponse() (response *DescribeGlossariesResponse) {
    response = &DescribeGlossariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlossaries
// 查询术语库列表。
//
// 
//
// 查询当前应用下的术语库列表。支持分页、过滤和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaries(request *DescribeGlossariesRequest) (response *DescribeGlossariesResponse, err error) {
    return c.DescribeGlossariesWithContext(context.Background(), request)
}

// DescribeGlossaries
// 查询术语库列表。
//
// 
//
// 查询当前应用下的术语库列表。支持分页、过滤和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossariesWithContext(ctx context.Context, request *DescribeGlossariesRequest) (response *DescribeGlossariesResponse, err error) {
    if request == nil {
        request = NewDescribeGlossariesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeGlossaries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlossaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlossariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlossaryEntriesRequest() (request *DescribeGlossaryEntriesRequest) {
    request = &DescribeGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeGlossaryEntries")
    
    
    return
}

func NewDescribeGlossaryEntriesResponse() (response *DescribeGlossaryEntriesResponse) {
    response = &DescribeGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlossaryEntries
// 查询术语条目列表。
//
// 
//
// 查询指定术语库下的术语条目。支持分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaryEntries(request *DescribeGlossaryEntriesRequest) (response *DescribeGlossaryEntriesResponse, err error) {
    return c.DescribeGlossaryEntriesWithContext(context.Background(), request)
}

// DescribeGlossaryEntries
// 查询术语条目列表。
//
// 
//
// 查询指定术语库下的术语条目。支持分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaryEntriesWithContext(ctx context.Context, request *DescribeGlossaryEntriesRequest) (response *DescribeGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewDescribeGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelEndpointListRequest() (request *DescribeModelEndpointListRequest) {
    request = &DescribeModelEndpointListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeModelEndpointList")
    
    
    return
}

func NewDescribeModelEndpointListResponse() (response *DescribeModelEndpointListResponse) {
    response = &DescribeModelEndpointListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelEndpointList
// 查询模型接入点列表。
//
// 
//
// 以模型为基准展示所有在线文本类型模型的接入点概览，支持按状态、计费方式、创建来源等条件筛选，使用 Offset/Limit 分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeModelEndpointList(request *DescribeModelEndpointListRequest) (response *DescribeModelEndpointListResponse, err error) {
    return c.DescribeModelEndpointListWithContext(context.Background(), request)
}

// DescribeModelEndpointList
// 查询模型接入点列表。
//
// 
//
// 以模型为基准展示所有在线文本类型模型的接入点概览，支持按状态、计费方式、创建来源等条件筛选，使用 Offset/Limit 分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeModelEndpointListWithContext(ctx context.Context, request *DescribeModelEndpointListRequest) (response *DescribeModelEndpointListResponse, err error) {
    if request == nil {
        request = NewDescribeModelEndpointListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeModelEndpointList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelEndpointList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelEndpointListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelListRequest() (request *DescribeModelListRequest) {
    request = &DescribeModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeModelList")
    
    
    return
}

func NewDescribeModelListResponse() (response *DescribeModelListResponse) {
    response = &DescribeModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelList
// 查询模型列表。
//
// 
//
// 支持按模型 ID、模型名称、模型能力等条件筛选，支持分页和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeModelList(request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    return c.DescribeModelListWithContext(context.Background(), request)
}

// DescribeModelList
// 查询模型列表。
//
// 
//
// 支持按模型 ID、模型名称、模型能力等条件筛选，支持分页和排序。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeModelListWithContext(ctx context.Context, request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    if request == nil {
        request = NewDescribeModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanRequest() (request *DescribeTokenPlanRequest) {
    request = &DescribeTokenPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlan")
    
    
    return
}

func NewDescribeTokenPlanResponse() (response *DescribeTokenPlanResponse) {
    response = &DescribeTokenPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlan
// 查询 TokenPlan 套餐详情。
//
// 
//
// 返回套餐基本信息及额度中心主额度包余量。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlan(request *DescribeTokenPlanRequest) (response *DescribeTokenPlanResponse, err error) {
    return c.DescribeTokenPlanWithContext(context.Background(), request)
}

// DescribeTokenPlan
// 查询 TokenPlan 套餐详情。
//
// 
//
// 返回套餐基本信息及额度中心主额度包余量。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanWithContext(ctx context.Context, request *DescribeTokenPlanRequest) (response *DescribeTokenPlanResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyRequest() (request *DescribeTokenPlanApiKeyRequest) {
    request = &DescribeTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKey")
    
    
    return
}

func NewDescribeTokenPlanApiKeyResponse() (response *DescribeTokenPlanApiKeyResponse) {
    response = &DescribeTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKey
// 查询 TokenPlan APIKey 详情。
//
// 
//
// 返回 APIKey 完整信息（含明文密钥）及子额度包余量。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKey(request *DescribeTokenPlanApiKeyRequest) (response *DescribeTokenPlanApiKeyResponse, err error) {
    return c.DescribeTokenPlanApiKeyWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKey
// 查询 TokenPlan APIKey 详情。
//
// 
//
// 返回 APIKey 完整信息（含明文密钥）及子额度包余量。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyRequest) (response *DescribeTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyListRequest() (request *DescribeTokenPlanApiKeyListRequest) {
    request = &DescribeTokenPlanApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeyList")
    
    
    return
}

func NewDescribeTokenPlanApiKeyListResponse() (response *DescribeTokenPlanApiKeyListResponse) {
    response = &DescribeTokenPlanApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeyList
// 查询 TokenPlan APIKey 列表。
//
// 
//
// 返回指定套餐下的 APIKey 列表，密钥已脱敏。主账号可查看全部，子账号仅可查看自己创建的。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyList(request *DescribeTokenPlanApiKeyListRequest) (response *DescribeTokenPlanApiKeyListResponse, err error) {
    return c.DescribeTokenPlanApiKeyListWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeyList
// 查询 TokenPlan APIKey 列表。
//
// 
//
// 返回指定套餐下的 APIKey 列表，密钥已脱敏。主账号可查看全部，子账号仅可查看自己创建的。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyListWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyListRequest) (response *DescribeTokenPlanApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeySecretRequest() (request *DescribeTokenPlanApiKeySecretRequest) {
    request = &DescribeTokenPlanApiKeySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeySecret")
    
    
    return
}

func NewDescribeTokenPlanApiKeySecretResponse() (response *DescribeTokenPlanApiKeySecretResponse) {
    response = &DescribeTokenPlanApiKeySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeySecret
// 查询 TokenPlan APIKey 密钥（明文）。
//
// 
//
// 返回指定 APIKey 的明文密钥值，请妥善保管。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeySecret(request *DescribeTokenPlanApiKeySecretRequest) (response *DescribeTokenPlanApiKeySecretResponse, err error) {
    return c.DescribeTokenPlanApiKeySecretWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeySecret
// 查询 TokenPlan APIKey 密钥（明文）。
//
// 
//
// 返回指定 APIKey 的明文密钥值，请妥善保管。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeySecretWithContext(ctx context.Context, request *DescribeTokenPlanApiKeySecretRequest) (response *DescribeTokenPlanApiKeySecretResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeySecretRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeySecret")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeySecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyUsageDetailRequest() (request *DescribeTokenPlanApiKeyUsageDetailRequest) {
    request = &DescribeTokenPlanApiKeyUsageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeyUsageDetail")
    
    
    return
}

func NewDescribeTokenPlanApiKeyUsageDetailResponse() (response *DescribeTokenPlanApiKeyUsageDetailResponse) {
    response = &DescribeTokenPlanApiKeyUsageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeyUsageDetail
// 查询 TokenPlan APIKey 调用明细。
//
// 
//
// 从 CLS 日志服务查询套餐下的调用明细，按 team_id 过滤，支持游标分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyUsageDetail(request *DescribeTokenPlanApiKeyUsageDetailRequest) (response *DescribeTokenPlanApiKeyUsageDetailResponse, err error) {
    return c.DescribeTokenPlanApiKeyUsageDetailWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeyUsageDetail
// 查询 TokenPlan APIKey 调用明细。
//
// 
//
// 从 CLS 日志服务查询套餐下的调用明细，按 team_id 过滤，支持游标分页。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyUsageDetailWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyUsageDetailRequest) (response *DescribeTokenPlanApiKeyUsageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyUsageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeyUsageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeyUsageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyUsageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanListRequest() (request *DescribeTokenPlanListRequest) {
    request = &DescribeTokenPlanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanList")
    
    
    return
}

func NewDescribeTokenPlanListResponse() (response *DescribeTokenPlanListResponse) {
    response = &DescribeTokenPlanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanList
// 查询 TokenPlan 套餐列表。
//
// 
//
// 支持分页、过滤和排序。主账号可查看全部，子账号仅可查看自己创建的。返回结果包含每个套餐关联的额度中心主额度包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeTokenPlanList(request *DescribeTokenPlanListRequest) (response *DescribeTokenPlanListResponse, err error) {
    return c.DescribeTokenPlanListWithContext(context.Background(), request)
}

// DescribeTokenPlanList
// 查询 TokenPlan 套餐列表。
//
// 
//
// 支持分页、过滤和排序。主账号可查看全部，子账号仅可查看自己创建的。返回结果包含每个套餐关联的额度中心主额度包详情。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeTokenPlanListWithContext(ctx context.Context, request *DescribeTokenPlanListRequest) (response *DescribeTokenPlanListResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageRankListRequest() (request *DescribeUsageRankListRequest) {
    request = &DescribeUsageRankListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeUsageRankList")
    
    
    return
}

func NewDescribeUsageRankListResponse() (response *DescribeUsageRankListResponse) {
    response = &DescribeUsageRankListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageRankList
// 查询用量排行列表。
//
// 
//
// 指标族（MetricType）
//
// - `tokens`（默认）：Token 用量统计。支持 Dimension = apikey / endpoint / model。
//
//   返回指标：TotalToken（总）/ InputTotalToken（输入）/ OutputTotalToken（输出）/ CacheTotalToken（读缓存）。
//
// - `search`：【待上线】联网搜索用量统计。支持 Dimension = apikey / endpoint / model。
//
//   返回指标：SearchRequestCount（搜索请求数）/ SearchCount（搜索引擎调用次数）。
//
// 
//
// 响应内容
//
// - MetricType 字段用于切换指标族，响应回显 MetricType 与 MetricKeys。
//
// - TotalStats：时间窗内全部对象的整段聚合值。
//
// - PageStats：当前翻页内对象的整段聚合值。
//
// - TopList：按MetricKeys[0]降序的对象列表，含整段聚合值与逐时间点曲线。
//
// 可能返回的错误码:
//  INTERNALERROR_BARADERROR = "InternalError.BaradError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERIODEXCEEDSSPAN = "InvalidParameter.PeriodExceedsSpan"
//  INVALIDPARAMETER_PERIODTOOFINEFORDATA = "InvalidParameter.PeriodTooFineForData"
//  INVALIDPARAMETER_TOOMANYOBJECTS = "InvalidParameter.TooManyObjects"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeUsageRankList(request *DescribeUsageRankListRequest) (response *DescribeUsageRankListResponse, err error) {
    return c.DescribeUsageRankListWithContext(context.Background(), request)
}

// DescribeUsageRankList
// 查询用量排行列表。
//
// 
//
// 指标族（MetricType）
//
// - `tokens`（默认）：Token 用量统计。支持 Dimension = apikey / endpoint / model。
//
//   返回指标：TotalToken（总）/ InputTotalToken（输入）/ OutputTotalToken（输出）/ CacheTotalToken（读缓存）。
//
// - `search`：【待上线】联网搜索用量统计。支持 Dimension = apikey / endpoint / model。
//
//   返回指标：SearchRequestCount（搜索请求数）/ SearchCount（搜索引擎调用次数）。
//
// 
//
// 响应内容
//
// - MetricType 字段用于切换指标族，响应回显 MetricType 与 MetricKeys。
//
// - TotalStats：时间窗内全部对象的整段聚合值。
//
// - PageStats：当前翻页内对象的整段聚合值。
//
// - TopList：按MetricKeys[0]降序的对象列表，含整段聚合值与逐时间点曲线。
//
// 可能返回的错误码:
//  INTERNALERROR_BARADERROR = "InternalError.BaradError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERIODEXCEEDSSPAN = "InvalidParameter.PeriodExceedsSpan"
//  INVALIDPARAMETER_PERIODTOOFINEFORDATA = "InvalidParameter.PeriodTooFineForData"
//  INVALIDPARAMETER_TOOMANYOBJECTS = "InvalidParameter.TooManyObjects"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeUsageRankListWithContext(ctx context.Context, request *DescribeUsageRankListRequest) (response *DescribeUsageRankListResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRankListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeUsageRankList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageRankList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageRankListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiKeyInfoRequest() (request *ModifyApiKeyInfoRequest) {
    request = &ModifyApiKeyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyApiKeyInfo")
    
    
    return
}

func NewModifyApiKeyInfoResponse() (response *ModifyApiKeyInfoResponse) {
    response = &ModifyApiKeyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiKeyInfo
// 更新 API 密钥信息。
//
// 
//
// 更新 API 密钥的备注信息、 IP 白名单和 Token 限额（修改限额推荐使用QuotaDesired参数）。所有可选参数不传表示不修改。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyInfo(request *ModifyApiKeyInfoRequest) (response *ModifyApiKeyInfoResponse, err error) {
    return c.ModifyApiKeyInfoWithContext(context.Background(), request)
}

// ModifyApiKeyInfo
// 更新 API 密钥信息。
//
// 
//
// 更新 API 密钥的备注信息、 IP 白名单和 Token 限额（修改限额推荐使用QuotaDesired参数）。所有可选参数不传表示不修改。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyInfoWithContext(ctx context.Context, request *ModifyApiKeyInfoRequest) (response *ModifyApiKeyInfoResponse, err error) {
    if request == nil {
        request = NewModifyApiKeyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyApiKeyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiKeyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiKeyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiKeyStatusRequest() (request *ModifyApiKeyStatusRequest) {
    request = &ModifyApiKeyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyApiKeyStatus")
    
    
    return
}

func NewModifyApiKeyStatusResponse() (response *ModifyApiKeyStatusResponse) {
    response = &ModifyApiKeyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiKeyStatus
// 更新 API 密钥的启用或禁用状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyStatus(request *ModifyApiKeyStatusRequest) (response *ModifyApiKeyStatusResponse, err error) {
    return c.ModifyApiKeyStatusWithContext(context.Background(), request)
}

// ModifyApiKeyStatus
// 更新 API 密钥的启用或禁用状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyStatusWithContext(ctx context.Context, request *ModifyApiKeyStatusRequest) (response *ModifyApiKeyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApiKeyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyApiKeyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiKeyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiKeyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEndpointRequest() (request *ModifyEndpointRequest) {
    request = &ModifyEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyEndpoint")
    
    
    return
}

func NewModifyEndpointResponse() (response *ModifyEndpointResponse) {
    response = &ModifyEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEndpoint
// 修改推理服务。
//
// 
//
// 修改推理服务的属性，支持修改服务名称、QPM/TPM 限流上限、TPM 包续费设置、智能路由开关和手动重试 TPM 购买。
//
// 
//
// 注意事项：
//
// - 不支持通过本接口切换计费类型（ChargeType），计费类型仅可在创建推理服务（CreateEndpoint）时指定。
//
// - 不支持通过本接口修改 TPM 预付费保障包的 quota（TpmInputLimit/TpmOutputLimit/TimeSpan），这些值仅可在创建推理服务时指定。
//
// - 当 RetryTPMPurchase 为 true 时，系统会异步重试 TPM 包购买，调用后需轮询推理服务状态确认结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENDPOINTNOTFOUND = "FailedOperation.EndpointNotFound"
//  FAILEDOPERATION_ENDPOINTNOTHEALTHY = "FailedOperation.EndpointNotHealthy"
//  FAILEDOPERATION_NOTPMPACKAGE = "FailedOperation.NoTPMPackage"
//  FAILEDOPERATION_OPERATIONDENIED = "FailedOperation.OperationDenied"
//  FAILEDOPERATION_SETRENEWFLAGFAILED = "FailedOperation.SetRenewFlagFailed"
//  FAILEDOPERATION_TPMPACKAGEPENDING = "FailedOperation.TPMPackagePending"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_QPMLIMITEXCEEDED = "InvalidParameter.QPMLimitExceeded"
//  INVALIDPARAMETER_TPMINPUTBELOWQUOTA = "InvalidParameter.TPMInputBelowQuota"
//  INVALIDPARAMETER_TPMINPUTLIMITEXCEEDED = "InvalidParameter.TPMInputLimitExceeded"
//  INVALIDPARAMETER_TPMLIMITEXCEEDED = "InvalidParameter.TPMLimitExceeded"
//  INVALIDPARAMETER_TPMOUTPUTBELOWQUOTA = "InvalidParameter.TPMOutputBelowQuota"
//  INVALIDPARAMETER_TPMOUTPUTLIMITEXCEEDED = "InvalidParameter.TPMOutputLimitExceeded"
//  INVALIDPARAMETER_TPMPREQUOTAMODIFYNOTSUPPORTED = "InvalidParameter.TPMPreQuotaModifyNotSupported"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyEndpoint(request *ModifyEndpointRequest) (response *ModifyEndpointResponse, err error) {
    return c.ModifyEndpointWithContext(context.Background(), request)
}

// ModifyEndpoint
// 修改推理服务。
//
// 
//
// 修改推理服务的属性，支持修改服务名称、QPM/TPM 限流上限、TPM 包续费设置、智能路由开关和手动重试 TPM 购买。
//
// 
//
// 注意事项：
//
// - 不支持通过本接口切换计费类型（ChargeType），计费类型仅可在创建推理服务（CreateEndpoint）时指定。
//
// - 不支持通过本接口修改 TPM 预付费保障包的 quota（TpmInputLimit/TpmOutputLimit/TimeSpan），这些值仅可在创建推理服务时指定。
//
// - 当 RetryTPMPurchase 为 true 时，系统会异步重试 TPM 包购买，调用后需轮询推理服务状态确认结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENDPOINTNOTFOUND = "FailedOperation.EndpointNotFound"
//  FAILEDOPERATION_ENDPOINTNOTHEALTHY = "FailedOperation.EndpointNotHealthy"
//  FAILEDOPERATION_NOTPMPACKAGE = "FailedOperation.NoTPMPackage"
//  FAILEDOPERATION_OPERATIONDENIED = "FailedOperation.OperationDenied"
//  FAILEDOPERATION_SETRENEWFLAGFAILED = "FailedOperation.SetRenewFlagFailed"
//  FAILEDOPERATION_TPMPACKAGEPENDING = "FailedOperation.TPMPackagePending"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_QPMLIMITEXCEEDED = "InvalidParameter.QPMLimitExceeded"
//  INVALIDPARAMETER_TPMINPUTBELOWQUOTA = "InvalidParameter.TPMInputBelowQuota"
//  INVALIDPARAMETER_TPMINPUTLIMITEXCEEDED = "InvalidParameter.TPMInputLimitExceeded"
//  INVALIDPARAMETER_TPMLIMITEXCEEDED = "InvalidParameter.TPMLimitExceeded"
//  INVALIDPARAMETER_TPMOUTPUTBELOWQUOTA = "InvalidParameter.TPMOutputBelowQuota"
//  INVALIDPARAMETER_TPMOUTPUTLIMITEXCEEDED = "InvalidParameter.TPMOutputLimitExceeded"
//  INVALIDPARAMETER_TPMPREQUOTAMODIFYNOTSUPPORTED = "InvalidParameter.TPMPreQuotaModifyNotSupported"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  RESOURCENOTFOUND_ENDPOINTNOTFOUND = "ResourceNotFound.EndpointNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyEndpointWithContext(ctx context.Context, request *ModifyEndpointRequest) (response *ModifyEndpointResponse, err error) {
    if request == nil {
        request = NewModifyEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlossaryEntriesRequest() (request *ModifyGlossaryEntriesRequest) {
    request = &ModifyGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyGlossaryEntries")
    
    
    return
}

func NewModifyGlossaryEntriesResponse() (response *ModifyGlossaryEntriesResponse) {
    response = &ModifyGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlossaryEntries
// 批量修改术语条目。
//
// 
//
// 在指定术语库下批量修改术语条目。单次最多修改 200 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGlossaryEntries(request *ModifyGlossaryEntriesRequest) (response *ModifyGlossaryEntriesResponse, err error) {
    return c.ModifyGlossaryEntriesWithContext(context.Background(), request)
}

// ModifyGlossaryEntries
// 批量修改术语条目。
//
// 
//
// 在指定术语库下批量修改术语条目。单次最多修改 200 条。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGlossaryEntriesWithContext(ctx context.Context, request *ModifyGlossaryEntriesRequest) (response *ModifyGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewModifyGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTokenPlanApiKeyRequest() (request *ModifyTokenPlanApiKeyRequest) {
    request = &ModifyTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyTokenPlanApiKey")
    
    
    return
}

func NewModifyTokenPlanApiKeyResponse() (response *ModifyTokenPlanApiKeyResponse) {
    response = &ModifyTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTokenPlanApiKey
// 修改 TokenPlan APIKey 配置（网关关注字段）。
//
// 
//
// 修改后自动通知网关更新缓存并同步额度中心。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKey(request *ModifyTokenPlanApiKeyRequest) (response *ModifyTokenPlanApiKeyResponse, err error) {
    return c.ModifyTokenPlanApiKeyWithContext(context.Background(), request)
}

// ModifyTokenPlanApiKey
// 修改 TokenPlan APIKey 配置（网关关注字段）。
//
// 
//
// 修改后自动通知网关更新缓存并同步额度中心。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeyWithContext(ctx context.Context, request *ModifyTokenPlanApiKeyRequest) (response *ModifyTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewModifyTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTokenPlanApiKeySecretRequest() (request *ModifyTokenPlanApiKeySecretRequest) {
    request = &ModifyTokenPlanApiKeySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyTokenPlanApiKeySecret")
    
    
    return
}

func NewModifyTokenPlanApiKeySecretResponse() (response *ModifyTokenPlanApiKeySecretResponse) {
    response = &ModifyTokenPlanApiKeySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTokenPlanApiKeySecret
// 重置 TokenPlan APIKey 密钥。
//
// 
//
// 重新生成密钥值，密钥版本号递增，旧密钥立即失效。APIKey ID 不变。重置后需通过 DescribeTokenPlanApiKeySecret 查询新密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeySecret(request *ModifyTokenPlanApiKeySecretRequest) (response *ModifyTokenPlanApiKeySecretResponse, err error) {
    return c.ModifyTokenPlanApiKeySecretWithContext(context.Background(), request)
}

// ModifyTokenPlanApiKeySecret
// 重置 TokenPlan APIKey 密钥。
//
// 
//
// 重新生成密钥值，密钥版本号递增，旧密钥立即失效。APIKey ID 不变。重置后需通过 DescribeTokenPlanApiKeySecret 查询新密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeySecretWithContext(ctx context.Context, request *ModifyTokenPlanApiKeySecretRequest) (response *ModifyTokenPlanApiKeySecretResponse, err error) {
    if request == nil {
        request = NewModifyTokenPlanApiKeySecretRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyTokenPlanApiKeySecret")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTokenPlanApiKeySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTokenPlanApiKeySecretResponse()
    err = c.Send(request, response)
    return
}

func NewRenewTokenPlanTeamOrderRequest() (request *RenewTokenPlanTeamOrderRequest) {
    request = &RenewTokenPlanTeamOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "RenewTokenPlanTeamOrder")
    
    
    return
}

func NewRenewTokenPlanTeamOrderResponse() (response *RenewTokenPlanTeamOrderResponse) {
    response = &RenewTokenPlanTeamOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewTokenPlanTeamOrder
// 续费套餐。
//
// 
//
// 对已有的 TokenPlan 套餐发起续费下单并完成支付，成功后返回大订单 ID 及关联的子订单、资源信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RenewTokenPlanTeamOrder(request *RenewTokenPlanTeamOrderRequest) (response *RenewTokenPlanTeamOrderResponse, err error) {
    return c.RenewTokenPlanTeamOrderWithContext(context.Background(), request)
}

// RenewTokenPlanTeamOrder
// 续费套餐。
//
// 
//
// 对已有的 TokenPlan 套餐发起续费下单并完成支付，成功后返回大订单 ID 及关联的子订单、资源信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RenewTokenPlanTeamOrderWithContext(ctx context.Context, request *RenewTokenPlanTeamOrderRequest) (response *RenewTokenPlanTeamOrderResponse, err error) {
    if request == nil {
        request = NewRenewTokenPlanTeamOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "RenewTokenPlanTeamOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewTokenPlanTeamOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewTokenPlanTeamOrderResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeTokenPlanTeamOrderRequest() (request *UpgradeTokenPlanTeamOrderRequest) {
    request = &UpgradeTokenPlanTeamOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "UpgradeTokenPlanTeamOrder")
    
    
    return
}

func NewUpgradeTokenPlanTeamOrderResponse() (response *UpgradeTokenPlanTeamOrderResponse) {
    response = &UpgradeTokenPlanTeamOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeTokenPlanTeamOrder
// 升配套餐。
//
// 
//
// 对已有的 TokenPlan 套餐发起升配下单并完成支付，扩容积分或 Token 额度，成功后返回大订单 ID 及关联的子订单、资源信息。新额度必须大于当前额度。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpgradeTokenPlanTeamOrder(request *UpgradeTokenPlanTeamOrderRequest) (response *UpgradeTokenPlanTeamOrderResponse, err error) {
    return c.UpgradeTokenPlanTeamOrderWithContext(context.Background(), request)
}

// UpgradeTokenPlanTeamOrder
// 升配套餐。
//
// 
//
// 对已有的 TokenPlan 套餐发起升配下单并完成支付，扩容积分或 Token 额度，成功后返回大订单 ID 及关联的子订单、资源信息。新额度必须大于当前额度。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpgradeTokenPlanTeamOrderWithContext(ctx context.Context, request *UpgradeTokenPlanTeamOrderRequest) (response *UpgradeTokenPlanTeamOrderResponse, err error) {
    if request == nil {
        request = NewUpgradeTokenPlanTeamOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "UpgradeTokenPlanTeamOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeTokenPlanTeamOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeTokenPlanTeamOrderResponse()
    err = c.Send(request, response)
    return
}
