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
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
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
