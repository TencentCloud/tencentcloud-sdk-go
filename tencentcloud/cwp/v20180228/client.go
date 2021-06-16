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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddLoginWhiteListRequest() (request *AddLoginWhiteListRequest) {
    request = &AddLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "AddLoginWhiteList")
    return
}

func NewAddLoginWhiteListResponse() (response *AddLoginWhiteListResponse) {
    response = &AddLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLoginWhiteList
// 本接口用于新增异地登录白名单规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddLoginWhiteList(request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListRequest()
    }
    response = NewAddLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddMachineTagRequest() (request *AddMachineTagRequest) {
    request = &AddMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "AddMachineTag")
    return
}

func NewAddMachineTagResponse() (response *AddMachineTagResponse) {
    response = &AddMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddMachineTag
// 增加机器关联标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddMachineTag(request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    if request == nil {
        request = NewAddMachineTagRequest()
    }
    response = NewAddMachineTagResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBaselineStrategy(request *CreateBaselineStrategyRequest) (response *CreateBaselineStrategyResponse, err error) {
    if request == nil {
        request = NewCreateBaselineStrategyRequest()
    }
    response = NewCreateBaselineStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenPortTaskRequest() (request *CreateOpenPortTaskRequest) {
    request = &CreateOpenPortTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "CreateOpenPortTask")
    return
}

func NewCreateOpenPortTaskResponse() (response *CreateOpenPortTaskResponse) {
    response = &CreateOpenPortTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOpenPortTask
// 本接口 (CreateOpenPortTask) 用于创建实时获取端口任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEOPENPORTTASK = "FailedOperation.CreateOpenPortTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOpenPortTask(request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    if request == nil {
        request = NewCreateOpenPortTaskRequest()
    }
    response = NewCreateOpenPortTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessTaskRequest() (request *CreateProcessTaskRequest) {
    request = &CreateProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "CreateProcessTask")
    return
}

func NewCreateProcessTaskResponse() (response *CreateProcessTaskResponse) {
    response = &CreateProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProcessTask
// 本接口 (CreateProcessTask) 用于创建实时拉取进程任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPROCESSTASK = "FailedOperation.CreateProcessTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProcessTask(request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateProcessTaskRequest()
    }
    response = NewCreateProcessTaskResponse()
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
// 添加房展防护服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CreateProtectServer(request *CreateProtectServerRequest) (response *CreateProtectServerResponse, err error) {
    if request == nil {
        request = NewCreateProtectServerRequest()
    }
    response = NewCreateProtectServerResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_NAMEHASREPETITION = "InvalidParameter.NameHasRepetition"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateSearchTemplate(request *CreateSearchTemplateRequest) (response *CreateSearchTemplateResponse, err error) {
    if request == nil {
        request = NewCreateSearchTemplateRequest()
    }
    response = NewCreateSearchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsualLoginPlacesRequest() (request *CreateUsualLoginPlacesRequest) {
    request = &CreateUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "CreateUsualLoginPlaces")
    return
}

func NewCreateUsualLoginPlacesResponse() (response *CreateUsualLoginPlacesResponse) {
    response = &CreateUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUsualLoginPlaces
// 此接口（CreateUsualLoginPlaces）用于添加常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUsualLoginPlaces(request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewCreateUsualLoginPlacesRequest()
    }
    response = NewCreateUsualLoginPlacesResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
func (c *Client) DeleteAttackLogs(request *DeleteAttackLogsRequest) (response *DeleteAttackLogsResponse, err error) {
    if request == nil {
        request = NewDeleteAttackLogsRequest()
    }
    response = NewDeleteAttackLogsResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
    response = NewDeleteMaliciousRequestsResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
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

func NewDeleteUsualLoginPlacesRequest() (request *DeleteUsualLoginPlacesRequest) {
    request = &DeleteUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DeleteUsualLoginPlaces")
    return
}

func NewDeleteUsualLoginPlacesResponse() (response *DeleteUsualLoginPlacesResponse) {
    response = &DeleteUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsualLoginPlaces
// 本接口（DeleteUsualLoginPlaces）用于删除常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteUsualLoginPlaces(request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteUsualLoginPlacesRequest()
    }
    response = NewDeleteUsualLoginPlacesResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDescribeAgentVulsRequest() (request *DescribeAgentVulsRequest) {
    request = &DescribeAgentVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAgentVuls")
    return
}

func NewDescribeAgentVulsResponse() (response *DescribeAgentVulsResponse) {
    response = &DescribeAgentVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentVuls
// 本接口 (DescribeAgentVuls) 用于获取单台主机的漏洞列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentVuls(request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVulsRequest()
    }
    response = NewDescribeAgentVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
    request = &DescribeAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeAlarmAttribute")
    return
}

func NewDescribeAlarmAttributeResponse() (response *DescribeAlarmAttributeResponse) {
    response = &DescribeAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmAttribute
// 本接口 (DescribeAlarmAttribute) 用于获取告警设置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAlarmAttribute(request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmAttributeRequest()
    }
    response = NewDescribeAlarmAttributeResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeAssetInfo(request *DescribeAssetInfoRequest) (response *DescribeAssetInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetInfoRequest()
    }
    response = NewDescribeAssetInfoResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeAssetRecentMachineInfo(request *DescribeAssetRecentMachineInfoRequest) (response *DescribeAssetRecentMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRecentMachineInfoRequest()
    }
    response = NewDescribeAssetRecentMachineInfoResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeAttackVulTypeList(request *DescribeAttackVulTypeListRequest) (response *DescribeAttackVulTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeAttackVulTypeListRequest()
    }
    response = NewDescribeAttackVulTypeListResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeBruteAttackList(request *DescribeBruteAttackListRequest) (response *DescribeBruteAttackListResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttackListRequest()
    }
    response = NewDescribeBruteAttackListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttacksRequest() (request *DescribeBruteAttacksRequest) {
    request = &DescribeBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeBruteAttacks")
    return
}

func NewDescribeBruteAttacksResponse() (response *DescribeBruteAttacksResponse) {
    response = &DescribeBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBruteAttacks
// 本接口{DescribeBruteAttacks}用于获取暴力破解事件列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeBruteAttacks(request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttacksRequest()
    }
    response = NewDescribeBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentInfoRequest() (request *DescribeComponentInfoRequest) {
    request = &DescribeComponentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeComponentInfo")
    return
}

func NewDescribeComponentInfoResponse() (response *DescribeComponentInfoResponse) {
    response = &DescribeComponentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponentInfo
// 本接口 (DescribeComponentInfo) 用于获取组件信息数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeComponentInfo(request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComponentInfoRequest()
    }
    response = NewDescribeComponentInfoResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
    response = NewDescribeComponentStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
    request = &DescribeComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeComponents")
    return
}

func NewDescribeComponentsResponse() (response *DescribeComponentsResponse) {
    response = &DescribeComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComponents
// 本接口 (DescribeComponents) 用于获取组件列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentsRequest()
    }
    response = NewDescribeComponentsResponse()
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
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeESHits(request *DescribeESHitsRequest) (response *DescribeESHitsResponse, err error) {
    if request == nil {
        request = NewDescribeESHitsRequest()
    }
    response = NewDescribeESHitsResponse()
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
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
// 获取主机安全相关统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
func (c *Client) DescribeHistoryService(request *DescribeHistoryServiceRequest) (response *DescribeHistoryServiceResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryServiceRequest()
    }
    response = NewDescribeHistoryServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImpactedHostsRequest() (request *DescribeImpactedHostsRequest) {
    request = &DescribeImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeImpactedHosts")
    return
}

func NewDescribeImpactedHostsResponse() (response *DescribeImpactedHostsResponse) {
    response = &DescribeImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImpactedHosts
// 本接口 (DescribeImpactedHosts) 用于获取漏洞受影响机器列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImpactedHosts(request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    if request == nil {
        request = NewDescribeImpactedHostsRequest()
    }
    response = NewDescribeImpactedHostsResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeLogStorageStatistic(request *DescribeLogStorageStatisticRequest) (response *DescribeLogStorageStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeLogStorageStatisticRequest()
    }
    response = NewDescribeLogStorageStatisticResponse()
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
// 本接口 (DescribeMachineList) 用于网页防篡改获取区域主机列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
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
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaliciousRequestsRequest() (request *DescribeMaliciousRequestsRequest) {
    request = &DescribeMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMaliciousRequests")
    return
}

func NewDescribeMaliciousRequestsResponse() (response *DescribeMaliciousRequestsResponse) {
    response = &DescribeMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaliciousRequests
// 本接口 (DescribeMaliciousRequests) 用于获取恶意请求数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeMaliciousRequests(request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestsRequest()
    }
    response = NewDescribeMaliciousRequestsResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDescribeMalwaresRequest() (request *DescribeMalwaresRequest) {
    request = &DescribeMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeMalwares")
    return
}

func NewDescribeMalwaresResponse() (response *DescribeMalwaresResponse) {
    response = &DescribeMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMalwares
// 本接口（DescribeMalwares）用于获取木马事件列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeMalwares(request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeMalwaresRequest()
    }
    response = NewDescribeMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNonlocalLoginPlacesRequest() (request *DescribeNonlocalLoginPlacesRequest) {
    request = &DescribeNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeNonlocalLoginPlaces")
    return
}

func NewDescribeNonlocalLoginPlacesResponse() (response *DescribeNonlocalLoginPlacesResponse) {
    response = &DescribeNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNonlocalLoginPlaces
// 本接口(DescribeNonlocalLoginPlaces)用于获取异地登录事件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeNonlocalLoginPlaces(request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeNonlocalLoginPlacesRequest()
    }
    response = NewDescribeNonlocalLoginPlacesResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
    response = NewDescribeOpenPortStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortTaskStatusRequest() (request *DescribeOpenPortTaskStatusRequest) {
    request = &DescribeOpenPortTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeOpenPortTaskStatus")
    return
}

func NewDescribeOpenPortTaskStatusResponse() (response *DescribeOpenPortTaskStatusResponse) {
    response = &DescribeOpenPortTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPortTaskStatus
// 本接口 (DescribeOpenPortTaskStatus) 用于获取实时拉取端口任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"
//  FAILEDOPERATION_OPENPORTTASKNOTFOUND = "FailedOperation.OpenPortTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOpenPortTaskStatus(request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortTaskStatusRequest()
    }
    response = NewDescribeOpenPortTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortsRequest() (request *DescribeOpenPortsRequest) {
    request = &DescribeOpenPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeOpenPorts")
    return
}

func NewDescribeOpenPortsResponse() (response *DescribeOpenPortsResponse) {
    response = &DescribeOpenPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOpenPorts
// 本接口 (DescribeOpenPorts) 用于获取端口列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPorts(request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortsRequest()
    }
    response = NewDescribeOpenPortsResponse()
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
// 本接口用于（DescribeOverviewStatistics）获取概览统计数据。
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
// 本接口 (DescribeProVersionInfo) 用于获取专业版信息。
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
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
    response = NewDescribeProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessTaskStatusRequest() (request *DescribeProcessTaskStatusRequest) {
    request = &DescribeProcessTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProcessTaskStatus")
    return
}

func NewDescribeProcessTaskStatusResponse() (response *DescribeProcessTaskStatusResponse) {
    response = &DescribeProcessTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcessTaskStatus
// 本接口 (DescribeProcessTaskStatus) 用于获取实时拉取进程任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PROCESSTASKNOTFOUND = "FailedOperation.ProcessTaskNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProcessTaskStatus(request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProcessTaskStatusRequest()
    }
    response = NewDescribeProcessTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeProcesses")
    return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
    response = &DescribeProcessesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProcesses
// 本接口 (DescribeProcesses) 用于获取进程列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    response = NewDescribeProcessesResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeRiskDnsList(request *DescribeRiskDnsListRequest) (response *DescribeRiskDnsListResponse, err error) {
    if request == nil {
        request = NewDescribeRiskDnsListRequest()
    }
    response = NewDescribeRiskDnsListResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScanMalwareSchedule(request *DescribeScanMalwareScheduleRequest) (response *DescribeScanMalwareScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeScanMalwareScheduleRequest()
    }
    response = NewDescribeScanMalwareScheduleResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
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
// 本接口 (DescribeSecurityDynamics) 用于获取安全事件消息数据。
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
// 概览页抽屉侧边弹窗：安全概览“立即处理”页面中的相关事件数统计接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
    response = NewDescribeSecurityTrendsResponse()
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
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
    response = NewDescribeUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulInfoRequest() (request *DescribeVulInfoRequest) {
    request = &DescribeVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulInfo")
    return
}

func NewDescribeVulInfoResponse() (response *DescribeVulInfoResponse) {
    response = &DescribeVulInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulInfo
// 本接口 (DescribeVulInfo) 用于获取漏洞详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVulInfo(request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoRequest()
    }
    response = NewDescribeVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanResultRequest() (request *DescribeVulScanResultRequest) {
    request = &DescribeVulScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVulScanResult")
    return
}

func NewDescribeVulScanResultResponse() (response *DescribeVulScanResultResponse) {
    response = &DescribeVulScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulScanResult
// 本接口 (DescribeVulScanResult) 用于获取漏洞检测结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulScanResult(request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanResultRequest()
    }
    response = NewDescribeVulScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeVuls")
    return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
    response = &DescribeVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVuls
// 本接口 (DescribeVuls) 用于获取漏洞列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    response = NewDescribeVulsResponse()
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
// 查询网站防篡改 概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
func (c *Client) DescribeWebPageGeneralize(request *DescribeWebPageGeneralizeRequest) (response *DescribeWebPageGeneralizeResponse, err error) {
    if request == nil {
        request = NewDescribeWebPageGeneralizeRequest()
    }
    response = NewDescribeWebPageGeneralizeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportBruteAttacksRequest() (request *DescribeWeeklyReportBruteAttacksRequest) {
    request = &DescribeWeeklyReportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReportBruteAttacks")
    return
}

func NewDescribeWeeklyReportBruteAttacksResponse() (response *DescribeWeeklyReportBruteAttacksResponse) {
    response = &DescribeWeeklyReportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportBruteAttacks
// 本接口 (DescribeWeeklyReportBruteAttacks) 用于获取专业周报密码破解数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportBruteAttacks(request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportBruteAttacksRequest()
    }
    response = NewDescribeWeeklyReportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportInfoRequest() (request *DescribeWeeklyReportInfoRequest) {
    request = &DescribeWeeklyReportInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReportInfo")
    return
}

func NewDescribeWeeklyReportInfoResponse() (response *DescribeWeeklyReportInfoResponse) {
    response = &DescribeWeeklyReportInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportInfo
// 本接口 (DescribeWeeklyReportInfo) 用于获取专业周报详情数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportInfo(request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportInfoRequest()
    }
    response = NewDescribeWeeklyReportInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportMalwaresRequest() (request *DescribeWeeklyReportMalwaresRequest) {
    request = &DescribeWeeklyReportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReportMalwares")
    return
}

func NewDescribeWeeklyReportMalwaresResponse() (response *DescribeWeeklyReportMalwaresResponse) {
    response = &DescribeWeeklyReportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportMalwares
// 本接口 (DescribeWeeklyReportMalwares) 用于获取专业周报木马数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportMalwares(request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportMalwaresRequest()
    }
    response = NewDescribeWeeklyReportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesRequest() (request *DescribeWeeklyReportNonlocalLoginPlacesRequest) {
    request = &DescribeWeeklyReportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReportNonlocalLoginPlaces")
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesResponse() (response *DescribeWeeklyReportNonlocalLoginPlacesResponse) {
    response = &DescribeWeeklyReportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportNonlocalLoginPlaces
// 本接口 (DescribeWeeklyReportNonlocalLoginPlaces) 用于获取专业周报异地登录数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportNonlocalLoginPlaces(request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportNonlocalLoginPlacesRequest()
    }
    response = NewDescribeWeeklyReportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportVulsRequest() (request *DescribeWeeklyReportVulsRequest) {
    request = &DescribeWeeklyReportVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReportVuls")
    return
}

func NewDescribeWeeklyReportVulsResponse() (response *DescribeWeeklyReportVulsResponse) {
    response = &DescribeWeeklyReportVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReportVuls
// 本接口 (DescribeWeeklyReportVuls) 用于专业版周报漏洞数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportVuls(request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportVulsRequest()
    }
    response = NewDescribeWeeklyReportVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportsRequest() (request *DescribeWeeklyReportsRequest) {
    request = &DescribeWeeklyReportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "DescribeWeeklyReports")
    return
}

func NewDescribeWeeklyReportsResponse() (response *DescribeWeeklyReportsResponse) {
    response = &DescribeWeeklyReportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWeeklyReports
// 本接口 (DescribeWeeklyReports) 用于获取周报列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReports(request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportsRequest()
    }
    response = NewDescribeWeeklyReportsResponse()
    err = c.Send(request, response)
    return
}

func NewEditBashRuleRequest() (request *EditBashRuleRequest) {
    request = &EditBashRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "EditBashRule")
    return
}

func NewEditBashRuleResponse() (response *EditBashRuleResponse) {
    response = &EditBashRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditBashRule
// 新增或修改高危命令规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditBashRule(request *EditBashRuleRequest) (response *EditBashRuleResponse, err error) {
    if request == nil {
        request = NewEditBashRuleRequest()
    }
    response = NewEditBashRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditPrivilegeRuleRequest() (request *EditPrivilegeRuleRequest) {
    request = &EditPrivilegeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "EditPrivilegeRule")
    return
}

func NewEditPrivilegeRuleResponse() (response *EditPrivilegeRuleResponse) {
    response = &EditPrivilegeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditPrivilegeRule
// 新增或修改本地提权规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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
func (c *Client) EditPrivilegeRule(request *EditPrivilegeRuleRequest) (response *EditPrivilegeRuleResponse, err error) {
    if request == nil {
        request = NewEditPrivilegeRuleRequest()
    }
    response = NewEditPrivilegeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditReverseShellRuleRequest() (request *EditReverseShellRuleRequest) {
    request = &EditReverseShellRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "EditReverseShellRule")
    return
}

func NewEditReverseShellRuleResponse() (response *EditReverseShellRuleResponse) {
    response = &EditReverseShellRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditReverseShellRule
// 编辑反弹Shell规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
//  INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"
//  INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditReverseShellRule(request *EditReverseShellRuleRequest) (response *EditReverseShellRuleResponse, err error) {
    if request == nil {
        request = NewEditReverseShellRuleRequest()
    }
    response = NewEditReverseShellRuleResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetCoreModuleList(request *ExportAssetCoreModuleListRequest) (response *ExportAssetCoreModuleListResponse, err error) {
    if request == nil {
        request = NewExportAssetCoreModuleListRequest()
    }
    response = NewExportAssetCoreModuleListResponse()
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
func (c *Client) ExportAttackLogs(request *ExportAttackLogsRequest) (response *ExportAttackLogsResponse, err error) {
    if request == nil {
        request = NewExportAttackLogsRequest()
    }
    response = NewExportAttackLogsResponse()
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
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
    response = NewExportBruteAttacksResponse()
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
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
func (c *Client) ExportPrivilegeEvents(request *ExportPrivilegeEventsRequest) (response *ExportPrivilegeEventsResponse, err error) {
    if request == nil {
        request = NewExportPrivilegeEventsRequest()
    }
    response = NewExportPrivilegeEventsResponse()
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ExportVulDetectionReport(request *ExportVulDetectionReportRequest) (response *ExportVulDetectionReportResponse, err error) {
    if request == nil {
        request = NewExportVulDetectionReportRequest()
    }
    response = NewExportVulDetectionReportResponse()
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewMisAlarmNonlocalLoginPlacesRequest() (request *MisAlarmNonlocalLoginPlacesRequest) {
    request = &MisAlarmNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "MisAlarmNonlocalLoginPlaces")
    return
}

func NewMisAlarmNonlocalLoginPlacesResponse() (response *MisAlarmNonlocalLoginPlacesResponse) {
    response = &MisAlarmNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MisAlarmNonlocalLoginPlaces
// 本接口{MisAlarmNonlocalLoginPlaces}将设置当前地点为常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) MisAlarmNonlocalLoginPlaces(request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewMisAlarmNonlocalLoginPlacesRequest()
    }
    response = NewMisAlarmNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmAttributeRequest() (request *ModifyAlarmAttributeRequest) {
    request = &ModifyAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyAlarmAttribute")
    return
}

func NewModifyAlarmAttributeResponse() (response *ModifyAlarmAttributeResponse) {
    response = &ModifyAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmAttribute
// 本接口（ModifyAlarmAttribute）用于修改告警设置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAlarmAttribute(request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmAttributeRequest()
    }
    response = NewModifyAlarmAttributeResponse()
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
// 本接口 (ModifyAutoOpenProVersionConfig) 用于设置新增主机自动开通专业版配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
    response = NewModifyAutoOpenProVersionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteListRequest() (request *ModifyLoginWhiteListRequest) {
    request = &ModifyLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "ModifyLoginWhiteList")
    return
}

func NewModifyLoginWhiteListResponse() (response *ModifyLoginWhiteListResponse) {
    response = &ModifyLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoginWhiteList
// 本接口用于编辑异地登录白名单规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyLoginWhiteList(request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteListRequest()
    }
    response = NewModifyLoginWhiteListResponse()
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
//  FAILEDOPERATION_APISERVERFAIL = "FailedOperation.APIServerFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    response = NewModifyProVersionRenewFlagResponse()
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
// 网站防篡改-修改网站防护设置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) ModifyWebPageProtectSetting(request *ModifyWebPageProtectSettingRequest) (response *ModifyWebPageProtectSettingResponse, err error) {
    if request == nil {
        request = NewModifyWebPageProtectSettingRequest()
    }
    response = NewModifyWebPageProtectSettingResponse()
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
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
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
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
//  FAILEDOPERATION_RECOVER = "FailedOperation.Recover"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR = "InternalError"
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
// 漏洞管理 - 一键检测
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
func (c *Client) ScanVul(request *ScanVulRequest) (response *ScanVulResponse, err error) {
    if request == nil {
        request = NewScanVulRequest()
    }
    response = NewScanVulResponse()
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"
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

func NewTrustMaliciousRequestRequest() (request *TrustMaliciousRequestRequest) {
    request = &TrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "TrustMaliciousRequest")
    return
}

func NewTrustMaliciousRequestResponse() (response *TrustMaliciousRequestResponse) {
    response = &TrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TrustMaliciousRequest
// 本接口 (TrustMaliciousRequest) 用于恶意请求添加信任。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) TrustMaliciousRequest(request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewTrustMaliciousRequestRequest()
    }
    response = NewTrustMaliciousRequestResponse()
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
    response = NewTrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMaliciousRequestRequest() (request *UntrustMaliciousRequestRequest) {
    request = &UntrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cwp", APIVersion, "UntrustMaliciousRequest")
    return
}

func NewUntrustMaliciousRequestResponse() (response *UntrustMaliciousRequestResponse) {
    response = &UntrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UntrustMaliciousRequest
// 本接口 (UntrustMaliciousRequest) 用于取消信任恶意请求。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMaliciousRequest(request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewUntrustMaliciousRequestRequest()
    }
    response = NewUntrustMaliciousRequestResponse()
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) UpdateBaselineStrategy(request *UpdateBaselineStrategyRequest) (response *UpdateBaselineStrategyResponse, err error) {
    if request == nil {
        request = NewUpdateBaselineStrategyRequest()
    }
    response = NewUpdateBaselineStrategyResponse()
    err = c.Send(request, response)
    return
}
