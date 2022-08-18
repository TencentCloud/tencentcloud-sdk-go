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

package v20210408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-08"

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


func NewAddSubStructureTasksRequest() (request *AddSubStructureTasksRequest) {
    request = &AddSubStructureTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "AddSubStructureTasks")
    
    
    return
}

func NewAddSubStructureTasksResponse() (response *AddSubStructureTasksResponse) {
    response = &AddSubStructureTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSubStructureTasks
// 如果主任务下的报告不满足需求，可以基于主任务批量添加子任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddSubStructureTasks(request *AddSubStructureTasksRequest) (response *AddSubStructureTasksResponse, err error) {
    return c.AddSubStructureTasksWithContext(context.Background(), request)
}

// AddSubStructureTasks
// 如果主任务下的报告不满足需求，可以基于主任务批量添加子任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddSubStructureTasksWithContext(ctx context.Context, request *AddSubStructureTasksRequest) (response *AddSubStructureTasksResponse, err error) {
    if request == nil {
        request = NewAddSubStructureTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSubStructureTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSubStructureTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoClassifyStructureTaskRequest() (request *CreateAutoClassifyStructureTaskRequest) {
    request = &CreateAutoClassifyStructureTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "CreateAutoClassifyStructureTask")
    
    
    return
}

func NewCreateAutoClassifyStructureTaskResponse() (response *CreateAutoClassifyStructureTaskResponse) {
    response = &CreateAutoClassifyStructureTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoClassifyStructureTask
// 本接口(CreateAutoClassifyStructureTask)基于提供的客户及保单信息，创建并启动结构化识别任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAutoClassifyStructureTask(request *CreateAutoClassifyStructureTaskRequest) (response *CreateAutoClassifyStructureTaskResponse, err error) {
    return c.CreateAutoClassifyStructureTaskWithContext(context.Background(), request)
}

// CreateAutoClassifyStructureTask
// 本接口(CreateAutoClassifyStructureTask)基于提供的客户及保单信息，创建并启动结构化识别任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAutoClassifyStructureTaskWithContext(ctx context.Context, request *CreateAutoClassifyStructureTaskRequest) (response *CreateAutoClassifyStructureTaskResponse, err error) {
    if request == nil {
        request = NewCreateAutoClassifyStructureTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoClassifyStructureTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoClassifyStructureTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStructureTaskRequest() (request *CreateStructureTaskRequest) {
    request = &CreateStructureTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "CreateStructureTask")
    
    
    return
}

func NewCreateStructureTaskResponse() (response *CreateStructureTaskResponse) {
    response = &CreateStructureTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStructureTask
// 本接口(CreateStructureTask)基于提供的客户及保单信息，创建并启动结构化识别任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateStructureTask(request *CreateStructureTaskRequest) (response *CreateStructureTaskResponse, err error) {
    return c.CreateStructureTaskWithContext(context.Background(), request)
}

// CreateStructureTask
// 本接口(CreateStructureTask)基于提供的客户及保单信息，创建并启动结构化识别任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateStructureTaskWithContext(ctx context.Context, request *CreateStructureTaskRequest) (response *CreateStructureTaskResponse, err error) {
    if request == nil {
        request = NewCreateStructureTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStructureTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStructureTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnderwriteTaskByIdRequest() (request *CreateUnderwriteTaskByIdRequest) {
    request = &CreateUnderwriteTaskByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "CreateUnderwriteTaskById")
    
    
    return
}

func NewCreateUnderwriteTaskByIdResponse() (response *CreateUnderwriteTaskByIdResponse) {
    response = &CreateUnderwriteTaskByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUnderwriteTaskById
// 本接口(CreateUnderwriteTaskById)用于根据结构化任务ID创建核保任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUnderwriteTaskById(request *CreateUnderwriteTaskByIdRequest) (response *CreateUnderwriteTaskByIdResponse, err error) {
    return c.CreateUnderwriteTaskByIdWithContext(context.Background(), request)
}

// CreateUnderwriteTaskById
// 本接口(CreateUnderwriteTaskById)用于根据结构化任务ID创建核保任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateUnderwriteTaskByIdWithContext(ctx context.Context, request *CreateUnderwriteTaskByIdRequest) (response *CreateUnderwriteTaskByIdResponse, err error) {
    if request == nil {
        request = NewCreateUnderwriteTaskByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUnderwriteTaskById require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUnderwriteTaskByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineUnderwriteRequest() (request *DescribeMachineUnderwriteRequest) {
    request = &DescribeMachineUnderwriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeMachineUnderwrite")
    
    
    return
}

func NewDescribeMachineUnderwriteResponse() (response *DescribeMachineUnderwriteResponse) {
    response = &DescribeMachineUnderwriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMachineUnderwrite
// 本接口(DescribeMachineUnderwrite)用于查询机器核保任务数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeMachineUnderwrite(request *DescribeMachineUnderwriteRequest) (response *DescribeMachineUnderwriteResponse, err error) {
    return c.DescribeMachineUnderwriteWithContext(context.Background(), request)
}

// DescribeMachineUnderwrite
// 本接口(DescribeMachineUnderwrite)用于查询机器核保任务数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeMachineUnderwriteWithContext(ctx context.Context, request *DescribeMachineUnderwriteRequest) (response *DescribeMachineUnderwriteResponse, err error) {
    if request == nil {
        request = NewDescribeMachineUnderwriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineUnderwrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineUnderwriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityScoreRequest() (request *DescribeQualityScoreRequest) {
    request = &DescribeQualityScoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeQualityScore")
    
    
    return
}

func NewDescribeQualityScoreResponse() (response *DescribeQualityScoreResponse) {
    response = &DescribeQualityScoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQualityScore
// 获取图片质量分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeQualityScore(request *DescribeQualityScoreRequest) (response *DescribeQualityScoreResponse, err error) {
    return c.DescribeQualityScoreWithContext(context.Background(), request)
}

// DescribeQualityScore
// 获取图片质量分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeQualityScoreWithContext(ctx context.Context, request *DescribeQualityScoreRequest) (response *DescribeQualityScoreResponse, err error) {
    if request == nil {
        request = NewDescribeQualityScoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityScoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportClassifyRequest() (request *DescribeReportClassifyRequest) {
    request = &DescribeReportClassifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeReportClassify")
    
    
    return
}

func NewDescribeReportClassifyResponse() (response *DescribeReportClassifyResponse) {
    response = &DescribeReportClassifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReportClassify
// 辅助用户对批量报告自动分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeReportClassify(request *DescribeReportClassifyRequest) (response *DescribeReportClassifyResponse, err error) {
    return c.DescribeReportClassifyWithContext(context.Background(), request)
}

// DescribeReportClassify
// 辅助用户对批量报告自动分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeReportClassifyWithContext(ctx context.Context, request *DescribeReportClassifyRequest) (response *DescribeReportClassifyResponse, err error) {
    if request == nil {
        request = NewDescribeReportClassifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportClassify require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportClassifyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructCompareDataRequest() (request *DescribeStructCompareDataRequest) {
    request = &DescribeStructCompareDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructCompareData")
    
    
    return
}

func NewDescribeStructCompareDataResponse() (response *DescribeStructCompareDataResponse) {
    response = &DescribeStructCompareDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructCompareData
// 结构化对比查询接口，对比结构化复核前后数据差异，查询识别正确率，召回率。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructCompareData(request *DescribeStructCompareDataRequest) (response *DescribeStructCompareDataResponse, err error) {
    return c.DescribeStructCompareDataWithContext(context.Background(), request)
}

// DescribeStructCompareData
// 结构化对比查询接口，对比结构化复核前后数据差异，查询识别正确率，召回率。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructCompareDataWithContext(ctx context.Context, request *DescribeStructCompareDataRequest) (response *DescribeStructCompareDataResponse, err error) {
    if request == nil {
        request = NewDescribeStructCompareDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStructCompareData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStructCompareDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureDifferenceRequest() (request *DescribeStructureDifferenceRequest) {
    request = &DescribeStructureDifferenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureDifference")
    
    
    return
}

func NewDescribeStructureDifferenceResponse() (response *DescribeStructureDifferenceResponse) {
    response = &DescribeStructureDifferenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureDifference
// 结构化复核差异查询接口，对比结构化复核前后数据差异，返回差异的部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructureDifference(request *DescribeStructureDifferenceRequest) (response *DescribeStructureDifferenceResponse, err error) {
    return c.DescribeStructureDifferenceWithContext(context.Background(), request)
}

// DescribeStructureDifference
// 结构化复核差异查询接口，对比结构化复核前后数据差异，返回差异的部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStructureDifferenceWithContext(ctx context.Context, request *DescribeStructureDifferenceRequest) (response *DescribeStructureDifferenceResponse, err error) {
    if request == nil {
        request = NewDescribeStructureDifferenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStructureDifference require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStructureDifferenceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureResultRequest() (request *DescribeStructureResultRequest) {
    request = &DescribeStructureResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureResult")
    
    
    return
}

func NewDescribeStructureResultResponse() (response *DescribeStructureResultResponse) {
    response = &DescribeStructureResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureResult
// 本接口(DescribeStructureResult)用于查询结构化结果接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeStructureResult(request *DescribeStructureResultRequest) (response *DescribeStructureResultResponse, err error) {
    return c.DescribeStructureResultWithContext(context.Background(), request)
}

// DescribeStructureResult
// 本接口(DescribeStructureResult)用于查询结构化结果接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) DescribeStructureResultWithContext(ctx context.Context, request *DescribeStructureResultRequest) (response *DescribeStructureResultResponse, err error) {
    if request == nil {
        request = NewDescribeStructureResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStructureResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStructureResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStructureTaskResultRequest() (request *DescribeStructureTaskResultRequest) {
    request = &DescribeStructureTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureTaskResult")
    
    
    return
}

func NewDescribeStructureTaskResultResponse() (response *DescribeStructureTaskResultResponse) {
    response = &DescribeStructureTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStructureTaskResult
// 依据任务ID获取结构化结果接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeStructureTaskResult(request *DescribeStructureTaskResultRequest) (response *DescribeStructureTaskResultResponse, err error) {
    return c.DescribeStructureTaskResultWithContext(context.Background(), request)
}

// DescribeStructureTaskResult
// 依据任务ID获取结构化结果接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeStructureTaskResultWithContext(ctx context.Context, request *DescribeStructureTaskResultRequest) (response *DescribeStructureTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeStructureTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStructureTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStructureTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnderwriteTaskRequest() (request *DescribeUnderwriteTaskRequest) {
    request = &DescribeUnderwriteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "DescribeUnderwriteTask")
    
    
    return
}

func NewDescribeUnderwriteTaskResponse() (response *DescribeUnderwriteTaskResponse) {
    response = &DescribeUnderwriteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnderwriteTask
// 本接口(DescribeUnderwriteTask)用于查询核保任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUnderwriteTask(request *DescribeUnderwriteTaskRequest) (response *DescribeUnderwriteTaskResponse, err error) {
    return c.DescribeUnderwriteTaskWithContext(context.Background(), request)
}

// DescribeUnderwriteTask
// 本接口(DescribeUnderwriteTask)用于查询核保任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUnderwriteTaskWithContext(ctx context.Context, request *DescribeUnderwriteTaskRequest) (response *DescribeUnderwriteTaskResponse, err error) {
    if request == nil {
        request = NewDescribeUnderwriteTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnderwriteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnderwriteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUploadMedicalFileRequest() (request *UploadMedicalFileRequest) {
    request = &UploadMedicalFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cii", APIVersion, "UploadMedicalFile")
    
    
    return
}

func NewUploadMedicalFileResponse() (response *UploadMedicalFileResponse) {
    response = &UploadMedicalFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadMedicalFile
// 上传医疗影像文件，可以用来做结构化。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ENTERPRISE = "AuthFailure.Enterprise"
//  AUTHFAILURE_PERSONAL = "AuthFailure.Personal"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UploadMedicalFile(request *UploadMedicalFileRequest) (response *UploadMedicalFileResponse, err error) {
    return c.UploadMedicalFileWithContext(context.Background(), request)
}

// UploadMedicalFile
// 上传医疗影像文件，可以用来做结构化。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ENTERPRISE = "AuthFailure.Enterprise"
//  AUTHFAILURE_PERSONAL = "AuthFailure.Personal"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UploadMedicalFileWithContext(ctx context.Context, request *UploadMedicalFileRequest) (response *UploadMedicalFileResponse, err error) {
    if request == nil {
        request = NewUploadMedicalFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadMedicalFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadMedicalFileResponse()
    err = c.Send(request, response)
    return
}
