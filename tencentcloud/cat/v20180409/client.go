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
//  FAILEDOPERATION_ERRPREPAIDRESOURCEEXPIRE = "FailedOperation.ErrPrePaidResourceExpire"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_ORDEROUTOFCREDIT = "FailedOperation.OrderOutOfCredit"
//  FAILEDOPERATION_PRERESOURCEIDFAILED = "FailedOperation.PreResourceIDFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  FAILEDOPERATION_TAGREQUIREDVERIFYFAILED = "FailedOperation.TagRequiredVerifyFailed"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  FAILEDOPERATION_TRIALTASKEXCEED = "FailedOperation.TrialTaskExceed"
//  FAILEDOPERATION_UNMARSHALRESPONSE = "FailedOperation.UnmarshalResponse"
//  FAILEDOPERATION_USERNOQCLOUDTAGFULLACCESS = "FailedOperation.UserNoQcloudTAGFullAccess"
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
//  FAILEDOPERATION_ERRPREPAIDRESOURCEEXPIRE = "FailedOperation.ErrPrePaidResourceExpire"
//  FAILEDOPERATION_NOVALIDNODES = "FailedOperation.NoValidNodes"
//  FAILEDOPERATION_ORDEROUTOFCREDIT = "FailedOperation.OrderOutOfCredit"
//  FAILEDOPERATION_PRERESOURCEIDFAILED = "FailedOperation.PreResourceIDFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  FAILEDOPERATION_TAGREQUIREDVERIFYFAILED = "FailedOperation.TagRequiredVerifyFailed"
//  FAILEDOPERATION_TASKNOTRUNNING = "FailedOperation.TaskNotRunning"
//  FAILEDOPERATION_TASKNOTSUSPENDED = "FailedOperation.TaskNotSuspended"
//  FAILEDOPERATION_TASKOPERATIONNOTALLOW = "FailedOperation.TaskOperationNotAllow"
//  FAILEDOPERATION_TASKTYPENOTSAME = "FailedOperation.TaskTypeNotSame"
//  FAILEDOPERATION_TRIALTASKEXCEED = "FailedOperation.TrialTaskExceed"
//  FAILEDOPERATION_UNMARSHALRESPONSE = "FailedOperation.UnmarshalResponse"
//  FAILEDOPERATION_USERNOQCLOUDTAGFULLACCESS = "FailedOperation.UserNoQcloudTAGFullAccess"
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

func NewDescribeNodesRequest() (request *DescribeNodesRequest) {
    request = &DescribeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cat", APIVersion, "DescribeNodes")
    
    
    return
}

func NewDescribeNodesResponse() (response *DescribeNodesResponse) {
    response = &DescribeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNodes
// 获取拨测节点
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
func (c *Client) DescribeNodes(request *DescribeNodesRequest) (response *DescribeNodesResponse, err error) {
    return c.DescribeNodesWithContext(context.Background(), request)
}

// DescribeNodes
// 获取拨测节点
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
func (c *Client) DescribeNodesWithContext(ctx context.Context, request *DescribeNodesRequest) (response *DescribeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodesResponse()
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
// 查询云拨测指标数据，指标支持使用sum,avg,max,min聚合函数进行指标数据查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeProbeMetricData(request *DescribeProbeMetricDataRequest) (response *DescribeProbeMetricDataResponse, err error) {
    return c.DescribeProbeMetricDataWithContext(context.Background(), request)
}

// DescribeProbeMetricData
// 查询云拨测指标数据，指标支持使用sum,avg,max,min聚合函数进行指标数据查询
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
//  FAILEDOPERATION_ERRPREPAIDRESOURCEEXPIRE = "FailedOperation.ErrPrePaidResourceExpire"
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
//  FAILEDOPERATION_ERRPREPAIDRESOURCEEXPIRE = "FailedOperation.ErrPrePaidResourceExpire"
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

func NewUpdateProbeTaskAttributesRequest() (request *UpdateProbeTaskAttributesRequest) {
    request = &UpdateProbeTaskAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cat", APIVersion, "UpdateProbeTaskAttributes")
    
    
    return
}

func NewUpdateProbeTaskAttributesResponse() (response *UpdateProbeTaskAttributesResponse) {
    response = &UpdateProbeTaskAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProbeTaskAttributes
// 更新探测任务属性
//
// 可能返回的错误码:
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
func (c *Client) UpdateProbeTaskAttributes(request *UpdateProbeTaskAttributesRequest) (response *UpdateProbeTaskAttributesResponse, err error) {
    return c.UpdateProbeTaskAttributesWithContext(context.Background(), request)
}

// UpdateProbeTaskAttributes
// 更新探测任务属性
//
// 可能返回的错误码:
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
func (c *Client) UpdateProbeTaskAttributesWithContext(ctx context.Context, request *UpdateProbeTaskAttributesRequest) (response *UpdateProbeTaskAttributesResponse, err error) {
    if request == nil {
        request = NewUpdateProbeTaskAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProbeTaskAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProbeTaskAttributesResponse()
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
//  FAILEDOPERATION_PRERESOURCEIDFAILED = "FailedOperation.PreResourceIDFailed"
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
//  FAILEDOPERATION_PRERESOURCEIDFAILED = "FailedOperation.PreResourceIDFailed"
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
