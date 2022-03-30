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

package v20180409

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-09"

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


func NewBindAlarmPolicyRequest() (request *BindAlarmPolicyRequest) {
    request = &BindAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "BindAlarmPolicy")
    
    
    return
}

func NewBindAlarmPolicyResponse() (response *BindAlarmPolicyResponse) {
    response = &BindAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAlarmPolicy
// 绑定拨测任务和告警策略组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindAlarmPolicy(request *BindAlarmPolicyRequest) (response *BindAlarmPolicyResponse, err error) {
    return c.BindAlarmPolicyWithContext(context.Background(), request)
}

// BindAlarmPolicy
// 绑定拨测任务和告警策略组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindAlarmPolicyWithContext(ctx context.Context, request *BindAlarmPolicyRequest) (response *BindAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewBindAlarmPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentGroupRequest() (request *CreateAgentGroupRequest) {
    request = &CreateAgentGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "CreateAgentGroup")
    
    
    return
}

func NewCreateAgentGroupResponse() (response *CreateAgentGroupResponse) {
    response = &CreateAgentGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAgentGroup
// 添加拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgentGroup(request *CreateAgentGroupRequest) (response *CreateAgentGroupResponse, err error) {
    return c.CreateAgentGroupWithContext(context.Background(), request)
}

// CreateAgentGroup
// 添加拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgentGroupWithContext(ctx context.Context, request *CreateAgentGroupRequest) (response *CreateAgentGroupResponse, err error) {
    if request == nil {
        request = NewCreateAgentGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProbeTasksRequest() (request *CreateProbeTasksRequest) {
    request = &CreateProbeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "CreateProbeTasks")
    
    
    return
}

func NewCreateProbeTasksResponse() (response *CreateProbeTasksResponse) {
    response = &CreateProbeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProbeTasks
// 批量创建拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_ORDEROUTOFCREDIT = "FailedOperation.OrderOutOfCredit"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  FAILEDOPERATION_TRIALTASKEXCEED = "FailedOperation.TrialTaskExceed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateProbeTasks(request *CreateProbeTasksRequest) (response *CreateProbeTasksResponse, err error) {
    return c.CreateProbeTasksWithContext(context.Background(), request)
}

// CreateProbeTasks
// 批量创建拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_ORDEROUTOFCREDIT = "FailedOperation.OrderOutOfCredit"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  FAILEDOPERATION_TRIALTASKEXCEED = "FailedOperation.TrialTaskExceed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateProbeTasksWithContext(ctx context.Context, request *CreateProbeTasksRequest) (response *CreateProbeTasksResponse, err error) {
    if request == nil {
        request = NewCreateProbeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProbeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProbeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskExRequest() (request *CreateTaskExRequest) {
    request = &CreateTaskExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "CreateTaskEx")
    
    
    return
}

func NewCreateTaskExResponse() (response *CreateTaskExResponse) {
    response = &CreateTaskExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTaskEx
// 创建拨测任务(扩展)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskEx(request *CreateTaskExRequest) (response *CreateTaskExResponse, err error) {
    return c.CreateTaskExWithContext(context.Background(), request)
}

// CreateTaskEx
// 创建拨测任务(扩展)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskExWithContext(ctx context.Context, request *CreateTaskExRequest) (response *CreateTaskExResponse, err error) {
    if request == nil {
        request = NewCreateTaskExRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskEx require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskExResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentGroupRequest() (request *DeleteAgentGroupRequest) {
    request = &DeleteAgentGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DeleteAgentGroup")
    
    
    return
}

func NewDeleteAgentGroupResponse() (response *DeleteAgentGroupResponse) {
    response = &DeleteAgentGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAgentGroup
// 删除拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAgentGroup(request *DeleteAgentGroupRequest) (response *DeleteAgentGroupResponse, err error) {
    return c.DeleteAgentGroupWithContext(context.Background(), request)
}

// DeleteAgentGroup
// 删除拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAgentGroupWithContext(ctx context.Context, request *DeleteAgentGroupRequest) (response *DeleteAgentGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAgentGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProbeTaskRequest() (request *DeleteProbeTaskRequest) {
    request = &DeleteProbeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DeleteProbeTask")
    
    
    return
}

func NewDeleteProbeTaskResponse() (response *DeleteProbeTaskResponse) {
    response = &DeleteProbeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProbeTask
// 删除拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteProbeTask(request *DeleteProbeTaskRequest) (response *DeleteProbeTaskResponse, err error) {
    return c.DeleteProbeTaskWithContext(context.Background(), request)
}

// DeleteProbeTask
// 删除拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteProbeTaskWithContext(ctx context.Context, request *DeleteProbeTaskRequest) (response *DeleteProbeTaskResponse, err error) {
    if request == nil {
        request = NewDeleteProbeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProbeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProbeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTasksRequest() (request *DeleteTasksRequest) {
    request = &DeleteTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DeleteTasks")
    
    
    return
}

func NewDeleteTasksResponse() (response *DeleteTasksResponse) {
    response = &DeleteTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTasks
// 删除多个拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTasks(request *DeleteTasksRequest) (response *DeleteTasksResponse, err error) {
    return c.DeleteTasksWithContext(context.Background(), request)
}

// DeleteTasks
// 删除多个拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTasksWithContext(ctx context.Context, request *DeleteTasksRequest) (response *DeleteTasksResponse, err error) {
    if request == nil {
        request = NewDeleteTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentGroupsRequest() (request *DescribeAgentGroupsRequest) {
    request = &DescribeAgentGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeAgentGroups")
    
    
    return
}

func NewDescribeAgentGroupsResponse() (response *DescribeAgentGroupsResponse) {
    response = &DescribeAgentGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentGroups
// 查询拨测分组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentGroups(request *DescribeAgentGroupsRequest) (response *DescribeAgentGroupsResponse, err error) {
    return c.DescribeAgentGroupsWithContext(context.Background(), request)
}

// DescribeAgentGroups
// 查询拨测分组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentGroupsWithContext(ctx context.Context, request *DescribeAgentGroupsRequest) (response *DescribeAgentGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentsRequest() (request *DescribeAgentsRequest) {
    request = &DescribeAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeAgents")
    
    
    return
}

func NewDescribeAgentsResponse() (response *DescribeAgentsResponse) {
    response = &DescribeAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgents
// 查询本用户可选的拨测点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgents(request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
    return c.DescribeAgentsWithContext(context.Background(), request)
}

// DescribeAgents
// 查询本用户可选的拨测点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgentsWithContext(ctx context.Context, request *DescribeAgentsRequest) (response *DescribeAgentsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmTopicRequest() (request *DescribeAlarmTopicRequest) {
    request = &DescribeAlarmTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarmTopic")
    
    
    return
}

func NewDescribeAlarmTopicResponse() (response *DescribeAlarmTopicResponse) {
    response = &DescribeAlarmTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmTopic
// 查询用户的告警主题列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmTopic(request *DescribeAlarmTopicRequest) (response *DescribeAlarmTopicResponse, err error) {
    return c.DescribeAlarmTopicWithContext(context.Background(), request)
}

// DescribeAlarmTopic
// 查询用户的告警主题列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmTopicWithContext(ctx context.Context, request *DescribeAlarmTopicRequest) (response *DescribeAlarmTopicResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
    request = &DescribeAlarmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarms")
    
    
    return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
    response = &DescribeAlarmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarms
// 查询拨测告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    return c.DescribeAlarmsWithContext(context.Background(), request)
}

// DescribeAlarms
// 查询拨测告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmsByTaskRequest() (request *DescribeAlarmsByTaskRequest) {
    request = &DescribeAlarmsByTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeAlarmsByTask")
    
    
    return
}

func NewDescribeAlarmsByTaskResponse() (response *DescribeAlarmsByTaskResponse) {
    response = &DescribeAlarmsByTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmsByTask
// 按任务查询拨测告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmsByTask(request *DescribeAlarmsByTaskRequest) (response *DescribeAlarmsByTaskResponse, err error) {
    return c.DescribeAlarmsByTaskWithContext(context.Background(), request)
}

// DescribeAlarmsByTask
// 按任务查询拨测告警列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmsByTaskWithContext(ctx context.Context, request *DescribeAlarmsByTaskRequest) (response *DescribeAlarmsByTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsByTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmsByTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmsByTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCatLogsRequest() (request *DescribeCatLogsRequest) {
    request = &DescribeCatLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeCatLogs")
    
    
    return
}

func NewDescribeCatLogsResponse() (response *DescribeCatLogsResponse) {
    response = &DescribeCatLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCatLogs
// 查询拨测流水
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCatLogs(request *DescribeCatLogsRequest) (response *DescribeCatLogsResponse, err error) {
    return c.DescribeCatLogsWithContext(context.Background(), request)
}

// DescribeCatLogs
// 查询拨测流水
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCatLogsWithContext(ctx context.Context, request *DescribeCatLogsRequest) (response *DescribeCatLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCatLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCatLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCatLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailedSingleProbeDataRequest() (request *DescribeDetailedSingleProbeDataRequest) {
    request = &DescribeDetailedSingleProbeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeDetailedSingleProbeData")
    
    
    return
}

func NewDescribeDetailedSingleProbeDataResponse() (response *DescribeDetailedSingleProbeDataResponse) {
    response = &DescribeDetailedSingleProbeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDetailedSingleProbeData
// 根据时间范围、任务ID、运营商等条件查询单次拨测详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDetailedSingleProbeData(request *DescribeDetailedSingleProbeDataRequest) (response *DescribeDetailedSingleProbeDataResponse, err error) {
    return c.DescribeDetailedSingleProbeDataWithContext(context.Background(), request)
}

// DescribeDetailedSingleProbeData
// 根据时间范围、任务ID、运营商等条件查询单次拨测详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDetailedSingleProbeDataWithContext(ctx context.Context, request *DescribeDetailedSingleProbeDataRequest) (response *DescribeDetailedSingleProbeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDetailedSingleProbeDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetailedSingleProbeData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetailedSingleProbeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProbeMetricDataRequest() (request *DescribeProbeMetricDataRequest) {
    request = &DescribeProbeMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeProbeMetricData")
    
    
    return
}

func NewDescribeProbeMetricDataResponse() (response *DescribeProbeMetricDataResponse) {
    response = &DescribeProbeMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProbeMetricData
// 列出云拨测指标详细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeProbeMetricData(request *DescribeProbeMetricDataRequest) (response *DescribeProbeMetricDataResponse, err error) {
    return c.DescribeProbeMetricDataWithContext(context.Background(), request)
}

// DescribeProbeMetricData
// 列出云拨测指标详细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeProbeMetricDataWithContext(ctx context.Context, request *DescribeProbeMetricDataRequest) (response *DescribeProbeMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeProbeMetricDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProbeMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProbeMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProbeNodesRequest() (request *DescribeProbeNodesRequest) {
    request = &DescribeProbeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeProbeNodes")
    
    
    return
}

func NewDescribeProbeNodesResponse() (response *DescribeProbeNodesResponse) {
    response = &DescribeProbeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProbeNodes
// 查询拨测节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProbeNodes(request *DescribeProbeNodesRequest) (response *DescribeProbeNodesResponse, err error) {
    return c.DescribeProbeNodesWithContext(context.Background(), request)
}

// DescribeProbeNodes
// 查询拨测节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProbeNodesWithContext(ctx context.Context, request *DescribeProbeNodesRequest) (response *DescribeProbeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeProbeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProbeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProbeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProbeTasksRequest() (request *DescribeProbeTasksRequest) {
    request = &DescribeProbeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeProbeTasks")
    
    
    return
}

func NewDescribeProbeTasksResponse() (response *DescribeProbeTasksResponse) {
    response = &DescribeProbeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProbeTasks
// 查询拨测任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProbeTasks(request *DescribeProbeTasksRequest) (response *DescribeProbeTasksResponse, err error) {
    return c.DescribeProbeTasksWithContext(context.Background(), request)
}

// DescribeProbeTasks
// 查询拨测任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProbeTasksWithContext(ctx context.Context, request *DescribeProbeTasksRequest) (response *DescribeProbeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeProbeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProbeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProbeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// 查询拨测任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// 查询拨测任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksByTypeRequest() (request *DescribeTasksByTypeRequest) {
    request = &DescribeTasksByTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeTasksByType")
    
    
    return
}

func NewDescribeTasksByTypeResponse() (response *DescribeTasksByTypeResponse) {
    response = &DescribeTasksByTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasksByType
// 按类型查询拨测任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTasksByType(request *DescribeTasksByTypeRequest) (response *DescribeTasksByTypeResponse, err error) {
    return c.DescribeTasksByTypeWithContext(context.Background(), request)
}

// DescribeTasksByType
// 按类型查询拨测任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTasksByTypeWithContext(ctx context.Context, request *DescribeTasksByTypeRequest) (response *DescribeTasksByTypeResponse, err error) {
    if request == nil {
        request = NewDescribeTasksByTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasksByType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksByTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserLimitRequest() (request *DescribeUserLimitRequest) {
    request = &DescribeUserLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "DescribeUserLimit")
    
    
    return
}

func NewDescribeUserLimitResponse() (response *DescribeUserLimitResponse) {
    response = &DescribeUserLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserLimit
// 获取用户可用资源限制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserLimit(request *DescribeUserLimitRequest) (response *DescribeUserLimitResponse, err error) {
    return c.DescribeUserLimitWithContext(context.Background(), request)
}

// DescribeUserLimit
// 获取用户可用资源限制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserLimitWithContext(ctx context.Context, request *DescribeUserLimitRequest) (response *DescribeUserLimitResponse, err error) {
    if request == nil {
        request = NewDescribeUserLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserLimitResponse()
    err = c.Send(request, response)
    return
}

func NewGetAvailRatioHistoryRequest() (request *GetAvailRatioHistoryRequest) {
    request = &GetAvailRatioHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetAvailRatioHistory")
    
    
    return
}

func NewGetAvailRatioHistoryResponse() (response *GetAvailRatioHistoryResponse) {
    response = &GetAvailRatioHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAvailRatioHistory
// 获取指定时刻的可用率地图信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetAvailRatioHistory(request *GetAvailRatioHistoryRequest) (response *GetAvailRatioHistoryResponse, err error) {
    return c.GetAvailRatioHistoryWithContext(context.Background(), request)
}

// GetAvailRatioHistory
// 获取指定时刻的可用率地图信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetAvailRatioHistoryWithContext(ctx context.Context, request *GetAvailRatioHistoryRequest) (response *GetAvailRatioHistoryResponse, err error) {
    if request == nil {
        request = NewGetAvailRatioHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAvailRatioHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAvailRatioHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetDailyAvailRatioRequest() (request *GetDailyAvailRatioRequest) {
    request = &GetDailyAvailRatioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetDailyAvailRatio")
    
    
    return
}

func NewGetDailyAvailRatioResponse() (response *GetDailyAvailRatioResponse) {
    response = &GetDailyAvailRatioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDailyAvailRatio
// 获取一天的整体可用率信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetDailyAvailRatio(request *GetDailyAvailRatioRequest) (response *GetDailyAvailRatioResponse, err error) {
    return c.GetDailyAvailRatioWithContext(context.Background(), request)
}

// GetDailyAvailRatio
// 获取一天的整体可用率信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetDailyAvailRatioWithContext(ctx context.Context, request *GetDailyAvailRatioRequest) (response *GetDailyAvailRatioResponse, err error) {
    if request == nil {
        request = NewGetDailyAvailRatioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDailyAvailRatio require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDailyAvailRatioResponse()
    err = c.Send(request, response)
    return
}

func NewGetRealAvailRatioRequest() (request *GetRealAvailRatioRequest) {
    request = &GetRealAvailRatioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetRealAvailRatio")
    
    
    return
}

func NewGetRealAvailRatioResponse() (response *GetRealAvailRatioResponse) {
    response = &GetRealAvailRatioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRealAvailRatio
// 获取实时可用率信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRealAvailRatio(request *GetRealAvailRatioRequest) (response *GetRealAvailRatioResponse, err error) {
    return c.GetRealAvailRatioWithContext(context.Background(), request)
}

// GetRealAvailRatio
// 获取实时可用率信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRealAvailRatioWithContext(ctx context.Context, request *GetRealAvailRatioRequest) (response *GetRealAvailRatioResponse, err error) {
    if request == nil {
        request = NewGetRealAvailRatioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRealAvailRatio require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRealAvailRatioResponse()
    err = c.Send(request, response)
    return
}

func NewGetRespTimeTrendExRequest() (request *GetRespTimeTrendExRequest) {
    request = &GetRespTimeTrendExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetRespTimeTrendEx")
    
    
    return
}

func NewGetRespTimeTrendExResponse() (response *GetRespTimeTrendExResponse) {
    response = &GetRespTimeTrendExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRespTimeTrendEx
// 查询拨测任务的走势数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRespTimeTrendEx(request *GetRespTimeTrendExRequest) (response *GetRespTimeTrendExResponse, err error) {
    return c.GetRespTimeTrendExWithContext(context.Background(), request)
}

// GetRespTimeTrendEx
// 查询拨测任务的走势数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRespTimeTrendExWithContext(ctx context.Context, request *GetRespTimeTrendExRequest) (response *GetRespTimeTrendExResponse, err error) {
    if request == nil {
        request = NewGetRespTimeTrendExRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRespTimeTrendEx require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRespTimeTrendExResponse()
    err = c.Send(request, response)
    return
}

func NewGetResultSummaryRequest() (request *GetResultSummaryRequest) {
    request = &GetResultSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetResultSummary")
    
    
    return
}

func NewGetResultSummaryResponse() (response *GetResultSummaryResponse) {
    response = &GetResultSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetResultSummary
// 获取任务列表的实时数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResultSummary(request *GetResultSummaryRequest) (response *GetResultSummaryResponse, err error) {
    return c.GetResultSummaryWithContext(context.Background(), request)
}

// GetResultSummary
// 获取任务列表的实时数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResultSummaryWithContext(ctx context.Context, request *GetResultSummaryRequest) (response *GetResultSummaryResponse, err error) {
    if request == nil {
        request = NewGetResultSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResultSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResultSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetReturnCodeHistoryRequest() (request *GetReturnCodeHistoryRequest) {
    request = &GetReturnCodeHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetReturnCodeHistory")
    
    
    return
}

func NewGetReturnCodeHistoryResponse() (response *GetReturnCodeHistoryResponse) {
    response = &GetReturnCodeHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetReturnCodeHistory
// 查询拨测任务的历史返回码信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetReturnCodeHistory(request *GetReturnCodeHistoryRequest) (response *GetReturnCodeHistoryResponse, err error) {
    return c.GetReturnCodeHistoryWithContext(context.Background(), request)
}

// GetReturnCodeHistory
// 查询拨测任务的历史返回码信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetReturnCodeHistoryWithContext(ctx context.Context, request *GetReturnCodeHistoryRequest) (response *GetReturnCodeHistoryResponse, err error) {
    if request == nil {
        request = NewGetReturnCodeHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetReturnCodeHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetReturnCodeHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetReturnCodeInfoRequest() (request *GetReturnCodeInfoRequest) {
    request = &GetReturnCodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetReturnCodeInfo")
    
    
    return
}

func NewGetReturnCodeInfoResponse() (response *GetReturnCodeInfoResponse) {
    response = &GetReturnCodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetReturnCodeInfo
// 查询拨测任务的返回码统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetReturnCodeInfo(request *GetReturnCodeInfoRequest) (response *GetReturnCodeInfoResponse, err error) {
    return c.GetReturnCodeInfoWithContext(context.Background(), request)
}

// GetReturnCodeInfo
// 查询拨测任务的返回码统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetReturnCodeInfoWithContext(ctx context.Context, request *GetReturnCodeInfoRequest) (response *GetReturnCodeInfoResponse, err error) {
    if request == nil {
        request = NewGetReturnCodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetReturnCodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetReturnCodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskTotalNumberRequest() (request *GetTaskTotalNumberRequest) {
    request = &GetTaskTotalNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "GetTaskTotalNumber")
    
    
    return
}

func NewGetTaskTotalNumberResponse() (response *GetTaskTotalNumberResponse) {
    response = &GetTaskTotalNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTaskTotalNumber
// 获取AppId下的拨测任务总数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTaskTotalNumber(request *GetTaskTotalNumberRequest) (response *GetTaskTotalNumberResponse, err error) {
    return c.GetTaskTotalNumberWithContext(context.Background(), request)
}

// GetTaskTotalNumber
// 获取AppId下的拨测任务总数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetTaskTotalNumberWithContext(ctx context.Context, request *GetTaskTotalNumberRequest) (response *GetTaskTotalNumberResponse, err error) {
    if request == nil {
        request = NewGetTaskTotalNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskTotalNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskTotalNumberResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentGroupRequest() (request *ModifyAgentGroupRequest) {
    request = &ModifyAgentGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "ModifyAgentGroup")
    
    
    return
}

func NewModifyAgentGroupResponse() (response *ModifyAgentGroupResponse) {
    response = &ModifyAgentGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAgentGroup
// 修改拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAgentGroup(request *ModifyAgentGroupRequest) (response *ModifyAgentGroupResponse, err error) {
    return c.ModifyAgentGroupWithContext(context.Background(), request)
}

// ModifyAgentGroup
// 修改拨测分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAgentGroupWithContext(ctx context.Context, request *ModifyAgentGroupRequest) (response *ModifyAgentGroupResponse, err error) {
    if request == nil {
        request = NewModifyAgentGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskExRequest() (request *ModifyTaskExRequest) {
    request = &ModifyTaskExRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "ModifyTaskEx")
    
    
    return
}

func NewModifyTaskExResponse() (response *ModifyTaskExResponse) {
    response = &ModifyTaskExResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTaskEx
// 修改拨测任务(扩展)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTaskEx(request *ModifyTaskExRequest) (response *ModifyTaskExResponse, err error) {
    return c.ModifyTaskExWithContext(context.Background(), request)
}

// ModifyTaskEx
// 修改拨测任务(扩展)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTaskExWithContext(ctx context.Context, request *ModifyTaskExRequest) (response *ModifyTaskExResponse, err error) {
    if request == nil {
        request = NewModifyTaskExRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskEx require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskExResponse()
    err = c.Send(request, response)
    return
}

func NewPauseTaskRequest() (request *PauseTaskRequest) {
    request = &PauseTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "PauseTask")
    
    
    return
}

func NewPauseTaskResponse() (response *PauseTaskResponse) {
    response = &PauseTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PauseTask
// 暂停拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseTask(request *PauseTaskRequest) (response *PauseTaskResponse, err error) {
    return c.PauseTaskWithContext(context.Background(), request)
}

// PauseTask
// 暂停拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseTaskWithContext(ctx context.Context, request *PauseTaskRequest) (response *PauseTaskResponse, err error) {
    if request == nil {
        request = NewPauseTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseTaskResponse()
    err = c.Send(request, response)
    return
}

func NewResumeProbeTaskRequest() (request *ResumeProbeTaskRequest) {
    request = &ResumeProbeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "ResumeProbeTask")
    
    
    return
}

func NewResumeProbeTaskResponse() (response *ResumeProbeTaskResponse) {
    response = &ResumeProbeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeProbeTask
// 恢复拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ResumeProbeTask(request *ResumeProbeTaskRequest) (response *ResumeProbeTaskResponse, err error) {
    return c.ResumeProbeTaskWithContext(context.Background(), request)
}

// ResumeProbeTask
// 恢复拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ResumeProbeTaskWithContext(ctx context.Context, request *ResumeProbeTaskRequest) (response *ResumeProbeTaskResponse, err error) {
    if request == nil {
        request = NewResumeProbeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeProbeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeProbeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRunTaskRequest() (request *RunTaskRequest) {
    request = &RunTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "RunTask")
    
    
    return
}

func NewRunTaskResponse() (response *RunTaskResponse) {
    response = &RunTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunTask
// 运行拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunTask(request *RunTaskRequest) (response *RunTaskResponse, err error) {
    return c.RunTaskWithContext(context.Background(), request)
}

// RunTask
// 运行拨测任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunTaskWithContext(ctx context.Context, request *RunTaskRequest) (response *RunTaskResponse, err error) {
    if request == nil {
        request = NewRunTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSuspendProbeTaskRequest() (request *SuspendProbeTaskRequest) {
    request = &SuspendProbeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "SuspendProbeTask")
    
    
    return
}

func NewSuspendProbeTaskResponse() (response *SuspendProbeTaskResponse) {
    response = &SuspendProbeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SuspendProbeTask
// 暂停拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SuspendProbeTask(request *SuspendProbeTaskRequest) (response *SuspendProbeTaskResponse, err error) {
    return c.SuspendProbeTaskWithContext(context.Background(), request)
}

// SuspendProbeTask
// 暂停拨测任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SuspendProbeTaskWithContext(ctx context.Context, request *SuspendProbeTaskRequest) (response *SuspendProbeTaskResponse, err error) {
    if request == nil {
        request = NewSuspendProbeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SuspendProbeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSuspendProbeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProbeTaskConfigurationListRequest() (request *UpdateProbeTaskConfigurationListRequest) {
    request = &UpdateProbeTaskConfigurationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "UpdateProbeTaskConfigurationList")
    
    
    return
}

func NewUpdateProbeTaskConfigurationListResponse() (response *UpdateProbeTaskConfigurationListResponse) {
    response = &UpdateProbeTaskConfigurationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProbeTaskConfigurationList
// 批量更新拨测任务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateProbeTaskConfigurationList(request *UpdateProbeTaskConfigurationListRequest) (response *UpdateProbeTaskConfigurationListResponse, err error) {
    return c.UpdateProbeTaskConfigurationListWithContext(context.Background(), request)
}

// UpdateProbeTaskConfigurationList
// 批量更新拨测任务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateProbeTaskConfigurationListWithContext(ctx context.Context, request *UpdateProbeTaskConfigurationListRequest) (response *UpdateProbeTaskConfigurationListResponse, err error) {
    if request == nil {
        request = NewUpdateProbeTaskConfigurationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProbeTaskConfigurationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProbeTaskConfigurationListResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyResultRequest() (request *VerifyResultRequest) {
    request = &VerifyResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cat", APIVersion, "VerifyResult")
    
    
    return
}

func NewVerifyResultResponse() (response *VerifyResultResponse) {
    response = &VerifyResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyResult
// 验证拨测任务，结果验证查询（验证成功的，才建议创建拨测任务）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyResult(request *VerifyResultRequest) (response *VerifyResultResponse, err error) {
    return c.VerifyResultWithContext(context.Background(), request)
}

// VerifyResult
// 验证拨测任务，结果验证查询（验证成功的，才建议创建拨测任务）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyResultWithContext(ctx context.Context, request *VerifyResultRequest) (response *VerifyResultResponse, err error) {
    if request == nil {
        request = NewVerifyResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyResultResponse()
    err = c.Send(request, response)
    return
}
