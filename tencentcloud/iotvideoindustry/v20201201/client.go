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

func NewCreateLiveChannelRequest() (request *CreateLiveChannelRequest) {
    request = &CreateLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateLiveChannel")
    
    
    return
}

func NewCreateLiveChannelResponse() (response *CreateLiveChannelResponse) {
    response = &CreateLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveChannel
// 创建直播频道
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) CreateLiveChannel(request *CreateLiveChannelRequest) (response *CreateLiveChannelResponse, err error) {
    if request == nil {
        request = NewCreateLiveChannelRequest()
    }
    
    response = NewCreateLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordPlanRequest() (request *CreateLiveRecordPlanRequest) {
    request = &CreateLiveRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateLiveRecordPlan")
    
    
    return
}

func NewCreateLiveRecordPlanResponse() (response *CreateLiveRecordPlanResponse) {
    response = &CreateLiveRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveRecordPlan
// 创建直播录制计划
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) CreateLiveRecordPlan(request *CreateLiveRecordPlanRequest) (response *CreateLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordPlanRequest()
    }
    
    response = NewCreateLiveRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMessageForwardRequest() (request *CreateMessageForwardRequest) {
    request = &CreateMessageForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateMessageForward")
    
    
    return
}

func NewCreateMessageForwardResponse() (response *CreateMessageForwardResponse) {
    response = &CreateMessageForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMessageForward
// 创建消息转发配置
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) CreateMessageForward(request *CreateMessageForwardRequest) (response *CreateMessageForwardResponse, err error) {
    if request == nil {
        request = NewCreateMessageForwardRequest()
    }
    
    response = NewCreateMessageForwardResponse()
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

func NewCreateSceneRequest() (request *CreateSceneRequest) {
    request = &CreateSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateScene")
    
    
    return
}

func NewCreateSceneResponse() (response *CreateSceneResponse) {
    response = &CreateSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScene
// 创建场景
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) CreateScene(request *CreateSceneRequest) (response *CreateSceneResponse, err error) {
    if request == nil {
        request = NewCreateSceneRequest()
    }
    
    response = NewCreateSceneResponse()
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

func NewDeleteChannelRequest() (request *DeleteChannelRequest) {
    request = &DeleteChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteChannel")
    
    
    return
}

func NewDeleteChannelResponse() (response *DeleteChannelResponse) {
    response = &DeleteChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteChannel
// 删除通道接口
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
//  INVALIDPARAMETER_DEVICEONLINE = "InvalidParameter.DeviceOnline"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteChannel(request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    if request == nil {
        request = NewDeleteChannelRequest()
    }
    
    response = NewDeleteChannelResponse()
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

func NewDeleteLiveChannelRequest() (request *DeleteLiveChannelRequest) {
    request = &DeleteLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteLiveChannel")
    
    
    return
}

func NewDeleteLiveChannelResponse() (response *DeleteLiveChannelResponse) {
    response = &DeleteLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveChannel
// 删除直播接口
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteLiveChannel(request *DeleteLiveChannelRequest) (response *DeleteLiveChannelResponse, err error) {
    if request == nil {
        request = NewDeleteLiveChannelRequest()
    }
    
    response = NewDeleteLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordPlanRequest() (request *DeleteLiveRecordPlanRequest) {
    request = &DeleteLiveRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteLiveRecordPlan")
    
    
    return
}

func NewDeleteLiveRecordPlanResponse() (response *DeleteLiveRecordPlanResponse) {
    response = &DeleteLiveRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveRecordPlan
// 删除直播录制计划
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteLiveRecordPlan(request *DeleteLiveRecordPlanRequest) (response *DeleteLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordPlanRequest()
    }
    
    response = NewDeleteLiveRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveVideoListRequest() (request *DeleteLiveVideoListRequest) {
    request = &DeleteLiveVideoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteLiveVideoList")
    
    
    return
}

func NewDeleteLiveVideoListResponse() (response *DeleteLiveVideoListResponse) {
    response = &DeleteLiveVideoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveVideoList
// 直播录像删除
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteLiveVideoList(request *DeleteLiveVideoListRequest) (response *DeleteLiveVideoListResponse, err error) {
    if request == nil {
        request = NewDeleteLiveVideoListRequest()
    }
    
    response = NewDeleteLiveVideoListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMessageForwardRequest() (request *DeleteMessageForwardRequest) {
    request = &DeleteMessageForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteMessageForward")
    
    
    return
}

func NewDeleteMessageForwardResponse() (response *DeleteMessageForwardResponse) {
    response = &DeleteMessageForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMessageForward
// 删除消息转发配置
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteMessageForward(request *DeleteMessageForwardRequest) (response *DeleteMessageForwardResponse, err error) {
    if request == nil {
        request = NewDeleteMessageForwardRequest()
    }
    
    response = NewDeleteMessageForwardResponse()
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

func NewDeleteSceneRequest() (request *DeleteSceneRequest) {
    request = &DeleteSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteScene")
    
    
    return
}

func NewDeleteSceneResponse() (response *DeleteSceneResponse) {
    response = &DeleteSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScene
// 删除场景
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteScene(request *DeleteSceneRequest) (response *DeleteSceneResponse, err error) {
    if request == nil {
        request = NewDeleteSceneRequest()
    }
    
    response = NewDeleteSceneResponse()
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

func NewDeleteVideoListRequest() (request *DeleteVideoListRequest) {
    request = &DeleteVideoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteVideoList")
    
    
    return
}

func NewDeleteVideoListResponse() (response *DeleteVideoListResponse) {
    response = &DeleteVideoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVideoList
// 删除录像存储列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DeleteVideoList(request *DeleteVideoListRequest) (response *DeleteVideoListResponse, err error) {
    if request == nil {
        request = NewDeleteVideoListRequest()
    }
    
    response = NewDeleteVideoListResponse()
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

func NewDescribeBindSceneDevicesRequest() (request *DescribeBindSceneDevicesRequest) {
    request = &DescribeBindSceneDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeBindSceneDevices")
    
    
    return
}

func NewDescribeBindSceneDevicesResponse() (response *DescribeBindSceneDevicesResponse) {
    response = &DescribeBindSceneDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBindSceneDevices
// 获取场景绑定设备列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeBindSceneDevices(request *DescribeBindSceneDevicesRequest) (response *DescribeBindSceneDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeBindSceneDevicesRequest()
    }
    
    response = NewDescribeBindSceneDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelsByLiveRecordPlanRequest() (request *DescribeChannelsByLiveRecordPlanRequest) {
    request = &DescribeChannelsByLiveRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeChannelsByLiveRecordPlan")
    
    
    return
}

func NewDescribeChannelsByLiveRecordPlanResponse() (response *DescribeChannelsByLiveRecordPlanResponse) {
    response = &DescribeChannelsByLiveRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelsByLiveRecordPlan
// 根据直播录制计划获取频道列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeChannelsByLiveRecordPlan(request *DescribeChannelsByLiveRecordPlanRequest) (response *DescribeChannelsByLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewDescribeChannelsByLiveRecordPlanRequest()
    }
    
    response = NewDescribeChannelsByLiveRecordPlanResponse()
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

func NewDescribeIPCChannelsRequest() (request *DescribeIPCChannelsRequest) {
    request = &DescribeIPCChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeIPCChannels")
    
    
    return
}

func NewDescribeIPCChannelsResponse() (response *DescribeIPCChannelsResponse) {
    response = &DescribeIPCChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIPCChannels
// 获取IPC设备下属通道
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
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
func (c *Client) DescribeIPCChannels(request *DescribeIPCChannelsRequest) (response *DescribeIPCChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeIPCChannelsRequest()
    }
    
    response = NewDescribeIPCChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveChannelRequest() (request *DescribeLiveChannelRequest) {
    request = &DescribeLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveChannel")
    
    
    return
}

func NewDescribeLiveChannelResponse() (response *DescribeLiveChannelResponse) {
    response = &DescribeLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveChannel
// 直播详情接口
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveChannel(request *DescribeLiveChannelRequest) (response *DescribeLiveChannelResponse, err error) {
    if request == nil {
        request = NewDescribeLiveChannelRequest()
    }
    
    response = NewDescribeLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveChannelListRequest() (request *DescribeLiveChannelListRequest) {
    request = &DescribeLiveChannelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveChannelList")
    
    
    return
}

func NewDescribeLiveChannelListResponse() (response *DescribeLiveChannelListResponse) {
    response = &DescribeLiveChannelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveChannelList
// 直播列表接口
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveChannelList(request *DescribeLiveChannelListRequest) (response *DescribeLiveChannelListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveChannelListRequest()
    }
    
    response = NewDescribeLiveChannelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordPlanByIdRequest() (request *DescribeLiveRecordPlanByIdRequest) {
    request = &DescribeLiveRecordPlanByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveRecordPlanById")
    
    
    return
}

func NewDescribeLiveRecordPlanByIdResponse() (response *DescribeLiveRecordPlanByIdResponse) {
    response = &DescribeLiveRecordPlanByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveRecordPlanById
// 获取直播录制计划详情
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveRecordPlanById(request *DescribeLiveRecordPlanByIdRequest) (response *DescribeLiveRecordPlanByIdResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordPlanByIdRequest()
    }
    
    response = NewDescribeLiveRecordPlanByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordPlanIdsRequest() (request *DescribeLiveRecordPlanIdsRequest) {
    request = &DescribeLiveRecordPlanIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveRecordPlanIds")
    
    
    return
}

func NewDescribeLiveRecordPlanIdsResponse() (response *DescribeLiveRecordPlanIdsResponse) {
    response = &DescribeLiveRecordPlanIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveRecordPlanIds
// 获取直播录制计划列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveRecordPlanIds(request *DescribeLiveRecordPlanIdsRequest) (response *DescribeLiveRecordPlanIdsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordPlanIdsRequest()
    }
    
    response = NewDescribeLiveRecordPlanIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamRequest() (request *DescribeLiveStreamRequest) {
    request = &DescribeLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveStream")
    
    
    return
}

func NewDescribeLiveStreamResponse() (response *DescribeLiveStreamResponse) {
    response = &DescribeLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStream
// 直播拉流接口
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveStream(request *DescribeLiveStreamRequest) (response *DescribeLiveStreamResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamRequest()
    }
    
    response = NewDescribeLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveVideoListRequest() (request *DescribeLiveVideoListRequest) {
    request = &DescribeLiveVideoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeLiveVideoList")
    
    
    return
}

func NewDescribeLiveVideoListResponse() (response *DescribeLiveVideoListResponse) {
    response = &DescribeLiveVideoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveVideoList
// 直播录像回放列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeLiveVideoList(request *DescribeLiveVideoListRequest) (response *DescribeLiveVideoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveVideoListRequest()
    }
    
    response = NewDescribeLiveVideoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageForwardRequest() (request *DescribeMessageForwardRequest) {
    request = &DescribeMessageForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeMessageForward")
    
    
    return
}

func NewDescribeMessageForwardResponse() (response *DescribeMessageForwardResponse) {
    response = &DescribeMessageForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMessageForward
// 查看消息转发配置详情
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeMessageForward(request *DescribeMessageForwardRequest) (response *DescribeMessageForwardResponse, err error) {
    if request == nil {
        request = NewDescribeMessageForwardRequest()
    }
    
    response = NewDescribeMessageForwardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageForwardsRequest() (request *DescribeMessageForwardsRequest) {
    request = &DescribeMessageForwardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeMessageForwards")
    
    
    return
}

func NewDescribeMessageForwardsResponse() (response *DescribeMessageForwardsResponse) {
    response = &DescribeMessageForwardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMessageForwards
// 查看消息转发配置列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeMessageForwards(request *DescribeMessageForwardsRequest) (response *DescribeMessageForwardsResponse, err error) {
    if request == nil {
        request = NewDescribeMessageForwardsRequest()
    }
    
    response = NewDescribeMessageForwardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordDatesByLiveRequest() (request *DescribeRecordDatesByLiveRequest) {
    request = &DescribeRecordDatesByLiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeRecordDatesByLive")
    
    
    return
}

func NewDescribeRecordDatesByLiveResponse() (response *DescribeRecordDatesByLiveResponse) {
    response = &DescribeRecordDatesByLiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordDatesByLive
// 直播录像存储日期列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeRecordDatesByLive(request *DescribeRecordDatesByLiveRequest) (response *DescribeRecordDatesByLiveResponse, err error) {
    if request == nil {
        request = NewDescribeRecordDatesByLiveRequest()
    }
    
    response = NewDescribeRecordDatesByLiveResponse()
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

func NewDescribeScenesRequest() (request *DescribeScenesRequest) {
    request = &DescribeScenesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeScenes")
    
    
    return
}

func NewDescribeScenesResponse() (response *DescribeScenesResponse) {
    response = &DescribeScenesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScenes
// 获取场景列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeScenes(request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    if request == nil {
        request = NewDescribeScenesRequest()
    }
    
    response = NewDescribeScenesResponse()
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

func NewDescribeSubscriptionStatusRequest() (request *DescribeSubscriptionStatusRequest) {
    request = &DescribeSubscriptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeSubscriptionStatus")
    
    
    return
}

func NewDescribeSubscriptionStatusResponse() (response *DescribeSubscriptionStatusResponse) {
    response = &DescribeSubscriptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubscriptionStatus
// 查询主设备订阅状态
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) DescribeSubscriptionStatus(request *DescribeSubscriptionStatusRequest) (response *DescribeSubscriptionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionStatusRequest()
    }
    
    response = NewDescribeSubscriptionStatusResponse()
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

func NewModifyBindPlanLiveChannelRequest() (request *ModifyBindPlanLiveChannelRequest) {
    request = &ModifyBindPlanLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyBindPlanLiveChannel")
    
    
    return
}

func NewModifyBindPlanLiveChannelResponse() (response *ModifyBindPlanLiveChannelResponse) {
    response = &ModifyBindPlanLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBindPlanLiveChannel
// 直播录制计划绑定解绑直播频道
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) ModifyBindPlanLiveChannel(request *ModifyBindPlanLiveChannelRequest) (response *ModifyBindPlanLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyBindPlanLiveChannelRequest()
    }
    
    response = NewModifyBindPlanLiveChannelResponse()
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

func NewModifyLiveChannelRequest() (request *ModifyLiveChannelRequest) {
    request = &ModifyLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyLiveChannel")
    
    
    return
}

func NewModifyLiveChannelResponse() (response *ModifyLiveChannelResponse) {
    response = &ModifyLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveChannel
// 编辑直播接口
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifyLiveChannel(request *ModifyLiveChannelRequest) (response *ModifyLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyLiveChannelRequest()
    }
    
    response = NewModifyLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordPlanRequest() (request *ModifyLiveRecordPlanRequest) {
    request = &ModifyLiveRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyLiveRecordPlan")
    
    
    return
}

func NewModifyLiveRecordPlanResponse() (response *ModifyLiveRecordPlanResponse) {
    response = &ModifyLiveRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveRecordPlan
// 编辑直播录制计划
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifyLiveRecordPlan(request *ModifyLiveRecordPlanRequest) (response *ModifyLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordPlanRequest()
    }
    
    response = NewModifyLiveRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveVideoRequest() (request *ModifyLiveVideoRequest) {
    request = &ModifyLiveVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyLiveVideo")
    
    
    return
}

func NewModifyLiveVideoResponse() (response *ModifyLiveVideoResponse) {
    response = &ModifyLiveVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveVideo
// 直播录像编辑
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifyLiveVideo(request *ModifyLiveVideoRequest) (response *ModifyLiveVideoResponse, err error) {
    if request == nil {
        request = NewModifyLiveVideoRequest()
    }
    
    response = NewModifyLiveVideoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMessageForwardRequest() (request *ModifyMessageForwardRequest) {
    request = &ModifyMessageForwardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyMessageForward")
    
    
    return
}

func NewModifyMessageForwardResponse() (response *ModifyMessageForwardResponse) {
    response = &ModifyMessageForwardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMessageForward
// 修改消息转发配置
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifyMessageForward(request *ModifyMessageForwardRequest) (response *ModifyMessageForwardResponse, err error) {
    if request == nil {
        request = NewModifyMessageForwardRequest()
    }
    
    response = NewModifyMessageForwardResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscriptionStatusRequest() (request *ModifySubscriptionStatusRequest) {
    request = &ModifySubscriptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifySubscriptionStatus")
    
    
    return
}

func NewModifySubscriptionStatusResponse() (response *ModifySubscriptionStatusResponse) {
    response = &ModifySubscriptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscriptionStatus
// 编辑设备订阅状态
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifySubscriptionStatus(request *ModifySubscriptionStatusRequest) (response *ModifySubscriptionStatusResponse, err error) {
    if request == nil {
        request = NewModifySubscriptionStatusRequest()
    }
    
    response = NewModifySubscriptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVideoInfoRequest() (request *ModifyVideoInfoRequest) {
    request = &ModifyVideoInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyVideoInfo")
    
    
    return
}

func NewModifyVideoInfoResponse() (response *ModifyVideoInfoResponse) {
    response = &ModifyVideoInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVideoInfo
// 修改录像存储列表
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
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
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
func (c *Client) ModifyVideoInfo(request *ModifyVideoInfoRequest) (response *ModifyVideoInfoResponse, err error) {
    if request == nil {
        request = NewModifyVideoInfoRequest()
    }
    
    response = NewModifyVideoInfoResponse()
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
