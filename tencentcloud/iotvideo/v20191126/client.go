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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewClearDeviceActiveCodeRequest() (request *ClearDeviceActiveCodeRequest) {
    request = &ClearDeviceActiveCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ClearDeviceActiveCode")
    
    
    return
}

func NewClearDeviceActiveCodeResponse() (response *ClearDeviceActiveCodeResponse) {
    response = &ClearDeviceActiveCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearDeviceActiveCode
// 清除设备激活码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ClearDeviceActiveCode(request *ClearDeviceActiveCodeRequest) (response *ClearDeviceActiveCodeResponse, err error) {
    return c.ClearDeviceActiveCodeWithContext(context.Background(), request)
}

// ClearDeviceActiveCode
// 清除设备激活码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ClearDeviceActiveCodeWithContext(ctx context.Context, request *ClearDeviceActiveCodeRequest) (response *ClearDeviceActiveCodeResponse, err error) {
    if request == nil {
        request = NewClearDeviceActiveCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearDeviceActiveCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearDeviceActiveCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAnonymousAccessTokenRequest() (request *CreateAnonymousAccessTokenRequest) {
    request = &CreateAnonymousAccessTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateAnonymousAccessToken")
    
    
    return
}

func NewCreateAnonymousAccessTokenResponse() (response *CreateAnonymousAccessTokenResponse) {
    response = &CreateAnonymousAccessTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAnonymousAccessToken
// 创建匿名访问Token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAnonymousAccessToken(request *CreateAnonymousAccessTokenRequest) (response *CreateAnonymousAccessTokenResponse, err error) {
    return c.CreateAnonymousAccessTokenWithContext(context.Background(), request)
}

// CreateAnonymousAccessToken
// 创建匿名访问Token
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAnonymousAccessTokenWithContext(ctx context.Context, request *CreateAnonymousAccessTokenRequest) (response *CreateAnonymousAccessTokenResponse, err error) {
    if request == nil {
        request = NewCreateAnonymousAccessTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAnonymousAccessToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAnonymousAccessTokenResponse()
    err = c.Send(request, response)
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

// CreateAppUsr
// 本接口（CreateAppUsr）用于接收由厂商云发送过来的注册请求,建立厂商云终端用户与IoT Video终端用户的映射关系。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppUsr(request *CreateAppUsrRequest) (response *CreateAppUsrResponse, err error) {
    return c.CreateAppUsrWithContext(context.Background(), request)
}

// CreateAppUsr
// 本接口（CreateAppUsr）用于接收由厂商云发送过来的注册请求,建立厂商云终端用户与IoT Video终端用户的映射关系。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppUsrWithContext(ctx context.Context, request *CreateAppUsrRequest) (response *CreateAppUsrResponse, err error) {
    if request == nil {
        request = NewCreateAppUsrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppUsr require credential")
    }

    request.SetContext(ctx)
    
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

// CreateBinding
// 本接口（CreateBinding）用于终端用户和设备进行绑定，具体的应用场景如下：
//
//     终端用户与设备具有“强关联”关系。用户与设备绑定之后，用户终端即具备了该设备的访问权限,访问或操作设备时，无需获取设备访问Token。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBinding(request *CreateBindingRequest) (response *CreateBindingResponse, err error) {
    return c.CreateBindingWithContext(context.Background(), request)
}

// CreateBinding
// 本接口（CreateBinding）用于终端用户和设备进行绑定，具体的应用场景如下：
//
//     终端用户与设备具有“强关联”关系。用户与设备绑定之后，用户终端即具备了该设备的访问权限,访问或操作设备时，无需获取设备访问Token。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBindingWithContext(ctx context.Context, request *CreateBindingRequest) (response *CreateBindingResponse, err error) {
    if request == nil {
        request = NewCreateBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBinding require credential")
    }

    request.SetContext(ctx)
    
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

// CreateDevToken
// 本接口（CreateDevToken）用于以下场景：
//
// 终端用户与设备没有强绑定关联关系;
//
// 允许终端用户短时或一次性临时访问设备;
//
// 当终端用户与设备有强绑定关系时，可以不用调用此接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDevToken(request *CreateDevTokenRequest) (response *CreateDevTokenResponse, err error) {
    return c.CreateDevTokenWithContext(context.Background(), request)
}

// CreateDevToken
// 本接口（CreateDevToken）用于以下场景：
//
// 终端用户与设备没有强绑定关联关系;
//
// 允许终端用户短时或一次性临时访问设备;
//
// 当终端用户与设备有强绑定关系时，可以不用调用此接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDevTokenWithContext(ctx context.Context, request *CreateDevTokenRequest) (response *CreateDevTokenResponse, err error) {
    if request == nil {
        request = NewCreateDevTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDevToken require credential")
    }

    request.SetContext(ctx)
    
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

// CreateDevices
// 本接口（CreateDevices）用于批量创建新的物联网视频通信设备。
//
// 注意：腾讯云不会对设备私钥进行保存，请自行保管好您的设备私钥。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDevices(request *CreateDevicesRequest) (response *CreateDevicesResponse, err error) {
    return c.CreateDevicesWithContext(context.Background(), request)
}

// CreateDevices
// 本接口（CreateDevices）用于批量创建新的物联网视频通信设备。
//
// 注意：腾讯云不会对设备私钥进行保存，请自行保管好您的设备私钥。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDevicesWithContext(ctx context.Context, request *CreateDevicesRequest) (response *CreateDevicesResponse, err error) {
    if request == nil {
        request = NewCreateDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDevices require credential")
    }

    request.SetContext(ctx)
    
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

// CreateGencode
// 本接口（CreateGencode）用于生成设备物模型源代码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGencode(request *CreateGencodeRequest) (response *CreateGencodeResponse, err error) {
    return c.CreateGencodeWithContext(context.Background(), request)
}

// CreateGencode
// 本接口（CreateGencode）用于生成设备物模型源代码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGencodeWithContext(ctx context.Context, request *CreateGencodeRequest) (response *CreateGencodeResponse, err error) {
    if request == nil {
        request = NewCreateGencodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGencode require credential")
    }

    request.SetContext(ctx)
    
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

// CreateIotDataType
// 本接口（CreateIotDataType）用于创建自定义物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotDataType(request *CreateIotDataTypeRequest) (response *CreateIotDataTypeResponse, err error) {
    return c.CreateIotDataTypeWithContext(context.Background(), request)
}

// CreateIotDataType
// 本接口（CreateIotDataType）用于创建自定义物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotDataTypeWithContext(ctx context.Context, request *CreateIotDataTypeRequest) (response *CreateIotDataTypeResponse, err error) {
    if request == nil {
        request = NewCreateIotDataTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIotDataType require credential")
    }

    request.SetContext(ctx)
    
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

// CreateIotModel
// 本接口（CreateIotModel）用于定义的物模型提交。
//
// 该接口实现了物模型草稿箱的功能，保存用户最后一次编辑的物模型数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotModel(request *CreateIotModelRequest) (response *CreateIotModelResponse, err error) {
    return c.CreateIotModelWithContext(context.Background(), request)
}

// CreateIotModel
// 本接口（CreateIotModel）用于定义的物模型提交。
//
// 该接口实现了物模型草稿箱的功能，保存用户最后一次编辑的物模型数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIotModelWithContext(ctx context.Context, request *CreateIotModelRequest) (response *CreateIotModelResponse, err error) {
    if request == nil {
        request = NewCreateIotModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIotModel require credential")
    }

    request.SetContext(ctx)
    
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

// CreateProduct
// 本接口（CreateProduct）用于创建一个新的物联网智能视频产品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    return c.CreateProductWithContext(context.Background(), request)
}

// CreateProduct
// 本接口（CreateProduct）用于创建一个新的物联网智能视频产品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// CreateStorage
// 该接口已经停止维护，请勿使用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStorage(request *CreateStorageRequest) (response *CreateStorageResponse, err error) {
    return c.CreateStorageWithContext(context.Background(), request)
}

// CreateStorage
// 该接口已经停止维护，请勿使用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStorageWithContext(ctx context.Context, request *CreateStorageRequest) (response *CreateStorageResponse, err error) {
    if request == nil {
        request = NewCreateStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStorageServiceRequest() (request *CreateStorageServiceRequest) {
    request = &CreateStorageServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateStorageService")
    
    
    return
}

func NewCreateStorageServiceResponse() (response *CreateStorageServiceResponse) {
    response = &CreateStorageServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStorageService
// 购买云存服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStorageService(request *CreateStorageServiceRequest) (response *CreateStorageServiceResponse, err error) {
    return c.CreateStorageServiceWithContext(context.Background(), request)
}

// CreateStorageService
// 购买云存服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateStorageServiceWithContext(ctx context.Context, request *CreateStorageServiceRequest) (response *CreateStorageServiceResponse, err error) {
    if request == nil {
        request = NewCreateStorageServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorageService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageServiceResponse()
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

// CreateTraceIds
// 本接口（CreateTraceIds）用于将设备加到日志跟踪白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTraceIds(request *CreateTraceIdsRequest) (response *CreateTraceIdsResponse, err error) {
    return c.CreateTraceIdsWithContext(context.Background(), request)
}

// CreateTraceIds
// 本接口（CreateTraceIds）用于将设备加到日志跟踪白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTraceIdsWithContext(ctx context.Context, request *CreateTraceIdsRequest) (response *CreateTraceIdsResponse, err error) {
    if request == nil {
        request = NewCreateTraceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTraceIds require credential")
    }

    request.SetContext(ctx)
    
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

// CreateUploadPath
// 本接口（CreateUploadPath）用于获取固件上传路径。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUploadPath(request *CreateUploadPathRequest) (response *CreateUploadPathResponse, err error) {
    return c.CreateUploadPathWithContext(context.Background(), request)
}

// CreateUploadPath
// 本接口（CreateUploadPath）用于获取固件上传路径。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUploadPathWithContext(ctx context.Context, request *CreateUploadPathRequest) (response *CreateUploadPathResponse, err error) {
    if request == nil {
        request = NewCreateUploadPathRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUploadPath require credential")
    }

    request.SetContext(ctx)
    
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

// CreateUsrToken
// 本接口（CreateUsrToken）用于终端用户获取IoT Video平台的accessToken，初始化SDK,连接到IoT Video接入服务器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUsrToken(request *CreateUsrTokenRequest) (response *CreateUsrTokenResponse, err error) {
    return c.CreateUsrTokenWithContext(context.Background(), request)
}

// CreateUsrToken
// 本接口（CreateUsrToken）用于终端用户获取IoT Video平台的accessToken，初始化SDK,连接到IoT Video接入服务器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUsrTokenWithContext(ctx context.Context, request *CreateUsrTokenRequest) (response *CreateUsrTokenResponse, err error) {
    if request == nil {
        request = NewCreateUsrTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsrToken require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteAppUsr
// 本接口（DeleteAppUsr）用于删除终端用户。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAppUsr(request *DeleteAppUsrRequest) (response *DeleteAppUsrResponse, err error) {
    return c.DeleteAppUsrWithContext(context.Background(), request)
}

// DeleteAppUsr
// 本接口（DeleteAppUsr）用于删除终端用户。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAppUsrWithContext(ctx context.Context, request *DeleteAppUsrRequest) (response *DeleteAppUsrResponse, err error) {
    if request == nil {
        request = NewDeleteAppUsrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppUsr require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteBinding
// 本接口（DeleteBinding）用于终端用户和设备进行解绑定。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteBinding(request *DeleteBindingRequest) (response *DeleteBindingResponse, err error) {
    return c.DeleteBindingWithContext(context.Background(), request)
}

// DeleteBinding
// 本接口（DeleteBinding）用于终端用户和设备进行解绑定。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteBindingWithContext(ctx context.Context, request *DeleteBindingRequest) (response *DeleteBindingResponse, err error) {
    if request == nil {
        request = NewDeleteBindingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBinding require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteDevice
// 本接口（DeleteDevice）用于删除设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DeleteIotDataType
// 本接口（DeleteIotDataType）用于删除自定义物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDataType(request *DeleteIotDataTypeRequest) (response *DeleteIotDataTypeResponse, err error) {
    return c.DeleteIotDataTypeWithContext(context.Background(), request)
}

// DeleteIotDataType
// 本接口（DeleteIotDataType）用于删除自定义物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIotDataTypeWithContext(ctx context.Context, request *DeleteIotDataTypeRequest) (response *DeleteIotDataTypeResponse, err error) {
    if request == nil {
        request = NewDeleteIotDataTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIotDataType require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteMessageQueue
// 本接口（DeleteMessageQueue）用于删除物联网智能视频产品的转发消息配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMessageQueue(request *DeleteMessageQueueRequest) (response *DeleteMessageQueueResponse, err error) {
    return c.DeleteMessageQueueWithContext(context.Background(), request)
}

// DeleteMessageQueue
// 本接口（DeleteMessageQueue）用于删除物联网智能视频产品的转发消息配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMessageQueueWithContext(ctx context.Context, request *DeleteMessageQueueRequest) (response *DeleteMessageQueueResponse, err error) {
    if request == nil {
        request = NewDeleteMessageQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMessageQueue require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteOtaVersion
// 本接口（DeleteOtaVersion）用于删除固件版本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOtaVersion(request *DeleteOtaVersionRequest) (response *DeleteOtaVersionResponse, err error) {
    return c.DeleteOtaVersionWithContext(context.Background(), request)
}

// DeleteOtaVersion
// 本接口（DeleteOtaVersion）用于删除固件版本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteOtaVersionWithContext(ctx context.Context, request *DeleteOtaVersionRequest) (response *DeleteOtaVersionResponse, err error) {
    if request == nil {
        request = NewDeleteOtaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOtaVersion require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网智能视频产品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网智能视频产品。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DeleteTraceIds
// 本接口（DeleteTraceIds）用于将设备从日志跟踪白名单中删除，该接口可批量操作，最多支持同时操作100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTraceIds(request *DeleteTraceIdsRequest) (response *DeleteTraceIdsResponse, err error) {
    return c.DeleteTraceIdsWithContext(context.Background(), request)
}

// DeleteTraceIds
// 本接口（DeleteTraceIds）用于将设备从日志跟踪白名单中删除，该接口可批量操作，最多支持同时操作100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTraceIdsWithContext(ctx context.Context, request *DeleteTraceIdsRequest) (response *DeleteTraceIdsResponse, err error) {
    if request == nil {
        request = NewDeleteTraceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTraceIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTraceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDeliverStorageServiceRequest() (request *DeliverStorageServiceRequest) {
    request = &DeliverStorageServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DeliverStorageService")
    
    
    return
}

func NewDeliverStorageServiceResponse() (response *DeliverStorageServiceResponse) {
    response = &DeliverStorageServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeliverStorageService
// 将已购买的云存服务转移到另一设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeliverStorageService(request *DeliverStorageServiceRequest) (response *DeliverStorageServiceResponse, err error) {
    return c.DeliverStorageServiceWithContext(context.Background(), request)
}

// DeliverStorageService
// 将已购买的云存服务转移到另一设备
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeliverStorageServiceWithContext(ctx context.Context, request *DeliverStorageServiceRequest) (response *DeliverStorageServiceResponse, err error) {
    if request == nil {
        request = NewDeliverStorageServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeliverStorageService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeliverStorageServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountBalanceRequest() (request *DescribeAccountBalanceRequest) {
    request = &DescribeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeAccountBalance")
    
    
    return
}

func NewDescribeAccountBalanceResponse() (response *DescribeAccountBalanceResponse) {
    response = &DescribeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountBalance
// 客户可通过本接口获取账户余额信息, 默认接口请求频率限制：1次/秒
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    return c.DescribeAccountBalanceWithContext(context.Background(), request)
}

// DescribeAccountBalance
// 客户可通过本接口获取账户余额信息, 默认接口请求频率限制：1次/秒
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalanceWithContext(ctx context.Context, request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountBalanceResponse()
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

// DescribeBindDev
// 本接口（DescribeBindDev）用于查询终端用户绑定的设备列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindDev(request *DescribeBindDevRequest) (response *DescribeBindDevResponse, err error) {
    return c.DescribeBindDevWithContext(context.Background(), request)
}

// DescribeBindDev
// 本接口（DescribeBindDev）用于查询终端用户绑定的设备列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindDevWithContext(ctx context.Context, request *DescribeBindDevRequest) (response *DescribeBindDevResponse, err error) {
    if request == nil {
        request = NewDescribeBindDevRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindDev require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBindUsr
// 本接口（DescribeBindUsr）用于查询设备被分享的所有用户列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindUsr(request *DescribeBindUsrRequest) (response *DescribeBindUsrResponse, err error) {
    return c.DescribeBindUsrWithContext(context.Background(), request)
}

// DescribeBindUsr
// 本接口（DescribeBindUsr）用于查询设备被分享的所有用户列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindUsrWithContext(ctx context.Context, request *DescribeBindUsrRequest) (response *DescribeBindUsrResponse, err error) {
    if request == nil {
        request = NewDescribeBindUsrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindUsr require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDevice
// 本接口（DescribeDevice）获取设备信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    return c.DescribeDeviceWithContext(context.Background(), request)
}

// DescribeDevice
// 本接口（DescribeDevice）获取设备信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DescribeDeviceModel
// 本接口（DescribeDeviceModel）用于获取设备物模型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceModel(request *DescribeDeviceModelRequest) (response *DescribeDeviceModelResponse, err error) {
    return c.DescribeDeviceModelWithContext(context.Background(), request)
}

// DescribeDeviceModel
// 本接口（DescribeDeviceModel）用于获取设备物模型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceModelWithContext(ctx context.Context, request *DescribeDeviceModelRequest) (response *DescribeDeviceModelResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceModel require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDevices
// 本接口（DescribeDevices）用于获取设备信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 本接口（DescribeDevices）用于获取设备信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DescribeIotDataType
// 本接口（DescribeIotDataType）用于查询自定义的物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDataType(request *DescribeIotDataTypeRequest) (response *DescribeIotDataTypeResponse, err error) {
    return c.DescribeIotDataTypeWithContext(context.Background(), request)
}

// DescribeIotDataType
// 本接口（DescribeIotDataType）用于查询自定义的物模型数据类型。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotDataTypeWithContext(ctx context.Context, request *DescribeIotDataTypeRequest) (response *DescribeIotDataTypeResponse, err error) {
    if request == nil {
        request = NewDescribeIotDataTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIotDataType require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeIotModel
// 本接口（DescribeIotModel）用于获取物模型定义详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotModel(request *DescribeIotModelRequest) (response *DescribeIotModelResponse, err error) {
    return c.DescribeIotModelWithContext(context.Background(), request)
}

// DescribeIotModel
// 本接口（DescribeIotModel）用于获取物模型定义详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotModelWithContext(ctx context.Context, request *DescribeIotModelRequest) (response *DescribeIotModelResponse, err error) {
    if request == nil {
        request = NewDescribeIotModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIotModel require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeIotModels
// 本接口（DescribeIotModels）用于列出物模型历史版本列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotModels(request *DescribeIotModelsRequest) (response *DescribeIotModelsResponse, err error) {
    return c.DescribeIotModelsWithContext(context.Background(), request)
}

// DescribeIotModels
// 本接口（DescribeIotModels）用于列出物模型历史版本列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIotModelsWithContext(ctx context.Context, request *DescribeIotModelsRequest) (response *DescribeIotModelsResponse, err error) {
    if request == nil {
        request = NewDescribeIotModelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIotModels require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeLogs
// 本接口（DescribeLogs）用于查询设备日志列表。
//
// 设备日志最长保留时长为15天,超期自动清除。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    return c.DescribeLogsWithContext(context.Background(), request)
}

// DescribeLogs
// 本接口（DescribeLogs）用于查询设备日志列表。
//
// 设备日志最长保留时长为15天,超期自动清除。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLogsWithContext(ctx context.Context, request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogs require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeMessageQueue
// 本接口（DescribeMessageQueue）用于查询物联网智能视频产品转发消息配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageQueue(request *DescribeMessageQueueRequest) (response *DescribeMessageQueueResponse, err error) {
    return c.DescribeMessageQueueWithContext(context.Background(), request)
}

// DescribeMessageQueue
// 本接口（DescribeMessageQueue）用于查询物联网智能视频产品转发消息配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageQueueWithContext(ctx context.Context, request *DescribeMessageQueueRequest) (response *DescribeMessageQueueResponse, err error) {
    if request == nil {
        request = NewDescribeMessageQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageQueue require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeModelDataRet
// 本接口（DescribeModelDataRet）用于根据TaskId获取对设备物模型操作最终响应的结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelDataRet(request *DescribeModelDataRetRequest) (response *DescribeModelDataRetResponse, err error) {
    return c.DescribeModelDataRetWithContext(context.Background(), request)
}

// DescribeModelDataRet
// 本接口（DescribeModelDataRet）用于根据TaskId获取对设备物模型操作最终响应的结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelDataRetWithContext(ctx context.Context, request *DescribeModelDataRetRequest) (response *DescribeModelDataRetResponse, err error) {
    if request == nil {
        request = NewDescribeModelDataRetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelDataRet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelDataRetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOsListRequest() (request *DescribeOsListRequest) {
    request = &DescribeOsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeOsList")
    
    
    return
}

func NewDescribeOsListResponse() (response *DescribeOsListResponse) {
    response = &DescribeOsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOsList
// 查看操作系统支持的芯片列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOsList(request *DescribeOsListRequest) (response *DescribeOsListResponse, err error) {
    return c.DescribeOsListWithContext(context.Background(), request)
}

// DescribeOsList
// 查看操作系统支持的芯片列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOsListWithContext(ctx context.Context, request *DescribeOsListRequest) (response *DescribeOsListResponse, err error) {
    if request == nil {
        request = NewDescribeOsListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOsListResponse()
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

// DescribeOtaVersions
// 本接口（DescribeOtaVersions）用于查询固件版本信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOtaVersions(request *DescribeOtaVersionsRequest) (response *DescribeOtaVersionsResponse, err error) {
    return c.DescribeOtaVersionsWithContext(context.Background(), request)
}

// DescribeOtaVersions
// 本接口（DescribeOtaVersions）用于查询固件版本信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeOtaVersionsWithContext(ctx context.Context, request *DescribeOtaVersionsRequest) (response *DescribeOtaVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeOtaVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOtaVersions require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeProduct
// 本接口（DescribeProduct）用于获取单个产品的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    return c.DescribeProductWithContext(context.Background(), request)
}

// DescribeProduct
// 本接口（DescribeProduct）用于获取单个产品的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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
// 本接口（DescribeProducts）用于列出用户账号下的物联网智能视频产品列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    return c.DescribeProductsWithContext(context.Background(), request)
}

// DescribeProducts
// 本接口（DescribeProducts）用于列出用户账号下的物联网智能视频产品列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DescribePubVersions
// 本接口（DescribePubVersions）用于获取某一产品发布过的全部固件版本。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePubVersions(request *DescribePubVersionsRequest) (response *DescribePubVersionsResponse, err error) {
    return c.DescribePubVersionsWithContext(context.Background(), request)
}

// DescribePubVersions
// 本接口（DescribePubVersions）用于获取某一产品发布过的全部固件版本。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePubVersionsWithContext(ctx context.Context, request *DescribePubVersionsRequest) (response *DescribePubVersionsResponse, err error) {
    if request == nil {
        request = NewDescribePubVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePubVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePubVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRechargeRecordsRequest() (request *DescribeRechargeRecordsRequest) {
    request = &DescribeRechargeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeRechargeRecords")
    
    
    return
}

func NewDescribeRechargeRecordsResponse() (response *DescribeRechargeRecordsResponse) {
    response = &DescribeRechargeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRechargeRecords
// 客户可通过本接口获取充值记录信息, 一次最多返回50条记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRechargeRecords(request *DescribeRechargeRecordsRequest) (response *DescribeRechargeRecordsResponse, err error) {
    return c.DescribeRechargeRecordsWithContext(context.Background(), request)
}

// DescribeRechargeRecords
// 客户可通过本接口获取充值记录信息, 一次最多返回50条记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRechargeRecordsWithContext(ctx context.Context, request *DescribeRechargeRecordsRequest) (response *DescribeRechargeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeRechargeRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRechargeRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRechargeRecordsResponse()
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

// DescribeRegistrationStatus
// 本接口（DescribeRegistrationStatus）用于查询终端用户的注册状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistrationStatus(request *DescribeRegistrationStatusRequest) (response *DescribeRegistrationStatusResponse, err error) {
    return c.DescribeRegistrationStatusWithContext(context.Background(), request)
}

// DescribeRegistrationStatus
// 本接口（DescribeRegistrationStatus）用于查询终端用户的注册状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistrationStatusWithContext(ctx context.Context, request *DescribeRegistrationStatusRequest) (response *DescribeRegistrationStatusResponse, err error) {
    if request == nil {
        request = NewDescribeRegistrationStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistrationStatus require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeRunLog
// 本接口（DescribeRunLog）用于获取设备运行日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRunLog(request *DescribeRunLogRequest) (response *DescribeRunLogResponse, err error) {
    return c.DescribeRunLogWithContext(context.Background(), request)
}

// DescribeRunLog
// 本接口（DescribeRunLog）用于获取设备运行日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRunLogWithContext(ctx context.Context, request *DescribeRunLogRequest) (response *DescribeRunLogResponse, err error) {
    if request == nil {
        request = NewDescribeRunLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRunLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageServiceRequest() (request *DescribeStorageServiceRequest) {
    request = &DescribeStorageServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeStorageService")
    
    
    return
}

func NewDescribeStorageServiceResponse() (response *DescribeStorageServiceResponse) {
    response = &DescribeStorageServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStorageService
// 查询云存服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStorageService(request *DescribeStorageServiceRequest) (response *DescribeStorageServiceResponse, err error) {
    return c.DescribeStorageServiceWithContext(context.Background(), request)
}

// DescribeStorageService
// 查询云存服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStorageServiceWithContext(ctx context.Context, request *DescribeStorageServiceRequest) (response *DescribeStorageServiceResponse, err error) {
    if request == nil {
        request = NewDescribeStorageServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamRequest() (request *DescribeStreamRequest) {
    request = &DescribeStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeStream")
    
    
    return
}

func NewDescribeStreamResponse() (response *DescribeStreamResponse) {
    response = &DescribeStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStream
// 请求设备直播流地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStream(request *DescribeStreamRequest) (response *DescribeStreamResponse, err error) {
    return c.DescribeStreamWithContext(context.Background(), request)
}

// DescribeStream
// 请求设备直播流地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"
//  FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"
//  FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStreamWithContext(ctx context.Context, request *DescribeStreamRequest) (response *DescribeStreamResponse, err error) {
    if request == nil {
        request = NewDescribeStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamResponse()
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

// DescribeTraceIds
// 本接口（DescribeTraceIds）用于查询设备日志跟踪白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTraceIds(request *DescribeTraceIdsRequest) (response *DescribeTraceIdsResponse, err error) {
    return c.DescribeTraceIdsWithContext(context.Background(), request)
}

// DescribeTraceIds
// 本接口（DescribeTraceIds）用于查询设备日志跟踪白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTraceIdsWithContext(ctx context.Context, request *DescribeTraceIdsRequest) (response *DescribeTraceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeTraceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTraceIds require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTraceStatus
// 本接口（DescribeTraceStatus）用于查询指定设备是否在白名单中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTraceStatus(request *DescribeTraceStatusRequest) (response *DescribeTraceStatusResponse, err error) {
    return c.DescribeTraceStatusWithContext(context.Background(), request)
}

// DescribeTraceStatus
// 本接口（DescribeTraceStatus）用于查询指定设备是否在白名单中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTraceStatusWithContext(ctx context.Context, request *DescribeTraceStatusRequest) (response *DescribeTraceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTraceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTraceStatus require credential")
    }

    request.SetContext(ctx)
    
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

// DisableDevice
// 本接口（DisableDevice）用于禁用设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableDevice(request *DisableDeviceRequest) (response *DisableDeviceResponse, err error) {
    return c.DisableDeviceWithContext(context.Background(), request)
}

// DisableDevice
// 本接口（DisableDevice）用于禁用设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableDeviceWithContext(ctx context.Context, request *DisableDeviceRequest) (response *DisableDeviceResponse, err error) {
    if request == nil {
        request = NewDisableDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableDevice require credential")
    }

    request.SetContext(ctx)
    
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

// DisableDeviceStream
// 本接口（DisableDeviceStream）用于停止设备推流，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableDeviceStream(request *DisableDeviceStreamRequest) (response *DisableDeviceStreamResponse, err error) {
    return c.DisableDeviceStreamWithContext(context.Background(), request)
}

// DisableDeviceStream
// 本接口（DisableDeviceStream）用于停止设备推流，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableDeviceStreamWithContext(ctx context.Context, request *DisableDeviceStreamRequest) (response *DisableDeviceStreamResponse, err error) {
    if request == nil {
        request = NewDisableDeviceStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableDeviceStream require credential")
    }

    request.SetContext(ctx)
    
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

// DisableOtaVersion
// 本接口（DisableOtaVersion）用于禁用固件版本。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableOtaVersion(request *DisableOtaVersionRequest) (response *DisableOtaVersionResponse, err error) {
    return c.DisableOtaVersionWithContext(context.Background(), request)
}

// DisableOtaVersion
// 本接口（DisableOtaVersion）用于禁用固件版本。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableOtaVersionWithContext(ctx context.Context, request *DisableOtaVersionRequest) (response *DisableOtaVersionResponse, err error) {
    if request == nil {
        request = NewDisableOtaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableOtaVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableOtaVersionResponse()
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

// ModifyDeviceAction
// 本接口（ModifyDeviceAction）用于修改设备物模型的行为（Action）。
//
// 
//
// 可对ctlVal数据属性进行写入,如:Action.takePhoto.ctlVal,设备在线且成功发送到设备才返回,物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
//
// 注意:
//
//   1.若设备当前不在线,会直接返回错误
//
//   2.若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒
//
//   3.value的内容必须与实际物模型的定义一致
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceAction(request *ModifyDeviceActionRequest) (response *ModifyDeviceActionResponse, err error) {
    return c.ModifyDeviceActionWithContext(context.Background(), request)
}

// ModifyDeviceAction
// 本接口（ModifyDeviceAction）用于修改设备物模型的行为（Action）。
//
// 
//
// 可对ctlVal数据属性进行写入,如:Action.takePhoto.ctlVal,设备在线且成功发送到设备才返回,物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
//
// 注意:
//
//   1.若设备当前不在线,会直接返回错误
//
//   2.若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒
//
//   3.value的内容必须与实际物模型的定义一致
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceActionWithContext(ctx context.Context, request *ModifyDeviceActionRequest) (response *ModifyDeviceActionResponse, err error) {
    if request == nil {
        request = NewModifyDeviceActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceAction require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyDeviceProperty
// 本接口（ModifyDeviceProperty）用于修改设备物模型的属性（ProWritable）。
//
// 可对setVal数据属性进行写入,如:
//
// ProWritable.Pos.setVal
//
// 对于嵌套类型的可写属性，可以仅对其部分数据内容进行写入，如:
//
// ProWritable.Pos.setVal.x;
//
// 可写属性云端写入成功即返回;云端向设备端发布属性变更参数;若当前设备不在线,在设备下次上线时会自动更新这些属性参数;
//
// 物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDeviceProperty(request *ModifyDevicePropertyRequest) (response *ModifyDevicePropertyResponse, err error) {
    return c.ModifyDevicePropertyWithContext(context.Background(), request)
}

// ModifyDeviceProperty
// 本接口（ModifyDeviceProperty）用于修改设备物模型的属性（ProWritable）。
//
// 可对setVal数据属性进行写入,如:
//
// ProWritable.Pos.setVal
//
// 对于嵌套类型的可写属性，可以仅对其部分数据内容进行写入，如:
//
// ProWritable.Pos.setVal.x;
//
// 可写属性云端写入成功即返回;云端向设备端发布属性变更参数;若当前设备不在线,在设备下次上线时会自动更新这些属性参数;
//
// 物模型写入数据时,不需要传入时标信息,平台以当前时标作为数据的时标更新物模型中的时标信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDevicePropertyWithContext(ctx context.Context, request *ModifyDevicePropertyRequest) (response *ModifyDevicePropertyResponse, err error) {
    if request == nil {
        request = NewModifyDevicePropertyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceProperty require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyProduct
// 本接口（ModifyProduct）用于编辑物联网智能视频产品的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyProduct(request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    return c.ModifyProductWithContext(context.Background(), request)
}

// ModifyProduct
// 本接口（ModifyProduct）用于编辑物联网智能视频产品的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

func NewModifyVerContentRequest() (request *ModifyVerContentRequest) {
    request = &ModifyVerContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "ModifyVerContent")
    
    
    return
}

func NewModifyVerContentResponse() (response *ModifyVerContentResponse) {
    response = &ModifyVerContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVerContent
// 编辑版本描述信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVerContent(request *ModifyVerContentRequest) (response *ModifyVerContentResponse, err error) {
    return c.ModifyVerContentWithContext(context.Background(), request)
}

// ModifyVerContent
// 编辑版本描述信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVerContentWithContext(ctx context.Context, request *ModifyVerContentRequest) (response *ModifyVerContentResponse, err error) {
    if request == nil {
        request = NewModifyVerContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVerContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVerContentResponse()
    err = c.Send(request, response)
    return
}

func NewRefundStorageServiceRequest() (request *RefundStorageServiceRequest) {
    request = &RefundStorageServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iotvideo", APIVersion, "RefundStorageService")
    
    
    return
}

func NewRefundStorageServiceResponse() (response *RefundStorageServiceResponse) {
    response = &RefundStorageServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefundStorageService
// 本接口（RefundStorageService）用于退订已购买的云存服务。
//
// 退订时，云存服务对应订单的处理方式 : 
//
// 1. 未开始的订单自动回到已付费订单池
//
// 2. 已开始的订单自动失效
//
// 3. 购买云存接口,优先从已付费订单池中分配订单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefundStorageService(request *RefundStorageServiceRequest) (response *RefundStorageServiceResponse, err error) {
    return c.RefundStorageServiceWithContext(context.Background(), request)
}

// RefundStorageService
// 本接口（RefundStorageService）用于退订已购买的云存服务。
//
// 退订时，云存服务对应订单的处理方式 : 
//
// 1. 未开始的订单自动回到已付费订单池
//
// 2. 已开始的订单自动失效
//
// 3. 购买云存接口,优先从已付费订单池中分配订单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RefundStorageServiceWithContext(ctx context.Context, request *RefundStorageServiceRequest) (response *RefundStorageServiceResponse, err error) {
    if request == nil {
        request = NewRefundStorageServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundStorageService require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundStorageServiceResponse()
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

// RunDevice
// 本接口（RunDevice）用于启用设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunDevice(request *RunDeviceRequest) (response *RunDeviceResponse, err error) {
    return c.RunDeviceWithContext(context.Background(), request)
}

// RunDevice
// 本接口（RunDevice）用于启用设备，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunDeviceWithContext(ctx context.Context, request *RunDeviceRequest) (response *RunDeviceResponse, err error) {
    if request == nil {
        request = NewRunDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunDevice require credential")
    }

    request.SetContext(ctx)
    
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

// RunDeviceStream
// 本接口（RunDeviceStream）用于开启设备推流，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunDeviceStream(request *RunDeviceStreamRequest) (response *RunDeviceStreamResponse, err error) {
    return c.RunDeviceStreamWithContext(context.Background(), request)
}

// RunDeviceStream
// 本接口（RunDeviceStream）用于开启设备推流，可进行批量操作，每次操作最多100台设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunDeviceStreamWithContext(ctx context.Context, request *RunDeviceStreamRequest) (response *RunDeviceStreamResponse, err error) {
    if request == nil {
        request = NewRunDeviceStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunDeviceStream require credential")
    }

    request.SetContext(ctx)
    
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

// RunIotModel
// 本接口（RunIotModel）用于对定义的物模型进行发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunIotModel(request *RunIotModelRequest) (response *RunIotModelResponse, err error) {
    return c.RunIotModelWithContext(context.Background(), request)
}

// RunIotModel
// 本接口（RunIotModel）用于对定义的物模型进行发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunIotModelWithContext(ctx context.Context, request *RunIotModelRequest) (response *RunIotModelResponse, err error) {
    if request == nil {
        request = NewRunIotModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunIotModel require credential")
    }

    request.SetContext(ctx)
    
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

// RunOtaVersion
// 本接口（RunOtaVersion）用于固件版本正式发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunOtaVersion(request *RunOtaVersionRequest) (response *RunOtaVersionResponse, err error) {
    return c.RunOtaVersionWithContext(context.Background(), request)
}

// RunOtaVersion
// 本接口（RunOtaVersion）用于固件版本正式发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunOtaVersionWithContext(ctx context.Context, request *RunOtaVersionRequest) (response *RunOtaVersionResponse, err error) {
    if request == nil {
        request = NewRunOtaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunOtaVersion require credential")
    }

    request.SetContext(ctx)
    
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

// RunTestOtaVersion
// 本接口（RunTestOtaVersion）用于固件版本测试发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunTestOtaVersion(request *RunTestOtaVersionRequest) (response *RunTestOtaVersionResponse, err error) {
    return c.RunTestOtaVersionWithContext(context.Background(), request)
}

// RunTestOtaVersion
// 本接口（RunTestOtaVersion）用于固件版本测试发布。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunTestOtaVersionWithContext(ctx context.Context, request *RunTestOtaVersionRequest) (response *RunTestOtaVersionResponse, err error) {
    if request == nil {
        request = NewRunTestOtaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunTestOtaVersion require credential")
    }

    request.SetContext(ctx)
    
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

// SendOnlineMsg
// 本接口（SendOnlineMsg）用于向设备发送在线消息。
//
// 注意：
//
// 若设备当前不在线,会直接返回错误;
//
// 若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒.waitresp非0情况下,会导致本接口阻塞3秒。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendOnlineMsg(request *SendOnlineMsgRequest) (response *SendOnlineMsgResponse, err error) {
    return c.SendOnlineMsgWithContext(context.Background(), request)
}

// SendOnlineMsg
// 本接口（SendOnlineMsg）用于向设备发送在线消息。
//
// 注意：
//
// 若设备当前不在线,会直接返回错误;
//
// 若设备网络出现异常时,消息发送可能超时,超时等待最长时间为3秒.waitresp非0情况下,会导致本接口阻塞3秒。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendOnlineMsgWithContext(ctx context.Context, request *SendOnlineMsgRequest) (response *SendOnlineMsgResponse, err error) {
    if request == nil {
        request = NewSendOnlineMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendOnlineMsg require credential")
    }

    request.SetContext(ctx)
    
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

// SetMessageQueue
// 本接口（SetMessageQueue）用于配置物联网智能视频产品的转发消息队列。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetMessageQueue(request *SetMessageQueueRequest) (response *SetMessageQueueResponse, err error) {
    return c.SetMessageQueueWithContext(context.Background(), request)
}

// SetMessageQueue
// 本接口（SetMessageQueue）用于配置物联网智能视频产品的转发消息队列。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetMessageQueueWithContext(ctx context.Context, request *SetMessageQueueRequest) (response *SetMessageQueueResponse, err error) {
    if request == nil {
        request = NewSetMessageQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetMessageQueue require credential")
    }

    request.SetContext(ctx)
    
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

// UpgradeDevice
// 本接口（UpgradeDevice）用于对设备进行固件升级。
//
// 该接口向指定的设备下发固件更新指令,可将固件升级到任意版本(可实现固件降级)。
//
// 警告:使能UpgradeNow参数存在一定的风险性！建议仅在debug场景下使用!
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeDevice(request *UpgradeDeviceRequest) (response *UpgradeDeviceResponse, err error) {
    return c.UpgradeDeviceWithContext(context.Background(), request)
}

// UpgradeDevice
// 本接口（UpgradeDevice）用于对设备进行固件升级。
//
// 该接口向指定的设备下发固件更新指令,可将固件升级到任意版本(可实现固件降级)。
//
// 警告:使能UpgradeNow参数存在一定的风险性！建议仅在debug场景下使用!
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeDeviceWithContext(ctx context.Context, request *UpgradeDeviceRequest) (response *UpgradeDeviceResponse, err error) {
    if request == nil {
        request = NewUpgradeDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDevice require credential")
    }

    request.SetContext(ctx)
    
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

// UploadOtaVersion
// 本接口（UploadOtaVersion）接收上传到控制台的固件版本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadOtaVersion(request *UploadOtaVersionRequest) (response *UploadOtaVersionResponse, err error) {
    return c.UploadOtaVersionWithContext(context.Background(), request)
}

// UploadOtaVersion
// 本接口（UploadOtaVersion）接收上传到控制台的固件版本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREFAILURE = "AuthFailure.SignatureFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadOtaVersionWithContext(ctx context.Context, request *UploadOtaVersionRequest) (response *UploadOtaVersionResponse, err error) {
    if request == nil {
        request = NewUploadOtaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadOtaVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadOtaVersionResponse()
    err = c.Send(request, response)
    return
}
