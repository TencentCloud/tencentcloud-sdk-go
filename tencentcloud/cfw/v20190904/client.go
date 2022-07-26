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
// 创建规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRules(request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    return c.CreateAcRulesWithContext(context.Background(), request)
}

// CreateAcRules
// 创建规则
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
// 创建防火墙实例和接入域名
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
// 创建防火墙实例和接入域名
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

func NewCreateSecurityGroupApiRulesRequest() (request *CreateSecurityGroupApiRulesRequest) {
    request = &CreateSecurityGroupApiRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "CreateSecurityGroupApiRules")
    
    
    return
}

func NewCreateSecurityGroupApiRulesResponse() (response *CreateSecurityGroupApiRulesResponse) {
    response = &CreateSecurityGroupApiRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroupApiRules
// 创建安全组API规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSecurityGroupApiRules(request *CreateSecurityGroupApiRulesRequest) (response *CreateSecurityGroupApiRulesResponse, err error) {
    return c.CreateSecurityGroupApiRulesWithContext(context.Background(), request)
}

// CreateSecurityGroupApiRules
// 创建安全组API规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSecurityGroupApiRulesWithContext(ctx context.Context, request *CreateSecurityGroupApiRulesRequest) (response *CreateSecurityGroupApiRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupApiRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupApiRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupApiRulesResponse()
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

func NewDeleteSecurityGroupAllRuleRequest() (request *DeleteSecurityGroupAllRuleRequest) {
    request = &DeleteSecurityGroupAllRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupAllRule")
    
    
    return
}

func NewDeleteSecurityGroupAllRuleResponse() (response *DeleteSecurityGroupAllRuleResponse) {
    response = &DeleteSecurityGroupAllRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroupAllRule
// 删除全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupAllRule(request *DeleteSecurityGroupAllRuleRequest) (response *DeleteSecurityGroupAllRuleResponse, err error) {
    return c.DeleteSecurityGroupAllRuleWithContext(context.Background(), request)
}

// DeleteSecurityGroupAllRule
// 删除全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupAllRuleWithContext(ctx context.Context, request *DeleteSecurityGroupAllRuleRequest) (response *DeleteSecurityGroupAllRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupAllRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupAllRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupAllRuleResponse()
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

func NewDeleteVpcInstanceRequest() (request *DeleteVpcInstanceRequest) {
    request = &DeleteVpcInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcInstance")
    
    
    return
}

func NewDeleteVpcInstanceResponse() (response *DeleteVpcInstanceResponse) {
    response = &DeleteVpcInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpcInstance
// 删除防火墙实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteVpcInstance(request *DeleteVpcInstanceRequest) (response *DeleteVpcInstanceResponse, err error) {
    return c.DeleteVpcInstanceWithContext(context.Background(), request)
}

// DeleteVpcInstance
// 删除防火墙实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteVpcInstanceWithContext(ctx context.Context, request *DeleteVpcInstanceRequest) (response *DeleteVpcInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteVpcInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcInstanceResponse()
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

func NewDescribeAddrTemplateListRequest() (request *DescribeAddrTemplateListRequest) {
    request = &DescribeAddrTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAddrTemplateList")
    
    
    return
}

func NewDescribeAddrTemplateListResponse() (response *DescribeAddrTemplateListResponse) {
    response = &DescribeAddrTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddrTemplateList
// 获取地址模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAddrTemplateList(request *DescribeAddrTemplateListRequest) (response *DescribeAddrTemplateListResponse, err error) {
    return c.DescribeAddrTemplateListWithContext(context.Background(), request)
}

// DescribeAddrTemplateList
// 获取地址模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAddrTemplateListWithContext(ctx context.Context, request *DescribeAddrTemplateListRequest) (response *DescribeAddrTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeAddrTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddrTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddrTemplateListResponse()
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
// ip防护状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIPStatusList(request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
    return c.DescribeIPStatusListWithContext(context.Background(), request)
}

// DescribeIPStatusList
// ip防护状态查询
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

func NewDescribeNatRuleOverviewRequest() (request *DescribeNatRuleOverviewRequest) {
    request = &DescribeNatRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatRuleOverview")
    
    
    return
}

func NewDescribeNatRuleOverviewResponse() (response *DescribeNatRuleOverviewResponse) {
    response = &DescribeNatRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatRuleOverview
// nat规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatRuleOverview(request *DescribeNatRuleOverviewRequest) (response *DescribeNatRuleOverviewResponse, err error) {
    return c.DescribeNatRuleOverviewWithContext(context.Background(), request)
}

// DescribeNatRuleOverview
// nat规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatRuleOverviewWithContext(ctx context.Context, request *DescribeNatRuleOverviewRequest) (response *DescribeNatRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeNatRuleOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatRuleOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatRuleOverviewResponse()
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
// DescribeSourceAsset-查询资产组全部资产信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSourceAsset(request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
    return c.DescribeSourceAssetWithContext(context.Background(), request)
}

// DescribeSourceAsset
// DescribeSourceAsset-查询资产组全部资产信息
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
// 防火墙开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchLists(request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    return c.DescribeSwitchListsWithContext(context.Background(), request)
}

// DescribeSwitchLists
// 防火墙开关列表
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

func NewDescribeSyncAssetStatusRequest() (request *DescribeSyncAssetStatusRequest) {
    request = &DescribeSyncAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSyncAssetStatus")
    
    
    return
}

func NewDescribeSyncAssetStatusResponse() (response *DescribeSyncAssetStatusResponse) {
    response = &DescribeSyncAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSyncAssetStatus
// 同步资产状态查询-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSyncAssetStatus(request *DescribeSyncAssetStatusRequest) (response *DescribeSyncAssetStatusResponse, err error) {
    return c.DescribeSyncAssetStatusWithContext(context.Background(), request)
}

// DescribeSyncAssetStatus
// 同步资产状态查询-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSyncAssetStatusWithContext(ctx context.Context, request *DescribeSyncAssetStatusRequest) (response *DescribeSyncAssetStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSyncAssetStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSyncAssetStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSyncAssetStatusResponse()
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

func NewDescribeVpcRuleOverviewRequest() (request *DescribeVpcRuleOverviewRequest) {
    request = &DescribeVpcRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcRuleOverview")
    
    
    return
}

func NewDescribeVpcRuleOverviewResponse() (response *DescribeVpcRuleOverviewResponse) {
    response = &DescribeVpcRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcRuleOverview
// vpc规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcRuleOverview(request *DescribeVpcRuleOverviewRequest) (response *DescribeVpcRuleOverviewResponse, err error) {
    return c.DescribeVpcRuleOverviewWithContext(context.Background(), request)
}

// DescribeVpcRuleOverview
// vpc规则列表概况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcRuleOverviewWithContext(ctx context.Context, request *DescribeVpcRuleOverviewRequest) (response *DescribeVpcRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeVpcRuleOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcRuleOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcRuleOverviewResponse()
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

func NewModifyAllSwitchStatusRequest() (request *ModifyAllSwitchStatusRequest) {
    request = &ModifyAllSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllSwitchStatus")
    
    
    return
}

func NewModifyAllSwitchStatusResponse() (response *ModifyAllSwitchStatusResponse) {
    response = &ModifyAllSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAllSwitchStatus
// 一键开启和关闭
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllSwitchStatus(request *ModifyAllSwitchStatusRequest) (response *ModifyAllSwitchStatusResponse, err error) {
    return c.ModifyAllSwitchStatusWithContext(context.Background(), request)
}

// ModifyAllSwitchStatus
// 一键开启和关闭
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllSwitchStatusWithContext(ctx context.Context, request *ModifyAllSwitchStatusRequest) (response *ModifyAllSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllSwitchStatusResponse()
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
// 支持对拦截列表、忽略列表如下操作：
//
// 批量增加拦截IP、忽略IP/域名
//
// 批量删除拦截IP、忽略IP/域名
//
// 批量修改拦截IP、忽略IP/域名生效事件
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
// 支持对拦截列表、忽略列表如下操作：
//
// 批量增加拦截IP、忽略IP/域名
//
// 批量删除拦截IP、忽略IP/域名
//
// 批量修改拦截IP、忽略IP/域名生效事件
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

func NewModifyItemSwitchStatusRequest() (request *ModifyItemSwitchStatusRequest) {
    request = &ModifyItemSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyItemSwitchStatus")
    
    
    return
}

func NewModifyItemSwitchStatusResponse() (response *ModifyItemSwitchStatusResponse) {
    response = &ModifyItemSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyItemSwitchStatus
// 修改单个防火墙开关
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyItemSwitchStatus(request *ModifyItemSwitchStatusRequest) (response *ModifyItemSwitchStatusResponse, err error) {
    return c.ModifyItemSwitchStatusWithContext(context.Background(), request)
}

// ModifyItemSwitchStatus
// 修改单个防火墙开关
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyItemSwitchStatusWithContext(ctx context.Context, request *ModifyItemSwitchStatusRequest) (response *ModifyItemSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyItemSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyItemSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyItemSwitchStatusResponse()
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

func NewModifyPublicIPSwitchStatusRequest() (request *ModifyPublicIPSwitchStatusRequest) {
    request = &ModifyPublicIPSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyPublicIPSwitchStatus")
    
    
    return
}

func NewModifyPublicIPSwitchStatusResponse() (response *ModifyPublicIPSwitchStatusResponse) {
    response = &ModifyPublicIPSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPublicIPSwitchStatus
// 单个修改互联网边界防火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPublicIPSwitchStatus(request *ModifyPublicIPSwitchStatusRequest) (response *ModifyPublicIPSwitchStatusResponse, err error) {
    return c.ModifyPublicIPSwitchStatusWithContext(context.Background(), request)
}

// ModifyPublicIPSwitchStatus
// 单个修改互联网边界防火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPublicIPSwitchStatusWithContext(ctx context.Context, request *ModifyPublicIPSwitchStatusRequest) (response *ModifyPublicIPSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyPublicIPSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublicIPSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublicIPSwitchStatusResponse()
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

func NewModifySecurityGroupAllRuleStatusRequest() (request *ModifySecurityGroupAllRuleStatusRequest) {
    request = &ModifySecurityGroupAllRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupAllRuleStatus")
    
    
    return
}

func NewModifySecurityGroupAllRuleStatusResponse() (response *ModifySecurityGroupAllRuleStatusResponse) {
    response = &ModifySecurityGroupAllRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySecurityGroupAllRuleStatus(request *ModifySecurityGroupAllRuleStatusRequest) (response *ModifySecurityGroupAllRuleStatusResponse, err error) {
    return c.ModifySecurityGroupAllRuleStatusWithContext(context.Background(), request)
}

// ModifySecurityGroupAllRuleStatus
// 启用停用全部规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySecurityGroupAllRuleStatusWithContext(ctx context.Context, request *ModifySecurityGroupAllRuleStatusRequest) (response *ModifySecurityGroupAllRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAllRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupAllRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupAllRuleStatusResponse()
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

func NewModifyVPCSwitchStatusRequest() (request *ModifyVPCSwitchStatusRequest) {
    request = &ModifyVPCSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyVPCSwitchStatus")
    
    
    return
}

func NewModifyVPCSwitchStatusResponse() (response *ModifyVPCSwitchStatusResponse) {
    response = &ModifyVPCSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVPCSwitchStatus
// 单个修改VPC火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVPCSwitchStatus(request *ModifyVPCSwitchStatusRequest) (response *ModifyVPCSwitchStatusResponse, err error) {
    return c.ModifyVPCSwitchStatusWithContext(context.Background(), request)
}

// ModifyVPCSwitchStatus
// 单个修改VPC火墙开关
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyVPCSwitchStatusWithContext(ctx context.Context, request *ModifyVPCSwitchStatusRequest) (response *ModifyVPCSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyVPCSwitchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVPCSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVPCSwitchStatusResponse()
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

func NewRunSyncAssetRequest() (request *RunSyncAssetRequest) {
    request = &RunSyncAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfw", APIVersion, "RunSyncAsset")
    
    
    return
}

func NewRunSyncAssetResponse() (response *RunSyncAssetResponse) {
    response = &RunSyncAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunSyncAsset
// 同步资产-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RunSyncAsset(request *RunSyncAssetRequest) (response *RunSyncAssetResponse, err error) {
    return c.RunSyncAssetWithContext(context.Background(), request)
}

// RunSyncAsset
// 同步资产-互联网&VPC
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RunSyncAssetWithContext(ctx context.Context, request *RunSyncAssetRequest) (response *RunSyncAssetResponse, err error) {
    if request == nil {
        request = NewRunSyncAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunSyncAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunSyncAssetResponse()
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
