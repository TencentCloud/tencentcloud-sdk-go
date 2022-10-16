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

package v20210820

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-08-20"

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


func NewBatchDeleteTasksNewRequest() (request *BatchDeleteTasksNewRequest) {
    request = &BatchDeleteTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchDeleteTasksNew")
    
    
    return
}

func NewBatchDeleteTasksNewResponse() (response *BatchDeleteTasksNewResponse) {
    response = &BatchDeleteTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeleteTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量删除任务，仅对任务状态为”已停止“有效；
func (c *Client) BatchDeleteTasksNew(request *BatchDeleteTasksNewRequest) (response *BatchDeleteTasksNewResponse, err error) {
    return c.BatchDeleteTasksNewWithContext(context.Background(), request)
}

// BatchDeleteTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量删除任务，仅对任务状态为”已停止“有效；
func (c *Client) BatchDeleteTasksNewWithContext(ctx context.Context, request *BatchDeleteTasksNewRequest) (response *BatchDeleteTasksNewResponse, err error) {
    if request == nil {
        request = NewBatchDeleteTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyOwnersNewRequest() (request *BatchModifyOwnersNewRequest) {
    request = &BatchModifyOwnersNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchModifyOwnersNew")
    
    
    return
}

func NewBatchModifyOwnersNewResponse() (response *BatchModifyOwnersNewResponse) {
    response = &BatchModifyOwnersNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyOwnersNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量修改任务责任人
func (c *Client) BatchModifyOwnersNew(request *BatchModifyOwnersNewRequest) (response *BatchModifyOwnersNewResponse, err error) {
    return c.BatchModifyOwnersNewWithContext(context.Background(), request)
}

// BatchModifyOwnersNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量修改任务责任人
func (c *Client) BatchModifyOwnersNewWithContext(ctx context.Context, request *BatchModifyOwnersNewRequest) (response *BatchModifyOwnersNewResponse, err error) {
    if request == nil {
        request = NewBatchModifyOwnersNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyOwnersNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyOwnersNewResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopTasksNewRequest() (request *BatchStopTasksNewRequest) {
    request = &BatchStopTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStopTasksNew")
    
    
    return
}

func NewBatchStopTasksNewResponse() (response *BatchStopTasksNewResponse) {
    response = &BatchStopTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchStopTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopTasksNew(request *BatchStopTasksNewRequest) (response *BatchStopTasksNewResponse, err error) {
    return c.BatchStopTasksNewWithContext(context.Background(), request)
}

// BatchStopTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopTasksNewWithContext(ctx context.Context, request *BatchStopTasksNewRequest) (response *BatchStopTasksNewResponse, err error) {
    if request == nil {
        request = NewBatchStopTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomFunctionRequest() (request *CreateCustomFunctionRequest) {
    request = &CreateCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCustomFunction")
    
    
    return
}

func NewCreateCustomFunctionResponse() (response *CreateCustomFunctionResponse) {
    response = &CreateCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomFunction
//  创建用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCustomFunction(request *CreateCustomFunctionRequest) (response *CreateCustomFunctionResponse, err error) {
    return c.CreateCustomFunctionWithContext(context.Background(), request)
}

// CreateCustomFunction
//  创建用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCustomFunctionWithContext(ctx context.Context, request *CreateCustomFunctionRequest) (response *CreateCustomFunctionResponse, err error) {
    if request == nil {
        request = NewCreateCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataSourceRequest() (request *CreateDataSourceRequest) {
    request = &CreateDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDataSource")
    
    
    return
}

func NewCreateDataSourceResponse() (response *CreateDataSourceResponse) {
    response = &CreateDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    return c.CreateDataSourceWithContext(context.Background(), request)
}

// CreateDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSourceWithContext(ctx context.Context, request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    if request == nil {
        request = NewCreateDataSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateFolder")
    
    
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

// CreateFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomFunctionRequest() (request *DeleteCustomFunctionRequest) {
    request = &DeleteCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCustomFunction")
    
    
    return
}

func NewDeleteCustomFunctionResponse() (response *DeleteCustomFunctionResponse) {
    response = &DeleteCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomFunction
// 删除用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCustomFunction(request *DeleteCustomFunctionRequest) (response *DeleteCustomFunctionResponse, err error) {
    return c.DeleteCustomFunctionWithContext(context.Background(), request)
}

// DeleteCustomFunction
// 删除用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCustomFunctionWithContext(ctx context.Context, request *DeleteCustomFunctionRequest) (response *DeleteCustomFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataSourcesRequest() (request *DeleteDataSourcesRequest) {
    request = &DeleteDataSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteDataSources")
    
    
    return
}

func NewDeleteDataSourcesResponse() (response *DeleteDataSourcesResponse) {
    response = &DeleteDataSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDataSources
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSources(request *DeleteDataSourcesRequest) (response *DeleteDataSourcesResponse, err error) {
    return c.DeleteDataSourcesWithContext(context.Background(), request)
}

// DeleteDataSources
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSourcesWithContext(ctx context.Context, request *DeleteDataSourcesRequest) (response *DeleteDataSourcesResponse, err error) {
    if request == nil {
        request = NewDeleteDataSourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataSourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
    request = &DeleteFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteFolder")
    
    
    return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
    response = &DeleteFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    return c.DeleteFolderWithContext(context.Background(), request)
}

// DeleteFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    if request == nil {
        request = NewDeleteFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowNewRequest() (request *DeleteWorkflowNewRequest) {
    request = &DeleteWorkflowNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowNew")
    
    
    return
}

func NewDeleteWorkflowNewResponse() (response *DeleteWorkflowNewResponse) {
    response = &DeleteWorkflowNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowNew(request *DeleteWorkflowNewRequest) (response *DeleteWorkflowNewResponse, err error) {
    return c.DeleteWorkflowNewWithContext(context.Background(), request)
}

// DeleteWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowNewWithContext(ctx context.Context, request *DeleteWorkflowNewRequest) (response *DeleteWorkflowNewResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceListRequest() (request *DescribeDataSourceListRequest) {
    request = &DescribeDataSourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceList")
    
    
    return
}

func NewDescribeDataSourceListResponse() (response *DescribeDataSourceListResponse) {
    response = &DescribeDataSourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataSourceList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeDataSourceList(request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    return c.DescribeDataSourceListWithContext(context.Background(), request)
}

// DescribeDataSourceList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeDataSourceListWithContext(ctx context.Context, request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceWithoutInfoRequest() (request *DescribeDataSourceWithoutInfoRequest) {
    request = &DescribeDataSourceWithoutInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceWithoutInfo")
    
    
    return
}

func NewDescribeDataSourceWithoutInfoResponse() (response *DescribeDataSourceWithoutInfoResponse) {
    response = &DescribeDataSourceWithoutInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataSourceWithoutInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceWithoutInfo(request *DescribeDataSourceWithoutInfoRequest) (response *DescribeDataSourceWithoutInfoResponse, err error) {
    return c.DescribeDataSourceWithoutInfoWithContext(context.Background(), request)
}

// DescribeDataSourceWithoutInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceWithoutInfoWithContext(ctx context.Context, request *DescribeDataSourceWithoutInfoRequest) (response *DescribeDataSourceWithoutInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceWithoutInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceWithoutInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceWithoutInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasourceRequest() (request *DescribeDatasourceRequest) {
    request = &DescribeDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDatasource")
    
    
    return
}

func NewDescribeDatasourceResponse() (response *DescribeDatasourceResponse) {
    response = &DescribeDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatasource(request *DescribeDatasourceRequest) (response *DescribeDatasourceResponse, err error) {
    return c.DescribeDatasourceWithContext(context.Background(), request)
}

// DescribeDatasource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatasourceWithContext(ctx context.Context, request *DescribeDatasourceRequest) (response *DescribeDatasourceResponse, err error) {
    if request == nil {
        request = NewDescribeDatasourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDependTasksNewRequest() (request *DescribeDependTasksNewRequest) {
    request = &DescribeDependTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDependTasksNew")
    
    
    return
}

func NewDescribeDependTasksNewResponse() (response *DescribeDependTasksNewResponse) {
    response = &DescribeDependTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDependTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTasksNew(request *DescribeDependTasksNewRequest) (response *DescribeDependTasksNewResponse, err error) {
    return c.DescribeDependTasksNewWithContext(context.Background(), request)
}

// DescribeDependTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTasksNewWithContext(ctx context.Context, request *DescribeDependTasksNewRequest) (response *DescribeDependTasksNewResponse, err error) {
    if request == nil {
        request = NewDescribeDependTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDependTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDependTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFolderListRequest() (request *DescribeFolderListRequest) {
    request = &DescribeFolderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFolderList")
    
    
    return
}

func NewDescribeFolderListResponse() (response *DescribeFolderListResponse) {
    response = &DescribeFolderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFolderList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderList(request *DescribeFolderListRequest) (response *DescribeFolderListResponse, err error) {
    return c.DescribeFolderListWithContext(context.Background(), request)
}

// DescribeFolderList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderListWithContext(ctx context.Context, request *DescribeFolderListRequest) (response *DescribeFolderListResponse, err error) {
    if request == nil {
        request = NewDescribeFolderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFolderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFolderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFolderWorkflowListRequest() (request *DescribeFolderWorkflowListRequest) {
    request = &DescribeFolderWorkflowListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFolderWorkflowList")
    
    
    return
}

func NewDescribeFolderWorkflowListResponse() (response *DescribeFolderWorkflowListResponse) {
    response = &DescribeFolderWorkflowListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFolderWorkflowList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹下的工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderWorkflowList(request *DescribeFolderWorkflowListRequest) (response *DescribeFolderWorkflowListResponse, err error) {
    return c.DescribeFolderWorkflowListWithContext(context.Background(), request)
}

// DescribeFolderWorkflowList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹下的工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderWorkflowListWithContext(ctx context.Context, request *DescribeFolderWorkflowListRequest) (response *DescribeFolderWorkflowListResponse, err error) {
    if request == nil {
        request = NewDescribeFolderWorkflowListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFolderWorkflowList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFolderWorkflowListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionKindsRequest() (request *DescribeFunctionKindsRequest) {
    request = &DescribeFunctionKindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFunctionKinds")
    
    
    return
}

func NewDescribeFunctionKindsResponse() (response *DescribeFunctionKindsResponse) {
    response = &DescribeFunctionKindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFunctionKinds
// 查询函数分类
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionKinds(request *DescribeFunctionKindsRequest) (response *DescribeFunctionKindsResponse, err error) {
    return c.DescribeFunctionKindsWithContext(context.Background(), request)
}

// DescribeFunctionKinds
// 查询函数分类
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionKindsWithContext(ctx context.Context, request *DescribeFunctionKindsRequest) (response *DescribeFunctionKindsResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionKindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionKinds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionKindsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionTypesRequest() (request *DescribeFunctionTypesRequest) {
    request = &DescribeFunctionTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFunctionTypes")
    
    
    return
}

func NewDescribeFunctionTypesResponse() (response *DescribeFunctionTypesResponse) {
    response = &DescribeFunctionTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFunctionTypes
// 查询函数类型
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionTypes(request *DescribeFunctionTypesRequest) (response *DescribeFunctionTypesResponse, err error) {
    return c.DescribeFunctionTypesWithContext(context.Background(), request)
}

// DescribeFunctionTypes
// 查询函数类型
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionTypesWithContext(ctx context.Context, request *DescribeFunctionTypesRequest) (response *DescribeFunctionTypesResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLogs")
    
    
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLogs
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogsWithContext(ctx context.Context, request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationalFunctionsRequest() (request *DescribeOrganizationalFunctionsRequest) {
    request = &DescribeOrganizationalFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOrganizationalFunctions")
    
    
    return
}

func NewDescribeOrganizationalFunctionsResponse() (response *DescribeOrganizationalFunctionsResponse) {
    response = &DescribeOrganizationalFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganizationalFunctions
// 查询全量函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationalFunctions(request *DescribeOrganizationalFunctionsRequest) (response *DescribeOrganizationalFunctionsResponse, err error) {
    return c.DescribeOrganizationalFunctionsWithContext(context.Background(), request)
}

// DescribeOrganizationalFunctions
// 查询全量函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationalFunctionsWithContext(ctx context.Context, request *DescribeOrganizationalFunctionsRequest) (response *DescribeOrganizationalFunctionsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationalFunctionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationalFunctions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationalFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeProject")
    
    
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    return c.DescribeProjectWithContext(context.Background(), request)
}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedInstancesRequest() (request *DescribeRelatedInstancesRequest) {
    request = &DescribeRelatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRelatedInstances")
    
    
    return
}

func NewDescribeRelatedInstancesResponse() (response *DescribeRelatedInstancesResponse) {
    response = &DescribeRelatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstances(request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    return c.DescribeRelatedInstancesWithContext(context.Background(), request)
}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstancesWithContext(ctx context.Context, request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务具体详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务具体详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewDescribeTaskInstancesRequest() (request *DescribeTaskInstancesRequest) {
    request = &DescribeTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskInstances")
    
    
    return
}

func NewDescribeTaskInstancesResponse() (response *DescribeTaskInstancesResponse) {
    response = &DescribeTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstances(request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    return c.DescribeTaskInstancesWithContext(context.Background(), request)
}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstancesWithContext(ctx context.Context, request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskScriptRequest() (request *DescribeTaskScriptRequest) {
    request = &DescribeTaskScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskScript")
    
    
    return
}

func NewDescribeTaskScriptResponse() (response *DescribeTaskScriptResponse) {
    response = &DescribeTaskScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskScript(request *DescribeTaskScriptRequest) (response *DescribeTaskScriptResponse, err error) {
    return c.DescribeTaskScriptWithContext(context.Background(), request)
}

// DescribeTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskScriptWithContext(ctx context.Context, request *DescribeTaskScriptRequest) (response *DescribeTaskScriptResponse, err error) {
    if request == nil {
        request = NewDescribeTaskScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksByPageRequest() (request *DescribeTasksByPageRequest) {
    request = &DescribeTasksByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTasksByPage")
    
    
    return
}

func NewDescribeTasksByPageResponse() (response *DescribeTasksByPageResponse) {
    response = &DescribeTasksByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasksByPage
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据工作流分页查询任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTasksByPage(request *DescribeTasksByPageRequest) (response *DescribeTasksByPageResponse, err error) {
    return c.DescribeTasksByPageWithContext(context.Background(), request)
}

// DescribeTasksByPage
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据工作流分页查询任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTasksByPageWithContext(ctx context.Context, request *DescribeTasksByPageRequest) (response *DescribeTasksByPageResponse, err error) {
    if request == nil {
        request = NewDescribeTasksByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasksByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksByPageResponse()
    err = c.Send(request, response)
    return
}

func NewForceSucInstancesRequest() (request *ForceSucInstancesRequest) {
    request = &ForceSucInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ForceSucInstances")
    
    
    return
}

func NewForceSucInstancesResponse() (response *ForceSucInstancesResponse) {
    response = &ForceSucInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForceSucInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量置成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucInstances(request *ForceSucInstancesRequest) (response *ForceSucInstancesResponse, err error) {
    return c.ForceSucInstancesWithContext(context.Background(), request)
}

// ForceSucInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量置成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucInstancesWithContext(ctx context.Context, request *ForceSucInstancesRequest) (response *ForceSucInstancesResponse, err error) {
    if request == nil {
        request = NewForceSucInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForceSucInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewForceSucInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeTasksRequest() (request *FreezeTasksRequest) {
    request = &FreezeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeTasks")
    
    
    return
}

func NewFreezeTasksResponse() (response *FreezeTasksResponse) {
    response = &FreezeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FreezeTasks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasks(request *FreezeTasksRequest) (response *FreezeTasksResponse, err error) {
    return c.FreezeTasksWithContext(context.Background(), request)
}

// FreezeTasks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksWithContext(ctx context.Context, request *FreezeTasksRequest) (response *FreezeTasksResponse, err error) {
    if request == nil {
        request = NewFreezeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeTasksByMultiWorkflowRequest() (request *FreezeTasksByMultiWorkflowRequest) {
    request = &FreezeTasksByMultiWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeTasksByMultiWorkflow")
    
    
    return
}

func NewFreezeTasksByMultiWorkflowResponse() (response *FreezeTasksByMultiWorkflowResponse) {
    response = &FreezeTasksByMultiWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FreezeTasksByMultiWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 基于多个工作流进行批量冻结任务操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByMultiWorkflow(request *FreezeTasksByMultiWorkflowRequest) (response *FreezeTasksByMultiWorkflowResponse, err error) {
    return c.FreezeTasksByMultiWorkflowWithContext(context.Background(), request)
}

// FreezeTasksByMultiWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 基于多个工作流进行批量冻结任务操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByMultiWorkflowWithContext(ctx context.Context, request *FreezeTasksByMultiWorkflowRequest) (response *FreezeTasksByMultiWorkflowResponse, err error) {
    if request == nil {
        request = NewFreezeTasksByMultiWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeTasksByMultiWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeTasksByMultiWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewKillInstancesRequest() (request *KillInstancesRequest) {
    request = &KillInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillInstances")
    
    
    return
}

func NewKillInstancesResponse() (response *KillInstancesResponse) {
    response = &KillInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KillInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量终止操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillInstances(request *KillInstancesRequest) (response *KillInstancesResponse, err error) {
    return c.KillInstancesWithContext(context.Background(), request)
}

// KillInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量终止操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillInstancesWithContext(ctx context.Context, request *KillInstancesRequest) (response *KillInstancesResponse, err error) {
    if request == nil {
        request = NewKillInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpTasksNewRequest() (request *MakeUpTasksNewRequest) {
    request = &MakeUpTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpTasksNew")
    
    
    return
}

func NewMakeUpTasksNewResponse() (response *MakeUpTasksNewResponse) {
    response = &MakeUpTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MakeUpTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 任务批量补录，调度状态任务才可以补录；
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksNew(request *MakeUpTasksNewRequest) (response *MakeUpTasksNewResponse, err error) {
    return c.MakeUpTasksNewWithContext(context.Background(), request)
}

// MakeUpTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 任务批量补录，调度状态任务才可以补录；
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksNewWithContext(ctx context.Context, request *MakeUpTasksNewRequest) (response *MakeUpTasksNewResponse, err error) {
    if request == nil {
        request = NewMakeUpTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpWorkflowNewRequest() (request *MakeUpWorkflowNewRequest) {
    request = &MakeUpWorkflowNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpWorkflowNew")
    
    
    return
}

func NewMakeUpWorkflowNewResponse() (response *MakeUpWorkflowNewResponse) {
    response = &MakeUpWorkflowNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MakeUpWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 工作流下所有任务的补录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpWorkflowNew(request *MakeUpWorkflowNewRequest) (response *MakeUpWorkflowNewResponse, err error) {
    return c.MakeUpWorkflowNewWithContext(context.Background(), request)
}

// MakeUpWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 工作流下所有任务的补录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpWorkflowNewWithContext(ctx context.Context, request *MakeUpWorkflowNewRequest) (response *MakeUpWorkflowNewResponse, err error) {
    if request == nil {
        request = NewMakeUpWorkflowNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpWorkflowNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpWorkflowNewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataSourceRequest() (request *ModifyDataSourceRequest) {
    request = &ModifyDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyDataSource")
    
    
    return
}

func NewModifyDataSourceResponse() (response *ModifyDataSourceResponse) {
    response = &ModifyDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDataSource(request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
    return c.ModifyDataSourceWithContext(context.Background(), request)
}

// ModifyDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDataSourceWithContext(ctx context.Context, request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
    if request == nil {
        request = NewModifyDataSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFolderRequest() (request *ModifyFolderRequest) {
    request = &ModifyFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyFolder")
    
    
    return
}

func NewModifyFolderResponse() (response *ModifyFolderResponse) {
    response = &ModifyFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyFolder(request *ModifyFolderRequest) (response *ModifyFolderResponse, err error) {
    return c.ModifyFolderWithContext(context.Background(), request)
}

// ModifyFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyFolderWithContext(ctx context.Context, request *ModifyFolderRequest) (response *ModifyFolderResponse, err error) {
    if request == nil {
        request = NewModifyFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFolderResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskInfoRequest() (request *ModifyTaskInfoRequest) {
    request = &ModifyTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskInfo")
    
    
    return
}

func NewModifyTaskInfoResponse() (response *ModifyTaskInfoResponse) {
    response = &ModifyTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTaskInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskInfo(request *ModifyTaskInfoRequest) (response *ModifyTaskInfoResponse, err error) {
    return c.ModifyTaskInfoWithContext(context.Background(), request)
}

// ModifyTaskInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskInfoWithContext(ctx context.Context, request *ModifyTaskInfoRequest) (response *ModifyTaskInfoResponse, err error) {
    if request == nil {
        request = NewModifyTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskLinksRequest() (request *ModifyTaskLinksRequest) {
    request = &ModifyTaskLinksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskLinks")
    
    
    return
}

func NewModifyTaskLinksResponse() (response *ModifyTaskLinksResponse) {
    response = &ModifyTaskLinksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTaskLinks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 添加父任务依赖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskLinks(request *ModifyTaskLinksRequest) (response *ModifyTaskLinksResponse, err error) {
    return c.ModifyTaskLinksWithContext(context.Background(), request)
}

// ModifyTaskLinks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 添加父任务依赖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskLinksWithContext(ctx context.Context, request *ModifyTaskLinksRequest) (response *ModifyTaskLinksResponse, err error) {
    if request == nil {
        request = NewModifyTaskLinksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskLinks require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskLinksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskScriptRequest() (request *ModifyTaskScriptRequest) {
    request = &ModifyTaskScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskScript")
    
    
    return
}

func NewModifyTaskScriptResponse() (response *ModifyTaskScriptResponse) {
    response = &ModifyTaskScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskScript(request *ModifyTaskScriptRequest) (response *ModifyTaskScriptResponse, err error) {
    return c.ModifyTaskScriptWithContext(context.Background(), request)
}

// ModifyTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskScriptWithContext(ctx context.Context, request *ModifyTaskScriptRequest) (response *ModifyTaskScriptResponse, err error) {
    if request == nil {
        request = NewModifyTaskScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskScriptResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkflowInfoRequest() (request *ModifyWorkflowInfoRequest) {
    request = &ModifyWorkflowInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyWorkflowInfo")
    
    
    return
}

func NewModifyWorkflowInfoResponse() (response *ModifyWorkflowInfoResponse) {
    response = &ModifyWorkflowInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWorkflowInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowInfo(request *ModifyWorkflowInfoRequest) (response *ModifyWorkflowInfoResponse, err error) {
    return c.ModifyWorkflowInfoWithContext(context.Background(), request)
}

// ModifyWorkflowInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowInfoWithContext(ctx context.Context, request *ModifyWorkflowInfoRequest) (response *ModifyWorkflowInfoResponse, err error) {
    if request == nil {
        request = NewModifyWorkflowInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkflowInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkflowInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkflowScheduleRequest() (request *ModifyWorkflowScheduleRequest) {
    request = &ModifyWorkflowScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyWorkflowSchedule")
    
    
    return
}

func NewModifyWorkflowScheduleResponse() (response *ModifyWorkflowScheduleResponse) {
    response = &ModifyWorkflowScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWorkflowSchedule
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowSchedule(request *ModifyWorkflowScheduleRequest) (response *ModifyWorkflowScheduleResponse, err error) {
    return c.ModifyWorkflowScheduleWithContext(context.Background(), request)
}

// ModifyWorkflowSchedule
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowScheduleWithContext(ctx context.Context, request *ModifyWorkflowScheduleRequest) (response *ModifyWorkflowScheduleResponse, err error) {
    if request == nil {
        request = NewModifyWorkflowScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkflowSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkflowScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterEventRequest() (request *RegisterEventRequest) {
    request = &RegisterEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RegisterEvent")
    
    
    return
}

func NewRegisterEventResponse() (response *RegisterEventResponse) {
    response = &RegisterEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEvent(request *RegisterEventRequest) (response *RegisterEventResponse, err error) {
    return c.RegisterEventWithContext(context.Background(), request)
}

// RegisterEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventWithContext(ctx context.Context, request *RegisterEventRequest) (response *RegisterEventResponse, err error) {
    if request == nil {
        request = NewRegisterEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterEventResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterEventListenerRequest() (request *RegisterEventListenerRequest) {
    request = &RegisterEventListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RegisterEventListener")
    
    
    return
}

func NewRegisterEventListenerResponse() (response *RegisterEventListenerResponse) {
    response = &RegisterEventListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterEventListener
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventListener(request *RegisterEventListenerRequest) (response *RegisterEventListenerResponse, err error) {
    return c.RegisterEventListenerWithContext(context.Background(), request)
}

// RegisterEventListener
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventListenerWithContext(ctx context.Context, request *RegisterEventListenerRequest) (response *RegisterEventListenerResponse, err error) {
    if request == nil {
        request = NewRegisterEventListenerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterEventListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterEventListenerResponse()
    err = c.Send(request, response)
    return
}

func NewRerunInstancesRequest() (request *RerunInstancesRequest) {
    request = &RerunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunInstances")
    
    
    return
}

func NewRerunInstancesResponse() (response *RerunInstancesResponse) {
    response = &RerunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RerunInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunInstances(request *RerunInstancesRequest) (response *RerunInstancesResponse, err error) {
    return c.RerunInstancesWithContext(context.Background(), request)
}

// RerunInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunInstancesWithContext(ctx context.Context, request *RerunInstancesRequest) (response *RerunInstancesResponse, err error) {
    if request == nil {
        request = NewRerunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRunTaskRequest() (request *RunTaskRequest) {
    request = &RunTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunTask")
    
    
    return
}

func NewRunTaskResponse() (response *RunTaskResponse) {
    response = &RunTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 运行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunTask(request *RunTaskRequest) (response *RunTaskResponse, err error) {
    return c.RunTaskWithContext(context.Background(), request)
}

// RunTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 运行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewSaveCustomFunctionRequest() (request *SaveCustomFunctionRequest) {
    request = &SaveCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SaveCustomFunction")
    
    
    return
}

func NewSaveCustomFunctionResponse() (response *SaveCustomFunctionResponse) {
    response = &SaveCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SaveCustomFunction
// 保存用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveCustomFunction(request *SaveCustomFunctionRequest) (response *SaveCustomFunctionResponse, err error) {
    return c.SaveCustomFunctionWithContext(context.Background(), request)
}

// SaveCustomFunction
// 保存用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveCustomFunctionWithContext(ctx context.Context, request *SaveCustomFunctionRequest) (response *SaveCustomFunctionResponse, err error) {
    if request == nil {
        request = NewSaveCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewSetTaskAlarmNewRequest() (request *SetTaskAlarmNewRequest) {
    request = &SetTaskAlarmNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SetTaskAlarmNew")
    
    
    return
}

func NewSetTaskAlarmNewResponse() (response *SetTaskAlarmNewResponse) {
    response = &SetTaskAlarmNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTaskAlarmNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 设置任务告警，新建/更新告警信息（最新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetTaskAlarmNew(request *SetTaskAlarmNewRequest) (response *SetTaskAlarmNewResponse, err error) {
    return c.SetTaskAlarmNewWithContext(context.Background(), request)
}

// SetTaskAlarmNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 设置任务告警，新建/更新告警信息（最新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetTaskAlarmNewWithContext(ctx context.Context, request *SetTaskAlarmNewRequest) (response *SetTaskAlarmNewResponse, err error) {
    if request == nil {
        request = NewSetTaskAlarmNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTaskAlarmNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTaskAlarmNewResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCustomFunctionRequest() (request *SubmitCustomFunctionRequest) {
    request = &SubmitCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitCustomFunction")
    
    
    return
}

func NewSubmitCustomFunctionResponse() (response *SubmitCustomFunctionResponse) {
    response = &SubmitCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitCustomFunction
// 提交自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitCustomFunction(request *SubmitCustomFunctionRequest) (response *SubmitCustomFunctionResponse, err error) {
    return c.SubmitCustomFunctionWithContext(context.Background(), request)
}

// SubmitCustomFunction
// 提交自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitCustomFunctionWithContext(ctx context.Context, request *SubmitCustomFunctionRequest) (response *SubmitCustomFunctionResponse, err error) {
    if request == nil {
        request = NewSubmitCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTaskRequest() (request *SubmitTaskRequest) {
    request = &SubmitTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTask")
    
    
    return
}

func NewSubmitTaskResponse() (response *SubmitTaskResponse) {
    response = &SubmitTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitTask(request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    return c.SubmitTaskWithContext(context.Background(), request)
}

// SubmitTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitTaskWithContext(ctx context.Context, request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitWorkflowRequest() (request *SubmitWorkflowRequest) {
    request = &SubmitWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitWorkflow")
    
    
    return
}

func NewSubmitWorkflowResponse() (response *SubmitWorkflowResponse) {
    response = &SubmitWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitWorkflow(request *SubmitWorkflowRequest) (response *SubmitWorkflowResponse, err error) {
    return c.SubmitWorkflowWithContext(context.Background(), request)
}

// SubmitWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitWorkflowWithContext(ctx context.Context, request *SubmitWorkflowRequest) (response *SubmitWorkflowResponse, err error) {
    if request == nil {
        request = NewSubmitWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerEventRequest() (request *TriggerEventRequest) {
    request = &TriggerEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "TriggerEvent")
    
    
    return
}

func NewTriggerEventResponse() (response *TriggerEventResponse) {
    response = &TriggerEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TriggerEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 触发事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) TriggerEvent(request *TriggerEventRequest) (response *TriggerEventResponse, err error) {
    return c.TriggerEventWithContext(context.Background(), request)
}

// TriggerEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 触发事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) TriggerEventWithContext(ctx context.Context, request *TriggerEventRequest) (response *TriggerEventResponse, err error) {
    if request == nil {
        request = NewTriggerEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TriggerEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewTriggerEventResponse()
    err = c.Send(request, response)
    return
}
