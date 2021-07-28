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

package v20191010

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-10"

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


func NewFlowProductRemindRequest() (request *FlowProductRemindRequest) {
    request = &FlowProductRemindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("market", APIVersion, "FlowProductRemind")
    return
}

func NewFlowProductRemindResponse() (response *FlowProductRemindResponse) {
    response = &FlowProductRemindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FlowProductRemind
// 计量商品用量提醒，用于服务商调用云服务，云服务向客户发送提醒信息
func (c *Client) FlowProductRemind(request *FlowProductRemindRequest) (response *FlowProductRemindResponse, err error) {
    if request == nil {
        request = NewFlowProductRemindRequest()
    }
    response = NewFlowProductRemindResponse()
    err = c.Send(request, response)
    return
}

func NewGetUsagePlanUsageAmountRequest() (request *GetUsagePlanUsageAmountRequest) {
    request = &GetUsagePlanUsageAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("market", APIVersion, "GetUsagePlanUsageAmount")
    return
}

func NewGetUsagePlanUsageAmountResponse() (response *GetUsagePlanUsageAmountResponse) {
    response = &GetUsagePlanUsageAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUsagePlanUsageAmount
// 该接口可以根据InstanceId查询实例的api的使用情况。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetUsagePlanUsageAmount(request *GetUsagePlanUsageAmountRequest) (response *GetUsagePlanUsageAmountResponse, err error) {
    if request == nil {
        request = NewGetUsagePlanUsageAmountRequest()
    }
    response = NewGetUsagePlanUsageAmountResponse()
    err = c.Send(request, response)
    return
}

func NewSyncUserAndOrderInfoRequest() (request *SyncUserAndOrderInfoRequest) {
    request = &SyncUserAndOrderInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("market", APIVersion, "SyncUserAndOrderInfo")
    return
}

func NewSyncUserAndOrderInfoResponse() (response *SyncUserAndOrderInfoResponse) {
    response = &SyncUserAndOrderInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncUserAndOrderInfo
// 同步企微的用户信息和订单信息到云市场
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) SyncUserAndOrderInfo(request *SyncUserAndOrderInfoRequest) (response *SyncUserAndOrderInfoResponse, err error) {
    if request == nil {
        request = NewSyncUserAndOrderInfoRequest()
    }
    response = NewSyncUserAndOrderInfoResponse()
    err = c.Send(request, response)
    return
}
