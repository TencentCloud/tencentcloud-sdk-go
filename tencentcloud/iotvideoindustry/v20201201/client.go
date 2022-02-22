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
    "context"
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
func (c *Client) BindGroupDevices(request *BindGroupDevicesRequest) (response *BindGroupDevicesResponse, err error) {
    if request == nil {
        request = NewBindGroupDevicesRequest()
    }
    
    response = NewBindGroupDevicesResponse()
    err = c.Send(request, response)
    return
}

// BindGroupDevices
// 本接口(BindGroupDevices) 用于绑定设备到分组。
func (c *Client) BindGroupDevicesWithContext(ctx context.Context, request *BindGroupDevicesRequest) (response *BindGroupDevicesResponse, err error) {
    if request == nil {
        request = NewBindGroupDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindGroupDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewControlChannelLocalRecordRequest() (request *ControlChannelLocalRecordRequest) {
    request = &ControlChannelLocalRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlChannelLocalRecord")
    
    
    return
}

func NewControlChannelLocalRecordResponse() (response *ControlChannelLocalRecordResponse) {
    response = &ControlChannelLocalRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlChannelLocalRecord
// 本接口（ControlChannelLocalRecord）用于对通道本地回放流进行控制，包括暂停、播放、拉动、结束等
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCEUNAVAILABLE_GBPROTOCOLEXECEXCEPTION = "ResourceUnavailable.GBProtocolExecException"
func (c *Client) ControlChannelLocalRecord(request *ControlChannelLocalRecordRequest) (response *ControlChannelLocalRecordResponse, err error) {
    if request == nil {
        request = NewControlChannelLocalRecordRequest()
    }
    
    response = NewControlChannelLocalRecordResponse()
    err = c.Send(request, response)
    return
}

// ControlChannelLocalRecord
// 本接口（ControlChannelLocalRecord）用于对通道本地回放流进行控制，包括暂停、播放、拉动、结束等
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCEUNAVAILABLE_GBPROTOCOLEXECEXCEPTION = "ResourceUnavailable.GBProtocolExecException"
func (c *Client) ControlChannelLocalRecordWithContext(ctx context.Context, request *ControlChannelLocalRecordRequest) (response *ControlChannelLocalRecordResponse, err error) {
    if request == nil {
        request = NewControlChannelLocalRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewControlChannelLocalRecordResponse()
    err = c.Send(request, response)
    return
}

func NewControlChannelPTZRequest() (request *ControlChannelPTZRequest) {
    request = &ControlChannelPTZRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlChannelPTZ")
    
    
    return
}

func NewControlChannelPTZResponse() (response *ControlChannelPTZResponse) {
    response = &ControlChannelPTZResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlChannelPTZ
// 本接口(ControlChannelPTZ) 用于对支持GB28181 PTZ信令的设备进行指定通道的远程控制。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
func (c *Client) ControlChannelPTZ(request *ControlChannelPTZRequest) (response *ControlChannelPTZResponse, err error) {
    if request == nil {
        request = NewControlChannelPTZRequest()
    }
    
    response = NewControlChannelPTZResponse()
    err = c.Send(request, response)
    return
}

// ControlChannelPTZ
// 本接口(ControlChannelPTZ) 用于对支持GB28181 PTZ信令的设备进行指定通道的远程控制。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"
func (c *Client) ControlChannelPTZWithContext(ctx context.Context, request *ControlChannelPTZRequest) (response *ControlChannelPTZResponse, err error) {
    if request == nil {
        request = NewControlChannelPTZRequest()
    }
    request.SetContext(ctx)
    
    response = NewControlChannelPTZResponse()
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
// 请使用ControlChannelPTZ接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) ControlDevicePTZ(request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    if request == nil {
        request = NewControlDevicePTZRequest()
    }
    
    response = NewControlDevicePTZResponse()
    err = c.Send(request, response)
    return
}

// ControlDevicePTZ
// 本接口(ControlDevicePTZ) 用于对支持GB28181 PTZ信令的设备进行远程控制。
//
// 请使用ControlChannelPTZ接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) ControlDevicePTZWithContext(ctx context.Context, request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    if request == nil {
        request = NewControlDevicePTZRequest()
    }
    request.SetContext(ctx)
    
    response = NewControlDevicePTZResponse()
    err = c.Send(request, response)
    return
}

func NewControlHomePositionRequest() (request *ControlHomePositionRequest) {
    request = &ControlHomePositionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlHomePosition")
    
    
    return
}

func NewControlHomePositionResponse() (response *ControlHomePositionResponse) {
    response = &ControlHomePositionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlHomePosition
// 看守位控制
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) ControlHomePosition(request *ControlHomePositionRequest) (response *ControlHomePositionResponse, err error) {
    if request == nil {
        request = NewControlHomePositionRequest()
    }
    
    response = NewControlHomePositionResponse()
    err = c.Send(request, response)
    return
}

// ControlHomePosition
// 看守位控制
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) ControlHomePositionWithContext(ctx context.Context, request *ControlHomePositionRequest) (response *ControlHomePositionResponse, err error) {
    if request == nil {
        request = NewControlHomePositionRequest()
    }
    request.SetContext(ctx)
    
    response = NewControlHomePositionResponse()
    err = c.Send(request, response)
    return
}

func NewControlPresetRequest() (request *ControlPresetRequest) {
    request = &ControlPresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ControlPreset")
    
    
    return
}

func NewControlPresetResponse() (response *ControlPresetResponse) {
    response = &ControlPresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlPreset
// 预置位控制
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
func (c *Client) ControlPreset(request *ControlPresetRequest) (response *ControlPresetResponse, err error) {
    if request == nil {
        request = NewControlPresetRequest()
    }
    
    response = NewControlPresetResponse()
    err = c.Send(request, response)
    return
}

// ControlPreset
// 预置位控制
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
func (c *Client) ControlPresetWithContext(ctx context.Context, request *ControlPresetRequest) (response *ControlPresetResponse, err error) {
    if request == nil {
        request = NewControlPresetRequest()
    }
    request.SetContext(ctx)
    
    response = NewControlPresetResponse()
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
// 请使用ControlChannelLocalRecord接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
func (c *Client) ControlRecordStream(request *ControlRecordStreamRequest) (response *ControlRecordStreamResponse, err error) {
    if request == nil {
        request = NewControlRecordStreamRequest()
    }
    
    response = NewControlRecordStreamResponse()
    err = c.Send(request, response)
    return
}

// ControlRecordStream
// 对回放流进行控制，包括暂停、播放、拉动、结束等
//
// 请使用ControlChannelLocalRecord接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
func (c *Client) ControlRecordStreamWithContext(ctx context.Context, request *ControlRecordStreamRequest) (response *ControlRecordStreamResponse, err error) {
    if request == nil {
        request = NewControlRecordStreamRequest()
    }
    request.SetContext(ctx)
    
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

// CreateDevice
// 本接口(CreateDevice) 用于创建设备。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
func (c *Client) CreateDeviceWithContext(ctx context.Context, request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
func (c *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeviceGroupRequest()
    }
    
    response = NewCreateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateDeviceGroup
// 本接口(CreateDeviceGroup) 用于创建设备管理分组。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
//  UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"
func (c *Client) CreateDeviceGroupWithContext(ctx context.Context, request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeviceGroupRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) CreateLiveChannelWithContext(ctx context.Context, request *CreateLiveChannelRequest) (response *CreateLiveChannelResponse, err error) {
    if request == nil {
        request = NewCreateLiveChannelRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) CreateLiveRecordPlanWithContext(ctx context.Context, request *CreateLiveRecordPlanRequest) (response *CreateLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordPlanRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) CreateMessageForwardWithContext(ctx context.Context, request *CreateMessageForwardRequest) (response *CreateMessageForwardResponse, err error) {
    if request == nil {
        request = NewCreateMessageForwardRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用CreateRecordingPlan代替
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateRecordPlan(request *CreateRecordPlanRequest) (response *CreateRecordPlanResponse, err error) {
    if request == nil {
        request = NewCreateRecordPlanRequest()
    }
    
    response = NewCreateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

// CreateRecordPlan
// 本接口(CreateRecordPlan) 用于创建录制计划，使设备与时间模板绑定，以便及时启动录制
//
// 请使用CreateRecordingPlan代替
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateRecordPlanWithContext(ctx context.Context, request *CreateRecordPlanRequest) (response *CreateRecordPlanResponse, err error) {
    if request == nil {
        request = NewCreateRecordPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordingPlanRequest() (request *CreateRecordingPlanRequest) {
    request = &CreateRecordingPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "CreateRecordingPlan")
    
    
    return
}

func NewCreateRecordingPlanResponse() (response *CreateRecordingPlanResponse) {
    response = &CreateRecordingPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecordingPlan
// 本接口(CreateRecordingPlan) 用于创建录制计划，使通道与时间模板绑定，以便及时启动录制
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateRecordingPlan(request *CreateRecordingPlanRequest) (response *CreateRecordingPlanResponse, err error) {
    if request == nil {
        request = NewCreateRecordingPlanRequest()
    }
    
    response = NewCreateRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

// CreateRecordingPlan
// 本接口(CreateRecordingPlan) 用于创建录制计划，使通道与时间模板绑定，以便及时启动录制
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateRecordingPlanWithContext(ctx context.Context, request *CreateRecordingPlanRequest) (response *CreateRecordingPlanResponse, err error) {
    if request == nil {
        request = NewCreateRecordingPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRecordingPlanResponse()
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
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) CreateScene(request *CreateSceneRequest) (response *CreateSceneResponse, err error) {
    if request == nil {
        request = NewCreateSceneRequest()
    }
    
    response = NewCreateSceneResponse()
    err = c.Send(request, response)
    return
}

// CreateScene
// 创建场景
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
func (c *Client) CreateSceneWithContext(ctx context.Context, request *CreateSceneRequest) (response *CreateSceneResponse, err error) {
    if request == nil {
        request = NewCreateSceneRequest()
    }
    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateTimeTemplate(request *CreateTimeTemplateRequest) (response *CreateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTimeTemplateRequest()
    }
    
    response = NewCreateTimeTemplateResponse()
    err = c.Send(request, response)
    return
}

// CreateTimeTemplate
// 本接口(CreateTimeTemplate) 用于根据模板描述的具体录制时间片段，创建定制化的时间模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) CreateTimeTemplateWithContext(ctx context.Context, request *CreateTimeTemplateRequest) (response *CreateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTimeTemplateRequest()
    }
    request.SetContext(ctx)
    
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
// 本接口用于删除设备下的通道
//
// 注意： 在线状态的设备不允许删除
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
func (c *Client) DeleteChannel(request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    if request == nil {
        request = NewDeleteChannelRequest()
    }
    
    response = NewDeleteChannelResponse()
    err = c.Send(request, response)
    return
}

// DeleteChannel
// 本接口用于删除设备下的通道
//
// 注意： 在线状态的设备不允许删除
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
func (c *Client) DeleteChannelWithContext(ctx context.Context, request *DeleteChannelRequest) (response *DeleteChannelResponse, err error) {
    if request == nil {
        request = NewDeleteChannelRequest()
    }
    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

// DeleteDevice
// 本接口(DeleteDevice)用于删除设备。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
func (c *Client) DeleteDeviceWithContext(ctx context.Context, request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
func (c *Client) DeleteDeviceGroup(request *DeleteDeviceGroupRequest) (response *DeleteDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupRequest()
    }
    
    response = NewDeleteDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteDeviceGroup
// 本接口(DeleteDeviceGroup)用于删除分组。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
func (c *Client) DeleteDeviceGroupWithContext(ctx context.Context, request *DeleteDeviceGroupRequest) (response *DeleteDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceGroupRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DeleteLiveChannelWithContext(ctx context.Context, request *DeleteLiveChannelRequest) (response *DeleteLiveChannelResponse, err error) {
    if request == nil {
        request = NewDeleteLiveChannelRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DeleteLiveRecordPlanWithContext(ctx context.Context, request *DeleteLiveRecordPlanRequest) (response *DeleteLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordPlanRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DeleteLiveVideoListWithContext(ctx context.Context, request *DeleteLiveVideoListRequest) (response *DeleteLiveVideoListResponse, err error) {
    if request == nil {
        request = NewDeleteLiveVideoListRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DeleteMessageForwardWithContext(ctx context.Context, request *DeleteMessageForwardRequest) (response *DeleteMessageForwardResponse, err error) {
    if request == nil {
        request = NewDeleteMessageForwardRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DeleteRecordingPlan接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordPlanRequest()
    }
    
    response = NewDeleteRecordPlanResponse()
    err = c.Send(request, response)
    return
}

// DeleteRecordPlan
// 本接口(DeleteRecordPlan)用于删除录制计划
//
// 录制计划删除的同时，会停止该录制计划下的全部录制任务。
//
// 请使用DeleteRecordingPlan接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlanWithContext(ctx context.Context, request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordingPlanRequest() (request *DeleteRecordingPlanRequest) {
    request = &DeleteRecordingPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteRecordingPlan")
    
    
    return
}

func NewDeleteRecordingPlanResponse() (response *DeleteRecordingPlanResponse) {
    response = &DeleteRecordingPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordingPlan
// 本接口(DeleteRecordingPlan)用于删除录制计划
//
// 录制计划删除的同时，会停止该录制计划下的全部录制任务。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordingPlan(request *DeleteRecordingPlanRequest) (response *DeleteRecordingPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingPlanRequest()
    }
    
    response = NewDeleteRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

// DeleteRecordingPlan
// 本接口(DeleteRecordingPlan)用于删除录制计划
//
// 录制计划删除的同时，会停止该录制计划下的全部录制任务。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordingPlanWithContext(ctx context.Context, request *DeleteRecordingPlanRequest) (response *DeleteRecordingPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRecordingPlanResponse()
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
func (c *Client) DeleteSceneWithContext(ctx context.Context, request *DeleteSceneRequest) (response *DeleteSceneResponse, err error) {
    if request == nil {
        request = NewDeleteSceneRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteTimeTemplate(request *DeleteTimeTemplateRequest) (response *DeleteTimeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTimeTemplateRequest()
    }
    
    response = NewDeleteTimeTemplateResponse()
    err = c.Send(request, response)
    return
}

// DeleteTimeTemplate
// 本接口(DeleteTimeTemplate) 用于删除时间模板。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteTimeTemplateWithContext(ctx context.Context, request *DeleteTimeTemplateRequest) (response *DeleteTimeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTimeTemplateRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteVideoList(request *DeleteVideoListRequest) (response *DeleteVideoListResponse, err error) {
    if request == nil {
        request = NewDeleteVideoListRequest()
    }
    
    response = NewDeleteVideoListResponse()
    err = c.Send(request, response)
    return
}

// DeleteVideoList
// 删除录像存储列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteVideoListWithContext(ctx context.Context, request *DeleteVideoListRequest) (response *DeleteVideoListResponse, err error) {
    if request == nil {
        request = NewDeleteVideoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteVideoListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWarningRequest() (request *DeleteWarningRequest) {
    request = &DeleteWarningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DeleteWarning")
    
    
    return
}

func NewDeleteWarningResponse() (response *DeleteWarningResponse) {
    response = &DeleteWarningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWarning
// 设备告警-删除告警
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteWarning(request *DeleteWarningRequest) (response *DeleteWarningResponse, err error) {
    if request == nil {
        request = NewDeleteWarningRequest()
    }
    
    response = NewDeleteWarningResponse()
    err = c.Send(request, response)
    return
}

// DeleteWarning
// 设备告警-删除告警
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) DeleteWarningWithContext(ctx context.Context, request *DeleteWarningRequest) (response *DeleteWarningResponse, err error) {
    if request == nil {
        request = NewDeleteWarningRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteWarningResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalEventsRequest() (request *DescribeAbnormalEventsRequest) {
    request = &DescribeAbnormalEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeAbnormalEvents")
    
    
    return
}

func NewDescribeAbnormalEventsResponse() (response *DescribeAbnormalEventsResponse) {
    response = &DescribeAbnormalEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalEvents
// 获取异常事件统计
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
func (c *Client) DescribeAbnormalEvents(request *DescribeAbnormalEventsRequest) (response *DescribeAbnormalEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventsRequest()
    }
    
    response = NewDescribeAbnormalEventsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAbnormalEvents
// 获取异常事件统计
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
func (c *Client) DescribeAbnormalEventsWithContext(ctx context.Context, request *DescribeAbnormalEventsRequest) (response *DescribeAbnormalEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAbnormalEventsResponse()
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
// 请使用DescribeDevicesList接口
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
func (c *Client) DescribeAllDeviceList(request *DescribeAllDeviceListRequest) (response *DescribeAllDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeAllDeviceListRequest()
    }
    
    response = NewDescribeAllDeviceListResponse()
    err = c.Send(request, response)
    return
}

// DescribeAllDeviceList
// 本接口(DescribeAllDeviceList) 用于获取设备列表。
//
// 请使用DescribeDevicesList接口
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
func (c *Client) DescribeAllDeviceListWithContext(ctx context.Context, request *DescribeAllDeviceListRequest) (response *DescribeAllDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeAllDeviceListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAllDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindSceneChannelsRequest() (request *DescribeBindSceneChannelsRequest) {
    request = &DescribeBindSceneChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeBindSceneChannels")
    
    
    return
}

func NewDescribeBindSceneChannelsResponse() (response *DescribeBindSceneChannelsResponse) {
    response = &DescribeBindSceneChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBindSceneChannels
// 获取场景绑定通道列表
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
//  INVALIDPARAMETERVALUE_DEVICEDATAMAPERROR = "InvalidParameterValue.DeviceDataMapError"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_RULELIMIT = "InvalidParameterValue.RuleLimit"
//  INVALIDPARAMETERVALUE_RULENOTEXIST = "InvalidParameterValue.RuleNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
//  INVALIDPARAMETERVALUE_TYPENOTSUPPORT = "InvalidParameterValue.TypeNotSupport"
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
//  UNSUPPORTEDOPERATION_LIVECHANNELBINDEXIST = "UnsupportedOperation.LiveChannelBindExist"
//  UNSUPPORTEDOPERATION_NOPERMISSION = "UnsupportedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_RULEDUPKEYEXIST = "UnsupportedOperation.RuleDupKeyExist"
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeBindSceneChannels(request *DescribeBindSceneChannelsRequest) (response *DescribeBindSceneChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeBindSceneChannelsRequest()
    }
    
    response = NewDescribeBindSceneChannelsResponse()
    err = c.Send(request, response)
    return
}

// DescribeBindSceneChannels
// 获取场景绑定通道列表
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
//  INVALIDPARAMETERVALUE_DEVICEDATAMAPERROR = "InvalidParameterValue.DeviceDataMapError"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_RULELIMIT = "InvalidParameterValue.RuleLimit"
//  INVALIDPARAMETERVALUE_RULENOTEXIST = "InvalidParameterValue.RuleNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
//  INVALIDPARAMETERVALUE_TYPENOTSUPPORT = "InvalidParameterValue.TypeNotSupport"
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
//  UNSUPPORTEDOPERATION_LIVECHANNELBINDEXIST = "UnsupportedOperation.LiveChannelBindExist"
//  UNSUPPORTEDOPERATION_NOPERMISSION = "UnsupportedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_RULEDUPKEYEXIST = "UnsupportedOperation.RuleDupKeyExist"
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeBindSceneChannelsWithContext(ctx context.Context, request *DescribeBindSceneChannelsRequest) (response *DescribeBindSceneChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeBindSceneChannelsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBindSceneChannelsResponse()
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
//  INVALIDPARAMETER_DEVICEONLINE = "InvalidParameter.DeviceOnline"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEDATAMAPERROR = "InvalidParameterValue.DeviceDataMapError"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_RULELIMIT = "InvalidParameterValue.RuleLimit"
//  INVALIDPARAMETERVALUE_RULENOTEXIST = "InvalidParameterValue.RuleNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
//  INVALIDPARAMETERVALUE_TYPENOTSUPPORT = "InvalidParameterValue.TypeNotSupport"
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
//  UNSUPPORTEDOPERATION_LIVECHANNELBINDEXIST = "UnsupportedOperation.LiveChannelBindExist"
//  UNSUPPORTEDOPERATION_NOPERMISSION = "UnsupportedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_RULEDUPKEYEXIST = "UnsupportedOperation.RuleDupKeyExist"
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
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
//  INVALIDPARAMETER_DEVICEONLINE = "InvalidParameter.DeviceOnline"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"
//  INVALIDPARAMETERVALUE_DEVICEDATAMAPERROR = "InvalidParameterValue.DeviceDataMapError"
//  INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"
//  INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"
//  INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"
//  INVALIDPARAMETERVALUE_RULELIMIT = "InvalidParameterValue.RuleLimit"
//  INVALIDPARAMETERVALUE_RULENOTEXIST = "InvalidParameterValue.RuleNotExist"
//  INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"
//  INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"
//  INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
//  INVALIDPARAMETERVALUE_TYPENOTSUPPORT = "InvalidParameterValue.TypeNotSupport"
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
//  UNSUPPORTEDOPERATION_LIVECHANNELBINDEXIST = "UnsupportedOperation.LiveChannelBindExist"
//  UNSUPPORTEDOPERATION_NOPERMISSION = "UnsupportedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_RULEDUPKEYEXIST = "UnsupportedOperation.RuleDupKeyExist"
//  UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"
//  UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"
//  UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
//  UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeBindSceneDevicesWithContext(ctx context.Context, request *DescribeBindSceneDevicesRequest) (response *DescribeBindSceneDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeBindSceneDevicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBindSceneDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelLiveStreamURLRequest() (request *DescribeChannelLiveStreamURLRequest) {
    request = &DescribeChannelLiveStreamURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeChannelLiveStreamURL")
    
    
    return
}

func NewDescribeChannelLiveStreamURLResponse() (response *DescribeChannelLiveStreamURLResponse) {
    response = &DescribeChannelLiveStreamURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelLiveStreamURL
// 本接口(DescribeChannelLiveStreamURL)用于获取设备指定通道实时流地址，地址是动态生成，如重新播放需要调用此接口重新获取最新播放地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelLiveStreamURL(request *DescribeChannelLiveStreamURLRequest) (response *DescribeChannelLiveStreamURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelLiveStreamURLRequest()
    }
    
    response = NewDescribeChannelLiveStreamURLResponse()
    err = c.Send(request, response)
    return
}

// DescribeChannelLiveStreamURL
// 本接口(DescribeChannelLiveStreamURL)用于获取设备指定通道实时流地址，地址是动态生成，如重新播放需要调用此接口重新获取最新播放地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelLiveStreamURLWithContext(ctx context.Context, request *DescribeChannelLiveStreamURLRequest) (response *DescribeChannelLiveStreamURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelLiveStreamURLRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChannelLiveStreamURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelLocalRecordURLRequest() (request *DescribeChannelLocalRecordURLRequest) {
    request = &DescribeChannelLocalRecordURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeChannelLocalRecordURL")
    
    
    return
}

func NewDescribeChannelLocalRecordURLResponse() (response *DescribeChannelLocalRecordURLResponse) {
    response = &DescribeChannelLocalRecordURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelLocalRecordURL
// 本接口（DescribeChannelLocalRecordURL）用于将NVR等设备对应通道本地回放文件，通过GB28181信令推送至云端，并生成对应的实时视频流URL，流地址URL是动态生成，如需重新播放请重新调用此接口获取最新地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelLocalRecordURL(request *DescribeChannelLocalRecordURLRequest) (response *DescribeChannelLocalRecordURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelLocalRecordURLRequest()
    }
    
    response = NewDescribeChannelLocalRecordURLResponse()
    err = c.Send(request, response)
    return
}

// DescribeChannelLocalRecordURL
// 本接口（DescribeChannelLocalRecordURL）用于将NVR等设备对应通道本地回放文件，通过GB28181信令推送至云端，并生成对应的实时视频流URL，流地址URL是动态生成，如需重新播放请重新调用此接口获取最新地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelLocalRecordURLWithContext(ctx context.Context, request *DescribeChannelLocalRecordURLRequest) (response *DescribeChannelLocalRecordURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelLocalRecordURLRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChannelLocalRecordURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelStreamURLRequest() (request *DescribeChannelStreamURLRequest) {
    request = &DescribeChannelStreamURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeChannelStreamURL")
    
    
    return
}

func NewDescribeChannelStreamURLResponse() (response *DescribeChannelStreamURLResponse) {
    response = &DescribeChannelStreamURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelStreamURL
// 本接口(DescribeChannelStreamURL)用于获取设备指定通道实时流地址，地址是动态生成，如重新播放需要调用此接口重新获取最新播放地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelStreamURL(request *DescribeChannelStreamURLRequest) (response *DescribeChannelStreamURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelStreamURLRequest()
    }
    
    response = NewDescribeChannelStreamURLResponse()
    err = c.Send(request, response)
    return
}

// DescribeChannelStreamURL
// 本接口(DescribeChannelStreamURL)用于获取设备指定通道实时流地址，地址是动态生成，如重新播放需要调用此接口重新获取最新播放地址。
//
// 正常推流，如未设置对应录制计划，且180s无人观看此流，将会被自动掐断。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeChannelStreamURLWithContext(ctx context.Context, request *DescribeChannelStreamURLRequest) (response *DescribeChannelStreamURLResponse, err error) {
    if request == nil {
        request = NewDescribeChannelStreamURLRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChannelStreamURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelsRequest() (request *DescribeChannelsRequest) {
    request = &DescribeChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeChannels")
    
    
    return
}

func NewDescribeChannelsResponse() (response *DescribeChannelsResponse) {
    response = &DescribeChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannels
// 本接口（DescribeChannels）用于获取设备下属通道列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeChannels(request *DescribeChannelsRequest) (response *DescribeChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeChannelsRequest()
    }
    
    response = NewDescribeChannelsResponse()
    err = c.Send(request, response)
    return
}

// DescribeChannels
// 本接口（DescribeChannels）用于获取设备下属通道列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeChannelsWithContext(ctx context.Context, request *DescribeChannelsRequest) (response *DescribeChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeChannelsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChannelsResponse()
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
func (c *Client) DescribeChannelsByLiveRecordPlanWithContext(ctx context.Context, request *DescribeChannelsByLiveRecordPlanRequest) (response *DescribeChannelsByLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewDescribeChannelsByLiveRecordPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChannelsByLiveRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentDeviceDataRequest() (request *DescribeCurrentDeviceDataRequest) {
    request = &DescribeCurrentDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeCurrentDeviceData")
    
    
    return
}

func NewDescribeCurrentDeviceDataResponse() (response *DescribeCurrentDeviceDataResponse) {
    response = &DescribeCurrentDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCurrentDeviceData
// 查询设备统计当前信息
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
func (c *Client) DescribeCurrentDeviceData(request *DescribeCurrentDeviceDataRequest) (response *DescribeCurrentDeviceDataResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentDeviceDataRequest()
    }
    
    response = NewDescribeCurrentDeviceDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeCurrentDeviceData
// 查询设备统计当前信息
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
func (c *Client) DescribeCurrentDeviceDataWithContext(ctx context.Context, request *DescribeCurrentDeviceDataRequest) (response *DescribeCurrentDeviceDataResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentDeviceDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCurrentDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDevice")
    
    
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevice
// 获取指定设备详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

// DescribeDevice
// 获取指定设备详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeDeviceWithContext(ctx context.Context, request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceEventRequest() (request *DescribeDeviceEventRequest) {
    request = &DescribeDeviceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDeviceEvent")
    
    
    return
}

func NewDescribeDeviceEventResponse() (response *DescribeDeviceEventResponse) {
    response = &DescribeDeviceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceEvent
// 获取设备事件
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
func (c *Client) DescribeDeviceEvent(request *DescribeDeviceEventRequest) (response *DescribeDeviceEventResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceEventRequest()
    }
    
    response = NewDescribeDeviceEventResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceEvent
// 获取设备事件
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
func (c *Client) DescribeDeviceEventWithContext(ctx context.Context, request *DescribeDeviceEventRequest) (response *DescribeDeviceEventResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceEventRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceEventResponse()
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
func (c *Client) DescribeDeviceGroup(request *DescribeDeviceGroupRequest) (response *DescribeDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupRequest()
    }
    
    response = NewDescribeDeviceGroupResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeDeviceGroupWithContext(ctx context.Context, request *DescribeDeviceGroupRequest) (response *DescribeDeviceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceListRequest() (request *DescribeDeviceListRequest) {
    request = &DescribeDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDeviceList")
    
    
    return
}

func NewDescribeDeviceListResponse() (response *DescribeDeviceListResponse) {
    response = &DescribeDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceList
// 本接口(DescribeDevicesList) 用于获取设备列表，支持模糊搜索
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
func (c *Client) DescribeDeviceList(request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceListRequest()
    }
    
    response = NewDescribeDeviceListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceList
// 本接口(DescribeDevicesList) 用于获取设备列表，支持模糊搜索
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
func (c *Client) DescribeDeviceListWithContext(ctx context.Context, request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceMonitorDataRequest() (request *DescribeDeviceMonitorDataRequest) {
    request = &DescribeDeviceMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeDeviceMonitorData")
    
    
    return
}

func NewDescribeDeviceMonitorDataResponse() (response *DescribeDeviceMonitorDataResponse) {
    response = &DescribeDeviceMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceMonitorData
// 查询设备统计monitor信息
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
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
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
func (c *Client) DescribeDeviceMonitorData(request *DescribeDeviceMonitorDataRequest) (response *DescribeDeviceMonitorDataResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceMonitorDataRequest()
    }
    
    response = NewDescribeDeviceMonitorDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceMonitorData
// 查询设备统计monitor信息
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
//  INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"
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
func (c *Client) DescribeDeviceMonitorDataWithContext(ctx context.Context, request *DescribeDeviceMonitorDataRequest) (response *DescribeDeviceMonitorDataResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceMonitorDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDeviceMonitorDataResponse()
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
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeDevicePassWord(request *DescribeDevicePassWordRequest) (response *DescribeDevicePassWordResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePassWordRequest()
    }
    
    response = NewDescribeDevicePassWordResponse()
    err = c.Send(request, response)
    return
}

// DescribeDevicePassWord
// 本接口(DescribeDevicePassWord)用于查询设备密码。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeDevicePassWordWithContext(ctx context.Context, request *DescribeDevicePassWordRequest) (response *DescribeDevicePassWordResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePassWordRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DescribeChannelStreamURL接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeDeviceStreams(request *DescribeDeviceStreamsRequest) (response *DescribeDeviceStreamsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStreamsRequest()
    }
    
    response = NewDescribeDeviceStreamsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDeviceStreams
// 本接口(DescribeDeviceStreams)用于获取设备实时流地址。
//
// 请使用DescribeChannelStreamURL接口
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
//  UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
func (c *Client) DescribeDeviceStreamsWithContext(ctx context.Context, request *DescribeDeviceStreamsRequest) (response *DescribeDeviceStreamsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStreamsRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupById(request *DescribeGroupByIdRequest) (response *DescribeGroupByIdResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByIdRequest()
    }
    
    response = NewDescribeGroupByIdResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupById
// 本接口(DescribeGroupById)用于根据分组ID查询分组。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupByIdWithContext(ctx context.Context, request *DescribeGroupByIdRequest) (response *DescribeGroupByIdResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByIdRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupByPath(request *DescribeGroupByPathRequest) (response *DescribeGroupByPathResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByPathRequest()
    }
    
    response = NewDescribeGroupByPathResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupByPath
// 根据分组路径查询分组
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupByPathWithContext(ctx context.Context, request *DescribeGroupByPathRequest) (response *DescribeGroupByPathResponse, err error) {
    if request == nil {
        request = NewDescribeGroupByPathRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupDevices(request *DescribeGroupDevicesRequest) (response *DescribeGroupDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupDevicesRequest()
    }
    
    response = NewDescribeGroupDevicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupDevices
// 本接口(DescribeGroupDevices)用于查询分组下的设备列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupDevicesWithContext(ctx context.Context, request *DescribeGroupDevicesRequest) (response *DescribeGroupDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupDevicesRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroups
// 本接口(DescribeGroups)用于批量查询分组信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DescribeChannels接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeIPCChannels(request *DescribeIPCChannelsRequest) (response *DescribeIPCChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeIPCChannelsRequest()
    }
    
    response = NewDescribeIPCChannelsResponse()
    err = c.Send(request, response)
    return
}

// DescribeIPCChannels
// 获取IPC设备下属通道
//
// 请使用DescribeChannels接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeIPCChannelsWithContext(ctx context.Context, request *DescribeIPCChannelsRequest) (response *DescribeIPCChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeIPCChannelsRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveChannelWithContext(ctx context.Context, request *DescribeLiveChannelRequest) (response *DescribeLiveChannelResponse, err error) {
    if request == nil {
        request = NewDescribeLiveChannelRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveChannelListWithContext(ctx context.Context, request *DescribeLiveChannelListRequest) (response *DescribeLiveChannelListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveChannelListRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveRecordPlanByIdWithContext(ctx context.Context, request *DescribeLiveRecordPlanByIdRequest) (response *DescribeLiveRecordPlanByIdResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordPlanByIdRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveRecordPlanIdsWithContext(ctx context.Context, request *DescribeLiveRecordPlanIdsRequest) (response *DescribeLiveRecordPlanIdsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordPlanIdsRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveStreamWithContext(ctx context.Context, request *DescribeLiveStreamRequest) (response *DescribeLiveStreamResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeLiveVideoListWithContext(ctx context.Context, request *DescribeLiveVideoListRequest) (response *DescribeLiveVideoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveVideoListRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeMessageForwardWithContext(ctx context.Context, request *DescribeMessageForwardRequest) (response *DescribeMessageForwardResponse, err error) {
    if request == nil {
        request = NewDescribeMessageForwardRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeMessageForwardsWithContext(ctx context.Context, request *DescribeMessageForwardsRequest) (response *DescribeMessageForwardsResponse, err error) {
    if request == nil {
        request = NewDescribeMessageForwardsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMessageForwardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorDataByDateRequest() (request *DescribeMonitorDataByDateRequest) {
    request = &DescribeMonitorDataByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeMonitorDataByDate")
    
    
    return
}

func NewDescribeMonitorDataByDateResponse() (response *DescribeMonitorDataByDateResponse) {
    response = &DescribeMonitorDataByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMonitorDataByDate
// 运营中心-设备录像存储统计
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
func (c *Client) DescribeMonitorDataByDate(request *DescribeMonitorDataByDateRequest) (response *DescribeMonitorDataByDateResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorDataByDateRequest()
    }
    
    response = NewDescribeMonitorDataByDateResponse()
    err = c.Send(request, response)
    return
}

// DescribeMonitorDataByDate
// 运营中心-设备录像存储统计
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
func (c *Client) DescribeMonitorDataByDateWithContext(ctx context.Context, request *DescribeMonitorDataByDateRequest) (response *DescribeMonitorDataByDateResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorDataByDateRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMonitorDataByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePresetListRequest() (request *DescribePresetListRequest) {
    request = &DescribePresetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribePresetList")
    
    
    return
}

func NewDescribePresetListResponse() (response *DescribePresetListResponse) {
    response = &DescribePresetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePresetList
// 获取预置位列表
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
func (c *Client) DescribePresetList(request *DescribePresetListRequest) (response *DescribePresetListResponse, err error) {
    if request == nil {
        request = NewDescribePresetListRequest()
    }
    
    response = NewDescribePresetListResponse()
    err = c.Send(request, response)
    return
}

// DescribePresetList
// 获取预置位列表
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
func (c *Client) DescribePresetListWithContext(ctx context.Context, request *DescribePresetListRequest) (response *DescribePresetListResponse, err error) {
    if request == nil {
        request = NewDescribePresetListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePresetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordDatesByChannelRequest() (request *DescribeRecordDatesByChannelRequest) {
    request = &DescribeRecordDatesByChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeRecordDatesByChannel")
    
    
    return
}

func NewDescribeRecordDatesByChannelResponse() (response *DescribeRecordDatesByChannelResponse) {
    response = &DescribeRecordDatesByChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordDatesByChannel
// 本接口(DescribeRecordDatesByChannel)用于查询设备含有录像文件的日期列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordDatesByChannel(request *DescribeRecordDatesByChannelRequest) (response *DescribeRecordDatesByChannelResponse, err error) {
    if request == nil {
        request = NewDescribeRecordDatesByChannelRequest()
    }
    
    response = NewDescribeRecordDatesByChannelResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordDatesByChannel
// 本接口(DescribeRecordDatesByChannel)用于查询设备含有录像文件的日期列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordDatesByChannelWithContext(ctx context.Context, request *DescribeRecordDatesByChannelRequest) (response *DescribeRecordDatesByChannelResponse, err error) {
    if request == nil {
        request = NewDescribeRecordDatesByChannelRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordDatesByChannelResponse()
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
func (c *Client) DescribeRecordDatesByLiveWithContext(ctx context.Context, request *DescribeRecordDatesByLiveRequest) (response *DescribeRecordDatesByLiveResponse, err error) {
    if request == nil {
        request = NewDescribeRecordDatesByLiveRequest()
    }
    request.SetContext(ctx)
    
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
// 获取回放视频流地址
//
// 请使用DescribeChannelLocalRecordURL接口
//
// 
//
// RecordId和StartTime/EndTime互斥
//
// 当存在RecordId时，StartTime和EndTime无效
//
// 当RecordId为空，StartTime和EndTime生效
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordStream(request *DescribeRecordStreamRequest) (response *DescribeRecordStreamResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStreamRequest()
    }
    
    response = NewDescribeRecordStreamResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordStream
// 获取回放视频流地址
//
// 请使用DescribeChannelLocalRecordURL接口
//
// 
//
// RecordId和StartTime/EndTime互斥
//
// 当存在RecordId时，StartTime和EndTime无效
//
// 当RecordId为空，StartTime和EndTime生效
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordStreamWithContext(ctx context.Context, request *DescribeRecordStreamRequest) (response *DescribeRecordStreamResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStreamRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingPlanByIdRequest() (request *DescribeRecordingPlanByIdRequest) {
    request = &DescribeRecordingPlanByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeRecordingPlanById")
    
    
    return
}

func NewDescribeRecordingPlanByIdResponse() (response *DescribeRecordingPlanByIdResponse) {
    response = &DescribeRecordingPlanByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordingPlanById
// 本接口(DescribeRecordingPlanById)用于根据录制计划ID获取录制计划。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordingPlanById(request *DescribeRecordingPlanByIdRequest) (response *DescribeRecordingPlanByIdResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingPlanByIdRequest()
    }
    
    response = NewDescribeRecordingPlanByIdResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordingPlanById
// 本接口(DescribeRecordingPlanById)用于根据录制计划ID获取录制计划。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordingPlanByIdWithContext(ctx context.Context, request *DescribeRecordingPlanByIdRequest) (response *DescribeRecordingPlanByIdResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingPlanByIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordingPlanByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingPlansRequest() (request *DescribeRecordingPlansRequest) {
    request = &DescribeRecordingPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeRecordingPlans")
    
    
    return
}

func NewDescribeRecordingPlansResponse() (response *DescribeRecordingPlansResponse) {
    response = &DescribeRecordingPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordingPlans
// 本接口(DescribeRecordingPlans)用于获取用户的全部录制计划。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordingPlans(request *DescribeRecordingPlansRequest) (response *DescribeRecordingPlansResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingPlansRequest()
    }
    
    response = NewDescribeRecordingPlansResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordingPlans
// 本接口(DescribeRecordingPlans)用于获取用户的全部录制计划。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeRecordingPlansWithContext(ctx context.Context, request *DescribeRecordingPlansRequest) (response *DescribeRecordingPlansResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingPlansRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordingPlansResponse()
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
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeSIPServer(request *DescribeSIPServerRequest) (response *DescribeSIPServerResponse, err error) {
    if request == nil {
        request = NewDescribeSIPServerRequest()
    }
    
    response = NewDescribeSIPServerResponse()
    err = c.Send(request, response)
    return
}

// DescribeSIPServer
// 本接口用于获取SIP服务器相关配置，用户可以通过这些配置项，将设备通过GB28181协议注册到本服务。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"
//  INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeSIPServerWithContext(ctx context.Context, request *DescribeSIPServerRequest) (response *DescribeSIPServerResponse, err error) {
    if request == nil {
        request = NewDescribeSIPServerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSIPServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSceneRequest() (request *DescribeSceneRequest) {
    request = &DescribeSceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeScene")
    
    
    return
}

func NewDescribeSceneResponse() (response *DescribeSceneResponse) {
    response = &DescribeSceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScene
// 场景详情
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
func (c *Client) DescribeScene(request *DescribeSceneRequest) (response *DescribeSceneResponse, err error) {
    if request == nil {
        request = NewDescribeSceneRequest()
    }
    
    response = NewDescribeSceneResponse()
    err = c.Send(request, response)
    return
}

// DescribeScene
// 场景详情
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
func (c *Client) DescribeSceneWithContext(ctx context.Context, request *DescribeSceneRequest) (response *DescribeSceneResponse, err error) {
    if request == nil {
        request = NewDescribeSceneRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSceneResponse()
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
func (c *Client) DescribeScenesWithContext(ctx context.Context, request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    if request == nil {
        request = NewDescribeScenesRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeStatisticDetails(request *DescribeStatisticDetailsRequest) (response *DescribeStatisticDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDetailsRequest()
    }
    
    response = NewDescribeStatisticDetailsResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeStatisticDetailsWithContext(ctx context.Context, request *DescribeStatisticDetailsRequest) (response *DescribeStatisticDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDetailsRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeStatisticSummary(request *DescribeStatisticSummaryRequest) (response *DescribeStatisticSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticSummaryRequest()
    }
    
    response = NewDescribeStatisticSummaryResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeStatisticSummaryWithContext(ctx context.Context, request *DescribeStatisticSummaryRequest) (response *DescribeStatisticSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticSummaryRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeSubGroups(request *DescribeSubGroupsRequest) (response *DescribeSubGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSubGroupsRequest()
    }
    
    response = NewDescribeSubGroupsResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeSubGroupsWithContext(ctx context.Context, request *DescribeSubGroupsRequest) (response *DescribeSubGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSubGroupsRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) DescribeSubscriptionStatusWithContext(ctx context.Context, request *DescribeSubscriptionStatusRequest) (response *DescribeSubscriptionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionStatusRequest()
    }
    request.SetContext(ctx)
    
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
// 根据时间获取云端录制文件列表
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) DescribeVideoList(request *DescribeVideoListRequest) (response *DescribeVideoListResponse, err error) {
    if request == nil {
        request = NewDescribeVideoListRequest()
    }
    
    response = NewDescribeVideoListResponse()
    err = c.Send(request, response)
    return
}

// DescribeVideoList
// 根据时间获取云端录制文件列表
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) DescribeVideoListWithContext(ctx context.Context, request *DescribeVideoListRequest) (response *DescribeVideoListResponse, err error) {
    if request == nil {
        request = NewDescribeVideoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeVideoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoListByChannelRequest() (request *DescribeVideoListByChannelRequest) {
    request = &DescribeVideoListByChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeVideoListByChannel")
    
    
    return
}

func NewDescribeVideoListByChannelResponse() (response *DescribeVideoListByChannelResponse) {
    response = &DescribeVideoListByChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoListByChannel
// 本接口(DescribeVideoListByChannel)用于查询指定通道的录制文件列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeVideoListByChannel(request *DescribeVideoListByChannelRequest) (response *DescribeVideoListByChannelResponse, err error) {
    if request == nil {
        request = NewDescribeVideoListByChannelRequest()
    }
    
    response = NewDescribeVideoListByChannelResponse()
    err = c.Send(request, response)
    return
}

// DescribeVideoListByChannel
// 本接口(DescribeVideoListByChannel)用于查询指定通道的录制文件列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeVideoListByChannelWithContext(ctx context.Context, request *DescribeVideoListByChannelRequest) (response *DescribeVideoListByChannelResponse, err error) {
    if request == nil {
        request = NewDescribeVideoListByChannelRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeVideoListByChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarnModRequest() (request *DescribeWarnModRequest) {
    request = &DescribeWarnModRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeWarnMod")
    
    
    return
}

func NewDescribeWarnModResponse() (response *DescribeWarnModResponse) {
    response = &DescribeWarnModResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWarnMod
// 告警等级列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeWarnMod(request *DescribeWarnModRequest) (response *DescribeWarnModResponse, err error) {
    if request == nil {
        request = NewDescribeWarnModRequest()
    }
    
    response = NewDescribeWarnModResponse()
    err = c.Send(request, response)
    return
}

// DescribeWarnMod
// 告警等级列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeWarnModWithContext(ctx context.Context, request *DescribeWarnModRequest) (response *DescribeWarnModResponse, err error) {
    if request == nil {
        request = NewDescribeWarnModRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeWarnModResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWarningsRequest() (request *DescribeWarningsRequest) {
    request = &DescribeWarningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeWarnings")
    
    
    return
}

func NewDescribeWarningsResponse() (response *DescribeWarningsResponse) {
    response = &DescribeWarningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWarnings
// 获取告警列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeWarnings(request *DescribeWarningsRequest) (response *DescribeWarningsResponse, err error) {
    if request == nil {
        request = NewDescribeWarningsRequest()
    }
    
    response = NewDescribeWarningsResponse()
    err = c.Send(request, response)
    return
}

// DescribeWarnings
// 获取告警列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeWarningsWithContext(ctx context.Context, request *DescribeWarningsRequest) (response *DescribeWarningsResponse, err error) {
    if request == nil {
        request = NewDescribeWarningsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeWarningsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeXP2PDataRequest() (request *DescribeXP2PDataRequest) {
    request = &DescribeXP2PDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "DescribeXP2PData")
    
    
    return
}

func NewDescribeXP2PDataResponse() (response *DescribeXP2PDataResponse) {
    response = &DescribeXP2PDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeXP2PData
// 获取X-P2P的统计数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeXP2PData(request *DescribeXP2PDataRequest) (response *DescribeXP2PDataResponse, err error) {
    if request == nil {
        request = NewDescribeXP2PDataRequest()
    }
    
    response = NewDescribeXP2PDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeXP2PData
// 获取X-P2P的统计数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) DescribeXP2PDataWithContext(ctx context.Context, request *DescribeXP2PDataRequest) (response *DescribeXP2PDataResponse, err error) {
    if request == nil {
        request = NewDescribeXP2PDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeXP2PDataResponse()
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
// 请使用DescribeRecordDatesByChannel接口
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) GetRecordDatesByDev(request *GetRecordDatesByDevRequest) (response *GetRecordDatesByDevResponse, err error) {
    if request == nil {
        request = NewGetRecordDatesByDevRequest()
    }
    
    response = NewGetRecordDatesByDevResponse()
    err = c.Send(request, response)
    return
}

// GetRecordDatesByDev
// 本接口(GetRecordDatesByDev)用于查询设备含有录像文件的日期列表。
//
// 请使用DescribeRecordDatesByChannel接口
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
func (c *Client) GetRecordDatesByDevWithContext(ctx context.Context, request *GetRecordDatesByDevRequest) (response *GetRecordDatesByDevResponse, err error) {
    if request == nil {
        request = NewGetRecordDatesByDevRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) GetRecordPlanByDevWithContext(ctx context.Context, request *GetRecordPlanByDevRequest) (response *GetRecordPlanByDevResponse, err error) {
    if request == nil {
        request = NewGetRecordPlanByDevRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DescribeRecordingPlanById接口
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
func (c *Client) GetRecordPlanById(request *GetRecordPlanByIdRequest) (response *GetRecordPlanByIdResponse, err error) {
    if request == nil {
        request = NewGetRecordPlanByIdRequest()
    }
    
    response = NewGetRecordPlanByIdResponse()
    err = c.Send(request, response)
    return
}

// GetRecordPlanById
// 本接口(GetRecordPlanById)用于根据录制计划ID获取录制计划。
//
// 请使用DescribeRecordingPlanById接口
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
func (c *Client) GetRecordPlanByIdWithContext(ctx context.Context, request *GetRecordPlanByIdRequest) (response *GetRecordPlanByIdResponse, err error) {
    if request == nil {
        request = NewGetRecordPlanByIdRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DescribeRecordingPlans接口
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
func (c *Client) GetRecordPlans(request *GetRecordPlansRequest) (response *GetRecordPlansResponse, err error) {
    if request == nil {
        request = NewGetRecordPlansRequest()
    }
    
    response = NewGetRecordPlansResponse()
    err = c.Send(request, response)
    return
}

// GetRecordPlans
// 本接口(GetRecordPlans)用于获取用户的全部录制计划。
//
// 请使用DescribeRecordingPlans接口
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
func (c *Client) GetRecordPlansWithContext(ctx context.Context, request *GetRecordPlansRequest) (response *GetRecordPlansResponse, err error) {
    if request == nil {
        request = NewGetRecordPlansRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) GetTimeTemplateById(request *GetTimeTemplateByIdRequest) (response *GetTimeTemplateByIdResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplateByIdRequest()
    }
    
    response = NewGetTimeTemplateByIdResponse()
    err = c.Send(request, response)
    return
}

// GetTimeTemplateById
// 本接口(GetTimeTemplateById)用于根据模板ID获取时间模板详情。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) GetTimeTemplateByIdWithContext(ctx context.Context, request *GetTimeTemplateByIdRequest) (response *GetTimeTemplateByIdResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplateByIdRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) GetTimeTemplates(request *GetTimeTemplatesRequest) (response *GetTimeTemplatesResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplatesRequest()
    }
    
    response = NewGetTimeTemplatesResponse()
    err = c.Send(request, response)
    return
}

// GetTimeTemplates
// 本接口(GetTimeTemplates)用于获取时间模板列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) GetTimeTemplatesWithContext(ctx context.Context, request *GetTimeTemplatesRequest) (response *GetTimeTemplatesResponse, err error) {
    if request == nil {
        request = NewGetTimeTemplatesRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用DescribeVideoListByChannel接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) GetVideoListByCon(request *GetVideoListByConRequest) (response *GetVideoListByConResponse, err error) {
    if request == nil {
        request = NewGetVideoListByConRequest()
    }
    
    response = NewGetVideoListByConResponse()
    err = c.Send(request, response)
    return
}

// GetVideoListByCon
// 本接口(GetVideoListByCon)用于查询设备的录制文件列表
//
// 请使用DescribeVideoListByChannel接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) GetVideoListByConWithContext(ctx context.Context, request *GetVideoListByConRequest) (response *GetVideoListByConResponse, err error) {
    if request == nil {
        request = NewGetVideoListByConRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) ModifyBindPlanLiveChannel(request *ModifyBindPlanLiveChannelRequest) (response *ModifyBindPlanLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyBindPlanLiveChannelRequest()
    }
    
    response = NewModifyBindPlanLiveChannelResponse()
    err = c.Send(request, response)
    return
}

// ModifyBindPlanLiveChannel
// 直播录制计划绑定解绑直播频道
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"
//  UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"
func (c *Client) ModifyBindPlanLiveChannelWithContext(ctx context.Context, request *ModifyBindPlanLiveChannelRequest) (response *ModifyBindPlanLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyBindPlanLiveChannelRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyBindPlanLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBindRecordingPlanRequest() (request *ModifyBindRecordingPlanRequest) {
    request = &ModifyBindRecordingPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyBindRecordingPlan")
    
    
    return
}

func NewModifyBindRecordingPlanResponse() (response *ModifyBindRecordingPlanResponse) {
    response = &ModifyBindRecordingPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBindRecordingPlan
// 本接口(ModifyBindRecordingPlan)用于更新录制计划绑定的通道
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
func (c *Client) ModifyBindRecordingPlan(request *ModifyBindRecordingPlanRequest) (response *ModifyBindRecordingPlanResponse, err error) {
    if request == nil {
        request = NewModifyBindRecordingPlanRequest()
    }
    
    response = NewModifyBindRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

// ModifyBindRecordingPlan
// 本接口(ModifyBindRecordingPlan)用于更新录制计划绑定的通道
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
func (c *Client) ModifyBindRecordingPlanWithContext(ctx context.Context, request *ModifyBindRecordingPlanRequest) (response *ModifyBindRecordingPlanResponse, err error) {
    if request == nil {
        request = NewModifyBindRecordingPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyBindRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBindSceneChannelsRequest() (request *ModifyBindSceneChannelsRequest) {
    request = &ModifyBindSceneChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyBindSceneChannels")
    
    
    return
}

func NewModifyBindSceneChannelsResponse() (response *ModifyBindSceneChannelsResponse) {
    response = &ModifyBindSceneChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBindSceneChannels
// 场景绑定解绑通道接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
func (c *Client) ModifyBindSceneChannels(request *ModifyBindSceneChannelsRequest) (response *ModifyBindSceneChannelsResponse, err error) {
    if request == nil {
        request = NewModifyBindSceneChannelsRequest()
    }
    
    response = NewModifyBindSceneChannelsResponse()
    err = c.Send(request, response)
    return
}

// ModifyBindSceneChannels
// 场景绑定解绑通道接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"
func (c *Client) ModifyBindSceneChannelsWithContext(ctx context.Context, request *ModifyBindSceneChannelsRequest) (response *ModifyBindSceneChannelsResponse, err error) {
    if request == nil {
        request = NewModifyBindSceneChannelsRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyBindSceneChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBindSceneDeviceRequest() (request *ModifyBindSceneDeviceRequest) {
    request = &ModifyBindSceneDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyBindSceneDevice")
    
    
    return
}

func NewModifyBindSceneDeviceResponse() (response *ModifyBindSceneDeviceResponse) {
    response = &ModifyBindSceneDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBindSceneDevice
// 场景绑定/解绑通道接口
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
func (c *Client) ModifyBindSceneDevice(request *ModifyBindSceneDeviceRequest) (response *ModifyBindSceneDeviceResponse, err error) {
    if request == nil {
        request = NewModifyBindSceneDeviceRequest()
    }
    
    response = NewModifyBindSceneDeviceResponse()
    err = c.Send(request, response)
    return
}

// ModifyBindSceneDevice
// 场景绑定/解绑通道接口
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
func (c *Client) ModifyBindSceneDeviceWithContext(ctx context.Context, request *ModifyBindSceneDeviceRequest) (response *ModifyBindSceneDeviceResponse, err error) {
    if request == nil {
        request = NewModifyBindSceneDeviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyBindSceneDeviceResponse()
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
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
func (c *Client) ModifyDeviceData(request *ModifyDeviceDataRequest) (response *ModifyDeviceDataResponse, err error) {
    if request == nil {
        request = NewModifyDeviceDataRequest()
    }
    
    response = NewModifyDeviceDataResponse()
    err = c.Send(request, response)
    return
}

// ModifyDeviceData
// 本接口(ModifyDeviceData)用于编辑设备信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"
func (c *Client) ModifyDeviceDataWithContext(ctx context.Context, request *ModifyDeviceDataRequest) (response *ModifyDeviceDataResponse, err error) {
    if request == nil {
        request = NewModifyDeviceDataRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) ModifyLiveChannelWithContext(ctx context.Context, request *ModifyLiveChannelRequest) (response *ModifyLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyLiveChannelRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) ModifyLiveRecordPlanWithContext(ctx context.Context, request *ModifyLiveRecordPlanRequest) (response *ModifyLiveRecordPlanResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordPlanRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) ModifyLiveVideoWithContext(ctx context.Context, request *ModifyLiveVideoRequest) (response *ModifyLiveVideoResponse, err error) {
    if request == nil {
        request = NewModifyLiveVideoRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) ModifyMessageForwardWithContext(ctx context.Context, request *ModifyMessageForwardRequest) (response *ModifyMessageForwardResponse, err error) {
    if request == nil {
        request = NewModifyMessageForwardRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyMessageForwardResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPresetRequest() (request *ModifyPresetRequest) {
    request = &ModifyPresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyPreset")
    
    
    return
}

func NewModifyPresetResponse() (response *ModifyPresetResponse) {
    response = &ModifyPresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPreset
// 编辑预置位信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPreset(request *ModifyPresetRequest) (response *ModifyPresetResponse, err error) {
    if request == nil {
        request = NewModifyPresetRequest()
    }
    
    response = NewModifyPresetResponse()
    err = c.Send(request, response)
    return
}

// ModifyPreset
// 编辑预置位信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPresetWithContext(ctx context.Context, request *ModifyPresetRequest) (response *ModifyPresetResponse, err error) {
    if request == nil {
        request = NewModifyPresetRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPresetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordingPlanRequest() (request *ModifyRecordingPlanRequest) {
    request = &ModifyRecordingPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyRecordingPlan")
    
    
    return
}

func NewModifyRecordingPlanResponse() (response *ModifyRecordingPlanResponse) {
    response = &ModifyRecordingPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecordingPlan
// 本接口(ModifyRecordingPlan)用于更新录制计划。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRecordingPlan(request *ModifyRecordingPlanRequest) (response *ModifyRecordingPlanResponse, err error) {
    if request == nil {
        request = NewModifyRecordingPlanRequest()
    }
    
    response = NewModifyRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

// ModifyRecordingPlan
// 本接口(ModifyRecordingPlan)用于更新录制计划。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRecordingPlanWithContext(ctx context.Context, request *ModifyRecordingPlanRequest) (response *ModifyRecordingPlanResponse, err error) {
    if request == nil {
        request = NewModifyRecordingPlanRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRecordingPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifySceneRequest() (request *ModifySceneRequest) {
    request = &ModifySceneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ModifyScene")
    
    
    return
}

func NewModifySceneResponse() (response *ModifySceneResponse) {
    response = &ModifySceneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyScene
// 修改场景
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
func (c *Client) ModifyScene(request *ModifySceneRequest) (response *ModifySceneResponse, err error) {
    if request == nil {
        request = NewModifySceneRequest()
    }
    
    response = NewModifySceneResponse()
    err = c.Send(request, response)
    return
}

// ModifyScene
// 修改场景
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
func (c *Client) ModifySceneWithContext(ctx context.Context, request *ModifySceneRequest) (response *ModifySceneResponse, err error) {
    if request == nil {
        request = NewModifySceneRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifySceneResponse()
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
func (c *Client) ModifySubscriptionStatusWithContext(ctx context.Context, request *ModifySubscriptionStatusRequest) (response *ModifySubscriptionStatusResponse, err error) {
    if request == nil {
        request = NewModifySubscriptionStatusRequest()
    }
    request.SetContext(ctx)
    
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
func (c *Client) ModifyVideoInfoWithContext(ctx context.Context, request *ModifyVideoInfoRequest) (response *ModifyVideoInfoResponse, err error) {
    if request == nil {
        request = NewModifyVideoInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyVideoInfoResponse()
    err = c.Send(request, response)
    return
}

func NewResetWarningRequest() (request *ResetWarningRequest) {
    request = &ResetWarningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideoindustry", APIVersion, "ResetWarning")
    
    
    return
}

func NewResetWarningResponse() (response *ResetWarningResponse) {
    response = &ResetWarningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetWarning
// 重置设备告警
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
func (c *Client) ResetWarning(request *ResetWarningRequest) (response *ResetWarningResponse, err error) {
    if request == nil {
        request = NewResetWarningRequest()
    }
    
    response = NewResetWarningResponse()
    err = c.Send(request, response)
    return
}

// ResetWarning
// 重置设备告警
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
func (c *Client) ResetWarningWithContext(ctx context.Context, request *ResetWarningRequest) (response *ResetWarningResponse, err error) {
    if request == nil {
        request = NewResetWarningRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetWarningResponse()
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
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
func (c *Client) UpdateDeviceGroup(request *UpdateDeviceGroupRequest) (response *UpdateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceGroupRequest()
    }
    
    response = NewUpdateDeviceGroupResponse()
    err = c.Send(request, response)
    return
}

// UpdateDeviceGroup
// 本接口(UpdateDeviceGroup)用于修改分组信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"
func (c *Client) UpdateDeviceGroupWithContext(ctx context.Context, request *UpdateDeviceGroupRequest) (response *UpdateDeviceGroupResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceGroupRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UpdateDevicePassWord(request *UpdateDevicePassWordRequest) (response *UpdateDevicePassWordResponse, err error) {
    if request == nil {
        request = NewUpdateDevicePassWordRequest()
    }
    
    response = NewUpdateDevicePassWordResponse()
    err = c.Send(request, response)
    return
}

// UpdateDevicePassWord
// 本接口(UpdateDevicePassWord)用于修改设备密码。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) UpdateDevicePassWordWithContext(ctx context.Context, request *UpdateDevicePassWordRequest) (response *UpdateDevicePassWordResponse, err error) {
    if request == nil {
        request = NewUpdateDevicePassWordRequest()
    }
    request.SetContext(ctx)
    
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
// 请使用 ModifyRecordingPlan接口和ModifyBindRecordingPlan接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordPlanRequest()
    }
    
    response = NewUpdateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

// UpdateRecordPlan
// 本接口(UpdateRecordPlan)用于更新录制计划。
//
// 请使用 ModifyRecordingPlan接口和ModifyBindRecordingPlan接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) UpdateRecordPlanWithContext(ctx context.Context, request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordPlanRequest()
    }
    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) UpdateTimeTemplate(request *UpdateTimeTemplateRequest) (response *UpdateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateTimeTemplateRequest()
    }
    
    response = NewUpdateTimeTemplateResponse()
    err = c.Send(request, response)
    return
}

// UpdateTimeTemplate
// 本接口(UpdateTimeTemplate)用于更新时间模板。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
//  UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"
//  UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"
func (c *Client) UpdateTimeTemplateWithContext(ctx context.Context, request *UpdateTimeTemplateRequest) (response *UpdateTimeTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateTimeTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateTimeTemplateResponse()
    err = c.Send(request, response)
    return
}
