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
// 增加访问控制（自定义策略）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    return c.AddCustomRuleWithContext(context.Background(), request)
}

// AddCustomRule
// 增加访问控制（自定义策略）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewAddDomainWhiteRuleRequest() (request *AddDomainWhiteRuleRequest) {
    request = &AddDomainWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddDomainWhiteRule")
    
    
    return
}

func NewAddDomainWhiteRuleResponse() (response *AddDomainWhiteRuleResponse) {
    response = &AddDomainWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDomainWhiteRule
// 增加域名规则白名单
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddDomainWhiteRule(request *AddDomainWhiteRuleRequest) (response *AddDomainWhiteRuleResponse, err error) {
    return c.AddDomainWhiteRuleWithContext(context.Background(), request)
}

// AddDomainWhiteRule
// 增加域名规则白名单
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddDomainWhiteRuleWithContext(ctx context.Context, request *AddDomainWhiteRuleRequest) (response *AddDomainWhiteRuleResponse, err error) {
    if request == nil {
        request = NewAddDomainWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDomainWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDomainWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddSpartaProtectionRequest() (request *AddSpartaProtectionRequest) {
    request = &AddSpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddSpartaProtection")
    
    
    return
}

func NewAddSpartaProtectionResponse() (response *AddSpartaProtectionResponse) {
    response = &AddSpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSpartaProtection
// 添加Spart防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddSpartaProtection(request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
    return c.AddSpartaProtectionWithContext(context.Background(), request)
}

// AddSpartaProtection
// 添加Spart防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddSpartaProtectionWithContext(ctx context.Context, request *AddSpartaProtectionRequest) (response *AddSpartaProtectionResponse, err error) {
    if request == nil {
        request = NewAddSpartaProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSpartaProtectionResponse()
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
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAttackDownloadRecord(request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
    return c.DeleteAttackDownloadRecordWithContext(context.Background(), request)
}

// DeleteAttackDownloadRecord
// 删除攻击日志下载任务记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
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

func NewDeleteDomainWhiteRulesRequest() (request *DeleteDomainWhiteRulesRequest) {
    request = &DeleteDomainWhiteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteDomainWhiteRules")
    
    
    return
}

func NewDeleteDomainWhiteRulesResponse() (response *DeleteDomainWhiteRulesResponse) {
    response = &DeleteDomainWhiteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomainWhiteRules
// 删除域名规则白名单
//
// 
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDomainWhiteRules(request *DeleteDomainWhiteRulesRequest) (response *DeleteDomainWhiteRulesResponse, err error) {
    return c.DeleteDomainWhiteRulesWithContext(context.Background(), request)
}

// DeleteDomainWhiteRules
// 删除域名规则白名单
//
// 
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDomainWhiteRulesWithContext(ctx context.Context, request *DeleteDomainWhiteRulesRequest) (response *DeleteDomainWhiteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteDomainWhiteRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainWhiteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainWhiteRulesResponse()
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

func NewDeleteIpAccessControlRequest() (request *DeleteIpAccessControlRequest) {
    request = &DeleteIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteIpAccessControl")
    
    
    return
}

func NewDeleteIpAccessControlResponse() (response *DeleteIpAccessControlResponse) {
    response = &DeleteIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIpAccessControl
// Waf IP黑白名单Delete接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIpAccessControl(request *DeleteIpAccessControlRequest) (response *DeleteIpAccessControlResponse, err error) {
    return c.DeleteIpAccessControlWithContext(context.Background(), request)
}

// DeleteIpAccessControl
// Waf IP黑白名单Delete接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIpAccessControlWithContext(ctx context.Context, request *DeleteIpAccessControlRequest) (response *DeleteIpAccessControlResponse, err error) {
    if request == nil {
        request = NewDeleteIpAccessControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIpAccessControlResponse()
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    return c.DeleteSessionWithContext(context.Background(), request)
}

// DeleteSession
// 删除CC攻击的session设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeAttackOverviewRequest() (request *DescribeAttackOverviewRequest) {
    request = &DescribeAttackOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackOverview")
    
    
    return
}

func NewDescribeAttackOverviewResponse() (response *DescribeAttackOverviewResponse) {
    response = &DescribeAttackOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackOverview
// 攻击总览
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
func (c *Client) DescribeAttackOverview(request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
    return c.DescribeAttackOverviewWithContext(context.Background(), request)
}

// DescribeAttackOverview
// 攻击总览
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
func (c *Client) DescribeAttackOverviewWithContext(ctx context.Context, request *DescribeAttackOverviewRequest) (response *DescribeAttackOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAttackOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoDenyIPRequest() (request *DescribeAutoDenyIPRequest) {
    request = &DescribeAutoDenyIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAutoDenyIP")
    
    
    return
}

func NewDescribeAutoDenyIPResponse() (response *DescribeAutoDenyIPResponse) {
    response = &DescribeAutoDenyIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoDenyIP
// 描述WAF自动封禁IP详情,对齐自动封堵状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAutoDenyIP(request *DescribeAutoDenyIPRequest) (response *DescribeAutoDenyIPResponse, err error) {
    return c.DescribeAutoDenyIPWithContext(context.Background(), request)
}

// DescribeAutoDenyIP
// 描述WAF自动封禁IP详情,对齐自动封堵状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAutoDenyIPWithContext(ctx context.Context, request *DescribeAutoDenyIPRequest) (response *DescribeAutoDenyIPResponse, err error) {
    if request == nil {
        request = NewDescribeAutoDenyIPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoDenyIP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoDenyIPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomWhiteRuleRequest() (request *DescribeCustomWhiteRuleRequest) {
    request = &DescribeCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomWhiteRule")
    
    
    return
}

func NewDescribeCustomWhiteRuleResponse() (response *DescribeCustomWhiteRuleResponse) {
    response = &DescribeCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomWhiteRule
// 获取防护配置中的精准白名单策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomWhiteRule(request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
    return c.DescribeCustomWhiteRuleWithContext(context.Background(), request)
}

// DescribeCustomWhiteRule
// 获取防护配置中的精准白名单策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomWhiteRuleWithContext(ctx context.Context, request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDescribeCustomWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainDetailsSaasRequest() (request *DescribeDomainDetailsSaasRequest) {
    request = &DescribeDomainDetailsSaasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsSaas")
    
    
    return
}

func NewDescribeDomainDetailsSaasResponse() (response *DescribeDomainDetailsSaasResponse) {
    response = &DescribeDomainDetailsSaasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainDetailsSaas
// 查询单个saas域名详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomainDetailsSaas(request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
    return c.DescribeDomainDetailsSaasWithContext(context.Background(), request)
}

// DescribeDomainDetailsSaas
// 查询单个saas域名详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomainDetailsSaasWithContext(ctx context.Context, request *DescribeDomainDetailsSaasRequest) (response *DescribeDomainDetailsSaasResponse, err error) {
    if request == nil {
        request = NewDescribeDomainDetailsSaasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainDetailsSaas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainDetailsSaasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainWhiteRulesRequest() (request *DescribeDomainWhiteRulesRequest) {
    request = &DescribeDomainWhiteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainWhiteRules")
    
    
    return
}

func NewDescribeDomainWhiteRulesResponse() (response *DescribeDomainWhiteRulesResponse) {
    response = &DescribeDomainWhiteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainWhiteRules
// 获取域名的规则白名单
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainWhiteRules(request *DescribeDomainWhiteRulesRequest) (response *DescribeDomainWhiteRulesResponse, err error) {
    return c.DescribeDomainWhiteRulesWithContext(context.Background(), request)
}

// DescribeDomainWhiteRules
// 获取域名的规则白名单
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainWhiteRulesWithContext(ctx context.Context, request *DescribeDomainWhiteRulesRequest) (response *DescribeDomainWhiteRulesResponse, err error) {
    if request == nil {
        request = NewDescribeDomainWhiteRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainWhiteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainWhiteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomains
// 查询用户所有域名的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// 查询用户所有域名的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
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

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 查询用户所有实例的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询用户所有实例的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpAccessControlRequest() (request *DescribeIpAccessControlRequest) {
    request = &DescribeIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeIpAccessControl")
    
    
    return
}

func NewDescribeIpAccessControlResponse() (response *DescribeIpAccessControlResponse) {
    response = &DescribeIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpAccessControl
// Waf ip黑白名单查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIpAccessControl(request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
    return c.DescribeIpAccessControlWithContext(context.Background(), request)
}

// DescribeIpAccessControl
// Waf ip黑白名单查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIpAccessControlWithContext(ctx context.Context, request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
    if request == nil {
        request = NewDescribeIpAccessControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpHitItemsRequest() (request *DescribeIpHitItemsRequest) {
    request = &DescribeIpHitItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeIpHitItems")
    
    
    return
}

func NewDescribeIpHitItemsResponse() (response *DescribeIpHitItemsResponse) {
    response = &DescribeIpHitItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpHitItems
// Waf  IP封堵状态查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIpHitItems(request *DescribeIpHitItemsRequest) (response *DescribeIpHitItemsResponse, err error) {
    return c.DescribeIpHitItemsWithContext(context.Background(), request)
}

// DescribeIpHitItems
// Waf  IP封堵状态查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIpHitItemsWithContext(ctx context.Context, request *DescribeIpHitItemsRequest) (response *DescribeIpHitItemsResponse, err error) {
    if request == nil {
        request = NewDescribeIpHitItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpHitItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpHitItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyStatusRequest() (request *DescribePolicyStatusRequest) {
    request = &DescribePolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePolicyStatus")
    
    
    return
}

func NewDescribePolicyStatusResponse() (response *DescribePolicyStatusResponse) {
    response = &DescribePolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePolicyStatus
// 获取防护状态以及生效的实例id
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyStatus(request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
    return c.DescribePolicyStatusWithContext(context.Background(), request)
}

// DescribePolicyStatus
// 获取防护状态以及生效的实例id
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyStatusWithContext(ctx context.Context, request *DescribePolicyStatusRequest) (response *DescribePolicyStatusResponse, err error) {
    if request == nil {
        request = NewDescribePolicyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleLimitRequest() (request *DescribeRuleLimitRequest) {
    request = &DescribeRuleLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeRuleLimit")
    
    
    return
}

func NewDescribeRuleLimitResponse() (response *DescribeRuleLimitResponse) {
    response = &DescribeRuleLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuleLimit
// 获取各个模块具体的规格限制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRuleLimit(request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
    return c.DescribeRuleLimitWithContext(context.Background(), request)
}

// DescribeRuleLimit
// 获取各个模块具体的规格限制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRuleLimitWithContext(ctx context.Context, request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
    if request == nil {
        request = NewDescribeRuleLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCdcClbWafRegionsRequest() (request *DescribeUserCdcClbWafRegionsRequest) {
    request = &DescribeUserCdcClbWafRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserCdcClbWafRegions")
    
    
    return
}

func NewDescribeUserCdcClbWafRegionsResponse() (response *DescribeUserCdcClbWafRegionsResponse) {
    response = &DescribeUserCdcClbWafRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserCdcClbWafRegions
// 在CDC场景下，负载均衡型WAF的添加、编辑域名配置的时候，需要展示CDC负载均衡型WAF（cdc-clb-waf)支持的地域列表，通过DescribeUserCdcClbWafRegions既可以获得当前对客户已经开放的地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCdcClbWafRegions(request *DescribeUserCdcClbWafRegionsRequest) (response *DescribeUserCdcClbWafRegionsResponse, err error) {
    return c.DescribeUserCdcClbWafRegionsWithContext(context.Background(), request)
}

// DescribeUserCdcClbWafRegions
// 在CDC场景下，负载均衡型WAF的添加、编辑域名配置的时候，需要展示CDC负载均衡型WAF（cdc-clb-waf)支持的地域列表，通过DescribeUserCdcClbWafRegions既可以获得当前对客户已经开放的地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCdcClbWafRegionsWithContext(ctx context.Context, request *DescribeUserCdcClbWafRegionsRequest) (response *DescribeUserCdcClbWafRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeUserCdcClbWafRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCdcClbWafRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCdcClbWafRegionsResponse()
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

func NewDescribeWafAutoDenyRulesRequest() (request *DescribeWafAutoDenyRulesRequest) {
    request = &DescribeWafAutoDenyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWafAutoDenyRules")
    
    
    return
}

func NewDescribeWafAutoDenyRulesResponse() (response *DescribeWafAutoDenyRulesResponse) {
    response = &DescribeWafAutoDenyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWafAutoDenyRules
// 返回ip惩罚规则详细信息
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWafAutoDenyRules(request *DescribeWafAutoDenyRulesRequest) (response *DescribeWafAutoDenyRulesResponse, err error) {
    return c.DescribeWafAutoDenyRulesWithContext(context.Background(), request)
}

// DescribeWafAutoDenyRules
// 返回ip惩罚规则详细信息
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWafAutoDenyRulesWithContext(ctx context.Context, request *DescribeWafAutoDenyRulesRequest) (response *DescribeWafAutoDenyRulesResponse, err error) {
    if request == nil {
        request = NewDescribeWafAutoDenyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafAutoDenyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafAutoDenyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafAutoDenyStatusRequest() (request *DescribeWafAutoDenyStatusRequest) {
    request = &DescribeWafAutoDenyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWafAutoDenyStatus")
    
    
    return
}

func NewDescribeWafAutoDenyStatusResponse() (response *DescribeWafAutoDenyStatusResponse) {
    response = &DescribeWafAutoDenyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWafAutoDenyStatus
// 描述WAF自动封禁模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWafAutoDenyStatus(request *DescribeWafAutoDenyStatusRequest) (response *DescribeWafAutoDenyStatusResponse, err error) {
    return c.DescribeWafAutoDenyStatusWithContext(context.Background(), request)
}

// DescribeWafAutoDenyStatus
// 描述WAF自动封禁模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWafAutoDenyStatusWithContext(ctx context.Context, request *DescribeWafAutoDenyStatusRequest) (response *DescribeWafAutoDenyStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWafAutoDenyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafAutoDenyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafAutoDenyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWafThreatenIntelligenceRequest() (request *DescribeWafThreatenIntelligenceRequest) {
    request = &DescribeWafThreatenIntelligenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWafThreatenIntelligence")
    
    
    return
}

func NewDescribeWafThreatenIntelligenceResponse() (response *DescribeWafThreatenIntelligenceResponse) {
    response = &DescribeWafThreatenIntelligenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWafThreatenIntelligence
// 描述WAF威胁情报封禁模块配置详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWafThreatenIntelligence(request *DescribeWafThreatenIntelligenceRequest) (response *DescribeWafThreatenIntelligenceResponse, err error) {
    return c.DescribeWafThreatenIntelligenceWithContext(context.Background(), request)
}

// DescribeWafThreatenIntelligence
// 描述WAF威胁情报封禁模块配置详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWafThreatenIntelligenceWithContext(ctx context.Context, request *DescribeWafThreatenIntelligenceRequest) (response *DescribeWafThreatenIntelligenceResponse, err error) {
    if request == nil {
        request = NewDescribeWafThreatenIntelligenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafThreatenIntelligence require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafThreatenIntelligenceResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttackDownloadRecordsRequest() (request *GetAttackDownloadRecordsRequest) {
    request = &GetAttackDownloadRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetAttackDownloadRecords")
    
    
    return
}

func NewGetAttackDownloadRecordsResponse() (response *GetAttackDownloadRecordsResponse) {
    response = &GetAttackDownloadRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAttackDownloadRecords
// 查询下载攻击日志任务记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetAttackDownloadRecords(request *GetAttackDownloadRecordsRequest) (response *GetAttackDownloadRecordsResponse, err error) {
    return c.GetAttackDownloadRecordsWithContext(context.Background(), request)
}

// GetAttackDownloadRecords
// 查询下载攻击日志任务记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetAttackDownloadRecordsWithContext(ctx context.Context, request *GetAttackDownloadRecordsRequest) (response *GetAttackDownloadRecordsResponse, err error) {
    if request == nil {
        request = NewGetAttackDownloadRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttackDownloadRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttackDownloadRecordsResponse()
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

func NewModifyAreaBanStatusRequest() (request *ModifyAreaBanStatusRequest) {
    request = &ModifyAreaBanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAreaBanStatus")
    
    
    return
}

func NewModifyAreaBanStatusResponse() (response *ModifyAreaBanStatusResponse) {
    response = &ModifyAreaBanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAreaBanStatus
// 修改防护域名的地域封禁状态
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
func (c *Client) ModifyAreaBanStatus(request *ModifyAreaBanStatusRequest) (response *ModifyAreaBanStatusResponse, err error) {
    return c.ModifyAreaBanStatusWithContext(context.Background(), request)
}

// ModifyAreaBanStatus
// 修改防护域名的地域封禁状态
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
func (c *Client) ModifyAreaBanStatusWithContext(ctx context.Context, request *ModifyAreaBanStatusRequest) (response *ModifyAreaBanStatusResponse, err error) {
    if request == nil {
        request = NewModifyAreaBanStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAreaBanStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAreaBanStatusResponse()
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
// 开启或禁用访问控制（自定义策略）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    return c.ModifyCustomRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomRuleStatus
// 开启或禁用访问控制（自定义策略）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewModifyDomainWhiteRuleRequest() (request *ModifyDomainWhiteRuleRequest) {
    request = &ModifyDomainWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainWhiteRule")
    
    
    return
}

func NewModifyDomainWhiteRuleResponse() (response *ModifyDomainWhiteRuleResponse) {
    response = &ModifyDomainWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainWhiteRule
// 更改某一条规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDomainWhiteRule(request *ModifyDomainWhiteRuleRequest) (response *ModifyDomainWhiteRuleResponse, err error) {
    return c.ModifyDomainWhiteRuleWithContext(context.Background(), request)
}

// ModifyDomainWhiteRule
// 更改某一条规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDomainWhiteRuleWithContext(ctx context.Context, request *ModifyDomainWhiteRuleRequest) (response *ModifyDomainWhiteRuleResponse, err error) {
    if request == nil {
        request = NewModifyDomainWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWafAutoDenyRulesRequest() (request *ModifyWafAutoDenyRulesRequest) {
    request = &ModifyWafAutoDenyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyWafAutoDenyRules")
    
    
    return
}

func NewModifyWafAutoDenyRulesResponse() (response *ModifyWafAutoDenyRulesResponse) {
    response = &ModifyWafAutoDenyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWafAutoDenyRules
// 修改ip惩罚规则
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWafAutoDenyRules(request *ModifyWafAutoDenyRulesRequest) (response *ModifyWafAutoDenyRulesResponse, err error) {
    return c.ModifyWafAutoDenyRulesWithContext(context.Background(), request)
}

// ModifyWafAutoDenyRules
// 修改ip惩罚规则
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWafAutoDenyRulesWithContext(ctx context.Context, request *ModifyWafAutoDenyRulesRequest) (response *ModifyWafAutoDenyRulesResponse, err error) {
    if request == nil {
        request = NewModifyWafAutoDenyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWafAutoDenyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWafAutoDenyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWafAutoDenyStatusRequest() (request *ModifyWafAutoDenyStatusRequest) {
    request = &ModifyWafAutoDenyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyWafAutoDenyStatus")
    
    
    return
}

func NewModifyWafAutoDenyStatusResponse() (response *ModifyWafAutoDenyStatusResponse) {
    response = &ModifyWafAutoDenyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWafAutoDenyStatus
// 配置WAF自动封禁模块状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyWafAutoDenyStatus(request *ModifyWafAutoDenyStatusRequest) (response *ModifyWafAutoDenyStatusResponse, err error) {
    return c.ModifyWafAutoDenyStatusWithContext(context.Background(), request)
}

// ModifyWafAutoDenyStatus
// 配置WAF自动封禁模块状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyWafAutoDenyStatusWithContext(ctx context.Context, request *ModifyWafAutoDenyStatusRequest) (response *ModifyWafAutoDenyStatusResponse, err error) {
    if request == nil {
        request = NewModifyWafAutoDenyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWafAutoDenyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWafAutoDenyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWafThreatenIntelligenceRequest() (request *ModifyWafThreatenIntelligenceRequest) {
    request = &ModifyWafThreatenIntelligenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyWafThreatenIntelligence")
    
    
    return
}

func NewModifyWafThreatenIntelligenceResponse() (response *ModifyWafThreatenIntelligenceResponse) {
    response = &ModifyWafThreatenIntelligenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWafThreatenIntelligence
// 配置WAF威胁情报封禁模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyWafThreatenIntelligence(request *ModifyWafThreatenIntelligenceRequest) (response *ModifyWafThreatenIntelligenceResponse, err error) {
    return c.ModifyWafThreatenIntelligenceWithContext(context.Background(), request)
}

// ModifyWafThreatenIntelligence
// 配置WAF威胁情报封禁模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyWafThreatenIntelligenceWithContext(ctx context.Context, request *ModifyWafThreatenIntelligenceRequest) (response *ModifyWafThreatenIntelligenceResponse, err error) {
    if request == nil {
        request = NewModifyWafThreatenIntelligenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWafThreatenIntelligence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWafThreatenIntelligenceResponse()
    err = c.Send(request, response)
    return
}

func NewPostAttackDownloadTaskRequest() (request *PostAttackDownloadTaskRequest) {
    request = &PostAttackDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "PostAttackDownloadTask")
    
    
    return
}

func NewPostAttackDownloadTaskResponse() (response *PostAttackDownloadTaskResponse) {
    response = &PostAttackDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PostAttackDownloadTask
// 创建搜索下载攻击日志任务，使用CLS新版本的搜索下载getlog接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PostAttackDownloadTask(request *PostAttackDownloadTaskRequest) (response *PostAttackDownloadTaskResponse, err error) {
    return c.PostAttackDownloadTaskWithContext(context.Background(), request)
}

// PostAttackDownloadTask
// 创建搜索下载攻击日志任务，使用CLS新版本的搜索下载getlog接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PostAttackDownloadTaskWithContext(ctx context.Context, request *PostAttackDownloadTaskRequest) (response *PostAttackDownloadTaskResponse, err error) {
    if request == nil {
        request = NewPostAttackDownloadTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PostAttackDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewPostAttackDownloadTaskResponse()
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

func NewSearchAttackLogRequest() (request *SearchAttackLogRequest) {
    request = &SearchAttackLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "SearchAttackLog")
    
    
    return
}

func NewSearchAttackLogResponse() (response *SearchAttackLogResponse) {
    response = &SearchAttackLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchAttackLog
// 新版本CLS接口存在参数变化，query改成了query_string支持lucence语法接口搜索查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
func (c *Client) SearchAttackLog(request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
    return c.SearchAttackLogWithContext(context.Background(), request)
}

// SearchAttackLog
// 新版本CLS接口存在参数变化，query改成了query_string支持lucence语法接口搜索查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
func (c *Client) SearchAttackLogWithContext(ctx context.Context, request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
    if request == nil {
        request = NewSearchAttackLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchAttackLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchAttackLogResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDomainRulesRequest() (request *SwitchDomainRulesRequest) {
    request = &SwitchDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "SwitchDomainRules")
    
    
    return
}

func NewSwitchDomainRulesResponse() (response *SwitchDomainRulesResponse) {
    response = &SwitchDomainRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDomainRules
// 切换域名的规则开关
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SwitchDomainRules(request *SwitchDomainRulesRequest) (response *SwitchDomainRulesResponse, err error) {
    return c.SwitchDomainRulesWithContext(context.Background(), request)
}

// SwitchDomainRules
// 切换域名的规则开关
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SwitchDomainRulesWithContext(ctx context.Context, request *SwitchDomainRulesRequest) (response *SwitchDomainRulesResponse, err error) {
    if request == nil {
        request = NewSwitchDomainRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDomainRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewUpsertIpAccessControlRequest() (request *UpsertIpAccessControlRequest) {
    request = &UpsertIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "UpsertIpAccessControl")
    
    
    return
}

func NewUpsertIpAccessControlResponse() (response *UpsertIpAccessControlResponse) {
    response = &UpsertIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpsertIpAccessControl
// Waf IP黑白名单Upsert接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertIpAccessControl(request *UpsertIpAccessControlRequest) (response *UpsertIpAccessControlResponse, err error) {
    return c.UpsertIpAccessControlWithContext(context.Background(), request)
}

// UpsertIpAccessControl
// Waf IP黑白名单Upsert接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertIpAccessControlWithContext(ctx context.Context, request *UpsertIpAccessControlRequest) (response *UpsertIpAccessControlResponse, err error) {
    if request == nil {
        request = NewUpsertIpAccessControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpsertIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpsertIpAccessControlResponse()
    err = c.Send(request, response)
    return
}
