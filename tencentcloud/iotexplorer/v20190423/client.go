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

package v20190423

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-23"

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


func NewCallDeviceActionAsyncRequest() (request *CallDeviceActionAsyncRequest) {
    request = &CallDeviceActionAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CallDeviceActionAsync")
    return
}

func NewCallDeviceActionAsyncResponse() (response *CallDeviceActionAsyncResponse) {
    response = &CallDeviceActionAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供给用户异步调用设备行为的能力
func (c *Client) CallDeviceActionAsync(request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionAsyncRequest()
    }
    response = NewCallDeviceActionAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCallDeviceActionSyncRequest() (request *CallDeviceActionSyncRequest) {
    request = &CallDeviceActionSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CallDeviceActionSync")
    return
}

func NewCallDeviceActionSyncResponse() (response *CallDeviceActionSyncResponse) {
    response = &CallDeviceActionSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为用户提供同步调用设备行为的能力。
func (c *Client) CallDeviceActionSync(request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionSyncRequest()
    }
    response = NewCallDeviceActionSyncResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceDataRequest() (request *ControlDeviceDataRequest) {
    request = &ControlDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ControlDeviceData")
    return
}

func NewControlDeviceDataResponse() (response *ControlDeviceDataResponse) {
    response = &ControlDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据设备产品ID、设备名称，设置控制设备的属性数据。
func (c *Client) ControlDeviceData(request *ControlDeviceDataRequest) (response *ControlDeviceDataResponse, err error) {
    if request == nil {
        request = NewControlDeviceDataRequest()
    }
    response = NewControlDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateDevice")
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建设备
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoRaFrequencyRequest() (request *CreateLoRaFrequencyRequest) {
    request = &CreateLoRaFrequencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateLoRaFrequency")
    return
}

func NewCreateLoRaFrequencyResponse() (response *CreateLoRaFrequencyResponse) {
    response = &CreateLoRaFrequencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建 LoRa 自定义频点
func (c *Client) CreateLoRaFrequency(request *CreateLoRaFrequencyRequest) (response *CreateLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewCreateLoRaFrequencyRequest()
    }
    response = NewCreateLoRaFrequencyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoRaGatewayRequest() (request *CreateLoRaGatewayRequest) {
    request = &CreateLoRaGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateLoRaGateway")
    return
}

func NewCreateLoRaGatewayResponse() (response *CreateLoRaGatewayResponse) {
    response = &CreateLoRaGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建新 LoRa 网关设备接口
func (c *Client) CreateLoRaGateway(request *CreateLoRaGatewayRequest) (response *CreateLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewCreateLoRaGatewayRequest()
    }
    response = NewCreateLoRaGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateProject")
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为用户提供新建项目的能力，用于集中管理产品和应用。
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStudioProductRequest() (request *CreateStudioProductRequest) {
    request = &CreateStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateStudioProduct")
    return
}

func NewCreateStudioProductResponse() (response *CreateStudioProductResponse) {
    response = &CreateStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为用户提供新建产品的能力，用于管理用户的设备
func (c *Client) CreateStudioProduct(request *CreateStudioProductRequest) (response *CreateStudioProductResponse, err error) {
    if request == nil {
        request = NewCreateStudioProductRequest()
    }
    response = NewCreateStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRuleRequest() (request *CreateTopicRuleRequest) {
    request = &CreateTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateTopicRule")
    return
}

func NewCreateTopicRuleResponse() (response *CreateTopicRuleResponse) {
    response = &CreateTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建规则
func (c *Client) CreateTopicRule(request *CreateTopicRuleRequest) (response *CreateTopicRuleResponse, err error) {
    if request == nil {
        request = NewCreateTopicRuleRequest()
    }
    response = NewCreateTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteDevice")
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

func NewDeleteLoRaFrequencyRequest() (request *DeleteLoRaFrequencyRequest) {
    request = &DeleteLoRaFrequencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteLoRaFrequency")
    return
}

func NewDeleteLoRaFrequencyResponse() (response *DeleteLoRaFrequencyResponse) {
    response = &DeleteLoRaFrequencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供删除LoRa自定义频点的能力
func (c *Client) DeleteLoRaFrequency(request *DeleteLoRaFrequencyRequest) (response *DeleteLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewDeleteLoRaFrequencyRequest()
    }
    response = NewDeleteLoRaFrequencyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoRaGatewayRequest() (request *DeleteLoRaGatewayRequest) {
    request = &DeleteLoRaGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteLoRaGateway")
    return
}

func NewDeleteLoRaGatewayResponse() (response *DeleteLoRaGatewayResponse) {
    response = &DeleteLoRaGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除  LoRa 网关的接口
func (c *Client) DeleteLoRaGateway(request *DeleteLoRaGatewayRequest) (response *DeleteLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteLoRaGatewayRequest()
    }
    response = NewDeleteLoRaGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteProject")
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供删除某个项目的能力
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStudioProductRequest() (request *DeleteStudioProductRequest) {
    request = &DeleteStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteStudioProduct")
    return
}

func NewDeleteStudioProductResponse() (response *DeleteStudioProductResponse) {
    response = &DeleteStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供删除某个项目下产品的能力
func (c *Client) DeleteStudioProduct(request *DeleteStudioProductRequest) (response *DeleteStudioProductResponse, err error) {
    if request == nil {
        request = NewDeleteStudioProductRequest()
    }
    response = NewDeleteStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRuleRequest() (request *DeleteTopicRuleRequest) {
    request = &DeleteTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteTopicRule")
    return
}

func NewDeleteTopicRuleResponse() (response *DeleteTopicRuleResponse) {
    response = &DeleteTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除规则
func (c *Client) DeleteTopicRule(request *DeleteTopicRuleRequest) (response *DeleteTopicRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRuleRequest()
    }
    response = NewDeleteTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDevice")
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查看某个设备的详细信息
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceDataRequest() (request *DescribeDeviceDataRequest) {
    request = &DescribeDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceData")
    return
}

func NewDescribeDeviceDataResponse() (response *DescribeDeviceDataResponse) {
    response = &DescribeDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据设备产品ID、设备名称，获取设备上报的属性数据。
func (c *Client) DescribeDeviceData(request *DescribeDeviceDataRequest) (response *DescribeDeviceDataResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataRequest()
    }
    response = NewDescribeDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceDataHistoryRequest() (request *DescribeDeviceDataHistoryRequest) {
    request = &DescribeDeviceDataHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceDataHistory")
    return
}

func NewDescribeDeviceDataHistoryResponse() (response *DescribeDeviceDataHistoryResponse) {
    response = &DescribeDeviceDataHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备在指定时间范围内上报的历史数据。
func (c *Client) DescribeDeviceDataHistory(request *DescribeDeviceDataHistoryRequest) (response *DescribeDeviceDataHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataHistoryRequest()
    }
    response = NewDescribeDeviceDataHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskRequest() (request *DescribeFirmwareTaskRequest) {
    request = &DescribeFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFirmwareTask")
    return
}

func NewDescribeFirmwareTaskResponse() (response *DescribeFirmwareTaskResponse) {
    response = &DescribeFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询固件升级任务列表
func (c *Client) DescribeFirmwareTask(request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    response = NewDescribeFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoRaFrequencyRequest() (request *DescribeLoRaFrequencyRequest) {
    request = &DescribeLoRaFrequencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeLoRaFrequency")
    return
}

func NewDescribeLoRaFrequencyResponse() (response *DescribeLoRaFrequencyResponse) {
    response = &DescribeLoRaFrequencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供查询LoRa自定义频点详情的能力
func (c *Client) DescribeLoRaFrequency(request *DescribeLoRaFrequencyRequest) (response *DescribeLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewDescribeLoRaFrequencyRequest()
    }
    response = NewDescribeLoRaFrequencyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelDefinitionRequest() (request *DescribeModelDefinitionRequest) {
    request = &DescribeModelDefinitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeModelDefinition")
    return
}

func NewDescribeModelDefinitionResponse() (response *DescribeModelDefinitionResponse) {
    response = &DescribeModelDefinitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询产品配置的数据模板信息
func (c *Client) DescribeModelDefinition(request *DescribeModelDefinitionRequest) (response *DescribeModelDefinitionResponse, err error) {
    if request == nil {
        request = NewDescribeModelDefinitionRequest()
    }
    response = NewDescribeModelDefinitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeProject")
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询项目详情
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStudioProductRequest() (request *DescribeStudioProductRequest) {
    request = &DescribeStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeStudioProduct")
    return
}

func NewDescribeStudioProductResponse() (response *DescribeStudioProductResponse) {
    response = &DescribeStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供查看产品详细信息的能力，包括产品的ID、数据协议、认证类型等重要参数
func (c *Client) DescribeStudioProduct(request *DescribeStudioProductRequest) (response *DescribeStudioProductResponse, err error) {
    if request == nil {
        request = NewDescribeStudioProductRequest()
    }
    response = NewDescribeStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRuleRequest() (request *DescribeTopicRuleRequest) {
    request = &DescribeTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeTopicRule")
    return
}

func NewDescribeTopicRuleResponse() (response *DescribeTopicRuleResponse) {
    response = &DescribeTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取规则信息
func (c *Client) DescribeTopicRule(request *DescribeTopicRuleRequest) (response *DescribeTopicRuleResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRuleRequest()
    }
    response = NewDescribeTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTopicRuleRequest() (request *DisableTopicRuleRequest) {
    request = &DisableTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DisableTopicRule")
    return
}

func NewDisableTopicRuleResponse() (response *DisableTopicRuleResponse) {
    response = &DisableTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁用规则
func (c *Client) DisableTopicRule(request *DisableTopicRuleRequest) (response *DisableTopicRuleResponse, err error) {
    if request == nil {
        request = NewDisableTopicRuleRequest()
    }
    response = NewDisableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTopicRuleRequest() (request *EnableTopicRuleRequest) {
    request = &EnableTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "EnableTopicRule")
    return
}

func NewEnableTopicRuleResponse() (response *EnableTopicRuleResponse) {
    response = &EnableTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用规则
func (c *Client) EnableTopicRule(request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    if request == nil {
        request = NewEnableTopicRuleRequest()
    }
    response = NewEnableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetCOSURLRequest() (request *GetCOSURLRequest) {
    request = &GetCOSURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetCOSURL")
    return
}

func NewGetCOSURLResponse() (response *GetCOSURLResponse) {
    response = &GetCOSURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetCOSURL）用于获取固件存储在COS的URL 
func (c *Client) GetCOSURL(request *GetCOSURLRequest) (response *GetCOSURLResponse, err error) {
    if request == nil {
        request = NewGetCOSURLRequest()
    }
    response = NewGetCOSURLResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceListRequest() (request *GetDeviceListRequest) {
    request = &GetDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetDeviceList")
    return
}

func NewGetDeviceListResponse() (response *GetDeviceListResponse) {
    response = &GetDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询某个产品下的设备列表
func (c *Client) GetDeviceList(request *GetDeviceListRequest) (response *GetDeviceListResponse, err error) {
    if request == nil {
        request = NewGetDeviceListRequest()
    }
    response = NewGetDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewGetLoRaGatewayListRequest() (request *GetLoRaGatewayListRequest) {
    request = &GetLoRaGatewayListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetLoRaGatewayList")
    return
}

func NewGetLoRaGatewayListResponse() (response *GetLoRaGatewayListResponse) {
    response = &GetLoRaGatewayListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取 LoRa 网关列表接口
func (c *Client) GetLoRaGatewayList(request *GetLoRaGatewayListRequest) (response *GetLoRaGatewayListResponse, err error) {
    if request == nil {
        request = NewGetLoRaGatewayListRequest()
    }
    response = NewGetLoRaGatewayListResponse()
    err = c.Send(request, response)
    return
}

func NewGetProjectListRequest() (request *GetProjectListRequest) {
    request = &GetProjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetProjectList")
    return
}

func NewGetProjectListResponse() (response *GetProjectListResponse) {
    response = &GetProjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供查询用户所创建的项目列表查询功能。
func (c *Client) GetProjectList(request *GetProjectListRequest) (response *GetProjectListResponse, err error) {
    if request == nil {
        request = NewGetProjectListRequest()
    }
    response = NewGetProjectListResponse()
    err = c.Send(request, response)
    return
}

func NewGetStudioProductListRequest() (request *GetStudioProductListRequest) {
    request = &GetStudioProductListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetStudioProductList")
    return
}

func NewGetStudioProductListResponse() (response *GetStudioProductListResponse) {
    response = &GetStudioProductListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供查询某个项目下所有产品信息的能力。
func (c *Client) GetStudioProductList(request *GetStudioProductListRequest) (response *GetStudioProductListResponse, err error) {
    if request == nil {
        request = NewGetStudioProductListRequest()
    }
    response = NewGetStudioProductListResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicRuleListRequest() (request *GetTopicRuleListRequest) {
    request = &GetTopicRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetTopicRuleList")
    return
}

func NewGetTopicRuleListResponse() (response *GetTopicRuleListResponse) {
    response = &GetTopicRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取规则列表
func (c *Client) GetTopicRuleList(request *GetTopicRuleListRequest) (response *GetTopicRuleListResponse, err error) {
    if request == nil {
        request = NewGetTopicRuleListRequest()
    }
    response = NewGetTopicRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewListEventHistoryRequest() (request *ListEventHistoryRequest) {
    request = &ListEventHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ListEventHistory")
    return
}

func NewListEventHistoryResponse() (response *ListEventHistoryResponse) {
    response = &ListEventHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备的历史事件
func (c *Client) ListEventHistory(request *ListEventHistoryRequest) (response *ListEventHistoryResponse, err error) {
    if request == nil {
        request = NewListEventHistoryRequest()
    }
    response = NewListEventHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewListFirmwaresRequest() (request *ListFirmwaresRequest) {
    request = &ListFirmwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ListFirmwares")
    return
}

func NewListFirmwaresResponse() (response *ListFirmwaresResponse) {
    response = &ListFirmwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListFirmwares）用于获取固件列表 
func (c *Client) ListFirmwares(request *ListFirmwaresRequest) (response *ListFirmwaresResponse, err error) {
    if request == nil {
        request = NewListFirmwaresRequest()
    }
    response = NewListFirmwaresResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoRaFrequencyRequest() (request *ModifyLoRaFrequencyRequest) {
    request = &ModifyLoRaFrequencyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyLoRaFrequency")
    return
}

func NewModifyLoRaFrequencyResponse() (response *ModifyLoRaFrequencyResponse) {
    response = &ModifyLoRaFrequencyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改LoRa自定义频点
func (c *Client) ModifyLoRaFrequency(request *ModifyLoRaFrequencyRequest) (response *ModifyLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewModifyLoRaFrequencyRequest()
    }
    response = NewModifyLoRaFrequencyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoRaGatewayRequest() (request *ModifyLoRaGatewayRequest) {
    request = &ModifyLoRaGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyLoRaGateway")
    return
}

func NewModifyLoRaGatewayResponse() (response *ModifyLoRaGatewayResponse) {
    response = &ModifyLoRaGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改 LoRa 网关信息
func (c *Client) ModifyLoRaGateway(request *ModifyLoRaGatewayRequest) (response *ModifyLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewModifyLoRaGatewayRequest()
    }
    response = NewModifyLoRaGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelDefinitionRequest() (request *ModifyModelDefinitionRequest) {
    request = &ModifyModelDefinitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyModelDefinition")
    return
}

func NewModifyModelDefinitionResponse() (response *ModifyModelDefinitionResponse) {
    response = &ModifyModelDefinitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供修改产品的数据模板的能力
func (c *Client) ModifyModelDefinition(request *ModifyModelDefinitionRequest) (response *ModifyModelDefinitionResponse, err error) {
    if request == nil {
        request = NewModifyModelDefinitionRequest()
    }
    response = NewModifyModelDefinitionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyProject")
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改项目
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStudioProductRequest() (request *ModifyStudioProductRequest) {
    request = &ModifyStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyStudioProduct")
    return
}

func NewModifyStudioProductResponse() (response *ModifyStudioProductResponse) {
    response = &ModifyStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供修改产品的名称和描述等信息的能力，对于已发布产品不允许进行修改。
func (c *Client) ModifyStudioProduct(request *ModifyStudioProductRequest) (response *ModifyStudioProductResponse, err error) {
    if request == nil {
        request = NewModifyStudioProductRequest()
    }
    response = NewModifyStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRuleRequest() (request *ModifyTopicRuleRequest) {
    request = &ModifyTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyTopicRule")
    return
}

func NewModifyTopicRuleResponse() (response *ModifyTopicRuleResponse) {
    response = &ModifyTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改规则
func (c *Client) ModifyTopicRule(request *ModifyTopicRuleRequest) (response *ModifyTopicRuleResponse, err error) {
    if request == nil {
        request = NewModifyTopicRuleRequest()
    }
    response = NewModifyTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewPublishMessageRequest() (request *PublishMessageRequest) {
    request = &PublishMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PublishMessage")
    return
}

func NewPublishMessageResponse() (response *PublishMessageResponse) {
    response = &PublishMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（PublishMessage）用于使用自定义透传协议进行设备远控
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    if request == nil {
        request = NewPublishMessageRequest()
    }
    response = NewPublishMessageResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseStudioProductRequest() (request *ReleaseStudioProductRequest) {
    request = &ReleaseStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ReleaseStudioProduct")
    return
}

func NewReleaseStudioProductResponse() (response *ReleaseStudioProductResponse) {
    response = &ReleaseStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 产品开发完成并测试通过后，通过发布产品将产品设置为发布状态
func (c *Client) ReleaseStudioProduct(request *ReleaseStudioProductRequest) (response *ReleaseStudioProductResponse, err error) {
    if request == nil {
        request = NewReleaseStudioProductRequest()
    }
    response = NewReleaseStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStudioProductRequest() (request *SearchStudioProductRequest) {
    request = &SearchStudioProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "SearchStudioProduct")
    return
}

func NewSearchStudioProductResponse() (response *SearchStudioProductResponse) {
    response = &SearchStudioProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提供根据产品名称查找产品的能力
func (c *Client) SearchStudioProduct(request *SearchStudioProductRequest) (response *SearchStudioProductResponse, err error) {
    if request == nil {
        request = NewSearchStudioProductRequest()
    }
    response = NewSearchStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewSearchTopicRuleRequest() (request *SearchTopicRuleRequest) {
    request = &SearchTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "SearchTopicRule")
    return
}

func NewSearchTopicRuleResponse() (response *SearchTopicRuleResponse) {
    response = &SearchTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 搜索规则
func (c *Client) SearchTopicRule(request *SearchTopicRuleRequest) (response *SearchTopicRuleResponse, err error) {
    if request == nil {
        request = NewSearchTopicRuleRequest()
    }
    response = NewSearchTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFirmwareRequest() (request *UpdateFirmwareRequest) {
    request = &UpdateFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UpdateFirmware")
    return
}

func NewUpdateFirmwareResponse() (response *UpdateFirmwareResponse) {
    response = &UpdateFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateFirmware）用于对指定设备发起固件升级请求 
func (c *Client) UpdateFirmware(request *UpdateFirmwareRequest) (response *UpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewUpdateFirmwareRequest()
    }
    response = NewUpdateFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFirmwareRequest() (request *UploadFirmwareRequest) {
    request = &UploadFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UploadFirmware")
    return
}

func NewUploadFirmwareResponse() (response *UploadFirmwareResponse) {
    response = &UploadFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UploadFirmware）用于上传设备固件至平台
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}
