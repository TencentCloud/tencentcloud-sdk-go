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

package v20200722

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-22"

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


func NewCreateFlowServiceRequest() (request *CreateFlowServiceRequest) {
    request = &CreateFlowServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "CreateFlowService")
    
    
    return
}

func NewCreateFlowServiceResponse() (response *CreateFlowServiceResponse) {
    response = &CreateFlowServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowService
// 该接口用于生成状态机服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateFlowService(request *CreateFlowServiceRequest) (response *CreateFlowServiceResponse, err error) {
    return c.CreateFlowServiceWithContext(context.Background(), request)
}

// CreateFlowService
// 该接口用于生成状态机服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateFlowServiceWithContext(ctx context.Context, request *CreateFlowServiceRequest) (response *CreateFlowServiceResponse, err error) {
    if request == nil {
        request = NewCreateFlowServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExecutionRequest() (request *DescribeExecutionRequest) {
    request = &DescribeExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "DescribeExecution")
    
    
    return
}

func NewDescribeExecutionResponse() (response *DescribeExecutionResponse) {
    response = &DescribeExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExecution
// 查询执行详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExecution(request *DescribeExecutionRequest) (response *DescribeExecutionResponse, err error) {
    return c.DescribeExecutionWithContext(context.Background(), request)
}

// DescribeExecution
// 查询执行详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExecutionWithContext(ctx context.Context, request *DescribeExecutionRequest) (response *DescribeExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExecutionHistoryRequest() (request *DescribeExecutionHistoryRequest) {
    request = &DescribeExecutionHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "DescribeExecutionHistory")
    
    
    return
}

func NewDescribeExecutionHistoryResponse() (response *DescribeExecutionHistoryResponse) {
    response = &DescribeExecutionHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExecutionHistory
// 一次执行会有很多步骤，经过很多节点，这个接口描述某一次执行的事件的历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeExecutionHistory(request *DescribeExecutionHistoryRequest) (response *DescribeExecutionHistoryResponse, err error) {
    return c.DescribeExecutionHistoryWithContext(context.Background(), request)
}

// DescribeExecutionHistory
// 一次执行会有很多步骤，经过很多节点，这个接口描述某一次执行的事件的历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeExecutionHistoryWithContext(ctx context.Context, request *DescribeExecutionHistoryRequest) (response *DescribeExecutionHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExecutionHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExecutionHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExecutionsRequest() (request *DescribeExecutionsRequest) {
    request = &DescribeExecutionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "DescribeExecutions")
    
    
    return
}

func NewDescribeExecutionsResponse() (response *DescribeExecutionsResponse) {
    response = &DescribeExecutionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExecutions
// 对状态机的执行历史进行描述.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeExecutions(request *DescribeExecutionsRequest) (response *DescribeExecutionsResponse, err error) {
    return c.DescribeExecutionsWithContext(context.Background(), request)
}

// DescribeExecutions
// 对状态机的执行历史进行描述.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeExecutionsWithContext(ctx context.Context, request *DescribeExecutionsRequest) (response *DescribeExecutionsResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExecutions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExecutionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowServiceDetailRequest() (request *DescribeFlowServiceDetailRequest) {
    request = &DescribeFlowServiceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "DescribeFlowServiceDetail")
    
    
    return
}

func NewDescribeFlowServiceDetailResponse() (response *DescribeFlowServiceDetailResponse) {
    response = &DescribeFlowServiceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowServiceDetail
// 查询该用户指定状态机下的详情数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowServiceDetail(request *DescribeFlowServiceDetailRequest) (response *DescribeFlowServiceDetailResponse, err error) {
    return c.DescribeFlowServiceDetailWithContext(context.Background(), request)
}

// DescribeFlowServiceDetail
// 查询该用户指定状态机下的详情数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowServiceDetailWithContext(ctx context.Context, request *DescribeFlowServiceDetailRequest) (response *DescribeFlowServiceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeFlowServiceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowServiceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowServiceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowServicesRequest() (request *DescribeFlowServicesRequest) {
    request = &DescribeFlowServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "DescribeFlowServices")
    
    
    return
}

func NewDescribeFlowServicesResponse() (response *DescribeFlowServicesResponse) {
    response = &DescribeFlowServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowServices
// 查询指定用户下所有状态机，以列表形式返回
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowServices(request *DescribeFlowServicesRequest) (response *DescribeFlowServicesResponse, err error) {
    return c.DescribeFlowServicesWithContext(context.Background(), request)
}

// DescribeFlowServices
// 查询指定用户下所有状态机，以列表形式返回
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowServicesWithContext(ctx context.Context, request *DescribeFlowServicesRequest) (response *DescribeFlowServicesResponse, err error) {
    if request == nil {
        request = NewDescribeFlowServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowServicesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlowServiceRequest() (request *ModifyFlowServiceRequest) {
    request = &ModifyFlowServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "ModifyFlowService")
    
    
    return
}

func NewModifyFlowServiceResponse() (response *ModifyFlowServiceResponse) {
    response = &ModifyFlowServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFlowService
// 该接口用于修改状态机
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFlowService(request *ModifyFlowServiceRequest) (response *ModifyFlowServiceResponse, err error) {
    return c.ModifyFlowServiceWithContext(context.Background(), request)
}

// ModifyFlowService
// 该接口用于修改状态机
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFlowServiceWithContext(ctx context.Context, request *ModifyFlowServiceRequest) (response *ModifyFlowServiceResponse, err error) {
    if request == nil {
        request = NewModifyFlowServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlowService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFlowServiceResponse()
    err = c.Send(request, response)
    return
}

func NewStartExecutionRequest() (request *StartExecutionRequest) {
    request = &StartExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "StartExecution")
    
    
    return
}

func NewStartExecutionResponse() (response *StartExecutionResponse) {
    response = &StartExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartExecution
// 为指定的状态机启动一次执行
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartExecution(request *StartExecutionRequest) (response *StartExecutionResponse, err error) {
    return c.StartExecutionWithContext(context.Background(), request)
}

// StartExecution
// 为指定的状态机启动一次执行
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartExecutionWithContext(ctx context.Context, request *StartExecutionRequest) (response *StartExecutionResponse, err error) {
    if request == nil {
        request = NewStartExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewStopExecutionRequest() (request *StopExecutionRequest) {
    request = &StopExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asw", APIVersion, "StopExecution")
    
    
    return
}

func NewStopExecutionResponse() (response *StopExecutionResponse) {
    response = &StopExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopExecution
// 终止某个状态机
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopExecution(request *StopExecutionRequest) (response *StopExecutionResponse, err error) {
    return c.StopExecutionWithContext(context.Background(), request)
}

// StopExecution
// 终止某个状态机
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopExecutionWithContext(ctx context.Context, request *StopExecutionRequest) (response *StopExecutionResponse, err error) {
    if request == nil {
        request = NewStopExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopExecution require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopExecutionResponse()
    err = c.Send(request, response)
    return
}
