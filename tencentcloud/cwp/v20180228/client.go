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

package v20180228

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

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


func NewAddLoginWhiteListsRequest() (request *AddLoginWhiteListsRequest) {
    request = &AddLoginWhiteListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "AddLoginWhiteLists")
    
    
    return
}

func NewAddLoginWhiteListsResponse() (response *AddLoginWhiteListsResponse) {
    response = &AddLoginWhiteListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLoginWhiteLists
// 批量添加异地登录白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddLoginWhiteLists(request *AddLoginWhiteListsRequest) (response *AddLoginWhiteListsResponse, err error) {
    return c.AddLoginWhiteListsWithContext(context.Background(), request)
}

// AddLoginWhiteLists
// 批量添加异地登录白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddLoginWhiteListsWithContext(ctx context.Context, request *AddLoginWhiteListsRequest) (response *AddLoginWhiteListsResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLoginWhiteLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLoginWhiteListsResponse()
    err = c.Send(request, response)
    return
}

func NewCancelIgnoreVulRequest() (request *CancelIgnoreVulRequest) {
    request = &CancelIgnoreVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CancelIgnoreVul")
    
    
    return
}

func NewCancelIgnoreVulResponse() (response *CancelIgnoreVulResponse) {
    response = &CancelIgnoreVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelIgnoreVul
// 取消漏洞忽略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CancelIgnoreVul(request *CancelIgnoreVulRequest) (response *CancelIgnoreVulResponse, err error) {
    return c.CancelIgnoreVulWithContext(context.Background(), request)
}

// CancelIgnoreVul
// 取消漏洞忽略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CancelIgnoreVulWithContext(ctx context.Context, request *CancelIgnoreVulRequest) (response *CancelIgnoreVulResponse, err error) {
    if request == nil {
        request = NewCancelIgnoreVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelIgnoreVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelIgnoreVulResponse()
    err = c.Send(request, response)
    return
}

func NewChangeRuleEventsIgnoreStatusRequest() (request *ChangeRuleEventsIgnoreStatusRequest) {
    request = &ChangeRuleEventsIgnoreStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ChangeRuleEventsIgnoreStatus")
    
    
    return
}

func NewChangeRuleEventsIgnoreStatusResponse() (response *ChangeRuleEventsIgnoreStatusResponse) {
    response = &ChangeRuleEventsIgnoreStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeRuleEventsIgnoreStatus
// 根据检测项id或事件id批量忽略事件或取消忽略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ChangeRuleEventsIgnoreStatus(request *ChangeRuleEventsIgnoreStatusRequest) (response *ChangeRuleEventsIgnoreStatusResponse, err error) {
    return c.ChangeRuleEventsIgnoreStatusWithContext(context.Background(), request)
}

// ChangeRuleEventsIgnoreStatus
// 根据检测项id或事件id批量忽略事件或取消忽略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ChangeRuleEventsIgnoreStatusWithContext(ctx context.Context, request *ChangeRuleEventsIgnoreStatusRequest) (response *ChangeRuleEventsIgnoreStatusResponse, err error) {
    if request == nil {
        request = NewChangeRuleEventsIgnoreStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeRuleEventsIgnoreStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeRuleEventsIgnoreStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChangeStrategyEnableStatusRequest() (request *ChangeStrategyEnableStatusRequest) {
    request = &ChangeStrategyEnableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ChangeStrategyEnableStatus")
    
    
    return
}

func NewChangeStrategyEnableStatusResponse() (response *ChangeStrategyEnableStatusResponse) {
    response = &ChangeStrategyEnableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeStrategyEnableStatus
// 根据策略id修改策略可用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ChangeStrategyEnableStatus(request *ChangeStrategyEnableStatusRequest) (response *ChangeStrategyEnableStatusResponse, err error) {
    return c.ChangeStrategyEnableStatusWithContext(context.Background(), request)
}

// ChangeStrategyEnableStatus
// 根据策略id修改策略可用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ChangeStrategyEnableStatusWithContext(ctx context.Context, request *ChangeStrategyEnableStatusRequest) (response *ChangeStrategyEnableStatusResponse, err error) {
    if request == nil {
        request = NewChangeStrategyEnableStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeStrategyEnableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeStrategyEnableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBashPolicyParamsRequest() (request *CheckBashPolicyParamsRequest) {
    request = &CheckBashPolicyParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CheckBashPolicyParams")
    
    
    return
}

func NewCheckBashPolicyParamsResponse() (response *CheckBashPolicyParamsResponse) {
    response = &CheckBashPolicyParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBashPolicyParams
// 校验高危命令用户规则新增和编辑时的参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckBashPolicyParams(request *CheckBashPolicyParamsRequest) (response *CheckBashPolicyParamsResponse, err error) {
    return c.CheckBashPolicyParamsWithContext(context.Background(), request)
}

// CheckBashPolicyParams
// 校验高危命令用户规则新增和编辑时的参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckBashPolicyParamsWithContext(ctx context.Context, request *CheckBashPolicyParamsRequest) (response *CheckBashPolicyParamsResponse, err error) {
    if request == nil {
        request = NewCheckBashPolicyParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckBashPolicyParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckBashPolicyParamsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBashRuleParamsRequest() (request *CheckBashRuleParamsRequest) {
    request = &CheckBashRuleParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CheckBashRuleParams")
    
    
    return
}

func NewCheckBashRuleParamsResponse() (response *CheckBashRuleParamsResponse) {
    response = &CheckBashRuleParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBashRuleParams
// 校验高危命令用户规则新增和编辑时的参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckBashRuleParams(request *CheckBashRuleParamsRequest) (response *CheckBashRuleParamsResponse, err error) {
    return c.CheckBashRuleParamsWithContext(context.Background(), request)
}

// CheckBashRuleParams
// 校验高危命令用户规则新增和编辑时的参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckBashRuleParamsWithContext(ctx context.Context, request *CheckBashRuleParamsRequest) (response *CheckBashRuleParamsResponse, err error) {
    if request == nil {
        request = NewCheckBashRuleParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckBashRuleParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckBashRuleParamsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFileTamperRuleRequest() (request *CheckFileTamperRuleRequest) {
    request = &CheckFileTamperRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CheckFileTamperRule")
    
    
    return
}

func NewCheckFileTamperRuleResponse() (response *CheckFileTamperRuleResponse) {
    response = &CheckFileTamperRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckFileTamperRule
// 检验核心文件监控前端新增和编辑时的规则参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckFileTamperRule(request *CheckFileTamperRuleRequest) (response *CheckFileTamperRuleResponse, err error) {
    return c.CheckFileTamperRuleWithContext(context.Background(), request)
}

// CheckFileTamperRule
// 检验核心文件监控前端新增和编辑时的规则参数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CheckFileTamperRuleWithContext(ctx context.Context, request *CheckFileTamperRuleRequest) (response *CheckFileTamperRuleResponse, err error) {
    if request == nil {
        request = NewCheckFileTamperRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFileTamperRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFileTamperRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFirstScanBaselineRequest() (request *CheckFirstScanBaselineRequest) {
    request = &CheckFirstScanBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CheckFirstScanBaseline")
    
    
    return
}

func NewCheckFirstScanBaselineResponse() (response *CheckFirstScanBaselineResponse) {
    response = &CheckFirstScanBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckFirstScanBaseline
// 查询基线是否第一次检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CheckFirstScanBaseline(request *CheckFirstScanBaselineRequest) (response *CheckFirstScanBaselineResponse, err error) {
    return c.CheckFirstScanBaselineWithContext(context.Background(), request)
}

// CheckFirstScanBaseline
// 查询基线是否第一次检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CheckFirstScanBaselineWithContext(ctx context.Context, request *CheckFirstScanBaselineRequest) (response *CheckFirstScanBaselineResponse, err error) {
    if request == nil {
        request = NewCheckFirstScanBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFirstScanBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFirstScanBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLogKafkaConnectionStateRequest() (request *CheckLogKafkaConnectionStateRequest) {
    request = &CheckLogKafkaConnectionStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CheckLogKafkaConnectionState")
    
    
    return
}

func NewCheckLogKafkaConnectionStateResponse() (response *CheckLogKafkaConnectionStateResponse) {
    response = &CheckLogKafkaConnectionStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckLogKafkaConnectionState
// 检查日志投递kafka连通性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckLogKafkaConnectionState(request *CheckLogKafkaConnectionStateRequest) (response *CheckLogKafkaConnectionStateResponse, err error) {
    return c.CheckLogKafkaConnectionStateWithContext(context.Background(), request)
}

// CheckLogKafkaConnectionState
// 检查日志投递kafka连通性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckLogKafkaConnectionStateWithContext(ctx context.Context, request *CheckLogKafkaConnectionStateRequest) (response *CheckLogKafkaConnectionStateResponse, err error) {
    if request == nil {
        request = NewCheckLogKafkaConnectionStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckLogKafkaConnectionState require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckLogKafkaConnectionStateResponse()
    err = c.Send(request, response)
    return
}

func NewClearLocalStorageRequest() (request *ClearLocalStorageRequest) {
    request = &ClearLocalStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ClearLocalStorage")
    
    
    return
}

func NewClearLocalStorageResponse() (response *ClearLocalStorageResponse) {
    response = &ClearLocalStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearLocalStorage
// 清理本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ClearLocalStorage(request *ClearLocalStorageRequest) (response *ClearLocalStorageResponse, err error) {
    return c.ClearLocalStorageWithContext(context.Background(), request)
}

// ClearLocalStorage
// 清理本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ClearLocalStorageWithContext(ctx context.Context, request *ClearLocalStorageRequest) (response *ClearLocalStorageResponse, err error) {
    if request == nil {
        request = NewClearLocalStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearLocalStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearLocalStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBanWhiteListRequest() (request *CreateBanWhiteListRequest) {
    request = &CreateBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateBanWhiteList")
    
    
    return
}

func NewCreateBanWhiteListResponse() (response *CreateBanWhiteListResponse) {
    response = &CreateBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBanWhiteList
// 添加阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBanWhiteList(request *CreateBanWhiteListRequest) (response *CreateBanWhiteListResponse, err error) {
    return c.CreateBanWhiteListWithContext(context.Background(), request)
}

// CreateBanWhiteList
// 添加阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBanWhiteListWithContext(ctx context.Context, request *CreateBanWhiteListRequest) (response *CreateBanWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateBanWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBanWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBaselineStrategyRequest() (request *CreateBaselineStrategyRequest) {
    request = &CreateBaselineStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateBaselineStrategy")
    
    
    return
}

func NewCreateBaselineStrategyResponse() (response *CreateBaselineStrategyResponse) {
    response = &CreateBaselineStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBaselineStrategy
// 根据策略信息创建基线策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  FAILEDOPERATION_TOOMANYSTRATEGY = "FailedOperation.TooManyStrategy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBaselineStrategy(request *CreateBaselineStrategyRequest) (response *CreateBaselineStrategyResponse, err error) {
    return c.CreateBaselineStrategyWithContext(context.Background(), request)
}

// CreateBaselineStrategy
// 根据策略信息创建基线策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  FAILEDOPERATION_TOOMANYSTRATEGY = "FailedOperation.TooManyStrategy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBaselineStrategyWithContext(ctx context.Context, request *CreateBaselineStrategyRequest) (response *CreateBaselineStrategyResponse, err error) {
    if request == nil {
        request = NewCreateBaselineStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBaselineStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBaselineStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBuyBindTaskRequest() (request *CreateBuyBindTaskRequest) {
    request = &CreateBuyBindTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateBuyBindTask")
    
    
    return
}

func NewCreateBuyBindTaskResponse() (response *CreateBuyBindTaskResponse) {
    response = &CreateBuyBindTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBuyBindTask
// 新购授权自动绑定任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateBuyBindTask(request *CreateBuyBindTaskRequest) (response *CreateBuyBindTaskResponse, err error) {
    return c.CreateBuyBindTaskWithContext(context.Background(), request)
}

// CreateBuyBindTask
// 新购授权自动绑定任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateBuyBindTaskWithContext(ctx context.Context, request *CreateBuyBindTaskRequest) (response *CreateBuyBindTaskResponse, err error) {
    if request == nil {
        request = NewCreateBuyBindTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBuyBindTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBuyBindTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudProtectServiceOrderRecordRequest() (request *CreateCloudProtectServiceOrderRecordRequest) {
    request = &CreateCloudProtectServiceOrderRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateCloudProtectServiceOrderRecord")
    
    
    return
}

func NewCreateCloudProtectServiceOrderRecordResponse() (response *CreateCloudProtectServiceOrderRecordResponse) {
    response = &CreateCloudProtectServiceOrderRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudProtectServiceOrderRecord
// 云护航服务使用完成后，该接口可以确认收货
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCloudProtectServiceOrderRecord(request *CreateCloudProtectServiceOrderRecordRequest) (response *CreateCloudProtectServiceOrderRecordResponse, err error) {
    return c.CreateCloudProtectServiceOrderRecordWithContext(context.Background(), request)
}

// CreateCloudProtectServiceOrderRecord
// 云护航服务使用完成后，该接口可以确认收货
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCloudProtectServiceOrderRecordWithContext(ctx context.Context, request *CreateCloudProtectServiceOrderRecordRequest) (response *CreateCloudProtectServiceOrderRecordResponse, err error) {
    if request == nil {
        request = NewCreateCloudProtectServiceOrderRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudProtectServiceOrderRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudProtectServiceOrderRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmergencyVulScanRequest() (request *CreateEmergencyVulScanRequest) {
    request = &CreateEmergencyVulScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateEmergencyVulScan")
    
    
    return
}

func NewCreateEmergencyVulScanResponse() (response *CreateEmergencyVulScanResponse) {
    response = &CreateEmergencyVulScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEmergencyVulScan
// 创建应急漏洞扫描任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEmergencyVulScan(request *CreateEmergencyVulScanRequest) (response *CreateEmergencyVulScanResponse, err error) {
    return c.CreateEmergencyVulScanWithContext(context.Background(), request)
}

// CreateEmergencyVulScan
// 创建应急漏洞扫描任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEmergencyVulScanWithContext(ctx context.Context, request *CreateEmergencyVulScanRequest) (response *CreateEmergencyVulScanResponse, err error) {
    if request == nil {
        request = NewCreateEmergencyVulScanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmergencyVulScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmergencyVulScanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIncidentBacktrackingRequest() (request *CreateIncidentBacktrackingRequest) {
    request = &CreateIncidentBacktrackingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateIncidentBacktracking")
    
    
    return
}

func NewCreateIncidentBacktrackingResponse() (response *CreateIncidentBacktrackingResponse) {
    response = &CreateIncidentBacktrackingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIncidentBacktracking
// 对旗舰版机器单次触发事件调查及告警回溯
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIncidentBacktracking(request *CreateIncidentBacktrackingRequest) (response *CreateIncidentBacktrackingResponse, err error) {
    return c.CreateIncidentBacktrackingWithContext(context.Background(), request)
}

// CreateIncidentBacktracking
// 对旗舰版机器单次触发事件调查及告警回溯
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateIncidentBacktrackingWithContext(ctx context.Context, request *CreateIncidentBacktrackingRequest) (response *CreateIncidentBacktrackingResponse, err error) {
    if request == nil {
        request = NewCreateIncidentBacktrackingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIncidentBacktracking require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIncidentBacktrackingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLicenseOrderRequest() (request *CreateLicenseOrderRequest) {
    request = &CreateLicenseOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateLicenseOrder")
    
    
    return
}

func NewCreateLicenseOrderResponse() (response *CreateLicenseOrderResponse) {
    response = &CreateLicenseOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLicenseOrder
// CreateLicenseOrder 该接口可以创建专业版/旗舰版订单
//
// 支持预付费后付费创建
//
// 后付费订单直接创建成功
//
// 预付费订单仅下单不支付,需要调用计费支付接口进行支付
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLicenseOrder(request *CreateLicenseOrderRequest) (response *CreateLicenseOrderResponse, err error) {
    return c.CreateLicenseOrderWithContext(context.Background(), request)
}

// CreateLicenseOrder
// CreateLicenseOrder 该接口可以创建专业版/旗舰版订单
//
// 支持预付费后付费创建
//
// 后付费订单直接创建成功
//
// 预付费订单仅下单不支付,需要调用计费支付接口进行支付
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLicenseOrderWithContext(ctx context.Context, request *CreateLicenseOrderRequest) (response *CreateLicenseOrderResponse, err error) {
    if request == nil {
        request = NewCreateLicenseOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLicenseOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLicenseOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogExportRequest() (request *CreateLogExportRequest) {
    request = &CreateLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateLogExport")
    
    
    return
}

func NewCreateLogExportResponse() (response *CreateLogExportResponse) {
    response = &CreateLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLogExport
// 创建日志下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLogExport(request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    return c.CreateLogExportWithContext(context.Background(), request)
}

// CreateLogExport
// 创建日志下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLogExportWithContext(ctx context.Context, request *CreateLogExportRequest) (response *CreateLogExportResponse, err error) {
    if request == nil {
        request = NewCreateLogExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMaliciousRequestWhiteListRequest() (request *CreateMaliciousRequestWhiteListRequest) {
    request = &CreateMaliciousRequestWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateMaliciousRequestWhiteList")
    
    
    return
}

func NewCreateMaliciousRequestWhiteListResponse() (response *CreateMaliciousRequestWhiteListResponse) {
    response = &CreateMaliciousRequestWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMaliciousRequestWhiteList
// 添加恶意请求白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMaliciousRequestWhiteList(request *CreateMaliciousRequestWhiteListRequest) (response *CreateMaliciousRequestWhiteListResponse, err error) {
    return c.CreateMaliciousRequestWhiteListWithContext(context.Background(), request)
}

// CreateMaliciousRequestWhiteList
// 添加恶意请求白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMaliciousRequestWhiteListWithContext(ctx context.Context, request *CreateMaliciousRequestWhiteListRequest) (response *CreateMaliciousRequestWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateMaliciousRequestWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMaliciousRequestWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMaliciousRequestWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMalwareWhiteListRequest() (request *CreateMalwareWhiteListRequest) {
    request = &CreateMalwareWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateMalwareWhiteList")
    
    
    return
}

func NewCreateMalwareWhiteListResponse() (response *CreateMalwareWhiteListResponse) {
    response = &CreateMalwareWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMalwareWhiteList
// 创建木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) CreateMalwareWhiteList(request *CreateMalwareWhiteListRequest) (response *CreateMalwareWhiteListResponse, err error) {
    return c.CreateMalwareWhiteListWithContext(context.Background(), request)
}

// CreateMalwareWhiteList
// 创建木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) CreateMalwareWhiteListWithContext(ctx context.Context, request *CreateMalwareWhiteListRequest) (response *CreateMalwareWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateMalwareWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMalwareWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMalwareWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetAttackWhiteListRequest() (request *CreateNetAttackWhiteListRequest) {
    request = &CreateNetAttackWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateNetAttackWhiteList")
    
    
    return
}

func NewCreateNetAttackWhiteListResponse() (response *CreateNetAttackWhiteListResponse) {
    response = &CreateNetAttackWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetAttackWhiteList
// 创建网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) CreateNetAttackWhiteList(request *CreateNetAttackWhiteListRequest) (response *CreateNetAttackWhiteListResponse, err error) {
    return c.CreateNetAttackWhiteListWithContext(context.Background(), request)
}

// CreateNetAttackWhiteList
// 创建网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) CreateNetAttackWhiteListWithContext(ctx context.Context, request *CreateNetAttackWhiteListRequest) (response *CreateNetAttackWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateNetAttackWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetAttackWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetAttackWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProtectServerRequest() (request *CreateProtectServerRequest) {
    request = &CreateProtectServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateProtectServer")
    
    
    return
}

func NewCreateProtectServerResponse() (response *CreateProtectServerResponse) {
    response = &CreateProtectServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProtectServer
// 添加网站防护服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateProtectServer(request *CreateProtectServerRequest) (response *CreateProtectServerResponse, err error) {
    return c.CreateProtectServerWithContext(context.Background(), request)
}

// CreateProtectServer
// 添加网站防护服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateProtectServerWithContext(ctx context.Context, request *CreateProtectServerRequest) (response *CreateProtectServerResponse, err error) {
    if request == nil {
        request = NewCreateProtectServerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProtectServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProtectServerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRansomDefenseStrategyRequest() (request *CreateRansomDefenseStrategyRequest) {
    request = &CreateRansomDefenseStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateRansomDefenseStrategy")
    
    
    return
}

func NewCreateRansomDefenseStrategyResponse() (response *CreateRansomDefenseStrategyResponse) {
    response = &CreateRansomDefenseStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRansomDefenseStrategy
// 创建或修改防勒索策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRansomDefenseStrategy(request *CreateRansomDefenseStrategyRequest) (response *CreateRansomDefenseStrategyResponse, err error) {
    return c.CreateRansomDefenseStrategyWithContext(context.Background(), request)
}

// CreateRansomDefenseStrategy
// 创建或修改防勒索策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRansomDefenseStrategyWithContext(ctx context.Context, request *CreateRansomDefenseStrategyRequest) (response *CreateRansomDefenseStrategyResponse, err error) {
    if request == nil {
        request = NewCreateRansomDefenseStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRansomDefenseStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRansomDefenseStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScanMalwareSettingRequest() (request *CreateScanMalwareSettingRequest) {
    request = &CreateScanMalwareSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateScanMalwareSetting")
    
    
    return
}

func NewCreateScanMalwareSettingResponse() (response *CreateScanMalwareSettingResponse) {
    response = &CreateScanMalwareSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScanMalwareSetting
// 该接口可以对入侵检测-文件查杀扫描检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateScanMalwareSetting(request *CreateScanMalwareSettingRequest) (response *CreateScanMalwareSettingResponse, err error) {
    return c.CreateScanMalwareSettingWithContext(context.Background(), request)
}

// CreateScanMalwareSetting
// 该接口可以对入侵检测-文件查杀扫描检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateScanMalwareSettingWithContext(ctx context.Context, request *CreateScanMalwareSettingRequest) (response *CreateScanMalwareSettingResponse, err error) {
    if request == nil {
        request = NewCreateScanMalwareSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScanMalwareSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScanMalwareSettingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSearchLogRequest() (request *CreateSearchLogRequest) {
    request = &CreateSearchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateSearchLog")
    
    
    return
}

func NewCreateSearchLogResponse() (response *CreateSearchLogResponse) {
    response = &CreateSearchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSearchLog
// 添加历史搜索记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateSearchLog(request *CreateSearchLogRequest) (response *CreateSearchLogResponse, err error) {
    return c.CreateSearchLogWithContext(context.Background(), request)
}

// CreateSearchLog
// 添加历史搜索记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateSearchLogWithContext(ctx context.Context, request *CreateSearchLogRequest) (response *CreateSearchLogResponse, err error) {
    if request == nil {
        request = NewCreateSearchLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSearchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSearchLogResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSearchTemplateRequest() (request *CreateSearchTemplateRequest) {
    request = &CreateSearchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateSearchTemplate")
    
    
    return
}

func NewCreateSearchTemplateResponse() (response *CreateSearchTemplateResponse) {
    response = &CreateSearchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSearchTemplate
// 添加检索模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NAMEHASREPETITION = "InvalidParameter.NameHasRepetition"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateSearchTemplate(request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
    return c.CreateSearchTemplateWithContext(context.Background(), request)
}

// CreateSearchTemplate
// 添加检索模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NAMEHASREPETITION = "InvalidParameter.NameHasRepetition"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateSearchTemplateWithContext(ctx context.Context, request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSearchTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSearchTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSearchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVulFixRequest() (request *CreateVulFixRequest) {
    request = &CreateVulFixRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateVulFix")
    
    
    return
}

func NewCreateVulFixResponse() (response *CreateVulFixResponse) {
    response = &CreateVulFixResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVulFix
// 提交漏洞修护
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVulFix(request *CreateVulFixRequest) (response *CreateVulFixResponse, err error) {
    return c.CreateVulFixWithContext(context.Background(), request)
}

// CreateVulFix
// 提交漏洞修护
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVulFixWithContext(ctx context.Context, request *CreateVulFixRequest) (response *CreateVulFixResponse, err error) {
    if request == nil {
        request = NewCreateVulFixRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVulFix require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVulFixResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWhiteListOrderRequest() (request *CreateWhiteListOrderRequest) {
    request = &CreateWhiteListOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "CreateWhiteListOrder")
    
    
    return
}

func NewCreateWhiteListOrderResponse() (response *CreateWhiteListOrderResponse) {
    response = &CreateWhiteListOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWhiteListOrder
// 该接口可以创建白名单订单
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateWhiteListOrder(request *CreateWhiteListOrderRequest) (response *CreateWhiteListOrderResponse, err error) {
    return c.CreateWhiteListOrderWithContext(context.Background(), request)
}

// CreateWhiteListOrder
// 该接口可以创建白名单订单
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateWhiteListOrderWithContext(ctx context.Context, request *CreateWhiteListOrderRequest) (response *CreateWhiteListOrderResponse, err error) {
    if request == nil {
        request = NewCreateWhiteListOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWhiteListOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWhiteListOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllJavaMemShellsRequest() (request *DeleteAllJavaMemShellsRequest) {
    request = &DeleteAllJavaMemShellsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteAllJavaMemShells")
    
    
    return
}

func NewDeleteAllJavaMemShellsResponse() (response *DeleteAllJavaMemShellsResponse) {
    response = &DeleteAllJavaMemShellsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAllJavaMemShells
// 删除全部java内存马事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAllJavaMemShells(request *DeleteAllJavaMemShellsRequest) (response *DeleteAllJavaMemShellsResponse, err error) {
    return c.DeleteAllJavaMemShellsWithContext(context.Background(), request)
}

// DeleteAllJavaMemShells
// 删除全部java内存马事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAllJavaMemShellsWithContext(ctx context.Context, request *DeleteAllJavaMemShellsRequest) (response *DeleteAllJavaMemShellsResponse, err error) {
    if request == nil {
        request = NewDeleteAllJavaMemShellsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllJavaMemShells require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllJavaMemShellsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttackLogsRequest() (request *DeleteAttackLogsRequest) {
    request = &DeleteAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteAttackLogs")
    
    
    return
}

func NewDeleteAttackLogsResponse() (response *DeleteAttackLogsResponse) {
    response = &DeleteAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAttackLogs
// 删除网络攻击日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAttackLogs(request *DeleteAttackLogsRequest) (response *DeleteAttackLogsResponse, err error) {
    return c.DeleteAttackLogsWithContext(context.Background(), request)
}

// DeleteAttackLogs
// 删除网络攻击日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAttackLogsWithContext(ctx context.Context, request *DeleteAttackLogsRequest) (response *DeleteAttackLogsResponse, err error) {
    if request == nil {
        request = NewDeleteAttackLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttackLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBanWhiteListRequest() (request *DeleteBanWhiteListRequest) {
    request = &DeleteBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBanWhiteList")
    
    
    return
}

func NewDeleteBanWhiteListResponse() (response *DeleteBanWhiteListResponse) {
    response = &DeleteBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBanWhiteList
// 删除阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBanWhiteList(request *DeleteBanWhiteListRequest) (response *DeleteBanWhiteListResponse, err error) {
    return c.DeleteBanWhiteListWithContext(context.Background(), request)
}

// DeleteBanWhiteList
// 删除阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBanWhiteListWithContext(ctx context.Context, request *DeleteBanWhiteListRequest) (response *DeleteBanWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteBanWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBanWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselinePolicyRequest() (request *DeleteBaselinePolicyRequest) {
    request = &DeleteBaselinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselinePolicy")
    
    
    return
}

func NewDeleteBaselinePolicyResponse() (response *DeleteBaselinePolicyResponse) {
    response = &DeleteBaselinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBaselinePolicy
// 删除基线策略配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBaselinePolicy(request *DeleteBaselinePolicyRequest) (response *DeleteBaselinePolicyResponse, err error) {
    return c.DeleteBaselinePolicyWithContext(context.Background(), request)
}

// DeleteBaselinePolicy
// 删除基线策略配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBaselinePolicyWithContext(ctx context.Context, request *DeleteBaselinePolicyRequest) (response *DeleteBaselinePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteBaselinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaselinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselineRuleRequest() (request *DeleteBaselineRuleRequest) {
    request = &DeleteBaselineRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineRule")
    
    
    return
}

func NewDeleteBaselineRuleResponse() (response *DeleteBaselineRuleResponse) {
    response = &DeleteBaselineRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBaselineRule
// 删除基线规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteBaselineRule(request *DeleteBaselineRuleRequest) (response *DeleteBaselineRuleResponse, err error) {
    return c.DeleteBaselineRuleWithContext(context.Background(), request)
}

// DeleteBaselineRule
// 删除基线规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteBaselineRuleWithContext(ctx context.Context, request *DeleteBaselineRuleRequest) (response *DeleteBaselineRuleResponse, err error) {
    if request == nil {
        request = NewDeleteBaselineRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaselineRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselineRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselineRuleIgnoreRequest() (request *DeleteBaselineRuleIgnoreRequest) {
    request = &DeleteBaselineRuleIgnoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineRuleIgnore")
    
    
    return
}

func NewDeleteBaselineRuleIgnoreResponse() (response *DeleteBaselineRuleIgnoreResponse) {
    response = &DeleteBaselineRuleIgnoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBaselineRuleIgnore
// 删除基线忽略规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteBaselineRuleIgnore(request *DeleteBaselineRuleIgnoreRequest) (response *DeleteBaselineRuleIgnoreResponse, err error) {
    return c.DeleteBaselineRuleIgnoreWithContext(context.Background(), request)
}

// DeleteBaselineRuleIgnore
// 删除基线忽略规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteBaselineRuleIgnoreWithContext(ctx context.Context, request *DeleteBaselineRuleIgnoreRequest) (response *DeleteBaselineRuleIgnoreResponse, err error) {
    if request == nil {
        request = NewDeleteBaselineRuleIgnoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaselineRuleIgnore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselineRuleIgnoreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselineStrategyRequest() (request *DeleteBaselineStrategyRequest) {
    request = &DeleteBaselineStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineStrategy")
    
    
    return
}

func NewDeleteBaselineStrategyResponse() (response *DeleteBaselineStrategyResponse) {
    response = &DeleteBaselineStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBaselineStrategy
// 根据基线策略id删除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBaselineStrategy(request *DeleteBaselineStrategyRequest) (response *DeleteBaselineStrategyResponse, err error) {
    return c.DeleteBaselineStrategyWithContext(context.Background(), request)
}

// DeleteBaselineStrategy
// 根据基线策略id删除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBaselineStrategyWithContext(ctx context.Context, request *DeleteBaselineStrategyRequest) (response *DeleteBaselineStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteBaselineStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaselineStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselineStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselineWeakPasswordRequest() (request *DeleteBaselineWeakPasswordRequest) {
    request = &DeleteBaselineWeakPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBaselineWeakPassword")
    
    
    return
}

func NewDeleteBaselineWeakPasswordResponse() (response *DeleteBaselineWeakPasswordResponse) {
    response = &DeleteBaselineWeakPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBaselineWeakPassword
// 删除基线弱口令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteBaselineWeakPassword(request *DeleteBaselineWeakPasswordRequest) (response *DeleteBaselineWeakPasswordResponse, err error) {
    return c.DeleteBaselineWeakPasswordWithContext(context.Background(), request)
}

// DeleteBaselineWeakPassword
// 删除基线弱口令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteBaselineWeakPasswordWithContext(ctx context.Context, request *DeleteBaselineWeakPasswordRequest) (response *DeleteBaselineWeakPasswordResponse, err error) {
    if request == nil {
        request = NewDeleteBaselineWeakPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaselineWeakPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselineWeakPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBashEventsRequest() (request *DeleteBashEventsRequest) {
    request = &DeleteBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashEvents")
    
    
    return
}

func NewDeleteBashEventsResponse() (response *DeleteBashEventsResponse) {
    response = &DeleteBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBashEvents
// 根据Ids删除高危命令事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashEvents(request *DeleteBashEventsRequest) (response *DeleteBashEventsResponse, err error) {
    return c.DeleteBashEventsWithContext(context.Background(), request)
}

// DeleteBashEvents
// 根据Ids删除高危命令事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashEventsWithContext(ctx context.Context, request *DeleteBashEventsRequest) (response *DeleteBashEventsResponse, err error) {
    if request == nil {
        request = NewDeleteBashEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBashEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBashPoliciesRequest() (request *DeleteBashPoliciesRequest) {
    request = &DeleteBashPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashPolicies")
    
    
    return
}

func NewDeleteBashPoliciesResponse() (response *DeleteBashPoliciesResponse) {
    response = &DeleteBashPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBashPolicies
// 删除高危命令策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashPolicies(request *DeleteBashPoliciesRequest) (response *DeleteBashPoliciesResponse, err error) {
    return c.DeleteBashPoliciesWithContext(context.Background(), request)
}

// DeleteBashPolicies
// 删除高危命令策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashPoliciesWithContext(ctx context.Context, request *DeleteBashPoliciesRequest) (response *DeleteBashPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteBashPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBashPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBashPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBashRulesRequest() (request *DeleteBashRulesRequest) {
    request = &DeleteBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBashRules")
    
    
    return
}

func NewDeleteBashRulesResponse() (response *DeleteBashRulesResponse) {
    response = &DeleteBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBashRules
// 删除高危命令规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashRules(request *DeleteBashRulesRequest) (response *DeleteBashRulesResponse, err error) {
    return c.DeleteBashRulesWithContext(context.Background(), request)
}

// DeleteBashRules
// 删除高危命令规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBashRulesWithContext(ctx context.Context, request *DeleteBashRulesRequest) (response *DeleteBashRulesResponse, err error) {
    if request == nil {
        request = NewDeleteBashRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBashRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
    request = &DeleteBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteBruteAttacks")
    
    
    return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
    response = &DeleteBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBruteAttacks
// 本接口 (DeleteBruteAttacks) 用于删除暴力破解记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    return c.DeleteBruteAttacksWithContext(context.Background(), request)
}

// DeleteBruteAttacks
// 本接口 (DeleteBruteAttacks) 用于删除暴力破解记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBruteAttacksWithContext(ctx context.Context, request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDeleteBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLicenseRecordRequest() (request *DeleteLicenseRecordRequest) {
    request = &DeleteLicenseRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteLicenseRecord")
    
    
    return
}

func NewDeleteLicenseRecordResponse() (response *DeleteLicenseRecordResponse) {
    response = &DeleteLicenseRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLicenseRecord
// 对授权管理-订单列表内已过期的订单进行删除.(删除后的订单不在统计范畴内)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLicenseRecord(request *DeleteLicenseRecordRequest) (response *DeleteLicenseRecordResponse, err error) {
    return c.DeleteLicenseRecordWithContext(context.Background(), request)
}

// DeleteLicenseRecord
// 对授权管理-订单列表内已过期的订单进行删除.(删除后的订单不在统计范畴内)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLicenseRecordWithContext(ctx context.Context, request *DeleteLicenseRecordRequest) (response *DeleteLicenseRecordResponse, err error) {
    if request == nil {
        request = NewDeleteLicenseRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLicenseRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLicenseRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLicenseRecordAllRequest() (request *DeleteLicenseRecordAllRequest) {
    request = &DeleteLicenseRecordAllRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteLicenseRecordAll")
    
    
    return
}

func NewDeleteLicenseRecordAllResponse() (response *DeleteLicenseRecordAllResponse) {
    response = &DeleteLicenseRecordAllResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLicenseRecordAll
// 删除授权全部记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLicenseRecordAll(request *DeleteLicenseRecordAllRequest) (response *DeleteLicenseRecordAllResponse, err error) {
    return c.DeleteLicenseRecordAllWithContext(context.Background(), request)
}

// DeleteLicenseRecordAll
// 删除授权全部记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLicenseRecordAllWithContext(ctx context.Context, request *DeleteLicenseRecordAllRequest) (response *DeleteLicenseRecordAllResponse, err error) {
    if request == nil {
        request = NewDeleteLicenseRecordAllRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLicenseRecordAll require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLicenseRecordAllResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogExportRequest() (request *DeleteLogExportRequest) {
    request = &DeleteLogExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteLogExport")
    
    
    return
}

func NewDeleteLogExportResponse() (response *DeleteLogExportResponse) {
    response = &DeleteLogExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLogExport
// 删除日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLogExport(request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    return c.DeleteLogExportWithContext(context.Background(), request)
}

// DeleteLogExport
// 删除日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLogExportWithContext(ctx context.Context, request *DeleteLogExportRequest) (response *DeleteLogExportResponse, err error) {
    if request == nil {
        request = NewDeleteLogExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
    request = &DeleteLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteLoginWhiteList")
    
    
    return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
    response = &DeleteLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoginWhiteList
// 本接口用于删除异地登录白名单规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    return c.DeleteLoginWhiteListWithContext(context.Background(), request)
}

// DeleteLoginWhiteList
// 本接口用于删除异地登录白名单规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLoginWhiteListWithContext(ctx context.Context, request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
    request = &DeleteMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachine")
    
    
    return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
    response = &DeleteMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMachine
// 本接口（DeleteMachine）用于卸载主机安全客户端。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    return c.DeleteMachineWithContext(context.Background(), request)
}

// DeleteMachine
// 本接口（DeleteMachine）用于卸载主机安全客户端。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineWithContext(ctx context.Context, request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    if request == nil {
        request = NewDeleteMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineClearHistoryRequest() (request *DeleteMachineClearHistoryRequest) {
    request = &DeleteMachineClearHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachineClearHistory")
    
    
    return
}

func NewDeleteMachineClearHistoryResponse() (response *DeleteMachineClearHistoryResponse) {
    response = &DeleteMachineClearHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMachineClearHistory
// 删除机器清理记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineClearHistory(request *DeleteMachineClearHistoryRequest) (response *DeleteMachineClearHistoryResponse, err error) {
    return c.DeleteMachineClearHistoryWithContext(context.Background(), request)
}

// DeleteMachineClearHistory
// 删除机器清理记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineClearHistoryWithContext(ctx context.Context, request *DeleteMachineClearHistoryRequest) (response *DeleteMachineClearHistoryResponse, err error) {
    if request == nil {
        request = NewDeleteMachineClearHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineClearHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineClearHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineTagRequest() (request *DeleteMachineTagRequest) {
    request = &DeleteMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMachineTag")
    
    
    return
}

func NewDeleteMachineTagResponse() (response *DeleteMachineTagResponse) {
    response = &DeleteMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMachineTag
// 删除服务器关联的标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineTag(request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    return c.DeleteMachineTagWithContext(context.Background(), request)
}

// DeleteMachineTag
// 删除服务器关联的标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMachineTagWithContext(ctx context.Context, request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    if request == nil {
        request = NewDeleteMachineTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMaliciousRequestWhiteListRequest() (request *DeleteMaliciousRequestWhiteListRequest) {
    request = &DeleteMaliciousRequestWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaliciousRequestWhiteList")
    
    
    return
}

func NewDeleteMaliciousRequestWhiteListResponse() (response *DeleteMaliciousRequestWhiteListResponse) {
    response = &DeleteMaliciousRequestWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMaliciousRequestWhiteList
// 删除恶意请求白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMaliciousRequestWhiteList(request *DeleteMaliciousRequestWhiteListRequest) (response *DeleteMaliciousRequestWhiteListResponse, err error) {
    return c.DeleteMaliciousRequestWhiteListWithContext(context.Background(), request)
}

// DeleteMaliciousRequestWhiteList
// 删除恶意请求白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMaliciousRequestWhiteListWithContext(ctx context.Context, request *DeleteMaliciousRequestWhiteListRequest) (response *DeleteMaliciousRequestWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMaliciousRequestWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMaliciousRequestWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMaliciousRequestsRequest() (request *DeleteMaliciousRequestsRequest) {
    request = &DeleteMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMaliciousRequests")
    
    
    return
}

func NewDeleteMaliciousRequestsResponse() (response *DeleteMaliciousRequestsResponse) {
    response = &DeleteMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMaliciousRequests
// 本接口 (DeleteMaliciousRequests) 用于删除恶意请求记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    return c.DeleteMaliciousRequestsWithContext(context.Background(), request)
}

// DeleteMaliciousRequests
// 本接口 (DeleteMaliciousRequests) 用于删除恶意请求记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMaliciousRequestsWithContext(ctx context.Context, request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwareScanTaskRequest() (request *DeleteMalwareScanTaskRequest) {
    request = &DeleteMalwareScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwareScanTask")
    
    
    return
}

func NewDeleteMalwareScanTaskResponse() (response *DeleteMalwareScanTaskResponse) {
    response = &DeleteMalwareScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMalwareScanTask
// 入侵管理-终止扫描任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMalwareScanTask(request *DeleteMalwareScanTaskRequest) (response *DeleteMalwareScanTaskResponse, err error) {
    return c.DeleteMalwareScanTaskWithContext(context.Background(), request)
}

// DeleteMalwareScanTask
// 入侵管理-终止扫描任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteMalwareScanTaskWithContext(ctx context.Context, request *DeleteMalwareScanTaskRequest) (response *DeleteMalwareScanTaskResponse, err error) {
    if request == nil {
        request = NewDeleteMalwareScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMalwareScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMalwareScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwareWhiteListRequest() (request *DeleteMalwareWhiteListRequest) {
    request = &DeleteMalwareWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwareWhiteList")
    
    
    return
}

func NewDeleteMalwareWhiteListResponse() (response *DeleteMalwareWhiteListResponse) {
    response = &DeleteMalwareWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMalwareWhiteList
// 删除木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMalwareWhiteList(request *DeleteMalwareWhiteListRequest) (response *DeleteMalwareWhiteListResponse, err error) {
    return c.DeleteMalwareWhiteListWithContext(context.Background(), request)
}

// DeleteMalwareWhiteList
// 删除木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMalwareWhiteListWithContext(ctx context.Context, request *DeleteMalwareWhiteListRequest) (response *DeleteMalwareWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteMalwareWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMalwareWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMalwareWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
    request = &DeleteMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteMalwares")
    
    
    return
}

func NewDeleteMalwaresResponse() (response *DeleteMalwaresResponse) {
    response = &DeleteMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMalwares
// 本接口 (DeleteMalwares) 用于删除木马记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    return c.DeleteMalwaresWithContext(context.Background(), request)
}

// DeleteMalwares
// 本接口 (DeleteMalwares) 用于删除木马记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMalwaresWithContext(ctx context.Context, request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    if request == nil {
        request = NewDeleteMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetAttackWhiteListRequest() (request *DeleteNetAttackWhiteListRequest) {
    request = &DeleteNetAttackWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteNetAttackWhiteList")
    
    
    return
}

func NewDeleteNetAttackWhiteListResponse() (response *DeleteNetAttackWhiteListResponse) {
    response = &DeleteNetAttackWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNetAttackWhiteList
// 删除网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetAttackWhiteList(request *DeleteNetAttackWhiteListRequest) (response *DeleteNetAttackWhiteListResponse, err error) {
    return c.DeleteNetAttackWhiteListWithContext(context.Background(), request)
}

// DeleteNetAttackWhiteList
// 删除网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetAttackWhiteListWithContext(ctx context.Context, request *DeleteNetAttackWhiteListRequest) (response *DeleteNetAttackWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteNetAttackWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetAttackWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetAttackWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNonlocalLoginPlacesRequest() (request *DeleteNonlocalLoginPlacesRequest) {
    request = &DeleteNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteNonlocalLoginPlaces")
    
    
    return
}

func NewDeleteNonlocalLoginPlacesResponse() (response *DeleteNonlocalLoginPlacesResponse) {
    response = &DeleteNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNonlocalLoginPlaces
// 本接口 (DeleteNonlocalLoginPlaces) 用于删除异地登录记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    return c.DeleteNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DeleteNonlocalLoginPlaces
// 本接口 (DeleteNonlocalLoginPlaces) 用于删除异地登录记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNonlocalLoginPlacesWithContext(ctx context.Context, request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivilegeEventsRequest() (request *DeletePrivilegeEventsRequest) {
    request = &DeletePrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeletePrivilegeEvents")
    
    
    return
}

func NewDeletePrivilegeEventsResponse() (response *DeletePrivilegeEventsResponse) {
    response = &DeletePrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrivilegeEvents
// 根据Ids删除本地提权
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeletePrivilegeEvents(request *DeletePrivilegeEventsRequest) (response *DeletePrivilegeEventsResponse, err error) {
    return c.DeletePrivilegeEventsWithContext(context.Background(), request)
}

// DeletePrivilegeEvents
// 根据Ids删除本地提权
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeletePrivilegeEventsWithContext(ctx context.Context, request *DeletePrivilegeEventsRequest) (response *DeletePrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewDeletePrivilegeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivilegeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivilegeRulesRequest() (request *DeletePrivilegeRulesRequest) {
    request = &DeletePrivilegeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeletePrivilegeRules")
    
    
    return
}

func NewDeletePrivilegeRulesResponse() (response *DeletePrivilegeRulesResponse) {
    response = &DeletePrivilegeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrivilegeRules
// 删除本地提权规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrivilegeRules(request *DeletePrivilegeRulesRequest) (response *DeletePrivilegeRulesResponse, err error) {
    return c.DeletePrivilegeRulesWithContext(context.Background(), request)
}

// DeletePrivilegeRules
// 删除本地提权规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrivilegeRulesWithContext(ctx context.Context, request *DeletePrivilegeRulesRequest) (response *DeletePrivilegeRulesResponse, err error) {
    if request == nil {
        request = NewDeletePrivilegeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivilegeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivilegeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProtectDirRequest() (request *DeleteProtectDirRequest) {
    request = &DeleteProtectDirRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteProtectDir")
    
    
    return
}

func NewDeleteProtectDirResponse() (response *DeleteProtectDirResponse) {
    response = &DeleteProtectDirResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProtectDir
// 删除防护网站
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProtectDir(request *DeleteProtectDirRequest) (response *DeleteProtectDirResponse, err error) {
    return c.DeleteProtectDirWithContext(context.Background(), request)
}

// DeleteProtectDir
// 删除防护网站
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProtectDirWithContext(ctx context.Context, request *DeleteProtectDirRequest) (response *DeleteProtectDirResponse, err error) {
    if request == nil {
        request = NewDeleteProtectDirRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProtectDir require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProtectDirResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
    request = &DeleteReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteReverseShellEvents")
    
    
    return
}

func NewDeleteReverseShellEventsResponse() (response *DeleteReverseShellEventsResponse) {
    response = &DeleteReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReverseShellEvents
// 根据Ids删除反弹Shell事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteReverseShellEvents(request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
    return c.DeleteReverseShellEventsWithContext(context.Background(), request)
}

// DeleteReverseShellEvents
// 根据Ids删除反弹Shell事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteReverseShellEventsWithContext(ctx context.Context, request *DeleteReverseShellEventsRequest) (response *DeleteReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReverseShellEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReverseShellRulesRequest() (request *DeleteReverseShellRulesRequest) {
    request = &DeleteReverseShellRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteReverseShellRules")
    
    
    return
}

func NewDeleteReverseShellRulesResponse() (response *DeleteReverseShellRulesResponse) {
    response = &DeleteReverseShellRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReverseShellRules
// 删除反弹Shell规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteReverseShellRules(request *DeleteReverseShellRulesRequest) (response *DeleteReverseShellRulesResponse, err error) {
    return c.DeleteReverseShellRulesWithContext(context.Background(), request)
}

// DeleteReverseShellRules
// 删除反弹Shell规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteReverseShellRulesWithContext(ctx context.Context, request *DeleteReverseShellRulesRequest) (response *DeleteReverseShellRulesResponse, err error) {
    if request == nil {
        request = NewDeleteReverseShellRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReverseShellRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReverseShellRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskDnsEventRequest() (request *DeleteRiskDnsEventRequest) {
    request = &DeleteRiskDnsEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteRiskDnsEvent")
    
    
    return
}

func NewDeleteRiskDnsEventResponse() (response *DeleteRiskDnsEventResponse) {
    response = &DeleteRiskDnsEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRiskDnsEvent
// 删除恶意请求事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRiskDnsEvent(request *DeleteRiskDnsEventRequest) (response *DeleteRiskDnsEventResponse, err error) {
    return c.DeleteRiskDnsEventWithContext(context.Background(), request)
}

// DeleteRiskDnsEvent
// 删除恶意请求事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRiskDnsEventWithContext(ctx context.Context, request *DeleteRiskDnsEventRequest) (response *DeleteRiskDnsEventResponse, err error) {
    if request == nil {
        request = NewDeleteRiskDnsEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskDnsEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskDnsEventResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRiskDnsPolicyRequest() (request *DeleteRiskDnsPolicyRequest) {
    request = &DeleteRiskDnsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteRiskDnsPolicy")
    
    
    return
}

func NewDeleteRiskDnsPolicyResponse() (response *DeleteRiskDnsPolicyResponse) {
    response = &DeleteRiskDnsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRiskDnsPolicy
// 删除恶意请求策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteRiskDnsPolicy(request *DeleteRiskDnsPolicyRequest) (response *DeleteRiskDnsPolicyResponse, err error) {
    return c.DeleteRiskDnsPolicyWithContext(context.Background(), request)
}

// DeleteRiskDnsPolicy
// 删除恶意请求策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteRiskDnsPolicyWithContext(ctx context.Context, request *DeleteRiskDnsPolicyRequest) (response *DeleteRiskDnsPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteRiskDnsPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRiskDnsPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRiskDnsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScanTaskRequest() (request *DeleteScanTaskRequest) {
    request = &DeleteScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteScanTask")
    
    
    return
}

func NewDeleteScanTaskResponse() (response *DeleteScanTaskResponse) {
    response = &DeleteScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScanTask
// DeleteScanTask 该接口可以对指定类型的扫描任务进行停止扫描;
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteScanTask(request *DeleteScanTaskRequest) (response *DeleteScanTaskResponse, err error) {
    return c.DeleteScanTaskWithContext(context.Background(), request)
}

// DeleteScanTask
// DeleteScanTask 该接口可以对指定类型的扫描任务进行停止扫描;
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteScanTaskWithContext(ctx context.Context, request *DeleteScanTaskRequest) (response *DeleteScanTaskResponse, err error) {
    if request == nil {
        request = NewDeleteScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSearchTemplateRequest() (request *DeleteSearchTemplateRequest) {
    request = &DeleteSearchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteSearchTemplate")
    
    
    return
}

func NewDeleteSearchTemplateResponse() (response *DeleteSearchTemplateResponse) {
    response = &DeleteSearchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSearchTemplate
// 删除检索模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSearchTemplate(request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
    return c.DeleteSearchTemplateWithContext(context.Background(), request)
}

// DeleteSearchTemplate
// 删除检索模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSearchTemplateWithContext(ctx context.Context, request *DeleteSearchTemplateRequest) (response *DeleteSearchTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSearchTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSearchTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSearchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
    request = &DeleteTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteTags")
    
    
    return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
    response = &DeleteTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTags
// 删除标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTags(request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
    return c.DeleteTagsWithContext(context.Background(), request)
}

// DeleteTags
// 删除标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTagsWithContext(ctx context.Context, request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
    if request == nil {
        request = NewDeleteTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebHookPolicyRequest() (request *DeleteWebHookPolicyRequest) {
    request = &DeleteWebHookPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookPolicy")
    
    
    return
}

func NewDeleteWebHookPolicyResponse() (response *DeleteWebHookPolicyResponse) {
    response = &DeleteWebHookPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWebHookPolicy
// 删除告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookPolicy(request *DeleteWebHookPolicyRequest) (response *DeleteWebHookPolicyResponse, err error) {
    return c.DeleteWebHookPolicyWithContext(context.Background(), request)
}

// DeleteWebHookPolicy
// 删除告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookPolicyWithContext(ctx context.Context, request *DeleteWebHookPolicyRequest) (response *DeleteWebHookPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteWebHookPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebHookPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebHookPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebHookReceiverRequest() (request *DeleteWebHookReceiverRequest) {
    request = &DeleteWebHookReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookReceiver")
    
    
    return
}

func NewDeleteWebHookReceiverResponse() (response *DeleteWebHookReceiverResponse) {
    response = &DeleteWebHookReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWebHookReceiver
// 删除告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookReceiver(request *DeleteWebHookReceiverRequest) (response *DeleteWebHookReceiverResponse, err error) {
    return c.DeleteWebHookReceiverWithContext(context.Background(), request)
}

// DeleteWebHookReceiver
// 删除告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookReceiverWithContext(ctx context.Context, request *DeleteWebHookReceiverRequest) (response *DeleteWebHookReceiverResponse, err error) {
    if request == nil {
        request = NewDeleteWebHookReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebHookReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebHookReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebHookRuleRequest() (request *DeleteWebHookRuleRequest) {
    request = &DeleteWebHookRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebHookRule")
    
    
    return
}

func NewDeleteWebHookRuleResponse() (response *DeleteWebHookRuleResponse) {
    response = &DeleteWebHookRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWebHookRule
// 删除企微机器人规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookRule(request *DeleteWebHookRuleRequest) (response *DeleteWebHookRuleResponse, err error) {
    return c.DeleteWebHookRuleWithContext(context.Background(), request)
}

// DeleteWebHookRule
// 删除企微机器人规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWebHookRuleWithContext(ctx context.Context, request *DeleteWebHookRuleRequest) (response *DeleteWebHookRuleResponse, err error) {
    if request == nil {
        request = NewDeleteWebHookRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebHookRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebHookRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebPageEventLogRequest() (request *DeleteWebPageEventLogRequest) {
    request = &DeleteWebPageEventLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteWebPageEventLog")
    
    
    return
}

func NewDeleteWebPageEventLogResponse() (response *DeleteWebPageEventLogResponse) {
    response = &DeleteWebPageEventLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWebPageEventLog
// 网站防篡改-删除事件记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWebPageEventLog(request *DeleteWebPageEventLogRequest) (response *DeleteWebPageEventLogResponse, err error) {
    return c.DeleteWebPageEventLogWithContext(context.Background(), request)
}

// DeleteWebPageEventLog
// 网站防篡改-删除事件记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWebPageEventLogWithContext(ctx context.Context, request *DeleteWebPageEventLogRequest) (response *DeleteWebPageEventLogResponse, err error) {
    if request == nil {
        request = NewDeleteWebPageEventLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebPageEventLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebPageEventLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeABTestConfigRequest() (request *DescribeABTestConfigRequest) {
    request = &DescribeABTestConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeABTestConfig")
    
    
    return
}

func NewDescribeABTestConfigResponse() (response *DescribeABTestConfigResponse) {
    response = &DescribeABTestConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeABTestConfig
// 获取用户当前灰度配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeABTestConfig(request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
    return c.DescribeABTestConfigWithContext(context.Background(), request)
}

// DescribeABTestConfig
// 获取用户当前灰度配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeABTestConfigWithContext(ctx context.Context, request *DescribeABTestConfigRequest) (response *DescribeABTestConfigResponse, err error) {
    if request == nil {
        request = NewDescribeABTestConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeABTestConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeABTestConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAESKeyRequest() (request *DescribeAESKeyRequest) {
    request = &DescribeAESKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAESKey")
    
    
    return
}

func NewDescribeAESKeyResponse() (response *DescribeAESKeyResponse) {
    response = &DescribeAESKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAESKey
// 获取配置的aeskey和aesiv
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAESKey(request *DescribeAESKeyRequest) (response *DescribeAESKeyResponse, err error) {
    return c.DescribeAESKeyWithContext(context.Background(), request)
}

// DescribeAESKey
// 获取配置的aeskey和aesiv
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAESKeyWithContext(ctx context.Context, request *DescribeAESKeyRequest) (response *DescribeAESKeyResponse, err error) {
    if request == nil {
        request = NewDescribeAESKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAESKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAESKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
    request = &DescribeAccountStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAccountStatistics")
    
    
    return
}

func NewDescribeAccountStatisticsResponse() (response *DescribeAccountStatisticsResponse) {
    response = &DescribeAccountStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountStatistics
// 本接口 (DescribeAccountStatistics) 用于获取帐号统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountStatistics(request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    return c.DescribeAccountStatisticsWithContext(context.Background(), request)
}

// DescribeAccountStatistics
// 本接口 (DescribeAccountStatistics) 用于获取帐号统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountStatisticsWithContext(ctx context.Context, request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentInstallCommandRequest() (request *DescribeAgentInstallCommandRequest) {
    request = &DescribeAgentInstallCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallCommand")
    
    
    return
}

func NewDescribeAgentInstallCommandResponse() (response *DescribeAgentInstallCommandResponse) {
    response = &DescribeAgentInstallCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentInstallCommand
// 获取agent安装命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeAgentInstallCommand(request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
    return c.DescribeAgentInstallCommandWithContext(context.Background(), request)
}

// DescribeAgentInstallCommand
// 获取agent安装命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeAgentInstallCommandWithContext(ctx context.Context, request *DescribeAgentInstallCommandRequest) (response *DescribeAgentInstallCommandResponse, err error) {
    if request == nil {
        request = NewDescribeAgentInstallCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentInstallCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentInstallCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentInstallationTokenRequest() (request *DescribeAgentInstallationTokenRequest) {
    request = &DescribeAgentInstallationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentInstallationToken")
    
    
    return
}

func NewDescribeAgentInstallationTokenResponse() (response *DescribeAgentInstallationTokenResponse) {
    response = &DescribeAgentInstallationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentInstallationToken
// 混合云安装agent token获取
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentInstallationToken(request *DescribeAgentInstallationTokenRequest) (response *DescribeAgentInstallationTokenResponse, err error) {
    return c.DescribeAgentInstallationTokenWithContext(context.Background(), request)
}

// DescribeAgentInstallationToken
// 混合云安装agent token获取
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentInstallationTokenWithContext(ctx context.Context, request *DescribeAgentInstallationTokenRequest) (response *DescribeAgentInstallationTokenResponse, err error) {
    if request == nil {
        request = NewDescribeAgentInstallationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentInstallationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentInstallationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmIncidentNodesRequest() (request *DescribeAlarmIncidentNodesRequest) {
    request = &DescribeAlarmIncidentNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAlarmIncidentNodes")
    
    
    return
}

func NewDescribeAlarmIncidentNodesResponse() (response *DescribeAlarmIncidentNodesResponse) {
    response = &DescribeAlarmIncidentNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmIncidentNodes
// 获取告警点所在事件的所有节点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAlarmIncidentNodes(request *DescribeAlarmIncidentNodesRequest) (response *DescribeAlarmIncidentNodesResponse, err error) {
    return c.DescribeAlarmIncidentNodesWithContext(context.Background(), request)
}

// DescribeAlarmIncidentNodes
// 获取告警点所在事件的所有节点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAlarmIncidentNodesWithContext(ctx context.Context, request *DescribeAlarmIncidentNodesRequest) (response *DescribeAlarmIncidentNodesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmIncidentNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmIncidentNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmIncidentNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmVertexIdRequest() (request *DescribeAlarmVertexIdRequest) {
    request = &DescribeAlarmVertexIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAlarmVertexId")
    
    
    return
}

func NewDescribeAlarmVertexIdResponse() (response *DescribeAlarmVertexIdResponse) {
    response = &DescribeAlarmVertexIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmVertexId
// 查询告警点id列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAlarmVertexId(request *DescribeAlarmVertexIdRequest) (response *DescribeAlarmVertexIdResponse, err error) {
    return c.DescribeAlarmVertexIdWithContext(context.Background(), request)
}

// DescribeAlarmVertexId
// 查询告警点id列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAlarmVertexIdWithContext(ctx context.Context, request *DescribeAlarmVertexIdRequest) (response *DescribeAlarmVertexIdResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmVertexIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmVertexId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmVertexIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetAppCountRequest() (request *DescribeAssetAppCountRequest) {
    request = &DescribeAssetAppCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppCount")
    
    
    return
}

func NewDescribeAssetAppCountResponse() (response *DescribeAssetAppCountResponse) {
    response = &DescribeAssetAppCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetAppCount
// 获取所有软件应用数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetAppCount(request *DescribeAssetAppCountRequest) (response *DescribeAssetAppCountResponse, err error) {
    return c.DescribeAssetAppCountWithContext(context.Background(), request)
}

// DescribeAssetAppCount
// 获取所有软件应用数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetAppCountWithContext(ctx context.Context, request *DescribeAssetAppCountRequest) (response *DescribeAssetAppCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetAppCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetAppCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetAppCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetAppListRequest() (request *DescribeAssetAppListRequest) {
    request = &DescribeAssetAppListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppList")
    
    
    return
}

func NewDescribeAssetAppListResponse() (response *DescribeAssetAppListResponse) {
    response = &DescribeAssetAppListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetAppList
// 查询应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppList(request *DescribeAssetAppListRequest) (response *DescribeAssetAppListResponse, err error) {
    return c.DescribeAssetAppListWithContext(context.Background(), request)
}

// DescribeAssetAppList
// 查询应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetAppListWithContext(ctx context.Context, request *DescribeAssetAppListRequest) (response *DescribeAssetAppListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetAppListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetAppList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetAppListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetAppProcessListRequest() (request *DescribeAssetAppProcessListRequest) {
    request = &DescribeAssetAppProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetAppProcessList")
    
    
    return
}

func NewDescribeAssetAppProcessListResponse() (response *DescribeAssetAppProcessListResponse) {
    response = &DescribeAssetAppProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetAppProcessList
// 获取软件关联进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetAppProcessList(request *DescribeAssetAppProcessListRequest) (response *DescribeAssetAppProcessListResponse, err error) {
    return c.DescribeAssetAppProcessListWithContext(context.Background(), request)
}

// DescribeAssetAppProcessList
// 获取软件关联进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetAppProcessListWithContext(ctx context.Context, request *DescribeAssetAppProcessListRequest) (response *DescribeAssetAppProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetAppProcessListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetAppProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetAppProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetCoreModuleInfoRequest() (request *DescribeAssetCoreModuleInfoRequest) {
    request = &DescribeAssetCoreModuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetCoreModuleInfo")
    
    
    return
}

func NewDescribeAssetCoreModuleInfoResponse() (response *DescribeAssetCoreModuleInfoResponse) {
    response = &DescribeAssetCoreModuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetCoreModuleInfo
// 获取内核模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetCoreModuleInfo(request *DescribeAssetCoreModuleInfoRequest) (response *DescribeAssetCoreModuleInfoResponse, err error) {
    return c.DescribeAssetCoreModuleInfoWithContext(context.Background(), request)
}

// DescribeAssetCoreModuleInfo
// 获取内核模块详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetCoreModuleInfoWithContext(ctx context.Context, request *DescribeAssetCoreModuleInfoRequest) (response *DescribeAssetCoreModuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetCoreModuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetCoreModuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetCoreModuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetCoreModuleListRequest() (request *DescribeAssetCoreModuleListRequest) {
    request = &DescribeAssetCoreModuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetCoreModuleList")
    
    
    return
}

func NewDescribeAssetCoreModuleListResponse() (response *DescribeAssetCoreModuleListResponse) {
    response = &DescribeAssetCoreModuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetCoreModuleList
// 查询资产管理内核模块列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetCoreModuleList(request *DescribeAssetCoreModuleListRequest) (response *DescribeAssetCoreModuleListResponse, err error) {
    return c.DescribeAssetCoreModuleListWithContext(context.Background(), request)
}

// DescribeAssetCoreModuleList
// 查询资产管理内核模块列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetCoreModuleListWithContext(ctx context.Context, request *DescribeAssetCoreModuleListRequest) (response *DescribeAssetCoreModuleListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetCoreModuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetCoreModuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetCoreModuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDatabaseCountRequest() (request *DescribeAssetDatabaseCountRequest) {
    request = &DescribeAssetDatabaseCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseCount")
    
    
    return
}

func NewDescribeAssetDatabaseCountResponse() (response *DescribeAssetDatabaseCountResponse) {
    response = &DescribeAssetDatabaseCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDatabaseCount
// 获取所有数据库数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDatabaseCount(request *DescribeAssetDatabaseCountRequest) (response *DescribeAssetDatabaseCountResponse, err error) {
    return c.DescribeAssetDatabaseCountWithContext(context.Background(), request)
}

// DescribeAssetDatabaseCount
// 获取所有数据库数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDatabaseCountWithContext(ctx context.Context, request *DescribeAssetDatabaseCountRequest) (response *DescribeAssetDatabaseCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDatabaseCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDatabaseCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDatabaseCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDatabaseInfoRequest() (request *DescribeAssetDatabaseInfoRequest) {
    request = &DescribeAssetDatabaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseInfo")
    
    
    return
}

func NewDescribeAssetDatabaseInfoResponse() (response *DescribeAssetDatabaseInfoResponse) {
    response = &DescribeAssetDatabaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDatabaseInfo
// 获取资产管理数据库详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDatabaseInfo(request *DescribeAssetDatabaseInfoRequest) (response *DescribeAssetDatabaseInfoResponse, err error) {
    return c.DescribeAssetDatabaseInfoWithContext(context.Background(), request)
}

// DescribeAssetDatabaseInfo
// 获取资产管理数据库详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDatabaseInfoWithContext(ctx context.Context, request *DescribeAssetDatabaseInfoRequest) (response *DescribeAssetDatabaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDatabaseInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDatabaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDatabaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDatabaseListRequest() (request *DescribeAssetDatabaseListRequest) {
    request = &DescribeAssetDatabaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDatabaseList")
    
    
    return
}

func NewDescribeAssetDatabaseListResponse() (response *DescribeAssetDatabaseListResponse) {
    response = &DescribeAssetDatabaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDatabaseList
// 查询资产管理数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDatabaseList(request *DescribeAssetDatabaseListRequest) (response *DescribeAssetDatabaseListResponse, err error) {
    return c.DescribeAssetDatabaseListWithContext(context.Background(), request)
}

// DescribeAssetDatabaseList
// 查询资产管理数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDatabaseListWithContext(ctx context.Context, request *DescribeAssetDatabaseListRequest) (response *DescribeAssetDatabaseListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDatabaseListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDatabaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDatabaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDiskListRequest() (request *DescribeAssetDiskListRequest) {
    request = &DescribeAssetDiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetDiskList")
    
    
    return
}

func NewDescribeAssetDiskListResponse() (response *DescribeAssetDiskListResponse) {
    response = &DescribeAssetDiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDiskList
// 获取主机磁盘分区列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDiskList(request *DescribeAssetDiskListRequest) (response *DescribeAssetDiskListResponse, err error) {
    return c.DescribeAssetDiskListWithContext(context.Background(), request)
}

// DescribeAssetDiskList
// 获取主机磁盘分区列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetDiskListWithContext(ctx context.Context, request *DescribeAssetDiskListRequest) (response *DescribeAssetDiskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDiskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetEnvListRequest() (request *DescribeAssetEnvListRequest) {
    request = &DescribeAssetEnvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetEnvList")
    
    
    return
}

func NewDescribeAssetEnvListResponse() (response *DescribeAssetEnvListResponse) {
    response = &DescribeAssetEnvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetEnvList
// 查询资产管理环境变量列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetEnvList(request *DescribeAssetEnvListRequest) (response *DescribeAssetEnvListResponse, err error) {
    return c.DescribeAssetEnvListWithContext(context.Background(), request)
}

// DescribeAssetEnvList
// 查询资产管理环境变量列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetEnvListWithContext(ctx context.Context, request *DescribeAssetEnvListRequest) (response *DescribeAssetEnvListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetEnvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetEnvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetEnvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetHostTotalCountRequest() (request *DescribeAssetHostTotalCountRequest) {
    request = &DescribeAssetHostTotalCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetHostTotalCount")
    
    
    return
}

func NewDescribeAssetHostTotalCountResponse() (response *DescribeAssetHostTotalCountResponse) {
    response = &DescribeAssetHostTotalCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetHostTotalCount
// 获取主机所有资源数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetHostTotalCount(request *DescribeAssetHostTotalCountRequest) (response *DescribeAssetHostTotalCountResponse, err error) {
    return c.DescribeAssetHostTotalCountWithContext(context.Background(), request)
}

// DescribeAssetHostTotalCount
// 获取主机所有资源数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetHostTotalCountWithContext(ctx context.Context, request *DescribeAssetHostTotalCountRequest) (response *DescribeAssetHostTotalCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetHostTotalCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetHostTotalCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetHostTotalCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetInfoRequest() (request *DescribeAssetInfoRequest) {
    request = &DescribeAssetInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetInfo")
    
    
    return
}

func NewDescribeAssetInfoResponse() (response *DescribeAssetInfoResponse) {
    response = &DescribeAssetInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetInfo
// 获取资产数量： 主机数、账号数、端口数、进程数、软件数、数据库数、Web应用数、Web框架数、Web服务数、Web站点数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetInfo(request *DescribeAssetInfoRequest) (response *DescribeAssetInfoResponse, err error) {
    return c.DescribeAssetInfoWithContext(context.Background(), request)
}

// DescribeAssetInfo
// 获取资产数量： 主机数、账号数、端口数、进程数、软件数、数据库数、Web应用数、Web框架数、Web服务数、Web站点数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetInfoWithContext(ctx context.Context, request *DescribeAssetInfoRequest) (response *DescribeAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetInitServiceListRequest() (request *DescribeAssetInitServiceListRequest) {
    request = &DescribeAssetInitServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetInitServiceList")
    
    
    return
}

func NewDescribeAssetInitServiceListResponse() (response *DescribeAssetInitServiceListResponse) {
    response = &DescribeAssetInitServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetInitServiceList
// 查询资产管理启动服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetInitServiceList(request *DescribeAssetInitServiceListRequest) (response *DescribeAssetInitServiceListResponse, err error) {
    return c.DescribeAssetInitServiceListWithContext(context.Background(), request)
}

// DescribeAssetInitServiceList
// 查询资产管理启动服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetInitServiceListWithContext(ctx context.Context, request *DescribeAssetInitServiceListRequest) (response *DescribeAssetInitServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetInitServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetInitServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetInitServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetJarInfoRequest() (request *DescribeAssetJarInfoRequest) {
    request = &DescribeAssetJarInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetJarInfo")
    
    
    return
}

func NewDescribeAssetJarInfoResponse() (response *DescribeAssetJarInfoResponse) {
    response = &DescribeAssetJarInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetJarInfo
// 获取Jar包详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetJarInfo(request *DescribeAssetJarInfoRequest) (response *DescribeAssetJarInfoResponse, err error) {
    return c.DescribeAssetJarInfoWithContext(context.Background(), request)
}

// DescribeAssetJarInfo
// 获取Jar包详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetJarInfoWithContext(ctx context.Context, request *DescribeAssetJarInfoRequest) (response *DescribeAssetJarInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetJarInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetJarInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetJarInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetJarListRequest() (request *DescribeAssetJarListRequest) {
    request = &DescribeAssetJarListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetJarList")
    
    
    return
}

func NewDescribeAssetJarListResponse() (response *DescribeAssetJarListResponse) {
    response = &DescribeAssetJarListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetJarList
// 查询Jar包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetJarList(request *DescribeAssetJarListRequest) (response *DescribeAssetJarListResponse, err error) {
    return c.DescribeAssetJarListWithContext(context.Background(), request)
}

// DescribeAssetJarList
// 查询Jar包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetJarListWithContext(ctx context.Context, request *DescribeAssetJarListRequest) (response *DescribeAssetJarListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetJarListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetJarList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetJarListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetLoadInfoRequest() (request *DescribeAssetLoadInfoRequest) {
    request = &DescribeAssetLoadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetLoadInfo")
    
    
    return
}

func NewDescribeAssetLoadInfoResponse() (response *DescribeAssetLoadInfoResponse) {
    response = &DescribeAssetLoadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetLoadInfo
// 获取系统负载、内存使用率、硬盘使用率情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetLoadInfo(request *DescribeAssetLoadInfoRequest) (response *DescribeAssetLoadInfoResponse, err error) {
    return c.DescribeAssetLoadInfoWithContext(context.Background(), request)
}

// DescribeAssetLoadInfo
// 获取系统负载、内存使用率、硬盘使用率情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetLoadInfoWithContext(ctx context.Context, request *DescribeAssetLoadInfoRequest) (response *DescribeAssetLoadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetLoadInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetLoadInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetLoadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetMachineDetailRequest() (request *DescribeAssetMachineDetailRequest) {
    request = &DescribeAssetMachineDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineDetail")
    
    
    return
}

func NewDescribeAssetMachineDetailResponse() (response *DescribeAssetMachineDetailResponse) {
    response = &DescribeAssetMachineDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetMachineDetail
// 获取资产管理主机资源详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineDetail(request *DescribeAssetMachineDetailRequest) (response *DescribeAssetMachineDetailResponse, err error) {
    return c.DescribeAssetMachineDetailWithContext(context.Background(), request)
}

// DescribeAssetMachineDetail
// 获取资产管理主机资源详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineDetailWithContext(ctx context.Context, request *DescribeAssetMachineDetailRequest) (response *DescribeAssetMachineDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetMachineDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetMachineDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetMachineDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetMachineListRequest() (request *DescribeAssetMachineListRequest) {
    request = &DescribeAssetMachineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineList")
    
    
    return
}

func NewDescribeAssetMachineListResponse() (response *DescribeAssetMachineListResponse) {
    response = &DescribeAssetMachineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetMachineList
// 获取资产指纹页面的资源监控列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineList(request *DescribeAssetMachineListRequest) (response *DescribeAssetMachineListResponse, err error) {
    return c.DescribeAssetMachineListWithContext(context.Background(), request)
}

// DescribeAssetMachineList
// 获取资产指纹页面的资源监控列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineListWithContext(ctx context.Context, request *DescribeAssetMachineListRequest) (response *DescribeAssetMachineListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetMachineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetMachineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetMachineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetMachineTagTopRequest() (request *DescribeAssetMachineTagTopRequest) {
    request = &DescribeAssetMachineTagTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetMachineTagTop")
    
    
    return
}

func NewDescribeAssetMachineTagTopResponse() (response *DescribeAssetMachineTagTopResponse) {
    response = &DescribeAssetMachineTagTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetMachineTagTop
// 获取主机标签Top5
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineTagTop(request *DescribeAssetMachineTagTopRequest) (response *DescribeAssetMachineTagTopResponse, err error) {
    return c.DescribeAssetMachineTagTopWithContext(context.Background(), request)
}

// DescribeAssetMachineTagTop
// 获取主机标签Top5
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetMachineTagTopWithContext(ctx context.Context, request *DescribeAssetMachineTagTopRequest) (response *DescribeAssetMachineTagTopResponse, err error) {
    if request == nil {
        request = NewDescribeAssetMachineTagTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetMachineTagTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetMachineTagTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetPlanTaskListRequest() (request *DescribeAssetPlanTaskListRequest) {
    request = &DescribeAssetPlanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPlanTaskList")
    
    
    return
}

func NewDescribeAssetPlanTaskListResponse() (response *DescribeAssetPlanTaskListResponse) {
    response = &DescribeAssetPlanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetPlanTaskList
// 查询资产管理计划任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPlanTaskList(request *DescribeAssetPlanTaskListRequest) (response *DescribeAssetPlanTaskListResponse, err error) {
    return c.DescribeAssetPlanTaskListWithContext(context.Background(), request)
}

// DescribeAssetPlanTaskList
// 查询资产管理计划任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetPlanTaskListWithContext(ctx context.Context, request *DescribeAssetPlanTaskListRequest) (response *DescribeAssetPlanTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetPlanTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetPlanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetPlanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetPortCountRequest() (request *DescribeAssetPortCountRequest) {
    request = &DescribeAssetPortCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPortCount")
    
    
    return
}

func NewDescribeAssetPortCountResponse() (response *DescribeAssetPortCountResponse) {
    response = &DescribeAssetPortCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetPortCount
// 获取所有端口数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetPortCount(request *DescribeAssetPortCountRequest) (response *DescribeAssetPortCountResponse, err error) {
    return c.DescribeAssetPortCountWithContext(context.Background(), request)
}

// DescribeAssetPortCount
// 获取所有端口数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetPortCountWithContext(ctx context.Context, request *DescribeAssetPortCountRequest) (response *DescribeAssetPortCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetPortCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetPortCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetPortCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetPortInfoListRequest() (request *DescribeAssetPortInfoListRequest) {
    request = &DescribeAssetPortInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetPortInfoList")
    
    
    return
}

func NewDescribeAssetPortInfoListResponse() (response *DescribeAssetPortInfoListResponse) {
    response = &DescribeAssetPortInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetPortInfoList
// 获取资产管理端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetPortInfoList(request *DescribeAssetPortInfoListRequest) (response *DescribeAssetPortInfoListResponse, err error) {
    return c.DescribeAssetPortInfoListWithContext(context.Background(), request)
}

// DescribeAssetPortInfoList
// 获取资产管理端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetPortInfoListWithContext(ctx context.Context, request *DescribeAssetPortInfoListRequest) (response *DescribeAssetPortInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetPortInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetPortInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetPortInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessCountRequest() (request *DescribeAssetProcessCountRequest) {
    request = &DescribeAssetProcessCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetProcessCount")
    
    
    return
}

func NewDescribeAssetProcessCountResponse() (response *DescribeAssetProcessCountResponse) {
    response = &DescribeAssetProcessCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetProcessCount
// 获取所有进程数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetProcessCount(request *DescribeAssetProcessCountRequest) (response *DescribeAssetProcessCountResponse, err error) {
    return c.DescribeAssetProcessCountWithContext(context.Background(), request)
}

// DescribeAssetProcessCount
// 获取所有进程数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetProcessCountWithContext(ctx context.Context, request *DescribeAssetProcessCountRequest) (response *DescribeAssetProcessCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetProcessInfoListRequest() (request *DescribeAssetProcessInfoListRequest) {
    request = &DescribeAssetProcessInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetProcessInfoList")
    
    
    return
}

func NewDescribeAssetProcessInfoListResponse() (response *DescribeAssetProcessInfoListResponse) {
    response = &DescribeAssetProcessInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetProcessInfoList
// 获取资产管理进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetProcessInfoList(request *DescribeAssetProcessInfoListRequest) (response *DescribeAssetProcessInfoListResponse, err error) {
    return c.DescribeAssetProcessInfoListWithContext(context.Background(), request)
}

// DescribeAssetProcessInfoList
// 获取资产管理进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetProcessInfoListWithContext(ctx context.Context, request *DescribeAssetProcessInfoListRequest) (response *DescribeAssetProcessInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetProcessInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetProcessInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetProcessInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetRecentMachineInfoRequest() (request *DescribeAssetRecentMachineInfoRequest) {
    request = &DescribeAssetRecentMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetRecentMachineInfo")
    
    
    return
}

func NewDescribeAssetRecentMachineInfoResponse() (response *DescribeAssetRecentMachineInfoResponse) {
    response = &DescribeAssetRecentMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetRecentMachineInfo
// 获取主机最近趋势情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAssetRecentMachineInfo(request *DescribeAssetRecentMachineInfoRequest) (response *DescribeAssetRecentMachineInfoResponse, err error) {
    return c.DescribeAssetRecentMachineInfoWithContext(context.Background(), request)
}

// DescribeAssetRecentMachineInfo
// 获取主机最近趋势情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAssetRecentMachineInfoWithContext(ctx context.Context, request *DescribeAssetRecentMachineInfoRequest) (response *DescribeAssetRecentMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRecentMachineInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetRecentMachineInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetRecentMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSystemPackageListRequest() (request *DescribeAssetSystemPackageListRequest) {
    request = &DescribeAssetSystemPackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetSystemPackageList")
    
    
    return
}

func NewDescribeAssetSystemPackageListResponse() (response *DescribeAssetSystemPackageListResponse) {
    response = &DescribeAssetSystemPackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetSystemPackageList
// 获取资产管理系统安装包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetSystemPackageList(request *DescribeAssetSystemPackageListRequest) (response *DescribeAssetSystemPackageListResponse, err error) {
    return c.DescribeAssetSystemPackageListWithContext(context.Background(), request)
}

// DescribeAssetSystemPackageList
// 获取资产管理系统安装包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetSystemPackageListWithContext(ctx context.Context, request *DescribeAssetSystemPackageListRequest) (response *DescribeAssetSystemPackageListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSystemPackageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetSystemPackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetSystemPackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetTotalCountRequest() (request *DescribeAssetTotalCountRequest) {
    request = &DescribeAssetTotalCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTotalCount")
    
    
    return
}

func NewDescribeAssetTotalCountResponse() (response *DescribeAssetTotalCountResponse) {
    response = &DescribeAssetTotalCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetTotalCount
// 获取所有资源数量：主机、账号、端口、进程、软件、数据库、Web应用、Web框架、Web服务、Web站点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetTotalCount(request *DescribeAssetTotalCountRequest) (response *DescribeAssetTotalCountResponse, err error) {
    return c.DescribeAssetTotalCountWithContext(context.Background(), request)
}

// DescribeAssetTotalCount
// 获取所有资源数量：主机、账号、端口、进程、软件、数据库、Web应用、Web框架、Web服务、Web站点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetTotalCountWithContext(ctx context.Context, request *DescribeAssetTotalCountRequest) (response *DescribeAssetTotalCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetTotalCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetTotalCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetTotalCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetTypeTopRequest() (request *DescribeAssetTypeTopRequest) {
    request = &DescribeAssetTypeTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTypeTop")
    
    
    return
}

func NewDescribeAssetTypeTopResponse() (response *DescribeAssetTypeTopResponse) {
    response = &DescribeAssetTypeTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetTypeTop
// 获取各种类型资源Top5
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetTypeTop(request *DescribeAssetTypeTopRequest) (response *DescribeAssetTypeTopResponse, err error) {
    return c.DescribeAssetTypeTopWithContext(context.Background(), request)
}

// DescribeAssetTypeTop
// 获取各种类型资源Top5
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetTypeTopWithContext(ctx context.Context, request *DescribeAssetTypeTopRequest) (response *DescribeAssetTypeTopResponse, err error) {
    if request == nil {
        request = NewDescribeAssetTypeTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetTypeTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetTypeTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetTypesRequest() (request *DescribeAssetTypesRequest) {
    request = &DescribeAssetTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetTypes")
    
    
    return
}

func NewDescribeAssetTypesResponse() (response *DescribeAssetTypesResponse) {
    response = &DescribeAssetTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetTypes
// 获取资产指纹类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeAssetTypes(request *DescribeAssetTypesRequest) (response *DescribeAssetTypesResponse, err error) {
    return c.DescribeAssetTypesWithContext(context.Background(), request)
}

// DescribeAssetTypes
// 获取资产指纹类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeAssetTypesWithContext(ctx context.Context, request *DescribeAssetTypesRequest) (response *DescribeAssetTypesResponse, err error) {
    if request == nil {
        request = NewDescribeAssetTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetUserCountRequest() (request *DescribeAssetUserCountRequest) {
    request = &DescribeAssetUserCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserCount")
    
    
    return
}

func NewDescribeAssetUserCountResponse() (response *DescribeAssetUserCountResponse) {
    response = &DescribeAssetUserCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetUserCount
// 获取所有账号数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetUserCount(request *DescribeAssetUserCountRequest) (response *DescribeAssetUserCountResponse, err error) {
    return c.DescribeAssetUserCountWithContext(context.Background(), request)
}

// DescribeAssetUserCount
// 获取所有账号数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetUserCountWithContext(ctx context.Context, request *DescribeAssetUserCountRequest) (response *DescribeAssetUserCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetUserCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetUserCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetUserCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetUserInfoRequest() (request *DescribeAssetUserInfoRequest) {
    request = &DescribeAssetUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserInfo")
    
    
    return
}

func NewDescribeAssetUserInfoResponse() (response *DescribeAssetUserInfoResponse) {
    response = &DescribeAssetUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetUserInfo
// 获取主机账号详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetUserInfo(request *DescribeAssetUserInfoRequest) (response *DescribeAssetUserInfoResponse, err error) {
    return c.DescribeAssetUserInfoWithContext(context.Background(), request)
}

// DescribeAssetUserInfo
// 获取主机账号详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetUserInfoWithContext(ctx context.Context, request *DescribeAssetUserInfoRequest) (response *DescribeAssetUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetUserKeyListRequest() (request *DescribeAssetUserKeyListRequest) {
    request = &DescribeAssetUserKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserKeyList")
    
    
    return
}

func NewDescribeAssetUserKeyListResponse() (response *DescribeAssetUserKeyListResponse) {
    response = &DescribeAssetUserKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetUserKeyList
// 获取主机账号Key列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetUserKeyList(request *DescribeAssetUserKeyListRequest) (response *DescribeAssetUserKeyListResponse, err error) {
    return c.DescribeAssetUserKeyListWithContext(context.Background(), request)
}

// DescribeAssetUserKeyList
// 获取主机账号Key列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetUserKeyListWithContext(ctx context.Context, request *DescribeAssetUserKeyListRequest) (response *DescribeAssetUserKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetUserKeyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetUserKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetUserKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetUserListRequest() (request *DescribeAssetUserListRequest) {
    request = &DescribeAssetUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetUserList")
    
    
    return
}

func NewDescribeAssetUserListResponse() (response *DescribeAssetUserListResponse) {
    response = &DescribeAssetUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetUserList
// 获取账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAssetUserList(request *DescribeAssetUserListRequest) (response *DescribeAssetUserListResponse, err error) {
    return c.DescribeAssetUserListWithContext(context.Background(), request)
}

// DescribeAssetUserList
// 获取账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAssetUserListWithContext(ctx context.Context, request *DescribeAssetUserListRequest) (response *DescribeAssetUserListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebAppCountRequest() (request *DescribeAssetWebAppCountRequest) {
    request = &DescribeAssetWebAppCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppCount")
    
    
    return
}

func NewDescribeAssetWebAppCountResponse() (response *DescribeAssetWebAppCountResponse) {
    response = &DescribeAssetWebAppCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebAppCount
// 获取所有Web应用数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppCount(request *DescribeAssetWebAppCountRequest) (response *DescribeAssetWebAppCountResponse, err error) {
    return c.DescribeAssetWebAppCountWithContext(context.Background(), request)
}

// DescribeAssetWebAppCount
// 获取所有Web应用数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppCountWithContext(ctx context.Context, request *DescribeAssetWebAppCountRequest) (response *DescribeAssetWebAppCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebAppCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebAppCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebAppCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebAppListRequest() (request *DescribeAssetWebAppListRequest) {
    request = &DescribeAssetWebAppListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppList")
    
    
    return
}

func NewDescribeAssetWebAppListResponse() (response *DescribeAssetWebAppListResponse) {
    response = &DescribeAssetWebAppListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebAppList
// 获取资产管理Web应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppList(request *DescribeAssetWebAppListRequest) (response *DescribeAssetWebAppListResponse, err error) {
    return c.DescribeAssetWebAppListWithContext(context.Background(), request)
}

// DescribeAssetWebAppList
// 获取资产管理Web应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppListWithContext(ctx context.Context, request *DescribeAssetWebAppListRequest) (response *DescribeAssetWebAppListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebAppListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebAppList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebAppListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebAppPluginListRequest() (request *DescribeAssetWebAppPluginListRequest) {
    request = &DescribeAssetWebAppPluginListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebAppPluginList")
    
    
    return
}

func NewDescribeAssetWebAppPluginListResponse() (response *DescribeAssetWebAppPluginListResponse) {
    response = &DescribeAssetWebAppPluginListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebAppPluginList
// 获取资产管理Web应用插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppPluginList(request *DescribeAssetWebAppPluginListRequest) (response *DescribeAssetWebAppPluginListResponse, err error) {
    return c.DescribeAssetWebAppPluginListWithContext(context.Background(), request)
}

// DescribeAssetWebAppPluginList
// 获取资产管理Web应用插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebAppPluginListWithContext(ctx context.Context, request *DescribeAssetWebAppPluginListRequest) (response *DescribeAssetWebAppPluginListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebAppPluginListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebAppPluginList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebAppPluginListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebFrameCountRequest() (request *DescribeAssetWebFrameCountRequest) {
    request = &DescribeAssetWebFrameCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebFrameCount")
    
    
    return
}

func NewDescribeAssetWebFrameCountResponse() (response *DescribeAssetWebFrameCountResponse) {
    response = &DescribeAssetWebFrameCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebFrameCount
// 获取所有Web框架数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebFrameCount(request *DescribeAssetWebFrameCountRequest) (response *DescribeAssetWebFrameCountResponse, err error) {
    return c.DescribeAssetWebFrameCountWithContext(context.Background(), request)
}

// DescribeAssetWebFrameCount
// 获取所有Web框架数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebFrameCountWithContext(ctx context.Context, request *DescribeAssetWebFrameCountRequest) (response *DescribeAssetWebFrameCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebFrameCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebFrameCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebFrameCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebFrameListRequest() (request *DescribeAssetWebFrameListRequest) {
    request = &DescribeAssetWebFrameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebFrameList")
    
    
    return
}

func NewDescribeAssetWebFrameListResponse() (response *DescribeAssetWebFrameListResponse) {
    response = &DescribeAssetWebFrameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebFrameList
// 获取资产管理Web框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebFrameList(request *DescribeAssetWebFrameListRequest) (response *DescribeAssetWebFrameListResponse, err error) {
    return c.DescribeAssetWebFrameListWithContext(context.Background(), request)
}

// DescribeAssetWebFrameList
// 获取资产管理Web框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebFrameListWithContext(ctx context.Context, request *DescribeAssetWebFrameListRequest) (response *DescribeAssetWebFrameListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebFrameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebFrameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebFrameListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebLocationCountRequest() (request *DescribeAssetWebLocationCountRequest) {
    request = &DescribeAssetWebLocationCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationCount")
    
    
    return
}

func NewDescribeAssetWebLocationCountResponse() (response *DescribeAssetWebLocationCountResponse) {
    response = &DescribeAssetWebLocationCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebLocationCount
// 获取所有Web站点数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationCount(request *DescribeAssetWebLocationCountRequest) (response *DescribeAssetWebLocationCountResponse, err error) {
    return c.DescribeAssetWebLocationCountWithContext(context.Background(), request)
}

// DescribeAssetWebLocationCount
// 获取所有Web站点数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationCountWithContext(ctx context.Context, request *DescribeAssetWebLocationCountRequest) (response *DescribeAssetWebLocationCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebLocationCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebLocationCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebLocationCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebLocationInfoRequest() (request *DescribeAssetWebLocationInfoRequest) {
    request = &DescribeAssetWebLocationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationInfo")
    
    
    return
}

func NewDescribeAssetWebLocationInfoResponse() (response *DescribeAssetWebLocationInfoResponse) {
    response = &DescribeAssetWebLocationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebLocationInfo
// 获取Web站点详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationInfo(request *DescribeAssetWebLocationInfoRequest) (response *DescribeAssetWebLocationInfoResponse, err error) {
    return c.DescribeAssetWebLocationInfoWithContext(context.Background(), request)
}

// DescribeAssetWebLocationInfo
// 获取Web站点详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationInfoWithContext(ctx context.Context, request *DescribeAssetWebLocationInfoRequest) (response *DescribeAssetWebLocationInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebLocationInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebLocationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebLocationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebLocationListRequest() (request *DescribeAssetWebLocationListRequest) {
    request = &DescribeAssetWebLocationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationList")
    
    
    return
}

func NewDescribeAssetWebLocationListResponse() (response *DescribeAssetWebLocationListResponse) {
    response = &DescribeAssetWebLocationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebLocationList
// 获取Web站点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationList(request *DescribeAssetWebLocationListRequest) (response *DescribeAssetWebLocationListResponse, err error) {
    return c.DescribeAssetWebLocationListWithContext(context.Background(), request)
}

// DescribeAssetWebLocationList
// 获取Web站点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationListWithContext(ctx context.Context, request *DescribeAssetWebLocationListRequest) (response *DescribeAssetWebLocationListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebLocationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebLocationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebLocationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebLocationPathListRequest() (request *DescribeAssetWebLocationPathListRequest) {
    request = &DescribeAssetWebLocationPathListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebLocationPathList")
    
    
    return
}

func NewDescribeAssetWebLocationPathListResponse() (response *DescribeAssetWebLocationPathListResponse) {
    response = &DescribeAssetWebLocationPathListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebLocationPathList
// 获取Web站点虚拟目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationPathList(request *DescribeAssetWebLocationPathListRequest) (response *DescribeAssetWebLocationPathListResponse, err error) {
    return c.DescribeAssetWebLocationPathListWithContext(context.Background(), request)
}

// DescribeAssetWebLocationPathList
// 获取Web站点虚拟目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebLocationPathListWithContext(ctx context.Context, request *DescribeAssetWebLocationPathListRequest) (response *DescribeAssetWebLocationPathListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebLocationPathListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebLocationPathList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebLocationPathListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebServiceCountRequest() (request *DescribeAssetWebServiceCountRequest) {
    request = &DescribeAssetWebServiceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceCount")
    
    
    return
}

func NewDescribeAssetWebServiceCountResponse() (response *DescribeAssetWebServiceCountResponse) {
    response = &DescribeAssetWebServiceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebServiceCount
// 获取所有Web服务数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebServiceCount(request *DescribeAssetWebServiceCountRequest) (response *DescribeAssetWebServiceCountResponse, err error) {
    return c.DescribeAssetWebServiceCountWithContext(context.Background(), request)
}

// DescribeAssetWebServiceCount
// 获取所有Web服务数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebServiceCountWithContext(ctx context.Context, request *DescribeAssetWebServiceCountRequest) (response *DescribeAssetWebServiceCountResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebServiceCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebServiceCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebServiceCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebServiceInfoListRequest() (request *DescribeAssetWebServiceInfoListRequest) {
    request = &DescribeAssetWebServiceInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceInfoList")
    
    
    return
}

func NewDescribeAssetWebServiceInfoListResponse() (response *DescribeAssetWebServiceInfoListResponse) {
    response = &DescribeAssetWebServiceInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebServiceInfoList
// 查询资产管理Web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceInfoList(request *DescribeAssetWebServiceInfoListRequest) (response *DescribeAssetWebServiceInfoListResponse, err error) {
    return c.DescribeAssetWebServiceInfoListWithContext(context.Background(), request)
}

// DescribeAssetWebServiceInfoList
// 查询资产管理Web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetWebServiceInfoListWithContext(ctx context.Context, request *DescribeAssetWebServiceInfoListRequest) (response *DescribeAssetWebServiceInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebServiceInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebServiceInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebServiceInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetWebServiceProcessListRequest() (request *DescribeAssetWebServiceProcessListRequest) {
    request = &DescribeAssetWebServiceProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAssetWebServiceProcessList")
    
    
    return
}

func NewDescribeAssetWebServiceProcessListResponse() (response *DescribeAssetWebServiceProcessListResponse) {
    response = &DescribeAssetWebServiceProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetWebServiceProcessList
// 获取Web服务关联进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebServiceProcessList(request *DescribeAssetWebServiceProcessListRequest) (response *DescribeAssetWebServiceProcessListResponse, err error) {
    return c.DescribeAssetWebServiceProcessListWithContext(context.Background(), request)
}

// DescribeAssetWebServiceProcessList
// 获取Web服务关联进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAssetWebServiceProcessListWithContext(ctx context.Context, request *DescribeAssetWebServiceProcessListRequest) (response *DescribeAssetWebServiceProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetWebServiceProcessListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetWebServiceProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetWebServiceProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackEventInfoRequest() (request *DescribeAttackEventInfoRequest) {
    request = &DescribeAttackEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackEventInfo")
    
    
    return
}

func NewDescribeAttackEventInfoResponse() (response *DescribeAttackEventInfoResponse) {
    response = &DescribeAttackEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackEventInfo
// 网络攻击事件详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackEventInfo(request *DescribeAttackEventInfoRequest) (response *DescribeAttackEventInfoResponse, err error) {
    return c.DescribeAttackEventInfoWithContext(context.Background(), request)
}

// DescribeAttackEventInfo
// 网络攻击事件详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackEventInfoWithContext(ctx context.Context, request *DescribeAttackEventInfoRequest) (response *DescribeAttackEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAttackEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackEventsRequest() (request *DescribeAttackEventsRequest) {
    request = &DescribeAttackEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackEvents")
    
    
    return
}

func NewDescribeAttackEventsResponse() (response *DescribeAttackEventsResponse) {
    response = &DescribeAttackEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackEvents
// 按分页形式展示网络攻击检测事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackEvents(request *DescribeAttackEventsRequest) (response *DescribeAttackEventsResponse, err error) {
    return c.DescribeAttackEventsWithContext(context.Background(), request)
}

// DescribeAttackEvents
// 按分页形式展示网络攻击检测事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackEventsWithContext(ctx context.Context, request *DescribeAttackEventsRequest) (response *DescribeAttackEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackLogInfoRequest() (request *DescribeAttackLogInfoRequest) {
    request = &DescribeAttackLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackLogInfo")
    
    
    return
}

func NewDescribeAttackLogInfoResponse() (response *DescribeAttackLogInfoResponse) {
    response = &DescribeAttackLogInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackLogInfo
// 网络攻击日志详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackLogInfo(request *DescribeAttackLogInfoRequest) (response *DescribeAttackLogInfoResponse, err error) {
    return c.DescribeAttackLogInfoWithContext(context.Background(), request)
}

// DescribeAttackLogInfo
// 网络攻击日志详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackLogInfoWithContext(ctx context.Context, request *DescribeAttackLogInfoRequest) (response *DescribeAttackLogInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAttackLogInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackLogInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackLogInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackLogsRequest() (request *DescribeAttackLogsRequest) {
    request = &DescribeAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackLogs")
    
    
    return
}

func NewDescribeAttackLogsResponse() (response *DescribeAttackLogsResponse) {
    response = &DescribeAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackLogs
// 按分页形式展示网络攻击日志列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackLogs(request *DescribeAttackLogsRequest) (response *DescribeAttackLogsResponse, err error) {
    return c.DescribeAttackLogsWithContext(context.Background(), request)
}

// DescribeAttackLogs
// 按分页形式展示网络攻击日志列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackLogsWithContext(ctx context.Context, request *DescribeAttackLogsRequest) (response *DescribeAttackLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackSourceRequest() (request *DescribeAttackSourceRequest) {
    request = &DescribeAttackSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackSource")
    
    
    return
}

func NewDescribeAttackSourceResponse() (response *DescribeAttackSourceResponse) {
    response = &DescribeAttackSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackSource
// 查询攻击溯源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackSource(request *DescribeAttackSourceRequest) (response *DescribeAttackSourceResponse, err error) {
    return c.DescribeAttackSourceWithContext(context.Background(), request)
}

// DescribeAttackSource
// 查询攻击溯源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackSourceWithContext(ctx context.Context, request *DescribeAttackSourceRequest) (response *DescribeAttackSourceResponse, err error) {
    if request == nil {
        request = NewDescribeAttackSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackSourceEventsRequest() (request *DescribeAttackSourceEventsRequest) {
    request = &DescribeAttackSourceEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackSourceEvents")
    
    
    return
}

func NewDescribeAttackSourceEventsResponse() (response *DescribeAttackSourceEventsResponse) {
    response = &DescribeAttackSourceEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackSourceEvents
// 查询攻击溯源事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackSourceEvents(request *DescribeAttackSourceEventsRequest) (response *DescribeAttackSourceEventsResponse, err error) {
    return c.DescribeAttackSourceEventsWithContext(context.Background(), request)
}

// DescribeAttackSourceEvents
// 查询攻击溯源事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackSourceEventsWithContext(ctx context.Context, request *DescribeAttackSourceEventsRequest) (response *DescribeAttackSourceEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackSourceEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackSourceEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackSourceEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackStatisticsRequest() (request *DescribeAttackStatisticsRequest) {
    request = &DescribeAttackStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackStatistics")
    
    
    return
}

func NewDescribeAttackStatisticsResponse() (response *DescribeAttackStatisticsResponse) {
    response = &DescribeAttackStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackStatistics
// 网络攻击数据统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackStatistics(request *DescribeAttackStatisticsRequest) (response *DescribeAttackStatisticsResponse, err error) {
    return c.DescribeAttackStatisticsWithContext(context.Background(), request)
}

// DescribeAttackStatistics
// 网络攻击数据统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackStatisticsWithContext(ctx context.Context, request *DescribeAttackStatisticsRequest) (response *DescribeAttackStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackTopRequest() (request *DescribeAttackTopRequest) {
    request = &DescribeAttackTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackTop")
    
    
    return
}

func NewDescribeAttackTopResponse() (response *DescribeAttackTopResponse) {
    response = &DescribeAttackTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackTop
// 网络攻击top5数据列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackTop(request *DescribeAttackTopRequest) (response *DescribeAttackTopResponse, err error) {
    return c.DescribeAttackTopWithContext(context.Background(), request)
}

// DescribeAttackTop
// 网络攻击top5数据列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackTopWithContext(ctx context.Context, request *DescribeAttackTopRequest) (response *DescribeAttackTopResponse, err error) {
    if request == nil {
        request = NewDescribeAttackTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackTrendsRequest() (request *DescribeAttackTrendsRequest) {
    request = &DescribeAttackTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackTrends")
    
    
    return
}

func NewDescribeAttackTrendsResponse() (response *DescribeAttackTrendsResponse) {
    response = &DescribeAttackTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackTrends
// 网络攻击趋势数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackTrends(request *DescribeAttackTrendsRequest) (response *DescribeAttackTrendsResponse, err error) {
    return c.DescribeAttackTrendsWithContext(context.Background(), request)
}

// DescribeAttackTrends
// 网络攻击趋势数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackTrendsWithContext(ctx context.Context, request *DescribeAttackTrendsRequest) (response *DescribeAttackTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeAttackTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackVulTypeListRequest() (request *DescribeAttackVulTypeListRequest) {
    request = &DescribeAttackVulTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAttackVulTypeList")
    
    
    return
}

func NewDescribeAttackVulTypeListResponse() (response *DescribeAttackVulTypeListResponse) {
    response = &DescribeAttackVulTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAttackVulTypeList
// 获取网络攻击威胁类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAttackVulTypeList(request *DescribeAttackVulTypeListRequest) (response *DescribeAttackVulTypeListResponse, err error) {
    return c.DescribeAttackVulTypeListWithContext(context.Background(), request)
}

// DescribeAttackVulTypeList
// 获取网络攻击威胁类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAttackVulTypeListWithContext(ctx context.Context, request *DescribeAttackVulTypeListRequest) (response *DescribeAttackVulTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeAttackVulTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttackVulTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttackVulTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableExpertServiceDetailRequest() (request *DescribeAvailableExpertServiceDetailRequest) {
    request = &DescribeAvailableExpertServiceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAvailableExpertServiceDetail")
    
    
    return
}

func NewDescribeAvailableExpertServiceDetailResponse() (response *DescribeAvailableExpertServiceDetailResponse) {
    response = &DescribeAvailableExpertServiceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableExpertServiceDetail
// 专家服务-可用订单详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAvailableExpertServiceDetail(request *DescribeAvailableExpertServiceDetailRequest) (response *DescribeAvailableExpertServiceDetailResponse, err error) {
    return c.DescribeAvailableExpertServiceDetailWithContext(context.Background(), request)
}

// DescribeAvailableExpertServiceDetail
// 专家服务-可用订单详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAvailableExpertServiceDetailWithContext(ctx context.Context, request *DescribeAvailableExpertServiceDetailRequest) (response *DescribeAvailableExpertServiceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableExpertServiceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableExpertServiceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableExpertServiceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanModeRequest() (request *DescribeBanModeRequest) {
    request = &DescribeBanModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanMode")
    
    
    return
}

func NewDescribeBanModeResponse() (response *DescribeBanModeResponse) {
    response = &DescribeBanModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBanMode
// 获取爆破阻断模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBanMode(request *DescribeBanModeRequest) (response *DescribeBanModeResponse, err error) {
    return c.DescribeBanModeWithContext(context.Background(), request)
}

// DescribeBanMode
// 获取爆破阻断模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBanModeWithContext(ctx context.Context, request *DescribeBanModeRequest) (response *DescribeBanModeResponse, err error) {
    if request == nil {
        request = NewDescribeBanModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBanMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBanModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanRegionsRequest() (request *DescribeBanRegionsRequest) {
    request = &DescribeBanRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanRegions")
    
    
    return
}

func NewDescribeBanRegionsResponse() (response *DescribeBanRegionsResponse) {
    response = &DescribeBanRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBanRegions
// 获取阻断地域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBanRegions(request *DescribeBanRegionsRequest) (response *DescribeBanRegionsResponse, err error) {
    return c.DescribeBanRegionsWithContext(context.Background(), request)
}

// DescribeBanRegions
// 获取阻断地域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBanRegionsWithContext(ctx context.Context, request *DescribeBanRegionsRequest) (response *DescribeBanRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeBanRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBanRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBanRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanStatusRequest() (request *DescribeBanStatusRequest) {
    request = &DescribeBanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanStatus")
    
    
    return
}

func NewDescribeBanStatusResponse() (response *DescribeBanStatusResponse) {
    response = &DescribeBanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBanStatus
// 获取阻断按钮状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBanStatus(request *DescribeBanStatusRequest) (response *DescribeBanStatusResponse, err error) {
    return c.DescribeBanStatusWithContext(context.Background(), request)
}

// DescribeBanStatus
// 获取阻断按钮状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBanStatusWithContext(ctx context.Context, request *DescribeBanStatusRequest) (response *DescribeBanStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBanStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBanStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBanWhiteListRequest() (request *DescribeBanWhiteListRequest) {
    request = &DescribeBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBanWhiteList")
    
    
    return
}

func NewDescribeBanWhiteListResponse() (response *DescribeBanWhiteListResponse) {
    response = &DescribeBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBanWhiteList
// 获取阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBanWhiteList(request *DescribeBanWhiteListRequest) (response *DescribeBanWhiteListResponse, err error) {
    return c.DescribeBanWhiteListWithContext(context.Background(), request)
}

// DescribeBanWhiteList
// 获取阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBanWhiteListWithContext(ctx context.Context, request *DescribeBanWhiteListRequest) (response *DescribeBanWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeBanWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBanWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineAnalysisDataRequest() (request *DescribeBaselineAnalysisDataRequest) {
    request = &DescribeBaselineAnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineAnalysisData")
    
    
    return
}

func NewDescribeBaselineAnalysisDataResponse() (response *DescribeBaselineAnalysisDataResponse) {
    response = &DescribeBaselineAnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineAnalysisData
// 根据基线策略id查询基线策略数据概览统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineAnalysisData(request *DescribeBaselineAnalysisDataRequest) (response *DescribeBaselineAnalysisDataResponse, err error) {
    return c.DescribeBaselineAnalysisDataWithContext(context.Background(), request)
}

// DescribeBaselineAnalysisData
// 根据基线策略id查询基线策略数据概览统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineAnalysisDataWithContext(ctx context.Context, request *DescribeBaselineAnalysisDataRequest) (response *DescribeBaselineAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineAnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineAnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineAnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineBasicInfoRequest() (request *DescribeBaselineBasicInfoRequest) {
    request = &DescribeBaselineBasicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineBasicInfo")
    
    
    return
}

func NewDescribeBaselineBasicInfoResponse() (response *DescribeBaselineBasicInfoResponse) {
    response = &DescribeBaselineBasicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineBasicInfo
// 查询基线基础信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineBasicInfo(request *DescribeBaselineBasicInfoRequest) (response *DescribeBaselineBasicInfoResponse, err error) {
    return c.DescribeBaselineBasicInfoWithContext(context.Background(), request)
}

// DescribeBaselineBasicInfo
// 查询基线基础信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineBasicInfoWithContext(ctx context.Context, request *DescribeBaselineBasicInfoRequest) (response *DescribeBaselineBasicInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineBasicInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineBasicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineBasicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineDefaultStrategyListRequest() (request *DescribeBaselineDefaultStrategyListRequest) {
    request = &DescribeBaselineDefaultStrategyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDefaultStrategyList")
    
    
    return
}

func NewDescribeBaselineDefaultStrategyListResponse() (response *DescribeBaselineDefaultStrategyListResponse) {
    response = &DescribeBaselineDefaultStrategyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineDefaultStrategyList
// 查询基线默认策略列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBaselineDefaultStrategyList(request *DescribeBaselineDefaultStrategyListRequest) (response *DescribeBaselineDefaultStrategyListResponse, err error) {
    return c.DescribeBaselineDefaultStrategyListWithContext(context.Background(), request)
}

// DescribeBaselineDefaultStrategyList
// 查询基线默认策略列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBaselineDefaultStrategyListWithContext(ctx context.Context, request *DescribeBaselineDefaultStrategyListRequest) (response *DescribeBaselineDefaultStrategyListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineDefaultStrategyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineDefaultStrategyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineDefaultStrategyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineDetailRequest() (request *DescribeBaselineDetailRequest) {
    request = &DescribeBaselineDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetail")
    
    
    return
}

func NewDescribeBaselineDetailResponse() (response *DescribeBaselineDetailResponse) {
    response = &DescribeBaselineDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineDetail
// 根据基线id查询基线详情接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineDetail(request *DescribeBaselineDetailRequest) (response *DescribeBaselineDetailResponse, err error) {
    return c.DescribeBaselineDetailWithContext(context.Background(), request)
}

// DescribeBaselineDetail
// 根据基线id查询基线详情接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineDetailWithContext(ctx context.Context, request *DescribeBaselineDetailRequest) (response *DescribeBaselineDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineDetectListRequest() (request *DescribeBaselineDetectListRequest) {
    request = &DescribeBaselineDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetectList")
    
    
    return
}

func NewDescribeBaselineDetectListResponse() (response *DescribeBaselineDetectListResponse) {
    response = &DescribeBaselineDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineDetectList
// 获取基线检测详情记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDetectList(request *DescribeBaselineDetectListRequest) (response *DescribeBaselineDetectListResponse, err error) {
    return c.DescribeBaselineDetectListWithContext(context.Background(), request)
}

// DescribeBaselineDetectList
// 获取基线检测详情记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDetectListWithContext(ctx context.Context, request *DescribeBaselineDetectListRequest) (response *DescribeBaselineDetectListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineDetectOverviewRequest() (request *DescribeBaselineDetectOverviewRequest) {
    request = &DescribeBaselineDetectOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDetectOverview")
    
    
    return
}

func NewDescribeBaselineDetectOverviewResponse() (response *DescribeBaselineDetectOverviewResponse) {
    response = &DescribeBaselineDetectOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineDetectOverview
// 获取基线检测概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDetectOverview(request *DescribeBaselineDetectOverviewRequest) (response *DescribeBaselineDetectOverviewResponse, err error) {
    return c.DescribeBaselineDetectOverviewWithContext(context.Background(), request)
}

// DescribeBaselineDetectOverview
// 获取基线检测概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDetectOverviewWithContext(ctx context.Context, request *DescribeBaselineDetectOverviewRequest) (response *DescribeBaselineDetectOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineDetectOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineDetectOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineDetectOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineDownloadListRequest() (request *DescribeBaselineDownloadListRequest) {
    request = &DescribeBaselineDownloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineDownloadList")
    
    
    return
}

func NewDescribeBaselineDownloadListResponse() (response *DescribeBaselineDownloadListResponse) {
    response = &DescribeBaselineDownloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineDownloadList
// 获取基线下载列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDownloadList(request *DescribeBaselineDownloadListRequest) (response *DescribeBaselineDownloadListResponse, err error) {
    return c.DescribeBaselineDownloadListWithContext(context.Background(), request)
}

// DescribeBaselineDownloadList
// 获取基线下载列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineDownloadListWithContext(ctx context.Context, request *DescribeBaselineDownloadListRequest) (response *DescribeBaselineDownloadListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineDownloadListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineDownloadList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineDownloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineEffectHostListRequest() (request *DescribeBaselineEffectHostListRequest) {
    request = &DescribeBaselineEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineEffectHostList")
    
    
    return
}

func NewDescribeBaselineEffectHostListResponse() (response *DescribeBaselineEffectHostListResponse) {
    response = &DescribeBaselineEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineEffectHostList
// 根据基线id查询基线影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineEffectHostList(request *DescribeBaselineEffectHostListRequest) (response *DescribeBaselineEffectHostListResponse, err error) {
    return c.DescribeBaselineEffectHostListWithContext(context.Background(), request)
}

// DescribeBaselineEffectHostList
// 根据基线id查询基线影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineEffectHostListWithContext(ctx context.Context, request *DescribeBaselineEffectHostListRequest) (response *DescribeBaselineEffectHostListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineFixListRequest() (request *DescribeBaselineFixListRequest) {
    request = &DescribeBaselineFixListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineFixList")
    
    
    return
}

func NewDescribeBaselineFixListResponse() (response *DescribeBaselineFixListResponse) {
    response = &DescribeBaselineFixListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineFixList
// 获取基线修复列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineFixList(request *DescribeBaselineFixListRequest) (response *DescribeBaselineFixListResponse, err error) {
    return c.DescribeBaselineFixListWithContext(context.Background(), request)
}

// DescribeBaselineFixList
// 获取基线修复列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineFixListWithContext(ctx context.Context, request *DescribeBaselineFixListRequest) (response *DescribeBaselineFixListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineFixListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineFixList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineFixListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineHostDetectListRequest() (request *DescribeBaselineHostDetectListRequest) {
    request = &DescribeBaselineHostDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostDetectList")
    
    
    return
}

func NewDescribeBaselineHostDetectListResponse() (response *DescribeBaselineHostDetectListResponse) {
    response = &DescribeBaselineHostDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineHostDetectList
// 获取基线检测主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostDetectList(request *DescribeBaselineHostDetectListRequest) (response *DescribeBaselineHostDetectListResponse, err error) {
    return c.DescribeBaselineHostDetectListWithContext(context.Background(), request)
}

// DescribeBaselineHostDetectList
// 获取基线检测主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostDetectListWithContext(ctx context.Context, request *DescribeBaselineHostDetectListRequest) (response *DescribeBaselineHostDetectListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineHostDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineHostDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineHostDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineHostIgnoreListRequest() (request *DescribeBaselineHostIgnoreListRequest) {
    request = &DescribeBaselineHostIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostIgnoreList")
    
    
    return
}

func NewDescribeBaselineHostIgnoreListResponse() (response *DescribeBaselineHostIgnoreListResponse) {
    response = &DescribeBaselineHostIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineHostIgnoreList
// 获取忽略规则主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostIgnoreList(request *DescribeBaselineHostIgnoreListRequest) (response *DescribeBaselineHostIgnoreListResponse, err error) {
    return c.DescribeBaselineHostIgnoreListWithContext(context.Background(), request)
}

// DescribeBaselineHostIgnoreList
// 获取忽略规则主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostIgnoreListWithContext(ctx context.Context, request *DescribeBaselineHostIgnoreListRequest) (response *DescribeBaselineHostIgnoreListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineHostIgnoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineHostIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineHostIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineHostRiskTopRequest() (request *DescribeBaselineHostRiskTopRequest) {
    request = &DescribeBaselineHostRiskTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostRiskTop")
    
    
    return
}

func NewDescribeBaselineHostRiskTopResponse() (response *DescribeBaselineHostRiskTopResponse) {
    response = &DescribeBaselineHostRiskTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineHostRiskTop
// 获取基线服务器风险TOP5
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostRiskTop(request *DescribeBaselineHostRiskTopRequest) (response *DescribeBaselineHostRiskTopResponse, err error) {
    return c.DescribeBaselineHostRiskTopWithContext(context.Background(), request)
}

// DescribeBaselineHostRiskTop
// 获取基线服务器风险TOP5
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostRiskTopWithContext(ctx context.Context, request *DescribeBaselineHostRiskTopRequest) (response *DescribeBaselineHostRiskTopResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineHostRiskTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineHostRiskTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineHostRiskTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineHostTopRequest() (request *DescribeBaselineHostTopRequest) {
    request = &DescribeBaselineHostTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineHostTop")
    
    
    return
}

func NewDescribeBaselineHostTopResponse() (response *DescribeBaselineHostTopResponse) {
    response = &DescribeBaselineHostTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineHostTop
// 接口返回TopN的风险服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineHostTop(request *DescribeBaselineHostTopRequest) (response *DescribeBaselineHostTopResponse, err error) {
    return c.DescribeBaselineHostTopWithContext(context.Background(), request)
}

// DescribeBaselineHostTop
// 接口返回TopN的风险服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineHostTopWithContext(ctx context.Context, request *DescribeBaselineHostTopRequest) (response *DescribeBaselineHostTopResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineHostTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineHostTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineHostTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineItemDetectListRequest() (request *DescribeBaselineItemDetectListRequest) {
    request = &DescribeBaselineItemDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemDetectList")
    
    
    return
}

func NewDescribeBaselineItemDetectListResponse() (response *DescribeBaselineItemDetectListResponse) {
    response = &DescribeBaselineItemDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineItemDetectList
// 获取基线检测项的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemDetectList(request *DescribeBaselineItemDetectListRequest) (response *DescribeBaselineItemDetectListResponse, err error) {
    return c.DescribeBaselineItemDetectListWithContext(context.Background(), request)
}

// DescribeBaselineItemDetectList
// 获取基线检测项的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemDetectListWithContext(ctx context.Context, request *DescribeBaselineItemDetectListRequest) (response *DescribeBaselineItemDetectListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineItemDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineItemDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineItemDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineItemIgnoreListRequest() (request *DescribeBaselineItemIgnoreListRequest) {
    request = &DescribeBaselineItemIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemIgnoreList")
    
    
    return
}

func NewDescribeBaselineItemIgnoreListResponse() (response *DescribeBaselineItemIgnoreListResponse) {
    response = &DescribeBaselineItemIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineItemIgnoreList
// 获取忽略规则项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemIgnoreList(request *DescribeBaselineItemIgnoreListRequest) (response *DescribeBaselineItemIgnoreListResponse, err error) {
    return c.DescribeBaselineItemIgnoreListWithContext(context.Background(), request)
}

// DescribeBaselineItemIgnoreList
// 获取忽略规则项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemIgnoreListWithContext(ctx context.Context, request *DescribeBaselineItemIgnoreListRequest) (response *DescribeBaselineItemIgnoreListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineItemIgnoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineItemIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineItemIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineItemInfoRequest() (request *DescribeBaselineItemInfoRequest) {
    request = &DescribeBaselineItemInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemInfo")
    
    
    return
}

func NewDescribeBaselineItemInfoResponse() (response *DescribeBaselineItemInfoResponse) {
    response = &DescribeBaselineItemInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineItemInfo
// 获取基线检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemInfo(request *DescribeBaselineItemInfoRequest) (response *DescribeBaselineItemInfoResponse, err error) {
    return c.DescribeBaselineItemInfoWithContext(context.Background(), request)
}

// DescribeBaselineItemInfo
// 获取基线检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemInfoWithContext(ctx context.Context, request *DescribeBaselineItemInfoRequest) (response *DescribeBaselineItemInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineItemInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineItemInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineItemInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineItemListRequest() (request *DescribeBaselineItemListRequest) {
    request = &DescribeBaselineItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemList")
    
    
    return
}

func NewDescribeBaselineItemListResponse() (response *DescribeBaselineItemListResponse) {
    response = &DescribeBaselineItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineItemList
// 获取基线项检测结果列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemList(request *DescribeBaselineItemListRequest) (response *DescribeBaselineItemListResponse, err error) {
    return c.DescribeBaselineItemListWithContext(context.Background(), request)
}

// DescribeBaselineItemList
// 获取基线项检测结果列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemListWithContext(ctx context.Context, request *DescribeBaselineItemListRequest) (response *DescribeBaselineItemListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineItemRiskTopRequest() (request *DescribeBaselineItemRiskTopRequest) {
    request = &DescribeBaselineItemRiskTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineItemRiskTop")
    
    
    return
}

func NewDescribeBaselineItemRiskTopResponse() (response *DescribeBaselineItemRiskTopResponse) {
    response = &DescribeBaselineItemRiskTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineItemRiskTop
// 获取基线检测项TOP5
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemRiskTop(request *DescribeBaselineItemRiskTopRequest) (response *DescribeBaselineItemRiskTopResponse, err error) {
    return c.DescribeBaselineItemRiskTopWithContext(context.Background(), request)
}

// DescribeBaselineItemRiskTop
// 获取基线检测项TOP5
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemRiskTopWithContext(ctx context.Context, request *DescribeBaselineItemRiskTopRequest) (response *DescribeBaselineItemRiskTopResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineItemRiskTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineItemRiskTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineItemRiskTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineListRequest() (request *DescribeBaselineListRequest) {
    request = &DescribeBaselineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineList")
    
    
    return
}

func NewDescribeBaselineListResponse() (response *DescribeBaselineListResponse) {
    response = &DescribeBaselineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineList
// 查询基线列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineList(request *DescribeBaselineListRequest) (response *DescribeBaselineListResponse, err error) {
    return c.DescribeBaselineListWithContext(context.Background(), request)
}

// DescribeBaselineList
// 查询基线列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineListWithContext(ctx context.Context, request *DescribeBaselineListRequest) (response *DescribeBaselineListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselinePolicyListRequest() (request *DescribeBaselinePolicyListRequest) {
    request = &DescribeBaselinePolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselinePolicyList")
    
    
    return
}

func NewDescribeBaselinePolicyListResponse() (response *DescribeBaselinePolicyListResponse) {
    response = &DescribeBaselinePolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselinePolicyList
// 获取基线策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselinePolicyList(request *DescribeBaselinePolicyListRequest) (response *DescribeBaselinePolicyListResponse, err error) {
    return c.DescribeBaselinePolicyListWithContext(context.Background(), request)
}

// DescribeBaselinePolicyList
// 获取基线策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselinePolicyListWithContext(ctx context.Context, request *DescribeBaselinePolicyListRequest) (response *DescribeBaselinePolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselinePolicyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselinePolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselinePolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineRuleRequest() (request *DescribeBaselineRuleRequest) {
    request = &DescribeBaselineRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRule")
    
    
    return
}

func NewDescribeBaselineRuleResponse() (response *DescribeBaselineRuleResponse) {
    response = &DescribeBaselineRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineRule
// 根据基线id查询下属检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineRule(request *DescribeBaselineRuleRequest) (response *DescribeBaselineRuleResponse, err error) {
    return c.DescribeBaselineRuleWithContext(context.Background(), request)
}

// DescribeBaselineRule
// 根据基线id查询下属检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineRuleWithContext(ctx context.Context, request *DescribeBaselineRuleRequest) (response *DescribeBaselineRuleResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineRuleCategoryListRequest() (request *DescribeBaselineRuleCategoryListRequest) {
    request = &DescribeBaselineRuleCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleCategoryList")
    
    
    return
}

func NewDescribeBaselineRuleCategoryListResponse() (response *DescribeBaselineRuleCategoryListResponse) {
    response = &DescribeBaselineRuleCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineRuleCategoryList
// 获取基线分类列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBaselineRuleCategoryList(request *DescribeBaselineRuleCategoryListRequest) (response *DescribeBaselineRuleCategoryListResponse, err error) {
    return c.DescribeBaselineRuleCategoryListWithContext(context.Background(), request)
}

// DescribeBaselineRuleCategoryList
// 获取基线分类列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBaselineRuleCategoryListWithContext(ctx context.Context, request *DescribeBaselineRuleCategoryListRequest) (response *DescribeBaselineRuleCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineRuleCategoryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineRuleCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineRuleCategoryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineRuleDetectListRequest() (request *DescribeBaselineRuleDetectListRequest) {
    request = &DescribeBaselineRuleDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleDetectList")
    
    
    return
}

func NewDescribeBaselineRuleDetectListResponse() (response *DescribeBaselineRuleDetectListResponse) {
    response = &DescribeBaselineRuleDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineRuleDetectList
// 获取基线规则检测列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleDetectList(request *DescribeBaselineRuleDetectListRequest) (response *DescribeBaselineRuleDetectListResponse, err error) {
    return c.DescribeBaselineRuleDetectListWithContext(context.Background(), request)
}

// DescribeBaselineRuleDetectList
// 获取基线规则检测列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleDetectListWithContext(ctx context.Context, request *DescribeBaselineRuleDetectListRequest) (response *DescribeBaselineRuleDetectListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineRuleDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineRuleDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineRuleDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineRuleIgnoreListRequest() (request *DescribeBaselineRuleIgnoreListRequest) {
    request = &DescribeBaselineRuleIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleIgnoreList")
    
    
    return
}

func NewDescribeBaselineRuleIgnoreListResponse() (response *DescribeBaselineRuleIgnoreListResponse) {
    response = &DescribeBaselineRuleIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineRuleIgnoreList
// 获取基线忽略规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleIgnoreList(request *DescribeBaselineRuleIgnoreListRequest) (response *DescribeBaselineRuleIgnoreListResponse, err error) {
    return c.DescribeBaselineRuleIgnoreListWithContext(context.Background(), request)
}

// DescribeBaselineRuleIgnoreList
// 获取基线忽略规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleIgnoreListWithContext(ctx context.Context, request *DescribeBaselineRuleIgnoreListRequest) (response *DescribeBaselineRuleIgnoreListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineRuleIgnoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineRuleIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineRuleIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineRuleListRequest() (request *DescribeBaselineRuleListRequest) {
    request = &DescribeBaselineRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineRuleList")
    
    
    return
}

func NewDescribeBaselineRuleListResponse() (response *DescribeBaselineRuleListResponse) {
    response = &DescribeBaselineRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineRuleList
// 获取基线规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleList(request *DescribeBaselineRuleListRequest) (response *DescribeBaselineRuleListResponse, err error) {
    return c.DescribeBaselineRuleListWithContext(context.Background(), request)
}

// DescribeBaselineRuleList
// 获取基线规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleListWithContext(ctx context.Context, request *DescribeBaselineRuleListRequest) (response *DescribeBaselineRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineScanScheduleRequest() (request *DescribeBaselineScanScheduleRequest) {
    request = &DescribeBaselineScanScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineScanSchedule")
    
    
    return
}

func NewDescribeBaselineScanScheduleResponse() (response *DescribeBaselineScanScheduleResponse) {
    response = &DescribeBaselineScanScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineScanSchedule
// 根据任务id查询基线检测进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineScanSchedule(request *DescribeBaselineScanScheduleRequest) (response *DescribeBaselineScanScheduleResponse, err error) {
    return c.DescribeBaselineScanScheduleWithContext(context.Background(), request)
}

// DescribeBaselineScanSchedule
// 根据任务id查询基线检测进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineScanScheduleWithContext(ctx context.Context, request *DescribeBaselineScanScheduleRequest) (response *DescribeBaselineScanScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineScanScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineScanSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineScanScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineStrategyDetailRequest() (request *DescribeBaselineStrategyDetailRequest) {
    request = &DescribeBaselineStrategyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineStrategyDetail")
    
    
    return
}

func NewDescribeBaselineStrategyDetailResponse() (response *DescribeBaselineStrategyDetailResponse) {
    response = &DescribeBaselineStrategyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineStrategyDetail
// 根据基线策略id查询策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineStrategyDetail(request *DescribeBaselineStrategyDetailRequest) (response *DescribeBaselineStrategyDetailResponse, err error) {
    return c.DescribeBaselineStrategyDetailWithContext(context.Background(), request)
}

// DescribeBaselineStrategyDetail
// 根据基线策略id查询策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineStrategyDetailWithContext(ctx context.Context, request *DescribeBaselineStrategyDetailRequest) (response *DescribeBaselineStrategyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineStrategyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineStrategyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineStrategyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineStrategyListRequest() (request *DescribeBaselineStrategyListRequest) {
    request = &DescribeBaselineStrategyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineStrategyList")
    
    
    return
}

func NewDescribeBaselineStrategyListResponse() (response *DescribeBaselineStrategyListResponse) {
    response = &DescribeBaselineStrategyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineStrategyList
// 查询一个用户下的基线策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineStrategyList(request *DescribeBaselineStrategyListRequest) (response *DescribeBaselineStrategyListResponse, err error) {
    return c.DescribeBaselineStrategyListWithContext(context.Background(), request)
}

// DescribeBaselineStrategyList
// 查询一个用户下的基线策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineStrategyListWithContext(ctx context.Context, request *DescribeBaselineStrategyListRequest) (response *DescribeBaselineStrategyListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineStrategyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineStrategyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineStrategyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineTopRequest() (request *DescribeBaselineTopRequest) {
    request = &DescribeBaselineTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineTop")
    
    
    return
}

func NewDescribeBaselineTopResponse() (response *DescribeBaselineTopResponse) {
    response = &DescribeBaselineTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineTop
// 根据策略id查询基线检测项TOP
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineTop(request *DescribeBaselineTopRequest) (response *DescribeBaselineTopResponse, err error) {
    return c.DescribeBaselineTopWithContext(context.Background(), request)
}

// DescribeBaselineTop
// 根据策略id查询基线检测项TOP
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBaselineTopWithContext(ctx context.Context, request *DescribeBaselineTopRequest) (response *DescribeBaselineTopResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineWeakPasswordListRequest() (request *DescribeBaselineWeakPasswordListRequest) {
    request = &DescribeBaselineWeakPasswordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBaselineWeakPasswordList")
    
    
    return
}

func NewDescribeBaselineWeakPasswordListResponse() (response *DescribeBaselineWeakPasswordListResponse) {
    response = &DescribeBaselineWeakPasswordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaselineWeakPasswordList
// 获取基线弱口令列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineWeakPasswordList(request *DescribeBaselineWeakPasswordListRequest) (response *DescribeBaselineWeakPasswordListResponse, err error) {
    return c.DescribeBaselineWeakPasswordListWithContext(context.Background(), request)
}

// DescribeBaselineWeakPasswordList
// 获取基线弱口令列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineWeakPasswordListWithContext(ctx context.Context, request *DescribeBaselineWeakPasswordListRequest) (response *DescribeBaselineWeakPasswordListResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineWeakPasswordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineWeakPasswordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineWeakPasswordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashEventsRequest() (request *DescribeBashEventsRequest) {
    request = &DescribeBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEvents")
    
    
    return
}

func NewDescribeBashEventsResponse() (response *DescribeBashEventsResponse) {
    response = &DescribeBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashEvents
// 获取高危命令列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashEvents(request *DescribeBashEventsRequest) (response *DescribeBashEventsResponse, err error) {
    return c.DescribeBashEventsWithContext(context.Background(), request)
}

// DescribeBashEvents
// 获取高危命令列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashEventsWithContext(ctx context.Context, request *DescribeBashEventsRequest) (response *DescribeBashEventsResponse, err error) {
    if request == nil {
        request = NewDescribeBashEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashEventsInfoRequest() (request *DescribeBashEventsInfoRequest) {
    request = &DescribeBashEventsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsInfo")
    
    
    return
}

func NewDescribeBashEventsInfoResponse() (response *DescribeBashEventsInfoResponse) {
    response = &DescribeBashEventsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashEventsInfo
// 查询高危命令事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBashEventsInfo(request *DescribeBashEventsInfoRequest) (response *DescribeBashEventsInfoResponse, err error) {
    return c.DescribeBashEventsInfoWithContext(context.Background(), request)
}

// DescribeBashEventsInfo
// 查询高危命令事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBashEventsInfoWithContext(ctx context.Context, request *DescribeBashEventsInfoRequest) (response *DescribeBashEventsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBashEventsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashEventsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashEventsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashEventsInfoNewRequest() (request *DescribeBashEventsInfoNewRequest) {
    request = &DescribeBashEventsInfoNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsInfoNew")
    
    
    return
}

func NewDescribeBashEventsInfoNewResponse() (response *DescribeBashEventsInfoNewResponse) {
    response = &DescribeBashEventsInfoNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashEventsInfoNew
// 查询高危命令事件详情(新)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBashEventsInfoNew(request *DescribeBashEventsInfoNewRequest) (response *DescribeBashEventsInfoNewResponse, err error) {
    return c.DescribeBashEventsInfoNewWithContext(context.Background(), request)
}

// DescribeBashEventsInfoNew
// 查询高危命令事件详情(新)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeBashEventsInfoNewWithContext(ctx context.Context, request *DescribeBashEventsInfoNewRequest) (response *DescribeBashEventsInfoNewResponse, err error) {
    if request == nil {
        request = NewDescribeBashEventsInfoNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashEventsInfoNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashEventsInfoNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashEventsNewRequest() (request *DescribeBashEventsNewRequest) {
    request = &DescribeBashEventsNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashEventsNew")
    
    
    return
}

func NewDescribeBashEventsNewResponse() (response *DescribeBashEventsNewResponse) {
    response = &DescribeBashEventsNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashEventsNew
// 获取高危命令列表(新)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashEventsNew(request *DescribeBashEventsNewRequest) (response *DescribeBashEventsNewResponse, err error) {
    return c.DescribeBashEventsNewWithContext(context.Background(), request)
}

// DescribeBashEventsNew
// 获取高危命令列表(新)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashEventsNewWithContext(ctx context.Context, request *DescribeBashEventsNewRequest) (response *DescribeBashEventsNewResponse, err error) {
    if request == nil {
        request = NewDescribeBashEventsNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashEventsNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashEventsNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashPoliciesRequest() (request *DescribeBashPoliciesRequest) {
    request = &DescribeBashPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashPolicies")
    
    
    return
}

func NewDescribeBashPoliciesResponse() (response *DescribeBashPoliciesResponse) {
    response = &DescribeBashPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashPolicies
// 获取高危命令策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashPolicies(request *DescribeBashPoliciesRequest) (response *DescribeBashPoliciesResponse, err error) {
    return c.DescribeBashPoliciesWithContext(context.Background(), request)
}

// DescribeBashPolicies
// 获取高危命令策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashPoliciesWithContext(ctx context.Context, request *DescribeBashPoliciesRequest) (response *DescribeBashPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeBashPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBashRulesRequest() (request *DescribeBashRulesRequest) {
    request = &DescribeBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBashRules")
    
    
    return
}

func NewDescribeBashRulesResponse() (response *DescribeBashRulesResponse) {
    response = &DescribeBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBashRules
// 获取高危命令规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashRules(request *DescribeBashRulesRequest) (response *DescribeBashRulesResponse, err error) {
    return c.DescribeBashRulesWithContext(context.Background(), request)
}

// DescribeBashRules
// 获取高危命令规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBashRulesWithContext(ctx context.Context, request *DescribeBashRulesRequest) (response *DescribeBashRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBashRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBashRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttackListRequest() (request *DescribeBruteAttackListRequest) {
    request = &DescribeBruteAttackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBruteAttackList")
    
    
    return
}

func NewDescribeBruteAttackListResponse() (response *DescribeBruteAttackListResponse) {
    response = &DescribeBruteAttackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBruteAttackList
// 获取密码破解列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBruteAttackList(request *DescribeBruteAttackListRequest) (response *DescribeBruteAttackListResponse, err error) {
    return c.DescribeBruteAttackListWithContext(context.Background(), request)
}

// DescribeBruteAttackList
// 获取密码破解列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBruteAttackListWithContext(ctx context.Context, request *DescribeBruteAttackListRequest) (response *DescribeBruteAttackListResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttackListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBruteAttackList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBruteAttackListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttackRulesRequest() (request *DescribeBruteAttackRulesRequest) {
    request = &DescribeBruteAttackRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBruteAttackRules")
    
    
    return
}

func NewDescribeBruteAttackRulesResponse() (response *DescribeBruteAttackRulesResponse) {
    response = &DescribeBruteAttackRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBruteAttackRules
// 获取爆破破解规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBruteAttackRules(request *DescribeBruteAttackRulesRequest) (response *DescribeBruteAttackRulesResponse, err error) {
    return c.DescribeBruteAttackRulesWithContext(context.Background(), request)
}

// DescribeBruteAttackRules
// 获取爆破破解规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBruteAttackRulesWithContext(ctx context.Context, request *DescribeBruteAttackRulesRequest) (response *DescribeBruteAttackRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttackRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBruteAttackRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBruteAttackRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCanFixVulMachineRequest() (request *DescribeCanFixVulMachineRequest) {
    request = &DescribeCanFixVulMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeCanFixVulMachine")
    
    
    return
}

func NewDescribeCanFixVulMachineResponse() (response *DescribeCanFixVulMachineResponse) {
    response = &DescribeCanFixVulMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCanFixVulMachine
// 漏洞修护-查询可修护主机信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCanFixVulMachine(request *DescribeCanFixVulMachineRequest) (response *DescribeCanFixVulMachineResponse, err error) {
    return c.DescribeCanFixVulMachineWithContext(context.Background(), request)
}

// DescribeCanFixVulMachine
// 漏洞修护-查询可修护主机信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCanFixVulMachineWithContext(ctx context.Context, request *DescribeCanFixVulMachineRequest) (response *DescribeCanFixVulMachineResponse, err error) {
    if request == nil {
        request = NewDescribeCanFixVulMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCanFixVulMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCanFixVulMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCanNotSeparateMachineRequest() (request *DescribeCanNotSeparateMachineRequest) {
    request = &DescribeCanNotSeparateMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeCanNotSeparateMachine")
    
    
    return
}

func NewDescribeCanNotSeparateMachineResponse() (response *DescribeCanNotSeparateMachineResponse) {
    response = &DescribeCanNotSeparateMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCanNotSeparateMachine
// 获取木马不可隔离的主机
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCanNotSeparateMachine(request *DescribeCanNotSeparateMachineRequest) (response *DescribeCanNotSeparateMachineResponse, err error) {
    return c.DescribeCanNotSeparateMachineWithContext(context.Background(), request)
}

// DescribeCanNotSeparateMachine
// 获取木马不可隔离的主机
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCanNotSeparateMachineWithContext(ctx context.Context, request *DescribeCanNotSeparateMachineRequest) (response *DescribeCanNotSeparateMachineResponse, err error) {
    if request == nil {
        request = NewDescribeCanNotSeparateMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCanNotSeparateMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCanNotSeparateMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientExceptionRequest() (request *DescribeClientExceptionRequest) {
    request = &DescribeClientExceptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeClientException")
    
    
    return
}

func NewDescribeClientExceptionResponse() (response *DescribeClientExceptionResponse) {
    response = &DescribeClientExceptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClientException
// 获取客户端异常事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClientException(request *DescribeClientExceptionRequest) (response *DescribeClientExceptionResponse, err error) {
    return c.DescribeClientExceptionWithContext(context.Background(), request)
}

// DescribeClientException
// 获取客户端异常事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClientExceptionWithContext(ctx context.Context, request *DescribeClientExceptionRequest) (response *DescribeClientExceptionResponse, err error) {
    if request == nil {
        request = NewDescribeClientExceptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientException require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientExceptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudProtectServiceOrderListRequest() (request *DescribeCloudProtectServiceOrderListRequest) {
    request = &DescribeCloudProtectServiceOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeCloudProtectServiceOrderList")
    
    
    return
}

func NewDescribeCloudProtectServiceOrderListResponse() (response *DescribeCloudProtectServiceOrderListResponse) {
    response = &DescribeCloudProtectServiceOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudProtectServiceOrderList
// 查询云护航服务订单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudProtectServiceOrderList(request *DescribeCloudProtectServiceOrderListRequest) (response *DescribeCloudProtectServiceOrderListResponse, err error) {
    return c.DescribeCloudProtectServiceOrderListWithContext(context.Background(), request)
}

// DescribeCloudProtectServiceOrderList
// 查询云护航服务订单列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudProtectServiceOrderListWithContext(ctx context.Context, request *DescribeCloudProtectServiceOrderListRequest) (response *DescribeCloudProtectServiceOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudProtectServiceOrderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudProtectServiceOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudProtectServiceOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
    request = &DescribeComponentStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeComponentStatistics")
    
    
    return
}

func NewDescribeComponentStatisticsResponse() (response *DescribeComponentStatisticsResponse) {
    response = &DescribeComponentStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponentStatistics
// 本接口 (DescribeComponentStatistics) 用于获取组件统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    return c.DescribeComponentStatisticsWithContext(context.Background(), request)
}

// DescribeComponentStatistics
// 本接口 (DescribeComponentStatistics) 用于获取组件统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeComponentStatisticsWithContext(ctx context.Context, request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponentStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefenceEventDetailRequest() (request *DescribeDefenceEventDetailRequest) {
    request = &DescribeDefenceEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeDefenceEventDetail")
    
    
    return
}

func NewDescribeDefenceEventDetailResponse() (response *DescribeDefenceEventDetailResponse) {
    response = &DescribeDefenceEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefenceEventDetail
// 获取漏洞防御事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDefenceEventDetail(request *DescribeDefenceEventDetailRequest) (response *DescribeDefenceEventDetailResponse, err error) {
    return c.DescribeDefenceEventDetailWithContext(context.Background(), request)
}

// DescribeDefenceEventDetail
// 获取漏洞防御事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDefenceEventDetailWithContext(ctx context.Context, request *DescribeDefenceEventDetailRequest) (response *DescribeDefenceEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDefenceEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefenceEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefenceEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectInstallCommandRequest() (request *DescribeDirectConnectInstallCommandRequest) {
    request = &DescribeDirectConnectInstallCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeDirectConnectInstallCommand")
    
    
    return
}

func NewDescribeDirectConnectInstallCommandResponse() (response *DescribeDirectConnectInstallCommandResponse) {
    response = &DescribeDirectConnectInstallCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDirectConnectInstallCommand
// 获取专线agent安装命令，包含token
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnectInstallCommand(request *DescribeDirectConnectInstallCommandRequest) (response *DescribeDirectConnectInstallCommandResponse, err error) {
    return c.DescribeDirectConnectInstallCommandWithContext(context.Background(), request)
}

// DescribeDirectConnectInstallCommand
// 获取专线agent安装命令，包含token
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnectInstallCommandWithContext(ctx context.Context, request *DescribeDirectConnectInstallCommandRequest) (response *DescribeDirectConnectInstallCommandResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectInstallCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDirectConnectInstallCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDirectConnectInstallCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeESAggregationsRequest() (request *DescribeESAggregationsRequest) {
    request = &DescribeESAggregationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeESAggregations")
    
    
    return
}

func NewDescribeESAggregationsResponse() (response *DescribeESAggregationsResponse) {
    response = &DescribeESAggregationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeESAggregations
// 获取ES字段聚合结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeESAggregations(request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
    return c.DescribeESAggregationsWithContext(context.Background(), request)
}

// DescribeESAggregations
// 获取ES字段聚合结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeESAggregationsWithContext(ctx context.Context, request *DescribeESAggregationsRequest) (response *DescribeESAggregationsResponse, err error) {
    if request == nil {
        request = NewDescribeESAggregationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeESAggregations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeESAggregationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmergencyResponseListRequest() (request *DescribeEmergencyResponseListRequest) {
    request = &DescribeEmergencyResponseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeEmergencyResponseList")
    
    
    return
}

func NewDescribeEmergencyResponseListResponse() (response *DescribeEmergencyResponseListResponse) {
    response = &DescribeEmergencyResponseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEmergencyResponseList
// 专家服务-应急响应列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEmergencyResponseList(request *DescribeEmergencyResponseListRequest) (response *DescribeEmergencyResponseListResponse, err error) {
    return c.DescribeEmergencyResponseListWithContext(context.Background(), request)
}

// DescribeEmergencyResponseList
// 专家服务-应急响应列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEmergencyResponseListWithContext(ctx context.Context, request *DescribeEmergencyResponseListRequest) (response *DescribeEmergencyResponseListResponse, err error) {
    if request == nil {
        request = NewDescribeEmergencyResponseListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmergencyResponseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmergencyResponseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEmergencyVulListRequest() (request *DescribeEmergencyVulListRequest) {
    request = &DescribeEmergencyVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeEmergencyVulList")
    
    
    return
}

func NewDescribeEmergencyVulListResponse() (response *DescribeEmergencyVulListResponse) {
    response = &DescribeEmergencyVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEmergencyVulList
// 获取应急漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEmergencyVulList(request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
    return c.DescribeEmergencyVulListWithContext(context.Background(), request)
}

// DescribeEmergencyVulList
// 获取应急漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEmergencyVulListWithContext(ctx context.Context, request *DescribeEmergencyVulListRequest) (response *DescribeEmergencyVulListResponse, err error) {
    if request == nil {
        request = NewDescribeEmergencyVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEmergencyVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEmergencyVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventByTableRequest() (request *DescribeEventByTableRequest) {
    request = &DescribeEventByTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeEventByTable")
    
    
    return
}

func NewDescribeEventByTableResponse() (response *DescribeEventByTableResponse) {
    response = &DescribeEventByTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEventByTable
// 根据事件表名和id查询告警事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEventByTable(request *DescribeEventByTableRequest) (response *DescribeEventByTableResponse, err error) {
    return c.DescribeEventByTableWithContext(context.Background(), request)
}

// DescribeEventByTable
// 根据事件表名和id查询告警事件详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeEventByTableWithContext(ctx context.Context, request *DescribeEventByTableRequest) (response *DescribeEventByTableResponse, err error) {
    if request == nil {
        request = NewDescribeEventByTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventByTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventByTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExpertServiceListRequest() (request *DescribeExpertServiceListRequest) {
    request = &DescribeExpertServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeExpertServiceList")
    
    
    return
}

func NewDescribeExpertServiceListResponse() (response *DescribeExpertServiceListResponse) {
    response = &DescribeExpertServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExpertServiceList
// 专家服务-安全管家列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExpertServiceList(request *DescribeExpertServiceListRequest) (response *DescribeExpertServiceListResponse, err error) {
    return c.DescribeExpertServiceListWithContext(context.Background(), request)
}

// DescribeExpertServiceList
// 专家服务-安全管家列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExpertServiceListWithContext(ctx context.Context, request *DescribeExpertServiceListRequest) (response *DescribeExpertServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeExpertServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExpertServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExpertServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExpertServiceOrderListRequest() (request *DescribeExpertServiceOrderListRequest) {
    request = &DescribeExpertServiceOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeExpertServiceOrderList")
    
    
    return
}

func NewDescribeExpertServiceOrderListResponse() (response *DescribeExpertServiceOrderListResponse) {
    response = &DescribeExpertServiceOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExpertServiceOrderList
// 专家服务-专家服务订单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExpertServiceOrderList(request *DescribeExpertServiceOrderListRequest) (response *DescribeExpertServiceOrderListResponse, err error) {
    return c.DescribeExpertServiceOrderListWithContext(context.Background(), request)
}

// DescribeExpertServiceOrderList
// 专家服务-专家服务订单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExpertServiceOrderListWithContext(ctx context.Context, request *DescribeExpertServiceOrderListRequest) (response *DescribeExpertServiceOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeExpertServiceOrderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExpertServiceOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExpertServiceOrderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportMachinesRequest() (request *DescribeExportMachinesRequest) {
    request = &DescribeExportMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeExportMachines")
    
    
    return
}

func NewDescribeExportMachinesResponse() (response *DescribeExportMachinesResponse) {
    response = &DescribeExportMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExportMachines
// 本接口 (DescribeExportMachines) 用于导出区域主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExportMachines(request *DescribeExportMachinesRequest) (response *DescribeExportMachinesResponse, err error) {
    return c.DescribeExportMachinesWithContext(context.Background(), request)
}

// DescribeExportMachines
// 本接口 (DescribeExportMachines) 用于导出区域主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeExportMachinesWithContext(ctx context.Context, request *DescribeExportMachinesRequest) (response *DescribeExportMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeExportMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFastAnalysisRequest() (request *DescribeFastAnalysisRequest) {
    request = &DescribeFastAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFastAnalysis")
    
    
    return
}

func NewDescribeFastAnalysisResponse() (response *DescribeFastAnalysisResponse) {
    response = &DescribeFastAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFastAnalysis
// 日志快速分析统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFastAnalysis(request *DescribeFastAnalysisRequest) (response *DescribeFastAnalysisResponse, err error) {
    return c.DescribeFastAnalysisWithContext(context.Background(), request)
}

// DescribeFastAnalysis
// 日志快速分析统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFastAnalysisWithContext(ctx context.Context, request *DescribeFastAnalysisRequest) (response *DescribeFastAnalysisResponse, err error) {
    if request == nil {
        request = NewDescribeFastAnalysisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFastAnalysis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFastAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTamperEventRuleInfoRequest() (request *DescribeFileTamperEventRuleInfoRequest) {
    request = &DescribeFileTamperEventRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperEventRuleInfo")
    
    
    return
}

func NewDescribeFileTamperEventRuleInfoResponse() (response *DescribeFileTamperEventRuleInfoResponse) {
    response = &DescribeFileTamperEventRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTamperEventRuleInfo
// 查看产生事件时规则详情接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperEventRuleInfo(request *DescribeFileTamperEventRuleInfoRequest) (response *DescribeFileTamperEventRuleInfoResponse, err error) {
    return c.DescribeFileTamperEventRuleInfoWithContext(context.Background(), request)
}

// DescribeFileTamperEventRuleInfo
// 查看产生事件时规则详情接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperEventRuleInfoWithContext(ctx context.Context, request *DescribeFileTamperEventRuleInfoRequest) (response *DescribeFileTamperEventRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFileTamperEventRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTamperEventRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTamperEventRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTamperEventsRequest() (request *DescribeFileTamperEventsRequest) {
    request = &DescribeFileTamperEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperEvents")
    
    
    return
}

func NewDescribeFileTamperEventsResponse() (response *DescribeFileTamperEventsResponse) {
    response = &DescribeFileTamperEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTamperEvents
// 核心文件监控事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFileTamperEvents(request *DescribeFileTamperEventsRequest) (response *DescribeFileTamperEventsResponse, err error) {
    return c.DescribeFileTamperEventsWithContext(context.Background(), request)
}

// DescribeFileTamperEvents
// 核心文件监控事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFileTamperEventsWithContext(ctx context.Context, request *DescribeFileTamperEventsRequest) (response *DescribeFileTamperEventsResponse, err error) {
    if request == nil {
        request = NewDescribeFileTamperEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTamperEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTamperEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTamperRuleCountRequest() (request *DescribeFileTamperRuleCountRequest) {
    request = &DescribeFileTamperRuleCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRuleCount")
    
    
    return
}

func NewDescribeFileTamperRuleCountResponse() (response *DescribeFileTamperRuleCountResponse) {
    response = &DescribeFileTamperRuleCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTamperRuleCount
// 查询主机关联文件监控规则数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperRuleCount(request *DescribeFileTamperRuleCountRequest) (response *DescribeFileTamperRuleCountResponse, err error) {
    return c.DescribeFileTamperRuleCountWithContext(context.Background(), request)
}

// DescribeFileTamperRuleCount
// 查询主机关联文件监控规则数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperRuleCountWithContext(ctx context.Context, request *DescribeFileTamperRuleCountRequest) (response *DescribeFileTamperRuleCountResponse, err error) {
    if request == nil {
        request = NewDescribeFileTamperRuleCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTamperRuleCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTamperRuleCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTamperRuleInfoRequest() (request *DescribeFileTamperRuleInfoRequest) {
    request = &DescribeFileTamperRuleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRuleInfo")
    
    
    return
}

func NewDescribeFileTamperRuleInfoResponse() (response *DescribeFileTamperRuleInfoResponse) {
    response = &DescribeFileTamperRuleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTamperRuleInfo
// 查询某个监控规则的详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperRuleInfo(request *DescribeFileTamperRuleInfoRequest) (response *DescribeFileTamperRuleInfoResponse, err error) {
    return c.DescribeFileTamperRuleInfoWithContext(context.Background(), request)
}

// DescribeFileTamperRuleInfo
// 查询某个监控规则的详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFileTamperRuleInfoWithContext(ctx context.Context, request *DescribeFileTamperRuleInfoRequest) (response *DescribeFileTamperRuleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFileTamperRuleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTamperRuleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTamperRuleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTamperRulesRequest() (request *DescribeFileTamperRulesRequest) {
    request = &DescribeFileTamperRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeFileTamperRules")
    
    
    return
}

func NewDescribeFileTamperRulesResponse() (response *DescribeFileTamperRulesResponse) {
    response = &DescribeFileTamperRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTamperRules
// 核心文件监控规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFileTamperRules(request *DescribeFileTamperRulesRequest) (response *DescribeFileTamperRulesResponse, err error) {
    return c.DescribeFileTamperRulesWithContext(context.Background(), request)
}

// DescribeFileTamperRules
// 核心文件监控规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFileTamperRulesWithContext(ctx context.Context, request *DescribeFileTamperRulesRequest) (response *DescribeFileTamperRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFileTamperRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTamperRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTamperRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralStatRequest() (request *DescribeGeneralStatRequest) {
    request = &DescribeGeneralStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeGeneralStat")
    
    
    return
}

func NewDescribeGeneralStatResponse() (response *DescribeGeneralStatResponse) {
    response = &DescribeGeneralStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGeneralStat
// 获取主机相关统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeGeneralStat(request *DescribeGeneralStatRequest) (response *DescribeGeneralStatResponse, err error) {
    return c.DescribeGeneralStatWithContext(context.Background(), request)
}

// DescribeGeneralStat
// 获取主机相关统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeGeneralStatWithContext(ctx context.Context, request *DescribeGeneralStatRequest) (response *DescribeGeneralStatResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
    request = &DescribeHistoryAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeHistoryAccounts")
    
    
    return
}

func NewDescribeHistoryAccountsResponse() (response *DescribeHistoryAccountsResponse) {
    response = &DescribeHistoryAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryAccounts
// 本接口 (DescribeHistoryAccounts) 用于获取帐号变更历史列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistoryAccounts(request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    return c.DescribeHistoryAccountsWithContext(context.Background(), request)
}

// DescribeHistoryAccounts
// 本接口 (DescribeHistoryAccounts) 用于获取帐号变更历史列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHistoryAccountsWithContext(ctx context.Context, request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistoryAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistoryAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryServiceRequest() (request *DescribeHistoryServiceRequest) {
    request = &DescribeHistoryServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeHistoryService")
    
    
    return
}

func NewDescribeHistoryServiceResponse() (response *DescribeHistoryServiceResponse) {
    response = &DescribeHistoryServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryService
// 查询日志检索服务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeHistoryService(request *DescribeHistoryServiceRequest) (response *DescribeHistoryServiceResponse, err error) {
    return c.DescribeHistoryServiceWithContext(context.Background(), request)
}

// DescribeHistoryService
// 查询日志检索服务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeHistoryServiceWithContext(ctx context.Context, request *DescribeHistoryServiceRequest) (response *DescribeHistoryServiceResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistoryService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistoryServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostInfoRequest() (request *DescribeHostInfoRequest) {
    request = &DescribeHostInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeHostInfo")
    
    
    return
}

func NewDescribeHostInfoResponse() (response *DescribeHostInfoResponse) {
    response = &DescribeHostInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostInfo
// 主机信息与标签信息查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHostInfo(request *DescribeHostInfoRequest) (response *DescribeHostInfoResponse, err error) {
    return c.DescribeHostInfoWithContext(context.Background(), request)
}

// DescribeHostInfo
// 主机信息与标签信息查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHostInfoWithContext(ctx context.Context, request *DescribeHostInfoRequest) (response *DescribeHostInfoResponse, err error) {
    if request == nil {
        request = NewDescribeHostInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostLoginListRequest() (request *DescribeHostLoginListRequest) {
    request = &DescribeHostLoginListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeHostLoginList")
    
    
    return
}

func NewDescribeHostLoginListResponse() (response *DescribeHostLoginListResponse) {
    response = &DescribeHostLoginListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostLoginList
// 获取登录审计列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHostLoginList(request *DescribeHostLoginListRequest) (response *DescribeHostLoginListResponse, err error) {
    return c.DescribeHostLoginListWithContext(context.Background(), request)
}

// DescribeHostLoginList
// 获取登录审计列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHostLoginListWithContext(ctx context.Context, request *DescribeHostLoginListRequest) (response *DescribeHostLoginListResponse, err error) {
    if request == nil {
        request = NewDescribeHostLoginListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostLoginList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostLoginListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHotVulTopRequest() (request *DescribeHotVulTopRequest) {
    request = &DescribeHotVulTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeHotVulTop")
    
    
    return
}

func NewDescribeHotVulTopResponse() (response *DescribeHotVulTopResponse) {
    response = &DescribeHotVulTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHotVulTop
// 获取全网热点漏洞
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHotVulTop(request *DescribeHotVulTopRequest) (response *DescribeHotVulTopResponse, err error) {
    return c.DescribeHotVulTopWithContext(context.Background(), request)
}

// DescribeHotVulTop
// 获取全网热点漏洞
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHotVulTopWithContext(ctx context.Context, request *DescribeHotVulTopRequest) (response *DescribeHotVulTopResponse, err error) {
    if request == nil {
        request = NewDescribeHotVulTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHotVulTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHotVulTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIgnoreBaselineRuleRequest() (request *DescribeIgnoreBaselineRuleRequest) {
    request = &DescribeIgnoreBaselineRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreBaselineRule")
    
    
    return
}

func NewDescribeIgnoreBaselineRuleResponse() (response *DescribeIgnoreBaselineRuleResponse) {
    response = &DescribeIgnoreBaselineRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIgnoreBaselineRule
// 查询已经忽略的检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIgnoreBaselineRule(request *DescribeIgnoreBaselineRuleRequest) (response *DescribeIgnoreBaselineRuleResponse, err error) {
    return c.DescribeIgnoreBaselineRuleWithContext(context.Background(), request)
}

// DescribeIgnoreBaselineRule
// 查询已经忽略的检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIgnoreBaselineRuleWithContext(ctx context.Context, request *DescribeIgnoreBaselineRuleRequest) (response *DescribeIgnoreBaselineRuleResponse, err error) {
    if request == nil {
        request = NewDescribeIgnoreBaselineRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIgnoreBaselineRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIgnoreBaselineRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIgnoreHostAndItemConfigRequest() (request *DescribeIgnoreHostAndItemConfigRequest) {
    request = &DescribeIgnoreHostAndItemConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreHostAndItemConfig")
    
    
    return
}

func NewDescribeIgnoreHostAndItemConfigResponse() (response *DescribeIgnoreHostAndItemConfigResponse) {
    response = &DescribeIgnoreHostAndItemConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIgnoreHostAndItemConfig
// 获取一键忽略受影响的检测项和主机信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgnoreHostAndItemConfig(request *DescribeIgnoreHostAndItemConfigRequest) (response *DescribeIgnoreHostAndItemConfigResponse, err error) {
    return c.DescribeIgnoreHostAndItemConfigWithContext(context.Background(), request)
}

// DescribeIgnoreHostAndItemConfig
// 获取一键忽略受影响的检测项和主机信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgnoreHostAndItemConfigWithContext(ctx context.Context, request *DescribeIgnoreHostAndItemConfigRequest) (response *DescribeIgnoreHostAndItemConfigResponse, err error) {
    if request == nil {
        request = NewDescribeIgnoreHostAndItemConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIgnoreHostAndItemConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIgnoreHostAndItemConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIgnoreRuleEffectHostListRequest() (request *DescribeIgnoreRuleEffectHostListRequest) {
    request = &DescribeIgnoreRuleEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeIgnoreRuleEffectHostList")
    
    
    return
}

func NewDescribeIgnoreRuleEffectHostListResponse() (response *DescribeIgnoreRuleEffectHostListResponse) {
    response = &DescribeIgnoreRuleEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIgnoreRuleEffectHostList
// 根据检测项id与筛选条件查询忽略检测项影响主机列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIgnoreRuleEffectHostList(request *DescribeIgnoreRuleEffectHostListRequest) (response *DescribeIgnoreRuleEffectHostListResponse, err error) {
    return c.DescribeIgnoreRuleEffectHostListWithContext(context.Background(), request)
}

// DescribeIgnoreRuleEffectHostList
// 根据检测项id与筛选条件查询忽略检测项影响主机列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIgnoreRuleEffectHostListWithContext(ctx context.Context, request *DescribeIgnoreRuleEffectHostListRequest) (response *DescribeIgnoreRuleEffectHostListResponse, err error) {
    if request == nil {
        request = NewDescribeIgnoreRuleEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIgnoreRuleEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIgnoreRuleEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImportMachineInfoRequest() (request *DescribeImportMachineInfoRequest) {
    request = &DescribeImportMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeImportMachineInfo")
    
    
    return
}

func NewDescribeImportMachineInfoResponse() (response *DescribeImportMachineInfoResponse) {
    response = &DescribeImportMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImportMachineInfo
// 查询批量导入机器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeImportMachineInfo(request *DescribeImportMachineInfoRequest) (response *DescribeImportMachineInfoResponse, err error) {
    return c.DescribeImportMachineInfoWithContext(context.Background(), request)
}

// DescribeImportMachineInfo
// 查询批量导入机器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeImportMachineInfoWithContext(ctx context.Context, request *DescribeImportMachineInfoRequest) (response *DescribeImportMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeImportMachineInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImportMachineInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImportMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
    request = &DescribeIndexListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeIndexList")
    
    
    return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
    response = &DescribeIndexListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    return c.DescribeIndexListWithContext(context.Background(), request)
}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIndexListWithContext(ctx context.Context, request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJavaMemShellInfoRequest() (request *DescribeJavaMemShellInfoRequest) {
    request = &DescribeJavaMemShellInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellInfo")
    
    
    return
}

func NewDescribeJavaMemShellInfoResponse() (response *DescribeJavaMemShellInfoResponse) {
    response = &DescribeJavaMemShellInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJavaMemShellInfo
// 查询java内存马事件详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJavaMemShellInfo(request *DescribeJavaMemShellInfoRequest) (response *DescribeJavaMemShellInfoResponse, err error) {
    return c.DescribeJavaMemShellInfoWithContext(context.Background(), request)
}

// DescribeJavaMemShellInfo
// 查询java内存马事件详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJavaMemShellInfoWithContext(ctx context.Context, request *DescribeJavaMemShellInfoRequest) (response *DescribeJavaMemShellInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJavaMemShellInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJavaMemShellInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJavaMemShellInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJavaMemShellListRequest() (request *DescribeJavaMemShellListRequest) {
    request = &DescribeJavaMemShellListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellList")
    
    
    return
}

func NewDescribeJavaMemShellListResponse() (response *DescribeJavaMemShellListResponse) {
    response = &DescribeJavaMemShellListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJavaMemShellList
// 查询java内存马事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellList(request *DescribeJavaMemShellListRequest) (response *DescribeJavaMemShellListResponse, err error) {
    return c.DescribeJavaMemShellListWithContext(context.Background(), request)
}

// DescribeJavaMemShellList
// 查询java内存马事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellListWithContext(ctx context.Context, request *DescribeJavaMemShellListRequest) (response *DescribeJavaMemShellListResponse, err error) {
    if request == nil {
        request = NewDescribeJavaMemShellListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJavaMemShellList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJavaMemShellListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJavaMemShellPluginInfoRequest() (request *DescribeJavaMemShellPluginInfoRequest) {
    request = &DescribeJavaMemShellPluginInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellPluginInfo")
    
    
    return
}

func NewDescribeJavaMemShellPluginInfoResponse() (response *DescribeJavaMemShellPluginInfoResponse) {
    response = &DescribeJavaMemShellPluginInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJavaMemShellPluginInfo
// 查询给定主机java内存马插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellPluginInfo(request *DescribeJavaMemShellPluginInfoRequest) (response *DescribeJavaMemShellPluginInfoResponse, err error) {
    return c.DescribeJavaMemShellPluginInfoWithContext(context.Background(), request)
}

// DescribeJavaMemShellPluginInfo
// 查询给定主机java内存马插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellPluginInfoWithContext(ctx context.Context, request *DescribeJavaMemShellPluginInfoRequest) (response *DescribeJavaMemShellPluginInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJavaMemShellPluginInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJavaMemShellPluginInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJavaMemShellPluginInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJavaMemShellPluginListRequest() (request *DescribeJavaMemShellPluginListRequest) {
    request = &DescribeJavaMemShellPluginListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeJavaMemShellPluginList")
    
    
    return
}

func NewDescribeJavaMemShellPluginListResponse() (response *DescribeJavaMemShellPluginListResponse) {
    response = &DescribeJavaMemShellPluginListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJavaMemShellPluginList
// 查询java内存马插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellPluginList(request *DescribeJavaMemShellPluginListRequest) (response *DescribeJavaMemShellPluginListResponse, err error) {
    return c.DescribeJavaMemShellPluginListWithContext(context.Background(), request)
}

// DescribeJavaMemShellPluginList
// 查询java内存马插件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeJavaMemShellPluginListWithContext(ctx context.Context, request *DescribeJavaMemShellPluginListRequest) (response *DescribeJavaMemShellPluginListResponse, err error) {
    if request == nil {
        request = NewDescribeJavaMemShellPluginListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJavaMemShellPluginList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJavaMemShellPluginListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseRequest() (request *DescribeLicenseRequest) {
    request = &DescribeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicense")
    
    
    return
}

func NewDescribeLicenseResponse() (response *DescribeLicenseResponse) {
    response = &DescribeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicense
// 查询授权信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicense(request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
    return c.DescribeLicenseWithContext(context.Background(), request)
}

// DescribeLicense
// 查询授权信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseWithContext(ctx context.Context, request *DescribeLicenseRequest) (response *DescribeLicenseResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseBindListRequest() (request *DescribeLicenseBindListRequest) {
    request = &DescribeLicenseBindListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseBindList")
    
    
    return
}

func NewDescribeLicenseBindListResponse() (response *DescribeLicenseBindListResponse) {
    response = &DescribeLicenseBindListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicenseBindList
// 该接口可以获取设置中心-授权管理,某个授权下已绑定的授权机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseBindList(request *DescribeLicenseBindListRequest) (response *DescribeLicenseBindListResponse, err error) {
    return c.DescribeLicenseBindListWithContext(context.Background(), request)
}

// DescribeLicenseBindList
// 该接口可以获取设置中心-授权管理,某个授权下已绑定的授权机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseBindListWithContext(ctx context.Context, request *DescribeLicenseBindListRequest) (response *DescribeLicenseBindListResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseBindListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseBindList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseBindListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseBindScheduleRequest() (request *DescribeLicenseBindScheduleRequest) {
    request = &DescribeLicenseBindScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseBindSchedule")
    
    
    return
}

func NewDescribeLicenseBindScheduleResponse() (response *DescribeLicenseBindScheduleResponse) {
    response = &DescribeLicenseBindScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicenseBindSchedule
// 查询授权绑定任务的进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseBindSchedule(request *DescribeLicenseBindScheduleRequest) (response *DescribeLicenseBindScheduleResponse, err error) {
    return c.DescribeLicenseBindScheduleWithContext(context.Background(), request)
}

// DescribeLicenseBindSchedule
// 查询授权绑定任务的进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseBindScheduleWithContext(ctx context.Context, request *DescribeLicenseBindScheduleRequest) (response *DescribeLicenseBindScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseBindScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseBindSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseBindScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseGeneralRequest() (request *DescribeLicenseGeneralRequest) {
    request = &DescribeLicenseGeneralRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseGeneral")
    
    
    return
}

func NewDescribeLicenseGeneralResponse() (response *DescribeLicenseGeneralResponse) {
    response = &DescribeLicenseGeneralResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicenseGeneral
// 授权管理-授权概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLicenseGeneral(request *DescribeLicenseGeneralRequest) (response *DescribeLicenseGeneralResponse, err error) {
    return c.DescribeLicenseGeneralWithContext(context.Background(), request)
}

// DescribeLicenseGeneral
// 授权管理-授权概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLicenseGeneralWithContext(ctx context.Context, request *DescribeLicenseGeneralRequest) (response *DescribeLicenseGeneralResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseGeneralRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseGeneral require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseGeneralResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseListRequest() (request *DescribeLicenseListRequest) {
    request = &DescribeLicenseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseList")
    
    
    return
}

func NewDescribeLicenseListResponse() (response *DescribeLicenseListResponse) {
    response = &DescribeLicenseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicenseList
// 获取用户所有授权订单信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseList(request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
    return c.DescribeLicenseListWithContext(context.Background(), request)
}

// DescribeLicenseList
// 获取用户所有授权订单信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseListWithContext(ctx context.Context, request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseWhiteConfigRequest() (request *DescribeLicenseWhiteConfigRequest) {
    request = &DescribeLicenseWhiteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLicenseWhiteConfig")
    
    
    return
}

func NewDescribeLicenseWhiteConfigResponse() (response *DescribeLicenseWhiteConfigResponse) {
    response = &DescribeLicenseWhiteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLicenseWhiteConfig
// 查询授权白名单的可用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseWhiteConfig(request *DescribeLicenseWhiteConfigRequest) (response *DescribeLicenseWhiteConfigResponse, err error) {
    return c.DescribeLicenseWhiteConfigWithContext(context.Background(), request)
}

// DescribeLicenseWhiteConfig
// 查询授权白名单的可用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLicenseWhiteConfigWithContext(ctx context.Context, request *DescribeLicenseWhiteConfigRequest) (response *DescribeLicenseWhiteConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseWhiteConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseWhiteConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseWhiteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogDeliveryKafkaOptionsRequest() (request *DescribeLogDeliveryKafkaOptionsRequest) {
    request = &DescribeLogDeliveryKafkaOptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogDeliveryKafkaOptions")
    
    
    return
}

func NewDescribeLogDeliveryKafkaOptionsResponse() (response *DescribeLogDeliveryKafkaOptionsResponse) {
    response = &DescribeLogDeliveryKafkaOptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogDeliveryKafkaOptions
// 查询日志投递kafka可选项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogDeliveryKafkaOptions(request *DescribeLogDeliveryKafkaOptionsRequest) (response *DescribeLogDeliveryKafkaOptionsResponse, err error) {
    return c.DescribeLogDeliveryKafkaOptionsWithContext(context.Background(), request)
}

// DescribeLogDeliveryKafkaOptions
// 查询日志投递kafka可选项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogDeliveryKafkaOptionsWithContext(ctx context.Context, request *DescribeLogDeliveryKafkaOptionsRequest) (response *DescribeLogDeliveryKafkaOptionsResponse, err error) {
    if request == nil {
        request = NewDescribeLogDeliveryKafkaOptionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogDeliveryKafkaOptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogDeliveryKafkaOptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogExportsRequest() (request *DescribeLogExportsRequest) {
    request = &DescribeLogExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogExports")
    
    
    return
}

func NewDescribeLogExportsResponse() (response *DescribeLogExportsResponse) {
    response = &DescribeLogExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogExports
// 获取日志下载任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogExports(request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    return c.DescribeLogExportsWithContext(context.Background(), request)
}

// DescribeLogExports
// 获取日志下载任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogExportsWithContext(ctx context.Context, request *DescribeLogExportsRequest) (response *DescribeLogExportsResponse, err error) {
    if request == nil {
        request = NewDescribeLogExportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
    request = &DescribeLogHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogHistogram")
    
    
    return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
    response = &DescribeLogHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogHistogram
// 获取日志直方图信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    return c.DescribeLogHistogramWithContext(context.Background(), request)
}

// DescribeLogHistogram
// 获取日志直方图信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogHistogramWithContext(ctx context.Context, request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeLogHistogramRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogIndexRequest() (request *DescribeLogIndexRequest) {
    request = &DescribeLogIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogIndex")
    
    
    return
}

func NewDescribeLogIndexResponse() (response *DescribeLogIndexResponse) {
    response = &DescribeLogIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogIndex
// 查询索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogIndex(request *DescribeLogIndexRequest) (response *DescribeLogIndexResponse, err error) {
    return c.DescribeLogIndexWithContext(context.Background(), request)
}

// DescribeLogIndex
// 查询索引
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogIndexWithContext(ctx context.Context, request *DescribeLogIndexRequest) (response *DescribeLogIndexResponse, err error) {
    if request == nil {
        request = NewDescribeLogIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogKafkaDeliverInfoRequest() (request *DescribeLogKafkaDeliverInfoRequest) {
    request = &DescribeLogKafkaDeliverInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogKafkaDeliverInfo")
    
    
    return
}

func NewDescribeLogKafkaDeliverInfoResponse() (response *DescribeLogKafkaDeliverInfoResponse) {
    response = &DescribeLogKafkaDeliverInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogKafkaDeliverInfo
// 获取kafka投递信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogKafkaDeliverInfo(request *DescribeLogKafkaDeliverInfoRequest) (response *DescribeLogKafkaDeliverInfoResponse, err error) {
    return c.DescribeLogKafkaDeliverInfoWithContext(context.Background(), request)
}

// DescribeLogKafkaDeliverInfo
// 获取kafka投递信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogKafkaDeliverInfoWithContext(ctx context.Context, request *DescribeLogKafkaDeliverInfoRequest) (response *DescribeLogKafkaDeliverInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLogKafkaDeliverInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogKafkaDeliverInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogKafkaDeliverInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogStorageConfigRequest() (request *DescribeLogStorageConfigRequest) {
    request = &DescribeLogStorageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageConfig")
    
    
    return
}

func NewDescribeLogStorageConfigResponse() (response *DescribeLogStorageConfigResponse) {
    response = &DescribeLogStorageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogStorageConfig
// 获取日志存储配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogStorageConfig(request *DescribeLogStorageConfigRequest) (response *DescribeLogStorageConfigResponse, err error) {
    return c.DescribeLogStorageConfigWithContext(context.Background(), request)
}

// DescribeLogStorageConfig
// 获取日志存储配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogStorageConfigWithContext(ctx context.Context, request *DescribeLogStorageConfigRequest) (response *DescribeLogStorageConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLogStorageConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogStorageConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogStorageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogStorageRecordRequest() (request *DescribeLogStorageRecordRequest) {
    request = &DescribeLogStorageRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageRecord")
    
    
    return
}

func NewDescribeLogStorageRecordResponse() (response *DescribeLogStorageRecordResponse) {
    response = &DescribeLogStorageRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogStorageRecord
// 获取日志存储量记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogStorageRecord(request *DescribeLogStorageRecordRequest) (response *DescribeLogStorageRecordResponse, err error) {
    return c.DescribeLogStorageRecordWithContext(context.Background(), request)
}

// DescribeLogStorageRecord
// 获取日志存储量记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogStorageRecordWithContext(ctx context.Context, request *DescribeLogStorageRecordRequest) (response *DescribeLogStorageRecordResponse, err error) {
    if request == nil {
        request = NewDescribeLogStorageRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogStorageRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogStorageRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogStorageStatisticRequest() (request *DescribeLogStorageStatisticRequest) {
    request = &DescribeLogStorageStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogStorageStatistic")
    
    
    return
}

func NewDescribeLogStorageStatisticResponse() (response *DescribeLogStorageStatisticResponse) {
    response = &DescribeLogStorageStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogStorageStatistic
// 获取日志检索容量使用统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
    return c.DescribeLogStorageStatisticWithContext(context.Background(), request)
}

// DescribeLogStorageStatistic
// 获取日志检索容量使用统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLogStorageStatisticWithContext(ctx context.Context, request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeLogStorageStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogStorageStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogStorageStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogTypeRequest() (request *DescribeLogTypeRequest) {
    request = &DescribeLogTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLogType")
    
    
    return
}

func NewDescribeLogTypeResponse() (response *DescribeLogTypeResponse) {
    response = &DescribeLogTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogType
// 日志分析功能-获取日志类型，使用该接口返回的结果暂时可过滤的日志类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLogType(request *DescribeLogTypeRequest) (response *DescribeLogTypeResponse, err error) {
    return c.DescribeLogTypeWithContext(context.Background(), request)
}

// DescribeLogType
// 日志分析功能-获取日志类型，使用该接口返回的结果暂时可过滤的日志类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLogTypeWithContext(ctx context.Context, request *DescribeLogTypeRequest) (response *DescribeLogTypeResponse, err error) {
    if request == nil {
        request = NewDescribeLogTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteCombinedListRequest() (request *DescribeLoginWhiteCombinedListRequest) {
    request = &DescribeLoginWhiteCombinedListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteCombinedList")
    
    
    return
}

func NewDescribeLoginWhiteCombinedListResponse() (response *DescribeLoginWhiteCombinedListResponse) {
    response = &DescribeLoginWhiteCombinedListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoginWhiteCombinedList
// 获取异地登录白名单合并后列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoginWhiteCombinedList(request *DescribeLoginWhiteCombinedListRequest) (response *DescribeLoginWhiteCombinedListResponse, err error) {
    return c.DescribeLoginWhiteCombinedListWithContext(context.Background(), request)
}

// DescribeLoginWhiteCombinedList
// 获取异地登录白名单合并后列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoginWhiteCombinedListWithContext(ctx context.Context, request *DescribeLoginWhiteCombinedListRequest) (response *DescribeLoginWhiteCombinedListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteCombinedListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginWhiteCombinedList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginWhiteCombinedListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteHostListRequest() (request *DescribeLoginWhiteHostListRequest) {
    request = &DescribeLoginWhiteHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteHostList")
    
    
    return
}

func NewDescribeLoginWhiteHostListResponse() (response *DescribeLoginWhiteHostListResponse) {
    response = &DescribeLoginWhiteHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoginWhiteHostList
// 查询合并后白名单机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLoginWhiteHostList(request *DescribeLoginWhiteHostListRequest) (response *DescribeLoginWhiteHostListResponse, err error) {
    return c.DescribeLoginWhiteHostListWithContext(context.Background(), request)
}

// DescribeLoginWhiteHostList
// 查询合并后白名单机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLoginWhiteHostListWithContext(ctx context.Context, request *DescribeLoginWhiteHostListRequest) (response *DescribeLoginWhiteHostListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginWhiteHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginWhiteHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
    request = &DescribeLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeLoginWhiteList")
    
    
    return
}

func NewDescribeLoginWhiteListResponse() (response *DescribeLoginWhiteListResponse) {
    response = &DescribeLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoginWhiteList
// 获取异地登录白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLoginWhiteList(request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    return c.DescribeLoginWhiteListWithContext(context.Background(), request)
}

// DescribeLoginWhiteList
// 获取异地登录白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLoginWhiteListWithContext(ctx context.Context, request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineClearHistoryRequest() (request *DescribeMachineClearHistoryRequest) {
    request = &DescribeMachineClearHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineClearHistory")
    
    
    return
}

func NewDescribeMachineClearHistoryResponse() (response *DescribeMachineClearHistoryResponse) {
    response = &DescribeMachineClearHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineClearHistory
// 查询机器清理历史记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineClearHistory(request *DescribeMachineClearHistoryRequest) (response *DescribeMachineClearHistoryResponse, err error) {
    return c.DescribeMachineClearHistoryWithContext(context.Background(), request)
}

// DescribeMachineClearHistory
// 查询机器清理历史记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineClearHistoryWithContext(ctx context.Context, request *DescribeMachineClearHistoryRequest) (response *DescribeMachineClearHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeMachineClearHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineClearHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineClearHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineDefenseCntRequest() (request *DescribeMachineDefenseCntRequest) {
    request = &DescribeMachineDefenseCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineDefenseCnt")
    
    
    return
}

func NewDescribeMachineDefenseCntResponse() (response *DescribeMachineDefenseCntResponse) {
    response = &DescribeMachineDefenseCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineDefenseCnt
// 查询主机高级防御事件数统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineDefenseCnt(request *DescribeMachineDefenseCntRequest) (response *DescribeMachineDefenseCntResponse, err error) {
    return c.DescribeMachineDefenseCntWithContext(context.Background(), request)
}

// DescribeMachineDefenseCnt
// 查询主机高级防御事件数统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineDefenseCntWithContext(ctx context.Context, request *DescribeMachineDefenseCntRequest) (response *DescribeMachineDefenseCntResponse, err error) {
    if request == nil {
        request = NewDescribeMachineDefenseCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineDefenseCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineDefenseCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineFileTamperRulesRequest() (request *DescribeMachineFileTamperRulesRequest) {
    request = &DescribeMachineFileTamperRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineFileTamperRules")
    
    
    return
}

func NewDescribeMachineFileTamperRulesResponse() (response *DescribeMachineFileTamperRulesResponse) {
    response = &DescribeMachineFileTamperRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineFileTamperRules
// 查询主机相关核心文件监控规则列 表   
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineFileTamperRules(request *DescribeMachineFileTamperRulesRequest) (response *DescribeMachineFileTamperRulesResponse, err error) {
    return c.DescribeMachineFileTamperRulesWithContext(context.Background(), request)
}

// DescribeMachineFileTamperRules
// 查询主机相关核心文件监控规则列 表   
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineFileTamperRulesWithContext(ctx context.Context, request *DescribeMachineFileTamperRulesRequest) (response *DescribeMachineFileTamperRulesResponse, err error) {
    if request == nil {
        request = NewDescribeMachineFileTamperRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineFileTamperRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineFileTamperRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineGeneralRequest() (request *DescribeMachineGeneralRequest) {
    request = &DescribeMachineGeneralRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineGeneral")
    
    
    return
}

func NewDescribeMachineGeneralResponse() (response *DescribeMachineGeneralResponse) {
    response = &DescribeMachineGeneralResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineGeneral
// 查询主机概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMachineGeneral(request *DescribeMachineGeneralRequest) (response *DescribeMachineGeneralResponse, err error) {
    return c.DescribeMachineGeneralWithContext(context.Background(), request)
}

// DescribeMachineGeneral
// 查询主机概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMachineGeneralWithContext(ctx context.Context, request *DescribeMachineGeneralRequest) (response *DescribeMachineGeneralResponse, err error) {
    if request == nil {
        request = NewDescribeMachineGeneralRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineGeneral require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineGeneralResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineInfoRequest() (request *DescribeMachineInfoRequest) {
    request = &DescribeMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineInfo")
    
    
    return
}

func NewDescribeMachineInfoResponse() (response *DescribeMachineInfoResponse) {
    response = &DescribeMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineInfo
// 本接口（DescribeMachineInfo）用于获取机器详细信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineInfo(request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    return c.DescribeMachineInfoWithContext(context.Background(), request)
}

// DescribeMachineInfo
// 本接口（DescribeMachineInfo）用于获取机器详细信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineInfoWithContext(ctx context.Context, request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMachineInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineLicenseDetailRequest() (request *DescribeMachineLicenseDetailRequest) {
    request = &DescribeMachineLicenseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineLicenseDetail")
    
    
    return
}

func NewDescribeMachineLicenseDetailResponse() (response *DescribeMachineLicenseDetailResponse) {
    response = &DescribeMachineLicenseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineLicenseDetail
// 本接口 (DescribeMachineLicenseDetail)查询机器授权信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineLicenseDetail(request *DescribeMachineLicenseDetailRequest) (response *DescribeMachineLicenseDetailResponse, err error) {
    return c.DescribeMachineLicenseDetailWithContext(context.Background(), request)
}

// DescribeMachineLicenseDetail
// 本接口 (DescribeMachineLicenseDetail)查询机器授权信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineLicenseDetailWithContext(ctx context.Context, request *DescribeMachineLicenseDetailRequest) (response *DescribeMachineLicenseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMachineLicenseDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineLicenseDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineLicenseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineListRequest() (request *DescribeMachineListRequest) {
    request = &DescribeMachineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineList")
    
    
    return
}

func NewDescribeMachineListResponse() (response *DescribeMachineListResponse) {
    response = &DescribeMachineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineList
// 用于网页防篡改获取区域主机列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineList(request *DescribeMachineListRequest) (response *DescribeMachineListResponse, err error) {
    return c.DescribeMachineListWithContext(context.Background(), request)
}

// DescribeMachineList
// 用于网页防篡改获取区域主机列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineListWithContext(ctx context.Context, request *DescribeMachineListRequest) (response *DescribeMachineListResponse, err error) {
    if request == nil {
        request = NewDescribeMachineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineOsListRequest() (request *DescribeMachineOsListRequest) {
    request = &DescribeMachineOsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineOsList")
    
    
    return
}

func NewDescribeMachineOsListResponse() (response *DescribeMachineOsListResponse) {
    response = &DescribeMachineOsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineOsList
// 查询可筛选操作系统列表.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMachineOsList(request *DescribeMachineOsListRequest) (response *DescribeMachineOsListResponse, err error) {
    return c.DescribeMachineOsListWithContext(context.Background(), request)
}

// DescribeMachineOsList
// 查询可筛选操作系统列表.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMachineOsListWithContext(ctx context.Context, request *DescribeMachineOsListRequest) (response *DescribeMachineOsListResponse, err error) {
    if request == nil {
        request = NewDescribeMachineOsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineOsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineOsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRegionListRequest() (request *DescribeMachineRegionListRequest) {
    request = &DescribeMachineRegionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRegionList")
    
    
    return
}

func NewDescribeMachineRegionListResponse() (response *DescribeMachineRegionListResponse) {
    response = &DescribeMachineRegionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineRegionList
// 查询主机地域列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineRegionList(request *DescribeMachineRegionListRequest) (response *DescribeMachineRegionListResponse, err error) {
    return c.DescribeMachineRegionListWithContext(context.Background(), request)
}

// DescribeMachineRegionList
// 查询主机地域列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachineRegionListWithContext(ctx context.Context, request *DescribeMachineRegionListRequest) (response *DescribeMachineRegionListResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRegionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineRegionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineRegionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRegionsRequest() (request *DescribeMachineRegionsRequest) {
    request = &DescribeMachineRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRegions")
    
    
    return
}

func NewDescribeMachineRegionsResponse() (response *DescribeMachineRegionsResponse) {
    response = &DescribeMachineRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineRegions
// 获取机器地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineRegions(request *DescribeMachineRegionsRequest) (response *DescribeMachineRegionsResponse, err error) {
    return c.DescribeMachineRegionsWithContext(context.Background(), request)
}

// DescribeMachineRegions
// 获取机器地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineRegionsWithContext(ctx context.Context, request *DescribeMachineRegionsRequest) (response *DescribeMachineRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRiskCntRequest() (request *DescribeMachineRiskCntRequest) {
    request = &DescribeMachineRiskCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineRiskCnt")
    
    
    return
}

func NewDescribeMachineRiskCntResponse() (response *DescribeMachineRiskCntResponse) {
    response = &DescribeMachineRiskCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineRiskCnt
// 查询主机入侵检测事件统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineRiskCnt(request *DescribeMachineRiskCntRequest) (response *DescribeMachineRiskCntResponse, err error) {
    return c.DescribeMachineRiskCntWithContext(context.Background(), request)
}

// DescribeMachineRiskCnt
// 查询主机入侵检测事件统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachineRiskCntWithContext(ctx context.Context, request *DescribeMachineRiskCntRequest) (response *DescribeMachineRiskCntResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRiskCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineRiskCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineRiskCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineSnapshotRequest() (request *DescribeMachineSnapshotRequest) {
    request = &DescribeMachineSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachineSnapshot")
    
    
    return
}

func NewDescribeMachineSnapshotResponse() (response *DescribeMachineSnapshotResponse) {
    response = &DescribeMachineSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineSnapshot
// 漏洞修护-查询主机创建的快照
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMachineSnapshot(request *DescribeMachineSnapshotRequest) (response *DescribeMachineSnapshotResponse, err error) {
    return c.DescribeMachineSnapshotWithContext(context.Background(), request)
}

// DescribeMachineSnapshot
// 漏洞修护-查询主机创建的快照
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMachineSnapshotWithContext(ctx context.Context, request *DescribeMachineSnapshotRequest) (response *DescribeMachineSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeMachineSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachines")
    
    
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachines
// 本接口 (DescribeMachines) 用于获取区域主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    return c.DescribeMachinesWithContext(context.Background(), request)
}

// DescribeMachines
// 本接口 (DescribeMachines) 用于获取区域主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachinesWithContext(ctx context.Context, request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesSimpleRequest() (request *DescribeMachinesSimpleRequest) {
    request = &DescribeMachinesSimpleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMachinesSimple")
    
    
    return
}

func NewDescribeMachinesSimpleResponse() (response *DescribeMachinesSimpleResponse) {
    response = &DescribeMachinesSimpleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachinesSimple
// 本接口 (DescribeMachinesSimple) 用于获取主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachinesSimple(request *DescribeMachinesSimpleRequest) (response *DescribeMachinesSimpleResponse, err error) {
    return c.DescribeMachinesSimpleWithContext(context.Background(), request)
}

// DescribeMachinesSimple
// 本接口 (DescribeMachinesSimple) 用于获取主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMachinesSimpleWithContext(ctx context.Context, request *DescribeMachinesSimpleRequest) (response *DescribeMachinesSimpleResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesSimpleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachinesSimple require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesSimpleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalWareListRequest() (request *DescribeMalWareListRequest) {
    request = &DescribeMalWareListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalWareList")
    
    
    return
}

func NewDescribeMalWareListResponse() (response *DescribeMalWareListResponse) {
    response = &DescribeMalWareListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalWareList
// 入侵检测获取木马列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMalWareList(request *DescribeMalWareListRequest) (response *DescribeMalWareListResponse, err error) {
    return c.DescribeMalWareListWithContext(context.Background(), request)
}

// DescribeMalWareList
// 入侵检测获取木马列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMalWareListWithContext(ctx context.Context, request *DescribeMalWareListRequest) (response *DescribeMalWareListResponse, err error) {
    if request == nil {
        request = NewDescribeMalWareListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalWareList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalWareListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaliciousRequestWhiteListRequest() (request *DescribeMaliciousRequestWhiteListRequest) {
    request = &DescribeMaliciousRequestWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaliciousRequestWhiteList")
    
    
    return
}

func NewDescribeMaliciousRequestWhiteListResponse() (response *DescribeMaliciousRequestWhiteListResponse) {
    response = &DescribeMaliciousRequestWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaliciousRequestWhiteList
// 查询恶意请求白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMaliciousRequestWhiteList(request *DescribeMaliciousRequestWhiteListRequest) (response *DescribeMaliciousRequestWhiteListResponse, err error) {
    return c.DescribeMaliciousRequestWhiteListWithContext(context.Background(), request)
}

// DescribeMaliciousRequestWhiteList
// 查询恶意请求白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMaliciousRequestWhiteListWithContext(ctx context.Context, request *DescribeMaliciousRequestWhiteListRequest) (response *DescribeMaliciousRequestWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaliciousRequestWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaliciousRequestWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareFileRequest() (request *DescribeMalwareFileRequest) {
    request = &DescribeMalwareFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareFile")
    
    
    return
}

func NewDescribeMalwareFileResponse() (response *DescribeMalwareFileResponse) {
    response = &DescribeMalwareFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareFile
// 获取木马文件下载地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareFile(request *DescribeMalwareFileRequest) (response *DescribeMalwareFileResponse, err error) {
    return c.DescribeMalwareFileWithContext(context.Background(), request)
}

// DescribeMalwareFile
// 获取木马文件下载地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareFileWithContext(ctx context.Context, request *DescribeMalwareFileRequest) (response *DescribeMalwareFileResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareInfoRequest() (request *DescribeMalwareInfoRequest) {
    request = &DescribeMalwareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareInfo")
    
    
    return
}

func NewDescribeMalwareInfoResponse() (response *DescribeMalwareInfoResponse) {
    response = &DescribeMalwareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareInfo
// 查看恶意文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareInfo(request *DescribeMalwareInfoRequest) (response *DescribeMalwareInfoResponse, err error) {
    return c.DescribeMalwareInfoWithContext(context.Background(), request)
}

// DescribeMalwareInfo
// 查看恶意文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareInfoWithContext(ctx context.Context, request *DescribeMalwareInfoRequest) (response *DescribeMalwareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareRiskOverviewRequest() (request *DescribeMalwareRiskOverviewRequest) {
    request = &DescribeMalwareRiskOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareRiskOverview")
    
    
    return
}

func NewDescribeMalwareRiskOverviewResponse() (response *DescribeMalwareRiskOverviewResponse) {
    response = &DescribeMalwareRiskOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareRiskOverview
// 获取文件查杀概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareRiskOverview(request *DescribeMalwareRiskOverviewRequest) (response *DescribeMalwareRiskOverviewResponse, err error) {
    return c.DescribeMalwareRiskOverviewWithContext(context.Background(), request)
}

// DescribeMalwareRiskOverview
// 获取文件查杀概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareRiskOverviewWithContext(ctx context.Context, request *DescribeMalwareRiskOverviewRequest) (response *DescribeMalwareRiskOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareRiskOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareRiskOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareRiskOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareRiskWarningRequest() (request *DescribeMalwareRiskWarningRequest) {
    request = &DescribeMalwareRiskWarningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareRiskWarning")
    
    
    return
}

func NewDescribeMalwareRiskWarningResponse() (response *DescribeMalwareRiskWarningResponse) {
    response = &DescribeMalwareRiskWarningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareRiskWarning
// 打开入侵检测-恶意文件检测,弹出风险预警内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareRiskWarning(request *DescribeMalwareRiskWarningRequest) (response *DescribeMalwareRiskWarningResponse, err error) {
    return c.DescribeMalwareRiskWarningWithContext(context.Background(), request)
}

// DescribeMalwareRiskWarning
// 打开入侵检测-恶意文件检测,弹出风险预警内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareRiskWarningWithContext(ctx context.Context, request *DescribeMalwareRiskWarningRequest) (response *DescribeMalwareRiskWarningResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareRiskWarningRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareRiskWarning require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareRiskWarningResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareTimingScanSettingRequest() (request *DescribeMalwareTimingScanSettingRequest) {
    request = &DescribeMalwareTimingScanSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareTimingScanSetting")
    
    
    return
}

func NewDescribeMalwareTimingScanSettingResponse() (response *DescribeMalwareTimingScanSettingResponse) {
    response = &DescribeMalwareTimingScanSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareTimingScanSetting
// 查询定时扫描配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareTimingScanSetting(request *DescribeMalwareTimingScanSettingRequest) (response *DescribeMalwareTimingScanSettingResponse, err error) {
    return c.DescribeMalwareTimingScanSettingWithContext(context.Background(), request)
}

// DescribeMalwareTimingScanSetting
// 查询定时扫描配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeMalwareTimingScanSettingWithContext(ctx context.Context, request *DescribeMalwareTimingScanSettingRequest) (response *DescribeMalwareTimingScanSettingResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareTimingScanSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareTimingScanSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareTimingScanSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareWhiteListRequest() (request *DescribeMalwareWhiteListRequest) {
    request = &DescribeMalwareWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareWhiteList")
    
    
    return
}

func NewDescribeMalwareWhiteListResponse() (response *DescribeMalwareWhiteListResponse) {
    response = &DescribeMalwareWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareWhiteList
// 获取木马白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareWhiteList(request *DescribeMalwareWhiteListRequest) (response *DescribeMalwareWhiteListResponse, err error) {
    return c.DescribeMalwareWhiteListWithContext(context.Background(), request)
}

// DescribeMalwareWhiteList
// 获取木马白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareWhiteListWithContext(ctx context.Context, request *DescribeMalwareWhiteListRequest) (response *DescribeMalwareWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwareWhiteListAffectListRequest() (request *DescribeMalwareWhiteListAffectListRequest) {
    request = &DescribeMalwareWhiteListAffectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwareWhiteListAffectList")
    
    
    return
}

func NewDescribeMalwareWhiteListAffectListResponse() (response *DescribeMalwareWhiteListAffectListResponse) {
    response = &DescribeMalwareWhiteListAffectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwareWhiteListAffectList
// 获取木马白名单受影响列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareWhiteListAffectList(request *DescribeMalwareWhiteListAffectListRequest) (response *DescribeMalwareWhiteListAffectListResponse, err error) {
    return c.DescribeMalwareWhiteListAffectListWithContext(context.Background(), request)
}

// DescribeMalwareWhiteListAffectList
// 获取木马白名单受影响列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMalwareWhiteListAffectListWithContext(ctx context.Context, request *DescribeMalwareWhiteListAffectListRequest) (response *DescribeMalwareWhiteListAffectListResponse, err error) {
    if request == nil {
        request = NewDescribeMalwareWhiteListAffectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwareWhiteListAffectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwareWhiteListAffectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonthInspectionReportRequest() (request *DescribeMonthInspectionReportRequest) {
    request = &DescribeMonthInspectionReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMonthInspectionReport")
    
    
    return
}

func NewDescribeMonthInspectionReportResponse() (response *DescribeMonthInspectionReportResponse) {
    response = &DescribeMonthInspectionReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMonthInspectionReport
// 专家服务-安全管家月巡检报告下载
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMonthInspectionReport(request *DescribeMonthInspectionReportRequest) (response *DescribeMonthInspectionReportResponse, err error) {
    return c.DescribeMonthInspectionReportWithContext(context.Background(), request)
}

// DescribeMonthInspectionReport
// 专家服务-安全管家月巡检报告下载
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMonthInspectionReportWithContext(ctx context.Context, request *DescribeMonthInspectionReportRequest) (response *DescribeMonthInspectionReportResponse, err error) {
    if request == nil {
        request = NewDescribeMonthInspectionReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonthInspectionReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonthInspectionReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetAttackSettingRequest() (request *DescribeNetAttackSettingRequest) {
    request = &DescribeNetAttackSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeNetAttackSetting")
    
    
    return
}

func NewDescribeNetAttackSettingResponse() (response *DescribeNetAttackSettingResponse) {
    response = &DescribeNetAttackSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetAttackSetting
// 查询网络攻击设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetAttackSetting(request *DescribeNetAttackSettingRequest) (response *DescribeNetAttackSettingResponse, err error) {
    return c.DescribeNetAttackSettingWithContext(context.Background(), request)
}

// DescribeNetAttackSetting
// 查询网络攻击设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetAttackSettingWithContext(ctx context.Context, request *DescribeNetAttackSettingRequest) (response *DescribeNetAttackSettingResponse, err error) {
    if request == nil {
        request = NewDescribeNetAttackSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetAttackSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetAttackSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetAttackWhiteListRequest() (request *DescribeNetAttackWhiteListRequest) {
    request = &DescribeNetAttackWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeNetAttackWhiteList")
    
    
    return
}

func NewDescribeNetAttackWhiteListResponse() (response *DescribeNetAttackWhiteListResponse) {
    response = &DescribeNetAttackWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetAttackWhiteList
// 获取网络攻击白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetAttackWhiteList(request *DescribeNetAttackWhiteListRequest) (response *DescribeNetAttackWhiteListResponse, err error) {
    return c.DescribeNetAttackWhiteListWithContext(context.Background(), request)
}

// DescribeNetAttackWhiteList
// 获取网络攻击白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetAttackWhiteListWithContext(ctx context.Context, request *DescribeNetAttackWhiteListRequest) (response *DescribeNetAttackWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeNetAttackWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetAttackWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetAttackWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
    request = &DescribeOpenPortStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeOpenPortStatistics")
    
    
    return
}

func NewDescribeOpenPortStatisticsResponse() (response *DescribeOpenPortStatisticsResponse) {
    response = &DescribeOpenPortStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPortStatistics
// 本接口 (DescribeOpenPortStatistics) 用于获取端口统计列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    return c.DescribeOpenPortStatisticsWithContext(context.Background(), request)
}

// DescribeOpenPortStatistics
// 本接口 (DescribeOpenPortStatistics) 用于获取端口统计列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPortStatisticsWithContext(ctx context.Context, request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPortStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
    request = &DescribeOverviewStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeOverviewStatistics")
    
    
    return
}

func NewDescribeOverviewStatisticsResponse() (response *DescribeOverviewStatisticsResponse) {
    response = &DescribeOverviewStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewStatistics
// 获取概览统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    return c.DescribeOverviewStatisticsWithContext(context.Background(), request)
}

// DescribeOverviewStatistics
// 获取概览统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOverviewStatisticsWithContext(ctx context.Context, request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivilegeEventInfoRequest() (request *DescribePrivilegeEventInfoRequest) {
    request = &DescribePrivilegeEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeEventInfo")
    
    
    return
}

func NewDescribePrivilegeEventInfoResponse() (response *DescribePrivilegeEventInfoResponse) {
    response = &DescribePrivilegeEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivilegeEventInfo
// 本地提权信息详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePrivilegeEventInfo(request *DescribePrivilegeEventInfoRequest) (response *DescribePrivilegeEventInfoResponse, err error) {
    return c.DescribePrivilegeEventInfoWithContext(context.Background(), request)
}

// DescribePrivilegeEventInfo
// 本地提权信息详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribePrivilegeEventInfoWithContext(ctx context.Context, request *DescribePrivilegeEventInfoRequest) (response *DescribePrivilegeEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribePrivilegeEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivilegeEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivilegeEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivilegeEventsRequest() (request *DescribePrivilegeEventsRequest) {
    request = &DescribePrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeEvents")
    
    
    return
}

func NewDescribePrivilegeEventsResponse() (response *DescribePrivilegeEventsResponse) {
    response = &DescribePrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivilegeEvents
// 获取本地提权事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrivilegeEvents(request *DescribePrivilegeEventsRequest) (response *DescribePrivilegeEventsResponse, err error) {
    return c.DescribePrivilegeEventsWithContext(context.Background(), request)
}

// DescribePrivilegeEvents
// 获取本地提权事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrivilegeEventsWithContext(ctx context.Context, request *DescribePrivilegeEventsRequest) (response *DescribePrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewDescribePrivilegeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivilegeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivilegeRulesRequest() (request *DescribePrivilegeRulesRequest) {
    request = &DescribePrivilegeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribePrivilegeRules")
    
    
    return
}

func NewDescribePrivilegeRulesResponse() (response *DescribePrivilegeRulesResponse) {
    response = &DescribePrivilegeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivilegeRules
// 获取本地提权规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrivilegeRules(request *DescribePrivilegeRulesRequest) (response *DescribePrivilegeRulesResponse, err error) {
    return c.DescribePrivilegeRulesWithContext(context.Background(), request)
}

// DescribePrivilegeRules
// 获取本地提权规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrivilegeRulesWithContext(ctx context.Context, request *DescribePrivilegeRulesRequest) (response *DescribePrivilegeRulesResponse, err error) {
    if request == nil {
        request = NewDescribePrivilegeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivilegeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivilegeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProVersionInfo")
    
    
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProVersionInfo
// 用于获取专业版概览信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    return c.DescribeProVersionInfoWithContext(context.Background(), request)
}

// DescribeProVersionInfo
// 用于获取专业版概览信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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

func NewDescribeProVersionStatusRequest() (request *DescribeProVersionStatusRequest) {
    request = &DescribeProVersionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProVersionStatus")
    
    
    return
}

func NewDescribeProVersionStatusResponse() (response *DescribeProVersionStatusResponse) {
    response = &DescribeProVersionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProVersionStatus
// 用于获取单台主机或所有主机是否开通专业版状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProVersionStatus(request *DescribeProVersionStatusRequest) (response *DescribeProVersionStatusResponse, err error) {
    return c.DescribeProVersionStatusWithContext(context.Background(), request)
}

// DescribeProVersionStatus
// 用于获取单台主机或所有主机是否开通专业版状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProVersionStatusWithContext(ctx context.Context, request *DescribeProVersionStatusRequest) (response *DescribeProVersionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProVersionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProVersionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
    request = &DescribeProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProcessStatistics")
    
    
    return
}

func NewDescribeProcessStatisticsResponse() (response *DescribeProcessStatisticsResponse) {
    response = &DescribeProcessStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcessStatistics
// 本接口 (DescribeProcessStatistics) 用于获取进程统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    return c.DescribeProcessStatisticsWithContext(context.Background(), request)
}

// DescribeProcessStatistics
// 本接口 (DescribeProcessStatistics) 用于获取进程统计列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessStatisticsWithContext(ctx context.Context, request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductStatusRequest() (request *DescribeProductStatusRequest) {
    request = &DescribeProductStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProductStatus")
    
    
    return
}

func NewDescribeProductStatusResponse() (response *DescribeProductStatusResponse) {
    response = &DescribeProductStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductStatus
// 产品试用状态查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProductStatus(request *DescribeProductStatusRequest) (response *DescribeProductStatusResponse, err error) {
    return c.DescribeProductStatusWithContext(context.Background(), request)
}

// DescribeProductStatus
// 产品试用状态查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProductStatusWithContext(ctx context.Context, request *DescribeProductStatusRequest) (response *DescribeProductStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProductStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectDirListRequest() (request *DescribeProtectDirListRequest) {
    request = &DescribeProtectDirListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProtectDirList")
    
    
    return
}

func NewDescribeProtectDirListResponse() (response *DescribeProtectDirListResponse) {
    response = &DescribeProtectDirListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProtectDirList
// 网页防篡改防护目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectDirList(request *DescribeProtectDirListRequest) (response *DescribeProtectDirListResponse, err error) {
    return c.DescribeProtectDirListWithContext(context.Background(), request)
}

// DescribeProtectDirList
// 网页防篡改防护目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectDirListWithContext(ctx context.Context, request *DescribeProtectDirListRequest) (response *DescribeProtectDirListResponse, err error) {
    if request == nil {
        request = NewDescribeProtectDirListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectDirList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectDirListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectDirRelatedServerRequest() (request *DescribeProtectDirRelatedServerRequest) {
    request = &DescribeProtectDirRelatedServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProtectDirRelatedServer")
    
    
    return
}

func NewDescribeProtectDirRelatedServerResponse() (response *DescribeProtectDirRelatedServerResponse) {
    response = &DescribeProtectDirRelatedServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProtectDirRelatedServer
// 查询防护目录关联服务器列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectDirRelatedServer(request *DescribeProtectDirRelatedServerRequest) (response *DescribeProtectDirRelatedServerResponse, err error) {
    return c.DescribeProtectDirRelatedServerWithContext(context.Background(), request)
}

// DescribeProtectDirRelatedServer
// 查询防护目录关联服务器列表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectDirRelatedServerWithContext(ctx context.Context, request *DescribeProtectDirRelatedServerRequest) (response *DescribeProtectDirRelatedServerResponse, err error) {
    if request == nil {
        request = NewDescribeProtectDirRelatedServerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectDirRelatedServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectDirRelatedServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectNetListRequest() (request *DescribeProtectNetListRequest) {
    request = &DescribeProtectNetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProtectNetList")
    
    
    return
}

func NewDescribeProtectNetListResponse() (response *DescribeProtectNetListResponse) {
    response = &DescribeProtectNetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProtectNetList
// 专家服务-旗舰重保列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectNetList(request *DescribeProtectNetListRequest) (response *DescribeProtectNetListResponse, err error) {
    return c.DescribeProtectNetListWithContext(context.Background(), request)
}

// DescribeProtectNetList
// 专家服务-旗舰重保列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectNetListWithContext(ctx context.Context, request *DescribeProtectNetListRequest) (response *DescribeProtectNetListResponse, err error) {
    if request == nil {
        request = NewDescribeProtectNetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectNetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectNetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicProxyInstallCommandRequest() (request *DescribePublicProxyInstallCommandRequest) {
    request = &DescribePublicProxyInstallCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribePublicProxyInstallCommand")
    
    
    return
}

func NewDescribePublicProxyInstallCommandResponse() (response *DescribePublicProxyInstallCommandResponse) {
    response = &DescribePublicProxyInstallCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicProxyInstallCommand
// 获取公网接入代理安装命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribePublicProxyInstallCommand(request *DescribePublicProxyInstallCommandRequest) (response *DescribePublicProxyInstallCommandResponse, err error) {
    return c.DescribePublicProxyInstallCommandWithContext(context.Background(), request)
}

// DescribePublicProxyInstallCommand
// 获取公网接入代理安装命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribePublicProxyInstallCommandWithContext(ctx context.Context, request *DescribePublicProxyInstallCommandRequest) (response *DescribePublicProxyInstallCommandResponse, err error) {
    if request == nil {
        request = NewDescribePublicProxyInstallCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicProxyInstallCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicProxyInstallCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseBackupListRequest() (request *DescribeRansomDefenseBackupListRequest) {
    request = &DescribeRansomDefenseBackupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseBackupList")
    
    
    return
}

func NewDescribeRansomDefenseBackupListResponse() (response *DescribeRansomDefenseBackupListResponse) {
    response = &DescribeRansomDefenseBackupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseBackupList
// 查询主机快照备份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseBackupList(request *DescribeRansomDefenseBackupListRequest) (response *DescribeRansomDefenseBackupListResponse, err error) {
    return c.DescribeRansomDefenseBackupListWithContext(context.Background(), request)
}

// DescribeRansomDefenseBackupList
// 查询主机快照备份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseBackupListWithContext(ctx context.Context, request *DescribeRansomDefenseBackupListRequest) (response *DescribeRansomDefenseBackupListResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseBackupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseBackupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseBackupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseEventsListRequest() (request *DescribeRansomDefenseEventsListRequest) {
    request = &DescribeRansomDefenseEventsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseEventsList")
    
    
    return
}

func NewDescribeRansomDefenseEventsListResponse() (response *DescribeRansomDefenseEventsListResponse) {
    response = &DescribeRansomDefenseEventsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseEventsList
// 查询防勒索事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseEventsList(request *DescribeRansomDefenseEventsListRequest) (response *DescribeRansomDefenseEventsListResponse, err error) {
    return c.DescribeRansomDefenseEventsListWithContext(context.Background(), request)
}

// DescribeRansomDefenseEventsList
// 查询防勒索事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseEventsListWithContext(ctx context.Context, request *DescribeRansomDefenseEventsListRequest) (response *DescribeRansomDefenseEventsListResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseEventsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseEventsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseEventsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseMachineListRequest() (request *DescribeRansomDefenseMachineListRequest) {
    request = &DescribeRansomDefenseMachineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseMachineList")
    
    
    return
}

func NewDescribeRansomDefenseMachineListResponse() (response *DescribeRansomDefenseMachineListResponse) {
    response = &DescribeRansomDefenseMachineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseMachineList
// 查询备份详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseMachineList(request *DescribeRansomDefenseMachineListRequest) (response *DescribeRansomDefenseMachineListResponse, err error) {
    return c.DescribeRansomDefenseMachineListWithContext(context.Background(), request)
}

// DescribeRansomDefenseMachineList
// 查询备份详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseMachineListWithContext(ctx context.Context, request *DescribeRansomDefenseMachineListRequest) (response *DescribeRansomDefenseMachineListResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseMachineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseMachineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseMachineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseMachineStrategyInfoRequest() (request *DescribeRansomDefenseMachineStrategyInfoRequest) {
    request = &DescribeRansomDefenseMachineStrategyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseMachineStrategyInfo")
    
    
    return
}

func NewDescribeRansomDefenseMachineStrategyInfoResponse() (response *DescribeRansomDefenseMachineStrategyInfoResponse) {
    response = &DescribeRansomDefenseMachineStrategyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseMachineStrategyInfo
// 获取主机绑定策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseMachineStrategyInfo(request *DescribeRansomDefenseMachineStrategyInfoRequest) (response *DescribeRansomDefenseMachineStrategyInfoResponse, err error) {
    return c.DescribeRansomDefenseMachineStrategyInfoWithContext(context.Background(), request)
}

// DescribeRansomDefenseMachineStrategyInfo
// 获取主机绑定策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseMachineStrategyInfoWithContext(ctx context.Context, request *DescribeRansomDefenseMachineStrategyInfoRequest) (response *DescribeRansomDefenseMachineStrategyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseMachineStrategyInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseMachineStrategyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseMachineStrategyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseRollBackTaskListRequest() (request *DescribeRansomDefenseRollBackTaskListRequest) {
    request = &DescribeRansomDefenseRollBackTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseRollBackTaskList")
    
    
    return
}

func NewDescribeRansomDefenseRollBackTaskListResponse() (response *DescribeRansomDefenseRollBackTaskListResponse) {
    response = &DescribeRansomDefenseRollBackTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseRollBackTaskList
// 查询回滚任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseRollBackTaskList(request *DescribeRansomDefenseRollBackTaskListRequest) (response *DescribeRansomDefenseRollBackTaskListResponse, err error) {
    return c.DescribeRansomDefenseRollBackTaskListWithContext(context.Background(), request)
}

// DescribeRansomDefenseRollBackTaskList
// 查询回滚任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseRollBackTaskListWithContext(ctx context.Context, request *DescribeRansomDefenseRollBackTaskListRequest) (response *DescribeRansomDefenseRollBackTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseRollBackTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseRollBackTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseRollBackTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseStateRequest() (request *DescribeRansomDefenseStateRequest) {
    request = &DescribeRansomDefenseStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseState")
    
    
    return
}

func NewDescribeRansomDefenseStateResponse() (response *DescribeRansomDefenseStateResponse) {
    response = &DescribeRansomDefenseStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseState
// 获取用户防勒索趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseState(request *DescribeRansomDefenseStateRequest) (response *DescribeRansomDefenseStateResponse, err error) {
    return c.DescribeRansomDefenseStateWithContext(context.Background(), request)
}

// DescribeRansomDefenseState
// 获取用户防勒索趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseStateWithContext(ctx context.Context, request *DescribeRansomDefenseStateRequest) (response *DescribeRansomDefenseStateResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseStrategyDetailRequest() (request *DescribeRansomDefenseStrategyDetailRequest) {
    request = &DescribeRansomDefenseStrategyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyDetail")
    
    
    return
}

func NewDescribeRansomDefenseStrategyDetailResponse() (response *DescribeRansomDefenseStrategyDetailResponse) {
    response = &DescribeRansomDefenseStrategyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseStrategyDetail
// 获取策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseStrategyDetail(request *DescribeRansomDefenseStrategyDetailRequest) (response *DescribeRansomDefenseStrategyDetailResponse, err error) {
    return c.DescribeRansomDefenseStrategyDetailWithContext(context.Background(), request)
}

// DescribeRansomDefenseStrategyDetail
// 获取策略详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseStrategyDetailWithContext(ctx context.Context, request *DescribeRansomDefenseStrategyDetailRequest) (response *DescribeRansomDefenseStrategyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseStrategyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseStrategyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseStrategyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseStrategyListRequest() (request *DescribeRansomDefenseStrategyListRequest) {
    request = &DescribeRansomDefenseStrategyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyList")
    
    
    return
}

func NewDescribeRansomDefenseStrategyListResponse() (response *DescribeRansomDefenseStrategyListResponse) {
    response = &DescribeRansomDefenseStrategyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseStrategyList
// 查询防勒索策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseStrategyList(request *DescribeRansomDefenseStrategyListRequest) (response *DescribeRansomDefenseStrategyListResponse, err error) {
    return c.DescribeRansomDefenseStrategyListWithContext(context.Background(), request)
}

// DescribeRansomDefenseStrategyList
// 查询防勒索策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseStrategyListWithContext(ctx context.Context, request *DescribeRansomDefenseStrategyListRequest) (response *DescribeRansomDefenseStrategyListResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseStrategyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseStrategyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseStrategyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseStrategyMachinesRequest() (request *DescribeRansomDefenseStrategyMachinesRequest) {
    request = &DescribeRansomDefenseStrategyMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseStrategyMachines")
    
    
    return
}

func NewDescribeRansomDefenseStrategyMachinesResponse() (response *DescribeRansomDefenseStrategyMachinesResponse) {
    response = &DescribeRansomDefenseStrategyMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseStrategyMachines
// 查询防勒索策略绑定机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseStrategyMachines(request *DescribeRansomDefenseStrategyMachinesRequest) (response *DescribeRansomDefenseStrategyMachinesResponse, err error) {
    return c.DescribeRansomDefenseStrategyMachinesWithContext(context.Background(), request)
}

// DescribeRansomDefenseStrategyMachines
// 查询防勒索策略绑定机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRansomDefenseStrategyMachinesWithContext(ctx context.Context, request *DescribeRansomDefenseStrategyMachinesRequest) (response *DescribeRansomDefenseStrategyMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseStrategyMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseStrategyMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseStrategyMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRansomDefenseTrendRequest() (request *DescribeRansomDefenseTrendRequest) {
    request = &DescribeRansomDefenseTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRansomDefenseTrend")
    
    
    return
}

func NewDescribeRansomDefenseTrendResponse() (response *DescribeRansomDefenseTrendResponse) {
    response = &DescribeRansomDefenseTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRansomDefenseTrend
// 获取全网勒索态势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseTrend(request *DescribeRansomDefenseTrendRequest) (response *DescribeRansomDefenseTrendResponse, err error) {
    return c.DescribeRansomDefenseTrendWithContext(context.Background(), request)
}

// DescribeRansomDefenseTrend
// 获取全网勒索态势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRansomDefenseTrendWithContext(ctx context.Context, request *DescribeRansomDefenseTrendRequest) (response *DescribeRansomDefenseTrendResponse, err error) {
    if request == nil {
        request = NewDescribeRansomDefenseTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRansomDefenseTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRansomDefenseTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecommendedProtectCpuRequest() (request *DescribeRecommendedProtectCpuRequest) {
    request = &DescribeRecommendedProtectCpuRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRecommendedProtectCpu")
    
    
    return
}

func NewDescribeRecommendedProtectCpuResponse() (response *DescribeRecommendedProtectCpuResponse) {
    response = &DescribeRecommendedProtectCpuResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecommendedProtectCpu
// 查询推荐购买防护核数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecommendedProtectCpu(request *DescribeRecommendedProtectCpuRequest) (response *DescribeRecommendedProtectCpuResponse, err error) {
    return c.DescribeRecommendedProtectCpuWithContext(context.Background(), request)
}

// DescribeRecommendedProtectCpu
// 查询推荐购买防护核数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecommendedProtectCpuWithContext(ctx context.Context, request *DescribeRecommendedProtectCpuRequest) (response *DescribeRecommendedProtectCpuResponse, err error) {
    if request == nil {
        request = NewDescribeRecommendedProtectCpuRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecommendedProtectCpu require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecommendedProtectCpuResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventInfoRequest() (request *DescribeReverseShellEventInfoRequest) {
    request = &DescribeReverseShellEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellEventInfo")
    
    
    return
}

func NewDescribeReverseShellEventInfoResponse() (response *DescribeReverseShellEventInfoResponse) {
    response = &DescribeReverseShellEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellEventInfo
// 反弹shell信息详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeReverseShellEventInfo(request *DescribeReverseShellEventInfoRequest) (response *DescribeReverseShellEventInfoResponse, err error) {
    return c.DescribeReverseShellEventInfoWithContext(context.Background(), request)
}

// DescribeReverseShellEventInfo
// 反弹shell信息详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeReverseShellEventInfoWithContext(ctx context.Context, request *DescribeReverseShellEventInfoRequest) (response *DescribeReverseShellEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
    request = &DescribeReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellEvents")
    
    
    return
}

func NewDescribeReverseShellEventsResponse() (response *DescribeReverseShellEventsResponse) {
    response = &DescribeReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellEvents
// 获取反弹Shell列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellEvents(request *DescribeReverseShellEventsRequest) (response *DescribeReverseShellEventsResponse, err error) {
    return c.DescribeReverseShellEventsWithContext(context.Background(), request)
}

// DescribeReverseShellEvents
// 获取反弹Shell列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewDescribeReverseShellRulesRequest() (request *DescribeReverseShellRulesRequest) {
    request = &DescribeReverseShellRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeReverseShellRules")
    
    
    return
}

func NewDescribeReverseShellRulesResponse() (response *DescribeReverseShellRulesResponse) {
    response = &DescribeReverseShellRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReverseShellRules
// 获取反弹Shell规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellRules(request *DescribeReverseShellRulesRequest) (response *DescribeReverseShellRulesResponse, err error) {
    return c.DescribeReverseShellRulesWithContext(context.Background(), request)
}

// DescribeReverseShellRules
// 获取反弹Shell规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReverseShellRulesWithContext(ctx context.Context, request *DescribeReverseShellRulesRequest) (response *DescribeReverseShellRulesResponse, err error) {
    if request == nil {
        request = NewDescribeReverseShellRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReverseShellRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReverseShellRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskBatchStatusRequest() (request *DescribeRiskBatchStatusRequest) {
    request = &DescribeRiskBatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskBatchStatus")
    
    
    return
}

func NewDescribeRiskBatchStatusResponse() (response *DescribeRiskBatchStatusResponse) {
    response = &DescribeRiskBatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskBatchStatus
// 查询入侵检测事件更新状态任务是否完成
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskBatchStatus(request *DescribeRiskBatchStatusRequest) (response *DescribeRiskBatchStatusResponse, err error) {
    return c.DescribeRiskBatchStatusWithContext(context.Background(), request)
}

// DescribeRiskBatchStatus
// 查询入侵检测事件更新状态任务是否完成
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskBatchStatusWithContext(ctx context.Context, request *DescribeRiskBatchStatusRequest) (response *DescribeRiskBatchStatusResponse, err error) {
    if request == nil {
        request = NewDescribeRiskBatchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskBatchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskBatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDnsEventInfoRequest() (request *DescribeRiskDnsEventInfoRequest) {
    request = &DescribeRiskDnsEventInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsEventInfo")
    
    
    return
}

func NewDescribeRiskDnsEventInfoResponse() (response *DescribeRiskDnsEventInfoResponse) {
    response = &DescribeRiskDnsEventInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskDnsEventInfo
// 查询恶意请求事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskDnsEventInfo(request *DescribeRiskDnsEventInfoRequest) (response *DescribeRiskDnsEventInfoResponse, err error) {
    return c.DescribeRiskDnsEventInfoWithContext(context.Background(), request)
}

// DescribeRiskDnsEventInfo
// 查询恶意请求事件详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskDnsEventInfoWithContext(ctx context.Context, request *DescribeRiskDnsEventInfoRequest) (response *DescribeRiskDnsEventInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsEventInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDnsEventInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDnsEventInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDnsEventListRequest() (request *DescribeRiskDnsEventListRequest) {
    request = &DescribeRiskDnsEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsEventList")
    
    
    return
}

func NewDescribeRiskDnsEventListResponse() (response *DescribeRiskDnsEventListResponse) {
    response = &DescribeRiskDnsEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskDnsEventList
// 获取恶意请求事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskDnsEventList(request *DescribeRiskDnsEventListRequest) (response *DescribeRiskDnsEventListResponse, err error) {
    return c.DescribeRiskDnsEventListWithContext(context.Background(), request)
}

// DescribeRiskDnsEventList
// 获取恶意请求事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskDnsEventListWithContext(ctx context.Context, request *DescribeRiskDnsEventListRequest) (response *DescribeRiskDnsEventListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDnsEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDnsEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDnsInfoRequest() (request *DescribeRiskDnsInfoRequest) {
    request = &DescribeRiskDnsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsInfo")
    
    
    return
}

func NewDescribeRiskDnsInfoResponse() (response *DescribeRiskDnsInfoResponse) {
    response = &DescribeRiskDnsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskDnsInfo
// 查询恶意请求详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskDnsInfo(request *DescribeRiskDnsInfoRequest) (response *DescribeRiskDnsInfoResponse, err error) {
    return c.DescribeRiskDnsInfoWithContext(context.Background(), request)
}

// DescribeRiskDnsInfo
// 查询恶意请求详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRiskDnsInfoWithContext(ctx context.Context, request *DescribeRiskDnsInfoRequest) (response *DescribeRiskDnsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDnsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDnsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDnsListRequest() (request *DescribeRiskDnsListRequest) {
    request = &DescribeRiskDnsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsList")
    
    
    return
}

func NewDescribeRiskDnsListResponse() (response *DescribeRiskDnsListResponse) {
    response = &DescribeRiskDnsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskDnsList
// 入侵检测，获取恶意请求列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRiskDnsList(request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
    return c.DescribeRiskDnsListWithContext(context.Background(), request)
}

// DescribeRiskDnsList
// 入侵检测，获取恶意请求列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeRiskDnsListWithContext(ctx context.Context, request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDnsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDnsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskDnsPolicyListRequest() (request *DescribeRiskDnsPolicyListRequest) {
    request = &DescribeRiskDnsPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskDnsPolicyList")
    
    
    return
}

func NewDescribeRiskDnsPolicyListResponse() (response *DescribeRiskDnsPolicyListResponse) {
    response = &DescribeRiskDnsPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskDnsPolicyList
// 获取恶意请求策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskDnsPolicyList(request *DescribeRiskDnsPolicyListRequest) (response *DescribeRiskDnsPolicyListResponse, err error) {
    return c.DescribeRiskDnsPolicyListWithContext(context.Background(), request)
}

// DescribeRiskDnsPolicyList
// 获取恶意请求策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskDnsPolicyListWithContext(ctx context.Context, request *DescribeRiskDnsPolicyListRequest) (response *DescribeRiskDnsPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsPolicyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskDnsPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskDnsPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRiskProcessEventsRequest() (request *DescribeRiskProcessEventsRequest) {
    request = &DescribeRiskProcessEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeRiskProcessEvents")
    
    
    return
}

func NewDescribeRiskProcessEventsResponse() (response *DescribeRiskProcessEventsResponse) {
    response = &DescribeRiskProcessEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRiskProcessEvents
// 获取异常进程列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskProcessEvents(request *DescribeRiskProcessEventsRequest) (response *DescribeRiskProcessEventsResponse, err error) {
    return c.DescribeRiskProcessEventsWithContext(context.Background(), request)
}

// DescribeRiskProcessEvents
// 获取异常进程列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRiskProcessEventsWithContext(ctx context.Context, request *DescribeRiskProcessEventsRequest) (response *DescribeRiskProcessEventsResponse, err error) {
    if request == nil {
        request = NewDescribeRiskProcessEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRiskProcessEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRiskProcessEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeInfoRequest() (request *DescribeSafeInfoRequest) {
    request = &DescribeSafeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSafeInfo")
    
    
    return
}

func NewDescribeSafeInfoResponse() (response *DescribeSafeInfoResponse) {
    response = &DescribeSafeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSafeInfo
// 查询安全通知信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeSafeInfo(request *DescribeSafeInfoRequest) (response *DescribeSafeInfoResponse, err error) {
    return c.DescribeSafeInfoWithContext(context.Background(), request)
}

// DescribeSafeInfo
// 查询安全通知信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeSafeInfoWithContext(ctx context.Context, request *DescribeSafeInfoRequest) (response *DescribeSafeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSafeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanMalwareScheduleRequest() (request *DescribeScanMalwareScheduleRequest) {
    request = &DescribeScanMalwareScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanMalwareSchedule")
    
    
    return
}

func NewDescribeScanMalwareScheduleResponse() (response *DescribeScanMalwareScheduleResponse) {
    response = &DescribeScanMalwareScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanMalwareSchedule
// 查询木马扫描进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanMalwareSchedule(request *DescribeScanMalwareScheduleRequest) (response *DescribeScanMalwareScheduleResponse, err error) {
    return c.DescribeScanMalwareScheduleWithContext(context.Background(), request)
}

// DescribeScanMalwareSchedule
// 查询木马扫描进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanMalwareScheduleWithContext(ctx context.Context, request *DescribeScanMalwareScheduleRequest) (response *DescribeScanMalwareScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeScanMalwareScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanMalwareSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanMalwareScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanScheduleRequest() (request *DescribeScanScheduleRequest) {
    request = &DescribeScanScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanSchedule")
    
    
    return
}

func NewDescribeScanScheduleResponse() (response *DescribeScanScheduleResponse) {
    response = &DescribeScanScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanSchedule
// 根据taskid查询检测进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanSchedule(request *DescribeScanScheduleRequest) (response *DescribeScanScheduleResponse, err error) {
    return c.DescribeScanScheduleWithContext(context.Background(), request)
}

// DescribeScanSchedule
// 根据taskid查询检测进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanScheduleWithContext(ctx context.Context, request *DescribeScanScheduleRequest) (response *DescribeScanScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeScanScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanStateRequest() (request *DescribeScanStateRequest) {
    request = &DescribeScanStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanState")
    
    
    return
}

func NewDescribeScanStateResponse() (response *DescribeScanStateResponse) {
    response = &DescribeScanStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanState
// DescribeScanState 该接口能查询对应模块正在进行的扫描任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanState(request *DescribeScanStateRequest) (response *DescribeScanStateResponse, err error) {
    return c.DescribeScanStateWithContext(context.Background(), request)
}

// DescribeScanState
// DescribeScanState 该接口能查询对应模块正在进行的扫描任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanStateWithContext(ctx context.Context, request *DescribeScanStateRequest) (response *DescribeScanStateResponse, err error) {
    if request == nil {
        request = NewDescribeScanStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskDetailsRequest() (request *DescribeScanTaskDetailsRequest) {
    request = &DescribeScanTaskDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanTaskDetails")
    
    
    return
}

func NewDescribeScanTaskDetailsResponse() (response *DescribeScanTaskDetailsResponse) {
    response = &DescribeScanTaskDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskDetails
// DescribeScanTaskDetails 查询扫描任务详情 , 可以查询扫描进度信息/异常;
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScanTaskDetails(request *DescribeScanTaskDetailsRequest) (response *DescribeScanTaskDetailsResponse, err error) {
    return c.DescribeScanTaskDetailsWithContext(context.Background(), request)
}

// DescribeScanTaskDetails
// DescribeScanTaskDetails 查询扫描任务详情 , 可以查询扫描进度信息/异常;
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScanTaskDetailsWithContext(ctx context.Context, request *DescribeScanTaskDetailsRequest) (response *DescribeScanTaskDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskStatusRequest() (request *DescribeScanTaskStatusRequest) {
    request = &DescribeScanTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanTaskStatus")
    
    
    return
}

func NewDescribeScanTaskStatusResponse() (response *DescribeScanTaskStatusResponse) {
    response = &DescribeScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskStatus
// DescribeScanTaskStatus 查询机器扫描状态列表用于过滤筛选
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanTaskStatus(request *DescribeScanTaskStatusRequest) (response *DescribeScanTaskStatusResponse, err error) {
    return c.DescribeScanTaskStatusWithContext(context.Background(), request)
}

// DescribeScanTaskStatus
// DescribeScanTaskStatus 查询机器扫描状态列表用于过滤筛选
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanTaskStatusWithContext(ctx context.Context, request *DescribeScanTaskStatusRequest) (response *DescribeScanTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanVulSettingRequest() (request *DescribeScanVulSettingRequest) {
    request = &DescribeScanVulSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScanVulSetting")
    
    
    return
}

func NewDescribeScanVulSettingResponse() (response *DescribeScanVulSettingResponse) {
    response = &DescribeScanVulSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanVulSetting
// 查询定期检测的配置
//
// 可能返回的错误码:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScanVulSetting(request *DescribeScanVulSettingRequest) (response *DescribeScanVulSettingResponse, err error) {
    return c.DescribeScanVulSettingWithContext(context.Background(), request)
}

// DescribeScanVulSetting
// 查询定期检测的配置
//
// 可能返回的错误码:
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScanVulSettingWithContext(ctx context.Context, request *DescribeScanVulSettingRequest) (response *DescribeScanVulSettingResponse, err error) {
    if request == nil {
        request = NewDescribeScanVulSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanVulSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanVulSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenAttackHotspotRequest() (request *DescribeScreenAttackHotspotRequest) {
    request = &DescribeScreenAttackHotspotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenAttackHotspot")
    
    
    return
}

func NewDescribeScreenAttackHotspotResponse() (response *DescribeScreenAttackHotspotResponse) {
    response = &DescribeScreenAttackHotspotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenAttackHotspot
// 大屏可视化获取全网攻击热点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenAttackHotspot(request *DescribeScreenAttackHotspotRequest) (response *DescribeScreenAttackHotspotResponse, err error) {
    return c.DescribeScreenAttackHotspotWithContext(context.Background(), request)
}

// DescribeScreenAttackHotspot
// 大屏可视化获取全网攻击热点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenAttackHotspotWithContext(ctx context.Context, request *DescribeScreenAttackHotspotRequest) (response *DescribeScreenAttackHotspotResponse, err error) {
    if request == nil {
        request = NewDescribeScreenAttackHotspotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenAttackHotspot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenAttackHotspotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenBroadcastsRequest() (request *DescribeScreenBroadcastsRequest) {
    request = &DescribeScreenBroadcastsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenBroadcasts")
    
    
    return
}

func NewDescribeScreenBroadcastsResponse() (response *DescribeScreenBroadcastsResponse) {
    response = &DescribeScreenBroadcastsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenBroadcasts
// 大屏可视化安全播报
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenBroadcasts(request *DescribeScreenBroadcastsRequest) (response *DescribeScreenBroadcastsResponse, err error) {
    return c.DescribeScreenBroadcastsWithContext(context.Background(), request)
}

// DescribeScreenBroadcasts
// 大屏可视化安全播报
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenBroadcastsWithContext(ctx context.Context, request *DescribeScreenBroadcastsRequest) (response *DescribeScreenBroadcastsResponse, err error) {
    if request == nil {
        request = NewDescribeScreenBroadcastsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenBroadcasts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenBroadcastsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenDefenseTrendsRequest() (request *DescribeScreenDefenseTrendsRequest) {
    request = &DescribeScreenDefenseTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenDefenseTrends")
    
    
    return
}

func NewDescribeScreenDefenseTrendsResponse() (response *DescribeScreenDefenseTrendsResponse) {
    response = &DescribeScreenDefenseTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenDefenseTrends
// 大屏可视化防趋势接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenDefenseTrends(request *DescribeScreenDefenseTrendsRequest) (response *DescribeScreenDefenseTrendsResponse, err error) {
    return c.DescribeScreenDefenseTrendsWithContext(context.Background(), request)
}

// DescribeScreenDefenseTrends
// 大屏可视化防趋势接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenDefenseTrendsWithContext(ctx context.Context, request *DescribeScreenDefenseTrendsRequest) (response *DescribeScreenDefenseTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeScreenDefenseTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenDefenseTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenDefenseTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenEmergentMsgRequest() (request *DescribeScreenEmergentMsgRequest) {
    request = &DescribeScreenEmergentMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenEmergentMsg")
    
    
    return
}

func NewDescribeScreenEmergentMsgResponse() (response *DescribeScreenEmergentMsgResponse) {
    response = &DescribeScreenEmergentMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenEmergentMsg
// 大屏可视化紧急通知
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenEmergentMsg(request *DescribeScreenEmergentMsgRequest) (response *DescribeScreenEmergentMsgResponse, err error) {
    return c.DescribeScreenEmergentMsgWithContext(context.Background(), request)
}

// DescribeScreenEmergentMsg
// 大屏可视化紧急通知
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenEmergentMsgWithContext(ctx context.Context, request *DescribeScreenEmergentMsgRequest) (response *DescribeScreenEmergentMsgResponse, err error) {
    if request == nil {
        request = NewDescribeScreenEmergentMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenEmergentMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenEmergentMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenEventsCntRequest() (request *DescribeScreenEventsCntRequest) {
    request = &DescribeScreenEventsCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenEventsCnt")
    
    
    return
}

func NewDescribeScreenEventsCntResponse() (response *DescribeScreenEventsCntResponse) {
    response = &DescribeScreenEventsCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenEventsCnt
// 大屏可视化获取安全概览相关事件统计数据接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenEventsCnt(request *DescribeScreenEventsCntRequest) (response *DescribeScreenEventsCntResponse, err error) {
    return c.DescribeScreenEventsCntWithContext(context.Background(), request)
}

// DescribeScreenEventsCnt
// 大屏可视化获取安全概览相关事件统计数据接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenEventsCntWithContext(ctx context.Context, request *DescribeScreenEventsCntRequest) (response *DescribeScreenEventsCntResponse, err error) {
    if request == nil {
        request = NewDescribeScreenEventsCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenEventsCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenEventsCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenGeneralStatRequest() (request *DescribeScreenGeneralStatRequest) {
    request = &DescribeScreenGeneralStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenGeneralStat")
    
    
    return
}

func NewDescribeScreenGeneralStatResponse() (response *DescribeScreenGeneralStatResponse) {
    response = &DescribeScreenGeneralStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenGeneralStat
// 大屏可视化获取主机相关统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenGeneralStat(request *DescribeScreenGeneralStatRequest) (response *DescribeScreenGeneralStatResponse, err error) {
    return c.DescribeScreenGeneralStatWithContext(context.Background(), request)
}

// DescribeScreenGeneralStat
// 大屏可视化获取主机相关统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenGeneralStatWithContext(ctx context.Context, request *DescribeScreenGeneralStatRequest) (response *DescribeScreenGeneralStatResponse, err error) {
    if request == nil {
        request = NewDescribeScreenGeneralStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenGeneralStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenGeneralStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenHostInvasionRequest() (request *DescribeScreenHostInvasionRequest) {
    request = &DescribeScreenHostInvasionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenHostInvasion")
    
    
    return
}

func NewDescribeScreenHostInvasionResponse() (response *DescribeScreenHostInvasionResponse) {
    response = &DescribeScreenHostInvasionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenHostInvasion
// 大屏可视化主机入侵详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenHostInvasion(request *DescribeScreenHostInvasionRequest) (response *DescribeScreenHostInvasionResponse, err error) {
    return c.DescribeScreenHostInvasionWithContext(context.Background(), request)
}

// DescribeScreenHostInvasion
// 大屏可视化主机入侵详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenHostInvasionWithContext(ctx context.Context, request *DescribeScreenHostInvasionRequest) (response *DescribeScreenHostInvasionResponse, err error) {
    if request == nil {
        request = NewDescribeScreenHostInvasionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenHostInvasion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenHostInvasionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenMachineRegionsRequest() (request *DescribeScreenMachineRegionsRequest) {
    request = &DescribeScreenMachineRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenMachineRegions")
    
    
    return
}

func NewDescribeScreenMachineRegionsResponse() (response *DescribeScreenMachineRegionsResponse) {
    response = &DescribeScreenMachineRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenMachineRegions
// 大屏可视化主机区域选项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenMachineRegions(request *DescribeScreenMachineRegionsRequest) (response *DescribeScreenMachineRegionsResponse, err error) {
    return c.DescribeScreenMachineRegionsWithContext(context.Background(), request)
}

// DescribeScreenMachineRegions
// 大屏可视化主机区域选项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenMachineRegionsWithContext(ctx context.Context, request *DescribeScreenMachineRegionsRequest) (response *DescribeScreenMachineRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeScreenMachineRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenMachineRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenMachineRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenMachinesRequest() (request *DescribeScreenMachinesRequest) {
    request = &DescribeScreenMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenMachines")
    
    
    return
}

func NewDescribeScreenMachinesResponse() (response *DescribeScreenMachinesResponse) {
    response = &DescribeScreenMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenMachines
// 大屏可视化主机区域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenMachines(request *DescribeScreenMachinesRequest) (response *DescribeScreenMachinesResponse, err error) {
    return c.DescribeScreenMachinesWithContext(context.Background(), request)
}

// DescribeScreenMachines
// 大屏可视化主机区域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenMachinesWithContext(ctx context.Context, request *DescribeScreenMachinesRequest) (response *DescribeScreenMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeScreenMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenProtectionCntRequest() (request *DescribeScreenProtectionCntRequest) {
    request = &DescribeScreenProtectionCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenProtectionCnt")
    
    
    return
}

func NewDescribeScreenProtectionCntResponse() (response *DescribeScreenProtectionCntResponse) {
    response = &DescribeScreenProtectionCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenProtectionCnt
// 大屏可视化主机安全防护引擎介绍
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenProtectionCnt(request *DescribeScreenProtectionCntRequest) (response *DescribeScreenProtectionCntResponse, err error) {
    return c.DescribeScreenProtectionCntWithContext(context.Background(), request)
}

// DescribeScreenProtectionCnt
// 大屏可视化主机安全防护引擎介绍
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenProtectionCntWithContext(ctx context.Context, request *DescribeScreenProtectionCntRequest) (response *DescribeScreenProtectionCntResponse, err error) {
    if request == nil {
        request = NewDescribeScreenProtectionCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenProtectionCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenProtectionCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenProtectionStatRequest() (request *DescribeScreenProtectionStatRequest) {
    request = &DescribeScreenProtectionStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenProtectionStat")
    
    
    return
}

func NewDescribeScreenProtectionStatResponse() (response *DescribeScreenProtectionStatResponse) {
    response = &DescribeScreenProtectionStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenProtectionStat
// 大屏获取安全防护状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenProtectionStat(request *DescribeScreenProtectionStatRequest) (response *DescribeScreenProtectionStatResponse, err error) {
    return c.DescribeScreenProtectionStatWithContext(context.Background(), request)
}

// DescribeScreenProtectionStat
// 大屏获取安全防护状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScreenProtectionStatWithContext(ctx context.Context, request *DescribeScreenProtectionStatRequest) (response *DescribeScreenProtectionStatResponse, err error) {
    if request == nil {
        request = NewDescribeScreenProtectionStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenProtectionStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenProtectionStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenRiskAssetsTopRequest() (request *DescribeScreenRiskAssetsTopRequest) {
    request = &DescribeScreenRiskAssetsTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeScreenRiskAssetsTop")
    
    
    return
}

func NewDescribeScreenRiskAssetsTopResponse() (response *DescribeScreenRiskAssetsTopResponse) {
    response = &DescribeScreenRiskAssetsTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenRiskAssetsTop
// 大屏可视化风险资产top5（今日），统计今日风险资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenRiskAssetsTop(request *DescribeScreenRiskAssetsTopRequest) (response *DescribeScreenRiskAssetsTopResponse, err error) {
    return c.DescribeScreenRiskAssetsTopWithContext(context.Background(), request)
}

// DescribeScreenRiskAssetsTop
// 大屏可视化风险资产top5（今日），统计今日风险资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeScreenRiskAssetsTopWithContext(ctx context.Context, request *DescribeScreenRiskAssetsTopRequest) (response *DescribeScreenRiskAssetsTopResponse, err error) {
    if request == nil {
        request = NewDescribeScreenRiskAssetsTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenRiskAssetsTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenRiskAssetsTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchLogsRequest() (request *DescribeSearchLogsRequest) {
    request = &DescribeSearchLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchLogs")
    
    
    return
}

func NewDescribeSearchLogsResponse() (response *DescribeSearchLogsResponse) {
    response = &DescribeSearchLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSearchLogs
// 获取历史搜索记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSearchLogs(request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
    return c.DescribeSearchLogsWithContext(context.Background(), request)
}

// DescribeSearchLogs
// 获取历史搜索记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSearchLogsWithContext(ctx context.Context, request *DescribeSearchLogsRequest) (response *DescribeSearchLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSearchLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchTemplatesRequest() (request *DescribeSearchTemplatesRequest) {
    request = &DescribeSearchTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchTemplates")
    
    
    return
}

func NewDescribeSearchTemplatesResponse() (response *DescribeSearchTemplatesResponse) {
    response = &DescribeSearchTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSearchTemplates
// 获取快速检索列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
    return c.DescribeSearchTemplatesWithContext(context.Background(), request)
}

// DescribeSearchTemplates
// 获取快速检索列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSearchTemplatesWithContext(ctx context.Context, request *DescribeSearchTemplatesRequest) (response *DescribeSearchTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeSearchTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityBroadcastInfoRequest() (request *DescribeSecurityBroadcastInfoRequest) {
    request = &DescribeSecurityBroadcastInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityBroadcastInfo")
    
    
    return
}

func NewDescribeSecurityBroadcastInfoResponse() (response *DescribeSecurityBroadcastInfoResponse) {
    response = &DescribeSecurityBroadcastInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityBroadcastInfo
// 查询安全播报文章信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityBroadcastInfo(request *DescribeSecurityBroadcastInfoRequest) (response *DescribeSecurityBroadcastInfoResponse, err error) {
    return c.DescribeSecurityBroadcastInfoWithContext(context.Background(), request)
}

// DescribeSecurityBroadcastInfo
// 查询安全播报文章信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityBroadcastInfoWithContext(ctx context.Context, request *DescribeSecurityBroadcastInfoRequest) (response *DescribeSecurityBroadcastInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityBroadcastInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityBroadcastInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityBroadcastInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityBroadcastsRequest() (request *DescribeSecurityBroadcastsRequest) {
    request = &DescribeSecurityBroadcastsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityBroadcasts")
    
    
    return
}

func NewDescribeSecurityBroadcastsResponse() (response *DescribeSecurityBroadcastsResponse) {
    response = &DescribeSecurityBroadcastsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityBroadcasts
// 安全播报列表页
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeSecurityBroadcasts(request *DescribeSecurityBroadcastsRequest) (response *DescribeSecurityBroadcastsResponse, err error) {
    return c.DescribeSecurityBroadcastsWithContext(context.Background(), request)
}

// DescribeSecurityBroadcasts
// 安全播报列表页
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeSecurityBroadcastsWithContext(ctx context.Context, request *DescribeSecurityBroadcastsRequest) (response *DescribeSecurityBroadcastsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityBroadcastsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityBroadcasts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityBroadcastsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
    request = &DescribeSecurityDynamicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityDynamics")
    
    
    return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
    response = &DescribeSecurityDynamicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityDynamics
// 本接口 (DescribeSecurityDynamics) 用于获取安全事件动态消息数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    return c.DescribeSecurityDynamicsWithContext(context.Background(), request)
}

// DescribeSecurityDynamics
// 本接口 (DescribeSecurityDynamics) 用于获取安全事件动态消息数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityDynamicsWithContext(ctx context.Context, request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityDynamicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityDynamics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityDynamicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityEventStatRequest() (request *DescribeSecurityEventStatRequest) {
    request = &DescribeSecurityEventStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityEventStat")
    
    
    return
}

func NewDescribeSecurityEventStatResponse() (response *DescribeSecurityEventStatResponse) {
    response = &DescribeSecurityEventStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityEventStat
// 获取安全事件统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityEventStat(request *DescribeSecurityEventStatRequest) (response *DescribeSecurityEventStatResponse, err error) {
    return c.DescribeSecurityEventStatWithContext(context.Background(), request)
}

// DescribeSecurityEventStat
// 获取安全事件统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityEventStatWithContext(ctx context.Context, request *DescribeSecurityEventStatRequest) (response *DescribeSecurityEventStatResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityEventStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityEventStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityEventStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityEventsCntRequest() (request *DescribeSecurityEventsCntRequest) {
    request = &DescribeSecurityEventsCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityEventsCnt")
    
    
    return
}

func NewDescribeSecurityEventsCntResponse() (response *DescribeSecurityEventsCntResponse) {
    response = &DescribeSecurityEventsCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityEventsCnt
// 获取安全概览相关事件统计数据接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityEventsCnt(request *DescribeSecurityEventsCntRequest) (response *DescribeSecurityEventsCntResponse, err error) {
    return c.DescribeSecurityEventsCntWithContext(context.Background(), request)
}

// DescribeSecurityEventsCnt
// 获取安全概览相关事件统计数据接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityEventsCntWithContext(ctx context.Context, request *DescribeSecurityEventsCntRequest) (response *DescribeSecurityEventsCntResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityEventsCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityEventsCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityEventsCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityProtectionStatRequest() (request *DescribeSecurityProtectionStatRequest) {
    request = &DescribeSecurityProtectionStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityProtectionStat")
    
    
    return
}

func NewDescribeSecurityProtectionStatResponse() (response *DescribeSecurityProtectionStatResponse) {
    response = &DescribeSecurityProtectionStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityProtectionStat
// 获取安全防护状态汇总
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSecurityProtectionStat(request *DescribeSecurityProtectionStatRequest) (response *DescribeSecurityProtectionStatResponse, err error) {
    return c.DescribeSecurityProtectionStatWithContext(context.Background(), request)
}

// DescribeSecurityProtectionStat
// 获取安全防护状态汇总
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSecurityProtectionStatWithContext(ctx context.Context, request *DescribeSecurityProtectionStatRequest) (response *DescribeSecurityProtectionStatResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityProtectionStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityProtectionStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityProtectionStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
    request = &DescribeSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSecurityTrends")
    
    
    return
}

func NewDescribeSecurityTrendsResponse() (response *DescribeSecurityTrendsResponse) {
    response = &DescribeSecurityTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityTrends
// 本接口 (DescribeSecurityTrends) 用于获取安全事件统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    return c.DescribeSecurityTrendsWithContext(context.Background(), request)
}

// DescribeSecurityTrends
// 本接口 (DescribeSecurityTrends) 用于获取安全事件统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSecurityTrendsWithContext(ctx context.Context, request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerRelatedDirInfoRequest() (request *DescribeServerRelatedDirInfoRequest) {
    request = &DescribeServerRelatedDirInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeServerRelatedDirInfo")
    
    
    return
}

func NewDescribeServerRelatedDirInfoResponse() (response *DescribeServerRelatedDirInfoResponse) {
    response = &DescribeServerRelatedDirInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServerRelatedDirInfo
// 查询服务区关联目录详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeServerRelatedDirInfo(request *DescribeServerRelatedDirInfoRequest) (response *DescribeServerRelatedDirInfoResponse, err error) {
    return c.DescribeServerRelatedDirInfoWithContext(context.Background(), request)
}

// DescribeServerRelatedDirInfo
// 查询服务区关联目录详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeServerRelatedDirInfoWithContext(ctx context.Context, request *DescribeServerRelatedDirInfoRequest) (response *DescribeServerRelatedDirInfoResponse, err error) {
    if request == nil {
        request = NewDescribeServerRelatedDirInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerRelatedDirInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerRelatedDirInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServersAndRiskAndFirstInfoRequest() (request *DescribeServersAndRiskAndFirstInfoRequest) {
    request = &DescribeServersAndRiskAndFirstInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeServersAndRiskAndFirstInfo")
    
    
    return
}

func NewDescribeServersAndRiskAndFirstInfoResponse() (response *DescribeServersAndRiskAndFirstInfoResponse) {
    response = &DescribeServersAndRiskAndFirstInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServersAndRiskAndFirstInfo
// 获取待处理风险文件数+影响服务器数+是否试用检测+最近检测时间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeServersAndRiskAndFirstInfo(request *DescribeServersAndRiskAndFirstInfoRequest) (response *DescribeServersAndRiskAndFirstInfoResponse, err error) {
    return c.DescribeServersAndRiskAndFirstInfoWithContext(context.Background(), request)
}

// DescribeServersAndRiskAndFirstInfo
// 获取待处理风险文件数+影响服务器数+是否试用检测+最近检测时间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeServersAndRiskAndFirstInfoWithContext(ctx context.Context, request *DescribeServersAndRiskAndFirstInfoRequest) (response *DescribeServersAndRiskAndFirstInfoResponse, err error) {
    if request == nil {
        request = NewDescribeServersAndRiskAndFirstInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServersAndRiskAndFirstInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServersAndRiskAndFirstInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyExistRequest() (request *DescribeStrategyExistRequest) {
    request = &DescribeStrategyExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeStrategyExist")
    
    
    return
}

func NewDescribeStrategyExistResponse() (response *DescribeStrategyExistResponse) {
    response = &DescribeStrategyExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStrategyExist
// 根据策略名查询策略是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStrategyExist(request *DescribeStrategyExistRequest) (response *DescribeStrategyExistResponse, err error) {
    return c.DescribeStrategyExistWithContext(context.Background(), request)
}

// DescribeStrategyExist
// 根据策略名查询策略是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStrategyExistWithContext(ctx context.Context, request *DescribeStrategyExistRequest) (response *DescribeStrategyExistResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategyExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyExistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
    request = &DescribeTagMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeTagMachines")
    
    
    return
}

func NewDescribeTagMachinesResponse() (response *DescribeTagMachinesResponse) {
    response = &DescribeTagMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagMachines
// 获取指定标签关联的服务器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagMachines(request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    return c.DescribeTagMachinesWithContext(context.Background(), request)
}

// DescribeTagMachines
// 获取指定标签关联的服务器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagMachinesWithContext(ctx context.Context, request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeTagMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeTags")
    
    
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTags
// 获取所有主机标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    return c.DescribeTagsWithContext(context.Background(), request)
}

// DescribeTags
// 获取所有主机标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTagsWithContext(ctx context.Context, request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrialReportRequest() (request *DescribeTrialReportRequest) {
    request = &DescribeTrialReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeTrialReport")
    
    
    return
}

func NewDescribeTrialReportResponse() (response *DescribeTrialReportResponse) {
    response = &DescribeTrialReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrialReport
// 查询主机安全授权试用报告(仅限控制台申领的)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrialReport(request *DescribeTrialReportRequest) (response *DescribeTrialReportResponse, err error) {
    return c.DescribeTrialReportWithContext(context.Background(), request)
}

// DescribeTrialReport
// 查询主机安全授权试用报告(仅限控制台申领的)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrialReportWithContext(ctx context.Context, request *DescribeTrialReportRequest) (response *DescribeTrialReportResponse, err error) {
    if request == nil {
        request = NewDescribeTrialReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrialReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrialReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUndoVulCountsRequest() (request *DescribeUndoVulCountsRequest) {
    request = &DescribeUndoVulCountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeUndoVulCounts")
    
    
    return
}

func NewDescribeUndoVulCountsResponse() (response *DescribeUndoVulCountsResponse) {
    response = &DescribeUndoVulCountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUndoVulCounts
// 获取漏洞管理模块指定类型的待处理漏洞数、主机数和非专业版主机数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUndoVulCounts(request *DescribeUndoVulCountsRequest) (response *DescribeUndoVulCountsResponse, err error) {
    return c.DescribeUndoVulCountsWithContext(context.Background(), request)
}

// DescribeUndoVulCounts
// 获取漏洞管理模块指定类型的待处理漏洞数、主机数和非专业版主机数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUndoVulCountsWithContext(ctx context.Context, request *DescribeUndoVulCountsRequest) (response *DescribeUndoVulCountsResponse, err error) {
    if request == nil {
        request = NewDescribeUndoVulCountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUndoVulCounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUndoVulCountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersConfigRequest() (request *DescribeUsersConfigRequest) {
    request = &DescribeUsersConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeUsersConfig")
    
    
    return
}

func NewDescribeUsersConfigResponse() (response *DescribeUsersConfigResponse) {
    response = &DescribeUsersConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsersConfig
// 用于查询用户自定义配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsersConfig(request *DescribeUsersConfigRequest) (response *DescribeUsersConfigResponse, err error) {
    return c.DescribeUsersConfigWithContext(context.Background(), request)
}

// DescribeUsersConfig
// 用于查询用户自定义配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsersConfigWithContext(ctx context.Context, request *DescribeUsersConfigRequest) (response *DescribeUsersConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUsersConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsersConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
    request = &DescribeUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeUsualLoginPlaces")
    
    
    return
}

func NewDescribeUsualLoginPlacesResponse() (response *DescribeUsualLoginPlacesResponse) {
    response = &DescribeUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsualLoginPlaces
// 此接口（DescribeUsualLoginPlaces）用于查询常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    return c.DescribeUsualLoginPlacesWithContext(context.Background(), request)
}

// DescribeUsualLoginPlaces
// 此接口（DescribeUsualLoginPlaces）用于查询常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUsualLoginPlacesWithContext(ctx context.Context, request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVdbAndPocInfoRequest() (request *DescribeVdbAndPocInfoRequest) {
    request = &DescribeVdbAndPocInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVdbAndPocInfo")
    
    
    return
}

func NewDescribeVdbAndPocInfoResponse() (response *DescribeVdbAndPocInfoResponse) {
    response = &DescribeVdbAndPocInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVdbAndPocInfo
// 获取病毒库及POC的更新信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVdbAndPocInfo(request *DescribeVdbAndPocInfoRequest) (response *DescribeVdbAndPocInfoResponse, err error) {
    return c.DescribeVdbAndPocInfoWithContext(context.Background(), request)
}

// DescribeVdbAndPocInfo
// 获取病毒库及POC的更新信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVdbAndPocInfoWithContext(ctx context.Context, request *DescribeVdbAndPocInfoRequest) (response *DescribeVdbAndPocInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVdbAndPocInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVdbAndPocInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVdbAndPocInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionCompareChartRequest() (request *DescribeVersionCompareChartRequest) {
    request = &DescribeVersionCompareChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVersionCompareChart")
    
    
    return
}

func NewDescribeVersionCompareChartResponse() (response *DescribeVersionCompareChartResponse) {
    response = &DescribeVersionCompareChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVersionCompareChart
// 获取版本对比信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVersionCompareChart(request *DescribeVersionCompareChartRequest) (response *DescribeVersionCompareChartResponse, err error) {
    return c.DescribeVersionCompareChartWithContext(context.Background(), request)
}

// DescribeVersionCompareChart
// 获取版本对比信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVersionCompareChartWithContext(ctx context.Context, request *DescribeVersionCompareChartRequest) (response *DescribeVersionCompareChartResponse, err error) {
    if request == nil {
        request = NewDescribeVersionCompareChartRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVersionCompareChart require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVersionCompareChartResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionStatisticsRequest() (request *DescribeVersionStatisticsRequest) {
    request = &DescribeVersionStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVersionStatistics")
    
    
    return
}

func NewDescribeVersionStatisticsResponse() (response *DescribeVersionStatisticsResponse) {
    response = &DescribeVersionStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVersionStatistics
// 用于统计专业版和基础版机器数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVersionStatistics(request *DescribeVersionStatisticsRequest) (response *DescribeVersionStatisticsResponse, err error) {
    return c.DescribeVersionStatisticsWithContext(context.Background(), request)
}

// DescribeVersionStatistics
// 用于统计专业版和基础版机器数。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVersionStatisticsWithContext(ctx context.Context, request *DescribeVersionStatisticsRequest) (response *DescribeVersionStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeVersionStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVersionStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVersionStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVertexDetailRequest() (request *DescribeVertexDetailRequest) {
    request = &DescribeVertexDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVertexDetail")
    
    
    return
}

func NewDescribeVertexDetailResponse() (response *DescribeVertexDetailResponse) {
    response = &DescribeVertexDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVertexDetail
// 获取指定点属性信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVertexDetail(request *DescribeVertexDetailRequest) (response *DescribeVertexDetailResponse, err error) {
    return c.DescribeVertexDetailWithContext(context.Background(), request)
}

// DescribeVertexDetail
// 获取指定点属性信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVertexDetailWithContext(ctx context.Context, request *DescribeVertexDetailRequest) (response *DescribeVertexDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVertexDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVertexDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVertexDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulCountByDatesRequest() (request *DescribeVulCountByDatesRequest) {
    request = &DescribeVulCountByDatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulCountByDates")
    
    
    return
}

func NewDescribeVulCountByDatesResponse() (response *DescribeVulCountByDatesResponse) {
    response = &DescribeVulCountByDatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulCountByDates
// 漏洞管理模块，获取近日指定类型的漏洞数量和主机数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulCountByDates(request *DescribeVulCountByDatesRequest) (response *DescribeVulCountByDatesResponse, err error) {
    return c.DescribeVulCountByDatesWithContext(context.Background(), request)
}

// DescribeVulCountByDates
// 漏洞管理模块，获取近日指定类型的漏洞数量和主机数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulCountByDatesWithContext(ctx context.Context, request *DescribeVulCountByDatesRequest) (response *DescribeVulCountByDatesResponse, err error) {
    if request == nil {
        request = NewDescribeVulCountByDatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulCountByDates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulCountByDatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulCveIdInfoRequest() (request *DescribeVulCveIdInfoRequest) {
    request = &DescribeVulCveIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulCveIdInfo")
    
    
    return
}

func NewDescribeVulCveIdInfoResponse() (response *DescribeVulCveIdInfoResponse) {
    response = &DescribeVulCveIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulCveIdInfo
// CveId查询漏洞详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulCveIdInfo(request *DescribeVulCveIdInfoRequest) (response *DescribeVulCveIdInfoResponse, err error) {
    return c.DescribeVulCveIdInfoWithContext(context.Background(), request)
}

// DescribeVulCveIdInfo
// CveId查询漏洞详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulCveIdInfoWithContext(ctx context.Context, request *DescribeVulCveIdInfoRequest) (response *DescribeVulCveIdInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulCveIdInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulCveIdInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulCveIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceEventRequest() (request *DescribeVulDefenceEventRequest) {
    request = &DescribeVulDefenceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceEvent")
    
    
    return
}

func NewDescribeVulDefenceEventResponse() (response *DescribeVulDefenceEventResponse) {
    response = &DescribeVulDefenceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefenceEvent
// 获取漏洞防御事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceEvent(request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
    return c.DescribeVulDefenceEventWithContext(context.Background(), request)
}

// DescribeVulDefenceEvent
// 获取漏洞防御事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceEventWithContext(ctx context.Context, request *DescribeVulDefenceEventRequest) (response *DescribeVulDefenceEventResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceListRequest() (request *DescribeVulDefenceListRequest) {
    request = &DescribeVulDefenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceList")
    
    
    return
}

func NewDescribeVulDefenceListResponse() (response *DescribeVulDefenceListResponse) {
    response = &DescribeVulDefenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefenceList
// 查询漏洞防御列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceList(request *DescribeVulDefenceListRequest) (response *DescribeVulDefenceListResponse, err error) {
    return c.DescribeVulDefenceListWithContext(context.Background(), request)
}

// DescribeVulDefenceList
// 查询漏洞防御列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceListWithContext(ctx context.Context, request *DescribeVulDefenceListRequest) (response *DescribeVulDefenceListResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceOverviewRequest() (request *DescribeVulDefenceOverviewRequest) {
    request = &DescribeVulDefenceOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceOverview")
    
    
    return
}

func NewDescribeVulDefenceOverviewResponse() (response *DescribeVulDefenceOverviewResponse) {
    response = &DescribeVulDefenceOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefenceOverview
// 获取漏洞防御概览信息，包括事件趋势及插件开启情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceOverview(request *DescribeVulDefenceOverviewRequest) (response *DescribeVulDefenceOverviewResponse, err error) {
    return c.DescribeVulDefenceOverviewWithContext(context.Background(), request)
}

// DescribeVulDefenceOverview
// 获取漏洞防御概览信息，包括事件趋势及插件开启情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulDefenceOverviewWithContext(ctx context.Context, request *DescribeVulDefenceOverviewRequest) (response *DescribeVulDefenceOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefencePluginDetailRequest() (request *DescribeVulDefencePluginDetailRequest) {
    request = &DescribeVulDefencePluginDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginDetail")
    
    
    return
}

func NewDescribeVulDefencePluginDetailResponse() (response *DescribeVulDefencePluginDetailResponse) {
    response = &DescribeVulDefencePluginDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefencePluginDetail
// 获取单台主机漏洞防御插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginDetail(request *DescribeVulDefencePluginDetailRequest) (response *DescribeVulDefencePluginDetailResponse, err error) {
    return c.DescribeVulDefencePluginDetailWithContext(context.Background(), request)
}

// DescribeVulDefencePluginDetail
// 获取单台主机漏洞防御插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginDetailWithContext(ctx context.Context, request *DescribeVulDefencePluginDetailRequest) (response *DescribeVulDefencePluginDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefencePluginDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefencePluginDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefencePluginDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefencePluginExceptionCountRequest() (request *DescribeVulDefencePluginExceptionCountRequest) {
    request = &DescribeVulDefencePluginExceptionCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginExceptionCount")
    
    
    return
}

func NewDescribeVulDefencePluginExceptionCountResponse() (response *DescribeVulDefencePluginExceptionCountResponse) {
    response = &DescribeVulDefencePluginExceptionCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefencePluginExceptionCount
// 获取当前异常插件数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginExceptionCount(request *DescribeVulDefencePluginExceptionCountRequest) (response *DescribeVulDefencePluginExceptionCountResponse, err error) {
    return c.DescribeVulDefencePluginExceptionCountWithContext(context.Background(), request)
}

// DescribeVulDefencePluginExceptionCount
// 获取当前异常插件数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginExceptionCountWithContext(ctx context.Context, request *DescribeVulDefencePluginExceptionCountRequest) (response *DescribeVulDefencePluginExceptionCountResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefencePluginExceptionCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefencePluginExceptionCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefencePluginExceptionCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefencePluginStatusRequest() (request *DescribeVulDefencePluginStatusRequest) {
    request = &DescribeVulDefencePluginStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefencePluginStatus")
    
    
    return
}

func NewDescribeVulDefencePluginStatusResponse() (response *DescribeVulDefencePluginStatusResponse) {
    response = &DescribeVulDefencePluginStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefencePluginStatus
// 获取各主机漏洞防御插件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginStatus(request *DescribeVulDefencePluginStatusRequest) (response *DescribeVulDefencePluginStatusResponse, err error) {
    return c.DescribeVulDefencePluginStatusWithContext(context.Background(), request)
}

// DescribeVulDefencePluginStatus
// 获取各主机漏洞防御插件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefencePluginStatusWithContext(ctx context.Context, request *DescribeVulDefencePluginStatusRequest) (response *DescribeVulDefencePluginStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefencePluginStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefencePluginStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefencePluginStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDefenceSettingRequest() (request *DescribeVulDefenceSettingRequest) {
    request = &DescribeVulDefenceSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulDefenceSetting")
    
    
    return
}

func NewDescribeVulDefenceSettingResponse() (response *DescribeVulDefenceSettingResponse) {
    response = &DescribeVulDefenceSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDefenceSetting
// 获取当前漏洞防御插件设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefenceSetting(request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
    return c.DescribeVulDefenceSettingWithContext(context.Background(), request)
}

// DescribeVulDefenceSetting
// 获取当前漏洞防御插件设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulDefenceSettingWithContext(ctx context.Context, request *DescribeVulDefenceSettingRequest) (response *DescribeVulDefenceSettingResponse, err error) {
    if request == nil {
        request = NewDescribeVulDefenceSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDefenceSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDefenceSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulEffectHostListRequest() (request *DescribeVulEffectHostListRequest) {
    request = &DescribeVulEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEffectHostList")
    
    
    return
}

func NewDescribeVulEffectHostListResponse() (response *DescribeVulEffectHostListResponse) {
    response = &DescribeVulEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulEffectHostList
// 漏洞影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulEffectHostList(request *DescribeVulEffectHostListRequest) (response *DescribeVulEffectHostListResponse, err error) {
    return c.DescribeVulEffectHostListWithContext(context.Background(), request)
}

// DescribeVulEffectHostList
// 漏洞影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulEffectHostListWithContext(ctx context.Context, request *DescribeVulEffectHostListRequest) (response *DescribeVulEffectHostListResponse, err error) {
    if request == nil {
        request = NewDescribeVulEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulEffectModulesRequest() (request *DescribeVulEffectModulesRequest) {
    request = &DescribeVulEffectModulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEffectModules")
    
    
    return
}

func NewDescribeVulEffectModulesResponse() (response *DescribeVulEffectModulesResponse) {
    response = &DescribeVulEffectModulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulEffectModules
// 漏洞影响组件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulEffectModules(request *DescribeVulEffectModulesRequest) (response *DescribeVulEffectModulesResponse, err error) {
    return c.DescribeVulEffectModulesWithContext(context.Background(), request)
}

// DescribeVulEffectModules
// 漏洞影响组件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulEffectModulesWithContext(ctx context.Context, request *DescribeVulEffectModulesRequest) (response *DescribeVulEffectModulesResponse, err error) {
    if request == nil {
        request = NewDescribeVulEffectModulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulEffectModules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulEffectModulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulEmergentMsgRequest() (request *DescribeVulEmergentMsgRequest) {
    request = &DescribeVulEmergentMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulEmergentMsg")
    
    
    return
}

func NewDescribeVulEmergentMsgResponse() (response *DescribeVulEmergentMsgResponse) {
    response = &DescribeVulEmergentMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulEmergentMsg
// 获取漏洞紧急通知
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulEmergentMsg(request *DescribeVulEmergentMsgRequest) (response *DescribeVulEmergentMsgResponse, err error) {
    return c.DescribeVulEmergentMsgWithContext(context.Background(), request)
}

// DescribeVulEmergentMsg
// 获取漏洞紧急通知
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulEmergentMsgWithContext(ctx context.Context, request *DescribeVulEmergentMsgRequest) (response *DescribeVulEmergentMsgResponse, err error) {
    if request == nil {
        request = NewDescribeVulEmergentMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulEmergentMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulEmergentMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulFixStatusRequest() (request *DescribeVulFixStatusRequest) {
    request = &DescribeVulFixStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulFixStatus")
    
    
    return
}

func NewDescribeVulFixStatusResponse() (response *DescribeVulFixStatusResponse) {
    response = &DescribeVulFixStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulFixStatus
// 漏洞修护-查找主机漏洞修护进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulFixStatus(request *DescribeVulFixStatusRequest) (response *DescribeVulFixStatusResponse, err error) {
    return c.DescribeVulFixStatusWithContext(context.Background(), request)
}

// DescribeVulFixStatus
// 漏洞修护-查找主机漏洞修护进度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulFixStatusWithContext(ctx context.Context, request *DescribeVulFixStatusRequest) (response *DescribeVulFixStatusResponse, err error) {
    if request == nil {
        request = NewDescribeVulFixStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulFixStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulFixStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulHostCountScanTimeRequest() (request *DescribeVulHostCountScanTimeRequest) {
    request = &DescribeVulHostCountScanTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulHostCountScanTime")
    
    
    return
}

func NewDescribeVulHostCountScanTimeResponse() (response *DescribeVulHostCountScanTimeResponse) {
    response = &DescribeVulHostCountScanTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulHostCountScanTime
// 获取待处理漏洞数+影响主机数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulHostCountScanTime(request *DescribeVulHostCountScanTimeRequest) (response *DescribeVulHostCountScanTimeResponse, err error) {
    return c.DescribeVulHostCountScanTimeWithContext(context.Background(), request)
}

// DescribeVulHostCountScanTime
// 获取待处理漏洞数+影响主机数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulHostCountScanTimeWithContext(ctx context.Context, request *DescribeVulHostCountScanTimeRequest) (response *DescribeVulHostCountScanTimeResponse, err error) {
    if request == nil {
        request = NewDescribeVulHostCountScanTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulHostCountScanTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulHostCountScanTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulHostTopRequest() (request *DescribeVulHostTopRequest) {
    request = &DescribeVulHostTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulHostTop")
    
    
    return
}

func NewDescribeVulHostTopResponse() (response *DescribeVulHostTopResponse) {
    response = &DescribeVulHostTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulHostTop
// 获取服务器风险top列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulHostTop(request *DescribeVulHostTopRequest) (response *DescribeVulHostTopResponse, err error) {
    return c.DescribeVulHostTopWithContext(context.Background(), request)
}

// DescribeVulHostTop
// 获取服务器风险top列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulHostTopWithContext(ctx context.Context, request *DescribeVulHostTopRequest) (response *DescribeVulHostTopResponse, err error) {
    if request == nil {
        request = NewDescribeVulHostTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulHostTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulHostTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulInfoCvssRequest() (request *DescribeVulInfoCvssRequest) {
    request = &DescribeVulInfoCvssRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulInfoCvss")
    
    
    return
}

func NewDescribeVulInfoCvssResponse() (response *DescribeVulInfoCvssResponse) {
    response = &DescribeVulInfoCvssResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulInfoCvss
// 漏洞详情，带CVSS版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulInfoCvss(request *DescribeVulInfoCvssRequest) (response *DescribeVulInfoCvssResponse, err error) {
    return c.DescribeVulInfoCvssWithContext(context.Background(), request)
}

// DescribeVulInfoCvss
// 漏洞详情，带CVSS版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulInfoCvssWithContext(ctx context.Context, request *DescribeVulInfoCvssRequest) (response *DescribeVulInfoCvssResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoCvssRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulInfoCvss require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulInfoCvssResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulLabelsRequest() (request *DescribeVulLabelsRequest) {
    request = &DescribeVulLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulLabels")
    
    
    return
}

func NewDescribeVulLabelsResponse() (response *DescribeVulLabelsResponse) {
    response = &DescribeVulLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulLabels
// 获取用户漏洞所有标签列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulLabels(request *DescribeVulLabelsRequest) (response *DescribeVulLabelsResponse, err error) {
    return c.DescribeVulLabelsWithContext(context.Background(), request)
}

// DescribeVulLabels
// 获取用户漏洞所有标签列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulLabelsWithContext(ctx context.Context, request *DescribeVulLabelsRequest) (response *DescribeVulLabelsResponse, err error) {
    if request == nil {
        request = NewDescribeVulLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulLevelCountRequest() (request *DescribeVulLevelCountRequest) {
    request = &DescribeVulLevelCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulLevelCount")
    
    
    return
}

func NewDescribeVulLevelCountResponse() (response *DescribeVulLevelCountResponse) {
    response = &DescribeVulLevelCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulLevelCount
// 漏洞数量等级分布统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulLevelCount(request *DescribeVulLevelCountRequest) (response *DescribeVulLevelCountResponse, err error) {
    return c.DescribeVulLevelCountWithContext(context.Background(), request)
}

// DescribeVulLevelCount
// 漏洞数量等级分布统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeVulLevelCountWithContext(ctx context.Context, request *DescribeVulLevelCountRequest) (response *DescribeVulLevelCountResponse, err error) {
    if request == nil {
        request = NewDescribeVulLevelCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulLevelCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulLevelCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulListRequest() (request *DescribeVulListRequest) {
    request = &DescribeVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulList")
    
    
    return
}

func NewDescribeVulListResponse() (response *DescribeVulListResponse) {
    response = &DescribeVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulList
// 获取漏洞列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    return c.DescribeVulListWithContext(context.Background(), request)
}

// DescribeVulList
// 获取漏洞列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulListWithContext(ctx context.Context, request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    if request == nil {
        request = NewDescribeVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulOverviewRequest() (request *DescribeVulOverviewRequest) {
    request = &DescribeVulOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulOverview")
    
    
    return
}

func NewDescribeVulOverviewResponse() (response *DescribeVulOverviewResponse) {
    response = &DescribeVulOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulOverview
// 获取漏洞概览数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulOverview(request *DescribeVulOverviewRequest) (response *DescribeVulOverviewResponse, err error) {
    return c.DescribeVulOverviewWithContext(context.Background(), request)
}

// DescribeVulOverview
// 获取漏洞概览数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVulOverviewWithContext(ctx context.Context, request *DescribeVulOverviewRequest) (response *DescribeVulOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeVulOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulStoreListRequest() (request *DescribeVulStoreListRequest) {
    request = &DescribeVulStoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulStoreList")
    
    
    return
}

func NewDescribeVulStoreListResponse() (response *DescribeVulStoreListResponse) {
    response = &DescribeVulStoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulStoreList
// 获取漏洞库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVulStoreList(request *DescribeVulStoreListRequest) (response *DescribeVulStoreListResponse, err error) {
    return c.DescribeVulStoreListWithContext(context.Background(), request)
}

// DescribeVulStoreList
// 获取漏洞库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVulStoreListWithContext(ctx context.Context, request *DescribeVulStoreListRequest) (response *DescribeVulStoreListResponse, err error) {
    if request == nil {
        request = NewDescribeVulStoreListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulStoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulStoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulTopRequest() (request *DescribeVulTopRequest) {
    request = &DescribeVulTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulTop")
    
    
    return
}

func NewDescribeVulTopResponse() (response *DescribeVulTopResponse) {
    response = &DescribeVulTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulTop
// 漏洞top统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulTop(request *DescribeVulTopRequest) (response *DescribeVulTopResponse, err error) {
    return c.DescribeVulTopWithContext(context.Background(), request)
}

// DescribeVulTop
// 漏洞top统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulTopWithContext(ctx context.Context, request *DescribeVulTopRequest) (response *DescribeVulTopResponse, err error) {
    if request == nil {
        request = NewDescribeVulTopRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulTopResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulTrendRequest() (request *DescribeVulTrendRequest) {
    request = &DescribeVulTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulTrend")
    
    
    return
}

func NewDescribeVulTrendResponse() (response *DescribeVulTrendResponse) {
    response = &DescribeVulTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulTrend
// 获取漏洞态势信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulTrend(request *DescribeVulTrendRequest) (response *DescribeVulTrendResponse, err error) {
    return c.DescribeVulTrendWithContext(context.Background(), request)
}

// DescribeVulTrend
// 获取漏洞态势信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulTrendWithContext(ctx context.Context, request *DescribeVulTrendRequest) (response *DescribeVulTrendResponse, err error) {
    if request == nil {
        request = NewDescribeVulTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarningHostConfigRequest() (request *DescribeWarningHostConfigRequest) {
    request = &DescribeWarningHostConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWarningHostConfig")
    
    
    return
}

func NewDescribeWarningHostConfigResponse() (response *DescribeWarningHostConfigResponse) {
    response = &DescribeWarningHostConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWarningHostConfig
// 查询告警机器范围配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWarningHostConfig(request *DescribeWarningHostConfigRequest) (response *DescribeWarningHostConfigResponse, err error) {
    return c.DescribeWarningHostConfigWithContext(context.Background(), request)
}

// DescribeWarningHostConfig
// 查询告警机器范围配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWarningHostConfigWithContext(ctx context.Context, request *DescribeWarningHostConfigRequest) (response *DescribeWarningHostConfigResponse, err error) {
    if request == nil {
        request = NewDescribeWarningHostConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWarningHostConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWarningHostConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarningListRequest() (request *DescribeWarningListRequest) {
    request = &DescribeWarningListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWarningList")
    
    
    return
}

func NewDescribeWarningListResponse() (response *DescribeWarningListResponse) {
    response = &DescribeWarningListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWarningList
// 获取当前用户告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWarningList(request *DescribeWarningListRequest) (response *DescribeWarningListResponse, err error) {
    return c.DescribeWarningListWithContext(context.Background(), request)
}

// DescribeWarningList
// 获取当前用户告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWarningListWithContext(ctx context.Context, request *DescribeWarningListRequest) (response *DescribeWarningListResponse, err error) {
    if request == nil {
        request = NewDescribeWarningListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWarningList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWarningListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebHookPolicyRequest() (request *DescribeWebHookPolicyRequest) {
    request = &DescribeWebHookPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookPolicy")
    
    
    return
}

func NewDescribeWebHookPolicyResponse() (response *DescribeWebHookPolicyResponse) {
    response = &DescribeWebHookPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebHookPolicy
// 查询告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookPolicy(request *DescribeWebHookPolicyRequest) (response *DescribeWebHookPolicyResponse, err error) {
    return c.DescribeWebHookPolicyWithContext(context.Background(), request)
}

// DescribeWebHookPolicy
// 查询告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookPolicyWithContext(ctx context.Context, request *DescribeWebHookPolicyRequest) (response *DescribeWebHookPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeWebHookPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebHookPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebHookPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebHookReceiverRequest() (request *DescribeWebHookReceiverRequest) {
    request = &DescribeWebHookReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookReceiver")
    
    
    return
}

func NewDescribeWebHookReceiverResponse() (response *DescribeWebHookReceiverResponse) {
    response = &DescribeWebHookReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebHookReceiver
// 查询告警接收人列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookReceiver(request *DescribeWebHookReceiverRequest) (response *DescribeWebHookReceiverResponse, err error) {
    return c.DescribeWebHookReceiverWithContext(context.Background(), request)
}

// DescribeWebHookReceiver
// 查询告警接收人列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookReceiverWithContext(ctx context.Context, request *DescribeWebHookReceiverRequest) (response *DescribeWebHookReceiverResponse, err error) {
    if request == nil {
        request = NewDescribeWebHookReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebHookReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebHookReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebHookReceiverUsageRequest() (request *DescribeWebHookReceiverUsageRequest) {
    request = &DescribeWebHookReceiverUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookReceiverUsage")
    
    
    return
}

func NewDescribeWebHookReceiverUsageResponse() (response *DescribeWebHookReceiverUsageResponse) {
    response = &DescribeWebHookReceiverUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebHookReceiverUsage
// 查询指定告警接收人的关联策略使用信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookReceiverUsage(request *DescribeWebHookReceiverUsageRequest) (response *DescribeWebHookReceiverUsageResponse, err error) {
    return c.DescribeWebHookReceiverUsageWithContext(context.Background(), request)
}

// DescribeWebHookReceiverUsage
// 查询指定告警接收人的关联策略使用信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebHookReceiverUsageWithContext(ctx context.Context, request *DescribeWebHookReceiverUsageRequest) (response *DescribeWebHookReceiverUsageResponse, err error) {
    if request == nil {
        request = NewDescribeWebHookReceiverUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebHookReceiverUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebHookReceiverUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebHookRuleRequest() (request *DescribeWebHookRuleRequest) {
    request = &DescribeWebHookRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookRule")
    
    
    return
}

func NewDescribeWebHookRuleResponse() (response *DescribeWebHookRuleResponse) {
    response = &DescribeWebHookRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebHookRule
// 获取企微机器人规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebHookRule(request *DescribeWebHookRuleRequest) (response *DescribeWebHookRuleResponse, err error) {
    return c.DescribeWebHookRuleWithContext(context.Background(), request)
}

// DescribeWebHookRule
// 获取企微机器人规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebHookRuleWithContext(ctx context.Context, request *DescribeWebHookRuleRequest) (response *DescribeWebHookRuleResponse, err error) {
    if request == nil {
        request = NewDescribeWebHookRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebHookRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebHookRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebHookRulesRequest() (request *DescribeWebHookRulesRequest) {
    request = &DescribeWebHookRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebHookRules")
    
    
    return
}

func NewDescribeWebHookRulesResponse() (response *DescribeWebHookRulesResponse) {
    response = &DescribeWebHookRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebHookRules
// 获取企微机器人规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebHookRules(request *DescribeWebHookRulesRequest) (response *DescribeWebHookRulesResponse, err error) {
    return c.DescribeWebHookRulesWithContext(context.Background(), request)
}

// DescribeWebHookRules
// 获取企微机器人规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebHookRulesWithContext(ctx context.Context, request *DescribeWebHookRulesRequest) (response *DescribeWebHookRulesResponse, err error) {
    if request == nil {
        request = NewDescribeWebHookRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebHookRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebHookRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebPageEventListRequest() (request *DescribeWebPageEventListRequest) {
    request = &DescribeWebPageEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageEventList")
    
    
    return
}

func NewDescribeWebPageEventListResponse() (response *DescribeWebPageEventListResponse) {
    response = &DescribeWebPageEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebPageEventList
// 查询篡改事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWebPageEventList(request *DescribeWebPageEventListRequest) (response *DescribeWebPageEventListResponse, err error) {
    return c.DescribeWebPageEventListWithContext(context.Background(), request)
}

// DescribeWebPageEventList
// 查询篡改事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeWebPageEventListWithContext(ctx context.Context, request *DescribeWebPageEventListRequest) (response *DescribeWebPageEventListResponse, err error) {
    if request == nil {
        request = NewDescribeWebPageEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebPageEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebPageEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebPageGeneralizeRequest() (request *DescribeWebPageGeneralizeRequest) {
    request = &DescribeWebPageGeneralizeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageGeneralize")
    
    
    return
}

func NewDescribeWebPageGeneralizeResponse() (response *DescribeWebPageGeneralizeResponse) {
    response = &DescribeWebPageGeneralizeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebPageGeneralize
// 查询网站防篡改概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageGeneralize(request *DescribeWebPageGeneralizeRequest) (response *DescribeWebPageGeneralizeResponse, err error) {
    return c.DescribeWebPageGeneralizeWithContext(context.Background(), request)
}

// DescribeWebPageGeneralize
// 查询网站防篡改概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageGeneralizeWithContext(ctx context.Context, request *DescribeWebPageGeneralizeRequest) (response *DescribeWebPageGeneralizeResponse, err error) {
    if request == nil {
        request = NewDescribeWebPageGeneralizeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebPageGeneralize require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebPageGeneralizeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebPageProtectStatRequest() (request *DescribeWebPageProtectStatRequest) {
    request = &DescribeWebPageProtectStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageProtectStat")
    
    
    return
}

func NewDescribeWebPageProtectStatResponse() (response *DescribeWebPageProtectStatResponse) {
    response = &DescribeWebPageProtectStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebPageProtectStat
// 网站防篡改-查询动态防护信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageProtectStat(request *DescribeWebPageProtectStatRequest) (response *DescribeWebPageProtectStatResponse, err error) {
    return c.DescribeWebPageProtectStatWithContext(context.Background(), request)
}

// DescribeWebPageProtectStat
// 网站防篡改-查询动态防护信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageProtectStatWithContext(ctx context.Context, request *DescribeWebPageProtectStatRequest) (response *DescribeWebPageProtectStatResponse, err error) {
    if request == nil {
        request = NewDescribeWebPageProtectStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebPageProtectStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebPageProtectStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebPageServiceInfoRequest() (request *DescribeWebPageServiceInfoRequest) {
    request = &DescribeWebPageServiceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWebPageServiceInfo")
    
    
    return
}

func NewDescribeWebPageServiceInfoResponse() (response *DescribeWebPageServiceInfoResponse) {
    response = &DescribeWebPageServiceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebPageServiceInfo
// 网站防篡改-查询网页防篡改服务器购买信息及服务器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageServiceInfo(request *DescribeWebPageServiceInfoRequest) (response *DescribeWebPageServiceInfoResponse, err error) {
    return c.DescribeWebPageServiceInfoWithContext(context.Background(), request)
}

// DescribeWebPageServiceInfo
// 网站防篡改-查询网页防篡改服务器购买信息及服务器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWebPageServiceInfoWithContext(ctx context.Context, request *DescribeWebPageServiceInfoRequest) (response *DescribeWebPageServiceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWebPageServiceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebPageServiceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebPageServiceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyOrderRequest() (request *DestroyOrderRequest) {
    request = &DestroyOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DestroyOrder")
    
    
    return
}

func NewDestroyOrderResponse() (response *DestroyOrderResponse) {
    response = &DestroyOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyOrder
// DestroyOrder  该接口可以对资源销毁.
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DestroyOrder(request *DestroyOrderRequest) (response *DestroyOrderResponse, err error) {
    return c.DestroyOrderWithContext(context.Background(), request)
}

// DestroyOrder
// DestroyOrder  该接口可以对资源销毁.
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DestroyOrderWithContext(ctx context.Context, request *DestroyOrderRequest) (response *DestroyOrderResponse, err error) {
    if request == nil {
        request = NewDestroyOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyOrderResponse()
    err = c.Send(request, response)
    return
}

func NewEditBashRulesRequest() (request *EditBashRulesRequest) {
    request = &EditBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "EditBashRules")
    
    
    return
}

func NewEditBashRulesResponse() (response *EditBashRulesResponse) {
    response = &EditBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditBashRules
// 新增或修改高危命令规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditBashRules(request *EditBashRulesRequest) (response *EditBashRulesResponse, err error) {
    return c.EditBashRulesWithContext(context.Background(), request)
}

// EditBashRules
// 新增或修改高危命令规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditBashRulesWithContext(ctx context.Context, request *EditBashRulesRequest) (response *EditBashRulesResponse, err error) {
    if request == nil {
        request = NewEditBashRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditBashRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewEditPrivilegeRulesRequest() (request *EditPrivilegeRulesRequest) {
    request = &EditPrivilegeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "EditPrivilegeRules")
    
    
    return
}

func NewEditPrivilegeRulesResponse() (response *EditPrivilegeRulesResponse) {
    response = &EditPrivilegeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditPrivilegeRules
// 新增或修改本地提权规则（支持多服务器选择）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditPrivilegeRules(request *EditPrivilegeRulesRequest) (response *EditPrivilegeRulesResponse, err error) {
    return c.EditPrivilegeRulesWithContext(context.Background(), request)
}

// EditPrivilegeRules
// 新增或修改本地提权规则（支持多服务器选择）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditPrivilegeRulesWithContext(ctx context.Context, request *EditPrivilegeRulesRequest) (response *EditPrivilegeRulesResponse, err error) {
    if request == nil {
        request = NewEditPrivilegeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditPrivilegeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditPrivilegeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewEditReverseShellRulesRequest() (request *EditReverseShellRulesRequest) {
    request = &EditReverseShellRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "EditReverseShellRules")
    
    
    return
}

func NewEditReverseShellRulesResponse() (response *EditReverseShellRulesResponse) {
    response = &EditReverseShellRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditReverseShellRules
// 编辑反弹Shell规则（支持多服务器选择）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditReverseShellRules(request *EditReverseShellRulesRequest) (response *EditReverseShellRulesResponse, err error) {
    return c.EditReverseShellRulesWithContext(context.Background(), request)
}

// EditReverseShellRules
// 编辑反弹Shell规则（支持多服务器选择）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditReverseShellRulesWithContext(ctx context.Context, request *EditReverseShellRulesRequest) (response *EditReverseShellRulesResponse, err error) {
    if request == nil {
        request = NewEditReverseShellRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditReverseShellRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditReverseShellRulesResponse()
    err = c.Send(request, response)
    return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
    request = &EditTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "EditTags")
    
    
    return
}

func NewEditTagsResponse() (response *EditTagsResponse) {
    response = &EditTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditTags
// 新增或编辑标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
    return c.EditTagsWithContext(context.Background(), request)
}

// EditTags
// 新增或编辑标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditTagsWithContext(ctx context.Context, request *EditTagsRequest) (response *EditTagsResponse, err error) {
    if request == nil {
        request = NewEditTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditTagsResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetAppListRequest() (request *ExportAssetAppListRequest) {
    request = &ExportAssetAppListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetAppList")
    
    
    return
}

func NewExportAssetAppListResponse() (response *ExportAssetAppListResponse) {
    response = &ExportAssetAppListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetAppList
// 导出资产管理应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetAppList(request *ExportAssetAppListRequest) (response *ExportAssetAppListResponse, err error) {
    return c.ExportAssetAppListWithContext(context.Background(), request)
}

// ExportAssetAppList
// 导出资产管理应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetAppListWithContext(ctx context.Context, request *ExportAssetAppListRequest) (response *ExportAssetAppListResponse, err error) {
    if request == nil {
        request = NewExportAssetAppListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetAppList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetAppListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetCoreModuleListRequest() (request *ExportAssetCoreModuleListRequest) {
    request = &ExportAssetCoreModuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetCoreModuleList")
    
    
    return
}

func NewExportAssetCoreModuleListResponse() (response *ExportAssetCoreModuleListResponse) {
    response = &ExportAssetCoreModuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetCoreModuleList
// 导出资产管理内核模块列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetCoreModuleList(request *ExportAssetCoreModuleListRequest) (response *ExportAssetCoreModuleListResponse, err error) {
    return c.ExportAssetCoreModuleListWithContext(context.Background(), request)
}

// ExportAssetCoreModuleList
// 导出资产管理内核模块列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetCoreModuleListWithContext(ctx context.Context, request *ExportAssetCoreModuleListRequest) (response *ExportAssetCoreModuleListResponse, err error) {
    if request == nil {
        request = NewExportAssetCoreModuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetCoreModuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetCoreModuleListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetDatabaseListRequest() (request *ExportAssetDatabaseListRequest) {
    request = &ExportAssetDatabaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetDatabaseList")
    
    
    return
}

func NewExportAssetDatabaseListResponse() (response *ExportAssetDatabaseListResponse) {
    response = &ExportAssetDatabaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetDatabaseList
// 导出资产管理数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetDatabaseList(request *ExportAssetDatabaseListRequest) (response *ExportAssetDatabaseListResponse, err error) {
    return c.ExportAssetDatabaseListWithContext(context.Background(), request)
}

// ExportAssetDatabaseList
// 导出资产管理数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetDatabaseListWithContext(ctx context.Context, request *ExportAssetDatabaseListRequest) (response *ExportAssetDatabaseListResponse, err error) {
    if request == nil {
        request = NewExportAssetDatabaseListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetDatabaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetDatabaseListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetEnvListRequest() (request *ExportAssetEnvListRequest) {
    request = &ExportAssetEnvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetEnvList")
    
    
    return
}

func NewExportAssetEnvListResponse() (response *ExportAssetEnvListResponse) {
    response = &ExportAssetEnvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetEnvList
// 导出资产管理环境变量列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetEnvList(request *ExportAssetEnvListRequest) (response *ExportAssetEnvListResponse, err error) {
    return c.ExportAssetEnvListWithContext(context.Background(), request)
}

// ExportAssetEnvList
// 导出资产管理环境变量列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetEnvListWithContext(ctx context.Context, request *ExportAssetEnvListRequest) (response *ExportAssetEnvListResponse, err error) {
    if request == nil {
        request = NewExportAssetEnvListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetEnvList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetEnvListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetInitServiceListRequest() (request *ExportAssetInitServiceListRequest) {
    request = &ExportAssetInitServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetInitServiceList")
    
    
    return
}

func NewExportAssetInitServiceListResponse() (response *ExportAssetInitServiceListResponse) {
    response = &ExportAssetInitServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetInitServiceList
// 导出资产管理启动服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetInitServiceList(request *ExportAssetInitServiceListRequest) (response *ExportAssetInitServiceListResponse, err error) {
    return c.ExportAssetInitServiceListWithContext(context.Background(), request)
}

// ExportAssetInitServiceList
// 导出资产管理启动服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetInitServiceListWithContext(ctx context.Context, request *ExportAssetInitServiceListRequest) (response *ExportAssetInitServiceListResponse, err error) {
    if request == nil {
        request = NewExportAssetInitServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetInitServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetInitServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetJarListRequest() (request *ExportAssetJarListRequest) {
    request = &ExportAssetJarListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetJarList")
    
    
    return
}

func NewExportAssetJarListResponse() (response *ExportAssetJarListResponse) {
    response = &ExportAssetJarListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetJarList
// 导出Jar包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetJarList(request *ExportAssetJarListRequest) (response *ExportAssetJarListResponse, err error) {
    return c.ExportAssetJarListWithContext(context.Background(), request)
}

// ExportAssetJarList
// 导出Jar包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetJarListWithContext(ctx context.Context, request *ExportAssetJarListRequest) (response *ExportAssetJarListResponse, err error) {
    if request == nil {
        request = NewExportAssetJarListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetJarList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetJarListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetMachineDetailRequest() (request *ExportAssetMachineDetailRequest) {
    request = &ExportAssetMachineDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetMachineDetail")
    
    
    return
}

func NewExportAssetMachineDetailResponse() (response *ExportAssetMachineDetailResponse) {
    response = &ExportAssetMachineDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetMachineDetail
// 导出资产管理主机资源详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetMachineDetail(request *ExportAssetMachineDetailRequest) (response *ExportAssetMachineDetailResponse, err error) {
    return c.ExportAssetMachineDetailWithContext(context.Background(), request)
}

// ExportAssetMachineDetail
// 导出资产管理主机资源详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetMachineDetailWithContext(ctx context.Context, request *ExportAssetMachineDetailRequest) (response *ExportAssetMachineDetailResponse, err error) {
    if request == nil {
        request = NewExportAssetMachineDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetMachineDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetMachineDetailResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetMachineListRequest() (request *ExportAssetMachineListRequest) {
    request = &ExportAssetMachineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetMachineList")
    
    
    return
}

func NewExportAssetMachineListResponse() (response *ExportAssetMachineListResponse) {
    response = &ExportAssetMachineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetMachineList
// 导出资源监控列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetMachineList(request *ExportAssetMachineListRequest) (response *ExportAssetMachineListResponse, err error) {
    return c.ExportAssetMachineListWithContext(context.Background(), request)
}

// ExportAssetMachineList
// 导出资源监控列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetMachineListWithContext(ctx context.Context, request *ExportAssetMachineListRequest) (response *ExportAssetMachineListResponse, err error) {
    if request == nil {
        request = NewExportAssetMachineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetMachineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetMachineListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetPlanTaskListRequest() (request *ExportAssetPlanTaskListRequest) {
    request = &ExportAssetPlanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetPlanTaskList")
    
    
    return
}

func NewExportAssetPlanTaskListResponse() (response *ExportAssetPlanTaskListResponse) {
    response = &ExportAssetPlanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetPlanTaskList
// 导出资产管理计划任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetPlanTaskList(request *ExportAssetPlanTaskListRequest) (response *ExportAssetPlanTaskListResponse, err error) {
    return c.ExportAssetPlanTaskListWithContext(context.Background(), request)
}

// ExportAssetPlanTaskList
// 导出资产管理计划任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetPlanTaskListWithContext(ctx context.Context, request *ExportAssetPlanTaskListRequest) (response *ExportAssetPlanTaskListResponse, err error) {
    if request == nil {
        request = NewExportAssetPlanTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetPlanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetPlanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetPortInfoListRequest() (request *ExportAssetPortInfoListRequest) {
    request = &ExportAssetPortInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetPortInfoList")
    
    
    return
}

func NewExportAssetPortInfoListResponse() (response *ExportAssetPortInfoListResponse) {
    response = &ExportAssetPortInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetPortInfoList
// 导出资产管理端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetPortInfoList(request *ExportAssetPortInfoListRequest) (response *ExportAssetPortInfoListResponse, err error) {
    return c.ExportAssetPortInfoListWithContext(context.Background(), request)
}

// ExportAssetPortInfoList
// 导出资产管理端口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetPortInfoListWithContext(ctx context.Context, request *ExportAssetPortInfoListRequest) (response *ExportAssetPortInfoListResponse, err error) {
    if request == nil {
        request = NewExportAssetPortInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetPortInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetPortInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetProcessInfoListRequest() (request *ExportAssetProcessInfoListRequest) {
    request = &ExportAssetProcessInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetProcessInfoList")
    
    
    return
}

func NewExportAssetProcessInfoListResponse() (response *ExportAssetProcessInfoListResponse) {
    response = &ExportAssetProcessInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetProcessInfoList
// 导出资产管理进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetProcessInfoList(request *ExportAssetProcessInfoListRequest) (response *ExportAssetProcessInfoListResponse, err error) {
    return c.ExportAssetProcessInfoListWithContext(context.Background(), request)
}

// ExportAssetProcessInfoList
// 导出资产管理进程列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetProcessInfoListWithContext(ctx context.Context, request *ExportAssetProcessInfoListRequest) (response *ExportAssetProcessInfoListResponse, err error) {
    if request == nil {
        request = NewExportAssetProcessInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetProcessInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetProcessInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetRecentMachineInfoRequest() (request *ExportAssetRecentMachineInfoRequest) {
    request = &ExportAssetRecentMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetRecentMachineInfo")
    
    
    return
}

func NewExportAssetRecentMachineInfoResponse() (response *ExportAssetRecentMachineInfoResponse) {
    response = &ExportAssetRecentMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetRecentMachineInfo
// 导出主机最近趋势情况（最长最近90天）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetRecentMachineInfo(request *ExportAssetRecentMachineInfoRequest) (response *ExportAssetRecentMachineInfoResponse, err error) {
    return c.ExportAssetRecentMachineInfoWithContext(context.Background(), request)
}

// ExportAssetRecentMachineInfo
// 导出主机最近趋势情况（最长最近90天）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetRecentMachineInfoWithContext(ctx context.Context, request *ExportAssetRecentMachineInfoRequest) (response *ExportAssetRecentMachineInfoResponse, err error) {
    if request == nil {
        request = NewExportAssetRecentMachineInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetRecentMachineInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetRecentMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetSystemPackageListRequest() (request *ExportAssetSystemPackageListRequest) {
    request = &ExportAssetSystemPackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetSystemPackageList")
    
    
    return
}

func NewExportAssetSystemPackageListResponse() (response *ExportAssetSystemPackageListResponse) {
    response = &ExportAssetSystemPackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetSystemPackageList
// 导出资产管理系统安装包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetSystemPackageList(request *ExportAssetSystemPackageListRequest) (response *ExportAssetSystemPackageListResponse, err error) {
    return c.ExportAssetSystemPackageListWithContext(context.Background(), request)
}

// ExportAssetSystemPackageList
// 导出资产管理系统安装包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetSystemPackageListWithContext(ctx context.Context, request *ExportAssetSystemPackageListRequest) (response *ExportAssetSystemPackageListResponse, err error) {
    if request == nil {
        request = NewExportAssetSystemPackageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetSystemPackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetSystemPackageListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetUserListRequest() (request *ExportAssetUserListRequest) {
    request = &ExportAssetUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetUserList")
    
    
    return
}

func NewExportAssetUserListResponse() (response *ExportAssetUserListResponse) {
    response = &ExportAssetUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetUserList
// 导出账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
func (c *Client) ExportAssetUserList(request *ExportAssetUserListRequest) (response *ExportAssetUserListResponse, err error) {
    return c.ExportAssetUserListWithContext(context.Background(), request)
}

// ExportAssetUserList
// 导出账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
func (c *Client) ExportAssetUserListWithContext(ctx context.Context, request *ExportAssetUserListRequest) (response *ExportAssetUserListResponse, err error) {
    if request == nil {
        request = NewExportAssetUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetUserListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetWebAppListRequest() (request *ExportAssetWebAppListRequest) {
    request = &ExportAssetWebAppListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebAppList")
    
    
    return
}

func NewExportAssetWebAppListResponse() (response *ExportAssetWebAppListResponse) {
    response = &ExportAssetWebAppListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetWebAppList
// 导出资产管理Web应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebAppList(request *ExportAssetWebAppListRequest) (response *ExportAssetWebAppListResponse, err error) {
    return c.ExportAssetWebAppListWithContext(context.Background(), request)
}

// ExportAssetWebAppList
// 导出资产管理Web应用列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebAppListWithContext(ctx context.Context, request *ExportAssetWebAppListRequest) (response *ExportAssetWebAppListResponse, err error) {
    if request == nil {
        request = NewExportAssetWebAppListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetWebAppList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetWebAppListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetWebFrameListRequest() (request *ExportAssetWebFrameListRequest) {
    request = &ExportAssetWebFrameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebFrameList")
    
    
    return
}

func NewExportAssetWebFrameListResponse() (response *ExportAssetWebFrameListResponse) {
    response = &ExportAssetWebFrameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetWebFrameList
// 导出资产管理Web框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebFrameList(request *ExportAssetWebFrameListRequest) (response *ExportAssetWebFrameListResponse, err error) {
    return c.ExportAssetWebFrameListWithContext(context.Background(), request)
}

// ExportAssetWebFrameList
// 导出资产管理Web框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebFrameListWithContext(ctx context.Context, request *ExportAssetWebFrameListRequest) (response *ExportAssetWebFrameListResponse, err error) {
    if request == nil {
        request = NewExportAssetWebFrameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetWebFrameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetWebFrameListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetWebLocationListRequest() (request *ExportAssetWebLocationListRequest) {
    request = &ExportAssetWebLocationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebLocationList")
    
    
    return
}

func NewExportAssetWebLocationListResponse() (response *ExportAssetWebLocationListResponse) {
    response = &ExportAssetWebLocationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetWebLocationList
// 导出Web站点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebLocationList(request *ExportAssetWebLocationListRequest) (response *ExportAssetWebLocationListResponse, err error) {
    return c.ExportAssetWebLocationListWithContext(context.Background(), request)
}

// ExportAssetWebLocationList
// 导出Web站点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportAssetWebLocationListWithContext(ctx context.Context, request *ExportAssetWebLocationListRequest) (response *ExportAssetWebLocationListResponse, err error) {
    if request == nil {
        request = NewExportAssetWebLocationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetWebLocationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetWebLocationListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetWebServiceInfoListRequest() (request *ExportAssetWebServiceInfoListRequest) {
    request = &ExportAssetWebServiceInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAssetWebServiceInfoList")
    
    
    return
}

func NewExportAssetWebServiceInfoListResponse() (response *ExportAssetWebServiceInfoListResponse) {
    response = &ExportAssetWebServiceInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAssetWebServiceInfoList
// 导出资产管理Web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetWebServiceInfoList(request *ExportAssetWebServiceInfoListRequest) (response *ExportAssetWebServiceInfoListResponse, err error) {
    return c.ExportAssetWebServiceInfoListWithContext(context.Background(), request)
}

// ExportAssetWebServiceInfoList
// 导出资产管理Web服务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetWebServiceInfoListWithContext(ctx context.Context, request *ExportAssetWebServiceInfoListRequest) (response *ExportAssetWebServiceInfoListResponse, err error) {
    if request == nil {
        request = NewExportAssetWebServiceInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetWebServiceInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetWebServiceInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewExportAttackEventsRequest() (request *ExportAttackEventsRequest) {
    request = &ExportAttackEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAttackEvents")
    
    
    return
}

func NewExportAttackEventsResponse() (response *ExportAttackEventsResponse) {
    response = &ExportAttackEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAttackEvents
// 导出网络攻击事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAttackEvents(request *ExportAttackEventsRequest) (response *ExportAttackEventsResponse, err error) {
    return c.ExportAttackEventsWithContext(context.Background(), request)
}

// ExportAttackEvents
// 导出网络攻击事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAttackEventsWithContext(ctx context.Context, request *ExportAttackEventsRequest) (response *ExportAttackEventsResponse, err error) {
    if request == nil {
        request = NewExportAttackEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAttackEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAttackEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportAttackLogsRequest() (request *ExportAttackLogsRequest) {
    request = &ExportAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportAttackLogs")
    
    
    return
}

func NewExportAttackLogsResponse() (response *ExportAttackLogsResponse) {
    response = &ExportAttackLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportAttackLogs
// 导出网络攻击日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAttackLogs(request *ExportAttackLogsRequest) (response *ExportAttackLogsResponse, err error) {
    return c.ExportAttackLogsWithContext(context.Background(), request)
}

// ExportAttackLogs
// 导出网络攻击日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAttackLogsWithContext(ctx context.Context, request *ExportAttackLogsRequest) (response *ExportAttackLogsResponse, err error) {
    if request == nil {
        request = NewExportAttackLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAttackLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAttackLogsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineEffectHostListRequest() (request *ExportBaselineEffectHostListRequest) {
    request = &ExportBaselineEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineEffectHostList")
    
    
    return
}

func NewExportBaselineEffectHostListResponse() (response *ExportBaselineEffectHostListResponse) {
    response = &ExportBaselineEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineEffectHostList
// 导出基线影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportBaselineEffectHostList(request *ExportBaselineEffectHostListRequest) (response *ExportBaselineEffectHostListResponse, err error) {
    return c.ExportBaselineEffectHostListWithContext(context.Background(), request)
}

// ExportBaselineEffectHostList
// 导出基线影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportBaselineEffectHostListWithContext(ctx context.Context, request *ExportBaselineEffectHostListRequest) (response *ExportBaselineEffectHostListResponse, err error) {
    if request == nil {
        request = NewExportBaselineEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineFixListRequest() (request *ExportBaselineFixListRequest) {
    request = &ExportBaselineFixListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineFixList")
    
    
    return
}

func NewExportBaselineFixListResponse() (response *ExportBaselineFixListResponse) {
    response = &ExportBaselineFixListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineFixList
// 导出修复列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineFixList(request *ExportBaselineFixListRequest) (response *ExportBaselineFixListResponse, err error) {
    return c.ExportBaselineFixListWithContext(context.Background(), request)
}

// ExportBaselineFixList
// 导出修复列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineFixListWithContext(ctx context.Context, request *ExportBaselineFixListRequest) (response *ExportBaselineFixListResponse, err error) {
    if request == nil {
        request = NewExportBaselineFixListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineFixList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineFixListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineHostDetectListRequest() (request *ExportBaselineHostDetectListRequest) {
    request = &ExportBaselineHostDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineHostDetectList")
    
    
    return
}

func NewExportBaselineHostDetectListResponse() (response *ExportBaselineHostDetectListResponse) {
    response = &ExportBaselineHostDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineHostDetectList
// 导出基线主机检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineHostDetectList(request *ExportBaselineHostDetectListRequest) (response *ExportBaselineHostDetectListResponse, err error) {
    return c.ExportBaselineHostDetectListWithContext(context.Background(), request)
}

// ExportBaselineHostDetectList
// 导出基线主机检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineHostDetectListWithContext(ctx context.Context, request *ExportBaselineHostDetectListRequest) (response *ExportBaselineHostDetectListResponse, err error) {
    if request == nil {
        request = NewExportBaselineHostDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineHostDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineHostDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineItemDetectListRequest() (request *ExportBaselineItemDetectListRequest) {
    request = &ExportBaselineItemDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineItemDetectList")
    
    
    return
}

func NewExportBaselineItemDetectListResponse() (response *ExportBaselineItemDetectListResponse) {
    response = &ExportBaselineItemDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineItemDetectList
// 导出基线检测项
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineItemDetectList(request *ExportBaselineItemDetectListRequest) (response *ExportBaselineItemDetectListResponse, err error) {
    return c.ExportBaselineItemDetectListWithContext(context.Background(), request)
}

// ExportBaselineItemDetectList
// 导出基线检测项
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineItemDetectListWithContext(ctx context.Context, request *ExportBaselineItemDetectListRequest) (response *ExportBaselineItemDetectListResponse, err error) {
    if request == nil {
        request = NewExportBaselineItemDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineItemDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineItemDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineItemListRequest() (request *ExportBaselineItemListRequest) {
    request = &ExportBaselineItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineItemList")
    
    
    return
}

func NewExportBaselineItemListResponse() (response *ExportBaselineItemListResponse) {
    response = &ExportBaselineItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineItemList
// 导出检测项结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineItemList(request *ExportBaselineItemListRequest) (response *ExportBaselineItemListResponse, err error) {
    return c.ExportBaselineItemListWithContext(context.Background(), request)
}

// ExportBaselineItemList
// 导出检测项结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineItemListWithContext(ctx context.Context, request *ExportBaselineItemListRequest) (response *ExportBaselineItemListResponse, err error) {
    if request == nil {
        request = NewExportBaselineItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineItemListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineListRequest() (request *ExportBaselineListRequest) {
    request = &ExportBaselineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineList")
    
    
    return
}

func NewExportBaselineListResponse() (response *ExportBaselineListResponse) {
    response = &ExportBaselineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineList
// 导出基线列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportBaselineList(request *ExportBaselineListRequest) (response *ExportBaselineListResponse, err error) {
    return c.ExportBaselineListWithContext(context.Background(), request)
}

// ExportBaselineList
// 导出基线列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportBaselineListWithContext(ctx context.Context, request *ExportBaselineListRequest) (response *ExportBaselineListResponse, err error) {
    if request == nil {
        request = NewExportBaselineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineRuleDetectListRequest() (request *ExportBaselineRuleDetectListRequest) {
    request = &ExportBaselineRuleDetectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineRuleDetectList")
    
    
    return
}

func NewExportBaselineRuleDetectListResponse() (response *ExportBaselineRuleDetectListResponse) {
    response = &ExportBaselineRuleDetectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineRuleDetectList
// 导出基线检测规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineRuleDetectList(request *ExportBaselineRuleDetectListRequest) (response *ExportBaselineRuleDetectListResponse, err error) {
    return c.ExportBaselineRuleDetectListWithContext(context.Background(), request)
}

// ExportBaselineRuleDetectList
// 导出基线检测规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineRuleDetectListWithContext(ctx context.Context, request *ExportBaselineRuleDetectListRequest) (response *ExportBaselineRuleDetectListResponse, err error) {
    if request == nil {
        request = NewExportBaselineRuleDetectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineRuleDetectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineRuleDetectListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBaselineWeakPasswordListRequest() (request *ExportBaselineWeakPasswordListRequest) {
    request = &ExportBaselineWeakPasswordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBaselineWeakPasswordList")
    
    
    return
}

func NewExportBaselineWeakPasswordListResponse() (response *ExportBaselineWeakPasswordListResponse) {
    response = &ExportBaselineWeakPasswordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBaselineWeakPasswordList
// 导出弱口令配置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineWeakPasswordList(request *ExportBaselineWeakPasswordListRequest) (response *ExportBaselineWeakPasswordListResponse, err error) {
    return c.ExportBaselineWeakPasswordListWithContext(context.Background(), request)
}

// ExportBaselineWeakPasswordList
// 导出弱口令配置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportBaselineWeakPasswordListWithContext(ctx context.Context, request *ExportBaselineWeakPasswordListRequest) (response *ExportBaselineWeakPasswordListResponse, err error) {
    if request == nil {
        request = NewExportBaselineWeakPasswordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBaselineWeakPasswordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBaselineWeakPasswordListResponse()
    err = c.Send(request, response)
    return
}

func NewExportBashEventsRequest() (request *ExportBashEventsRequest) {
    request = &ExportBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBashEvents")
    
    
    return
}

func NewExportBashEventsResponse() (response *ExportBashEventsResponse) {
    response = &ExportBashEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBashEvents
// 导出高危命令事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashEvents(request *ExportBashEventsRequest) (response *ExportBashEventsResponse, err error) {
    return c.ExportBashEventsWithContext(context.Background(), request)
}

// ExportBashEvents
// 导出高危命令事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashEventsWithContext(ctx context.Context, request *ExportBashEventsRequest) (response *ExportBashEventsResponse, err error) {
    if request == nil {
        request = NewExportBashEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBashEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBashEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBashEventsNewRequest() (request *ExportBashEventsNewRequest) {
    request = &ExportBashEventsNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBashEventsNew")
    
    
    return
}

func NewExportBashEventsNewResponse() (response *ExportBashEventsNewResponse) {
    response = &ExportBashEventsNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBashEventsNew
// 导出高危命令事件(新)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashEventsNew(request *ExportBashEventsNewRequest) (response *ExportBashEventsNewResponse, err error) {
    return c.ExportBashEventsNewWithContext(context.Background(), request)
}

// ExportBashEventsNew
// 导出高危命令事件(新)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashEventsNewWithContext(ctx context.Context, request *ExportBashEventsNewRequest) (response *ExportBashEventsNewResponse, err error) {
    if request == nil {
        request = NewExportBashEventsNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBashEventsNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBashEventsNewResponse()
    err = c.Send(request, response)
    return
}

func NewExportBashPoliciesRequest() (request *ExportBashPoliciesRequest) {
    request = &ExportBashPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBashPolicies")
    
    
    return
}

func NewExportBashPoliciesResponse() (response *ExportBashPoliciesResponse) {
    response = &ExportBashPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBashPolicies
// 导出高危命令策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashPolicies(request *ExportBashPoliciesRequest) (response *ExportBashPoliciesResponse, err error) {
    return c.ExportBashPoliciesWithContext(context.Background(), request)
}

// ExportBashPolicies
// 导出高危命令策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBashPoliciesWithContext(ctx context.Context, request *ExportBashPoliciesRequest) (response *ExportBashPoliciesResponse, err error) {
    if request == nil {
        request = NewExportBashPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBashPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBashPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewExportBruteAttacksRequest() (request *ExportBruteAttacksRequest) {
    request = &ExportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportBruteAttacks")
    
    
    return
}

func NewExportBruteAttacksResponse() (response *ExportBruteAttacksResponse) {
    response = &ExportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBruteAttacks
// 本接口 (ExportBruteAttacks) 用于导出密码破解记录成CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    return c.ExportBruteAttacksWithContext(context.Background(), request)
}

// ExportBruteAttacks
// 本接口 (ExportBruteAttacks) 用于导出密码破解记录成CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportBruteAttacksWithContext(ctx context.Context, request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewExportFileTamperEventsRequest() (request *ExportFileTamperEventsRequest) {
    request = &ExportFileTamperEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportFileTamperEvents")
    
    
    return
}

func NewExportFileTamperEventsResponse() (response *ExportFileTamperEventsResponse) {
    response = &ExportFileTamperEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportFileTamperEvents
// 导出核心文件事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportFileTamperEvents(request *ExportFileTamperEventsRequest) (response *ExportFileTamperEventsResponse, err error) {
    return c.ExportFileTamperEventsWithContext(context.Background(), request)
}

// ExportFileTamperEvents
// 导出核心文件事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportFileTamperEventsWithContext(ctx context.Context, request *ExportFileTamperEventsRequest) (response *ExportFileTamperEventsResponse, err error) {
    if request == nil {
        request = NewExportFileTamperEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportFileTamperEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportFileTamperEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportFileTamperRulesRequest() (request *ExportFileTamperRulesRequest) {
    request = &ExportFileTamperRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportFileTamperRules")
    
    
    return
}

func NewExportFileTamperRulesResponse() (response *ExportFileTamperRulesResponse) {
    response = &ExportFileTamperRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportFileTamperRules
// 导出核心文件监控规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportFileTamperRules(request *ExportFileTamperRulesRequest) (response *ExportFileTamperRulesResponse, err error) {
    return c.ExportFileTamperRulesWithContext(context.Background(), request)
}

// ExportFileTamperRules
// 导出核心文件监控规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportFileTamperRulesWithContext(ctx context.Context, request *ExportFileTamperRulesRequest) (response *ExportFileTamperRulesResponse, err error) {
    if request == nil {
        request = NewExportFileTamperRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportFileTamperRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportFileTamperRulesResponse()
    err = c.Send(request, response)
    return
}

func NewExportIgnoreBaselineRuleRequest() (request *ExportIgnoreBaselineRuleRequest) {
    request = &ExportIgnoreBaselineRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportIgnoreBaselineRule")
    
    
    return
}

func NewExportIgnoreBaselineRuleResponse() (response *ExportIgnoreBaselineRuleResponse) {
    response = &ExportIgnoreBaselineRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportIgnoreBaselineRule
// 导出已忽略基线检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportIgnoreBaselineRule(request *ExportIgnoreBaselineRuleRequest) (response *ExportIgnoreBaselineRuleResponse, err error) {
    return c.ExportIgnoreBaselineRuleWithContext(context.Background(), request)
}

// ExportIgnoreBaselineRule
// 导出已忽略基线检测项信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportIgnoreBaselineRuleWithContext(ctx context.Context, request *ExportIgnoreBaselineRuleRequest) (response *ExportIgnoreBaselineRuleResponse, err error) {
    if request == nil {
        request = NewExportIgnoreBaselineRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportIgnoreBaselineRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportIgnoreBaselineRuleResponse()
    err = c.Send(request, response)
    return
}

func NewExportIgnoreRuleEffectHostListRequest() (request *ExportIgnoreRuleEffectHostListRequest) {
    request = &ExportIgnoreRuleEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportIgnoreRuleEffectHostList")
    
    
    return
}

func NewExportIgnoreRuleEffectHostListResponse() (response *ExportIgnoreRuleEffectHostListResponse) {
    response = &ExportIgnoreRuleEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportIgnoreRuleEffectHostList
// 根据检测项id导出忽略检测项影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportIgnoreRuleEffectHostList(request *ExportIgnoreRuleEffectHostListRequest) (response *ExportIgnoreRuleEffectHostListResponse, err error) {
    return c.ExportIgnoreRuleEffectHostListWithContext(context.Background(), request)
}

// ExportIgnoreRuleEffectHostList
// 根据检测项id导出忽略检测项影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportIgnoreRuleEffectHostListWithContext(ctx context.Context, request *ExportIgnoreRuleEffectHostListRequest) (response *ExportIgnoreRuleEffectHostListResponse, err error) {
    if request == nil {
        request = NewExportIgnoreRuleEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportIgnoreRuleEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportIgnoreRuleEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewExportJavaMemShellPluginsRequest() (request *ExportJavaMemShellPluginsRequest) {
    request = &ExportJavaMemShellPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportJavaMemShellPlugins")
    
    
    return
}

func NewExportJavaMemShellPluginsResponse() (response *ExportJavaMemShellPluginsResponse) {
    response = &ExportJavaMemShellPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportJavaMemShellPlugins
// 导出java内存马插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportJavaMemShellPlugins(request *ExportJavaMemShellPluginsRequest) (response *ExportJavaMemShellPluginsResponse, err error) {
    return c.ExportJavaMemShellPluginsWithContext(context.Background(), request)
}

// ExportJavaMemShellPlugins
// 导出java内存马插件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportJavaMemShellPluginsWithContext(ctx context.Context, request *ExportJavaMemShellPluginsRequest) (response *ExportJavaMemShellPluginsResponse, err error) {
    if request == nil {
        request = NewExportJavaMemShellPluginsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportJavaMemShellPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportJavaMemShellPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewExportJavaMemShellsRequest() (request *ExportJavaMemShellsRequest) {
    request = &ExportJavaMemShellsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportJavaMemShells")
    
    
    return
}

func NewExportJavaMemShellsResponse() (response *ExportJavaMemShellsResponse) {
    response = &ExportJavaMemShellsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportJavaMemShells
// 导出java内存马事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportJavaMemShells(request *ExportJavaMemShellsRequest) (response *ExportJavaMemShellsResponse, err error) {
    return c.ExportJavaMemShellsWithContext(context.Background(), request)
}

// ExportJavaMemShells
// 导出java内存马事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportJavaMemShellsWithContext(ctx context.Context, request *ExportJavaMemShellsRequest) (response *ExportJavaMemShellsResponse, err error) {
    if request == nil {
        request = NewExportJavaMemShellsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportJavaMemShells require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportJavaMemShellsResponse()
    err = c.Send(request, response)
    return
}

func NewExportLicenseDetailRequest() (request *ExportLicenseDetailRequest) {
    request = &ExportLicenseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportLicenseDetail")
    
    
    return
}

func NewExportLicenseDetailResponse() (response *ExportLicenseDetailResponse) {
    response = &ExportLicenseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportLicenseDetail
// 导出授权列表对应的绑定信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportLicenseDetail(request *ExportLicenseDetailRequest) (response *ExportLicenseDetailResponse, err error) {
    return c.ExportLicenseDetailWithContext(context.Background(), request)
}

// ExportLicenseDetail
// 导出授权列表对应的绑定信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportLicenseDetailWithContext(ctx context.Context, request *ExportLicenseDetailRequest) (response *ExportLicenseDetailResponse, err error) {
    if request == nil {
        request = NewExportLicenseDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportLicenseDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportLicenseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
    request = &ExportMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportMaliciousRequests")
    
    
    return
}

func NewExportMaliciousRequestsResponse() (response *ExportMaliciousRequestsResponse) {
    response = &ExportMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportMaliciousRequests
// 本接口 (ExportMaliciousRequests) 用于导出下载恶意请求文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportMaliciousRequests(request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    return c.ExportMaliciousRequestsWithContext(context.Background(), request)
}

// ExportMaliciousRequests
// 本接口 (ExportMaliciousRequests) 用于导出下载恶意请求文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportMaliciousRequestsWithContext(ctx context.Context, request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewExportMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewExportMalwaresRequest() (request *ExportMalwaresRequest) {
    request = &ExportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportMalwares")
    
    
    return
}

func NewExportMalwaresResponse() (response *ExportMalwaresResponse) {
    response = &ExportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportMalwares
// 本接口 (ExportMalwares) 用于导出木马记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    return c.ExportMalwaresWithContext(context.Background(), request)
}

// ExportMalwares
// 本接口 (ExportMalwares) 用于导出木马记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportMalwaresWithContext(ctx context.Context, request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    if request == nil {
        request = NewExportMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewExportNonlocalLoginPlacesRequest() (request *ExportNonlocalLoginPlacesRequest) {
    request = &ExportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportNonlocalLoginPlaces")
    
    
    return
}

func NewExportNonlocalLoginPlacesResponse() (response *ExportNonlocalLoginPlacesResponse) {
    response = &ExportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportNonlocalLoginPlaces
// 本接口 (ExportNonlocalLoginPlaces) 用于导出异地登录事件记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportNonlocalLoginPlaces(request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    return c.ExportNonlocalLoginPlacesWithContext(context.Background(), request)
}

// ExportNonlocalLoginPlaces
// 本接口 (ExportNonlocalLoginPlaces) 用于导出异地登录事件记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportNonlocalLoginPlacesWithContext(ctx context.Context, request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewExportNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewExportPrivilegeEventsRequest() (request *ExportPrivilegeEventsRequest) {
    request = &ExportPrivilegeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportPrivilegeEvents")
    
    
    return
}

func NewExportPrivilegeEventsResponse() (response *ExportPrivilegeEventsResponse) {
    response = &ExportPrivilegeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportPrivilegeEvents
// 导出本地提权事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportPrivilegeEvents(request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
    return c.ExportPrivilegeEventsWithContext(context.Background(), request)
}

// ExportPrivilegeEvents
// 导出本地提权事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ExportPrivilegeEventsWithContext(ctx context.Context, request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewExportPrivilegeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportPrivilegeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportPrivilegeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportProtectDirListRequest() (request *ExportProtectDirListRequest) {
    request = &ExportProtectDirListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportProtectDirList")
    
    
    return
}

func NewExportProtectDirListResponse() (response *ExportProtectDirListResponse) {
    response = &ExportProtectDirListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportProtectDirList
// 导出网页防篡改防护目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportProtectDirList(request *ExportProtectDirListRequest) (response *ExportProtectDirListResponse, err error) {
    return c.ExportProtectDirListWithContext(context.Background(), request)
}

// ExportProtectDirList
// 导出网页防篡改防护目录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportProtectDirListWithContext(ctx context.Context, request *ExportProtectDirListRequest) (response *ExportProtectDirListResponse, err error) {
    if request == nil {
        request = NewExportProtectDirListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportProtectDirList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportProtectDirListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRansomDefenseBackupListRequest() (request *ExportRansomDefenseBackupListRequest) {
    request = &ExportRansomDefenseBackupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseBackupList")
    
    
    return
}

func NewExportRansomDefenseBackupListResponse() (response *ExportRansomDefenseBackupListResponse) {
    response = &ExportRansomDefenseBackupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRansomDefenseBackupList
// 导出主机快照备份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseBackupList(request *ExportRansomDefenseBackupListRequest) (response *ExportRansomDefenseBackupListResponse, err error) {
    return c.ExportRansomDefenseBackupListWithContext(context.Background(), request)
}

// ExportRansomDefenseBackupList
// 导出主机快照备份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseBackupListWithContext(ctx context.Context, request *ExportRansomDefenseBackupListRequest) (response *ExportRansomDefenseBackupListResponse, err error) {
    if request == nil {
        request = NewExportRansomDefenseBackupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRansomDefenseBackupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRansomDefenseBackupListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRansomDefenseEventsListRequest() (request *ExportRansomDefenseEventsListRequest) {
    request = &ExportRansomDefenseEventsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseEventsList")
    
    
    return
}

func NewExportRansomDefenseEventsListResponse() (response *ExportRansomDefenseEventsListResponse) {
    response = &ExportRansomDefenseEventsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRansomDefenseEventsList
// 导出防勒索事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseEventsList(request *ExportRansomDefenseEventsListRequest) (response *ExportRansomDefenseEventsListResponse, err error) {
    return c.ExportRansomDefenseEventsListWithContext(context.Background(), request)
}

// ExportRansomDefenseEventsList
// 导出防勒索事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseEventsListWithContext(ctx context.Context, request *ExportRansomDefenseEventsListRequest) (response *ExportRansomDefenseEventsListResponse, err error) {
    if request == nil {
        request = NewExportRansomDefenseEventsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRansomDefenseEventsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRansomDefenseEventsListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRansomDefenseMachineListRequest() (request *ExportRansomDefenseMachineListRequest) {
    request = &ExportRansomDefenseMachineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseMachineList")
    
    
    return
}

func NewExportRansomDefenseMachineListResponse() (response *ExportRansomDefenseMachineListResponse) {
    response = &ExportRansomDefenseMachineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRansomDefenseMachineList
// 导出备份详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseMachineList(request *ExportRansomDefenseMachineListRequest) (response *ExportRansomDefenseMachineListResponse, err error) {
    return c.ExportRansomDefenseMachineListWithContext(context.Background(), request)
}

// ExportRansomDefenseMachineList
// 导出备份详情列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseMachineListWithContext(ctx context.Context, request *ExportRansomDefenseMachineListRequest) (response *ExportRansomDefenseMachineListResponse, err error) {
    if request == nil {
        request = NewExportRansomDefenseMachineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRansomDefenseMachineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRansomDefenseMachineListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRansomDefenseStrategyListRequest() (request *ExportRansomDefenseStrategyListRequest) {
    request = &ExportRansomDefenseStrategyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseStrategyList")
    
    
    return
}

func NewExportRansomDefenseStrategyListResponse() (response *ExportRansomDefenseStrategyListResponse) {
    response = &ExportRansomDefenseStrategyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRansomDefenseStrategyList
// 导出防勒索策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseStrategyList(request *ExportRansomDefenseStrategyListRequest) (response *ExportRansomDefenseStrategyListResponse, err error) {
    return c.ExportRansomDefenseStrategyListWithContext(context.Background(), request)
}

// ExportRansomDefenseStrategyList
// 导出防勒索策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseStrategyListWithContext(ctx context.Context, request *ExportRansomDefenseStrategyListRequest) (response *ExportRansomDefenseStrategyListResponse, err error) {
    if request == nil {
        request = NewExportRansomDefenseStrategyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRansomDefenseStrategyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRansomDefenseStrategyListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRansomDefenseStrategyMachinesRequest() (request *ExportRansomDefenseStrategyMachinesRequest) {
    request = &ExportRansomDefenseStrategyMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRansomDefenseStrategyMachines")
    
    
    return
}

func NewExportRansomDefenseStrategyMachinesResponse() (response *ExportRansomDefenseStrategyMachinesResponse) {
    response = &ExportRansomDefenseStrategyMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRansomDefenseStrategyMachines
// 导出勒索防御策略绑定机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseStrategyMachines(request *ExportRansomDefenseStrategyMachinesRequest) (response *ExportRansomDefenseStrategyMachinesResponse, err error) {
    return c.ExportRansomDefenseStrategyMachinesWithContext(context.Background(), request)
}

// ExportRansomDefenseStrategyMachines
// 导出勒索防御策略绑定机器列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportRansomDefenseStrategyMachinesWithContext(ctx context.Context, request *ExportRansomDefenseStrategyMachinesRequest) (response *ExportRansomDefenseStrategyMachinesResponse, err error) {
    if request == nil {
        request = NewExportRansomDefenseStrategyMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRansomDefenseStrategyMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRansomDefenseStrategyMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewExportReverseShellEventsRequest() (request *ExportReverseShellEventsRequest) {
    request = &ExportReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportReverseShellEvents")
    
    
    return
}

func NewExportReverseShellEventsResponse() (response *ExportReverseShellEventsResponse) {
    response = &ExportReverseShellEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportReverseShellEvents
// 导出反弹Shell事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportReverseShellEvents(request *ExportReverseShellEventsRequest) (response *ExportReverseShellEventsResponse, err error) {
    return c.ExportReverseShellEventsWithContext(context.Background(), request)
}

// ExportReverseShellEvents
// 导出反弹Shell事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportReverseShellEventsWithContext(ctx context.Context, request *ExportReverseShellEventsRequest) (response *ExportReverseShellEventsResponse, err error) {
    if request == nil {
        request = NewExportReverseShellEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportReverseShellEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportReverseShellEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportRiskDnsEventListRequest() (request *ExportRiskDnsEventListRequest) {
    request = &ExportRiskDnsEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskDnsEventList")
    
    
    return
}

func NewExportRiskDnsEventListResponse() (response *ExportRiskDnsEventListResponse) {
    response = &ExportRiskDnsEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRiskDnsEventList
// 导出恶意请求事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportRiskDnsEventList(request *ExportRiskDnsEventListRequest) (response *ExportRiskDnsEventListResponse, err error) {
    return c.ExportRiskDnsEventListWithContext(context.Background(), request)
}

// ExportRiskDnsEventList
// 导出恶意请求事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportRiskDnsEventListWithContext(ctx context.Context, request *ExportRiskDnsEventListRequest) (response *ExportRiskDnsEventListResponse, err error) {
    if request == nil {
        request = NewExportRiskDnsEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRiskDnsEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRiskDnsEventListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRiskDnsPolicyListRequest() (request *ExportRiskDnsPolicyListRequest) {
    request = &ExportRiskDnsPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskDnsPolicyList")
    
    
    return
}

func NewExportRiskDnsPolicyListResponse() (response *ExportRiskDnsPolicyListResponse) {
    response = &ExportRiskDnsPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRiskDnsPolicyList
// 导出恶意请求策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ExportRiskDnsPolicyList(request *ExportRiskDnsPolicyListRequest) (response *ExportRiskDnsPolicyListResponse, err error) {
    return c.ExportRiskDnsPolicyListWithContext(context.Background(), request)
}

// ExportRiskDnsPolicyList
// 导出恶意请求策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ExportRiskDnsPolicyListWithContext(ctx context.Context, request *ExportRiskDnsPolicyListRequest) (response *ExportRiskDnsPolicyListResponse, err error) {
    if request == nil {
        request = NewExportRiskDnsPolicyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRiskDnsPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRiskDnsPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewExportRiskProcessEventsRequest() (request *ExportRiskProcessEventsRequest) {
    request = &ExportRiskProcessEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportRiskProcessEvents")
    
    
    return
}

func NewExportRiskProcessEventsResponse() (response *ExportRiskProcessEventsResponse) {
    response = &ExportRiskProcessEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportRiskProcessEvents
// 导出异常进程事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportRiskProcessEvents(request *ExportRiskProcessEventsRequest) (response *ExportRiskProcessEventsResponse, err error) {
    return c.ExportRiskProcessEventsWithContext(context.Background(), request)
}

// ExportRiskProcessEvents
// 导出异常进程事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportRiskProcessEventsWithContext(ctx context.Context, request *ExportRiskProcessEventsRequest) (response *ExportRiskProcessEventsResponse, err error) {
    if request == nil {
        request = NewExportRiskProcessEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRiskProcessEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRiskProcessEventsResponse()
    err = c.Send(request, response)
    return
}

func NewExportScanTaskDetailsRequest() (request *ExportScanTaskDetailsRequest) {
    request = &ExportScanTaskDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportScanTaskDetails")
    
    
    return
}

func NewExportScanTaskDetailsResponse() (response *ExportScanTaskDetailsResponse) {
    response = &ExportScanTaskDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportScanTaskDetails
// 根据任务id导出指定扫描任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportScanTaskDetails(request *ExportScanTaskDetailsRequest) (response *ExportScanTaskDetailsResponse, err error) {
    return c.ExportScanTaskDetailsWithContext(context.Background(), request)
}

// ExportScanTaskDetails
// 根据任务id导出指定扫描任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportScanTaskDetailsWithContext(ctx context.Context, request *ExportScanTaskDetailsRequest) (response *ExportScanTaskDetailsResponse, err error) {
    if request == nil {
        request = NewExportScanTaskDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportScanTaskDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportScanTaskDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewExportSecurityTrendsRequest() (request *ExportSecurityTrendsRequest) {
    request = &ExportSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportSecurityTrends")
    
    
    return
}

func NewExportSecurityTrendsResponse() (response *ExportSecurityTrendsResponse) {
    response = &ExportSecurityTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportSecurityTrends
// 导出风险趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ExportSecurityTrends(request *ExportSecurityTrendsRequest) (response *ExportSecurityTrendsResponse, err error) {
    return c.ExportSecurityTrendsWithContext(context.Background(), request)
}

// ExportSecurityTrends
// 导出风险趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ExportSecurityTrendsWithContext(ctx context.Context, request *ExportSecurityTrendsRequest) (response *ExportSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewExportSecurityTrendsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportSecurityTrends require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportSecurityTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewExportTasksRequest() (request *ExportTasksRequest) {
    request = &ExportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportTasks")
    
    
    return
}

func NewExportTasksResponse() (response *ExportTasksResponse) {
    response = &ExportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportTasks
// 用于异步导出数据量大的日志文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportTasks(request *ExportTasksRequest) (response *ExportTasksResponse, err error) {
    return c.ExportTasksWithContext(context.Background(), request)
}

// ExportTasks
// 用于异步导出数据量大的日志文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportTasksWithContext(ctx context.Context, request *ExportTasksRequest) (response *ExportTasksResponse, err error) {
    if request == nil {
        request = NewExportTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulDefenceEventRequest() (request *ExportVulDefenceEventRequest) {
    request = &ExportVulDefenceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefenceEvent")
    
    
    return
}

func NewExportVulDefenceEventResponse() (response *ExportVulDefenceEventResponse) {
    response = &ExportVulDefenceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulDefenceEvent
// 导出漏洞防御事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefenceEvent(request *ExportVulDefenceEventRequest) (response *ExportVulDefenceEventResponse, err error) {
    return c.ExportVulDefenceEventWithContext(context.Background(), request)
}

// ExportVulDefenceEvent
// 导出漏洞防御事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefenceEventWithContext(ctx context.Context, request *ExportVulDefenceEventRequest) (response *ExportVulDefenceEventResponse, err error) {
    if request == nil {
        request = NewExportVulDefenceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulDefenceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulDefenceEventResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulDefenceListRequest() (request *ExportVulDefenceListRequest) {
    request = &ExportVulDefenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefenceList")
    
    
    return
}

func NewExportVulDefenceListResponse() (response *ExportVulDefenceListResponse) {
    response = &ExportVulDefenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulDefenceList
// 导出漏洞防御列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefenceList(request *ExportVulDefenceListRequest) (response *ExportVulDefenceListResponse, err error) {
    return c.ExportVulDefenceListWithContext(context.Background(), request)
}

// ExportVulDefenceList
// 导出漏洞防御列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefenceListWithContext(ctx context.Context, request *ExportVulDefenceListRequest) (response *ExportVulDefenceListResponse, err error) {
    if request == nil {
        request = NewExportVulDefenceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulDefenceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulDefenceListResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulDefencePluginEventRequest() (request *ExportVulDefencePluginEventRequest) {
    request = &ExportVulDefencePluginEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDefencePluginEvent")
    
    
    return
}

func NewExportVulDefencePluginEventResponse() (response *ExportVulDefencePluginEventResponse) {
    response = &ExportVulDefencePluginEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulDefencePluginEvent
// 导出漏洞防御插件事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefencePluginEvent(request *ExportVulDefencePluginEventRequest) (response *ExportVulDefencePluginEventResponse, err error) {
    return c.ExportVulDefencePluginEventWithContext(context.Background(), request)
}

// ExportVulDefencePluginEvent
// 导出漏洞防御插件事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDefencePluginEventWithContext(ctx context.Context, request *ExportVulDefencePluginEventRequest) (response *ExportVulDefencePluginEventResponse, err error) {
    if request == nil {
        request = NewExportVulDefencePluginEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulDefencePluginEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulDefencePluginEventResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulDetectionExcelRequest() (request *ExportVulDetectionExcelRequest) {
    request = &ExportVulDetectionExcelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDetectionExcel")
    
    
    return
}

func NewExportVulDetectionExcelResponse() (response *ExportVulDetectionExcelResponse) {
    response = &ExportVulDetectionExcelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulDetectionExcel
// 导出本次漏洞检测Excel
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDetectionExcel(request *ExportVulDetectionExcelRequest) (response *ExportVulDetectionExcelResponse, err error) {
    return c.ExportVulDetectionExcelWithContext(context.Background(), request)
}

// ExportVulDetectionExcel
// 导出本次漏洞检测Excel
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDetectionExcelWithContext(ctx context.Context, request *ExportVulDetectionExcelRequest) (response *ExportVulDetectionExcelResponse, err error) {
    if request == nil {
        request = NewExportVulDetectionExcelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulDetectionExcel require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulDetectionExcelResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulDetectionReportRequest() (request *ExportVulDetectionReportRequest) {
    request = &ExportVulDetectionReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulDetectionReport")
    
    
    return
}

func NewExportVulDetectionReportResponse() (response *ExportVulDetectionReportResponse) {
    response = &ExportVulDetectionReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulDetectionReport
// 导出漏洞检测报告。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDetectionReport(request *ExportVulDetectionReportRequest) (response *ExportVulDetectionReportResponse, err error) {
    return c.ExportVulDetectionReportWithContext(context.Background(), request)
}

// ExportVulDetectionReport
// 导出漏洞检测报告。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDetectionReportWithContext(ctx context.Context, request *ExportVulDetectionReportRequest) (response *ExportVulDetectionReportResponse, err error) {
    if request == nil {
        request = NewExportVulDetectionReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulDetectionReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulDetectionReportResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulEffectHostListRequest() (request *ExportVulEffectHostListRequest) {
    request = &ExportVulEffectHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulEffectHostList")
    
    
    return
}

func NewExportVulEffectHostListResponse() (response *ExportVulEffectHostListResponse) {
    response = &ExportVulEffectHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulEffectHostList
// 导出漏洞影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ExportVulEffectHostList(request *ExportVulEffectHostListRequest) (response *ExportVulEffectHostListResponse, err error) {
    return c.ExportVulEffectHostListWithContext(context.Background(), request)
}

// ExportVulEffectHostList
// 导出漏洞影响主机列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ExportVulEffectHostListWithContext(ctx context.Context, request *ExportVulEffectHostListRequest) (response *ExportVulEffectHostListResponse, err error) {
    if request == nil {
        request = NewExportVulEffectHostListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulEffectHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulEffectHostListResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulInfoRequest() (request *ExportVulInfoRequest) {
    request = &ExportVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulInfo")
    
    
    return
}

func NewExportVulInfoResponse() (response *ExportVulInfoResponse) {
    response = &ExportVulInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulInfo
// 导出漏洞信息，包括影响主机列表，组件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ExportVulInfo(request *ExportVulInfoRequest) (response *ExportVulInfoResponse, err error) {
    return c.ExportVulInfoWithContext(context.Background(), request)
}

// ExportVulInfo
// 导出漏洞信息，包括影响主机列表，组件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ExportVulInfoWithContext(ctx context.Context, request *ExportVulInfoRequest) (response *ExportVulInfoResponse, err error) {
    if request == nil {
        request = NewExportVulInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewExportVulListRequest() (request *ExportVulListRequest) {
    request = &ExportVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportVulList")
    
    
    return
}

func NewExportVulListResponse() (response *ExportVulListResponse) {
    response = &ExportVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportVulList
// 漏洞管理-导出漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulList(request *ExportVulListRequest) (response *ExportVulListResponse, err error) {
    return c.ExportVulListWithContext(context.Background(), request)
}

// ExportVulList
// 漏洞管理-导出漏洞列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulListWithContext(ctx context.Context, request *ExportVulListRequest) (response *ExportVulListResponse, err error) {
    if request == nil {
        request = NewExportVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportVulListResponse()
    err = c.Send(request, response)
    return
}

func NewExportWebPageEventListRequest() (request *ExportWebPageEventListRequest) {
    request = &ExportWebPageEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ExportWebPageEventList")
    
    
    return
}

func NewExportWebPageEventListResponse() (response *ExportWebPageEventListResponse) {
    response = &ExportWebPageEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportWebPageEventList
// 导出篡改事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportWebPageEventList(request *ExportWebPageEventListRequest) (response *ExportWebPageEventListResponse, err error) {
    return c.ExportWebPageEventListWithContext(context.Background(), request)
}

// ExportWebPageEventList
// 导出篡改事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportWebPageEventListWithContext(ctx context.Context, request *ExportWebPageEventListRequest) (response *ExportWebPageEventListResponse, err error) {
    if request == nil {
        request = NewExportWebPageEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportWebPageEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportWebPageEventListResponse()
    err = c.Send(request, response)
    return
}

func NewFixBaselineDetectRequest() (request *FixBaselineDetectRequest) {
    request = &FixBaselineDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "FixBaselineDetect")
    
    
    return
}

func NewFixBaselineDetectResponse() (response *FixBaselineDetectResponse) {
    response = &FixBaselineDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FixBaselineDetect
// 修复基线检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) FixBaselineDetect(request *FixBaselineDetectRequest) (response *FixBaselineDetectResponse, err error) {
    return c.FixBaselineDetectWithContext(context.Background(), request)
}

// FixBaselineDetect
// 修复基线检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) FixBaselineDetectWithContext(ctx context.Context, request *FixBaselineDetectRequest) (response *FixBaselineDetectResponse, err error) {
    if request == nil {
        request = NewFixBaselineDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FixBaselineDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewFixBaselineDetectResponse()
    err = c.Send(request, response)
    return
}

func NewGetLocalStorageItemRequest() (request *GetLocalStorageItemRequest) {
    request = &GetLocalStorageItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "GetLocalStorageItem")
    
    
    return
}

func NewGetLocalStorageItemResponse() (response *GetLocalStorageItemResponse) {
    response = &GetLocalStorageItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetLocalStorageItem
// 获取本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetLocalStorageItem(request *GetLocalStorageItemRequest) (response *GetLocalStorageItemResponse, err error) {
    return c.GetLocalStorageItemWithContext(context.Background(), request)
}

// GetLocalStorageItem
// 获取本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetLocalStorageItemWithContext(ctx context.Context, request *GetLocalStorageItemRequest) (response *GetLocalStorageItemResponse, err error) {
    if request == nil {
        request = NewGetLocalStorageItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLocalStorageItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLocalStorageItemResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
    request = &IgnoreImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "IgnoreImpactedHosts")
    
    
    return
}

func NewIgnoreImpactedHostsResponse() (response *IgnoreImpactedHostsResponse) {
    response = &IgnoreImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IgnoreImpactedHosts
// 本接口 (IgnoreImpactedHosts) 用于忽略漏洞。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    return c.IgnoreImpactedHostsWithContext(context.Background(), request)
}

// IgnoreImpactedHosts
// 本接口 (IgnoreImpactedHosts) 用于忽略漏洞。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) IgnoreImpactedHostsWithContext(ctx context.Context, request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    if request == nil {
        request = NewIgnoreImpactedHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreImpactedHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewKeysLocalStorageRequest() (request *KeysLocalStorageRequest) {
    request = &KeysLocalStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "KeysLocalStorage")
    
    
    return
}

func NewKeysLocalStorageResponse() (response *KeysLocalStorageResponse) {
    response = &KeysLocalStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KeysLocalStorage
// 获取本地存储键值列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) KeysLocalStorage(request *KeysLocalStorageRequest) (response *KeysLocalStorageResponse, err error) {
    return c.KeysLocalStorageWithContext(context.Background(), request)
}

// KeysLocalStorage
// 获取本地存储键值列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) KeysLocalStorageWithContext(ctx context.Context, request *KeysLocalStorageRequest) (response *KeysLocalStorageResponse, err error) {
    if request == nil {
        request = NewKeysLocalStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KeysLocalStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewKeysLocalStorageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
    request = &ModifyAutoOpenProVersionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyAutoOpenProVersionConfig")
    
    
    return
}

func NewModifyAutoOpenProVersionConfigResponse() (response *ModifyAutoOpenProVersionConfigResponse) {
    response = &ModifyAutoOpenProVersionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoOpenProVersionConfig
//  用于设置新增主机自动开通专业防护配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    return c.ModifyAutoOpenProVersionConfigWithContext(context.Background(), request)
}

// ModifyAutoOpenProVersionConfig
//  用于设置新增主机自动开通专业防护配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAutoOpenProVersionConfigWithContext(ctx context.Context, request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoOpenProVersionConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoOpenProVersionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanModeRequest() (request *ModifyBanModeRequest) {
    request = &ModifyBanModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanMode")
    
    
    return
}

func NewModifyBanModeResponse() (response *ModifyBanModeResponse) {
    response = &ModifyBanModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBanMode
// 修改爆破阻断模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyBanMode(request *ModifyBanModeRequest) (response *ModifyBanModeResponse, err error) {
    return c.ModifyBanModeWithContext(context.Background(), request)
}

// ModifyBanMode
// 修改爆破阻断模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyBanModeWithContext(ctx context.Context, request *ModifyBanModeRequest) (response *ModifyBanModeResponse, err error) {
    if request == nil {
        request = NewModifyBanModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBanMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBanModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanStatusRequest() (request *ModifyBanStatusRequest) {
    request = &ModifyBanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanStatus")
    
    
    return
}

func NewModifyBanStatusResponse() (response *ModifyBanStatusResponse) {
    response = &ModifyBanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBanStatus
// 设置阻断开关状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBanStatus(request *ModifyBanStatusRequest) (response *ModifyBanStatusResponse, err error) {
    return c.ModifyBanStatusWithContext(context.Background(), request)
}

// ModifyBanStatus
// 设置阻断开关状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBanStatusWithContext(ctx context.Context, request *ModifyBanStatusRequest) (response *ModifyBanStatusResponse, err error) {
    if request == nil {
        request = NewModifyBanStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBanStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBanWhiteListRequest() (request *ModifyBanWhiteListRequest) {
    request = &ModifyBanWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBanWhiteList")
    
    
    return
}

func NewModifyBanWhiteListResponse() (response *ModifyBanWhiteListResponse) {
    response = &ModifyBanWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBanWhiteList
// 修改阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBanWhiteList(request *ModifyBanWhiteListRequest) (response *ModifyBanWhiteListResponse, err error) {
    return c.ModifyBanWhiteListWithContext(context.Background(), request)
}

// ModifyBanWhiteList
// 修改阻断白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBanWhiteListWithContext(ctx context.Context, request *ModifyBanWhiteListRequest) (response *ModifyBanWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyBanWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBanWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBanWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselinePolicyRequest() (request *ModifyBaselinePolicyRequest) {
    request = &ModifyBaselinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselinePolicy")
    
    
    return
}

func NewModifyBaselinePolicyResponse() (response *ModifyBaselinePolicyResponse) {
    response = &ModifyBaselinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBaselinePolicy
// 更改基线策略设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselinePolicy(request *ModifyBaselinePolicyRequest) (response *ModifyBaselinePolicyResponse, err error) {
    return c.ModifyBaselinePolicyWithContext(context.Background(), request)
}

// ModifyBaselinePolicy
// 更改基线策略设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselinePolicyWithContext(ctx context.Context, request *ModifyBaselinePolicyRequest) (response *ModifyBaselinePolicyResponse, err error) {
    if request == nil {
        request = NewModifyBaselinePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselinePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselinePolicyStateRequest() (request *ModifyBaselinePolicyStateRequest) {
    request = &ModifyBaselinePolicyStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselinePolicyState")
    
    
    return
}

func NewModifyBaselinePolicyStateResponse() (response *ModifyBaselinePolicyStateResponse) {
    response = &ModifyBaselinePolicyStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBaselinePolicyState
// 更改基线策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselinePolicyState(request *ModifyBaselinePolicyStateRequest) (response *ModifyBaselinePolicyStateResponse, err error) {
    return c.ModifyBaselinePolicyStateWithContext(context.Background(), request)
}

// ModifyBaselinePolicyState
// 更改基线策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselinePolicyStateWithContext(ctx context.Context, request *ModifyBaselinePolicyStateRequest) (response *ModifyBaselinePolicyStateResponse, err error) {
    if request == nil {
        request = NewModifyBaselinePolicyStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselinePolicyState require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselinePolicyStateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselineRuleRequest() (request *ModifyBaselineRuleRequest) {
    request = &ModifyBaselineRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineRule")
    
    
    return
}

func NewModifyBaselineRuleResponse() (response *ModifyBaselineRuleResponse) {
    response = &ModifyBaselineRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBaselineRule
// 更改基线检测规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineRule(request *ModifyBaselineRuleRequest) (response *ModifyBaselineRuleResponse, err error) {
    return c.ModifyBaselineRuleWithContext(context.Background(), request)
}

// ModifyBaselineRule
// 更改基线检测规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineRuleWithContext(ctx context.Context, request *ModifyBaselineRuleRequest) (response *ModifyBaselineRuleResponse, err error) {
    if request == nil {
        request = NewModifyBaselineRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselineRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselineRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselineRuleIgnoreRequest() (request *ModifyBaselineRuleIgnoreRequest) {
    request = &ModifyBaselineRuleIgnoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineRuleIgnore")
    
    
    return
}

func NewModifyBaselineRuleIgnoreResponse() (response *ModifyBaselineRuleIgnoreResponse) {
    response = &ModifyBaselineRuleIgnoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBaselineRuleIgnore
// 更改基线忽略规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineRuleIgnore(request *ModifyBaselineRuleIgnoreRequest) (response *ModifyBaselineRuleIgnoreResponse, err error) {
    return c.ModifyBaselineRuleIgnoreWithContext(context.Background(), request)
}

// ModifyBaselineRuleIgnore
// 更改基线忽略规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineRuleIgnoreWithContext(ctx context.Context, request *ModifyBaselineRuleIgnoreRequest) (response *ModifyBaselineRuleIgnoreResponse, err error) {
    if request == nil {
        request = NewModifyBaselineRuleIgnoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselineRuleIgnore require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselineRuleIgnoreResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselineWeakPasswordRequest() (request *ModifyBaselineWeakPasswordRequest) {
    request = &ModifyBaselineWeakPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBaselineWeakPassword")
    
    
    return
}

func NewModifyBaselineWeakPasswordResponse() (response *ModifyBaselineWeakPasswordResponse) {
    response = &ModifyBaselineWeakPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBaselineWeakPassword
// 更改或新增弱口令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineWeakPassword(request *ModifyBaselineWeakPasswordRequest) (response *ModifyBaselineWeakPasswordResponse, err error) {
    return c.ModifyBaselineWeakPasswordWithContext(context.Background(), request)
}

// ModifyBaselineWeakPassword
// 更改或新增弱口令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyBaselineWeakPasswordWithContext(ctx context.Context, request *ModifyBaselineWeakPasswordRequest) (response *ModifyBaselineWeakPasswordResponse, err error) {
    if request == nil {
        request = NewModifyBaselineWeakPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselineWeakPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselineWeakPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBashPolicyRequest() (request *ModifyBashPolicyRequest) {
    request = &ModifyBashPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBashPolicy")
    
    
    return
}

func NewModifyBashPolicyResponse() (response *ModifyBashPolicyResponse) {
    response = &ModifyBashPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBashPolicy
// 新增或修改高危命令策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBashPolicy(request *ModifyBashPolicyRequest) (response *ModifyBashPolicyResponse, err error) {
    return c.ModifyBashPolicyWithContext(context.Background(), request)
}

// ModifyBashPolicy
// 新增或修改高危命令策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBashPolicyWithContext(ctx context.Context, request *ModifyBashPolicyRequest) (response *ModifyBashPolicyResponse, err error) {
    if request == nil {
        request = NewModifyBashPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBashPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBashPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBashPolicyStatusRequest() (request *ModifyBashPolicyStatusRequest) {
    request = &ModifyBashPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBashPolicyStatus")
    
    
    return
}

func NewModifyBashPolicyStatusResponse() (response *ModifyBashPolicyStatusResponse) {
    response = &ModifyBashPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBashPolicyStatus
// 切换高危命令策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBashPolicyStatus(request *ModifyBashPolicyStatusRequest) (response *ModifyBashPolicyStatusResponse, err error) {
    return c.ModifyBashPolicyStatusWithContext(context.Background(), request)
}

// ModifyBashPolicyStatus
// 切换高危命令策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyBashPolicyStatusWithContext(ctx context.Context, request *ModifyBashPolicyStatusRequest) (response *ModifyBashPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyBashPolicyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBashPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBashPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBruteAttackRulesRequest() (request *ModifyBruteAttackRulesRequest) {
    request = &ModifyBruteAttackRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyBruteAttackRules")
    
    
    return
}

func NewModifyBruteAttackRulesResponse() (response *ModifyBruteAttackRulesResponse) {
    response = &ModifyBruteAttackRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBruteAttackRules
// 修改暴力破解规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyBruteAttackRules(request *ModifyBruteAttackRulesRequest) (response *ModifyBruteAttackRulesResponse, err error) {
    return c.ModifyBruteAttackRulesWithContext(context.Background(), request)
}

// ModifyBruteAttackRules
// 修改暴力破解规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyBruteAttackRulesWithContext(ctx context.Context, request *ModifyBruteAttackRulesRequest) (response *ModifyBruteAttackRulesResponse, err error) {
    if request == nil {
        request = NewModifyBruteAttackRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBruteAttackRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBruteAttackRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEventAttackStatusRequest() (request *ModifyEventAttackStatusRequest) {
    request = &ModifyEventAttackStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyEventAttackStatus")
    
    
    return
}

func NewModifyEventAttackStatusResponse() (response *ModifyEventAttackStatusResponse) {
    response = &ModifyEventAttackStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEventAttackStatus
// 修改网络攻击事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEventAttackStatus(request *ModifyEventAttackStatusRequest) (response *ModifyEventAttackStatusResponse, err error) {
    return c.ModifyEventAttackStatusWithContext(context.Background(), request)
}

// ModifyEventAttackStatus
// 修改网络攻击事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEventAttackStatusWithContext(ctx context.Context, request *ModifyEventAttackStatusRequest) (response *ModifyEventAttackStatusResponse, err error) {
    if request == nil {
        request = NewModifyEventAttackStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEventAttackStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEventAttackStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileTamperEventsRequest() (request *ModifyFileTamperEventsRequest) {
    request = &ModifyFileTamperEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperEvents")
    
    
    return
}

func NewModifyFileTamperEventsResponse() (response *ModifyFileTamperEventsResponse) {
    response = &ModifyFileTamperEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFileTamperEvents
// 核心文件事件更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyFileTamperEvents(request *ModifyFileTamperEventsRequest) (response *ModifyFileTamperEventsResponse, err error) {
    return c.ModifyFileTamperEventsWithContext(context.Background(), request)
}

// ModifyFileTamperEvents
// 核心文件事件更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyFileTamperEventsWithContext(ctx context.Context, request *ModifyFileTamperEventsRequest) (response *ModifyFileTamperEventsResponse, err error) {
    if request == nil {
        request = NewModifyFileTamperEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileTamperEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileTamperEventsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileTamperRuleRequest() (request *ModifyFileTamperRuleRequest) {
    request = &ModifyFileTamperRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperRule")
    
    
    return
}

func NewModifyFileTamperRuleResponse() (response *ModifyFileTamperRuleResponse) {
    response = &ModifyFileTamperRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFileTamperRule
// 编辑、新增核心文件监控规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyFileTamperRule(request *ModifyFileTamperRuleRequest) (response *ModifyFileTamperRuleResponse, err error) {
    return c.ModifyFileTamperRuleWithContext(context.Background(), request)
}

// ModifyFileTamperRule
// 编辑、新增核心文件监控规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyFileTamperRuleWithContext(ctx context.Context, request *ModifyFileTamperRuleRequest) (response *ModifyFileTamperRuleResponse, err error) {
    if request == nil {
        request = NewModifyFileTamperRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileTamperRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileTamperRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileTamperRuleStatusRequest() (request *ModifyFileTamperRuleStatusRequest) {
    request = &ModifyFileTamperRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyFileTamperRuleStatus")
    
    
    return
}

func NewModifyFileTamperRuleStatusResponse() (response *ModifyFileTamperRuleStatusResponse) {
    response = &ModifyFileTamperRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFileTamperRuleStatus
// 核心文件规则状态更新，支持批量删除 关闭
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFileTamperRuleStatus(request *ModifyFileTamperRuleStatusRequest) (response *ModifyFileTamperRuleStatusResponse, err error) {
    return c.ModifyFileTamperRuleStatusWithContext(context.Background(), request)
}

// ModifyFileTamperRuleStatus
// 核心文件规则状态更新，支持批量删除 关闭
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFileTamperRuleStatusWithContext(ctx context.Context, request *ModifyFileTamperRuleStatusRequest) (response *ModifyFileTamperRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyFileTamperRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileTamperRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileTamperRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJavaMemShellPluginSwitchRequest() (request *ModifyJavaMemShellPluginSwitchRequest) {
    request = &ModifyJavaMemShellPluginSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyJavaMemShellPluginSwitch")
    
    
    return
}

func NewModifyJavaMemShellPluginSwitchResponse() (response *ModifyJavaMemShellPluginSwitchResponse) {
    response = &ModifyJavaMemShellPluginSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyJavaMemShellPluginSwitch
// 开关java内存马插件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJavaMemShellPluginSwitch(request *ModifyJavaMemShellPluginSwitchRequest) (response *ModifyJavaMemShellPluginSwitchResponse, err error) {
    return c.ModifyJavaMemShellPluginSwitchWithContext(context.Background(), request)
}

// ModifyJavaMemShellPluginSwitch
// 开关java内存马插件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJavaMemShellPluginSwitchWithContext(ctx context.Context, request *ModifyJavaMemShellPluginSwitchRequest) (response *ModifyJavaMemShellPluginSwitchResponse, err error) {
    if request == nil {
        request = NewModifyJavaMemShellPluginSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJavaMemShellPluginSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJavaMemShellPluginSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJavaMemShellsStatusRequest() (request *ModifyJavaMemShellsStatusRequest) {
    request = &ModifyJavaMemShellsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyJavaMemShellsStatus")
    
    
    return
}

func NewModifyJavaMemShellsStatusResponse() (response *ModifyJavaMemShellsStatusResponse) {
    response = &ModifyJavaMemShellsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyJavaMemShellsStatus
// 修改java内存马事件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJavaMemShellsStatus(request *ModifyJavaMemShellsStatusRequest) (response *ModifyJavaMemShellsStatusResponse, err error) {
    return c.ModifyJavaMemShellsStatusWithContext(context.Background(), request)
}

// ModifyJavaMemShellsStatus
// 修改java内存马事件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyJavaMemShellsStatusWithContext(ctx context.Context, request *ModifyJavaMemShellsStatusRequest) (response *ModifyJavaMemShellsStatusResponse, err error) {
    if request == nil {
        request = NewModifyJavaMemShellsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJavaMemShellsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJavaMemShellsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLicenseBindsRequest() (request *ModifyLicenseBindsRequest) {
    request = &ModifyLicenseBindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseBinds")
    
    
    return
}

func NewModifyLicenseBindsResponse() (response *ModifyLicenseBindsResponse) {
    response = &ModifyLicenseBindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLicenseBinds
// 设置中心-授权管理 对某个授权批量绑定机器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseBinds(request *ModifyLicenseBindsRequest) (response *ModifyLicenseBindsResponse, err error) {
    return c.ModifyLicenseBindsWithContext(context.Background(), request)
}

// ModifyLicenseBinds
// 设置中心-授权管理 对某个授权批量绑定机器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseBindsWithContext(ctx context.Context, request *ModifyLicenseBindsRequest) (response *ModifyLicenseBindsResponse, err error) {
    if request == nil {
        request = NewModifyLicenseBindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLicenseBinds require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLicenseBindsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLicenseOrderRequest() (request *ModifyLicenseOrderRequest) {
    request = &ModifyLicenseOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseOrder")
    
    
    return
}

func NewModifyLicenseOrderResponse() (response *ModifyLicenseOrderResponse) {
    response = &ModifyLicenseOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLicenseOrder
// 编辑《主机安全-按量计费》授权订单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseOrder(request *ModifyLicenseOrderRequest) (response *ModifyLicenseOrderResponse, err error) {
    return c.ModifyLicenseOrderWithContext(context.Background(), request)
}

// ModifyLicenseOrder
// 编辑《主机安全-按量计费》授权订单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseOrderWithContext(ctx context.Context, request *ModifyLicenseOrderRequest) (response *ModifyLicenseOrderResponse, err error) {
    if request == nil {
        request = NewModifyLicenseOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLicenseOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLicenseOrderResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLicenseUnBindsRequest() (request *ModifyLicenseUnBindsRequest) {
    request = &ModifyLicenseUnBindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLicenseUnBinds")
    
    
    return
}

func NewModifyLicenseUnBindsResponse() (response *ModifyLicenseUnBindsResponse) {
    response = &ModifyLicenseUnBindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLicenseUnBinds
// 设置中心-授权管理 对某个授权批量解绑机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseUnBinds(request *ModifyLicenseUnBindsRequest) (response *ModifyLicenseUnBindsResponse, err error) {
    return c.ModifyLicenseUnBindsWithContext(context.Background(), request)
}

// ModifyLicenseUnBinds
// 设置中心-授权管理 对某个授权批量解绑机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseUnBindsWithContext(ctx context.Context, request *ModifyLicenseUnBindsRequest) (response *ModifyLicenseUnBindsResponse, err error) {
    if request == nil {
        request = NewModifyLicenseUnBindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLicenseUnBinds require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLicenseUnBindsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogKafkaAccessRequest() (request *ModifyLogKafkaAccessRequest) {
    request = &ModifyLogKafkaAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaAccess")
    
    
    return
}

func NewModifyLogKafkaAccessResponse() (response *ModifyLogKafkaAccessResponse) {
    response = &ModifyLogKafkaAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLogKafkaAccess
// 新增或修改日志投递kafka接入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaAccess(request *ModifyLogKafkaAccessRequest) (response *ModifyLogKafkaAccessResponse, err error) {
    return c.ModifyLogKafkaAccessWithContext(context.Background(), request)
}

// ModifyLogKafkaAccess
// 新增或修改日志投递kafka接入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaAccessWithContext(ctx context.Context, request *ModifyLogKafkaAccessRequest) (response *ModifyLogKafkaAccessResponse, err error) {
    if request == nil {
        request = NewModifyLogKafkaAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogKafkaAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogKafkaAccessResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogKafkaDeliverTypeRequest() (request *ModifyLogKafkaDeliverTypeRequest) {
    request = &ModifyLogKafkaDeliverTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaDeliverType")
    
    
    return
}

func NewModifyLogKafkaDeliverTypeResponse() (response *ModifyLogKafkaDeliverTypeResponse) {
    response = &ModifyLogKafkaDeliverTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLogKafkaDeliverType
// 修改指定日志类别投递配置、开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaDeliverType(request *ModifyLogKafkaDeliverTypeRequest) (response *ModifyLogKafkaDeliverTypeResponse, err error) {
    return c.ModifyLogKafkaDeliverTypeWithContext(context.Background(), request)
}

// ModifyLogKafkaDeliverType
// 修改指定日志类别投递配置、开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaDeliverTypeWithContext(ctx context.Context, request *ModifyLogKafkaDeliverTypeRequest) (response *ModifyLogKafkaDeliverTypeResponse, err error) {
    if request == nil {
        request = NewModifyLogKafkaDeliverTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogKafkaDeliverType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogKafkaDeliverTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogKafkaStateRequest() (request *ModifyLogKafkaStateRequest) {
    request = &ModifyLogKafkaStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogKafkaState")
    
    
    return
}

func NewModifyLogKafkaStateResponse() (response *ModifyLogKafkaStateResponse) {
    response = &ModifyLogKafkaStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLogKafkaState
// 修改日志投递状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaState(request *ModifyLogKafkaStateRequest) (response *ModifyLogKafkaStateResponse, err error) {
    return c.ModifyLogKafkaStateWithContext(context.Background(), request)
}

// ModifyLogKafkaState
// 修改日志投递状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLogKafkaStateWithContext(ctx context.Context, request *ModifyLogKafkaStateRequest) (response *ModifyLogKafkaStateResponse, err error) {
    if request == nil {
        request = NewModifyLogKafkaStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogKafkaState require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogKafkaStateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogStorageConfigRequest() (request *ModifyLogStorageConfigRequest) {
    request = &ModifyLogStorageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLogStorageConfig")
    
    
    return
}

func NewModifyLogStorageConfigResponse() (response *ModifyLogStorageConfigResponse) {
    response = &ModifyLogStorageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLogStorageConfig
// 修改日志存储配置
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLogStorageConfig(request *ModifyLogStorageConfigRequest) (response *ModifyLogStorageConfigResponse, err error) {
    return c.ModifyLogStorageConfigWithContext(context.Background(), request)
}

// ModifyLogStorageConfig
// 修改日志存储配置
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLogStorageConfigWithContext(ctx context.Context, request *ModifyLogStorageConfigRequest) (response *ModifyLogStorageConfigResponse, err error) {
    if request == nil {
        request = NewModifyLogStorageConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogStorageConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogStorageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteInfoRequest() (request *ModifyLoginWhiteInfoRequest) {
    request = &ModifyLoginWhiteInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLoginWhiteInfo")
    
    
    return
}

func NewModifyLoginWhiteInfoResponse() (response *ModifyLoginWhiteInfoResponse) {
    response = &ModifyLoginWhiteInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoginWhiteInfo
// 更新登录审计白名单信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLoginWhiteInfo(request *ModifyLoginWhiteInfoRequest) (response *ModifyLoginWhiteInfoResponse, err error) {
    return c.ModifyLoginWhiteInfoWithContext(context.Background(), request)
}

// ModifyLoginWhiteInfo
// 更新登录审计白名单信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLoginWhiteInfoWithContext(ctx context.Context, request *ModifyLoginWhiteInfoRequest) (response *ModifyLoginWhiteInfoResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoginWhiteInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoginWhiteInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteRecordRequest() (request *ModifyLoginWhiteRecordRequest) {
    request = &ModifyLoginWhiteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLoginWhiteRecord")
    
    
    return
}

func NewModifyLoginWhiteRecordResponse() (response *ModifyLoginWhiteRecordResponse) {
    response = &ModifyLoginWhiteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoginWhiteRecord
// 更新合并后登录审计白名单信息（服务器列表数目应小于1000）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLoginWhiteRecord(request *ModifyLoginWhiteRecordRequest) (response *ModifyLoginWhiteRecordResponse, err error) {
    return c.ModifyLoginWhiteRecordWithContext(context.Background(), request)
}

// ModifyLoginWhiteRecord
// 更新合并后登录审计白名单信息（服务器列表数目应小于1000）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_RULEHOSTDUPLICATEERR = "InvalidParameter.RuleHostDuplicateErr"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLoginWhiteRecordWithContext(ctx context.Context, request *ModifyLoginWhiteRecordRequest) (response *ModifyLoginWhiteRecordResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoginWhiteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoginWhiteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachineAutoClearConfigRequest() (request *ModifyMachineAutoClearConfigRequest) {
    request = &ModifyMachineAutoClearConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyMachineAutoClearConfig")
    
    
    return
}

func NewModifyMachineAutoClearConfigResponse() (response *ModifyMachineAutoClearConfigResponse) {
    response = &ModifyMachineAutoClearConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMachineAutoClearConfig
// 修改机器清理配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMachineAutoClearConfig(request *ModifyMachineAutoClearConfigRequest) (response *ModifyMachineAutoClearConfigResponse, err error) {
    return c.ModifyMachineAutoClearConfigWithContext(context.Background(), request)
}

// ModifyMachineAutoClearConfig
// 修改机器清理配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMachineAutoClearConfigWithContext(ctx context.Context, request *ModifyMachineAutoClearConfigRequest) (response *ModifyMachineAutoClearConfigResponse, err error) {
    if request == nil {
        request = NewModifyMachineAutoClearConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineAutoClearConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineAutoClearConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachineRemarkRequest() (request *ModifyMachineRemarkRequest) {
    request = &ModifyMachineRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyMachineRemark")
    
    
    return
}

func NewModifyMachineRemarkResponse() (response *ModifyMachineRemarkResponse) {
    response = &ModifyMachineRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMachineRemark
// 修改主机备注信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyMachineRemark(request *ModifyMachineRemarkRequest) (response *ModifyMachineRemarkResponse, err error) {
    return c.ModifyMachineRemarkWithContext(context.Background(), request)
}

// ModifyMachineRemark
// 修改主机备注信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyMachineRemarkWithContext(ctx context.Context, request *ModifyMachineRemarkRequest) (response *ModifyMachineRemarkResponse, err error) {
    if request == nil {
        request = NewModifyMachineRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaliciousRequestWhiteListRequest() (request *ModifyMaliciousRequestWhiteListRequest) {
    request = &ModifyMaliciousRequestWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyMaliciousRequestWhiteList")
    
    
    return
}

func NewModifyMaliciousRequestWhiteListResponse() (response *ModifyMaliciousRequestWhiteListResponse) {
    response = &ModifyMaliciousRequestWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMaliciousRequestWhiteList
// 更新恶意请求白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMaliciousRequestWhiteList(request *ModifyMaliciousRequestWhiteListRequest) (response *ModifyMaliciousRequestWhiteListResponse, err error) {
    return c.ModifyMaliciousRequestWhiteListWithContext(context.Background(), request)
}

// ModifyMaliciousRequestWhiteList
// 更新恶意请求白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMaliciousRequestWhiteListWithContext(ctx context.Context, request *ModifyMaliciousRequestWhiteListRequest) (response *ModifyMaliciousRequestWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyMaliciousRequestWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaliciousRequestWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaliciousRequestWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMalwareTimingScanSettingsRequest() (request *ModifyMalwareTimingScanSettingsRequest) {
    request = &ModifyMalwareTimingScanSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyMalwareTimingScanSettings")
    
    
    return
}

func NewModifyMalwareTimingScanSettingsResponse() (response *ModifyMalwareTimingScanSettingsResponse) {
    response = &ModifyMalwareTimingScanSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMalwareTimingScanSettings
// 定时扫描设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMalwareTimingScanSettings(request *ModifyMalwareTimingScanSettingsRequest) (response *ModifyMalwareTimingScanSettingsResponse, err error) {
    return c.ModifyMalwareTimingScanSettingsWithContext(context.Background(), request)
}

// ModifyMalwareTimingScanSettings
// 定时扫描设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMalwareTimingScanSettingsWithContext(ctx context.Context, request *ModifyMalwareTimingScanSettingsRequest) (response *ModifyMalwareTimingScanSettingsResponse, err error) {
    if request == nil {
        request = NewModifyMalwareTimingScanSettingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMalwareTimingScanSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMalwareTimingScanSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMalwareWhiteListRequest() (request *ModifyMalwareWhiteListRequest) {
    request = &ModifyMalwareWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyMalwareWhiteList")
    
    
    return
}

func NewModifyMalwareWhiteListResponse() (response *ModifyMalwareWhiteListResponse) {
    response = &ModifyMalwareWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMalwareWhiteList
// 编辑木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) ModifyMalwareWhiteList(request *ModifyMalwareWhiteListRequest) (response *ModifyMalwareWhiteListResponse, err error) {
    return c.ModifyMalwareWhiteListWithContext(context.Background(), request)
}

// ModifyMalwareWhiteList
// 编辑木马白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) ModifyMalwareWhiteListWithContext(ctx context.Context, request *ModifyMalwareWhiteListRequest) (response *ModifyMalwareWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyMalwareWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMalwareWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMalwareWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetAttackSettingRequest() (request *ModifyNetAttackSettingRequest) {
    request = &ModifyNetAttackSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyNetAttackSetting")
    
    
    return
}

func NewModifyNetAttackSettingResponse() (response *ModifyNetAttackSettingResponse) {
    response = &ModifyNetAttackSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetAttackSetting
// 修改网络攻击设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyNetAttackSetting(request *ModifyNetAttackSettingRequest) (response *ModifyNetAttackSettingResponse, err error) {
    return c.ModifyNetAttackSettingWithContext(context.Background(), request)
}

// ModifyNetAttackSetting
// 修改网络攻击设置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyNetAttackSettingWithContext(ctx context.Context, request *ModifyNetAttackSettingRequest) (response *ModifyNetAttackSettingResponse, err error) {
    if request == nil {
        request = NewModifyNetAttackSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetAttackSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetAttackSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetAttackWhiteListRequest() (request *ModifyNetAttackWhiteListRequest) {
    request = &ModifyNetAttackWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyNetAttackWhiteList")
    
    
    return
}

func NewModifyNetAttackWhiteListResponse() (response *ModifyNetAttackWhiteListResponse) {
    response = &ModifyNetAttackWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetAttackWhiteList
// 编辑网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) ModifyNetAttackWhiteList(request *ModifyNetAttackWhiteListRequest) (response *ModifyNetAttackWhiteListResponse, err error) {
    return c.ModifyNetAttackWhiteListWithContext(context.Background(), request)
}

// ModifyNetAttackWhiteList
// 编辑网络攻击白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) ModifyNetAttackWhiteListWithContext(ctx context.Context, request *ModifyNetAttackWhiteListRequest) (response *ModifyNetAttackWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyNetAttackWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetAttackWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetAttackWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOrderAttributeRequest() (request *ModifyOrderAttributeRequest) {
    request = &ModifyOrderAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyOrderAttribute")
    
    
    return
}

func NewModifyOrderAttributeResponse() (response *ModifyOrderAttributeResponse) {
    response = &ModifyOrderAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyOrderAttribute
// 对订单属性编辑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrderAttribute(request *ModifyOrderAttributeRequest) (response *ModifyOrderAttributeResponse, err error) {
    return c.ModifyOrderAttributeWithContext(context.Background(), request)
}

// ModifyOrderAttribute
// 对订单属性编辑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrderAttributeWithContext(ctx context.Context, request *ModifyOrderAttributeRequest) (response *ModifyOrderAttributeResponse, err error) {
    if request == nil {
        request = NewModifyOrderAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOrderAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOrderAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRansomDefenseEventsStatusRequest() (request *ModifyRansomDefenseEventsStatusRequest) {
    request = &ModifyRansomDefenseEventsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyRansomDefenseEventsStatus")
    
    
    return
}

func NewModifyRansomDefenseEventsStatusResponse() (response *ModifyRansomDefenseEventsStatusResponse) {
    response = &ModifyRansomDefenseEventsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRansomDefenseEventsStatus
// 修改防勒索事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRansomDefenseEventsStatus(request *ModifyRansomDefenseEventsStatusRequest) (response *ModifyRansomDefenseEventsStatusResponse, err error) {
    return c.ModifyRansomDefenseEventsStatusWithContext(context.Background(), request)
}

// ModifyRansomDefenseEventsStatus
// 修改防勒索事件状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRansomDefenseEventsStatusWithContext(ctx context.Context, request *ModifyRansomDefenseEventsStatusRequest) (response *ModifyRansomDefenseEventsStatusResponse, err error) {
    if request == nil {
        request = NewModifyRansomDefenseEventsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRansomDefenseEventsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRansomDefenseEventsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRansomDefenseStrategyStatusRequest() (request *ModifyRansomDefenseStrategyStatusRequest) {
    request = &ModifyRansomDefenseStrategyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyRansomDefenseStrategyStatus")
    
    
    return
}

func NewModifyRansomDefenseStrategyStatusResponse() (response *ModifyRansomDefenseStrategyStatusResponse) {
    response = &ModifyRansomDefenseStrategyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRansomDefenseStrategyStatus
// 批量修改防勒索策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRansomDefenseStrategyStatus(request *ModifyRansomDefenseStrategyStatusRequest) (response *ModifyRansomDefenseStrategyStatusResponse, err error) {
    return c.ModifyRansomDefenseStrategyStatusWithContext(context.Background(), request)
}

// ModifyRansomDefenseStrategyStatus
// 批量修改防勒索策略状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRansomDefenseStrategyStatusWithContext(ctx context.Context, request *ModifyRansomDefenseStrategyStatusRequest) (response *ModifyRansomDefenseStrategyStatusResponse, err error) {
    if request == nil {
        request = NewModifyRansomDefenseStrategyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRansomDefenseStrategyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRansomDefenseStrategyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskDnsPolicyRequest() (request *ModifyRiskDnsPolicyRequest) {
    request = &ModifyRiskDnsPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskDnsPolicy")
    
    
    return
}

func NewModifyRiskDnsPolicyResponse() (response *ModifyRiskDnsPolicyResponse) {
    response = &ModifyRiskDnsPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRiskDnsPolicy
// 更改恶意请求策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRiskDnsPolicy(request *ModifyRiskDnsPolicyRequest) (response *ModifyRiskDnsPolicyResponse, err error) {
    return c.ModifyRiskDnsPolicyWithContext(context.Background(), request)
}

// ModifyRiskDnsPolicy
// 更改恶意请求策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRiskDnsPolicyWithContext(ctx context.Context, request *ModifyRiskDnsPolicyRequest) (response *ModifyRiskDnsPolicyResponse, err error) {
    if request == nil {
        request = NewModifyRiskDnsPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskDnsPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskDnsPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskDnsPolicyStatusRequest() (request *ModifyRiskDnsPolicyStatusRequest) {
    request = &ModifyRiskDnsPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskDnsPolicyStatus")
    
    
    return
}

func NewModifyRiskDnsPolicyStatusResponse() (response *ModifyRiskDnsPolicyStatusResponse) {
    response = &ModifyRiskDnsPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRiskDnsPolicyStatus
// 更改恶意请求策略状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyRiskDnsPolicyStatus(request *ModifyRiskDnsPolicyStatusRequest) (response *ModifyRiskDnsPolicyStatusResponse, err error) {
    return c.ModifyRiskDnsPolicyStatusWithContext(context.Background(), request)
}

// ModifyRiskDnsPolicyStatus
// 更改恶意请求策略状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyRiskDnsPolicyStatusWithContext(ctx context.Context, request *ModifyRiskDnsPolicyStatusRequest) (response *ModifyRiskDnsPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskDnsPolicyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskDnsPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskDnsPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRiskEventsStatusRequest() (request *ModifyRiskEventsStatusRequest) {
    request = &ModifyRiskEventsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyRiskEventsStatus")
    
    
    return
}

func NewModifyRiskEventsStatusResponse() (response *ModifyRiskEventsStatusResponse) {
    response = &ModifyRiskEventsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRiskEventsStatus
// 入侵检测所有事件的状态，包括：文件查杀，异常登录，密码破解，高危命令，反弹shell，本地提取
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRiskEventsStatus(request *ModifyRiskEventsStatusRequest) (response *ModifyRiskEventsStatusResponse, err error) {
    return c.ModifyRiskEventsStatusWithContext(context.Background(), request)
}

// ModifyRiskEventsStatus
// 入侵检测所有事件的状态，包括：文件查杀，异常登录，密码破解，高危命令，反弹shell，本地提取
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRiskEventsStatusWithContext(ctx context.Context, request *ModifyRiskEventsStatusRequest) (response *ModifyRiskEventsStatusResponse, err error) {
    if request == nil {
        request = NewModifyRiskEventsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRiskEventsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRiskEventsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsersConfigRequest() (request *ModifyUsersConfigRequest) {
    request = &ModifyUsersConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyUsersConfig")
    
    
    return
}

func NewModifyUsersConfigResponse() (response *ModifyUsersConfigResponse) {
    response = &ModifyUsersConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUsersConfig
// 用于创建/修改用户自定义配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUsersConfig(request *ModifyUsersConfigRequest) (response *ModifyUsersConfigResponse, err error) {
    return c.ModifyUsersConfigWithContext(context.Background(), request)
}

// ModifyUsersConfig
// 用于创建/修改用户自定义配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUsersConfigWithContext(ctx context.Context, request *ModifyUsersConfigRequest) (response *ModifyUsersConfigResponse, err error) {
    if request == nil {
        request = NewModifyUsersConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUsersConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUsersConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVulDefenceEventStatusRequest() (request *ModifyVulDefenceEventStatusRequest) {
    request = &ModifyVulDefenceEventStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyVulDefenceEventStatus")
    
    
    return
}

func NewModifyVulDefenceEventStatusResponse() (response *ModifyVulDefenceEventStatusResponse) {
    response = &ModifyVulDefenceEventStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVulDefenceEventStatus
// 修改漏洞防御事件状态（修复漏洞通过其他接口实现）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyVulDefenceEventStatus(request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
    return c.ModifyVulDefenceEventStatusWithContext(context.Background(), request)
}

// ModifyVulDefenceEventStatus
// 修改漏洞防御事件状态（修复漏洞通过其他接口实现）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyVulDefenceEventStatusWithContext(ctx context.Context, request *ModifyVulDefenceEventStatusRequest) (response *ModifyVulDefenceEventStatusResponse, err error) {
    if request == nil {
        request = NewModifyVulDefenceEventStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVulDefenceEventStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVulDefenceEventStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVulDefenceSettingRequest() (request *ModifyVulDefenceSettingRequest) {
    request = &ModifyVulDefenceSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyVulDefenceSetting")
    
    
    return
}

func NewModifyVulDefenceSettingResponse() (response *ModifyVulDefenceSettingResponse) {
    response = &ModifyVulDefenceSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVulDefenceSetting
// 修改漏洞防御插件设置
//
// 1）新增主机自动加入，scope为1，quuids为空
//
// 2）全量旗舰版不自动加入，scope为0，quuids为当前quuid列表，
//
// 3）给定quuid列表，scope为0，quuids为用户选择quuid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyVulDefenceSetting(request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
    return c.ModifyVulDefenceSettingWithContext(context.Background(), request)
}

// ModifyVulDefenceSetting
// 修改漏洞防御插件设置
//
// 1）新增主机自动加入，scope为1，quuids为空
//
// 2）全量旗舰版不自动加入，scope为0，quuids为当前quuid列表，
//
// 3）给定quuid列表，scope为0，quuids为用户选择quuid
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyVulDefenceSettingWithContext(ctx context.Context, request *ModifyVulDefenceSettingRequest) (response *ModifyVulDefenceSettingResponse, err error) {
    if request == nil {
        request = NewModifyVulDefenceSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVulDefenceSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVulDefenceSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWarningHostConfigRequest() (request *ModifyWarningHostConfigRequest) {
    request = &ModifyWarningHostConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWarningHostConfig")
    
    
    return
}

func NewModifyWarningHostConfigResponse() (response *ModifyWarningHostConfigResponse) {
    response = &ModifyWarningHostConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWarningHostConfig
// 修改告警机器范围配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWarningHostConfig(request *ModifyWarningHostConfigRequest) (response *ModifyWarningHostConfigResponse, err error) {
    return c.ModifyWarningHostConfigWithContext(context.Background(), request)
}

// ModifyWarningHostConfig
// 修改告警机器范围配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWarningHostConfigWithContext(ctx context.Context, request *ModifyWarningHostConfigRequest) (response *ModifyWarningHostConfigResponse, err error) {
    if request == nil {
        request = NewModifyWarningHostConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWarningHostConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWarningHostConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWarningSettingRequest() (request *ModifyWarningSettingRequest) {
    request = &ModifyWarningSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWarningSetting")
    
    
    return
}

func NewModifyWarningSettingResponse() (response *ModifyWarningSettingResponse) {
    response = &ModifyWarningSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWarningSetting
// 修改告警设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWarningSetting(request *ModifyWarningSettingRequest) (response *ModifyWarningSettingResponse, err error) {
    return c.ModifyWarningSettingWithContext(context.Background(), request)
}

// ModifyWarningSetting
// 修改告警设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWarningSettingWithContext(ctx context.Context, request *ModifyWarningSettingRequest) (response *ModifyWarningSettingResponse, err error) {
    if request == nil {
        request = NewModifyWarningSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWarningSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWarningSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebHookPolicyRequest() (request *ModifyWebHookPolicyRequest) {
    request = &ModifyWebHookPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookPolicy")
    
    
    return
}

func NewModifyWebHookPolicyResponse() (response *ModifyWebHookPolicyResponse) {
    response = &ModifyWebHookPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebHookPolicy
// 新增或修改告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebHookPolicy(request *ModifyWebHookPolicyRequest) (response *ModifyWebHookPolicyResponse, err error) {
    return c.ModifyWebHookPolicyWithContext(context.Background(), request)
}

// ModifyWebHookPolicy
// 新增或修改告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebHookPolicyWithContext(ctx context.Context, request *ModifyWebHookPolicyRequest) (response *ModifyWebHookPolicyResponse, err error) {
    if request == nil {
        request = NewModifyWebHookPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebHookPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebHookPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebHookPolicyStatusRequest() (request *ModifyWebHookPolicyStatusRequest) {
    request = &ModifyWebHookPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookPolicyStatus")
    
    
    return
}

func NewModifyWebHookPolicyStatusResponse() (response *ModifyWebHookPolicyStatusResponse) {
    response = &ModifyWebHookPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebHookPolicyStatus
// 修改告警策略开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebHookPolicyStatus(request *ModifyWebHookPolicyStatusRequest) (response *ModifyWebHookPolicyStatusResponse, err error) {
    return c.ModifyWebHookPolicyStatusWithContext(context.Background(), request)
}

// ModifyWebHookPolicyStatus
// 修改告警策略开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebHookPolicyStatusWithContext(ctx context.Context, request *ModifyWebHookPolicyStatusRequest) (response *ModifyWebHookPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyWebHookPolicyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebHookPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebHookPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebHookReceiverRequest() (request *ModifyWebHookReceiverRequest) {
    request = &ModifyWebHookReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookReceiver")
    
    
    return
}

func NewModifyWebHookReceiverResponse() (response *ModifyWebHookReceiverResponse) {
    response = &ModifyWebHookReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebHookReceiver
// 新增或更新告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWebHookReceiver(request *ModifyWebHookReceiverRequest) (response *ModifyWebHookReceiverResponse, err error) {
    return c.ModifyWebHookReceiverWithContext(context.Background(), request)
}

// ModifyWebHookReceiver
// 新增或更新告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWebHookReceiverWithContext(ctx context.Context, request *ModifyWebHookReceiverRequest) (response *ModifyWebHookReceiverResponse, err error) {
    if request == nil {
        request = NewModifyWebHookReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebHookReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebHookReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebHookRuleRequest() (request *ModifyWebHookRuleRequest) {
    request = &ModifyWebHookRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookRule")
    
    
    return
}

func NewModifyWebHookRuleResponse() (response *ModifyWebHookRuleResponse) {
    response = &ModifyWebHookRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebHookRule
// 新增或修改企微机器人规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebHookRule(request *ModifyWebHookRuleRequest) (response *ModifyWebHookRuleResponse, err error) {
    return c.ModifyWebHookRuleWithContext(context.Background(), request)
}

// ModifyWebHookRule
// 新增或修改企微机器人规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebHookRuleWithContext(ctx context.Context, request *ModifyWebHookRuleRequest) (response *ModifyWebHookRuleResponse, err error) {
    if request == nil {
        request = NewModifyWebHookRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebHookRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebHookRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebHookRuleStatusRequest() (request *ModifyWebHookRuleStatusRequest) {
    request = &ModifyWebHookRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebHookRuleStatus")
    
    
    return
}

func NewModifyWebHookRuleStatusResponse() (response *ModifyWebHookRuleStatusResponse) {
    response = &ModifyWebHookRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebHookRuleStatus
// 修改企微机器人规则状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebHookRuleStatus(request *ModifyWebHookRuleStatusRequest) (response *ModifyWebHookRuleStatusResponse, err error) {
    return c.ModifyWebHookRuleStatusWithContext(context.Background(), request)
}

// ModifyWebHookRuleStatus
// 修改企微机器人规则状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebHookRuleStatusWithContext(ctx context.Context, request *ModifyWebHookRuleStatusRequest) (response *ModifyWebHookRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyWebHookRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebHookRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebHookRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebPageProtectDirRequest() (request *ModifyWebPageProtectDirRequest) {
    request = &ModifyWebPageProtectDirRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectDir")
    
    
    return
}

func NewModifyWebPageProtectDirResponse() (response *ModifyWebPageProtectDirResponse) {
    response = &ModifyWebPageProtectDirResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebPageProtectDir
// 创建/修改网站防护目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_LICENSEEXCEEDED = "FailedOperation.LicenseExceeded"
//  FAILEDOPERATION_PROTECTSTARTFAIL = "FailedOperation.ProtectStartFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyWebPageProtectDir(request *ModifyWebPageProtectDirRequest) (response *ModifyWebPageProtectDirResponse, err error) {
    return c.ModifyWebPageProtectDirWithContext(context.Background(), request)
}

// ModifyWebPageProtectDir
// 创建/修改网站防护目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_LICENSEEXCEEDED = "FailedOperation.LicenseExceeded"
//  FAILEDOPERATION_PROTECTSTARTFAIL = "FailedOperation.ProtectStartFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyWebPageProtectDirWithContext(ctx context.Context, request *ModifyWebPageProtectDirRequest) (response *ModifyWebPageProtectDirResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectDirRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebPageProtectDir require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebPageProtectDirResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebPageProtectSettingRequest() (request *ModifyWebPageProtectSettingRequest) {
    request = &ModifyWebPageProtectSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectSetting")
    
    
    return
}

func NewModifyWebPageProtectSettingResponse() (response *ModifyWebPageProtectSettingResponse) {
    response = &ModifyWebPageProtectSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebPageProtectSetting
// 修改网站防护设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebPageProtectSetting(request *ModifyWebPageProtectSettingRequest) (response *ModifyWebPageProtectSettingResponse, err error) {
    return c.ModifyWebPageProtectSettingWithContext(context.Background(), request)
}

// ModifyWebPageProtectSetting
// 修改网站防护设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebPageProtectSettingWithContext(ctx context.Context, request *ModifyWebPageProtectSettingRequest) (response *ModifyWebPageProtectSettingResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebPageProtectSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebPageProtectSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebPageProtectSwitchRequest() (request *ModifyWebPageProtectSwitchRequest) {
    request = &ModifyWebPageProtectSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyWebPageProtectSwitch")
    
    
    return
}

func NewModifyWebPageProtectSwitchResponse() (response *ModifyWebPageProtectSwitchResponse) {
    response = &ModifyWebPageProtectSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebPageProtectSwitch
// 网站防篡改防护设置开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_LICENSEEXCEEDED = "FailedOperation.LicenseExceeded"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebPageProtectSwitch(request *ModifyWebPageProtectSwitchRequest) (response *ModifyWebPageProtectSwitchResponse, err error) {
    return c.ModifyWebPageProtectSwitchWithContext(context.Background(), request)
}

// ModifyWebPageProtectSwitch
// 网站防篡改防护设置开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_LICENSEEXCEEDED = "FailedOperation.LicenseExceeded"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyWebPageProtectSwitchWithContext(ctx context.Context, request *ModifyWebPageProtectSwitchRequest) (response *ModifyWebPageProtectSwitchResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebPageProtectSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebPageProtectSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewRansomDefenseRollbackRequest() (request *RansomDefenseRollbackRequest) {
    request = &RansomDefenseRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RansomDefenseRollback")
    
    
    return
}

func NewRansomDefenseRollbackResponse() (response *RansomDefenseRollbackResponse) {
    response = &RansomDefenseRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RansomDefenseRollback
// 防勒索快照回滚
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RansomDefenseRollback(request *RansomDefenseRollbackRequest) (response *RansomDefenseRollbackResponse, err error) {
    return c.RansomDefenseRollbackWithContext(context.Background(), request)
}

// RansomDefenseRollback
// 防勒索快照回滚
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RansomDefenseRollbackWithContext(ctx context.Context, request *RansomDefenseRollbackRequest) (response *RansomDefenseRollbackResponse, err error) {
    if request == nil {
        request = NewRansomDefenseRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RansomDefenseRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewRansomDefenseRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
    request = &RecoverMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RecoverMalwares")
    
    
    return
}

func NewRecoverMalwaresResponse() (response *RecoverMalwaresResponse) {
    response = &RecoverMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverMalwares
// 本接口（RecoverMalwares）用于批量恢复已经被隔离的木马文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECOVER = "FailedOperation.Recover"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RecoverMalwares(request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    return c.RecoverMalwaresWithContext(context.Background(), request)
}

// RecoverMalwares
// 本接口（RecoverMalwares）用于批量恢复已经被隔离的木马文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECOVER = "FailedOperation.Recover"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RecoverMalwaresWithContext(ctx context.Context, request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    if request == nil {
        request = NewRecoverMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveLocalStorageItemRequest() (request *RemoveLocalStorageItemRequest) {
    request = &RemoveLocalStorageItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RemoveLocalStorageItem")
    
    
    return
}

func NewRemoveLocalStorageItemResponse() (response *RemoveLocalStorageItemResponse) {
    response = &RemoveLocalStorageItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveLocalStorageItem
// 删除本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RemoveLocalStorageItem(request *RemoveLocalStorageItemRequest) (response *RemoveLocalStorageItemResponse, err error) {
    return c.RemoveLocalStorageItemWithContext(context.Background(), request)
}

// RemoveLocalStorageItem
// 删除本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RemoveLocalStorageItemWithContext(ctx context.Context, request *RemoveLocalStorageItemRequest) (response *RemoveLocalStorageItemResponse, err error) {
    if request == nil {
        request = NewRemoveLocalStorageItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveLocalStorageItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveLocalStorageItemResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveMachineRequest() (request *RemoveMachineRequest) {
    request = &RemoveMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RemoveMachine")
    
    
    return
}

func NewRemoveMachineResponse() (response *RemoveMachineResponse) {
    response = &RemoveMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveMachine
// 删除主机所有记录，目前只支持非腾讯云主机，且需要主机在离线状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveMachine(request *RemoveMachineRequest) (response *RemoveMachineResponse, err error) {
    return c.RemoveMachineWithContext(context.Background(), request)
}

// RemoveMachine
// 删除主机所有记录，目前只支持非腾讯云主机，且需要主机在离线状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveMachineWithContext(ctx context.Context, request *RemoveMachineRequest) (response *RemoveMachineResponse, err error) {
    if request == nil {
        request = NewRemoveMachineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveMachineResponse()
    err = c.Send(request, response)
    return
}

func NewRetryCreateSnapshotRequest() (request *RetryCreateSnapshotRequest) {
    request = &RetryCreateSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RetryCreateSnapshot")
    
    
    return
}

func NewRetryCreateSnapshotResponse() (response *RetryCreateSnapshotResponse) {
    response = &RetryCreateSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RetryCreateSnapshot
// 快照创建失败时可以重试创建快照并且自动进行漏洞修复
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RetryCreateSnapshot(request *RetryCreateSnapshotRequest) (response *RetryCreateSnapshotResponse, err error) {
    return c.RetryCreateSnapshotWithContext(context.Background(), request)
}

// RetryCreateSnapshot
// 快照创建失败时可以重试创建快照并且自动进行漏洞修复
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RetryCreateSnapshotWithContext(ctx context.Context, request *RetryCreateSnapshotRequest) (response *RetryCreateSnapshotResponse, err error) {
    if request == nil {
        request = NewRetryCreateSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryCreateSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryCreateSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewRetryVulFixRequest() (request *RetryVulFixRequest) {
    request = &RetryVulFixRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "RetryVulFix")
    
    
    return
}

func NewRetryVulFixResponse() (response *RetryVulFixResponse) {
    response = &RetryVulFixResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RetryVulFix
// 修复失败时单独对某一个主机修复漏洞
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RetryVulFix(request *RetryVulFixRequest) (response *RetryVulFixResponse, err error) {
    return c.RetryVulFixWithContext(context.Background(), request)
}

// RetryVulFix
// 修复失败时单独对某一个主机修复漏洞
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RetryVulFixWithContext(ctx context.Context, request *RetryVulFixRequest) (response *RetryVulFixResponse, err error) {
    if request == nil {
        request = NewRetryVulFixRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryVulFix require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryVulFixResponse()
    err = c.Send(request, response)
    return
}

func NewScanAssetRequest() (request *ScanAssetRequest) {
    request = &ScanAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanAsset")
    
    
    return
}

func NewScanAssetResponse() (response *ScanAssetResponse) {
    response = &ScanAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanAsset
// 资产指纹启动扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanAsset(request *ScanAssetRequest) (response *ScanAssetResponse, err error) {
    return c.ScanAssetWithContext(context.Background(), request)
}

// ScanAsset
// 资产指纹启动扫描
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanAssetWithContext(ctx context.Context, request *ScanAssetRequest) (response *ScanAssetResponse, err error) {
    if request == nil {
        request = NewScanAssetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanAssetResponse()
    err = c.Send(request, response)
    return
}

func NewScanBaselineRequest() (request *ScanBaselineRequest) {
    request = &ScanBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanBaseline")
    
    
    return
}

func NewScanBaselineResponse() (response *ScanBaselineResponse) {
    response = &ScanBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanBaseline
// 基线检测与基线重新检测接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
func (c *Client) ScanBaseline(request *ScanBaselineRequest) (response *ScanBaselineResponse, err error) {
    return c.ScanBaselineWithContext(context.Background(), request)
}

// ScanBaseline
// 基线检测与基线重新检测接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
func (c *Client) ScanBaselineWithContext(ctx context.Context, request *ScanBaselineRequest) (response *ScanBaselineResponse, err error) {
    if request == nil {
        request = NewScanBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewScanTaskAgainRequest() (request *ScanTaskAgainRequest) {
    request = &ScanTaskAgainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanTaskAgain")
    
    
    return
}

func NewScanTaskAgainResponse() (response *ScanTaskAgainResponse) {
    response = &ScanTaskAgainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanTaskAgain
// ScanTaskAgain  重新开始扫描任务，可以指定机器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
func (c *Client) ScanTaskAgain(request *ScanTaskAgainRequest) (response *ScanTaskAgainResponse, err error) {
    return c.ScanTaskAgainWithContext(context.Background(), request)
}

// ScanTaskAgain
// ScanTaskAgain  重新开始扫描任务，可以指定机器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINENOTFOUND = "ResourceNotFound.MachineNotFound"
func (c *Client) ScanTaskAgainWithContext(ctx context.Context, request *ScanTaskAgainRequest) (response *ScanTaskAgainResponse, err error) {
    if request == nil {
        request = NewScanTaskAgainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanTaskAgain require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanTaskAgainResponse()
    err = c.Send(request, response)
    return
}

func NewScanVulRequest() (request *ScanVulRequest) {
    request = &ScanVulRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanVul")
    
    
    return
}

func NewScanVulResponse() (response *ScanVulResponse) {
    response = &ScanVulResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanVul
//  一键检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanVul(request *ScanVulRequest) (response *ScanVulResponse, err error) {
    return c.ScanVulWithContext(context.Background(), request)
}

// ScanVul
//  一键检测
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanVulWithContext(ctx context.Context, request *ScanVulRequest) (response *ScanVulResponse, err error) {
    if request == nil {
        request = NewScanVulRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanVul require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanVulResponse()
    err = c.Send(request, response)
    return
}

func NewScanVulAgainRequest() (request *ScanVulAgainRequest) {
    request = &ScanVulAgainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanVulAgain")
    
    
    return
}

func NewScanVulAgainResponse() (response *ScanVulAgainResponse) {
    response = &ScanVulAgainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanVulAgain
// 漏洞管理-重新检测接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanVulAgain(request *ScanVulAgainRequest) (response *ScanVulAgainResponse, err error) {
    return c.ScanVulAgainWithContext(context.Background(), request)
}

// ScanVulAgain
// 漏洞管理-重新检测接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanVulAgainWithContext(ctx context.Context, request *ScanVulAgainRequest) (response *ScanVulAgainResponse, err error) {
    if request == nil {
        request = NewScanVulAgainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanVulAgain require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanVulAgainResponse()
    err = c.Send(request, response)
    return
}

func NewScanVulSettingRequest() (request *ScanVulSettingRequest) {
    request = &ScanVulSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "ScanVulSetting")
    
    
    return
}

func NewScanVulSettingResponse() (response *ScanVulSettingResponse) {
    response = &ScanVulSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanVulSetting
// 定期扫描漏洞设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ScanVulSetting(request *ScanVulSettingRequest) (response *ScanVulSettingResponse, err error) {
    return c.ScanVulSettingWithContext(context.Background(), request)
}

// ScanVulSetting
// 定期扫描漏洞设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ScanVulSettingWithContext(ctx context.Context, request *ScanVulSettingRequest) (response *ScanVulSettingResponse, err error) {
    if request == nil {
        request = NewScanVulSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanVulSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanVulSettingResponse()
    err = c.Send(request, response)
    return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
    request = &SearchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SearchLog")
    
    
    return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
    response = &SearchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchLog
// 查询日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
    return c.SearchLogWithContext(context.Background(), request)
}

// SearchLog
// 查询日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SearchLogWithContext(ctx context.Context, request *SearchLogRequest) (response *SearchLogResponse, err error) {
    if request == nil {
        request = NewSearchLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchLogResponse()
    err = c.Send(request, response)
    return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
    request = &SeparateMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SeparateMalwares")
    
    
    return
}

func NewSeparateMalwaresResponse() (response *SeparateMalwaresResponse) {
    response = &SeparateMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SeparateMalwares
// 本接口（SeparateMalwares）用于隔离木马。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PARTSEPARATE = "FailedOperation.PartSeparate"
//  FAILEDOPERATION_SINGLESEPARATE = "FailedOperation.SingleSeparate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SeparateMalwares(request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    return c.SeparateMalwaresWithContext(context.Background(), request)
}

// SeparateMalwares
// 本接口（SeparateMalwares）用于隔离木马。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PARTSEPARATE = "FailedOperation.PartSeparate"
//  FAILEDOPERATION_SINGLESEPARATE = "FailedOperation.SingleSeparate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SeparateMalwaresWithContext(ctx context.Context, request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    if request == nil {
        request = NewSeparateMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SeparateMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewSeparateMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewSetBashEventsStatusRequest() (request *SetBashEventsStatusRequest) {
    request = &SetBashEventsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SetBashEventsStatus")
    
    
    return
}

func NewSetBashEventsStatusResponse() (response *SetBashEventsStatusResponse) {
    response = &SetBashEventsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetBashEventsStatus
// 设置高危命令事件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetBashEventsStatus(request *SetBashEventsStatusRequest) (response *SetBashEventsStatusResponse, err error) {
    return c.SetBashEventsStatusWithContext(context.Background(), request)
}

// SetBashEventsStatus
// 设置高危命令事件状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetBashEventsStatusWithContext(ctx context.Context, request *SetBashEventsStatusRequest) (response *SetBashEventsStatusResponse, err error) {
    if request == nil {
        request = NewSetBashEventsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetBashEventsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetBashEventsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSetLocalStorageExpireRequest() (request *SetLocalStorageExpireRequest) {
    request = &SetLocalStorageExpireRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SetLocalStorageExpire")
    
    
    return
}

func NewSetLocalStorageExpireResponse() (response *SetLocalStorageExpireResponse) {
    response = &SetLocalStorageExpireResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetLocalStorageExpire
// 设置本地存储过期时间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) SetLocalStorageExpire(request *SetLocalStorageExpireRequest) (response *SetLocalStorageExpireResponse, err error) {
    return c.SetLocalStorageExpireWithContext(context.Background(), request)
}

// SetLocalStorageExpire
// 设置本地存储过期时间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) SetLocalStorageExpireWithContext(ctx context.Context, request *SetLocalStorageExpireRequest) (response *SetLocalStorageExpireResponse, err error) {
    if request == nil {
        request = NewSetLocalStorageExpireRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLocalStorageExpire require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLocalStorageExpireResponse()
    err = c.Send(request, response)
    return
}

func NewSetLocalStorageItemRequest() (request *SetLocalStorageItemRequest) {
    request = &SetLocalStorageItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SetLocalStorageItem")
    
    
    return
}

func NewSetLocalStorageItemResponse() (response *SetLocalStorageItemResponse) {
    response = &SetLocalStorageItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetLocalStorageItem
// 设置本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) SetLocalStorageItem(request *SetLocalStorageItemRequest) (response *SetLocalStorageItemResponse, err error) {
    return c.SetLocalStorageItemWithContext(context.Background(), request)
}

// SetLocalStorageItem
// 设置本地存储数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) SetLocalStorageItemWithContext(ctx context.Context, request *SetLocalStorageItemRequest) (response *SetLocalStorageItemResponse, err error) {
    if request == nil {
        request = NewSetLocalStorageItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLocalStorageItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLocalStorageItemResponse()
    err = c.Send(request, response)
    return
}

func NewStartBaselineDetectRequest() (request *StartBaselineDetectRequest) {
    request = &StartBaselineDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "StartBaselineDetect")
    
    
    return
}

func NewStartBaselineDetectResponse() (response *StartBaselineDetectResponse) {
    response = &StartBaselineDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartBaselineDetect
// 检测基线
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartBaselineDetect(request *StartBaselineDetectRequest) (response *StartBaselineDetectResponse, err error) {
    return c.StartBaselineDetectWithContext(context.Background(), request)
}

// StartBaselineDetect
// 检测基线
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StartBaselineDetectWithContext(ctx context.Context, request *StartBaselineDetectRequest) (response *StartBaselineDetectResponse, err error) {
    if request == nil {
        request = NewStartBaselineDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBaselineDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBaselineDetectResponse()
    err = c.Send(request, response)
    return
}

func NewStopAssetScanRequest() (request *StopAssetScanRequest) {
    request = &StopAssetScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "StopAssetScan")
    
    
    return
}

func NewStopAssetScanResponse() (response *StopAssetScanResponse) {
    response = &StopAssetScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopAssetScan
// 停止资产扫描任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) StopAssetScan(request *StopAssetScanRequest) (response *StopAssetScanResponse, err error) {
    return c.StopAssetScanWithContext(context.Background(), request)
}

// StopAssetScan
// 停止资产扫描任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) StopAssetScanWithContext(ctx context.Context, request *StopAssetScanRequest) (response *StopAssetScanResponse, err error) {
    if request == nil {
        request = NewStopAssetScanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAssetScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAssetScanResponse()
    err = c.Send(request, response)
    return
}

func NewStopBaselineDetectRequest() (request *StopBaselineDetectRequest) {
    request = &StopBaselineDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "StopBaselineDetect")
    
    
    return
}

func NewStopBaselineDetectResponse() (response *StopBaselineDetectResponse) {
    response = &StopBaselineDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopBaselineDetect
// 停止基线检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopBaselineDetect(request *StopBaselineDetectRequest) (response *StopBaselineDetectResponse, err error) {
    return c.StopBaselineDetectWithContext(context.Background(), request)
}

// StopBaselineDetect
// 停止基线检测
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopBaselineDetectWithContext(ctx context.Context, request *StopBaselineDetectRequest) (response *StopBaselineDetectResponse, err error) {
    if request == nil {
        request = NewStopBaselineDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopBaselineDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopBaselineDetectResponse()
    err = c.Send(request, response)
    return
}

func NewStopNoticeBanTipsRequest() (request *StopNoticeBanTipsRequest) {
    request = &StopNoticeBanTipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "StopNoticeBanTips")
    
    
    return
}

func NewStopNoticeBanTipsResponse() (response *StopNoticeBanTipsResponse) {
    response = &StopNoticeBanTipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopNoticeBanTips
// 不再提醒爆破阻断提示弹窗
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) StopNoticeBanTips(request *StopNoticeBanTipsRequest) (response *StopNoticeBanTipsResponse, err error) {
    return c.StopNoticeBanTipsWithContext(context.Background(), request)
}

// StopNoticeBanTips
// 不再提醒爆破阻断提示弹窗
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) StopNoticeBanTipsWithContext(ctx context.Context, request *StopNoticeBanTipsRequest) (response *StopNoticeBanTipsResponse, err error) {
    if request == nil {
        request = NewStopNoticeBanTipsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopNoticeBanTips require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopNoticeBanTipsResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchBashRulesRequest() (request *SwitchBashRulesRequest) {
    request = &SwitchBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SwitchBashRules")
    
    
    return
}

func NewSwitchBashRulesResponse() (response *SwitchBashRulesResponse) {
    response = &SwitchBashRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchBashRules
// 切换高危命令规则状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SwitchBashRules(request *SwitchBashRulesRequest) (response *SwitchBashRulesResponse, err error) {
    return c.SwitchBashRulesWithContext(context.Background(), request)
}

// SwitchBashRules
// 切换高危命令规则状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SwitchBashRulesWithContext(ctx context.Context, request *SwitchBashRulesRequest) (response *SwitchBashRulesResponse, err error) {
    if request == nil {
        request = NewSwitchBashRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchBashRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchBashRulesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncAssetScanRequest() (request *SyncAssetScanRequest) {
    request = &SyncAssetScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SyncAssetScan")
    
    
    return
}

func NewSyncAssetScanResponse() (response *SyncAssetScanResponse) {
    response = &SyncAssetScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncAssetScan
// 同步资产扫描信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetScan(request *SyncAssetScanRequest) (response *SyncAssetScanResponse, err error) {
    return c.SyncAssetScanWithContext(context.Background(), request)
}

// SyncAssetScan
// 同步资产扫描信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncAssetScanWithContext(ctx context.Context, request *SyncAssetScanRequest) (response *SyncAssetScanResponse, err error) {
    if request == nil {
        request = NewSyncAssetScanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncAssetScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncAssetScanResponse()
    err = c.Send(request, response)
    return
}

func NewSyncBaselineDetectSummaryRequest() (request *SyncBaselineDetectSummaryRequest) {
    request = &SyncBaselineDetectSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SyncBaselineDetectSummary")
    
    
    return
}

func NewSyncBaselineDetectSummaryResponse() (response *SyncBaselineDetectSummaryResponse) {
    response = &SyncBaselineDetectSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncBaselineDetectSummary
// 同步基线检测进度概要
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) SyncBaselineDetectSummary(request *SyncBaselineDetectSummaryRequest) (response *SyncBaselineDetectSummaryResponse, err error) {
    return c.SyncBaselineDetectSummaryWithContext(context.Background(), request)
}

// SyncBaselineDetectSummary
// 同步基线检测进度概要
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) SyncBaselineDetectSummaryWithContext(ctx context.Context, request *SyncBaselineDetectSummaryRequest) (response *SyncBaselineDetectSummaryResponse, err error) {
    if request == nil {
        request = NewSyncBaselineDetectSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncBaselineDetectSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncBaselineDetectSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewSyncMachinesRequest() (request *SyncMachinesRequest) {
    request = &SyncMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "SyncMachines")
    
    
    return
}

func NewSyncMachinesResponse() (response *SyncMachinesResponse) {
    response = &SyncMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncMachines
// 同步机器信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SyncMachines(request *SyncMachinesRequest) (response *SyncMachinesResponse, err error) {
    return c.SyncMachinesWithContext(context.Background(), request)
}

// SyncMachines
// 同步机器信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SyncMachinesWithContext(ctx context.Context, request *SyncMachinesRequest) (response *SyncMachinesResponse, err error) {
    if request == nil {
        request = NewSyncMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewTestWebHookRuleRequest() (request *TestWebHookRuleRequest) {
    request = &TestWebHookRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "TestWebHookRule")
    
    
    return
}

func NewTestWebHookRuleResponse() (response *TestWebHookRuleResponse) {
    response = &TestWebHookRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TestWebHookRule
// 测试企微机器人规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TestWebHookRule(request *TestWebHookRuleRequest) (response *TestWebHookRuleResponse, err error) {
    return c.TestWebHookRuleWithContext(context.Background(), request)
}

// TestWebHookRule
// 测试企微机器人规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TestWebHookRuleWithContext(ctx context.Context, request *TestWebHookRuleRequest) (response *TestWebHookRuleResponse, err error) {
    if request == nil {
        request = NewTestWebHookRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TestWebHookRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewTestWebHookRuleResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
    request = &TrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "TrustMalwares")
    
    
    return
}

func NewTrustMalwaresResponse() (response *TrustMalwaresResponse) {
    response = &TrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TrustMalwares
// 本接口(TrustMalwares)将被识别木马文件设为信任。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    return c.TrustMalwaresWithContext(context.Background(), request)
}

// TrustMalwares
// 本接口(TrustMalwares)将被识别木马文件设为信任。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) TrustMalwaresWithContext(ctx context.Context, request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrustMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
    request = &UntrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "UntrustMalwares")
    
    
    return
}

func NewUntrustMalwaresResponse() (response *UntrustMalwaresResponse) {
    response = &UntrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UntrustMalwares
// 本接口（UntrustMalwares）用于取消信任木马文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMalwares(request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    return c.UntrustMalwaresWithContext(context.Background(), request)
}

// UntrustMalwares
// 本接口（UntrustMalwares）用于取消信任木马文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMalwaresWithContext(ctx context.Context, request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    if request == nil {
        request = NewUntrustMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UntrustMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewUntrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBaselineStrategyRequest() (request *UpdateBaselineStrategyRequest) {
    request = &UpdateBaselineStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "UpdateBaselineStrategy")
    
    
    return
}

func NewUpdateBaselineStrategyResponse() (response *UpdateBaselineStrategyResponse) {
    response = &UpdateBaselineStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBaselineStrategy
// 根据基线策略id更新策略信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateBaselineStrategy(request *UpdateBaselineStrategyRequest) (response *UpdateBaselineStrategyResponse, err error) {
    return c.UpdateBaselineStrategyWithContext(context.Background(), request)
}

// UpdateBaselineStrategy
// 根据基线策略id更新策略信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateBaselineStrategyWithContext(ctx context.Context, request *UpdateBaselineStrategyRequest) (response *UpdateBaselineStrategyResponse, err error) {
    if request == nil {
        request = NewUpdateBaselineStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateBaselineStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateBaselineStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateMachineTagsRequest() (request *UpdateMachineTagsRequest) {
    request = &UpdateMachineTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "UpdateMachineTags")
    
    
    return
}

func NewUpdateMachineTagsResponse() (response *UpdateMachineTagsResponse) {
    response = &UpdateMachineTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateMachineTags
// 关联机器标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateMachineTags(request *UpdateMachineTagsRequest) (response *UpdateMachineTagsResponse, err error) {
    return c.UpdateMachineTagsWithContext(context.Background(), request)
}

// UpdateMachineTags
// 关联机器标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateMachineTagsWithContext(ctx context.Context, request *UpdateMachineTagsRequest) (response *UpdateMachineTagsResponse, err error) {
    if request == nil {
        request = NewUpdateMachineTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateMachineTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateMachineTagsResponse()
    err = c.Send(request, response)
    return
}
