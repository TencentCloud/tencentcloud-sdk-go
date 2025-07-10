// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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


func NewActivateTWeCallLicenseRequest() (request *ActivateTWeCallLicenseRequest) {
    request = &ActivateTWeCallLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ActivateTWeCallLicense")
    
    
    return
}

func NewActivateTWeCallLicenseResponse() (response *ActivateTWeCallLicenseResponse) {
    response = &ActivateTWeCallLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateTWeCallLicense
// 激活
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateTWeCallLicense(request *ActivateTWeCallLicenseRequest) (response *ActivateTWeCallLicenseResponse, err error) {
    return c.ActivateTWeCallLicenseWithContext(context.Background(), request)
}

// ActivateTWeCallLicense
// 激活
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateTWeCallLicenseWithContext(ctx context.Context, request *ActivateTWeCallLicenseRequest) (response *ActivateTWeCallLicenseResponse, err error) {
    if request == nil {
        request = NewActivateTWeCallLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateTWeCallLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateTWeCallLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewBindCloudStorageUserRequest() (request *BindCloudStorageUserRequest) {
    request = &BindCloudStorageUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "BindCloudStorageUser")
    
    
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
// 批量绑定子产品。
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
// 批量绑定子产品。
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

func NewCancelAssignTWeCallLicenseRequest() (request *CancelAssignTWeCallLicenseRequest) {
    request = &CancelAssignTWeCallLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CancelAssignTWeCallLicense")
    
    
    return
}

func NewCancelAssignTWeCallLicenseResponse() (response *CancelAssignTWeCallLicenseResponse) {
    response = &CancelAssignTWeCallLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelAssignTWeCallLicense
// 取消分配
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAssignTWeCallLicense(request *CancelAssignTWeCallLicenseRequest) (response *CancelAssignTWeCallLicenseResponse, err error) {
    return c.CancelAssignTWeCallLicenseWithContext(context.Background(), request)
}

// CancelAssignTWeCallLicense
// 取消分配
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelAssignTWeCallLicenseWithContext(ctx context.Context, request *CancelAssignTWeCallLicenseRequest) (response *CancelAssignTWeCallLicenseResponse, err error) {
    if request == nil {
        request = NewCancelAssignTWeCallLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAssignTWeCallLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAssignTWeCallLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewChangeP2PRouteRequest() (request *ChangeP2PRouteRequest) {
    request = &ChangeP2PRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ChangeP2PRoute")
    
    
    return
}

func NewChangeP2PRouteResponse() (response *ChangeP2PRouteResponse) {
    response = &ChangeP2PRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeP2PRoute
// p2p路线切换（此接口目前处于内测接口，可以联系申请加白 ）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChangeP2PRoute(request *ChangeP2PRouteRequest) (response *ChangeP2PRouteResponse, err error) {
    return c.ChangeP2PRouteWithContext(context.Background(), request)
}

// ChangeP2PRoute
// p2p路线切换（此接口目前处于内测接口，可以联系申请加白 ）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChangeP2PRouteWithContext(ctx context.Context, request *ChangeP2PRouteRequest) (response *ChangeP2PRouteResponse, err error) {
    if request == nil {
        request = NewChangeP2PRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeP2PRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeP2PRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFirmwareUpdateRequest() (request *CheckFirmwareUpdateRequest) {
    request = &CheckFirmwareUpdateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CheckFirmwareUpdate")
    
    
    return
}

func NewCheckFirmwareUpdateResponse() (response *CheckFirmwareUpdateResponse) {
    response = &CheckFirmwareUpdateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckFirmwareUpdate
// 本接口（CheckFirmwareUpdate）用于查询设备可升级固件版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) CheckFirmwareUpdate(request *CheckFirmwareUpdateRequest) (response *CheckFirmwareUpdateResponse, err error) {
    return c.CheckFirmwareUpdateWithContext(context.Background(), request)
}

// CheckFirmwareUpdate
// 本接口（CheckFirmwareUpdate）用于查询设备可升级固件版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) CheckFirmwareUpdateWithContext(ctx context.Context, request *CheckFirmwareUpdateRequest) (response *CheckFirmwareUpdateResponse, err error) {
    if request == nil {
        request = NewCheckFirmwareUpdateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFirmwareUpdate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFirmwareUpdateResponse()
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
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

func NewCreateAISearchTaskAsyncRequest() (request *CreateAISearchTaskAsyncRequest) {
    request = &CreateAISearchTaskAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateAISearchTaskAsync")
    
    
    return
}

func NewCreateAISearchTaskAsyncResponse() (response *CreateAISearchTaskAsyncResponse) {
    response = &CreateAISearchTaskAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAISearchTaskAsync
// 创建视频语义异步搜索任务
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
func (c *Client) CreateAISearchTaskAsync(request *CreateAISearchTaskAsyncRequest) (response *CreateAISearchTaskAsyncResponse, err error) {
    return c.CreateAISearchTaskAsyncWithContext(context.Background(), request)
}

// CreateAISearchTaskAsync
// 创建视频语义异步搜索任务
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
func (c *Client) CreateAISearchTaskAsyncWithContext(ctx context.Context, request *CreateAISearchTaskAsyncRequest) (response *CreateAISearchTaskAsyncResponse, err error) {
    if request == nil {
        request = NewCreateAISearchTaskAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAISearchTaskAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAISearchTaskAsyncResponse()
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

func NewCreateCloudStorageAIServiceRequest() (request *CreateCloudStorageAIServiceRequest) {
    request = &CreateCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateCloudStorageAIService")
    
    
    return
}

func NewCreateCloudStorageAIServiceResponse() (response *CreateCloudStorageAIServiceResponse) {
    response = &CreateCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudStorageAIService
// 开通设备云存AI分析服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAIPACKAGEEXPIRETIMEEXCEEDED = "FailedOperation.CloudStorageAIPackageExpireTimeExceeded"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_CLOUDSTORAGEPACKAGEREQUIRED = "FailedOperation.CloudStoragePackageRequired"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLOUDSTORAGEAIPACKAGEIDNOTEXIST = "InvalidParameterValue.CloudStorageAIPackageIdNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLOUDSTORAGEPACKAGETIMEMISMATCH = "UnsupportedOperation.CloudStoragePackageTimeMismatch"
//  UNSUPPORTEDOPERATION_CLOUDSTORAGEPACKAGETYPEMISMATCH = "UnsupportedOperation.CloudStoragePackageTypeMismatch"
func (c *Client) CreateCloudStorageAIService(request *CreateCloudStorageAIServiceRequest) (response *CreateCloudStorageAIServiceResponse, err error) {
    return c.CreateCloudStorageAIServiceWithContext(context.Background(), request)
}

// CreateCloudStorageAIService
// 开通设备云存AI分析服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAIPACKAGEEXPIRETIMEEXCEEDED = "FailedOperation.CloudStorageAIPackageExpireTimeExceeded"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_CLOUDSTORAGEPACKAGEREQUIRED = "FailedOperation.CloudStoragePackageRequired"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLOUDSTORAGEAIPACKAGEIDNOTEXIST = "InvalidParameterValue.CloudStorageAIPackageIdNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLOUDSTORAGEPACKAGETIMEMISMATCH = "UnsupportedOperation.CloudStoragePackageTimeMismatch"
//  UNSUPPORTEDOPERATION_CLOUDSTORAGEPACKAGETYPEMISMATCH = "UnsupportedOperation.CloudStoragePackageTypeMismatch"
func (c *Client) CreateCloudStorageAIServiceWithContext(ctx context.Context, request *CreateCloudStorageAIServiceRequest) (response *CreateCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewCreateCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudStorageAIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudStorageAIServiceTaskRequest() (request *CreateCloudStorageAIServiceTaskRequest) {
    request = &CreateCloudStorageAIServiceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateCloudStorageAIServiceTask")
    
    
    return
}

func NewCreateCloudStorageAIServiceTaskResponse() (response *CreateCloudStorageAIServiceTaskResponse) {
    response = &CreateCloudStorageAIServiceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudStorageAIServiceTask
// 创建设备云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) CreateCloudStorageAIServiceTask(request *CreateCloudStorageAIServiceTaskRequest) (response *CreateCloudStorageAIServiceTaskResponse, err error) {
    return c.CreateCloudStorageAIServiceTaskWithContext(context.Background(), request)
}

// CreateCloudStorageAIServiceTask
// 创建设备云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) CreateCloudStorageAIServiceTaskWithContext(ctx context.Context, request *CreateCloudStorageAIServiceTaskRequest) (response *CreateCloudStorageAIServiceTaskResponse, err error) {
    if request == nil {
        request = NewCreateCloudStorageAIServiceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudStorageAIServiceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudStorageAIServiceTaskResponse()
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

func NewCreateDeviceChannelRequest() (request *CreateDeviceChannelRequest) {
    request = &CreateDeviceChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateDeviceChannel")
    
    
    return
}

func NewCreateDeviceChannelResponse() (response *CreateDeviceChannelResponse) {
    response = &CreateDeviceChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceChannel
// 创建设备通道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceChannel(request *CreateDeviceChannelRequest) (response *CreateDeviceChannelResponse, err error) {
    return c.CreateDeviceChannelWithContext(context.Background(), request)
}

// CreateDeviceChannel
// 创建设备通道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeviceChannelWithContext(ctx context.Context, request *CreateDeviceChannelRequest) (response *CreateDeviceChannelResponse, err error) {
    if request == nil {
        request = NewCreateDeviceChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExternalSourceAIServiceTaskRequest() (request *CreateExternalSourceAIServiceTaskRequest) {
    request = &CreateExternalSourceAIServiceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateExternalSourceAIServiceTask")
    
    
    return
}

func NewCreateExternalSourceAIServiceTaskResponse() (response *CreateExternalSourceAIServiceTaskResponse) {
    response = &CreateExternalSourceAIServiceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExternalSourceAIServiceTask
// 创建外部视频 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) CreateExternalSourceAIServiceTask(request *CreateExternalSourceAIServiceTaskRequest) (response *CreateExternalSourceAIServiceTaskResponse, err error) {
    return c.CreateExternalSourceAIServiceTaskWithContext(context.Background(), request)
}

// CreateExternalSourceAIServiceTask
// 创建外部视频 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) CreateExternalSourceAIServiceTaskWithContext(ctx context.Context, request *CreateExternalSourceAIServiceTaskRequest) (response *CreateExternalSourceAIServiceTaskResponse, err error) {
    if request == nil {
        request = NewCreateExternalSourceAIServiceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExternalSourceAIServiceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExternalSourceAIServiceTaskResponse()
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
// > 创建围栏绑定信息。
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
// > 创建围栏绑定信息。
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

func NewCreateFreeCloudStorageRequest() (request *CreateFreeCloudStorageRequest) {
    request = &CreateFreeCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateFreeCloudStorage")
    
    
    return
}

func NewCreateFreeCloudStorageResponse() (response *CreateFreeCloudStorageResponse) {
    response = &CreateFreeCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFreeCloudStorage
// 开通云存卡服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFreeCloudStorage(request *CreateFreeCloudStorageRequest) (response *CreateFreeCloudStorageResponse, err error) {
    return c.CreateFreeCloudStorageWithContext(context.Background(), request)
}

// CreateFreeCloudStorage
// 开通云存卡服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFreeCloudStorageWithContext(ctx context.Context, request *CreateFreeCloudStorageRequest) (response *CreateFreeCloudStorageResponse, err error) {
    if request == nil {
        request = NewCreateFreeCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFreeCloudStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFreeCloudStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIotVideoCloudStorageRequest() (request *CreateIotVideoCloudStorageRequest) {
    request = &CreateIotVideoCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateIotVideoCloudStorage")
    
    
    return
}

func NewCreateIotVideoCloudStorageResponse() (response *CreateIotVideoCloudStorageResponse) {
    response = &CreateIotVideoCloudStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIotVideoCloudStorage
// 开通云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotVideoCloudStorage(request *CreateIotVideoCloudStorageRequest) (response *CreateIotVideoCloudStorageResponse, err error) {
    return c.CreateIotVideoCloudStorageWithContext(context.Background(), request)
}

// CreateIotVideoCloudStorage
// 开通云存服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotVideoCloudStorageWithContext(ctx context.Context, request *CreateIotVideoCloudStorageRequest) (response *CreateIotVideoCloudStorageResponse, err error) {
    if request == nil {
        request = NewCreateIotVideoCloudStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIotVideoCloudStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIotVideoCloudStorageResponse()
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
// 创建围栏。
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
// 创建围栏。
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
// 创建位置空间。
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
// 创建位置空间。
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

func NewCreateTRTCSignaturesWithRoomIdRequest() (request *CreateTRTCSignaturesWithRoomIdRequest) {
    request = &CreateTRTCSignaturesWithRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateTRTCSignaturesWithRoomId")
    
    
    return
}

func NewCreateTRTCSignaturesWithRoomIdResponse() (response *CreateTRTCSignaturesWithRoomIdResponse) {
    response = &CreateTRTCSignaturesWithRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTRTCSignaturesWithRoomId
// 创建TRTC通话参数
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ROOMIDEXIST = "InvalidParameter.RoomIdExist"
//  INVALIDPARAMETER_USERIDOVERLIMIT = "InvalidParameter.UserIdOverLimit"
//  UNSUPPORTEDOPERATION_TRTCSERVICENOTOPEN = "UnsupportedOperation.TRTCServiceNotOpen"
func (c *Client) CreateTRTCSignaturesWithRoomId(request *CreateTRTCSignaturesWithRoomIdRequest) (response *CreateTRTCSignaturesWithRoomIdResponse, err error) {
    return c.CreateTRTCSignaturesWithRoomIdWithContext(context.Background(), request)
}

// CreateTRTCSignaturesWithRoomId
// 创建TRTC通话参数
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ROOMIDEXIST = "InvalidParameter.RoomIdExist"
//  INVALIDPARAMETER_USERIDOVERLIMIT = "InvalidParameter.UserIdOverLimit"
//  UNSUPPORTEDOPERATION_TRTCSERVICENOTOPEN = "UnsupportedOperation.TRTCServiceNotOpen"
func (c *Client) CreateTRTCSignaturesWithRoomIdWithContext(ctx context.Context, request *CreateTRTCSignaturesWithRoomIdRequest) (response *CreateTRTCSignaturesWithRoomIdResponse, err error) {
    if request == nil {
        request = NewCreateTRTCSignaturesWithRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTRTCSignaturesWithRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTRTCSignaturesWithRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTWeSeeRecognitionTaskRequest() (request *CreateTWeSeeRecognitionTaskRequest) {
    request = &CreateTWeSeeRecognitionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "CreateTWeSeeRecognitionTask")
    
    
    return
}

func NewCreateTWeSeeRecognitionTaskResponse() (response *CreateTWeSeeRecognitionTaskResponse) {
    response = &CreateTWeSeeRecognitionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTWeSeeRecognitionTask
// 创建 TWeSee 语义理解任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTWeSeeRecognitionTask(request *CreateTWeSeeRecognitionTaskRequest) (response *CreateTWeSeeRecognitionTaskResponse, err error) {
    return c.CreateTWeSeeRecognitionTaskWithContext(context.Background(), request)
}

// CreateTWeSeeRecognitionTask
// 创建 TWeSee 语义理解任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTWeSeeRecognitionTaskWithContext(ctx context.Context, request *CreateTWeSeeRecognitionTaskRequest) (response *CreateTWeSeeRecognitionTaskResponse, err error) {
    if request == nil {
        request = NewCreateTWeSeeRecognitionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTWeSeeRecognitionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTWeSeeRecognitionTaskResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDeleteCloudStorageEventRequest() (request *DeleteCloudStorageEventRequest) {
    request = &DeleteCloudStorageEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DeleteCloudStorageEvent")
    
    
    return
}

func NewDeleteCloudStorageEventResponse() (response *DeleteCloudStorageEventResponse) {
    response = &DeleteCloudStorageEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudStorageEvent
// 删除云存事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudStorageEvent(request *DeleteCloudStorageEventRequest) (response *DeleteCloudStorageEventResponse, err error) {
    return c.DeleteCloudStorageEventWithContext(context.Background(), request)
}

// DeleteCloudStorageEvent
// 删除云存事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudStorageEventWithContext(ctx context.Context, request *DeleteCloudStorageEventRequest) (response *DeleteCloudStorageEventResponse, err error) {
    if request == nil {
        request = NewDeleteCloudStorageEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudStorageEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudStorageEventResponse()
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
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
// 删除围栏绑定信息。
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
// 删除围栏绑定信息。
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
// 删除围栏。
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
// 删除围栏。
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
// 删除位置空间。
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
// 删除位置空间。
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
// 提供删除某个项目的能力。
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
// 提供删除某个项目的能力。
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
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGateWayProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASSHARED = "UnsupportedOperation.ProductHasShared"
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
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGateWayProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASSHARED = "UnsupportedOperation.ProductHasShared"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
// 删除规则。
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
// 删除规则。
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

func NewDescribeAISearchTaskAsyncRequest() (request *DescribeAISearchTaskAsyncRequest) {
    request = &DescribeAISearchTaskAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeAISearchTaskAsync")
    
    
    return
}

func NewDescribeAISearchTaskAsyncResponse() (response *DescribeAISearchTaskAsyncResponse) {
    response = &DescribeAISearchTaskAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAISearchTaskAsync
// 获取视频语义异步搜索任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAISearchTaskAsync(request *DescribeAISearchTaskAsyncRequest) (response *DescribeAISearchTaskAsyncResponse, err error) {
    return c.DescribeAISearchTaskAsyncWithContext(context.Background(), request)
}

// DescribeAISearchTaskAsync
// 获取视频语义异步搜索任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAISearchTaskAsyncWithContext(ctx context.Context, request *DescribeAISearchTaskAsyncRequest) (response *DescribeAISearchTaskAsyncResponse, err error) {
    if request == nil {
        request = NewDescribeAISearchTaskAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAISearchTaskAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAISearchTaskAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActivateDeviceRequest() (request *DescribeActivateDeviceRequest) {
    request = &DescribeActivateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeActivateDevice")
    
    
    return
}

func NewDescribeActivateDeviceResponse() (response *DescribeActivateDeviceResponse) {
    response = &DescribeActivateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeActivateDevice
// 获取设备激活详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeActivateDevice(request *DescribeActivateDeviceRequest) (response *DescribeActivateDeviceResponse, err error) {
    return c.DescribeActivateDeviceWithContext(context.Background(), request)
}

// DescribeActivateDevice
// 获取设备激活详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeActivateDeviceWithContext(ctx context.Context, request *DescribeActivateDeviceRequest) (response *DescribeActivateDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeActivateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActivateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActivateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActivateLicenseServiceRequest() (request *DescribeActivateLicenseServiceRequest) {
    request = &DescribeActivateLicenseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeActivateLicenseService")
    
    
    return
}

func NewDescribeActivateLicenseServiceResponse() (response *DescribeActivateLicenseServiceResponse) {
    response = &DescribeActivateLicenseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeActivateLicenseService
// 获取增值服务激活码详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeActivateLicenseService(request *DescribeActivateLicenseServiceRequest) (response *DescribeActivateLicenseServiceResponse, err error) {
    return c.DescribeActivateLicenseServiceWithContext(context.Background(), request)
}

// DescribeActivateLicenseService
// 获取增值服务激活码详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeActivateLicenseServiceWithContext(ctx context.Context, request *DescribeActivateLicenseServiceRequest) (response *DescribeActivateLicenseServiceResponse, err error) {
    if request == nil {
        request = NewDescribeActivateLicenseServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActivateLicenseService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActivateLicenseServiceResponse()
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

func NewDescribeCloudStorageRequest() (request *DescribeCloudStorageRequest) {
    request = &DescribeCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorage")
    
    
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

func NewDescribeCloudStorageAIServiceRequest() (request *DescribeCloudStorageAIServiceRequest) {
    request = &DescribeCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageAIService")
    
    
    return
}

func NewDescribeCloudStorageAIServiceResponse() (response *DescribeCloudStorageAIServiceResponse) {
    response = &DescribeCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageAIService
// 查询指定设备的云存 AI 服务开通状态与参数配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageAIService(request *DescribeCloudStorageAIServiceRequest) (response *DescribeCloudStorageAIServiceResponse, err error) {
    return c.DescribeCloudStorageAIServiceWithContext(context.Background(), request)
}

// DescribeCloudStorageAIService
// 查询指定设备的云存 AI 服务开通状态与参数配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageAIServiceWithContext(ctx context.Context, request *DescribeCloudStorageAIServiceRequest) (response *DescribeCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageAIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageAIServiceCallbackRequest() (request *DescribeCloudStorageAIServiceCallbackRequest) {
    request = &DescribeCloudStorageAIServiceCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageAIServiceCallback")
    
    
    return
}

func NewDescribeCloudStorageAIServiceCallbackResponse() (response *DescribeCloudStorageAIServiceCallbackResponse) {
    response = &DescribeCloudStorageAIServiceCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageAIServiceCallback
// 查询云存AI分析回调配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageAIServiceCallback(request *DescribeCloudStorageAIServiceCallbackRequest) (response *DescribeCloudStorageAIServiceCallbackResponse, err error) {
    return c.DescribeCloudStorageAIServiceCallbackWithContext(context.Background(), request)
}

// DescribeCloudStorageAIServiceCallback
// 查询云存AI分析回调配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageAIServiceCallbackWithContext(ctx context.Context, request *DescribeCloudStorageAIServiceCallbackRequest) (response *DescribeCloudStorageAIServiceCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageAIServiceCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageAIServiceCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageAIServiceCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageAIServiceTaskRequest() (request *DescribeCloudStorageAIServiceTaskRequest) {
    request = &DescribeCloudStorageAIServiceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageAIServiceTask")
    
    
    return
}

func NewDescribeCloudStorageAIServiceTaskResponse() (response *DescribeCloudStorageAIServiceTaskResponse) {
    response = &DescribeCloudStorageAIServiceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageAIServiceTask
// 查询指定的云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASK = "ResourceNotFound.CloudStorageAIServiceTask"
func (c *Client) DescribeCloudStorageAIServiceTask(request *DescribeCloudStorageAIServiceTaskRequest) (response *DescribeCloudStorageAIServiceTaskResponse, err error) {
    return c.DescribeCloudStorageAIServiceTaskWithContext(context.Background(), request)
}

// DescribeCloudStorageAIServiceTask
// 查询指定的云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASK = "ResourceNotFound.CloudStorageAIServiceTask"
func (c *Client) DescribeCloudStorageAIServiceTaskWithContext(ctx context.Context, request *DescribeCloudStorageAIServiceTaskRequest) (response *DescribeCloudStorageAIServiceTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageAIServiceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageAIServiceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageAIServiceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageAIServiceTasksRequest() (request *DescribeCloudStorageAIServiceTasksRequest) {
    request = &DescribeCloudStorageAIServiceTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageAIServiceTasks")
    
    
    return
}

func NewDescribeCloudStorageAIServiceTasksResponse() (response *DescribeCloudStorageAIServiceTasksResponse) {
    response = &DescribeCloudStorageAIServiceTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageAIServiceTasks
// 查询指定设备的云存 AI 分析任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudStorageAIServiceTasks(request *DescribeCloudStorageAIServiceTasksRequest) (response *DescribeCloudStorageAIServiceTasksResponse, err error) {
    return c.DescribeCloudStorageAIServiceTasksWithContext(context.Background(), request)
}

// DescribeCloudStorageAIServiceTasks
// 查询指定设备的云存 AI 分析任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudStorageAIServiceTasksWithContext(ctx context.Context, request *DescribeCloudStorageAIServiceTasksRequest) (response *DescribeCloudStorageAIServiceTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageAIServiceTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageAIServiceTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageAIServiceTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageDateRequest() (request *DescribeCloudStorageDateRequest) {
    request = &DescribeCloudStorageDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageDate")
    
    
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageEvents")
    
    
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

func NewDescribeCloudStorageEventsWithAITasksRequest() (request *DescribeCloudStorageEventsWithAITasksRequest) {
    request = &DescribeCloudStorageEventsWithAITasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageEventsWithAITasks")
    
    
    return
}

func NewDescribeCloudStorageEventsWithAITasksResponse() (response *DescribeCloudStorageEventsWithAITasksResponse) {
    response = &DescribeCloudStorageEventsWithAITasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageEventsWithAITasks
// 拉取云存事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageEventsWithAITasks(request *DescribeCloudStorageEventsWithAITasksRequest) (response *DescribeCloudStorageEventsWithAITasksResponse, err error) {
    return c.DescribeCloudStorageEventsWithAITasksWithContext(context.Background(), request)
}

// DescribeCloudStorageEventsWithAITasks
// 拉取云存事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageEventsWithAITasksWithContext(ctx context.Context, request *DescribeCloudStorageEventsWithAITasksRequest) (response *DescribeCloudStorageEventsWithAITasksResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageEventsWithAITasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageEventsWithAITasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageEventsWithAITasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageMultiThumbnailRequest() (request *DescribeCloudStorageMultiThumbnailRequest) {
    request = &DescribeCloudStorageMultiThumbnailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageMultiThumbnail")
    
    
    return
}

func NewDescribeCloudStorageMultiThumbnailResponse() (response *DescribeCloudStorageMultiThumbnailResponse) {
    response = &DescribeCloudStorageMultiThumbnailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageMultiThumbnail
// 拉取多个云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageMultiThumbnail(request *DescribeCloudStorageMultiThumbnailRequest) (response *DescribeCloudStorageMultiThumbnailResponse, err error) {
    return c.DescribeCloudStorageMultiThumbnailWithContext(context.Background(), request)
}

// DescribeCloudStorageMultiThumbnail
// 拉取多个云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageMultiThumbnailWithContext(ctx context.Context, request *DescribeCloudStorageMultiThumbnailRequest) (response *DescribeCloudStorageMultiThumbnailResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageMultiThumbnailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageMultiThumbnail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageMultiThumbnailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageOrderRequest() (request *DescribeCloudStorageOrderRequest) {
    request = &DescribeCloudStorageOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageOrder")
    
    
    return
}

func NewDescribeCloudStorageOrderResponse() (response *DescribeCloudStorageOrderResponse) {
    response = &DescribeCloudStorageOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageOrder
// 查询云存服务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageOrder(request *DescribeCloudStorageOrderRequest) (response *DescribeCloudStorageOrderResponse, err error) {
    return c.DescribeCloudStorageOrderWithContext(context.Background(), request)
}

// DescribeCloudStorageOrder
// 查询云存服务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageOrderWithContext(ctx context.Context, request *DescribeCloudStorageOrderRequest) (response *DescribeCloudStorageOrderResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStoragePackageConsumeDetailsRequest() (request *DescribeCloudStoragePackageConsumeDetailsRequest) {
    request = &DescribeCloudStoragePackageConsumeDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStoragePackageConsumeDetails")
    
    
    return
}

func NewDescribeCloudStoragePackageConsumeDetailsResponse() (response *DescribeCloudStoragePackageConsumeDetailsResponse) {
    response = &DescribeCloudStoragePackageConsumeDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStoragePackageConsumeDetails
// 获取云存套餐包消耗详细记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStoragePackageConsumeDetails(request *DescribeCloudStoragePackageConsumeDetailsRequest) (response *DescribeCloudStoragePackageConsumeDetailsResponse, err error) {
    return c.DescribeCloudStoragePackageConsumeDetailsWithContext(context.Background(), request)
}

// DescribeCloudStoragePackageConsumeDetails
// 获取云存套餐包消耗详细记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStoragePackageConsumeDetailsWithContext(ctx context.Context, request *DescribeCloudStoragePackageConsumeDetailsRequest) (response *DescribeCloudStoragePackageConsumeDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStoragePackageConsumeDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStoragePackageConsumeDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStoragePackageConsumeDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStoragePackageConsumeStatsRequest() (request *DescribeCloudStoragePackageConsumeStatsRequest) {
    request = &DescribeCloudStoragePackageConsumeStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStoragePackageConsumeStats")
    
    
    return
}

func NewDescribeCloudStoragePackageConsumeStatsResponse() (response *DescribeCloudStoragePackageConsumeStatsResponse) {
    response = &DescribeCloudStoragePackageConsumeStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStoragePackageConsumeStats
// 获取云存套餐包消耗统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStoragePackageConsumeStats(request *DescribeCloudStoragePackageConsumeStatsRequest) (response *DescribeCloudStoragePackageConsumeStatsResponse, err error) {
    return c.DescribeCloudStoragePackageConsumeStatsWithContext(context.Background(), request)
}

// DescribeCloudStoragePackageConsumeStats
// 获取云存套餐包消耗统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStoragePackageConsumeStatsWithContext(ctx context.Context, request *DescribeCloudStoragePackageConsumeStatsRequest) (response *DescribeCloudStoragePackageConsumeStatsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStoragePackageConsumeStatsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStoragePackageConsumeStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStoragePackageConsumeStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageStreamDataRequest() (request *DescribeCloudStorageStreamDataRequest) {
    request = &DescribeCloudStorageStreamDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageStreamData")
    
    
    return
}

func NewDescribeCloudStorageStreamDataResponse() (response *DescribeCloudStorageStreamDataResponse) {
    response = &DescribeCloudStorageStreamDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageStreamData
// 获取设备图片流数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageStreamData(request *DescribeCloudStorageStreamDataRequest) (response *DescribeCloudStorageStreamDataResponse, err error) {
    return c.DescribeCloudStorageStreamDataWithContext(context.Background(), request)
}

// DescribeCloudStorageStreamData
// 获取设备图片流数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageStreamDataWithContext(ctx context.Context, request *DescribeCloudStorageStreamDataRequest) (response *DescribeCloudStorageStreamDataResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageStreamDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageStreamData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageStreamDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageThumbnailRequest() (request *DescribeCloudStorageThumbnailRequest) {
    request = &DescribeCloudStorageThumbnailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageThumbnail")
    
    
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

func NewDescribeCloudStorageThumbnailListRequest() (request *DescribeCloudStorageThumbnailListRequest) {
    request = &DescribeCloudStorageThumbnailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageThumbnailList")
    
    
    return
}

func NewDescribeCloudStorageThumbnailListResponse() (response *DescribeCloudStorageThumbnailListResponse) {
    response = &DescribeCloudStorageThumbnailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudStorageThumbnailList
// 批量拉取云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageThumbnailList(request *DescribeCloudStorageThumbnailListRequest) (response *DescribeCloudStorageThumbnailListResponse, err error) {
    return c.DescribeCloudStorageThumbnailListWithContext(context.Background(), request)
}

// DescribeCloudStorageThumbnailList
// 批量拉取云存事件缩略图
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudStorageThumbnailListWithContext(ctx context.Context, request *DescribeCloudStorageThumbnailListRequest) (response *DescribeCloudStorageThumbnailListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudStorageThumbnailListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudStorageThumbnailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudStorageThumbnailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudStorageTimeRequest() (request *DescribeCloudStorageTimeRequest) {
    request = &DescribeCloudStorageTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageTime")
    
    
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
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCloudStorageUsers")
    
    
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

func NewDescribeCsReportCountDataInfoRequest() (request *DescribeCsReportCountDataInfoRequest) {
    request = &DescribeCsReportCountDataInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeCsReportCountDataInfo")
    
    
    return
}

func NewDescribeCsReportCountDataInfoResponse() (response *DescribeCsReportCountDataInfoResponse) {
    response = &DescribeCsReportCountDataInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCsReportCountDataInfo
// 获取云存上报统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCsReportCountDataInfo(request *DescribeCsReportCountDataInfoRequest) (response *DescribeCsReportCountDataInfoResponse, err error) {
    return c.DescribeCsReportCountDataInfoWithContext(context.Background(), request)
}

// DescribeCsReportCountDataInfo
// 获取云存上报统计信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCsReportCountDataInfoWithContext(ctx context.Context, request *DescribeCsReportCountDataInfoRequest) (response *DescribeCsReportCountDataInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCsReportCountDataInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCsReportCountDataInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCsReportCountDataInfoResponse()
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

func NewDescribeDeviceFirmWareRequest() (request *DescribeDeviceFirmWareRequest) {
    request = &DescribeDeviceFirmWareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceFirmWare")
    
    
    return
}

func NewDescribeDeviceFirmWareResponse() (response *DescribeDeviceFirmWareResponse) {
    response = &DescribeDeviceFirmWareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceFirmWare
// 获取设备固件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEFIRMWARENOTEXIST = "ResourceNotFound.DeviceFirmWareNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeDeviceFirmWare(request *DescribeDeviceFirmWareRequest) (response *DescribeDeviceFirmWareResponse, err error) {
    return c.DescribeDeviceFirmWareWithContext(context.Background(), request)
}

// DescribeDeviceFirmWare
// 获取设备固件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEFIRMWARENOTEXIST = "ResourceNotFound.DeviceFirmWareNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeDeviceFirmWareWithContext(ctx context.Context, request *DescribeDeviceFirmWareRequest) (response *DescribeDeviceFirmWareResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceFirmWareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceFirmWare require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceFirmWareResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceFirmwaresRequest() (request *DescribeDeviceFirmwaresRequest) {
    request = &DescribeDeviceFirmwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDeviceFirmwares")
    
    
    return
}

func NewDescribeDeviceFirmwaresResponse() (response *DescribeDeviceFirmwaresResponse) {
    response = &DescribeDeviceFirmwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceFirmwares
// 获取设备当前固件信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeDeviceFirmwares(request *DescribeDeviceFirmwaresRequest) (response *DescribeDeviceFirmwaresResponse, err error) {
    return c.DescribeDeviceFirmwaresWithContext(context.Background(), request)
}

// DescribeDeviceFirmwares
// 获取设备当前固件信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICENOTEXIST = "InvalidParameterValue.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeDeviceFirmwaresWithContext(ctx context.Context, request *DescribeDeviceFirmwaresRequest) (response *DescribeDeviceFirmwaresResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceFirmwaresRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceFirmwares require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceFirmwaresResponse()
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
// 获取实时位置解析依赖于teg位置服务，近30天调用只有2个个人账号调用，产品推下线
//
// 
//
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
// 获取实时位置解析依赖于teg位置服务，近30天调用只有2个个人账号调用，产品推下线
//
// 
//
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

func NewDescribeDevicePackagesRequest() (request *DescribeDevicePackagesRequest) {
    request = &DescribeDevicePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeDevicePackages")
    
    
    return
}

func NewDescribeDevicePackagesResponse() (response *DescribeDevicePackagesResponse) {
    response = &DescribeDevicePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDevicePackages
// 根据设备信息拉取有效套餐列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevicePackages(request *DescribeDevicePackagesRequest) (response *DescribeDevicePackagesResponse, err error) {
    return c.DescribeDevicePackagesWithContext(context.Background(), request)
}

// DescribeDevicePackages
// 根据设备信息拉取有效套餐列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevicePackagesWithContext(ctx context.Context, request *DescribeDevicePackagesRequest) (response *DescribeDevicePackagesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePackagesResponse()
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
// 获取围栏绑定信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOFENCE = "UnauthorizedOperation.NoPermissionToStudioFence"
func (c *Client) DescribeFenceBindList(request *DescribeFenceBindListRequest) (response *DescribeFenceBindListResponse, err error) {
    return c.DescribeFenceBindListWithContext(context.Background(), request)
}

// DescribeFenceBindList
// 获取围栏绑定信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FENCENOTEXIST = "ResourceNotFound.FenceNotExist"
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
// 获取围栏告警事件列表。
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
// 获取围栏告警事件列表。
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

func NewDescribeFirmwareRequest() (request *DescribeFirmwareRequest) {
    request = &DescribeFirmwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFirmware")
    
    
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
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeFirmware(request *DescribeFirmwareRequest) (response *DescribeFirmwareResponse, err error) {
    return c.DescribeFirmwareWithContext(context.Background(), request)
}

// DescribeFirmware
// 查询固件信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
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

func NewDescribeFirmwareUpdateStatusRequest() (request *DescribeFirmwareUpdateStatusRequest) {
    request = &DescribeFirmwareUpdateStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFirmwareUpdateStatus")
    
    
    return
}

func NewDescribeFirmwareUpdateStatusResponse() (response *DescribeFirmwareUpdateStatusResponse) {
    response = &DescribeFirmwareUpdateStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFirmwareUpdateStatus
// 本接口（DescribeFirmwareUpdateStatus）用于查询设备固件升级状态及进度。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOPERMISSION = "InvalidParameterValue.NoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFirmwareUpdateStatus(request *DescribeFirmwareUpdateStatusRequest) (response *DescribeFirmwareUpdateStatusResponse, err error) {
    return c.DescribeFirmwareUpdateStatusWithContext(context.Background(), request)
}

// DescribeFirmwareUpdateStatus
// 本接口（DescribeFirmwareUpdateStatus）用于查询设备固件升级状态及进度。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOPERMISSION = "InvalidParameterValue.NoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFirmwareUpdateStatusWithContext(ctx context.Context, request *DescribeFirmwareUpdateStatusRequest) (response *DescribeFirmwareUpdateStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFirmwareUpdateStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirmwareUpdateStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirmwareUpdateStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFreeCloudStorageNumRequest() (request *DescribeFreeCloudStorageNumRequest) {
    request = &DescribeFreeCloudStorageNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeFreeCloudStorageNum")
    
    
    return
}

func NewDescribeFreeCloudStorageNumResponse() (response *DescribeFreeCloudStorageNumResponse) {
    response = &DescribeFreeCloudStorageNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFreeCloudStorageNum
// 查询云存卡套餐信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFreeCloudStorageNum(request *DescribeFreeCloudStorageNumRequest) (response *DescribeFreeCloudStorageNumResponse, err error) {
    return c.DescribeFreeCloudStorageNumWithContext(context.Background(), request)
}

// DescribeFreeCloudStorageNum
// 查询云存卡套餐信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFreeCloudStorageNumWithContext(ctx context.Context, request *DescribeFreeCloudStorageNumRequest) (response *DescribeFreeCloudStorageNumResponse, err error) {
    if request == nil {
        request = NewDescribeFreeCloudStorageNumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFreeCloudStorageNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFreeCloudStorageNumResponse()
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
// 用于获取网关可绑定或解绑的子产品。
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
// 用于获取网关可绑定或解绑的子产品。
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

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// 公共实例过期时间 0001-01-01T00:00:00Z，公共实例是永久有效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// 公共实例过期时间 0001-01-01T00:00:00Z，公共实例是永久有效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOINSTANCE = "UnauthorizedOperation.NoPermissionToStudioInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_UNPAIDORDER = "UnsupportedOperation.UnpaidOrder"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
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

func NewDescribeP2PRouteRequest() (request *DescribeP2PRouteRequest) {
    request = &DescribeP2PRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeP2PRoute")
    
    
    return
}

func NewDescribeP2PRouteResponse() (response *DescribeP2PRouteResponse) {
    response = &DescribeP2PRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeP2PRoute
// 当前p2p线路
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeP2PRoute(request *DescribeP2PRouteRequest) (response *DescribeP2PRouteResponse, err error) {
    return c.DescribeP2PRouteWithContext(context.Background(), request)
}

// DescribeP2PRoute
// 当前p2p线路
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeP2PRouteWithContext(ctx context.Context, request *DescribeP2PRouteRequest) (response *DescribeP2PRouteResponse, err error) {
    if request == nil {
        request = NewDescribeP2PRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeP2PRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeP2PRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackageConsumeTaskRequest() (request *DescribePackageConsumeTaskRequest) {
    request = &DescribePackageConsumeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribePackageConsumeTask")
    
    
    return
}

func NewDescribePackageConsumeTaskResponse() (response *DescribePackageConsumeTaskResponse) {
    response = &DescribePackageConsumeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePackageConsumeTask
// 查询套餐消耗记录详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageConsumeTask(request *DescribePackageConsumeTaskRequest) (response *DescribePackageConsumeTaskResponse, err error) {
    return c.DescribePackageConsumeTaskWithContext(context.Background(), request)
}

// DescribePackageConsumeTask
// 查询套餐消耗记录详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageConsumeTaskWithContext(ctx context.Context, request *DescribePackageConsumeTaskRequest) (response *DescribePackageConsumeTaskResponse, err error) {
    if request == nil {
        request = NewDescribePackageConsumeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackageConsumeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackageConsumeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackageConsumeTasksRequest() (request *DescribePackageConsumeTasksRequest) {
    request = &DescribePackageConsumeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribePackageConsumeTasks")
    
    
    return
}

func NewDescribePackageConsumeTasksResponse() (response *DescribePackageConsumeTasksResponse) {
    response = &DescribePackageConsumeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePackageConsumeTasks
// 查询套餐消耗记录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageConsumeTasks(request *DescribePackageConsumeTasksRequest) (response *DescribePackageConsumeTasksResponse, err error) {
    return c.DescribePackageConsumeTasksWithContext(context.Background(), request)
}

// DescribePackageConsumeTasks
// 查询套餐消耗记录列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageConsumeTasksWithContext(ctx context.Context, request *DescribePackageConsumeTasksRequest) (response *DescribePackageConsumeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePackageConsumeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackageConsumeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackageConsumeTasksResponse()
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

func NewDescribeProductCloudStorageAIServiceRequest() (request *DescribeProductCloudStorageAIServiceRequest) {
    request = &DescribeProductCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeProductCloudStorageAIService")
    
    
    return
}

func NewDescribeProductCloudStorageAIServiceResponse() (response *DescribeProductCloudStorageAIServiceResponse) {
    response = &DescribeProductCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductCloudStorageAIService
// 查询指定产品的云存 AI 服务开通状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductCloudStorageAIService(request *DescribeProductCloudStorageAIServiceRequest) (response *DescribeProductCloudStorageAIServiceResponse, err error) {
    return c.DescribeProductCloudStorageAIServiceWithContext(context.Background(), request)
}

// DescribeProductCloudStorageAIService
// 查询指定产品的云存 AI 服务开通状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductCloudStorageAIServiceWithContext(ctx context.Context, request *DescribeProductCloudStorageAIServiceRequest) (response *DescribeProductCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewDescribeProductCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductCloudStorageAIServiceResponse()
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
// 获取位置空间中围栏告警事件列表。
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
// 获取位置空间中围栏告警事件列表。
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

func NewDescribeTWeSeeConfigRequest() (request *DescribeTWeSeeConfigRequest) {
    request = &DescribeTWeSeeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeTWeSeeConfig")
    
    
    return
}

func NewDescribeTWeSeeConfigResponse() (response *DescribeTWeSeeConfigResponse) {
    response = &DescribeTWeSeeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTWeSeeConfig
// 拉取 TWeSee 配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeTWeSeeConfig(request *DescribeTWeSeeConfigRequest) (response *DescribeTWeSeeConfigResponse, err error) {
    return c.DescribeTWeSeeConfigWithContext(context.Background(), request)
}

// DescribeTWeSeeConfig
// 拉取 TWeSee 配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeTWeSeeConfigWithContext(ctx context.Context, request *DescribeTWeSeeConfigRequest) (response *DescribeTWeSeeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeTWeSeeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTWeSeeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTWeSeeConfigResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewDescribeUnbindedDevicesRequest() (request *DescribeUnbindedDevicesRequest) {
    request = &DescribeUnbindedDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeUnbindedDevices")
    
    
    return
}

func NewDescribeUnbindedDevicesResponse() (response *DescribeUnbindedDevicesResponse) {
    response = &DescribeUnbindedDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnbindedDevices
// 获取未绑定的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeUnbindedDevices(request *DescribeUnbindedDevicesRequest) (response *DescribeUnbindedDevicesResponse, err error) {
    return c.DescribeUnbindedDevicesWithContext(context.Background(), request)
}

// DescribeUnbindedDevices
// 获取未绑定的设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeUnbindedDevicesWithContext(ctx context.Context, request *DescribeUnbindedDevicesRequest) (response *DescribeUnbindedDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeUnbindedDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnbindedDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnbindedDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoLicenseRequest() (request *DescribeVideoLicenseRequest) {
    request = &DescribeVideoLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DescribeVideoLicense")
    
    
    return
}

func NewDescribeVideoLicenseResponse() (response *DescribeVideoLicenseResponse) {
    response = &DescribeVideoLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoLicense
// 用于查询视频激活码统计概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeVideoLicense(request *DescribeVideoLicenseRequest) (response *DescribeVideoLicenseResponse, err error) {
    return c.DescribeVideoLicenseWithContext(context.Background(), request)
}

// DescribeVideoLicense
// 用于查询视频激活码统计概览
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DescribeVideoLicenseWithContext(ctx context.Context, request *DescribeVideoLicenseRequest) (response *DescribeVideoLicenseResponse, err error) {
    if request == nil {
        request = NewDescribeVideoLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoLicenseResponse()
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
//  UNSUPPORTEDOPERATION_CANNOTREBINDFAMILY = "UnsupportedOperation.CannotReBindFamily"
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
//  UNSUPPORTEDOPERATION_CANNOTREBINDFAMILY = "UnsupportedOperation.CannotReBindFamily"
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

func NewDismissRoomByStrRoomIdFromTRTCRequest() (request *DismissRoomByStrRoomIdFromTRTCRequest) {
    request = &DismissRoomByStrRoomIdFromTRTCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "DismissRoomByStrRoomIdFromTRTC")
    
    
    return
}

func NewDismissRoomByStrRoomIdFromTRTCResponse() (response *DismissRoomByStrRoomIdFromTRTCResponse) {
    response = &DismissRoomByStrRoomIdFromTRTCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DismissRoomByStrRoomIdFromTRTC
// 解散TRTC房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTRTCFAIL = "FailedOperation.RequestTRTCFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOTRTCROOM = "UnauthorizedOperation.NoPermissionToTRTCRoom"
//  UNSUPPORTEDOPERATION_TRTCSERVICENOTOPEN = "UnsupportedOperation.TRTCServiceNotOpen"
func (c *Client) DismissRoomByStrRoomIdFromTRTC(request *DismissRoomByStrRoomIdFromTRTCRequest) (response *DismissRoomByStrRoomIdFromTRTCResponse, err error) {
    return c.DismissRoomByStrRoomIdFromTRTCWithContext(context.Background(), request)
}

// DismissRoomByStrRoomIdFromTRTC
// 解散TRTC房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTRTCFAIL = "FailedOperation.RequestTRTCFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOTRTCROOM = "UnauthorizedOperation.NoPermissionToTRTCRoom"
//  UNSUPPORTEDOPERATION_TRTCSERVICENOTOPEN = "UnsupportedOperation.TRTCServiceNotOpen"
func (c *Client) DismissRoomByStrRoomIdFromTRTCWithContext(ctx context.Context, request *DismissRoomByStrRoomIdFromTRTCRequest) (response *DismissRoomByStrRoomIdFromTRTCResponse, err error) {
    if request == nil {
        request = NewDismissRoomByStrRoomIdFromTRTCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomIdFromTRTC require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdFromTRTCResponse()
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

func NewGenerateCloudStorageAIServiceTaskFileURLRequest() (request *GenerateCloudStorageAIServiceTaskFileURLRequest) {
    request = &GenerateCloudStorageAIServiceTaskFileURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GenerateCloudStorageAIServiceTaskFileURL")
    
    
    return
}

func NewGenerateCloudStorageAIServiceTaskFileURLResponse() (response *GenerateCloudStorageAIServiceTaskFileURLResponse) {
    response = &GenerateCloudStorageAIServiceTaskFileURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateCloudStorageAIServiceTaskFileURL
// 获取云存 AI 分析任务输出文件的下载地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASK = "ResourceNotFound.CloudStorageAIServiceTask"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASKFILE = "ResourceNotFound.CloudStorageAIServiceTaskFile"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCloudStorageAIServiceTaskFileURL(request *GenerateCloudStorageAIServiceTaskFileURLRequest) (response *GenerateCloudStorageAIServiceTaskFileURLResponse, err error) {
    return c.GenerateCloudStorageAIServiceTaskFileURLWithContext(context.Background(), request)
}

// GenerateCloudStorageAIServiceTaskFileURL
// 获取云存 AI 分析任务输出文件的下载地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASK = "ResourceNotFound.CloudStorageAIServiceTask"
//  RESOURCENOTFOUND_CLOUDSTORAGEAISERVICETASKFILE = "ResourceNotFound.CloudStorageAIServiceTaskFile"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCloudStorageAIServiceTaskFileURLWithContext(ctx context.Context, request *GenerateCloudStorageAIServiceTaskFileURLRequest) (response *GenerateCloudStorageAIServiceTaskFileURLResponse, err error) {
    if request == nil {
        request = NewGenerateCloudStorageAIServiceTaskFileURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateCloudStorageAIServiceTaskFileURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateCloudStorageAIServiceTaskFileURLResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateSignedVideoURLRequest() (request *GenerateSignedVideoURLRequest) {
    request = &GenerateSignedVideoURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GenerateSignedVideoURL")
    
    
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

func NewGetAuthMiniProgramAppListRequest() (request *GetAuthMiniProgramAppListRequest) {
    request = &GetAuthMiniProgramAppListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetAuthMiniProgramAppList")
    
    
    return
}

func NewGetAuthMiniProgramAppListResponse() (response *GetAuthMiniProgramAppListResponse) {
    response = &GetAuthMiniProgramAppListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAuthMiniProgramAppList
// 查询小程序列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthMiniProgramAppList(request *GetAuthMiniProgramAppListRequest) (response *GetAuthMiniProgramAppListResponse, err error) {
    return c.GetAuthMiniProgramAppListWithContext(context.Background(), request)
}

// GetAuthMiniProgramAppList
// 查询小程序列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAuthMiniProgramAppListWithContext(ctx context.Context, request *GetAuthMiniProgramAppListRequest) (response *GetAuthMiniProgramAppListResponse, err error) {
    if request == nil {
        request = NewGetAuthMiniProgramAppListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAuthMiniProgramAppList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAuthMiniProgramAppListResponse()
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
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
//  INVALIDPARAMETERVALUE_GETPRODUCTSLISTERROR = "InvalidParameterValue.GetProductsListError"
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
//  INVALIDPARAMETERVALUE_GETPRODUCTSLISTERROR = "InvalidParameterValue.GetProductsListError"
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

func NewGetDeviceSumStatisticsRequest() (request *GetDeviceSumStatisticsRequest) {
    request = &GetDeviceSumStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetDeviceSumStatistics")
    
    
    return
}

func NewGetDeviceSumStatisticsResponse() (response *GetDeviceSumStatisticsResponse) {
    response = &GetDeviceSumStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDeviceSumStatistics
// 拉取设备统计汇总数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
func (c *Client) GetDeviceSumStatistics(request *GetDeviceSumStatisticsRequest) (response *GetDeviceSumStatisticsResponse, err error) {
    return c.GetDeviceSumStatisticsWithContext(context.Background(), request)
}

// GetDeviceSumStatistics
// 拉取设备统计汇总数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
func (c *Client) GetDeviceSumStatisticsWithContext(ctx context.Context, request *GetDeviceSumStatisticsRequest) (response *GetDeviceSumStatisticsResponse, err error) {
    if request == nil {
        request = NewGetDeviceSumStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceSumStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceSumStatisticsResponse()
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
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
//  UNAUTHORIZEDOPERATION_NOVERIFIED = "UnauthorizedOperation.NoVerified"
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

func NewGetTWeCallActiveStatusRequest() (request *GetTWeCallActiveStatusRequest) {
    request = &GetTWeCallActiveStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetTWeCallActiveStatus")
    
    
    return
}

func NewGetTWeCallActiveStatusResponse() (response *GetTWeCallActiveStatusResponse) {
    response = &GetTWeCallActiveStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTWeCallActiveStatus
// 查询激活状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTWeCallActiveStatus(request *GetTWeCallActiveStatusRequest) (response *GetTWeCallActiveStatusResponse, err error) {
    return c.GetTWeCallActiveStatusWithContext(context.Background(), request)
}

// GetTWeCallActiveStatus
// 查询激活状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTWeCallActiveStatusWithContext(ctx context.Context, request *GetTWeCallActiveStatusRequest) (response *GetTWeCallActiveStatusResponse, err error) {
    if request == nil {
        request = NewGetTWeCallActiveStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTWeCallActiveStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTWeCallActiveStatusResponse()
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

func NewGetWechatDeviceTicketRequest() (request *GetWechatDeviceTicketRequest) {
    request = &GetWechatDeviceTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "GetWechatDeviceTicket")
    
    
    return
}

func NewGetWechatDeviceTicketResponse() (response *GetWechatDeviceTicketResponse) {
    response = &GetWechatDeviceTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWechatDeviceTicket
// 查询微信设备授权票据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetWechatDeviceTicket(request *GetWechatDeviceTicketRequest) (response *GetWechatDeviceTicketResponse, err error) {
    return c.GetWechatDeviceTicketWithContext(context.Background(), request)
}

// GetWechatDeviceTicket
// 查询微信设备授权票据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetWechatDeviceTicketWithContext(ctx context.Context, request *GetWechatDeviceTicketRequest) (response *GetWechatDeviceTicketResponse, err error) {
    if request == nil {
        request = NewGetWechatDeviceTicketRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWechatDeviceTicket require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWechatDeviceTicketResponse()
    err = c.Send(request, response)
    return
}

func NewInheritCloudStorageUserRequest() (request *InheritCloudStorageUserRequest) {
    request = &InheritCloudStorageUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InheritCloudStorageUser")
    
    
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

func NewInvokeAISearchServiceRequest() (request *InvokeAISearchServiceRequest) {
    request = &InvokeAISearchServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InvokeAISearchService")
    
    
    return
}

func NewInvokeAISearchServiceResponse() (response *InvokeAISearchServiceResponse) {
    response = &InvokeAISearchServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeAISearchService
// 视频语义搜索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeAISearchService(request *InvokeAISearchServiceRequest) (response *InvokeAISearchServiceResponse, err error) {
    return c.InvokeAISearchServiceWithContext(context.Background(), request)
}

// InvokeAISearchService
// 视频语义搜索
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeAISearchServiceWithContext(ctx context.Context, request *InvokeAISearchServiceRequest) (response *InvokeAISearchServiceResponse, err error) {
    if request == nil {
        request = NewInvokeAISearchServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeAISearchService require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeAISearchServiceResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeCloudStorageAIServiceTaskRequest() (request *InvokeCloudStorageAIServiceTaskRequest) {
    request = &InvokeCloudStorageAIServiceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InvokeCloudStorageAIServiceTask")
    
    
    return
}

func NewInvokeCloudStorageAIServiceTaskResponse() (response *InvokeCloudStorageAIServiceTaskResponse) {
    response = &InvokeCloudStorageAIServiceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeCloudStorageAIServiceTask
// 同步执行设备云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) InvokeCloudStorageAIServiceTask(request *InvokeCloudStorageAIServiceTaskRequest) (response *InvokeCloudStorageAIServiceTaskResponse, err error) {
    return c.InvokeCloudStorageAIServiceTaskWithContext(context.Background(), request)
}

// InvokeCloudStorageAIServiceTask
// 同步执行设备云存 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) InvokeCloudStorageAIServiceTaskWithContext(ctx context.Context, request *InvokeCloudStorageAIServiceTaskRequest) (response *InvokeCloudStorageAIServiceTaskResponse, err error) {
    if request == nil {
        request = NewInvokeCloudStorageAIServiceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeCloudStorageAIServiceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeCloudStorageAIServiceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeExternalSourceAIServiceTaskRequest() (request *InvokeExternalSourceAIServiceTaskRequest) {
    request = &InvokeExternalSourceAIServiceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InvokeExternalSourceAIServiceTask")
    
    
    return
}

func NewInvokeExternalSourceAIServiceTaskResponse() (response *InvokeExternalSourceAIServiceTaskResponse) {
    response = &InvokeExternalSourceAIServiceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeExternalSourceAIServiceTask
// 创建外部视频 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) InvokeExternalSourceAIServiceTask(request *InvokeExternalSourceAIServiceTaskRequest) (response *InvokeExternalSourceAIServiceTaskResponse, err error) {
    return c.InvokeExternalSourceAIServiceTaskWithContext(context.Background(), request)
}

// InvokeExternalSourceAIServiceTask
// 创建外部视频 AI 分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICETASKALREADYEXISTS = "FailedOperation.CloudStorageAIServiceTaskAlreadyExists"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
func (c *Client) InvokeExternalSourceAIServiceTaskWithContext(ctx context.Context, request *InvokeExternalSourceAIServiceTaskRequest) (response *InvokeExternalSourceAIServiceTaskResponse, err error) {
    if request == nil {
        request = NewInvokeExternalSourceAIServiceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeExternalSourceAIServiceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeExternalSourceAIServiceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeTWeSeeRecognitionTaskRequest() (request *InvokeTWeSeeRecognitionTaskRequest) {
    request = &InvokeTWeSeeRecognitionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InvokeTWeSeeRecognitionTask")
    
    
    return
}

func NewInvokeTWeSeeRecognitionTaskResponse() (response *InvokeTWeSeeRecognitionTaskResponse) {
    response = &InvokeTWeSeeRecognitionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeTWeSeeRecognitionTask
// 同步执行 TWeSee 语义理解任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeTWeSeeRecognitionTask(request *InvokeTWeSeeRecognitionTaskRequest) (response *InvokeTWeSeeRecognitionTaskResponse, err error) {
    return c.InvokeTWeSeeRecognitionTaskWithContext(context.Background(), request)
}

// InvokeTWeSeeRecognitionTask
// 同步执行 TWeSee 语义理解任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeTWeSeeRecognitionTaskWithContext(ctx context.Context, request *InvokeTWeSeeRecognitionTaskRequest) (response *InvokeTWeSeeRecognitionTaskResponse, err error) {
    if request == nil {
        request = NewInvokeTWeSeeRecognitionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeTWeSeeRecognitionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeTWeSeeRecognitionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeVideosKeywordsAnalyzerRequest() (request *InvokeVideosKeywordsAnalyzerRequest) {
    request = &InvokeVideosKeywordsAnalyzerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "InvokeVideosKeywordsAnalyzer")
    
    
    return
}

func NewInvokeVideosKeywordsAnalyzerResponse() (response *InvokeVideosKeywordsAnalyzerResponse) {
    response = &InvokeVideosKeywordsAnalyzerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeVideosKeywordsAnalyzer
// 获取某个时间段的视频内容关键字
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeVideosKeywordsAnalyzer(request *InvokeVideosKeywordsAnalyzerRequest) (response *InvokeVideosKeywordsAnalyzerResponse, err error) {
    return c.InvokeVideosKeywordsAnalyzerWithContext(context.Background(), request)
}

// InvokeVideosKeywordsAnalyzer
// 获取某个时间段的视频内容关键字
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT_CLOUDSTORAGEAISERVICETASKQUOTAINSUFFICIENT = "ResourceInsufficient.CloudStorageAIServiceTaskQuotaInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InvokeVideosKeywordsAnalyzerWithContext(ctx context.Context, request *InvokeVideosKeywordsAnalyzerRequest) (response *InvokeVideosKeywordsAnalyzerResponse, err error) {
    if request == nil {
        request = NewInvokeVideosKeywordsAnalyzerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeVideosKeywordsAnalyzer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeVideosKeywordsAnalyzerResponse()
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// 更新应用信息
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
//  INVALIDPARAMETERVALUE_PUSHENVIRONMENTINVALID = "InvalidParameterValue.PushEnvironmentInvalid"
//  INVALIDPARAMETERVALUE_TPNSANDROIDVALIDATIONFAILED = "InvalidParameterValue.TPNSAndroidValidationFailed"
//  INVALIDPARAMETERVALUE_TPNSIOSVALIDATIONFAILED = "InvalidParameterValue.TPNSiOSValidationFailed"
//  RESOURCENOTFOUND_APPNOTEXISTS = "ResourceNotFound.AppNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// 更新应用信息
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
//  INVALIDPARAMETERVALUE_PUSHENVIRONMENTINVALID = "InvalidParameterValue.PushEnvironmentInvalid"
//  INVALIDPARAMETERVALUE_TPNSANDROIDVALIDATIONFAILED = "InvalidParameterValue.TPNSAndroidValidationFailed"
//  INVALIDPARAMETERVALUE_TPNSIOSVALIDATIONFAILED = "InvalidParameterValue.TPNSiOSValidationFailed"
//  RESOURCENOTFOUND_APPNOTEXISTS = "ResourceNotFound.AppNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudStorageAIServiceRequest() (request *ModifyCloudStorageAIServiceRequest) {
    request = &ModifyCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyCloudStorageAIService")
    
    
    return
}

func NewModifyCloudStorageAIServiceResponse() (response *ModifyCloudStorageAIServiceResponse) {
    response = &ModifyCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudStorageAIService
// 修改指定设备的云存 AI 服务参数配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudStorageAIService(request *ModifyCloudStorageAIServiceRequest) (response *ModifyCloudStorageAIServiceResponse, err error) {
    return c.ModifyCloudStorageAIServiceWithContext(context.Background(), request)
}

// ModifyCloudStorageAIService
// 修改指定设备的云存 AI 服务参数配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudStorageAIServiceWithContext(ctx context.Context, request *ModifyCloudStorageAIServiceRequest) (response *ModifyCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewModifyCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudStorageAIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudStorageAIServiceCallbackRequest() (request *ModifyCloudStorageAIServiceCallbackRequest) {
    request = &ModifyCloudStorageAIServiceCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyCloudStorageAIServiceCallback")
    
    
    return
}

func NewModifyCloudStorageAIServiceCallbackResponse() (response *ModifyCloudStorageAIServiceCallbackResponse) {
    response = &ModifyCloudStorageAIServiceCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudStorageAIServiceCallback
// 修改云存AI分析回调配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudStorageAIServiceCallback(request *ModifyCloudStorageAIServiceCallbackRequest) (response *ModifyCloudStorageAIServiceCallbackResponse, err error) {
    return c.ModifyCloudStorageAIServiceCallbackWithContext(context.Background(), request)
}

// ModifyCloudStorageAIServiceCallback
// 修改云存AI分析回调配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDSTORAGEAISERVICENOTENABLED = "FailedOperation.CloudStorageAIServiceNotEnabled"
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudStorageAIServiceCallbackWithContext(ctx context.Context, request *ModifyCloudStorageAIServiceCallbackRequest) (response *ModifyCloudStorageAIServiceCallbackResponse, err error) {
    if request == nil {
        request = NewModifyCloudStorageAIServiceCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudStorageAIServiceCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudStorageAIServiceCallbackResponse()
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
// 更新围栏绑定信息。
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
// 更新围栏绑定信息。
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
// 更新围栏。
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
// 更新围栏。
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
// 更新位置空间。
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
// 更新位置空间。
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

func NewModifyProductCloudStorageAIServiceRequest() (request *ModifyProductCloudStorageAIServiceRequest) {
    request = &ModifyProductCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyProductCloudStorageAIService")
    
    
    return
}

func NewModifyProductCloudStorageAIServiceResponse() (response *ModifyProductCloudStorageAIServiceResponse) {
    response = &ModifyProductCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProductCloudStorageAIService
// 修改指定产品的云存 AI 服务开通状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProductCloudStorageAIService(request *ModifyProductCloudStorageAIServiceRequest) (response *ModifyProductCloudStorageAIServiceResponse, err error) {
    return c.ModifyProductCloudStorageAIServiceWithContext(context.Background(), request)
}

// ModifyProductCloudStorageAIService
// 修改指定产品的云存 AI 服务开通状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProductCloudStorageAIServiceWithContext(ctx context.Context, request *ModifyProductCloudStorageAIServiceRequest) (response *ModifyProductCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewModifyProductCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProductCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProductCloudStorageAIServiceResponse()
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
// 修改项目。
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
// 修改项目。
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
//  UNSUPPORTEDOPERATION_VIDEOPRODUCTNOTEXIST = "UnsupportedOperation.VideoProductNotExist"
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
//  UNSUPPORTEDOPERATION_VIDEOPRODUCTNOTEXIST = "UnsupportedOperation.VideoProductNotExist"
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

func NewModifyTWeSeeConfigRequest() (request *ModifyTWeSeeConfigRequest) {
    request = &ModifyTWeSeeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ModifyTWeSeeConfig")
    
    
    return
}

func NewModifyTWeSeeConfigResponse() (response *ModifyTWeSeeConfigResponse) {
    response = &ModifyTWeSeeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTWeSeeConfig
// 修改 TWeSee 配置
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
//  UNSUPPORTEDOPERATION_VIDEOPRODUCTNOTEXIST = "UnsupportedOperation.VideoProductNotExist"
func (c *Client) ModifyTWeSeeConfig(request *ModifyTWeSeeConfigRequest) (response *ModifyTWeSeeConfigResponse, err error) {
    return c.ModifyTWeSeeConfigWithContext(context.Background(), request)
}

// ModifyTWeSeeConfig
// 修改 TWeSee 配置
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
//  UNSUPPORTEDOPERATION_VIDEOPRODUCTNOTEXIST = "UnsupportedOperation.VideoProductNotExist"
func (c *Client) ModifyTWeSeeConfigWithContext(ctx context.Context, request *ModifyTWeSeeConfigRequest) (response *ModifyTWeSeeConfigResponse, err error) {
    if request == nil {
        request = NewModifyTWeSeeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTWeSeeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTWeSeeConfigResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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

func NewPauseTWeCallDeviceRequest() (request *PauseTWeCallDeviceRequest) {
    request = &PauseTWeCallDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PauseTWeCallDevice")
    
    
    return
}

func NewPauseTWeCallDeviceResponse() (response *PauseTWeCallDeviceResponse) {
    response = &PauseTWeCallDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseTWeCallDevice
// 暂停设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PauseTWeCallDevice(request *PauseTWeCallDeviceRequest) (response *PauseTWeCallDeviceResponse, err error) {
    return c.PauseTWeCallDeviceWithContext(context.Background(), request)
}

// PauseTWeCallDevice
// 暂停设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PauseTWeCallDeviceWithContext(ctx context.Context, request *PauseTWeCallDeviceRequest) (response *PauseTWeCallDeviceResponse, err error) {
    if request == nil {
        request = NewPauseTWeCallDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseTWeCallDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseTWeCallDeviceResponse()
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
// 发布广播消息、发布RRPC消息属于早期服务，目前已停止维护，需要从官网下线。
//
// 
//
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
// 发布广播消息、发布RRPC消息属于早期服务，目前已停止维护，需要从官网下线。
//
// 
//
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

func NewPublishFirmwareUpdateMessageRequest() (request *PublishFirmwareUpdateMessageRequest) {
    request = &PublishFirmwareUpdateMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "PublishFirmwareUpdateMessage")
    
    
    return
}

func NewPublishFirmwareUpdateMessageResponse() (response *PublishFirmwareUpdateMessageResponse) {
    response = &PublishFirmwareUpdateMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PublishFirmwareUpdateMessage
// 本接口（PublishFirmwareUpdateMessage）用于用户确认升级后，云端向设备发起固件升级请求。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEINFOOUTDATED = "FailedOperation.DeviceInfoOutdated"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOPERMISSION = "InvalidParameterValue.NoPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) PublishFirmwareUpdateMessage(request *PublishFirmwareUpdateMessageRequest) (response *PublishFirmwareUpdateMessageResponse, err error) {
    return c.PublishFirmwareUpdateMessageWithContext(context.Background(), request)
}

// PublishFirmwareUpdateMessage
// 本接口（PublishFirmwareUpdateMessage）用于用户确认升级后，云端向设备发起固件升级请求。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEINFOOUTDATED = "FailedOperation.DeviceInfoOutdated"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOPERMISSION = "InvalidParameterValue.NoPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) PublishFirmwareUpdateMessageWithContext(ctx context.Context, request *PublishFirmwareUpdateMessageRequest) (response *PublishFirmwareUpdateMessageResponse, err error) {
    if request == nil {
        request = NewPublishFirmwareUpdateMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishFirmwareUpdateMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishFirmwareUpdateMessageResponse()
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
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
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
//  UNSUPPORTEDOPERATION_LORANOUPLINK = "UnsupportedOperation.LoRaNoUpLink"
//  UNSUPPORTEDOPERATION_LORANOTACTIVATE = "UnsupportedOperation.LoRaNotActivate"
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
// 发布广播消息、发布RRPC消息属于早期服务，目前已停止维护，需要从官网下线。
//
// 
//
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) PublishRRPCMessage(request *PublishRRPCMessageRequest) (response *PublishRRPCMessageResponse, err error) {
    return c.PublishRRPCMessageWithContext(context.Background(), request)
}

// PublishRRPCMessage
// 发布广播消息、发布RRPC消息属于早期服务，目前已停止维护，需要从官网下线。
//
// 
//
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
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

func NewRemoveUserByRoomIdFromTRTCRequest() (request *RemoveUserByRoomIdFromTRTCRequest) {
    request = &RemoveUserByRoomIdFromTRTCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "RemoveUserByRoomIdFromTRTC")
    
    
    return
}

func NewRemoveUserByRoomIdFromTRTCResponse() (response *RemoveUserByRoomIdFromTRTCResponse) {
    response = &RemoveUserByRoomIdFromTRTCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUserByRoomIdFromTRTC
// TRTC操作，将用户从房间移出
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTRTCFAIL = "FailedOperation.RequestTRTCFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOTRTCROOM = "UnauthorizedOperation.NoPermissionToTRTCRoom"
func (c *Client) RemoveUserByRoomIdFromTRTC(request *RemoveUserByRoomIdFromTRTCRequest) (response *RemoveUserByRoomIdFromTRTCResponse, err error) {
    return c.RemoveUserByRoomIdFromTRTCWithContext(context.Background(), request)
}

// RemoveUserByRoomIdFromTRTC
// TRTC操作，将用户从房间移出
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTRTCFAIL = "FailedOperation.RequestTRTCFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOTRTCROOM = "UnauthorizedOperation.NoPermissionToTRTCRoom"
func (c *Client) RemoveUserByRoomIdFromTRTCWithContext(ctx context.Context, request *RemoveUserByRoomIdFromTRTCRequest) (response *RemoveUserByRoomIdFromTRTCResponse, err error) {
    if request == nil {
        request = NewRemoveUserByRoomIdFromTRTCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByRoomIdFromTRTC require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByRoomIdFromTRTCResponse()
    err = c.Send(request, response)
    return
}

func NewResetCloudStorageRequest() (request *ResetCloudStorageRequest) {
    request = &ResetCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ResetCloudStorage")
    
    
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

func NewResetCloudStorageAIServiceRequest() (request *ResetCloudStorageAIServiceRequest) {
    request = &ResetCloudStorageAIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ResetCloudStorageAIService")
    
    
    return
}

func NewResetCloudStorageAIServiceResponse() (response *ResetCloudStorageAIServiceResponse) {
    response = &ResetCloudStorageAIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetCloudStorageAIService
// 重置指定设备的云存 AI 服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorageAIService(request *ResetCloudStorageAIServiceRequest) (response *ResetCloudStorageAIServiceResponse, err error) {
    return c.ResetCloudStorageAIServiceWithContext(context.Background(), request)
}

// ResetCloudStorageAIService
// 重置指定设备的云存 AI 服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTIOTVIDEOSERVICENOTENABLED = "FailedOperation.ProductIotVideoServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorageAIServiceWithContext(ctx context.Context, request *ResetCloudStorageAIServiceRequest) (response *ResetCloudStorageAIServiceResponse, err error) {
    if request == nil {
        request = NewResetCloudStorageAIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetCloudStorageAIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetCloudStorageAIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewResetCloudStorageEventRequest() (request *ResetCloudStorageEventRequest) {
    request = &ResetCloudStorageEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ResetCloudStorageEvent")
    
    
    return
}

func NewResetCloudStorageEventResponse() (response *ResetCloudStorageEventResponse) {
    response = &ResetCloudStorageEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetCloudStorageEvent
// 重置云存事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorageEvent(request *ResetCloudStorageEventRequest) (response *ResetCloudStorageEventResponse, err error) {
    return c.ResetCloudStorageEventWithContext(context.Background(), request)
}

// ResetCloudStorageEvent
// 重置云存事件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetCloudStorageEventWithContext(ctx context.Context, request *ResetCloudStorageEventRequest) (response *ResetCloudStorageEventResponse, err error) {
    if request == nil {
        request = NewResetCloudStorageEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetCloudStorageEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetCloudStorageEventResponse()
    err = c.Send(request, response)
    return
}

func NewResetTWeCallDeviceRequest() (request *ResetTWeCallDeviceRequest) {
    request = &ResetTWeCallDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ResetTWeCallDevice")
    
    
    return
}

func NewResetTWeCallDeviceResponse() (response *ResetTWeCallDeviceResponse) {
    response = &ResetTWeCallDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetTWeCallDevice
// 重置设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetTWeCallDevice(request *ResetTWeCallDeviceRequest) (response *ResetTWeCallDeviceResponse, err error) {
    return c.ResetTWeCallDeviceWithContext(context.Background(), request)
}

// ResetTWeCallDevice
// 重置设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetTWeCallDeviceWithContext(ctx context.Context, request *ResetTWeCallDeviceRequest) (response *ResetTWeCallDeviceResponse, err error) {
    if request == nil {
        request = NewResetTWeCallDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetTWeCallDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetTWeCallDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewResumeWeCallDeviceRequest() (request *ResumeWeCallDeviceRequest) {
    request = &ResumeWeCallDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "ResumeWeCallDevice")
    
    
    return
}

func NewResumeWeCallDeviceResponse() (response *ResumeWeCallDeviceResponse) {
    response = &ResumeWeCallDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeWeCallDevice
// 恢复设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeWeCallDevice(request *ResumeWeCallDeviceRequest) (response *ResumeWeCallDeviceResponse, err error) {
    return c.ResumeWeCallDeviceWithContext(context.Background(), request)
}

// ResumeWeCallDevice
// 恢复设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeWeCallDeviceWithContext(ctx context.Context, request *ResumeWeCallDeviceRequest) (response *ResumeWeCallDeviceResponse, err error) {
    if request == nil {
        request = NewResumeWeCallDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeWeCallDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeWeCallDeviceResponse()
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
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
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
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
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

func NewTransferCloudStorageRequest() (request *TransferCloudStorageRequest) {
    request = &TransferCloudStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "TransferCloudStorage")
    
    
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

func NewTransferTWeCallDeviceRequest() (request *TransferTWeCallDeviceRequest) {
    request = &TransferTWeCallDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "TransferTWeCallDevice")
    
    
    return
}

func NewTransferTWeCallDeviceResponse() (response *TransferTWeCallDeviceResponse) {
    response = &TransferTWeCallDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TransferTWeCallDevice
// 转移设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TransferTWeCallDevice(request *TransferTWeCallDeviceRequest) (response *TransferTWeCallDeviceResponse, err error) {
    return c.TransferTWeCallDeviceWithContext(context.Background(), request)
}

// TransferTWeCallDevice
// 转移设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TransferTWeCallDeviceWithContext(ctx context.Context, request *TransferTWeCallDeviceRequest) (response *TransferTWeCallDeviceResponse, err error) {
    if request == nil {
        request = NewTransferTWeCallDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferTWeCallDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferTWeCallDeviceResponse()
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
// 批量解绑子产品。
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
// 批量解绑子产品。
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

func NewUpdateDeviceTWeCallAuthorizeStatusRequest() (request *UpdateDeviceTWeCallAuthorizeStatusRequest) {
    request = &UpdateDeviceTWeCallAuthorizeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotexplorer", APIVersion, "UpdateDeviceTWeCallAuthorizeStatus")
    
    
    return
}

func NewUpdateDeviceTWeCallAuthorizeStatusResponse() (response *UpdateDeviceTWeCallAuthorizeStatusResponse) {
    response = &UpdateDeviceTWeCallAuthorizeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDeviceTWeCallAuthorizeStatus
// 更新用户对设备的TweCall授权状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateDeviceTWeCallAuthorizeStatus(request *UpdateDeviceTWeCallAuthorizeStatusRequest) (response *UpdateDeviceTWeCallAuthorizeStatusResponse, err error) {
    return c.UpdateDeviceTWeCallAuthorizeStatusWithContext(context.Background(), request)
}

// UpdateDeviceTWeCallAuthorizeStatus
// 更新用户对设备的TweCall授权状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateDeviceTWeCallAuthorizeStatusWithContext(ctx context.Context, request *UpdateDeviceTWeCallAuthorizeStatusRequest) (response *UpdateDeviceTWeCallAuthorizeStatusResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceTWeCallAuthorizeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceTWeCallAuthorizeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceTWeCallAuthorizeStatusResponse()
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
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
