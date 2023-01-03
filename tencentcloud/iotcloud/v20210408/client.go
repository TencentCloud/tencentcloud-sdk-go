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

package v20210408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-08"

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
//  FAILEDOPERATION_BINDDEVICEPERONCEOVERLIMIT = "FailedOperation.BindDevicePerOnceOverLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) BindDevices(request *BindDevicesRequest) (response *BindDevicesResponse, err error) {
    return c.BindDevicesWithContext(context.Background(), request)
}

// BindDevices
// 本接口（BindDevices）用于网关设备批量绑定子设备 
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDDEVICEOVERLIMIT = "FailedOperation.BindDeviceOverLimit"
//  FAILEDOPERATION_BINDDEVICEPERONCEOVERLIMIT = "FailedOperation.BindDevicePerOnceOverLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDevices require credential")
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
    return c.CancelDeviceFirmwareTaskWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDeviceFirmwareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelDeviceFirmwareTaskResponse()
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
//  FAILEDOPERATION_ALREADYDISTRIBUTIONDEVICE = "FailedOperation.AlreadyDistributionDevice"
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    return c.CreateDeviceWithContext(context.Background(), request)
}

// CreateDevice
// 本接口（CreateDevice）用于新建一个物联网通信设备。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALREADYDISTRIBUTIONDEVICE = "FailedOperation.AlreadyDistributionDevice"
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDeviceWithContext(ctx context.Context, request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) CreateMultiDevicesTask(request *CreateMultiDevicesTaskRequest) (response *CreateMultiDevicesTaskResponse, err error) {
    return c.CreateMultiDevicesTaskWithContext(context.Background(), request)
}

// CreateMultiDevicesTask
// 本接口（CreateMultiDevicesTask）用于创建产品级别的批量创建设备任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) CreateMultiDevicesTaskWithContext(ctx context.Context, request *CreateMultiDevicesTaskRequest) (response *CreateMultiDevicesTaskResponse, err error) {
    if request == nil {
        request = NewCreateMultiDevicesTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiDevicesTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiDevicesTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateCARequest() (request *CreatePrivateCARequest) {
    request = &CreatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreatePrivateCA")
    
    
    return
}

func NewCreatePrivateCAResponse() (response *CreatePrivateCAResponse) {
    response = &CreatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateCA
// 创建私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CACERTLIMIT = "LimitExceeded.CACertLimit"
//  LIMITEXCEEDED_CACERTNAMEREPEAT = "LimitExceeded.CACertNameRepeat"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) CreatePrivateCA(request *CreatePrivateCARequest) (response *CreatePrivateCAResponse, err error) {
    return c.CreatePrivateCAWithContext(context.Background(), request)
}

// CreatePrivateCA
// 创建私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CACERTLIMIT = "LimitExceeded.CACertLimit"
//  LIMITEXCEEDED_CACERTNAMEREPEAT = "LimitExceeded.CACertNameRepeat"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) CreatePrivateCAWithContext(ctx context.Context, request *CreatePrivateCARequest) (response *CreatePrivateCAResponse, err error) {
    if request == nil {
        request = NewCreatePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateCAResponse()
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
//  FAILEDOPERATION_ACCOUNTISOLATED = "FailedOperation.AccountIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  INVALIDPARAMETERVALUE_TIDPRODUCTALREADYEXIST = "InvalidParameterValue.TidProductAlreadyExist"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_THINGMODELNOTEXIST = "ResourceNotFound.ThingModelNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    return c.CreateProductWithContext(context.Background(), request)
}

// CreateProduct
// 本接口（CreateProduct）用于创建一个新的物联网通信产品 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTISOLATED = "FailedOperation.AccountIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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
    return c.CreateTopicPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopicPolicy require credential")
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
    return c.CreateTopicRuleWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopicRule require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除物联网通信设备。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceResourceRequest() (request *DeleteDeviceResourceRequest) {
    request = &DeleteDeviceResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteDeviceResource")
    
    
    return
}

func NewDeleteDeviceResourceResponse() (response *DeleteDeviceResourceResponse) {
    response = &DeleteDeviceResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDeviceResource
// 本接口（DeleteDeviceResource）用于删除设备资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) DeleteDeviceResource(request *DeleteDeviceResourceRequest) (response *DeleteDeviceResourceResponse, err error) {
    return c.DeleteDeviceResourceWithContext(context.Background(), request)
}

// DeleteDeviceResource
// 本接口（DeleteDeviceResource）用于删除设备资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) DeleteDeviceResourceWithContext(ctx context.Context, request *DeleteDeviceResourceRequest) (response *DeleteDeviceResourceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceShadowRequest() (request *DeleteDeviceShadowRequest) {
    request = &DeleteDeviceShadowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteDeviceShadow")
    
    
    return
}

func NewDeleteDeviceShadowResponse() (response *DeleteDeviceShadowResponse) {
    response = &DeleteDeviceShadowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDeviceShadow
// 本接口（DeleteDeviceShadow）用于删除设备影子 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
func (c *Client) DeleteDeviceShadow(request *DeleteDeviceShadowRequest) (response *DeleteDeviceShadowResponse, err error) {
    return c.DeleteDeviceShadowWithContext(context.Background(), request)
}

// DeleteDeviceShadow
// 本接口（DeleteDeviceShadow）用于删除设备影子 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
func (c *Client) DeleteDeviceShadowWithContext(ctx context.Context, request *DeleteDeviceShadowRequest) (response *DeleteDeviceShadowResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceShadowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceShadow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceShadowResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateCARequest() (request *DeletePrivateCARequest) {
    request = &DeletePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeletePrivateCA")
    
    
    return
}

func NewDeletePrivateCAResponse() (response *DeletePrivateCAResponse) {
    response = &DeletePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrivateCA
// 删除私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DeletePrivateCA(request *DeletePrivateCARequest) (response *DeletePrivateCAResponse, err error) {
    return c.DeletePrivateCAWithContext(context.Background(), request)
}

// DeletePrivateCA
// 删除私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DeletePrivateCAWithContext(ctx context.Context, request *DeletePrivateCARequest) (response *DeletePrivateCAResponse, err error) {
    if request == nil {
        request = NewDeletePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateCAResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DELETETIDFAIL = "UnauthorizedOperation.DeleteTidFail"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网通信产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DELETETIDFAIL = "UnauthorizedOperation.DeleteTidFail"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
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

func NewDeleteProductPrivateCARequest() (request *DeleteProductPrivateCARequest) {
    request = &DeleteProductPrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteProductPrivateCA")
    
    
    return
}

func NewDeleteProductPrivateCAResponse() (response *DeleteProductPrivateCAResponse) {
    response = &DeleteProductPrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProductPrivateCA
// 删除产品的私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
func (c *Client) DeleteProductPrivateCA(request *DeleteProductPrivateCARequest) (response *DeleteProductPrivateCAResponse, err error) {
    return c.DeleteProductPrivateCAWithContext(context.Background(), request)
}

// DeleteProductPrivateCA
// 删除产品的私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
func (c *Client) DeleteProductPrivateCAWithContext(ctx context.Context, request *DeleteProductPrivateCARequest) (response *DeleteProductPrivateCAResponse, err error) {
    if request == nil {
        request = NewDeleteProductPrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProductPrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProductPrivateCAResponse()
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
    return c.DeleteTopicRuleWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicRuleResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    return c.DescribeDeviceWithContext(context.Background(), request)
}

// DescribeDevice
// 本接口（DescribeDevice）用于查看设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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
    return c.DescribeDeviceClientKeyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceClientKey require credential")
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
    return c.DescribeDeviceResourceWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceResource require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResources(request *DescribeDeviceResourcesRequest) (response *DescribeDeviceResourcesResponse, err error) {
    return c.DescribeDeviceResourcesWithContext(context.Background(), request)
}

// DescribeDeviceResources
// 本接口（DescribeDeviceResources）用于查询设备资源列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceResourcesWithContext(ctx context.Context, request *DescribeDeviceResourcesRequest) (response *DescribeDeviceResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceResources require credential")
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
    return c.DescribeDeviceShadowWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceShadow require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 本接口（DescribeDevices）用于查询物联网通信设备的设备列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeFirmware(request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    return c.DescribeFirmwareWithContext(context.Background(), request)
}

// DescribeFirmware
// 查询固件信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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
    return c.DescribeFirmwareTaskWithContext(context.Background(), request)
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
    return c.DescribeFirmwareTaskDevicesWithContext(context.Background(), request)
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
    return c.DescribeFirmwareTaskDistributionWithContext(context.Background(), request)
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
    return c.DescribeFirmwareTaskStatisticsWithContext(context.Background(), request)
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
    return c.DescribeFirmwareTasksWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirmwareTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayBindDevicesRequest() (request *DescribeGatewayBindDevicesRequest) {
    request = &DescribeGatewayBindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeGatewayBindDevices")
    
    
    return
}

func NewDescribeGatewayBindDevicesResponse() (response *DescribeGatewayBindDevicesResponse) {
    response = &DescribeGatewayBindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayBindDevices
// 本接口（DescribeGatewayBindDevices）用于获取网关绑定的子设备列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeGatewayBindDevices(request *DescribeGatewayBindDevicesRequest) (response *DescribeGatewayBindDevicesResponse, err error) {
    return c.DescribeGatewayBindDevicesWithContext(context.Background(), request)
}

// DescribeGatewayBindDevices
// 本接口（DescribeGatewayBindDevices）用于获取网关绑定的子设备列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeGatewayBindDevicesWithContext(ctx context.Context, request *DescribeGatewayBindDevicesRequest) (response *DescribeGatewayBindDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayBindDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayBindDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayBindDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCARequest() (request *DescribePrivateCARequest) {
    request = &DescribePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCA")
    
    
    return
}

func NewDescribePrivateCAResponse() (response *DescribePrivateCAResponse) {
    response = &DescribePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCA
// 查询私有化CA信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCA(request *DescribePrivateCARequest) (response *DescribePrivateCAResponse, err error) {
    return c.DescribePrivateCAWithContext(context.Background(), request)
}

// DescribePrivateCA
// 查询私有化CA信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCAWithContext(ctx context.Context, request *DescribePrivateCARequest) (response *DescribePrivateCAResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCABindedProductsRequest() (request *DescribePrivateCABindedProductsRequest) {
    request = &DescribePrivateCABindedProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCABindedProducts")
    
    
    return
}

func NewDescribePrivateCABindedProductsResponse() (response *DescribePrivateCABindedProductsResponse) {
    response = &DescribePrivateCABindedProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCABindedProducts
// 查询私有CA绑定的产品列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) DescribePrivateCABindedProducts(request *DescribePrivateCABindedProductsRequest) (response *DescribePrivateCABindedProductsResponse, err error) {
    return c.DescribePrivateCABindedProductsWithContext(context.Background(), request)
}

// DescribePrivateCABindedProducts
// 查询私有CA绑定的产品列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) DescribePrivateCABindedProductsWithContext(ctx context.Context, request *DescribePrivateCABindedProductsRequest) (response *DescribePrivateCABindedProductsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCABindedProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCABindedProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCABindedProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCAsRequest() (request *DescribePrivateCAsRequest) {
    request = &DescribePrivateCAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCAs")
    
    
    return
}

func NewDescribePrivateCAsResponse() (response *DescribePrivateCAsResponse) {
    response = &DescribePrivateCAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCAs
// 查询私有CA证书列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) DescribePrivateCAs(request *DescribePrivateCAsRequest) (response *DescribePrivateCAsResponse, err error) {
    return c.DescribePrivateCAsWithContext(context.Background(), request)
}

// DescribePrivateCAs
// 查询私有CA证书列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) DescribePrivateCAsWithContext(ctx context.Context, request *DescribePrivateCAsRequest) (response *DescribePrivateCAsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCAsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductRequest() (request *DescribeProductRequest) {
    request = &DescribeProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProduct")
    
    
    return
}

func NewDescribeProductResponse() (response *DescribeProductResponse) {
    response = &DescribeProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProduct
// 本接口（DescribeProduct）用于查看产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    return c.DescribeProductWithContext(context.Background(), request)
}

// DescribeProduct
// 本接口（DescribeProduct）用于查看产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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

func NewDescribeProductCARequest() (request *DescribeProductCARequest) {
    request = &DescribeProductCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductCA")
    
    
    return
}

func NewDescribeProductCAResponse() (response *DescribeProductCAResponse) {
    response = &DescribeProductCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductCA
// 查询产品绑定的CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductCA(request *DescribeProductCARequest) (response *DescribeProductCAResponse, err error) {
    return c.DescribeProductCAWithContext(context.Background(), request)
}

// DescribeProductCA
// 查询产品绑定的CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductCAWithContext(ctx context.Context, request *DescribeProductCARequest) (response *DescribeProductCAResponse, err error) {
    if request == nil {
        request = NewDescribeProductCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductCAResponse()
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
    return c.DescribeProductResourceWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductResource require credential")
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
    return c.DescribeProductResourcesWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductResources require credential")
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
    return c.DescribeProductTaskWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductTask require credential")
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
    return c.DescribeProductTasksWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductTasks require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    return c.DescribeProductsWithContext(context.Background(), request)
}

// DescribeProducts
// 本接口（DescribeProducts）用于列出产品列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribePushResourceTaskStatistics(request *DescribePushResourceTaskStatisticsRequest) (response *DescribePushResourceTaskStatisticsResponse, err error) {
    return c.DescribePushResourceTaskStatisticsWithContext(context.Background(), request)
}

// DescribePushResourceTaskStatistics
// 查询推送资源任务统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribePushResourceTaskStatisticsWithContext(ctx context.Context, request *DescribePushResourceTaskStatisticsRequest) (response *DescribePushResourceTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribePushResourceTaskStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushResourceTaskStatistics require credential")
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeResourceTasks(request *DescribeResourceTasksRequest) (response *DescribeResourceTasksResponse, err error) {
    return c.DescribeResourceTasksWithContext(context.Background(), request)
}

// DescribeResourceTasks
// 查询资源推送任务列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
func (c *Client) DescribeResourceTasksWithContext(ctx context.Context, request *DescribeResourceTasksRequest) (response *DescribeResourceTasksResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTasksResponse()
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
    return c.DisableTopicRuleWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadDeviceResourceRequest() (request *DownloadDeviceResourceRequest) {
    request = &DownloadDeviceResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "DownloadDeviceResource")
    
    
    return
}

func NewDownloadDeviceResourceResponse() (response *DownloadDeviceResourceResponse) {
    response = &DownloadDeviceResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadDeviceResource
// 本接口（DownloadDeviceResource）用于下载设备资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEFILENOTMATCH = "FailedOperation.ResourceFileNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_RESOURCEFILENOTEXIST = "ResourceNotFound.ResourceFileNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) DownloadDeviceResource(request *DownloadDeviceResourceRequest) (response *DownloadDeviceResourceResponse, err error) {
    return c.DownloadDeviceResourceWithContext(context.Background(), request)
}

// DownloadDeviceResource
// 本接口（DownloadDeviceResource）用于下载设备资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEFILENOTMATCH = "FailedOperation.ResourceFileNotMatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_RESOURCEFILENOTEXIST = "ResourceNotFound.ResourceFileNotExist"
//  UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"
func (c *Client) DownloadDeviceResourceWithContext(ctx context.Context, request *DownloadDeviceResourceRequest) (response *DownloadDeviceResourceResponse, err error) {
    if request == nil {
        request = NewDownloadDeviceResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadDeviceResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadDeviceResourceResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
func (c *Client) EditFirmware(request *EditFirmwareRequest) (response *EditFirmwareResponse, err error) {
    return c.EditFirmwareWithContext(context.Background(), request)
}

// EditFirmware
// 编辑固件信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) EnableTopicRule(request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    return c.EnableTopicRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) EnableTopicRuleWithContext(ctx context.Context, request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    if request == nil {
        request = NewEnableTopicRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetAllVersionRequest() (request *GetAllVersionRequest) {
    request = &GetAllVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "GetAllVersion")
    
    
    return
}

func NewGetAllVersionResponse() (response *GetAllVersionResponse) {
    response = &GetAllVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAllVersion
// 本接口（GetAllVersion）用于获取所有的版本列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAllVersion(request *GetAllVersionRequest) (response *GetAllVersionResponse, err error) {
    return c.GetAllVersionWithContext(context.Background(), request)
}

// GetAllVersion
// 本接口（GetAllVersion）用于获取所有的版本列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAllVersionWithContext(ctx context.Context, request *GetAllVersionRequest) (response *GetAllVersionResponse, err error) {
    if request == nil {
        request = NewGetAllVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAllVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAllVersionResponse()
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
    return c.GetCOSURLWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCOSURL require credential")
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
    return c.GetUserResourceInfoWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUserResourceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserResourceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewListFirmwaresRequest() (request *ListFirmwaresRequest) {
    request = &ListFirmwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListFirmwares")
    
    
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

func NewListLogRequest() (request *ListLogRequest) {
    request = &ListLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListLog")
    
    
    return
}

func NewListLogResponse() (response *ListLogResponse) {
    response = &ListLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLog
// 本接口（ListLog）用于查看日志信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLog(request *ListLogRequest) (response *ListLogResponse, err error) {
    return c.ListLogWithContext(context.Background(), request)
}

// ListLog
// 本接口（ListLog）用于查看日志信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogWithContext(ctx context.Context, request *ListLogRequest) (response *ListLogResponse, err error) {
    if request == nil {
        request = NewListLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogResponse()
    err = c.Send(request, response)
    return
}

func NewListLogPayloadRequest() (request *ListLogPayloadRequest) {
    request = &ListLogPayloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListLogPayload")
    
    
    return
}

func NewListLogPayloadResponse() (response *ListLogPayloadResponse) {
    response = &ListLogPayloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLogPayload
// 获取日志内容列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogPayload(request *ListLogPayloadRequest) (response *ListLogPayloadResponse, err error) {
    return c.ListLogPayloadWithContext(context.Background(), request)
}

// ListLogPayload
// 获取日志内容列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogPayloadWithContext(ctx context.Context, request *ListLogPayloadRequest) (response *ListLogPayloadResponse, err error) {
    if request == nil {
        request = NewListLogPayloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLogPayload require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogPayloadResponse()
    err = c.Send(request, response)
    return
}

func NewListSDKLogRequest() (request *ListSDKLogRequest) {
    request = &ListSDKLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListSDKLog")
    
    
    return
}

func NewListSDKLogResponse() (response *ListSDKLogResponse) {
    response = &ListSDKLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSDKLog
// 获取设备上报的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) ListSDKLog(request *ListSDKLogRequest) (response *ListSDKLogResponse, err error) {
    return c.ListSDKLogWithContext(context.Background(), request)
}

// ListSDKLog
// 获取设备上报的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) ListSDKLogWithContext(ctx context.Context, request *ListSDKLogRequest) (response *ListSDKLogResponse, err error) {
    if request == nil {
        request = NewListSDKLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSDKLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSDKLogResponse()
    err = c.Send(request, response)
    return
}

func NewListTopicRulesRequest() (request *ListTopicRulesRequest) {
    request = &ListTopicRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListTopicRules")
    
    
    return
}

func NewListTopicRulesResponse() (response *ListTopicRulesResponse) {
    response = &ListTopicRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopicRules
// 本接口（ListTopicRules）用于分页获取规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) ListTopicRules(request *ListTopicRulesRequest) (response *ListTopicRulesResponse, err error) {
    return c.ListTopicRulesWithContext(context.Background(), request)
}

// ListTopicRules
// 本接口（ListTopicRules）用于分页获取规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
func (c *Client) ListTopicRulesWithContext(ctx context.Context, request *ListTopicRulesRequest) (response *ListTopicRulesResponse, err error) {
    if request == nil {
        request = NewListTopicRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopicRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopicRulesResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessage(request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    return c.PublishBroadcastMessageWithContext(context.Background(), request)
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BROADCASTTASKISRUNNING = "FailedOperation.BroadcastTaskIsRunning"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessageWithContext(ctx context.Context, request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    if request == nil {
        request = NewPublishBroadcastMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishBroadcastMessage require credential")
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
    return c.PublishMessageWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishMessage require credential")
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
    return c.PublishRRPCMessageWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishRRPCMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishRRPCMessageResponse()
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
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_OPERATIONDENIED = "InvalidParameterValue.OperationDenied"
//  INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReplaceTopicRule(request *ReplaceTopicRuleRequest) (response *ReplaceTopicRuleResponse, err error) {
    return c.ReplaceTopicRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_OPERATIONDENIED = "InvalidParameterValue.OperationDenied"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceTopicRule require credential")
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
    return c.ResetDeviceStateWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDeviceState require credential")
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
    return c.RetryDeviceFirmwareTaskWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDeviceFirmwareTask require credential")
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
    return c.SetProductsForbiddenStatusWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetProductsForbiddenStatus require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UnbindDevices(request *UnbindDevicesRequest) (response *UnbindDevicesResponse, err error) {
    return c.UnbindDevicesWithContext(context.Background(), request)
}

// UnbindDevices
// 本接口（UnbindDevices）用于网关设备批量解绑子设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UnbindDevicesWithContext(ctx context.Context, request *UnbindDevicesRequest) (response *UnbindDevicesResponse, err error) {
    if request == nil {
        request = NewUnbindDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDevices require credential")
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
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateDeviceAvailableState(request *UpdateDeviceAvailableStateRequest) (response *UpdateDeviceAvailableStateResponse, err error) {
    return c.UpdateDeviceAvailableStateWithContext(context.Background(), request)
}

// UpdateDeviceAvailableState
// 启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateDeviceAvailableStateWithContext(ctx context.Context, request *UpdateDeviceAvailableStateRequest) (response *UpdateDeviceAvailableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceAvailableStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceAvailableState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceAvailableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceLogLevelRequest() (request *UpdateDeviceLogLevelRequest) {
    request = &UpdateDeviceLogLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDeviceLogLevel")
    
    
    return
}

func NewUpdateDeviceLogLevelResponse() (response *UpdateDeviceLogLevelResponse) {
    response = &UpdateDeviceLogLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceLogLevel
// 设置设备上报的日志级别  
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) UpdateDeviceLogLevel(request *UpdateDeviceLogLevelRequest) (response *UpdateDeviceLogLevelResponse, err error) {
    return c.UpdateDeviceLogLevelWithContext(context.Background(), request)
}

// UpdateDeviceLogLevel
// 设置设备上报的日志级别  
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) UpdateDeviceLogLevelWithContext(ctx context.Context, request *UpdateDeviceLogLevelRequest) (response *UpdateDeviceLogLevelResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceLogLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceLogLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceLogLevelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicePSKRequest() (request *UpdateDevicePSKRequest) {
    request = &UpdateDevicePSKRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDevicePSK")
    
    
    return
}

func NewUpdateDevicePSKResponse() (response *UpdateDevicePSKResponse) {
    response = &UpdateDevicePSKResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicePSK
// 本接口（UpdateDevicePSK）用于更新设备的PSK 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
func (c *Client) UpdateDevicePSK(request *UpdateDevicePSKRequest) (response *UpdateDevicePSKResponse, err error) {
    return c.UpdateDevicePSKWithContext(context.Background(), request)
}

// UpdateDevicePSK
// 本接口（UpdateDevicePSK）用于更新设备的PSK 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
func (c *Client) UpdateDevicePSKWithContext(ctx context.Context, request *UpdateDevicePSKRequest) (response *UpdateDevicePSKResponse, err error) {
    if request == nil {
        request = NewUpdateDevicePSKRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevicePSK require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDevicePSKResponse()
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
//  INVALIDPARAMETERVALUE_NOTMERGEABLE = "InvalidParameterValue.NotMergeAble"
//  INVALIDPARAMETERVALUE_PREFIXINVALID = "InvalidParameterValue.PrefixInvalid"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) UpdateDeviceShadow(request *UpdateDeviceShadowRequest) (response *UpdateDeviceShadowResponse, err error) {
    return c.UpdateDeviceShadowWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_NOTMERGEABLE = "InvalidParameterValue.NotMergeAble"
//  INVALIDPARAMETERVALUE_PREFIXINVALID = "InvalidParameterValue.PrefixInvalid"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
func (c *Client) UpdateDeviceShadowWithContext(ctx context.Context, request *UpdateDeviceShadowRequest) (response *UpdateDeviceShadowResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceShadowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceShadow require credential")
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateDevicesEnableState(request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    return c.UpdateDevicesEnableStateWithContext(context.Background(), request)
}

// UpdateDevicesEnableState
// 批量启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateDevicesEnableStateWithContext(ctx context.Context, request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevicesEnableState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDevicesEnableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrivateCARequest() (request *UpdatePrivateCARequest) {
    request = &UpdatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdatePrivateCA")
    
    
    return
}

func NewUpdatePrivateCAResponse() (response *UpdatePrivateCAResponse) {
    response = &UpdatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePrivateCA
// 更新私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) UpdatePrivateCA(request *UpdatePrivateCARequest) (response *UpdatePrivateCAResponse, err error) {
    return c.UpdatePrivateCAWithContext(context.Background(), request)
}

// UpdatePrivateCA
// 更新私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) UpdatePrivateCAWithContext(ctx context.Context, request *UpdatePrivateCARequest) (response *UpdatePrivateCAResponse, err error) {
    if request == nil {
        request = NewUpdatePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductDynamicRegisterRequest() (request *UpdateProductDynamicRegisterRequest) {
    request = &UpdateProductDynamicRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateProductDynamicRegister")
    
    
    return
}

func NewUpdateProductDynamicRegisterResponse() (response *UpdateProductDynamicRegisterResponse) {
    response = &UpdateProductDynamicRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProductDynamicRegister
// 更新产品动态注册的配置 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateProductDynamicRegister(request *UpdateProductDynamicRegisterRequest) (response *UpdateProductDynamicRegisterResponse, err error) {
    return c.UpdateProductDynamicRegisterWithContext(context.Background(), request)
}

// UpdateProductDynamicRegister
// 更新产品动态注册的配置 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateProductDynamicRegisterWithContext(ctx context.Context, request *UpdateProductDynamicRegisterRequest) (response *UpdateProductDynamicRegisterResponse, err error) {
    if request == nil {
        request = NewUpdateProductDynamicRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProductDynamicRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProductDynamicRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductPrivateCARequest() (request *UpdateProductPrivateCARequest) {
    request = &UpdateProductPrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateProductPrivateCA")
    
    
    return
}

func NewUpdateProductPrivateCAResponse() (response *UpdateProductPrivateCAResponse) {
    response = &UpdateProductPrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProductPrivateCA
// 更新产品的私有CA
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CACERTNOTSUPPORT = "LimitExceeded.CACertNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) UpdateProductPrivateCA(request *UpdateProductPrivateCARequest) (response *UpdateProductPrivateCAResponse, err error) {
    return c.UpdateProductPrivateCAWithContext(context.Background(), request)
}

// UpdateProductPrivateCA
// 更新产品的私有CA
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CACERTNOTSUPPORT = "LimitExceeded.CACertNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) UpdateProductPrivateCAWithContext(ctx context.Context, request *UpdateProductPrivateCARequest) (response *UpdateProductPrivateCAResponse, err error) {
    if request == nil {
        request = NewUpdateProductPrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProductPrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProductPrivateCAResponse()
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
    return c.UpdateTopicPolicyWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTopicPolicy require credential")
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
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
