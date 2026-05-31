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
// 从 CLS 日志服务查询套餐下的调用明细，按 pkg_id 过滤，支持游标分页。
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
// 从 CLS 日志服务查询套餐下的调用明细，按 pkg_id 过滤，支持游标分页。
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
// 按 apikey、endpoint、model 三个维度统计指定时间窗内的用量排行，返回顶部数据卡所需的 PageStats/TotalStats、左侧 Top 列表（含每对象整段累计值）、右侧色块趋势图所需的逐点曲线。前端通过 Offset 翻页、ShowAll 切换 CSV 全量导出模式。
//
// 
//
// MetricType 字段用于切换指标族，本期支持 tokens；接口预留以支持后续指标族扩展。响应回显 MetricType 与 MetricKeys（实际参与渲染的 metric key 列表，顺序固定 [Total, Input, Output]），前端按此渲染顶部数据卡与趋势图，无需硬编码 key 名。
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
// 按 apikey、endpoint、model 三个维度统计指定时间窗内的用量排行，返回顶部数据卡所需的 PageStats/TotalStats、左侧 Top 列表（含每对象整段累计值）、右侧色块趋势图所需的逐点曲线。前端通过 Offset 翻页、ShowAll 切换 CSV 全量导出模式。
//
// 
//
// MetricType 字段用于切换指标族，本期支持 tokens；接口预留以支持后续指标族扩展。响应回显 MetricType 与 MetricKeys（实际参与渲染的 metric key 列表，顺序固定 [Total, Input, Output]），前端按此渲染顶部数据卡与趋势图，无需硬编码 key 名。
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
