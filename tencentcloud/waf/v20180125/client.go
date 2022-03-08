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

package v20180125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-01-25"

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


func NewAddCustomRuleRequest() (request *AddCustomRuleRequest) {
    request = &AddCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "AddCustomRule")
    
    
    return
}

func NewAddCustomRuleResponse() (response *AddCustomRuleResponse) {
    response = &AddCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCustomRule
// 增加自定义策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    return c.AddCustomRuleWithContext(context.Background(), request)
}

// AddCustomRule
// 增加自定义策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) AddCustomRuleWithContext(ctx context.Context, request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    if request == nil {
        request = NewAddCustomRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessExportRequest() (request *CreateAccessExportRequest) {
    request = &CreateAccessExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "CreateAccessExport")
    
    
    return
}

func NewCreateAccessExportResponse() (response *CreateAccessExportResponse) {
    response = &CreateAccessExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccessExport
// 本接口用于创建访问日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessExport(request *CreateAccessExportRequest) (response *CreateAccessExportResponse, err error) {
    return c.CreateAccessExportWithContext(context.Background(), request)
}

// CreateAccessExport
// 本接口用于创建访问日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAccessExportWithContext(ctx context.Context, request *CreateAccessExportRequest) (response *CreateAccessExportResponse, err error) {
    if request == nil {
        request = NewCreateAccessExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAttackDownloadTaskRequest() (request *CreateAttackDownloadTaskRequest) {
    request = &CreateAttackDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "CreateAttackDownloadTask")
    
    
    return
}

func NewCreateAttackDownloadTaskResponse() (response *CreateAttackDownloadTaskResponse) {
    response = &CreateAttackDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAttackDownloadTask
// 创建攻击日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateAttackDownloadTask(request *CreateAttackDownloadTaskRequest) (response *CreateAttackDownloadTaskResponse, err error) {
    return c.CreateAttackDownloadTaskWithContext(context.Background(), request)
}

// CreateAttackDownloadTask
// 创建攻击日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateAttackDownloadTaskWithContext(ctx context.Context, request *CreateAttackDownloadTaskRequest) (response *CreateAttackDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateAttackDownloadTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAttackDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAttackDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessExportRequest() (request *DeleteAccessExportRequest) {
    request = &DeleteAccessExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAccessExport")
    
    
    return
}

func NewDeleteAccessExportResponse() (response *DeleteAccessExportResponse) {
    response = &DeleteAccessExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessExport
// 本接口用于删除访问日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAccessExport(request *DeleteAccessExportRequest) (response *DeleteAccessExportResponse, err error) {
    return c.DeleteAccessExportWithContext(context.Background(), request)
}

// DeleteAccessExport
// 本接口用于删除访问日志导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAccessExportWithContext(ctx context.Context, request *DeleteAccessExportRequest) (response *DeleteAccessExportResponse, err error) {
    if request == nil {
        request = NewDeleteAccessExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttackDownloadRecordRequest() (request *DeleteAttackDownloadRecordRequest) {
    request = &DeleteAttackDownloadRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAttackDownloadRecord")
    
    
    return
}

func NewDeleteAttackDownloadRecordResponse() (response *DeleteAttackDownloadRecordResponse) {
    response = &DeleteAttackDownloadRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAttackDownloadRecord
// 删除攻击日志下载任务记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAttackDownloadRecord(request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
    return c.DeleteAttackDownloadRecordWithContext(context.Background(), request)
}

// DeleteAttackDownloadRecord
// 删除攻击日志下载任务记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAttackDownloadRecordWithContext(ctx context.Context, request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
    if request == nil {
        request = NewDeleteAttackDownloadRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttackDownloadRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttackDownloadRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDownloadRecordRequest() (request *DeleteDownloadRecordRequest) {
    request = &DeleteDownloadRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteDownloadRecord")
    
    
    return
}

func NewDeleteDownloadRecordResponse() (response *DeleteDownloadRecordResponse) {
    response = &DeleteDownloadRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDownloadRecord
// 删除访问日志下载记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteDownloadRecord(request *DeleteDownloadRecordRequest) (response *DeleteDownloadRecordResponse, err error) {
    return c.DeleteDownloadRecordWithContext(context.Background(), request)
}

// DeleteDownloadRecord
// 删除访问日志下载记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteDownloadRecordWithContext(ctx context.Context, request *DeleteDownloadRecordRequest) (response *DeleteDownloadRecordResponse, err error) {
    if request == nil {
        request = NewDeleteDownloadRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDownloadRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDownloadRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
    request = &DeleteSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteSession")
    
    
    return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
    response = &DeleteSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSession
// 删除CC攻击的session设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    return c.DeleteSessionWithContext(context.Background(), request)
}

// DeleteSession
// 删除CC攻击的session设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteSessionWithContext(ctx context.Context, request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    if request == nil {
        request = NewDeleteSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessExportsRequest() (request *DescribeAccessExportsRequest) {
    request = &DescribeAccessExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessExports")
    
    
    return
}

func NewDescribeAccessExportsResponse() (response *DescribeAccessExportsResponse) {
    response = &DescribeAccessExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessExports
// 本接口用于获取访问日志导出列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessExports(request *DescribeAccessExportsRequest) (response *DescribeAccessExportsResponse, err error) {
    return c.DescribeAccessExportsWithContext(context.Background(), request)
}

// DescribeAccessExports
// 本接口用于获取访问日志导出列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessExportsWithContext(ctx context.Context, request *DescribeAccessExportsRequest) (response *DescribeAccessExportsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessExportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessFastAnalysisRequest() (request *DescribeAccessFastAnalysisRequest) {
    request = &DescribeAccessFastAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessFastAnalysis")
    
    
    return
}

func NewDescribeAccessFastAnalysisResponse() (response *DescribeAccessFastAnalysisResponse) {
    response = &DescribeAccessFastAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessFastAnalysis
// 本接口用于访问日志的快速分析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessFastAnalysis(request *DescribeAccessFastAnalysisRequest) (response *DescribeAccessFastAnalysisResponse, err error) {
    return c.DescribeAccessFastAnalysisWithContext(context.Background(), request)
}

// DescribeAccessFastAnalysis
// 本接口用于访问日志的快速分析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessFastAnalysisWithContext(ctx context.Context, request *DescribeAccessFastAnalysisRequest) (response *DescribeAccessFastAnalysisResponse, err error) {
    if request == nil {
        request = NewDescribeAccessFastAnalysisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessFastAnalysis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessFastAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessIndexRequest() (request *DescribeAccessIndexRequest) {
    request = &DescribeAccessIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessIndex")
    
    
    return
}

func NewDescribeAccessIndexResponse() (response *DescribeAccessIndexResponse) {
    response = &DescribeAccessIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessIndex
// 本接口用于获取访问日志索引配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessIndex(request *DescribeAccessIndexRequest) (response *DescribeAccessIndexResponse, err error) {
    return c.DescribeAccessIndexWithContext(context.Background(), request)
}

// DescribeAccessIndex
// 本接口用于获取访问日志索引配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccessIndexWithContext(ctx context.Context, request *DescribeAccessIndexRequest) (response *DescribeAccessIndexResponse, err error) {
    if request == nil {
        request = NewDescribeAccessIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomRulesRequest() (request *DescribeCustomRulesRequest) {
    request = &DescribeCustomRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRules")
    
    
    return
}

func NewDescribeCustomRulesResponse() (response *DescribeCustomRulesResponse) {
    response = &DescribeCustomRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomRules
// 获取防护配置中的自定义策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomRules(request *DescribeCustomRulesRequest) (response *DescribeCustomRulesResponse, err error) {
    return c.DescribeCustomRulesWithContext(context.Background(), request)
}

// DescribeCustomRules
// 获取防护配置中的自定义策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCustomRulesWithContext(ctx context.Context, request *DescribeCustomRulesRequest) (response *DescribeCustomRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowTrendRequest() (request *DescribeFlowTrendRequest) {
    request = &DescribeFlowTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeFlowTrend")
    
    
    return
}

func NewDescribeFlowTrendResponse() (response *DescribeFlowTrendResponse) {
    response = &DescribeFlowTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowTrend
// 获取waf流量访问趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeFlowTrend(request *DescribeFlowTrendRequest) (response *DescribeFlowTrendResponse, err error) {
    return c.DescribeFlowTrendWithContext(context.Background(), request)
}

// DescribeFlowTrend
// 获取waf流量访问趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeFlowTrendWithContext(ctx context.Context, request *DescribeFlowTrendRequest) (response *DescribeFlowTrendResponse, err error) {
    if request == nil {
        request = NewDescribeFlowTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserClbWafRegionsRequest() (request *DescribeUserClbWafRegionsRequest) {
    request = &DescribeUserClbWafRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserClbWafRegions")
    
    
    return
}

func NewDescribeUserClbWafRegionsResponse() (response *DescribeUserClbWafRegionsResponse) {
    response = &DescribeUserClbWafRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserClbWafRegions
// 在负载均衡型WAF的添加、编辑域名配置的时候，需要展示负载均衡型WAF（clb-waf)支持的地域列表，通过DescribeUserClbWafRegions既可以获得当前对客户已经开放的地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserClbWafRegions(request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
    return c.DescribeUserClbWafRegionsWithContext(context.Background(), request)
}

// DescribeUserClbWafRegions
// 在负载均衡型WAF的添加、编辑域名配置的时候，需要展示负载均衡型WAF（clb-waf)支持的地域列表，通过DescribeUserClbWafRegions既可以获得当前对客户已经开放的地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserClbWafRegionsWithContext(ctx context.Context, request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeUserClbWafRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserClbWafRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserClbWafRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessPeriodRequest() (request *ModifyAccessPeriodRequest) {
    request = &ModifyAccessPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAccessPeriod")
    
    
    return
}

func NewModifyAccessPeriodResponse() (response *ModifyAccessPeriodResponse) {
    response = &ModifyAccessPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessPeriod
// 本接口用于修改访问日志保存期限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessPeriod(request *ModifyAccessPeriodRequest) (response *ModifyAccessPeriodResponse, err error) {
    return c.ModifyAccessPeriodWithContext(context.Background(), request)
}

// ModifyAccessPeriod
// 本接口用于修改访问日志保存期限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAccessPeriodWithContext(ctx context.Context, request *ModifyAccessPeriodRequest) (response *ModifyAccessPeriodResponse, err error) {
    if request == nil {
        request = NewModifyAccessPeriodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomRuleStatusRequest() (request *ModifyCustomRuleStatusRequest) {
    request = &ModifyCustomRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRuleStatus")
    
    
    return
}

func NewModifyCustomRuleStatusResponse() (response *ModifyCustomRuleStatusResponse) {
    response = &ModifyCustomRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomRuleStatus
// 开启或禁用自定义策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    return c.ModifyCustomRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomRuleStatus
// 开启或禁用自定义策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCustomRuleStatusWithContext(ctx context.Context, request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyCustomRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSearchAccessLogRequest() (request *SearchAccessLogRequest) {
    request = &SearchAccessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "SearchAccessLog")
    
    
    return
}

func NewSearchAccessLogResponse() (response *SearchAccessLogResponse) {
    response = &SearchAccessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchAccessLog
// 本接口用于搜索WAF访问日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchAccessLog(request *SearchAccessLogRequest) (response *SearchAccessLogResponse, err error) {
    return c.SearchAccessLogWithContext(context.Background(), request)
}

// SearchAccessLog
// 本接口用于搜索WAF访问日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchAccessLogWithContext(ctx context.Context, request *SearchAccessLogRequest) (response *SearchAccessLogResponse, err error) {
    if request == nil {
        request = NewSearchAccessLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchAccessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchAccessLogResponse()
    err = c.Send(request, response)
    return
}
