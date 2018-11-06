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

package v20180423

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-23"

type Client struct {
    common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
    return
}


func NewBindPsaTagRequest() (request *BindPsaTagRequest) {
    request = &BindPsaTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "BindPsaTag")
    return
}

func NewBindPsaTagResponse() (response *BindPsaTagResponse) {
    response = &BindPsaTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为预授权规则绑定标签
func (c *Client) BindPsaTag(request *BindPsaTagRequest) (response *BindPsaTagResponse, err error) {
    if request == nil {
        request = NewBindPsaTagRequest()
    }
    response = NewBindPsaTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePsaRegulationRequest() (request *CreatePsaRegulationRequest) {
    request = &CreatePsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "CreatePsaRegulation")
    return
}

func NewCreatePsaRegulationResponse() (response *CreatePsaRegulationResponse) {
    response = &CreatePsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建预授权规则
func (c *Client) CreatePsaRegulation(request *CreatePsaRegulationRequest) (response *CreatePsaRegulationResponse, err error) {
    if request == nil {
        request = NewCreatePsaRegulationRequest()
    }
    response = NewCreatePsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSpotDeviceRequest() (request *CreateSpotDeviceRequest) {
    request = &CreateSpotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "CreateSpotDevice")
    return
}

func NewCreateSpotDeviceResponse() (response *CreateSpotDeviceResponse) {
    response = &CreateSpotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建黑石竞价实例
func (c *Client) CreateSpotDevice(request *CreateSpotDeviceRequest) (response *CreateSpotDeviceResponse, err error) {
    if request == nil {
        request = NewCreateSpotDeviceRequest()
    }
    response = NewCreateSpotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserCmdRequest() (request *CreateUserCmdRequest) {
    request = &CreateUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "CreateUserCmd")
    return
}

func NewCreateUserCmdResponse() (response *CreateUserCmdResponse) {
    response = &CreateUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建自定义脚本
func (c *Client) CreateUserCmd(request *CreateUserCmdRequest) (response *CreateUserCmdResponse, err error) {
    if request == nil {
        request = NewCreateUserCmdRequest()
    }
    response = NewCreateUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePsaRegulationRequest() (request *DeletePsaRegulationRequest) {
    request = &DeletePsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DeletePsaRegulation")
    return
}

func NewDeletePsaRegulationResponse() (response *DeletePsaRegulationResponse) {
    response = &DeletePsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除预授权规则
func (c *Client) DeletePsaRegulation(request *DeletePsaRegulationRequest) (response *DeletePsaRegulationResponse, err error) {
    if request == nil {
        request = NewDeletePsaRegulationRequest()
    }
    response = NewDeletePsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevices")
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询物理服务器，可以按照实例，业务IP等过滤
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePsaRegulationsRequest() (request *DescribePsaRegulationsRequest) {
    request = &DescribePsaRegulationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribePsaRegulations")
    return
}

func NewDescribePsaRegulationsResponse() (response *DescribePsaRegulationsResponse) {
    response = &DescribePsaRegulationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取预授权规则列表
func (c *Client) DescribePsaRegulations(request *DescribePsaRegulationsRequest) (response *DescribePsaRegulationsResponse, err error) {
    if request == nil {
        request = NewDescribePsaRegulationsRequest()
    }
    response = NewDescribePsaRegulationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepairTaskConstantRequest() (request *DescribeRepairTaskConstantRequest) {
    request = &DescribeRepairTaskConstantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeRepairTaskConstant")
    return
}

func NewDescribeRepairTaskConstantResponse() (response *DescribeRepairTaskConstantResponse) {
    response = &DescribeRepairTaskConstantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 维修任务配置获取
func (c *Client) DescribeRepairTaskConstant(request *DescribeRepairTaskConstantRequest) (response *DescribeRepairTaskConstantResponse, err error) {
    if request == nil {
        request = NewDescribeRepairTaskConstantRequest()
    }
    response = NewDescribeRepairTaskConstantResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeTaskInfo")
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户维修任务列表及详细信息<br>
// <br>
// TaskStatus（任务状态ID）与状态中文名的对应关系如下：<br>
// 1：未授权<br>
// 2：处理中<br>
// 3：待确认<br>
// 4：未授权-暂不处理<br>
// 5：已恢复<br>
// 6：待确认-未恢复<br>
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskOperationLogRequest() (request *DescribeTaskOperationLogRequest) {
    request = &DescribeTaskOperationLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeTaskOperationLog")
    return
}

func NewDescribeTaskOperationLogResponse() (response *DescribeTaskOperationLogResponse) {
    response = &DescribeTaskOperationLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取维修任务操作日志
func (c *Client) DescribeTaskOperationLog(request *DescribeTaskOperationLogRequest) (response *DescribeTaskOperationLogResponse, err error) {
    if request == nil {
        request = NewDescribeTaskOperationLogRequest()
    }
    response = NewDescribeTaskOperationLogResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPsaRegulationRequest() (request *ModifyPsaRegulationRequest) {
    request = &ModifyPsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyPsaRegulation")
    return
}

func NewModifyPsaRegulationResponse() (response *ModifyPsaRegulationResponse) {
    response = &ModifyPsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 允许修改规则信息及关联故障类型
func (c *Client) ModifyPsaRegulation(request *ModifyPsaRegulationRequest) (response *ModifyPsaRegulationResponse, err error) {
    if request == nil {
        request = NewModifyPsaRegulationRequest()
    }
    response = NewModifyPsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewRepairTaskControlRequest() (request *RepairTaskControlRequest) {
    request = &RepairTaskControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "RepairTaskControl")
    return
}

func NewRepairTaskControlResponse() (response *RepairTaskControlResponse) {
    response = &RepairTaskControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口用于操作维修任务<br>
// 入参TaskId为维修任务ID<br>
// 入参Operate表示对维修任务的操作，支持如下取值：<br>
// AuthorizeRepair（授权维修）<br>
// Ignore（暂不提醒）<br>
// ConfirmRecovered（维修完成后，确认故障恢复）<br>
// ConfirmUnRecovered（维修完成后，确认故障未恢复）<br>
// <br>
// 操作约束（当前任务状态(TaskStatus)->对应可执行的操作）：<br>
// 未授权(1)->授权维修；暂不处理<br>
// 暂不处理(4)->授权维修<br>
// 待确认(3)->确认故障恢复；确认故障未恢复<br>
// 未恢复(6)->确认故障恢复<br>
// <br>
// 对于Ping不可达故障的任务，还允许：<br>
// 未授权->确认故障恢复<br>
// 暂不处理->确认故障恢复<br>
// <br>
// 处理中与已恢复状态的任务不允许进行操作。<br>
// <br>
// 详细信息请访问：https://cloud.tencent.com/document/product/386/18190
func (c *Client) RepairTaskControl(request *RepairTaskControlRequest) (response *RepairTaskControlResponse, err error) {
    if request == nil {
        request = NewRepairTaskControlRequest()
    }
    response = NewRepairTaskControlResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPsaTagRequest() (request *UnbindPsaTagRequest) {
    request = &UnbindPsaTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "UnbindPsaTag")
    return
}

func NewUnbindPsaTagResponse() (response *UnbindPsaTagResponse) {
    response = &UnbindPsaTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解除标签与预授权规则的绑定
func (c *Client) UnbindPsaTag(request *UnbindPsaTagRequest) (response *UnbindPsaTagResponse, err error) {
    if request == nil {
        request = NewUnbindPsaTagRequest()
    }
    response = NewUnbindPsaTagResponse()
    err = c.Send(request, response)
    return
}
