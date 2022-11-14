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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyAIModelRequest() (request *ApplyAIModelRequest) {
    request = &ApplyAIModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ApplyAIModel")
    
    
    return
}

func NewApplyAIModelResponse() (response *ApplyAIModelResponse) {
    response = &ApplyAIModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyAIModel
// 申请AI模型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyAIModel(request *ApplyAIModelRequest) (response *ApplyAIModelResponse, err error) {
    return c.ApplyAIModelWithContext(context.Background(), request)
}

// ApplyAIModel
// 申请AI模型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyAIModelWithContext(ctx context.Context, request *ApplyAIModelRequest) (response *ApplyAIModelResponse, err error) {
    if request == nil {
        request = NewApplyAIModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyAIModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyAIModelResponse()
    err = c.Send(request, response)
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

// BatchUpdateFirmware
// 本接口（BatchUpdateFirmware）用于批量更新设备固件 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) BatchUpdateFirmware(request *BatchUpdateFirmwareRequest) (response *BatchUpdateFirmwareResponse, err error) {
    return c.BatchUpdateFirmwareWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchUpdateFirmware require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchUpdateFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewBindCloudStorageUserRequest() (request *BindCloudStorageUserRequest) {
    request = &BindCloudStorageUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "BindCloudStorageUser")
    
    
    return
}

func NewBindCloudStorageUserResponse() (response *BindCloudStorageUserResponse) {
    response = &BindCloudStorageUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindCloudStorageUser
// 绑定云存用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindCloudStorageUser(request *BindCloudStorageUserRequest) (response *BindCloudStorageUserResponse, err error) {
    return c.BindCloudStorageUserWithContext(context.Background(), request)
}

// BindCloudStorageUser
// 绑定云存用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindCloudStorageUserWithContext(ctx context.Context, request *BindCloudStorageUserRequest) (response *BindCloudStorageUserResponse, err error) {
    if request == nil {
        request = NewBindCloudStorageUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindCloudStorageUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindCloudStorageUserResponse()
    err = c.Send(request, response)
    return
}

func NewCancelAIModelApplicationRequest() (request *CancelAIModelApplicationRequest) {
    request = &CancelAIModelApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CancelAIModelApplication")
    
    
    return
}

func NewCancelAIModelApplicationResponse() (response *CancelAIModelApplicationResponse) {
    response = &CancelAIModelApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelAIModelApplication
// 取消AI模型申请
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAIModelApplication(request *CancelAIModelApplicationRequest) (response *CancelAIModelApplicationResponse, err error) {
    return c.CancelAIModelApplicationWithContext(context.Background(), request)
}

// CancelAIModelApplication
// 取消AI模型申请
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAIModelApplicationWithContext(ctx context.Context, request *CancelAIModelApplicationRequest) (response *CancelAIModelApplicationResponse, err error) {
    if request == nil {
        request = NewCancelAIModelApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAIModelApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAIModelApplicationResponse()
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

// CancelDeviceFirmwareTask
// 本接口用于取消设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) CancelDeviceFirmwareTask(request *CancelDeviceFirmwareTaskRequest) (response *CancelDeviceFirmwareTaskResponse, err error) {
    return c.CancelDeviceFirmwareTaskWithContext(context.Background(), request)
}

// CancelDeviceFirmwareTask
// 本接口用于取消设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) CancelDeviceFirmwareTaskWithContext(ctx context.Context, request *CancelDeviceFirmwareTaskRequest) (response *CancelDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewCancelDeviceFirmwareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDeviceFirmwareTask require credential")
    }

    request.SetContext(ctx)
    
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

// CheckForwardAuth
// 判断是否开启的转发的权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckForwardAuth(request *CheckForwardAuthRequest) (response *CheckForwardAuthResponse, err error) {
    return c.CheckForwardAuthWithContext(context.Background(), request)
}

// CheckForwardAuth
// 判断是否开启的转发的权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckForwardAuthWithContext(ctx context.Context, request *CheckForwardAuthRequest) (response *CheckForwardAuthResponse, err error) {
    if request == nil {
        request = NewCheckForwardAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckForwardAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckForwardAuthResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceDataRequest() (request *ControlDeviceDataRequest) {
    request = &ControlDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ControlDeviceData")
    
    
    return
}

func NewControlDeviceDataResponse() (response *ControlDeviceDataResponse) {
    response = &ControlDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlDeviceData
// 根据设备产品ID、设备名称，设置控制设备的属性数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ControlDeviceData(request *ControlDeviceDataRequest) (response *ControlDeviceDataResponse, err error) {
    return c.ControlDeviceDataWithContext(context.Background(), request)
}

// ControlDeviceData
// 根据设备产品ID、设备名称，设置控制设备的属性数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ControlDeviceDataWithContext(ctx context.Context, request *ControlDeviceDataRequest) (response *ControlDeviceDataResponse, err error) {
    if request == nil {
        request = NewControlDeviceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDeviceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIDetectionRequest() (request *CreateAIDetectionRequest) {
    request = &CreateAIDetectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateAIDetection")
    
    
    return
}

func NewCreateAIDetectionResponse() (response *CreateAIDetectionResponse) {
    response = &CreateAIDetectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAIDetection
// 发起AI推理请求
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAIDetection(request *CreateAIDetectionRequest) (response *CreateAIDetectionResponse, err error) {
    return c.CreateAIDetectionWithContext(context.Background(), request)
}

// CreateAIDetection
// 发起AI推理请求
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAIDetectionWithContext(ctx context.Context, request *CreateAIDetectionRequest) (response *CreateAIDetectionResponse, err error) {
    if request == nil {
        request = NewCreateAIDetectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIDetection require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIDetectionResponse()
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

// CreateBatch
// 创建批次
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatch(request *CreateBatchRequest) (response *CreateBatchResponse, err error) {
    return c.CreateBatchWithContext(context.Background(), request)
}

// CreateBatch
// 创建批次
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatchWithContext(ctx context.Context, request *CreateBatchRequest) (response *CreateBatchResponse, err error) {
    if request == nil {
        request = NewCreateBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCOSCredentialsRequest() (request *CreateCOSCredentialsRequest) {
    request = &CreateCOSCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateCOSCredentials")
    
    
    return
}

func NewCreateCOSCredentialsResponse() (response *CreateCOSCredentialsResponse) {
    response = &CreateCOSCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCOSCredentials
// 创建COS上传密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCOSCredentials(request *CreateCOSCredentialsRequest) (response *CreateCOSCredentialsResponse, err error) {
    return c.CreateCOSCredentialsWithContext(context.Background(), request)
}

// CreateCOSCredentials
// 创建COS上传密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCOSCredentialsWithContext(ctx context.Context, request *CreateCOSCredentialsRequest) (response *CreateCOSCredentialsResponse, err error) {
    if request == nil {
        request = NewCreateCOSCredentialsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCOSCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCOSCredentialsResponse()
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

// CreateCloudStorage
// 开通云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudStorage(request *CreateCloudStorageRequest) (response *CreateCloudStorageResponse, err error) {
    return c.CreateCloudStorageWithContext(context.Background(), request)
}

// CreateCloudStorage
// 开通云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudStorageWithContext(ctx context.Context, request *CreateCloudStorageRequest) (response *CreateCloudStorageResponse, err error) {
    if request == nil {
        request = NewCreateCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataForwardRequest() (request *CreateDataForwardRequest) {
    request = &CreateDataForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateDataForward")
    
    
    return
}

func NewCreateDataForwardResponse() (response *CreateDataForwardResponse) {
    response = &CreateDataForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataForward
// 创建数据转发
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataForward(request *CreateDataForwardRequest) (response *CreateDataForwardResponse, err error) {
    return c.CreateDataForwardWithContext(context.Background(), request)
}

// CreateDataForward
// 创建数据转发
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataForwardWithContext(ctx context.Context, request *CreateDataForwardRequest) (response *CreateDataForwardResponse, err error) {
    if request == nil {
        request = NewCreateDataForwardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataForward require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataForwardResponse()
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

// CreateForwardRule
// 创建转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateForwardRule(request *CreateForwardRuleRequest) (response *CreateForwardRuleResponse, err error) {
    return c.CreateForwardRuleWithContext(context.Background(), request)
}

// CreateForwardRule
// 创建转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateForwardRuleWithContext(ctx context.Context, request *CreateForwardRuleRequest) (response *CreateForwardRuleResponse, err error) {
    if request == nil {
        request = NewCreateForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateForwardRule require credential")
    }

    request.SetContext(ctx)
    
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

// CreateProduct
// 创建产品
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    return c.CreateProductWithContext(context.Background(), request)
}

// CreateProduct
// 创建产品
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProductWithContext(ctx context.Context, request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProduct require credential")
    }

    request.SetContext(ctx)
    
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

// CreateTaskFileUrl
// 本接口（CreateTaskFileUrl）用于获取产品级任务文件上传链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTaskFileUrl(request *CreateTaskFileUrlRequest) (response *CreateTaskFileUrlResponse, err error) {
    return c.CreateTaskFileUrlWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFileUrl require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteDevice
// 删除设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 删除设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDeviceWithContext(ctx context.Context, request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevice require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteFirmware
// 本接口（DeleteFirmware）用于删除固件 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFirmware(request *DeleteFirmwareRequest) (response *DeleteFirmwareResponse, err error) {
    return c.DeleteFirmwareWithContext(context.Background(), request)
}

// DeleteFirmware
// 本接口（DeleteFirmware）用于删除固件 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteFirmwareWithContext(ctx context.Context, request *DeleteFirmwareRequest) (response *DeleteFirmwareResponse, err error) {
    if request == nil {
        request = NewDeleteFirmwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFirmware require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteForwardRule
// 删除转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteForwardRule(request *DeleteForwardRuleRequest) (response *DeleteForwardRuleResponse, err error) {
    return c.DeleteForwardRuleWithContext(context.Background(), request)
}

// DeleteForwardRule
// 删除转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteForwardRuleWithContext(ctx context.Context, request *DeleteForwardRuleRequest) (response *DeleteForwardRuleResponse, err error) {
    if request == nil {
        request = NewDeleteForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteForwardRule require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteProduct
// 删除产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 删除产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteProductWithContext(ctx context.Context, request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIModelApplicationsRequest() (request *DescribeAIModelApplicationsRequest) {
    request = &DescribeAIModelApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeAIModelApplications")
    
    
    return
}

func NewDescribeAIModelApplicationsResponse() (response *DescribeAIModelApplicationsResponse) {
    response = &DescribeAIModelApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIModelApplications
// 用户AI模型申请记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelApplications(request *DescribeAIModelApplicationsRequest) (response *DescribeAIModelApplicationsResponse, err error) {
    return c.DescribeAIModelApplicationsWithContext(context.Background(), request)
}

// DescribeAIModelApplications
// 用户AI模型申请记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelApplicationsWithContext(ctx context.Context, request *DescribeAIModelApplicationsRequest) (response *DescribeAIModelApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeAIModelApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIModelApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIModelApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIModelChannelRequest() (request *DescribeAIModelChannelRequest) {
    request = &DescribeAIModelChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeAIModelChannel")
    
    
    return
}

func NewDescribeAIModelChannelResponse() (response *DescribeAIModelChannelResponse) {
    response = &DescribeAIModelChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIModelChannel
// 查看AI推理结果推送配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelChannel(request *DescribeAIModelChannelRequest) (response *DescribeAIModelChannelResponse, err error) {
    return c.DescribeAIModelChannelWithContext(context.Background(), request)
}

// DescribeAIModelChannel
// 查看AI推理结果推送配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelChannelWithContext(ctx context.Context, request *DescribeAIModelChannelRequest) (response *DescribeAIModelChannelResponse, err error) {
    if request == nil {
        request = NewDescribeAIModelChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIModelChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIModelChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIModelUsageRequest() (request *DescribeAIModelUsageRequest) {
    request = &DescribeAIModelUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeAIModelUsage")
    
    
    return
}

func NewDescribeAIModelUsageResponse() (response *DescribeAIModelUsageResponse) {
    response = &DescribeAIModelUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIModelUsage
// 查看AI模型资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelUsage(request *DescribeAIModelUsageRequest) (response *DescribeAIModelUsageResponse, err error) {
    return c.DescribeAIModelUsageWithContext(context.Background(), request)
}

// DescribeAIModelUsage
// 查看AI模型资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelUsageWithContext(ctx context.Context, request *DescribeAIModelUsageRequest) (response *DescribeAIModelUsageResponse, err error) {
    if request == nil {
        request = NewDescribeAIModelUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIModelUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIModelUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIModelsRequest() (request *DescribeAIModelsRequest) {
    request = &DescribeAIModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeAIModels")
    
    
    return
}

func NewDescribeAIModelsResponse() (response *DescribeAIModelsResponse) {
    response = &DescribeAIModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAIModels
// 拉取AI模型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModels(request *DescribeAIModelsRequest) (response *DescribeAIModelsResponse, err error) {
    return c.DescribeAIModelsWithContext(context.Background(), request)
}

// DescribeAIModels
// 拉取AI模型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAIModelsWithContext(ctx context.Context, request *DescribeAIModelsRequest) (response *DescribeAIModelsResponse, err error) {
    if request == nil {
        request = NewDescribeAIModelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIModelsResponse()
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

// DescribeBalance
// 查询账户余额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBalance(request *DescribeBalanceRequest) (response *DescribeBalanceResponse, err error) {
    return c.DescribeBalanceWithContext(context.Background(), request)
}

// DescribeBalance
// 查询账户余额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBalanceWithContext(ctx context.Context, request *DescribeBalanceRequest) (response *DescribeBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBalance require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBalanceTransactions
// 拉取账户流水
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBalanceTransactions(request *DescribeBalanceTransactionsRequest) (response *DescribeBalanceTransactionsResponse, err error) {
    return c.DescribeBalanceTransactionsWithContext(context.Background(), request)
}

// DescribeBalanceTransactions
// 拉取账户流水
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBalanceTransactionsWithContext(ctx context.Context, request *DescribeBalanceTransactionsRequest) (response *DescribeBalanceTransactionsResponse, err error) {
    if request == nil {
        request = NewDescribeBalanceTransactionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBalanceTransactions require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBatch
// 获取批次详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBatch(request *DescribeBatchRequest) (response *DescribeBatchResponse, err error) {
    return c.DescribeBatchWithContext(context.Background(), request)
}

// DescribeBatch
// 获取批次详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBatchWithContext(ctx context.Context, request *DescribeBatchRequest) (response *DescribeBatchResponse, err error) {
    if request == nil {
        request = NewDescribeBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatch require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBatchs
// 获取批次列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBatchs(request *DescribeBatchsRequest) (response *DescribeBatchsResponse, err error) {
    return c.DescribeBatchsWithContext(context.Background(), request)
}

// DescribeBatchs
// 获取批次列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBatchsWithContext(ctx context.Context, request *DescribeBatchsRequest) (response *DescribeBatchsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchs require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCategory
// 获取Category详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCategory(request *DescribeCategoryRequest) (response *DescribeCategoryResponse, err error) {
    return c.DescribeCategoryWithContext(context.Background(), request)
}

// DescribeCategory
// 获取Category详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCategoryWithContext(ctx context.Context, request *DescribeCategoryRequest) (response *DescribeCategoryResponse, err error) {
    if request == nil {
        request = NewDescribeCategoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCategory require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCloudStorage
// 获取设备云存服务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorage(request *DescribeCloudStorageRequest) (response *DescribeCloudStorageResponse, err error) {
    return c.DescribeCloudStorageWithContext(context.Background(), request)
}

// DescribeCloudStorage
// 获取设备云存服务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageWithContext(ctx context.Context, request *DescribeCloudStorageRequest) (response *DescribeCloudStorageResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorage require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCloudStorageDate
// 获取具有云存的日期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageDate(request *DescribeCloudStorageDateRequest) (response *DescribeCloudStorageDateResponse, err error) {
    return c.DescribeCloudStorageDateWithContext(context.Background(), request)
}

// DescribeCloudStorageDate
// 获取具有云存的日期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageDateWithContext(ctx context.Context, request *DescribeCloudStorageDateRequest) (response *DescribeCloudStorageDateResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageDate require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCloudStorageEvents
// 拉取云存事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageEvents(request *DescribeCloudStorageEventsRequest) (response *DescribeCloudStorageEventsResponse, err error) {
    return c.DescribeCloudStorageEventsWithContext(context.Background(), request)
}

// DescribeCloudStorageEvents
// 拉取云存事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageEventsWithContext(ctx context.Context, request *DescribeCloudStorageEventsRequest) (response *DescribeCloudStorageEventsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageEvents require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCloudStorageThumbnail
// 拉取云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageThumbnail(request *DescribeCloudStorageThumbnailRequest) (response *DescribeCloudStorageThumbnailResponse, err error) {
    return c.DescribeCloudStorageThumbnailWithContext(context.Background(), request)
}

// DescribeCloudStorageThumbnail
// 拉取云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageThumbnailWithContext(ctx context.Context, request *DescribeCloudStorageThumbnailRequest) (response *DescribeCloudStorageThumbnailResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageThumbnailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageThumbnail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeCloudStorageTime
// 获取某一天云存时间轴
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageTime(request *DescribeCloudStorageTimeRequest) (response *DescribeCloudStorageTimeResponse, err error) {
    return c.DescribeCloudStorageTimeWithContext(context.Background(), request)
}

// DescribeCloudStorageTime
// 获取某一天云存时间轴
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageTimeWithContext(ctx context.Context, request *DescribeCloudStorageTimeRequest) (response *DescribeCloudStorageTimeResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageUsersRequest() (request *DescribeCloudStorageUsersRequest) {
    request = &DescribeCloudStorageUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeCloudStorageUsers")
    
    
    return
}

func NewDescribeCloudStorageUsersResponse() (response *DescribeCloudStorageUsersResponse) {
    response = &DescribeCloudStorageUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudStorageUsers
// 拉取云存用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageUsers(request *DescribeCloudStorageUsersRequest) (response *DescribeCloudStorageUsersResponse, err error) {
    return c.DescribeCloudStorageUsersWithContext(context.Background(), request)
}

// DescribeCloudStorageUsers
// 拉取云存用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageUsersWithContext(ctx context.Context, request *DescribeCloudStorageUsersRequest) (response *DescribeCloudStorageUsersResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataForwardListRequest() (request *DescribeDataForwardListRequest) {
    request = &DescribeDataForwardListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDataForwardList")
    
    
    return
}

func NewDescribeDataForwardListResponse() (response *DescribeDataForwardListResponse) {
    response = &DescribeDataForwardListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataForwardList
// 获取数据转发列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataForwardList(request *DescribeDataForwardListRequest) (response *DescribeDataForwardListResponse, err error) {
    return c.DescribeDataForwardListWithContext(context.Background(), request)
}

// DescribeDataForwardList
// 获取数据转发列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataForwardListWithContext(ctx context.Context, request *DescribeDataForwardListRequest) (response *DescribeDataForwardListResponse, err error) {
    if request == nil {
        request = NewDescribeDataForwardListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataForwardList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataForwardListResponse()
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

// DescribeDevice
// 查看设备详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    return c.DescribeDeviceWithContext(context.Background(), request)
}

// DescribeDevice
// 查看设备详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceWithContext(ctx context.Context, request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevice require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDeviceActionHistory
// 为用户提供获取动作历史的能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceActionHistory(request *DescribeDeviceActionHistoryRequest) (response *DescribeDeviceActionHistoryResponse, err error) {
    return c.DescribeDeviceActionHistoryWithContext(context.Background(), request)
}

// DescribeDeviceActionHistory
// 为用户提供获取动作历史的能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceActionHistoryWithContext(ctx context.Context, request *DescribeDeviceActionHistoryRequest) (response *DescribeDeviceActionHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceActionHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceActionHistory require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDeviceCommLog
// 获取设备在指定时间范围内的通讯日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceCommLog(request *DescribeDeviceCommLogRequest) (response *DescribeDeviceCommLogResponse, err error) {
    return c.DescribeDeviceCommLogWithContext(context.Background(), request)
}

// DescribeDeviceCommLog
// 获取设备在指定时间范围内的通讯日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceCommLogWithContext(ctx context.Context, request *DescribeDeviceCommLogRequest) (response *DescribeDeviceCommLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceCommLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceCommLog require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDeviceData
// 获取设备属性数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceData(request *DescribeDeviceDataRequest) (response *DescribeDeviceDataResponse, err error) {
    return c.DescribeDeviceDataWithContext(context.Background(), request)
}

// DescribeDeviceData
// 获取设备属性数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceDataWithContext(ctx context.Context, request *DescribeDeviceDataRequest) (response *DescribeDeviceDataResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceData require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDeviceDataHistory
// 获取设备在指定时间范围内上报的历史数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceDataHistory(request *DescribeDeviceDataHistoryRequest) (response *DescribeDeviceDataHistoryResponse, err error) {
    return c.DescribeDeviceDataHistoryWithContext(context.Background(), request)
}

// DescribeDeviceDataHistory
// 获取设备在指定时间范围内上报的历史数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceDataHistoryWithContext(ctx context.Context, request *DescribeDeviceDataHistoryRequest) (response *DescribeDeviceDataHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceDataHistory require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDeviceEventHistory
// 获取设备的历史事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceEventHistory(request *DescribeDeviceEventHistoryRequest) (response *DescribeDeviceEventHistoryResponse, err error) {
    return c.DescribeDeviceEventHistoryWithContext(context.Background(), request)
}

// DescribeDeviceEventHistory
// 获取设备的历史事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceEventHistoryWithContext(ctx context.Context, request *DescribeDeviceEventHistoryRequest) (response *DescribeDeviceEventHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceEventHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceEventHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceEventHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceStatusLogRequest() (request *DescribeDeviceStatusLogRequest) {
    request = &DescribeDeviceStatusLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceStatusLog")
    
    
    return
}

func NewDescribeDeviceStatusLogResponse() (response *DescribeDeviceStatusLogResponse) {
    response = &DescribeDeviceStatusLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceStatusLog
// 获取设备上下线日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceStatusLog(request *DescribeDeviceStatusLogRequest) (response *DescribeDeviceStatusLogResponse, err error) {
    return c.DescribeDeviceStatusLogWithContext(context.Background(), request)
}

// DescribeDeviceStatusLog
// 获取设备上下线日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceStatusLogWithContext(ctx context.Context, request *DescribeDeviceStatusLogRequest) (response *DescribeDeviceStatusLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStatusLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceStatusLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceStatusLogResponse()
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

// DescribeDevices
// 获取设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 获取设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmware")
    
    
    return
}

func NewDescribeFirmwareResponse() (response *DescribeFirmwareResponse) {
    response = &DescribeFirmwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmware
// 本接口（DescribeFirmware）用于查询固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) DescribeFirmware(request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    return c.DescribeFirmwareWithContext(context.Background(), request)
}

// DescribeFirmware
// 本接口（DescribeFirmware）用于查询固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
func (c *Client) DescribeFirmwareWithContext(ctx context.Context, request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmware require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTask")
    
    
    return
}

func NewDescribeFirmwareTaskResponse() (response *DescribeFirmwareTaskResponse) {
    response = &DescribeFirmwareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTask
// 此接口查询固件升级任务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTask(request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    return c.DescribeFirmwareTaskWithContext(context.Background(), request)
}

// DescribeFirmwareTask
// 此接口查询固件升级任务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
func (c *Client) DescribeFirmwareTaskWithContext(ctx context.Context, request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTask require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskDevices")
    
    
    return
}

func NewDescribeFirmwareTaskDevicesResponse() (response *DescribeFirmwareTaskDevicesResponse) {
    response = &DescribeFirmwareTaskDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskDevices
// 本接口用于查询固件升级任务的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskDevices(request *DescribeFirmwareTaskDevicesRequest) (response *DescribeFirmwareTaskDevicesResponse, err error) {
    return c.DescribeFirmwareTaskDevicesWithContext(context.Background(), request)
}

// DescribeFirmwareTaskDevices
// 本接口用于查询固件升级任务的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskDevicesWithContext(ctx context.Context, request *DescribeFirmwareTaskDevicesRequest) (response *DescribeFirmwareTaskDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTaskDevices require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskDistribution")
    
    
    return
}

func NewDescribeFirmwareTaskDistributionResponse() (response *DescribeFirmwareTaskDistributionResponse) {
    response = &DescribeFirmwareTaskDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskDistribution
// 本接口用于查询固件升级任务状态分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskDistribution(request *DescribeFirmwareTaskDistributionRequest) (response *DescribeFirmwareTaskDistributionResponse, err error) {
    return c.DescribeFirmwareTaskDistributionWithContext(context.Background(), request)
}

// DescribeFirmwareTaskDistribution
// 本接口用于查询固件升级任务状态分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskDistributionWithContext(ctx context.Context, request *DescribeFirmwareTaskDistributionRequest) (response *DescribeFirmwareTaskDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskDistributionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTaskDistribution require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTaskStatistics")
    
    
    return
}

func NewDescribeFirmwareTaskStatisticsResponse() (response *DescribeFirmwareTaskStatisticsResponse) {
    response = &DescribeFirmwareTaskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTaskStatistics
// 本接口用于查询固件升级任务统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskStatistics(request *DescribeFirmwareTaskStatisticsRequest) (response *DescribeFirmwareTaskStatisticsResponse, err error) {
    return c.DescribeFirmwareTaskStatisticsWithContext(context.Background(), request)
}

// DescribeFirmwareTaskStatistics
// 本接口用于查询固件升级任务统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTaskStatisticsWithContext(ctx context.Context, request *DescribeFirmwareTaskStatisticsRequest) (response *DescribeFirmwareTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTaskStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTaskStatistics require credential")
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
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeFirmwareTasks")
    
    
    return
}

func NewDescribeFirmwareTasksResponse() (response *DescribeFirmwareTasksResponse) {
    response = &DescribeFirmwareTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirmwareTasks
// 本接口用于查询固件升级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTasks(request *DescribeFirmwareTasksRequest) (response *DescribeFirmwareTasksResponse, err error) {
    return c.DescribeFirmwareTasksWithContext(context.Background(), request)
}

// DescribeFirmwareTasks
// 本接口用于查询固件升级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFirmwareTasksWithContext(ctx context.Context, request *DescribeFirmwareTasksRequest) (response *DescribeFirmwareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTasks require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeForwardRule
// 获取产品转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeForwardRule(request *DescribeForwardRuleRequest) (response *DescribeForwardRuleResponse, err error) {
    return c.DescribeForwardRuleWithContext(context.Background(), request)
}

// DescribeForwardRule
// 获取产品转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeForwardRuleWithContext(ctx context.Context, request *DescribeForwardRuleRequest) (response *DescribeForwardRuleResponse, err error) {
    if request == nil {
        request = NewDescribeForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardRule require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeModelDefinition
// 查询产品配置的数据模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelDefinition(request *DescribeModelDefinitionRequest) (response *DescribeModelDefinitionResponse, err error) {
    return c.DescribeModelDefinitionWithContext(context.Background(), request)
}

// DescribeModelDefinition
// 查询产品配置的数据模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelDefinitionWithContext(ctx context.Context, request *DescribeModelDefinitionRequest) (response *DescribeModelDefinitionResponse, err error) {
    if request == nil {
        request = NewDescribeModelDefinitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelDefinition require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeProduct
// 获取产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    return c.DescribeProductWithContext(context.Background(), request)
}

// DescribeProduct
// 获取产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductWithContext(ctx context.Context, request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    if request == nil {
        request = NewDescribeProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductDynamicRegisterRequest() (request *DescribeProductDynamicRegisterRequest) {
    request = &DescribeProductDynamicRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeProductDynamicRegister")
    
    
    return
}

func NewDescribeProductDynamicRegisterResponse() (response *DescribeProductDynamicRegisterResponse) {
    response = &DescribeProductDynamicRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductDynamicRegister
// 获取产品动态注册详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductDynamicRegister(request *DescribeProductDynamicRegisterRequest) (response *DescribeProductDynamicRegisterResponse, err error) {
    return c.DescribeProductDynamicRegisterWithContext(context.Background(), request)
}

// DescribeProductDynamicRegister
// 获取产品动态注册详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductDynamicRegisterWithContext(ctx context.Context, request *DescribeProductDynamicRegisterRequest) (response *DescribeProductDynamicRegisterResponse, err error) {
    if request == nil {
        request = NewDescribeProductDynamicRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductDynamicRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductDynamicRegisterResponse()
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

// DescribeProducts
// 获取产品列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    return c.DescribeProductsWithContext(context.Background(), request)
}

// DescribeProducts
// 获取产品列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductsWithContext(ctx context.Context, request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSDKLogRequest() (request *DescribeSDKLogRequest) {
    request = &DescribeSDKLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeSDKLog")
    
    
    return
}

func NewDescribeSDKLogResponse() (response *DescribeSDKLogResponse) {
    response = &DescribeSDKLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSDKLog
// 获取设备sdk日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSDKLog(request *DescribeSDKLogRequest) (response *DescribeSDKLogResponse, err error) {
    return c.DescribeSDKLogWithContext(context.Background(), request)
}

// DescribeSDKLog
// 获取设备sdk日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSDKLogWithContext(ctx context.Context, request *DescribeSDKLogRequest) (response *DescribeSDKLogResponse, err error) {
    if request == nil {
        request = NewDescribeSDKLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSDKLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSDKLogResponse()
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

// EditFirmware
// 本接口用于编辑固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
func (c *Client) EditFirmware(request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    return c.EditFirmwareWithContext(context.Background(), request)
}

// EditFirmware
// 本接口用于编辑固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
func (c *Client) EditFirmwareWithContext(ctx context.Context, request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    if request == nil {
        request = NewEditFirmwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditFirmware require credential")
    }

    request.SetContext(ctx)
    
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

// GenerateSignedVideoURL
// 获取视频防盗链播放URL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateSignedVideoURL(request *GenerateSignedVideoURLRequest) (response *GenerateSignedVideoURLResponse, err error) {
    return c.GenerateSignedVideoURLWithContext(context.Background(), request)
}

// GenerateSignedVideoURL
// 获取视频防盗链播放URL
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateSignedVideoURLWithContext(ctx context.Context, request *GenerateSignedVideoURLRequest) (response *GenerateSignedVideoURLResponse, err error) {
    if request == nil {
        request = NewGenerateSignedVideoURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateSignedVideoURL require credential")
    }

    request.SetContext(ctx)
    
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

// GetAllFirmwareVersion
// 本接口（GetAllFirmwareVersion）用于获取所有的版本列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAllFirmwareVersion(request *GetAllFirmwareVersionRequest) (response *GetAllFirmwareVersionResponse, err error) {
    return c.GetAllFirmwareVersionWithContext(context.Background(), request)
}

// GetAllFirmwareVersion
// 本接口（GetAllFirmwareVersion）用于获取所有的版本列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAllFirmwareVersionWithContext(ctx context.Context, request *GetAllFirmwareVersionRequest) (response *GetAllFirmwareVersionResponse, err error) {
    if request == nil {
        request = NewGetAllFirmwareVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAllFirmwareVersion require credential")
    }

    request.SetContext(ctx)
    
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

// GetFirmwareURL
// 本接口（GetFirmwareURL）用于获取固件存储的URL 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFirmwareURL(request *GetFirmwareURLRequest) (response *GetFirmwareURLResponse, err error) {
    return c.GetFirmwareURLWithContext(context.Background(), request)
}

// GetFirmwareURL
// 本接口（GetFirmwareURL）用于获取固件存储的URL 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFirmwareURLWithContext(ctx context.Context, request *GetFirmwareURLRequest) (response *GetFirmwareURLResponse, err error) {
    if request == nil {
        request = NewGetFirmwareURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFirmwareURL require credential")
    }

    request.SetContext(ctx)
    
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

// ImportModelDefinition
// 导入其它产品的数据模板，覆盖现有数据模板的物模型和产品分类信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImportModelDefinition(request *ImportModelDefinitionRequest) (response *ImportModelDefinitionResponse, err error) {
    return c.ImportModelDefinitionWithContext(context.Background(), request)
}

// ImportModelDefinition
// 导入其它产品的数据模板，覆盖现有数据模板的物模型和产品分类信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImportModelDefinitionWithContext(ctx context.Context, request *ImportModelDefinitionRequest) (response *ImportModelDefinitionResponse, err error) {
    if request == nil {
        request = NewImportModelDefinitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportModelDefinition require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportModelDefinitionResponse()
    err = c.Send(request, response)
    return
}

func NewInheritCloudStorageUserRequest() (request *InheritCloudStorageUserRequest) {
    request = &InheritCloudStorageUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "InheritCloudStorageUser")
    
    
    return
}

func NewInheritCloudStorageUserResponse() (response *InheritCloudStorageUserResponse) {
    response = &InheritCloudStorageUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InheritCloudStorageUser
// 继承云存用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InheritCloudStorageUser(request *InheritCloudStorageUserRequest) (response *InheritCloudStorageUserResponse, err error) {
    return c.InheritCloudStorageUserWithContext(context.Background(), request)
}

// InheritCloudStorageUser
// 继承云存用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InheritCloudStorageUserWithContext(ctx context.Context, request *InheritCloudStorageUserRequest) (response *InheritCloudStorageUserResponse, err error) {
    if request == nil {
        request = NewInheritCloudStorageUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InheritCloudStorageUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewInheritCloudStorageUserResponse()
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

// ListFirmwares
// 本接口（ListFirmwares）用于获取固件列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListFirmwares(request *ListFirmwaresRequest) (response *ListFirmwaresResponse, err error) {
    return c.ListFirmwaresWithContext(context.Background(), request)
}

// ListFirmwares
// 本接口（ListFirmwares）用于获取固件列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListFirmwaresWithContext(ctx context.Context, request *ListFirmwaresRequest) (response *ListFirmwaresResponse, err error) {
    if request == nil {
        request = NewListFirmwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListFirmwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewListFirmwaresResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataForwardRequest() (request *ModifyDataForwardRequest) {
    request = &ModifyDataForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDataForward")
    
    
    return
}

func NewModifyDataForwardResponse() (response *ModifyDataForwardResponse) {
    response = &ModifyDataForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDataForward
// 修改数据转发
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataForward(request *ModifyDataForwardRequest) (response *ModifyDataForwardResponse, err error) {
    return c.ModifyDataForwardWithContext(context.Background(), request)
}

// ModifyDataForward
// 修改数据转发
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataForwardWithContext(ctx context.Context, request *ModifyDataForwardRequest) (response *ModifyDataForwardResponse, err error) {
    if request == nil {
        request = NewModifyDataForwardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataForward require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataForwardResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataForwardStatusRequest() (request *ModifyDataForwardStatusRequest) {
    request = &ModifyDataForwardStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDataForwardStatus")
    
    
    return
}

func NewModifyDataForwardStatusResponse() (response *ModifyDataForwardStatusResponse) {
    response = &ModifyDataForwardStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDataForwardStatus
// 设置数据转发状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataForwardStatus(request *ModifyDataForwardStatusRequest) (response *ModifyDataForwardStatusResponse, err error) {
    return c.ModifyDataForwardStatusWithContext(context.Background(), request)
}

// ModifyDataForwardStatus
// 设置数据转发状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataForwardStatusWithContext(ctx context.Context, request *ModifyDataForwardStatusRequest) (response *ModifyDataForwardStatusResponse, err error) {
    if request == nil {
        request = NewModifyDataForwardStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataForwardStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataForwardStatusResponse()
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

// ModifyDevice
// 修改设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    return c.ModifyDeviceWithContext(context.Background(), request)
}

// ModifyDevice
// 修改设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceWithContext(ctx context.Context, request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    if request == nil {
        request = NewModifyDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceLogLevelRequest() (request *ModifyDeviceLogLevelRequest) {
    request = &ModifyDeviceLogLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyDeviceLogLevel")
    
    
    return
}

func NewModifyDeviceLogLevelResponse() (response *ModifyDeviceLogLevelResponse) {
    response = &ModifyDeviceLogLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDeviceLogLevel
// 更新设备日志级别
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceLogLevel(request *ModifyDeviceLogLevelRequest) (response *ModifyDeviceLogLevelResponse, err error) {
    return c.ModifyDeviceLogLevelWithContext(context.Background(), request)
}

// ModifyDeviceLogLevel
// 更新设备日志级别
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceLogLevelWithContext(ctx context.Context, request *ModifyDeviceLogLevelRequest) (response *ModifyDeviceLogLevelResponse, err error) {
    if request == nil {
        request = NewModifyDeviceLogLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceLogLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceLogLevelResponse()
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

// ModifyForwardRule
// 修改转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyForwardRule(request *ModifyForwardRuleRequest) (response *ModifyForwardRuleResponse, err error) {
    return c.ModifyForwardRuleWithContext(context.Background(), request)
}

// ModifyForwardRule
// 修改转发规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyForwardRuleWithContext(ctx context.Context, request *ModifyForwardRuleRequest) (response *ModifyForwardRuleResponse, err error) {
    if request == nil {
        request = NewModifyForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyForwardRule require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyModelDefinition
// 提供修改产品的数据模板的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyModelDefinition(request *ModifyModelDefinitionRequest) (response *ModifyModelDefinitionResponse, err error) {
    return c.ModifyModelDefinitionWithContext(context.Background(), request)
}

// ModifyModelDefinition
// 提供修改产品的数据模板的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyModelDefinitionWithContext(ctx context.Context, request *ModifyModelDefinitionRequest) (response *ModifyModelDefinitionResponse, err error) {
    if request == nil {
        request = NewModifyModelDefinitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelDefinition require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyProduct
// 修改产品信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProduct(request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    return c.ModifyProductWithContext(context.Background(), request)
}

// ModifyProduct
// 修改产品信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProductWithContext(ctx context.Context, request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    if request == nil {
        request = NewModifyProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProductResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProductDynamicRegisterRequest() (request *ModifyProductDynamicRegisterRequest) {
    request = &ModifyProductDynamicRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyProductDynamicRegister")
    
    
    return
}

func NewModifyProductDynamicRegisterResponse() (response *ModifyProductDynamicRegisterResponse) {
    response = &ModifyProductDynamicRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProductDynamicRegister
// 修改产品动态注册
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProductDynamicRegister(request *ModifyProductDynamicRegisterRequest) (response *ModifyProductDynamicRegisterResponse, err error) {
    return c.ModifyProductDynamicRegisterWithContext(context.Background(), request)
}

// ModifyProductDynamicRegister
// 修改产品动态注册
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProductDynamicRegisterWithContext(ctx context.Context, request *ModifyProductDynamicRegisterRequest) (response *ModifyProductDynamicRegisterResponse, err error) {
    if request == nil {
        request = NewModifyProductDynamicRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProductDynamicRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProductDynamicRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewPublishMessageRequest() (request *PublishMessageRequest) {
    request = &PublishMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "PublishMessage")
    
    
    return
}

func NewPublishMessageResponse() (response *PublishMessageResponse) {
    response = &PublishMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishMessage
// 本接口（PublishMessage）用于使用自定义透传协议进行设备远控
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    return c.PublishMessageWithContext(context.Background(), request)
}

// PublishMessage
// 本接口（PublishMessage）用于使用自定义透传协议进行设备远控
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PublishMessageWithContext(ctx context.Context, request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    if request == nil {
        request = NewPublishMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishMessageResponse()
    err = c.Send(request, response)
    return
}

func NewReportAliveDeviceRequest() (request *ReportAliveDeviceRequest) {
    request = &ReportAliveDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ReportAliveDevice")
    
    
    return
}

func NewReportAliveDeviceResponse() (response *ReportAliveDeviceResponse) {
    response = &ReportAliveDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportAliveDevice
// 上报活跃设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReportAliveDevice(request *ReportAliveDeviceRequest) (response *ReportAliveDeviceResponse, err error) {
    return c.ReportAliveDeviceWithContext(context.Background(), request)
}

// ReportAliveDevice
// 上报活跃设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReportAliveDeviceWithContext(ctx context.Context, request *ReportAliveDeviceRequest) (response *ReportAliveDeviceResponse, err error) {
    if request == nil {
        request = NewReportAliveDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportAliveDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportAliveDeviceResponse()
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

// ResetCloudStorage
// 重置云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorage(request *ResetCloudStorageRequest) (response *ResetCloudStorageResponse, err error) {
    return c.ResetCloudStorageWithContext(context.Background(), request)
}

// ResetCloudStorage
// 重置云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorageWithContext(ctx context.Context, request *ResetCloudStorageRequest) (response *ResetCloudStorageResponse, err error) {
    if request == nil {
        request = NewResetCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetCloudStorage require credential")
    }

    request.SetContext(ctx)
    
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

// RetryDeviceFirmwareTask
// 本接口用于重试设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) RetryDeviceFirmwareTask(request *RetryDeviceFirmwareTaskRequest) (response *RetryDeviceFirmwareTaskResponse, err error) {
    return c.RetryDeviceFirmwareTaskWithContext(context.Background(), request)
}

// RetryDeviceFirmwareTask
// 本接口用于重试设备升级任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
func (c *Client) RetryDeviceFirmwareTaskWithContext(ctx context.Context, request *RetryDeviceFirmwareTaskRequest) (response *RetryDeviceFirmwareTaskResponse, err error) {
    if request == nil {
        request = NewRetryDeviceFirmwareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDeviceFirmwareTask require credential")
    }

    request.SetContext(ctx)
    
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

// SetForwardAuth
// 设置转发权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetForwardAuth(request *SetForwardAuthRequest) (response *SetForwardAuthResponse, err error) {
    return c.SetForwardAuthWithContext(context.Background(), request)
}

// SetForwardAuth
// 设置转发权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetForwardAuthWithContext(ctx context.Context, request *SetForwardAuthRequest) (response *SetForwardAuthResponse, err error) {
    if request == nil {
        request = NewSetForwardAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetForwardAuth require credential")
    }

    request.SetContext(ctx)
    
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

// TransferCloudStorage
// 转移云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TransferCloudStorage(request *TransferCloudStorageRequest) (response *TransferCloudStorageResponse, err error) {
    return c.TransferCloudStorageWithContext(context.Background(), request)
}

// TransferCloudStorage
// 转移云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TransferCloudStorageWithContext(ctx context.Context, request *TransferCloudStorageRequest) (response *TransferCloudStorageResponse, err error) {
    if request == nil {
        request = NewTransferCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferCloudStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIModelChannelRequest() (request *UpdateAIModelChannelRequest) {
    request = &UpdateAIModelChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "UpdateAIModelChannel")
    
    
    return
}

func NewUpdateAIModelChannelResponse() (response *UpdateAIModelChannelResponse) {
    response = &UpdateAIModelChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAIModelChannel
// 更新AI推理结果推送配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAIModelChannel(request *UpdateAIModelChannelRequest) (response *UpdateAIModelChannelResponse, err error) {
    return c.UpdateAIModelChannelWithContext(context.Background(), request)
}

// UpdateAIModelChannel
// 更新AI推理结果推送配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAIModelChannelWithContext(ctx context.Context, request *UpdateAIModelChannelRequest) (response *UpdateAIModelChannelResponse, err error) {
    if request == nil {
        request = NewUpdateAIModelChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIModelChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIModelChannelResponse()
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

// UploadFirmware
// 本接口（UploadFirmware）用于上传设备固件信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    return c.UploadFirmwareWithContext(context.Background(), request)
}

// UploadFirmware
// 本接口（UploadFirmware）用于上传设备固件信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
func (c *Client) UploadFirmwareWithContext(ctx context.Context, request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFirmware require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}

func NewWakeUpDeviceRequest() (request *WakeUpDeviceRequest) {
    request = &WakeUpDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "WakeUpDevice")
    
    
    return
}

func NewWakeUpDeviceResponse() (response *WakeUpDeviceResponse) {
    response = &WakeUpDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WakeUpDevice
// 设备唤醒
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) WakeUpDevice(request *WakeUpDeviceRequest) (response *WakeUpDeviceResponse, err error) {
    return c.WakeUpDeviceWithContext(context.Background(), request)
}

// WakeUpDevice
// 设备唤醒
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) WakeUpDeviceWithContext(ctx context.Context, request *WakeUpDeviceRequest) (response *WakeUpDeviceResponse, err error) {
    if request == nil {
        request = NewWakeUpDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WakeUpDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewWakeUpDeviceResponse()
    err = c.Send(request, response)
    return
}
