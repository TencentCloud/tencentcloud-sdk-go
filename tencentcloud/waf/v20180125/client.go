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


func NewAddAntiFakeUrlRequest() (request *AddAntiFakeUrlRequest) {
    request = &AddAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddAntiFakeUrl")
    
    
    return
}

func NewAddAntiFakeUrlResponse() (response *AddAntiFakeUrlResponse) {
    response = &AddAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAntiFakeUrl
// 添加防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
func (c *Client) AddAntiFakeUrl(request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
    return c.AddAntiFakeUrlWithContext(context.Background(), request)
}

// AddAntiFakeUrl
// 添加防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
func (c *Client) AddAntiFakeUrlWithContext(ctx context.Context, request *AddAntiFakeUrlRequest) (response *AddAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewAddAntiFakeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewAddAntiInfoLeakRulesRequest() (request *AddAntiInfoLeakRulesRequest) {
    request = &AddAntiInfoLeakRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddAntiInfoLeakRules")
    
    
    return
}

func NewAddAntiInfoLeakRulesResponse() (response *AddAntiInfoLeakRulesResponse) {
    response = &AddAntiInfoLeakRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAntiInfoLeakRules
// 添加信息防泄漏规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAntiInfoLeakRules(request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
    return c.AddAntiInfoLeakRulesWithContext(context.Background(), request)
}

// AddAntiInfoLeakRules
// 添加信息防泄漏规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAntiInfoLeakRulesWithContext(ctx context.Context, request *AddAntiInfoLeakRulesRequest) (response *AddAntiInfoLeakRulesResponse, err error) {
    if request == nil {
        request = NewAddAntiInfoLeakRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAntiInfoLeakRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAntiInfoLeakRulesResponse()
    err = c.Send(request, response)
    return
}

func NewAddAttackWhiteRuleRequest() (request *AddAttackWhiteRuleRequest) {
    request = &AddAttackWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddAttackWhiteRule")
    
    
    return
}

func NewAddAttackWhiteRuleResponse() (response *AddAttackWhiteRuleResponse) {
    response = &AddAttackWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAttackWhiteRule
// 供用户控制台调用，增加Tiga规则引擎白名单。
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
func (c *Client) AddAttackWhiteRule(request *AddAttackWhiteRuleRequest) (response *AddAttackWhiteRuleResponse, err error) {
    return c.AddAttackWhiteRuleWithContext(context.Background(), request)
}

// AddAttackWhiteRule
// 供用户控制台调用，增加Tiga规则引擎白名单。
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
func (c *Client) AddAttackWhiteRuleWithContext(ctx context.Context, request *AddAttackWhiteRuleRequest) (response *AddAttackWhiteRuleResponse, err error) {
    if request == nil {
        request = NewAddAttackWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAttackWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAttackWhiteRuleResponse()
    err = c.Send(request, response)
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
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
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
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
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

func NewAddCustomWhiteRuleRequest() (request *AddCustomWhiteRuleRequest) {
    request = &AddCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "AddCustomWhiteRule")
    
    
    return
}

func NewAddCustomWhiteRuleResponse() (response *AddCustomWhiteRuleResponse) {
    response = &AddCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCustomWhiteRule
// 增加精准白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomWhiteRule(request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
    return c.AddCustomWhiteRuleWithContext(context.Background(), request)
}

// AddCustomWhiteRule
// 增加精准白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCustomWhiteRuleWithContext(ctx context.Context, request *AddCustomWhiteRuleRequest) (response *AddCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewAddCustomWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCustomWhiteRuleResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) AddDomainWhiteRule(request *AddDomainWhiteRuleRequest) (response *AddDomainWhiteRuleResponse, err error) {
    return c.AddDomainWhiteRuleWithContext(context.Background(), request)
}

// AddDomainWhiteRule
// 增加域名规则白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
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
// 添加SaaS型WAF防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_DOMAINEXCEEDSLIMITERR = "InvalidParameter.DomainExceedsLimitErr"
//  INVALIDPARAMETER_DOMAINNOTRECORD = "InvalidParameter.DomainNotRecord"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETER_XFFRESETPARAMETERERR = "InvalidParameter.XFFResetParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_EMPTYERR = "ResourceInUse.EmptyErr"
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
// 添加SaaS型WAF防护域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_DOMAINEXCEEDSLIMITERR = "InvalidParameter.DomainExceedsLimitErr"
//  INVALIDPARAMETER_DOMAINNOTRECORD = "InvalidParameter.DomainNotRecord"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
//  INVALIDPARAMETER_XFFRESETPARAMETERERR = "InvalidParameter.XFFResetParameterErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_EMPTYERR = "ResourceInUse.EmptyErr"
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

func NewCreateHostRequest() (request *CreateHostRequest) {
    request = &CreateHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "CreateHost")
    
    
    return
}

func NewCreateHostResponse() (response *CreateHostResponse) {
    response = &CreateHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHost
// clb-waf中添加防护域名
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
func (c *Client) CreateHost(request *CreateHostRequest) (response *CreateHostResponse, err error) {
    return c.CreateHostWithContext(context.Background(), request)
}

// CreateHost
// clb-waf中添加防护域名
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
func (c *Client) CreateHostWithContext(ctx context.Context, request *CreateHostRequest) (response *CreateHostResponse, err error) {
    if request == nil {
        request = NewCreateHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHostResponse()
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

func NewDeleteAntiFakeUrlRequest() (request *DeleteAntiFakeUrlRequest) {
    request = &DeleteAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiFakeUrl")
    
    
    return
}

func NewDeleteAntiFakeUrlResponse() (response *DeleteAntiFakeUrlResponse) {
    response = &DeleteAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAntiFakeUrl
// 删除防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAntiFakeUrl(request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
    return c.DeleteAntiFakeUrlWithContext(context.Background(), request)
}

// DeleteAntiFakeUrl
// 删除防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAntiFakeUrlWithContext(ctx context.Context, request *DeleteAntiFakeUrlRequest) (response *DeleteAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewDeleteAntiFakeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAntiInfoLeakRuleRequest() (request *DeleteAntiInfoLeakRuleRequest) {
    request = &DeleteAntiInfoLeakRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAntiInfoLeakRule")
    
    
    return
}

func NewDeleteAntiInfoLeakRuleResponse() (response *DeleteAntiInfoLeakRuleResponse) {
    response = &DeleteAntiInfoLeakRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAntiInfoLeakRule
// 信息防泄漏删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DeleteAntiInfoLeakRule(request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
    return c.DeleteAntiInfoLeakRuleWithContext(context.Background(), request)
}

// DeleteAntiInfoLeakRule
// 信息防泄漏删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DeleteAntiInfoLeakRuleWithContext(ctx context.Context, request *DeleteAntiInfoLeakRuleRequest) (response *DeleteAntiInfoLeakRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAntiInfoLeakRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAntiInfoLeakRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAntiInfoLeakRuleResponse()
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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

func NewDeleteAttackWhiteRuleRequest() (request *DeleteAttackWhiteRuleRequest) {
    request = &DeleteAttackWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAttackWhiteRule")
    
    
    return
}

func NewDeleteAttackWhiteRuleResponse() (response *DeleteAttackWhiteRuleResponse) {
    response = &DeleteAttackWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAttackWhiteRule
// 供用户控制台调用，删除Tiga规则引擎白名单。
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
func (c *Client) DeleteAttackWhiteRule(request *DeleteAttackWhiteRuleRequest) (response *DeleteAttackWhiteRuleResponse, err error) {
    return c.DeleteAttackWhiteRuleWithContext(context.Background(), request)
}

// DeleteAttackWhiteRule
// 供用户控制台调用，删除Tiga规则引擎白名单。
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
func (c *Client) DeleteAttackWhiteRuleWithContext(ctx context.Context, request *DeleteAttackWhiteRuleRequest) (response *DeleteAttackWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAttackWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttackWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttackWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCRuleRequest() (request *DeleteCCRuleRequest) {
    request = &DeleteCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCCRule")
    
    
    return
}

func NewDeleteCCRuleResponse() (response *DeleteCCRuleResponse) {
    response = &DeleteCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCCRule
// Waf  CC V2 Delete接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRule(request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
    return c.DeleteCCRuleWithContext(context.Background(), request)
}

// DeleteCCRule
// Waf  CC V2 Delete接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCCRuleWithContext(ctx context.Context, request *DeleteCCRuleRequest) (response *DeleteCCRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCCRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomRuleRequest() (request *DeleteCustomRuleRequest) {
    request = &DeleteCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomRule")
    
    
    return
}

func NewDeleteCustomRuleResponse() (response *DeleteCustomRuleResponse) {
    response = &DeleteCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomRule
// 删除自定义规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomRule(request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
    return c.DeleteCustomRuleWithContext(context.Background(), request)
}

// DeleteCustomRule
// 删除自定义规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomRuleWithContext(ctx context.Context, request *DeleteCustomRuleRequest) (response *DeleteCustomRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCustomRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomWhiteRuleRequest() (request *DeleteCustomWhiteRuleRequest) {
    request = &DeleteCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteCustomWhiteRule")
    
    
    return
}

func NewDeleteCustomWhiteRuleResponse() (response *DeleteCustomWhiteRuleResponse) {
    response = &DeleteCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomWhiteRule
// 删除精准白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomWhiteRule(request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
    return c.DeleteCustomWhiteRuleWithContext(context.Background(), request)
}

// DeleteCustomWhiteRule
// 删除精准白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomWhiteRuleWithContext(ctx context.Context, request *DeleteCustomWhiteRuleRequest) (response *DeleteCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCustomWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomWhiteRuleResponse()
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
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) DeleteDomainWhiteRules(request *DeleteDomainWhiteRulesRequest) (response *DeleteDomainWhiteRulesResponse, err error) {
    return c.DeleteDomainWhiteRulesWithContext(context.Background(), request)
}

// DeleteDomainWhiteRules
// 删除域名规则白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
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

func NewDeleteHostRequest() (request *DeleteHostRequest) {
    request = &DeleteHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteHost")
    
    
    return
}

func NewDeleteHostResponse() (response *DeleteHostResponse) {
    response = &DeleteHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHost
// 删除CLB-WAF防护域名
//
// 支持批量操作
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
func (c *Client) DeleteHost(request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
    return c.DeleteHostWithContext(context.Background(), request)
}

// DeleteHost
// 删除CLB-WAF防护域名
//
// 支持批量操作
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
func (c *Client) DeleteHostWithContext(ctx context.Context, request *DeleteHostRequest) (response *DeleteHostResponse, err error) {
    if request == nil {
        request = NewDeleteHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHostResponse()
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
//  FAILEDOPERATION_THENUMBEROFONETIMEDELETIONSREACHEDTHEUPPERLIMIT = "FailedOperation.TheNumberOfOneTimeDeletionsReachedTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIpAccessControl(request *DeleteIpAccessControlRequest) (response *DeleteIpAccessControlResponse, err error) {
    return c.DeleteIpAccessControlWithContext(context.Background(), request)
}

// DeleteIpAccessControl
// Waf IP黑白名单Delete接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_THENUMBEROFONETIMEDELETIONSREACHEDTHEUPPERLIMIT = "FailedOperation.TheNumberOfOneTimeDeletionsReachedTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
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
//  FAILEDOPERATION_SESSIONINUSED = "FailedOperation.SessionInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    return c.DeleteSessionWithContext(context.Background(), request)
}

// DeleteSession
// 删除CC攻击的session设置
//
// 可能返回的错误码:
//  FAILEDOPERATION_SESSIONINUSED = "FailedOperation.SessionInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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

func NewDeleteSpartaProtectionRequest() (request *DeleteSpartaProtectionRequest) {
    request = &DeleteSpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DeleteSpartaProtection")
    
    
    return
}

func NewDeleteSpartaProtectionResponse() (response *DeleteSpartaProtectionResponse) {
    response = &DeleteSpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSpartaProtection
// Saas型WAF删除防护域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteSpartaProtection(request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
    return c.DeleteSpartaProtectionWithContext(context.Background(), request)
}

// DeleteSpartaProtection
// Saas型WAF删除防护域名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER_UNAUTHORIZEDOPERATIONPARAMETERERR = "InvalidParameter.UnauthorizedOperationParameterErr"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteSpartaProtectionWithContext(ctx context.Context, request *DeleteSpartaProtectionRequest) (response *DeleteSpartaProtectionResponse, err error) {
    if request == nil {
        request = NewDeleteSpartaProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSpartaProtectionResponse()
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

func NewDescribeAccessHistogramRequest() (request *DescribeAccessHistogramRequest) {
    request = &DescribeAccessHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAccessHistogram")
    
    
    return
}

func NewDescribeAccessHistogramResponse() (response *DescribeAccessHistogramResponse) {
    response = &DescribeAccessHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccessHistogram
// 本接口用于访问日志柱状趋势图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
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
func (c *Client) DescribeAccessHistogram(request *DescribeAccessHistogramRequest) (response *DescribeAccessHistogramResponse, err error) {
    return c.DescribeAccessHistogramWithContext(context.Background(), request)
}

// DescribeAccessHistogram
// 本接口用于访问日志柱状趋势图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
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
func (c *Client) DescribeAccessHistogramWithContext(ctx context.Context, request *DescribeAccessHistogramRequest) (response *DescribeAccessHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeAccessHistogramRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessHistogramResponse()
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

func NewDescribeAntiFakeRulesRequest() (request *DescribeAntiFakeRulesRequest) {
    request = &DescribeAntiFakeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiFakeRules")
    
    
    return
}

func NewDescribeAntiFakeRulesResponse() (response *DescribeAntiFakeRulesResponse) {
    response = &DescribeAntiFakeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiFakeRules
// 获取防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAntiFakeRules(request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
    return c.DescribeAntiFakeRulesWithContext(context.Background(), request)
}

// DescribeAntiFakeRules
// 获取防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAntiFakeRulesWithContext(ctx context.Context, request *DescribeAntiFakeRulesRequest) (response *DescribeAntiFakeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAntiFakeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiFakeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiFakeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAntiFakeUrlRequest() (request *DescribeAntiFakeUrlRequest) {
    request = &DescribeAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiFakeUrl")
    
    
    return
}

func NewDescribeAntiFakeUrlResponse() (response *DescribeAntiFakeUrlResponse) {
    response = &DescribeAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiFakeUrl
// 获取防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAntiFakeUrl(request *DescribeAntiFakeUrlRequest) (response *DescribeAntiFakeUrlResponse, err error) {
    return c.DescribeAntiFakeUrlWithContext(context.Background(), request)
}

// DescribeAntiFakeUrl
// 获取防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAntiFakeUrlWithContext(ctx context.Context, request *DescribeAntiFakeUrlRequest) (response *DescribeAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewDescribeAntiFakeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAntiInfoLeakRulesRequest() (request *DescribeAntiInfoLeakRulesRequest) {
    request = &DescribeAntiInfoLeakRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiInfoLeakRules")
    
    
    return
}

func NewDescribeAntiInfoLeakRulesResponse() (response *DescribeAntiInfoLeakRulesResponse) {
    response = &DescribeAntiInfoLeakRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiInfoLeakRules
// 老接口已经不再使用。
//
// 
//
// 获取信息防泄漏规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAntiInfoLeakRules(request *DescribeAntiInfoLeakRulesRequest) (response *DescribeAntiInfoLeakRulesResponse, err error) {
    return c.DescribeAntiInfoLeakRulesWithContext(context.Background(), request)
}

// DescribeAntiInfoLeakRules
// 老接口已经不再使用。
//
// 
//
// 获取信息防泄漏规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAntiInfoLeakRulesWithContext(ctx context.Context, request *DescribeAntiInfoLeakRulesRequest) (response *DescribeAntiInfoLeakRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAntiInfoLeakRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiInfoLeakRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiInfoLeakRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAntiInfoLeakageRulesRequest() (request *DescribeAntiInfoLeakageRulesRequest) {
    request = &DescribeAntiInfoLeakageRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAntiInfoLeakageRules")
    
    
    return
}

func NewDescribeAntiInfoLeakageRulesResponse() (response *DescribeAntiInfoLeakageRulesResponse) {
    response = &DescribeAntiInfoLeakageRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAntiInfoLeakageRules
// 取得信息防泄漏规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeAntiInfoLeakageRules(request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
    return c.DescribeAntiInfoLeakageRulesWithContext(context.Background(), request)
}

// DescribeAntiInfoLeakageRules
// 取得信息防泄漏规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeAntiInfoLeakageRulesWithContext(ctx context.Context, request *DescribeAntiInfoLeakageRulesRequest) (response *DescribeAntiInfoLeakageRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAntiInfoLeakageRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAntiInfoLeakageRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAntiInfoLeakageRulesResponse()
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

func NewDescribeAttackTypeRequest() (request *DescribeAttackTypeRequest) {
    request = &DescribeAttackTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackType")
    
    
    return
}

func NewDescribeAttackTypeResponse() (response *DescribeAttackTypeResponse) {
    response = &DescribeAttackTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttackType
// 查询指定域名TOP N攻击类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAttackType(request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
    return c.DescribeAttackTypeWithContext(context.Background(), request)
}

// DescribeAttackType
// 查询指定域名TOP N攻击类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAttackTypeWithContext(ctx context.Context, request *DescribeAttackTypeRequest) (response *DescribeAttackTypeResponse, err error) {
    if request == nil {
        request = NewDescribeAttackTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackWhiteRuleRequest() (request *DescribeAttackWhiteRuleRequest) {
    request = &DescribeAttackWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeAttackWhiteRule")
    
    
    return
}

func NewDescribeAttackWhiteRuleResponse() (response *DescribeAttackWhiteRuleResponse) {
    response = &DescribeAttackWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttackWhiteRule
// 获取用户规则白名单列表
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
func (c *Client) DescribeAttackWhiteRule(request *DescribeAttackWhiteRuleRequest) (response *DescribeAttackWhiteRuleResponse, err error) {
    return c.DescribeAttackWhiteRuleWithContext(context.Background(), request)
}

// DescribeAttackWhiteRule
// 获取用户规则白名单列表
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
func (c *Client) DescribeAttackWhiteRuleWithContext(ctx context.Context, request *DescribeAttackWhiteRuleRequest) (response *DescribeAttackWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDescribeAttackWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackWhiteRuleResponse()
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
// 接口已废弃
//
// 
//
// 描述WAF自动封禁IP详情,对齐自动封堵状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAutoDenyIP(request *DescribeAutoDenyIPRequest) (response *DescribeAutoDenyIPResponse, err error) {
    return c.DescribeAutoDenyIPWithContext(context.Background(), request)
}

// DescribeAutoDenyIP
// 接口已废弃
//
// 
//
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

func NewDescribeBatchIpAccessControlRequest() (request *DescribeBatchIpAccessControlRequest) {
    request = &DescribeBatchIpAccessControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeBatchIpAccessControl")
    
    
    return
}

func NewDescribeBatchIpAccessControlResponse() (response *DescribeBatchIpAccessControlResponse) {
    response = &DescribeBatchIpAccessControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchIpAccessControl
// Waf 多域名ip黑白名单查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBatchIpAccessControl(request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
    return c.DescribeBatchIpAccessControlWithContext(context.Background(), request)
}

// DescribeBatchIpAccessControl
// Waf 多域名ip黑白名单查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBatchIpAccessControlWithContext(ctx context.Context, request *DescribeBatchIpAccessControlRequest) (response *DescribeBatchIpAccessControlResponse, err error) {
    if request == nil {
        request = NewDescribeBatchIpAccessControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchIpAccessControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchIpAccessControlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCRuleRequest() (request *DescribeCCRuleRequest) {
    request = &DescribeCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRule")
    
    
    return
}

func NewDescribeCCRuleResponse() (response *DescribeCCRuleResponse) {
    response = &DescribeCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCCRule
// Waf  CC V2 Query接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCCRule(request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
    return c.DescribeCCRuleWithContext(context.Background(), request)
}

// DescribeCCRule
// Waf  CC V2 Query接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCCRuleWithContext(ctx context.Context, request *DescribeCCRuleRequest) (response *DescribeCCRuleResponse, err error) {
    if request == nil {
        request = NewDescribeCCRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCRuleListRequest() (request *DescribeCCRuleListRequest) {
    request = &DescribeCCRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCCRuleList")
    
    
    return
}

func NewDescribeCCRuleListResponse() (response *DescribeCCRuleListResponse) {
    response = &DescribeCCRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCCRuleList
// 根据多条件查询CC规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCRuleList(request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
    return c.DescribeCCRuleListWithContext(context.Background(), request)
}

// DescribeCCRuleList
// 根据多条件查询CC规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCCRuleListWithContext(ctx context.Context, request *DescribeCCRuleListRequest) (response *DescribeCCRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeCCRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateVerifyResultRequest() (request *DescribeCertificateVerifyResultRequest) {
    request = &DescribeCertificateVerifyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCertificateVerifyResult")
    
    
    return
}

func NewDescribeCertificateVerifyResultResponse() (response *DescribeCertificateVerifyResultResponse) {
    response = &DescribeCertificateVerifyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateVerifyResult
// 获取证书的检查结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCertificateVerifyResult(request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
    return c.DescribeCertificateVerifyResultWithContext(context.Background(), request)
}

// DescribeCertificateVerifyResult
// 获取证书的检查结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCertificateVerifyResultWithContext(ctx context.Context, request *DescribeCertificateVerifyResultRequest) (response *DescribeCertificateVerifyResultResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateVerifyResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateVerifyResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateVerifyResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCiphersDetailRequest() (request *DescribeCiphersDetailRequest) {
    request = &DescribeCiphersDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCiphersDetail")
    
    
    return
}

func NewDescribeCiphersDetailResponse() (response *DescribeCiphersDetailResponse) {
    response = &DescribeCiphersDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCiphersDetail
// Saas型WAF接入查询加密套件信息
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
func (c *Client) DescribeCiphersDetail(request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
    return c.DescribeCiphersDetailWithContext(context.Background(), request)
}

// DescribeCiphersDetail
// Saas型WAF接入查询加密套件信息
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
func (c *Client) DescribeCiphersDetailWithContext(ctx context.Context, request *DescribeCiphersDetailRequest) (response *DescribeCiphersDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCiphersDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCiphersDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCiphersDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomRuleListRequest() (request *DescribeCustomRuleListRequest) {
    request = &DescribeCustomRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRuleList")
    
    
    return
}

func NewDescribeCustomRuleListResponse() (response *DescribeCustomRuleListResponse) {
    response = &DescribeCustomRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomRuleList
// 获取防护配置中的访问控制策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomRuleList(request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
    return c.DescribeCustomRuleListWithContext(context.Background(), request)
}

// DescribeCustomRuleList
// 获取防护配置中的访问控制策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomRuleListWithContext(ctx context.Context, request *DescribeCustomRuleListRequest) (response *DescribeCustomRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomRuleListResponse()
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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

func NewDescribeDomainCountInfoRequest() (request *DescribeDomainCountInfoRequest) {
    request = &DescribeDomainCountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainCountInfo")
    
    
    return
}

func NewDescribeDomainCountInfoResponse() (response *DescribeDomainCountInfoResponse) {
    response = &DescribeDomainCountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainCountInfo
// 获取域名概况
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
func (c *Client) DescribeDomainCountInfo(request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
    return c.DescribeDomainCountInfoWithContext(context.Background(), request)
}

// DescribeDomainCountInfo
// 获取域名概况
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
func (c *Client) DescribeDomainCountInfoWithContext(ctx context.Context, request *DescribeDomainCountInfoRequest) (response *DescribeDomainCountInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainCountInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainCountInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainCountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainDetailsClbRequest() (request *DescribeDomainDetailsClbRequest) {
    request = &DescribeDomainDetailsClbRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainDetailsClb")
    
    
    return
}

func NewDescribeDomainDetailsClbResponse() (response *DescribeDomainDetailsClbResponse) {
    response = &DescribeDomainDetailsClbResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainDetailsClb
// 获取一个clbwaf域名详情
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
func (c *Client) DescribeDomainDetailsClb(request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
    return c.DescribeDomainDetailsClbWithContext(context.Background(), request)
}

// DescribeDomainDetailsClb
// 获取一个clbwaf域名详情
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
func (c *Client) DescribeDomainDetailsClbWithContext(ctx context.Context, request *DescribeDomainDetailsClbRequest) (response *DescribeDomainDetailsClbResponse, err error) {
    if request == nil {
        request = NewDescribeDomainDetailsClbRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainDetailsClb require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainDetailsClbResponse()
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
// 查询单个saaswaf域名详情
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
// 查询单个saaswaf域名详情
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

func NewDescribeDomainRulesRequest() (request *DescribeDomainRulesRequest) {
    request = &DescribeDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainRules")
    
    
    return
}

func NewDescribeDomainRulesResponse() (response *DescribeDomainRulesResponse) {
    response = &DescribeDomainRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainRules
// 拉取域名的防护规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) DescribeDomainRules(request *DescribeDomainRulesRequest) (response *DescribeDomainRulesResponse, err error) {
    return c.DescribeDomainRulesWithContext(context.Background(), request)
}

// DescribeDomainRules
// 拉取域名的防护规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) DescribeDomainRulesWithContext(ctx context.Context, request *DescribeDomainRulesRequest) (response *DescribeDomainRulesResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainVerifyResultRequest() (request *DescribeDomainVerifyResultRequest) {
    request = &DescribeDomainVerifyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeDomainVerifyResult")
    
    
    return
}

func NewDescribeDomainVerifyResultResponse() (response *DescribeDomainVerifyResultResponse) {
    response = &DescribeDomainVerifyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainVerifyResult
// 获取添加域名操作的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeDomainVerifyResult(request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
    return c.DescribeDomainVerifyResultWithContext(context.Background(), request)
}

// DescribeDomainVerifyResult
// 获取添加域名操作的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeDomainVerifyResultWithContext(ctx context.Context, request *DescribeDomainVerifyResultRequest) (response *DescribeDomainVerifyResultResponse, err error) {
    if request == nil {
        request = NewDescribeDomainVerifyResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainVerifyResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainVerifyResultResponse()
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
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDomainWhiteRules(request *DescribeDomainWhiteRulesRequest) (response *DescribeDomainWhiteRulesResponse, err error) {
    return c.DescribeDomainWhiteRulesWithContext(context.Background(), request)
}

// DescribeDomainWhiteRules
// 获取域名的规则白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

func NewDescribeFindDomainListRequest() (request *DescribeFindDomainListRequest) {
    request = &DescribeFindDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeFindDomainList")
    
    
    return
}

func NewDescribeFindDomainListResponse() (response *DescribeFindDomainListResponse) {
    response = &DescribeFindDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFindDomainList
// 获取发现域名列表接口
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
func (c *Client) DescribeFindDomainList(request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
    return c.DescribeFindDomainListWithContext(context.Background(), request)
}

// DescribeFindDomainList
// 获取发现域名列表接口
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
func (c *Client) DescribeFindDomainListWithContext(ctx context.Context, request *DescribeFindDomainListRequest) (response *DescribeFindDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeFindDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFindDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFindDomainListResponse()
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
//  FAILEDOPERATION_CLICKHOUSEOPERATIONFAILED = "FailedOperation.ClickHouseOperationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNKNOWNACTION = "InvalidParameter.UnknownAction"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) DescribeFlowTrend(request *DescribeFlowTrendRequest) (response *DescribeFlowTrendResponse, err error) {
    return c.DescribeFlowTrendWithContext(context.Background(), request)
}

// DescribeFlowTrend
// 获取waf流量访问趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLICKHOUSEOPERATIONFAILED = "FailedOperation.ClickHouseOperationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNKNOWNACTION = "InvalidParameter.UnknownAction"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
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

func NewDescribeHistogramRequest() (request *DescribeHistogramRequest) {
    request = &DescribeHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHistogram")
    
    
    return
}

func NewDescribeHistogramResponse() (response *DescribeHistogramResponse) {
    response = &DescribeHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHistogram
// 查询多种条件的聚类分析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistogram(request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
    return c.DescribeHistogramWithContext(context.Background(), request)
}

// DescribeHistogram
// 查询多种条件的聚类分析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistogramWithContext(ctx context.Context, request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeHistogramRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostRequest() (request *DescribeHostRequest) {
    request = &DescribeHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHost")
    
    
    return
}

func NewDescribeHostResponse() (response *DescribeHostResponse) {
    response = &DescribeHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHost
// clb-waf获取防护域名详情
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
func (c *Client) DescribeHost(request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
    return c.DescribeHostWithContext(context.Background(), request)
}

// DescribeHost
// clb-waf获取防护域名详情
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
func (c *Client) DescribeHostWithContext(ctx context.Context, request *DescribeHostRequest) (response *DescribeHostResponse, err error) {
    if request == nil {
        request = NewDescribeHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostLimitRequest() (request *DescribeHostLimitRequest) {
    request = &DescribeHostLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHostLimit")
    
    
    return
}

func NewDescribeHostLimitResponse() (response *DescribeHostLimitResponse) {
    response = &DescribeHostLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostLimit
// 添加域名的首先验证是否购买了套餐，是否没有达到购买套餐的限制，域名是否已经添加
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
func (c *Client) DescribeHostLimit(request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
    return c.DescribeHostLimitWithContext(context.Background(), request)
}

// DescribeHostLimit
// 添加域名的首先验证是否购买了套餐，是否没有达到购买套餐的限制，域名是否已经添加
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
func (c *Client) DescribeHostLimitWithContext(ctx context.Context, request *DescribeHostLimitRequest) (response *DescribeHostLimitResponse, err error) {
    if request == nil {
        request = NewDescribeHostLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
    request = &DescribeHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeHosts")
    
    
    return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
    response = &DescribeHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHosts
// clb-waf中获取防护域名列表
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
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    return c.DescribeHostsWithContext(context.Background(), request)
}

// DescribeHosts
// clb-waf中获取防护域名列表
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
func (c *Client) DescribeHostsWithContext(ctx context.Context, request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsResponse()
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) DescribeIpAccessControl(request *DescribeIpAccessControlRequest) (response *DescribeIpAccessControlResponse, err error) {
    return c.DescribeIpAccessControlWithContext(context.Background(), request)
}

// DescribeIpAccessControl
// Waf ip黑白名单查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIpHitItems(request *DescribeIpHitItemsRequest) (response *DescribeIpHitItemsResponse, err error) {
    return c.DescribeIpHitItemsWithContext(context.Background(), request)
}

// DescribeIpHitItems
// Waf  IP封堵状态查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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

func NewDescribeObjectsRequest() (request *DescribeObjectsRequest) {
    request = &DescribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeObjects")
    
    
    return
}

func NewDescribeObjectsResponse() (response *DescribeObjectsResponse) {
    response = &DescribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeObjects
// 查看防护对象列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjects(request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
    return c.DescribeObjectsWithContext(context.Background(), request)
}

// DescribeObjects
// 查看防护对象列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjectsWithContext(ctx context.Context, request *DescribeObjectsRequest) (response *DescribeObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakPointsRequest() (request *DescribePeakPointsRequest) {
    request = &DescribePeakPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePeakPoints")
    
    
    return
}

func NewDescribePeakPointsResponse() (response *DescribePeakPointsResponse) {
    response = &DescribePeakPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePeakPoints
// 查询业务和攻击概要趋势
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
func (c *Client) DescribePeakPoints(request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
    return c.DescribePeakPointsWithContext(context.Background(), request)
}

// DescribePeakPoints
// 查询业务和攻击概要趋势
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
func (c *Client) DescribePeakPointsWithContext(ctx context.Context, request *DescribePeakPointsRequest) (response *DescribePeakPointsResponse, err error) {
    if request == nil {
        request = NewDescribePeakPointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePeakPoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePeakPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakValueRequest() (request *DescribePeakValueRequest) {
    request = &DescribePeakValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePeakValue")
    
    
    return
}

func NewDescribePeakValueResponse() (response *DescribePeakValueResponse) {
    response = &DescribePeakValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePeakValue
// 获取业务和攻击概览峰值
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
func (c *Client) DescribePeakValue(request *DescribePeakValueRequest) (response *DescribePeakValueResponse, err error) {
    return c.DescribePeakValueWithContext(context.Background(), request)
}

// DescribePeakValue
// 获取业务和攻击概览峰值
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
func (c *Client) DescribePeakValueWithContext(ctx context.Context, request *DescribePeakValueRequest) (response *DescribePeakValueResponse, err error) {
    if request == nil {
        request = NewDescribePeakValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePeakValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePeakValueResponse()
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

func NewDescribePortsRequest() (request *DescribePortsRequest) {
    request = &DescribePortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribePorts")
    
    
    return
}

func NewDescribePortsResponse() (response *DescribePortsResponse) {
    response = &DescribePortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePorts
// 获取Saas型WAF防护端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribePorts(request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    return c.DescribePortsWithContext(context.Background(), request)
}

// DescribePorts
// 获取Saas型WAF防护端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribePortsWithContext(ctx context.Context, request *DescribePortsRequest) (response *DescribePortsResponse, err error) {
    if request == nil {
        request = NewDescribePortsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortsResponse()
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleLimit(request *DescribeRuleLimitRequest) (response *DescribeRuleLimitResponse, err error) {
    return c.DescribeRuleLimitWithContext(context.Background(), request)
}

// DescribeRuleLimit
// 获取各个模块具体的规格限制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewDescribeSessionRequest() (request *DescribeSessionRequest) {
    request = &DescribeSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeSession")
    
    
    return
}

func NewDescribeSessionResponse() (response *DescribeSessionResponse) {
    response = &DescribeSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSession
// Waf 会话定义查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSession(request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    return c.DescribeSessionWithContext(context.Background(), request)
}

// DescribeSession
// Waf 会话定义查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSessionWithContext(ctx context.Context, request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    if request == nil {
        request = NewDescribeSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpartaProtectionInfoRequest() (request *DescribeSpartaProtectionInfoRequest) {
    request = &DescribeSpartaProtectionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeSpartaProtectionInfo")
    
    
    return
}

func NewDescribeSpartaProtectionInfoResponse() (response *DescribeSpartaProtectionInfoResponse) {
    response = &DescribeSpartaProtectionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpartaProtectionInfo
// waf斯巴达-获取防护域名信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpartaProtectionInfo(request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
    return c.DescribeSpartaProtectionInfoWithContext(context.Background(), request)
}

// DescribeSpartaProtectionInfo
// waf斯巴达-获取防护域名信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpartaProtectionInfoWithContext(ctx context.Context, request *DescribeSpartaProtectionInfoRequest) (response *DescribeSpartaProtectionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpartaProtectionInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpartaProtectionInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpartaProtectionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTlsVersionRequest() (request *DescribeTlsVersionRequest) {
    request = &DescribeTlsVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeTlsVersion")
    
    
    return
}

func NewDescribeTlsVersionResponse() (response *DescribeTlsVersionResponse) {
    response = &DescribeTlsVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTlsVersion
// 查询用户TLS版本
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
func (c *Client) DescribeTlsVersion(request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
    return c.DescribeTlsVersionWithContext(context.Background(), request)
}

// DescribeTlsVersion
// 查询用户TLS版本
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
func (c *Client) DescribeTlsVersionWithContext(ctx context.Context, request *DescribeTlsVersionRequest) (response *DescribeTlsVersionResponse, err error) {
    if request == nil {
        request = NewDescribeTlsVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTlsVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTlsVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopAttackDomainRequest() (request *DescribeTopAttackDomainRequest) {
    request = &DescribeTopAttackDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeTopAttackDomain")
    
    
    return
}

func NewDescribeTopAttackDomainResponse() (response *DescribeTopAttackDomainResponse) {
    response = &DescribeTopAttackDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopAttackDomain
// 查询Top5的攻击域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTopAttackDomain(request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
    return c.DescribeTopAttackDomainWithContext(context.Background(), request)
}

// DescribeTopAttackDomain
// 查询Top5的攻击域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTopAttackDomainWithContext(ctx context.Context, request *DescribeTopAttackDomainRequest) (response *DescribeTopAttackDomainResponse, err error) {
    if request == nil {
        request = NewDescribeTopAttackDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopAttackDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopAttackDomainResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUserCdcClbWafRegions(request *DescribeUserCdcClbWafRegionsRequest) (response *DescribeUserCdcClbWafRegionsResponse, err error) {
    return c.DescribeUserCdcClbWafRegionsWithContext(context.Background(), request)
}

// DescribeUserCdcClbWafRegions
// 在CDC场景下，负载均衡型WAF的添加、编辑域名配置的时候，需要展示CDC负载均衡型WAF（cdc-clb-waf)支持的地域列表，通过DescribeUserCdcClbWafRegions既可以获得当前对客户已经开放的地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewDescribeUserDomainInfoRequest() (request *DescribeUserDomainInfoRequest) {
    request = &DescribeUserDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserDomainInfo")
    
    
    return
}

func NewDescribeUserDomainInfoResponse() (response *DescribeUserDomainInfoResponse) {
    response = &DescribeUserDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDomainInfo
// 查询saas和clb的域名信息
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
func (c *Client) DescribeUserDomainInfo(request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
    return c.DescribeUserDomainInfoWithContext(context.Background(), request)
}

// DescribeUserDomainInfo
// 查询saas和clb的域名信息
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
func (c *Client) DescribeUserDomainInfoWithContext(ctx context.Context, request *DescribeUserDomainInfoRequest) (response *DescribeUserDomainInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserLevelRequest() (request *DescribeUserLevelRequest) {
    request = &DescribeUserLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserLevel")
    
    
    return
}

func NewDescribeUserLevelResponse() (response *DescribeUserLevelResponse) {
    response = &DescribeUserLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserLevel
// 获取用户防护规则等级
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserLevel(request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
    return c.DescribeUserLevelWithContext(context.Background(), request)
}

// DescribeUserLevel
// 获取用户防护规则等级
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeUserLevelWithContext(ctx context.Context, request *DescribeUserLevelRequest) (response *DescribeUserLevelResponse, err error) {
    if request == nil {
        request = NewDescribeUserLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserLevelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserSignatureRuleRequest() (request *DescribeUserSignatureRuleRequest) {
    request = &DescribeUserSignatureRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserSignatureRule")
    
    
    return
}

func NewDescribeUserSignatureRuleResponse() (response *DescribeUserSignatureRuleResponse) {
    response = &DescribeUserSignatureRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserSignatureRule
// 获取用户特征规则列表
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
func (c *Client) DescribeUserSignatureRule(request *DescribeUserSignatureRuleRequest) (response *DescribeUserSignatureRuleResponse, err error) {
    return c.DescribeUserSignatureRuleWithContext(context.Background(), request)
}

// DescribeUserSignatureRule
// 获取用户特征规则列表
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
func (c *Client) DescribeUserSignatureRuleWithContext(ctx context.Context, request *DescribeUserSignatureRuleRequest) (response *DescribeUserSignatureRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUserSignatureRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserSignatureRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserSignatureRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVipInfoRequest() (request *DescribeVipInfoRequest) {
    request = &DescribeVipInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeVipInfo")
    
    
    return
}

func NewDescribeVipInfoResponse() (response *DescribeVipInfoResponse) {
    response = &DescribeVipInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVipInfo
// 根据过滤条件查询VIP信息
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
func (c *Client) DescribeVipInfo(request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
    return c.DescribeVipInfoWithContext(context.Background(), request)
}

// DescribeVipInfo
// 根据过滤条件查询VIP信息
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
func (c *Client) DescribeVipInfoWithContext(ctx context.Context, request *DescribeVipInfoRequest) (response *DescribeVipInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVipInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVipInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVipInfoResponse()
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

func NewDescribeWafInfoRequest() (request *DescribeWafInfoRequest) {
    request = &DescribeWafInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWafInfo")
    
    
    return
}

func NewDescribeWafInfoResponse() (response *DescribeWafInfoResponse) {
    response = &DescribeWafInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWafInfo
// 获取负载均衡绑定的WAF信息，可以根据租户负载均衡实例ID、负载均衡监听器ID、负载均衡的域名信息来查询对应绑定的 Waf的状态信息。
//
// 查询的范围：负载均衡实例ID、负载均衡实例ID+监听器ID、负载均衡实例ID+监听器ID+域名。
//
// 可能的错误码：ResourceNotFound（没有找到对应的资源）、UnsupportedRegion（目前clb-waf只支持北京、广州、上海、成都、重庆、香港地域）。
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
func (c *Client) DescribeWafInfo(request *DescribeWafInfoRequest) (response *DescribeWafInfoResponse, err error) {
    return c.DescribeWafInfoWithContext(context.Background(), request)
}

// DescribeWafInfo
// 获取负载均衡绑定的WAF信息，可以根据租户负载均衡实例ID、负载均衡监听器ID、负载均衡的域名信息来查询对应绑定的 Waf的状态信息。
//
// 查询的范围：负载均衡实例ID、负载均衡实例ID+监听器ID、负载均衡实例ID+监听器ID+域名。
//
// 可能的错误码：ResourceNotFound（没有找到对应的资源）、UnsupportedRegion（目前clb-waf只支持北京、广州、上海、成都、重庆、香港地域）。
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
func (c *Client) DescribeWafInfoWithContext(ctx context.Context, request *DescribeWafInfoRequest) (response *DescribeWafInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWafInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWafInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWafInfoResponse()
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
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWafThreatenIntelligence(request *DescribeWafThreatenIntelligenceRequest) (response *DescribeWafThreatenIntelligenceResponse, err error) {
    return c.DescribeWafThreatenIntelligenceWithContext(context.Background(), request)
}

// DescribeWafThreatenIntelligence
// 描述WAF威胁情报封禁模块配置详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
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

func NewDescribeWebshellStatusRequest() (request *DescribeWebshellStatusRequest) {
    request = &DescribeWebshellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "DescribeWebshellStatus")
    
    
    return
}

func NewDescribeWebshellStatusResponse() (response *DescribeWebshellStatusResponse) {
    response = &DescribeWebshellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebshellStatus
// 获取域名的webshell状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWebshellStatus(request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
    return c.DescribeWebshellStatusWithContext(context.Background(), request)
}

// DescribeWebshellStatus
// 获取域名的webshell状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeWebshellStatusWithContext(ctx context.Context, request *DescribeWebshellStatusRequest) (response *DescribeWebshellStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWebshellStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebshellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebshellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewFreshAntiFakeUrlRequest() (request *FreshAntiFakeUrlRequest) {
    request = &FreshAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "FreshAntiFakeUrl")
    
    
    return
}

func NewFreshAntiFakeUrlResponse() (response *FreshAntiFakeUrlResponse) {
    response = &FreshAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreshAntiFakeUrl
// 刷新防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) FreshAntiFakeUrl(request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
    return c.FreshAntiFakeUrlWithContext(context.Background(), request)
}

// FreshAntiFakeUrl
// 刷新防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) FreshAntiFakeUrlWithContext(ctx context.Context, request *FreshAntiFakeUrlRequest) (response *FreshAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewFreshAntiFakeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreshAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreshAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateDealsAndPayNewRequest() (request *GenerateDealsAndPayNewRequest) {
    request = &GenerateDealsAndPayNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GenerateDealsAndPayNew")
    
    
    return
}

func NewGenerateDealsAndPayNewResponse() (response *GenerateDealsAndPayNewResponse) {
    response = &GenerateDealsAndPayNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateDealsAndPayNew
// 计费资源购买、续费下单接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
func (c *Client) GenerateDealsAndPayNew(request *GenerateDealsAndPayNewRequest) (response *GenerateDealsAndPayNewResponse, err error) {
    return c.GenerateDealsAndPayNewWithContext(context.Background(), request)
}

// GenerateDealsAndPayNew
// 计费资源购买、续费下单接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
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
func (c *Client) GenerateDealsAndPayNewWithContext(ctx context.Context, request *GenerateDealsAndPayNewRequest) (response *GenerateDealsAndPayNewResponse, err error) {
    if request == nil {
        request = NewGenerateDealsAndPayNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateDealsAndPayNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateDealsAndPayNewResponse()
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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

func NewGetAttackHistogramRequest() (request *GetAttackHistogramRequest) {
    request = &GetAttackHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetAttackHistogram")
    
    
    return
}

func NewGetAttackHistogramResponse() (response *GetAttackHistogramResponse) {
    response = &GetAttackHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAttackHistogram
// 生成攻击日志的产生时间柱状图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackHistogram(request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
    return c.GetAttackHistogramWithContext(context.Background(), request)
}

// GetAttackHistogram
// 生成攻击日志的产生时间柱状图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackHistogramWithContext(ctx context.Context, request *GetAttackHistogramRequest) (response *GetAttackHistogramResponse, err error) {
    if request == nil {
        request = NewGetAttackHistogramRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttackHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttackHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttackTotalCountRequest() (request *GetAttackTotalCountRequest) {
    request = &GetAttackTotalCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetAttackTotalCount")
    
    
    return
}

func NewGetAttackTotalCountResponse() (response *GetAttackTotalCountResponse) {
    response = &GetAttackTotalCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAttackTotalCount
// 按照条件查询展示攻击总次数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackTotalCount(request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
    return c.GetAttackTotalCountWithContext(context.Background(), request)
}

// GetAttackTotalCount
// 按照条件查询展示攻击总次数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetAttackTotalCountWithContext(ctx context.Context, request *GetAttackTotalCountRequest) (response *GetAttackTotalCountResponse, err error) {
    if request == nil {
        request = NewGetAttackTotalCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttackTotalCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttackTotalCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceQpsLimitRequest() (request *GetInstanceQpsLimitRequest) {
    request = &GetInstanceQpsLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "GetInstanceQpsLimit")
    
    
    return
}

func NewGetInstanceQpsLimitResponse() (response *GetInstanceQpsLimitResponse) {
    response = &GetInstanceQpsLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetInstanceQpsLimit
// 获取套餐实例的弹性qps上限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetInstanceQpsLimit(request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
    return c.GetInstanceQpsLimitWithContext(context.Background(), request)
}

// GetInstanceQpsLimit
// 获取套餐实例的弹性qps上限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_SQLSYNTAXERR = "InvalidParameter.SQLSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
func (c *Client) GetInstanceQpsLimitWithContext(ctx context.Context, request *GetInstanceQpsLimitRequest) (response *GetInstanceQpsLimitResponse, err error) {
    if request == nil {
        request = NewGetInstanceQpsLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetInstanceQpsLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetInstanceQpsLimitResponse()
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
// 本接口用于修改访问日志保存期限及大字段是否存储
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
// 本接口用于修改访问日志保存期限及大字段是否存储
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

func NewModifyAntiFakeUrlRequest() (request *ModifyAntiFakeUrlRequest) {
    request = &ModifyAntiFakeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrl")
    
    
    return
}

func NewModifyAntiFakeUrlResponse() (response *ModifyAntiFakeUrlResponse) {
    response = &ModifyAntiFakeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiFakeUrl
// 编辑防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiFakeUrl(request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
    return c.ModifyAntiFakeUrlWithContext(context.Background(), request)
}

// ModifyAntiFakeUrl
// 编辑防篡改url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiFakeUrlWithContext(ctx context.Context, request *ModifyAntiFakeUrlRequest) (response *ModifyAntiFakeUrlResponse, err error) {
    if request == nil {
        request = NewModifyAntiFakeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiFakeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiFakeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiFakeUrlStatusRequest() (request *ModifyAntiFakeUrlStatusRequest) {
    request = &ModifyAntiFakeUrlStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiFakeUrlStatus")
    
    
    return
}

func NewModifyAntiFakeUrlStatusResponse() (response *ModifyAntiFakeUrlStatusResponse) {
    response = &ModifyAntiFakeUrlStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiFakeUrlStatus
// 切换防篡改开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAntiFakeUrlStatus(request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
    return c.ModifyAntiFakeUrlStatusWithContext(context.Background(), request)
}

// ModifyAntiFakeUrlStatus
// 切换防篡改开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAntiFakeUrlStatusWithContext(ctx context.Context, request *ModifyAntiFakeUrlStatusRequest) (response *ModifyAntiFakeUrlStatusResponse, err error) {
    if request == nil {
        request = NewModifyAntiFakeUrlStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiFakeUrlStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiFakeUrlStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiInfoLeakRuleStatusRequest() (request *ModifyAntiInfoLeakRuleStatusRequest) {
    request = &ModifyAntiInfoLeakRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRuleStatus")
    
    
    return
}

func NewModifyAntiInfoLeakRuleStatusResponse() (response *ModifyAntiInfoLeakRuleStatusResponse) {
    response = &ModifyAntiInfoLeakRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiInfoLeakRuleStatus
// 信息防泄漏切换规则开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRuleStatus(request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
    return c.ModifyAntiInfoLeakRuleStatusWithContext(context.Background(), request)
}

// ModifyAntiInfoLeakRuleStatus
// 信息防泄漏切换规则开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRuleStatusWithContext(ctx context.Context, request *ModifyAntiInfoLeakRuleStatusRequest) (response *ModifyAntiInfoLeakRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAntiInfoLeakRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiInfoLeakRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiInfoLeakRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAntiInfoLeakRulesRequest() (request *ModifyAntiInfoLeakRulesRequest) {
    request = &ModifyAntiInfoLeakRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAntiInfoLeakRules")
    
    
    return
}

func NewModifyAntiInfoLeakRulesResponse() (response *ModifyAntiInfoLeakRulesResponse) {
    response = &ModifyAntiInfoLeakRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAntiInfoLeakRules
// 编辑信息防泄漏规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRules(request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
    return c.ModifyAntiInfoLeakRulesWithContext(context.Background(), request)
}

// ModifyAntiInfoLeakRules
// 编辑信息防泄漏规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyAntiInfoLeakRulesWithContext(ctx context.Context, request *ModifyAntiInfoLeakRulesRequest) (response *ModifyAntiInfoLeakRulesResponse, err error) {
    if request == nil {
        request = NewModifyAntiInfoLeakRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAntiInfoLeakRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAntiInfoLeakRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiAnalyzeStatusRequest() (request *ModifyApiAnalyzeStatusRequest) {
    request = &ModifyApiAnalyzeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyApiAnalyzeStatus")
    
    
    return
}

func NewModifyApiAnalyzeStatusResponse() (response *ModifyApiAnalyzeStatusResponse) {
    response = &ModifyApiAnalyzeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiAnalyzeStatus
// api分析页面开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyApiAnalyzeStatus(request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
    return c.ModifyApiAnalyzeStatusWithContext(context.Background(), request)
}

// ModifyApiAnalyzeStatus
// api分析页面开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyApiAnalyzeStatusWithContext(ctx context.Context, request *ModifyApiAnalyzeStatusRequest) (response *ModifyApiAnalyzeStatusResponse, err error) {
    if request == nil {
        request = NewModifyApiAnalyzeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiAnalyzeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiAnalyzeStatusResponse()
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
func (c *Client) ModifyAreaBanStatus(request *ModifyAreaBanStatusRequest) (response *ModifyAreaBanStatusResponse, err error) {
    return c.ModifyAreaBanStatusWithContext(context.Background(), request)
}

// ModifyAreaBanStatus
// 修改防护域名的地域封禁状态
//
// 可能返回的错误码:
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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

func NewModifyAttackWhiteRuleRequest() (request *ModifyAttackWhiteRuleRequest) {
    request = &ModifyAttackWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyAttackWhiteRule")
    
    
    return
}

func NewModifyAttackWhiteRuleResponse() (response *ModifyAttackWhiteRuleResponse) {
    response = &ModifyAttackWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttackWhiteRule
// 供用户控制台调用，修改Tiga规则引擎白名单。
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
func (c *Client) ModifyAttackWhiteRule(request *ModifyAttackWhiteRuleRequest) (response *ModifyAttackWhiteRuleResponse, err error) {
    return c.ModifyAttackWhiteRuleWithContext(context.Background(), request)
}

// ModifyAttackWhiteRule
// 供用户控制台调用，修改Tiga规则引擎白名单。
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
func (c *Client) ModifyAttackWhiteRuleWithContext(ctx context.Context, request *ModifyAttackWhiteRuleRequest) (response *ModifyAttackWhiteRuleResponse, err error) {
    if request == nil {
        request = NewModifyAttackWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAttackWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAttackWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBotStatusRequest() (request *ModifyBotStatusRequest) {
    request = &ModifyBotStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyBotStatus")
    
    
    return
}

func NewModifyBotStatusResponse() (response *ModifyBotStatusResponse) {
    response = &ModifyBotStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBotStatus
// Bot_V2 bot总开关更新
//
// 可能返回的错误码:
//  AUTHFAILURE_ERRCODENOPURCHASED = "AuthFailure.ErrCodeNoPurchased"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyBotStatus(request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
    return c.ModifyBotStatusWithContext(context.Background(), request)
}

// ModifyBotStatus
// Bot_V2 bot总开关更新
//
// 可能返回的错误码:
//  AUTHFAILURE_ERRCODENOPURCHASED = "AuthFailure.ErrCodeNoPurchased"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyBotStatusWithContext(ctx context.Context, request *ModifyBotStatusRequest) (response *ModifyBotStatusResponse, err error) {
    if request == nil {
        request = NewModifyBotStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBotStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBotStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomRuleRequest() (request *ModifyCustomRuleRequest) {
    request = &ModifyCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRule")
    
    
    return
}

func NewModifyCustomRuleResponse() (response *ModifyCustomRuleResponse) {
    response = &ModifyCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomRule
// 编辑自定义规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRule(request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
    return c.ModifyCustomRuleWithContext(context.Background(), request)
}

// ModifyCustomRule
// 编辑自定义规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleWithContext(ctx context.Context, request *ModifyCustomRuleRequest) (response *ModifyCustomRuleResponse, err error) {
    if request == nil {
        request = NewModifyCustomRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomRuleResponse()
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
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    return c.ModifyCustomRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomRuleStatus
// 开启或禁用访问控制（自定义策略）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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

func NewModifyCustomWhiteRuleRequest() (request *ModifyCustomWhiteRuleRequest) {
    request = &ModifyCustomWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRule")
    
    
    return
}

func NewModifyCustomWhiteRuleResponse() (response *ModifyCustomWhiteRuleResponse) {
    response = &ModifyCustomWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomWhiteRule
// 编辑精准白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRule(request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
    return c.ModifyCustomWhiteRuleWithContext(context.Background(), request)
}

// ModifyCustomWhiteRule
// 编辑精准白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleWithContext(ctx context.Context, request *ModifyCustomWhiteRuleRequest) (response *ModifyCustomWhiteRuleResponse, err error) {
    if request == nil {
        request = NewModifyCustomWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomWhiteRuleStatusRequest() (request *ModifyCustomWhiteRuleStatusRequest) {
    request = &ModifyCustomWhiteRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomWhiteRuleStatus")
    
    
    return
}

func NewModifyCustomWhiteRuleStatusResponse() (response *ModifyCustomWhiteRuleStatusResponse) {
    response = &ModifyCustomWhiteRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomWhiteRuleStatus
// 开启或禁用精准白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleStatus(request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
    return c.ModifyCustomWhiteRuleStatusWithContext(context.Background(), request)
}

// ModifyCustomWhiteRuleStatus
// 开启或禁用精准白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCustomWhiteRuleStatusWithContext(ctx context.Context, request *ModifyCustomWhiteRuleStatusRequest) (response *ModifyCustomWhiteRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyCustomWhiteRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomWhiteRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomWhiteRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainIpv6StatusRequest() (request *ModifyDomainIpv6StatusRequest) {
    request = &ModifyDomainIpv6StatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainIpv6Status")
    
    
    return
}

func NewModifyDomainIpv6StatusResponse() (response *ModifyDomainIpv6StatusResponse) {
    response = &ModifyDomainIpv6StatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainIpv6Status
// 切换ipv6开关
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
//  RESOURCEUNAVAILABLE_DOMAINIPV6INCONFIGERR = "ResourceUnavailable.DomainIpv6InConfigErr"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
    return c.ModifyDomainIpv6StatusWithContext(context.Background(), request)
}

// ModifyDomainIpv6Status
// 切换ipv6开关
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
//  RESOURCEUNAVAILABLE_DOMAINIPV6INCONFIGERR = "ResourceUnavailable.DomainIpv6InConfigErr"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDomainIpv6StatusWithContext(ctx context.Context, request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainIpv6StatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainIpv6Status require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainIpv6StatusResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) ModifyDomainWhiteRule(request *ModifyDomainWhiteRuleRequest) (response *ModifyDomainWhiteRuleResponse, err error) {
    return c.ModifyDomainWhiteRuleWithContext(context.Background(), request)
}

// ModifyDomainWhiteRule
// 更改某一条规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
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

func NewModifyDomainsCLSStatusRequest() (request *ModifyDomainsCLSStatusRequest) {
    request = &ModifyDomainsCLSStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyDomainsCLSStatus")
    
    
    return
}

func NewModifyDomainsCLSStatusResponse() (response *ModifyDomainsCLSStatusResponse) {
    response = &ModifyDomainsCLSStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainsCLSStatus
// 修改域名列表的访问日志开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
func (c *Client) ModifyDomainsCLSStatus(request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
    return c.ModifyDomainsCLSStatusWithContext(context.Background(), request)
}

// ModifyDomainsCLSStatus
// 修改域名列表的访问日志开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
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
func (c *Client) ModifyDomainsCLSStatusWithContext(ctx context.Context, request *ModifyDomainsCLSStatusRequest) (response *ModifyDomainsCLSStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainsCLSStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainsCLSStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainsCLSStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostRequest() (request *ModifyHostRequest) {
    request = &ModifyHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHost")
    
    
    return
}

func NewModifyHostResponse() (response *ModifyHostResponse) {
    response = &ModifyHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHost
// clb-waf编辑防护域名配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHost(request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
    return c.ModifyHostWithContext(context.Background(), request)
}

// ModifyHost
// clb-waf编辑防护域名配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyHostWithContext(ctx context.Context, request *ModifyHostRequest) (response *ModifyHostResponse, err error) {
    if request == nil {
        request = NewModifyHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostFlowModeRequest() (request *ModifyHostFlowModeRequest) {
    request = &ModifyHostFlowModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostFlowMode")
    
    
    return
}

func NewModifyHostFlowModeResponse() (response *ModifyHostFlowModeResponse) {
    response = &ModifyHostFlowModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostFlowMode
// clb-waf 设置防护域名的流量模式
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
func (c *Client) ModifyHostFlowMode(request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
    return c.ModifyHostFlowModeWithContext(context.Background(), request)
}

// ModifyHostFlowMode
// clb-waf 设置防护域名的流量模式
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
func (c *Client) ModifyHostFlowModeWithContext(ctx context.Context, request *ModifyHostFlowModeRequest) (response *ModifyHostFlowModeResponse, err error) {
    if request == nil {
        request = NewModifyHostFlowModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostFlowMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostFlowModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostModeRequest() (request *ModifyHostModeRequest) {
    request = &ModifyHostModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostMode")
    
    
    return
}

func NewModifyHostModeResponse() (response *ModifyHostModeResponse) {
    response = &ModifyHostModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostMode
// clb-waf设置防护域名防护状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostMode(request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
    return c.ModifyHostModeWithContext(context.Background(), request)
}

// ModifyHostMode
// clb-waf设置防护域名防护状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHostModeWithContext(ctx context.Context, request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
    if request == nil {
        request = NewModifyHostModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostStatusRequest() (request *ModifyHostStatusRequest) {
    request = &ModifyHostStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyHostStatus")
    
    
    return
}

func NewModifyHostStatusResponse() (response *ModifyHostStatusResponse) {
    response = &ModifyHostStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostStatus
// clb-waf 设置防护域名WAF开关
//
// 支持批量操作。
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
func (c *Client) ModifyHostStatus(request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
    return c.ModifyHostStatusWithContext(context.Background(), request)
}

// ModifyHostStatus
// clb-waf 设置防护域名WAF开关
//
// 支持批量操作。
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
func (c *Client) ModifyHostStatusWithContext(ctx context.Context, request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
    if request == nil {
        request = NewModifyHostStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceElasticModeRequest() (request *ModifyInstanceElasticModeRequest) {
    request = &ModifyInstanceElasticModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceElasticMode")
    
    
    return
}

func NewModifyInstanceElasticModeResponse() (response *ModifyInstanceElasticModeResponse) {
    response = &ModifyInstanceElasticModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceElasticMode
// 修改实例的QPS弹性计费开关
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
func (c *Client) ModifyInstanceElasticMode(request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
    return c.ModifyInstanceElasticModeWithContext(context.Background(), request)
}

// ModifyInstanceElasticMode
// 修改实例的QPS弹性计费开关
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
func (c *Client) ModifyInstanceElasticModeWithContext(ctx context.Context, request *ModifyInstanceElasticModeRequest) (response *ModifyInstanceElasticModeResponse, err error) {
    if request == nil {
        request = NewModifyInstanceElasticModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceElasticMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceElasticModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// 修改实例的名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// 修改实例的名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceQpsLimitRequest() (request *ModifyInstanceQpsLimitRequest) {
    request = &ModifyInstanceQpsLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceQpsLimit")
    
    
    return
}

func NewModifyInstanceQpsLimitResponse() (response *ModifyInstanceQpsLimitResponse) {
    response = &ModifyInstanceQpsLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceQpsLimit
// 设置套餐实例的弹性qps上限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceQpsLimit(request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
    return c.ModifyInstanceQpsLimitWithContext(context.Background(), request)
}

// ModifyInstanceQpsLimit
// 设置套餐实例的弹性qps上限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyInstanceQpsLimitWithContext(ctx context.Context, request *ModifyInstanceQpsLimitRequest) (response *ModifyInstanceQpsLimitResponse, err error) {
    if request == nil {
        request = NewModifyInstanceQpsLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceQpsLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceQpsLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRenewFlagRequest() (request *ModifyInstanceRenewFlagRequest) {
    request = &ModifyInstanceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyInstanceRenewFlag")
    
    
    return
}

func NewModifyInstanceRenewFlagResponse() (response *ModifyInstanceRenewFlagResponse) {
    response = &ModifyInstanceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceRenewFlag
// 修改实例的自动续费开关
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
func (c *Client) ModifyInstanceRenewFlag(request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
    return c.ModifyInstanceRenewFlagWithContext(context.Background(), request)
}

// ModifyInstanceRenewFlag
// 修改实例的自动续费开关
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
func (c *Client) ModifyInstanceRenewFlagWithContext(ctx context.Context, request *ModifyInstanceRenewFlagRequest) (response *ModifyInstanceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleStatusRequest() (request *ModifyModuleStatusRequest) {
    request = &ModifyModuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyModuleStatus")
    
    
    return
}

func NewModifyModuleStatusResponse() (response *ModifyModuleStatusResponse) {
    response = &ModifyModuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleStatus
// 设置某个domain下基础安全模块的开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyModuleStatus(request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
    return c.ModifyModuleStatusWithContext(context.Background(), request)
}

// ModifyModuleStatus
// 设置某个domain下基础安全模块的开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyModuleStatusWithContext(ctx context.Context, request *ModifyModuleStatusRequest) (response *ModifyModuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyModuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyObjectRequest() (request *ModifyObjectRequest) {
    request = &ModifyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyObject")
    
    
    return
}

func NewModifyObjectResponse() (response *ModifyObjectResponse) {
    response = &ModifyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyObject
// 修改防护对象
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyObject(request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
    return c.ModifyObjectWithContext(context.Background(), request)
}

// ModifyObject
// 修改防护对象
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyObjectWithContext(ctx context.Context, request *ModifyObjectRequest) (response *ModifyObjectResponse, err error) {
    if request == nil {
        request = NewModifyObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProtectionStatusRequest() (request *ModifyProtectionStatusRequest) {
    request = &ModifyProtectionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyProtectionStatus")
    
    
    return
}

func NewModifyProtectionStatusResponse() (response *ModifyProtectionStatusResponse) {
    response = &ModifyProtectionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProtectionStatus
// waf斯巴达-waf开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProtectionStatus(request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
    return c.ModifyProtectionStatusWithContext(context.Background(), request)
}

// ModifyProtectionStatus
// waf斯巴达-waf开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProtectionStatusWithContext(ctx context.Context, request *ModifyProtectionStatusRequest) (response *ModifyProtectionStatusResponse, err error) {
    if request == nil {
        request = NewModifyProtectionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProtectionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProtectionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpartaProtectionRequest() (request *ModifySpartaProtectionRequest) {
    request = &ModifySpartaProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtection")
    
    
    return
}

func NewModifySpartaProtectionResponse() (response *ModifySpartaProtectionResponse) {
    response = &ModifySpartaProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpartaProtection
// 修改域名配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_SUPPORTTLSCONFFAILED = "InvalidParameter.SupportTLSConfFailed"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
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
func (c *Client) ModifySpartaProtection(request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
    return c.ModifySpartaProtectionWithContext(context.Background(), request)
}

// ModifySpartaProtection
// 修改域名配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ASYNCHRONOUSCALLFAILED = "InternalError.AsynchronousCallFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATIONPARAMETERERR = "InvalidParameter.CertificationParameterErr"
//  INVALIDPARAMETER_PORTPARAMETERERR = "InvalidParameter.PortParameterErr"
//  INVALIDPARAMETER_PROTECTIONDOMAINPARAMETERERR = "InvalidParameter.ProtectionDomainParameterErr"
//  INVALIDPARAMETER_SUPPORTTLSCONFFAILED = "InvalidParameter.SupportTLSConfFailed"
//  INVALIDPARAMETER_TLSPARAMETERERR = "InvalidParameter.TLSParameterErr"
//  INVALIDPARAMETER_UPSTREAMPARAMETERERR = "InvalidParameter.UpstreamParameterErr"
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
func (c *Client) ModifySpartaProtectionWithContext(ctx context.Context, request *ModifySpartaProtectionRequest) (response *ModifySpartaProtectionResponse, err error) {
    if request == nil {
        request = NewModifySpartaProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpartaProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpartaProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpartaProtectionModeRequest() (request *ModifySpartaProtectionModeRequest) {
    request = &ModifySpartaProtectionModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifySpartaProtectionMode")
    
    
    return
}

func NewModifySpartaProtectionModeResponse() (response *ModifySpartaProtectionModeResponse) {
    response = &ModifySpartaProtectionModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySpartaProtectionMode
// 设置waf防护状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySpartaProtectionMode(request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
    return c.ModifySpartaProtectionModeWithContext(context.Background(), request)
}

// ModifySpartaProtectionMode
// 设置waf防护状态
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySpartaProtectionModeWithContext(ctx context.Context, request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
    if request == nil {
        request = NewModifySpartaProtectionModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpartaProtectionMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpartaProtectionModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserLevelRequest() (request *ModifyUserLevelRequest) {
    request = &ModifyUserLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyUserLevel")
    
    
    return
}

func NewModifyUserLevelResponse() (response *ModifyUserLevelResponse) {
    response = &ModifyUserLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserLevel
// 修改用户防护规则等级
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserLevel(request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
    return c.ModifyUserLevelWithContext(context.Background(), request)
}

// ModifyUserLevel
// 修改用户防护规则等级
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserLevelWithContext(ctx context.Context, request *ModifyUserLevelRequest) (response *ModifyUserLevelResponse, err error) {
    if request == nil {
        request = NewModifyUserLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserSignatureRuleRequest() (request *ModifyUserSignatureRuleRequest) {
    request = &ModifyUserSignatureRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyUserSignatureRule")
    
    
    return
}

func NewModifyUserSignatureRuleResponse() (response *ModifyUserSignatureRuleResponse) {
    response = &ModifyUserSignatureRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserSignatureRule
// 修改用户防护规则，开启关闭具体的某条规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserSignatureRule(request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
    return c.ModifyUserSignatureRuleWithContext(context.Background(), request)
}

// ModifyUserSignatureRule
// 修改用户防护规则，开启关闭具体的某条规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyUserSignatureRuleWithContext(ctx context.Context, request *ModifyUserSignatureRuleRequest) (response *ModifyUserSignatureRuleResponse, err error) {
    if request == nil {
        request = NewModifyUserSignatureRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserSignatureRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserSignatureRuleResponse()
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
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWafThreatenIntelligence(request *ModifyWafThreatenIntelligenceRequest) (response *ModifyWafThreatenIntelligenceResponse, err error) {
    return c.ModifyWafThreatenIntelligenceWithContext(context.Background(), request)
}

// ModifyWafThreatenIntelligence
// 配置WAF威胁情报封禁模块详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewModifyWebshellStatusRequest() (request *ModifyWebshellStatusRequest) {
    request = &ModifyWebshellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "ModifyWebshellStatus")
    
    
    return
}

func NewModifyWebshellStatusResponse() (response *ModifyWebshellStatusResponse) {
    response = &ModifyWebshellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWebshellStatus
// 设置域名的webshell状态。
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWebshellStatus(request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
    return c.ModifyWebshellStatusWithContext(context.Background(), request)
}

// ModifyWebshellStatus
// 设置域名的webshell状态。
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyWebshellStatusWithContext(ctx context.Context, request *ModifyWebshellStatusRequest) (response *ModifyWebshellStatusResponse, err error) {
    if request == nil {
        request = NewModifyWebshellStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebshellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebshellStatusResponse()
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
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

func NewRefreshAccessCheckResultRequest() (request *RefreshAccessCheckResultRequest) {
    request = &RefreshAccessCheckResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "RefreshAccessCheckResult")
    
    
    return
}

func NewRefreshAccessCheckResultResponse() (response *RefreshAccessCheckResultResponse) {
    response = &RefreshAccessCheckResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefreshAccessCheckResult
// 刷新接入检查的结果，后台会生成接入检查任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefreshAccessCheckResult(request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
    return c.RefreshAccessCheckResultWithContext(context.Background(), request)
}

// RefreshAccessCheckResult
// 刷新接入检查的结果，后台会生成接入检查任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefreshAccessCheckResultWithContext(ctx context.Context, request *RefreshAccessCheckResultRequest) (response *RefreshAccessCheckResultResponse, err error) {
    if request == nil {
        request = NewRefreshAccessCheckResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshAccessCheckResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshAccessCheckResultResponse()
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
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
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
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCERTIFICATE = "InvalidParameter.InvalidCertificate"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYCERTBYSSLIDFAILED = "InvalidParameter.QueryCertBySSLIDFailed"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
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
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchAttackLog(request *SearchAttackLogRequest) (response *SearchAttackLogResponse, err error) {
    return c.SearchAttackLogWithContext(context.Background(), request)
}

// SearchAttackLog
// 新版本CLS接口存在参数变化，query改成了query_string支持lucence语法接口搜索查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSDBOPERATIONFAILED = "FailedOperation.CLSDBOperationFailed"
//  FAILEDOPERATION_CLSINTERNALERROR = "FailedOperation.CLSInternalError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERR = "InternalError.UnknownErr"
//  INVALIDPARAMETER_LOGICERR = "InvalidParameter.LogicErr"
//  INVALIDPARAMETER_QUERYSTRINGSYNTAXERR = "InvalidParameter.QueryStringSyntaxErr"
//  INVALIDPARAMETER_TYPEMISMATCH = "InvalidParameter.TypeMismatch"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
func (c *Client) SwitchDomainRules(request *SwitchDomainRulesRequest) (response *SwitchDomainRulesResponse, err error) {
    return c.SwitchDomainRulesWithContext(context.Background(), request)
}

// SwitchDomainRules
// 切换域名的规则开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MYSQLDBOPERATIONFAILED = "FailedOperation.MysqlDBOperationFailed"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDREQUEST = "UnsupportedOperation.InvalidRequest"
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

func NewSwitchElasticModeRequest() (request *SwitchElasticModeRequest) {
    request = &SwitchElasticModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "SwitchElasticMode")
    
    
    return
}

func NewSwitchElasticModeResponse() (response *SwitchElasticModeResponse) {
    response = &SwitchElasticModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchElasticMode
// 切换弹性的开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SwitchElasticMode(request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
    return c.SwitchElasticModeWithContext(context.Background(), request)
}

// SwitchElasticMode
// 切换弹性的开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SwitchElasticModeWithContext(ctx context.Context, request *SwitchElasticModeRequest) (response *SwitchElasticModeResponse, err error) {
    if request == nil {
        request = NewSwitchElasticModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchElasticMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchElasticModeResponse()
    err = c.Send(request, response)
    return
}

func NewUpsertCCRuleRequest() (request *UpsertCCRuleRequest) {
    request = &UpsertCCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "UpsertCCRule")
    
    
    return
}

func NewUpsertCCRuleResponse() (response *UpsertCCRuleResponse) {
    response = &UpsertCCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpsertCCRule
// Waf  CC V2 Upsert接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertCCRule(request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
    return c.UpsertCCRuleWithContext(context.Background(), request)
}

// UpsertCCRule
// Waf  CC V2 Upsert接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertCCRuleWithContext(ctx context.Context, request *UpsertCCRuleRequest) (response *UpsertCCRuleResponse, err error) {
    if request == nil {
        request = NewUpsertCCRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpsertCCRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpsertCCRuleResponse()
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
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertIpAccessControl(request *UpsertIpAccessControlRequest) (response *UpsertIpAccessControlResponse, err error) {
    return c.UpsertIpAccessControlWithContext(context.Background(), request)
}

// UpsertIpAccessControl
// Waf IP黑白名单Upsert接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_THENUMBEROFADDEDBLACKANDWHITELISTEXCEEDSTHEUPPERLIMIT = "FailedOperation.TheNumberOfAddedBlackAndWhiteListExceedsTheUpperLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DBErr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SPECIFICATIONERR = "LimitExceeded.SpecificationErr"
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

func NewUpsertSessionRequest() (request *UpsertSessionRequest) {
    request = &UpsertSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("waf", APIVersion, "UpsertSession")
    
    
    return
}

func NewUpsertSessionResponse() (response *UpsertSessionResponse) {
    response = &UpsertSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpsertSession
// Waf  会话定义 Upsert接口
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertSession(request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
    return c.UpsertSessionWithContext(context.Background(), request)
}

// UpsertSession
// Waf  会话定义 Upsert接口
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpsertSessionWithContext(ctx context.Context, request *UpsertSessionRequest) (response *UpsertSessionResponse, err error) {
    if request == nil {
        request = NewUpsertSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpsertSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpsertSessionResponse()
    err = c.Send(request, response)
    return
}
