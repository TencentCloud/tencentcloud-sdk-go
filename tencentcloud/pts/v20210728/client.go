// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20210728

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-28"

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


func NewAbortCronJobsRequest() (request *AbortCronJobsRequest) {
    request = &AbortCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AbortCronJobs")
    
    
    return
}

func NewAbortCronJobsResponse() (response *AbortCronJobsResponse) {
    response = &AbortCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortCronJobs
// 停止定时任务
func (c *Client) AbortCronJobs(request *AbortCronJobsRequest) (response *AbortCronJobsResponse, err error) {
    return c.AbortCronJobsWithContext(context.Background(), request)
}

// AbortCronJobs
// 停止定时任务
func (c *Client) AbortCronJobsWithContext(ctx context.Context, request *AbortCronJobsRequest) (response *AbortCronJobsResponse, err error) {
    if request == nil {
        request = NewAbortCronJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "AbortCronJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewAbortJobRequest() (request *AbortJobRequest) {
    request = &AbortJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AbortJob")
    
    
    return
}

func NewAbortJobResponse() (response *AbortJobResponse) {
    response = &AbortJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortJob
// 停止任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_JOBSTATUSNOTRUNNING = "FailedOperation.JobStatusNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AbortJob(request *AbortJobRequest) (response *AbortJobResponse, err error) {
    return c.AbortJobWithContext(context.Background(), request)
}

// AbortJob
// 停止任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_JOBSTATUSNOTRUNNING = "FailedOperation.JobStatusNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AbortJobWithContext(ctx context.Context, request *AbortJobRequest) (response *AbortJobResponse, err error) {
    if request == nil {
        request = NewAbortJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "AbortJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortJobResponse()
    err = c.Send(request, response)
    return
}

func NewAdjustJobSpeedRequest() (request *AdjustJobSpeedRequest) {
    request = &AdjustJobSpeedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "AdjustJobSpeed")
    
    
    return
}

func NewAdjustJobSpeedResponse() (response *AdjustJobSpeedResponse) {
    response = &AdjustJobSpeedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AdjustJobSpeed
// 调整任务的目标RPS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AdjustJobSpeed(request *AdjustJobSpeedRequest) (response *AdjustJobSpeedResponse, err error) {
    return c.AdjustJobSpeedWithContext(context.Background(), request)
}

// AdjustJobSpeed
// 调整任务的目标RPS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AdjustJobSpeedWithContext(ctx context.Context, request *AdjustJobSpeedRequest) (response *AdjustJobSpeedResponse, err error) {
    if request == nil {
        request = NewAdjustJobSpeedRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "AdjustJobSpeed")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdjustJobSpeed require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdjustJobSpeedResponse()
    err = c.Send(request, response)
    return
}

func NewCopyScenarioRequest() (request *CopyScenarioRequest) {
    request = &CopyScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CopyScenario")
    
    
    return
}

func NewCopyScenarioResponse() (response *CopyScenarioResponse) {
    response = &CopyScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyScenario
// 复制场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) CopyScenario(request *CopyScenarioRequest) (response *CopyScenarioResponse, err error) {
    return c.CopyScenarioWithContext(context.Background(), request)
}

// CopyScenario
// 复制场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) CopyScenarioWithContext(ctx context.Context, request *CopyScenarioRequest) (response *CopyScenarioResponse, err error) {
    if request == nil {
        request = NewCopyScenarioRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CopyScenario")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyScenarioResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertChannelRequest() (request *CreateAlertChannelRequest) {
    request = &CreateAlertChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateAlertChannel")
    
    
    return
}

func NewCreateAlertChannelResponse() (response *CreateAlertChannelResponse) {
    response = &CreateAlertChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertChannel
// 创建告警通知接收组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateAlertChannel(request *CreateAlertChannelRequest) (response *CreateAlertChannelResponse, err error) {
    return c.CreateAlertChannelWithContext(context.Background(), request)
}

// CreateAlertChannel
// 创建告警通知接收组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateAlertChannelWithContext(ctx context.Context, request *CreateAlertChannelRequest) (response *CreateAlertChannelResponse, err error) {
    if request == nil {
        request = NewCreateAlertChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateAlertChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCronJobRequest() (request *CreateCronJobRequest) {
    request = &CreateCronJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateCronJob")
    
    
    return
}

func NewCreateCronJobResponse() (response *CreateCronJobResponse) {
    response = &CreateCronJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCronJob
// 创建定时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateCronJob(request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
    return c.CreateCronJobWithContext(context.Background(), request)
}

// CreateCronJob
// 创建定时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateCronJobWithContext(ctx context.Context, request *CreateCronJobRequest) (response *CreateCronJobResponse, err error) {
    if request == nil {
        request = NewCreateCronJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateCronJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCronJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCronJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateEnvironment")
    
    
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateEnvironment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileRequest() (request *CreateFileRequest) {
    request = &CreateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateFile")
    
    
    return
}

func NewCreateFileResponse() (response *CreateFileResponse) {
    response = &CreateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFile
// 创建文件
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateFile(request *CreateFileRequest) (response *CreateFileResponse, err error) {
    return c.CreateFileWithContext(context.Background(), request)
}

// CreateFile
// 创建文件
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
func (c *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest) (response *CreateFileResponse, err error) {
    if request == nil {
        request = NewCreateFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScenarioRequest() (request *CreateScenarioRequest) {
    request = &CreateScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "CreateScenario")
    
    
    return
}

func NewCreateScenarioResponse() (response *CreateScenarioResponse) {
    response = &CreateScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScenario
// 创建场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateScenario(request *CreateScenarioRequest) (response *CreateScenarioResponse, err error) {
    return c.CreateScenarioWithContext(context.Background(), request)
}

// CreateScenario
// 创建场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateScenarioWithContext(ctx context.Context, request *CreateScenarioRequest) (response *CreateScenarioResponse, err error) {
    if request == nil {
        request = NewCreateScenarioRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "CreateScenario")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScenarioResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlertChannelRequest() (request *DeleteAlertChannelRequest) {
    request = &DeleteAlertChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteAlertChannel")
    
    
    return
}

func NewDeleteAlertChannelResponse() (response *DeleteAlertChannelResponse) {
    response = &DeleteAlertChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlertChannel
// 删除告警通知接收组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAlertChannel(request *DeleteAlertChannelRequest) (response *DeleteAlertChannelResponse, err error) {
    return c.DeleteAlertChannelWithContext(context.Background(), request)
}

// DeleteAlertChannel
// 删除告警通知接收组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteAlertChannelWithContext(ctx context.Context, request *DeleteAlertChannelRequest) (response *DeleteAlertChannelResponse, err error) {
    if request == nil {
        request = NewDeleteAlertChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteAlertChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlertChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlertChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCronJobsRequest() (request *DeleteCronJobsRequest) {
    request = &DeleteCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteCronJobs")
    
    
    return
}

func NewDeleteCronJobsResponse() (response *DeleteCronJobsResponse) {
    response = &DeleteCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCronJobs
// 删除定时任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCronJobs(request *DeleteCronJobsRequest) (response *DeleteCronJobsResponse, err error) {
    return c.DeleteCronJobsWithContext(context.Background(), request)
}

// DeleteCronJobs
// 删除定时任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteCronJobsWithContext(ctx context.Context, request *DeleteCronJobsRequest) (response *DeleteCronJobsResponse, err error) {
    if request == nil {
        request = NewDeleteCronJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteCronJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentsRequest() (request *DeleteEnvironmentsRequest) {
    request = &DeleteEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteEnvironments")
    
    
    return
}

func NewDeleteEnvironmentsResponse() (response *DeleteEnvironmentsResponse) {
    response = &DeleteEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnvironments
// 删除环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    return c.DeleteEnvironmentsWithContext(context.Background(), request)
}

// DeleteEnvironments
// 删除环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteEnvironmentsWithContext(ctx context.Context, request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFilesRequest() (request *DeleteFilesRequest) {
    request = &DeleteFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteFiles")
    
    
    return
}

func NewDeleteFilesResponse() (response *DeleteFilesResponse) {
    response = &DeleteFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFiles
// 删除文件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFiles(request *DeleteFilesRequest) (response *DeleteFilesResponse, err error) {
    return c.DeleteFilesWithContext(context.Background(), request)
}

// DeleteFiles
// 删除文件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFilesWithContext(ctx context.Context, request *DeleteFilesRequest) (response *DeleteFilesResponse, err error) {
    if request == nil {
        request = NewDeleteFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobsRequest() (request *DeleteJobsRequest) {
    request = &DeleteJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteJobs")
    
    
    return
}

func NewDeleteJobsResponse() (response *DeleteJobsResponse) {
    response = &DeleteJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJobs
// 删除任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobs(request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    return c.DeleteJobsWithContext(context.Background(), request)
}

// DeleteJobs
// 删除任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobsWithContext(ctx context.Context, request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    if request == nil {
        request = NewDeleteJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectsRequest() (request *DeleteProjectsRequest) {
    request = &DeleteProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteProjects")
    
    
    return
}

func NewDeleteProjectsResponse() (response *DeleteProjectsResponse) {
    response = &DeleteProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjects
// 删除项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProjects(request *DeleteProjectsRequest) (response *DeleteProjectsResponse, err error) {
    return c.DeleteProjectsWithContext(context.Background(), request)
}

// DeleteProjects
// 删除项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteProjectsWithContext(ctx context.Context, request *DeleteProjectsRequest) (response *DeleteProjectsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScenariosRequest() (request *DeleteScenariosRequest) {
    request = &DeleteScenariosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DeleteScenarios")
    
    
    return
}

func NewDeleteScenariosResponse() (response *DeleteScenariosResponse) {
    response = &DeleteScenariosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScenarios
// 删除场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteScenarios(request *DeleteScenariosRequest) (response *DeleteScenariosResponse, err error) {
    return c.DeleteScenariosWithContext(context.Background(), request)
}

// DeleteScenarios
// 删除场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteScenariosWithContext(ctx context.Context, request *DeleteScenariosRequest) (response *DeleteScenariosResponse, err error) {
    if request == nil {
        request = NewDeleteScenariosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DeleteScenarios")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScenarios require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScenariosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertChannelsRequest() (request *DescribeAlertChannelsRequest) {
    request = &DescribeAlertChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAlertChannels")
    
    
    return
}

func NewDescribeAlertChannelsResponse() (response *DescribeAlertChannelsResponse) {
    response = &DescribeAlertChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertChannels
// 查询告警通知接收组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertChannels(request *DescribeAlertChannelsRequest) (response *DescribeAlertChannelsResponse, err error) {
    return c.DescribeAlertChannelsWithContext(context.Background(), request)
}

// DescribeAlertChannels
// 查询告警通知接收组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertChannelsWithContext(ctx context.Context, request *DescribeAlertChannelsRequest) (response *DescribeAlertChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeAlertChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeAlertChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRecordsRequest() (request *DescribeAlertRecordsRequest) {
    request = &DescribeAlertRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAlertRecords")
    
    
    return
}

func NewDescribeAlertRecordsResponse() (response *DescribeAlertRecordsResponse) {
    response = &DescribeAlertRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRecords
// 返回告警历史项的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertRecords(request *DescribeAlertRecordsRequest) (response *DescribeAlertRecordsResponse, err error) {
    return c.DescribeAlertRecordsWithContext(context.Background(), request)
}

// DescribeAlertRecords
// 返回告警历史项的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAlertRecordsWithContext(ctx context.Context, request *DescribeAlertRecordsRequest) (response *DescribeAlertRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeAlertRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableMetricsRequest() (request *DescribeAvailableMetricsRequest) {
    request = &DescribeAvailableMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeAvailableMetrics")
    
    
    return
}

func NewDescribeAvailableMetricsResponse() (response *DescribeAvailableMetricsResponse) {
    response = &DescribeAvailableMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableMetrics
// 查询系统支持的指标
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAvailableMetrics(request *DescribeAvailableMetricsRequest) (response *DescribeAvailableMetricsResponse, err error) {
    return c.DescribeAvailableMetricsWithContext(context.Background(), request)
}

// DescribeAvailableMetrics
// 查询系统支持的指标
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeAvailableMetricsWithContext(ctx context.Context, request *DescribeAvailableMetricsRequest) (response *DescribeAvailableMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeAvailableMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckSummaryRequest() (request *DescribeCheckSummaryRequest) {
    request = &DescribeCheckSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeCheckSummary")
    
    
    return
}

func NewDescribeCheckSummaryResponse() (response *DescribeCheckSummaryResponse) {
    response = &DescribeCheckSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckSummary
// 查询检查点汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCheckSummary(request *DescribeCheckSummaryRequest) (response *DescribeCheckSummaryResponse, err error) {
    return c.DescribeCheckSummaryWithContext(context.Background(), request)
}

// DescribeCheckSummary
// 查询检查点汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCheckSummaryWithContext(ctx context.Context, request *DescribeCheckSummaryRequest) (response *DescribeCheckSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCheckSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeCheckSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCronJobsRequest() (request *DescribeCronJobsRequest) {
    request = &DescribeCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeCronJobs")
    
    
    return
}

func NewDescribeCronJobsResponse() (response *DescribeCronJobsResponse) {
    response = &DescribeCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCronJobs
// 列出定时任务，非必填数组为空就默认全选
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeCronJobs(request *DescribeCronJobsRequest) (response *DescribeCronJobsResponse, err error) {
    return c.DescribeCronJobsWithContext(context.Background(), request)
}

// DescribeCronJobs
// 列出定时任务，非必填数组为空就默认全选
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) DescribeCronJobsWithContext(ctx context.Context, request *DescribeCronJobsRequest) (response *DescribeCronJobsResponse, err error) {
    if request == nil {
        request = NewDescribeCronJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeCronJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// 查看环境列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// 查看环境列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorSummaryRequest() (request *DescribeErrorSummaryRequest) {
    request = &DescribeErrorSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeErrorSummary")
    
    
    return
}

func NewDescribeErrorSummaryResponse() (response *DescribeErrorSummaryResponse) {
    response = &DescribeErrorSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeErrorSummary
// 查询错误详情汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeErrorSummary(request *DescribeErrorSummaryRequest) (response *DescribeErrorSummaryResponse, err error) {
    return c.DescribeErrorSummaryWithContext(context.Background(), request)
}

// DescribeErrorSummary
// 查询错误详情汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeErrorSummaryWithContext(ctx context.Context, request *DescribeErrorSummaryRequest) (response *DescribeErrorSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeErrorSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeErrorSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilesRequest() (request *DescribeFilesRequest) {
    request = &DescribeFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeFiles")
    
    
    return
}

func NewDescribeFilesResponse() (response *DescribeFilesResponse) {
    response = &DescribeFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFiles
// 查询文件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFiles(request *DescribeFilesRequest) (response *DescribeFilesResponse, err error) {
    return c.DescribeFilesWithContext(context.Background(), request)
}

// DescribeFiles
// 查询文件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFilesWithContext(ctx context.Context, request *DescribeFilesRequest) (response *DescribeFilesResponse, err error) {
    if request == nil {
        request = NewDescribeFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// 查询任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// 查询任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLabelValuesRequest() (request *DescribeLabelValuesRequest) {
    request = &DescribeLabelValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeLabelValues")
    
    
    return
}

func NewDescribeLabelValuesResponse() (response *DescribeLabelValuesResponse) {
    response = &DescribeLabelValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLabelValues
// 查询标签内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLabelValues(request *DescribeLabelValuesRequest) (response *DescribeLabelValuesResponse, err error) {
    return c.DescribeLabelValuesWithContext(context.Background(), request)
}

// DescribeLabelValues
// 查询标签内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLabelValuesWithContext(ctx context.Context, request *DescribeLabelValuesRequest) (response *DescribeLabelValuesResponse, err error) {
    if request == nil {
        request = NewDescribeLabelValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeLabelValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLabelValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLabelValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricLabelWithValuesRequest() (request *DescribeMetricLabelWithValuesRequest) {
    request = &DescribeMetricLabelWithValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeMetricLabelWithValues")
    
    
    return
}

func NewDescribeMetricLabelWithValuesResponse() (response *DescribeMetricLabelWithValuesResponse) {
    response = &DescribeMetricLabelWithValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricLabelWithValues
// 查询指标所有的label及values值
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricLabelWithValues(request *DescribeMetricLabelWithValuesRequest) (response *DescribeMetricLabelWithValuesResponse, err error) {
    return c.DescribeMetricLabelWithValuesWithContext(context.Background(), request)
}

// DescribeMetricLabelWithValues
// 查询指标所有的label及values值
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricLabelWithValuesWithContext(ctx context.Context, request *DescribeMetricLabelWithValuesRequest) (response *DescribeMetricLabelWithValuesResponse, err error) {
    if request == nil {
        request = NewDescribeMetricLabelWithValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeMetricLabelWithValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricLabelWithValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricLabelWithValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNormalLogsRequest() (request *DescribeNormalLogsRequest) {
    request = &DescribeNormalLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeNormalLogs")
    
    
    return
}

func NewDescribeNormalLogsResponse() (response *DescribeNormalLogsResponse) {
    response = &DescribeNormalLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNormalLogs
// 压测过程日志包括引擎输出日志及用户输出日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNormalLogs(request *DescribeNormalLogsRequest) (response *DescribeNormalLogsResponse, err error) {
    return c.DescribeNormalLogsWithContext(context.Background(), request)
}

// DescribeNormalLogs
// 压测过程日志包括引擎输出日志及用户输出日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNormalLogsWithContext(ctx context.Context, request *DescribeNormalLogsRequest) (response *DescribeNormalLogsResponse, err error) {
    if request == nil {
        request = NewDescribeNormalLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeNormalLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNormalLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNormalLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjects
// 查询项目列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// 查询项目列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegions
// 查询地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 查询地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRequestSummaryRequest() (request *DescribeRequestSummaryRequest) {
    request = &DescribeRequestSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeRequestSummary")
    
    
    return
}

func NewDescribeRequestSummaryResponse() (response *DescribeRequestSummaryResponse) {
    response = &DescribeRequestSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRequestSummary
// 查询请求汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRequestSummary(request *DescribeRequestSummaryRequest) (response *DescribeRequestSummaryResponse, err error) {
    return c.DescribeRequestSummaryWithContext(context.Background(), request)
}

// DescribeRequestSummary
// 查询请求汇总信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRequestSummaryWithContext(ctx context.Context, request *DescribeRequestSummaryRequest) (response *DescribeRequestSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeRequestSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeRequestSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRequestSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRequestSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleBatchQueryRequest() (request *DescribeSampleBatchQueryRequest) {
    request = &DescribeSampleBatchQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleBatchQuery")
    
    
    return
}

func NewDescribeSampleBatchQueryResponse() (response *DescribeSampleBatchQueryResponse) {
    response = &DescribeSampleBatchQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleBatchQuery
// 批量查询指标，返回固定时间点指标内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleBatchQuery(request *DescribeSampleBatchQueryRequest) (response *DescribeSampleBatchQueryResponse, err error) {
    return c.DescribeSampleBatchQueryWithContext(context.Background(), request)
}

// DescribeSampleBatchQuery
// 批量查询指标，返回固定时间点指标内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleBatchQueryWithContext(ctx context.Context, request *DescribeSampleBatchQueryRequest) (response *DescribeSampleBatchQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleBatchQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeSampleBatchQuery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleBatchQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleBatchQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleLogsRequest() (request *DescribeSampleLogsRequest) {
    request = &DescribeSampleLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleLogs")
    
    
    return
}

func NewDescribeSampleLogsResponse() (response *DescribeSampleLogsResponse) {
    response = &DescribeSampleLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleLogs
// 查询采样日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleLogs(request *DescribeSampleLogsRequest) (response *DescribeSampleLogsResponse, err error) {
    return c.DescribeSampleLogsWithContext(context.Background(), request)
}

// DescribeSampleLogs
// 查询采样日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleLogsWithContext(ctx context.Context, request *DescribeSampleLogsRequest) (response *DescribeSampleLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSampleLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeSampleLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleMatrixBatchQueryRequest() (request *DescribeSampleMatrixBatchQueryRequest) {
    request = &DescribeSampleMatrixBatchQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleMatrixBatchQuery")
    
    
    return
}

func NewDescribeSampleMatrixBatchQueryResponse() (response *DescribeSampleMatrixBatchQueryResponse) {
    response = &DescribeSampleMatrixBatchQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleMatrixBatchQuery
// 批量查询指标矩阵
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleMatrixBatchQuery(request *DescribeSampleMatrixBatchQueryRequest) (response *DescribeSampleMatrixBatchQueryResponse, err error) {
    return c.DescribeSampleMatrixBatchQueryWithContext(context.Background(), request)
}

// DescribeSampleMatrixBatchQuery
// 批量查询指标矩阵
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSampleMatrixBatchQueryWithContext(ctx context.Context, request *DescribeSampleMatrixBatchQueryRequest) (response *DescribeSampleMatrixBatchQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleMatrixBatchQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeSampleMatrixBatchQuery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleMatrixBatchQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleMatrixBatchQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleMatrixQueryRequest() (request *DescribeSampleMatrixQueryRequest) {
    request = &DescribeSampleMatrixQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleMatrixQuery")
    
    
    return
}

func NewDescribeSampleMatrixQueryResponse() (response *DescribeSampleMatrixQueryResponse) {
    response = &DescribeSampleMatrixQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleMatrixQuery
// 查询指标矩阵
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSampleMatrixQuery(request *DescribeSampleMatrixQueryRequest) (response *DescribeSampleMatrixQueryResponse, err error) {
    return c.DescribeSampleMatrixQueryWithContext(context.Background(), request)
}

// DescribeSampleMatrixQuery
// 查询指标矩阵
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSampleMatrixQueryWithContext(ctx context.Context, request *DescribeSampleMatrixQueryRequest) (response *DescribeSampleMatrixQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleMatrixQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeSampleMatrixQuery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleMatrixQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleMatrixQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSampleQueryRequest() (request *DescribeSampleQueryRequest) {
    request = &DescribeSampleQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeSampleQuery")
    
    
    return
}

func NewDescribeSampleQueryResponse() (response *DescribeSampleQueryResponse) {
    response = &DescribeSampleQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSampleQuery
// 查询指标，返回固定时间点指标内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleQuery(request *DescribeSampleQueryRequest) (response *DescribeSampleQueryResponse, err error) {
    return c.DescribeSampleQueryWithContext(context.Background(), request)
}

// DescribeSampleQuery
// 查询指标，返回固定时间点指标内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSampleQueryWithContext(ctx context.Context, request *DescribeSampleQueryRequest) (response *DescribeSampleQueryResponse, err error) {
    if request == nil {
        request = NewDescribeSampleQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeSampleQuery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSampleQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSampleQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenarioWithJobsRequest() (request *DescribeScenarioWithJobsRequest) {
    request = &DescribeScenarioWithJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeScenarioWithJobs")
    
    
    return
}

func NewDescribeScenarioWithJobsResponse() (response *DescribeScenarioWithJobsResponse) {
    response = &DescribeScenarioWithJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScenarioWithJobs
// 查询场景配置并附带已经执行的任务内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarioWithJobs(request *DescribeScenarioWithJobsRequest) (response *DescribeScenarioWithJobsResponse, err error) {
    return c.DescribeScenarioWithJobsWithContext(context.Background(), request)
}

// DescribeScenarioWithJobs
// 查询场景配置并附带已经执行的任务内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarioWithJobsWithContext(ctx context.Context, request *DescribeScenarioWithJobsRequest) (response *DescribeScenarioWithJobsResponse, err error) {
    if request == nil {
        request = NewDescribeScenarioWithJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeScenarioWithJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenarioWithJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenarioWithJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenariosRequest() (request *DescribeScenariosRequest) {
    request = &DescribeScenariosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "DescribeScenarios")
    
    
    return
}

func NewDescribeScenariosResponse() (response *DescribeScenariosResponse) {
    response = &DescribeScenariosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScenarios
// 查询场景列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenarios(request *DescribeScenariosRequest) (response *DescribeScenariosResponse, err error) {
    return c.DescribeScenariosWithContext(context.Background(), request)
}

// DescribeScenarios
// 查询场景列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeScenariosWithContext(ctx context.Context, request *DescribeScenariosRequest) (response *DescribeScenariosResponse, err error) {
    if request == nil {
        request = NewDescribeScenariosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "DescribeScenarios")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenarios require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenariosResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateTmpKeyRequest() (request *GenerateTmpKeyRequest) {
    request = &GenerateTmpKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "GenerateTmpKey")
    
    
    return
}

func NewGenerateTmpKeyResponse() (response *GenerateTmpKeyResponse) {
    response = &GenerateTmpKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateTmpKey
// 生成临时COS凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GenerateTmpKey(request *GenerateTmpKeyRequest) (response *GenerateTmpKeyResponse, err error) {
    return c.GenerateTmpKeyWithContext(context.Background(), request)
}

// GenerateTmpKey
// 生成临时COS凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GenerateTmpKeyWithContext(ctx context.Context, request *GenerateTmpKeyRequest) (response *GenerateTmpKeyResponse, err error) {
    if request == nil {
        request = NewGenerateTmpKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "GenerateTmpKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateTmpKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateTmpKeyResponse()
    err = c.Send(request, response)
    return
}

func NewRestartCronJobsRequest() (request *RestartCronJobsRequest) {
    request = &RestartCronJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "RestartCronJobs")
    
    
    return
}

func NewRestartCronJobsResponse() (response *RestartCronJobsResponse) {
    response = &RestartCronJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartCronJobs
// 重启状态为已中止的定时任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RestartCronJobs(request *RestartCronJobsRequest) (response *RestartCronJobsResponse, err error) {
    return c.RestartCronJobsWithContext(context.Background(), request)
}

// RestartCronJobs
// 重启状态为已中止的定时任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RestartCronJobsWithContext(ctx context.Context, request *RestartCronJobsRequest) (response *RestartCronJobsResponse, err error) {
    if request == nil {
        request = NewRestartCronJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "RestartCronJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartCronJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartCronJobsResponse()
    err = c.Send(request, response)
    return
}

func NewStartJobRequest() (request *StartJobRequest) {
    request = &StartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "StartJob")
    
    
    return
}

func NewStartJobResponse() (response *StartJobResponse) {
    response = &StartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartJob
// 创建并启动任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOTASKSINJOB = "FailedOperation.NoTasksInJob"
//  FAILEDOPERATION_NOTSUPPORTEDINENV = "FailedOperation.NotSupportedInEnv"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) StartJob(request *StartJobRequest) (response *StartJobResponse, err error) {
    return c.StartJobWithContext(context.Background(), request)
}

// StartJob
// 创建并启动任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_NOTASKSINJOB = "FailedOperation.NoTasksInJob"
//  FAILEDOPERATION_NOTSUPPORTEDINENV = "FailedOperation.NotSupportedInEnv"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) StartJobWithContext(ctx context.Context, request *StartJobRequest) (response *StartJobResponse, err error) {
    if request == nil {
        request = NewStartJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "StartJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCronJobRequest() (request *UpdateCronJobRequest) {
    request = &UpdateCronJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateCronJob")
    
    
    return
}

func NewUpdateCronJobResponse() (response *UpdateCronJobResponse) {
    response = &UpdateCronJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCronJob
// 更新定时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateCronJob(request *UpdateCronJobRequest) (response *UpdateCronJobResponse, err error) {
    return c.UpdateCronJobWithContext(context.Background(), request)
}

// UpdateCronJob
// 更新定时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateCronJobWithContext(ctx context.Context, request *UpdateCronJobRequest) (response *UpdateCronJobResponse, err error) {
    if request == nil {
        request = NewUpdateCronJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateCronJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCronJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCronJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEnvironmentRequest() (request *UpdateEnvironmentRequest) {
    request = &UpdateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateEnvironment")
    
    
    return
}

func NewUpdateEnvironmentResponse() (response *UpdateEnvironmentResponse) {
    response = &UpdateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateEnvironment
// 更新环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateEnvironment(request *UpdateEnvironmentRequest) (response *UpdateEnvironmentResponse, err error) {
    return c.UpdateEnvironmentWithContext(context.Background(), request)
}

// UpdateEnvironment
// 更新环境
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateEnvironmentWithContext(ctx context.Context, request *UpdateEnvironmentRequest) (response *UpdateEnvironmentResponse, err error) {
    if request == nil {
        request = NewUpdateEnvironmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateEnvironment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFileScenarioRelationRequest() (request *UpdateFileScenarioRelationRequest) {
    request = &UpdateFileScenarioRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateFileScenarioRelation")
    
    
    return
}

func NewUpdateFileScenarioRelationResponse() (response *UpdateFileScenarioRelationResponse) {
    response = &UpdateFileScenarioRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFileScenarioRelation
// 更新关联文件场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) UpdateFileScenarioRelation(request *UpdateFileScenarioRelationRequest) (response *UpdateFileScenarioRelationResponse, err error) {
    return c.UpdateFileScenarioRelationWithContext(context.Background(), request)
}

// UpdateFileScenarioRelation
// 更新关联文件场景
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) UpdateFileScenarioRelationWithContext(ctx context.Context, request *UpdateFileScenarioRelationRequest) (response *UpdateFileScenarioRelationResponse, err error) {
    if request == nil {
        request = NewUpdateFileScenarioRelationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateFileScenarioRelation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFileScenarioRelation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFileScenarioRelationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJobRequest() (request *UpdateJobRequest) {
    request = &UpdateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateJob")
    
    
    return
}

func NewUpdateJobResponse() (response *UpdateJobResponse) {
    response = &UpdateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateJob
// 更新任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    return c.UpdateJobWithContext(context.Background(), request)
}

// UpdateJob
// 更新任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateJobWithContext(ctx context.Context, request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    if request == nil {
        request = NewUpdateJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProjectRequest() (request *UpdateProjectRequest) {
    request = &UpdateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateProject")
    
    
    return
}

func NewUpdateProjectResponse() (response *UpdateProjectResponse) {
    response = &UpdateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateProject
// 更新项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    return c.UpdateProjectWithContext(context.Background(), request)
}

// UpdateProject
// 更新项目
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    if request == nil {
        request = NewUpdateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScenarioRequest() (request *UpdateScenarioRequest) {
    request = &UpdateScenarioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pts", APIVersion, "UpdateScenario")
    
    
    return
}

func NewUpdateScenarioResponse() (response *UpdateScenarioResponse) {
    response = &UpdateScenarioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScenario
// 更新场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateScenario(request *UpdateScenarioRequest) (response *UpdateScenarioResponse, err error) {
    return c.UpdateScenarioWithContext(context.Background(), request)
}

// UpdateScenario
// 更新场景
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateScenarioWithContext(ctx context.Context, request *UpdateScenarioRequest) (response *UpdateScenarioResponse, err error) {
    if request == nil {
        request = NewUpdateScenarioRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "pts", APIVersion, "UpdateScenario")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScenario require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScenarioResponse()
    err = c.Send(request, response)
    return
}
