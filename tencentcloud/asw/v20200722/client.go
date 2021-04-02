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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 该接口用于生成状态机服务
func (c *Client) CreateFlowService(request *CreateFlowServiceRequest) (response *CreateFlowServiceResponse, err error) {
    if request == nil {
        request = NewCreateFlowServiceRequest()
    }
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

// 查询执行详细信息
func (c *Client) DescribeExecution(request *DescribeExecutionRequest) (response *DescribeExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionRequest()
    }
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

// 一次执行会有很多步骤，经过很多节点，这个接口描述某一次执行的事件的历史
func (c *Client) DescribeExecutionHistory(request *DescribeExecutionHistoryRequest) (response *DescribeExecutionHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionHistoryRequest()
    }
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

// 对状态机的执行历史进行描述.
func (c *Client) DescribeExecutions(request *DescribeExecutionsRequest) (response *DescribeExecutionsResponse, err error) {
    if request == nil {
        request = NewDescribeExecutionsRequest()
    }
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

// 查询该用户指定状态机下的详情数据。
func (c *Client) DescribeFlowServiceDetail(request *DescribeFlowServiceDetailRequest) (response *DescribeFlowServiceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeFlowServiceDetailRequest()
    }
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

// 查询指定用户下所有状态机，以列表形式返回
func (c *Client) DescribeFlowServices(request *DescribeFlowServicesRequest) (response *DescribeFlowServicesResponse, err error) {
    if request == nil {
        request = NewDescribeFlowServicesRequest()
    }
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

// 该接口用于修改状态机
func (c *Client) ModifyFlowService(request *ModifyFlowServiceRequest) (response *ModifyFlowServiceResponse, err error) {
    if request == nil {
        request = NewModifyFlowServiceRequest()
    }
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

// 为指定的状态机启动一次执行
func (c *Client) StartExecution(request *StartExecutionRequest) (response *StartExecutionResponse, err error) {
    if request == nil {
        request = NewStartExecutionRequest()
    }
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

// 终止某个状态机
func (c *Client) StopExecution(request *StopExecutionRequest) (response *StopExecutionResponse, err error) {
    if request == nil {
        request = NewStopExecutionRequest()
    }
    response = NewStopExecutionResponse()
    err = c.Send(request, response)
    return
}
