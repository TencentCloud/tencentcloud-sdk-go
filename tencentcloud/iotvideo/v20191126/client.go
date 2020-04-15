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

package v20191126

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-26"

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


func NewCreateAppUsrRequest() (request *CreateAppUsrRequest) {
    request = &CreateAppUsrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateAppUsr")
    return
}

func NewCreateAppUsrResponse() (response *CreateAppUsrResponse) {
    response = &CreateAppUsrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAppUsr）用于接收由厂商云发送过来的注册请求,建立厂商云终端用户与IoT Video终端用户的映射关系。
func (c *Client) CreateAppUsr(request *CreateAppUsrRequest) (response *CreateAppUsrResponse, err error) {
    if request == nil {
        request = NewCreateAppUsrRequest()
    }
    response = NewCreateAppUsrResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBindingRequest() (request *CreateBindingRequest) {
    request = &CreateBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateBinding")
    return
}

func NewCreateBindingResponse() (response *CreateBindingResponse) {
    response = &CreateBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateBinding）用于终端用户和设备进行绑定，具体的应用场景如下：
//     终端用户与设备具有“强关联”关系。用户与设备绑定之后，用户终端即具备了该设备的访问权限,访问或操作设备时，无需获取设备访问Token。
func (c *Client) CreateBinding(request *CreateBindingRequest) (response *CreateBindingResponse, err error) {
    if request == nil {
        request = NewCreateBindingRequest()
    }
    response = NewCreateBindingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDevTokenRequest() (request *CreateDevTokenRequest) {
    request = &CreateDevTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateDevToken")
    return
}

func NewCreateDevTokenResponse() (response *CreateDevTokenResponse) {
    response = &CreateDevTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDevToken）用于以下场景：
// 终端用户与设备没有强绑定关联关系;
// 允许终端用户短时或一次性临时访问设备;
// 当终端用户与设备有强绑定关系时，可以不用调用此接口
func (c *Client) CreateDevToken(request *CreateDevTokenRequest) (response *CreateDevTokenResponse, err error) {
    if request == nil {
        request = NewCreateDevTokenRequest()
    }
    response = NewCreateDevTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDevicesRequest() (request *CreateDevicesRequest) {
    request = &CreateDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateDevices")
    return
}

func NewCreateDevicesResponse() (response *CreateDevicesResponse) {
    response = &CreateDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDevices）用于批量创建新的物联网视频通信设备。
// 注意：腾讯云不会对设备私钥进行保存，请自行保管好您的设备私钥。
func (c *Client) CreateDevices(request *CreateDevicesRequest) (response *CreateDevicesResponse, err error) {
    if request == nil {
        request = NewCreateDevicesRequest()
    }
    response = NewCreateDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGencodeRequest() (request *CreateGencodeRequest) {
    request = &CreateGencodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateGencode")
    return
}

func NewCreateGencodeResponse() (response *CreateGencodeResponse) {
    response = &CreateGencodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateGencode）用于生成设备物模型源代码
func (c *Client) CreateGencode(request *CreateGencodeRequest) (response *CreateGencodeResponse, err error) {
    if request == nil {
        request = NewCreateGencodeRequest()
    }
    response = NewCreateGencodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIotDataTypeRequest() (request *CreateIotDataTypeRequest) {
    request = &CreateIotDataTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateIotDataType")
    return
}

func NewCreateIotDataTypeResponse() (response *CreateIotDataTypeResponse) {
    response = &CreateIotDataTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateIotDataType）用于创建自定义物模型数据类型。
func (c *Client) CreateIotDataType(request *CreateIotDataTypeRequest) (response *CreateIotDataTypeResponse, err error) {
    if request == nil {
        request = NewCreateIotDataTypeRequest()
    }
    response = NewCreateIotDataTypeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIotModelRequest() (request *CreateIotModelRequest) {
    request = &CreateIotModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateIotModel")
    return
}

func NewCreateIotModelResponse() (response *CreateIotModelResponse) {
    response = &CreateIotModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateIotModel）用于定义的物模型提交。
// 该接口实现了物模型草稿箱的功能，保存用户最后一次编辑的物模型数据。
func (c *Client) CreateIotModel(request *CreateIotModelRequest) (response *CreateIotModelResponse, err error) {
    if request == nil {
        request = NewCreateIotModelRequest()
    }
    response = NewCreateIotModelResponse()
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

// 本接口（CreateProduct）用于创建一个新的物联网智能视频产品。
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStorageRequest() (request *CreateStorageRequest) {
    request = &CreateStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateStorage")
    return
}

func NewCreateStorageResponse() (response *CreateStorageResponse) {
    response = &CreateStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateStorage）用于购买云存套餐。
func (c *Client) CreateStorage(request *CreateStorageRequest) (response *CreateStorageResponse, err error) {
    if request == nil {
        request = NewCreateStorageRequest()
    }
    response = NewCreateStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTraceIdsRequest() (request *CreateTraceIdsRequest) {
    request = &CreateTraceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateTraceIds")
    return
}

func NewCreateTraceIdsResponse() (response *CreateTraceIdsResponse) {
    response = &CreateTraceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateTraceIds）用于将设备加到日志跟踪白名单。
func (c *Client) CreateTraceIds(request *CreateTraceIdsRequest) (response *CreateTraceIdsResponse, err error) {
    if request == nil {
        request = NewCreateTraceIdsRequest()
    }
    response = NewCreateTraceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUploadPathRequest() (request *CreateUploadPathRequest) {
    request = &CreateUploadPathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateUploadPath")
    return
}

func NewCreateUploadPathResponse() (response *CreateUploadPathResponse) {
    response = &CreateUploadPathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateUploadPath）用于获取固件上传路径。
func (c *Client) CreateUploadPath(request *CreateUploadPathRequest) (response *CreateUploadPathResponse, err error) {
    if request == nil {
        request = NewCreateUploadPathRequest()
    }
    response = NewCreateUploadPathResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsrTokenRequest() (request *CreateUsrTokenRequest) {
    request = &CreateUsrTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateUsrToken")
    return
}

func NewCreateUsrTokenResponse() (response *CreateUsrTokenResponse) {
    response = &CreateUsrTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateUsrToken）用于终端用户获取IoT Video平台的accessToken，初始化SDK,连接到IoT Video接入服务器。
func (c *Client) CreateUsrToken(request *CreateUsrTokenRequest) (response *CreateUsrTokenResponse, err error) {
    if request == nil {
        request = NewCreateUsrTokenRequest()
    }
    response = NewCreateUsrTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppUsrRequest() (request *DeleteAppUsrRequest) {
    request = &DeleteAppUsrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteAppUsr")
    return
}

func NewDeleteAppUsrResponse() (response *DeleteAppUsrResponse) {
    response = &DeleteAppUsrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAppUsr）用于删除终端用户。
func (c *Client) DeleteAppUsr(request *DeleteAppUsrRequest) (response *DeleteAppUsrResponse, err error) {
    if request == nil {
        request = NewDeleteAppUsrRequest()
    }
    response = NewDeleteAppUsrResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBindingRequest() (request *DeleteBindingRequest) {
    request = &DeleteBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteBinding")
    return
}

func NewDeleteBindingResponse() (response *DeleteBindingResponse) {
    response = &DeleteBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteBinding）用于终端用户和设备进行解绑定。
func (c *Client) DeleteBinding(request *DeleteBindingRequest) (response *DeleteBindingResponse, err error) {
    if request == nil {
        request = NewDeleteBindingRequest()
    }
    response = NewDeleteBindingResponse()
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

// 本接口（DeleteDevice）用于删除设备，可进行批量操作，每次操作最多100台设备。
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIotDataTypeRequest() (request *DeleteIotDataTypeRequest) {
    request = &DeleteIotDataTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteIotDataType")
    return
}

func NewDeleteIotDataTypeResponse() (response *DeleteIotDataTypeResponse) {
    response = &DeleteIotDataTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteIotDataType）用于删除自定义物模型数据类型。
func (c *Client) DeleteIotDataType(request *DeleteIotDataTypeRequest) (response *DeleteIotDataTypeResponse, err error) {
    if request == nil {
        request = NewDeleteIotDataTypeRequest()
    }
    response = NewDeleteIotDataTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMessageQueueRequest() (request *DeleteMessageQueueRequest) {
    request = &DeleteMessageQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteMessageQueue")
    return
}

func NewDeleteMessageQueueResponse() (response *DeleteMessageQueueResponse) {
    response = &DeleteMessageQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteMessageQueue）用于删除物联网智能视频产品的转发消息配置信息。
func (c *Client) DeleteMessageQueue(request *DeleteMessageQueueRequest) (response *DeleteMessageQueueResponse, err error) {
    if request == nil {
        request = NewDeleteMessageQueueRequest()
    }
    response = NewDeleteMessageQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOtaVersionRequest() (request *DeleteOtaVersionRequest) {
    request = &DeleteOtaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteOtaVersion")
    return
}

func NewDeleteOtaVersionResponse() (response *DeleteOtaVersionResponse) {
    response = &DeleteOtaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteOtaVersion）用于删除固件版本信息。
func (c *Client) DeleteOtaVersion(request *DeleteOtaVersionRequest) (response *DeleteOtaVersionResponse, err error) {
    if request == nil {
        request = NewDeleteOtaVersionRequest()
    }
    response = NewDeleteOtaVersionResponse()
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

// 本接口（DeleteProduct）用于删除一个物联网智能视频产品。
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTraceIdsRequest() (request *DeleteTraceIdsRequest) {
    request = &DeleteTraceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeleteTraceIds")
    return
}

func NewDeleteTraceIdsResponse() (response *DeleteTraceIdsResponse) {
    response = &DeleteTraceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteTraceIds）用于将设备从日志跟踪白名单中删除，该接口可批量操作，最多支持同时操作100台设备。
func (c *Client) DeleteTraceIds(request *DeleteTraceIdsRequest) (response *DeleteTraceIdsResponse, err error) {
    if request == nil {
        request = NewDeleteTraceIdsRequest()
    }
    response = NewDeleteTraceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindDevRequest() (request *DescribeBindDevRequest) {
    request = &DescribeBindDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBindDev")
    return
}

func NewDescribeBindDevResponse() (response *DescribeBindDevResponse) {
    response = &DescribeBindDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBindDev）用于查询终端用户绑定的设备列表。
func (c *Client) DescribeBindDev(request *DescribeBindDevRequest) (response *DescribeBindDevResponse, err error) {
    if request == nil {
        request = NewDescribeBindDevRequest()
    }
    response = NewDescribeBindDevResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindUsrRequest() (request *DescribeBindUsrRequest) {
    request = &DescribeBindUsrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeBindUsr")
    return
}

func NewDescribeBindUsrResponse() (response *DescribeBindUsrResponse) {
    response = &DescribeBindUsrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBindUsr）用于查询设备被分享的所有用户列表。
func (c *Client) DescribeBindUsr(request *DescribeBindUsrRequest) (response *DescribeBindUsrResponse, err error) {
    if request == nil {
        request = NewDescribeBindUsrRequest()
    }
    response = NewDescribeBindUsrResponse()
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

// 本接口（DescribeDevice）获取设备信息。
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceModelRequest() (request *DescribeDeviceModelRequest) {
    request = &DescribeDeviceModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceModel")
    return
}

func NewDescribeDeviceModelResponse() (response *DescribeDeviceModelResponse) {
    response = &DescribeDeviceModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDeviceModel）用于获取设备物模型。
func (c *Client) DescribeDeviceModel(request *DescribeDeviceModelRequest) (response *DescribeDeviceModelResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceModelRequest()
    }
    response = NewDescribeDeviceModelResponse()
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

// 本接口（DescribeDevices）用于获取设备信息列表。
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIotDataTypeRequest() (request *DescribeIotDataTypeRequest) {
    request = &DescribeIotDataTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeIotDataType")
    return
}

func NewDescribeIotDataTypeResponse() (response *DescribeIotDataTypeResponse) {
    response = &DescribeIotDataTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeIotDataType）用于查询自定义的物模型数据类型。
func (c *Client) DescribeIotDataType(request *DescribeIotDataTypeRequest) (response *DescribeIotDataTypeResponse, err error) {
    if request == nil {
        request = NewDescribeIotDataTypeRequest()
    }
    response = NewDescribeIotDataTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIotModelRequest() (request *DescribeIotModelRequest) {
    request = &DescribeIotModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeIotModel")
    return
}

func NewDescribeIotModelResponse() (response *DescribeIotModelResponse) {
    response = &DescribeIotModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeIotModel）用于获取物模型定义详情。
func (c *Client) DescribeIotModel(request *DescribeIotModelRequest) (response *DescribeIotModelResponse, err error) {
    if request == nil {
        request = NewDescribeIotModelRequest()
    }
    response = NewDescribeIotModelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIotModelsRequest() (request *DescribeIotModelsRequest) {
    request = &DescribeIotModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeIotModels")
    return
}

func NewDescribeIotModelsResponse() (response *DescribeIotModelsResponse) {
    response = &DescribeIotModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeIotModels）用于列出物模型历史版本列表。
func (c *Client) DescribeIotModels(request *DescribeIotModelsRequest) (response *DescribeIotModelsResponse, err error) {
    if request == nil {
        request = NewDescribeIotModelsRequest()
    }
    response = NewDescribeIotModelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
    request = &DescribeLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeLogs")
    return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
    response = &DescribeLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeLogs）用于查询设备日志列表。
// 设备日志最长保留时长为15天,超期自动清除。
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    response = NewDescribeLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageQueueRequest() (request *DescribeMessageQueueRequest) {
    request = &DescribeMessageQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeMessageQueue")
    return
}

func NewDescribeMessageQueueResponse() (response *DescribeMessageQueueResponse) {
    response = &DescribeMessageQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMessageQueue）用于查询物联网智能视频产品转发消息配置。
func (c *Client) DescribeMessageQueue(request *DescribeMessageQueueRequest) (response *DescribeMessageQueueResponse, err error) {
    if request == nil {
        request = NewDescribeMessageQueueRequest()
    }
    response = NewDescribeMessageQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelDataRetRequest() (request *DescribeModelDataRetRequest) {
    request = &DescribeModelDataRetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeModelDataRet")
    return
}

func NewDescribeModelDataRetResponse() (response *DescribeModelDataRetResponse) {
    response = &DescribeModelDataRetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeModelDataRet）用于根据TaskId获取对设备物模型操作最终响应的结果。
func (c *Client) DescribeModelDataRet(request *DescribeModelDataRetRequest) (response *DescribeModelDataRetResponse, err error) {
    if request == nil {
        request = NewDescribeModelDataRetRequest()
    }
    response = NewDescribeModelDataRetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOtaVersionsRequest() (request *DescribeOtaVersionsRequest) {
    request = &DescribeOtaVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeOtaVersions")
    return
}

func NewDescribeOtaVersionsResponse() (response *DescribeOtaVersionsResponse) {
    response = &DescribeOtaVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeOtaVersions）用于查询固件版本信息列表。
func (c *Client) DescribeOtaVersions(request *DescribeOtaVersionsRequest) (response *DescribeOtaVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeOtaVersionsRequest()
    }
    response = NewDescribeOtaVersionsResponse()
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

// 本接口（DescribeProduct）用于获取单个产品的详细信息。
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

// 本接口（DescribeProducts）用于列出用户账号下的物联网智能视频产品列表。
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePubVersionsRequest() (request *DescribePubVersionsRequest) {
    request = &DescribePubVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribePubVersions")
    return
}

func NewDescribePubVersionsResponse() (response *DescribePubVersionsResponse) {
    response = &DescribePubVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePubVersions）用于获取某一产品发布过的全部固件版本。
func (c *Client) DescribePubVersions(request *DescribePubVersionsRequest) (response *DescribePubVersionsResponse, err error) {
    if request == nil {
        request = NewDescribePubVersionsRequest()
    }
    response = NewDescribePubVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistrationStatusRequest() (request *DescribeRegistrationStatusRequest) {
    request = &DescribeRegistrationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeRegistrationStatus")
    return
}

func NewDescribeRegistrationStatusResponse() (response *DescribeRegistrationStatusResponse) {
    response = &DescribeRegistrationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRegistrationStatus）用于查询终端用户的注册状态。
func (c *Client) DescribeRegistrationStatus(request *DescribeRegistrationStatusRequest) (response *DescribeRegistrationStatusResponse, err error) {
    if request == nil {
        request = NewDescribeRegistrationStatusRequest()
    }
    response = NewDescribeRegistrationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRunLogRequest() (request *DescribeRunLogRequest) {
    request = &DescribeRunLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeRunLog")
    return
}

func NewDescribeRunLogResponse() (response *DescribeRunLogResponse) {
    response = &DescribeRunLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRunLog）用于获取设备运行日志。
func (c *Client) DescribeRunLog(request *DescribeRunLogRequest) (response *DescribeRunLogResponse, err error) {
    if request == nil {
        request = NewDescribeRunLogRequest()
    }
    response = NewDescribeRunLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTraceIdsRequest() (request *DescribeTraceIdsRequest) {
    request = &DescribeTraceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeTraceIds")
    return
}

func NewDescribeTraceIdsResponse() (response *DescribeTraceIdsResponse) {
    response = &DescribeTraceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeTraceIds）用于查询设备日志跟踪白名单。
func (c *Client) DescribeTraceIds(request *DescribeTraceIdsRequest) (response *DescribeTraceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeTraceIdsRequest()
    }
    response = NewDescribeTraceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTraceStatusRequest() (request *DescribeTraceStatusRequest) {
    request = &DescribeTraceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeTraceStatus")
    return
}

func NewDescribeTraceStatusResponse() (response *DescribeTraceStatusResponse) {
    response = &DescribeTraceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeTraceStatus）用于查询指定设备是否在白名单中。
func (c *Client) DescribeTraceStatus(request *DescribeTraceStatusRequest) (response *DescribeTraceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTraceStatusRequest()
    }
    response = NewDescribeTraceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableDeviceRequest() (request *DisableDeviceRequest) {
    request = &DisableDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DisableDevice")
    return
}

func NewDisableDeviceResponse() (response *DisableDeviceResponse) {
    response = &DisableDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DisableDevice）用于禁用设备，可进行批量操作，每次操作最多100台设备。
func (c *Client) DisableDevice(request *DisableDeviceRequest) (response *DisableDeviceResponse, err error) {
    if request == nil {
        request = NewDisableDeviceRequest()
    }
    response = NewDisableDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableDeviceStreamRequest() (request *DisableDeviceStreamRequest) {
    request = &DisableDeviceStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DisableDeviceStream")
    return
}

func NewDisableDeviceStreamResponse() (response *DisableDeviceStreamResponse) {
    response = &DisableDeviceStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DisableDeviceStream）用于停止设备推流，可进行批量操作，每次操作最多100台设备。
func (c *Client) DisableDeviceStream(request *DisableDeviceStreamRequest) (response *DisableDeviceStreamResponse, err error) {
    if request == nil {
        request = NewDisableDeviceStreamRequest()
    }
    response = NewDisableDeviceStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDisableOtaVersionRequest() (request *DisableOtaVersionRequest) {
    request = &DisableOtaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DisableOtaVersion")
    return
}

func NewDisableOtaVersionResponse() (response *DisableOtaVersionResponse) {
    response = &DisableOtaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DisableOtaVersion）用于禁用固件版本。
func (c *Client) DisableOtaVersion(request *DisableOtaVersionRequest) (response *DisableOtaVersionResponse, err error) {
    if request == nil {
        request = NewDisableOtaVersionRequest()
    }
    response = NewDisableOtaVersionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceActionRequest() (request *ModifyDeviceActionRequest) {
    request = &ModifyDeviceActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDeviceAction")
    return
}

func NewModifyDeviceActionResponse() (response *ModifyDeviceActionResponse) {
    response = &ModifyDeviceActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDeviceAction）用于修改设备物模型的行为（Action）。
// 
// 可对ctlVal数据属性进行写入,如:Action.takePhoto.ctlVal,设备在线且成功发送到设备才返回,物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
// 注意:
//   1.若设备当前不在线,会直接返回错误
//   2.若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒
//   3.value的内容必须与实际物模型的定义一致
func (c *Client) ModifyDeviceAction(request *ModifyDeviceActionRequest) (response *ModifyDeviceActionResponse, err error) {
    if request == nil {
        request = NewModifyDeviceActionRequest()
    }
    response = NewModifyDeviceActionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDevicePropertyRequest() (request *ModifyDevicePropertyRequest) {
    request = &ModifyDevicePropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDeviceProperty")
    return
}

func NewModifyDevicePropertyResponse() (response *ModifyDevicePropertyResponse) {
    response = &ModifyDevicePropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDeviceProperty）用于修改设备物模型的属性（ProWritable）。
// 可对setVal数据属性进行写入,如:
// ProWritable.Pos.setVal
// 对于嵌套类型的可写属性，可以仅对其部分数据内容进行写入，如:
// ProWritable.Pos.setVal.x;
// 可写属性云端写入成功即返回;云端向设备端发布属性变更参数;若当前设备不在线,在设备下次上线时会自动更新这些属性参数;
// 物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
func (c *Client) ModifyDeviceProperty(request *ModifyDevicePropertyRequest) (response *ModifyDevicePropertyResponse, err error) {
    if request == nil {
        request = NewModifyDevicePropertyRequest()
    }
    response = NewModifyDevicePropertyResponse()
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

// 本接口（ModifyProduct）用于编辑物联网智能视频产品的相关信息。
func (c *Client) ModifyProduct(request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    if request == nil {
        request = NewModifyProductRequest()
    }
    response = NewModifyProductResponse()
    err = c.Send(request, response)
    return
}

func NewRunDeviceRequest() (request *RunDeviceRequest) {
    request = &RunDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RunDevice")
    return
}

func NewRunDeviceResponse() (response *RunDeviceResponse) {
    response = &RunDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunDevice）用于启用设备，可进行批量操作，每次操作最多100台设备。
func (c *Client) RunDevice(request *RunDeviceRequest) (response *RunDeviceResponse, err error) {
    if request == nil {
        request = NewRunDeviceRequest()
    }
    response = NewRunDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewRunDeviceStreamRequest() (request *RunDeviceStreamRequest) {
    request = &RunDeviceStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RunDeviceStream")
    return
}

func NewRunDeviceStreamResponse() (response *RunDeviceStreamResponse) {
    response = &RunDeviceStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunDeviceStream）用于开启设备推流，可进行批量操作，每次操作最多100台设备。
func (c *Client) RunDeviceStream(request *RunDeviceStreamRequest) (response *RunDeviceStreamResponse, err error) {
    if request == nil {
        request = NewRunDeviceStreamRequest()
    }
    response = NewRunDeviceStreamResponse()
    err = c.Send(request, response)
    return
}

func NewRunIotModelRequest() (request *RunIotModelRequest) {
    request = &RunIotModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RunIotModel")
    return
}

func NewRunIotModelResponse() (response *RunIotModelResponse) {
    response = &RunIotModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunIotModel）用于对定义的物模型进行发布。
func (c *Client) RunIotModel(request *RunIotModelRequest) (response *RunIotModelResponse, err error) {
    if request == nil {
        request = NewRunIotModelRequest()
    }
    response = NewRunIotModelResponse()
    err = c.Send(request, response)
    return
}

func NewRunOtaVersionRequest() (request *RunOtaVersionRequest) {
    request = &RunOtaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RunOtaVersion")
    return
}

func NewRunOtaVersionResponse() (response *RunOtaVersionResponse) {
    response = &RunOtaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunOtaVersion）用于固件版本正式发布。
func (c *Client) RunOtaVersion(request *RunOtaVersionRequest) (response *RunOtaVersionResponse, err error) {
    if request == nil {
        request = NewRunOtaVersionRequest()
    }
    response = NewRunOtaVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRunTestOtaVersionRequest() (request *RunTestOtaVersionRequest) {
    request = &RunTestOtaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "RunTestOtaVersion")
    return
}

func NewRunTestOtaVersionResponse() (response *RunTestOtaVersionResponse) {
    response = &RunTestOtaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunTestOtaVersion）用于固件版本测试发布。
func (c *Client) RunTestOtaVersion(request *RunTestOtaVersionRequest) (response *RunTestOtaVersionResponse, err error) {
    if request == nil {
        request = NewRunTestOtaVersionRequest()
    }
    response = NewRunTestOtaVersionResponse()
    err = c.Send(request, response)
    return
}

func NewSendOnlineMsgRequest() (request *SendOnlineMsgRequest) {
    request = &SendOnlineMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "SendOnlineMsg")
    return
}

func NewSendOnlineMsgResponse() (response *SendOnlineMsgResponse) {
    response = &SendOnlineMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SendOnlineMsg）用于向设备发送在线消息。
// 注意：
// 若设备当前不在线,会直接返回错误;
// 若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒.waitresp非0情况下,会导致本接口阻塞3秒。
func (c *Client) SendOnlineMsg(request *SendOnlineMsgRequest) (response *SendOnlineMsgResponse, err error) {
    if request == nil {
        request = NewSendOnlineMsgRequest()
    }
    response = NewSendOnlineMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSetMessageQueueRequest() (request *SetMessageQueueRequest) {
    request = &SetMessageQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "SetMessageQueue")
    return
}

func NewSetMessageQueueResponse() (response *SetMessageQueueResponse) {
    response = &SetMessageQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SetMessageQueue）用于配置物联网智能视频产品的转发消息队列。
func (c *Client) SetMessageQueue(request *SetMessageQueueRequest) (response *SetMessageQueueResponse, err error) {
    if request == nil {
        request = NewSetMessageQueueRequest()
    }
    response = NewSetMessageQueueResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDeviceRequest() (request *UpgradeDeviceRequest) {
    request = &UpgradeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "UpgradeDevice")
    return
}

func NewUpgradeDeviceResponse() (response *UpgradeDeviceResponse) {
    response = &UpgradeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeDevice）用于对设备进行固件升级。
// 该接口向指定的设备下发固件更新指令,可将固件升级到任意版本(可实现固件降级)。
// 警告:使能UpgradeNow参数存在一定的风险性！建议仅在debug场景下使用!
func (c *Client) UpgradeDevice(request *UpgradeDeviceRequest) (response *UpgradeDeviceResponse, err error) {
    if request == nil {
        request = NewUpgradeDeviceRequest()
    }
    response = NewUpgradeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewUploadOtaVersionRequest() (request *UploadOtaVersionRequest) {
    request = &UploadOtaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "UploadOtaVersion")
    return
}

func NewUploadOtaVersionResponse() (response *UploadOtaVersionResponse) {
    response = &UploadOtaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UploadOtaVersion）接收上传到控制台的固件版本信息。
func (c *Client) UploadOtaVersion(request *UploadOtaVersionRequest) (response *UploadOtaVersionResponse, err error) {
    if request == nil {
        request = NewUploadOtaVersionRequest()
    }
    response = NewUploadOtaVersionResponse()
    err = c.Send(request, response)
    return
}
