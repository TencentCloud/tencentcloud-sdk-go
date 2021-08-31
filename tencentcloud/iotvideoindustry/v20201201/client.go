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

package v20201201

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-01"

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


func NewBindGroupDevicesRequest() (request *BindGroupDevicesRequest) {
    request = &BindGroupDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "BindGroupDevices")
    return
}

func NewBindGroupDevicesResponse() (response *BindGroupDevicesResponse) {
    response = &BindGroupDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindGroupDevices
// 本接口(BindGroupDevices) 用于绑定设备到分组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) BindGroupDevices(request *BindGroupDevicesRequest) (response *BindGroupDevicesResponse, err error) {
    if request == nil {
        request = NewBindGroupDevicesRequest()
    }
    response = NewBindGroupDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewControlDevicePTZRequest() (request *ControlDevicePTZRequest) {
    request = &ControlDevicePTZRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlDevicePTZ")
    return
}

func NewControlDevicePTZResponse() (response *ControlDevicePTZResponse) {
    response = &ControlDevicePTZResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlDevicePTZ
// 本接口(ControlDevicePTZ) 用于对支持GB28181 PTZ信令的设备进行远程控制。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) ControlDevicePTZ(request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    if request == nil {
        request = NewControlDevicePTZRequest()
    }
    response = NewControlDevicePTZResponse()
    err = c.Send(request, response)
    return
}

func NewControlRecordStreamRequest() (request *ControlRecordStreamRequest) {
    request = &ControlRecordStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlRecordStream")
    return
}

func NewControlRecordStreamResponse() (response *ControlRecordStreamResponse) {
    response = &ControlRecordStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlRecordStream
// 对回放流进行控制，包括暂停、播放、拉动、结束等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) ControlRecordStream(request *ControlRecordStreamRequest) (response *ControlRecordStreamResponse, err error) {
    if request == nil {
        request = NewControlRecordStreamRequest()
    }
    response = NewControlRecordStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateDevice")
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDevice
// 本接口(CreateDevice) 用于创建设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceGroupRequest() (request *CreateDeviceGroupRequest) {
    request = &CreateDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateDeviceGroup")
    return
}

func NewCreateDeviceGroupResponse() (response *CreateDeviceGroupResponse) {
    response = &CreateDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDeviceGroup
// 本接口(CreateDeviceGroup) 用于创建设备管理分组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeviceGroupRequest()
    }
    response = NewCreateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordPlanRequest() (request *CreateRecordPlanRequest) {
    request = &CreateRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateRecordPlan")
    return
}

func NewCreateRecordPlanResponse() (response *CreateRecordPlanResponse) {
    response = &CreateRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecordPlan
// 本接口(CreateRecordPlan) 用于创建录制计划，使设备与时间模板绑定，以便及时启动录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) CreateRecordPlan(request *CreateRecordPlanRequest) (response *CreateRecordPlanResponse, err error) {
    if request == nil {
        request = NewCreateRecordPlanRequest()
    }
    response = NewCreateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTimeTemplateRequest() (request *CreateTimeTemplateRequest) {
    request = &CreateTimeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateTimeTemplate")
    return
}

func NewCreateTimeTemplateResponse() (response *CreateTimeTemplateResponse) {
    response = &CreateTimeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTimeTemplate
// 本接口(CreateTimeTemplate) 用于根据模板描述的具体录制时间片段，创建定制化的时间模板。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateTimeTemplate(request *CreateTimeTemplateRequest) (response *CreateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTimeTemplateRequest()
    }
    response = NewCreateTimeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteDevice")
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// 本接口(DeleteDevice)用于删除设备。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceGroupRequest() (request *DeleteDeviceGroupRequest) {
    request = &DeleteDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteDeviceGroup")
    return
}

func NewDeleteDeviceGroupResponse() (response *DeleteDeviceGroupResponse) {
    response = &DeleteDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDeviceGroup
// 本接口(DeleteDeviceGroup)用于删除分组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) DeleteDeviceGroup(request *DeleteDeviceGroupRequest) (response *DeleteDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupRequest()
    }
    response = NewDeleteDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordPlanRequest() (request *DeleteRecordPlanRequest) {
    request = &DeleteRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteRecordPlan")
    return
}

func NewDeleteRecordPlanResponse() (response *DeleteRecordPlanResponse) {
    response = &DeleteRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordPlan
// 本接口(DeleteRecordPlan)用于删除录制计划
//
// 录制计划删除的同时，会停止该录制计划下的全部录制任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordPlanRequest()
    }
    response = NewDeleteRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTimeTemplateRequest() (request *DeleteTimeTemplateRequest) {
    request = &DeleteTimeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteTimeTemplate")
    return
}

func NewDeleteTimeTemplateResponse() (response *DeleteTimeTemplateResponse) {
    response = &DeleteTimeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTimeTemplate
// 本接口(DeleteTimeTemplate) 用于删除时间模板。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) DeleteTimeTemplate(request *DeleteTimeTemplateRequest) (response *DeleteTimeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTimeTemplateRequest()
    }
    response = NewDeleteTimeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllDeviceListRequest() (request *DescribeAllDeviceListRequest) {
    request = &DescribeAllDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeAllDeviceList")
    return
}

func NewDescribeAllDeviceListResponse() (response *DescribeAllDeviceListResponse) {
    response = &DescribeAllDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllDeviceList
// 本接口(DescribeAllDeviceList) 用于获取设备列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeAllDeviceList(request *DescribeAllDeviceListRequest) (response *DescribeAllDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeAllDeviceListRequest()
    }
    response = NewDescribeAllDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceGroupRequest() (request *DescribeDeviceGroupRequest) {
    request = &DescribeDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDeviceGroup")
    return
}

func NewDescribeDeviceGroupResponse() (response *DescribeDeviceGroupResponse) {
    response = &DescribeDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceGroup
// 本接口(DescribeDeviceGroup)用于根据设备ID查询设备所在分组信息，可批量查询。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeDeviceGroup(request *DescribeDeviceGroupRequest) (response *DescribeDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupRequest()
    }
    response = NewDescribeDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePassWordRequest() (request *DescribeDevicePassWordRequest) {
    request = &DescribeDevicePassWordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDevicePassWord")
    return
}

func NewDescribeDevicePassWordResponse() (response *DescribeDevicePassWordResponse) {
    response = &DescribeDevicePassWordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePassWord
// 本接口(DescribeDevicePassWord)用于查询设备密码。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeDevicePassWord(request *DescribeDevicePassWordRequest) (response *DescribeDevicePassWordResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePassWordRequest()
    }
    response = NewDescribeDevicePassWordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceStreamsRequest() (request *DescribeDeviceStreamsRequest) {
    request = &DescribeDeviceStreamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDeviceStreams")
    return
}

func NewDescribeDeviceStreamsResponse() (response *DescribeDeviceStreamsResponse) {
    response = &DescribeDeviceStreamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceStreams
// 本接口(DescribeDeviceStreams)用于获取设备实时流地址。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeDeviceStreams(request *DescribeDeviceStreamsRequest) (response *DescribeDeviceStreamsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStreamsRequest()
    }
    response = NewDescribeDeviceStreamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupByIdRequest() (request *DescribeGroupByIdRequest) {
    request = &DescribeGroupByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeGroupById")
    return
}

func NewDescribeGroupByIdResponse() (response *DescribeGroupByIdResponse) {
    response = &DescribeGroupByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupById
// 本接口(DescribeGroupById)用于根据分组ID查询分组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) DescribeGroupById(request *DescribeGroupByIdRequest) (response *DescribeGroupByIdResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByIdRequest()
    }
    response = NewDescribeGroupByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupByPathRequest() (request *DescribeGroupByPathRequest) {
    request = &DescribeGroupByPathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeGroupByPath")
    return
}

func NewDescribeGroupByPathResponse() (response *DescribeGroupByPathResponse) {
    response = &DescribeGroupByPathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupByPath
// 根据分组路径查询分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) DescribeGroupByPath(request *DescribeGroupByPathRequest) (response *DescribeGroupByPathResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByPathRequest()
    }
    response = NewDescribeGroupByPathResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupDevicesRequest() (request *DescribeGroupDevicesRequest) {
    request = &DescribeGroupDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeGroupDevices")
    return
}

func NewDescribeGroupDevicesResponse() (response *DescribeGroupDevicesResponse) {
    response = &DescribeGroupDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupDevices
// 本接口(DescribeGroupDevices)用于查询分组下的设备列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeGroupDevices(request *DescribeGroupDevicesRequest) (response *DescribeGroupDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupDevicesRequest()
    }
    response = NewDescribeGroupDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeGroups")
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroups
// 本接口(DescribeGroups)用于批量查询分组信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordStreamRequest() (request *DescribeRecordStreamRequest) {
    request = &DescribeRecordStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeRecordStream")
    return
}

func NewDescribeRecordStreamResponse() (response *DescribeRecordStreamResponse) {
    response = &DescribeRecordStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordStream
// 获取回放视频流(NVR录制用)
//
// RecordId和StartTime/EndTime互斥
//
// 当存在RecordId时，StartTime和EndTime无效
//
// 当RecordId为空，StartTime和EndTime生效
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeRecordStream(request *DescribeRecordStreamRequest) (response *DescribeRecordStreamResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStreamRequest()
    }
    response = NewDescribeRecordStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSIPServerRequest() (request *DescribeSIPServerRequest) {
    request = &DescribeSIPServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeSIPServer")
    return
}

func NewDescribeSIPServerResponse() (response *DescribeSIPServerResponse) {
    response = &DescribeSIPServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSIPServer
// 本接口用于获取SIP服务器相关配置，用户可以通过这些配置项，将设备通过GB28181协议注册到本服务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeSIPServer(request *DescribeSIPServerRequest) (response *DescribeSIPServerResponse, err error) {
    if request == nil {
        request = NewDescribeSIPServerRequest()
    }
    response = NewDescribeSIPServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticDetailsRequest() (request *DescribeStatisticDetailsRequest) {
    request = &DescribeStatisticDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeStatisticDetails")
    return
}

func NewDescribeStatisticDetailsResponse() (response *DescribeStatisticDetailsResponse) {
    response = &DescribeStatisticDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStatisticDetails
// 本接口(DescribeStatisticDetails)用于查询指定统计项详情，返回结果按天为单位聚合，支持的最大时间查询范围为31天。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeStatisticDetails(request *DescribeStatisticDetailsRequest) (response *DescribeStatisticDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDetailsRequest()
    }
    response = NewDescribeStatisticDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticSummaryRequest() (request *DescribeStatisticSummaryRequest) {
    request = &DescribeStatisticSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeStatisticSummary")
    return
}

func NewDescribeStatisticSummaryResponse() (response *DescribeStatisticSummaryResponse) {
    response = &DescribeStatisticSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStatisticSummary
// 本接口(DescribeStatisticSummary)用于查询用户昨日的概览数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeStatisticSummary(request *DescribeStatisticSummaryRequest) (response *DescribeStatisticSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticSummaryRequest()
    }
    response = NewDescribeStatisticSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubGroupsRequest() (request *DescribeSubGroupsRequest) {
    request = &DescribeSubGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeSubGroups")
    return
}

func NewDescribeSubGroupsResponse() (response *DescribeSubGroupsResponse) {
    response = &DescribeSubGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubGroups
// 本接口(DescribeSubGroups)用于查询分组下的子分组列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) DescribeSubGroups(request *DescribeSubGroupsRequest) (response *DescribeSubGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSubGroupsRequest()
    }
    response = NewDescribeSubGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoListRequest() (request *DescribeVideoListRequest) {
    request = &DescribeVideoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeVideoList")
    return
}

func NewDescribeVideoListResponse() (response *DescribeVideoListResponse) {
    response = &DescribeVideoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoList
// 根据时间获取回放文件列表(云端录制用)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeVideoList(request *DescribeVideoListRequest) (response *DescribeVideoListResponse, err error) {
    if request == nil {
        request = NewDescribeVideoListRequest()
    }
    response = NewDescribeVideoListResponse()
    err = c.Send(request, response)
    return
}

func NewGetRecordDatesByDevRequest() (request *GetRecordDatesByDevRequest) {
    request = &GetRecordDatesByDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetRecordDatesByDev")
    return
}

func NewGetRecordDatesByDevResponse() (response *GetRecordDatesByDevResponse) {
    response = &GetRecordDatesByDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRecordDatesByDev
// 本接口(GetRecordDatesByDev)用于查询设备含有录像文件的日期列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) GetRecordDatesByDev(request *GetRecordDatesByDevRequest) (response *GetRecordDatesByDevResponse, err error) {
    if request == nil {
        request = NewGetRecordDatesByDevRequest()
    }
    response = NewGetRecordDatesByDevResponse()
    err = c.Send(request, response)
    return
}

func NewGetRecordPlanByDevRequest() (request *GetRecordPlanByDevRequest) {
    request = &GetRecordPlanByDevRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetRecordPlanByDev")
    return
}

func NewGetRecordPlanByDevResponse() (response *GetRecordPlanByDevResponse) {
    response = &GetRecordPlanByDevResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRecordPlanByDev
// 本接口(GetRecordPlanByDev)用于根据设备ID查询其绑定的录制计划.
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) GetRecordPlanByDev(request *GetRecordPlanByDevRequest) (response *GetRecordPlanByDevResponse, err error) {
    if request == nil {
        request = NewGetRecordPlanByDevRequest()
    }
    response = NewGetRecordPlanByDevResponse()
    err = c.Send(request, response)
    return
}

func NewGetRecordPlanByIdRequest() (request *GetRecordPlanByIdRequest) {
    request = &GetRecordPlanByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetRecordPlanById")
    return
}

func NewGetRecordPlanByIdResponse() (response *GetRecordPlanByIdResponse) {
    response = &GetRecordPlanByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRecordPlanById
// 本接口(GetRecordPlanById)用于根据录制计划ID获取录制计划。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) GetRecordPlanById(request *GetRecordPlanByIdRequest) (response *GetRecordPlanByIdResponse, err error) {
    if request == nil {
        request = NewGetRecordPlanByIdRequest()
    }
    response = NewGetRecordPlanByIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetRecordPlansRequest() (request *GetRecordPlansRequest) {
    request = &GetRecordPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetRecordPlans")
    return
}

func NewGetRecordPlansResponse() (response *GetRecordPlansResponse) {
    response = &GetRecordPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRecordPlans
// 本接口(GetRecordPlans)用于获取用户的全部录制计划。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetRecordPlans(request *GetRecordPlansRequest) (response *GetRecordPlansResponse, err error) {
    if request == nil {
        request = NewGetRecordPlansRequest()
    }
    response = NewGetRecordPlansResponse()
    err = c.Send(request, response)
    return
}

func NewGetTimeTemplateByIdRequest() (request *GetTimeTemplateByIdRequest) {
    request = &GetTimeTemplateByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetTimeTemplateById")
    return
}

func NewGetTimeTemplateByIdResponse() (response *GetTimeTemplateByIdResponse) {
    response = &GetTimeTemplateByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTimeTemplateById
// 本接口(GetTimeTemplateById)用于根据模板ID获取时间模板详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) GetTimeTemplateById(request *GetTimeTemplateByIdRequest) (response *GetTimeTemplateByIdResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplateByIdRequest()
    }
    response = NewGetTimeTemplateByIdResponse()
    err = c.Send(request, response)
    return
}

func NewGetTimeTemplatesRequest() (request *GetTimeTemplatesRequest) {
    request = &GetTimeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetTimeTemplates")
    return
}

func NewGetTimeTemplatesResponse() (response *GetTimeTemplatesResponse) {
    response = &GetTimeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTimeTemplates
// 本接口(GetTimeTemplates)用于获取时间模板列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) GetTimeTemplates(request *GetTimeTemplatesRequest) (response *GetTimeTemplatesResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplatesRequest()
    }
    response = NewGetTimeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewGetVideoListByConRequest() (request *GetVideoListByConRequest) {
    request = &GetVideoListByConRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "GetVideoListByCon")
    return
}

func NewGetVideoListByConResponse() (response *GetVideoListByConResponse) {
    response = &GetVideoListByConResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetVideoListByCon
// 本接口(GetVideoListByCon)用于查询设备的录制文件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) GetVideoListByCon(request *GetVideoListByConRequest) (response *GetVideoListByConResponse, err error) {
    if request == nil {
        request = NewGetVideoListByConRequest()
    }
    response = NewGetVideoListByConResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceDataRequest() (request *ModifyDeviceDataRequest) {
    request = &ModifyDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyDeviceData")
    return
}

func NewModifyDeviceDataResponse() (response *ModifyDeviceDataResponse) {
    response = &ModifyDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDeviceData
// 本接口(ModifyDeviceData)用于编辑设备信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) ModifyDeviceData(request *ModifyDeviceDataRequest) (response *ModifyDeviceDataResponse, err error) {
    if request == nil {
        request = NewModifyDeviceDataRequest()
    }
    response = NewModifyDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceGroupRequest() (request *UpdateDeviceGroupRequest) {
    request = &UpdateDeviceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "UpdateDeviceGroup")
    return
}

func NewUpdateDeviceGroupResponse() (response *UpdateDeviceGroupResponse) {
    response = &UpdateDeviceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceGroup
// 本接口(UpdateDeviceGroup)用于修改分组信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) UpdateDeviceGroup(request *UpdateDeviceGroupRequest) (response *UpdateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceGroupRequest()
    }
    response = NewUpdateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicePassWordRequest() (request *UpdateDevicePassWordRequest) {
    request = &UpdateDevicePassWordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "UpdateDevicePassWord")
    return
}

func NewUpdateDevicePassWordResponse() (response *UpdateDevicePassWordResponse) {
    response = &UpdateDevicePassWordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicePassWord
// 本接口(UpdateDevicePassWord)用于修改设备密码。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) UpdateDevicePassWord(request *UpdateDevicePassWordRequest) (response *UpdateDevicePassWordResponse, err error) {
    if request == nil {
        request = NewUpdateDevicePassWordRequest()
    }
    response = NewUpdateDevicePassWordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordPlanRequest() (request *UpdateRecordPlanRequest) {
    request = &UpdateRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "UpdateRecordPlan")
    return
}

func NewUpdateRecordPlanResponse() (response *UpdateRecordPlanResponse) {
    response = &UpdateRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecordPlan
// 本接口(UpdateRecordPlan)用于更新录制计划。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordPlanRequest()
    }
    response = NewUpdateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTimeTemplateRequest() (request *UpdateTimeTemplateRequest) {
    request = &UpdateTimeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "UpdateTimeTemplate")
    return
}

func NewUpdateTimeTemplateResponse() (response *UpdateTimeTemplateResponse) {
    response = &UpdateTimeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateTimeTemplate
// 本接口(UpdateTimeTemplate)用于更新时间模板。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
func (c *Client) UpdateTimeTemplate(request *UpdateTimeTemplateRequest) (response *UpdateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateTimeTemplateRequest()
    }
    response = NewUpdateTimeTemplateResponse()
    err = c.Send(request, response)
    return
}
