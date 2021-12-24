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
// 此接口（CreateConsoleLoginUrl）用于创建电子签控制台登录链接。若企业未激活，调用同步企业信息、同步经办人信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrl(request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    if request == nil {
        request = NewCreateConsoleLoginUrlRequest()
    }
    
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
// 接口（CreateFlowsByTemplates）用于使用多个模板批量创建流程
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
func (c *Client) CreateFlowsByTemplates(request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    if request == nil {
        request = NewCreateFlowsByTemplatesRequest()
    }
    
    response = NewCreateFlowsByTemplatesResponse()
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
// 根据流程Id批量创建签署参与者签署H5链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSignUrls(request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    if request == nil {
        request = NewCreateSignUrlsRequest()
    }
    
    response = NewCreateSignUrlsResponse()
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
// 此接口（DescribeFlowDetailInfo）用于查询合同(流程)的详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowDetailInfo(request *DescribeFlowDetailInfoRequest) (response *DescribeFlowDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFlowDetailInfoRequest()
    }
    
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
// 根据流程信息批量获取资源下载链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceUrlsByFlows(request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUrlsByFlowsRequest()
    }
    
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
// 通过此接口（DescribeTemplates）查询该企业在电子签渠道版中配置的有效模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTemplatesRequest()
    }
    
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
// 此接口（DescribeUsage）用于获取渠道所有合作企业流量消耗情况。
//
//  注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsage(request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRequest()
    }
    
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
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载确认页面链接，支持客户合同（流程）归类打包下载。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrl(request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    if request == nil {
        request = NewGetDownloadFlowUrlRequest()
    }
    
    response = NewGetDownloadFlowUrlResponse()
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
// 此接口（OperateChannelTemplate）用于渠道侧将模板库中的模板对合作企业进行查询和设置, 其中包括可见性的修改以及对合作企业的指定.
//
// 1、同步标识=select时：
//
// 输入规则：“可见标识”、“指定合作企业列表”为空。
//
// 处理规则：返回指定模版的可见标识、指定合作企业列表。
//
// 2、同步标识=update时：
//
// 输入规则：“可见标识”、“指定合作企业列表”非必填输入。
//
// 处理规则：
//
// 若“可见标识”=空，不做处理，返回当前的可见标识。
//
// 若“可见标识”=所有合作企业，不取“指定合作企业列表”的值处理。
//
// 若“可见标识”=指定合作企业，取“指定合作企业列表”的值进行更新/插入。
//
// 3、同步标识=delete时：
//
// 输入规则：“可见标识”、“指定合作企业列表”非必填输入。
//
// 处理规则：
//
// 仅取“指定合作企业列表”的值进行删除处理，为空时不做处。
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplate(request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    if request == nil {
        request = NewOperateChannelTemplateRequest()
    }
    
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
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
func (c *Client) PrepareFlows(request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    if request == nil {
        request = NewPrepareFlowsRequest()
    }
    
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
// 此接口（SyncProxyOrganization）用于同步渠道侧企业信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganization(request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationRequest()
    }
    
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
// 此接口（SyncProxyOrganizationOperators）用于同步渠道合作企业经办人列表
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperators(request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationOperatorsRequest()
    }
    
    response = NewSyncProxyOrganizationOperatorsResponse()
    err = c.Send(request, response)
    return
}
