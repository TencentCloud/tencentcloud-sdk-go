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
func (c *Client) CallDeviceActionAsync(request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionAsyncRequest()
    }
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
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) CallDeviceActionSync(request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionSyncRequest()
    }
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERTAIONERROR = "InternalError.DBOpertaionError"
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELDEFINEEVENTPROPNAMEERROR = "InvalidParameterValue.ModelDefineEventPropNameError"
//  INVALIDPARAMETERVALUE_MODELDEFINEINVALID = "InvalidParameterValue.ModelDefineInvalid"
//  INVALIDPARAMETERVALUE_MODELDEFINENIL = "InvalidParameterValue.ModelDefineNil"
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
    if request == nil {
        request = NewControlDeviceDataRequest()
    }
    response = NewControlDeviceDataResponse()
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
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    response = NewCreateDeviceResponse()
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
    if request == nil {
        request = NewCreateLoRaFrequencyRequest()
    }
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  RESOURCENOTFOUND_GATEWAYDUPKEYEXIST = "ResourceNotFound.GatewayDupKeyExist"
func (c *Client) CreateLoRaGateway(request *CreateLoRaGatewayRequest) (response *CreateLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewCreateLoRaGatewayRequest()
    }
    response = NewCreateLoRaGatewayResponse()
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
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
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
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnsupportedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_PRODUCTDUPKEYEXIST = "UnsupportedOperation.ProductDupKeyExist"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) CreateStudioProduct(request *CreateStudioProductRequest) (response *CreateStudioProductResponse, err error) {
    if request == nil {
        request = NewCreateStudioProductRequest()
    }
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
func (c *Client) CreateTopicPolicy(request *CreateTopicPolicyRequest) (response *CreateTopicPolicyResponse, err error) {
    if request == nil {
        request = NewCreateTopicPolicyRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTopicRule(request *CreateTopicRuleRequest) (response *CreateTopicRuleResponse, err error) {
    if request == nil {
        request = NewCreateTopicRuleRequest()
    }
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
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
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
    if request == nil {
        request = NewDeleteDevicesRequest()
    }
    response = NewDeleteDevicesResponse()
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
    if request == nil {
        request = NewDeleteLoRaFrequencyRequest()
    }
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
    if request == nil {
        request = NewDeleteLoRaGatewayRequest()
    }
    response = NewDeleteLoRaGatewayResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_POOLEXISTUNDERPROJECT = "UnsupportedOperation.PoolExistUnderProject"
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
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
    if request == nil {
        request = NewDeleteStudioProductRequest()
    }
    response = NewDeleteStudioProductResponse()
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
    if request == nil {
        request = NewDeleteTopicRuleRequest()
    }
    response = NewDeleteTopicRuleResponse()
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    response = NewDescribeDeviceResponse()
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
func (c *Client) DescribeDeviceDataHistory(request *DescribeDeviceDataHistoryRequest) (response *DescribeDeviceDataHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataHistoryRequest()
    }
    response = NewDescribeDeviceDataHistoryResponse()
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
    if request == nil {
        request = NewDescribeFirmwareTaskRequest()
    }
    response = NewDescribeFirmwareTaskResponse()
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
    if request == nil {
        request = NewDescribeLoRaFrequencyRequest()
    }
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
    if request == nil {
        request = NewDescribeModelDefinitionRequest()
    }
    response = NewDescribeModelDefinitionResponse()
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
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    response = NewDescribeProjectResponse()
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
func (c *Client) DescribeStudioProduct(request *DescribeStudioProductRequest) (response *DescribeStudioProductResponse, err error) {
    if request == nil {
        request = NewDescribeStudioProductRequest()
    }
    response = NewDescribeStudioProductResponse()
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
    if request == nil {
        request = NewDescribeTopicRuleRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_APPNOPERMISSION = "UnauthorizedOperation.AppNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) DirectBindDeviceInFamily(request *DirectBindDeviceInFamilyRequest) (response *DirectBindDeviceInFamilyResponse, err error) {
    if request == nil {
        request = NewDirectBindDeviceInFamilyRequest()
    }
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
    if request == nil {
        request = NewDisableTopicRuleRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"
//  INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableTopicRule(request *EnableTopicRuleRequest) (response *EnableTopicRuleResponse, err error) {
    if request == nil {
        request = NewEnableTopicRuleRequest()
    }
    response = NewEnableTopicRuleResponse()
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
// 本接口（GetCOSURL）用于获取固件存储在COS的URL 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetCOSURL(request *GetCOSURLRequest) (response *GetCOSURLResponse, err error) {
    if request == nil {
        request = NewGetCOSURLRequest()
    }
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
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetDeviceList(request *GetDeviceListRequest) (response *GetDeviceListResponse, err error) {
    if request == nil {
        request = NewGetDeviceListRequest()
    }
    response = NewGetDeviceListResponse()
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
    if request == nil {
        request = NewGetLoRaGatewayListRequest()
    }
    response = NewGetLoRaGatewayListResponse()
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
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOINSTANCE = "UnauthorizedOperation.NoPermissionToInstance"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) GetProjectList(request *GetProjectListRequest) (response *GetProjectListResponse, err error) {
    if request == nil {
        request = NewGetProjectListRequest()
    }
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
    if request == nil {
        request = NewGetStudioProductListRequest()
    }
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
    if request == nil {
        request = NewGetTopicRuleListRequest()
    }
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
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOPROJECT = "UnauthorizedOperation.NoPermissionToProject"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ListEventHistory(request *ListEventHistoryRequest) (response *ListEventHistoryResponse, err error) {
    if request == nil {
        request = NewListEventHistoryRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) ListFirmwares(request *ListFirmwaresRequest) (response *ListFirmwaresResponse, err error) {
    if request == nil {
        request = NewListFirmwaresRequest()
    }
    response = NewListFirmwaresResponse()
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
    if request == nil {
        request = NewModifyLoRaFrequencyRequest()
    }
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
//  INTERNALERROR_INTERNALRPCERROR = "InternalError.InternalRPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_GATEWAYNOTEXIST = "ResourceNotFound.GatewayNotExist"
func (c *Client) ModifyLoRaGateway(request *ModifyLoRaGatewayRequest) (response *ModifyLoRaGatewayResponse, err error) {
    if request == nil {
        request = NewModifyLoRaGatewayRequest()
    }
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
//  UNSUPPORTEDOPERATION_PRODUCTEXISTUNDERPROJECT = "UnsupportedOperation.ProductExistUnderProject"
func (c *Client) ModifyModelDefinition(request *ModifyModelDefinitionRequest) (response *ModifyModelDefinitionResponse, err error) {
    if request == nil {
        request = NewModifyModelDefinitionRequest()
    }
    response = NewModifyModelDefinitionResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
//  UNSUPPORTEDOPERATION_PROJECTDUPKEYEXIST = "UnsupportedOperation.ProjectDupKeyExist"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    response = NewModifyProjectResponse()
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
func (c *Client) ModifyStudioProduct(request *ModifyStudioProductRequest) (response *ModifyStudioProductResponse, err error) {
    if request == nil {
        request = NewModifyStudioProductRequest()
    }
    response = NewModifyStudioProductResponse()
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
//  INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"
//  INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"
//  INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"
//  INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"
//  RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyTopicRule(request *ModifyTopicRuleRequest) (response *ModifyTopicRuleResponse, err error) {
    if request == nil {
        request = NewModifyTopicRuleRequest()
    }
    response = NewModifyTopicRuleResponse()
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
//  FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"
//  FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
    if request == nil {
        request = NewPublishMessageRequest()
    }
    response = NewPublishMessageResponse()
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
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
func (c *Client) ReleaseStudioProduct(request *ReleaseStudioProductRequest) (response *ReleaseStudioProductResponse, err error) {
    if request == nil {
        request = NewReleaseStudioProductRequest()
    }
    response = NewReleaseStudioProductResponse()
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
//  INVALIDPARAMETERVALUE_PRODUCTPARMSERROR = "InvalidParameterValue.ProductParmsError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchStudioProduct(request *SearchStudioProductRequest) (response *SearchStudioProductResponse, err error) {
    if request == nil {
        request = NewSearchStudioProductRequest()
    }
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
    if request == nil {
        request = NewSearchTopicRuleRequest()
    }
    response = NewSearchTopicRuleResponse()
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
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONTOSTUDIOPRODUCT = "UnauthorizedOperation.NoPermissionToStudioProduct"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UpdateFirmware(request *UpdateFirmwareRequest) (response *UpdateFirmwareResponse, err error) {
    if request == nil {
        request = NewUpdateFirmwareRequest()
    }
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
// 本接口（UploadFirmware）用于上传设备固件至平台
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FIRMWAREALREADYEXIST = "InvalidParameter.FirmwareAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"
//  LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"
//  RESOURCENOTFOUND_STUDIOPRODUCTNOTEXIST = "ResourceNotFound.StudioProductNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEISOLATED = "UnsupportedOperation.InstanceIsolated"
func (c *Client) UploadFirmware(request *UploadFirmwareRequest) (response *UploadFirmwareResponse, err error) {
    if request == nil {
        request = NewUploadFirmwareRequest()
    }
    response = NewUploadFirmwareResponse()
    err = c.Send(request, response)
    return
}
