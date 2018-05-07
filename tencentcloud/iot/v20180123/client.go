// Copyright 1999-2018 Tencent Ltd.
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
package v20180123

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-01-23"

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


func NewActivateRuleRequest() (request *ActivateRuleRequest) {
    request = &ActivateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "ActivateRule")
    return
}

func NewActivateRuleResponse() (response *ActivateRuleResponse) {
    response = &ActivateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用规则
func (c *Client) ActivateRule(request *ActivateRuleRequest) (response *ActivateRuleResponse, err error) {
    if request == nil {
        request = NewActivateRuleRequest()
    }
    response = NewActivateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddDeviceRequest() (request *AddDeviceRequest) {
    request = &AddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "AddDevice")
    return
}

func NewAddDeviceResponse() (response *AddDeviceResponse) {
    response = &AddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增设备
func (c *Client) AddDevice(request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    if request == nil {
        request = NewAddDeviceRequest()
    }
    response = NewAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAddProductRequest() (request *AddProductRequest) {
    request = &AddProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "AddProduct")
    return
}

func NewAddProductResponse() (response *AddProductResponse) {
    response = &AddProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增产品
func (c *Client) AddProduct(request *AddProductRequest) (response *AddProductResponse, err error) {
    if request == nil {
        request = NewAddProductRequest()
    }
    response = NewAddProductResponse()
    err = c.Send(request, response)
    return
}

func NewAddRuleRequest() (request *AddRuleRequest) {
    request = &AddRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "AddRule")
    return
}

func NewAddRuleResponse() (response *AddRuleResponse) {
    response = &AddRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增规则
func (c *Client) AddRule(request *AddRuleRequest) (response *AddRuleResponse, err error) {
    if request == nil {
        request = NewAddRuleRequest()
    }
    response = NewAddRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddTopicRequest() (request *AddTopicRequest) {
    request = &AddTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "AddTopic")
    return
}

func NewAddTopicResponse() (response *AddTopicResponse) {
    response = &AddTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增Topic
func (c *Client) AddTopic(request *AddTopicRequest) (response *AddTopicResponse, err error) {
    if request == nil {
        request = NewAddTopicRequest()
    }
    response = NewAddTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeactivateRuleRequest() (request *DeactivateRuleRequest) {
    request = &DeactivateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "DeactivateRule")
    return
}

func NewDeactivateRuleResponse() (response *DeactivateRuleResponse) {
    response = &DeactivateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用规则
func (c *Client) DeactivateRule(request *DeactivateRuleRequest) (response *DeactivateRuleResponse, err error) {
    if request == nil {
        request = NewDeactivateRuleRequest()
    }
    response = NewDeactivateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "DeleteDevice")
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除设备
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "DeleteProduct")
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除产品
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除规则
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "DeleteTopic")
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除Topic
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataHistoryRequest() (request *GetDataHistoryRequest) {
    request = &GetDataHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDataHistory")
    return
}

func NewGetDataHistoryResponse() (response *GetDataHistoryResponse) {
    response = &GetDataHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取数据历史
func (c *Client) GetDataHistory(request *GetDataHistoryRequest) (response *GetDataHistoryResponse, err error) {
    if request == nil {
        request = NewGetDataHistoryRequest()
    }
    response = NewGetDataHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceRequest() (request *GetDeviceRequest) {
    request = &GetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDevice")
    return
}

func NewGetDeviceResponse() (response *GetDeviceResponse) {
    response = &GetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备信息
func (c *Client) GetDevice(request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    if request == nil {
        request = NewGetDeviceRequest()
    }
    response = NewGetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceDataRequest() (request *GetDeviceDataRequest) {
    request = &GetDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceData")
    return
}

func NewGetDeviceDataResponse() (response *GetDeviceDataResponse) {
    response = &GetDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备数据
func (c *Client) GetDeviceData(request *GetDeviceDataRequest) (response *GetDeviceDataResponse, err error) {
    if request == nil {
        request = NewGetDeviceDataRequest()
    }
    response = NewGetDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceLogRequest() (request *GetDeviceLogRequest) {
    request = &GetDeviceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceLog")
    return
}

func NewGetDeviceLogResponse() (response *GetDeviceLogResponse) {
    response = &GetDeviceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备日志
func (c *Client) GetDeviceLog(request *GetDeviceLogRequest) (response *GetDeviceLogResponse, err error) {
    if request == nil {
        request = NewGetDeviceLogRequest()
    }
    response = NewGetDeviceLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceStatusesRequest() (request *GetDeviceStatusesRequest) {
    request = &GetDeviceStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceStatuses")
    return
}

func NewGetDeviceStatusesResponse() (response *GetDeviceStatusesResponse) {
    response = &GetDeviceStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量获取设备状态
func (c *Client) GetDeviceStatuses(request *GetDeviceStatusesRequest) (response *GetDeviceStatusesResponse, err error) {
    if request == nil {
        request = NewGetDeviceStatusesRequest()
    }
    response = NewGetDeviceStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicesRequest() (request *GetDevicesRequest) {
    request = &GetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetDevices")
    return
}

func NewGetDevicesResponse() (response *GetDevicesResponse) {
    response = &GetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备列表
func (c *Client) GetDevices(request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    if request == nil {
        request = NewGetDevicesRequest()
    }
    response = NewGetDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetProductRequest() (request *GetProductRequest) {
    request = &GetProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetProduct")
    return
}

func NewGetProductResponse() (response *GetProductResponse) {
    response = &GetProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品信息
func (c *Client) GetProduct(request *GetProductRequest) (response *GetProductResponse, err error) {
    if request == nil {
        request = NewGetProductRequest()
    }
    response = NewGetProductResponse()
    err = c.Send(request, response)
    return
}

func NewGetProductsRequest() (request *GetProductsRequest) {
    request = &GetProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetProducts")
    return
}

func NewGetProductsResponse() (response *GetProductsResponse) {
    response = &GetProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户名下的产品列表
func (c *Client) GetProducts(request *GetProductsRequest) (response *GetProductsResponse, err error) {
    if request == nil {
        request = NewGetProductsRequest()
    }
    response = NewGetProductsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRuleRequest() (request *GetRuleRequest) {
    request = &GetRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetRule")
    return
}

func NewGetRuleResponse() (response *GetRuleResponse) {
    response = &GetRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取转发规则信息
func (c *Client) GetRule(request *GetRuleRequest) (response *GetRuleResponse, err error) {
    if request == nil {
        request = NewGetRuleRequest()
    }
    response = NewGetRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetRulesRequest() (request *GetRulesRequest) {
    request = &GetRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetRules")
    return
}

func NewGetRulesResponse() (response *GetRulesResponse) {
    response = &GetRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取转发规则列表
func (c *Client) GetRules(request *GetRulesRequest) (response *GetRulesResponse, err error) {
    if request == nil {
        request = NewGetRulesRequest()
    }
    response = NewGetRulesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicRequest() (request *GetTopicRequest) {
    request = &GetTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetTopic")
    return
}

func NewGetTopicResponse() (response *GetTopicResponse) {
    response = &GetTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Topic信息
func (c *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
    if request == nil {
        request = NewGetTopicRequest()
    }
    response = NewGetTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicsRequest() (request *GetTopicsRequest) {
    request = &GetTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetTopics")
    return
}

func NewGetTopicsResponse() (response *GetTopicsResponse) {
    response = &GetTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Topic列表
func (c *Client) GetTopics(request *GetTopicsRequest) (response *GetTopicsResponse, err error) {
    if request == nil {
        request = NewGetTopicsRequest()
    }
    response = NewGetTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "GetUser")
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取用户信息
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    if request == nil {
        request = NewGetUserRequest()
    }
    response = NewGetUserResponse()
    err = c.Send(request, response)
    return
}

func NewIssueDeviceControlRequest() (request *IssueDeviceControlRequest) {
    request = &IssueDeviceControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "IssueDeviceControl")
    return
}

func NewIssueDeviceControlResponse() (response *IssueDeviceControlResponse) {
    response = &IssueDeviceControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下发设备控制指令
func (c *Client) IssueDeviceControl(request *IssueDeviceControlRequest) (response *IssueDeviceControlResponse, err error) {
    if request == nil {
        request = NewIssueDeviceControlRequest()
    }
    response = NewIssueDeviceControlResponse()
    err = c.Send(request, response)
    return
}

func NewPublishMsgRequest() (request *PublishMsgRequest) {
    request = &PublishMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "PublishMsg")
    return
}

func NewPublishMsgResponse() (response *PublishMsgResponse) {
    response = &PublishMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 向Topic发布消息
func (c *Client) PublishMsg(request *PublishMsgRequest) (response *PublishMsgResponse, err error) {
    if request == nil {
        request = NewPublishMsgRequest()
    }
    response = NewPublishMsgResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceRequest() (request *ResetDeviceRequest) {
    request = &ResetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "ResetDevice")
    return
}

func NewResetDeviceResponse() (response *ResetDeviceResponse) {
    response = &ResetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重置设备
func (c *Client) ResetDevice(request *ResetDeviceRequest) (response *ResetDeviceResponse, err error) {
    if request == nil {
        request = NewResetDeviceRequest()
    }
    response = NewResetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductRequest() (request *UpdateProductRequest) {
    request = &UpdateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "UpdateProduct")
    return
}

func NewUpdateProductResponse() (response *UpdateProductResponse) {
    response = &UpdateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新产品信息
func (c *Client) UpdateProduct(request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
    if request == nil {
        request = NewUpdateProductRequest()
    }
    response = NewUpdateProductResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRuleRequest() (request *UpdateRuleRequest) {
    request = &UpdateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iot", APIVersion, "UpdateRule")
    return
}

func NewUpdateRuleResponse() (response *UpdateRuleResponse) {
    response = &UpdateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新规则
func (c *Client) UpdateRule(request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRuleRequest()
    }
    response = NewUpdateRuleResponse()
    err = c.Send(request, response)
    return
}
