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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBaselinePolicy(request *DeleteBaselinePolicyRequest) (response *DeleteBaselinePolicyResponse, err error) {
    return c.DeleteBaselinePolicyWithContext(context.Background(), request)
}

// DeleteBaselinePolicy
// 删除基线策略配置
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
// 本接口（DeleteMachine）用于卸载云镜客户端。
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
// 本接口（DeleteMachine）用于卸载云镜客户端。
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineHostIgnoreList(request *DescribeBaselineHostIgnoreListRequest) (response *DescribeBaselineHostIgnoreListResponse, err error) {
    return c.DescribeBaselineHostIgnoreListWithContext(context.Background(), request)
}

// DescribeBaselineHostIgnoreList
// 获取忽略规则主机列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemIgnoreList(request *DescribeBaselineItemIgnoreListRequest) (response *DescribeBaselineItemIgnoreListResponse, err error) {
    return c.DescribeBaselineItemIgnoreListWithContext(context.Background(), request)
}

// DescribeBaselineItemIgnoreList
// 获取忽略规则项列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineItemList(request *DescribeBaselineItemListRequest) (response *DescribeBaselineItemListResponse, err error) {
    return c.DescribeBaselineItemListWithContext(context.Background(), request)
}

// DescribeBaselineItemList
// 获取基线项检测结果列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselinePolicyList(request *DescribeBaselinePolicyListRequest) (response *DescribeBaselinePolicyListResponse, err error) {
    return c.DescribeBaselinePolicyListWithContext(context.Background(), request)
}

// DescribeBaselinePolicyList
// 获取基线策略列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineRuleList(request *DescribeBaselineRuleListRequest) (response *DescribeBaselineRuleListResponse, err error) {
    return c.DescribeBaselineRuleListWithContext(context.Background(), request)
}

// DescribeBaselineRuleList
// 获取基线规则列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBaselineWeakPasswordList(request *DescribeBaselineWeakPasswordListRequest) (response *DescribeBaselineWeakPasswordListResponse, err error) {
    return c.DescribeBaselineWeakPasswordListWithContext(context.Background(), request)
}

// DescribeBaselineWeakPasswordList
// 获取基线弱口令列表
//
// 可能返回的错误码:
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
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeClientException(request *DescribeClientExceptionRequest) (response *DescribeClientExceptionResponse, err error) {
    return c.DescribeClientExceptionWithContext(context.Background(), request)
}

// DescribeClientException
// 获取客户端异常事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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

func NewDescribeSaveOrUpdateWarningsRequest() (request *DescribeSaveOrUpdateWarningsRequest) {
    request = &DescribeSaveOrUpdateWarningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSaveOrUpdateWarnings")
    
    
    return
}

func NewDescribeSaveOrUpdateWarningsResponse() (response *DescribeSaveOrUpdateWarningsResponse) {
    response = &DescribeSaveOrUpdateWarningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSaveOrUpdateWarnings
// 更新或者插入用户告警设置(该接口废弃,请调用 ModifyWarningSetting )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSaveOrUpdateWarnings(request *DescribeSaveOrUpdateWarningsRequest) (response *DescribeSaveOrUpdateWarningsResponse, err error) {
    return c.DescribeSaveOrUpdateWarningsWithContext(context.Background(), request)
}

// DescribeSaveOrUpdateWarnings
// 更新或者插入用户告警设置(该接口废弃,请调用 ModifyWarningSetting )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSaveOrUpdateWarningsWithContext(ctx context.Context, request *DescribeSaveOrUpdateWarningsRequest) (response *DescribeSaveOrUpdateWarningsResponse, err error) {
    if request == nil {
        request = NewDescribeSaveOrUpdateWarningsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSaveOrUpdateWarnings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSaveOrUpdateWarningsResponse()
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

func NewDescribeSearchExportListRequest() (request *DescribeSearchExportListRequest) {
    request = &DescribeSearchExportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeSearchExportList")
    
    
    return
}

func NewDescribeSearchExportListResponse() (response *DescribeSearchExportListResponse) {
    response = &DescribeSearchExportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSearchExportList
// 导出ES查询文档列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSearchExportList(request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
    return c.DescribeSearchExportListWithContext(context.Background(), request)
}

// DescribeSearchExportList
// 导出ES查询文档列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSearchExportListWithContext(ctx context.Context, request *DescribeSearchExportListRequest) (response *DescribeSearchExportListResponse, err error) {
    if request == nil {
        request = NewDescribeSearchExportListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchExportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchExportListResponse()
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
func (c *Client) DestroyOrder(request *DestroyOrderRequest) (response *DestroyOrderResponse, err error) {
    return c.DestroyOrderWithContext(context.Background(), request)
}

// DestroyOrder
// DestroyOrder  该接口可以对资源销毁.
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
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
