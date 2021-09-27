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
    if request == nil {
        request = NewCancelIgnoreVulRequest()
    }
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
    if request == nil {
        request = NewChangeRuleEventsIgnoreStatusRequest()
    }
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
    if request == nil {
        request = NewCheckBashRuleParamsRequest()
    }
    response = NewCheckBashRuleParamsResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
    request = &CloseProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "CloseProVersion")
    return
}

func NewCloseProVersionResponse() (response *CloseProVersionResponse) {
    response = &CloseProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseProVersion
// 本接口 (CloseProVersion) 用于关闭专业版。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOSEPROVERSION = "FailedOperation.CloseProVersion"
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  FAILEDOPERATION_PREPAYMODE = "FailedOperation.PrePayMode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CloseProVersion(request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    if request == nil {
        request = NewCloseProVersionRequest()
    }
    response = NewCloseProVersionResponse()
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
    if request == nil {
        request = NewCreateBaselineStrategyRequest()
    }
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEmergencyVulScan(request *CreateEmergencyVulScanRequest) (response *CreateEmergencyVulScanResponse, err error) {
    if request == nil {
        request = NewCreateEmergencyVulScanRequest()
    }
    response = NewCreateEmergencyVulScanResponse()
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
    if request == nil {
        request = NewCreateProtectServerRequest()
    }
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateScanMalwareSetting(request *CreateScanMalwareSettingRequest) (response *CreateScanMalwareSettingResponse, err error) {
    if request == nil {
        request = NewCreateScanMalwareSettingRequest()
    }
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
    if request == nil {
        request = NewCreateSearchLogRequest()
    }
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
    if request == nil {
        request = NewCreateSearchTemplateRequest()
    }
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
    if request == nil {
        request = NewDeleteAttackLogsRequest()
    }
    response = NewDeleteAttackLogsResponse()
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
    if request == nil {
        request = NewDeleteBaselineStrategyRequest()
    }
    response = NewDeleteBaselineStrategyResponse()
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
    if request == nil {
        request = NewDeleteBashEventsRequest()
    }
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
    if request == nil {
        request = NewDeleteBashRulesRequest()
    }
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
    if request == nil {
        request = NewDeleteBruteAttacksRequest()
    }
    response = NewDeleteBruteAttacksResponse()
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
    if request == nil {
        request = NewDeleteLoginWhiteListRequest()
    }
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
    if request == nil {
        request = NewDeleteMachineRequest()
    }
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
    if request == nil {
        request = NewDeleteMachineTagRequest()
    }
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
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
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
    if request == nil {
        request = NewDeleteMalwareScanTaskRequest()
    }
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
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    if request == nil {
        request = NewDeleteMalwaresRequest()
    }
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
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteNonlocalLoginPlacesRequest()
    }
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
    if request == nil {
        request = NewDeletePrivilegeEventsRequest()
    }
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
    if request == nil {
        request = NewDeletePrivilegeRulesRequest()
    }
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
    if request == nil {
        request = NewDeleteProtectDirRequest()
    }
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
    if request == nil {
        request = NewDeleteReverseShellEventsRequest()
    }
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
    if request == nil {
        request = NewDeleteReverseShellRulesRequest()
    }
    response = NewDeleteReverseShellRulesResponse()
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
    if request == nil {
        request = NewDeleteSearchTemplateRequest()
    }
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
    if request == nil {
        request = NewDeleteTagsRequest()
    }
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
    if request == nil {
        request = NewDeleteWebPageEventLogRequest()
    }
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
    if request == nil {
        request = NewDescribeAccountStatisticsRequest()
    }
    response = NewDescribeAccountStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 该接口已废弃
//
// 
//
// 本接口 (DescribeAccounts) 用于获取帐号列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
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
    if request == nil {
        request = NewDescribeAssetAppListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetAppProcessListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetCoreModuleInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetCoreModuleListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetDatabaseInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetDatabaseListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetEnvListRequest()
    }
    response = NewDescribeAssetEnvListResponse()
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
    if request == nil {
        request = NewDescribeAssetInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetInitServiceListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetJarInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetJarListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetMachineDetailRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetMachineListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetPlanTaskListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetPortInfoListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetProcessInfoListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetRecentMachineInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetSystemPackageListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetUserInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetUserListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebAppListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebAppPluginListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebFrameListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebLocationInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebLocationListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebServiceInfoListRequest()
    }
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
    if request == nil {
        request = NewDescribeAssetWebServiceProcessListRequest()
    }
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
    if request == nil {
        request = NewDescribeAttackLogInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeAttackLogsRequest()
    }
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
    if request == nil {
        request = NewDescribeAttackVulTypeListRequest()
    }
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
    if request == nil {
        request = NewDescribeAvailableExpertServiceDetailRequest()
    }
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
    if request == nil {
        request = NewDescribeBanModeRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBanRegions(request *DescribeBanRegionsRequest) (response *DescribeBanRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeBanRegionsRequest()
    }
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
    if request == nil {
        request = NewDescribeBanStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeBanWhiteListRequest()
    }
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
func (c *Client) DescribeBaselineAnalysisData(request *DescribeBaselineAnalysisDataRequest) (response *DescribeBaselineAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineAnalysisDataRequest()
    }
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
    if request == nil {
        request = NewDescribeBaselineBasicInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeBaselineDetailRequest()
    }
    response = NewDescribeBaselineDetailResponse()
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
    if request == nil {
        request = NewDescribeBaselineEffectHostListRequest()
    }
    response = NewDescribeBaselineEffectHostListResponse()
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
    if request == nil {
        request = NewDescribeBaselineHostTopRequest()
    }
    response = NewDescribeBaselineHostTopResponse()
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
    if request == nil {
        request = NewDescribeBaselineListRequest()
    }
    response = NewDescribeBaselineListResponse()
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
    if request == nil {
        request = NewDescribeBaselineRuleRequest()
    }
    response = NewDescribeBaselineRuleResponse()
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
    if request == nil {
        request = NewDescribeBaselineScanScheduleRequest()
    }
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
    if request == nil {
        request = NewDescribeBaselineStrategyDetailRequest()
    }
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
    if request == nil {
        request = NewDescribeBaselineStrategyListRequest()
    }
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
    if request == nil {
        request = NewDescribeBaselineTopRequest()
    }
    response = NewDescribeBaselineTopResponse()
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
    if request == nil {
        request = NewDescribeBashEventsRequest()
    }
    response = NewDescribeBashEventsResponse()
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
    if request == nil {
        request = NewDescribeBashRulesRequest()
    }
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
func (c *Client) DescribeBruteAttackList(request *DescribeBruteAttackListRequest) (response *DescribeBruteAttackListResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttackListRequest()
    }
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
    if request == nil {
        request = NewDescribeBruteAttackRulesRequest()
    }
    response = NewDescribeBruteAttackRulesResponse()
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
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
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
    if request == nil {
        request = NewDescribeESAggregationsRequest()
    }
    response = NewDescribeESAggregationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeESHitsRequest() (request *DescribeESHitsRequest) {
    request = &DescribeESHitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeESHits")
    return
}

func NewDescribeESHitsResponse() (response *DescribeESHitsResponse) {
    response = &DescribeESHitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeESHits
// 获取ES查询文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
    if request == nil {
        request = NewDescribeESHitsRequest()
    }
    response = NewDescribeESHitsResponse()
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
    if request == nil {
        request = NewDescribeEmergencyResponseListRequest()
    }
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
    if request == nil {
        request = NewDescribeEmergencyVulListRequest()
    }
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
    if request == nil {
        request = NewDescribeExpertServiceListRequest()
    }
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
    if request == nil {
        request = NewDescribeExpertServiceOrderListRequest()
    }
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
    if request == nil {
        request = NewDescribeExportMachinesRequest()
    }
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
    if request == nil {
        request = NewDescribeGeneralStatRequest()
    }
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
    if request == nil {
        request = NewDescribeHistoryAccountsRequest()
    }
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
    if request == nil {
        request = NewDescribeHistoryServiceRequest()
    }
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
    if request == nil {
        request = NewDescribeHostLoginListRequest()
    }
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
    if request == nil {
        request = NewDescribeIgnoreBaselineRuleRequest()
    }
    response = NewDescribeIgnoreBaselineRuleResponse()
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
    if request == nil {
        request = NewDescribeIgnoreRuleEffectHostListRequest()
    }
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
    if request == nil {
        request = NewDescribeImportMachineInfoRequest()
    }
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
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    response = NewDescribeIndexListResponse()
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
    if request == nil {
        request = NewDescribeLogStorageStatisticRequest()
    }
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
func (c *Client) DescribeLoginWhiteCombinedList(request *DescribeLoginWhiteCombinedListRequest) (response *DescribeLoginWhiteCombinedListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteCombinedListRequest()
    }
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
    if request == nil {
        request = NewDescribeLoginWhiteListRequest()
    }
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
    if request == nil {
        request = NewDescribeMachineInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeMachineListRequest()
    }
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
    if request == nil {
        request = NewDescribeMachineOsListRequest()
    }
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
    if request == nil {
        request = NewDescribeMachineRegionsRequest()
    }
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
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
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
    if request == nil {
        request = NewDescribeMalWareListRequest()
    }
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
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMaliciousRequestWhiteList(request *DescribeMaliciousRequestWhiteListRequest) (response *DescribeMaliciousRequestWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestWhiteListRequest()
    }
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
    if request == nil {
        request = NewDescribeMalwareFileRequest()
    }
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
    if request == nil {
        request = NewDescribeMalwareInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeMalwareRiskWarningRequest()
    }
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
    if request == nil {
        request = NewDescribeMalwareTimingScanSettingRequest()
    }
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
    if request == nil {
        request = NewDescribeMonthInspectionReportRequest()
    }
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
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
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
    if request == nil {
        request = NewDescribeOverviewStatisticsRequest()
    }
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
    if request == nil {
        request = NewDescribePrivilegeEventsRequest()
    }
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
    if request == nil {
        request = NewDescribePrivilegeRulesRequest()
    }
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
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeProVersionStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
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
    if request == nil {
        request = NewDescribeProtectDirListRequest()
    }
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
    if request == nil {
        request = NewDescribeProtectDirRelatedServerRequest()
    }
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
// 专家服务-旗舰护网列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProtectNetList(request *DescribeProtectNetListRequest) (response *DescribeProtectNetListResponse, err error) {
    if request == nil {
        request = NewDescribeProtectNetListRequest()
    }
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
    if request == nil {
        request = NewDescribeReverseShellEventsRequest()
    }
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
    if request == nil {
        request = NewDescribeReverseShellRulesRequest()
    }
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
    if request == nil {
        request = NewDescribeRiskDnsListRequest()
    }
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
    if request == nil {
        request = NewDescribeSaveOrUpdateWarningsRequest()
    }
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
    if request == nil {
        request = NewDescribeScanMalwareScheduleRequest()
    }
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
    if request == nil {
        request = NewDescribeScanScheduleRequest()
    }
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
    if request == nil {
        request = NewDescribeScanStateRequest()
    }
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
    if request == nil {
        request = NewDescribeScanTaskDetailsRequest()
    }
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
    if request == nil {
        request = NewDescribeScanTaskStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeScanVulSettingRequest()
    }
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
    if request == nil {
        request = NewDescribeSearchExportListRequest()
    }
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
    if request == nil {
        request = NewDescribeSearchLogsRequest()
    }
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
    if request == nil {
        request = NewDescribeSearchTemplatesRequest()
    }
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
    if request == nil {
        request = NewDescribeSecurityDynamicsRequest()
    }
    response = NewDescribeSecurityDynamicsResponse()
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
func (c *Client) DescribeSecurityEventsCnt(request *DescribeSecurityEventsCntRequest) (response *DescribeSecurityEventsCntResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityEventsCntRequest()
    }
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
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
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
    if request == nil {
        request = NewDescribeServerRelatedDirInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeServersAndRiskAndFirstInfoRequest()
    }
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
    if request == nil {
        request = NewDescribeStrategyExistRequest()
    }
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
    if request == nil {
        request = NewDescribeTagMachinesRequest()
    }
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
    if request == nil {
        request = NewDescribeTagsRequest()
    }
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
    if request == nil {
        request = NewDescribeUndoVulCountsRequest()
    }
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
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
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
    if request == nil {
        request = NewDescribeVersionStatisticsRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulCountByDates(request *DescribeVulCountByDatesRequest) (response *DescribeVulCountByDatesResponse, err error) {
    if request == nil {
        request = NewDescribeVulCountByDatesRequest()
    }
    response = NewDescribeVulCountByDatesResponse()
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
    if request == nil {
        request = NewDescribeVulHostCountScanTimeRequest()
    }
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
    if request == nil {
        request = NewDescribeVulHostTopRequest()
    }
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
    if request == nil {
        request = NewDescribeVulInfoCvssRequest()
    }
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
    if request == nil {
        request = NewDescribeVulLevelCountRequest()
    }
    response = NewDescribeVulLevelCountResponse()
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
    if request == nil {
        request = NewDescribeVulTopRequest()
    }
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
    if request == nil {
        request = NewDescribeWarningListRequest()
    }
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
    if request == nil {
        request = NewDescribeWebPageEventListRequest()
    }
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
    if request == nil {
        request = NewDescribeWebPageGeneralizeRequest()
    }
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
    if request == nil {
        request = NewDescribeWebPageProtectStatRequest()
    }
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
    if request == nil {
        request = NewDescribeWebPageServiceInfoRequest()
    }
    response = NewDescribeWebPageServiceInfoResponse()
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
    if request == nil {
        request = NewEditBashRulesRequest()
    }
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
    if request == nil {
        request = NewEditTagsRequest()
    }
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
    if request == nil {
        request = NewExportAssetCoreModuleListRequest()
    }
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
    if request == nil {
        request = NewExportAssetWebServiceInfoListRequest()
    }
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
    if request == nil {
        request = NewExportAttackLogsRequest()
    }
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
    if request == nil {
        request = NewExportBaselineEffectHostListRequest()
    }
    response = NewExportBaselineEffectHostListResponse()
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
    if request == nil {
        request = NewExportBaselineListRequest()
    }
    response = NewExportBaselineListResponse()
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
    if request == nil {
        request = NewExportBashEventsRequest()
    }
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
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
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
    if request == nil {
        request = NewExportIgnoreBaselineRuleRequest()
    }
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
    if request == nil {
        request = NewExportIgnoreRuleEffectHostListRequest()
    }
    response = NewExportIgnoreRuleEffectHostListResponse()
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
    if request == nil {
        request = NewExportMaliciousRequestsRequest()
    }
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
    if request == nil {
        request = NewExportMalwaresRequest()
    }
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
    if request == nil {
        request = NewExportNonlocalLoginPlacesRequest()
    }
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
    if request == nil {
        request = NewExportPrivilegeEventsRequest()
    }
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
    if request == nil {
        request = NewExportProtectDirListRequest()
    }
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
    if request == nil {
        request = NewExportReverseShellEventsRequest()
    }
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
    if request == nil {
        request = NewExportScanTaskDetailsRequest()
    }
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
    if request == nil {
        request = NewExportSecurityTrendsRequest()
    }
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
    if request == nil {
        request = NewExportTasksRequest()
    }
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
    if request == nil {
        request = NewExportVulDetectionExcelRequest()
    }
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
    if request == nil {
        request = NewExportVulDetectionReportRequest()
    }
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
    if request == nil {
        request = NewExportVulEffectHostListRequest()
    }
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
    if request == nil {
        request = NewExportVulListRequest()
    }
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
    if request == nil {
        request = NewExportWebPageEventListRequest()
    }
    response = NewExportWebPageEventListResponse()
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
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    if request == nil {
        request = NewIgnoreImpactedHostsRequest()
    }
    response = NewIgnoreImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceOpenProVersionPrepaidRequest() (request *InquiryPriceOpenProVersionPrepaidRequest) {
    request = &InquiryPriceOpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "InquiryPriceOpenProVersionPrepaid")
    return
}

func NewInquiryPriceOpenProVersionPrepaidResponse() (response *InquiryPriceOpenProVersionPrepaidResponse) {
    response = &InquiryPriceOpenProVersionPrepaidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceOpenProVersionPrepaid
// 本接口 (InquiryPriceOpenProVersionPrepaid) 用于开通专业版询价(预付费)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INQUIRYPRICE = "FailedOperation.InquiryPrice"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) InquiryPriceOpenProVersionPrepaid(request *InquiryPriceOpenProVersionPrepaidRequest) (response *InquiryPriceOpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewInquiryPriceOpenProVersionPrepaidRequest()
    }
    response = NewInquiryPriceOpenProVersionPrepaidResponse()
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
//  用于设置新增主机自动开通专业版配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
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
    if request == nil {
        request = NewModifyBanModeRequest()
    }
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
    if request == nil {
        request = NewModifyBanStatusRequest()
    }
    response = NewModifyBanStatusResponse()
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
    if request == nil {
        request = NewModifyBruteAttackRulesRequest()
    }
    response = NewModifyBruteAttackRulesResponse()
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
    if request == nil {
        request = NewModifyMalwareTimingScanSettingsRequest()
    }
    response = NewModifyMalwareTimingScanSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
    request = &ModifyProVersionRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyProVersionRenewFlag")
    return
}

func NewModifyProVersionRenewFlagResponse() (response *ModifyProVersionRenewFlagResponse) {
    response = &ModifyProVersionRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProVersionRenewFlag
// 本接口 (ModifyProVersionRenewFlag) 用于修改专业版包年包月续费标识。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    response = NewModifyProVersionRenewFlagResponse()
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
    if request == nil {
        request = NewModifyWarningSettingRequest()
    }
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
    if request == nil {
        request = NewModifyWebPageProtectDirRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyWebPageProtectSetting(request *ModifyWebPageProtectSettingRequest) (response *ModifyWebPageProtectSettingResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectSettingRequest()
    }
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
func (c *Client) ModifyWebPageProtectSwitch(request *ModifyWebPageProtectSwitchRequest) (response *ModifyWebPageProtectSwitchResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectSwitchRequest()
    }
    response = NewModifyWebPageProtectSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
    request = &OpenProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "OpenProVersion")
    return
}

func NewOpenProVersionResponse() (response *OpenProVersionResponse) {
    response = &OpenProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenProVersion
// 本接口 (OpenProVersion) 用于开通专业版。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) OpenProVersion(request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    if request == nil {
        request = NewOpenProVersionRequest()
    }
    response = NewOpenProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionPrepaidRequest() (request *OpenProVersionPrepaidRequest) {
    request = &OpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "OpenProVersionPrepaid")
    return
}

func NewOpenProVersionPrepaidResponse() (response *OpenProVersionPrepaidResponse) {
    response = &OpenProVersionPrepaidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenProVersionPrepaid
// 本接口 (OpenProVersionPrepaid) 用于开通专业版(包年包月)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) OpenProVersionPrepaid(request *OpenProVersionPrepaidRequest) (response *OpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewOpenProVersionPrepaidRequest()
    }
    response = NewOpenProVersionPrepaidResponse()
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
    if request == nil {
        request = NewRecoverMalwaresRequest()
    }
    response = NewRecoverMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewRenewProVersionRequest() (request *RenewProVersionRequest) {
    request = &RenewProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "RenewProVersion")
    return
}

func NewRenewProVersionResponse() (response *RenewProVersionResponse) {
    response = &RenewProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewProVersion
// 本接口 (RenewProVersion) 用于续费专业版(包年包月)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewProVersion(request *RenewProVersionRequest) (response *RenewProVersionResponse, err error) {
    if request == nil {
        request = NewRenewProVersionRequest()
    }
    response = NewRenewProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRescanImpactedHostRequest() (request *RescanImpactedHostRequest) {
    request = &RescanImpactedHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "RescanImpactedHost")
    return
}

func NewRescanImpactedHostResponse() (response *RescanImpactedHostResponse) {
    response = &RescanImpactedHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RescanImpactedHost
// 该接口已废弃
//
// 
//
// 本接口 (RescanImpactedHost) 用于漏洞重新检测。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"
//  FAILEDOPERATION_RESCANVULPROCESSINUSE = "FailedOperation.RescanVulProcessInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RescanImpactedHost(request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    if request == nil {
        request = NewRescanImpactedHostRequest()
    }
    response = NewRescanImpactedHostResponse()
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
    if request == nil {
        request = NewScanAssetRequest()
    }
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
    if request == nil {
        request = NewScanVulRequest()
    }
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanVulAgain(request *ScanVulAgainRequest) (response *ScanVulAgainResponse, err error) {
    if request == nil {
        request = NewScanVulAgainRequest()
    }
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
    if request == nil {
        request = NewScanVulSettingRequest()
    }
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
    if request == nil {
        request = NewSeparateMalwaresRequest()
    }
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
    if request == nil {
        request = NewSetBashEventsStatusRequest()
    }
    response = NewSetBashEventsStatusResponse()
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
    if request == nil {
        request = NewStopNoticeBanTipsRequest()
    }
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
    if request == nil {
        request = NewSwitchBashRulesRequest()
    }
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
func (c *Client) SyncAssetScan(request *SyncAssetScanRequest) (response *SyncAssetScanResponse, err error) {
    if request == nil {
        request = NewSyncAssetScanRequest()
    }
    response = NewSyncAssetScanResponse()
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
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
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
    if request == nil {
        request = NewUntrustMalwaresRequest()
    }
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
    if request == nil {
        request = NewUpdateBaselineStrategyRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateMachineTags(request *UpdateMachineTagsRequest) (response *UpdateMachineTagsResponse, err error) {
    if request == nil {
        request = NewUpdateMachineTagsRequest()
    }
    response = NewUpdateMachineTagsResponse()
    err = c.Send(request, response)
    return
}
