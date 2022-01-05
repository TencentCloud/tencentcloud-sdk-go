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

package v20180614

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-14"

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


func NewBatchUpdateFirmwareRequest() (request *BatchUpdateFirmwareRequest) {
    request = &BatchUpdateFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "BatchUpdateFirmware")
    
    
    return
}

func NewBatchUpdateFirmwareResponse() (response *BatchUpdateFirmwareResponse) {
    response = &BatchUpdateFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchUpdateFirmware
// 本接口（BatchUpdateFirmware）用于批量更新设备固件 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) BatchUpdateFirmware(request *BatchUpdateFirmwareRequest) (response *BatchUpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewBatchUpdateFirmwareRequest()
    }
    
    response = NewBatchUpdateFirmwareResponse()
    err = c.Send(request, response)
    return
}

// BatchUpdateFirmware
// 本接口（BatchUpdateFirmware）用于批量更新设备固件 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) BatchUpdateFirmwareWithContext(ctx context.Context, request *BatchUpdateFirmwareRequest) (response *BatchUpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewBatchUpdateFirmwareRequest()
    }
    request.SetContext(ctx)
    
    response = NewBatchUpdateFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewBindDevicesRequest() (request *BindDevicesRequest) {
    request = &BindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "BindDevices")
    
    
    return
}

func NewBindDevicesResponse() (response *BindDevicesResponse) {
    response = &BindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindDevices
// 本接口（BindDevices）用于网关设备批量绑定子设备 
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDDEVICEOVERLIMIT = "FailedOperation.BindDeviceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) BindDevices(request *BindDevicesRequest) (response *BindDevicesResponse, err error) {
    if request == nil {
        request = NewBindDevicesRequest()
    }
    
    response = NewBindDevicesResponse()
    err = c.Send(request, response)
    return
}

// BindDevices
// 本接口（BindDevices）用于网关设备批量绑定子设备 
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDDEVICEOVERLIMIT = "FailedOperation.BindDeviceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) BindDevicesWithContext(ctx context.Context, request *BindDevicesRequest) (response *BindDevicesResponse, err error) {
    if request == nil {
        request = NewBindDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewCancelDeviceFirmwareTaskRequest() (request *CancelDeviceFirmwareTaskRequest) {
    request = &CancelDeviceFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CancelDeviceFirmwareTask")
    
    
    return
}

func NewCancelDeviceFirmwareTaskResponse() (response *CancelDeviceFirmwareTaskResponse) {
    response = &CancelDeviceFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelDeviceFirmwareTask
// 取消设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) CancelDeviceFirmwareTask(request *CancelDeviceFirmwareTaskRequest) (response *CancelDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewCancelDeviceFirmwareTaskRequest()
    }
    
    response = NewCancelDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

// CancelDeviceFirmwareTask
// 取消设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) CancelDeviceFirmwareTaskWithContext(ctx context.Context, request *CancelDeviceFirmwareTaskRequest) (response *CancelDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewCancelDeviceFirmwareTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCancelDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CancelTask")
    
    
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelTask
// 本接口（CancelTask）用于取消一个未被调度的任务。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKALREADYSTARTED = "FailedOperation.AsyncTaskAlreadyStarted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

// CancelTask
// 本接口（CancelTask）用于取消一个未被调度的任务。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKALREADYSTARTED = "FailedOperation.AsyncTaskAlreadyStarted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateDevice")
    
    
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDevice
// 本接口（CreateDevice）用于新建一个物联网通信设备。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

// CreateDevice
// 本接口（CreateDevice）用于新建一个物联网通信设备。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDeviceWithContext(ctx context.Context, request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoraDeviceRequest() (request *CreateLoraDeviceRequest) {
    request = &CreateLoraDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateLoraDevice")
    
    
    return
}

func NewCreateLoraDeviceResponse() (response *CreateLoraDeviceResponse) {
    response = &CreateLoraDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoraDevice
// 创建lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
func (c *Client) CreateLoraDevice(request *CreateLoraDeviceRequest) (response *CreateLoraDeviceResponse, err error) {
    if request == nil {
        request = NewCreateLoraDeviceRequest()
    }
    
    response = NewCreateLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

// CreateLoraDevice
// 创建lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
func (c *Client) CreateLoraDeviceWithContext(ctx context.Context, request *CreateLoraDeviceRequest) (response *CreateLoraDeviceResponse, err error) {
    if request == nil {
        request = NewCreateLoraDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiDeviceRequest() (request *CreateMultiDeviceRequest) {
    request = &CreateMultiDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateMultiDevice")
    
    
    return
}

func NewCreateMultiDeviceResponse() (response *CreateMultiDeviceResponse) {
    response = &CreateMultiDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMultiDevice
// 本接口（CreateMultiDevice）用于批量创建物联云设备。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateMultiDevice(request *CreateMultiDeviceRequest) (response *CreateMultiDeviceResponse, err error) {
    if request == nil {
        request = NewCreateMultiDeviceRequest()
    }
    
    response = NewCreateMultiDeviceResponse()
    err = c.Send(request, response)
    return
}

// CreateMultiDevice
// 本接口（CreateMultiDevice）用于批量创建物联云设备。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateMultiDeviceWithContext(ctx context.Context, request *CreateMultiDeviceRequest) (response *CreateMultiDeviceResponse, err error) {
    if request == nil {
        request = NewCreateMultiDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateMultiDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiDevicesTaskRequest() (request *CreateMultiDevicesTaskRequest) {
    request = &CreateMultiDevicesTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateMultiDevicesTask")
    
    
    return
}

func NewCreateMultiDevicesTaskResponse() (response *CreateMultiDevicesTaskResponse) {
    response = &CreateMultiDevicesTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMultiDevicesTask
// 本接口（CreateMultiDevicesTask）用于创建产品级别的批量创建设备任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateMultiDevicesTask(request *CreateMultiDevicesTaskRequest) (response *CreateMultiDevicesTaskResponse, err error) {
    if request == nil {
        request = NewCreateMultiDevicesTaskRequest()
    }
    
    response = NewCreateMultiDevicesTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateMultiDevicesTask
// 本接口（CreateMultiDevicesTask）用于创建产品级别的批量创建设备任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateMultiDevicesTaskWithContext(ctx context.Context, request *CreateMultiDevicesTaskRequest) (response *CreateMultiDevicesTaskResponse, err error) {
    if request == nil {
        request = NewCreateMultiDevicesTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateMultiDevicesTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProductRequest() (request *CreateProductRequest) {
    request = &CreateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateProduct")
    
    
    return
}

func NewCreateProductResponse() (response *CreateProductResponse) {
    response = &CreateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProduct
// 本接口（CreateProduct）用于创建一个新的物联网通信产品 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  INVALIDPARAMETERVALUE_TIDPRODUCTALREADYEXIST = "InvalidParameterValue.TidProductAlreadyExist"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_THINGMODELNOTEXIST = "ResourceNotFound.ThingModelNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

// CreateProduct
// 本接口（CreateProduct）用于创建一个新的物联网通信产品 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  INVALIDPARAMETERVALUE_TIDPRODUCTALREADYEXIST = "InvalidParameterValue.TidProductAlreadyExist"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_THINGMODELNOTEXIST = "ResourceNotFound.ThingModelNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) CreateProductWithContext(ctx context.Context, request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// 本接口（CreateTask）用于创建一个批量任务。目前此接口可以创建批量更新影子以及批量下发消息的任务 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PENGINGORPROCESSINGTASKSEXCEEDLIMIT = "LimitExceeded.PengingOrProcessingTasksExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateTask
// 本接口（CreateTask）用于创建一个批量任务。目前此接口可以创建批量更新影子以及批量下发消息的任务 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PENGINGORPROCESSINGTASKSEXCEEDLIMIT = "LimitExceeded.PengingOrProcessingTasksExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFileUrlRequest() (request *CreateTaskFileUrlRequest) {
    request = &CreateTaskFileUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateTaskFileUrl")
    
    
    return
}

func NewCreateTaskFileUrlResponse() (response *CreateTaskFileUrlResponse) {
    response = &CreateTaskFileUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTaskFileUrl
// 本接口（CreateTaskFileUrl）用于获取产品级任务文件上传链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTaskFileUrl(request *CreateTaskFileUrlRequest) (response *CreateTaskFileUrlResponse, err error) {
    if request == nil {
        request = NewCreateTaskFileUrlRequest()
    }
    
    response = NewCreateTaskFileUrlResponse()
    err = c.Send(request, response)
    return
}

// CreateTaskFileUrl
// 本接口（CreateTaskFileUrl）用于获取产品级任务文件上传链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTaskFileUrlWithContext(ctx context.Context, request *CreateTaskFileUrlRequest) (response *CreateTaskFileUrlResponse, err error) {
    if request == nil {
        request = NewCreateTaskFileUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTaskFileUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicPolicyRequest() (request *CreateTopicPolicyRequest) {
    request = &CreateTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateTopicPolicy")
    
    
    return
}

func NewCreateTopicPolicyResponse() (response *CreateTopicPolicyResponse) {
    response = &CreateTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopicPolicy
// 本接口（CreateTopicPolicy）用于创建一个Topic 
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNOTBIND = "FailedOperation.ProductNotBind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  LIMITEXCEEDED_TOPICPOLICYEXCEEDLIMIT = "LimitExceeded.TopicPolicyExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateTopicPolicy(request *CreateTopicPolicyRequest) (response *CreateTopicPolicyResponse, err error) {
    if request == nil {
        request = NewCreateTopicPolicyRequest()
    }
    
    response = NewCreateTopicPolicyResponse()
    err = c.Send(request, response)
    return
}

// CreateTopicPolicy
// 本接口（CreateTopicPolicy）用于创建一个Topic 
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNOTBIND = "FailedOperation.ProductNotBind"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  LIMITEXCEEDED_TOPICPOLICYEXCEEDLIMIT = "LimitExceeded.TopicPolicyExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) CreateTopicPolicyWithContext(ctx context.Context, request *CreateTopicPolicyRequest) (response *CreateTopicPolicyResponse, err error) {
    if request == nil {
        request = NewCreateTopicPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTopicPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRuleRequest() (request *CreateTopicRuleRequest) {
    request = &CreateTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateTopicRule")
    
    
    return
}

func NewCreateTopicRuleResponse() (response *CreateTopicRuleResponse) {
    response = &CreateTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopicRule
// 本接口（CreateTopicRule）用于创建一个规则 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateTopicRule(request *CreateTopicRuleRequest) (response *CreateTopicRuleResponse, err error) {
    if request == nil {
        request = NewCreateTopicRuleRequest()
    }
    
    response = NewCreateTopicRuleResponse()
    err = c.Send(request, response)
    return
}

// CreateTopicRule
// 本接口（CreateTopicRule）用于创建一个规则 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateTopicRuleWithContext(ctx context.Context, request *CreateTopicRuleRequest) (response *CreateTopicRuleResponse, err error) {
    if request == nil {
        request = NewCreateTopicRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteDevice")
    
    
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除物联网通信设备。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除物联网通信设备。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDeviceWithContext(ctx context.Context, request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoraDeviceRequest() (request *DeleteLoraDeviceRequest) {
    request = &DeleteLoraDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteLoraDevice")
    
    
    return
}

func NewDeleteLoraDeviceResponse() (response *DeleteLoraDeviceResponse) {
    response = &DeleteLoraDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoraDevice
// 删除lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DeleteLoraDevice(request *DeleteLoraDeviceRequest) (response *DeleteLoraDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteLoraDeviceRequest()
    }
    
    response = NewDeleteLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

// DeleteLoraDevice
// 删除lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DeleteLoraDeviceWithContext(ctx context.Context, request *DeleteLoraDeviceRequest) (response *DeleteLoraDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteLoraDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteProduct")
    
    
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网通信产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网通信产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProductWithContext(ctx context.Context, request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRuleRequest() (request *DeleteTopicRuleRequest) {
    request = &DeleteTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteTopicRule")
    
    
    return
}

func NewDeleteTopicRuleResponse() (response *DeleteTopicRuleResponse) {
    response = &DeleteTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopicRule
// 本接口（DeleteTopicRule）用于删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DeleteTopicRule(request *DeleteTopicRuleRequest) (response *DeleteTopicRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRuleRequest()
    }
    
    response = NewDeleteTopicRuleResponse()
    err = c.Send(request, response)
    return
}

// DeleteTopicRule
// 本接口（DeleteTopicRule）用于删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DeleteTopicRuleWithContext(ctx context.Context, request *DeleteTopicRuleRequest) (response *DeleteTopicRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllDevicesRequest() (request *DescribeAllDevicesRequest) {
    request = &DescribeAllDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeAllDevices")
    
    
    return
}

func NewDescribeAllDevicesResponse() (response *DescribeAllDevicesResponse) {
    response = &DescribeAllDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllDevices
// 查询所有设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DescribeAllDevices(request *DescribeAllDevicesRequest) (response *DescribeAllDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeAllDevicesRequest()
    }
    
    response = NewDescribeAllDevicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAllDevices
// 查询所有设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DescribeAllDevicesWithContext(ctx context.Context, request *DescribeAllDevicesRequest) (response *DescribeAllDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeAllDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAllDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevice")
    
    
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevice
// 本接口（DescribeDevice）用于查看设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

// DescribeDevice
// 本接口（DescribeDevice）用于查看设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceWithContext(ctx context.Context, request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceClientKeyRequest() (request *DescribeDeviceClientKeyRequest) {
    request = &DescribeDeviceClientKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDeviceClientKey")
    
    
    return
}

func NewDescribeDeviceClientKeyResponse() (response *DescribeDeviceClientKeyResponse) {
    response = &DescribeDeviceClientKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceClientKey
// 获取证书认证类型设备的私钥，刚生成或者重置设备后仅可调用一次 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNSUPPORTEDOPERATION_CLIENTCERTALREADYGOT = "UnsupportedOperation.ClientCertAlreadyGot"
//  UNSUPPORTEDOPERATION_WRONGPRODUCTAUTHTYPE = "UnsupportedOperation.WrongProductAuthType"
func (c *Client) DescribeDeviceClientKey(request *DescribeDeviceClientKeyRequest) (response *DescribeDeviceClientKeyResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClientKeyRequest()
    }
    
    response = NewDescribeDeviceClientKeyResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceClientKey
// 获取证书认证类型设备的私钥，刚生成或者重置设备后仅可调用一次 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNSUPPORTEDOPERATION_CLIENTCERTALREADYGOT = "UnsupportedOperation.ClientCertAlreadyGot"
//  UNSUPPORTEDOPERATION_WRONGPRODUCTAUTHTYPE = "UnsupportedOperation.WrongProductAuthType"
func (c *Client) DescribeDeviceClientKeyWithContext(ctx context.Context, request *DescribeDeviceClientKeyRequest) (response *DescribeDeviceClientKeyResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClientKeyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceClientKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceResourceRequest() (request *DescribeDeviceResourceRequest) {
    request = &DescribeDeviceResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDeviceResource")
    
    
    return
}

func NewDescribeDeviceResourceResponse() (response *DescribeDeviceResourceResponse) {
    response = &DescribeDeviceResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceResource
// 本接口（DescribeDeviceResource）用于查询设备资源详情。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResource(request *DescribeDeviceResourceRequest) (response *DescribeDeviceResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceResourceRequest()
    }
    
    response = NewDescribeDeviceResourceResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceResource
// 本接口（DescribeDeviceResource）用于查询设备资源详情。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResourceWithContext(ctx context.Context, request *DescribeDeviceResourceRequest) (response *DescribeDeviceResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceResourcesRequest() (request *DescribeDeviceResourcesRequest) {
    request = &DescribeDeviceResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDeviceResources")
    
    
    return
}

func NewDescribeDeviceResourcesResponse() (response *DescribeDeviceResourcesResponse) {
    response = &DescribeDeviceResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceResources
// 本接口（DescribeDeviceResources）用于查询设备资源列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResources(request *DescribeDeviceResourcesRequest) (response *DescribeDeviceResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceResourcesRequest()
    }
    
    response = NewDescribeDeviceResourcesResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceResources
// 本接口（DescribeDeviceResources）用于查询设备资源列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResourcesWithContext(ctx context.Context, request *DescribeDeviceResourcesRequest) (response *DescribeDeviceResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceResourcesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceShadowRequest() (request *DescribeDeviceShadowRequest) {
    request = &DescribeDeviceShadowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDeviceShadow")
    
    
    return
}

func NewDescribeDeviceShadowResponse() (response *DescribeDeviceShadowResponse) {
    response = &DescribeDeviceShadowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceShadow
// 本接口（DescribeDeviceShadow）用于查询虚拟设备信息。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMINCOMPLETE = "InvalidParameterValue.ParamIncomplete"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) DescribeDeviceShadow(request *DescribeDeviceShadowRequest) (response *DescribeDeviceShadowResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceShadowRequest()
    }
    
    response = NewDescribeDeviceShadowResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceShadow
// 本接口（DescribeDeviceShadow）用于查询虚拟设备信息。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMINCOMPLETE = "InvalidParameterValue.ParamIncomplete"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) DescribeDeviceShadowWithContext(ctx context.Context, request *DescribeDeviceShadowRequest) (response *DescribeDeviceShadowResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceShadowRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceShadowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevices")
    
    
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevices
// 本接口（DescribeDevices）用于查询物联网通信设备的设备列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeDevices
// 本接口（DescribeDevices）用于查询物联网通信设备的设备列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareRequest() (request *DescribeFirmwareRequest) {
    request = &DescribeFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmware")
    
    
    return
}

func NewDescribeFirmwareResponse() (response *DescribeFirmwareResponse) {
    response = &DescribeFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmware
// 查询固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) DescribeFirmware(request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareRequest()
    }
    
    response = NewDescribeFirmwareResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmware
// 查询固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) DescribeFirmwareWithContext(ctx context.Context, request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskRequest() (request *DescribeFirmwareTaskRequest) {
    request = &DescribeFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmwareTask")
    
    
    return
}

func NewDescribeFirmwareTaskResponse() (response *DescribeFirmwareTaskResponse) {
    response = &DescribeFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTask
// 查询固件升级任务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTask(request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    
    response = NewDescribeFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmwareTask
// 查询固件升级任务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskWithContext(ctx context.Context, request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskDevicesRequest() (request *DescribeFirmwareTaskDevicesRequest) {
    request = &DescribeFirmwareTaskDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmwareTaskDevices")
    
    
    return
}

func NewDescribeFirmwareTaskDevicesResponse() (response *DescribeFirmwareTaskDevicesResponse) {
    response = &DescribeFirmwareTaskDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskDevices
// 查询固件升级任务的设备列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskDevices(request *DescribeFirmwareTaskDevicesRequest) (response *DescribeFirmwareTaskDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDevicesRequest()
    }
    
    response = NewDescribeFirmwareTaskDevicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmwareTaskDevices
// 查询固件升级任务的设备列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskDevicesWithContext(ctx context.Context, request *DescribeFirmwareTaskDevicesRequest) (response *DescribeFirmwareTaskDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTaskDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskDistributionRequest() (request *DescribeFirmwareTaskDistributionRequest) {
    request = &DescribeFirmwareTaskDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmwareTaskDistribution")
    
    
    return
}

func NewDescribeFirmwareTaskDistributionResponse() (response *DescribeFirmwareTaskDistributionResponse) {
    response = &DescribeFirmwareTaskDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskDistribution
// 查询固件升级任务状态分布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskDistribution(request *DescribeFirmwareTaskDistributionRequest) (response *DescribeFirmwareTaskDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDistributionRequest()
    }
    
    response = NewDescribeFirmwareTaskDistributionResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmwareTaskDistribution
// 查询固件升级任务状态分布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskDistributionWithContext(ctx context.Context, request *DescribeFirmwareTaskDistributionRequest) (response *DescribeFirmwareTaskDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDistributionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTaskDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTaskStatisticsRequest() (request *DescribeFirmwareTaskStatisticsRequest) {
    request = &DescribeFirmwareTaskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmwareTaskStatistics")
    
    
    return
}

func NewDescribeFirmwareTaskStatisticsResponse() (response *DescribeFirmwareTaskStatisticsResponse) {
    response = &DescribeFirmwareTaskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskStatistics
// 查询固件升级任务统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskStatistics(request *DescribeFirmwareTaskStatisticsRequest) (response *DescribeFirmwareTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskStatisticsRequest()
    }
    
    response = NewDescribeFirmwareTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmwareTaskStatistics
// 查询固件升级任务统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskStatisticsWithContext(ctx context.Context, request *DescribeFirmwareTaskStatisticsRequest) (response *DescribeFirmwareTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirmwareTasksRequest() (request *DescribeFirmwareTasksRequest) {
    request = &DescribeFirmwareTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeFirmwareTasks")
    
    
    return
}

func NewDescribeFirmwareTasksResponse() (response *DescribeFirmwareTasksResponse) {
    response = &DescribeFirmwareTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTasks
// 查询固件升级任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTasks(request *DescribeFirmwareTasksRequest) (response *DescribeFirmwareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTasksRequest()
    }
    
    response = NewDescribeFirmwareTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeFirmwareTasks
// 查询固件升级任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTasksWithContext(ctx context.Context, request *DescribeFirmwareTasksRequest) (response *DescribeFirmwareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoraDeviceRequest() (request *DescribeLoraDeviceRequest) {
    request = &DescribeLoraDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeLoraDevice")
    
    
    return
}

func NewDescribeLoraDeviceResponse() (response *DescribeLoraDeviceResponse) {
    response = &DescribeLoraDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoraDevice
// 获取lora类型设备的详细信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  INVALIDPARAMETERVALUE_QUERYDEVICEFAIL = "InvalidParameterValue.QueryDeviceFail"
//  INVALIDPARAMETERVALUE_QUERYLORADEVICEFAIL = "InvalidParameterValue.QueryLoraDeviceFail"
//  INVALIDPARAMETERVALUE_QUERYNOTLORADEVICEFAIL = "InvalidParameterValue.QueryNotLoraDeviceFail"
//  INVALIDPARAMETERVALUE_REGISTERLORAINFOERROR = "InvalidParameterValue.RegisterLoraInfoError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeLoraDevice(request *DescribeLoraDeviceRequest) (response *DescribeLoraDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeLoraDeviceRequest()
    }
    
    response = NewDescribeLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

// DescribeLoraDevice
// 获取lora类型设备的详细信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  INVALIDPARAMETERVALUE_QUERYDEVICEFAIL = "InvalidParameterValue.QueryDeviceFail"
//  INVALIDPARAMETERVALUE_QUERYLORADEVICEFAIL = "InvalidParameterValue.QueryLoraDeviceFail"
//  INVALIDPARAMETERVALUE_QUERYNOTLORADEVICEFAIL = "InvalidParameterValue.QueryNotLoraDeviceFail"
//  INVALIDPARAMETERVALUE_REGISTERLORAINFOERROR = "InvalidParameterValue.RegisterLoraInfoError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeLoraDeviceWithContext(ctx context.Context, request *DescribeLoraDeviceRequest) (response *DescribeLoraDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeLoraDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLoraDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiDevTaskRequest() (request *DescribeMultiDevTaskRequest) {
    request = &DescribeMultiDevTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeMultiDevTask")
    
    
    return
}

func NewDescribeMultiDevTaskResponse() (response *DescribeMultiDevTaskResponse) {
    response = &DescribeMultiDevTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMultiDevTask
// 本接口（DescribeMultiDevTask）用于查询批量创建设备任务的执行状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKIDNOTMATCH = "FailedOperation.TaskIDNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TASKIDNOTMATCH = "InvalidParameterValue.TaskIDNotMatch"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeMultiDevTask(request *DescribeMultiDevTaskRequest) (response *DescribeMultiDevTaskResponse, err error) {
    if request == nil {
        request = NewDescribeMultiDevTaskRequest()
    }
    
    response = NewDescribeMultiDevTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeMultiDevTask
// 本接口（DescribeMultiDevTask）用于查询批量创建设备任务的执行状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKIDNOTMATCH = "FailedOperation.TaskIDNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TASKIDNOTMATCH = "InvalidParameterValue.TaskIDNotMatch"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeMultiDevTaskWithContext(ctx context.Context, request *DescribeMultiDevTaskRequest) (response *DescribeMultiDevTaskResponse, err error) {
    if request == nil {
        request = NewDescribeMultiDevTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMultiDevTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiDevicesRequest() (request *DescribeMultiDevicesRequest) {
    request = &DescribeMultiDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeMultiDevices")
    
    
    return
}

func NewDescribeMultiDevicesResponse() (response *DescribeMultiDevicesResponse) {
    response = &DescribeMultiDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMultiDevices
// 本接口（DescribeMultiDevices）用于查询批量创建设备的执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKIDNOTMATCH = "FailedOperation.TaskIDNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
//  UNAUTHORIZEDOPERATION_CREATEMULTIDEVICETASKNOTFINISHED = "UnauthorizedOperation.CreateMultiDeviceTaskNotFinished"
func (c *Client) DescribeMultiDevices(request *DescribeMultiDevicesRequest) (response *DescribeMultiDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeMultiDevicesRequest()
    }
    
    response = NewDescribeMultiDevicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeMultiDevices
// 本接口（DescribeMultiDevices）用于查询批量创建设备的执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKIDNOTMATCH = "FailedOperation.TaskIDNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
//  UNAUTHORIZEDOPERATION_CREATEMULTIDEVICETASKNOTFINISHED = "UnauthorizedOperation.CreateMultiDeviceTaskNotFinished"
func (c *Client) DescribeMultiDevicesWithContext(ctx context.Context, request *DescribeMultiDevicesRequest) (response *DescribeMultiDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeMultiDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMultiDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductResourceRequest() (request *DescribeProductResourceRequest) {
    request = &DescribeProductResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductResource")
    
    
    return
}

func NewDescribeProductResourceResponse() (response *DescribeProductResourceResponse) {
    response = &DescribeProductResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductResource
// 本接口（DescribeProductResource）用于查询产品资源详情。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeProductResource(request *DescribeProductResourceRequest) (response *DescribeProductResourceResponse, err error) {
    if request == nil {
        request = NewDescribeProductResourceRequest()
    }
    
    response = NewDescribeProductResourceResponse()
    err = c.Send(request, response)
    return
}

// DescribeProductResource
// 本接口（DescribeProductResource）用于查询产品资源详情。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeProductResourceWithContext(ctx context.Context, request *DescribeProductResourceRequest) (response *DescribeProductResourceResponse, err error) {
    if request == nil {
        request = NewDescribeProductResourceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductResourcesRequest() (request *DescribeProductResourcesRequest) {
    request = &DescribeProductResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductResources")
    
    
    return
}

func NewDescribeProductResourcesResponse() (response *DescribeProductResourcesResponse) {
    response = &DescribeProductResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductResources
// 本接口（DescribeProductResources）用于查询产品资源列表。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTRESOURCEDUPLICATE = "FailedOperation.ProductResourceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeProductResources(request *DescribeProductResourcesRequest) (response *DescribeProductResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeProductResourcesRequest()
    }
    
    response = NewDescribeProductResourcesResponse()
    err = c.Send(request, response)
    return
}

// DescribeProductResources
// 本接口（DescribeProductResources）用于查询产品资源列表。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTRESOURCEDUPLICATE = "FailedOperation.ProductResourceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeProductResourcesWithContext(ctx context.Context, request *DescribeProductResourcesRequest) (response *DescribeProductResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeProductResourcesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductTaskRequest() (request *DescribeProductTaskRequest) {
    request = &DescribeProductTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductTask")
    
    
    return
}

func NewDescribeProductTaskResponse() (response *DescribeProductTaskResponse) {
    response = &DescribeProductTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductTask
// 本接口（DescribeProductTask）用于查看产品级别的任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeProductTask(request *DescribeProductTaskRequest) (response *DescribeProductTaskResponse, err error) {
    if request == nil {
        request = NewDescribeProductTaskRequest()
    }
    
    response = NewDescribeProductTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeProductTask
// 本接口（DescribeProductTask）用于查看产品级别的任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeProductTaskWithContext(ctx context.Context, request *DescribeProductTaskRequest) (response *DescribeProductTaskResponse, err error) {
    if request == nil {
        request = NewDescribeProductTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductTasksRequest() (request *DescribeProductTasksRequest) {
    request = &DescribeProductTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductTasks")
    
    
    return
}

func NewDescribeProductTasksResponse() (response *DescribeProductTasksResponse) {
    response = &DescribeProductTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductTasks
// 本接口（DescribeProductTasks）用于查看产品级别的任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProductTasks(request *DescribeProductTasksRequest) (response *DescribeProductTasksResponse, err error) {
    if request == nil {
        request = NewDescribeProductTasksRequest()
    }
    
    response = NewDescribeProductTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeProductTasks
// 本接口（DescribeProductTasks）用于查看产品级别的任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProductTasksWithContext(ctx context.Context, request *DescribeProductTasksRequest) (response *DescribeProductTasksResponse, err error) {
    if request == nil {
        request = NewDescribeProductTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
    request = &DescribeProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProducts")
    
    
    return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
    response = &DescribeProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProducts
// 本接口（DescribeProducts）用于列出产品列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProducts
// 本接口（DescribeProducts）用于列出产品列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeProductsWithContext(ctx context.Context, request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushResourceTaskStatisticsRequest() (request *DescribePushResourceTaskStatisticsRequest) {
    request = &DescribePushResourceTaskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePushResourceTaskStatistics")
    
    
    return
}

func NewDescribePushResourceTaskStatisticsResponse() (response *DescribePushResourceTaskStatisticsResponse) {
    response = &DescribePushResourceTaskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePushResourceTaskStatistics
// 查询推送资源任务统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribePushResourceTaskStatistics(request *DescribePushResourceTaskStatisticsRequest) (response *DescribePushResourceTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribePushResourceTaskStatisticsRequest()
    }
    
    response = NewDescribePushResourceTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribePushResourceTaskStatistics
// 查询推送资源任务统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribePushResourceTaskStatisticsWithContext(ctx context.Context, request *DescribePushResourceTaskStatisticsRequest) (response *DescribePushResourceTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribePushResourceTaskStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePushResourceTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTasksRequest() (request *DescribeResourceTasksRequest) {
    request = &DescribeResourceTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeResourceTasks")
    
    
    return
}

func NewDescribeResourceTasksResponse() (response *DescribeResourceTasksResponse) {
    response = &DescribeResourceTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceTasks
// 查询资源推送任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeResourceTasks(request *DescribeResourceTasksRequest) (response *DescribeResourceTasksResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTasksRequest()
    }
    
    response = NewDescribeResourceTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeResourceTasks
// 查询资源推送任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeResourceTasksWithContext(ctx context.Context, request *DescribeResourceTasksRequest) (response *DescribeResourceTasksResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeResourceTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 本接口（DescribeTask）用于查询一个已创建任务的详情，任务保留一个月 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeTask
// 本接口（DescribeTask）用于查询一个已创建任务的详情，任务保留一个月 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 本接口（DescribeTasks）用于查询已创建的任务列表，任务保留一个月 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeTasks
// 本接口（DescribeTasks）用于查询已创建的任务列表，任务保留一个月 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTopicRuleRequest() (request *DisableTopicRuleRequest) {
    request = &DisableTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DisableTopicRule")
    
    
    return
}

func NewDisableTopicRuleResponse() (response *DisableTopicRuleResponse) {
    response = &DisableTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableTopicRule
// 本接口（DisableTopicRule）用于禁用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYDISABLED = "FailedOperation.RuleAlreadyDisabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DisableTopicRule(request *DisableTopicRuleRequest) (response *DisableTopicRuleResponse, err error) {
    if request == nil {
        request = NewDisableTopicRuleRequest()
    }
    
    response = NewDisableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

// DisableTopicRule
// 本接口（DisableTopicRule）用于禁用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYDISABLED = "FailedOperation.RuleAlreadyDisabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) DisableTopicRuleWithContext(ctx context.Context, request *DisableTopicRuleRequest) (response *DisableTopicRuleResponse, err error) {
    if request == nil {
        request = NewDisableTopicRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEditFirmwareRequest() (request *EditFirmwareRequest) {
    request = &EditFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "EditFirmware")
    
    
    return
}

func NewEditFirmwareResponse() (response *EditFirmwareResponse) {
    response = &EditFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EditFirmware
// 编辑固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
func (c *Client) EditFirmware(request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    if request == nil {
        request = NewEditFirmwareRequest()
    }
    
    response = NewEditFirmwareResponse()
    err = c.Send(request, response)
    return
}

// EditFirmware
// 编辑固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
func (c *Client) EditFirmwareWithContext(ctx context.Context, request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    if request == nil {
        request = NewEditFirmwareRequest()
    }
    request.SetContext(ctx)
    
    response = NewEditFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTopicRuleRequest() (request *EnableTopicRuleRequest) {
    request = &EnableTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "EnableTopicRule")
    
    
    return
}

func NewEnableTopicRuleResponse() (response *EnableTopicRuleResponse) {
    response = &EnableTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableTopicRule
// 本接口（EnableTopicRule）用于启用规则 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATIONOFFUNCTIONITEM = "FailedOperation.DuplicationOfFunctionItem"
//  FAILEDOPERATION_FUNCTIONFILENOTEXIST = "FailedOperation.FunctionFileNotExist"
//  FAILEDOPERATION_PROXYIPISNOTENOUGH = "FailedOperation.ProxyIPIsNotEnough"
//  FAILEDOPERATION_RULEALREADYENABLED = "FailedOperation.RuleAlreadyEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) EnableTopicRule(request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    if request == nil {
        request = NewEnableTopicRuleRequest()
    }
    
    response = NewEnableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

// EnableTopicRule
// 本接口（EnableTopicRule）用于启用规则 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATIONOFFUNCTIONITEM = "FailedOperation.DuplicationOfFunctionItem"
//  FAILEDOPERATION_FUNCTIONFILENOTEXIST = "FailedOperation.FunctionFileNotExist"
//  FAILEDOPERATION_PROXYIPISNOTENOUGH = "FailedOperation.ProxyIPIsNotEnough"
//  FAILEDOPERATION_RULEALREADYENABLED = "FailedOperation.RuleAlreadyEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) EnableTopicRuleWithContext(ctx context.Context, request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    if request == nil {
        request = NewEnableTopicRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetCOSURLRequest() (request *GetCOSURLRequest) {
    request = &GetCOSURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "GetCOSURL")
    
    
    return
}

func NewGetCOSURLResponse() (response *GetCOSURLResponse) {
    response = &GetCOSURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCOSURL
// 本接口（GetCOSURL）用于获取固件存储在COS的URL 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetCOSURL(request *GetCOSURLRequest) (response *GetCOSURLResponse, err error) {
    if request == nil {
        request = NewGetCOSURLRequest()
    }
    
    response = NewGetCOSURLResponse()
    err = c.Send(request, response)
    return
}

// GetCOSURL
// 本接口（GetCOSURL）用于获取固件存储在COS的URL 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetCOSURLWithContext(ctx context.Context, request *GetCOSURLRequest) (response *GetCOSURLResponse, err error) {
    if request == nil {
        request = NewGetCOSURLRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetCOSURLResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserResourceInfoRequest() (request *GetUserResourceInfoRequest) {
    request = &GetUserResourceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "GetUserResourceInfo")
    
    
    return
}

func NewGetUserResourceInfoResponse() (response *GetUserResourceInfoResponse) {
    response = &GetUserResourceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUserResourceInfo
// 本接口（GetUserResourceInfo）用于查询用户资源使用信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) GetUserResourceInfo(request *GetUserResourceInfoRequest) (response *GetUserResourceInfoResponse, err error) {
    if request == nil {
        request = NewGetUserResourceInfoRequest()
    }
    
    response = NewGetUserResourceInfoResponse()
    err = c.Send(request, response)
    return
}

// GetUserResourceInfo
// 本接口（GetUserResourceInfo）用于查询用户资源使用信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) GetUserResourceInfoWithContext(ctx context.Context, request *GetUserResourceInfoRequest) (response *GetUserResourceInfoResponse, err error) {
    if request == nil {
        request = NewGetUserResourceInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetUserResourceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewPublishAsDeviceRequest() (request *PublishAsDeviceRequest) {
    request = &PublishAsDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishAsDevice")
    
    
    return
}

func NewPublishAsDeviceResponse() (response *PublishAsDeviceResponse) {
    response = &PublishAsDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishAsDevice
// 模拟lora类型的设备端向服务器端发送消息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) PublishAsDevice(request *PublishAsDeviceRequest) (response *PublishAsDeviceResponse, err error) {
    if request == nil {
        request = NewPublishAsDeviceRequest()
    }
    
    response = NewPublishAsDeviceResponse()
    err = c.Send(request, response)
    return
}

// PublishAsDevice
// 模拟lora类型的设备端向服务器端发送消息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) PublishAsDeviceWithContext(ctx context.Context, request *PublishAsDeviceRequest) (response *PublishAsDeviceResponse, err error) {
    if request == nil {
        request = NewPublishAsDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishAsDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewPublishBroadcastMessageRequest() (request *PublishBroadcastMessageRequest) {
    request = &PublishBroadcastMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishBroadcastMessage")
    
    
    return
}

func NewPublishBroadcastMessageResponse() (response *PublishBroadcastMessageResponse) {
    response = &PublishBroadcastMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BROADCASTTASKISRUNNING = "FailedOperation.BroadcastTaskIsRunning"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessage(request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    if request == nil {
        request = NewPublishBroadcastMessageRequest()
    }
    
    response = NewPublishBroadcastMessageResponse()
    err = c.Send(request, response)
    return
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BROADCASTTASKISRUNNING = "FailedOperation.BroadcastTaskIsRunning"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessageWithContext(ctx context.Context, request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    if request == nil {
        request = NewPublishBroadcastMessageRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishBroadcastMessageResponse()
    err = c.Send(request, response)
    return
}

func NewPublishMessageRequest() (request *PublishMessageRequest) {
    request = &PublishMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishMessage")
    
    
    return
}

func NewPublishMessageResponse() (response *PublishMessageResponse) {
    response = &PublishMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishMessage
// 本接口（PublishMessage）用于向某个主题发消息。 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_INVALIDMSGLEN = "FailedOperation.InvalidMsgLen"
//  FAILEDOPERATION_INVALIDTOPICNAME = "FailedOperation.InvalidTopicName"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"
//  LIMITEXCEEDED_OFFLINEMESSAGEEXCEEDLIMIT = "LimitExceeded.OfflineMessageExceedLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    if request == nil {
        request = NewPublishMessageRequest()
    }
    
    response = NewPublishMessageResponse()
    err = c.Send(request, response)
    return
}

// PublishMessage
// 本接口（PublishMessage）用于向某个主题发消息。 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_INVALIDMSGLEN = "FailedOperation.InvalidMsgLen"
//  FAILEDOPERATION_INVALIDTOPICNAME = "FailedOperation.InvalidTopicName"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"
//  LIMITEXCEEDED_OFFLINEMESSAGEEXCEEDLIMIT = "LimitExceeded.OfflineMessageExceedLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) PublishMessageWithContext(ctx context.Context, request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    if request == nil {
        request = NewPublishMessageRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishMessageResponse()
    err = c.Send(request, response)
    return
}

func NewPublishRRPCMessageRequest() (request *PublishRRPCMessageRequest) {
    request = &PublishRRPCMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishRRPCMessage")
    
    
    return
}

func NewPublishRRPCMessageResponse() (response *PublishRRPCMessageResponse) {
    response = &PublishRRPCMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishRRPCMessage
// 发布RRPC消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_RRPCTIMEOUT = "FailedOperation.RRPCTimeout"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  LIMITEXCEEDED_OFFLINEMESSAGEEXCEEDLIMIT = "LimitExceeded.OfflineMessageExceedLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) PublishRRPCMessage(request *PublishRRPCMessageRequest) (response *PublishRRPCMessageResponse, err error) {
    if request == nil {
        request = NewPublishRRPCMessageRequest()
    }
    
    response = NewPublishRRPCMessageResponse()
    err = c.Send(request, response)
    return
}

// PublishRRPCMessage
// 发布RRPC消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_RRPCTIMEOUT = "FailedOperation.RRPCTimeout"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  LIMITEXCEEDED_OFFLINEMESSAGEEXCEEDLIMIT = "LimitExceeded.OfflineMessageExceedLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) PublishRRPCMessageWithContext(ctx context.Context, request *PublishRRPCMessageRequest) (response *PublishRRPCMessageResponse, err error) {
    if request == nil {
        request = NewPublishRRPCMessageRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishRRPCMessageResponse()
    err = c.Send(request, response)
    return
}

func NewPublishToDeviceRequest() (request *PublishToDeviceRequest) {
    request = &PublishToDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishToDevice")
    
    
    return
}

func NewPublishToDeviceResponse() (response *PublishToDeviceResponse) {
    response = &PublishToDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishToDevice
// 服务器端下发消息给lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) PublishToDevice(request *PublishToDeviceRequest) (response *PublishToDeviceResponse, err error) {
    if request == nil {
        request = NewPublishToDeviceRequest()
    }
    
    response = NewPublishToDeviceResponse()
    err = c.Send(request, response)
    return
}

// PublishToDevice
// 服务器端下发消息给lora类型的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) PublishToDeviceWithContext(ctx context.Context, request *PublishToDeviceRequest) (response *PublishToDeviceResponse, err error) {
    if request == nil {
        request = NewPublishToDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishToDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceTopicRuleRequest() (request *ReplaceTopicRuleRequest) {
    request = &ReplaceTopicRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "ReplaceTopicRule")
    
    
    return
}

func NewReplaceTopicRuleResponse() (response *ReplaceTopicRuleResponse) {
    response = &ReplaceTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceTopicRule
// 本接口（ReplaceTopicRule）用于修改替换规则 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReplaceTopicRule(request *ReplaceTopicRuleRequest) (response *ReplaceTopicRuleResponse, err error) {
    if request == nil {
        request = NewReplaceTopicRuleRequest()
    }
    
    response = NewReplaceTopicRuleResponse()
    err = c.Send(request, response)
    return
}

// ReplaceTopicRule
// 本接口（ReplaceTopicRule）用于修改替换规则 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReplaceTopicRuleWithContext(ctx context.Context, request *ReplaceTopicRuleRequest) (response *ReplaceTopicRuleResponse, err error) {
    if request == nil {
        request = NewReplaceTopicRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewReplaceTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceStateRequest() (request *ResetDeviceStateRequest) {
    request = &ResetDeviceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "ResetDeviceState")
    
    
    return
}

func NewResetDeviceStateResponse() (response *ResetDeviceStateResponse) {
    response = &ResetDeviceStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetDeviceState
// 重置设备的连接状态 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) ResetDeviceState(request *ResetDeviceStateRequest) (response *ResetDeviceStateResponse, err error) {
    if request == nil {
        request = NewResetDeviceStateRequest()
    }
    
    response = NewResetDeviceStateResponse()
    err = c.Send(request, response)
    return
}

// ResetDeviceState
// 重置设备的连接状态 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) ResetDeviceStateWithContext(ctx context.Context, request *ResetDeviceStateRequest) (response *ResetDeviceStateResponse, err error) {
    if request == nil {
        request = NewResetDeviceStateRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetDeviceStateResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDeviceFirmwareTaskRequest() (request *RetryDeviceFirmwareTaskRequest) {
    request = &RetryDeviceFirmwareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "RetryDeviceFirmwareTask")
    
    
    return
}

func NewRetryDeviceFirmwareTaskResponse() (response *RetryDeviceFirmwareTaskResponse) {
    response = &RetryDeviceFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RetryDeviceFirmwareTask
// 重试设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) RetryDeviceFirmwareTask(request *RetryDeviceFirmwareTaskRequest) (response *RetryDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewRetryDeviceFirmwareTaskRequest()
    }
    
    response = NewRetryDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

// RetryDeviceFirmwareTask
// 重试设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) RetryDeviceFirmwareTaskWithContext(ctx context.Context, request *RetryDeviceFirmwareTaskRequest) (response *RetryDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewRetryDeviceFirmwareTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewRetryDeviceFirmwareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSetProductsForbiddenStatusRequest() (request *SetProductsForbiddenStatusRequest) {
    request = &SetProductsForbiddenStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "SetProductsForbiddenStatus")
    
    
    return
}

func NewSetProductsForbiddenStatusResponse() (response *SetProductsForbiddenStatusResponse) {
    response = &SetProductsForbiddenStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetProductsForbiddenStatus
// 批量设置产品禁用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) SetProductsForbiddenStatus(request *SetProductsForbiddenStatusRequest) (response *SetProductsForbiddenStatusResponse, err error) {
    if request == nil {
        request = NewSetProductsForbiddenStatusRequest()
    }
    
    response = NewSetProductsForbiddenStatusResponse()
    err = c.Send(request, response)
    return
}

// SetProductsForbiddenStatus
// 批量设置产品禁用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) SetProductsForbiddenStatusWithContext(ctx context.Context, request *SetProductsForbiddenStatusRequest) (response *SetProductsForbiddenStatusResponse, err error) {
    if request == nil {
        request = NewSetProductsForbiddenStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewSetProductsForbiddenStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDevicesRequest() (request *UnbindDevicesRequest) {
    request = &UnbindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UnbindDevices")
    
    
    return
}

func NewUnbindDevicesResponse() (response *UnbindDevicesResponse) {
    response = &UnbindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindDevices
// 本接口（UnbindDevices）用于网关设备批量解绑子设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UnbindDevices(request *UnbindDevicesRequest) (response *UnbindDevicesResponse, err error) {
    if request == nil {
        request = NewUnbindDevicesRequest()
    }
    
    response = NewUnbindDevicesResponse()
    err = c.Send(request, response)
    return
}

// UnbindDevices
// 本接口（UnbindDevices）用于网关设备批量解绑子设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UnbindDevicesWithContext(ctx context.Context, request *UnbindDevicesRequest) (response *UnbindDevicesResponse, err error) {
    if request == nil {
        request = NewUnbindDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewUnbindDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceAvailableStateRequest() (request *UpdateDeviceAvailableStateRequest) {
    request = &UpdateDeviceAvailableStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDeviceAvailableState")
    
    
    return
}

func NewUpdateDeviceAvailableStateResponse() (response *UpdateDeviceAvailableStateResponse) {
    response = &UpdateDeviceAvailableStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceAvailableState
// 启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDeviceAvailableState(request *UpdateDeviceAvailableStateRequest) (response *UpdateDeviceAvailableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceAvailableStateRequest()
    }
    
    response = NewUpdateDeviceAvailableStateResponse()
    err = c.Send(request, response)
    return
}

// UpdateDeviceAvailableState
// 启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDeviceAvailableStateWithContext(ctx context.Context, request *UpdateDeviceAvailableStateRequest) (response *UpdateDeviceAvailableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceAvailableStateRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateDeviceAvailableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceShadowRequest() (request *UpdateDeviceShadowRequest) {
    request = &UpdateDeviceShadowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDeviceShadow")
    
    
    return
}

func NewUpdateDeviceShadowResponse() (response *UpdateDeviceShadowResponse) {
    response = &UpdateDeviceShadowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceShadow
// 本接口（UpdateDeviceShadow）用于更新虚拟设备信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEVERSIONNOTMATCH = "FailedOperation.UpdateVersionNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDJSON = "InvalidParameterValue.InvalidJSON"
//  INVALIDPARAMETERVALUE_JSONHASINVALIDNODE = "InvalidParameterValue.JSONHasInvalidNode"
//  INVALIDPARAMETERVALUE_JSONSIZEEXCEEDLIMIT = "InvalidParameterValue.JSONSizeExceedLimit"
//  INVALIDPARAMETERVALUE_PREFIXINVALID = "InvalidParameterValue.PrefixInvalid"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) UpdateDeviceShadow(request *UpdateDeviceShadowRequest) (response *UpdateDeviceShadowResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceShadowRequest()
    }
    
    response = NewUpdateDeviceShadowResponse()
    err = c.Send(request, response)
    return
}

// UpdateDeviceShadow
// 本接口（UpdateDeviceShadow）用于更新虚拟设备信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEVERSIONNOTMATCH = "FailedOperation.UpdateVersionNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDJSON = "InvalidParameterValue.InvalidJSON"
//  INVALIDPARAMETERVALUE_JSONHASINVALIDNODE = "InvalidParameterValue.JSONHasInvalidNode"
//  INVALIDPARAMETERVALUE_JSONSIZEEXCEEDLIMIT = "InvalidParameterValue.JSONSizeExceedLimit"
//  INVALIDPARAMETERVALUE_PREFIXINVALID = "InvalidParameterValue.PrefixInvalid"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) UpdateDeviceShadowWithContext(ctx context.Context, request *UpdateDeviceShadowRequest) (response *UpdateDeviceShadowResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceShadowRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateDeviceShadowResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicesEnableStateRequest() (request *UpdateDevicesEnableStateRequest) {
    request = &UpdateDevicesEnableStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDevicesEnableState")
    
    
    return
}

func NewUpdateDevicesEnableStateResponse() (response *UpdateDevicesEnableStateResponse) {
    response = &UpdateDevicesEnableStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicesEnableState
// 批量启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDevicesEnableState(request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
    
    response = NewUpdateDevicesEnableStateResponse()
    err = c.Send(request, response)
    return
}

// UpdateDevicesEnableState
// 批量启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDevicesEnableStateWithContext(ctx context.Context, request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateDevicesEnableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTopicPolicyRequest() (request *UpdateTopicPolicyRequest) {
    request = &UpdateTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateTopicPolicy")
    
    
    return
}

func NewUpdateTopicPolicyResponse() (response *UpdateTopicPolicyResponse) {
    response = &UpdateTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateTopicPolicy
// 本接口（UpdateTopicPolicy）用于更新Topic信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) UpdateTopicPolicy(request *UpdateTopicPolicyRequest) (response *UpdateTopicPolicyResponse, err error) {
    if request == nil {
        request = NewUpdateTopicPolicyRequest()
    }
    
    response = NewUpdateTopicPolicyResponse()
    err = c.Send(request, response)
    return
}

// UpdateTopicPolicy
// 本接口（UpdateTopicPolicy）用于更新Topic信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) UpdateTopicPolicyWithContext(ctx context.Context, request *UpdateTopicPolicyRequest) (response *UpdateTopicPolicyResponse, err error) {
    if request == nil {
        request = NewUpdateTopicPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateTopicPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFirmwareRequest() (request *UploadFirmwareRequest) {
    request = &UploadFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UploadFirmware")
    
    
    return
}

func NewUploadFirmwareResponse() (response *UploadFirmwareResponse) {
    response = &UploadFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadFirmware
// 本接口（UploadFirmware）用于上传设备固件信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}

// UploadFirmware
// 本接口（UploadFirmware）用于上传设备固件信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
func (c *Client) UploadFirmwareWithContext(ctx context.Context, request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    request.SetContext(ctx)
    
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}
