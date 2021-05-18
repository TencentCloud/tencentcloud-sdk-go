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

package v20201215

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-15"

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


func NewBatchUpdateFirmwareRequest() (request *BatchUpdateFirmwareRequest) {
    request = &BatchUpdateFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "BatchUpdateFirmware")
    return
}

func NewBatchUpdateFirmwareResponse() (response *BatchUpdateFirmwareResponse) {
    response = &BatchUpdateFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（BatchUpdateFirmware）用于批量更新设备固件 
func (c *Client) BatchUpdateFirmware(request *BatchUpdateFirmwareRequest) (response *BatchUpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewBatchUpdateFirmwareRequest()
    }
    response = NewBatchUpdateFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewCancelDeviceFirmwareTaskRequest() (request *CancelDeviceFirmwareTaskRequest) {
    request = &CancelDeviceFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CancelDeviceFirmwareTask")
    return
}

func NewCancelDeviceFirmwareTaskResponse() (response *CancelDeviceFirmwareTaskResponse) {
    response = &CancelDeviceFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于取消设备升级任务
func (c *Client) CancelDeviceFirmwareTask(request *CancelDeviceFirmwareTaskRequest) (response *CancelDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewCancelDeviceFirmwareTaskRequest()
    }
    response = NewCancelDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCheckForwardAuthRequest() (request *CheckForwardAuthRequest) {
    request = &CheckForwardAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CheckForwardAuth")
    return
}

func NewCheckForwardAuthResponse() (response *CheckForwardAuthResponse) {
    response = &CheckForwardAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 判断是否开启的转发的权限
func (c *Client) CheckForwardAuth(request *CheckForwardAuthRequest) (response *CheckForwardAuthResponse, err error) {
    if request == nil {
        request = NewCheckForwardAuthRequest()
    }
    response = NewCheckForwardAuthResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchRequest() (request *CreateBatchRequest) {
    request = &CreateBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateBatch")
    return
}

func NewCreateBatchResponse() (response *CreateBatchResponse) {
    response = &CreateBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建批次
func (c *Client) CreateBatch(request *CreateBatchRequest) (response *CreateBatchResponse, err error) {
    if request == nil {
        request = NewCreateBatchRequest()
    }
    response = NewCreateBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudStorageRequest() (request *CreateCloudStorageRequest) {
    request = &CreateCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateCloudStorage")
    return
}

func NewCreateCloudStorageResponse() (response *CreateCloudStorageResponse) {
    response = &CreateCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开通云存服务
func (c *Client) CreateCloudStorage(request *CreateCloudStorageRequest) (response *CreateCloudStorageResponse, err error) {
    if request == nil {
        request = NewCreateCloudStorageRequest()
    }
    response = NewCreateCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateForwardRuleRequest() (request *CreateForwardRuleRequest) {
    request = &CreateForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateForwardRule")
    return
}

func NewCreateForwardRuleResponse() (response *CreateForwardRuleResponse) {
    response = &CreateForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建转发规则
func (c *Client) CreateForwardRule(request *CreateForwardRuleRequest) (response *CreateForwardRuleResponse, err error) {
    if request == nil {
        request = NewCreateForwardRuleRequest()
    }
    response = NewCreateForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProductRequest() (request *CreateProductRequest) {
    request = &CreateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateProduct")
    return
}

func NewCreateProductResponse() (response *CreateProductResponse) {
    response = &CreateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建产品
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFileUrlRequest() (request *CreateTaskFileUrlRequest) {
    request = &CreateTaskFileUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateTaskFileUrl")
    return
}

func NewCreateTaskFileUrlResponse() (response *CreateTaskFileUrlResponse) {
    response = &CreateTaskFileUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateTaskFileUrl）用于获取产品级任务文件上传链接
func (c *Client) CreateTaskFileUrl(request *CreateTaskFileUrlRequest) (response *CreateTaskFileUrlResponse, err error) {
    if request == nil {
        request = NewCreateTaskFileUrlRequest()
    }
    response = NewCreateTaskFileUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteDevice")
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

func NewDeleteFirmwareRequest() (request *DeleteFirmwareRequest) {
    request = &DeleteFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteFirmware")
    return
}

func NewDeleteFirmwareResponse() (response *DeleteFirmwareResponse) {
    response = &DeleteFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteFirmware）用于删除固件 
func (c *Client) DeleteFirmware(request *DeleteFirmwareRequest) (response *DeleteFirmwareResponse, err error) {
    if request == nil {
        request = NewDeleteFirmwareRequest()
    }
    response = NewDeleteFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteForwardRuleRequest() (request *DeleteForwardRuleRequest) {
    request = &DeleteForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteForwardRule")
    return
}

func NewDeleteForwardRuleResponse() (response *DeleteForwardRuleResponse) {
    response = &DeleteForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除转发规则
func (c *Client) DeleteForwardRule(request *DeleteForwardRuleRequest) (response *DeleteForwardRuleResponse, err error) {
    if request == nil {
        request = NewDeleteForwardRuleRequest()
    }
    response = NewDeleteForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteProduct")
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

func NewDescribeBalanceRequest() (request *DescribeBalanceRequest) {
    request = &DescribeBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBalance")
    return
}

func NewDescribeBalanceResponse() (response *DescribeBalanceResponse) {
    response = &DescribeBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询账户余额
func (c *Client) DescribeBalance(request *DescribeBalanceRequest) (response *DescribeBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeBalanceRequest()
    }
    response = NewDescribeBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBalanceTransactionsRequest() (request *DescribeBalanceTransactionsRequest) {
    request = &DescribeBalanceTransactionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBalanceTransactions")
    return
}

func NewDescribeBalanceTransactionsResponse() (response *DescribeBalanceTransactionsResponse) {
    response = &DescribeBalanceTransactionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取账户流水
func (c *Client) DescribeBalanceTransactions(request *DescribeBalanceTransactionsRequest) (response *DescribeBalanceTransactionsResponse, err error) {
    if request == nil {
        request = NewDescribeBalanceTransactionsRequest()
    }
    response = NewDescribeBalanceTransactionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchRequest() (request *DescribeBatchRequest) {
    request = &DescribeBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBatch")
    return
}

func NewDescribeBatchResponse() (response *DescribeBatchResponse) {
    response = &DescribeBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取批次详情
func (c *Client) DescribeBatch(request *DescribeBatchRequest) (response *DescribeBatchResponse, err error) {
    if request == nil {
        request = NewDescribeBatchRequest()
    }
    response = NewDescribeBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchsRequest() (request *DescribeBatchsRequest) {
    request = &DescribeBatchsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBatchs")
    return
}

func NewDescribeBatchsResponse() (response *DescribeBatchsResponse) {
    response = &DescribeBatchsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取批次列表
func (c *Client) DescribeBatchs(request *DescribeBatchsRequest) (response *DescribeBatchsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchsRequest()
    }
    response = NewDescribeBatchsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCategoryRequest() (request *DescribeCategoryRequest) {
    request = &DescribeCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCategory")
    return
}

func NewDescribeCategoryResponse() (response *DescribeCategoryResponse) {
    response = &DescribeCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Category详情
func (c *Client) DescribeCategory(request *DescribeCategoryRequest) (response *DescribeCategoryResponse, err error) {
    if request == nil {
        request = NewDescribeCategoryRequest()
    }
    response = NewDescribeCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageRequest() (request *DescribeCloudStorageRequest) {
    request = &DescribeCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorage")
    return
}

func NewDescribeCloudStorageResponse() (response *DescribeCloudStorageResponse) {
    response = &DescribeCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备云存服务详情
func (c *Client) DescribeCloudStorage(request *DescribeCloudStorageRequest) (response *DescribeCloudStorageResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageRequest()
    }
    response = NewDescribeCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageDateRequest() (request *DescribeCloudStorageDateRequest) {
    request = &DescribeCloudStorageDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorageDate")
    return
}

func NewDescribeCloudStorageDateResponse() (response *DescribeCloudStorageDateResponse) {
    response = &DescribeCloudStorageDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取具有云存的日期
func (c *Client) DescribeCloudStorageDate(request *DescribeCloudStorageDateRequest) (response *DescribeCloudStorageDateResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageDateRequest()
    }
    response = NewDescribeCloudStorageDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageEventsRequest() (request *DescribeCloudStorageEventsRequest) {
    request = &DescribeCloudStorageEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorageEvents")
    return
}

func NewDescribeCloudStorageEventsResponse() (response *DescribeCloudStorageEventsResponse) {
    response = &DescribeCloudStorageEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取云存事件列表
func (c *Client) DescribeCloudStorageEvents(request *DescribeCloudStorageEventsRequest) (response *DescribeCloudStorageEventsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageEventsRequest()
    }
    response = NewDescribeCloudStorageEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageThumbnailRequest() (request *DescribeCloudStorageThumbnailRequest) {
    request = &DescribeCloudStorageThumbnailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorageThumbnail")
    return
}

func NewDescribeCloudStorageThumbnailResponse() (response *DescribeCloudStorageThumbnailResponse) {
    response = &DescribeCloudStorageThumbnailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取云存事件缩略图
func (c *Client) DescribeCloudStorageThumbnail(request *DescribeCloudStorageThumbnailRequest) (response *DescribeCloudStorageThumbnailResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageThumbnailRequest()
    }
    response = NewDescribeCloudStorageThumbnailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageTimeRequest() (request *DescribeCloudStorageTimeRequest) {
    request = &DescribeCloudStorageTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorageTime")
    return
}

func NewDescribeCloudStorageTimeResponse() (response *DescribeCloudStorageTimeResponse) {
    response = &DescribeCloudStorageTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取某一天云存时间轴
func (c *Client) DescribeCloudStorageTime(request *DescribeCloudStorageTimeRequest) (response *DescribeCloudStorageTimeResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageTimeRequest()
    }
    response = NewDescribeCloudStorageTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDevice")
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看设备详情
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceActionHistoryRequest() (request *DescribeDeviceActionHistoryRequest) {
    request = &DescribeDeviceActionHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceActionHistory")
    return
}

func NewDescribeDeviceActionHistoryResponse() (response *DescribeDeviceActionHistoryResponse) {
    response = &DescribeDeviceActionHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为用户提供获取动作历史的能力。
func (c *Client) DescribeDeviceActionHistory(request *DescribeDeviceActionHistoryRequest) (response *DescribeDeviceActionHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceActionHistoryRequest()
    }
    response = NewDescribeDeviceActionHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceCommLogRequest() (request *DescribeDeviceCommLogRequest) {
    request = &DescribeDeviceCommLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceCommLog")
    return
}

func NewDescribeDeviceCommLogResponse() (response *DescribeDeviceCommLogResponse) {
    response = &DescribeDeviceCommLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备在指定时间范围内的通讯日志
func (c *Client) DescribeDeviceCommLog(request *DescribeDeviceCommLogRequest) (response *DescribeDeviceCommLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceCommLogRequest()
    }
    response = NewDescribeDeviceCommLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceDataRequest() (request *DescribeDeviceDataRequest) {
    request = &DescribeDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceData")
    return
}

func NewDescribeDeviceDataResponse() (response *DescribeDeviceDataResponse) {
    response = &DescribeDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备属性数据
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
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceDataHistory")
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

func NewDescribeDeviceEventHistoryRequest() (request *DescribeDeviceEventHistoryRequest) {
    request = &DescribeDeviceEventHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceEventHistory")
    return
}

func NewDescribeDeviceEventHistoryResponse() (response *DescribeDeviceEventHistoryResponse) {
    response = &DescribeDeviceEventHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备的历史事件
func (c *Client) DescribeDeviceEventHistory(request *DescribeDeviceEventHistoryRequest) (response *DescribeDeviceEventHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceEventHistoryRequest()
    }
    response = NewDescribeDeviceEventHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDevices")
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取设备列表
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareRequest() (request *DescribeFirmwareRequest) {
    request = &DescribeFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmware")
    return
}

func NewDescribeFirmwareResponse() (response *DescribeFirmwareResponse) {
    response = &DescribeFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFirmware）用于查询固件信息
func (c *Client) DescribeFirmware(request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareRequest()
    }
    response = NewDescribeFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskRequest() (request *DescribeFirmwareTaskRequest) {
    request = &DescribeFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTask")
    return
}

func NewDescribeFirmwareTaskResponse() (response *DescribeFirmwareTaskResponse) {
    response = &DescribeFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 此接口查询固件升级任务详情
func (c *Client) DescribeFirmwareTask(request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    response = NewDescribeFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskDevicesRequest() (request *DescribeFirmwareTaskDevicesRequest) {
    request = &DescribeFirmwareTaskDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskDevices")
    return
}

func NewDescribeFirmwareTaskDevicesResponse() (response *DescribeFirmwareTaskDevicesResponse) {
    response = &DescribeFirmwareTaskDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询固件升级任务的设备列表
func (c *Client) DescribeFirmwareTaskDevices(request *DescribeFirmwareTaskDevicesRequest) (response *DescribeFirmwareTaskDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDevicesRequest()
    }
    response = NewDescribeFirmwareTaskDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskDistributionRequest() (request *DescribeFirmwareTaskDistributionRequest) {
    request = &DescribeFirmwareTaskDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskDistribution")
    return
}

func NewDescribeFirmwareTaskDistributionResponse() (response *DescribeFirmwareTaskDistributionResponse) {
    response = &DescribeFirmwareTaskDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询固件升级任务状态分布
func (c *Client) DescribeFirmwareTaskDistribution(request *DescribeFirmwareTaskDistributionRequest) (response *DescribeFirmwareTaskDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDistributionRequest()
    }
    response = NewDescribeFirmwareTaskDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskStatisticsRequest() (request *DescribeFirmwareTaskStatisticsRequest) {
    request = &DescribeFirmwareTaskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskStatistics")
    return
}

func NewDescribeFirmwareTaskStatisticsResponse() (response *DescribeFirmwareTaskStatisticsResponse) {
    response = &DescribeFirmwareTaskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询固件升级任务统计信息
func (c *Client) DescribeFirmwareTaskStatistics(request *DescribeFirmwareTaskStatisticsRequest) (response *DescribeFirmwareTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskStatisticsRequest()
    }
    response = NewDescribeFirmwareTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTasksRequest() (request *DescribeFirmwareTasksRequest) {
    request = &DescribeFirmwareTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTasks")
    return
}

func NewDescribeFirmwareTasksResponse() (response *DescribeFirmwareTasksResponse) {
    response = &DescribeFirmwareTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于查询固件升级任务列表
func (c *Client) DescribeFirmwareTasks(request *DescribeFirmwareTasksRequest) (response *DescribeFirmwareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTasksRequest()
    }
    response = NewDescribeFirmwareTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardRuleRequest() (request *DescribeForwardRuleRequest) {
    request = &DescribeForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeForwardRule")
    return
}

func NewDescribeForwardRuleResponse() (response *DescribeForwardRuleResponse) {
    response = &DescribeForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品转发规则
func (c *Client) DescribeForwardRule(request *DescribeForwardRuleRequest) (response *DescribeForwardRuleResponse, err error) {
    if request == nil {
        request = NewDescribeForwardRuleRequest()
    }
    response = NewDescribeForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelDefinitionRequest() (request *DescribeModelDefinitionRequest) {
    request = &DescribeModelDefinitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeModelDefinition")
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

func NewDescribeProductRequest() (request *DescribeProductRequest) {
    request = &DescribeProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeProduct")
    return
}

func NewDescribeProductResponse() (response *DescribeProductResponse) {
    response = &DescribeProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品详情
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    if request == nil {
        request = NewDescribeProductRequest()
    }
    response = NewDescribeProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
    request = &DescribeProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeProducts")
    return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
    response = &DescribeProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取产品列表
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewEditFirmwareRequest() (request *EditFirmwareRequest) {
    request = &EditFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "EditFirmware")
    return
}

func NewEditFirmwareResponse() (response *EditFirmwareResponse) {
    response = &EditFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于编辑固件信息
func (c *Client) EditFirmware(request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    if request == nil {
        request = NewEditFirmwareRequest()
    }
    response = NewEditFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateSignedVideoURLRequest() (request *GenerateSignedVideoURLRequest) {
    request = &GenerateSignedVideoURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "GenerateSignedVideoURL")
    return
}

func NewGenerateSignedVideoURLResponse() (response *GenerateSignedVideoURLResponse) {
    response = &GenerateSignedVideoURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取视频防盗链播放URL
func (c *Client) GenerateSignedVideoURL(request *GenerateSignedVideoURLRequest) (response *GenerateSignedVideoURLResponse, err error) {
    if request == nil {
        request = NewGenerateSignedVideoURLRequest()
    }
    response = NewGenerateSignedVideoURLResponse()
    err = c.Send(request, response)
    return
}

func NewGetAllFirmwareVersionRequest() (request *GetAllFirmwareVersionRequest) {
    request = &GetAllFirmwareVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "GetAllFirmwareVersion")
    return
}

func NewGetAllFirmwareVersionResponse() (response *GetAllFirmwareVersionResponse) {
    response = &GetAllFirmwareVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetAllFirmwareVersion）用于获取所有的版本列表 
func (c *Client) GetAllFirmwareVersion(request *GetAllFirmwareVersionRequest) (response *GetAllFirmwareVersionResponse, err error) {
    if request == nil {
        request = NewGetAllFirmwareVersionRequest()
    }
    response = NewGetAllFirmwareVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetFirmwareURLRequest() (request *GetFirmwareURLRequest) {
    request = &GetFirmwareURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "GetFirmwareURL")
    return
}

func NewGetFirmwareURLResponse() (response *GetFirmwareURLResponse) {
    response = &GetFirmwareURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetFirmwareURL）用于获取固件存储的URL 
func (c *Client) GetFirmwareURL(request *GetFirmwareURLRequest) (response *GetFirmwareURLResponse, err error) {
    if request == nil {
        request = NewGetFirmwareURLRequest()
    }
    response = NewGetFirmwareURLResponse()
    err = c.Send(request, response)
    return
}

func NewImportModelDefinitionRequest() (request *ImportModelDefinitionRequest) {
    request = &ImportModelDefinitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ImportModelDefinition")
    return
}

func NewImportModelDefinitionResponse() (response *ImportModelDefinitionResponse) {
    response = &ImportModelDefinitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 导入其它产品的数据模板，覆盖现有数据模板的物模型和产品分类信息
func (c *Client) ImportModelDefinition(request *ImportModelDefinitionRequest) (response *ImportModelDefinitionResponse, err error) {
    if request == nil {
        request = NewImportModelDefinitionRequest()
    }
    response = NewImportModelDefinitionResponse()
    err = c.Send(request, response)
    return
}

func NewListFirmwaresRequest() (request *ListFirmwaresRequest) {
    request = &ListFirmwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ListFirmwares")
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

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
    request = &ModifyDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDevice")
    return
}

func NewModifyDeviceResponse() (response *ModifyDeviceResponse) {
    response = &ModifyDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改设备信息
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    if request == nil {
        request = NewModifyDeviceRequest()
    }
    response = NewModifyDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyForwardRuleRequest() (request *ModifyForwardRuleRequest) {
    request = &ModifyForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyForwardRule")
    return
}

func NewModifyForwardRuleResponse() (response *ModifyForwardRuleResponse) {
    response = &ModifyForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改转发规则
func (c *Client) ModifyForwardRule(request *ModifyForwardRuleRequest) (response *ModifyForwardRuleResponse, err error) {
    if request == nil {
        request = NewModifyForwardRuleRequest()
    }
    response = NewModifyForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelDefinitionRequest() (request *ModifyModelDefinitionRequest) {
    request = &ModifyModelDefinitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyModelDefinition")
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

func NewModifyProductRequest() (request *ModifyProductRequest) {
    request = &ModifyProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyProduct")
    return
}

func NewModifyProductResponse() (response *ModifyProductResponse) {
    response = &ModifyProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改产品信息
func (c *Client) ModifyProduct(request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    if request == nil {
        request = NewModifyProductRequest()
    }
    response = NewModifyProductResponse()
    err = c.Send(request, response)
    return
}

func NewResetCloudStorageRequest() (request *ResetCloudStorageRequest) {
    request = &ResetCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ResetCloudStorage")
    return
}

func NewResetCloudStorageResponse() (response *ResetCloudStorageResponse) {
    response = &ResetCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重置云存服务
func (c *Client) ResetCloudStorage(request *ResetCloudStorageRequest) (response *ResetCloudStorageResponse, err error) {
    if request == nil {
        request = NewResetCloudStorageRequest()
    }
    response = NewResetCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDeviceFirmwareTaskRequest() (request *RetryDeviceFirmwareTaskRequest) {
    request = &RetryDeviceFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RetryDeviceFirmwareTask")
    return
}

func NewRetryDeviceFirmwareTaskResponse() (response *RetryDeviceFirmwareTaskResponse) {
    response = &RetryDeviceFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于重试设备升级任务
func (c *Client) RetryDeviceFirmwareTask(request *RetryDeviceFirmwareTaskRequest) (response *RetryDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewRetryDeviceFirmwareTaskRequest()
    }
    response = NewRetryDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSetForwardAuthRequest() (request *SetForwardAuthRequest) {
    request = &SetForwardAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "SetForwardAuth")
    return
}

func NewSetForwardAuthResponse() (response *SetForwardAuthResponse) {
    response = &SetForwardAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置转发权限
func (c *Client) SetForwardAuth(request *SetForwardAuthRequest) (response *SetForwardAuthResponse, err error) {
    if request == nil {
        request = NewSetForwardAuthRequest()
    }
    response = NewSetForwardAuthResponse()
    err = c.Send(request, response)
    return
}

func NewTransferCloudStorageRequest() (request *TransferCloudStorageRequest) {
    request = &TransferCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "TransferCloudStorage")
    return
}

func NewTransferCloudStorageResponse() (response *TransferCloudStorageResponse) {
    response = &TransferCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 转移云存服务
func (c *Client) TransferCloudStorage(request *TransferCloudStorageRequest) (response *TransferCloudStorageResponse, err error) {
    if request == nil {
        request = NewTransferCloudStorageRequest()
    }
    response = NewTransferCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFirmwareRequest() (request *UploadFirmwareRequest) {
    request = &UploadFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "UploadFirmware")
    return
}

func NewUploadFirmwareResponse() (response *UploadFirmwareResponse) {
    response = &UploadFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UploadFirmware）用于上传设备固件信息 
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}
