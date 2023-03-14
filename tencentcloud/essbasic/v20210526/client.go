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

package v20210526

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-26"

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


func NewChannelBatchCancelFlowsRequest() (request *ChannelBatchCancelFlowsRequest) {
    request = &ChannelBatchCancelFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelBatchCancelFlows")
    
    
    return
}

func NewChannelBatchCancelFlowsResponse() (response *ChannelBatchCancelFlowsResponse) {
    response = &ChannelBatchCancelFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelBatchCancelFlows
// 指定需要批量撤销的签署流程Id，批量撤销合同
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口失败后返回错误信息
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelBatchCancelFlows(request *ChannelBatchCancelFlowsRequest) (response *ChannelBatchCancelFlowsResponse, err error) {
    return c.ChannelBatchCancelFlowsWithContext(context.Background(), request)
}

// ChannelBatchCancelFlows
// 指定需要批量撤销的签署流程Id，批量撤销合同
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口失败后返回错误信息
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelBatchCancelFlowsWithContext(ctx context.Context, request *ChannelBatchCancelFlowsRequest) (response *ChannelBatchCancelFlowsResponse, err error) {
    if request == nil {
        request = NewChannelBatchCancelFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelBatchCancelFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelBatchCancelFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCancelFlowRequest() (request *ChannelCancelFlowRequest) {
    request = &ChannelCancelFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCancelFlow")
    
    
    return
}

func NewChannelCancelFlowResponse() (response *ChannelCancelFlowResponse) {
    response = &ChannelCancelFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCancelFlow
// 撤销签署流程接口，可以撤回：未全部签署完成；不可以撤回（终态）：已全部签署完成、已拒签、已过期、已撤回。
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCancelFlow(request *ChannelCancelFlowRequest) (response *ChannelCancelFlowResponse, err error) {
    return c.ChannelCancelFlowWithContext(context.Background(), request)
}

// ChannelCancelFlow
// 撤销签署流程接口，可以撤回：未全部签署完成；不可以撤回（终态）：已全部签署完成、已拒签、已过期、已撤回。
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCancelFlowWithContext(ctx context.Context, request *ChannelCancelFlowRequest) (response *ChannelCancelFlowResponse, err error) {
    if request == nil {
        request = NewChannelCancelFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCancelFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCancelFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCancelMultiFlowSignQRCodeRequest() (request *ChannelCancelMultiFlowSignQRCodeRequest) {
    request = &ChannelCancelMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCancelMultiFlowSignQRCode")
    
    
    return
}

func NewChannelCancelMultiFlowSignQRCodeResponse() (response *ChannelCancelMultiFlowSignQRCodeResponse) {
    response = &ChannelCancelMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCancelMultiFlowSignQRCode
// 此接口（ChannelCancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCancelMultiFlowSignQRCode(request *ChannelCancelMultiFlowSignQRCodeRequest) (response *ChannelCancelMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCancelMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCancelMultiFlowSignQRCode
// 此接口（ChannelCancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCancelMultiFlowSignQRCodeWithContext(ctx context.Context, request *ChannelCancelMultiFlowSignQRCodeRequest) (response *ChannelCancelMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewChannelCancelMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCancelMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCancelMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBatchCancelFlowUrlRequest() (request *ChannelCreateBatchCancelFlowUrlRequest) {
    request = &ChannelCreateBatchCancelFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBatchCancelFlowUrl")
    
    
    return
}

func NewChannelCreateBatchCancelFlowUrlResponse() (response *ChannelCreateBatchCancelFlowUrlResponse) {
    response = &ChannelCreateBatchCancelFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateBatchCancelFlowUrl
// 指定需要批量撤销的签署流程Id，获取批量撤销链接 - 不建议使用此接口，可使用ChannelBatchCancelFlows
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
//
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
//
// 可以撤回：未全部签署完成；不可以撤回（终态）：已全部签署完成、已拒签、已过期、已撤回。
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchCancelFlowUrl(request *ChannelCreateBatchCancelFlowUrlRequest) (response *ChannelCreateBatchCancelFlowUrlResponse, err error) {
    return c.ChannelCreateBatchCancelFlowUrlWithContext(context.Background(), request)
}

// ChannelCreateBatchCancelFlowUrl
// 指定需要批量撤销的签署流程Id，获取批量撤销链接 - 不建议使用此接口，可使用ChannelBatchCancelFlows
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
//
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
//
// 可以撤回：未全部签署完成；不可以撤回（终态）：已全部签署完成、已拒签、已过期、已撤回。
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchCancelFlowUrlWithContext(ctx context.Context, request *ChannelCreateBatchCancelFlowUrlRequest) (response *ChannelCreateBatchCancelFlowUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateBatchCancelFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBatchCancelFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBatchCancelFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBoundFlowsRequest() (request *ChannelCreateBoundFlowsRequest) {
    request = &ChannelCreateBoundFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBoundFlows")
    
    
    return
}

func NewChannelCreateBoundFlowsResponse() (response *ChannelCreateBoundFlowsResponse) {
    response = &ChannelCreateBoundFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于子客领取合同，经办人需要有相应的角色，合同不能重复领取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateBoundFlows(request *ChannelCreateBoundFlowsRequest) (response *ChannelCreateBoundFlowsResponse, err error) {
    return c.ChannelCreateBoundFlowsWithContext(context.Background(), request)
}

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于子客领取合同，经办人需要有相应的角色，合同不能重复领取。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateBoundFlowsWithContext(ctx context.Context, request *ChannelCreateBoundFlowsRequest) (response *ChannelCreateBoundFlowsResponse, err error) {
    if request == nil {
        request = NewChannelCreateBoundFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBoundFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBoundFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateConvertTaskApiRequest() (request *ChannelCreateConvertTaskApiRequest) {
    request = &ChannelCreateConvertTaskApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateConvertTaskApi")
    
    
    return
}

func NewChannelCreateConvertTaskApiResponse() (response *ChannelCreateConvertTaskApiResponse) {
    response = &ChannelCreateConvertTaskApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateConvertTaskApi
// 创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateConvertTaskApi(request *ChannelCreateConvertTaskApiRequest) (response *ChannelCreateConvertTaskApiResponse, err error) {
    return c.ChannelCreateConvertTaskApiWithContext(context.Background(), request)
}

// ChannelCreateConvertTaskApi
// 创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateConvertTaskApiWithContext(ctx context.Context, request *ChannelCreateConvertTaskApiRequest) (response *ChannelCreateConvertTaskApiResponse, err error) {
    if request == nil {
        request = NewChannelCreateConvertTaskApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateConvertTaskApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateConvertTaskApiResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateEmbedWebUrlRequest() (request *ChannelCreateEmbedWebUrlRequest) {
    request = &ChannelCreateEmbedWebUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateEmbedWebUrl")
    
    
    return
}

func NewChannelCreateEmbedWebUrlResponse() (response *ChannelCreateEmbedWebUrlResponse) {
    response = &ChannelCreateEmbedWebUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateEmbedWebUrl
// 本接口（ChannelCreateEmbedWebUrl）用于创建嵌入web的链接
//
// 本接口支持创建：创建印章，创建模板，修改模板，预览模板，预览合同流程的web链接
//
// 进入web连接后与当前控制台操作保持一致
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateEmbedWebUrl(request *ChannelCreateEmbedWebUrlRequest) (response *ChannelCreateEmbedWebUrlResponse, err error) {
    return c.ChannelCreateEmbedWebUrlWithContext(context.Background(), request)
}

// ChannelCreateEmbedWebUrl
// 本接口（ChannelCreateEmbedWebUrl）用于创建嵌入web的链接
//
// 本接口支持创建：创建印章，创建模板，修改模板，预览模板，预览合同流程的web链接
//
// 进入web连接后与当前控制台操作保持一致
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateEmbedWebUrlWithContext(ctx context.Context, request *ChannelCreateEmbedWebUrlRequest) (response *ChannelCreateEmbedWebUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateEmbedWebUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateEmbedWebUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateEmbedWebUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowByFilesRequest() (request *ChannelCreateFlowByFilesRequest) {
    request = &ChannelCreateFlowByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowByFiles")
    
    
    return
}

func NewChannelCreateFlowByFilesResponse() (response *ChannelCreateFlowByFilesResponse) {
    response = &ChannelCreateFlowByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateFlowByFiles
// 接口（ChannelCreateFlowByFiles）用于通过文件创建签署流程。此接口静默签能力不可直接使用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  MISSINGPARAMETER_SEALID = "MissingParameter.SealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowByFiles(request *ChannelCreateFlowByFilesRequest) (response *ChannelCreateFlowByFilesResponse, err error) {
    return c.ChannelCreateFlowByFilesWithContext(context.Background(), request)
}

// ChannelCreateFlowByFiles
// 接口（ChannelCreateFlowByFiles）用于通过文件创建签署流程。此接口静默签能力不可直接使用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  MISSINGPARAMETER_SEALID = "MissingParameter.SealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowByFilesWithContext(ctx context.Context, request *ChannelCreateFlowByFilesRequest) (response *ChannelCreateFlowByFilesResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowGroupByFilesRequest() (request *ChannelCreateFlowGroupByFilesRequest) {
    request = &ChannelCreateFlowGroupByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowGroupByFiles")
    
    
    return
}

func NewChannelCreateFlowGroupByFilesResponse() (response *ChannelCreateFlowGroupByFilesResponse) {
    response = &ChannelCreateFlowGroupByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateFlowGroupByFiles
// 接口（ChannelCreateFlowGroupByFiles）用于通过多文件创建合同组签署流程。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByFiles(request *ChannelCreateFlowGroupByFilesRequest) (response *ChannelCreateFlowGroupByFilesResponse, err error) {
    return c.ChannelCreateFlowGroupByFilesWithContext(context.Background(), request)
}

// ChannelCreateFlowGroupByFiles
// 接口（ChannelCreateFlowGroupByFiles）用于通过多文件创建合同组签署流程。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByFilesWithContext(ctx context.Context, request *ChannelCreateFlowGroupByFilesRequest) (response *ChannelCreateFlowGroupByFilesResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowGroupByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowGroupByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowGroupByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowRemindsRequest() (request *ChannelCreateFlowRemindsRequest) {
    request = &ChannelCreateFlowRemindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowReminds")
    
    
    return
}

func NewChannelCreateFlowRemindsResponse() (response *ChannelCreateFlowRemindsResponse) {
    response = &ChannelCreateFlowRemindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateFlowReminds
// 指定需要批量催办的签署流程Id，批量催办合同，最多100个；接口失败后返回错误信息
//
// 注意:
//
// 该接口不可直接调用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowReminds(request *ChannelCreateFlowRemindsRequest) (response *ChannelCreateFlowRemindsResponse, err error) {
    return c.ChannelCreateFlowRemindsWithContext(context.Background(), request)
}

// ChannelCreateFlowReminds
// 指定需要批量催办的签署流程Id，批量催办合同，最多100个；接口失败后返回错误信息
//
// 注意:
//
// 该接口不可直接调用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowRemindsWithContext(ctx context.Context, request *ChannelCreateFlowRemindsRequest) (response *ChannelCreateFlowRemindsResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowRemindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowReminds require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowRemindsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowSignReviewRequest() (request *ChannelCreateFlowSignReviewRequest) {
    request = &ChannelCreateFlowSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowSignReview")
    
    
    return
}

func NewChannelCreateFlowSignReviewResponse() (response *ChannelCreateFlowSignReviewResponse) {
    response = &ChannelCreateFlowSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateFlowSignReview
// 提交企业签署流程审批结果
//
// 
//
// 在通过接口(CreateFlowsByTemplates 或者ChannelCreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateFlowSignReview(request *ChannelCreateFlowSignReviewRequest) (response *ChannelCreateFlowSignReviewResponse, err error) {
    return c.ChannelCreateFlowSignReviewWithContext(context.Background(), request)
}

// ChannelCreateFlowSignReview
// 提交企业签署流程审批结果
//
// 
//
// 在通过接口(CreateFlowsByTemplates 或者ChannelCreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateFlowSignReviewWithContext(ctx context.Context, request *ChannelCreateFlowSignReviewRequest) (response *ChannelCreateFlowSignReviewResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowSignUrlRequest() (request *ChannelCreateFlowSignUrlRequest) {
    request = &ChannelCreateFlowSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowSignUrl")
    
    
    return
}

func NewChannelCreateFlowSignUrlResponse() (response *ChannelCreateFlowSignUrlResponse) {
    response = &ChannelCreateFlowSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateFlowSignUrl
// 创建签署链接，请联系客户经理申请使用
//
// 该接口用于发起合同后，生成C端签署人的签署链接，点击跳转小程序完成签署
//
// 注意：该接口目前签署人类型仅支持个人签署方（PERSON）
//
// 注意：该接口可生成签署链接的C端签署人必须仅有手写签名和时间类型的签署控件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowSignUrl(request *ChannelCreateFlowSignUrlRequest) (response *ChannelCreateFlowSignUrlResponse, err error) {
    return c.ChannelCreateFlowSignUrlWithContext(context.Background(), request)
}

// ChannelCreateFlowSignUrl
// 创建签署链接，请联系客户经理申请使用
//
// 该接口用于发起合同后，生成C端签署人的签署链接，点击跳转小程序完成签署
//
// 注意：该接口目前签署人类型仅支持个人签署方（PERSON）
//
// 注意：该接口可生成签署链接的C端签署人必须仅有手写签名和时间类型的签署控件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowSignUrlWithContext(ctx context.Context, request *ChannelCreateFlowSignUrlRequest) (response *ChannelCreateFlowSignUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateMultiFlowSignQRCodeRequest() (request *ChannelCreateMultiFlowSignQRCodeRequest) {
    request = &ChannelCreateMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateMultiFlowSignQRCode")
    
    
    return
}

func NewChannelCreateMultiFlowSignQRCodeResponse() (response *ChannelCreateMultiFlowSignQRCodeResponse) {
    response = &ChannelCreateMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateMultiFlowSignQRCode
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫签署流程二维码。
//
// 适用的模版仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模版，且模版中发起方没有填写控件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateMultiFlowSignQRCode(request *ChannelCreateMultiFlowSignQRCodeRequest) (response *ChannelCreateMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCreateMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCreateMultiFlowSignQRCode
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫签署流程二维码。
//
// 适用的模版仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模版，且模版中发起方没有填写控件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateMultiFlowSignQRCodeWithContext(ctx context.Context, request *ChannelCreateMultiFlowSignQRCodeRequest) (response *ChannelCreateMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewChannelCreateMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreatePrepareFlowRequest() (request *ChannelCreatePrepareFlowRequest) {
    request = &ChannelCreatePrepareFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreatePrepareFlow")
    
    
    return
}

func NewChannelCreatePrepareFlowResponse() (response *ChannelCreatePrepareFlowResponse) {
    response = &ChannelCreatePrepareFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreatePrepareFlow
// 创建预发起合同
//
// 通过此接口指定：合同，签署人，填写控件信息，生成预创建合同链接，点击后跳转到web页面完成合同创建并发起
//
// 可指定合同信息不可更改，签署人信息不可更改
//
// 合同发起后，填写及签署流程与现有操作流程一致
//
// 注意：目前仅支持模板发起
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreatePrepareFlow(request *ChannelCreatePrepareFlowRequest) (response *ChannelCreatePrepareFlowResponse, err error) {
    return c.ChannelCreatePrepareFlowWithContext(context.Background(), request)
}

// ChannelCreatePrepareFlow
// 创建预发起合同
//
// 通过此接口指定：合同，签署人，填写控件信息，生成预创建合同链接，点击后跳转到web页面完成合同创建并发起
//
// 可指定合同信息不可更改，签署人信息不可更改
//
// 合同发起后，填写及签署流程与现有操作流程一致
//
// 注意：目前仅支持模板发起
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreatePrepareFlowWithContext(ctx context.Context, request *ChannelCreatePrepareFlowRequest) (response *ChannelCreatePrepareFlowResponse, err error) {
    if request == nil {
        request = NewChannelCreatePrepareFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreatePrepareFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreatePrepareFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateReleaseFlowRequest() (request *ChannelCreateReleaseFlowRequest) {
    request = &ChannelCreateReleaseFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateReleaseFlow")
    
    
    return
}

func NewChannelCreateReleaseFlowResponse() (response *ChannelCreateReleaseFlowResponse) {
    response = &ChannelCreateReleaseFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateReleaseFlow
// 发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
//
// 合同发起人必须在电子签已经进行实名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateReleaseFlow(request *ChannelCreateReleaseFlowRequest) (response *ChannelCreateReleaseFlowResponse, err error) {
    return c.ChannelCreateReleaseFlowWithContext(context.Background(), request)
}

// ChannelCreateReleaseFlow
// 发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
//
// 合同发起人必须在电子签已经进行实名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateReleaseFlowWithContext(ctx context.Context, request *ChannelCreateReleaseFlowRequest) (response *ChannelCreateReleaseFlowResponse, err error) {
    if request == nil {
        request = NewChannelCreateReleaseFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateReleaseFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateReleaseFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateSealPolicyRequest() (request *ChannelCreateSealPolicyRequest) {
    request = &ChannelCreateSealPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateSealPolicy")
    
    
    return
}

func NewChannelCreateSealPolicyResponse() (response *ChannelCreateSealPolicyResponse) {
    response = &ChannelCreateSealPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelCreateSealPolicy
// 将指定印章授权给第三方平台子客企业下的某些员工
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelCreateSealPolicy(request *ChannelCreateSealPolicyRequest) (response *ChannelCreateSealPolicyResponse, err error) {
    return c.ChannelCreateSealPolicyWithContext(context.Background(), request)
}

// ChannelCreateSealPolicy
// 将指定印章授权给第三方平台子客企业下的某些员工
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelCreateSealPolicyWithContext(ctx context.Context, request *ChannelCreateSealPolicyRequest) (response *ChannelCreateSealPolicyResponse, err error) {
    if request == nil {
        request = NewChannelCreateSealPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateSealPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateSealPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDeleteSealPoliciesRequest() (request *ChannelDeleteSealPoliciesRequest) {
    request = &ChannelDeleteSealPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDeleteSealPolicies")
    
    
    return
}

func NewChannelDeleteSealPoliciesResponse() (response *ChannelDeleteSealPoliciesResponse) {
    response = &ChannelDeleteSealPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelDeleteSealPolicies
// 删除指定印章下多个授权信息
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelDeleteSealPolicies(request *ChannelDeleteSealPoliciesRequest) (response *ChannelDeleteSealPoliciesResponse, err error) {
    return c.ChannelDeleteSealPoliciesWithContext(context.Background(), request)
}

// ChannelDeleteSealPolicies
// 删除指定印章下多个授权信息
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelDeleteSealPoliciesWithContext(ctx context.Context, request *ChannelDeleteSealPoliciesRequest) (response *ChannelDeleteSealPoliciesResponse, err error) {
    if request == nil {
        request = NewChannelDeleteSealPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDeleteSealPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDeleteSealPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeEmployeesRequest() (request *ChannelDescribeEmployeesRequest) {
    request = &ChannelDescribeEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeEmployees")
    
    
    return
}

func NewChannelDescribeEmployeesResponse() (response *ChannelDescribeEmployeesResponse) {
    response = &ChannelDescribeEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelDescribeEmployees
// 查询企业员工列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeEmployees(request *ChannelDescribeEmployeesRequest) (response *ChannelDescribeEmployeesResponse, err error) {
    return c.ChannelDescribeEmployeesWithContext(context.Background(), request)
}

// ChannelDescribeEmployees
// 查询企业员工列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeEmployeesWithContext(ctx context.Context, request *ChannelDescribeEmployeesRequest) (response *ChannelDescribeEmployeesResponse, err error) {
    if request == nil {
        request = NewChannelDescribeEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeOrganizationSealsRequest() (request *ChannelDescribeOrganizationSealsRequest) {
    request = &ChannelDescribeOrganizationSealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeOrganizationSeals")
    
    
    return
}

func NewChannelDescribeOrganizationSealsResponse() (response *ChannelDescribeOrganizationSealsResponse) {
    response = &ChannelDescribeOrganizationSealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelDescribeOrganizationSeals
// 查询子客企业电子印章，需要操作者具有管理印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeOrganizationSeals(request *ChannelDescribeOrganizationSealsRequest) (response *ChannelDescribeOrganizationSealsResponse, err error) {
    return c.ChannelDescribeOrganizationSealsWithContext(context.Background(), request)
}

// ChannelDescribeOrganizationSeals
// 查询子客企业电子印章，需要操作者具有管理印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeOrganizationSealsWithContext(ctx context.Context, request *ChannelDescribeOrganizationSealsRequest) (response *ChannelDescribeOrganizationSealsResponse, err error) {
    if request == nil {
        request = NewChannelDescribeOrganizationSealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeOrganizationSeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeOrganizationSealsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelGetTaskResultApiRequest() (request *ChannelGetTaskResultApiRequest) {
    request = &ChannelGetTaskResultApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelGetTaskResultApi")
    
    
    return
}

func NewChannelGetTaskResultApiResponse() (response *ChannelGetTaskResultApiResponse) {
    response = &ChannelGetTaskResultApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelGetTaskResultApi
// 查询转换任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelGetTaskResultApi(request *ChannelGetTaskResultApiRequest) (response *ChannelGetTaskResultApiResponse, err error) {
    return c.ChannelGetTaskResultApiWithContext(context.Background(), request)
}

// ChannelGetTaskResultApi
// 查询转换任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelGetTaskResultApiWithContext(ctx context.Context, request *ChannelGetTaskResultApiRequest) (response *ChannelGetTaskResultApiResponse, err error) {
    if request == nil {
        request = NewChannelGetTaskResultApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelGetTaskResultApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelGetTaskResultApiResponse()
    err = c.Send(request, response)
    return
}

func NewChannelUpdateSealStatusRequest() (request *ChannelUpdateSealStatusRequest) {
    request = &ChannelUpdateSealStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelUpdateSealStatus")
    
    
    return
}

func NewChannelUpdateSealStatusResponse() (response *ChannelUpdateSealStatusResponse) {
    response = &ChannelUpdateSealStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelUpdateSealStatus
// 本接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ChannelUpdateSealStatus(request *ChannelUpdateSealStatusRequest) (response *ChannelUpdateSealStatusResponse, err error) {
    return c.ChannelUpdateSealStatusWithContext(context.Background(), request)
}

// ChannelUpdateSealStatus
// 本接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ChannelUpdateSealStatusWithContext(ctx context.Context, request *ChannelUpdateSealStatusRequest) (response *ChannelUpdateSealStatusResponse, err error) {
    if request == nil {
        request = NewChannelUpdateSealStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelUpdateSealStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelUpdateSealStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChannelVerifyPdfRequest() (request *ChannelVerifyPdfRequest) {
    request = &ChannelVerifyPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelVerifyPdf")
    
    
    return
}

func NewChannelVerifyPdfResponse() (response *ChannelVerifyPdfResponse) {
    response = &ChannelVerifyPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChannelVerifyPdf
// 合同文件验签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelVerifyPdf(request *ChannelVerifyPdfRequest) (response *ChannelVerifyPdfResponse, err error) {
    return c.ChannelVerifyPdfWithContext(context.Background(), request)
}

// ChannelVerifyPdf
// 合同文件验签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelVerifyPdfWithContext(ctx context.Context, request *ChannelVerifyPdfRequest) (response *ChannelVerifyPdfResponse, err error) {
    if request == nil {
        request = NewChannelVerifyPdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelVerifyPdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelVerifyPdfResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelFlowEvidenceReportRequest() (request *CreateChannelFlowEvidenceReportRequest) {
    request = &CreateChannelFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateChannelFlowEvidenceReport")
    
    
    return
}

func NewCreateChannelFlowEvidenceReportResponse() (response *CreateChannelFlowEvidenceReportResponse) {
    response = &CreateChannelFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateChannelFlowEvidenceReport
// 创建出证报告，返回报告 ID。需要配合出证套餐才能调用。
//
// 出证需要一定时间，建议调用创建出证24小时之后再通过DescribeChannelFlowEvidenceReport进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelFlowEvidenceReport(request *CreateChannelFlowEvidenceReportRequest) (response *CreateChannelFlowEvidenceReportResponse, err error) {
    return c.CreateChannelFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateChannelFlowEvidenceReport
// 创建出证报告，返回报告 ID。需要配合出证套餐才能调用。
//
// 出证需要一定时间，建议调用创建出证24小时之后再通过DescribeChannelFlowEvidenceReport进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelFlowEvidenceReportWithContext(ctx context.Context, request *CreateChannelFlowEvidenceReportRequest) (response *CreateChannelFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewCreateChannelFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannelFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleLoginUrlRequest() (request *CreateConsoleLoginUrlRequest) {
    request = &CreateConsoleLoginUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateConsoleLoginUrl")
    
    
    return
}

func NewCreateConsoleLoginUrlResponse() (response *CreateConsoleLoginUrlResponse) {
    response = &CreateConsoleLoginUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConsoleLoginUrl
// 此接口（CreateConsoleLoginUrl）用于创建第三方平台子客企业控制台Web/移动登录链接。登录链接是子客控制台的唯一入口。
//
// 若子客企业未激活，会进入企业激活流程，首次参与激活流程的经办人会成为超管。（若企业激活过程中填写信息有误，需要重置激活流程，可以换一个经办人OpenId获取新的链接进入。）
//
// 若子客企业已激活，使用了新的经办人OpenId进入，则会进入经办人的实名流程。
//
// 若子客企业、经办人均已完成认证，则会直接进入子客Web控制台。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrl(request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    return c.CreateConsoleLoginUrlWithContext(context.Background(), request)
}

// CreateConsoleLoginUrl
// 此接口（CreateConsoleLoginUrl）用于创建第三方平台子客企业控制台Web/移动登录链接。登录链接是子客控制台的唯一入口。
//
// 若子客企业未激活，会进入企业激活流程，首次参与激活流程的经办人会成为超管。（若企业激活过程中填写信息有误，需要重置激活流程，可以换一个经办人OpenId获取新的链接进入。）
//
// 若子客企业已激活，使用了新的经办人OpenId进入，则会进入经办人的实名流程。
//
// 若子客企业、经办人均已完成认证，则会直接进入子客Web控制台。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrlWithContext(ctx context.Context, request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    if request == nil {
        request = NewCreateConsoleLoginUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleLoginUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleLoginUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowsByTemplatesRequest() (request *CreateFlowsByTemplatesRequest) {
    request = &CreateFlowsByTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFlowsByTemplates")
    
    
    return
}

func NewCreateFlowsByTemplatesResponse() (response *CreateFlowsByTemplatesResponse) {
    response = &CreateFlowsByTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowsByTemplates
// 接口（CreateFlowsByTemplates）用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
//
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小
//
// 合同发起人必须在电子签已经进行实名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowsByTemplates(request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    return c.CreateFlowsByTemplatesWithContext(context.Background(), request)
}

// CreateFlowsByTemplates
// 接口（CreateFlowsByTemplates）用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
//
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小
//
// 合同发起人必须在电子签已经进行实名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowsByTemplatesWithContext(ctx context.Context, request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    if request == nil {
        request = NewCreateFlowsByTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowsByTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowsByTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealByImageRequest() (request *CreateSealByImageRequest) {
    request = &CreateSealByImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSealByImage")
    
    
    return
}

func NewCreateSealByImageResponse() (response *CreateSealByImageResponse) {
    response = &CreateSealByImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSealByImage
// 通过图片为子客企业代创建印章，图片最大5MB
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSealByImage(request *CreateSealByImageRequest) (response *CreateSealByImageResponse, err error) {
    return c.CreateSealByImageWithContext(context.Background(), request)
}

// CreateSealByImage
// 通过图片为子客企业代创建印章，图片最大5MB
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSealByImageWithContext(ctx context.Context, request *CreateSealByImageRequest) (response *CreateSealByImageResponse, err error) {
    if request == nil {
        request = NewCreateSealByImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSealByImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealByImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignUrlsRequest() (request *CreateSignUrlsRequest) {
    request = &CreateSignUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSignUrls")
    
    
    return
}

func NewCreateSignUrlsResponse() (response *CreateSignUrlsResponse) {
    response = &CreateSignUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSignUrls
// 创建跳转小程序查看或签署的链接。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。（与短信提醒用户签署形式一样）。
//
// 
//
// 2. 通过链接Url打开H5引导页-->点击跳转到小程序-->签署完退出小程序-->回到H5引导页-->跳转到指定JumpUrl
//
// 设置EndPoint为CHANNEL，指定JumpUrl，得到链接打开即可。
//
// 
//
// 3. 客户App直接跳转到小程序-->小程序签署完成-->返回App
//
// 跳转到小程序的实现，参考官方文档
//
// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 使用CreateSignUrls，设置EndPoint为APP，得到path。
//
// 
//
// 4. 客户小程序直接跳到电子签小程序-->签署完成退出电子签小程序-->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为全屏、半屏两种方式）
//
// 全屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html）
//
// 半屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html）
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 使用CreateSignUrls，设置EndPoint为APP，得到path。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSignUrls(request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    return c.CreateSignUrlsWithContext(context.Background(), request)
}

// CreateSignUrls
// 创建跳转小程序查看或签署的链接。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。（与短信提醒用户签署形式一样）。
//
// 
//
// 2. 通过链接Url打开H5引导页-->点击跳转到小程序-->签署完退出小程序-->回到H5引导页-->跳转到指定JumpUrl
//
// 设置EndPoint为CHANNEL，指定JumpUrl，得到链接打开即可。
//
// 
//
// 3. 客户App直接跳转到小程序-->小程序签署完成-->返回App
//
// 跳转到小程序的实现，参考官方文档
//
// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 使用CreateSignUrls，设置EndPoint为APP，得到path。
//
// 
//
// 4. 客户小程序直接跳到电子签小程序-->签署完成退出电子签小程序-->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为全屏、半屏两种方式）
//
// 全屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html）
//
// 半屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html）
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 使用CreateSignUrls，设置EndPoint为APP，得到path。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSignUrlsWithContext(ctx context.Context, request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    if request == nil {
        request = NewCreateSignUrlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelFlowEvidenceReportRequest() (request *DescribeChannelFlowEvidenceReportRequest) {
    request = &DescribeChannelFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeChannelFlowEvidenceReport")
    
    
    return
}

func NewDescribeChannelFlowEvidenceReportResponse() (response *DescribeChannelFlowEvidenceReportResponse) {
    response = &DescribeChannelFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelFlowEvidenceReport
// 查询出证报告，返回报告 URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeChannelFlowEvidenceReport(request *DescribeChannelFlowEvidenceReportRequest) (response *DescribeChannelFlowEvidenceReportResponse, err error) {
    return c.DescribeChannelFlowEvidenceReportWithContext(context.Background(), request)
}

// DescribeChannelFlowEvidenceReport
// 查询出证报告，返回报告 URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeChannelFlowEvidenceReportWithContext(ctx context.Context, request *DescribeChannelFlowEvidenceReportRequest) (response *DescribeChannelFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewDescribeChannelFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtendedServiceAuthInfoRequest() (request *DescribeExtendedServiceAuthInfoRequest) {
    request = &DescribeExtendedServiceAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeExtendedServiceAuthInfo")
    
    
    return
}

func NewDescribeExtendedServiceAuthInfoResponse() (response *DescribeExtendedServiceAuthInfoResponse) {
    response = &DescribeExtendedServiceAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExtendedServiceAuthInfo
// 查询企业扩展服务授权信息，企业经办人需要是企业超管或者法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfo(request *DescribeExtendedServiceAuthInfoRequest) (response *DescribeExtendedServiceAuthInfoResponse, err error) {
    return c.DescribeExtendedServiceAuthInfoWithContext(context.Background(), request)
}

// DescribeExtendedServiceAuthInfo
// 查询企业扩展服务授权信息，企业经办人需要是企业超管或者法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfoWithContext(ctx context.Context, request *DescribeExtendedServiceAuthInfoRequest) (response *DescribeExtendedServiceAuthInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtendedServiceAuthInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtendedServiceAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtendedServiceAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowDetailInfoRequest() (request *DescribeFlowDetailInfoRequest) {
    request = &DescribeFlowDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFlowDetailInfo")
    
    
    return
}

func NewDescribeFlowDetailInfoResponse() (response *DescribeFlowDetailInfoResponse) {
    response = &DescribeFlowDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowDetailInfo
// 此接口（DescribeFlowDetailInfo）用于查询合同(签署流程)的详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowDetailInfo(request *DescribeFlowDetailInfoRequest) (response *DescribeFlowDetailInfoResponse, err error) {
    return c.DescribeFlowDetailInfoWithContext(context.Background(), request)
}

// DescribeFlowDetailInfo
// 此接口（DescribeFlowDetailInfo）用于查询合同(签署流程)的详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowDetailInfoWithContext(ctx context.Context, request *DescribeFlowDetailInfoRequest) (response *DescribeFlowDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFlowDetailInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUrlsByFlowsRequest() (request *DescribeResourceUrlsByFlowsRequest) {
    request = &DescribeResourceUrlsByFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeResourceUrlsByFlows")
    
    
    return
}

func NewDescribeResourceUrlsByFlowsResponse() (response *DescribeResourceUrlsByFlowsResponse) {
    response = &DescribeResourceUrlsByFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceUrlsByFlows
// 根据签署流程信息批量获取资源下载链接，可以下载签署中、签署完的合同，需合作企业先进行授权。
//
// 此接口直接返回下载的资源的url，与接口GetDownloadFlowUrl跳转到控制台的下载方式不同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_USEROPENID = "MissingParameter.UserOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
func (c *Client) DescribeResourceUrlsByFlows(request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    return c.DescribeResourceUrlsByFlowsWithContext(context.Background(), request)
}

// DescribeResourceUrlsByFlows
// 根据签署流程信息批量获取资源下载链接，可以下载签署中、签署完的合同，需合作企业先进行授权。
//
// 此接口直接返回下载的资源的url，与接口GetDownloadFlowUrl跳转到控制台的下载方式不同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_USEROPENID = "MissingParameter.UserOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
func (c *Client) DescribeResourceUrlsByFlowsWithContext(ctx context.Context, request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUrlsByFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUrlsByFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUrlsByFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
    request = &DescribeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeTemplates")
    
    
    return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
    response = &DescribeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTemplates
// 通过此接口（DescribeTemplates）查询该第三方平台子客企业在电子签拥有的有效模板，不包括第三方平台模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    return c.DescribeTemplatesWithContext(context.Background(), request)
}

// DescribeTemplates
// 通过此接口（DescribeTemplates）查询该第三方平台子客企业在电子签拥有的有效模板，不包括第三方平台模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageRequest() (request *DescribeUsageRequest) {
    request = &DescribeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeUsage")
    
    
    return
}

func NewDescribeUsageResponse() (response *DescribeUsageResponse) {
    response = &DescribeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsage
// 此接口（DescribeUsage）用于获取第三方平台所有合作企业流量消耗情况。
//
//  注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsage(request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    return c.DescribeUsageWithContext(context.Background(), request)
}

// DescribeUsage
// 此接口（DescribeUsage）用于获取第三方平台所有合作企业流量消耗情况。
//
//  注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsageWithContext(ctx context.Context, request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewGetDownloadFlowUrlRequest() (request *GetDownloadFlowUrlRequest) {
    request = &GetDownloadFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "GetDownloadFlowUrl")
    
    
    return
}

func NewGetDownloadFlowUrlResponse() (response *GetDownloadFlowUrlResponse) {
    response = &GetDownloadFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
//
// 当前接口限制最多合同（流程）50个.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWNUMEXCEED = "FailedOperation.FlowNumExceed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWINFO = "MissingParameter.FlowInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_DOWNLOADMORETHANONE = "OperationDenied.DownLoadMoreThanOne"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FILE = "ResourceNotFound.File"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_TEAMWORKORGANIZATION = "ResourceNotFound.TeamWorkOrganization"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrl(request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    return c.GetDownloadFlowUrlWithContext(context.Background(), request)
}

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
//
// 当前接口限制最多合同（流程）50个.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWNUMEXCEED = "FailedOperation.FlowNumExceed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWINFO = "MissingParameter.FlowInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_DOWNLOADMORETHANONE = "OperationDenied.DownLoadMoreThanOne"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FILE = "ResourceNotFound.File"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_TEAMWORKORGANIZATION = "ResourceNotFound.TeamWorkOrganization"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrlWithContext(ctx context.Context, request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    if request == nil {
        request = NewGetDownloadFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDownloadFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDownloadFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExtendedServiceRequest() (request *ModifyExtendedServiceRequest) {
    request = &ModifyExtendedServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyExtendedService")
    
    
    return
}

func NewModifyExtendedServiceResponse() (response *ModifyExtendedServiceResponse) {
    response = &ModifyExtendedServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyExtendedService
// 修改（操作）企业扩展服务 ，企业经办人需要是企业超管或者法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedService(request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    return c.ModifyExtendedServiceWithContext(context.Background(), request)
}

// ModifyExtendedService
// 修改（操作）企业扩展服务 ，企业经办人需要是企业超管或者法人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedServiceWithContext(ctx context.Context, request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    if request == nil {
        request = NewModifyExtendedServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExtendedService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExtendedServiceResponse()
    err = c.Send(request, response)
    return
}

func NewOperateChannelTemplateRequest() (request *OperateChannelTemplateRequest) {
    request = &OperateChannelTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "OperateChannelTemplate")
    
    
    return
}

func NewOperateChannelTemplateResponse() (response *OperateChannelTemplateResponse) {
    response = &OperateChannelTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业可见性的查询和设置，不会直接分配第三方应用平台模板给子客企业。
//
// 1、OperateType=select时：
//
// 查询第三方应用平台模板库
//
// 2、OperateType=update或者delete时：
//
// 对子客企业进行模板库中模板可见性的修改、删除操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATES = "MissingParameter.Templates"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTHTAG = "OperationDenied.AuthTag"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONAUTH = "ResourceNotFound.ApplicationAuth"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplate(request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    return c.OperateChannelTemplateWithContext(context.Background(), request)
}

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业可见性的查询和设置，不会直接分配第三方应用平台模板给子客企业。
//
// 1、OperateType=select时：
//
// 查询第三方应用平台模板库
//
// 2、OperateType=update或者delete时：
//
// 对子客企业进行模板库中模板可见性的修改、删除操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATES = "MissingParameter.Templates"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTHTAG = "OperationDenied.AuthTag"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONAUTH = "ResourceNotFound.ApplicationAuth"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplateWithContext(ctx context.Context, request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    if request == nil {
        request = NewOperateChannelTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateChannelTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateChannelTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPrepareFlowsRequest() (request *PrepareFlowsRequest) {
    request = &PrepareFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "PrepareFlows")
    
    
    return
}

func NewPrepareFlowsResponse() (response *PrepareFlowsResponse) {
    response = &PrepareFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
//
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
//
// 目前该接口只支持B2C，不建议使用，将会废弃。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) PrepareFlows(request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    return c.PrepareFlowsWithContext(context.Background(), request)
}

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
//
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
//
// 目前该接口只支持B2C，不建议使用，将会废弃。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) PrepareFlowsWithContext(ctx context.Context, request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    if request == nil {
        request = NewPrepareFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PrepareFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewPrepareFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewSyncProxyOrganizationRequest() (request *SyncProxyOrganizationRequest) {
    request = &SyncProxyOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "SyncProxyOrganization")
    
    
    return
}

func NewSyncProxyOrganizationResponse() (response *SyncProxyOrganizationResponse) {
    response = &SyncProxyOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncProxyOrganization
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，主要是子客企业的营业执照，便于子客企业开通过程中不用手动上传。若有需要调用此接口，需要在创建控制链接CreateConsoleLoginUrl之后即刻进行调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganization(request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    return c.SyncProxyOrganizationWithContext(context.Background(), request)
}

// SyncProxyOrganization
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，主要是子客企业的营业执照，便于子客企业开通过程中不用手动上传。若有需要调用此接口，需要在创建控制链接CreateConsoleLoginUrl之后即刻进行调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationWithContext(ctx context.Context, request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncProxyOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncProxyOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewSyncProxyOrganizationOperatorsRequest() (request *SyncProxyOrganizationOperatorsRequest) {
    request = &SyncProxyOrganizationOperatorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "SyncProxyOrganizationOperators")
    
    
    return
}

func NewSyncProxyOrganizationOperatorsResponse() (response *SyncProxyOrganizationOperatorsResponse) {
    response = &SyncProxyOrganizationOperatorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncProxyOrganizationOperators
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。子客Web控制台的组织架构管理，是依赖于第三方应用平台的，无法针对员工做新增/更新/离职等操作。
//
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_PROXYORGANIZATIONOPERATOR = "LimitExceeded.ProxyOrganizationOperator"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperators(request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    return c.SyncProxyOrganizationOperatorsWithContext(context.Background(), request)
}

// SyncProxyOrganizationOperators
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。子客Web控制台的组织架构管理，是依赖于第三方应用平台的，无法针对员工做新增/更新/离职等操作。
//
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_PROXYORGANIZATIONOPERATOR = "LimitExceeded.ProxyOrganizationOperator"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperatorsWithContext(ctx context.Context, request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationOperatorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncProxyOrganizationOperators require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncProxyOrganizationOperatorsResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFilesRequest() (request *UploadFilesRequest) {
    request = &UploadFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "UploadFiles")
    
    
    return
}

func NewUploadFilesResponse() (response *UploadFilesResponse) {
    response = &UploadFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 其中上传的文件，图片类型(png/jpg/jpeg)大小限制为5M，其他大小限制为60M。
//
// 调用时需要设置Domain, 正式环境为 file.ess.tencent.cn。
//
// 代码示例：
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 其中上传的文件，图片类型(png/jpg/jpeg)大小限制为5M，其他大小限制为60M。
//
// 调用时需要设置Domain, 正式环境为 file.ess.tencent.cn。
//
// 代码示例：
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFilesWithContext(ctx context.Context, request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    if request == nil {
        request = NewUploadFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFilesResponse()
    err = c.Send(request, response)
    return
}
