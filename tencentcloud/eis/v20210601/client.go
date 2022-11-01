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

package v20210601

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-06-01"

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


func NewGetRuntimeMCRequest() (request *GetRuntimeMCRequest) {
    request = &GetRuntimeMCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "GetRuntimeMC")
    
    
    return
}

func NewGetRuntimeMCResponse() (response *GetRuntimeMCResponse) {
    response = &GetRuntimeMCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRuntimeMC
// 获取运行时详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INTERNALERROR_COUNTRUNTIMEINSTANCESFAILED = "InternalError.CountRuntimeInstancesFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMENAMESPACEINVALID = "InvalidParameterValue.RuntimeNamespaceInvalid"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) GetRuntimeMC(request *GetRuntimeMCRequest) (response *GetRuntimeMCResponse, err error) {
    return c.GetRuntimeMCWithContext(context.Background(), request)
}

// GetRuntimeMC
// 获取运行时详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INTERNALERROR_COUNTRUNTIMEINSTANCESFAILED = "InternalError.CountRuntimeInstancesFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMENAMESPACEINVALID = "InvalidParameterValue.RuntimeNamespaceInvalid"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) GetRuntimeMCWithContext(ctx context.Context, request *GetRuntimeMCRequest) (response *GetRuntimeMCResponse, err error) {
    if request == nil {
        request = NewGetRuntimeMCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRuntimeMC require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRuntimeMCResponse()
    err = c.Send(request, response)
    return
}

func NewGetRuntimeResourceMonitorMetricMCRequest() (request *GetRuntimeResourceMonitorMetricMCRequest) {
    request = &GetRuntimeResourceMonitorMetricMCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "GetRuntimeResourceMonitorMetricMC")
    
    
    return
}

func NewGetRuntimeResourceMonitorMetricMCResponse() (response *GetRuntimeResourceMonitorMetricMCResponse) {
    response = &GetRuntimeResourceMonitorMetricMCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRuntimeResourceMonitorMetricMC
// 获取运行时资源监控详情，cpu，memory，bandwidth
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INNERLOGICTIMEOUT = "FailedOperation.InnerLogicTimeOut"
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INVALIDPARAMETERVALUE_INVALIDRUNTIMEMETRICSEARCHCONDITION = "InvalidParameterValue.InvalidRuntimeMetricSearchCondition"
//  INVALIDPARAMETERVALUE_NOTSUPPORTEDACTIONFORPUBLICRUNTIME = "InvalidParameterValue.NotSupportedActionForPublicRuntime"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMEMETRICRATENOTSUPPORT = "InvalidParameterValue.RuntimeMetricRateNotSupport"
//  INVALIDPARAMETERVALUE_RUNTIMENAMESPACEINVALID = "InvalidParameterValue.RuntimeNamespaceInvalid"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) GetRuntimeResourceMonitorMetricMC(request *GetRuntimeResourceMonitorMetricMCRequest) (response *GetRuntimeResourceMonitorMetricMCResponse, err error) {
    return c.GetRuntimeResourceMonitorMetricMCWithContext(context.Background(), request)
}

// GetRuntimeResourceMonitorMetricMC
// 获取运行时资源监控详情，cpu，memory，bandwidth
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INNERLOGICTIMEOUT = "FailedOperation.InnerLogicTimeOut"
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INVALIDPARAMETERVALUE_INVALIDRUNTIMEMETRICSEARCHCONDITION = "InvalidParameterValue.InvalidRuntimeMetricSearchCondition"
//  INVALIDPARAMETERVALUE_NOTSUPPORTEDACTIONFORPUBLICRUNTIME = "InvalidParameterValue.NotSupportedActionForPublicRuntime"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMEMETRICRATENOTSUPPORT = "InvalidParameterValue.RuntimeMetricRateNotSupport"
//  INVALIDPARAMETERVALUE_RUNTIMENAMESPACEINVALID = "InvalidParameterValue.RuntimeNamespaceInvalid"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) GetRuntimeResourceMonitorMetricMCWithContext(ctx context.Context, request *GetRuntimeResourceMonitorMetricMCRequest) (response *GetRuntimeResourceMonitorMetricMCResponse, err error) {
    if request == nil {
        request = NewGetRuntimeResourceMonitorMetricMCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRuntimeResourceMonitorMetricMC require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRuntimeResourceMonitorMetricMCResponse()
    err = c.Send(request, response)
    return
}

func NewListDeployableRuntimesMCRequest() (request *ListDeployableRuntimesMCRequest) {
    request = &ListDeployableRuntimesMCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "ListDeployableRuntimesMC")
    
    
    return
}

func NewListDeployableRuntimesMCResponse() (response *ListDeployableRuntimesMCResponse) {
    response = &ListDeployableRuntimesMCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListDeployableRuntimesMC
// 返回用户可用的运行时列表，发布应用时返回的运行时环境，仅shared和private运行时，无sandbox运行时，并且只有running/scaling状态的
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_UNSUPPORTEDOPERATIONTYPE = "FailedOperation.UnSupportedOperationType"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
func (c *Client) ListDeployableRuntimesMC(request *ListDeployableRuntimesMCRequest) (response *ListDeployableRuntimesMCResponse, err error) {
    return c.ListDeployableRuntimesMCWithContext(context.Background(), request)
}

// ListDeployableRuntimesMC
// 返回用户可用的运行时列表，发布应用时返回的运行时环境，仅shared和private运行时，无sandbox运行时，并且只有running/scaling状态的
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_UNSUPPORTEDOPERATIONTYPE = "FailedOperation.UnSupportedOperationType"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
func (c *Client) ListDeployableRuntimesMCWithContext(ctx context.Context, request *ListDeployableRuntimesMCRequest) (response *ListDeployableRuntimesMCResponse, err error) {
    if request == nil {
        request = NewListDeployableRuntimesMCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDeployableRuntimesMC require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDeployableRuntimesMCResponse()
    err = c.Send(request, response)
    return
}

func NewListRuntimeDeployedInstancesMCRequest() (request *ListRuntimeDeployedInstancesMCRequest) {
    request = &ListRuntimeDeployedInstancesMCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "ListRuntimeDeployedInstancesMC")
    
    
    return
}

func NewListRuntimeDeployedInstancesMCResponse() (response *ListRuntimeDeployedInstancesMCResponse) {
    response = &ListRuntimeDeployedInstancesMCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRuntimeDeployedInstancesMC
// 获取运行时部署的应用实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_LISTRUNTIMEINSTANCESFAILED = "InternalError.ListRuntimeInstancesFailed"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) ListRuntimeDeployedInstancesMC(request *ListRuntimeDeployedInstancesMCRequest) (response *ListRuntimeDeployedInstancesMCResponse, err error) {
    return c.ListRuntimeDeployedInstancesMCWithContext(context.Background(), request)
}

// ListRuntimeDeployedInstancesMC
// 获取运行时部署的应用实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_CHECKRUNTIMEFAILED = "InternalError.CheckRuntimeFailed"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_LISTRUNTIMEINSTANCESFAILED = "InternalError.ListRuntimeInstancesFailed"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  INVALIDPARAMETERVALUE_RUNTIMEALREADYDELETED = "InvalidParameterValue.RuntimeAlreadyDeleted"
//  INVALIDPARAMETERVALUE_RUNTIMEIDNOTEXIST = "InvalidParameterValue.RuntimeIdNotExist"
//  INVALIDPARAMETERVALUE_RUNTIMEZONENOTEXISTED = "InvalidParameterValue.RuntimeZoneNotExisted"
func (c *Client) ListRuntimeDeployedInstancesMCWithContext(ctx context.Context, request *ListRuntimeDeployedInstancesMCRequest) (response *ListRuntimeDeployedInstancesMCResponse, err error) {
    if request == nil {
        request = NewListRuntimeDeployedInstancesMCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRuntimeDeployedInstancesMC require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRuntimeDeployedInstancesMCResponse()
    err = c.Send(request, response)
    return
}

func NewListRuntimesMCRequest() (request *ListRuntimesMCRequest) {
    request = &ListRuntimesMCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eis", APIVersion, "ListRuntimesMC")
    
    
    return
}

func NewListRuntimesMCResponse() (response *ListRuntimesMCResponse) {
    response = &ListRuntimesMCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRuntimesMC
// 返回用户的运行时列表，运行时管理主页使用，包含沙箱、共享运行时及独立运行时环境，不包含已经删除的运行时
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  FAILEDOPERATION_UNSUPPORTEDOPERATIONTYPE = "FailedOperation.UnSupportedOperationType"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_LISTRUNTIMESFAILED = "InternalError.ListRuntimesFailed"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListRuntimesMC(request *ListRuntimesMCRequest) (response *ListRuntimesMCResponse, err error) {
    return c.ListRuntimesMCWithContext(context.Background(), request)
}

// ListRuntimesMC
// 返回用户的运行时列表，运行时管理主页使用，包含沙箱、共享运行时及独立运行时环境，不包含已经删除的运行时
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHENTICATIONFAILED = "FailedOperation.AuthenticationFailed"
//  FAILEDOPERATION_METACOMPILERERROR = "FailedOperation.MetaCompilerError"
//  FAILEDOPERATION_UNSUPPORTEDOPERATIONTYPE = "FailedOperation.UnSupportedOperationType"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_LISTRUNTIMESFAILED = "InternalError.ListRuntimesFailed"
//  INTERNALERROR_METACOMPILERERROR = "InternalError.MetaCompilerError"
//  INTERNALERROR_RPCPILOTSERVERERROR = "InternalError.RpcPilotServerError"
//  INVALIDPARAMETERVALUE_PILOTZONENOTSUPPORTED = "InvalidParameterValue.PilotZoneNotSupported"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListRuntimesMCWithContext(ctx context.Context, request *ListRuntimesMCRequest) (response *ListRuntimesMCResponse, err error) {
    if request == nil {
        request = NewListRuntimesMCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRuntimesMC require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRuntimesMCResponse()
    err = c.Send(request, response)
    return
}
