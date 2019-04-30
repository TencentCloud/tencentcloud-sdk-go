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

func NewBuyDevicesRequest() (request *BuyDevicesRequest) {
    request = &BuyDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "BuyDevices")
    return
}

func NewBuyDevicesResponse() (response *BuyDevicesResponse) {
    response = &BuyDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 购买黑石物理机
func (c *Client) BuyDevices(request *BuyDevicesRequest) (response *BuyDevicesResponse, err error) {
    if request == nil {
        request = NewBuyDevicesRequest()
    }
    response = NewBuyDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomImageRequest() (request *CreateCustomImageRequest) {
    request = &CreateCustomImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "CreateCustomImage")
    return
}

func NewCreateCustomImageResponse() (response *CreateCustomImageResponse) {
    response = &CreateCustomImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建自定义镜像<br>
// 每个AppId在每个可用区最多保留20个自定义镜像
func (c *Client) CreateCustomImage(request *CreateCustomImageRequest) (response *CreateCustomImageResponse, err error) {
    if request == nil {
        request = NewCreateCustomImageRequest()
    }
    response = NewCreateCustomImageResponse()
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

func NewDeleteCustomImagesRequest() (request *DeleteCustomImagesRequest) {
    request = &DeleteCustomImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DeleteCustomImages")
    return
}

func NewDeleteCustomImagesResponse() (response *DeleteCustomImagesResponse) {
    response = &DeleteCustomImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除自定义镜像<br>
// 正用于部署或重装中的镜像被删除后，镜像文件将保留一段时间，直到部署或重装结束
func (c *Client) DeleteCustomImages(request *DeleteCustomImagesRequest) (response *DeleteCustomImagesResponse, err error) {
    if request == nil {
        request = NewDeleteCustomImagesRequest()
    }
    response = NewDeleteCustomImagesResponse()
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

func NewDeleteUserCmdsRequest() (request *DeleteUserCmdsRequest) {
    request = &DeleteUserCmdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DeleteUserCmds")
    return
}

func NewDeleteUserCmdsResponse() (response *DeleteUserCmdsResponse) {
    response = &DeleteUserCmdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除自定义脚本
func (c *Client) DeleteUserCmds(request *DeleteUserCmdsRequest) (response *DeleteUserCmdsResponse, err error) {
    if request == nil {
        request = NewDeleteUserCmdsRequest()
    }
    response = NewDeleteUserCmdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImageProcessRequest() (request *DescribeCustomImageProcessRequest) {
    request = &DescribeCustomImageProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeCustomImageProcess")
    return
}

func NewDescribeCustomImageProcessResponse() (response *DescribeCustomImageProcessResponse) {
    response = &DescribeCustomImageProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询自定义镜像制作进度
func (c *Client) DescribeCustomImageProcess(request *DescribeCustomImageProcessRequest) (response *DescribeCustomImageProcessResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImageProcessRequest()
    }
    response = NewDescribeCustomImageProcessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImagesRequest() (request *DescribeCustomImagesRequest) {
    request = &DescribeCustomImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeCustomImages")
    return
}

func NewDescribeCustomImagesResponse() (response *DescribeCustomImagesResponse) {
    response = &DescribeCustomImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看自定义镜像列表
func (c *Client) DescribeCustomImages(request *DescribeCustomImagesRequest) (response *DescribeCustomImagesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImagesRequest()
    }
    response = NewDescribeCustomImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceClassRequest() (request *DescribeDeviceClassRequest) {
    request = &DescribeDeviceClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceClass")
    return
}

func NewDescribeDeviceClassResponse() (response *DescribeDeviceClassResponse) {
    response = &DescribeDeviceClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取获取设备类型
func (c *Client) DescribeDeviceClass(request *DescribeDeviceClassRequest) (response *DescribeDeviceClassResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClassRequest()
    }
    response = NewDescribeDeviceClassResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceClassPartitionRequest() (request *DescribeDeviceClassPartitionRequest) {
    request = &DescribeDeviceClassPartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceClassPartition")
    return
}

func NewDescribeDeviceClassPartitionResponse() (response *DescribeDeviceClassPartitionResponse) {
    response = &DescribeDeviceClassPartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询机型支持的RAID方式， 并返回系统盘的分区和逻辑盘的列表
func (c *Client) DescribeDeviceClassPartition(request *DescribeDeviceClassPartitionRequest) (response *DescribeDeviceClassPartitionResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClassPartitionRequest()
    }
    response = NewDescribeDeviceClassPartitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceHardwareInfoRequest() (request *DescribeDeviceHardwareInfoRequest) {
    request = &DescribeDeviceHardwareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceHardwareInfo")
    return
}

func NewDescribeDeviceHardwareInfoResponse() (response *DescribeDeviceHardwareInfoResponse) {
    response = &DescribeDeviceHardwareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询设备硬件配置信息，如 CPU 型号，内存大小，磁盘大小和数量
func (c *Client) DescribeDeviceHardwareInfo(request *DescribeDeviceHardwareInfoRequest) (response *DescribeDeviceHardwareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceHardwareInfoRequest()
    }
    response = NewDescribeDeviceHardwareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceInventoryRequest() (request *DescribeDeviceInventoryRequest) {
    request = &DescribeDeviceInventoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceInventory")
    return
}

func NewDescribeDeviceInventoryResponse() (response *DescribeDeviceInventoryResponse) {
    response = &DescribeDeviceInventoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询设备库存
func (c *Client) DescribeDeviceInventory(request *DescribeDeviceInventoryRequest) (response *DescribeDeviceInventoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceInventoryRequest()
    }
    response = NewDescribeDeviceInventoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceOperationLogRequest() (request *DescribeDeviceOperationLogRequest) {
    request = &DescribeDeviceOperationLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceOperationLog")
    return
}

func NewDescribeDeviceOperationLogResponse() (response *DescribeDeviceOperationLogResponse) {
    response = &DescribeDeviceOperationLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询设备操作日志， 如设备重启，重装，设置密码等操作
func (c *Client) DescribeDeviceOperationLog(request *DescribeDeviceOperationLogRequest) (response *DescribeDeviceOperationLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceOperationLogRequest()
    }
    response = NewDescribeDeviceOperationLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePartitionRequest() (request *DescribeDevicePartitionRequest) {
    request = &DescribeDevicePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePartition")
    return
}

func NewDescribeDevicePartitionResponse() (response *DescribeDevicePartitionResponse) {
    response = &DescribeDevicePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取物理机的分区格式
func (c *Client) DescribeDevicePartition(request *DescribeDevicePartitionRequest) (response *DescribeDevicePartitionResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePartitionRequest()
    }
    response = NewDescribeDevicePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePositionRequest() (request *DescribeDevicePositionRequest) {
    request = &DescribeDevicePositionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePosition")
    return
}

func NewDescribeDevicePositionResponse() (response *DescribeDevicePositionResponse) {
    response = &DescribeDevicePositionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务器所在的位置，如机架，上联交换机等信息
func (c *Client) DescribeDevicePosition(request *DescribeDevicePositionRequest) (response *DescribeDevicePositionResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePositionRequest()
    }
    response = NewDescribeDevicePositionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePriceInfoRequest() (request *DescribeDevicePriceInfoRequest) {
    request = &DescribeDevicePriceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePriceInfo")
    return
}

func NewDescribeDevicePriceInfoResponse() (response *DescribeDevicePriceInfoResponse) {
    response = &DescribeDevicePriceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询服务器价格信息，支持设备的批量查找，支持标准机型和弹性机型的混合查找
func (c *Client) DescribeDevicePriceInfo(request *DescribeDevicePriceInfoRequest) (response *DescribeDevicePriceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePriceInfoRequest()
    }
    response = NewDescribeDevicePriceInfoResponse()
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

func NewDescribeHardwareSpecificationRequest() (request *DescribeHardwareSpecificationRequest) {
    request = &DescribeHardwareSpecificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeHardwareSpecification")
    return
}

func NewDescribeHardwareSpecificationResponse() (response *DescribeHardwareSpecificationResponse) {
    response = &DescribeHardwareSpecificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询自定义机型部件信息，包括CpuId对应的型号，DiskTypeId对应的磁盘类型
func (c *Client) DescribeHardwareSpecification(request *DescribeHardwareSpecificationRequest) (response *DescribeHardwareSpecificationResponse, err error) {
    if request == nil {
        request = NewDescribeHardwareSpecificationRequest()
    }
    response = NewDescribeHardwareSpecificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostedDeviceOutBandInfoRequest() (request *DescribeHostedDeviceOutBandInfoRequest) {
    request = &DescribeHostedDeviceOutBandInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeHostedDeviceOutBandInfo")
    return
}

func NewDescribeHostedDeviceOutBandInfoResponse() (response *DescribeHostedDeviceOutBandInfoResponse) {
    response = &DescribeHostedDeviceOutBandInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询托管设备带外信息
func (c *Client) DescribeHostedDeviceOutBandInfo(request *DescribeHostedDeviceOutBandInfoRequest) (response *DescribeHostedDeviceOutBandInfoResponse, err error) {
    if request == nil {
        request = NewDescribeHostedDeviceOutBandInfoRequest()
    }
    response = NewDescribeHostedDeviceOutBandInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperationResultRequest() (request *DescribeOperationResultRequest) {
    request = &DescribeOperationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeOperationResult")
    return
}

func NewDescribeOperationResultResponse() (response *DescribeOperationResultResponse) {
    response = &DescribeOperationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取异步操作状态的完成状态
func (c *Client) DescribeOperationResult(request *DescribeOperationResultRequest) (response *DescribeOperationResultResponse, err error) {
    if request == nil {
        request = NewDescribeOperationResultRequest()
    }
    response = NewDescribeOperationResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOsInfoRequest() (request *DescribeOsInfoRequest) {
    request = &DescribeOsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeOsInfo")
    return
}

func NewDescribeOsInfoResponse() (response *DescribeOsInfoResponse) {
    response = &DescribeOsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定机型所支持的操作系统
func (c *Client) DescribeOsInfo(request *DescribeOsInfoRequest) (response *DescribeOsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOsInfoRequest()
    }
    response = NewDescribeOsInfoResponse()
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

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询地域以及可用区
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
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

func NewDescribeUserCmdTaskInfoRequest() (request *DescribeUserCmdTaskInfoRequest) {
    request = &DescribeUserCmdTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmdTaskInfo")
    return
}

func NewDescribeUserCmdTaskInfoResponse() (response *DescribeUserCmdTaskInfoResponse) {
    response = &DescribeUserCmdTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义脚本任务详细信息
func (c *Client) DescribeUserCmdTaskInfo(request *DescribeUserCmdTaskInfoRequest) (response *DescribeUserCmdTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdTaskInfoRequest()
    }
    response = NewDescribeUserCmdTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCmdTasksRequest() (request *DescribeUserCmdTasksRequest) {
    request = &DescribeUserCmdTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmdTasks")
    return
}

func NewDescribeUserCmdTasksResponse() (response *DescribeUserCmdTasksResponse) {
    response = &DescribeUserCmdTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义脚本任务列表
func (c *Client) DescribeUserCmdTasks(request *DescribeUserCmdTasksRequest) (response *DescribeUserCmdTasksResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdTasksRequest()
    }
    response = NewDescribeUserCmdTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCmdsRequest() (request *DescribeUserCmdsRequest) {
    request = &DescribeUserCmdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmds")
    return
}

func NewDescribeUserCmdsResponse() (response *DescribeUserCmdsResponse) {
    response = &DescribeUserCmdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义脚本信息列表
func (c *Client) DescribeUserCmds(request *DescribeUserCmdsRequest) (response *DescribeUserCmdsResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdsRequest()
    }
    response = NewDescribeUserCmdsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomImageAttributeRequest() (request *ModifyCustomImageAttributeRequest) {
    request = &ModifyCustomImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyCustomImageAttribute")
    return
}

func NewModifyCustomImageAttributeResponse() (response *ModifyCustomImageAttributeResponse) {
    response = &ModifyCustomImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于修改自定义镜像名或描述
func (c *Client) ModifyCustomImageAttribute(request *ModifyCustomImageAttributeRequest) (response *ModifyCustomImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomImageAttributeRequest()
    }
    response = NewModifyCustomImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceAliasesRequest() (request *ModifyDeviceAliasesRequest) {
    request = &ModifyDeviceAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyDeviceAliases")
    return
}

func NewModifyDeviceAliasesResponse() (response *ModifyDeviceAliasesResponse) {
    response = &ModifyDeviceAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改服务器名称
func (c *Client) ModifyDeviceAliases(request *ModifyDeviceAliasesRequest) (response *ModifyDeviceAliasesResponse, err error) {
    if request == nil {
        request = NewModifyDeviceAliasesRequest()
    }
    response = NewModifyDeviceAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceAutoRenewFlagRequest() (request *ModifyDeviceAutoRenewFlagRequest) {
    request = &ModifyDeviceAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyDeviceAutoRenewFlag")
    return
}

func NewModifyDeviceAutoRenewFlagResponse() (response *ModifyDeviceAutoRenewFlagResponse) {
    response = &ModifyDeviceAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改物理机服务器自动续费标志
func (c *Client) ModifyDeviceAutoRenewFlag(request *ModifyDeviceAutoRenewFlagRequest) (response *ModifyDeviceAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDeviceAutoRenewFlagRequest()
    }
    response = NewModifyDeviceAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLanIpRequest() (request *ModifyLanIpRequest) {
    request = &ModifyLanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyLanIp")
    return
}

func NewModifyLanIpResponse() (response *ModifyLanIpResponse) {
    response = &ModifyLanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改物理机内网IP（不重装系统）
func (c *Client) ModifyLanIp(request *ModifyLanIpRequest) (response *ModifyLanIpResponse, err error) {
    if request == nil {
        request = NewModifyLanIpRequest()
    }
    response = NewModifyLanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPayModePre2PostRequest() (request *ModifyPayModePre2PostRequest) {
    request = &ModifyPayModePre2PostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyPayModePre2Post")
    return
}

func NewModifyPayModePre2PostResponse() (response *ModifyPayModePre2PostResponse) {
    response = &ModifyPayModePre2PostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将设备的预付费模式修改为后付费计费模式，支持批量转换。（前提是客户要加入黑石物理机后付费计费的白名单，申请黑石物理机后付费可以联系腾讯云客服）
func (c *Client) ModifyPayModePre2Post(request *ModifyPayModePre2PostRequest) (response *ModifyPayModePre2PostResponse, err error) {
    if request == nil {
        request = NewModifyPayModePre2PostRequest()
    }
    response = NewModifyPayModePre2PostResponse()
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

func NewModifyUserCmdRequest() (request *ModifyUserCmdRequest) {
    request = &ModifyUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ModifyUserCmd")
    return
}

func NewModifyUserCmdResponse() (response *ModifyUserCmdResponse) {
    response = &ModifyUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改自定义脚本
func (c *Client) ModifyUserCmd(request *ModifyUserCmdRequest) (response *ModifyUserCmdResponse, err error) {
    if request == nil {
        request = NewModifyUserCmdRequest()
    }
    response = NewModifyUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineDevicesRequest() (request *OfflineDevicesRequest) {
    request = &OfflineDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "OfflineDevices")
    return
}

func NewOfflineDevicesResponse() (response *OfflineDevicesResponse) {
    response = &OfflineDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁黑石物理机实例：可以销毁物理机列表中的竞价实例，或回收站列表中所有计费模式的实例
func (c *Client) OfflineDevices(request *OfflineDevicesRequest) (response *OfflineDevicesResponse, err error) {
    if request == nil {
        request = NewOfflineDevicesRequest()
    }
    response = NewOfflineDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRebootDevicesRequest() (request *RebootDevicesRequest) {
    request = &RebootDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "RebootDevices")
    return
}

func NewRebootDevicesResponse() (response *RebootDevicesResponse) {
    response = &RebootDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重启机器
func (c *Client) RebootDevices(request *RebootDevicesRequest) (response *RebootDevicesResponse, err error) {
    if request == nil {
        request = NewRebootDevicesRequest()
    }
    response = NewRebootDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverDevicesRequest() (request *RecoverDevicesRequest) {
    request = &RecoverDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "RecoverDevices")
    return
}

func NewRecoverDevicesResponse() (response *RecoverDevicesResponse) {
    response = &RecoverDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复回收站中的物理机（仅限后付费的物理机）
func (c *Client) RecoverDevices(request *RecoverDevicesRequest) (response *RecoverDevicesResponse, err error) {
    if request == nil {
        request = NewRecoverDevicesRequest()
    }
    response = NewRecoverDevicesResponse()
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

func NewResetDevicePasswordRequest() (request *ResetDevicePasswordRequest) {
    request = &ResetDevicePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ResetDevicePassword")
    return
}

func NewResetDevicePasswordResponse() (response *ResetDevicePasswordResponse) {
    response = &ResetDevicePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重置服务器密码
func (c *Client) ResetDevicePassword(request *ResetDevicePasswordRequest) (response *ResetDevicePasswordResponse, err error) {
    if request == nil {
        request = NewResetDevicePasswordRequest()
    }
    response = NewResetDevicePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewReturnDevicesRequest() (request *ReturnDevicesRequest) {
    request = &ReturnDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ReturnDevices")
    return
}

func NewReturnDevicesResponse() (response *ReturnDevicesResponse) {
    response = &ReturnDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 退回物理机至回收站，支持批量退还不同计费模式的物理机（包括预付费、后付费、预付费转后付费）
func (c *Client) ReturnDevices(request *ReturnDevicesRequest) (response *ReturnDevicesResponse, err error) {
    if request == nil {
        request = NewReturnDevicesRequest()
    }
    response = NewReturnDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRunUserCmdRequest() (request *RunUserCmdRequest) {
    request = &RunUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "RunUserCmd")
    return
}

func NewRunUserCmdResponse() (response *RunUserCmdResponse) {
    response = &RunUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运行自定义脚本
func (c *Client) RunUserCmd(request *RunUserCmdRequest) (response *RunUserCmdResponse, err error) {
    if request == nil {
        request = NewRunUserCmdRequest()
    }
    response = NewRunUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewSetOutBandVpnAuthPasswordRequest() (request *SetOutBandVpnAuthPasswordRequest) {
    request = &SetOutBandVpnAuthPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "SetOutBandVpnAuthPassword")
    return
}

func NewSetOutBandVpnAuthPasswordResponse() (response *SetOutBandVpnAuthPasswordResponse) {
    response = &SetOutBandVpnAuthPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置带外VPN认证用户密码
func (c *Client) SetOutBandVpnAuthPassword(request *SetOutBandVpnAuthPasswordRequest) (response *SetOutBandVpnAuthPasswordResponse, err error) {
    if request == nil {
        request = NewSetOutBandVpnAuthPasswordRequest()
    }
    response = NewSetOutBandVpnAuthPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewShutdownDevicesRequest() (request *ShutdownDevicesRequest) {
    request = &ShutdownDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "ShutdownDevices")
    return
}

func NewShutdownDevicesResponse() (response *ShutdownDevicesResponse) {
    response = &ShutdownDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关闭服务器
func (c *Client) ShutdownDevices(request *ShutdownDevicesRequest) (response *ShutdownDevicesResponse, err error) {
    if request == nil {
        request = NewShutdownDevicesRequest()
    }
    response = NewShutdownDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewStartDevicesRequest() (request *StartDevicesRequest) {
    request = &StartDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bm", APIVersion, "StartDevices")
    return
}

func NewStartDevicesResponse() (response *StartDevicesResponse) {
    response = &StartDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启服务器
func (c *Client) StartDevices(request *StartDevicesRequest) (response *StartDevicesResponse, err error) {
    if request == nil {
        request = NewStartDevicesRequest()
    }
    response = NewStartDevicesResponse()
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
