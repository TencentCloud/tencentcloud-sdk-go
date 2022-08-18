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


func NewAddLoginWhiteListRequest() (request *AddLoginWhiteListRequest) {
    request = &AddLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "AddLoginWhiteList")
    
    
    return
}

func NewAddLoginWhiteListResponse() (response *AddLoginWhiteListResponse) {
    response = &AddLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLoginWhiteList
// 本接口（AddLoginWhiteList）用于添加白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) AddLoginWhiteList(request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    return c.AddLoginWhiteListWithContext(context.Background(), request)
}

// AddLoginWhiteList
// 本接口（AddLoginWhiteList）用于添加白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) AddLoginWhiteListWithContext(ctx context.Context, request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddMachineTagRequest() (request *AddMachineTagRequest) {
    request = &AddMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "AddMachineTag")
    
    
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
    return c.AddMachineTagWithContext(context.Background(), request)
}

// AddMachineTag
// 增加机器关联标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddMachineTagWithContext(ctx context.Context, request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    if request == nil {
        request = NewAddMachineTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMachineTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
    request = &CloseProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "CloseProVersion")
    
    
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
    return c.CloseProVersionWithContext(context.Background(), request)
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
func (c *Client) CloseProVersionWithContext(ctx context.Context, request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    if request == nil {
        request = NewCloseProVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBaselineStrategyRequest() (request *CreateBaselineStrategyRequest) {
    request = &CreateBaselineStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateBaselineStrategy")
    
    
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
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) CreateBaselineStrategy(request *CreateBaselineStrategyRequest) (response *CreateBaselineStrategyResponse, err error) {
    return c.CreateBaselineStrategyWithContext(context.Background(), request)
}

// CreateBaselineStrategy
// 根据策略信息创建基线策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"
//  FAILEDOPERATION_TOOMANYSTRATEGY = "FailedOperation.TooManyStrategy"
//  INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewCreateOpenPortTaskRequest() (request *CreateOpenPortTaskRequest) {
    request = &CreateOpenPortTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateOpenPortTask")
    
    
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
    return c.CreateOpenPortTaskWithContext(context.Background(), request)
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
func (c *Client) CreateOpenPortTaskWithContext(ctx context.Context, request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    if request == nil {
        request = NewCreateOpenPortTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpenPortTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpenPortTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessTaskRequest() (request *CreateProcessTaskRequest) {
    request = &CreateProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateProcessTask")
    
    
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
    return c.CreateProcessTaskWithContext(context.Background(), request)
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
func (c *Client) CreateProcessTaskWithContext(ctx context.Context, request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateProcessTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProcessTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsualLoginPlacesRequest() (request *CreateUsualLoginPlacesRequest) {
    request = &CreateUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateUsualLoginPlaces")
    
    
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
    return c.CreateUsualLoginPlacesWithContext(context.Background(), request)
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
func (c *Client) CreateUsualLoginPlacesWithContext(ctx context.Context, request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewCreateUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttackLogsRequest() (request *DeleteAttackLogsRequest) {
    request = &DeleteAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteAttackLogs")
    
    
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
    return c.DeleteAttackLogsWithContext(context.Background(), request)
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

func NewDeleteBashEventsRequest() (request *DeleteBashEventsRequest) {
    request = &DeleteBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBashEvents")
    
    
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
    return c.DeleteBashEventsWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBashRules")
    
    
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBruteAttacks")
    
    
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
    return c.DeleteBruteAttacksWithContext(context.Background(), request)
}

// DeleteBruteAttacks
// 本接口 (DeleteBruteAttacks) 用于删除暴力破解记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
    request = &DeleteLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteLoginWhiteList")
    
    
    return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
    response = &DeleteLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoginWhiteList
// 删除白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    return c.DeleteLoginWhiteListWithContext(context.Background(), request)
}

// DeleteLoginWhiteList
// 删除白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachine")
    
    
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
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    return c.DeleteMachineWithContext(context.Background(), request)
}

// DeleteMachine
// 本接口（DeleteMachine）用于卸载云镜客户端。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachineTag")
    
    
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
    return c.DeleteMachineTagWithContext(context.Background(), request)
}

// DeleteMachineTag
// 删除服务器关联的标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMaliciousRequests")
    
    
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
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    return c.DeleteMaliciousRequestsWithContext(context.Background(), request)
}

// DeleteMaliciousRequests
// 本接口 (DeleteMaliciousRequests) 用于删除恶意请求记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
    request = &DeleteMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMalwares")
    
    
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteNonlocalLoginPlaces")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    return c.DeleteNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DeleteNonlocalLoginPlaces
// 本接口 (DeleteNonlocalLoginPlaces) 用于删除异地登录记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeletePrivilegeEvents")
    
    
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
    return c.DeletePrivilegeEventsWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeletePrivilegeRules")
    
    
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
    return c.DeletePrivilegeRulesWithContext(context.Background(), request)
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

func NewDeleteReverseShellEventsRequest() (request *DeleteReverseShellEventsRequest) {
    request = &DeleteReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteReverseShellEvents")
    
    
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
    return c.DeleteReverseShellEventsWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteReverseShellRules")
    
    
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
    return c.DeleteReverseShellRulesWithContext(context.Background(), request)
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

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
    request = &DeleteTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteTags")
    
    
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
    return c.DeleteTagsWithContext(context.Background(), request)
}

// DeleteTags
// 删除标签
//
// 可能返回的错误码:
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

func NewDeleteUsualLoginPlacesRequest() (request *DeleteUsualLoginPlacesRequest) {
    request = &DeleteUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteUsualLoginPlaces")
    
    
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
    return c.DeleteUsualLoginPlacesWithContext(context.Background(), request)
}

// DeleteUsualLoginPlaces
// 本接口（DeleteUsualLoginPlaces）用于删除常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteUsualLoginPlacesWithContext(ctx context.Context, request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteUsualLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsualLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
    request = &DescribeAccountStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccountStatistics")
    
    
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
    return c.DescribeAccountStatisticsWithContext(context.Background(), request)
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

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccounts")
    
    
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
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口 (DescribeAccounts) 用于获取帐号列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVulsRequest() (request *DescribeAgentVulsRequest) {
    request = &DescribeAgentVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAgentVuls")
    
    
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
    return c.DescribeAgentVulsWithContext(context.Background(), request)
}

// DescribeAgentVuls
// 本接口 (DescribeAgentVuls) 用于获取单台主机的漏洞列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAgentVulsWithContext(ctx context.Context, request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
    request = &DescribeAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAlarmAttribute")
    
    
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
    return c.DescribeAlarmAttributeWithContext(context.Background(), request)
}

// DescribeAlarmAttribute
// 本接口 (DescribeAlarmAttribute) 用于获取告警设置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAlarmAttributeWithContext(ctx context.Context, request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttackLogInfoRequest() (request *DescribeAttackLogInfoRequest) {
    request = &DescribeAttackLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAttackLogInfo")
    
    
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
func (c *Client) DescribeAttackLogInfo(request *DescribeAttackLogInfoRequest) (response *DescribeAttackLogInfoResponse, err error) {
    return c.DescribeAttackLogInfoWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAttackLogs")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAttackLogs(request *DescribeAttackLogsRequest) (response *DescribeAttackLogsResponse, err error) {
    return c.DescribeAttackLogsWithContext(context.Background(), request)
}

// DescribeAttackLogs
// 按分页形式展示网络攻击日志列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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

func NewDescribeBashEventsRequest() (request *DescribeBashEventsRequest) {
    request = &DescribeBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBashEvents")
    
    
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDescribeBashRulesRequest() (request *DescribeBashRulesRequest) {
    request = &DescribeBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBashRules")
    
    
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
    return c.DescribeBashRulesWithContext(context.Background(), request)
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

func NewDescribeBruteAttacksRequest() (request *DescribeBruteAttacksRequest) {
    request = &DescribeBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBruteAttacks")
    
    
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
    return c.DescribeBruteAttacksWithContext(context.Background(), request)
}

// DescribeBruteAttacks
// 本接口{DescribeBruteAttacks}用于获取暴力破解事件列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeBruteAttacksWithContext(ctx context.Context, request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentInfoRequest() (request *DescribeComponentInfoRequest) {
    request = &DescribeComponentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentInfo")
    
    
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
    return c.DescribeComponentInfoWithContext(context.Background(), request)
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
func (c *Client) DescribeComponentInfoWithContext(ctx context.Context, request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComponentInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponentInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
    request = &DescribeComponentStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentStatistics")
    
    
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
    return c.DescribeComponentStatisticsWithContext(context.Background(), request)
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

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
    request = &DescribeComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponents")
    
    
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
    return c.DescribeComponentsWithContext(context.Background(), request)
}

// DescribeComponents
// 本接口 (DescribeComponents) 用于获取组件列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeComponentsWithContext(ctx context.Context, request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComponents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
    request = &DescribeHistoryAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeHistoryAccounts")
    
    
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
    return c.DescribeHistoryAccountsWithContext(context.Background(), request)
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

func NewDescribeImpactedHostsRequest() (request *DescribeImpactedHostsRequest) {
    request = &DescribeImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeImpactedHosts")
    
    
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
    return c.DescribeImpactedHostsWithContext(context.Background(), request)
}

// DescribeImpactedHosts
// 本接口 (DescribeImpactedHosts) 用于获取漏洞受影响机器列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImpactedHostsWithContext(ctx context.Context, request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    if request == nil {
        request = NewDescribeImpactedHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImpactedHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
    request = &DescribeLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeLoginWhiteList")
    
    
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
    return c.DescribeLoginWhiteListWithContext(context.Background(), request)
}

// DescribeLoginWhiteList
// 获取异地登录白名单列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachineInfo")
    
    
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

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachines")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    return c.DescribeMachinesWithContext(context.Background(), request)
}

// DescribeMachines
// 本接口 (DescribeMachines) 用于获取区域主机列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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

func NewDescribeMaliciousRequestsRequest() (request *DescribeMaliciousRequestsRequest) {
    request = &DescribeMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMaliciousRequests")
    
    
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
    return c.DescribeMaliciousRequestsWithContext(context.Background(), request)
}

// DescribeMaliciousRequests
// 本接口 (DescribeMaliciousRequests) 用于获取恶意请求数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeMaliciousRequestsWithContext(ctx context.Context, request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaliciousRequests require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwaresRequest() (request *DescribeMalwaresRequest) {
    request = &DescribeMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMalwares")
    
    
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
    return c.DescribeMalwaresWithContext(context.Background(), request)
}

// DescribeMalwares
// 本接口（DescribeMalwares）用于获取木马事件列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeMalwaresWithContext(ctx context.Context, request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNonlocalLoginPlacesRequest() (request *DescribeNonlocalLoginPlacesRequest) {
    request = &DescribeNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeNonlocalLoginPlaces")
    
    
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
    return c.DescribeNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DescribeNonlocalLoginPlaces
// 本接口(DescribeNonlocalLoginPlaces)用于获取异地登录事件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) DescribeNonlocalLoginPlacesWithContext(ctx context.Context, request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
    request = &DescribeOpenPortStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortStatistics")
    
    
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
    return c.DescribeOpenPortStatisticsWithContext(context.Background(), request)
}

// DescribeOpenPortStatistics
// 本接口 (DescribeOpenPortStatistics) 用于获取端口统计列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
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

func NewDescribeOpenPortTaskStatusRequest() (request *DescribeOpenPortTaskStatusRequest) {
    request = &DescribeOpenPortTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortTaskStatus")
    
    
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
    return c.DescribeOpenPortTaskStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeOpenPortTaskStatusWithContext(ctx context.Context, request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPortTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortsRequest() (request *DescribeOpenPortsRequest) {
    request = &DescribeOpenPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPorts")
    
    
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
    return c.DescribeOpenPortsWithContext(context.Background(), request)
}

// DescribeOpenPorts
// 本接口 (DescribeOpenPorts) 用于获取端口列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOpenPortsWithContext(ctx context.Context, request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpenPorts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpenPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
    request = &DescribeOverviewStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOverviewStatistics")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    return c.DescribeOverviewStatisticsWithContext(context.Background(), request)
}

// DescribeOverviewStatistics
// 本接口用于（DescribeOverviewStatistics）获取概览统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribePrivilegeEvents")
    
    
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
    return c.DescribePrivilegeEventsWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribePrivilegeRules")
    
    
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
    return c.DescribePrivilegeRulesWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProVersionInfo")
    
    
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
    return c.DescribeProVersionInfoWithContext(context.Background(), request)
}

// DescribeProVersionInfo
// 本接口 (DescribeProVersionInfo) 用于获取专业版信息。
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

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
    request = &DescribeProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessStatistics")
    
    
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

func NewDescribeProcessTaskStatusRequest() (request *DescribeProcessTaskStatusRequest) {
    request = &DescribeProcessTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessTaskStatus")
    
    
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
    return c.DescribeProcessTaskStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeProcessTaskStatusWithContext(ctx context.Context, request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProcessTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcessTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcesses")
    
    
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
    return c.DescribeProcessesWithContext(context.Background(), request)
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
func (c *Client) DescribeProcessesWithContext(ctx context.Context, request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProcesses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProcessesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReverseShellEventsRequest() (request *DescribeReverseShellEventsRequest) {
    request = &DescribeReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeReverseShellEvents")
    
    
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
    return c.DescribeReverseShellEventsWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeReverseShellRules")
    
    
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
    return c.DescribeReverseShellRulesWithContext(context.Background(), request)
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

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
    request = &DescribeSecurityDynamicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityDynamics")
    
    
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
    return c.DescribeSecurityDynamicsWithContext(context.Background(), request)
}

// DescribeSecurityDynamics
// 本接口 (DescribeSecurityDynamics) 用于获取安全事件消息数据。
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

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
    request = &DescribeSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityTrends")
    
    
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
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    return c.DescribeSecurityTrendsWithContext(context.Background(), request)
}

// DescribeSecurityTrends
// 本接口 (DescribeSecurityTrends) 用于获取安全事件统计数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
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

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
    request = &DescribeTagMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTagMachines")
    
    
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
    return c.DescribeTagMachinesWithContext(context.Background(), request)
}

// DescribeTagMachines
// 获取指定标签关联的服务器信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTags")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    return c.DescribeTagsWithContext(context.Background(), request)
}

// DescribeTags
// 获取所有主机标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
    request = &DescribeUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeUsualLoginPlaces")
    
    
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
    return c.DescribeUsualLoginPlacesWithContext(context.Background(), request)
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

func NewDescribeVulInfoRequest() (request *DescribeVulInfoRequest) {
    request = &DescribeVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulInfo")
    
    
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
    return c.DescribeVulInfoWithContext(context.Background(), request)
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
func (c *Client) DescribeVulInfoWithContext(ctx context.Context, request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanResultRequest() (request *DescribeVulScanResultRequest) {
    request = &DescribeVulScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulScanResult")
    
    
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
    return c.DescribeVulScanResultWithContext(context.Background(), request)
}

// DescribeVulScanResult
// 本接口 (DescribeVulScanResult) 用于获取漏洞检测结果。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulScanResultWithContext(ctx context.Context, request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulScanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVuls")
    
    
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
    return c.DescribeVulsWithContext(context.Background(), request)
}

// DescribeVuls
// 本接口 (DescribeVuls) 用于获取漏洞列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVulsWithContext(ctx context.Context, request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportBruteAttacksRequest() (request *DescribeWeeklyReportBruteAttacksRequest) {
    request = &DescribeWeeklyReportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportBruteAttacks")
    
    
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
    return c.DescribeWeeklyReportBruteAttacksWithContext(context.Background(), request)
}

// DescribeWeeklyReportBruteAttacks
// 本接口 (DescribeWeeklyReportBruteAttacks) 用于获取专业周报密码破解数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportBruteAttacksWithContext(ctx context.Context, request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportBruteAttacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportBruteAttacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportInfoRequest() (request *DescribeWeeklyReportInfoRequest) {
    request = &DescribeWeeklyReportInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportInfo")
    
    
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
    return c.DescribeWeeklyReportInfoWithContext(context.Background(), request)
}

// DescribeWeeklyReportInfo
// 本接口 (DescribeWeeklyReportInfo) 用于获取专业周报详情数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportInfoWithContext(ctx context.Context, request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportMalwaresRequest() (request *DescribeWeeklyReportMalwaresRequest) {
    request = &DescribeWeeklyReportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportMalwares")
    
    
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
    return c.DescribeWeeklyReportMalwaresWithContext(context.Background(), request)
}

// DescribeWeeklyReportMalwares
// 本接口 (DescribeWeeklyReportMalwares) 用于获取专业周报木马数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportMalwaresWithContext(ctx context.Context, request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportMalwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportMalwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesRequest() (request *DescribeWeeklyReportNonlocalLoginPlacesRequest) {
    request = &DescribeWeeklyReportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportNonlocalLoginPlaces")
    
    
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
    return c.DescribeWeeklyReportNonlocalLoginPlacesWithContext(context.Background(), request)
}

// DescribeWeeklyReportNonlocalLoginPlaces
// 本接口 (DescribeWeeklyReportNonlocalLoginPlaces) 用于获取专业周报异地登录数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportNonlocalLoginPlacesWithContext(ctx context.Context, request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportVulsRequest() (request *DescribeWeeklyReportVulsRequest) {
    request = &DescribeWeeklyReportVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportVuls")
    
    
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
    return c.DescribeWeeklyReportVulsWithContext(context.Background(), request)
}

// DescribeWeeklyReportVuls
// 本接口 (DescribeWeeklyReportVuls) 用于专业版周报漏洞数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) DescribeWeeklyReportVulsWithContext(ctx context.Context, request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportVulsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReportVuls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportsRequest() (request *DescribeWeeklyReportsRequest) {
    request = &DescribeWeeklyReportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReports")
    
    
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
    return c.DescribeWeeklyReportsWithContext(context.Background(), request)
}

// DescribeWeeklyReports
// 本接口 (DescribeWeeklyReports) 用于获取周报列表数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"
func (c *Client) DescribeWeeklyReportsWithContext(ctx context.Context, request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWeeklyReports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWeeklyReportsResponse()
    err = c.Send(request, response)
    return
}

func NewEditBashRuleRequest() (request *EditBashRuleRequest) {
    request = &EditBashRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "EditBashRule")
    
    
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
    return c.EditBashRuleWithContext(context.Background(), request)
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
func (c *Client) EditBashRuleWithContext(ctx context.Context, request *EditBashRuleRequest) (response *EditBashRuleResponse, err error) {
    if request == nil {
        request = NewEditBashRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditBashRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditBashRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditPrivilegeRuleRequest() (request *EditPrivilegeRuleRequest) {
    request = &EditPrivilegeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "EditPrivilegeRule")
    
    
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
    return c.EditPrivilegeRuleWithContext(context.Background(), request)
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
func (c *Client) EditPrivilegeRuleWithContext(ctx context.Context, request *EditPrivilegeRuleRequest) (response *EditPrivilegeRuleResponse, err error) {
    if request == nil {
        request = NewEditPrivilegeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditPrivilegeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditPrivilegeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditReverseShellRuleRequest() (request *EditReverseShellRuleRequest) {
    request = &EditReverseShellRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "EditReverseShellRule")
    
    
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
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditReverseShellRule(request *EditReverseShellRuleRequest) (response *EditReverseShellRuleResponse, err error) {
    return c.EditReverseShellRuleWithContext(context.Background(), request)
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
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditReverseShellRuleWithContext(ctx context.Context, request *EditReverseShellRuleRequest) (response *EditReverseShellRuleResponse, err error) {
    if request == nil {
        request = NewEditReverseShellRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditReverseShellRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditReverseShellRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
    request = &EditTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "EditTags")
    
    
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
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
    return c.EditTagsWithContext(context.Background(), request)
}

// EditTags
// 新增或编辑标签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"
//  LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewExportAttackLogsRequest() (request *ExportAttackLogsRequest) {
    request = &ExportAttackLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportAttackLogs")
    
    
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
    return c.ExportAttackLogsWithContext(context.Background(), request)
}

// ExportAttackLogs
// 导出网络攻击日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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

func NewExportBashEventsRequest() (request *ExportBashEventsRequest) {
    request = &ExportBashEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBashEvents")
    
    
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
    return c.ExportBashEventsWithContext(context.Background(), request)
}

// ExportBashEvents
// 导出高危命令事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBruteAttacks")
    
    
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
    return c.ExportBruteAttacksWithContext(context.Background(), request)
}

// ExportBruteAttacks
// 本接口 (ExportBruteAttacks) 用于导出密码破解记录成CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
    request = &ExportMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMaliciousRequests")
    
    
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
    return c.ExportMaliciousRequestsWithContext(context.Background(), request)
}

// ExportMaliciousRequests
// 本接口 (ExportMaliciousRequests) 用于导出下载恶意请求文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMalwares")
    
    
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
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    return c.ExportMalwaresWithContext(context.Background(), request)
}

// ExportMalwares
// 本接口 (ExportMalwares) 用于导出木马记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportNonlocalLoginPlaces")
    
    
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
    return c.ExportNonlocalLoginPlacesWithContext(context.Background(), request)
}

// ExportNonlocalLoginPlaces
// 本接口 (ExportNonlocalLoginPlaces) 用于导出异地登录事件记录CSV文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportPrivilegeEvents")
    
    
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
    return c.ExportPrivilegeEventsWithContext(context.Background(), request)
}

// ExportPrivilegeEvents
// 导出本地提权事件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORT = "FailedOperation.Export"
//  INTERNALERROR = "InternalError"
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

func NewExportReverseShellEventsRequest() (request *ExportReverseShellEventsRequest) {
    request = &ExportReverseShellEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportReverseShellEvents")
    
    
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
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

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
    request = &IgnoreImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "IgnoreImpactedHosts")
    
    
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
    return c.IgnoreImpactedHostsWithContext(context.Background(), request)
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

func NewInquiryPriceOpenProVersionPrepaidRequest() (request *InquiryPriceOpenProVersionPrepaidRequest) {
    request = &InquiryPriceOpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "InquiryPriceOpenProVersionPrepaid")
    
    
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
    return c.InquiryPriceOpenProVersionPrepaidWithContext(context.Background(), request)
}

// InquiryPriceOpenProVersionPrepaid
// 本接口 (InquiryPriceOpenProVersionPrepaid) 用于开通专业版询价(预付费)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INQUIRYPRICE = "FailedOperation.InquiryPrice"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) InquiryPriceOpenProVersionPrepaidWithContext(ctx context.Context, request *InquiryPriceOpenProVersionPrepaidRequest) (response *InquiryPriceOpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewInquiryPriceOpenProVersionPrepaidRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceOpenProVersionPrepaid require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceOpenProVersionPrepaidResponse()
    err = c.Send(request, response)
    return
}

func NewMisAlarmNonlocalLoginPlacesRequest() (request *MisAlarmNonlocalLoginPlacesRequest) {
    request = &MisAlarmNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "MisAlarmNonlocalLoginPlaces")
    
    
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
    return c.MisAlarmNonlocalLoginPlacesWithContext(context.Background(), request)
}

// MisAlarmNonlocalLoginPlaces
// 本接口{MisAlarmNonlocalLoginPlaces}将设置当前地点为常用登录地。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) MisAlarmNonlocalLoginPlacesWithContext(ctx context.Context, request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewMisAlarmNonlocalLoginPlacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MisAlarmNonlocalLoginPlaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewMisAlarmNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmAttributeRequest() (request *ModifyAlarmAttributeRequest) {
    request = &ModifyAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAlarmAttribute")
    
    
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
    return c.ModifyAlarmAttributeWithContext(context.Background(), request)
}

// ModifyAlarmAttribute
// 本接口（ModifyAlarmAttribute）用于修改告警设置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAlarmAttributeWithContext(ctx context.Context, request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
    request = &ModifyAutoOpenProVersionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAutoOpenProVersionConfig")
    
    
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
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    return c.ModifyAutoOpenProVersionConfigWithContext(context.Background(), request)
}

// ModifyAutoOpenProVersionConfig
// 本接口 (ModifyAutoOpenProVersionConfig) 用于设置新增主机自动开通专业版配置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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

func NewModifyLoginWhiteListRequest() (request *ModifyLoginWhiteListRequest) {
    request = &ModifyLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyLoginWhiteList")
    
    
    return
}

func NewModifyLoginWhiteListResponse() (response *ModifyLoginWhiteListResponse) {
    response = &ModifyLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoginWhiteList
// 编辑白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyLoginWhiteList(request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    return c.ModifyLoginWhiteListWithContext(context.Background(), request)
}

// ModifyLoginWhiteList
// 编辑白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) ModifyLoginWhiteListWithContext(ctx context.Context, request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoginWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
    request = &ModifyProVersionRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyProVersionRenewFlag")
    
    
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
    return c.ModifyProVersionRenewFlagWithContext(context.Background(), request)
}

// ModifyProVersionRenewFlag
// 本接口 (ModifyProVersionRenewFlag) 用于修改专业版包年包月续费标识。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyProVersionRenewFlagWithContext(ctx context.Context, request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProVersionRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProVersionRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
    request = &OpenProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersion")
    
    
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
    return c.OpenProVersionWithContext(context.Background(), request)
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
func (c *Client) OpenProVersionWithContext(ctx context.Context, request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    if request == nil {
        request = NewOpenProVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenProVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionPrepaidRequest() (request *OpenProVersionPrepaidRequest) {
    request = &OpenProVersionPrepaidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersionPrepaid")
    
    
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
    return c.OpenProVersionPrepaidWithContext(context.Background(), request)
}

// OpenProVersionPrepaid
// 本接口 (OpenProVersionPrepaid) 用于开通专业版(包年包月)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
func (c *Client) OpenProVersionPrepaidWithContext(ctx context.Context, request *OpenProVersionPrepaidRequest) (response *OpenProVersionPrepaidResponse, err error) {
    if request == nil {
        request = NewOpenProVersionPrepaidRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenProVersionPrepaid require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenProVersionPrepaidResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
    request = &RecoverMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "RecoverMalwares")
    
    
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
    return c.RecoverMalwaresWithContext(context.Background(), request)
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

func NewRenewProVersionRequest() (request *RenewProVersionRequest) {
    request = &RenewProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "RenewProVersion")
    
    
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
    return c.RenewProVersionWithContext(context.Background(), request)
}

// RenewProVersion
// 本接口 (RenewProVersion) 用于续费专业版(包年包月)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewProVersionWithContext(ctx context.Context, request *RenewProVersionRequest) (response *RenewProVersionResponse, err error) {
    if request == nil {
        request = NewRenewProVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewProVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRescanImpactedHostRequest() (request *RescanImpactedHostRequest) {
    request = &RescanImpactedHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "RescanImpactedHost")
    
    
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
    return c.RescanImpactedHostWithContext(context.Background(), request)
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
func (c *Client) RescanImpactedHostWithContext(ctx context.Context, request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    if request == nil {
        request = NewRescanImpactedHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RescanImpactedHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewRescanImpactedHostResponse()
    err = c.Send(request, response)
    return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
    request = &SeparateMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "SeparateMalwares")
    
    
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
    return c.SeparateMalwaresWithContext(context.Background(), request)
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
    
    request.Init().WithApiInfo("yunjing", APIVersion, "SetBashEventsStatus")
    
    
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
    return c.SetBashEventsStatusWithContext(context.Background(), request)
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

func NewSwitchBashRulesRequest() (request *SwitchBashRulesRequest) {
    request = &SwitchBashRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "SwitchBashRules")
    
    
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
    return c.SwitchBashRulesWithContext(context.Background(), request)
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

func NewTrustMaliciousRequestRequest() (request *TrustMaliciousRequestRequest) {
    request = &TrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMaliciousRequest")
    
    
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
    return c.TrustMaliciousRequestWithContext(context.Background(), request)
}

// TrustMaliciousRequest
// 本接口 (TrustMaliciousRequest) 用于恶意请求添加信任。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
func (c *Client) TrustMaliciousRequestWithContext(ctx context.Context, request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewTrustMaliciousRequestRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrustMaliciousRequest require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
    request = &TrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMalwares")
    
    
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
    return c.TrustMalwaresWithContext(context.Background(), request)
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

func NewUntrustMaliciousRequestRequest() (request *UntrustMaliciousRequestRequest) {
    request = &UntrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMaliciousRequest")
    
    
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
    return c.UntrustMaliciousRequestWithContext(context.Background(), request)
}

// UntrustMaliciousRequest
// 本接口 (UntrustMaliciousRequest) 用于取消信任恶意请求。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UntrustMaliciousRequestWithContext(ctx context.Context, request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewUntrustMaliciousRequestRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UntrustMaliciousRequest require credential")
    }

    request.SetContext(ctx)
    
    response = NewUntrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
    request = &UntrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMalwares")
    
    
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
    return c.UntrustMalwaresWithContext(context.Background(), request)
}

// UntrustMalwares
// 本接口（UntrustMalwares）用于取消信任木马文件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"
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
