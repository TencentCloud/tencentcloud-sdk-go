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

package v20201101

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-01"

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


func NewAddAssetImageRegistryRegistryDetailRequest() (request *AddAssetImageRegistryRegistryDetailRequest) {
    request = &AddAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewAddAssetImageRegistryRegistryDetailResponse() (response *AddAssetImageRegistryRegistryDetailResponse) {
    response = &AddAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddAssetImageRegistryRegistryDetail
// 新增单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddAssetImageRegistryRegistryDetail(request *AddAssetImageRegistryRegistryDetailRequest) (response *AddAssetImageRegistryRegistryDetailResponse, err error) {
    return c.AddAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// AddAssetImageRegistryRegistryDetail
// 新增单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *AddAssetImageRegistryRegistryDetailRequest) (response *AddAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewAddAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewAddCompliancePolicyItemToWhitelistRequest() (request *AddCompliancePolicyItemToWhitelistRequest) {
    request = &AddCompliancePolicyItemToWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddCompliancePolicyItemToWhitelist")
    
    
    return
}

func NewAddCompliancePolicyItemToWhitelistResponse() (response *AddCompliancePolicyItemToWhitelistResponse) {
    response = &AddCompliancePolicyItemToWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCompliancePolicyItemToWhitelist
// 将指定的检测项添加到白名单中，不显示未通过结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyItemToWhitelist(request *AddCompliancePolicyItemToWhitelistRequest) (response *AddCompliancePolicyItemToWhitelistResponse, err error) {
    return c.AddCompliancePolicyItemToWhitelistWithContext(context.Background(), request)
}

// AddCompliancePolicyItemToWhitelist
// 将指定的检测项添加到白名单中，不显示未通过结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) AddCompliancePolicyItemToWhitelistWithContext(ctx context.Context, request *AddCompliancePolicyItemToWhitelistRequest) (response *AddCompliancePolicyItemToWhitelistResponse, err error) {
    if request == nil {
        request = NewAddCompliancePolicyItemToWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCompliancePolicyItemToWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCompliancePolicyItemToWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditAbnormalProcessRuleRequest() (request *AddEditAbnormalProcessRuleRequest) {
    request = &AddEditAbnormalProcessRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditAbnormalProcessRule")
    
    
    return
}

func NewAddEditAbnormalProcessRuleResponse() (response *AddEditAbnormalProcessRuleResponse) {
    response = &AddEditAbnormalProcessRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEditAbnormalProcessRule
// 添加编辑运行时异常进程策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAbnormalProcessRule(request *AddEditAbnormalProcessRuleRequest) (response *AddEditAbnormalProcessRuleResponse, err error) {
    return c.AddEditAbnormalProcessRuleWithContext(context.Background(), request)
}

// AddEditAbnormalProcessRule
// 添加编辑运行时异常进程策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAbnormalProcessRuleWithContext(ctx context.Context, request *AddEditAbnormalProcessRuleRequest) (response *AddEditAbnormalProcessRuleResponse, err error) {
    if request == nil {
        request = NewAddEditAbnormalProcessRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditAbnormalProcessRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditAbnormalProcessRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditAccessControlRuleRequest() (request *AddEditAccessControlRuleRequest) {
    request = &AddEditAccessControlRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditAccessControlRule")
    
    
    return
}

func NewAddEditAccessControlRuleResponse() (response *AddEditAccessControlRuleResponse) {
    response = &AddEditAccessControlRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEditAccessControlRule
// 添加编辑运行时访问控制策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAccessControlRule(request *AddEditAccessControlRuleRequest) (response *AddEditAccessControlRuleResponse, err error) {
    return c.AddEditAccessControlRuleWithContext(context.Background(), request)
}

// AddEditAccessControlRule
// 添加编辑运行时访问控制策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddEditAccessControlRuleWithContext(ctx context.Context, request *AddEditAccessControlRuleRequest) (response *AddEditAccessControlRuleResponse, err error) {
    if request == nil {
        request = NewAddEditAccessControlRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditAccessControlRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditAccessControlRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditReverseShellWhiteListRequest() (request *AddEditReverseShellWhiteListRequest) {
    request = &AddEditReverseShellWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditReverseShellWhiteList")
    
    
    return
}

func NewAddEditReverseShellWhiteListResponse() (response *AddEditReverseShellWhiteListResponse) {
    response = &AddEditReverseShellWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEditReverseShellWhiteList
// 添加编辑运行时反弹shell白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRIPNOVALID = "InvalidParameter.ErrIpNoValid"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditReverseShellWhiteList(request *AddEditReverseShellWhiteListRequest) (response *AddEditReverseShellWhiteListResponse, err error) {
    return c.AddEditReverseShellWhiteListWithContext(context.Background(), request)
}

// AddEditReverseShellWhiteList
// 添加编辑运行时反弹shell白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRIPNOVALID = "InvalidParameter.ErrIpNoValid"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditReverseShellWhiteListWithContext(ctx context.Context, request *AddEditReverseShellWhiteListRequest) (response *AddEditReverseShellWhiteListResponse, err error) {
    if request == nil {
        request = NewAddEditReverseShellWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditReverseShellWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditReverseShellWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditRiskSyscallWhiteListRequest() (request *AddEditRiskSyscallWhiteListRequest) {
    request = &AddEditRiskSyscallWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditRiskSyscallWhiteList")
    
    
    return
}

func NewAddEditRiskSyscallWhiteListResponse() (response *AddEditRiskSyscallWhiteListResponse) {
    response = &AddEditRiskSyscallWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEditRiskSyscallWhiteList
// 添加编辑运行时高危系统调用白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditRiskSyscallWhiteList(request *AddEditRiskSyscallWhiteListRequest) (response *AddEditRiskSyscallWhiteListResponse, err error) {
    return c.AddEditRiskSyscallWhiteListWithContext(context.Background(), request)
}

// AddEditRiskSyscallWhiteList
// 添加编辑运行时高危系统调用白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"
//  FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"
//  FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"
//  FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"
func (c *Client) AddEditRiskSyscallWhiteListWithContext(ctx context.Context, request *AddEditRiskSyscallWhiteListRequest) (response *AddEditRiskSyscallWhiteListResponse, err error) {
    if request == nil {
        request = NewAddEditRiskSyscallWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditRiskSyscallWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditRiskSyscallWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddEditWarningRulesRequest() (request *AddEditWarningRulesRequest) {
    request = &AddEditWarningRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "AddEditWarningRules")
    
    
    return
}

func NewAddEditWarningRulesResponse() (response *AddEditWarningRulesResponse) {
    response = &AddEditWarningRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEditWarningRules
// 添加编辑告警策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddEditWarningRules(request *AddEditWarningRulesRequest) (response *AddEditWarningRulesResponse, err error) {
    return c.AddEditWarningRulesWithContext(context.Background(), request)
}

// AddEditWarningRules
// 添加编辑告警策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AddEditWarningRulesWithContext(ctx context.Context, request *AddEditWarningRulesRequest) (response *AddEditWarningRulesResponse, err error) {
    if request == nil {
        request = NewAddEditWarningRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEditWarningRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEditWarningRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCheckRepeatAssetImageRegistryRequest() (request *CheckRepeatAssetImageRegistryRequest) {
    request = &CheckRepeatAssetImageRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CheckRepeatAssetImageRegistry")
    
    
    return
}

func NewCheckRepeatAssetImageRegistryResponse() (response *CheckRepeatAssetImageRegistryResponse) {
    response = &CheckRepeatAssetImageRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckRepeatAssetImageRegistry
// 检查单个镜像仓库名是否重复
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckRepeatAssetImageRegistry(request *CheckRepeatAssetImageRegistryRequest) (response *CheckRepeatAssetImageRegistryResponse, err error) {
    return c.CheckRepeatAssetImageRegistryWithContext(context.Background(), request)
}

// CheckRepeatAssetImageRegistry
// 检查单个镜像仓库名是否重复
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckRepeatAssetImageRegistryWithContext(ctx context.Context, request *CheckRepeatAssetImageRegistryRequest) (response *CheckRepeatAssetImageRegistryResponse, err error) {
    if request == nil {
        request = NewCheckRepeatAssetImageRegistryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRepeatAssetImageRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRepeatAssetImageRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageRegistryScanTaskRequest() (request *CreateAssetImageRegistryScanTaskRequest) {
    request = &CreateAssetImageRegistryScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTask")
    
    
    return
}

func NewCreateAssetImageRegistryScanTaskResponse() (response *CreateAssetImageRegistryScanTaskResponse) {
    response = &CreateAssetImageRegistryScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssetImageRegistryScanTask
// 镜像仓库创建镜像扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTask(request *CreateAssetImageRegistryScanTaskRequest) (response *CreateAssetImageRegistryScanTaskResponse, err error) {
    return c.CreateAssetImageRegistryScanTaskWithContext(context.Background(), request)
}

// CreateAssetImageRegistryScanTask
// 镜像仓库创建镜像扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskWithContext(ctx context.Context, request *CreateAssetImageRegistryScanTaskRequest) (response *CreateAssetImageRegistryScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageRegistryScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageRegistryScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageRegistryScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageRegistryScanTaskOneKeyRequest() (request *CreateAssetImageRegistryScanTaskOneKeyRequest) {
    request = &CreateAssetImageRegistryScanTaskOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageRegistryScanTaskOneKey")
    
    
    return
}

func NewCreateAssetImageRegistryScanTaskOneKeyResponse() (response *CreateAssetImageRegistryScanTaskOneKeyResponse) {
    response = &CreateAssetImageRegistryScanTaskOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssetImageRegistryScanTaskOneKey
// 镜像仓库创建镜像一键扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskOneKey(request *CreateAssetImageRegistryScanTaskOneKeyRequest) (response *CreateAssetImageRegistryScanTaskOneKeyResponse, err error) {
    return c.CreateAssetImageRegistryScanTaskOneKeyWithContext(context.Background(), request)
}

// CreateAssetImageRegistryScanTaskOneKey
// 镜像仓库创建镜像一键扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageRegistryScanTaskOneKeyWithContext(ctx context.Context, request *CreateAssetImageRegistryScanTaskOneKeyRequest) (response *CreateAssetImageRegistryScanTaskOneKeyResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageRegistryScanTaskOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageRegistryScanTaskOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageRegistryScanTaskOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageScanSettingRequest() (request *CreateAssetImageScanSettingRequest) {
    request = &CreateAssetImageScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanSetting")
    
    
    return
}

func NewCreateAssetImageScanSettingResponse() (response *CreateAssetImageScanSettingResponse) {
    response = &CreateAssetImageScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssetImageScanSetting
// 添加容器安全镜像扫描设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAssetImageScanSetting(request *CreateAssetImageScanSettingRequest) (response *CreateAssetImageScanSettingResponse, err error) {
    return c.CreateAssetImageScanSettingWithContext(context.Background(), request)
}

// CreateAssetImageScanSetting
// 添加容器安全镜像扫描设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAssetImageScanSettingWithContext(ctx context.Context, request *CreateAssetImageScanSettingRequest) (response *CreateAssetImageScanSettingResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetImageScanTaskRequest() (request *CreateAssetImageScanTaskRequest) {
    request = &CreateAssetImageScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateAssetImageScanTask")
    
    
    return
}

func NewCreateAssetImageScanTaskResponse() (response *CreateAssetImageScanTaskResponse) {
    response = &CreateAssetImageScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssetImageScanTask
// 容器安全创建镜像扫描任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRALREADYSCANNING = "FailedOperation.ErrAlreadyScanning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageScanTask(request *CreateAssetImageScanTaskRequest) (response *CreateAssetImageScanTaskResponse, err error) {
    return c.CreateAssetImageScanTaskWithContext(context.Background(), request)
}

// CreateAssetImageScanTask
// 容器安全创建镜像扫描任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRALREADYSCANNING = "FailedOperation.ErrAlreadyScanning"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssetImageScanTaskWithContext(ctx context.Context, request *CreateAssetImageScanTaskRequest) (response *CreateAssetImageScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetImageScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetImageScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetImageScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCheckComponentRequest() (request *CreateCheckComponentRequest) {
    request = &CreateCheckComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateCheckComponent")
    
    
    return
}

func NewCreateCheckComponentResponse() (response *CreateCheckComponentResponse) {
    response = &CreateCheckComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCheckComponent
// 安装检查组件，创建防护容器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCheckComponent(request *CreateCheckComponentRequest) (response *CreateCheckComponentResponse, err error) {
    return c.CreateCheckComponentWithContext(context.Background(), request)
}

// CreateCheckComponent
// 安装检查组件，创建防护容器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCheckComponentWithContext(ctx context.Context, request *CreateCheckComponentRequest) (response *CreateCheckComponentResponse, err error) {
    if request == nil {
        request = NewCreateCheckComponentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCheckComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCheckComponentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterCheckTaskRequest() (request *CreateClusterCheckTaskRequest) {
    request = &CreateClusterCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateClusterCheckTask")
    
    
    return
}

func NewCreateClusterCheckTaskResponse() (response *CreateClusterCheckTaskResponse) {
    response = &CreateClusterCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterCheckTask
// 创建集群检查任务，用户检查用户的集群相关风险项
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClusterCheckTask(request *CreateClusterCheckTaskRequest) (response *CreateClusterCheckTaskResponse, err error) {
    return c.CreateClusterCheckTaskWithContext(context.Background(), request)
}

// CreateClusterCheckTask
// 创建集群检查任务，用户检查用户的集群相关风险项
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClusterCheckTaskWithContext(ctx context.Context, request *CreateClusterCheckTaskRequest) (response *CreateClusterCheckTaskResponse, err error) {
    if request == nil {
        request = NewCreateClusterCheckTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterCheckTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterCheckTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComplianceTaskRequest() (request *CreateComplianceTaskRequest) {
    request = &CreateComplianceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateComplianceTask")
    
    
    return
}

func NewCreateComplianceTaskResponse() (response *CreateComplianceTaskResponse) {
    response = &CreateComplianceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateComplianceTask
// 创建合规检查任务，在资产级别触发重新检测时使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplianceTask(request *CreateComplianceTaskRequest) (response *CreateComplianceTaskResponse, err error) {
    return c.CreateComplianceTaskWithContext(context.Background(), request)
}

// CreateComplianceTask
// 创建合规检查任务，在资产级别触发重新检测时使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateComplianceTaskWithContext(ctx context.Context, request *CreateComplianceTaskRequest) (response *CreateComplianceTaskResponse, err error) {
    if request == nil {
        request = NewCreateComplianceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComplianceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComplianceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExportComplianceStatusListJobRequest() (request *CreateExportComplianceStatusListJobRequest) {
    request = &CreateExportComplianceStatusListJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateExportComplianceStatusListJob")
    
    
    return
}

func NewCreateExportComplianceStatusListJobResponse() (response *CreateExportComplianceStatusListJobResponse) {
    response = &CreateExportComplianceStatusListJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateExportComplianceStatusListJob
// 创建一个导出安全合规信息的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateExportComplianceStatusListJob(request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
    return c.CreateExportComplianceStatusListJobWithContext(context.Background(), request)
}

// CreateExportComplianceStatusListJob
// 创建一个导出安全合规信息的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) CreateExportComplianceStatusListJobWithContext(ctx context.Context, request *CreateExportComplianceStatusListJobRequest) (response *CreateExportComplianceStatusListJobResponse, err error) {
    if request == nil {
        request = NewCreateExportComplianceStatusListJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExportComplianceStatusListJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportComplianceStatusListJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrModifyPostPayCoresRequest() (request *CreateOrModifyPostPayCoresRequest) {
    request = &CreateOrModifyPostPayCoresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateOrModifyPostPayCores")
    
    
    return
}

func NewCreateOrModifyPostPayCoresResponse() (response *CreateOrModifyPostPayCoresResponse) {
    response = &CreateOrModifyPostPayCoresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrModifyPostPayCores
// CreateOrModifyPostPayCores  创建或者编辑弹性计费上限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateOrModifyPostPayCores(request *CreateOrModifyPostPayCoresRequest) (response *CreateOrModifyPostPayCoresResponse, err error) {
    return c.CreateOrModifyPostPayCoresWithContext(context.Background(), request)
}

// CreateOrModifyPostPayCores
// CreateOrModifyPostPayCores  创建或者编辑弹性计费上限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateOrModifyPostPayCoresWithContext(ctx context.Context, request *CreateOrModifyPostPayCoresRequest) (response *CreateOrModifyPostPayCoresResponse, err error) {
    if request == nil {
        request = NewCreateOrModifyPostPayCoresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrModifyPostPayCores require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrModifyPostPayCoresResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRefreshTaskRequest() (request *CreateRefreshTaskRequest) {
    request = &CreateRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateRefreshTask")
    
    
    return
}

func NewCreateRefreshTaskResponse() (response *CreateRefreshTaskResponse) {
    response = &CreateRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRefreshTask
// 下发刷新任务，会刷新资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRefreshTask(request *CreateRefreshTaskRequest) (response *CreateRefreshTaskResponse, err error) {
    return c.CreateRefreshTaskWithContext(context.Background(), request)
}

// CreateRefreshTask
// 下发刷新任务，会刷新资产信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRefreshTaskWithContext(ctx context.Context, request *CreateRefreshTaskRequest) (response *CreateRefreshTaskResponse, err error) {
    if request == nil {
        request = NewCreateRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVirusScanAgainRequest() (request *CreateVirusScanAgainRequest) {
    request = &CreateVirusScanAgainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanAgain")
    
    
    return
}

func NewCreateVirusScanAgainResponse() (response *CreateVirusScanAgainResponse) {
    response = &CreateVirusScanAgainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVirusScanAgain
// 运行时文件查杀重新检测
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirusScanAgain(request *CreateVirusScanAgainRequest) (response *CreateVirusScanAgainResponse, err error) {
    return c.CreateVirusScanAgainWithContext(context.Background(), request)
}

// CreateVirusScanAgain
// 运行时文件查杀重新检测
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateVirusScanAgainWithContext(ctx context.Context, request *CreateVirusScanAgainRequest) (response *CreateVirusScanAgainResponse, err error) {
    if request == nil {
        request = NewCreateVirusScanAgainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVirusScanAgain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVirusScanAgainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVirusScanTaskRequest() (request *CreateVirusScanTaskRequest) {
    request = &CreateVirusScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "CreateVirusScanTask")
    
    
    return
}

func NewCreateVirusScanTaskResponse() (response *CreateVirusScanTaskResponse) {
    response = &CreateVirusScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVirusScanTask
// 运行时文件查杀一键扫描
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVirusScanTask(request *CreateVirusScanTaskRequest) (response *CreateVirusScanTaskResponse, err error) {
    return c.CreateVirusScanTaskWithContext(context.Background(), request)
}

// CreateVirusScanTask
// 运行时文件查杀一键扫描
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVirusScanTaskWithContext(ctx context.Context, request *CreateVirusScanTaskRequest) (response *CreateVirusScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateVirusScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVirusScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVirusScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAbnormalProcessRulesRequest() (request *DeleteAbnormalProcessRulesRequest) {
    request = &DeleteAbnormalProcessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteAbnormalProcessRules")
    
    
    return
}

func NewDeleteAbnormalProcessRulesResponse() (response *DeleteAbnormalProcessRulesResponse) {
    response = &DeleteAbnormalProcessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAbnormalProcessRules
// 删除运行异常进程策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAbnormalProcessRules(request *DeleteAbnormalProcessRulesRequest) (response *DeleteAbnormalProcessRulesResponse, err error) {
    return c.DeleteAbnormalProcessRulesWithContext(context.Background(), request)
}

// DeleteAbnormalProcessRules
// 删除运行异常进程策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAbnormalProcessRulesWithContext(ctx context.Context, request *DeleteAbnormalProcessRulesRequest) (response *DeleteAbnormalProcessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAbnormalProcessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAbnormalProcessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAbnormalProcessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessControlRulesRequest() (request *DeleteAccessControlRulesRequest) {
    request = &DeleteAccessControlRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteAccessControlRules")
    
    
    return
}

func NewDeleteAccessControlRulesResponse() (response *DeleteAccessControlRulesResponse) {
    response = &DeleteAccessControlRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessControlRules
// 删除运行访问控制策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAccessControlRules(request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
    return c.DeleteAccessControlRulesWithContext(context.Background(), request)
}

// DeleteAccessControlRules
// 删除运行访问控制策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAccessControlRulesWithContext(ctx context.Context, request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessControlRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessControlRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessControlRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompliancePolicyItemFromWhitelistRequest() (request *DeleteCompliancePolicyItemFromWhitelistRequest) {
    request = &DeleteCompliancePolicyItemFromWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteCompliancePolicyItemFromWhitelist")
    
    
    return
}

func NewDeleteCompliancePolicyItemFromWhitelistResponse() (response *DeleteCompliancePolicyItemFromWhitelistResponse) {
    response = &DeleteCompliancePolicyItemFromWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCompliancePolicyItemFromWhitelist
// 从白名单中删除将指定的检测项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyItemFromWhitelist(request *DeleteCompliancePolicyItemFromWhitelistRequest) (response *DeleteCompliancePolicyItemFromWhitelistResponse, err error) {
    return c.DeleteCompliancePolicyItemFromWhitelistWithContext(context.Background(), request)
}

// DeleteCompliancePolicyItemFromWhitelist
// 从白名单中删除将指定的检测项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DeleteCompliancePolicyItemFromWhitelistWithContext(ctx context.Context, request *DeleteCompliancePolicyItemFromWhitelistRequest) (response *DeleteCompliancePolicyItemFromWhitelistResponse, err error) {
    if request == nil {
        request = NewDeleteCompliancePolicyItemFromWhitelistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompliancePolicyItemFromWhitelist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompliancePolicyItemFromWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellWhiteListsRequest() (request *DeleteReverseShellWhiteListsRequest) {
    request = &DeleteReverseShellWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteReverseShellWhiteLists")
    
    
    return
}

func NewDeleteReverseShellWhiteListsResponse() (response *DeleteReverseShellWhiteListsResponse) {
    response = &DeleteReverseShellWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReverseShellWhiteLists
// 删除运行时反弹shell白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellWhiteLists(request *DeleteReverseShellWhiteListsRequest) (response *DeleteReverseShellWhiteListsResponse, err error) {
    return c.DeleteReverseShellWhiteListsWithContext(context.Background(), request)
}

// DeleteReverseShellWhiteLists
// 删除运行时反弹shell白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteReverseShellWhiteListsWithContext(ctx context.Context, request *DeleteReverseShellWhiteListsRequest) (response *DeleteReverseShellWhiteListsResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReverseShellWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReverseShellWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskSyscallWhiteListsRequest() (request *DeleteRiskSyscallWhiteListsRequest) {
    request = &DeleteRiskSyscallWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DeleteRiskSyscallWhiteLists")
    
    
    return
}

func NewDeleteRiskSyscallWhiteListsResponse() (response *DeleteRiskSyscallWhiteListsResponse) {
    response = &DeleteRiskSyscallWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRiskSyscallWhiteLists
// 删除运行时高危系统调用白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallWhiteLists(request *DeleteRiskSyscallWhiteListsRequest) (response *DeleteRiskSyscallWhiteListsResponse, err error) {
    return c.DeleteRiskSyscallWhiteListsWithContext(context.Background(), request)
}

// DeleteRiskSyscallWhiteLists
// 删除运行时高危系统调用白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteRiskSyscallWhiteListsWithContext(ctx context.Context, request *DeleteRiskSyscallWhiteListsRequest) (response *DeleteRiskSyscallWhiteListsResponse, err error) {
    if request == nil {
        request = NewDeleteRiskSyscallWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskSyscallWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskSyscallWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessDetailRequest() (request *DescribeAbnormalProcessDetailRequest) {
    request = &DescribeAbnormalProcessDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessDetail")
    
    
    return
}

func NewDescribeAbnormalProcessDetailResponse() (response *DescribeAbnormalProcessDetailResponse) {
    response = &DescribeAbnormalProcessDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessDetail
// 查询运行时异常进程事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessDetail(request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
    return c.DescribeAbnormalProcessDetailWithContext(context.Background(), request)
}

// DescribeAbnormalProcessDetail
// 查询运行时异常进程事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessDetailWithContext(ctx context.Context, request *DescribeAbnormalProcessDetailRequest) (response *DescribeAbnormalProcessDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessEventsRequest() (request *DescribeAbnormalProcessEventsRequest) {
    request = &DescribeAbnormalProcessEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEvents")
    
    
    return
}

func NewDescribeAbnormalProcessEventsResponse() (response *DescribeAbnormalProcessEventsResponse) {
    response = &DescribeAbnormalProcessEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessEvents
// 查询运行时异常进程事件列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEvents(request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
    return c.DescribeAbnormalProcessEventsWithContext(context.Background(), request)
}

// DescribeAbnormalProcessEvents
// 查询运行时异常进程事件列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsWithContext(ctx context.Context, request *DescribeAbnormalProcessEventsRequest) (response *DescribeAbnormalProcessEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessEventsExportRequest() (request *DescribeAbnormalProcessEventsExportRequest) {
    request = &DescribeAbnormalProcessEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessEventsExport")
    
    
    return
}

func NewDescribeAbnormalProcessEventsExportResponse() (response *DescribeAbnormalProcessEventsExportResponse) {
    response = &DescribeAbnormalProcessEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessEventsExport
// 查询运行时异常进程事件列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsExport(request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
    return c.DescribeAbnormalProcessEventsExportWithContext(context.Background(), request)
}

// DescribeAbnormalProcessEventsExport
// 查询运行时异常进程事件列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessEventsExportWithContext(ctx context.Context, request *DescribeAbnormalProcessEventsExportRequest) (response *DescribeAbnormalProcessEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRuleDetailRequest() (request *DescribeAbnormalProcessRuleDetailRequest) {
    request = &DescribeAbnormalProcessRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRuleDetail")
    
    
    return
}

func NewDescribeAbnormalProcessRuleDetailResponse() (response *DescribeAbnormalProcessRuleDetailResponse) {
    response = &DescribeAbnormalProcessRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessRuleDetail
// 查询运行时异常策略详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRuleDetail(request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
    return c.DescribeAbnormalProcessRuleDetailWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRuleDetail
// 查询运行时异常策略详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRuleDetailWithContext(ctx context.Context, request *DescribeAbnormalProcessRuleDetailRequest) (response *DescribeAbnormalProcessRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRulesRequest() (request *DescribeAbnormalProcessRulesRequest) {
    request = &DescribeAbnormalProcessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRules")
    
    
    return
}

func NewDescribeAbnormalProcessRulesResponse() (response *DescribeAbnormalProcessRulesResponse) {
    response = &DescribeAbnormalProcessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessRules
// 查询运行时异常进程策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRules(request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
    return c.DescribeAbnormalProcessRulesWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRules
// 查询运行时异常进程策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesWithContext(ctx context.Context, request *DescribeAbnormalProcessRulesRequest) (response *DescribeAbnormalProcessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalProcessRulesExportRequest() (request *DescribeAbnormalProcessRulesExportRequest) {
    request = &DescribeAbnormalProcessRulesExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAbnormalProcessRulesExport")
    
    
    return
}

func NewDescribeAbnormalProcessRulesExportResponse() (response *DescribeAbnormalProcessRulesExportResponse) {
    response = &DescribeAbnormalProcessRulesExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalProcessRulesExport
// 查询运行时异常进程策略列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesExport(request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
    return c.DescribeAbnormalProcessRulesExportWithContext(context.Background(), request)
}

// DescribeAbnormalProcessRulesExport
// 查询运行时异常进程策略列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAbnormalProcessRulesExportWithContext(ctx context.Context, request *DescribeAbnormalProcessRulesExportRequest) (response *DescribeAbnormalProcessRulesExportResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalProcessRulesExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalProcessRulesExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalProcessRulesExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlDetailRequest() (request *DescribeAccessControlDetailRequest) {
    request = &DescribeAccessControlDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlDetail")
    
    
    return
}

func NewDescribeAccessControlDetailResponse() (response *DescribeAccessControlDetailResponse) {
    response = &DescribeAccessControlDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlDetail
// 查询运行时访问控制事件的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlDetail(request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
    return c.DescribeAccessControlDetailWithContext(context.Background(), request)
}

// DescribeAccessControlDetail
// 查询运行时访问控制事件的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlDetailWithContext(ctx context.Context, request *DescribeAccessControlDetailRequest) (response *DescribeAccessControlDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlEventsRequest() (request *DescribeAccessControlEventsRequest) {
    request = &DescribeAccessControlEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEvents")
    
    
    return
}

func NewDescribeAccessControlEventsResponse() (response *DescribeAccessControlEventsResponse) {
    response = &DescribeAccessControlEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlEvents
// 查询运行时访问控制事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEvents(request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
    return c.DescribeAccessControlEventsWithContext(context.Background(), request)
}

// DescribeAccessControlEvents
// 查询运行时访问控制事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsWithContext(ctx context.Context, request *DescribeAccessControlEventsRequest) (response *DescribeAccessControlEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlEventsExportRequest() (request *DescribeAccessControlEventsExportRequest) {
    request = &DescribeAccessControlEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlEventsExport")
    
    
    return
}

func NewDescribeAccessControlEventsExportResponse() (response *DescribeAccessControlEventsExportResponse) {
    response = &DescribeAccessControlEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlEventsExport
// 查询运行时访问控制事件列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsExport(request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
    return c.DescribeAccessControlEventsExportWithContext(context.Background(), request)
}

// DescribeAccessControlEventsExport
// 查询运行时访问控制事件列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlEventsExportWithContext(ctx context.Context, request *DescribeAccessControlEventsExportRequest) (response *DescribeAccessControlEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRuleDetailRequest() (request *DescribeAccessControlRuleDetailRequest) {
    request = &DescribeAccessControlRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRuleDetail")
    
    
    return
}

func NewDescribeAccessControlRuleDetailResponse() (response *DescribeAccessControlRuleDetailResponse) {
    response = &DescribeAccessControlRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlRuleDetail
// 查询运行时访问控制策略详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRuleDetail(request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
    return c.DescribeAccessControlRuleDetailWithContext(context.Background(), request)
}

// DescribeAccessControlRuleDetail
// 查询运行时访问控制策略详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRuleDetailWithContext(ctx context.Context, request *DescribeAccessControlRuleDetailRequest) (response *DescribeAccessControlRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
    request = &DescribeAccessControlRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRules")
    
    
    return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
    response = &DescribeAccessControlRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlRules
// 查询运行访问控制策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
    return c.DescribeAccessControlRulesWithContext(context.Background(), request)
}

// DescribeAccessControlRules
// 查询运行访问控制策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesWithContext(ctx context.Context, request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessControlRulesExportRequest() (request *DescribeAccessControlRulesExportRequest) {
    request = &DescribeAccessControlRulesExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAccessControlRulesExport")
    
    
    return
}

func NewDescribeAccessControlRulesExportResponse() (response *DescribeAccessControlRulesExportResponse) {
    response = &DescribeAccessControlRulesExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessControlRulesExport
// 查询运行时访问控制策略列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesExport(request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
    return c.DescribeAccessControlRulesExportWithContext(context.Background(), request)
}

// DescribeAccessControlRulesExport
// 查询运行时访问控制策略列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessControlRulesExportWithContext(ctx context.Context, request *DescribeAccessControlRulesExportRequest) (response *DescribeAccessControlRulesExportResponse, err error) {
    if request == nil {
        request = NewDescribeAccessControlRulesExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessControlRulesExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessControlRulesExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedClusterCountRequest() (request *DescribeAffectedClusterCountRequest) {
    request = &DescribeAffectedClusterCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedClusterCount")
    
    
    return
}

func NewDescribeAffectedClusterCountResponse() (response *DescribeAffectedClusterCountResponse) {
    response = &DescribeAffectedClusterCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAffectedClusterCount
// 获取受影响的集群数量，返回数量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedClusterCount(request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
    return c.DescribeAffectedClusterCountWithContext(context.Background(), request)
}

// DescribeAffectedClusterCount
// 获取受影响的集群数量，返回数量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedClusterCountWithContext(ctx context.Context, request *DescribeAffectedClusterCountRequest) (response *DescribeAffectedClusterCountResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedClusterCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedClusterCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedClusterCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedNodeListRequest() (request *DescribeAffectedNodeListRequest) {
    request = &DescribeAffectedNodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedNodeList")
    
    
    return
}

func NewDescribeAffectedNodeListResponse() (response *DescribeAffectedNodeListResponse) {
    response = &DescribeAffectedNodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAffectedNodeList
// 查询节点类型的影响范围，返回节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedNodeList(request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
    return c.DescribeAffectedNodeListWithContext(context.Background(), request)
}

// DescribeAffectedNodeList
// 查询节点类型的影响范围，返回节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedNodeListWithContext(ctx context.Context, request *DescribeAffectedNodeListRequest) (response *DescribeAffectedNodeListResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedNodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedNodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedNodeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAffectedWorkloadListRequest() (request *DescribeAffectedWorkloadListRequest) {
    request = &DescribeAffectedWorkloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAffectedWorkloadList")
    
    
    return
}

func NewDescribeAffectedWorkloadListResponse() (response *DescribeAffectedWorkloadListResponse) {
    response = &DescribeAffectedWorkloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAffectedWorkloadList
// 查询workload类型的影响范围，返回workload列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedWorkloadList(request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
    return c.DescribeAffectedWorkloadListWithContext(context.Background(), request)
}

// DescribeAffectedWorkloadList
// 查询workload类型的影响范围，返回workload列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAffectedWorkloadListWithContext(ctx context.Context, request *DescribeAffectedWorkloadListRequest) (response *DescribeAffectedWorkloadListResponse, err error) {
    if request == nil {
        request = NewDescribeAffectedWorkloadListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAffectedWorkloadList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAffectedWorkloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetAppServiceListRequest() (request *DescribeAssetAppServiceListRequest) {
    request = &DescribeAssetAppServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetAppServiceList")
    
    
    return
}

func NewDescribeAssetAppServiceListResponse() (response *DescribeAssetAppServiceListResponse) {
    response = &DescribeAssetAppServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetAppServiceList
// 容器安全查询app服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppServiceList(request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
    return c.DescribeAssetAppServiceListWithContext(context.Background(), request)
}

// DescribeAssetAppServiceList
// 容器安全查询app服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppServiceListWithContext(ctx context.Context, request *DescribeAssetAppServiceListRequest) (response *DescribeAssetAppServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetAppServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetAppServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetAppServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetComponentListRequest() (request *DescribeAssetComponentListRequest) {
    request = &DescribeAssetComponentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetComponentList")
    
    
    return
}

func NewDescribeAssetComponentListResponse() (response *DescribeAssetComponentListResponse) {
    response = &DescribeAssetComponentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetComponentList
// 容器安全搜索查询容器组件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetComponentList(request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
    return c.DescribeAssetComponentListWithContext(context.Background(), request)
}

// DescribeAssetComponentList
// 容器安全搜索查询容器组件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetComponentListWithContext(ctx context.Context, request *DescribeAssetComponentListRequest) (response *DescribeAssetComponentListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetComponentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetComponentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetComponentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetContainerDetailRequest() (request *DescribeAssetContainerDetailRequest) {
    request = &DescribeAssetContainerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerDetail")
    
    
    return
}

func NewDescribeAssetContainerDetailResponse() (response *DescribeAssetContainerDetailResponse) {
    response = &DescribeAssetContainerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetContainerDetail
// 查询容器详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerDetail(request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
    return c.DescribeAssetContainerDetailWithContext(context.Background(), request)
}

// DescribeAssetContainerDetail
// 查询容器详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerDetailWithContext(ctx context.Context, request *DescribeAssetContainerDetailRequest) (response *DescribeAssetContainerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetContainerDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetContainerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetContainerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetContainerListRequest() (request *DescribeAssetContainerListRequest) {
    request = &DescribeAssetContainerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetContainerList")
    
    
    return
}

func NewDescribeAssetContainerListResponse() (response *DescribeAssetContainerListResponse) {
    response = &DescribeAssetContainerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetContainerList
// 搜索查询容器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerList(request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
    return c.DescribeAssetContainerListWithContext(context.Background(), request)
}

// DescribeAssetContainerList
// 搜索查询容器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetContainerListWithContext(ctx context.Context, request *DescribeAssetContainerListRequest) (response *DescribeAssetContainerListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetContainerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetContainerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetContainerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDBServiceListRequest() (request *DescribeAssetDBServiceListRequest) {
    request = &DescribeAssetDBServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetDBServiceList")
    
    
    return
}

func NewDescribeAssetDBServiceListResponse() (response *DescribeAssetDBServiceListResponse) {
    response = &DescribeAssetDBServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDBServiceList
// 容器安全查询db服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDBServiceList(request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
    return c.DescribeAssetDBServiceListWithContext(context.Background(), request)
}

// DescribeAssetDBServiceList
// 容器安全查询db服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDBServiceListWithContext(ctx context.Context, request *DescribeAssetDBServiceListRequest) (response *DescribeAssetDBServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDBServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDBServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDBServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHostDetailRequest() (request *DescribeAssetHostDetailRequest) {
    request = &DescribeAssetHostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostDetail")
    
    
    return
}

func NewDescribeAssetHostDetailResponse() (response *DescribeAssetHostDetailResponse) {
    response = &DescribeAssetHostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetHostDetail
// 查询主机详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostDetail(request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
    return c.DescribeAssetHostDetailWithContext(context.Background(), request)
}

// DescribeAssetHostDetail
// 查询主机详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostDetailWithContext(ctx context.Context, request *DescribeAssetHostDetailRequest) (response *DescribeAssetHostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHostDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetHostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetHostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHostListRequest() (request *DescribeAssetHostListRequest) {
    request = &DescribeAssetHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetHostList")
    
    
    return
}

func NewDescribeAssetHostListResponse() (response *DescribeAssetHostListResponse) {
    response = &DescribeAssetHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetHostList
// 容器安全搜索查询主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostList(request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
    return c.DescribeAssetHostListWithContext(context.Background(), request)
}

// DescribeAssetHostList
// 容器安全搜索查询主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetHostListWithContext(ctx context.Context, request *DescribeAssetHostListRequest) (response *DescribeAssetHostListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageBindRuleInfoRequest() (request *DescribeAssetImageBindRuleInfoRequest) {
    request = &DescribeAssetImageBindRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageBindRuleInfo")
    
    
    return
}

func NewDescribeAssetImageBindRuleInfoResponse() (response *DescribeAssetImageBindRuleInfoResponse) {
    response = &DescribeAssetImageBindRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageBindRuleInfo
// 镜像绑定规则列表信息，包含运行时访问控制和异常进程公用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageBindRuleInfo(request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
    return c.DescribeAssetImageBindRuleInfoWithContext(context.Background(), request)
}

// DescribeAssetImageBindRuleInfo
// 镜像绑定规则列表信息，包含运行时访问控制和异常进程公用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageBindRuleInfoWithContext(ctx context.Context, request *DescribeAssetImageBindRuleInfoRequest) (response *DescribeAssetImageBindRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageBindRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageBindRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageBindRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageDetailRequest() (request *DescribeAssetImageDetailRequest) {
    request = &DescribeAssetImageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageDetail")
    
    
    return
}

func NewDescribeAssetImageDetailResponse() (response *DescribeAssetImageDetailResponse) {
    response = &DescribeAssetImageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageDetail
// 查询镜像详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageDetail(request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
    return c.DescribeAssetImageDetailWithContext(context.Background(), request)
}

// DescribeAssetImageDetail
// 查询镜像详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageDetailWithContext(ctx context.Context, request *DescribeAssetImageDetailRequest) (response *DescribeAssetImageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageHostListRequest() (request *DescribeAssetImageHostListRequest) {
    request = &DescribeAssetImageHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageHostList")
    
    
    return
}

func NewDescribeAssetImageHostListResponse() (response *DescribeAssetImageHostListResponse) {
    response = &DescribeAssetImageHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageHostList
// 容器安全查询镜像关联主机
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageHostList(request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
    return c.DescribeAssetImageHostListWithContext(context.Background(), request)
}

// DescribeAssetImageHostList
// 容器安全查询镜像关联主机
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageHostListWithContext(ctx context.Context, request *DescribeAssetImageHostListRequest) (response *DescribeAssetImageHostListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageListRequest() (request *DescribeAssetImageListRequest) {
    request = &DescribeAssetImageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageList")
    
    
    return
}

func NewDescribeAssetImageListResponse() (response *DescribeAssetImageListResponse) {
    response = &DescribeAssetImageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageList
// 容器安全搜索查询镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageList(request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
    return c.DescribeAssetImageListWithContext(context.Background(), request)
}

// DescribeAssetImageList
// 容器安全搜索查询镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListWithContext(ctx context.Context, request *DescribeAssetImageListRequest) (response *DescribeAssetImageListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageListExportRequest() (request *DescribeAssetImageListExportRequest) {
    request = &DescribeAssetImageListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageListExport")
    
    
    return
}

func NewDescribeAssetImageListExportResponse() (response *DescribeAssetImageListExportResponse) {
    response = &DescribeAssetImageListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageListExport
// 容器安全搜索查询镜像列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListExport(request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
    return c.DescribeAssetImageListExportWithContext(context.Background(), request)
}

// DescribeAssetImageListExport
// 容器安全搜索查询镜像列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageListExportWithContext(ctx context.Context, request *DescribeAssetImageListExportRequest) (response *DescribeAssetImageListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryAssetStatusRequest() (request *DescribeAssetImageRegistryAssetStatusRequest) {
    request = &DescribeAssetImageRegistryAssetStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryAssetStatus")
    
    
    return
}

func NewDescribeAssetImageRegistryAssetStatusResponse() (response *DescribeAssetImageRegistryAssetStatusResponse) {
    response = &DescribeAssetImageRegistryAssetStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryAssetStatus
// 查看镜像仓库资产更新进度状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryAssetStatus(request *DescribeAssetImageRegistryAssetStatusRequest) (response *DescribeAssetImageRegistryAssetStatusResponse, err error) {
    return c.DescribeAssetImageRegistryAssetStatusWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryAssetStatus
// 查看镜像仓库资产更新进度状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryAssetStatusWithContext(ctx context.Context, request *DescribeAssetImageRegistryAssetStatusRequest) (response *DescribeAssetImageRegistryAssetStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryAssetStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryAssetStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryAssetStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryDetailRequest() (request *DescribeAssetImageRegistryDetailRequest) {
    request = &DescribeAssetImageRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryDetail")
    
    
    return
}

func NewDescribeAssetImageRegistryDetailResponse() (response *DescribeAssetImageRegistryDetailResponse) {
    response = &DescribeAssetImageRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryDetail
// 镜像仓库镜像仓库列表详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryDetail(request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
    return c.DescribeAssetImageRegistryDetailWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryDetail
// 镜像仓库镜像仓库列表详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryDetailWithContext(ctx context.Context, request *DescribeAssetImageRegistryDetailRequest) (response *DescribeAssetImageRegistryDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryListRequest() (request *DescribeAssetImageRegistryListRequest) {
    request = &DescribeAssetImageRegistryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryList")
    
    
    return
}

func NewDescribeAssetImageRegistryListResponse() (response *DescribeAssetImageRegistryListResponse) {
    response = &DescribeAssetImageRegistryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryList
// 镜像仓库镜像仓库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryList(request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
    return c.DescribeAssetImageRegistryListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryList
// 镜像仓库镜像仓库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListWithContext(ctx context.Context, request *DescribeAssetImageRegistryListRequest) (response *DescribeAssetImageRegistryListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryListExportRequest() (request *DescribeAssetImageRegistryListExportRequest) {
    request = &DescribeAssetImageRegistryListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryListExportResponse() (response *DescribeAssetImageRegistryListExportResponse) {
    response = &DescribeAssetImageRegistryListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryListExport
// 镜像仓库镜像列表导出
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListExport(request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
    return c.DescribeAssetImageRegistryListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryListExport
// 镜像仓库镜像列表导出
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryListExportRequest) (response *DescribeAssetImageRegistryListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRegistryDetailRequest() (request *DescribeAssetImageRegistryRegistryDetailRequest) {
    request = &DescribeAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewDescribeAssetImageRegistryRegistryDetailResponse() (response *DescribeAssetImageRegistryRegistryDetailResponse) {
    response = &DescribeAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryRegistryDetail
// 查看单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryDetail(request *DescribeAssetImageRegistryRegistryDetailRequest) (response *DescribeAssetImageRegistryRegistryDetailResponse, err error) {
    return c.DescribeAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRegistryDetail
// 查看单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *DescribeAssetImageRegistryRegistryDetailRequest) (response *DescribeAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRegistryListRequest() (request *DescribeAssetImageRegistryRegistryListRequest) {
    request = &DescribeAssetImageRegistryRegistryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRegistryList")
    
    
    return
}

func NewDescribeAssetImageRegistryRegistryListResponse() (response *DescribeAssetImageRegistryRegistryListResponse) {
    response = &DescribeAssetImageRegistryRegistryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryRegistryList
// 镜像仓库仓库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryList(request *DescribeAssetImageRegistryRegistryListRequest) (response *DescribeAssetImageRegistryRegistryListResponse, err error) {
    return c.DescribeAssetImageRegistryRegistryListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRegistryList
// 镜像仓库仓库列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRegistryListWithContext(ctx context.Context, request *DescribeAssetImageRegistryRegistryListRequest) (response *DescribeAssetImageRegistryRegistryListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRegistryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRegistryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRegistryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRiskInfoListRequest() (request *DescribeAssetImageRegistryRiskInfoListRequest) {
    request = &DescribeAssetImageRegistryRiskInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskInfoList")
    
    
    return
}

func NewDescribeAssetImageRegistryRiskInfoListResponse() (response *DescribeAssetImageRegistryRiskInfoListResponse) {
    response = &DescribeAssetImageRegistryRiskInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryRiskInfoList
// 镜像仓库查询镜像高危行为列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskInfoList(request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
    return c.DescribeAssetImageRegistryRiskInfoListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRiskInfoList
// 镜像仓库查询镜像高危行为列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskInfoListWithContext(ctx context.Context, request *DescribeAssetImageRegistryRiskInfoListRequest) (response *DescribeAssetImageRegistryRiskInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRiskInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRiskInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRiskInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryRiskListExportRequest() (request *DescribeAssetImageRegistryRiskListExportRequest) {
    request = &DescribeAssetImageRegistryRiskListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryRiskListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryRiskListExportResponse() (response *DescribeAssetImageRegistryRiskListExportResponse) {
    response = &DescribeAssetImageRegistryRiskListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryRiskListExport
// 镜像仓库敏感信息列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskListExport(request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
    return c.DescribeAssetImageRegistryRiskListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryRiskListExport
// 镜像仓库敏感信息列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryRiskListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryRiskListExportRequest) (response *DescribeAssetImageRegistryRiskListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryRiskListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryRiskListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryRiskListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyRequest() (request *DescribeAssetImageRegistryScanStatusOneKeyRequest) {
    request = &DescribeAssetImageRegistryScanStatusOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryScanStatusOneKey")
    
    
    return
}

func NewDescribeAssetImageRegistryScanStatusOneKeyResponse() (response *DescribeAssetImageRegistryScanStatusOneKeyResponse) {
    response = &DescribeAssetImageRegistryScanStatusOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryScanStatusOneKey
// 镜像仓库查询一键镜像扫描状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryScanStatusOneKey(request *DescribeAssetImageRegistryScanStatusOneKeyRequest) (response *DescribeAssetImageRegistryScanStatusOneKeyResponse, err error) {
    return c.DescribeAssetImageRegistryScanStatusOneKeyWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryScanStatusOneKey
// 镜像仓库查询一键镜像扫描状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryScanStatusOneKeyWithContext(ctx context.Context, request *DescribeAssetImageRegistryScanStatusOneKeyRequest) (response *DescribeAssetImageRegistryScanStatusOneKeyResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryScanStatusOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryScanStatusOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryScanStatusOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistrySummaryRequest() (request *DescribeAssetImageRegistrySummaryRequest) {
    request = &DescribeAssetImageRegistrySummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistrySummary")
    
    
    return
}

func NewDescribeAssetImageRegistrySummaryResponse() (response *DescribeAssetImageRegistrySummaryResponse) {
    response = &DescribeAssetImageRegistrySummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistrySummary
// 镜像仓库查询镜像统计信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistrySummary(request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
    return c.DescribeAssetImageRegistrySummaryWithContext(context.Background(), request)
}

// DescribeAssetImageRegistrySummary
// 镜像仓库查询镜像统计信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistrySummaryWithContext(ctx context.Context, request *DescribeAssetImageRegistrySummaryRequest) (response *DescribeAssetImageRegistrySummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistrySummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistrySummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistrySummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVirusListRequest() (request *DescribeAssetImageRegistryVirusListRequest) {
    request = &DescribeAssetImageRegistryVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusList")
    
    
    return
}

func NewDescribeAssetImageRegistryVirusListResponse() (response *DescribeAssetImageRegistryVirusListResponse) {
    response = &DescribeAssetImageRegistryVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryVirusList
// 镜像仓库查询木马病毒列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusList(request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
    return c.DescribeAssetImageRegistryVirusListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVirusList
// 镜像仓库查询木马病毒列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListWithContext(ctx context.Context, request *DescribeAssetImageRegistryVirusListRequest) (response *DescribeAssetImageRegistryVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVirusListExportRequest() (request *DescribeAssetImageRegistryVirusListExportRequest) {
    request = &DescribeAssetImageRegistryVirusListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVirusListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryVirusListExportResponse() (response *DescribeAssetImageRegistryVirusListExportResponse) {
    response = &DescribeAssetImageRegistryVirusListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryVirusListExport
// 镜像仓库木马信息列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListExport(request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
    return c.DescribeAssetImageRegistryVirusListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVirusListExport
// 镜像仓库木马信息列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVirusListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryVirusListExportRequest) (response *DescribeAssetImageRegistryVirusListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVirusListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVirusListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVirusListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVulListRequest() (request *DescribeAssetImageRegistryVulListRequest) {
    request = &DescribeAssetImageRegistryVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulList")
    
    
    return
}

func NewDescribeAssetImageRegistryVulListResponse() (response *DescribeAssetImageRegistryVulListResponse) {
    response = &DescribeAssetImageRegistryVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryVulList
// 镜像仓库查询镜像漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulList(request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
    return c.DescribeAssetImageRegistryVulListWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVulList
// 镜像仓库查询镜像漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListWithContext(ctx context.Context, request *DescribeAssetImageRegistryVulListRequest) (response *DescribeAssetImageRegistryVulListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRegistryVulListExportRequest() (request *DescribeAssetImageRegistryVulListExportRequest) {
    request = &DescribeAssetImageRegistryVulListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRegistryVulListExport")
    
    
    return
}

func NewDescribeAssetImageRegistryVulListExportResponse() (response *DescribeAssetImageRegistryVulListExportResponse) {
    response = &DescribeAssetImageRegistryVulListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRegistryVulListExport
// 镜像仓库漏洞列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListExport(request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
    return c.DescribeAssetImageRegistryVulListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRegistryVulListExport
// 镜像仓库漏洞列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRegistryVulListExportWithContext(ctx context.Context, request *DescribeAssetImageRegistryVulListExportRequest) (response *DescribeAssetImageRegistryVulListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRegistryVulListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRegistryVulListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRegistryVulListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRiskListRequest() (request *DescribeAssetImageRiskListRequest) {
    request = &DescribeAssetImageRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskList")
    
    
    return
}

func NewDescribeAssetImageRiskListResponse() (response *DescribeAssetImageRiskListResponse) {
    response = &DescribeAssetImageRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRiskList
// 容器安全查询镜像风险列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskList(request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
    return c.DescribeAssetImageRiskListWithContext(context.Background(), request)
}

// DescribeAssetImageRiskList
// 容器安全查询镜像风险列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListWithContext(ctx context.Context, request *DescribeAssetImageRiskListRequest) (response *DescribeAssetImageRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageRiskListExportRequest() (request *DescribeAssetImageRiskListExportRequest) {
    request = &DescribeAssetImageRiskListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageRiskListExport")
    
    
    return
}

func NewDescribeAssetImageRiskListExportResponse() (response *DescribeAssetImageRiskListExportResponse) {
    response = &DescribeAssetImageRiskListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageRiskListExport
// 容器安全搜索查询镜像风险列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListExport(request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
    return c.DescribeAssetImageRiskListExportWithContext(context.Background(), request)
}

// DescribeAssetImageRiskListExport
// 容器安全搜索查询镜像风险列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageRiskListExportWithContext(ctx context.Context, request *DescribeAssetImageRiskListExportRequest) (response *DescribeAssetImageRiskListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageRiskListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageRiskListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageRiskListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanSettingRequest() (request *DescribeAssetImageScanSettingRequest) {
    request = &DescribeAssetImageScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanSetting")
    
    
    return
}

func NewDescribeAssetImageScanSettingResponse() (response *DescribeAssetImageScanSettingResponse) {
    response = &DescribeAssetImageScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageScanSetting
// 获取镜像扫描设置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetImageScanSetting(request *DescribeAssetImageScanSettingRequest) (response *DescribeAssetImageScanSettingResponse, err error) {
    return c.DescribeAssetImageScanSettingWithContext(context.Background(), request)
}

// DescribeAssetImageScanSetting
// 获取镜像扫描设置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetImageScanSettingWithContext(ctx context.Context, request *DescribeAssetImageScanSettingRequest) (response *DescribeAssetImageScanSettingResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanStatusRequest() (request *DescribeAssetImageScanStatusRequest) {
    request = &DescribeAssetImageScanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanStatus")
    
    
    return
}

func NewDescribeAssetImageScanStatusResponse() (response *DescribeAssetImageScanStatusResponse) {
    response = &DescribeAssetImageScanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageScanStatus
// 容器安全查询镜像扫描状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanStatus(request *DescribeAssetImageScanStatusRequest) (response *DescribeAssetImageScanStatusResponse, err error) {
    return c.DescribeAssetImageScanStatusWithContext(context.Background(), request)
}

// DescribeAssetImageScanStatus
// 容器安全查询镜像扫描状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanStatusWithContext(ctx context.Context, request *DescribeAssetImageScanStatusRequest) (response *DescribeAssetImageScanStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageScanTaskRequest() (request *DescribeAssetImageScanTaskRequest) {
    request = &DescribeAssetImageScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageScanTask")
    
    
    return
}

func NewDescribeAssetImageScanTaskResponse() (response *DescribeAssetImageScanTaskResponse) {
    response = &DescribeAssetImageScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageScanTask
// 查询正在一键扫描的镜像扫描taskid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanTask(request *DescribeAssetImageScanTaskRequest) (response *DescribeAssetImageScanTaskResponse, err error) {
    return c.DescribeAssetImageScanTaskWithContext(context.Background(), request)
}

// DescribeAssetImageScanTask
// 查询正在一键扫描的镜像扫描taskid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageScanTaskWithContext(ctx context.Context, request *DescribeAssetImageScanTaskRequest) (response *DescribeAssetImageScanTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageSimpleListRequest() (request *DescribeAssetImageSimpleListRequest) {
    request = &DescribeAssetImageSimpleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageSimpleList")
    
    
    return
}

func NewDescribeAssetImageSimpleListResponse() (response *DescribeAssetImageSimpleListResponse) {
    response = &DescribeAssetImageSimpleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageSimpleList
// 容器安全搜索查询镜像简略信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageSimpleList(request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
    return c.DescribeAssetImageSimpleListWithContext(context.Background(), request)
}

// DescribeAssetImageSimpleList
// 容器安全搜索查询镜像简略信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageSimpleListWithContext(ctx context.Context, request *DescribeAssetImageSimpleListRequest) (response *DescribeAssetImageSimpleListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageSimpleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageSimpleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageSimpleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVirusListRequest() (request *DescribeAssetImageVirusListRequest) {
    request = &DescribeAssetImageVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusList")
    
    
    return
}

func NewDescribeAssetImageVirusListResponse() (response *DescribeAssetImageVirusListResponse) {
    response = &DescribeAssetImageVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageVirusList
// 容器安全查询镜像病毒列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusList(request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
    return c.DescribeAssetImageVirusListWithContext(context.Background(), request)
}

// DescribeAssetImageVirusList
// 容器安全查询镜像病毒列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListWithContext(ctx context.Context, request *DescribeAssetImageVirusListRequest) (response *DescribeAssetImageVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVirusListExportRequest() (request *DescribeAssetImageVirusListExportRequest) {
    request = &DescribeAssetImageVirusListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVirusListExport")
    
    
    return
}

func NewDescribeAssetImageVirusListExportResponse() (response *DescribeAssetImageVirusListExportResponse) {
    response = &DescribeAssetImageVirusListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageVirusListExport
// 容器安全搜索查询镜像木马列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListExport(request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
    return c.DescribeAssetImageVirusListExportWithContext(context.Background(), request)
}

// DescribeAssetImageVirusListExport
// 容器安全搜索查询镜像木马列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVirusListExportWithContext(ctx context.Context, request *DescribeAssetImageVirusListExportRequest) (response *DescribeAssetImageVirusListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVirusListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVirusListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVirusListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVulListRequest() (request *DescribeAssetImageVulListRequest) {
    request = &DescribeAssetImageVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulList")
    
    
    return
}

func NewDescribeAssetImageVulListResponse() (response *DescribeAssetImageVulListResponse) {
    response = &DescribeAssetImageVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageVulList
// 容器安全查询镜像漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulList(request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
    return c.DescribeAssetImageVulListWithContext(context.Background(), request)
}

// DescribeAssetImageVulList
// 容器安全查询镜像漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListWithContext(ctx context.Context, request *DescribeAssetImageVulListRequest) (response *DescribeAssetImageVulListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetImageVulListExportRequest() (request *DescribeAssetImageVulListExportRequest) {
    request = &DescribeAssetImageVulListExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetImageVulListExport")
    
    
    return
}

func NewDescribeAssetImageVulListExportResponse() (response *DescribeAssetImageVulListExportResponse) {
    response = &DescribeAssetImageVulListExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetImageVulListExport
// 容器安全搜索查询镜像漏洞列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListExport(request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
    return c.DescribeAssetImageVulListExportWithContext(context.Background(), request)
}

// DescribeAssetImageVulListExport
// 容器安全搜索查询镜像漏洞列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetImageVulListExportWithContext(ctx context.Context, request *DescribeAssetImageVulListExportRequest) (response *DescribeAssetImageVulListExportResponse, err error) {
    if request == nil {
        request = NewDescribeAssetImageVulListExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetImageVulListExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetImageVulListExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetPortListRequest() (request *DescribeAssetPortListRequest) {
    request = &DescribeAssetPortListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetPortList")
    
    
    return
}

func NewDescribeAssetPortListResponse() (response *DescribeAssetPortListResponse) {
    response = &DescribeAssetPortListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetPortList
// 容器安全搜索查询端口占用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPortList(request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
    return c.DescribeAssetPortListWithContext(context.Background(), request)
}

// DescribeAssetPortList
// 容器安全搜索查询端口占用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPortListWithContext(ctx context.Context, request *DescribeAssetPortListRequest) (response *DescribeAssetPortListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetPortListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetPortList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetPortListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessListRequest() (request *DescribeAssetProcessListRequest) {
    request = &DescribeAssetProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetProcessList")
    
    
    return
}

func NewDescribeAssetProcessListResponse() (response *DescribeAssetProcessListResponse) {
    response = &DescribeAssetProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetProcessList
// 容器安全搜索查询进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetProcessList(request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    return c.DescribeAssetProcessListWithContext(context.Background(), request)
}

// DescribeAssetProcessList
// 容器安全搜索查询进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetProcessListWithContext(ctx context.Context, request *DescribeAssetProcessListRequest) (response *DescribeAssetProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSummaryRequest() (request *DescribeAssetSummaryRequest) {
    request = &DescribeAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetSummary")
    
    
    return
}

func NewDescribeAssetSummaryResponse() (response *DescribeAssetSummaryResponse) {
    response = &DescribeAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetSummary
// 查询账户容器、镜像等统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSummary(request *DescribeAssetSummaryRequest) (response *DescribeAssetSummaryResponse, err error) {
    return c.DescribeAssetSummaryWithContext(context.Background(), request)
}

// DescribeAssetSummary
// 查询账户容器、镜像等统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetSummaryWithContext(ctx context.Context, request *DescribeAssetSummaryRequest) (response *DescribeAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebServiceListRequest() (request *DescribeAssetWebServiceListRequest) {
    request = &DescribeAssetWebServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeAssetWebServiceList")
    
    
    return
}

func NewDescribeAssetWebServiceListResponse() (response *DescribeAssetWebServiceListResponse) {
    response = &DescribeAssetWebServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebServiceList
// 容器安全查询web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceList(request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
    return c.DescribeAssetWebServiceListWithContext(context.Background(), request)
}

// DescribeAssetWebServiceList
// 容器安全查询web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceListWithContext(ctx context.Context, request *DescribeAssetWebServiceListRequest) (response *DescribeAssetWebServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckItemListRequest() (request *DescribeCheckItemListRequest) {
    request = &DescribeCheckItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCheckItemList")
    
    
    return
}

func NewDescribeCheckItemListResponse() (response *DescribeCheckItemListResponse) {
    response = &DescribeCheckItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCheckItemList
// 查询所有检查项接口，返回总数和检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCheckItemList(request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
    return c.DescribeCheckItemListWithContext(context.Background(), request)
}

// DescribeCheckItemList
// 查询所有检查项接口，返回总数和检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCheckItemListWithContext(ctx context.Context, request *DescribeCheckItemListRequest) (response *DescribeCheckItemListResponse, err error) {
    if request == nil {
        request = NewDescribeCheckItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterDetail")
    
    
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterDetail
// 查询单个集群的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// 查询单个集群的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterSummaryRequest() (request *DescribeClusterSummaryRequest) {
    request = &DescribeClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeClusterSummary")
    
    
    return
}

func NewDescribeClusterSummaryResponse() (response *DescribeClusterSummaryResponse) {
    response = &DescribeClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterSummary
// 查询用户集群资产总览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterSummary(request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
    return c.DescribeClusterSummaryWithContext(context.Background(), request)
}

// DescribeClusterSummary
// 查询用户集群资产总览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterSummaryWithContext(ctx context.Context, request *DescribeClusterSummaryRequest) (response *DescribeClusterSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeClusterSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetDetailInfoRequest() (request *DescribeComplianceAssetDetailInfoRequest) {
    request = &DescribeComplianceAssetDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetDetailInfo")
    
    
    return
}

func NewDescribeComplianceAssetDetailInfoResponse() (response *DescribeComplianceAssetDetailInfoResponse) {
    response = &DescribeComplianceAssetDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceAssetDetailInfo
// 查询某个资产的详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetDetailInfo(request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
    return c.DescribeComplianceAssetDetailInfoWithContext(context.Background(), request)
}

// DescribeComplianceAssetDetailInfo
// 查询某个资产的详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetDetailInfoWithContext(ctx context.Context, request *DescribeComplianceAssetDetailInfoRequest) (response *DescribeComplianceAssetDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetDetailInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetListRequest() (request *DescribeComplianceAssetListRequest) {
    request = &DescribeComplianceAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetList")
    
    
    return
}

func NewDescribeComplianceAssetListResponse() (response *DescribeComplianceAssetListResponse) {
    response = &DescribeComplianceAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceAssetList
// 查询某类资产的列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    return c.DescribeComplianceAssetListWithContext(context.Background(), request)
}

// DescribeComplianceAssetList
// 查询某类资产的列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetListWithContext(ctx context.Context, request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetPolicyItemListRequest() (request *DescribeComplianceAssetPolicyItemListRequest) {
    request = &DescribeComplianceAssetPolicyItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceAssetPolicyItemList")
    
    
    return
}

func NewDescribeComplianceAssetPolicyItemListResponse() (response *DescribeComplianceAssetPolicyItemListResponse) {
    response = &DescribeComplianceAssetPolicyItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceAssetPolicyItemList
// 查询某资产下的检测项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetPolicyItemList(request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
    return c.DescribeComplianceAssetPolicyItemListWithContext(context.Background(), request)
}

// DescribeComplianceAssetPolicyItemList
// 查询某资产下的检测项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceAssetPolicyItemListWithContext(ctx context.Context, request *DescribeComplianceAssetPolicyItemListRequest) (response *DescribeComplianceAssetPolicyItemListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetPolicyItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetPolicyItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetPolicyItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePeriodTaskListRequest() (request *DescribeCompliancePeriodTaskListRequest) {
    request = &DescribeCompliancePeriodTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePeriodTaskList")
    
    
    return
}

func NewDescribeCompliancePeriodTaskListResponse() (response *DescribeCompliancePeriodTaskListResponse) {
    response = &DescribeCompliancePeriodTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCompliancePeriodTaskList
// 查询合规检测的定时任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePeriodTaskList(request *DescribeCompliancePeriodTaskListRequest) (response *DescribeCompliancePeriodTaskListResponse, err error) {
    return c.DescribeCompliancePeriodTaskListWithContext(context.Background(), request)
}

// DescribeCompliancePeriodTaskList
// 查询合规检测的定时任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePeriodTaskListWithContext(ctx context.Context, request *DescribeCompliancePeriodTaskListRequest) (response *DescribeCompliancePeriodTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePeriodTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePeriodTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePeriodTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePolicyItemAffectedAssetListRequest() (request *DescribeCompliancePolicyItemAffectedAssetListRequest) {
    request = &DescribeCompliancePolicyItemAffectedAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedAssetList")
    
    
    return
}

func NewDescribeCompliancePolicyItemAffectedAssetListResponse() (response *DescribeCompliancePolicyItemAffectedAssetListResponse) {
    response = &DescribeCompliancePolicyItemAffectedAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCompliancePolicyItemAffectedAssetList
// 按照 检测项 → 资产 的两级层次展开的第二层级：资产层级。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedAssetList(request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
    return c.DescribeCompliancePolicyItemAffectedAssetListWithContext(context.Background(), request)
}

// DescribeCompliancePolicyItemAffectedAssetList
// 按照 检测项 → 资产 的两级层次展开的第二层级：资产层级。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedAssetListWithContext(ctx context.Context, request *DescribeCompliancePolicyItemAffectedAssetListRequest) (response *DescribeCompliancePolicyItemAffectedAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePolicyItemAffectedAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePolicyItemAffectedAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePolicyItemAffectedAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompliancePolicyItemAffectedSummaryRequest() (request *DescribeCompliancePolicyItemAffectedSummaryRequest) {
    request = &DescribeCompliancePolicyItemAffectedSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeCompliancePolicyItemAffectedSummary")
    
    
    return
}

func NewDescribeCompliancePolicyItemAffectedSummaryResponse() (response *DescribeCompliancePolicyItemAffectedSummaryResponse) {
    response = &DescribeCompliancePolicyItemAffectedSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCompliancePolicyItemAffectedSummary
// 按照 检测项 → 资产 的两级层次展开的第一层级：检测项层级。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedSummary(request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
    return c.DescribeCompliancePolicyItemAffectedSummaryWithContext(context.Background(), request)
}

// DescribeCompliancePolicyItemAffectedSummary
// 按照 检测项 → 资产 的两级层次展开的第一层级：检测项层级。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeCompliancePolicyItemAffectedSummaryWithContext(ctx context.Context, request *DescribeCompliancePolicyItemAffectedSummaryRequest) (response *DescribeCompliancePolicyItemAffectedSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCompliancePolicyItemAffectedSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompliancePolicyItemAffectedSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompliancePolicyItemAffectedSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceScanFailedAssetListRequest() (request *DescribeComplianceScanFailedAssetListRequest) {
    request = &DescribeComplianceScanFailedAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceScanFailedAssetList")
    
    
    return
}

func NewDescribeComplianceScanFailedAssetListResponse() (response *DescribeComplianceScanFailedAssetListResponse) {
    response = &DescribeComplianceScanFailedAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceScanFailedAssetList
// 按照 资产 → 检测项 二层结构展示的信息。这里查询第一层 资产的通过率汇总信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceScanFailedAssetList(request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
    return c.DescribeComplianceScanFailedAssetListWithContext(context.Background(), request)
}

// DescribeComplianceScanFailedAssetList
// 按照 资产 → 检测项 二层结构展示的信息。这里查询第一层 资产的通过率汇总信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceScanFailedAssetListWithContext(ctx context.Context, request *DescribeComplianceScanFailedAssetListRequest) (response *DescribeComplianceScanFailedAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceScanFailedAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceScanFailedAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceScanFailedAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceTaskAssetSummaryRequest() (request *DescribeComplianceTaskAssetSummaryRequest) {
    request = &DescribeComplianceTaskAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskAssetSummary")
    
    
    return
}

func NewDescribeComplianceTaskAssetSummaryResponse() (response *DescribeComplianceTaskAssetSummaryResponse) {
    response = &DescribeComplianceTaskAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceTaskAssetSummary
// 查询上次任务的资产通过率汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceTaskAssetSummary(request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
    return c.DescribeComplianceTaskAssetSummaryWithContext(context.Background(), request)
}

// DescribeComplianceTaskAssetSummary
// 查询上次任务的资产通过率汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceTaskAssetSummaryWithContext(ctx context.Context, request *DescribeComplianceTaskAssetSummaryRequest) (response *DescribeComplianceTaskAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceTaskAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceTaskAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceTaskAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceTaskPolicyItemSummaryListRequest() (request *DescribeComplianceTaskPolicyItemSummaryListRequest) {
    request = &DescribeComplianceTaskPolicyItemSummaryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceTaskPolicyItemSummaryList")
    
    
    return
}

func NewDescribeComplianceTaskPolicyItemSummaryListResponse() (response *DescribeComplianceTaskPolicyItemSummaryListResponse) {
    response = &DescribeComplianceTaskPolicyItemSummaryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceTaskPolicyItemSummaryList
// 查询最近一次任务发现的检测项的汇总信息列表，按照 检测项 → 资产 的两级层次展开。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeComplianceTaskPolicyItemSummaryList(request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
    return c.DescribeComplianceTaskPolicyItemSummaryListWithContext(context.Background(), request)
}

// DescribeComplianceTaskPolicyItemSummaryList
// 查询最近一次任务发现的检测项的汇总信息列表，按照 检测项 → 资产 的两级层次展开。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeComplianceTaskPolicyItemSummaryListWithContext(ctx context.Context, request *DescribeComplianceTaskPolicyItemSummaryListRequest) (response *DescribeComplianceTaskPolicyItemSummaryListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceTaskPolicyItemSummaryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceTaskPolicyItemSummaryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceTaskPolicyItemSummaryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceWhitelistItemListRequest() (request *DescribeComplianceWhitelistItemListRequest) {
    request = &DescribeComplianceWhitelistItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeComplianceWhitelistItemList")
    
    
    return
}

func NewDescribeComplianceWhitelistItemListResponse() (response *DescribeComplianceWhitelistItemListResponse) {
    response = &DescribeComplianceWhitelistItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceWhitelistItemList
// 查询白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceWhitelistItemList(request *DescribeComplianceWhitelistItemListRequest) (response *DescribeComplianceWhitelistItemListResponse, err error) {
    return c.DescribeComplianceWhitelistItemListWithContext(context.Background(), request)
}

// DescribeComplianceWhitelistItemList
// 查询白名单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeComplianceWhitelistItemListWithContext(ctx context.Context, request *DescribeComplianceWhitelistItemListRequest) (response *DescribeComplianceWhitelistItemListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceWhitelistItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceWhitelistItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceWhitelistItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerAssetSummaryRequest() (request *DescribeContainerAssetSummaryRequest) {
    request = &DescribeContainerAssetSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerAssetSummary")
    
    
    return
}

func NewDescribeContainerAssetSummaryResponse() (response *DescribeContainerAssetSummaryResponse) {
    response = &DescribeContainerAssetSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerAssetSummary
// 查询容器资产概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerAssetSummary(request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
    return c.DescribeContainerAssetSummaryWithContext(context.Background(), request)
}

// DescribeContainerAssetSummary
// 查询容器资产概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerAssetSummaryWithContext(ctx context.Context, request *DescribeContainerAssetSummaryRequest) (response *DescribeContainerAssetSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeContainerAssetSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerAssetSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerAssetSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerSecEventSummaryRequest() (request *DescribeContainerSecEventSummaryRequest) {
    request = &DescribeContainerSecEventSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeContainerSecEventSummary")
    
    
    return
}

func NewDescribeContainerSecEventSummaryResponse() (response *DescribeContainerSecEventSummaryResponse) {
    response = &DescribeContainerSecEventSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerSecEventSummary
// 查询容器安全未处理事件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerSecEventSummary(request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
    return c.DescribeContainerSecEventSummaryWithContext(context.Background(), request)
}

// DescribeContainerSecEventSummary
// 查询容器安全未处理事件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerSecEventSummaryWithContext(ctx context.Context, request *DescribeContainerSecEventSummaryRequest) (response *DescribeContainerSecEventSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeContainerSecEventSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerSecEventSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerSecEventSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventDetailRequest() (request *DescribeEscapeEventDetailRequest) {
    request = &DescribeEscapeEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventDetail")
    
    
    return
}

func NewDescribeEscapeEventDetailResponse() (response *DescribeEscapeEventDetailResponse) {
    response = &DescribeEscapeEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEscapeEventDetail
// DescribeEscapeEventDetail  查询容器逃逸事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeEscapeEventDetail(request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
    return c.DescribeEscapeEventDetailWithContext(context.Background(), request)
}

// DescribeEscapeEventDetail
// DescribeEscapeEventDetail  查询容器逃逸事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeEscapeEventDetailWithContext(ctx context.Context, request *DescribeEscapeEventDetailRequest) (response *DescribeEscapeEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventInfoRequest() (request *DescribeEscapeEventInfoRequest) {
    request = &DescribeEscapeEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventInfo")
    
    
    return
}

func NewDescribeEscapeEventInfoResponse() (response *DescribeEscapeEventInfoResponse) {
    response = &DescribeEscapeEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEscapeEventInfo
// DescribeEscapeEventInfo 查询容器逃逸事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventInfo(request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
    return c.DescribeEscapeEventInfoWithContext(context.Background(), request)
}

// DescribeEscapeEventInfo
// DescribeEscapeEventInfo 查询容器逃逸事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventInfoWithContext(ctx context.Context, request *DescribeEscapeEventInfoRequest) (response *DescribeEscapeEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeEventsExportRequest() (request *DescribeEscapeEventsExportRequest) {
    request = &DescribeEscapeEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeEventsExport")
    
    
    return
}

func NewDescribeEscapeEventsExportResponse() (response *DescribeEscapeEventsExportResponse) {
    response = &DescribeEscapeEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEscapeEventsExport
// DescribeEscapeEventsExport  查询容器逃逸事件列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventsExport(request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
    return c.DescribeEscapeEventsExportWithContext(context.Background(), request)
}

// DescribeEscapeEventsExport
// DescribeEscapeEventsExport  查询容器逃逸事件列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEscapeEventsExportWithContext(ctx context.Context, request *DescribeEscapeEventsExportRequest) (response *DescribeEscapeEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeRuleInfoRequest() (request *DescribeEscapeRuleInfoRequest) {
    request = &DescribeEscapeRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeRuleInfo")
    
    
    return
}

func NewDescribeEscapeRuleInfoResponse() (response *DescribeEscapeRuleInfoResponse) {
    response = &DescribeEscapeRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEscapeRuleInfo
// DescribeEscapeRuleInfo 查询容器逃逸扫描规则信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeRuleInfo(request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
    return c.DescribeEscapeRuleInfoWithContext(context.Background(), request)
}

// DescribeEscapeRuleInfo
// DescribeEscapeRuleInfo 查询容器逃逸扫描规则信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeRuleInfoWithContext(ctx context.Context, request *DescribeEscapeRuleInfoRequest) (response *DescribeEscapeRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEscapeSafeStateRequest() (request *DescribeEscapeSafeStateRequest) {
    request = &DescribeEscapeSafeStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeEscapeSafeState")
    
    
    return
}

func NewDescribeEscapeSafeStateResponse() (response *DescribeEscapeSafeStateResponse) {
    response = &DescribeEscapeSafeStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEscapeSafeState
// DescribeEscapeSafeState 查询容器逃逸安全状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeSafeState(request *DescribeEscapeSafeStateRequest) (response *DescribeEscapeSafeStateResponse, err error) {
    return c.DescribeEscapeSafeStateWithContext(context.Background(), request)
}

// DescribeEscapeSafeState
// DescribeEscapeSafeState 查询容器逃逸安全状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeEscapeSafeStateWithContext(ctx context.Context, request *DescribeEscapeSafeStateRequest) (response *DescribeEscapeSafeStateResponse, err error) {
    if request == nil {
        request = NewDescribeEscapeSafeStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEscapeSafeState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEscapeSafeStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportJobResultRequest() (request *DescribeExportJobResultRequest) {
    request = &DescribeExportJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeExportJobResult")
    
    
    return
}

func NewDescribeExportJobResultResponse() (response *DescribeExportJobResultResponse) {
    response = &DescribeExportJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExportJobResult
// 查询导出任务的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeExportJobResult(request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
    return c.DescribeExportJobResultWithContext(context.Background(), request)
}

// DescribeExportJobResult
// 查询导出任务的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeExportJobResultWithContext(ctx context.Context, request *DescribeExportJobResultRequest) (response *DescribeExportJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeExportJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAuthorizedInfoRequest() (request *DescribeImageAuthorizedInfoRequest) {
    request = &DescribeImageAuthorizedInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageAuthorizedInfo")
    
    
    return
}

func NewDescribeImageAuthorizedInfoResponse() (response *DescribeImageAuthorizedInfoResponse) {
    response = &DescribeImageAuthorizedInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageAuthorizedInfo
// DescribeImageAuthorizedInfo  查询镜像授权信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeImageAuthorizedInfo(request *DescribeImageAuthorizedInfoRequest) (response *DescribeImageAuthorizedInfoResponse, err error) {
    return c.DescribeImageAuthorizedInfoWithContext(context.Background(), request)
}

// DescribeImageAuthorizedInfo
// DescribeImageAuthorizedInfo  查询镜像授权信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeImageAuthorizedInfoWithContext(ctx context.Context, request *DescribeImageAuthorizedInfoRequest) (response *DescribeImageAuthorizedInfoResponse, err error) {
    if request == nil {
        request = NewDescribeImageAuthorizedInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAuthorizedInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAuthorizedInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRegistryTimingScanTaskRequest() (request *DescribeImageRegistryTimingScanTaskRequest) {
    request = &DescribeImageRegistryTimingScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRegistryTimingScanTask")
    
    
    return
}

func NewDescribeImageRegistryTimingScanTaskResponse() (response *DescribeImageRegistryTimingScanTaskResponse) {
    response = &DescribeImageRegistryTimingScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageRegistryTimingScanTask
// 镜像仓库查看定时任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryTimingScanTask(request *DescribeImageRegistryTimingScanTaskRequest) (response *DescribeImageRegistryTimingScanTaskResponse, err error) {
    return c.DescribeImageRegistryTimingScanTaskWithContext(context.Background(), request)
}

// DescribeImageRegistryTimingScanTask
// 镜像仓库查看定时任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeImageRegistryTimingScanTaskWithContext(ctx context.Context, request *DescribeImageRegistryTimingScanTaskRequest) (response *DescribeImageRegistryTimingScanTaskResponse, err error) {
    if request == nil {
        request = NewDescribeImageRegistryTimingScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRegistryTimingScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRegistryTimingScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRiskSummaryRequest() (request *DescribeImageRiskSummaryRequest) {
    request = &DescribeImageRiskSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskSummary")
    
    
    return
}

func NewDescribeImageRiskSummaryResponse() (response *DescribeImageRiskSummaryResponse) {
    response = &DescribeImageRiskSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageRiskSummary
// 查询本地镜像风险概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageRiskSummary(request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
    return c.DescribeImageRiskSummaryWithContext(context.Background(), request)
}

// DescribeImageRiskSummary
// 查询本地镜像风险概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageRiskSummaryWithContext(ctx context.Context, request *DescribeImageRiskSummaryRequest) (response *DescribeImageRiskSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRiskSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRiskSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRiskSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRiskTendencyRequest() (request *DescribeImageRiskTendencyRequest) {
    request = &DescribeImageRiskTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageRiskTendency")
    
    
    return
}

func NewDescribeImageRiskTendencyResponse() (response *DescribeImageRiskTendencyResponse) {
    response = &DescribeImageRiskTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageRiskTendency
// 查询容器安全本地镜像风险趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeImageRiskTendency(request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
    return c.DescribeImageRiskTendencyWithContext(context.Background(), request)
}

// DescribeImageRiskTendency
// 查询容器安全本地镜像风险趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeImageRiskTendencyWithContext(ctx context.Context, request *DescribeImageRiskTendencyRequest) (response *DescribeImageRiskTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeImageRiskTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRiskTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRiskTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSimpleListRequest() (request *DescribeImageSimpleListRequest) {
    request = &DescribeImageSimpleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeImageSimpleList")
    
    
    return
}

func NewDescribeImageSimpleListResponse() (response *DescribeImageSimpleListResponse) {
    response = &DescribeImageSimpleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageSimpleList
// DescribeImageSimpleList 查询全部镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageSimpleList(request *DescribeImageSimpleListRequest) (response *DescribeImageSimpleListResponse, err error) {
    return c.DescribeImageSimpleListWithContext(context.Background(), request)
}

// DescribeImageSimpleList
// DescribeImageSimpleList 查询全部镜像列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeImageSimpleListWithContext(ctx context.Context, request *DescribeImageSimpleListRequest) (response *DescribeImageSimpleListResponse, err error) {
    if request == nil {
        request = NewDescribeImageSimpleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageSimpleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageSimpleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostPayDetailRequest() (request *DescribePostPayDetailRequest) {
    request = &DescribePostPayDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePostPayDetail")
    
    
    return
}

func NewDescribePostPayDetailResponse() (response *DescribePostPayDetailResponse) {
    response = &DescribePostPayDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePostPayDetail
// DescribePostPayDetail  查询后付费详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribePostPayDetail(request *DescribePostPayDetailRequest) (response *DescribePostPayDetailResponse, err error) {
    return c.DescribePostPayDetailWithContext(context.Background(), request)
}

// DescribePostPayDetail
// DescribePostPayDetail  查询后付费详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribePostPayDetailWithContext(ctx context.Context, request *DescribePostPayDetailRequest) (response *DescribePostPayDetailResponse, err error) {
    if request == nil {
        request = NewDescribePostPayDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePostPayDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePostPayDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeProVersionInfo")
    
    
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProVersionInfo
// DescribeProVersionInfo  查询专业版需购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    return c.DescribeProVersionInfoWithContext(context.Background(), request)
}

// DescribeProVersionInfo
// DescribeProVersionInfo  查询专业版需购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeProVersionInfoWithContext(ctx context.Context, request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProVersionInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProVersionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurchaseStateInfoRequest() (request *DescribePurchaseStateInfoRequest) {
    request = &DescribePurchaseStateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribePurchaseStateInfo")
    
    
    return
}

func NewDescribePurchaseStateInfoResponse() (response *DescribePurchaseStateInfoResponse) {
    response = &DescribePurchaseStateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurchaseStateInfo
// DescribePurchaseStateInfo 查询容器安全服务已购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePurchaseStateInfo(request *DescribePurchaseStateInfoRequest) (response *DescribePurchaseStateInfoResponse, err error) {
    return c.DescribePurchaseStateInfoWithContext(context.Background(), request)
}

// DescribePurchaseStateInfo
// DescribePurchaseStateInfo 查询容器安全服务已购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePurchaseStateInfoWithContext(ctx context.Context, request *DescribePurchaseStateInfoRequest) (response *DescribePurchaseStateInfoResponse, err error) {
    if request == nil {
        request = NewDescribePurchaseStateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurchaseStateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurchaseStateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRefreshTaskRequest() (request *DescribeRefreshTaskRequest) {
    request = &DescribeRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRefreshTask")
    
    
    return
}

func NewDescribeRefreshTaskResponse() (response *DescribeRefreshTaskResponse) {
    response = &DescribeRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRefreshTask
// 查询刷新任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRefreshTask(request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
    return c.DescribeRefreshTaskWithContext(context.Background(), request)
}

// DescribeRefreshTask
// 查询刷新任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRefreshTaskWithContext(ctx context.Context, request *DescribeRefreshTaskRequest) (response *DescribeRefreshTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellDetailRequest() (request *DescribeReverseShellDetailRequest) {
    request = &DescribeReverseShellDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellDetail")
    
    
    return
}

func NewDescribeReverseShellDetailResponse() (response *DescribeReverseShellDetailResponse) {
    response = &DescribeReverseShellDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellDetail
// 查询运行时反弹shell事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellDetail(request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
    return c.DescribeReverseShellDetailWithContext(context.Background(), request)
}

// DescribeReverseShellDetail
// 查询运行时反弹shell事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellDetailWithContext(ctx context.Context, request *DescribeReverseShellDetailRequest) (response *DescribeReverseShellDetailResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
    request = &DescribeReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEvents")
    
    
    return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
    response = &DescribeReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellEvents
// 查询运行时反弹shell事件列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    return c.DescribeReverseShellEventsWithContext(context.Background(), request)
}

// DescribeReverseShellEvents
// 查询运行时反弹shell事件列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsWithContext(ctx context.Context, request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsExportRequest() (request *DescribeReverseShellEventsExportRequest) {
    request = &DescribeReverseShellEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellEventsExport")
    
    
    return
}

func NewDescribeReverseShellEventsExportResponse() (response *DescribeReverseShellEventsExportResponse) {
    response = &DescribeReverseShellEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellEventsExport
// 查询运行时反弹shell事件列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsExport(request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
    return c.DescribeReverseShellEventsExportWithContext(context.Background(), request)
}

// DescribeReverseShellEventsExport
// 查询运行时反弹shell事件列表信息导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEventsExportWithContext(ctx context.Context, request *DescribeReverseShellEventsExportRequest) (response *DescribeReverseShellEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellWhiteListDetailRequest() (request *DescribeReverseShellWhiteListDetailRequest) {
    request = &DescribeReverseShellWhiteListDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteListDetail")
    
    
    return
}

func NewDescribeReverseShellWhiteListDetailResponse() (response *DescribeReverseShellWhiteListDetailResponse) {
    response = &DescribeReverseShellWhiteListDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellWhiteListDetail
// 查询运行时反弹shell白名单详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellWhiteListDetail(request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
    return c.DescribeReverseShellWhiteListDetailWithContext(context.Background(), request)
}

// DescribeReverseShellWhiteListDetail
// 查询运行时反弹shell白名单详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeReverseShellWhiteListDetailWithContext(ctx context.Context, request *DescribeReverseShellWhiteListDetailRequest) (response *DescribeReverseShellWhiteListDetailResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellWhiteListDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellWhiteListDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellWhiteListDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellWhiteListsRequest() (request *DescribeReverseShellWhiteListsRequest) {
    request = &DescribeReverseShellWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeReverseShellWhiteLists")
    
    
    return
}

func NewDescribeReverseShellWhiteListsResponse() (response *DescribeReverseShellWhiteListsResponse) {
    response = &DescribeReverseShellWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellWhiteLists
// 查询运行时运行时反弹shell白名单列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReverseShellWhiteLists(request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
    return c.DescribeReverseShellWhiteListsWithContext(context.Background(), request)
}

// DescribeReverseShellWhiteLists
// 查询运行时运行时反弹shell白名单列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeReverseShellWhiteListsWithContext(ctx context.Context, request *DescribeReverseShellWhiteListsRequest) (response *DescribeReverseShellWhiteListsResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskListRequest() (request *DescribeRiskListRequest) {
    request = &DescribeRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskList")
    
    
    return
}

func NewDescribeRiskListResponse() (response *DescribeRiskListResponse) {
    response = &DescribeRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskList
// 查询最近一次任务发现的风险项的信息列表，支持根据特殊字段进行过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskList(request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    return c.DescribeRiskListWithContext(context.Background(), request)
}

// DescribeRiskList
// 查询最近一次任务发现的风险项的信息列表，支持根据特殊字段进行过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRiskListWithContext(ctx context.Context, request *DescribeRiskListRequest) (response *DescribeRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallDetailRequest() (request *DescribeRiskSyscallDetailRequest) {
    request = &DescribeRiskSyscallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallDetail")
    
    
    return
}

func NewDescribeRiskSyscallDetailResponse() (response *DescribeRiskSyscallDetailResponse) {
    response = &DescribeRiskSyscallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallDetail
// 查询高危系统调用事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallDetail(request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
    return c.DescribeRiskSyscallDetailWithContext(context.Background(), request)
}

// DescribeRiskSyscallDetail
// 查询高危系统调用事件详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallDetailWithContext(ctx context.Context, request *DescribeRiskSyscallDetailRequest) (response *DescribeRiskSyscallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallEventsRequest() (request *DescribeRiskSyscallEventsRequest) {
    request = &DescribeRiskSyscallEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEvents")
    
    
    return
}

func NewDescribeRiskSyscallEventsResponse() (response *DescribeRiskSyscallEventsResponse) {
    response = &DescribeRiskSyscallEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallEvents
// 查询运行时运行时高危系统调用列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEvents(request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
    return c.DescribeRiskSyscallEventsWithContext(context.Background(), request)
}

// DescribeRiskSyscallEvents
// 查询运行时运行时高危系统调用列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsWithContext(ctx context.Context, request *DescribeRiskSyscallEventsRequest) (response *DescribeRiskSyscallEventsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallEventsExportRequest() (request *DescribeRiskSyscallEventsExportRequest) {
    request = &DescribeRiskSyscallEventsExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallEventsExport")
    
    
    return
}

func NewDescribeRiskSyscallEventsExportResponse() (response *DescribeRiskSyscallEventsExportResponse) {
    response = &DescribeRiskSyscallEventsExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallEventsExport
// 运行时高危系统调用列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsExport(request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
    return c.DescribeRiskSyscallEventsExportWithContext(context.Background(), request)
}

// DescribeRiskSyscallEventsExport
// 运行时高危系统调用列表导出
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallEventsExportWithContext(ctx context.Context, request *DescribeRiskSyscallEventsExportRequest) (response *DescribeRiskSyscallEventsExportResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallEventsExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallEventsExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallEventsExportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallNamesRequest() (request *DescribeRiskSyscallNamesRequest) {
    request = &DescribeRiskSyscallNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallNames")
    
    
    return
}

func NewDescribeRiskSyscallNamesResponse() (response *DescribeRiskSyscallNamesResponse) {
    response = &DescribeRiskSyscallNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallNames
// 查询运行时高危系统调用系统名称列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallNames(request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
    return c.DescribeRiskSyscallNamesWithContext(context.Background(), request)
}

// DescribeRiskSyscallNames
// 查询运行时高危系统调用系统名称列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskSyscallNamesWithContext(ctx context.Context, request *DescribeRiskSyscallNamesRequest) (response *DescribeRiskSyscallNamesResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallNamesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallNames require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallNamesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallWhiteListDetailRequest() (request *DescribeRiskSyscallWhiteListDetailRequest) {
    request = &DescribeRiskSyscallWhiteListDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteListDetail")
    
    
    return
}

func NewDescribeRiskSyscallWhiteListDetailResponse() (response *DescribeRiskSyscallWhiteListDetailResponse) {
    response = &DescribeRiskSyscallWhiteListDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallWhiteListDetail
// 查询运行时高危系统调用白名单详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallWhiteListDetail(request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
    return c.DescribeRiskSyscallWhiteListDetailWithContext(context.Background(), request)
}

// DescribeRiskSyscallWhiteListDetail
// 查询运行时高危系统调用白名单详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) DescribeRiskSyscallWhiteListDetailWithContext(ctx context.Context, request *DescribeRiskSyscallWhiteListDetailRequest) (response *DescribeRiskSyscallWhiteListDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallWhiteListDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallWhiteListDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallWhiteListDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskSyscallWhiteListsRequest() (request *DescribeRiskSyscallWhiteListsRequest) {
    request = &DescribeRiskSyscallWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeRiskSyscallWhiteLists")
    
    
    return
}

func NewDescribeRiskSyscallWhiteListsResponse() (response *DescribeRiskSyscallWhiteListsResponse) {
    response = &DescribeRiskSyscallWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskSyscallWhiteLists
// 查询运行时高危系统调用白名单列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskSyscallWhiteLists(request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
    return c.DescribeRiskSyscallWhiteListsWithContext(context.Background(), request)
}

// DescribeRiskSyscallWhiteLists
// 查询运行时高危系统调用白名单列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeRiskSyscallWhiteListsWithContext(ctx context.Context, request *DescribeRiskSyscallWhiteListsRequest) (response *DescribeRiskSyscallWhiteListsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskSyscallWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskSyscallWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskSyscallWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecEventsTendencyRequest() (request *DescribeSecEventsTendencyRequest) {
    request = &DescribeSecEventsTendencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeSecEventsTendency")
    
    
    return
}

func NewDescribeSecEventsTendencyResponse() (response *DescribeSecEventsTendencyResponse) {
    response = &DescribeSecEventsTendencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecEventsTendency
// 查询容器运行时安全事件趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeSecEventsTendency(request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
    return c.DescribeSecEventsTendencyWithContext(context.Background(), request)
}

// DescribeSecEventsTendency
// 查询容器运行时安全事件趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"
func (c *Client) DescribeSecEventsTendencyWithContext(ctx context.Context, request *DescribeSecEventsTendencyRequest) (response *DescribeSecEventsTendencyResponse, err error) {
    if request == nil {
        request = NewDescribeSecEventsTendencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecEventsTendency require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecEventsTendencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultSummaryRequest() (request *DescribeTaskResultSummaryRequest) {
    request = &DescribeTaskResultSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeTaskResultSummary")
    
    
    return
}

func NewDescribeTaskResultSummaryResponse() (response *DescribeTaskResultSummaryResponse) {
    response = &DescribeTaskResultSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskResultSummary
// 查询检查结果总览，返回受影响的节点数量，返回7天的数据，总共7个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskResultSummary(request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
    return c.DescribeTaskResultSummaryWithContext(context.Background(), request)
}

// DescribeTaskResultSummary
// 查询检查结果总览，返回受影响的节点数量，返回7天的数据，总共7个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskResultSummaryWithContext(ctx context.Context, request *DescribeTaskResultSummaryRequest) (response *DescribeTaskResultSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResultSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnfinishRefreshTaskRequest() (request *DescribeUnfinishRefreshTaskRequest) {
    request = &DescribeUnfinishRefreshTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeUnfinishRefreshTask")
    
    
    return
}

func NewDescribeUnfinishRefreshTaskResponse() (response *DescribeUnfinishRefreshTaskResponse) {
    response = &DescribeUnfinishRefreshTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnfinishRefreshTask
// 查询未完成的刷新资产任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUnfinishRefreshTask(request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
    return c.DescribeUnfinishRefreshTaskWithContext(context.Background(), request)
}

// DescribeUnfinishRefreshTask
// 查询未完成的刷新资产任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUnfinishRefreshTaskWithContext(ctx context.Context, request *DescribeUnfinishRefreshTaskRequest) (response *DescribeUnfinishRefreshTaskResponse, err error) {
    if request == nil {
        request = NewDescribeUnfinishRefreshTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnfinishRefreshTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnfinishRefreshTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserClusterRequest() (request *DescribeUserClusterRequest) {
    request = &DescribeUserClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeUserCluster")
    
    
    return
}

func NewDescribeUserClusterResponse() (response *DescribeUserClusterResponse) {
    response = &DescribeUserClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserCluster
// 安全概览和集群安全页进入调用该接口，查询用户集群相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCluster(request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
    return c.DescribeUserClusterWithContext(context.Background(), request)
}

// DescribeUserCluster
// 安全概览和集群安全页进入调用该接口，查询用户集群相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserClusterWithContext(ctx context.Context, request *DescribeUserClusterRequest) (response *DescribeUserClusterResponse, err error) {
    if request == nil {
        request = NewDescribeUserClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeValueAddedSrvInfoRequest() (request *DescribeValueAddedSrvInfoRequest) {
    request = &DescribeValueAddedSrvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeValueAddedSrvInfo")
    
    
    return
}

func NewDescribeValueAddedSrvInfoResponse() (response *DescribeValueAddedSrvInfoResponse) {
    response = &DescribeValueAddedSrvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeValueAddedSrvInfo
// DescribeValueAddedSrvInfo查询增值服务需购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeValueAddedSrvInfo(request *DescribeValueAddedSrvInfoRequest) (response *DescribeValueAddedSrvInfoResponse, err error) {
    return c.DescribeValueAddedSrvInfoWithContext(context.Background(), request)
}

// DescribeValueAddedSrvInfo
// DescribeValueAddedSrvInfo查询增值服务需购买信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeValueAddedSrvInfoWithContext(ctx context.Context, request *DescribeValueAddedSrvInfoRequest) (response *DescribeValueAddedSrvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeValueAddedSrvInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeValueAddedSrvInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeValueAddedSrvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusDetailRequest() (request *DescribeVirusDetailRequest) {
    request = &DescribeVirusDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusDetail")
    
    
    return
}

func NewDescribeVirusDetailResponse() (response *DescribeVirusDetailResponse) {
    response = &DescribeVirusDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusDetail
// 运行时查询木马文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusDetail(request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
    return c.DescribeVirusDetailWithContext(context.Background(), request)
}

// DescribeVirusDetail
// 运行时查询木马文件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusDetailWithContext(ctx context.Context, request *DescribeVirusDetailRequest) (response *DescribeVirusDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVirusDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusListRequest() (request *DescribeVirusListRequest) {
    request = &DescribeVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusList")
    
    
    return
}

func NewDescribeVirusListResponse() (response *DescribeVirusListResponse) {
    response = &DescribeVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusList
// 查询运行时文件查杀事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusList(request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
    return c.DescribeVirusListWithContext(context.Background(), request)
}

// DescribeVirusList
// 查询运行时文件查杀事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusListWithContext(ctx context.Context, request *DescribeVirusListRequest) (response *DescribeVirusListResponse, err error) {
    if request == nil {
        request = NewDescribeVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusMonitorSettingRequest() (request *DescribeVirusMonitorSettingRequest) {
    request = &DescribeVirusMonitorSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusMonitorSetting")
    
    
    return
}

func NewDescribeVirusMonitorSettingResponse() (response *DescribeVirusMonitorSettingResponse) {
    response = &DescribeVirusMonitorSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusMonitorSetting
// 运行时查询文件查杀实时监控设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusMonitorSetting(request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
    return c.DescribeVirusMonitorSettingWithContext(context.Background(), request)
}

// DescribeVirusMonitorSetting
// 运行时查询文件查杀实时监控设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusMonitorSettingWithContext(ctx context.Context, request *DescribeVirusMonitorSettingRequest) (response *DescribeVirusMonitorSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusMonitorSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusMonitorSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusMonitorSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanSettingRequest() (request *DescribeVirusScanSettingRequest) {
    request = &DescribeVirusScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanSetting")
    
    
    return
}

func NewDescribeVirusScanSettingResponse() (response *DescribeVirusScanSettingResponse) {
    response = &DescribeVirusScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusScanSetting
// 运行时查询文件查杀设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanSetting(request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
    return c.DescribeVirusScanSettingWithContext(context.Background(), request)
}

// DescribeVirusScanSetting
// 运行时查询文件查杀设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanSettingWithContext(ctx context.Context, request *DescribeVirusScanSettingRequest) (response *DescribeVirusScanSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanTaskStatusRequest() (request *DescribeVirusScanTaskStatusRequest) {
    request = &DescribeVirusScanTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTaskStatus")
    
    
    return
}

func NewDescribeVirusScanTaskStatusResponse() (response *DescribeVirusScanTaskStatusResponse) {
    response = &DescribeVirusScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusScanTaskStatus
// 运行时查询文件查杀任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanTaskStatus(request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
    return c.DescribeVirusScanTaskStatusWithContext(context.Background(), request)
}

// DescribeVirusScanTaskStatus
// 运行时查询文件查杀任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusScanTaskStatusWithContext(ctx context.Context, request *DescribeVirusScanTaskStatusRequest) (response *DescribeVirusScanTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusScanTimeoutSettingRequest() (request *DescribeVirusScanTimeoutSettingRequest) {
    request = &DescribeVirusScanTimeoutSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusScanTimeoutSetting")
    
    
    return
}

func NewDescribeVirusScanTimeoutSettingResponse() (response *DescribeVirusScanTimeoutSettingResponse) {
    response = &DescribeVirusScanTimeoutSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusScanTimeoutSetting
// 运行时文件扫描超时设置查询
//
// 可能返回的错误码:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeVirusScanTimeoutSetting(request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
    return c.DescribeVirusScanTimeoutSettingWithContext(context.Background(), request)
}

// DescribeVirusScanTimeoutSetting
// 运行时文件扫描超时设置查询
//
// 可能返回的错误码:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeVirusScanTimeoutSettingWithContext(ctx context.Context, request *DescribeVirusScanTimeoutSettingRequest) (response *DescribeVirusScanTimeoutSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVirusScanTimeoutSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusScanTimeoutSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusScanTimeoutSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusSummaryRequest() (request *DescribeVirusSummaryRequest) {
    request = &DescribeVirusSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusSummary")
    
    
    return
}

func NewDescribeVirusSummaryResponse() (response *DescribeVirusSummaryResponse) {
    response = &DescribeVirusSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusSummary
// 运行时查询木马概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSummary(request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
    return c.DescribeVirusSummaryWithContext(context.Background(), request)
}

// DescribeVirusSummary
// 运行时查询木马概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusSummaryWithContext(ctx context.Context, request *DescribeVirusSummaryRequest) (response *DescribeVirusSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeVirusSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirusTaskListRequest() (request *DescribeVirusTaskListRequest) {
    request = &DescribeVirusTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeVirusTaskList")
    
    
    return
}

func NewDescribeVirusTaskListResponse() (response *DescribeVirusTaskListResponse) {
    response = &DescribeVirusTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVirusTaskList
// 运行时查询文件查杀任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusTaskList(request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
    return c.DescribeVirusTaskListWithContext(context.Background(), request)
}

// DescribeVirusTaskList
// 运行时查询文件查杀任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVirusTaskListWithContext(ctx context.Context, request *DescribeVirusTaskListRequest) (response *DescribeVirusTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeVirusTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirusTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirusTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarningRulesRequest() (request *DescribeWarningRulesRequest) {
    request = &DescribeWarningRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "DescribeWarningRules")
    
    
    return
}

func NewDescribeWarningRulesResponse() (response *DescribeWarningRulesResponse) {
    response = &DescribeWarningRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWarningRules
// 获取告警策略列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWarningRules(request *DescribeWarningRulesRequest) (response *DescribeWarningRulesResponse, err error) {
    return c.DescribeWarningRulesWithContext(context.Background(), request)
}

// DescribeWarningRules
// 获取告警策略列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWarningRulesWithContext(ctx context.Context, request *DescribeWarningRulesRequest) (response *DescribeWarningRulesResponse, err error) {
    if request == nil {
        request = NewDescribeWarningRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWarningRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWarningRulesResponse()
    err = c.Send(request, response)
    return
}

func NewExportVirusListRequest() (request *ExportVirusListRequest) {
    request = &ExportVirusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ExportVirusList")
    
    
    return
}

func NewExportVirusListResponse() (response *ExportVirusListResponse) {
    response = &ExportVirusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVirusList
// 运行时文件查杀事件列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportVirusList(request *ExportVirusListRequest) (response *ExportVirusListResponse, err error) {
    return c.ExportVirusListWithContext(context.Background(), request)
}

// ExportVirusList
// 运行时文件查杀事件列表导出
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportVirusListWithContext(ctx context.Context, request *ExportVirusListRequest) (response *ExportVirusListResponse, err error) {
    if request == nil {
        request = NewExportVirusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVirusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVirusListResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeUserComplianceEnvironmentRequest() (request *InitializeUserComplianceEnvironmentRequest) {
    request = &InitializeUserComplianceEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "InitializeUserComplianceEnvironment")
    
    
    return
}

func NewInitializeUserComplianceEnvironmentResponse() (response *InitializeUserComplianceEnvironmentResponse) {
    response = &InitializeUserComplianceEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitializeUserComplianceEnvironment
// 为客户初始化合规基线的使用环境，创建必要的数据和选项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InitializeUserComplianceEnvironment(request *InitializeUserComplianceEnvironmentRequest) (response *InitializeUserComplianceEnvironmentResponse, err error) {
    return c.InitializeUserComplianceEnvironmentWithContext(context.Background(), request)
}

// InitializeUserComplianceEnvironment
// 为客户初始化合规基线的使用环境，创建必要的数据和选项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InitializeUserComplianceEnvironmentWithContext(ctx context.Context, request *InitializeUserComplianceEnvironmentRequest) (response *InitializeUserComplianceEnvironmentResponse, err error) {
    if request == nil {
        request = NewInitializeUserComplianceEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitializeUserComplianceEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitializeUserComplianceEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAbnormalProcessRuleStatusRequest() (request *ModifyAbnormalProcessRuleStatusRequest) {
    request = &ModifyAbnormalProcessRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessRuleStatus")
    
    
    return
}

func NewModifyAbnormalProcessRuleStatusResponse() (response *ModifyAbnormalProcessRuleStatusResponse) {
    response = &ModifyAbnormalProcessRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAbnormalProcessRuleStatus
// 修改运行时异常进程策略的开启关闭状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessRuleStatus(request *ModifyAbnormalProcessRuleStatusRequest) (response *ModifyAbnormalProcessRuleStatusResponse, err error) {
    return c.ModifyAbnormalProcessRuleStatusWithContext(context.Background(), request)
}

// ModifyAbnormalProcessRuleStatus
// 修改运行时异常进程策略的开启关闭状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessRuleStatusWithContext(ctx context.Context, request *ModifyAbnormalProcessRuleStatusRequest) (response *ModifyAbnormalProcessRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAbnormalProcessRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAbnormalProcessRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAbnormalProcessRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAbnormalProcessStatusRequest() (request *ModifyAbnormalProcessStatusRequest) {
    request = &ModifyAbnormalProcessStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAbnormalProcessStatus")
    
    
    return
}

func NewModifyAbnormalProcessStatusResponse() (response *ModifyAbnormalProcessStatusResponse) {
    response = &ModifyAbnormalProcessStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAbnormalProcessStatus
// 修改异常进程事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessStatus(request *ModifyAbnormalProcessStatusRequest) (response *ModifyAbnormalProcessStatusResponse, err error) {
    return c.ModifyAbnormalProcessStatusWithContext(context.Background(), request)
}

// ModifyAbnormalProcessStatus
// 修改异常进程事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAbnormalProcessStatusWithContext(ctx context.Context, request *ModifyAbnormalProcessStatusRequest) (response *ModifyAbnormalProcessStatusResponse, err error) {
    if request == nil {
        request = NewModifyAbnormalProcessStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAbnormalProcessStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAbnormalProcessStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessControlRuleStatusRequest() (request *ModifyAccessControlRuleStatusRequest) {
    request = &ModifyAccessControlRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlRuleStatus")
    
    
    return
}

func NewModifyAccessControlRuleStatusResponse() (response *ModifyAccessControlRuleStatusResponse) {
    response = &ModifyAccessControlRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessControlRuleStatus
// 修改运行时访问控制策略的状态，启用或者禁用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlRuleStatus(request *ModifyAccessControlRuleStatusRequest) (response *ModifyAccessControlRuleStatusResponse, err error) {
    return c.ModifyAccessControlRuleStatusWithContext(context.Background(), request)
}

// ModifyAccessControlRuleStatus
// 修改运行时访问控制策略的状态，启用或者禁用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlRuleStatusWithContext(ctx context.Context, request *ModifyAccessControlRuleStatusRequest) (response *ModifyAccessControlRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessControlRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessControlRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessControlRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessControlStatusRequest() (request *ModifyAccessControlStatusRequest) {
    request = &ModifyAccessControlStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAccessControlStatus")
    
    
    return
}

func NewModifyAccessControlStatusResponse() (response *ModifyAccessControlStatusResponse) {
    response = &ModifyAccessControlStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessControlStatus
// 修改运行时访问控制事件状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlStatus(request *ModifyAccessControlStatusRequest) (response *ModifyAccessControlStatusResponse, err error) {
    return c.ModifyAccessControlStatusWithContext(context.Background(), request)
}

// ModifyAccessControlStatus
// 修改运行时访问控制事件状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessControlStatusWithContext(ctx context.Context, request *ModifyAccessControlStatusRequest) (response *ModifyAccessControlStatusResponse, err error) {
    if request == nil {
        request = NewModifyAccessControlStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessControlStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessControlStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetRequest() (request *ModifyAssetRequest) {
    request = &ModifyAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAsset")
    
    
    return
}

func NewModifyAssetResponse() (response *ModifyAssetResponse) {
    response = &ModifyAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAsset
// 容器安全主机资产刷新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAsset(request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
    return c.ModifyAssetWithContext(context.Background(), request)
}

// ModifyAsset
// 容器安全主机资产刷新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetWithContext(ctx context.Context, request *ModifyAssetRequest) (response *ModifyAssetResponse, err error) {
    if request == nil {
        request = NewModifyAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageRegistryScanStopRequest() (request *ModifyAssetImageRegistryScanStopRequest) {
    request = &ModifyAssetImageRegistryScanStopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStop")
    
    
    return
}

func NewModifyAssetImageRegistryScanStopResponse() (response *ModifyAssetImageRegistryScanStopResponse) {
    response = &ModifyAssetImageRegistryScanStopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAssetImageRegistryScanStop
// 镜像仓库停止镜像扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStop(request *ModifyAssetImageRegistryScanStopRequest) (response *ModifyAssetImageRegistryScanStopResponse, err error) {
    return c.ModifyAssetImageRegistryScanStopWithContext(context.Background(), request)
}

// ModifyAssetImageRegistryScanStop
// 镜像仓库停止镜像扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopWithContext(ctx context.Context, request *ModifyAssetImageRegistryScanStopRequest) (response *ModifyAssetImageRegistryScanStopResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageRegistryScanStopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageRegistryScanStop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageRegistryScanStopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageRegistryScanStopOneKeyRequest() (request *ModifyAssetImageRegistryScanStopOneKeyRequest) {
    request = &ModifyAssetImageRegistryScanStopOneKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageRegistryScanStopOneKey")
    
    
    return
}

func NewModifyAssetImageRegistryScanStopOneKeyResponse() (response *ModifyAssetImageRegistryScanStopOneKeyResponse) {
    response = &ModifyAssetImageRegistryScanStopOneKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAssetImageRegistryScanStopOneKey
// 镜像仓库停止镜像一键扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopOneKey(request *ModifyAssetImageRegistryScanStopOneKeyRequest) (response *ModifyAssetImageRegistryScanStopOneKeyResponse, err error) {
    return c.ModifyAssetImageRegistryScanStopOneKeyWithContext(context.Background(), request)
}

// ModifyAssetImageRegistryScanStopOneKey
// 镜像仓库停止镜像一键扫描任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageRegistryScanStopOneKeyWithContext(ctx context.Context, request *ModifyAssetImageRegistryScanStopOneKeyRequest) (response *ModifyAssetImageRegistryScanStopOneKeyResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageRegistryScanStopOneKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageRegistryScanStopOneKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageRegistryScanStopOneKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetImageScanStopRequest() (request *ModifyAssetImageScanStopRequest) {
    request = &ModifyAssetImageScanStopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyAssetImageScanStop")
    
    
    return
}

func NewModifyAssetImageScanStopResponse() (response *ModifyAssetImageScanStopResponse) {
    response = &ModifyAssetImageScanStopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAssetImageScanStop
// 容器安全停止镜像扫描
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageScanStop(request *ModifyAssetImageScanStopRequest) (response *ModifyAssetImageScanStopResponse, err error) {
    return c.ModifyAssetImageScanStopWithContext(context.Background(), request)
}

// ModifyAssetImageScanStop
// 容器安全停止镜像扫描
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAssetImageScanStopWithContext(ctx context.Context, request *ModifyAssetImageScanStopRequest) (response *ModifyAssetImageScanStopResponse, err error) {
    if request == nil {
        request = NewModifyAssetImageScanStopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetImageScanStop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetImageScanStopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompliancePeriodTaskRequest() (request *ModifyCompliancePeriodTaskRequest) {
    request = &ModifyCompliancePeriodTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyCompliancePeriodTask")
    
    
    return
}

func NewModifyCompliancePeriodTaskResponse() (response *ModifyCompliancePeriodTaskResponse) {
    response = &ModifyCompliancePeriodTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCompliancePeriodTask
// 修改定时任务的设置，包括检测周期、开启/禁用合规基准。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyCompliancePeriodTask(request *ModifyCompliancePeriodTaskRequest) (response *ModifyCompliancePeriodTaskResponse, err error) {
    return c.ModifyCompliancePeriodTaskWithContext(context.Background(), request)
}

// ModifyCompliancePeriodTask
// 修改定时任务的设置，包括检测周期、开启/禁用合规基准。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ModifyCompliancePeriodTaskWithContext(ctx context.Context, request *ModifyCompliancePeriodTaskRequest) (response *ModifyCompliancePeriodTaskResponse, err error) {
    if request == nil {
        request = NewModifyCompliancePeriodTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompliancePeriodTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompliancePeriodTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEscapeEventStatusRequest() (request *ModifyEscapeEventStatusRequest) {
    request = &ModifyEscapeEventStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeEventStatus")
    
    
    return
}

func NewModifyEscapeEventStatusResponse() (response *ModifyEscapeEventStatusResponse) {
    response = &ModifyEscapeEventStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEscapeEventStatus
// ModifyEscapeEventStatus  修改容器逃逸扫描事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) ModifyEscapeEventStatus(request *ModifyEscapeEventStatusRequest) (response *ModifyEscapeEventStatusResponse, err error) {
    return c.ModifyEscapeEventStatusWithContext(context.Background(), request)
}

// ModifyEscapeEventStatus
// ModifyEscapeEventStatus  修改容器逃逸扫描事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) ModifyEscapeEventStatusWithContext(ctx context.Context, request *ModifyEscapeEventStatusRequest) (response *ModifyEscapeEventStatusResponse, err error) {
    if request == nil {
        request = NewModifyEscapeEventStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEscapeEventStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEscapeEventStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEscapeRuleRequest() (request *ModifyEscapeRuleRequest) {
    request = &ModifyEscapeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyEscapeRule")
    
    
    return
}

func NewModifyEscapeRuleResponse() (response *ModifyEscapeRuleResponse) {
    response = &ModifyEscapeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEscapeRule
// ModifyEscapeRule  修改容器逃逸扫描规则信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyEscapeRule(request *ModifyEscapeRuleRequest) (response *ModifyEscapeRuleResponse, err error) {
    return c.ModifyEscapeRuleWithContext(context.Background(), request)
}

// ModifyEscapeRule
// ModifyEscapeRule  修改容器逃逸扫描规则信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyEscapeRuleWithContext(ctx context.Context, request *ModifyEscapeRuleRequest) (response *ModifyEscapeRuleResponse, err error) {
    if request == nil {
        request = NewModifyEscapeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEscapeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEscapeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReverseShellStatusRequest() (request *ModifyReverseShellStatusRequest) {
    request = &ModifyReverseShellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyReverseShellStatus")
    
    
    return
}

func NewModifyReverseShellStatusResponse() (response *ModifyReverseShellStatusResponse) {
    response = &ModifyReverseShellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyReverseShellStatus
// 修改反弹shell事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyReverseShellStatus(request *ModifyReverseShellStatusRequest) (response *ModifyReverseShellStatusResponse, err error) {
    return c.ModifyReverseShellStatusWithContext(context.Background(), request)
}

// ModifyReverseShellStatus
// 修改反弹shell事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyReverseShellStatusWithContext(ctx context.Context, request *ModifyReverseShellStatusRequest) (response *ModifyReverseShellStatusResponse, err error) {
    if request == nil {
        request = NewModifyReverseShellStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReverseShellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReverseShellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskSyscallStatusRequest() (request *ModifyRiskSyscallStatusRequest) {
    request = &ModifyRiskSyscallStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyRiskSyscallStatus")
    
    
    return
}

func NewModifyRiskSyscallStatusResponse() (response *ModifyRiskSyscallStatusResponse) {
    response = &ModifyRiskSyscallStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRiskSyscallStatus
// 修改高危系统调用事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyRiskSyscallStatus(request *ModifyRiskSyscallStatusRequest) (response *ModifyRiskSyscallStatusResponse, err error) {
    return c.ModifyRiskSyscallStatusWithContext(context.Background(), request)
}

// ModifyRiskSyscallStatus
// 修改高危系统调用事件的状态信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyRiskSyscallStatusWithContext(ctx context.Context, request *ModifyRiskSyscallStatusRequest) (response *ModifyRiskSyscallStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskSyscallStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskSyscallStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskSyscallStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusFileStatusRequest() (request *ModifyVirusFileStatusRequest) {
    request = &ModifyVirusFileStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusFileStatus")
    
    
    return
}

func NewModifyVirusFileStatusResponse() (response *ModifyVirusFileStatusResponse) {
    response = &ModifyVirusFileStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVirusFileStatus
// 运行时更新木马文件事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyVirusFileStatus(request *ModifyVirusFileStatusRequest) (response *ModifyVirusFileStatusResponse, err error) {
    return c.ModifyVirusFileStatusWithContext(context.Background(), request)
}

// ModifyVirusFileStatus
// 运行时更新木马文件事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"
func (c *Client) ModifyVirusFileStatusWithContext(ctx context.Context, request *ModifyVirusFileStatusRequest) (response *ModifyVirusFileStatusResponse, err error) {
    if request == nil {
        request = NewModifyVirusFileStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusFileStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusFileStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusMonitorSettingRequest() (request *ModifyVirusMonitorSettingRequest) {
    request = &ModifyVirusMonitorSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusMonitorSetting")
    
    
    return
}

func NewModifyVirusMonitorSettingResponse() (response *ModifyVirusMonitorSettingResponse) {
    response = &ModifyVirusMonitorSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVirusMonitorSetting
// 运行时更新文件查杀实时监控设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusMonitorSetting(request *ModifyVirusMonitorSettingRequest) (response *ModifyVirusMonitorSettingResponse, err error) {
    return c.ModifyVirusMonitorSettingWithContext(context.Background(), request)
}

// ModifyVirusMonitorSetting
// 运行时更新文件查杀实时监控设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusMonitorSettingWithContext(ctx context.Context, request *ModifyVirusMonitorSettingRequest) (response *ModifyVirusMonitorSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusMonitorSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusMonitorSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusMonitorSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusScanSettingRequest() (request *ModifyVirusScanSettingRequest) {
    request = &ModifyVirusScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanSetting")
    
    
    return
}

func NewModifyVirusScanSettingResponse() (response *ModifyVirusScanSettingResponse) {
    response = &ModifyVirusScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVirusScanSetting
// 运行时更新文件查杀设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanSetting(request *ModifyVirusScanSettingRequest) (response *ModifyVirusScanSettingResponse, err error) {
    return c.ModifyVirusScanSettingWithContext(context.Background(), request)
}

// ModifyVirusScanSetting
// 运行时更新文件查杀设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanSettingWithContext(ctx context.Context, request *ModifyVirusScanSettingRequest) (response *ModifyVirusScanSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirusScanTimeoutSettingRequest() (request *ModifyVirusScanTimeoutSettingRequest) {
    request = &ModifyVirusScanTimeoutSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ModifyVirusScanTimeoutSetting")
    
    
    return
}

func NewModifyVirusScanTimeoutSettingResponse() (response *ModifyVirusScanTimeoutSettingResponse) {
    response = &ModifyVirusScanTimeoutSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVirusScanTimeoutSetting
// 运行时文件扫描超时设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanTimeoutSetting(request *ModifyVirusScanTimeoutSettingRequest) (response *ModifyVirusScanTimeoutSettingResponse, err error) {
    return c.ModifyVirusScanTimeoutSettingWithContext(context.Background(), request)
}

// ModifyVirusScanTimeoutSetting
// 运行时文件扫描超时设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVirusScanTimeoutSettingWithContext(ctx context.Context, request *ModifyVirusScanTimeoutSettingRequest) (response *ModifyVirusScanTimeoutSettingResponse, err error) {
    if request == nil {
        request = NewModifyVirusScanTimeoutSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirusScanTimeoutSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirusScanTimeoutSettingResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAssetImageRegistryRegistryDetailRequest() (request *RemoveAssetImageRegistryRegistryDetailRequest) {
    request = &RemoveAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "RemoveAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewRemoveAssetImageRegistryRegistryDetailResponse() (response *RemoveAssetImageRegistryRegistryDetailResponse) {
    response = &RemoveAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveAssetImageRegistryRegistryDetail
// 删除单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveAssetImageRegistryRegistryDetail(request *RemoveAssetImageRegistryRegistryDetailRequest) (response *RemoveAssetImageRegistryRegistryDetailResponse, err error) {
    return c.RemoveAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// RemoveAssetImageRegistryRegistryDetail
// 删除单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RemoveAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *RemoveAssetImageRegistryRegistryDetailRequest) (response *RemoveAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewRemoveAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewRenewImageAuthorizeStateRequest() (request *RenewImageAuthorizeStateRequest) {
    request = &RenewImageAuthorizeStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "RenewImageAuthorizeState")
    
    
    return
}

func NewRenewImageAuthorizeStateResponse() (response *RenewImageAuthorizeStateResponse) {
    response = &RenewImageAuthorizeStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewImageAuthorizeState
// RenewImageAuthorizeState   授权镜像扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RenewImageAuthorizeState(request *RenewImageAuthorizeStateRequest) (response *RenewImageAuthorizeStateResponse, err error) {
    return c.RenewImageAuthorizeStateWithContext(context.Background(), request)
}

// RenewImageAuthorizeState
// RenewImageAuthorizeState   授权镜像扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RenewImageAuthorizeStateWithContext(ctx context.Context, request *RenewImageAuthorizeStateRequest) (response *RenewImageAuthorizeStateResponse, err error) {
    if request == nil {
        request = NewRenewImageAuthorizeStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewImageAuthorizeState require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewImageAuthorizeStateResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceAssetsRequest() (request *ScanComplianceAssetsRequest) {
    request = &ScanComplianceAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssets")
    
    
    return
}

func NewScanComplianceAssetsResponse() (response *ScanComplianceAssetsResponse) {
    response = &ScanComplianceAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanComplianceAssets
// 重新检测选定的资产
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssets(request *ScanComplianceAssetsRequest) (response *ScanComplianceAssetsResponse, err error) {
    return c.ScanComplianceAssetsWithContext(context.Background(), request)
}

// ScanComplianceAssets
// 重新检测选定的资产
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsWithContext(ctx context.Context, request *ScanComplianceAssetsRequest) (response *ScanComplianceAssetsResponse, err error) {
    if request == nil {
        request = NewScanComplianceAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceAssetsByPolicyItemRequest() (request *ScanComplianceAssetsByPolicyItemRequest) {
    request = &ScanComplianceAssetsByPolicyItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceAssetsByPolicyItem")
    
    
    return
}

func NewScanComplianceAssetsByPolicyItemResponse() (response *ScanComplianceAssetsByPolicyItemResponse) {
    response = &ScanComplianceAssetsByPolicyItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanComplianceAssetsByPolicyItem
// 用指定的检测项重新检测选定的资产，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsByPolicyItem(request *ScanComplianceAssetsByPolicyItemRequest) (response *ScanComplianceAssetsByPolicyItemResponse, err error) {
    return c.ScanComplianceAssetsByPolicyItemWithContext(context.Background(), request)
}

// ScanComplianceAssetsByPolicyItem
// 用指定的检测项重新检测选定的资产，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceAssetsByPolicyItemWithContext(ctx context.Context, request *ScanComplianceAssetsByPolicyItemRequest) (response *ScanComplianceAssetsByPolicyItemResponse, err error) {
    if request == nil {
        request = NewScanComplianceAssetsByPolicyItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceAssetsByPolicyItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceAssetsByPolicyItemResponse()
    err = c.Send(request, response)
    return
}

func NewScanCompliancePolicyItemsRequest() (request *ScanCompliancePolicyItemsRequest) {
    request = &ScanCompliancePolicyItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ScanCompliancePolicyItems")
    
    
    return
}

func NewScanCompliancePolicyItemsResponse() (response *ScanCompliancePolicyItemsResponse) {
    response = &ScanCompliancePolicyItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanCompliancePolicyItems
// 重新检测选的检测项下的所有资产，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanCompliancePolicyItems(request *ScanCompliancePolicyItemsRequest) (response *ScanCompliancePolicyItemsResponse, err error) {
    return c.ScanCompliancePolicyItemsWithContext(context.Background(), request)
}

// ScanCompliancePolicyItems
// 重新检测选的检测项下的所有资产，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanCompliancePolicyItemsWithContext(ctx context.Context, request *ScanCompliancePolicyItemsRequest) (response *ScanCompliancePolicyItemsResponse, err error) {
    if request == nil {
        request = NewScanCompliancePolicyItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanCompliancePolicyItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanCompliancePolicyItemsResponse()
    err = c.Send(request, response)
    return
}

func NewScanComplianceScanFailedAssetsRequest() (request *ScanComplianceScanFailedAssetsRequest) {
    request = &ScanComplianceScanFailedAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "ScanComplianceScanFailedAssets")
    
    
    return
}

func NewScanComplianceScanFailedAssetsResponse() (response *ScanComplianceScanFailedAssetsResponse) {
    response = &ScanComplianceScanFailedAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanComplianceScanFailedAssets
// 重新检测选定的检测失败的资产下的所有失败的检测项，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceScanFailedAssets(request *ScanComplianceScanFailedAssetsRequest) (response *ScanComplianceScanFailedAssetsResponse, err error) {
    return c.ScanComplianceScanFailedAssetsWithContext(context.Background(), request)
}

// ScanComplianceScanFailedAssets
// 重新检测选定的检测失败的资产下的所有失败的检测项，返回创建的合规检查任务的ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ScanComplianceScanFailedAssetsWithContext(ctx context.Context, request *ScanComplianceScanFailedAssetsRequest) (response *ScanComplianceScanFailedAssetsResponse, err error) {
    if request == nil {
        request = NewScanComplianceScanFailedAssetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanComplianceScanFailedAssets require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanComplianceScanFailedAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewSetCheckModeRequest() (request *SetCheckModeRequest) {
    request = &SetCheckModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "SetCheckMode")
    
    
    return
}

func NewSetCheckModeResponse() (response *SetCheckModeResponse) {
    response = &SetCheckModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetCheckMode
// 设置检测模式和自动检查
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCheckMode(request *SetCheckModeRequest) (response *SetCheckModeResponse, err error) {
    return c.SetCheckModeWithContext(context.Background(), request)
}

// SetCheckMode
// 设置检测模式和自动检查
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCheckModeWithContext(ctx context.Context, request *SetCheckModeRequest) (response *SetCheckModeResponse, err error) {
    if request == nil {
        request = NewSetCheckModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCheckMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCheckModeResponse()
    err = c.Send(request, response)
    return
}

func NewStopVirusScanTaskRequest() (request *StopVirusScanTaskRequest) {
    request = &StopVirusScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "StopVirusScanTask")
    
    
    return
}

func NewStopVirusScanTaskResponse() (response *StopVirusScanTaskResponse) {
    response = &StopVirusScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopVirusScanTask
// 运行时停止木马查杀任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopVirusScanTask(request *StopVirusScanTaskRequest) (response *StopVirusScanTaskResponse, err error) {
    return c.StopVirusScanTaskWithContext(context.Background(), request)
}

// StopVirusScanTask
// 运行时停止木马查杀任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopVirusScanTaskWithContext(ctx context.Context, request *StopVirusScanTaskRequest) (response *StopVirusScanTaskResponse, err error) {
    if request == nil {
        request = NewStopVirusScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopVirusScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopVirusScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSyncAssetImageRegistryAssetRequest() (request *SyncAssetImageRegistryAssetRequest) {
    request = &SyncAssetImageRegistryAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "SyncAssetImageRegistryAsset")
    
    
    return
}

func NewSyncAssetImageRegistryAssetResponse() (response *SyncAssetImageRegistryAssetResponse) {
    response = &SyncAssetImageRegistryAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncAssetImageRegistryAsset
// 镜像仓库资产刷新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetImageRegistryAsset(request *SyncAssetImageRegistryAssetRequest) (response *SyncAssetImageRegistryAssetResponse, err error) {
    return c.SyncAssetImageRegistryAssetWithContext(context.Background(), request)
}

// SyncAssetImageRegistryAsset
// 镜像仓库资产刷新
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetImageRegistryAssetWithContext(ctx context.Context, request *SyncAssetImageRegistryAssetRequest) (response *SyncAssetImageRegistryAssetResponse, err error) {
    if request == nil {
        request = NewSyncAssetImageRegistryAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncAssetImageRegistryAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncAssetImageRegistryAssetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssetImageRegistryRegistryDetailRequest() (request *UpdateAssetImageRegistryRegistryDetailRequest) {
    request = &UpdateAssetImageRegistryRegistryDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateAssetImageRegistryRegistryDetail")
    
    
    return
}

func NewUpdateAssetImageRegistryRegistryDetailResponse() (response *UpdateAssetImageRegistryRegistryDetailResponse) {
    response = &UpdateAssetImageRegistryRegistryDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAssetImageRegistryRegistryDetail
// 更新单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAssetImageRegistryRegistryDetail(request *UpdateAssetImageRegistryRegistryDetailRequest) (response *UpdateAssetImageRegistryRegistryDetailResponse, err error) {
    return c.UpdateAssetImageRegistryRegistryDetailWithContext(context.Background(), request)
}

// UpdateAssetImageRegistryRegistryDetail
// 更新单个镜像仓库详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateAssetImageRegistryRegistryDetailWithContext(ctx context.Context, request *UpdateAssetImageRegistryRegistryDetailRequest) (response *UpdateAssetImageRegistryRegistryDetailResponse, err error) {
    if request == nil {
        request = NewUpdateAssetImageRegistryRegistryDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAssetImageRegistryRegistryDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAssetImageRegistryRegistryDetailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageRegistryTimingScanTaskRequest() (request *UpdateImageRegistryTimingScanTaskRequest) {
    request = &UpdateImageRegistryTimingScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcss", APIVersion, "UpdateImageRegistryTimingScanTask")
    
    
    return
}

func NewUpdateImageRegistryTimingScanTaskResponse() (response *UpdateImageRegistryTimingScanTaskResponse) {
    response = &UpdateImageRegistryTimingScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateImageRegistryTimingScanTask
// 镜像仓库更新定时任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateImageRegistryTimingScanTask(request *UpdateImageRegistryTimingScanTaskRequest) (response *UpdateImageRegistryTimingScanTaskResponse, err error) {
    return c.UpdateImageRegistryTimingScanTaskWithContext(context.Background(), request)
}

// UpdateImageRegistryTimingScanTask
// 镜像仓库更新定时任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateImageRegistryTimingScanTaskWithContext(ctx context.Context, request *UpdateImageRegistryTimingScanTaskRequest) (response *UpdateImageRegistryTimingScanTaskResponse, err error) {
    if request == nil {
        request = NewUpdateImageRegistryTimingScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateImageRegistryTimingScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateImageRegistryTimingScanTaskResponse()
    err = c.Send(request, response)
    return
}
