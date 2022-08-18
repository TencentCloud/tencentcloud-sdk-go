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

package v20180226

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-26"

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


func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "CreateJob")
    
    
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateJob
// 创建训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    return c.CreateJobWithContext(context.Background(), request)
}

// CreateJob
// 创建训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelRequest() (request *CreateModelRequest) {
    request = &CreateModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "CreateModel")
    
    
    return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
    response = &CreateModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateModel
// 部署模型，用以对外提供服务。有两种部署模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"
//  LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"
//  LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"
//  MISSINGPARAMETER_CODE = "MissingParameter.Code"
//  RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) CreateModel(request *CreateModelRequest) (response *CreateModelResponse, err error) {
    return c.CreateModelWithContext(context.Background(), request)
}

// CreateModel
// 部署模型，用以对外提供服务。有两种部署模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"
//  INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"
//  INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"
//  INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"
//  LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"
//  LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"
//  LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"
//  MISSINGPARAMETER_CODE = "MissingParameter.Code"
//  RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest) (response *CreateModelResponse, err error) {
    if request == nil {
        request = NewCreateModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "DeleteJob")
    
    
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteJob
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    return c.DeleteJobWithContext(context.Background(), request)
}

// DeleteJob
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelRequest() (request *DeleteModelRequest) {
    request = &DeleteModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "DeleteModel")
    
    
    return
}

func NewDeleteModelResponse() (response *DeleteModelResponse) {
    response = &DeleteModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModel
// 删除指定的部署模型。模型有两种部署模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) DeleteModel(request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    return c.DeleteModelWithContext(context.Background(), request)
}

// DeleteModel
// 删除指定的部署模型。模型有两种部署模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    if request == nil {
        request = NewDeleteModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
    request = &DescribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "DescribeJob")
    
    
    return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
    response = &DescribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJob
// 获取训练任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    return c.DescribeJobWithContext(context.Background(), request)
}

// DescribeJob
// 获取训练任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    if request == nil {
        request = NewDescribeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelRequest() (request *DescribeModelRequest) {
    request = &DescribeModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "DescribeModel")
    
    
    return
}

func NewDescribeModelResponse() (response *DescribeModelResponse) {
    response = &DescribeModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModel
// 描述已经部署的某个模型。而模型部署有两种模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) DescribeModel(request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    return c.DescribeModelWithContext(context.Background(), request)
}

// DescribeModel
// 描述已经部署的某个模型。而模型部署有两种模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) DescribeModelWithContext(ctx context.Context, request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    if request == nil {
        request = NewDescribeModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelResponse()
    err = c.Send(request, response)
    return
}

func NewInstallAgentRequest() (request *InstallAgentRequest) {
    request = &InstallAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "InstallAgent")
    
    
    return
}

func NewInstallAgentResponse() (response *InstallAgentResponse) {
    response = &InstallAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InstallAgent
// 安装agent
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) InstallAgent(request *InstallAgentRequest) (response *InstallAgentResponse, err error) {
    return c.InstallAgentWithContext(context.Background(), request)
}

// InstallAgent
// 安装agent
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) InstallAgentWithContext(ctx context.Context, request *InstallAgentRequest) (response *InstallAgentResponse, err error) {
    if request == nil {
        request = NewInstallAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallAgentResponse()
    err = c.Send(request, response)
    return
}

func NewListJobsRequest() (request *ListJobsRequest) {
    request = &ListJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "ListJobs")
    
    
    return
}

func NewListJobsResponse() (response *ListJobsResponse) {
    response = &ListJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListJobs
// 列举训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) ListJobs(request *ListJobsRequest) (response *ListJobsResponse, err error) {
    return c.ListJobsWithContext(context.Background(), request)
}

// ListJobs
// 列举训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest) (response *ListJobsResponse, err error) {
    if request == nil {
        request = NewListJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListJobsResponse()
    err = c.Send(request, response)
    return
}

func NewListModelsRequest() (request *ListModelsRequest) {
    request = &ListModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "ListModels")
    
    
    return
}

func NewListModelsResponse() (response *ListModelsResponse) {
    response = &ListModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListModels
// 用以列举已经部署的模型。而部署有两种模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。不同部署模式下的模型分开列出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.Orderby"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) ListModels(request *ListModelsRequest) (response *ListModelsResponse, err error) {
    return c.ListModelsWithContext(context.Background(), request)
}

// ListModels
// 用以列举已经部署的模型。而部署有两种模式：`无服务器模式` 和 `集群模式`。`无服务器模式` 下，模型文件被部署到无服务器云函数，即 [SCF](https://cloud.tencent.com/product/scf)，用户可以在其控制台上进一步操作。`集群模式` 下，模型文件被部署到 TI-A 的计算集群中。不同部署模式下的模型分开列出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.Orderby"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
//  UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) ListModelsWithContext(ctx context.Context, request *ListModelsRequest) (response *ListModelsResponse, err error) {
    if request == nil {
        request = NewListModelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListModelsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryLogsRequest() (request *QueryLogsRequest) {
    request = &QueryLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tia", APIVersion, "QueryLogs")
    
    
    return
}

func NewQueryLogsResponse() (response *QueryLogsResponse) {
    response = &QueryLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryLogs
// 查询 TI-A 训练任务的日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) QueryLogs(request *QueryLogsRequest) (response *QueryLogsResponse, err error) {
    return c.QueryLogsWithContext(context.Background(), request)
}

// QueryLogs
// 查询 TI-A 训练任务的日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.TimeOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDVERSION = "UnsupportedOperation.UnsupportedVersion"
func (c *Client) QueryLogsWithContext(ctx context.Context, request *QueryLogsRequest) (response *QueryLogsResponse, err error) {
    if request == nil {
        request = NewQueryLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryLogsResponse()
    err = c.Send(request, response)
    return
}
