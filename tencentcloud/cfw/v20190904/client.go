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

package v20190904

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-04"

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


func NewAddAcRuleRequest() (request *AddAcRuleRequest) {
    request = &AddAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddAcRule")
    
    
    return
}

func NewAddAcRuleResponse() (response *AddAcRuleResponse) {
    response = &AddAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAcRule
// 添加互联网边界规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddAcRule(request *AddAcRuleRequest) (response *AddAcRuleResponse, err error) {
    return c.AddAcRuleWithContext(context.Background(), request)
}

// AddAcRule
// 添加互联网边界规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddAcRuleWithContext(ctx context.Context, request *AddAcRuleRequest) (response *AddAcRuleResponse, err error) {
    if request == nil {
        request = NewAddAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddAclRuleRequest() (request *AddAclRuleRequest) {
    request = &AddAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddAclRule")
    
    
    return
}

func NewAddAclRuleResponse() (response *AddAclRuleResponse) {
    response = &AddAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAclRule
// 添加互联网边界访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddAclRule(request *AddAclRuleRequest) (response *AddAclRuleResponse, err error) {
    return c.AddAclRuleWithContext(context.Background(), request)
}

// AddAclRule
// 添加互联网边界访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddAclRuleWithContext(ctx context.Context, request *AddAclRuleRequest) (response *AddAclRuleResponse, err error) {
    if request == nil {
        request = NewAddAclRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEnterpriseSecurityGroupRulesRequest() (request *AddEnterpriseSecurityGroupRulesRequest) {
    request = &AddEnterpriseSecurityGroupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddEnterpriseSecurityGroupRules")
    
    
    return
}

func NewAddEnterpriseSecurityGroupRulesResponse() (response *AddEnterpriseSecurityGroupRulesResponse) {
    response = &AddEnterpriseSecurityGroupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEnterpriseSecurityGroupRules
// 创建新企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddEnterpriseSecurityGroupRules(request *AddEnterpriseSecurityGroupRulesRequest) (response *AddEnterpriseSecurityGroupRulesResponse, err error) {
    return c.AddEnterpriseSecurityGroupRulesWithContext(context.Background(), request)
}

// AddEnterpriseSecurityGroupRules
// 创建新企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddEnterpriseSecurityGroupRulesWithContext(ctx context.Context, request *AddEnterpriseSecurityGroupRulesRequest) (response *AddEnterpriseSecurityGroupRulesResponse, err error) {
    if request == nil {
        request = NewAddEnterpriseSecurityGroupRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEnterpriseSecurityGroupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEnterpriseSecurityGroupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewAddNatAcRuleRequest() (request *AddNatAcRuleRequest) {
    request = &AddNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddNatAcRule")
    
    
    return
}

func NewAddNatAcRuleResponse() (response *AddNatAcRuleResponse) {
    response = &AddNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNatAcRule
// 添加nat访问控制规则(地域必填)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddNatAcRule(request *AddNatAcRuleRequest) (response *AddNatAcRuleResponse, err error) {
    return c.AddNatAcRuleWithContext(context.Background(), request)
}

// AddNatAcRule
// 添加nat访问控制规则(地域必填)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddNatAcRuleWithContext(ctx context.Context, request *AddNatAcRuleRequest) (response *AddNatAcRuleResponse, err error) {
    if request == nil {
        request = NewAddNatAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddVpcAcRuleRequest() (request *AddVpcAcRuleRequest) {
    request = &AddVpcAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddVpcAcRule")
    
    
    return
}

func NewAddVpcAcRuleResponse() (response *AddVpcAcRuleResponse) {
    response = &AddVpcAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddVpcAcRule
// 添加VPC内网间规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddVpcAcRule(request *AddVpcAcRuleRequest) (response *AddVpcAcRuleResponse, err error) {
    return c.AddVpcAcRuleWithContext(context.Background(), request)
}

// AddVpcAcRule
// 添加VPC内网间规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) AddVpcAcRuleWithContext(ctx context.Context, request *AddVpcAcRuleRequest) (response *AddVpcAcRuleResponse, err error) {
    if request == nil {
        request = NewAddVpcAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddVpcAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddVpcAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAcRulesRequest() (request *CreateAcRulesRequest) {
    request = &CreateAcRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAcRules")
    
    
    return
}

func NewCreateAcRulesResponse() (response *CreateAcRulesResponse) {
    response = &CreateAcRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAcRules
// 创建访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRules(request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    return c.CreateAcRulesWithContext(context.Background(), request)
}

// CreateAcRules
// 创建访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRulesWithContext(ctx context.Context, request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    if request == nil {
        request = NewCreateAcRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAcRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressTemplateRequest() (request *CreateAddressTemplateRequest) {
    request = &CreateAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAddressTemplate")
    
    
    return
}

func NewCreateAddressTemplateResponse() (response *CreateAddressTemplateResponse) {
    response = &CreateAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAddressTemplate
// 创建地址模板规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    return c.CreateAddressTemplateWithContext(context.Background(), request)
}

// CreateAddressTemplate
// 创建地址模板规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAddressTemplateWithContext(ctx context.Context, request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAddressTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertCenterIsolateRequest() (request *CreateAlertCenterIsolateRequest) {
    request = &CreateAlertCenterIsolateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAlertCenterIsolate")
    
    
    return
}

func NewCreateAlertCenterIsolateResponse() (response *CreateAlertCenterIsolateResponse) {
    response = &CreateAlertCenterIsolateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertCenterIsolate
// 用户告警中心-封隔离处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterIsolate(request *CreateAlertCenterIsolateRequest) (response *CreateAlertCenterIsolateResponse, err error) {
    return c.CreateAlertCenterIsolateWithContext(context.Background(), request)
}

// CreateAlertCenterIsolate
// 用户告警中心-封隔离处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterIsolateWithContext(ctx context.Context, request *CreateAlertCenterIsolateRequest) (response *CreateAlertCenterIsolateResponse, err error) {
    if request == nil {
        request = NewCreateAlertCenterIsolateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertCenterIsolate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertCenterIsolateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertCenterOmitRequest() (request *CreateAlertCenterOmitRequest) {
    request = &CreateAlertCenterOmitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAlertCenterOmit")
    
    
    return
}

func NewCreateAlertCenterOmitResponse() (response *CreateAlertCenterOmitResponse) {
    response = &CreateAlertCenterOmitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertCenterOmit
// 用户告警中心-忽略处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterOmit(request *CreateAlertCenterOmitRequest) (response *CreateAlertCenterOmitResponse, err error) {
    return c.CreateAlertCenterOmitWithContext(context.Background(), request)
}

// CreateAlertCenterOmit
// 用户告警中心-忽略处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterOmitWithContext(ctx context.Context, request *CreateAlertCenterOmitRequest) (response *CreateAlertCenterOmitResponse, err error) {
    if request == nil {
        request = NewCreateAlertCenterOmitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertCenterOmit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertCenterOmitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertCenterRuleRequest() (request *CreateAlertCenterRuleRequest) {
    request = &CreateAlertCenterRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAlertCenterRule")
    
    
    return
}

func NewCreateAlertCenterRuleResponse() (response *CreateAlertCenterRuleResponse) {
    response = &CreateAlertCenterRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertCenterRule
// 用户告警中心-封禁、放通处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterRule(request *CreateAlertCenterRuleRequest) (response *CreateAlertCenterRuleResponse, err error) {
    return c.CreateAlertCenterRuleWithContext(context.Background(), request)
}

// CreateAlertCenterRule
// 用户告警中心-封禁、放通处置按钮
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAlertCenterRuleWithContext(ctx context.Context, request *CreateAlertCenterRuleRequest) (response *CreateAlertCenterRuleResponse, err error) {
    if request == nil {
        request = NewCreateAlertCenterRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertCenterRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertCenterRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlockIgnoreRuleListRequest() (request *CreateBlockIgnoreRuleListRequest) {
    request = &CreateBlockIgnoreRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateBlockIgnoreRuleList")
    
    
    return
}

func NewCreateBlockIgnoreRuleListResponse() (response *CreateBlockIgnoreRuleListResponse) {
    response = &CreateBlockIgnoreRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBlockIgnoreRuleList
// 批量添加入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBlockIgnoreRuleList(request *CreateBlockIgnoreRuleListRequest) (response *CreateBlockIgnoreRuleListResponse, err error) {
    return c.CreateBlockIgnoreRuleListWithContext(context.Background(), request)
}

// CreateBlockIgnoreRuleList
// 批量添加入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBlockIgnoreRuleListWithContext(ctx context.Context, request *CreateBlockIgnoreRuleListRequest) (response *CreateBlockIgnoreRuleListResponse, err error) {
    if request == nil {
        request = NewCreateBlockIgnoreRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlockIgnoreRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlockIgnoreRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlockIgnoreRuleNewRequest() (request *CreateBlockIgnoreRuleNewRequest) {
    request = &CreateBlockIgnoreRuleNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateBlockIgnoreRuleNew")
    
    
    return
}

func NewCreateBlockIgnoreRuleNewResponse() (response *CreateBlockIgnoreRuleNewResponse) {
    response = &CreateBlockIgnoreRuleNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBlockIgnoreRuleNew
// 批量添加入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBlockIgnoreRuleNew(request *CreateBlockIgnoreRuleNewRequest) (response *CreateBlockIgnoreRuleNewResponse, err error) {
    return c.CreateBlockIgnoreRuleNewWithContext(context.Background(), request)
}

// CreateBlockIgnoreRuleNew
// 批量添加入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBlockIgnoreRuleNewWithContext(ctx context.Context, request *CreateBlockIgnoreRuleNewRequest) (response *CreateBlockIgnoreRuleNewResponse, err error) {
    if request == nil {
        request = NewCreateBlockIgnoreRuleNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlockIgnoreRuleNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlockIgnoreRuleNewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChooseVpcsRequest() (request *CreateChooseVpcsRequest) {
    request = &CreateChooseVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateChooseVpcs")
    
    
    return
}

func NewCreateChooseVpcsResponse() (response *CreateChooseVpcsResponse) {
    response = &CreateChooseVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChooseVpcs
// 创建、选择vpc
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateChooseVpcs(request *CreateChooseVpcsRequest) (response *CreateChooseVpcsResponse, err error) {
    return c.CreateChooseVpcsWithContext(context.Background(), request)
}

// CreateChooseVpcs
// 创建、选择vpc
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateChooseVpcsWithContext(ctx context.Context, request *CreateChooseVpcsRequest) (response *CreateChooseVpcsResponse, err error) {
    if request == nil {
        request = NewCreateChooseVpcsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChooseVpcs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChooseVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatabaseWhiteListRulesRequest() (request *CreateDatabaseWhiteListRulesRequest) {
    request = &CreateDatabaseWhiteListRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateDatabaseWhiteListRules")
    
    
    return
}

func NewCreateDatabaseWhiteListRulesResponse() (response *CreateDatabaseWhiteListRulesResponse) {
    response = &CreateDatabaseWhiteListRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatabaseWhiteListRules
// 创建暴露数据库白名单规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDatabaseWhiteListRules(request *CreateDatabaseWhiteListRulesRequest) (response *CreateDatabaseWhiteListRulesResponse, err error) {
    return c.CreateDatabaseWhiteListRulesWithContext(context.Background(), request)
}

// CreateDatabaseWhiteListRules
// 创建暴露数据库白名单规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDatabaseWhiteListRulesWithContext(ctx context.Context, request *CreateDatabaseWhiteListRulesRequest) (response *CreateDatabaseWhiteListRulesResponse, err error) {
    if request == nil {
        request = NewCreateDatabaseWhiteListRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatabaseWhiteListRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatabaseWhiteListRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIdsWhiteRuleRequest() (request *CreateIdsWhiteRuleRequest) {
    request = &CreateIdsWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateIdsWhiteRule")
    
    
    return
}

func NewCreateIdsWhiteRuleResponse() (response *CreateIdsWhiteRuleResponse) {
    response = &CreateIdsWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIdsWhiteRule
// 创建入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateIdsWhiteRule(request *CreateIdsWhiteRuleRequest) (response *CreateIdsWhiteRuleResponse, err error) {
    return c.CreateIdsWhiteRuleWithContext(context.Background(), request)
}

// CreateIdsWhiteRule
// 创建入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateIdsWhiteRuleWithContext(ctx context.Context, request *CreateIdsWhiteRuleRequest) (response *CreateIdsWhiteRuleResponse, err error) {
    if request == nil {
        request = NewCreateIdsWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIdsWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIdsWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatFwInstanceRequest() (request *CreateNatFwInstanceRequest) {
    request = &CreateNatFwInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstance")
    
    
    return
}

func NewCreateNatFwInstanceResponse() (response *CreateNatFwInstanceResponse) {
    response = &CreateNatFwInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatFwInstance
// 创建NAT防火墙实例（Region参数必填）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateNatFwInstance(request *CreateNatFwInstanceRequest) (response *CreateNatFwInstanceResponse, err error) {
    return c.CreateNatFwInstanceWithContext(context.Background(), request)
}

// CreateNatFwInstance
// 创建NAT防火墙实例（Region参数必填）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateNatFwInstanceWithContext(ctx context.Context, request *CreateNatFwInstanceRequest) (response *CreateNatFwInstanceResponse, err error) {
    if request == nil {
        request = NewCreateNatFwInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatFwInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatFwInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatFwInstanceWithDomainRequest() (request *CreateNatFwInstanceWithDomainRequest) {
    request = &CreateNatFwInstanceWithDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstanceWithDomain")
    
    
    return
}

func NewCreateNatFwInstanceWithDomainResponse() (response *CreateNatFwInstanceWithDomainResponse) {
    response = &CreateNatFwInstanceWithDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatFwInstanceWithDomain
// 创建防火墙实例和接入域名（Region参数必填）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateNatFwInstanceWithDomain(request *CreateNatFwInstanceWithDomainRequest) (response *CreateNatFwInstanceWithDomainResponse, err error) {
    return c.CreateNatFwInstanceWithDomainWithContext(context.Background(), request)
}

// CreateNatFwInstanceWithDomain
// 创建防火墙实例和接入域名（Region参数必填）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateNatFwInstanceWithDomainWithContext(ctx context.Context, request *CreateNatFwInstanceWithDomainRequest) (response *CreateNatFwInstanceWithDomainResponse, err error) {
    if request == nil {
        request = NewCreateNatFwInstanceWithDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatFwInstanceWithDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatFwInstanceWithDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRulesRequest() (request *CreateSecurityGroupRulesRequest) {
    request = &CreateSecurityGroupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateSecurityGroupRules")
    
    
    return
}

func NewCreateSecurityGroupRulesResponse() (response *CreateSecurityGroupRulesResponse) {
    response = &CreateSecurityGroupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupRules
// 创建企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityGroupRules(request *CreateSecurityGroupRulesRequest) (response *CreateSecurityGroupRulesResponse, err error) {
    return c.CreateSecurityGroupRulesWithContext(context.Background(), request)
}

// CreateSecurityGroupRules
// 创建企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityGroupRulesWithContext(ctx context.Context, request *CreateSecurityGroupRulesRequest) (response *CreateSecurityGroupRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcFwGroupRequest() (request *CreateVpcFwGroupRequest) {
    request = &CreateVpcFwGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateVpcFwGroup")
    
    
    return
}

func NewCreateVpcFwGroupResponse() (response *CreateVpcFwGroupResponse) {
    response = &CreateVpcFwGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVpcFwGroup
// 创建VPC间防火墙(防火墙组)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateVpcFwGroup(request *CreateVpcFwGroupRequest) (response *CreateVpcFwGroupResponse, err error) {
    return c.CreateVpcFwGroupWithContext(context.Background(), request)
}

// CreateVpcFwGroup
// 创建VPC间防火墙(防火墙组)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) CreateVpcFwGroupWithContext(ctx context.Context, request *CreateVpcFwGroupRequest) (response *CreateVpcFwGroupResponse, err error) {
    if request == nil {
        request = NewCreateVpcFwGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcFwGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpcFwGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAcRuleRequest() (request *DeleteAcRuleRequest) {
    request = &DeleteAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAcRule")
    
    
    return
}

func NewDeleteAcRuleResponse() (response *DeleteAcRuleResponse) {
    response = &DeleteAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAcRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAcRule(request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
    return c.DeleteAcRuleWithContext(context.Background(), request)
}

// DeleteAcRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAcRuleWithContext(ctx context.Context, request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressTemplateRequest() (request *DeleteAddressTemplateRequest) {
    request = &DeleteAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAddressTemplate")
    
    
    return
}

func NewDeleteAddressTemplateResponse() (response *DeleteAddressTemplateResponse) {
    response = &DeleteAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAddressTemplate
// 删除地址模板规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAddressTemplate(request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
    return c.DeleteAddressTemplateWithContext(context.Background(), request)
}

// DeleteAddressTemplate
// 删除地址模板规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAddressTemplateWithContext(ctx context.Context, request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddressTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllAccessControlRuleRequest() (request *DeleteAllAccessControlRuleRequest) {
    request = &DeleteAllAccessControlRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAllAccessControlRule")
    
    
    return
}

func NewDeleteAllAccessControlRuleResponse() (response *DeleteAllAccessControlRuleResponse) {
    response = &DeleteAllAccessControlRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllAccessControlRule
// 全部删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAllAccessControlRule(request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
    return c.DeleteAllAccessControlRuleWithContext(context.Background(), request)
}

// DeleteAllAccessControlRule
// 全部删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAllAccessControlRuleWithContext(ctx context.Context, request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAllAccessControlRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllAccessControlRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllAccessControlRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlockIgnoreRuleListRequest() (request *DeleteBlockIgnoreRuleListRequest) {
    request = &DeleteBlockIgnoreRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteBlockIgnoreRuleList")
    
    
    return
}

func NewDeleteBlockIgnoreRuleListResponse() (response *DeleteBlockIgnoreRuleListResponse) {
    response = &DeleteBlockIgnoreRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBlockIgnoreRuleList
// 批量删除入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteBlockIgnoreRuleList(request *DeleteBlockIgnoreRuleListRequest) (response *DeleteBlockIgnoreRuleListResponse, err error) {
    return c.DeleteBlockIgnoreRuleListWithContext(context.Background(), request)
}

// DeleteBlockIgnoreRuleList
// 批量删除入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteBlockIgnoreRuleListWithContext(ctx context.Context, request *DeleteBlockIgnoreRuleListRequest) (response *DeleteBlockIgnoreRuleListResponse, err error) {
    if request == nil {
        request = NewDeleteBlockIgnoreRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlockIgnoreRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlockIgnoreRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlockIgnoreRuleNewRequest() (request *DeleteBlockIgnoreRuleNewRequest) {
    request = &DeleteBlockIgnoreRuleNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteBlockIgnoreRuleNew")
    
    
    return
}

func NewDeleteBlockIgnoreRuleNewResponse() (response *DeleteBlockIgnoreRuleNewResponse) {
    response = &DeleteBlockIgnoreRuleNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBlockIgnoreRuleNew
// 批量删除入侵防御封禁列表、放通列表规则（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteBlockIgnoreRuleNew(request *DeleteBlockIgnoreRuleNewRequest) (response *DeleteBlockIgnoreRuleNewResponse, err error) {
    return c.DeleteBlockIgnoreRuleNewWithContext(context.Background(), request)
}

// DeleteBlockIgnoreRuleNew
// 批量删除入侵防御封禁列表、放通列表规则（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteBlockIgnoreRuleNewWithContext(ctx context.Context, request *DeleteBlockIgnoreRuleNewRequest) (response *DeleteBlockIgnoreRuleNewResponse, err error) {
    if request == nil {
        request = NewDeleteBlockIgnoreRuleNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlockIgnoreRuleNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlockIgnoreRuleNewResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIdsWhiteRuleRequest() (request *DeleteIdsWhiteRuleRequest) {
    request = &DeleteIdsWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteIdsWhiteRule")
    
    
    return
}

func NewDeleteIdsWhiteRuleResponse() (response *DeleteIdsWhiteRuleResponse) {
    response = &DeleteIdsWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIdsWhiteRule
// 删除入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteIdsWhiteRule(request *DeleteIdsWhiteRuleRequest) (response *DeleteIdsWhiteRuleResponse, err error) {
    return c.DeleteIdsWhiteRuleWithContext(context.Background(), request)
}

// DeleteIdsWhiteRule
// 删除入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteIdsWhiteRuleWithContext(ctx context.Context, request *DeleteIdsWhiteRuleRequest) (response *DeleteIdsWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteIdsWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIdsWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIdsWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatFwInstanceRequest() (request *DeleteNatFwInstanceRequest) {
    request = &DeleteNatFwInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteNatFwInstance")
    
    
    return
}

func NewDeleteNatFwInstanceResponse() (response *DeleteNatFwInstanceResponse) {
    response = &DeleteNatFwInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNatFwInstance
// 销毁防火墙实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteNatFwInstance(request *DeleteNatFwInstanceRequest) (response *DeleteNatFwInstanceResponse, err error) {
    return c.DeleteNatFwInstanceWithContext(context.Background(), request)
}

// DeleteNatFwInstance
// 销毁防火墙实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteNatFwInstanceWithContext(ctx context.Context, request *DeleteNatFwInstanceRequest) (response *DeleteNatFwInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteNatFwInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatFwInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNatFwInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRemoteAccessDomainRequest() (request *DeleteRemoteAccessDomainRequest) {
    request = &DeleteRemoteAccessDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteRemoteAccessDomain")
    
    
    return
}

func NewDeleteRemoteAccessDomainResponse() (response *DeleteRemoteAccessDomainResponse) {
    response = &DeleteRemoteAccessDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRemoteAccessDomain
// 删除远程运维域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRemoteAccessDomain(request *DeleteRemoteAccessDomainRequest) (response *DeleteRemoteAccessDomainResponse, err error) {
    return c.DeleteRemoteAccessDomainWithContext(context.Background(), request)
}

// DeleteRemoteAccessDomain
// 删除远程运维域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRemoteAccessDomainWithContext(ctx context.Context, request *DeleteRemoteAccessDomainRequest) (response *DeleteRemoteAccessDomainResponse, err error) {
    if request == nil {
        request = NewDeleteRemoteAccessDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRemoteAccessDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRemoteAccessDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
    request = &DeleteResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteResourceGroup")
    
    
    return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
    response = &DeleteResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceGroup
// DeleteResourceGroup-资产中心资产组删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    return c.DeleteResourceGroupWithContext(context.Background(), request)
}

// DeleteResourceGroup
// DeleteResourceGroup-资产中心资产组删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteResourceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
    request = &DeleteSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupRule")
    
    
    return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
    response = &DeleteSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroupRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    return c.DeleteSecurityGroupRuleWithContext(context.Background(), request)
}

// DeleteSecurityGroupRule
// 删除规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupRuleWithContext(ctx context.Context, request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcFwGroupRequest() (request *DeleteVpcFwGroupRequest) {
    request = &DeleteVpcFwGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcFwGroup")
    
    
    return
}

func NewDeleteVpcFwGroupResponse() (response *DeleteVpcFwGroupResponse) {
    response = &DeleteVpcFwGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVpcFwGroup
// 删除防火墙(组)，或者删除其中实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteVpcFwGroup(request *DeleteVpcFwGroupRequest) (response *DeleteVpcFwGroupResponse, err error) {
    return c.DeleteVpcFwGroupWithContext(context.Background(), request)
}

// DeleteVpcFwGroup
// 删除防火墙(组)，或者删除其中实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DeleteVpcFwGroupWithContext(ctx context.Context, request *DeleteVpcFwGroupRequest) (response *DeleteVpcFwGroupResponse, err error) {
    if request == nil {
        request = NewDeleteVpcFwGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcFwGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcFwGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAcListsRequest() (request *DescribeAcListsRequest) {
    request = &DescribeAcListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAcLists")
    
    
    return
}

func NewDescribeAcListsResponse() (response *DescribeAcListsResponse) {
    response = &DescribeAcListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAcLists
// 访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAcLists(request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
    return c.DescribeAcListsWithContext(context.Background(), request)
}

// DescribeAcLists
// 访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAcListsWithContext(ctx context.Context, request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
    if request == nil {
        request = NewDescribeAcListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAcLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAcListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAclRuleRequest() (request *DescribeAclRuleRequest) {
    request = &DescribeAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAclRule")
    
    
    return
}

func NewDescribeAclRuleResponse() (response *DescribeAclRuleResponse) {
    response = &DescribeAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAclRule
// 查询互联网边界访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAclRule(request *DescribeAclRuleRequest) (response *DescribeAclRuleResponse, err error) {
    return c.DescribeAclRuleWithContext(context.Background(), request)
}

// DescribeAclRule
// 查询互联网边界访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAclRuleWithContext(ctx context.Context, request *DescribeAclRuleRequest) (response *DescribeAclRuleResponse, err error) {
    if request == nil {
        request = NewDescribeAclRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressTemplateListRequest() (request *DescribeAddressTemplateListRequest) {
    request = &DescribeAddressTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAddressTemplateList")
    
    
    return
}

func NewDescribeAddressTemplateListResponse() (response *DescribeAddressTemplateListResponse) {
    response = &DescribeAddressTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddressTemplateList
// 查询地址模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAddressTemplateList(request *DescribeAddressTemplateListRequest) (response *DescribeAddressTemplateListResponse, err error) {
    return c.DescribeAddressTemplateListWithContext(context.Background(), request)
}

// DescribeAddressTemplateList
// 查询地址模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAddressTemplateListWithContext(ctx context.Context, request *DescribeAddressTemplateListRequest) (response *DescribeAddressTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddressTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSyncRequest() (request *DescribeAssetSyncRequest) {
    request = &DescribeAssetSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssetSync")
    
    
    return
}

func NewDescribeAssetSyncResponse() (response *DescribeAssetSyncResponse) {
    response = &DescribeAssetSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetSync
// 资产同步状态查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeAssetSync(request *DescribeAssetSyncRequest) (response *DescribeAssetSyncResponse, err error) {
    return c.DescribeAssetSyncWithContext(context.Background(), request)
}

// DescribeAssetSync
// 资产同步状态查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeAssetSyncWithContext(ctx context.Context, request *DescribeAssetSyncRequest) (response *DescribeAssetSyncResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssociatedInstanceListRequest() (request *DescribeAssociatedInstanceListRequest) {
    request = &DescribeAssociatedInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssociatedInstanceList")
    
    
    return
}

func NewDescribeAssociatedInstanceListResponse() (response *DescribeAssociatedInstanceListResponse) {
    response = &DescribeAssociatedInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssociatedInstanceList
// 获取安全组关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAssociatedInstanceList(request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
    return c.DescribeAssociatedInstanceListWithContext(context.Background(), request)
}

// DescribeAssociatedInstanceList
// 获取安全组关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAssociatedInstanceListWithContext(ctx context.Context, request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssociatedInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssociatedInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssociatedInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockByIpTimesListRequest() (request *DescribeBlockByIpTimesListRequest) {
    request = &DescribeBlockByIpTimesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockByIpTimesList")
    
    
    return
}

func NewDescribeBlockByIpTimesListResponse() (response *DescribeBlockByIpTimesListResponse) {
    response = &DescribeBlockByIpTimesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockByIpTimesList
// DescribeBlockByIpTimesList 告警中心阻断IP折线图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeBlockByIpTimesList(request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
    return c.DescribeBlockByIpTimesListWithContext(context.Background(), request)
}

// DescribeBlockByIpTimesList
// DescribeBlockByIpTimesList 告警中心阻断IP折线图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeBlockByIpTimesListWithContext(ctx context.Context, request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockByIpTimesListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockByIpTimesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockByIpTimesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIgnoreListRequest() (request *DescribeBlockIgnoreListRequest) {
    request = &DescribeBlockIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockIgnoreList")
    
    
    return
}

func NewDescribeBlockIgnoreListResponse() (response *DescribeBlockIgnoreListResponse) {
    response = &DescribeBlockIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockIgnoreList
// 查询入侵防御放通封禁列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBlockIgnoreList(request *DescribeBlockIgnoreListRequest) (response *DescribeBlockIgnoreListResponse, err error) {
    return c.DescribeBlockIgnoreListWithContext(context.Background(), request)
}

// DescribeBlockIgnoreList
// 查询入侵防御放通封禁列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBlockIgnoreListWithContext(ctx context.Context, request *DescribeBlockIgnoreListRequest) (response *DescribeBlockIgnoreListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIgnoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockStaticListRequest() (request *DescribeBlockStaticListRequest) {
    request = &DescribeBlockStaticListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockStaticList")
    
    
    return
}

func NewDescribeBlockStaticListResponse() (response *DescribeBlockStaticListResponse) {
    response = &DescribeBlockStaticListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockStaticList
// DescribeBlockStaticList 告警中心柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeBlockStaticList(request *DescribeBlockStaticListRequest) (response *DescribeBlockStaticListResponse, err error) {
    return c.DescribeBlockStaticListWithContext(context.Background(), request)
}

// DescribeBlockStaticList
// DescribeBlockStaticList 告警中心柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeBlockStaticListWithContext(ctx context.Context, request *DescribeBlockStaticListRequest) (response *DescribeBlockStaticListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockStaticListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockStaticList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockStaticListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfwEipsRequest() (request *DescribeCfwEipsRequest) {
    request = &DescribeCfwEipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeCfwEips")
    
    
    return
}

func NewDescribeCfwEipsResponse() (response *DescribeCfwEipsResponse) {
    response = &DescribeCfwEipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCfwEips
// 查询防火墙弹性公网IP
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeCfwEips(request *DescribeCfwEipsRequest) (response *DescribeCfwEipsResponse, err error) {
    return c.DescribeCfwEipsWithContext(context.Background(), request)
}

// DescribeCfwEips
// 查询防火墙弹性公网IP
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeCfwEipsWithContext(ctx context.Context, request *DescribeCfwEipsRequest) (response *DescribeCfwEipsResponse, err error) {
    if request == nil {
        request = NewDescribeCfwEipsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfwEips require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfwEipsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefenseSwitchRequest() (request *DescribeDefenseSwitchRequest) {
    request = &DescribeDefenseSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeDefenseSwitch")
    
    
    return
}

func NewDescribeDefenseSwitchResponse() (response *DescribeDefenseSwitchResponse) {
    response = &DescribeDefenseSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefenseSwitch
// 获取入侵防御按钮列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDefenseSwitch(request *DescribeDefenseSwitchRequest) (response *DescribeDefenseSwitchResponse, err error) {
    return c.DescribeDefenseSwitchWithContext(context.Background(), request)
}

// DescribeDefenseSwitch
// 获取入侵防御按钮列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDefenseSwitchWithContext(ctx context.Context, request *DescribeDefenseSwitchRequest) (response *DescribeDefenseSwitchResponse, err error) {
    if request == nil {
        request = NewDescribeDefenseSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefenseSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefenseSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnterpriseSGRuleProgressRequest() (request *DescribeEnterpriseSGRuleProgressRequest) {
    request = &DescribeEnterpriseSGRuleProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeEnterpriseSGRuleProgress")
    
    
    return
}

func NewDescribeEnterpriseSGRuleProgressResponse() (response *DescribeEnterpriseSGRuleProgressResponse) {
    response = &DescribeEnterpriseSGRuleProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnterpriseSGRuleProgress
// 查询新版安全组下发进度
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSGRuleProgress(request *DescribeEnterpriseSGRuleProgressRequest) (response *DescribeEnterpriseSGRuleProgressResponse, err error) {
    return c.DescribeEnterpriseSGRuleProgressWithContext(context.Background(), request)
}

// DescribeEnterpriseSGRuleProgress
// 查询新版安全组下发进度
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSGRuleProgressWithContext(ctx context.Context, request *DescribeEnterpriseSGRuleProgressRequest) (response *DescribeEnterpriseSGRuleProgressResponse, err error) {
    if request == nil {
        request = NewDescribeEnterpriseSGRuleProgressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterpriseSGRuleProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterpriseSGRuleProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnterpriseSecurityGroupRuleRequest() (request *DescribeEnterpriseSecurityGroupRuleRequest) {
    request = &DescribeEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeEnterpriseSecurityGroupRule")
    
    
    return
}

func NewDescribeEnterpriseSecurityGroupRuleResponse() (response *DescribeEnterpriseSecurityGroupRuleResponse) {
    response = &DescribeEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnterpriseSecurityGroupRule
// 查询新企业安全组规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSecurityGroupRule(request *DescribeEnterpriseSecurityGroupRuleRequest) (response *DescribeEnterpriseSecurityGroupRuleResponse, err error) {
    return c.DescribeEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// DescribeEnterpriseSecurityGroupRule
// 查询新企业安全组规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *DescribeEnterpriseSecurityGroupRuleRequest) (response *DescribeEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDescribeEnterpriseSecurityGroupRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnterpriseSecurityGroupRuleListRequest() (request *DescribeEnterpriseSecurityGroupRuleListRequest) {
    request = &DescribeEnterpriseSecurityGroupRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeEnterpriseSecurityGroupRuleList")
    
    
    return
}

func NewDescribeEnterpriseSecurityGroupRuleListResponse() (response *DescribeEnterpriseSecurityGroupRuleListResponse) {
    response = &DescribeEnterpriseSecurityGroupRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnterpriseSecurityGroupRuleList
// 查询新企业安全组规则  从node接口迁移   原接口DescribeSecurityGroupNewList
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSecurityGroupRuleList(request *DescribeEnterpriseSecurityGroupRuleListRequest) (response *DescribeEnterpriseSecurityGroupRuleListResponse, err error) {
    return c.DescribeEnterpriseSecurityGroupRuleListWithContext(context.Background(), request)
}

// DescribeEnterpriseSecurityGroupRuleList
// 查询新企业安全组规则  从node接口迁移   原接口DescribeSecurityGroupNewList
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeEnterpriseSecurityGroupRuleListWithContext(ctx context.Context, request *DescribeEnterpriseSecurityGroupRuleListRequest) (response *DescribeEnterpriseSecurityGroupRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeEnterpriseSecurityGroupRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterpriseSecurityGroupRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterpriseSecurityGroupRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFwEdgeIpsRequest() (request *DescribeFwEdgeIpsRequest) {
    request = &DescribeFwEdgeIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwEdgeIps")
    
    
    return
}

func NewDescribeFwEdgeIpsResponse() (response *DescribeFwEdgeIpsResponse) {
    response = &DescribeFwEdgeIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFwEdgeIps
// 串行防火墙IP开关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwEdgeIps(request *DescribeFwEdgeIpsRequest) (response *DescribeFwEdgeIpsResponse, err error) {
    return c.DescribeFwEdgeIpsWithContext(context.Background(), request)
}

// DescribeFwEdgeIps
// 串行防火墙IP开关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwEdgeIpsWithContext(ctx context.Context, request *DescribeFwEdgeIpsRequest) (response *DescribeFwEdgeIpsResponse, err error) {
    if request == nil {
        request = NewDescribeFwEdgeIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFwEdgeIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFwEdgeIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFwGroupInstanceInfoRequest() (request *DescribeFwGroupInstanceInfoRequest) {
    request = &DescribeFwGroupInstanceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwGroupInstanceInfo")
    
    
    return
}

func NewDescribeFwGroupInstanceInfoResponse() (response *DescribeFwGroupInstanceInfoResponse) {
    response = &DescribeFwGroupInstanceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFwGroupInstanceInfo
// 获取租户所有VPC防火墙(组)及VPC防火墙实例卡片信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwGroupInstanceInfo(request *DescribeFwGroupInstanceInfoRequest) (response *DescribeFwGroupInstanceInfoResponse, err error) {
    return c.DescribeFwGroupInstanceInfoWithContext(context.Background(), request)
}

// DescribeFwGroupInstanceInfo
// 获取租户所有VPC防火墙(组)及VPC防火墙实例卡片信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwGroupInstanceInfoWithContext(ctx context.Context, request *DescribeFwGroupInstanceInfoRequest) (response *DescribeFwGroupInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFwGroupInstanceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFwGroupInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFwGroupInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFwSyncStatusRequest() (request *DescribeFwSyncStatusRequest) {
    request = &DescribeFwSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeFwSyncStatus")
    
    
    return
}

func NewDescribeFwSyncStatusResponse() (response *DescribeFwSyncStatusResponse) {
    response = &DescribeFwSyncStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFwSyncStatus
// 获取防火墙同步状态，一般在执行同步操作后查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwSyncStatus(request *DescribeFwSyncStatusRequest) (response *DescribeFwSyncStatusResponse, err error) {
    return c.DescribeFwSyncStatusWithContext(context.Background(), request)
}

// DescribeFwSyncStatus
// 获取防火墙同步状态，一般在执行同步操作后查询
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeFwSyncStatusWithContext(ctx context.Context, request *DescribeFwSyncStatusRequest) (response *DescribeFwSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFwSyncStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFwSyncStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFwSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGuideScanInfoRequest() (request *DescribeGuideScanInfoRequest) {
    request = &DescribeGuideScanInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideScanInfo")
    
    
    return
}

func NewDescribeGuideScanInfoResponse() (response *DescribeGuideScanInfoResponse) {
    response = &DescribeGuideScanInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGuideScanInfo
// DescribeGuideScanInfo新手引导扫描接口信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeGuideScanInfo(request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
    return c.DescribeGuideScanInfoWithContext(context.Background(), request)
}

// DescribeGuideScanInfo
// DescribeGuideScanInfo新手引导扫描接口信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeGuideScanInfoWithContext(ctx context.Context, request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGuideScanInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGuideScanInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGuideScanInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStatusListRequest() (request *DescribeIPStatusListRequest) {
    request = &DescribeIPStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeIPStatusList")
    
    
    return
}

func NewDescribeIPStatusListResponse() (response *DescribeIPStatusListResponse) {
    response = &DescribeIPStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPStatusList
// IP防护状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIPStatusList(request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
    return c.DescribeIPStatusListWithContext(context.Background(), request)
}

// DescribeIPStatusList
// IP防护状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIPStatusListWithContext(ctx context.Context, request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeIPStatusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdsWhiteRuleRequest() (request *DescribeIdsWhiteRuleRequest) {
    request = &DescribeIdsWhiteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeIdsWhiteRule")
    
    
    return
}

func NewDescribeIdsWhiteRuleResponse() (response *DescribeIdsWhiteRuleResponse) {
    response = &DescribeIdsWhiteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdsWhiteRule
// 查询入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeIdsWhiteRule(request *DescribeIdsWhiteRuleRequest) (response *DescribeIdsWhiteRuleResponse, err error) {
    return c.DescribeIdsWhiteRuleWithContext(context.Background(), request)
}

// DescribeIdsWhiteRule
// 查询入侵防御规则白名单接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeIdsWhiteRuleWithContext(ctx context.Context, request *DescribeIdsWhiteRuleRequest) (response *DescribeIdsWhiteRuleResponse, err error) {
    if request == nil {
        request = NewDescribeIdsWhiteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdsWhiteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdsWhiteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
    request = &DescribeLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeLogs")
    
    
    return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
    response = &DescribeLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogs
// 日志审计日志查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    return c.DescribeLogsWithContext(context.Background(), request)
}

// DescribeLogs
// 日志审计日志查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeLogsWithContext(ctx context.Context, request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatAcRuleRequest() (request *DescribeNatAcRuleRequest) {
    request = &DescribeNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatAcRule")
    
    
    return
}

func NewDescribeNatAcRuleResponse() (response *DescribeNatAcRuleResponse) {
    response = &DescribeNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatAcRule
// 查询NAT访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatAcRule(request *DescribeNatAcRuleRequest) (response *DescribeNatAcRuleResponse, err error) {
    return c.DescribeNatAcRuleWithContext(context.Background(), request)
}

// DescribeNatAcRule
// 查询NAT访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatAcRuleWithContext(ctx context.Context, request *DescribeNatAcRuleRequest) (response *DescribeNatAcRuleResponse, err error) {
    if request == nil {
        request = NewDescribeNatAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwDnatRuleRequest() (request *DescribeNatFwDnatRuleRequest) {
    request = &DescribeNatFwDnatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwDnatRule")
    
    
    return
}

func NewDescribeNatFwDnatRuleResponse() (response *DescribeNatFwDnatRuleResponse) {
    response = &DescribeNatFwDnatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwDnatRule
// 查询Nat防火墙Dnat规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatFwDnatRule(request *DescribeNatFwDnatRuleRequest) (response *DescribeNatFwDnatRuleResponse, err error) {
    return c.DescribeNatFwDnatRuleWithContext(context.Background(), request)
}

// DescribeNatFwDnatRule
// 查询Nat防火墙Dnat规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatFwDnatRuleWithContext(ctx context.Context, request *DescribeNatFwDnatRuleRequest) (response *DescribeNatFwDnatRuleResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwDnatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwDnatRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwDnatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInfoCountRequest() (request *DescribeNatFwInfoCountRequest) {
    request = &DescribeNatFwInfoCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInfoCount")
    
    
    return
}

func NewDescribeNatFwInfoCountResponse() (response *DescribeNatFwInfoCountResponse) {
    response = &DescribeNatFwInfoCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInfoCount
// 获取当前用户接入nat防火墙的所有子网数及natfw实例个数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInfoCount(request *DescribeNatFwInfoCountRequest) (response *DescribeNatFwInfoCountResponse, err error) {
    return c.DescribeNatFwInfoCountWithContext(context.Background(), request)
}

// DescribeNatFwInfoCount
// 获取当前用户接入nat防火墙的所有子网数及natfw实例个数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInfoCountWithContext(ctx context.Context, request *DescribeNatFwInfoCountRequest) (response *DescribeNatFwInfoCountResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInfoCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInfoCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInfoCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstanceRequest() (request *DescribeNatFwInstanceRequest) {
    request = &DescribeNatFwInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstance")
    
    
    return
}

func NewDescribeNatFwInstanceResponse() (response *DescribeNatFwInstanceResponse) {
    response = &DescribeNatFwInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstance
// DescribeNatFwInstance 获取租户所有NAT实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstance(request *DescribeNatFwInstanceRequest) (response *DescribeNatFwInstanceResponse, err error) {
    return c.DescribeNatFwInstanceWithContext(context.Background(), request)
}

// DescribeNatFwInstance
// DescribeNatFwInstance 获取租户所有NAT实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstanceWithContext(ctx context.Context, request *DescribeNatFwInstanceRequest) (response *DescribeNatFwInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstanceWithRegionRequest() (request *DescribeNatFwInstanceWithRegionRequest) {
    request = &DescribeNatFwInstanceWithRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstanceWithRegion")
    
    
    return
}

func NewDescribeNatFwInstanceWithRegionResponse() (response *DescribeNatFwInstanceWithRegionResponse) {
    response = &DescribeNatFwInstanceWithRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstanceWithRegion
// GetNatFwInstanceWithRegion 获取租户新增运维的NAT实例，带上地域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstanceWithRegion(request *DescribeNatFwInstanceWithRegionRequest) (response *DescribeNatFwInstanceWithRegionResponse, err error) {
    return c.DescribeNatFwInstanceWithRegionWithContext(context.Background(), request)
}

// DescribeNatFwInstanceWithRegion
// GetNatFwInstanceWithRegion 获取租户新增运维的NAT实例，带上地域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstanceWithRegionWithContext(ctx context.Context, request *DescribeNatFwInstanceWithRegionRequest) (response *DescribeNatFwInstanceWithRegionResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstanceWithRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstanceWithRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstanceWithRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstancesInfoRequest() (request *DescribeNatFwInstancesInfoRequest) {
    request = &DescribeNatFwInstancesInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstancesInfo")
    
    
    return
}

func NewDescribeNatFwInstancesInfoResponse() (response *DescribeNatFwInstancesInfoResponse) {
    response = &DescribeNatFwInstancesInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstancesInfo
// GetNatInstance 获取租户所有NAT实例及实例卡片信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstancesInfo(request *DescribeNatFwInstancesInfoRequest) (response *DescribeNatFwInstancesInfoResponse, err error) {
    return c.DescribeNatFwInstancesInfoWithContext(context.Background(), request)
}

// DescribeNatFwInstancesInfo
// GetNatInstance 获取租户所有NAT实例及实例卡片信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwInstancesInfoWithContext(ctx context.Context, request *DescribeNatFwInstancesInfoRequest) (response *DescribeNatFwInstancesInfoResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstancesInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstancesInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstancesInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwVpcDnsLstRequest() (request *DescribeNatFwVpcDnsLstRequest) {
    request = &DescribeNatFwVpcDnsLstRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwVpcDnsLst")
    
    
    return
}

func NewDescribeNatFwVpcDnsLstResponse() (response *DescribeNatFwVpcDnsLstResponse) {
    response = &DescribeNatFwVpcDnsLstResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwVpcDnsLst
// 展示当前natfw 实例对应的vpc dns开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwVpcDnsLst(request *DescribeNatFwVpcDnsLstRequest) (response *DescribeNatFwVpcDnsLstResponse, err error) {
    return c.DescribeNatFwVpcDnsLstWithContext(context.Background(), request)
}

// DescribeNatFwVpcDnsLst
// 展示当前natfw 实例对应的vpc dns开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeNatFwVpcDnsLstWithContext(ctx context.Context, request *DescribeNatFwVpcDnsLstRequest) (response *DescribeNatFwVpcDnsLstResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwVpcDnsLstRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwVpcDnsLst require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwVpcDnsLstResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatSwitchListRequest() (request *DescribeNatSwitchListRequest) {
    request = &DescribeNatSwitchListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatSwitchList")
    
    
    return
}

func NewDescribeNatSwitchListResponse() (response *DescribeNatSwitchListResponse) {
    response = &DescribeNatSwitchListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatSwitchList
// 查询NAT边界防火墙开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNatSwitchList(request *DescribeNatSwitchListRequest) (response *DescribeNatSwitchListResponse, err error) {
    return c.DescribeNatSwitchListWithContext(context.Background(), request)
}

// DescribeNatSwitchList
// 查询NAT边界防火墙开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNatSwitchListWithContext(ctx context.Context, request *DescribeNatSwitchListRequest) (response *DescribeNatSwitchListResponse, err error) {
    if request == nil {
        request = NewDescribeNatSwitchListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatSwitchList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatSwitchListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceGroupRequest() (request *DescribeResourceGroupRequest) {
    request = &DescribeResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroup")
    
    
    return
}

func NewDescribeResourceGroupResponse() (response *DescribeResourceGroupResponse) {
    response = &DescribeResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceGroup
// DescribeResourceGroup资产中心资产树信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeResourceGroup(request *DescribeResourceGroupRequest) (response *DescribeResourceGroupResponse, err error) {
    return c.DescribeResourceGroupWithContext(context.Background(), request)
}

// DescribeResourceGroup
// DescribeResourceGroup资产中心资产树信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeResourceGroupWithContext(ctx context.Context, request *DescribeResourceGroupRequest) (response *DescribeResourceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeResourceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceGroupNewRequest() (request *DescribeResourceGroupNewRequest) {
    request = &DescribeResourceGroupNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroupNew")
    
    
    return
}

func NewDescribeResourceGroupNewResponse() (response *DescribeResourceGroupNewResponse) {
    response = &DescribeResourceGroupNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceGroupNew
// DescribeResourceGroupNew资产中心资产树信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeResourceGroupNew(request *DescribeResourceGroupNewRequest) (response *DescribeResourceGroupNewResponse, err error) {
    return c.DescribeResourceGroupNewWithContext(context.Background(), request)
}

// DescribeResourceGroupNew
// DescribeResourceGroupNew资产中心资产树信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeResourceGroupNewWithContext(ctx context.Context, request *DescribeResourceGroupNewRequest) (response *DescribeResourceGroupNewResponse, err error) {
    if request == nil {
        request = NewDescribeResourceGroupNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceGroupNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceGroupNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleOverviewRequest() (request *DescribeRuleOverviewRequest) {
    request = &DescribeRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeRuleOverview")
    
    
    return
}

func NewDescribeRuleOverviewResponse() (response *DescribeRuleOverviewResponse) {
    response = &DescribeRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleOverview
// 查询规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuleOverview(request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
    return c.DescribeRuleOverviewWithContext(context.Background(), request)
}

// DescribeRuleOverview
// 查询规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuleOverviewWithContext(ctx context.Context, request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRuleOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupListRequest() (request *DescribeSecurityGroupListRequest) {
    request = &DescribeSecurityGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSecurityGroupList")
    
    
    return
}

func NewDescribeSecurityGroupListResponse() (response *DescribeSecurityGroupListResponse) {
    response = &DescribeSecurityGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupList
// 查询安全组规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecurityGroupList(request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
    return c.DescribeSecurityGroupListWithContext(context.Background(), request)
}

// DescribeSecurityGroupList
// 查询安全组规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecurityGroupListWithContext(ctx context.Context, request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceAssetRequest() (request *DescribeSourceAssetRequest) {
    request = &DescribeSourceAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSourceAsset")
    
    
    return
}

func NewDescribeSourceAssetResponse() (response *DescribeSourceAssetResponse) {
    response = &DescribeSourceAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceAsset
// DescribeSourceAsset-查询全部资产信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSourceAsset(request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
    return c.DescribeSourceAssetWithContext(context.Background(), request)
}

// DescribeSourceAsset
// DescribeSourceAsset-查询全部资产信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSourceAssetWithContext(ctx context.Context, request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
    if request == nil {
        request = NewDescribeSourceAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSwitchListsRequest() (request *DescribeSwitchListsRequest) {
    request = &DescribeSwitchListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchLists")
    
    
    return
}

func NewDescribeSwitchListsResponse() (response *DescribeSwitchListsResponse) {
    response = &DescribeSwitchListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSwitchLists
// 防火墙开关列表，已废弃，请使用DescribeFwEdgeIps
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchLists(request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    return c.DescribeSwitchListsWithContext(context.Background(), request)
}

// DescribeSwitchLists
// 防火墙开关列表，已废弃，请使用DescribeFwEdgeIps
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchListsWithContext(ctx context.Context, request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    if request == nil {
        request = NewDescribeSwitchListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSwitchLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSwitchListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogInfoRequest() (request *DescribeTLogInfoRequest) {
    request = &DescribeTLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogInfo")
    
    
    return
}

func NewDescribeTLogInfoResponse() (response *DescribeTLogInfoResponse) {
    response = &DescribeTLogInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTLogInfo
// DescribeTLogInfo告警中心概况
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeTLogInfo(request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
    return c.DescribeTLogInfoWithContext(context.Background(), request)
}

// DescribeTLogInfo
// DescribeTLogInfo告警中心概况
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeTLogInfoWithContext(ctx context.Context, request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTLogInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTLogInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTLogInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogIpListRequest() (request *DescribeTLogIpListRequest) {
    request = &DescribeTLogIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogIpList")
    
    
    return
}

func NewDescribeTLogIpListResponse() (response *DescribeTLogIpListResponse) {
    response = &DescribeTLogIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTLogIpList
// DescribeTLogIpList告警中心IP柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeTLogIpList(request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
    return c.DescribeTLogIpListWithContext(context.Background(), request)
}

// DescribeTLogIpList
// DescribeTLogIpList告警中心IP柱形图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeTLogIpListWithContext(ctx context.Context, request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
    if request == nil {
        request = NewDescribeTLogIpListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTLogIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTLogIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableStatusRequest() (request *DescribeTableStatusRequest) {
    request = &DescribeTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTableStatus")
    
    
    return
}

func NewDescribeTableStatusResponse() (response *DescribeTableStatusResponse) {
    response = &DescribeTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableStatus
// 查询规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableStatus(request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
    return c.DescribeTableStatusWithContext(context.Background(), request)
}

// DescribeTableStatus
// 查询规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableStatusWithContext(ctx context.Context, request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTableStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnHandleEventTabListRequest() (request *DescribeUnHandleEventTabListRequest) {
    request = &DescribeUnHandleEventTabListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeUnHandleEventTabList")
    
    
    return
}

func NewDescribeUnHandleEventTabListResponse() (response *DescribeUnHandleEventTabListResponse) {
    response = &DescribeUnHandleEventTabListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnHandleEventTabList
// DescribeUnHandleEventTabList 告警中心伪攻击链事件未处置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeUnHandleEventTabList(request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
    return c.DescribeUnHandleEventTabListWithContext(context.Background(), request)
}

// DescribeUnHandleEventTabList
// DescribeUnHandleEventTabList 告警中心伪攻击链事件未处置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeUnHandleEventTabListWithContext(ctx context.Context, request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
    if request == nil {
        request = NewDescribeUnHandleEventTabListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnHandleEventTabList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnHandleEventTabListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcAcRuleRequest() (request *DescribeVpcAcRuleRequest) {
    request = &DescribeVpcAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcAcRule")
    
    
    return
}

func NewDescribeVpcAcRuleResponse() (response *DescribeVpcAcRuleResponse) {
    response = &DescribeVpcAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcAcRule
// 查询内网间访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcAcRule(request *DescribeVpcAcRuleRequest) (response *DescribeVpcAcRuleResponse, err error) {
    return c.DescribeVpcAcRuleWithContext(context.Background(), request)
}

// DescribeVpcAcRule
// 查询内网间访问控制列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcAcRuleWithContext(ctx context.Context, request *DescribeVpcAcRuleRequest) (response *DescribeVpcAcRuleResponse, err error) {
    if request == nil {
        request = NewDescribeVpcAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcFwGroupSwitchRequest() (request *DescribeVpcFwGroupSwitchRequest) {
    request = &DescribeVpcFwGroupSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcFwGroupSwitch")
    
    
    return
}

func NewDescribeVpcFwGroupSwitchResponse() (response *DescribeVpcFwGroupSwitchResponse) {
    response = &DescribeVpcFwGroupSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcFwGroupSwitch
// VPC防火墙(组)开关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeVpcFwGroupSwitch(request *DescribeVpcFwGroupSwitchRequest) (response *DescribeVpcFwGroupSwitchResponse, err error) {
    return c.DescribeVpcFwGroupSwitchWithContext(context.Background(), request)
}

// DescribeVpcFwGroupSwitch
// VPC防火墙(组)开关列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) DescribeVpcFwGroupSwitchWithContext(ctx context.Context, request *DescribeVpcFwGroupSwitchRequest) (response *DescribeVpcFwGroupSwitchResponse, err error) {
    if request == nil {
        request = NewDescribeVpcFwGroupSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcFwGroupSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcFwGroupSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewExpandCfwVerticalRequest() (request *ExpandCfwVerticalRequest) {
    request = &ExpandCfwVerticalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ExpandCfwVertical")
    
    
    return
}

func NewExpandCfwVerticalResponse() (response *ExpandCfwVerticalResponse) {
    response = &ExpandCfwVerticalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExpandCfwVertical
// 防火墙垂直扩容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ExpandCfwVertical(request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
    return c.ExpandCfwVerticalWithContext(context.Background(), request)
}

// ExpandCfwVertical
// 防火墙垂直扩容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ExpandCfwVerticalWithContext(ctx context.Context, request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
    if request == nil {
        request = NewExpandCfwVerticalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandCfwVertical require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandCfwVerticalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAcRuleRequest() (request *ModifyAcRuleRequest) {
    request = &ModifyAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAcRule")
    
    
    return
}

func NewModifyAcRuleResponse() (response *ModifyAcRuleResponse) {
    response = &ModifyAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAcRule
// 修改规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAcRule(request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
    return c.ModifyAcRuleWithContext(context.Background(), request)
}

// ModifyAcRule
// 修改规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAcRuleWithContext(ctx context.Context, request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAclRuleRequest() (request *ModifyAclRuleRequest) {
    request = &ModifyAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAclRule")
    
    
    return
}

func NewModifyAclRuleResponse() (response *ModifyAclRuleResponse) {
    response = &ModifyAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAclRule
// 修改互联网边界访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAclRule(request *ModifyAclRuleRequest) (response *ModifyAclRuleResponse, err error) {
    return c.ModifyAclRuleWithContext(context.Background(), request)
}

// ModifyAclRule
// 修改互联网边界访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAclRuleWithContext(ctx context.Context, request *ModifyAclRuleRequest) (response *ModifyAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyAclRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressTemplateRequest() (request *ModifyAddressTemplateRequest) {
    request = &ModifyAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAddressTemplate")
    
    
    return
}

func NewModifyAddressTemplateResponse() (response *ModifyAddressTemplateResponse) {
    response = &ModifyAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAddressTemplate
// 修改地址模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAddressTemplate(request *ModifyAddressTemplateRequest) (response *ModifyAddressTemplateResponse, err error) {
    return c.ModifyAddressTemplateWithContext(context.Background(), request)
}

// ModifyAddressTemplate
// 修改地址模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAddressTemplateWithContext(ctx context.Context, request *ModifyAddressTemplateRequest) (response *ModifyAddressTemplateResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllPublicIPSwitchStatusRequest() (request *ModifyAllPublicIPSwitchStatusRequest) {
    request = &ModifyAllPublicIPSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllPublicIPSwitchStatus")
    
    
    return
}

func NewModifyAllPublicIPSwitchStatusResponse() (response *ModifyAllPublicIPSwitchStatusResponse) {
    response = &ModifyAllPublicIPSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllPublicIPSwitchStatus
// 互联网边界防火墙一键开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllPublicIPSwitchStatus(request *ModifyAllPublicIPSwitchStatusRequest) (response *ModifyAllPublicIPSwitchStatusResponse, err error) {
    return c.ModifyAllPublicIPSwitchStatusWithContext(context.Background(), request)
}

// ModifyAllPublicIPSwitchStatus
// 互联网边界防火墙一键开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllPublicIPSwitchStatusWithContext(ctx context.Context, request *ModifyAllPublicIPSwitchStatusRequest) (response *ModifyAllPublicIPSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllPublicIPSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllPublicIPSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllPublicIPSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllRuleStatusRequest() (request *ModifyAllRuleStatusRequest) {
    request = &ModifyAllRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllRuleStatus")
    
    
    return
}

func NewModifyAllRuleStatusResponse() (response *ModifyAllRuleStatusResponse) {
    response = &ModifyAllRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAllRuleStatus(request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
    return c.ModifyAllRuleStatusWithContext(context.Background(), request)
}

// ModifyAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAllRuleStatusWithContext(ctx context.Context, request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllVPCSwitchStatusRequest() (request *ModifyAllVPCSwitchStatusRequest) {
    request = &ModifyAllVPCSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllVPCSwitchStatus")
    
    
    return
}

func NewModifyAllVPCSwitchStatusResponse() (response *ModifyAllVPCSwitchStatusResponse) {
    response = &ModifyAllVPCSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllVPCSwitchStatus
// VPC防火墙一键开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllVPCSwitchStatus(request *ModifyAllVPCSwitchStatusRequest) (response *ModifyAllVPCSwitchStatusResponse, err error) {
    return c.ModifyAllVPCSwitchStatusWithContext(context.Background(), request)
}

// ModifyAllVPCSwitchStatus
// VPC防火墙一键开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllVPCSwitchStatusWithContext(ctx context.Context, request *ModifyAllVPCSwitchStatusRequest) (response *ModifyAllVPCSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllVPCSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllVPCSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllVPCSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetScanRequest() (request *ModifyAssetScanRequest) {
    request = &ModifyAssetScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAssetScan")
    
    
    return
}

func NewModifyAssetScanResponse() (response *ModifyAssetScanResponse) {
    response = &ModifyAssetScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetScan
// 资产扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAssetScan(request *ModifyAssetScanRequest) (response *ModifyAssetScanResponse, err error) {
    return c.ModifyAssetScanWithContext(context.Background(), request)
}

// ModifyAssetScan
// 资产扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAssetScanWithContext(ctx context.Context, request *ModifyAssetScanRequest) (response *ModifyAssetScanResponse, err error) {
    if request == nil {
        request = NewModifyAssetScanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetScanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetSyncRequest() (request *ModifyAssetSyncRequest) {
    request = &ModifyAssetSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAssetSync")
    
    
    return
}

func NewModifyAssetSyncResponse() (response *ModifyAssetSyncResponse) {
    response = &ModifyAssetSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetSync
// 资产同步
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyAssetSync(request *ModifyAssetSyncRequest) (response *ModifyAssetSyncResponse, err error) {
    return c.ModifyAssetSyncWithContext(context.Background(), request)
}

// ModifyAssetSync
// 资产同步
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyAssetSyncWithContext(ctx context.Context, request *ModifyAssetSyncRequest) (response *ModifyAssetSyncResponse, err error) {
    if request == nil {
        request = NewModifyAssetSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetSyncResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIgnoreListRequest() (request *ModifyBlockIgnoreListRequest) {
    request = &ModifyBlockIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreList")
    
    
    return
}

func NewModifyBlockIgnoreListResponse() (response *ModifyBlockIgnoreListResponse) {
    response = &ModifyBlockIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockIgnoreList
// 支持对封禁列表、放通列表如下操作：
//
// 批量增加封禁IP、放通IP/域名
//
// 批量删除封禁IP、放通IP/域名
//
// 批量修改封禁IP、放通IP/域名生效事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreList(request *ModifyBlockIgnoreListRequest) (response *ModifyBlockIgnoreListResponse, err error) {
    return c.ModifyBlockIgnoreListWithContext(context.Background(), request)
}

// ModifyBlockIgnoreList
// 支持对封禁列表、放通列表如下操作：
//
// 批量增加封禁IP、放通IP/域名
//
// 批量删除封禁IP、放通IP/域名
//
// 批量修改封禁IP、放通IP/域名生效事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreListWithContext(ctx context.Context, request *ModifyBlockIgnoreListRequest) (response *ModifyBlockIgnoreListResponse, err error) {
    if request == nil {
        request = NewModifyBlockIgnoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIgnoreRuleRequest() (request *ModifyBlockIgnoreRuleRequest) {
    request = &ModifyBlockIgnoreRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreRule")
    
    
    return
}

func NewModifyBlockIgnoreRuleResponse() (response *ModifyBlockIgnoreRuleResponse) {
    response = &ModifyBlockIgnoreRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockIgnoreRule
// 编辑单条入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreRule(request *ModifyBlockIgnoreRuleRequest) (response *ModifyBlockIgnoreRuleResponse, err error) {
    return c.ModifyBlockIgnoreRuleWithContext(context.Background(), request)
}

// ModifyBlockIgnoreRule
// 编辑单条入侵防御封禁列表、放通列表规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreRuleWithContext(ctx context.Context, request *ModifyBlockIgnoreRuleRequest) (response *ModifyBlockIgnoreRuleResponse, err error) {
    if request == nil {
        request = NewModifyBlockIgnoreRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockIgnoreRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockIgnoreRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIgnoreRuleNewRequest() (request *ModifyBlockIgnoreRuleNewRequest) {
    request = &ModifyBlockIgnoreRuleNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreRuleNew")
    
    
    return
}

func NewModifyBlockIgnoreRuleNewResponse() (response *ModifyBlockIgnoreRuleNewResponse) {
    response = &ModifyBlockIgnoreRuleNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockIgnoreRuleNew
// 编辑单条入侵防御封禁列表、放通列表规则（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreRuleNew(request *ModifyBlockIgnoreRuleNewRequest) (response *ModifyBlockIgnoreRuleNewResponse, err error) {
    return c.ModifyBlockIgnoreRuleNewWithContext(context.Background(), request)
}

// ModifyBlockIgnoreRuleNew
// 编辑单条入侵防御封禁列表、放通列表规则（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreRuleNewWithContext(ctx context.Context, request *ModifyBlockIgnoreRuleNewRequest) (response *ModifyBlockIgnoreRuleNewResponse, err error) {
    if request == nil {
        request = NewModifyBlockIgnoreRuleNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockIgnoreRuleNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockIgnoreRuleNewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockTopRequest() (request *ModifyBlockTopRequest) {
    request = &ModifyBlockTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockTop")
    
    
    return
}

func NewModifyBlockTopResponse() (response *ModifyBlockTopResponse) {
    response = &ModifyBlockTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockTop
// ModifyBlockTop取消置顶接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyBlockTop(request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
    return c.ModifyBlockTopWithContext(context.Background(), request)
}

// ModifyBlockTop
// ModifyBlockTop取消置顶接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyBlockTopWithContext(ctx context.Context, request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
    if request == nil {
        request = NewModifyBlockTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockTopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEWRuleStatusRequest() (request *ModifyEWRuleStatusRequest) {
    request = &ModifyEWRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEWRuleStatus")
    
    
    return
}

func NewModifyEWRuleStatusResponse() (response *ModifyEWRuleStatusResponse) {
    response = &ModifyEWRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEWRuleStatus
// 启用停用VPC间规则或Nat边界规则
//
// VPC间规则需指定EdgeId。Nat边界规则需指定地域Region与Direction。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyEWRuleStatus(request *ModifyEWRuleStatusRequest) (response *ModifyEWRuleStatusResponse, err error) {
    return c.ModifyEWRuleStatusWithContext(context.Background(), request)
}

// ModifyEWRuleStatus
// 启用停用VPC间规则或Nat边界规则
//
// VPC间规则需指定EdgeId。Nat边界规则需指定地域Region与Direction。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyEWRuleStatusWithContext(ctx context.Context, request *ModifyEWRuleStatusRequest) (response *ModifyEWRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyEWRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEWRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEWRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeIpSwitchRequest() (request *ModifyEdgeIpSwitchRequest) {
    request = &ModifyEdgeIpSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEdgeIpSwitch")
    
    
    return
}

func NewModifyEdgeIpSwitchResponse() (response *ModifyEdgeIpSwitchResponse) {
    response = &ModifyEdgeIpSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEdgeIpSwitch
// 修改边界防火墙开关(旁路、串行)
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeIpSwitch(request *ModifyEdgeIpSwitchRequest) (response *ModifyEdgeIpSwitchResponse, err error) {
    return c.ModifyEdgeIpSwitchWithContext(context.Background(), request)
}

// ModifyEdgeIpSwitch
// 修改边界防火墙开关(旁路、串行)
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEdgeIpSwitchWithContext(ctx context.Context, request *ModifyEdgeIpSwitchRequest) (response *ModifyEdgeIpSwitchResponse, err error) {
    if request == nil {
        request = NewModifyEdgeIpSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeIpSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeIpSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnterpriseSecurityDispatchStatusRequest() (request *ModifyEnterpriseSecurityDispatchStatusRequest) {
    request = &ModifyEnterpriseSecurityDispatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEnterpriseSecurityDispatchStatus")
    
    
    return
}

func NewModifyEnterpriseSecurityDispatchStatusResponse() (response *ModifyEnterpriseSecurityDispatchStatusResponse) {
    response = &ModifyEnterpriseSecurityDispatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnterpriseSecurityDispatchStatus
// 修改企业安全组下发状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnterpriseSecurityDispatchStatus(request *ModifyEnterpriseSecurityDispatchStatusRequest) (response *ModifyEnterpriseSecurityDispatchStatusResponse, err error) {
    return c.ModifyEnterpriseSecurityDispatchStatusWithContext(context.Background(), request)
}

// ModifyEnterpriseSecurityDispatchStatus
// 修改企业安全组下发状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnterpriseSecurityDispatchStatusWithContext(ctx context.Context, request *ModifyEnterpriseSecurityDispatchStatusRequest) (response *ModifyEnterpriseSecurityDispatchStatusResponse, err error) {
    if request == nil {
        request = NewModifyEnterpriseSecurityDispatchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnterpriseSecurityDispatchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnterpriseSecurityDispatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnterpriseSecurityGroupRuleRequest() (request *ModifyEnterpriseSecurityGroupRuleRequest) {
    request = &ModifyEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEnterpriseSecurityGroupRule")
    
    
    return
}

func NewModifyEnterpriseSecurityGroupRuleResponse() (response *ModifyEnterpriseSecurityGroupRuleResponse) {
    response = &ModifyEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnterpriseSecurityGroupRule
// 编辑新企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEnterpriseSecurityGroupRule(request *ModifyEnterpriseSecurityGroupRuleRequest) (response *ModifyEnterpriseSecurityGroupRuleResponse, err error) {
    return c.ModifyEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// ModifyEnterpriseSecurityGroupRule
// 编辑新企业安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *ModifyEnterpriseSecurityGroupRuleRequest) (response *ModifyEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewModifyEnterpriseSecurityGroupRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFwGroupSwitchRequest() (request *ModifyFwGroupSwitchRequest) {
    request = &ModifyFwGroupSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyFwGroupSwitch")
    
    
    return
}

func NewModifyFwGroupSwitchResponse() (response *ModifyFwGroupSwitchResponse) {
    response = &ModifyFwGroupSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFwGroupSwitch
// 修改防火墙(组)开关(支持单点模式、多点模式、全互通模式)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyFwGroupSwitch(request *ModifyFwGroupSwitchRequest) (response *ModifyFwGroupSwitchResponse, err error) {
    return c.ModifyFwGroupSwitchWithContext(context.Background(), request)
}

// ModifyFwGroupSwitch
// 修改防火墙(组)开关(支持单点模式、多点模式、全互通模式)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyFwGroupSwitchWithContext(ctx context.Context, request *ModifyFwGroupSwitchRequest) (response *ModifyFwGroupSwitchResponse, err error) {
    if request == nil {
        request = NewModifyFwGroupSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFwGroupSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFwGroupSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatAcRuleRequest() (request *ModifyNatAcRuleRequest) {
    request = &ModifyNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatAcRule")
    
    
    return
}

func NewModifyNatAcRuleResponse() (response *ModifyNatAcRuleResponse) {
    response = &ModifyNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatAcRule
// 修改NAT访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatAcRule(request *ModifyNatAcRuleRequest) (response *ModifyNatAcRuleResponse, err error) {
    return c.ModifyNatAcRuleWithContext(context.Background(), request)
}

// ModifyNatAcRule
// 修改NAT访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatAcRuleWithContext(ctx context.Context, request *ModifyNatAcRuleRequest) (response *ModifyNatAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwReSelectRequest() (request *ModifyNatFwReSelectRequest) {
    request = &ModifyNatFwReSelectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwReSelect")
    
    
    return
}

func NewModifyNatFwReSelectResponse() (response *ModifyNatFwReSelectResponse) {
    response = &ModifyNatFwReSelectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwReSelect
// 防火墙实例重新选择vpc或nat
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatFwReSelect(request *ModifyNatFwReSelectRequest) (response *ModifyNatFwReSelectResponse, err error) {
    return c.ModifyNatFwReSelectWithContext(context.Background(), request)
}

// ModifyNatFwReSelect
// 防火墙实例重新选择vpc或nat
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatFwReSelectWithContext(ctx context.Context, request *ModifyNatFwReSelectRequest) (response *ModifyNatFwReSelectResponse, err error) {
    if request == nil {
        request = NewModifyNatFwReSelectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwReSelect require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwReSelectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwSwitchRequest() (request *ModifyNatFwSwitchRequest) {
    request = &ModifyNatFwSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwSwitch")
    
    
    return
}

func NewModifyNatFwSwitchResponse() (response *ModifyNatFwSwitchResponse) {
    response = &ModifyNatFwSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwSwitch
// 修改NAT防火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatFwSwitch(request *ModifyNatFwSwitchRequest) (response *ModifyNatFwSwitchResponse, err error) {
    return c.ModifyNatFwSwitchWithContext(context.Background(), request)
}

// ModifyNatFwSwitch
// 修改NAT防火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatFwSwitchWithContext(ctx context.Context, request *ModifyNatFwSwitchRequest) (response *ModifyNatFwSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNatFwSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwVpcDnsSwitchRequest() (request *ModifyNatFwVpcDnsSwitchRequest) {
    request = &ModifyNatFwVpcDnsSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwVpcDnsSwitch")
    
    
    return
}

func NewModifyNatFwVpcDnsSwitchResponse() (response *ModifyNatFwVpcDnsSwitchResponse) {
    response = &ModifyNatFwVpcDnsSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwVpcDnsSwitch
// nat 防火墙VPC DNS 开关切换
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNatFwVpcDnsSwitch(request *ModifyNatFwVpcDnsSwitchRequest) (response *ModifyNatFwVpcDnsSwitchResponse, err error) {
    return c.ModifyNatFwVpcDnsSwitchWithContext(context.Background(), request)
}

// ModifyNatFwVpcDnsSwitch
// nat 防火墙VPC DNS 开关切换
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNatFwVpcDnsSwitchWithContext(ctx context.Context, request *ModifyNatFwVpcDnsSwitchRequest) (response *ModifyNatFwVpcDnsSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNatFwVpcDnsSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwVpcDnsSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwVpcDnsSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatInstanceRequest() (request *ModifyNatInstanceRequest) {
    request = &ModifyNatInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatInstance")
    
    
    return
}

func NewModifyNatInstanceResponse() (response *ModifyNatInstanceResponse) {
    response = &ModifyNatInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatInstance
// 编辑NAT防火墙
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatInstance(request *ModifyNatInstanceRequest) (response *ModifyNatInstanceResponse, err error) {
    return c.ModifyNatInstanceWithContext(context.Background(), request)
}

// ModifyNatInstance
// 编辑NAT防火墙
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatInstanceWithContext(ctx context.Context, request *ModifyNatInstanceRequest) (response *ModifyNatInstanceResponse, err error) {
    if request == nil {
        request = NewModifyNatInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatSequenceRulesRequest() (request *ModifyNatSequenceRulesRequest) {
    request = &ModifyNatSequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatSequenceRules")
    
    
    return
}

func NewModifyNatSequenceRulesResponse() (response *ModifyNatSequenceRulesResponse) {
    response = &ModifyNatSequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatSequenceRules
// NAT防火墙规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatSequenceRules(request *ModifyNatSequenceRulesRequest) (response *ModifyNatSequenceRulesResponse, err error) {
    return c.ModifyNatSequenceRulesWithContext(context.Background(), request)
}

// ModifyNatSequenceRules
// NAT防火墙规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyNatSequenceRulesWithContext(ctx context.Context, request *ModifyNatSequenceRulesRequest) (response *ModifyNatSequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifyNatSequenceRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatSequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatSequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceGroupRequest() (request *ModifyResourceGroupRequest) {
    request = &ModifyResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyResourceGroup")
    
    
    return
}

func NewModifyResourceGroupResponse() (response *ModifyResourceGroupResponse) {
    response = &ModifyResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceGroup
// ModifyResourceGroup-资产中心资产组信息修改
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
    return c.ModifyResourceGroupWithContext(context.Background(), request)
}

// ModifyResourceGroup
// ModifyResourceGroup-资产中心资产组信息修改
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyResourceGroupWithContext(ctx context.Context, request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
    if request == nil {
        request = NewModifyResourceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRunSyncAssetRequest() (request *ModifyRunSyncAssetRequest) {
    request = &ModifyRunSyncAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyRunSyncAsset")
    
    
    return
}

func NewModifyRunSyncAssetResponse() (response *ModifyRunSyncAssetResponse) {
    response = &ModifyRunSyncAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRunSyncAsset
// 同步资产-互联网&VPC（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRunSyncAsset(request *ModifyRunSyncAssetRequest) (response *ModifyRunSyncAssetResponse, err error) {
    return c.ModifyRunSyncAssetWithContext(context.Background(), request)
}

// ModifyRunSyncAsset
// 同步资产-互联网&VPC（新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRunSyncAssetWithContext(ctx context.Context, request *ModifyRunSyncAssetRequest) (response *ModifyRunSyncAssetResponse, err error) {
    if request == nil {
        request = NewModifyRunSyncAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRunSyncAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRunSyncAssetResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupItemRuleStatusRequest() (request *ModifySecurityGroupItemRuleStatusRequest) {
    request = &ModifySecurityGroupItemRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupItemRuleStatus")
    
    
    return
}

func NewModifySecurityGroupItemRuleStatusResponse() (response *ModifySecurityGroupItemRuleStatusResponse) {
    response = &ModifySecurityGroupItemRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupItemRuleStatus
// 启用停用单条企业安全组规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupItemRuleStatus(request *ModifySecurityGroupItemRuleStatusRequest) (response *ModifySecurityGroupItemRuleStatusResponse, err error) {
    return c.ModifySecurityGroupItemRuleStatusWithContext(context.Background(), request)
}

// ModifySecurityGroupItemRuleStatus
// 启用停用单条企业安全组规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupItemRuleStatusWithContext(ctx context.Context, request *ModifySecurityGroupItemRuleStatusRequest) (response *ModifySecurityGroupItemRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupItemRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupItemRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupItemRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupRuleRequest() (request *ModifySecurityGroupRuleRequest) {
    request = &ModifySecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupRule")
    
    
    return
}

func NewModifySecurityGroupRuleResponse() (response *ModifySecurityGroupRuleResponse) {
    response = &ModifySecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupRule
// 编辑单条安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupRule(request *ModifySecurityGroupRuleRequest) (response *ModifySecurityGroupRuleResponse, err error) {
    return c.ModifySecurityGroupRuleWithContext(context.Background(), request)
}

// ModifySecurityGroupRule
// 编辑单条安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupRuleWithContext(ctx context.Context, request *ModifySecurityGroupRuleRequest) (response *ModifySecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupSequenceRulesRequest() (request *ModifySecurityGroupSequenceRulesRequest) {
    request = &ModifySecurityGroupSequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupSequenceRules")
    
    
    return
}

func NewModifySecurityGroupSequenceRulesResponse() (response *ModifySecurityGroupSequenceRulesResponse) {
    response = &ModifySecurityGroupSequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupSequenceRules
// 企业安全组规则快速排序
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupSequenceRules(request *ModifySecurityGroupSequenceRulesRequest) (response *ModifySecurityGroupSequenceRulesResponse, err error) {
    return c.ModifySecurityGroupSequenceRulesWithContext(context.Background(), request)
}

// ModifySecurityGroupSequenceRules
// 企业安全组规则快速排序
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupSequenceRulesWithContext(ctx context.Context, request *ModifySecurityGroupSequenceRulesRequest) (response *ModifySecurityGroupSequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupSequenceRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupSequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupSequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySequenceAclRulesRequest() (request *ModifySequenceAclRulesRequest) {
    request = &ModifySequenceAclRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySequenceAclRules")
    
    
    return
}

func NewModifySequenceAclRulesResponse() (response *ModifySequenceAclRulesResponse) {
    response = &ModifySequenceAclRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySequenceAclRules
// 互联网边界规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifySequenceAclRules(request *ModifySequenceAclRulesRequest) (response *ModifySequenceAclRulesResponse, err error) {
    return c.ModifySequenceAclRulesWithContext(context.Background(), request)
}

// ModifySequenceAclRules
// 互联网边界规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifySequenceAclRulesWithContext(ctx context.Context, request *ModifySequenceAclRulesRequest) (response *ModifySequenceAclRulesResponse, err error) {
    if request == nil {
        request = NewModifySequenceAclRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySequenceAclRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySequenceAclRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySequenceRulesRequest() (request *ModifySequenceRulesRequest) {
    request = &ModifySequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySequenceRules")
    
    
    return
}

func NewModifySequenceRulesResponse() (response *ModifySequenceRulesResponse) {
    response = &ModifySequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySequenceRules
// 修改规则执行顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySequenceRules(request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
    return c.ModifySequenceRulesWithContext(context.Background(), request)
}

// ModifySequenceRules
// 修改规则执行顺序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySequenceRulesWithContext(ctx context.Context, request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifySequenceRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStorageSettingRequest() (request *ModifyStorageSettingRequest) {
    request = &ModifyStorageSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyStorageSetting")
    
    
    return
}

func NewModifyStorageSettingResponse() (response *ModifyStorageSettingResponse) {
    response = &ModifyStorageSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStorageSetting
// 日志存储设置，可以修改存储时间和清空日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyStorageSetting(request *ModifyStorageSettingRequest) (response *ModifyStorageSettingResponse, err error) {
    return c.ModifyStorageSettingWithContext(context.Background(), request)
}

// ModifyStorageSetting
// 日志存储设置，可以修改存储时间和清空日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyStorageSettingWithContext(ctx context.Context, request *ModifyStorageSettingRequest) (response *ModifyStorageSettingResponse, err error) {
    if request == nil {
        request = NewModifyStorageSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStorageSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStorageSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableStatusRequest() (request *ModifyTableStatusRequest) {
    request = &ModifyTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyTableStatus")
    
    
    return
}

func NewModifyTableStatusResponse() (response *ModifyTableStatusResponse) {
    response = &ModifyTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableStatus
// 修改规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTableStatus(request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
    return c.ModifyTableStatusWithContext(context.Background(), request)
}

// ModifyTableStatus
// 修改规则表状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTableStatusWithContext(ctx context.Context, request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
    if request == nil {
        request = NewModifyTableStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAcRuleRequest() (request *ModifyVpcAcRuleRequest) {
    request = &ModifyVpcAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcAcRule")
    
    
    return
}

func NewModifyVpcAcRuleResponse() (response *ModifyVpcAcRuleResponse) {
    response = &ModifyVpcAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcAcRule
// 修改内网间访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVpcAcRule(request *ModifyVpcAcRuleRequest) (response *ModifyVpcAcRuleResponse, err error) {
    return c.ModifyVpcAcRuleWithContext(context.Background(), request)
}

// ModifyVpcAcRule
// 修改内网间访问控制规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVpcAcRuleWithContext(ctx context.Context, request *ModifyVpcAcRuleRequest) (response *ModifyVpcAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyVpcAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcFwGroupRequest() (request *ModifyVpcFwGroupRequest) {
    request = &ModifyVpcFwGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwGroup")
    
    
    return
}

func NewModifyVpcFwGroupResponse() (response *ModifyVpcFwGroupResponse) {
    response = &ModifyVpcFwGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcFwGroup
// 编辑VPC间防火墙(防火墙组)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyVpcFwGroup(request *ModifyVpcFwGroupRequest) (response *ModifyVpcFwGroupResponse, err error) {
    return c.ModifyVpcFwGroupWithContext(context.Background(), request)
}

// ModifyVpcFwGroup
// 编辑VPC间防火墙(防火墙组)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyVpcFwGroupWithContext(ctx context.Context, request *ModifyVpcFwGroupRequest) (response *ModifyVpcFwGroupResponse, err error) {
    if request == nil {
        request = NewModifyVpcFwGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcFwGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcFwGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcFwSequenceRulesRequest() (request *ModifyVpcFwSequenceRulesRequest) {
    request = &ModifyVpcFwSequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyVpcFwSequenceRules")
    
    
    return
}

func NewModifyVpcFwSequenceRulesResponse() (response *ModifyVpcFwSequenceRulesResponse) {
    response = &ModifyVpcFwSequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcFwSequenceRules
// vpc间规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyVpcFwSequenceRules(request *ModifyVpcFwSequenceRulesRequest) (response *ModifyVpcFwSequenceRulesResponse, err error) {
    return c.ModifyVpcFwSequenceRulesWithContext(context.Background(), request)
}

// ModifyVpcFwSequenceRules
// vpc间规则快速排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) ModifyVpcFwSequenceRulesWithContext(ctx context.Context, request *ModifyVpcFwSequenceRulesRequest) (response *ModifyVpcFwSequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifyVpcFwSequenceRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcFwSequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpcFwSequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAcRuleRequest() (request *RemoveAcRuleRequest) {
    request = &RemoveAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveAcRule")
    
    
    return
}

func NewRemoveAcRuleResponse() (response *RemoveAcRuleResponse) {
    response = &RemoveAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveAcRule
// 删除互联网边界规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveAcRule(request *RemoveAcRuleRequest) (response *RemoveAcRuleResponse, err error) {
    return c.RemoveAcRuleWithContext(context.Background(), request)
}

// RemoveAcRule
// 删除互联网边界规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveAcRuleWithContext(ctx context.Context, request *RemoveAcRuleRequest) (response *RemoveAcRuleResponse, err error) {
    if request == nil {
        request = NewRemoveAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAclRuleRequest() (request *RemoveAclRuleRequest) {
    request = &RemoveAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveAclRule")
    
    
    return
}

func NewRemoveAclRuleResponse() (response *RemoveAclRuleResponse) {
    response = &RemoveAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveAclRule
// 删除互联网边界访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveAclRule(request *RemoveAclRuleRequest) (response *RemoveAclRuleResponse, err error) {
    return c.RemoveAclRuleWithContext(context.Background(), request)
}

// RemoveAclRule
// 删除互联网边界访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveAclRuleWithContext(ctx context.Context, request *RemoveAclRuleRequest) (response *RemoveAclRuleResponse, err error) {
    if request == nil {
        request = NewRemoveAclRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveEnterpriseSecurityGroupRuleRequest() (request *RemoveEnterpriseSecurityGroupRuleRequest) {
    request = &RemoveEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveEnterpriseSecurityGroupRule")
    
    
    return
}

func NewRemoveEnterpriseSecurityGroupRuleResponse() (response *RemoveEnterpriseSecurityGroupRuleResponse) {
    response = &RemoveEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveEnterpriseSecurityGroupRule
// 删除新企业安全组规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveEnterpriseSecurityGroupRule(request *RemoveEnterpriseSecurityGroupRuleRequest) (response *RemoveEnterpriseSecurityGroupRuleResponse, err error) {
    return c.RemoveEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// RemoveEnterpriseSecurityGroupRule
// 删除新企业安全组规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *RemoveEnterpriseSecurityGroupRuleRequest) (response *RemoveEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewRemoveEnterpriseSecurityGroupRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveNatAcRuleRequest() (request *RemoveNatAcRuleRequest) {
    request = &RemoveNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveNatAcRule")
    
    
    return
}

func NewRemoveNatAcRuleResponse() (response *RemoveNatAcRuleResponse) {
    response = &RemoveNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveNatAcRule
// 删除NAT访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveNatAcRule(request *RemoveNatAcRuleRequest) (response *RemoveNatAcRuleResponse, err error) {
    return c.RemoveNatAcRuleWithContext(context.Background(), request)
}

// RemoveNatAcRule
// 删除NAT访问控制规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveNatAcRuleWithContext(ctx context.Context, request *RemoveNatAcRuleRequest) (response *RemoveNatAcRuleResponse, err error) {
    if request == nil {
        request = NewRemoveNatAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveVpcAcRuleRequest() (request *RemoveVpcAcRuleRequest) {
    request = &RemoveVpcAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveVpcAcRule")
    
    
    return
}

func NewRemoveVpcAcRuleResponse() (response *RemoveVpcAcRuleResponse) {
    response = &RemoveVpcAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveVpcAcRule
// 删除VPC间规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveVpcAcRule(request *RemoveVpcAcRuleRequest) (response *RemoveVpcAcRuleResponse, err error) {
    return c.RemoveVpcAcRuleWithContext(context.Background(), request)
}

// RemoveVpcAcRule
// 删除VPC间规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) RemoveVpcAcRuleWithContext(ctx context.Context, request *RemoveVpcAcRuleRequest) (response *RemoveVpcAcRuleResponse, err error) {
    if request == nil {
        request = NewRemoveVpcAcRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveVpcAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveVpcAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSetNatFwDnatRuleRequest() (request *SetNatFwDnatRuleRequest) {
    request = &SetNatFwDnatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwDnatRule")
    
    
    return
}

func NewSetNatFwDnatRuleResponse() (response *SetNatFwDnatRuleResponse) {
    response = &SetNatFwDnatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNatFwDnatRule
// 配置防火墙Dnat规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SetNatFwDnatRule(request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
    return c.SetNatFwDnatRuleWithContext(context.Background(), request)
}

// SetNatFwDnatRule
// 配置防火墙Dnat规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SetNatFwDnatRuleWithContext(ctx context.Context, request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
    if request == nil {
        request = NewSetNatFwDnatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNatFwDnatRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNatFwDnatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSetNatFwEipRequest() (request *SetNatFwEipRequest) {
    request = &SetNatFwEipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwEip")
    
    
    return
}

func NewSetNatFwEipResponse() (response *SetNatFwEipResponse) {
    response = &SetNatFwEipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNatFwEip
// 设置防火墙实例弹性公网ip，目前仅支持新增模式的防火墙实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SetNatFwEip(request *SetNatFwEipRequest) (response *SetNatFwEipResponse, err error) {
    return c.SetNatFwEipWithContext(context.Background(), request)
}

// SetNatFwEip
// 设置防火墙实例弹性公网ip，目前仅支持新增模式的防火墙实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SetNatFwEipWithContext(ctx context.Context, request *SetNatFwEipRequest) (response *SetNatFwEipResponse, err error) {
    if request == nil {
        request = NewSetNatFwEipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNatFwEip require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNatFwEipResponse()
    err = c.Send(request, response)
    return
}

func NewStopSecurityGroupRuleDispatchRequest() (request *StopSecurityGroupRuleDispatchRequest) {
    request = &StopSecurityGroupRuleDispatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "StopSecurityGroupRuleDispatch")
    
    
    return
}

func NewStopSecurityGroupRuleDispatchResponse() (response *StopSecurityGroupRuleDispatchResponse) {
    response = &StopSecurityGroupRuleDispatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSecurityGroupRuleDispatch
// 中止安全组规则下发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) StopSecurityGroupRuleDispatch(request *StopSecurityGroupRuleDispatchRequest) (response *StopSecurityGroupRuleDispatchResponse, err error) {
    return c.StopSecurityGroupRuleDispatchWithContext(context.Background(), request)
}

// StopSecurityGroupRuleDispatch
// 中止安全组规则下发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) StopSecurityGroupRuleDispatchWithContext(ctx context.Context, request *StopSecurityGroupRuleDispatchRequest) (response *StopSecurityGroupRuleDispatchResponse, err error) {
    if request == nil {
        request = NewStopSecurityGroupRuleDispatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSecurityGroupRuleDispatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSecurityGroupRuleDispatchResponse()
    err = c.Send(request, response)
    return
}

func NewSyncFwOperateRequest() (request *SyncFwOperateRequest) {
    request = &SyncFwOperateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "SyncFwOperate")
    
    
    return
}

func NewSyncFwOperateResponse() (response *SyncFwOperateResponse) {
    response = &SyncFwOperateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncFwOperate
// 同步防火墙操作，包括同步防火墙路由（若vpc，专线网关等增加了Cidr，需要手动同步一下路由使之在防火墙上生效）等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SyncFwOperate(request *SyncFwOperateRequest) (response *SyncFwOperateResponse, err error) {
    return c.SyncFwOperateWithContext(context.Background(), request)
}

// SyncFwOperate
// 同步防火墙操作，包括同步防火墙路由（若vpc，专线网关等增加了Cidr，需要手动同步一下路由使之在防火墙上生效）等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
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
func (c *Client) SyncFwOperateWithContext(ctx context.Context, request *SyncFwOperateRequest) (response *SyncFwOperateResponse, err error) {
    if request == nil {
        request = NewSyncFwOperateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncFwOperate require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncFwOperateResponse()
    err = c.Send(request, response)
    return
}
