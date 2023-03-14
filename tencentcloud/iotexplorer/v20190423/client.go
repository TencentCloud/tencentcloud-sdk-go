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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBindDevicesRequest() (request *BindDevicesRequest) {
    request = &BindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "BindDevices")
    
    
    return
}

func NewBindDevicesResponse() (response *BindDevicesResponse) {
    response = &BindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindDevices
// 批量绑定子设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) BindDevices(request *BindDevicesRequest) (response *BindDevicesResponse, err error) {
    return c.BindDevicesWithContext(context.Background(), request)
}

// BindDevices
// 批量绑定子设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewBindProductsRequest() (request *BindProductsRequest) {
    request = &BindProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "BindProducts")
    
    
    return
}

func NewBindProductsResponse() (response *BindProductsResponse) {
    response = &BindProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindProducts
// 批量绑定子产品
//
// 可能返回的错误码:
//  FAILEDOPERATION_SOMEPRODUCTISALREADYBINDED = "FailedOperation.SomeProductIsAlreadyBinded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRODUCTISNOTGATEWAY = "InvalidParameter.ProductIsNotGateway"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTISNOTGATEWAY = "InvalidParameterValue.ProductIsNotGateway"
//  LIMITEXCEEDED_BINDPRODUCTSEXCEEDLIMIT = "LimitExceeded.BindProductsExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindProducts(request *BindProductsRequest) (response *BindProductsResponse, err error) {
    return c.BindProductsWithContext(context.Background(), request)
}

// BindProducts
// 批量绑定子产品
//
// 可能返回的错误码:
//  FAILEDOPERATION_SOMEPRODUCTISALREADYBINDED = "FailedOperation.SomeProductIsAlreadyBinded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PRODUCTISNOTGATEWAY = "InvalidParameter.ProductIsNotGateway"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTISNOTGATEWAY = "InvalidParameterValue.ProductIsNotGateway"
//  LIMITEXCEEDED_BINDPRODUCTSEXCEEDLIMIT = "LimitExceeded.BindProductsExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindProductsWithContext(ctx context.Context, request *BindProductsRequest) (response *BindProductsResponse, err error) {
    if request == nil {
        request = NewBindProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindProductsResponse()
    err = c.Send(request, response)
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

// CallDeviceActionAsync
// 提供给用户异步调用设备行为的能力
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONUNREACHABLE = "FailedOperation.ActionUnreachable"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER_ACTIONINPUTPARAMSINVALID = "InvalidParameter.ActionInputParamsInvalid"
//  INVALIDPARAMETERVALUE_ACTIONNILORNOTEXIST = "InvalidParameterValue.ActionNilOrNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELPROPERTYNOTEXIST = "InvalidParameterValue.ModelPropertyNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CallDeviceActionAsync(request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    return c.CallDeviceActionAsyncWithContext(context.Background(), request)
}

// CallDeviceActionAsync
// 提供给用户异步调用设备行为的能力
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONUNREACHABLE = "FailedOperation.ActionUnreachable"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER_ACTIONINPUTPARAMSINVALID = "InvalidParameter.ActionInputParamsInvalid"
//  INVALIDPARAMETERVALUE_ACTIONNILORNOTEXIST = "InvalidParameterValue.ActionNilOrNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELPROPERTYNOTEXIST = "InvalidParameterValue.ModelPropertyNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CallDeviceActionAsyncWithContext(ctx context.Context, request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallDeviceActionAsync require credential")
    }

    request.SetContext(ctx)
    
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

// CallDeviceActionSync
// 为用户提供同步调用设备行为的能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONUNREACHABLE = "FailedOperation.ActionUnreachable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER_ACTIONINPUTPARAMSINVALID = "InvalidParameter.ActionInputParamsInvalid"
//  INVALIDPARAMETERVALUE_ACTIONNILORNOTEXIST = "InvalidParameterValue.ActionNilOrNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELPROPERTYNOTEXIST = "InvalidParameterValue.ModelPropertyNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CallDeviceActionSync(request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    return c.CallDeviceActionSyncWithContext(context.Background(), request)
}

// CallDeviceActionSync
// 为用户提供同步调用设备行为的能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONUNREACHABLE = "FailedOperation.ActionUnreachable"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER_ACTIONINPUTPARAMSINVALID = "InvalidParameter.ActionInputParamsInvalid"
//  INVALIDPARAMETERVALUE_ACTIONNILORNOTEXIST = "InvalidParameterValue.ActionNilOrNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELPROPERTYNOTEXIST = "InvalidParameterValue.ModelPropertyNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CallDeviceActionSyncWithContext(ctx context.Context, request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallDeviceActionSync require credential")
    }

    request.SetContext(ctx)
    
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

// ControlDeviceData
// 根据设备产品ID、设备名称，设置控制设备的属性数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
func (c *Client) ControlDeviceData(request *ControlDeviceDataRequest) (response *ControlDeviceDataResponse, err error) {
    return c.ControlDeviceDataWithContext(context.Background(), request)
}

// ControlDeviceData
// 根据设备产品ID、设备名称，设置控制设备的属性数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
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

func NewCreateBatchProductionRequest() (request *CreateBatchProductionRequest) {
    request = &CreateBatchProductionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateBatchProduction")
    
    
    return
}

func NewCreateBatchProductionResponse() (response *CreateBatchProductionResponse) {
    response = &CreateBatchProductionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBatchProduction
// 用于新建批量生产设备
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNOTRELEASED = "FailedOperation.ProductNotReleased"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRLLSYNCBROADCASTDEVICENAMELENGTHEXCEED = "InvalidParameterValue.ErrLLSyncBroadcastDeviceNameLengthExceed"
//  LIMITEXCEEDED_BATCHPRODUCTIONEXCEEDLIMIT = "LimitExceeded.BatchProductionExceedLimit"
//  LIMITEXCEEDED_BATCHPRODUCTIONNULL = "LimitExceeded.BatchProductionNull"
//  RESOURCEINSUFFICIENT_BATCHPRODUCTIONISRUNNING = "ResourceInsufficient.BatchProductionIsRunning"
//  RESOURCENOTFOUND_CANNOTGETFROMURL = "ResourceNotFound.CannotGetFromUrl"
//  RESOURCENOTFOUND_DEVICEDUPKEYEXIST = "ResourceNotFound.DeviceDupKeyExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICEEXCEEDLIMIT = "UnsupportedOperation.DeviceExceedLimit"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CreateBatchProduction(request *CreateBatchProductionRequest) (response *CreateBatchProductionResponse, err error) {
    return c.CreateBatchProductionWithContext(context.Background(), request)
}

// CreateBatchProduction
// 用于新建批量生产设备
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNOTRELEASED = "FailedOperation.ProductNotReleased"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRLLSYNCBROADCASTDEVICENAMELENGTHEXCEED = "InvalidParameterValue.ErrLLSyncBroadcastDeviceNameLengthExceed"
//  LIMITEXCEEDED_BATCHPRODUCTIONEXCEEDLIMIT = "LimitExceeded.BatchProductionExceedLimit"
//  LIMITEXCEEDED_BATCHPRODUCTIONNULL = "LimitExceeded.BatchProductionNull"
//  RESOURCEINSUFFICIENT_BATCHPRODUCTIONISRUNNING = "ResourceInsufficient.BatchProductionIsRunning"
//  RESOURCENOTFOUND_CANNOTGETFROMURL = "ResourceNotFound.CannotGetFromUrl"
//  RESOURCENOTFOUND_DEVICEDUPKEYEXIST = "ResourceNotFound.DeviceDupKeyExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICEEXCEEDLIMIT = "UnsupportedOperation.DeviceExceedLimit"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) CreateBatchProductionWithContext(ctx context.Context, request *CreateBatchProductionRequest) (response *CreateBatchProductionResponse, err error) {
    if request == nil {
        request = NewCreateBatchProductionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchProduction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchProductionResponse()
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

// CreateDevice
// 创建设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNAUTHORIZEDOPERATION_USERLICENSEEXCEEDLIMIT = "UnauthorizedOperation.UserLicenseExceedLimit"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICEEXCEEDLIMIT = "UnsupportedOperation.DeviceExceedLimit"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_VIDEOACCOUNTNOTEXIST = "UnsupportedOperation.VideoAccountNotExist"
//  UNSUPPORTEDOPERATION_VIDEOINSUFFICIENTLICENSES = "UnsupportedOperation.VideoInsufficientLicenses"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    return c.CreateDeviceWithContext(context.Background(), request)
}

// CreateDevice
// 创建设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNAUTHORIZEDOPERATION_USERLICENSEEXCEEDLIMIT = "UnauthorizedOperation.UserLicenseExceedLimit"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICEEXCEEDLIMIT = "UnsupportedOperation.DeviceExceedLimit"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_VIDEOACCOUNTNOTEXIST = "UnsupportedOperation.VideoAccountNotExist"
//  UNSUPPORTEDOPERATION_VIDEOINSUFFICIENTLICENSES = "UnsupportedOperation.VideoInsufficientLicenses"
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

func NewCreateFenceBindRequest() (request *CreateFenceBindRequest) {
    request = &CreateFenceBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateFenceBind")
    
    
    return
}

func NewCreateFenceBindResponse() (response *CreateFenceBindResponse) {
    response = &CreateFenceBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFenceBind
// 创建围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) CreateFenceBind(request *CreateFenceBindRequest) (response *CreateFenceBindResponse, err error) {
    return c.CreateFenceBindWithContext(context.Background(), request)
}

// CreateFenceBind
// 创建围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) CreateFenceBindWithContext(ctx context.Context, request *CreateFenceBindRequest) (response *CreateFenceBindResponse, err error) {
    if request == nil {
        request = NewCreateFenceBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFenceBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFenceBindResponse()
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

// CreateLoRaFrequency
// 创建 LoRa 自定义频点
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LORAFREQPARMSERROR = "InvalidParameterValue.LoRaFreqParmsError"
//  LIMITEXCEEDED_STUDIOLORAFREQEXCEEDLIMIT = "LimitExceeded.StudioLoRaFreqExceedLimit"
//  UNSUPPORTEDOPERATION_LORAFREQDUPKEYEXIST = "UnsupportedOperation.LoRaFreqDupKeyExist"
func (c *Client) CreateLoRaFrequency(request *CreateLoRaFrequencyRequest) (response *CreateLoRaFrequencyResponse, err error) {
    return c.CreateLoRaFrequencyWithContext(context.Background(), request)
}

// CreateLoRaFrequency
// 创建 LoRa 自定义频点
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LORAFREQPARMSERROR = "InvalidParameterValue.LoRaFreqParmsError"
//  LIMITEXCEEDED_STUDIOLORAFREQEXCEEDLIMIT = "LimitExceeded.StudioLoRaFreqExceedLimit"
//  UNSUPPORTEDOPERATION_LORAFREQDUPKEYEXIST = "UnsupportedOperation.LoRaFreqDupKeyExist"
func (c *Client) CreateLoRaFrequencyWithContext(ctx context.Context, request *CreateLoRaFrequencyRequest) (response *CreateLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewCreateLoRaFrequencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoRaFrequency require credential")
    }

    request.SetContext(ctx)
    
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

// CreateLoRaGateway
// 创建新 LoRa 网关设备接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALLORASERVERERROR = "InternalError.InternalLoRaServerError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_GATEWAYDUPKEYEXIST = "ResourceNotFound.GatewayDupKeyExist"
func (c *Client) CreateLoRaGateway(request *CreateLoRaGatewayRequest) (response *CreateLoRaGatewayResponse, err error) {
    return c.CreateLoRaGatewayWithContext(context.Background(), request)
}

// CreateLoRaGateway
// 创建新 LoRa 网关设备接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALLORASERVERERROR = "InternalError.InternalLoRaServerError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_GATEWAYDUPKEYEXIST = "ResourceNotFound.GatewayDupKeyExist"
func (c *Client) CreateLoRaGatewayWithContext(ctx context.Context, request *CreateLoRaGatewayRequest) (response *CreateLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewCreateLoRaGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoRaGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoRaGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePositionFenceRequest() (request *CreatePositionFenceRequest) {
    request = &CreatePositionFenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreatePositionFence")
    
    
    return
}

func NewCreatePositionFenceResponse() (response *CreatePositionFenceResponse) {
    response = &CreatePositionFenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePositionFence
// 创建围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FENCEDUPKEYEXIST = "UnsupportedOperation.FenceDupKeyExist"
func (c *Client) CreatePositionFence(request *CreatePositionFenceRequest) (response *CreatePositionFenceResponse, err error) {
    return c.CreatePositionFenceWithContext(context.Background(), request)
}

// CreatePositionFence
// 创建围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FENCEDUPKEYEXIST = "UnsupportedOperation.FenceDupKeyExist"
func (c *Client) CreatePositionFenceWithContext(ctx context.Context, request *CreatePositionFenceRequest) (response *CreatePositionFenceResponse, err error) {
    if request == nil {
        request = NewCreatePositionFenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePositionFence require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePositionFenceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePositionSpaceRequest() (request *CreatePositionSpaceRequest) {
    request = &CreatePositionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreatePositionSpace")
    
    
    return
}

func NewCreatePositionSpaceResponse() (response *CreatePositionSpaceResponse) {
    response = &CreatePositionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePositionSpace
// 创建位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SPACEDUPKEYEXIST = "UnsupportedOperation.SpaceDupKeyExist"
func (c *Client) CreatePositionSpace(request *CreatePositionSpaceRequest) (response *CreatePositionSpaceResponse, err error) {
    return c.CreatePositionSpaceWithContext(context.Background(), request)
}

// CreatePositionSpace
// 创建位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SPACEDUPKEYEXIST = "UnsupportedOperation.SpaceDupKeyExist"
func (c *Client) CreatePositionSpaceWithContext(ctx context.Context, request *CreatePositionSpaceRequest) (response *CreatePositionSpaceResponse, err error) {
    if request == nil {
        request = NewCreatePositionSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePositionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePositionSpaceResponse()
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

// CreateProject
// 为用户提供新建项目的能力，用于集中管理产品和应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  LIMITEXCEEDED_PROJECTEXCEEDLIMIT = "LimitExceeded.ProjectExceedLimit"
//  RESOURCENOTFOUND_DEVICEDUPKEYEXIST = "ResourceNotFound.DeviceDupKeyExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_MODULENOTEXIST = "ResourceNotFound.ModuleNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_POOLEXISTUNDERPROJECT = "UnsupportedOperation.PoolExistUnderProject"
//  UNSUPPORTEDOPERATION_PROJECTDUPKEYEXIST = "UnsupportedOperation.ProjectDupKeyExist"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 为用户提供新建项目的能力，用于集中管理产品和应用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  LIMITEXCEEDED_PROJECTEXCEEDLIMIT = "LimitExceeded.ProjectExceedLimit"
//  RESOURCENOTFOUND_DEVICEDUPKEYEXIST = "ResourceNotFound.DeviceDupKeyExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_MODULENOTEXIST = "ResourceNotFound.ModuleNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_POOLEXISTUNDERPROJECT = "UnsupportedOperation.PoolExistUnderProject"
//  UNSUPPORTEDOPERATION_PROJECTDUPKEYEXIST = "UnsupportedOperation.ProjectDupKeyExist"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
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

// CreateStudioProduct
// 为用户提供新建产品的能力，用于管理用户的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIDInvalid"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  LIMITEXCEEDED_THINGMODELEXCEEDLIMIT = "LimitExceeded.ThingModelExceedLimit"
//  MISSINGPARAMETER_MODELDEFINEEVENTTYPEERROR = "MissingParameter.ModelDefineEventTypeError"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTDUPKEYEXIST = "UnsupportedOperation.ProductDupKeyExist"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) CreateStudioProduct(request *CreateStudioProductRequest) (response *CreateStudioProductResponse, err error) {
    return c.CreateStudioProductWithContext(context.Background(), request)
}

// CreateStudioProduct
// 为用户提供新建产品的能力，用于管理用户的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIDInvalid"
//  LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  LIMITEXCEEDED_THINGMODELEXCEEDLIMIT = "LimitExceeded.ThingModelExceedLimit"
//  MISSINGPARAMETER_MODELDEFINEEVENTTYPEERROR = "MissingParameter.ModelDefineEventTypeError"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTDUPKEYEXIST = "UnsupportedOperation.ProductDupKeyExist"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) CreateStudioProductWithContext(ctx context.Context, request *CreateStudioProductRequest) (response *CreateStudioProductResponse, err error) {
    if request == nil {
        request = NewCreateStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStudioProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicPolicyRequest() (request *CreateTopicPolicyRequest) {
    request = &CreateTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateTopicPolicy")
    
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  LIMITEXCEEDED_TOPICPOLICYEXCEEDLIMIT = "LimitExceeded.TopicPolicyExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
func (c *Client) CreateTopicPolicy(request *CreateTopicPolicyRequest) (response *CreateTopicPolicyResponse, err error) {
    return c.CreateTopicPolicyWithContext(context.Background(), request)
}

// CreateTopicPolicy
// 本接口（CreateTopicPolicy）用于创建一个Topic 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  LIMITEXCEEDED_TOPICPOLICYEXCEEDLIMIT = "LimitExceeded.TopicPolicyExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateTopicRule")
    
    
    return
}

func NewCreateTopicRuleResponse() (response *CreateTopicRuleResponse) {
    response = &CreateTopicRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopicRule
// 创建规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTopicRule(request *CreateTopicRuleRequest) (response *CreateTopicRuleResponse, err error) {
    return c.CreateTopicRuleWithContext(context.Background(), request)
}

// CreateTopicRule
// 创建规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteDevice")
    
    
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 删除设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDeleteDevicesRequest() (request *DeleteDevicesRequest) {
    request = &DeleteDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteDevices")
    
    
    return
}

func NewDeleteDevicesResponse() (response *DeleteDevicesResponse) {
    response = &DeleteDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevices
// 批量删除设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    return c.DeleteDevicesWithContext(context.Background(), request)
}

// DeleteDevices
// 批量删除设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DeleteDevicesWithContext(ctx context.Context, request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    if request == nil {
        request = NewDeleteDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFenceBindRequest() (request *DeleteFenceBindRequest) {
    request = &DeleteFenceBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteFenceBind")
    
    
    return
}

func NewDeleteFenceBindResponse() (response *DeleteFenceBindResponse) {
    response = &DeleteFenceBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFenceBind
// 删除围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FENCEBINDNOTEXIST = "ResourceNotFound.FenceBindNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DeleteFenceBind(request *DeleteFenceBindRequest) (response *DeleteFenceBindResponse, err error) {
    return c.DeleteFenceBindWithContext(context.Background(), request)
}

// DeleteFenceBind
// 删除围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FENCEBINDNOTEXIST = "ResourceNotFound.FenceBindNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DeleteFenceBindWithContext(ctx context.Context, request *DeleteFenceBindRequest) (response *DeleteFenceBindResponse, err error) {
    if request == nil {
        request = NewDeleteFenceBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFenceBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFenceBindResponse()
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

// DeleteLoRaFrequency
// 提供删除LoRa自定义频点的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NODESEXISTUNDERVPN = "UnsupportedOperation.NodesExistUnderVPN"
//  UNSUPPORTEDOPERATION_STUDIOLORAFREQINUSED = "UnsupportedOperation.StudioLoRaFreqInUsed"
func (c *Client) DeleteLoRaFrequency(request *DeleteLoRaFrequencyRequest) (response *DeleteLoRaFrequencyResponse, err error) {
    return c.DeleteLoRaFrequencyWithContext(context.Background(), request)
}

// DeleteLoRaFrequency
// 提供删除LoRa自定义频点的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NODESEXISTUNDERVPN = "UnsupportedOperation.NodesExistUnderVPN"
//  UNSUPPORTEDOPERATION_STUDIOLORAFREQINUSED = "UnsupportedOperation.StudioLoRaFreqInUsed"
func (c *Client) DeleteLoRaFrequencyWithContext(ctx context.Context, request *DeleteLoRaFrequencyRequest) (response *DeleteLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewDeleteLoRaFrequencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoRaFrequency require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteLoRaGateway
// 删除  LoRa 网关的接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_GATEWAYNOTEXIST = "ResourceNotFound.GatewayNotExist"
func (c *Client) DeleteLoRaGateway(request *DeleteLoRaGatewayRequest) (response *DeleteLoRaGatewayResponse, err error) {
    return c.DeleteLoRaGatewayWithContext(context.Background(), request)
}

// DeleteLoRaGateway
// 删除  LoRa 网关的接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_GATEWAYNOTEXIST = "ResourceNotFound.GatewayNotExist"
func (c *Client) DeleteLoRaGatewayWithContext(ctx context.Context, request *DeleteLoRaGatewayRequest) (response *DeleteLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteLoRaGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoRaGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoRaGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePositionFenceRequest() (request *DeletePositionFenceRequest) {
    request = &DeletePositionFenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeletePositionFence")
    
    
    return
}

func NewDeletePositionFenceResponse() (response *DeletePositionFenceResponse) {
    response = &DeletePositionFenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePositionFence
// 删除围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
//  UNSUPPORTEDOPERATION_BINDSEXISTUNDERFENCE = "UnsupportedOperation.BindsExistUnderFence"
func (c *Client) DeletePositionFence(request *DeletePositionFenceRequest) (response *DeletePositionFenceResponse, err error) {
    return c.DeletePositionFenceWithContext(context.Background(), request)
}

// DeletePositionFence
// 删除围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
//  UNSUPPORTEDOPERATION_BINDSEXISTUNDERFENCE = "UnsupportedOperation.BindsExistUnderFence"
func (c *Client) DeletePositionFenceWithContext(ctx context.Context, request *DeletePositionFenceRequest) (response *DeletePositionFenceResponse, err error) {
    if request == nil {
        request = NewDeletePositionFenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePositionFence require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePositionFenceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePositionSpaceRequest() (request *DeletePositionSpaceRequest) {
    request = &DeletePositionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeletePositionSpace")
    
    
    return
}

func NewDeletePositionSpaceResponse() (response *DeletePositionSpaceResponse) {
    response = &DeletePositionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePositionSpace
// 删除位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_FENCEEXISTUNDERSPACE = "UnsupportedOperation.FenceExistUnderSpace"
func (c *Client) DeletePositionSpace(request *DeletePositionSpaceRequest) (response *DeletePositionSpaceResponse, err error) {
    return c.DeletePositionSpaceWithContext(context.Background(), request)
}

// DeletePositionSpace
// 删除位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_FENCEEXISTUNDERSPACE = "UnsupportedOperation.FenceExistUnderSpace"
func (c *Client) DeletePositionSpaceWithContext(ctx context.Context, request *DeletePositionSpaceRequest) (response *DeletePositionSpaceResponse, err error) {
    if request == nil {
        request = NewDeletePositionSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePositionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePositionSpaceResponse()
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

// DeleteProject
// 提供删除某个项目的能力
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ENABLESAASSERVICEEXISTUNDERPROJECT = "UnsupportedOperation.EnableSaasServiceExistUnderProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_POOLEXISTUNDERPROJECT = "UnsupportedOperation.PoolExistUnderProject"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 提供删除某个项目的能力
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ENABLESAASSERVICEEXISTUNDERPROJECT = "UnsupportedOperation.EnableSaasServiceExistUnderProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_POOLEXISTUNDERPROJECT = "UnsupportedOperation.PoolExistUnderProject"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteStudioProduct
// 提供删除某个项目下产品的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGateWayProduct"
func (c *Client) DeleteStudioProduct(request *DeleteStudioProductRequest) (response *DeleteStudioProductResponse, err error) {
    return c.DeleteStudioProductWithContext(context.Background(), request)
}

// DeleteStudioProduct
// 提供删除某个项目下产品的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGateWayProduct"
func (c *Client) DeleteStudioProductWithContext(ctx context.Context, request *DeleteStudioProductRequest) (response *DeleteStudioProductResponse, err error) {
    if request == nil {
        request = NewDeleteStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStudioProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicPolicyRequest() (request *DeleteTopicPolicyRequest) {
    request = &DeleteTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteTopicPolicy")
    
    
    return
}

func NewDeleteTopicPolicyResponse() (response *DeleteTopicPolicyResponse) {
    response = &DeleteTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopicPolicy
// 本接口（DeleteTopicPolicy）用于删除Topic 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) DeleteTopicPolicy(request *DeleteTopicPolicyRequest) (response *DeleteTopicPolicyResponse, err error) {
    return c.DeleteTopicPolicyWithContext(context.Background(), request)
}

// DeleteTopicPolicy
// 本接口（DeleteTopicPolicy）用于删除Topic 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) DeleteTopicPolicyWithContext(ctx context.Context, request *DeleteTopicPolicyRequest) (response *DeleteTopicPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteTopicPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopicPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicPolicyResponse()
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

// DeleteTopicRule
// 删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTopicRule(request *DeleteTopicRuleRequest) (response *DeleteTopicRuleResponse, err error) {
    return c.DeleteTopicRuleWithContext(context.Background(), request)
}

// DeleteTopicRule
// 删除规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeBatchProductionRequest() (request *DescribeBatchProductionRequest) {
    request = &DescribeBatchProductionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeBatchProduction")
    
    
    return
}

func NewDescribeBatchProductionResponse() (response *DescribeBatchProductionResponse) {
    response = &DescribeBatchProductionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchProduction
// 获取量产详情信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_BATCHPRODUCTIONNOTEXIST = "ResourceNotFound.BatchProductionNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeBatchProduction(request *DescribeBatchProductionRequest) (response *DescribeBatchProductionResponse, err error) {
    return c.DescribeBatchProductionWithContext(context.Background(), request)
}

// DescribeBatchProduction
// 获取量产详情信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_BATCHPRODUCTIONNOTEXIST = "ResourceNotFound.BatchProductionNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeBatchProductionWithContext(ctx context.Context, request *DescribeBatchProductionRequest) (response *DescribeBatchProductionResponse, err error) {
    if request == nil {
        request = NewDescribeBatchProductionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchProduction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchProductionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindedProductsRequest() (request *DescribeBindedProductsRequest) {
    request = &DescribeBindedProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeBindedProducts")
    
    
    return
}

func NewDescribeBindedProductsResponse() (response *DescribeBindedProductsResponse) {
    response = &DescribeBindedProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBindedProducts
// 获取网关产品已经绑定的子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeBindedProducts(request *DescribeBindedProductsRequest) (response *DescribeBindedProductsResponse, err error) {
    return c.DescribeBindedProductsWithContext(context.Background(), request)
}

// DescribeBindedProducts
// 获取网关产品已经绑定的子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeBindedProductsWithContext(ctx context.Context, request *DescribeBindedProductsRequest) (response *DescribeBindedProductsResponse, err error) {
    if request == nil {
        request = NewDescribeBindedProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindedProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindedProductsResponse()
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

// DescribeDevice
// 用于查看某个设备的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    return c.DescribeDeviceWithContext(context.Background(), request)
}

// DescribeDevice
// 用于查看某个设备的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDescribeDeviceBindGatewayRequest() (request *DescribeDeviceBindGatewayRequest) {
    request = &DescribeDeviceBindGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceBindGateway")
    
    
    return
}

func NewDescribeDeviceBindGatewayResponse() (response *DescribeDeviceBindGatewayResponse) {
    response = &DescribeDeviceBindGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceBindGateway
// 查询设备绑定的网关设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEHASNOTBINDGATEWAY = "InvalidParameterValue.DeviceHasNotBindGateway"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceBindGateway(request *DescribeDeviceBindGatewayRequest) (response *DescribeDeviceBindGatewayResponse, err error) {
    return c.DescribeDeviceBindGatewayWithContext(context.Background(), request)
}

// DescribeDeviceBindGateway
// 查询设备绑定的网关设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEHASNOTBINDGATEWAY = "InvalidParameterValue.DeviceHasNotBindGateway"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceBindGatewayWithContext(ctx context.Context, request *DescribeDeviceBindGatewayRequest) (response *DescribeDeviceBindGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceBindGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceBindGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceBindGatewayResponse()
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

// DescribeDeviceData
// 根据设备产品ID、设备名称，获取设备上报的属性数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeDeviceData(request *DescribeDeviceDataRequest) (response *DescribeDeviceDataResponse, err error) {
    return c.DescribeDeviceDataWithContext(context.Background(), request)
}

// DescribeDeviceData
// 根据设备产品ID、设备名称，获取设备上报的属性数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceDataHistory")
    
    
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
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeDeviceDataHistory(request *DescribeDeviceDataHistoryRequest) (response *DescribeDeviceDataHistoryResponse, err error) {
    return c.DescribeDeviceDataHistoryWithContext(context.Background(), request)
}

// DescribeDeviceDataHistory
// 获取设备在指定时间范围内上报的历史数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDescribeDeviceLocationSolveRequest() (request *DescribeDeviceLocationSolveRequest) {
    request = &DescribeDeviceLocationSolveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceLocationSolve")
    
    
    return
}

func NewDescribeDeviceLocationSolveResponse() (response *DescribeDeviceLocationSolveResponse) {
    response = &DescribeDeviceLocationSolveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceLocationSolve
// 获取实时位置解析
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MODULENOTEXIST = "ResourceNotFound.ModuleNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceLocationSolve(request *DescribeDeviceLocationSolveRequest) (response *DescribeDeviceLocationSolveResponse, err error) {
    return c.DescribeDeviceLocationSolveWithContext(context.Background(), request)
}

// DescribeDeviceLocationSolve
// 获取实时位置解析
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MODULENOTEXIST = "ResourceNotFound.ModuleNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceLocationSolveWithContext(ctx context.Context, request *DescribeDeviceLocationSolveRequest) (response *DescribeDeviceLocationSolveResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceLocationSolveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceLocationSolve require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceLocationSolveResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePositionListRequest() (request *DescribeDevicePositionListRequest) {
    request = &DescribeDevicePositionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDevicePositionList")
    
    
    return
}

func NewDescribeDevicePositionListResponse() (response *DescribeDevicePositionListResponse) {
    response = &DescribeDevicePositionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePositionList
// 获取设备位置列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
func (c *Client) DescribeDevicePositionList(request *DescribeDevicePositionListRequest) (response *DescribeDevicePositionListResponse, err error) {
    return c.DescribeDevicePositionListWithContext(context.Background(), request)
}

// DescribeDevicePositionList
// 获取设备位置列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
func (c *Client) DescribeDevicePositionListWithContext(ctx context.Context, request *DescribeDevicePositionListRequest) (response *DescribeDevicePositionListResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePositionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePositionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePositionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFenceBindListRequest() (request *DescribeFenceBindListRequest) {
    request = &DescribeFenceBindListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFenceBindList")
    
    
    return
}

func NewDescribeFenceBindListResponse() (response *DescribeFenceBindListResponse) {
    response = &DescribeFenceBindListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFenceBindList
// 获取围栏绑定信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DescribeFenceBindList(request *DescribeFenceBindListRequest) (response *DescribeFenceBindListResponse, err error) {
    return c.DescribeFenceBindListWithContext(context.Background(), request)
}

// DescribeFenceBindList
// 获取围栏绑定信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DescribeFenceBindListWithContext(ctx context.Context, request *DescribeFenceBindListRequest) (response *DescribeFenceBindListResponse, err error) {
    if request == nil {
        request = NewDescribeFenceBindListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFenceBindList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFenceBindListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFenceEventListRequest() (request *DescribeFenceEventListRequest) {
    request = &DescribeFenceEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFenceEventList")
    
    
    return
}

func NewDescribeFenceEventListResponse() (response *DescribeFenceEventListResponse) {
    response = &DescribeFenceEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFenceEventList
// 获取围栏告警事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DescribeFenceEventList(request *DescribeFenceEventListRequest) (response *DescribeFenceEventListResponse, err error) {
    return c.DescribeFenceEventListWithContext(context.Background(), request)
}

// DescribeFenceEventList
// 获取围栏告警事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DescribeFenceEventListWithContext(ctx context.Context, request *DescribeFenceEventListRequest) (response *DescribeFenceEventListResponse, err error) {
    if request == nil {
        request = NewDescribeFenceEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFenceEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFenceEventListResponse()
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

// DescribeFirmwareTask
// 查询固件升级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
func (c *Client) DescribeFirmwareTask(request *DescribeFirmwareTaskRequest) (response *DescribeFirmwareTaskResponse, err error) {
    return c.DescribeFirmwareTaskWithContext(context.Background(), request)
}

// DescribeFirmwareTask
// 查询固件升级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
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

func NewDescribeGatewayBindDevicesRequest() (request *DescribeGatewayBindDevicesRequest) {
    request = &DescribeGatewayBindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeGatewayBindDevices")
    
    
    return
}

func NewDescribeGatewayBindDevicesResponse() (response *DescribeGatewayBindDevicesResponse) {
    response = &DescribeGatewayBindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayBindDevices
// 获取网关绑定的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeGatewayBindDevices(request *DescribeGatewayBindDevicesRequest) (response *DescribeGatewayBindDevicesResponse, err error) {
    return c.DescribeGatewayBindDevicesWithContext(context.Background(), request)
}

// DescribeGatewayBindDevices
// 获取网关绑定的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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

func NewDescribeGatewaySubDeviceListRequest() (request *DescribeGatewaySubDeviceListRequest) {
    request = &DescribeGatewaySubDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeGatewaySubDeviceList")
    
    
    return
}

func NewDescribeGatewaySubDeviceListResponse() (response *DescribeGatewaySubDeviceListResponse) {
    response = &DescribeGatewaySubDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewaySubDeviceList
// 查询绑定到家庭的网关设备的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeGatewaySubDeviceList(request *DescribeGatewaySubDeviceListRequest) (response *DescribeGatewaySubDeviceListResponse, err error) {
    return c.DescribeGatewaySubDeviceListWithContext(context.Background(), request)
}

// DescribeGatewaySubDeviceList
// 查询绑定到家庭的网关设备的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeGatewaySubDeviceListWithContext(ctx context.Context, request *DescribeGatewaySubDeviceListRequest) (response *DescribeGatewaySubDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeGatewaySubDeviceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewaySubDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewaySubDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewaySubProductsRequest() (request *DescribeGatewaySubProductsRequest) {
    request = &DescribeGatewaySubProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeGatewaySubProducts")
    
    
    return
}

func NewDescribeGatewaySubProductsResponse() (response *DescribeGatewaySubProductsResponse) {
    response = &DescribeGatewaySubProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewaySubProducts
// 用于获取网关可绑定或解绑的子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_STAFFPOOLDUPNAMEEXIST = "UnsupportedOperation.StaffPoolDupNameExist"
func (c *Client) DescribeGatewaySubProducts(request *DescribeGatewaySubProductsRequest) (response *DescribeGatewaySubProductsResponse, err error) {
    return c.DescribeGatewaySubProductsWithContext(context.Background(), request)
}

// DescribeGatewaySubProducts
// 用于获取网关可绑定或解绑的子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_STAFFPOOLDUPNAMEEXIST = "UnsupportedOperation.StaffPoolDupNameExist"
func (c *Client) DescribeGatewaySubProductsWithContext(ctx context.Context, request *DescribeGatewaySubProductsRequest) (response *DescribeGatewaySubProductsResponse, err error) {
    if request == nil {
        request = NewDescribeGatewaySubProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewaySubProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewaySubProductsResponse()
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

// DescribeLoRaFrequency
// 提供查询LoRa自定义频点详情的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoRaFrequency(request *DescribeLoRaFrequencyRequest) (response *DescribeLoRaFrequencyResponse, err error) {
    return c.DescribeLoRaFrequencyWithContext(context.Background(), request)
}

// DescribeLoRaFrequency
// 提供查询LoRa自定义频点详情的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoRaFrequencyWithContext(ctx context.Context, request *DescribeLoRaFrequencyRequest) (response *DescribeLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewDescribeLoRaFrequencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoRaFrequency require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeModelDefinition
// 查询产品配置的数据模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_MODELDEFINEDUPID = "InvalidParameterValue.ModelDefineDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORMODEL = "InvalidParameterValue.ModelDefineErrorModel"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORTYPE = "InvalidParameterValue.ModelDefineErrorType"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSDUPID = "InvalidParameterValue.ModelDefineEventParamsDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSEXCEEDLIMIT = "InvalidParameterValue.ModelDefineEventParamsExceedLimit"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeModelDefinition(request *DescribeModelDefinitionRequest) (response *DescribeModelDefinitionResponse, err error) {
    return c.DescribeModelDefinitionWithContext(context.Background(), request)
}

// DescribeModelDefinition
// 查询产品配置的数据模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_MODELDEFINEDUPID = "InvalidParameterValue.ModelDefineDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORMODEL = "InvalidParameterValue.ModelDefineErrorModel"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORTYPE = "InvalidParameterValue.ModelDefineErrorType"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSDUPID = "InvalidParameterValue.ModelDefineEventParamsDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSEXCEEDLIMIT = "InvalidParameterValue.ModelDefineEventParamsExceedLimit"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDescribePositionFenceListRequest() (request *DescribePositionFenceListRequest) {
    request = &DescribePositionFenceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribePositionFenceList")
    
    
    return
}

func NewDescribePositionFenceListResponse() (response *DescribePositionFenceListResponse) {
    response = &DescribePositionFenceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePositionFenceList
// 获取围栏列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePositionFenceList(request *DescribePositionFenceListRequest) (response *DescribePositionFenceListResponse, err error) {
    return c.DescribePositionFenceListWithContext(context.Background(), request)
}

// DescribePositionFenceList
// 获取围栏列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePositionFenceListWithContext(ctx context.Context, request *DescribePositionFenceListRequest) (response *DescribePositionFenceListResponse, err error) {
    if request == nil {
        request = NewDescribePositionFenceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePositionFenceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePositionFenceListResponse()
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

// DescribeProject
// 查询项目详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    return c.DescribeProjectWithContext(context.Background(), request)
}

// DescribeProject
// 查询项目详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceFenceEventListRequest() (request *DescribeSpaceFenceEventListRequest) {
    request = &DescribeSpaceFenceEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeSpaceFenceEventList")
    
    
    return
}

func NewDescribeSpaceFenceEventListResponse() (response *DescribeSpaceFenceEventListResponse) {
    response = &DescribeSpaceFenceEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSpaceFenceEventList
// 获取位置空间中围栏告警事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSpaceFenceEventList(request *DescribeSpaceFenceEventListRequest) (response *DescribeSpaceFenceEventListResponse, err error) {
    return c.DescribeSpaceFenceEventListWithContext(context.Background(), request)
}

// DescribeSpaceFenceEventList
// 获取位置空间中围栏告警事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSpaceFenceEventListWithContext(ctx context.Context, request *DescribeSpaceFenceEventListRequest) (response *DescribeSpaceFenceEventListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceFenceEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceFenceEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceFenceEventListResponse()
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

// DescribeStudioProduct
// 提供查看产品详细信息的能力，包括产品的ID、数据协议、认证类型等重要参数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeStudioProduct(request *DescribeStudioProductRequest) (response *DescribeStudioProductResponse, err error) {
    return c.DescribeStudioProductWithContext(context.Background(), request)
}

// DescribeStudioProduct
// 提供查看产品详细信息的能力，包括产品的ID、数据协议、认证类型等重要参数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeStudioProductWithContext(ctx context.Context, request *DescribeStudioProductRequest) (response *DescribeStudioProductResponse, err error) {
    if request == nil {
        request = NewDescribeStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStudioProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicPolicyRequest() (request *DescribeTopicPolicyRequest) {
    request = &DescribeTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeTopicPolicy")
    
    
    return
}

func NewDescribeTopicPolicyResponse() (response *DescribeTopicPolicyResponse) {
    response = &DescribeTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicPolicy
// 本接口（DescribeTopicPolicy）用于查看Topic详细信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) DescribeTopicPolicy(request *DescribeTopicPolicyRequest) (response *DescribeTopicPolicyResponse, err error) {
    return c.DescribeTopicPolicyWithContext(context.Background(), request)
}

// DescribeTopicPolicy
// 本接口（DescribeTopicPolicy）用于查看Topic详细信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) DescribeTopicPolicyWithContext(ctx context.Context, request *DescribeTopicPolicyRequest) (response *DescribeTopicPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeTopicPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicPolicyResponse()
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

// DescribeTopicRule
// 获取规则信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicRule(request *DescribeTopicRuleRequest) (response *DescribeTopicRuleResponse, err error) {
    return c.DescribeTopicRuleWithContext(context.Background(), request)
}

// DescribeTopicRule
// 获取规则信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicRuleWithContext(ctx context.Context, request *DescribeTopicRuleRequest) (response *DescribeTopicRuleResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDirectBindDeviceInFamilyRequest() (request *DirectBindDeviceInFamilyRequest) {
    request = &DirectBindDeviceInFamilyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DirectBindDeviceInFamily")
    
    
    return
}

func NewDirectBindDeviceInFamilyResponse() (response *DirectBindDeviceInFamilyResponse) {
    response = &DirectBindDeviceInFamilyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DirectBindDeviceInFamily
// 直接绑定设备和家庭
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.APPNoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOFAMILY = "UnauthorizedOperation.NoPermissionToFamily"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICETYPE = "UnsupportedOperation.DeviceType"
func (c *Client) DirectBindDeviceInFamily(request *DirectBindDeviceInFamilyRequest) (response *DirectBindDeviceInFamilyResponse, err error) {
    return c.DirectBindDeviceInFamilyWithContext(context.Background(), request)
}

// DirectBindDeviceInFamily
// 直接绑定设备和家庭
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.APPNoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOFAMILY = "UnauthorizedOperation.NoPermissionToFamily"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICETYPE = "UnsupportedOperation.DeviceType"
func (c *Client) DirectBindDeviceInFamilyWithContext(ctx context.Context, request *DirectBindDeviceInFamilyRequest) (response *DirectBindDeviceInFamilyResponse, err error) {
    if request == nil {
        request = NewDirectBindDeviceInFamilyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DirectBindDeviceInFamily require credential")
    }

    request.SetContext(ctx)
    
    response = NewDirectBindDeviceInFamilyResponse()
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

// DisableTopicRule
// 禁用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYDISABLED = "FailedOperation.RuleAlreadyDisabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableTopicRule(request *DisableTopicRuleRequest) (response *DisableTopicRuleResponse, err error) {
    return c.DisableTopicRuleWithContext(context.Background(), request)
}

// DisableTopicRule
// 禁用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYDISABLED = "FailedOperation.RuleAlreadyDisabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

// EnableTopicRule
// 启用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYENABLED = "FailedOperation.RuleAlreadyEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableTopicRule(request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    return c.EnableTopicRuleWithContext(context.Background(), request)
}

// EnableTopicRule
// 启用规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYENABLED = "FailedOperation.RuleAlreadyEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewGenSingleDeviceSignatureOfPublicRequest() (request *GenSingleDeviceSignatureOfPublicRequest) {
    request = &GenSingleDeviceSignatureOfPublicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GenSingleDeviceSignatureOfPublic")
    
    
    return
}

func NewGenSingleDeviceSignatureOfPublicResponse() (response *GenSingleDeviceSignatureOfPublicResponse) {
    response = &GenSingleDeviceSignatureOfPublicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenSingleDeviceSignatureOfPublic
// 无
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GenSingleDeviceSignatureOfPublic(request *GenSingleDeviceSignatureOfPublicRequest) (response *GenSingleDeviceSignatureOfPublicResponse, err error) {
    return c.GenSingleDeviceSignatureOfPublicWithContext(context.Background(), request)
}

// GenSingleDeviceSignatureOfPublic
// 无
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GenSingleDeviceSignatureOfPublicWithContext(ctx context.Context, request *GenSingleDeviceSignatureOfPublicRequest) (response *GenSingleDeviceSignatureOfPublicResponse, err error) {
    if request == nil {
        request = NewGenSingleDeviceSignatureOfPublicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenSingleDeviceSignatureOfPublic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenSingleDeviceSignatureOfPublicResponse()
    err = c.Send(request, response)
    return
}

func NewGetBatchProductionsListRequest() (request *GetBatchProductionsListRequest) {
    request = &GetBatchProductionsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetBatchProductionsList")
    
    
    return
}

func NewGetBatchProductionsListResponse() (response *GetBatchProductionsListResponse) {
    response = &GetBatchProductionsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetBatchProductionsList
// 列出量产数据列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
func (c *Client) GetBatchProductionsList(request *GetBatchProductionsListRequest) (response *GetBatchProductionsListResponse, err error) {
    return c.GetBatchProductionsListWithContext(context.Background(), request)
}

// GetBatchProductionsList
// 列出量产数据列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
func (c *Client) GetBatchProductionsListWithContext(ctx context.Context, request *GetBatchProductionsListRequest) (response *GetBatchProductionsListResponse, err error) {
    if request == nil {
        request = NewGetBatchProductionsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetBatchProductionsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetBatchProductionsListResponse()
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

// GetCOSURL
// 本接口（GetCOSURL）用于获取固件COS存储的上传请求URL地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetCOSURL(request *GetCOSURLRequest) (response *GetCOSURLResponse, err error) {
    return c.GetCOSURLWithContext(context.Background(), request)
}

// GetCOSURL
// 本接口（GetCOSURL）用于获取固件COS存储的上传请求URL地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

// GetDeviceList
// 用于查询某个产品下的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetDeviceList(request *GetDeviceListRequest) (response *GetDeviceListResponse, err error) {
    return c.GetDeviceListWithContext(context.Background(), request)
}

// GetDeviceList
// 用于查询某个产品下的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetDeviceListWithContext(ctx context.Context, request *GetDeviceListRequest) (response *GetDeviceListResponse, err error) {
    if request == nil {
        request = NewGetDeviceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceLocationHistoryRequest() (request *GetDeviceLocationHistoryRequest) {
    request = &GetDeviceLocationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetDeviceLocationHistory")
    
    
    return
}

func NewGetDeviceLocationHistoryResponse() (response *GetDeviceLocationHistoryResponse) {
    response = &GetDeviceLocationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceLocationHistory
// 获取设备历史位置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetDeviceLocationHistory(request *GetDeviceLocationHistoryRequest) (response *GetDeviceLocationHistoryResponse, err error) {
    return c.GetDeviceLocationHistoryWithContext(context.Background(), request)
}

// GetDeviceLocationHistory
// 获取设备历史位置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_STARTTIMELATERENDTIME = "InvalidParameterValue.StartTimeLaterEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetDeviceLocationHistoryWithContext(ctx context.Context, request *GetDeviceLocationHistoryRequest) (response *GetDeviceLocationHistoryResponse, err error) {
    if request == nil {
        request = NewGetDeviceLocationHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceLocationHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceLocationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetFamilyDeviceUserListRequest() (request *GetFamilyDeviceUserListRequest) {
    request = &GetFamilyDeviceUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetFamilyDeviceUserList")
    
    
    return
}

func NewGetFamilyDeviceUserListResponse() (response *GetFamilyDeviceUserListResponse) {
    response = &GetFamilyDeviceUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFamilyDeviceUserList
// 用于获取设备绑定的用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetFamilyDeviceUserList(request *GetFamilyDeviceUserListRequest) (response *GetFamilyDeviceUserListResponse, err error) {
    return c.GetFamilyDeviceUserListWithContext(context.Background(), request)
}

// GetFamilyDeviceUserList
// 用于获取设备绑定的用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetFamilyDeviceUserListWithContext(ctx context.Context, request *GetFamilyDeviceUserListRequest) (response *GetFamilyDeviceUserListResponse, err error) {
    if request == nil {
        request = NewGetFamilyDeviceUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFamilyDeviceUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFamilyDeviceUserListResponse()
    err = c.Send(request, response)
    return
}

func NewGetGatewaySubDeviceListRequest() (request *GetGatewaySubDeviceListRequest) {
    request = &GetGatewaySubDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetGatewaySubDeviceList")
    
    
    return
}

func NewGetGatewaySubDeviceListResponse() (response *GetGatewaySubDeviceListResponse) {
    response = &GetGatewaySubDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGatewaySubDeviceList
// 获取指定网关设备的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetGatewaySubDeviceList(request *GetGatewaySubDeviceListRequest) (response *GetGatewaySubDeviceListResponse, err error) {
    return c.GetGatewaySubDeviceListWithContext(context.Background(), request)
}

// GetGatewaySubDeviceList
// 获取指定网关设备的子设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTBIND = "ResourceNotFound.DeviceNotBind"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) GetGatewaySubDeviceListWithContext(ctx context.Context, request *GetGatewaySubDeviceListRequest) (response *GetGatewaySubDeviceListResponse, err error) {
    if request == nil {
        request = NewGetGatewaySubDeviceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGatewaySubDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGatewaySubDeviceListResponse()
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

// GetLoRaGatewayList
// 获取 LoRa 网关列表接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetLoRaGatewayList(request *GetLoRaGatewayListRequest) (response *GetLoRaGatewayListResponse, err error) {
    return c.GetLoRaGatewayListWithContext(context.Background(), request)
}

// GetLoRaGatewayList
// 获取 LoRa 网关列表接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetLoRaGatewayListWithContext(ctx context.Context, request *GetLoRaGatewayListRequest) (response *GetLoRaGatewayListResponse, err error) {
    if request == nil {
        request = NewGetLoRaGatewayListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLoRaGatewayList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLoRaGatewayListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPositionSpaceListRequest() (request *GetPositionSpaceListRequest) {
    request = &GetPositionSpaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetPositionSpaceList")
    
    
    return
}

func NewGetPositionSpaceListResponse() (response *GetPositionSpaceListResponse) {
    response = &GetPositionSpaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPositionSpaceList
// 获取位置空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetPositionSpaceList(request *GetPositionSpaceListRequest) (response *GetPositionSpaceListResponse, err error) {
    return c.GetPositionSpaceListWithContext(context.Background(), request)
}

// GetPositionSpaceList
// 获取位置空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetPositionSpaceListWithContext(ctx context.Context, request *GetPositionSpaceListRequest) (response *GetPositionSpaceListResponse, err error) {
    if request == nil {
        request = NewGetPositionSpaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPositionSpaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPositionSpaceListResponse()
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

// GetProjectList
// 提供查询用户所创建的项目列表查询功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) GetProjectList(request *GetProjectListRequest) (response *GetProjectListResponse, err error) {
    return c.GetProjectListWithContext(context.Background(), request)
}

// GetProjectList
// 提供查询用户所创建的项目列表查询功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) GetProjectListWithContext(ctx context.Context, request *GetProjectListRequest) (response *GetProjectListResponse, err error) {
    if request == nil {
        request = NewGetProjectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProjectList require credential")
    }

    request.SetContext(ctx)
    
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

// GetStudioProductList
// 提供查询某个项目下所有产品信息的能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetStudioProductList(request *GetStudioProductListRequest) (response *GetStudioProductListResponse, err error) {
    return c.GetStudioProductListWithContext(context.Background(), request)
}

// GetStudioProductList
// 提供查询某个项目下所有产品信息的能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetStudioProductListWithContext(ctx context.Context, request *GetStudioProductListRequest) (response *GetStudioProductListResponse, err error) {
    if request == nil {
        request = NewGetStudioProductListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetStudioProductList require credential")
    }

    request.SetContext(ctx)
    
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

// GetTopicRuleList
// 获取规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTopicRuleList(request *GetTopicRuleListRequest) (response *GetTopicRuleListResponse, err error) {
    return c.GetTopicRuleListWithContext(context.Background(), request)
}

// GetTopicRuleList
// 获取规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTopicRuleListWithContext(ctx context.Context, request *GetTopicRuleListRequest) (response *GetTopicRuleListResponse, err error) {
    if request == nil {
        request = NewGetTopicRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTopicRuleList require credential")
    }

    request.SetContext(ctx)
    
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

// ListEventHistory
// 获取设备的历史事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.APPNoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ListEventHistory(request *ListEventHistoryRequest) (response *ListEventHistoryResponse, err error) {
    return c.ListEventHistoryWithContext(context.Background(), request)
}

// ListEventHistory
// 获取设备的历史事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.APPNoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ListEventHistoryWithContext(ctx context.Context, request *ListEventHistoryRequest) (response *ListEventHistoryResponse, err error) {
    if request == nil {
        request = NewListEventHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEventHistory require credential")
    }

    request.SetContext(ctx)
    
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

// ListFirmwares
// 本接口（ListFirmwares）用于获取固件列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ListFirmwares(request *ListFirmwaresRequest) (response *ListFirmwaresResponse, err error) {
    return c.ListFirmwaresWithContext(context.Background(), request)
}

// ListFirmwares
// 本接口（ListFirmwares）用于获取固件列表 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewListTopicPolicyRequest() (request *ListTopicPolicyRequest) {
    request = &ListTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ListTopicPolicy")
    
    
    return
}

func NewListTopicPolicyResponse() (response *ListTopicPolicyResponse) {
    response = &ListTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopicPolicy
// 本接口（ListTopicPolicy）用于获取Topic列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ListTopicPolicy(request *ListTopicPolicyRequest) (response *ListTopicPolicyResponse, err error) {
    return c.ListTopicPolicyWithContext(context.Background(), request)
}

// ListTopicPolicy
// 本接口（ListTopicPolicy）用于获取Topic列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ListTopicPolicyWithContext(ctx context.Context, request *ListTopicPolicyRequest) (response *ListTopicPolicyResponse, err error) {
    if request == nil {
        request = NewListTopicPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopicPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopicPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFenceBindRequest() (request *ModifyFenceBindRequest) {
    request = &ModifyFenceBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyFenceBind")
    
    
    return
}

func NewModifyFenceBindResponse() (response *ModifyFenceBindResponse) {
    response = &ModifyFenceBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFenceBind
// 更新围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) ModifyFenceBind(request *ModifyFenceBindRequest) (response *ModifyFenceBindResponse, err error) {
    return c.ModifyFenceBindWithContext(context.Background(), request)
}

// ModifyFenceBind
// 更新围栏绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) ModifyFenceBindWithContext(ctx context.Context, request *ModifyFenceBindRequest) (response *ModifyFenceBindResponse, err error) {
    if request == nil {
        request = NewModifyFenceBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFenceBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFenceBindResponse()
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

// ModifyLoRaFrequency
// 修改LoRa自定义频点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LORAFREQPARMSERROR = "InvalidParameterValue.LoRaFreqParmsError"
//  INVALIDPARAMETERVALUE_VPNPARMSERROR = "InvalidParameterValue.VPNParmsError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_LORAFREQDUPKEYEXIST = "UnsupportedOperation.LoRaFreqDupKeyExist"
//  UNSUPPORTEDOPERATION_VPNDUPKEYEXIST = "UnsupportedOperation.VPNDupKeyExist"
func (c *Client) ModifyLoRaFrequency(request *ModifyLoRaFrequencyRequest) (response *ModifyLoRaFrequencyResponse, err error) {
    return c.ModifyLoRaFrequencyWithContext(context.Background(), request)
}

// ModifyLoRaFrequency
// 修改LoRa自定义频点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LORAFREQPARMSERROR = "InvalidParameterValue.LoRaFreqParmsError"
//  INVALIDPARAMETERVALUE_VPNPARMSERROR = "InvalidParameterValue.VPNParmsError"
//  RESOURCENOTFOUND_STUDIOLORAFREQNOTEXIST = "ResourceNotFound.StudioLoRaFreqNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_LORAFREQDUPKEYEXIST = "UnsupportedOperation.LoRaFreqDupKeyExist"
//  UNSUPPORTEDOPERATION_VPNDUPKEYEXIST = "UnsupportedOperation.VPNDupKeyExist"
func (c *Client) ModifyLoRaFrequencyWithContext(ctx context.Context, request *ModifyLoRaFrequencyRequest) (response *ModifyLoRaFrequencyResponse, err error) {
    if request == nil {
        request = NewModifyLoRaFrequencyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoRaFrequency require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyLoRaGateway
// 修改 LoRa 网关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALLORASERVERERROR = "InternalError.InternalLoRaServerError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_GATEWAYNOTEXIST = "ResourceNotFound.GatewayNotExist"
func (c *Client) ModifyLoRaGateway(request *ModifyLoRaGatewayRequest) (response *ModifyLoRaGatewayResponse, err error) {
    return c.ModifyLoRaGatewayWithContext(context.Background(), request)
}

// ModifyLoRaGateway
// 修改 LoRa 网关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALLORASERVERERROR = "InternalError.InternalLoRaServerError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_GATEWAYNOTEXIST = "ResourceNotFound.GatewayNotExist"
func (c *Client) ModifyLoRaGatewayWithContext(ctx context.Context, request *ModifyLoRaGatewayRequest) (response *ModifyLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewModifyLoRaGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoRaGateway require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyModelDefinition
// 提供修改产品的数据模板的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPDESCRIPTIONTOOLONG = "InvalidParameterValue.AppDescriptionTooLong"
//  INVALIDPARAMETERVALUE_APPEXISTS = "InvalidParameterValue.AppExists"
//  INVALIDPARAMETERVALUE_APPNAMETOOLONG = "InvalidParameterValue.AppNameTooLong"
//  INVALIDPARAMETERVALUE_APPNOPERMISSION = "InvalidParameterValue.AppNoPermission"
//  INVALIDPARAMETERVALUE_APPNOTEXISTS = "InvalidParameterValue.AppNotExists"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_ERRORTASKNOTEXIST = "InvalidParameterValue.ErrorTaskNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_MODELDEFINEDUPID = "InvalidParameterValue.ModelDefineDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORMODEL = "InvalidParameterValue.ModelDefineErrorModel"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORTYPE = "InvalidParameterValue.ModelDefineErrorType"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSDUPID = "InvalidParameterValue.ModelDefineEventParamsDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSEXCEEDLIMIT = "InvalidParameterValue.ModelDefineEventParamsExceedLimit"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPERROR = "InvalidParameterValue.ModelDefineEventPropError"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTTYPEERROR = "InvalidParameterValue.ModelDefineEventTypeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_MSGCONTENTINVALID = "InvalidParameterValue.MsgContentInvalid"
//  INVALIDPARAMETERVALUE_MSGLEVELINVALID = "InvalidParameterValue.MsgLevelInvalid"
//  INVALIDPARAMETERVALUE_MSGTITLEINVALID = "InvalidParameterValue.MsgTitleInvalid"
//  INVALIDPARAMETERVALUE_MSGTYPEINVALID = "InvalidParameterValue.MsgTypeInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIDInvalid"
//  LIMITEXCEEDED_APPLICATIONEXCEEDLIMIT = "LimitExceeded.ApplicationExceedLimit"
//  LIMITEXCEEDED_PROJECTEXCEEDLIMIT = "LimitExceeded.ProjectExceedLimit"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  LIMITEXCEEDED_THINGMODELEXCEEDLIMIT = "LimitExceeded.ThingModelExceedLimit"
//  MISSINGPARAMETER_MODELDEFINEEVENTTYPEERROR = "MissingParameter.ModelDefineEventTypeError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) ModifyModelDefinition(request *ModifyModelDefinitionRequest) (response *ModifyModelDefinitionResponse, err error) {
    return c.ModifyModelDefinitionWithContext(context.Background(), request)
}

// ModifyModelDefinition
// 提供修改产品的数据模板的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPDESCRIPTIONTOOLONG = "InvalidParameterValue.AppDescriptionTooLong"
//  INVALIDPARAMETERVALUE_APPEXISTS = "InvalidParameterValue.AppExists"
//  INVALIDPARAMETERVALUE_APPNAMETOOLONG = "InvalidParameterValue.AppNameTooLong"
//  INVALIDPARAMETERVALUE_APPNOPERMISSION = "InvalidParameterValue.AppNoPermission"
//  INVALIDPARAMETERVALUE_APPNOTEXISTS = "InvalidParameterValue.AppNotExists"
//  INVALIDPARAMETERVALUE_DEVICENAMEINVALID = "InvalidParameterValue.DeviceNameInvalid"
//  INVALIDPARAMETERVALUE_ERRORTASKNOTEXIST = "InvalidParameterValue.ErrorTaskNotExist"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_MODELDEFINEDUPID = "InvalidParameterValue.ModelDefineDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORMODEL = "InvalidParameterValue.ModelDefineErrorModel"
//  INVALIDPARAMETERVALUE_MODELDEFINEERRORTYPE = "InvalidParameterValue.ModelDefineErrorType"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSDUPID = "InvalidParameterValue.ModelDefineEventParamsDupID"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPARAMSEXCEEDLIMIT = "InvalidParameterValue.ModelDefineEventParamsExceedLimit"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPERROR = "InvalidParameterValue.ModelDefineEventPropError"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTTYPEERROR = "InvalidParameterValue.ModelDefineEventTypeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPBOOLMAPPINGERROR = "InvalidParameterValue.ModelDefinePropBoolMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPENUMMAPPINGERROR = "InvalidParameterValue.ModelDefinePropEnumMappingError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEERROR = "InvalidParameterValue.ModelDefinePropRangeError"
//  INVALIDPARAMETERVALUE_MODELDEFINEPROPRANGEOVERFLOW = "InvalidParameterValue.ModelDefinePropRangeOverflow"
//  INVALIDPARAMETERVALUE_MSGCONTENTINVALID = "InvalidParameterValue.MsgContentInvalid"
//  INVALIDPARAMETERVALUE_MSGLEVELINVALID = "InvalidParameterValue.MsgLevelInvalid"
//  INVALIDPARAMETERVALUE_MSGTITLEINVALID = "InvalidParameterValue.MsgTitleInvalid"
//  INVALIDPARAMETERVALUE_MSGTYPEINVALID = "InvalidParameterValue.MsgTypeInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIDInvalid"
//  LIMITEXCEEDED_APPLICATIONEXCEEDLIMIT = "LimitExceeded.ApplicationExceedLimit"
//  LIMITEXCEEDED_PROJECTEXCEEDLIMIT = "LimitExceeded.ProjectExceedLimit"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  LIMITEXCEEDED_THINGMODELEXCEEDLIMIT = "LimitExceeded.ThingModelExceedLimit"
//  MISSINGPARAMETER_MODELDEFINEEVENTTYPEERROR = "MissingParameter.ModelDefineEventTypeError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
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

func NewModifyPositionFenceRequest() (request *ModifyPositionFenceRequest) {
    request = &ModifyPositionFenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyPositionFence")
    
    
    return
}

func NewModifyPositionFenceResponse() (response *ModifyPositionFenceResponse) {
    response = &ModifyPositionFenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPositionFence
// 更新围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) ModifyPositionFence(request *ModifyPositionFenceRequest) (response *ModifyPositionFenceResponse, err error) {
    return c.ModifyPositionFenceWithContext(context.Background(), request)
}

// ModifyPositionFence
// 更新围栏
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) ModifyPositionFenceWithContext(ctx context.Context, request *ModifyPositionFenceRequest) (response *ModifyPositionFenceResponse, err error) {
    if request == nil {
        request = NewModifyPositionFenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPositionFence require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPositionFenceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPositionSpaceRequest() (request *ModifyPositionSpaceRequest) {
    request = &ModifyPositionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyPositionSpace")
    
    
    return
}

func NewModifyPositionSpaceResponse() (response *ModifyPositionSpaceResponse) {
    response = &ModifyPositionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPositionSpace
// 更新位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SPACEDUPKEYEXIST = "UnsupportedOperation.SpaceDupKeyExist"
func (c *Client) ModifyPositionSpace(request *ModifyPositionSpaceRequest) (response *ModifyPositionSpaceResponse, err error) {
    return c.ModifyPositionSpaceWithContext(context.Background(), request)
}

// ModifyPositionSpace
// 更新位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SPACEDUPKEYEXIST = "UnsupportedOperation.SpaceDupKeyExist"
func (c *Client) ModifyPositionSpaceWithContext(ctx context.Context, request *ModifyPositionSpaceRequest) (response *ModifyPositionSpaceResponse, err error) {
    if request == nil {
        request = NewModifyPositionSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPositionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPositionSpaceResponse()
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

// ModifyProject
// 修改项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PROJECTDUPKEYEXIST = "UnsupportedOperation.ProjectDupKeyExist"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 修改项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE_PROJECTPARMSERROR = "InvalidParameterValue.ProjectParmsError"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PROJECTDUPKEYEXIST = "UnsupportedOperation.ProjectDupKeyExist"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifySpacePropertyRequest() (request *ModifySpacePropertyRequest) {
    request = &ModifySpacePropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifySpaceProperty")
    
    
    return
}

func NewModifySpacePropertyResponse() (response *ModifySpacePropertyResponse) {
    response = &ModifySpacePropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySpaceProperty
// 更新位置空间产品属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ModifySpaceProperty(request *ModifySpacePropertyRequest) (response *ModifySpacePropertyResponse, err error) {
    return c.ModifySpacePropertyWithContext(context.Background(), request)
}

// ModifySpaceProperty
// 更新位置空间产品属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SPACENOTEXIST = "ResourceNotFound.SpaceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ModifySpacePropertyWithContext(ctx context.Context, request *ModifySpacePropertyRequest) (response *ModifySpacePropertyResponse, err error) {
    if request == nil {
        request = NewModifySpacePropertyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySpaceProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySpacePropertyResponse()
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

// ModifyStudioProduct
// 提供修改产品的名称和描述等信息的能力，对于已发布产品不允许进行修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ModifyStudioProduct(request *ModifyStudioProductRequest) (response *ModifyStudioProductResponse, err error) {
    return c.ModifyStudioProductWithContext(context.Background(), request)
}

// ModifyStudioProduct
// 提供修改产品的名称和描述等信息的能力，对于已发布产品不允许进行修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINEDONTMATCHTEMPLATE = "InvalidParameterValue.ModelDefineDontMatchTemplate"
//  INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"
//  INVALIDPARAMETERVALUE_PRODUCTIDINVALID = "InvalidParameterValue.ProductIDInvalid"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  LIMITEXCEEDED_STUDIOPRODUCTEXCEEDLIMIT = "LimitExceeded.StudioProductExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ModifyStudioProductWithContext(ctx context.Context, request *ModifyStudioProductRequest) (response *ModifyStudioProductResponse, err error) {
    if request == nil {
        request = NewModifyStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStudioProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicPolicyRequest() (request *ModifyTopicPolicyRequest) {
    request = &ModifyTopicPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyTopicPolicy")
    
    
    return
}

func NewModifyTopicPolicyResponse() (response *ModifyTopicPolicyResponse) {
    response = &ModifyTopicPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTopicPolicy
// 本接口（UpdateTopicPolicy）用于更新Topic信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) ModifyTopicPolicy(request *ModifyTopicPolicyRequest) (response *ModifyTopicPolicyResponse, err error) {
    return c.ModifyTopicPolicyWithContext(context.Background(), request)
}

// ModifyTopicPolicy
// 本接口（UpdateTopicPolicy）用于更新Topic信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"
func (c *Client) ModifyTopicPolicyWithContext(ctx context.Context, request *ModifyTopicPolicyRequest) (response *ModifyTopicPolicyResponse, err error) {
    if request == nil {
        request = NewModifyTopicPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopicPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicPolicyResponse()
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

// ModifyTopicRule
// 修改规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_OPERATIONDENIED = "InvalidParameterValue.OperationDenied"
//  INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"
//  INVALIDPARAMETERVALUE_SELECTKEYFROMBINARYPAYLOAD = "InvalidParameterValue.SelectKeyFromBinaryPayload"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyTopicRule(request *ModifyTopicRuleRequest) (response *ModifyTopicRuleResponse, err error) {
    return c.ModifyTopicRuleWithContext(context.Background(), request)
}

// ModifyTopicRule
// 修改规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_OPERATIONDENIED = "InvalidParameterValue.OperationDenied"
//  INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"
//  INVALIDPARAMETERVALUE_SELECTKEYFROMBINARYPAYLOAD = "InvalidParameterValue.SelectKeyFromBinaryPayload"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyTopicRuleWithContext(ctx context.Context, request *ModifyTopicRuleRequest) (response *ModifyTopicRuleResponse, err error) {
    if request == nil {
        request = NewModifyTopicRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewPublishBroadcastMessageRequest() (request *PublishBroadcastMessageRequest) {
    request = &PublishBroadcastMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PublishBroadcastMessage")
    
    
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PublishBroadcastMessage(request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    return c.PublishBroadcastMessageWithContext(context.Background(), request)
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_BROADCASTTASKISRUNNING = "FailedOperation.BroadcastTaskIsRunning"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PublishMessage")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    return c.PublishMessageWithContext(context.Background(), request)
}

// PublishMessage
// 本接口（PublishMessage）用于使用自定义透传协议进行设备远控
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PublishRRPCMessage")
    
    
    return
}

func NewPublishRRPCMessageResponse() (response *PublishRRPCMessageResponse) {
    response = &PublishRRPCMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishRRPCMessage
// 下发RRPC消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_RRPCTIMEOUT = "FailedOperation.RRPCTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PublishRRPCMessage(request *PublishRRPCMessageRequest) (response *PublishRRPCMessageResponse, err error) {
    return c.PublishRRPCMessageWithContext(context.Background(), request)
}

// PublishRRPCMessage
// 下发RRPC消息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  FAILEDOPERATION_RRPCTIMEOUT = "FailedOperation.RRPCTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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

// ReleaseStudioProduct
// 产品开发完成并测试通过后，通过发布产品将产品设置为发布状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ReleaseStudioProduct(request *ReleaseStudioProductRequest) (response *ReleaseStudioProductResponse, err error) {
    return c.ReleaseStudioProductWithContext(context.Background(), request)
}

// ReleaseStudioProduct
// 产品开发完成并测试通过后，通过发布产品将产品设置为发布状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ReleaseStudioProductWithContext(ctx context.Context, request *ReleaseStudioProductRequest) (response *ReleaseStudioProductResponse, err error) {
    if request == nil {
        request = NewReleaseStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseStudioProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseStudioProductResponse()
    err = c.Send(request, response)
    return
}

func NewSearchPositionSpaceRequest() (request *SearchPositionSpaceRequest) {
    request = &SearchPositionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "SearchPositionSpace")
    
    
    return
}

func NewSearchPositionSpaceResponse() (response *SearchPositionSpaceResponse) {
    response = &SearchPositionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchPositionSpace
// 搜索位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchPositionSpace(request *SearchPositionSpaceRequest) (response *SearchPositionSpaceResponse, err error) {
    return c.SearchPositionSpaceWithContext(context.Background(), request)
}

// SearchPositionSpace
// 搜索位置空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchPositionSpaceWithContext(ctx context.Context, request *SearchPositionSpaceRequest) (response *SearchPositionSpaceResponse, err error) {
    if request == nil {
        request = NewSearchPositionSpaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchPositionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchPositionSpaceResponse()
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

// SearchStudioProduct
// 提供根据产品名称查找产品的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchStudioProduct(request *SearchStudioProductRequest) (response *SearchStudioProductResponse, err error) {
    return c.SearchStudioProductWithContext(context.Background(), request)
}

// SearchStudioProduct
// 提供根据产品名称查找产品的能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchStudioProductWithContext(ctx context.Context, request *SearchStudioProductRequest) (response *SearchStudioProductResponse, err error) {
    if request == nil {
        request = NewSearchStudioProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchStudioProduct require credential")
    }

    request.SetContext(ctx)
    
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

// SearchTopicRule
// 搜索规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchTopicRule(request *SearchTopicRuleRequest) (response *SearchTopicRuleResponse, err error) {
    return c.SearchTopicRuleWithContext(context.Background(), request)
}

// SearchTopicRule
// 搜索规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchTopicRuleWithContext(ctx context.Context, request *SearchTopicRuleRequest) (response *SearchTopicRuleResponse, err error) {
    if request == nil {
        request = NewSearchTopicRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchTopicRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchTopicRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDevicesRequest() (request *UnbindDevicesRequest) {
    request = &UnbindDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UnbindDevices")
    
    
    return
}

func NewUnbindDevicesResponse() (response *UnbindDevicesResponse) {
    response = &UnbindDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindDevices
// 批量解绑子设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UnbindDevices(request *UnbindDevicesRequest) (response *UnbindDevicesResponse, err error) {
    return c.UnbindDevicesWithContext(context.Background(), request)
}

// UnbindDevices
// 批量解绑子设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewUnbindProductsRequest() (request *UnbindProductsRequest) {
    request = &UnbindProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UnbindProducts")
    
    
    return
}

func NewUnbindProductsResponse() (response *UnbindProductsResponse) {
    response = &UnbindProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindProducts
// 批量解绑子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_EXISTBINDEDDEVICESUNDERGATEWAYPRODUCT = "UnsupportedOperation.ExistBindedDevicesUnderGatewayProduct"
func (c *Client) UnbindProducts(request *UnbindProductsRequest) (response *UnbindProductsResponse, err error) {
    return c.UnbindProductsWithContext(context.Background(), request)
}

// UnbindProducts
// 批量解绑子产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_EXISTBINDEDDEVICESUNDERGATEWAYPRODUCT = "UnsupportedOperation.ExistBindedDevicesUnderGatewayProduct"
func (c *Client) UnbindProductsWithContext(ctx context.Context, request *UnbindProductsRequest) (response *UnbindProductsResponse, err error) {
    if request == nil {
        request = NewUnbindProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindProductsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicesEnableStateRequest() (request *UpdateDevicesEnableStateRequest) {
    request = &UpdateDevicesEnableStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UpdateDevicesEnableState")
    
    
    return
}

func NewUpdateDevicesEnableStateResponse() (response *UpdateDevicesEnableStateResponse) {
    response = &UpdateDevicesEnableStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicesEnableState
// 批量禁用启用设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UpdateDevicesEnableState(request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    return c.UpdateDevicesEnableStateWithContext(context.Background(), request)
}

// UpdateDevicesEnableState
// 批量禁用启用设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

// UpdateFirmware
// 本接口（UpdateFirmware）用于对指定设备发起固件升级请求 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWAREISUPDATED = "FailedOperation.DeviceFirmwareIsUpdated"
//  FAILEDOPERATION_DEVICEINFOOUTDATED = "FailedOperation.DeviceInfoOutdated"
//  FAILEDOPERATION_OTHERUPDATETASKEXIST = "FailedOperation.OtherUpdateTaskExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UpdateFirmware(request *UpdateFirmwareRequest) (response *UpdateFirmwareResponse, err error) {
    return c.UpdateFirmwareWithContext(context.Background(), request)
}

// UpdateFirmware
// 本接口（UpdateFirmware）用于对指定设备发起固件升级请求 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFIRMWAREISUPDATED = "FailedOperation.DeviceFirmwareIsUpdated"
//  FAILEDOPERATION_DEVICEINFOOUTDATED = "FailedOperation.DeviceInfoOutdated"
//  FAILEDOPERATION_OTHERUPDATETASKEXIST = "FailedOperation.OtherUpdateTaskExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UpdateFirmwareWithContext(ctx context.Context, request *UpdateFirmwareRequest) (response *UpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewUpdateFirmwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFirmware require credential")
    }

    request.SetContext(ctx)
    
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

// UploadFirmware
// 本接口（UploadFirmware）用于创建设备固件版本信息，在平台用于固件版本升级、固件资源下发等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER_FIRMWAREALREADYEXIST = "InvalidParameter.FirmwareAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    return c.UploadFirmwareWithContext(context.Background(), request)
}

// UploadFirmware
// 本接口（UploadFirmware）用于创建设备固件版本信息，在平台用于固件版本升级、固件资源下发等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETER_FIRMWAREALREADYEXIST = "InvalidParameter.FirmwareAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
