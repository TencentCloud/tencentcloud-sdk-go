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

package v20230517

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-05-17"

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


func NewAddAITaskRequest() (request *AddAITaskRequest) {
    request = &AddAITaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddAITask")
    
    
    return
}

func NewAddAITaskResponse() (response *AddAITaskResponse) {
    response = &AddAITaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddAITask
// 添加AI任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CALLBACKURLCONTAINILLEGALCHARACTER = "InvalidParameterValue.CallbackURLContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELIDALREADYEXISTSINOTHERAITASKS = "InvalidParameterValue.ChannelIdAlreadyExistsInOtherAITasks"
//  INVALIDPARAMETERVALUE_CHANNELLISTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ChannelListContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELLISTMUSTBENOTEMPTY = "InvalidParameterValue.ChannelListMustBeNotEmpty"
//  INVALIDPARAMETERVALUE_CHANNELNUMBERMUSTBELESSTHANONETHOUSAND = "InvalidParameterValue.ChannelNumberMustBeLessThanOneThousand"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_HASDUPLICATEOPERTIMESLOT = "InvalidParameterValue.HasDuplicateOperTimeSlot"
//  INVALIDPARAMETERVALUE_INVALIDAITASKDESC = "InvalidParameterValue.InvalidAITaskDesc"
//  INVALIDPARAMETERVALUE_INVALIDAITASKNAME = "InvalidParameterValue.InvalidAITaskName"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDOPERTIMESLOTFORMAT = "InvalidParameterValue.InvalidOperTimeSlotFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATETAG = "InvalidParameterValue.InvalidTemplateTag"
//  INVALIDPARAMETERVALUE_INVALIDTIMEINTERVAL = "InvalidParameterValue.InvalidTimeInterval"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTCONTAINILLEGALCHARACTER = "InvalidParameterValue.OperTimeSlotContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTNUMBERMUSTBELESSTHANFIVE = "InvalidParameterValue.OperTimeSlotNumberMustBeLessThanFive"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTSTARTMUSTBELESSTHANEND = "InvalidParameterValue.OperTimeSlotStartMustBeLessThanEnd"
//  INVALIDPARAMETERVALUE_TEMPLATETAGMUSTBECONSISTENT = "InvalidParameterValue.TemplateTagMustBeConsistent"
func (c *Client) AddAITask(request *AddAITaskRequest) (response *AddAITaskResponse, err error) {
    return c.AddAITaskWithContext(context.Background(), request)
}

// AddAITask
// 添加AI任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CALLBACKURLCONTAINILLEGALCHARACTER = "InvalidParameterValue.CallbackURLContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELIDALREADYEXISTSINOTHERAITASKS = "InvalidParameterValue.ChannelIdAlreadyExistsInOtherAITasks"
//  INVALIDPARAMETERVALUE_CHANNELLISTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ChannelListContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELLISTMUSTBENOTEMPTY = "InvalidParameterValue.ChannelListMustBeNotEmpty"
//  INVALIDPARAMETERVALUE_CHANNELNUMBERMUSTBELESSTHANONETHOUSAND = "InvalidParameterValue.ChannelNumberMustBeLessThanOneThousand"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_HASDUPLICATEOPERTIMESLOT = "InvalidParameterValue.HasDuplicateOperTimeSlot"
//  INVALIDPARAMETERVALUE_INVALIDAITASKDESC = "InvalidParameterValue.InvalidAITaskDesc"
//  INVALIDPARAMETERVALUE_INVALIDAITASKNAME = "InvalidParameterValue.InvalidAITaskName"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDOPERTIMESLOTFORMAT = "InvalidParameterValue.InvalidOperTimeSlotFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATETAG = "InvalidParameterValue.InvalidTemplateTag"
//  INVALIDPARAMETERVALUE_INVALIDTIMEINTERVAL = "InvalidParameterValue.InvalidTimeInterval"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTCONTAINILLEGALCHARACTER = "InvalidParameterValue.OperTimeSlotContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTNUMBERMUSTBELESSTHANFIVE = "InvalidParameterValue.OperTimeSlotNumberMustBeLessThanFive"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTSTARTMUSTBELESSTHANEND = "InvalidParameterValue.OperTimeSlotStartMustBeLessThanEnd"
//  INVALIDPARAMETERVALUE_TEMPLATETAGMUSTBECONSISTENT = "InvalidParameterValue.TemplateTagMustBeConsistent"
func (c *Client) AddAITaskWithContext(ctx context.Context, request *AddAITaskRequest) (response *AddAITaskResponse, err error) {
    if request == nil {
        request = NewAddAITaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAITask require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAITaskResponse()
    err = c.Send(request, response)
    return
}

func NewAddOrganizationRequest() (request *AddOrganizationRequest) {
    request = &AddOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddOrganization")
    
    
    return
}

func NewAddOrganizationResponse() (response *AddOrganizationResponse) {
    response = &AddOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddOrganization
// 用于新增组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddOrganization(request *AddOrganizationRequest) (response *AddOrganizationResponse, err error) {
    return c.AddOrganizationWithContext(context.Background(), request)
}

// AddOrganization
// 用于新增组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddOrganizationWithContext(ctx context.Context, request *AddOrganizationRequest) (response *AddOrganizationResponse, err error) {
    if request == nil {
        request = NewAddOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordBackupPlanRequest() (request *AddRecordBackupPlanRequest) {
    request = &AddRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordBackupPlan")
    
    
    return
}

func NewAddRecordBackupPlanResponse() (response *AddRecordBackupPlanResponse) {
    response = &AddRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRecordBackupPlan
// 用于新增录像上云计划 （当前仅适用于通过GB28181协议和网关接入的设备/视频通道）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddRecordBackupPlan(request *AddRecordBackupPlanRequest) (response *AddRecordBackupPlanResponse, err error) {
    return c.AddRecordBackupPlanWithContext(context.Background(), request)
}

// AddRecordBackupPlan
// 用于新增录像上云计划 （当前仅适用于通过GB28181协议和网关接入的设备/视频通道）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddRecordBackupPlanWithContext(ctx context.Context, request *AddRecordBackupPlanRequest) (response *AddRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewAddRecordBackupPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordBackupTemplateRequest() (request *AddRecordBackupTemplateRequest) {
    request = &AddRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordBackupTemplate")
    
    
    return
}

func NewAddRecordBackupTemplateResponse() (response *AddRecordBackupTemplateResponse) {
    response = &AddRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRecordBackupTemplate
// 用于新增录像上云模板。
//
// > 该功能本质是拉取设备本地录像数据上云（即存在 IPC 摄像头存储卡或 NVR 硬盘中的录像），操作时需先设定录像时间段（即想要上云的设备本地录像），再设定上云时间段和上云倍速，平台将于上云时间段倍速拉取设备对应前一天的录像时间段数据。
//
// 
//
// > 设定需至少满足（上云时间段=前一天的录像时间段/上云倍速），建议上云时间段可多设定10%左右的时间，避免因网络波动导致数据拉取不完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordBackupTemplate(request *AddRecordBackupTemplateRequest) (response *AddRecordBackupTemplateResponse, err error) {
    return c.AddRecordBackupTemplateWithContext(context.Background(), request)
}

// AddRecordBackupTemplate
// 用于新增录像上云模板。
//
// > 该功能本质是拉取设备本地录像数据上云（即存在 IPC 摄像头存储卡或 NVR 硬盘中的录像），操作时需先设定录像时间段（即想要上云的设备本地录像），再设定上云时间段和上云倍速，平台将于上云时间段倍速拉取设备对应前一天的录像时间段数据。
//
// 
//
// > 设定需至少满足（上云时间段=前一天的录像时间段/上云倍速），建议上云时间段可多设定10%左右的时间，避免因网络波动导致数据拉取不完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordBackupTemplateWithContext(ctx context.Context, request *AddRecordBackupTemplateRequest) (response *AddRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewAddRecordBackupTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordPlanRequest() (request *AddRecordPlanRequest) {
    request = &AddRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordPlan")
    
    
    return
}

func NewAddRecordPlanResponse() (response *AddRecordPlanResponse) {
    response = &AddRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRecordPlan
// 用于新增实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_EMPTYTEMPLATEID = "InvalidParameterValue.EmptyTemplateId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) AddRecordPlan(request *AddRecordPlanRequest) (response *AddRecordPlanResponse, err error) {
    return c.AddRecordPlanWithContext(context.Background(), request)
}

// AddRecordPlan
// 用于新增实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_EMPTYTEMPLATEID = "InvalidParameterValue.EmptyTemplateId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) AddRecordPlanWithContext(ctx context.Context, request *AddRecordPlanRequest) (response *AddRecordPlanResponse, err error) {
    if request == nil {
        request = NewAddRecordPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordRetrieveTaskRequest() (request *AddRecordRetrieveTaskRequest) {
    request = &AddRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordRetrieveTask")
    
    
    return
}

func NewAddRecordRetrieveTaskResponse() (response *AddRecordRetrieveTaskResponse) {
    response = &AddRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRecordRetrieveTask
// 用于新建取回任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATION = "InvalidParameterValue.InvalidExpiration"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVALMODE = "InvalidParameterValue.InvalidRetrievalMode"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_RETRIEVETASKCHANNELSEXCEEDSRANGE = "InvalidParameterValue.RetrieveTaskChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_MISSINGRETRIEVETASKPARAM = "MissingParameter.MissingRetrieveTaskParam"
//  RESOURCEINUSE_RETRIEVETASKNAMEREPEAT = "ResourceInUse.RetrieveTaskNameRepeat"
func (c *Client) AddRecordRetrieveTask(request *AddRecordRetrieveTaskRequest) (response *AddRecordRetrieveTaskResponse, err error) {
    return c.AddRecordRetrieveTaskWithContext(context.Background(), request)
}

// AddRecordRetrieveTask
// 用于新建取回任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATION = "InvalidParameterValue.InvalidExpiration"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVALMODE = "InvalidParameterValue.InvalidRetrievalMode"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_RETRIEVETASKCHANNELSEXCEEDSRANGE = "InvalidParameterValue.RetrieveTaskChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_MISSINGRETRIEVETASKPARAM = "MissingParameter.MissingRetrieveTaskParam"
//  RESOURCEINUSE_RETRIEVETASKNAMEREPEAT = "ResourceInUse.RetrieveTaskNameRepeat"
func (c *Client) AddRecordRetrieveTaskWithContext(ctx context.Context, request *AddRecordRetrieveTaskRequest) (response *AddRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewAddRecordRetrieveTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordTemplateRequest() (request *AddRecordTemplateRequest) {
    request = &AddRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordTemplate")
    
    
    return
}

func NewAddRecordTemplateResponse() (response *AddRecordTemplateResponse) {
    response = &AddRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRecordTemplate
// 用于新增实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordTemplate(request *AddRecordTemplateRequest) (response *AddRecordTemplateResponse, err error) {
    return c.AddRecordTemplateWithContext(context.Background(), request)
}

// AddRecordTemplate
// 用于新增实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordTemplateWithContext(ctx context.Context, request *AddRecordTemplateRequest) (response *AddRecordTemplateResponse, err error) {
    if request == nil {
        request = NewAddRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAddStreamAuthRequest() (request *AddStreamAuthRequest) {
    request = &AddStreamAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddStreamAuth")
    
    
    return
}

func NewAddStreamAuthResponse() (response *AddStreamAuthResponse) {
    response = &AddStreamAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddStreamAuth
// 用于设置推拉流鉴权配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPULLSTATE = "InvalidParameterValue.InvalidPullState"
//  INVALIDPARAMETERVALUE_INVALIDPUSHSTATE = "InvalidParameterValue.InvalidPushState"
//  INVALIDPARAMETERVALUE_INVALIDSECRET = "InvalidParameterValue.InvalidSecret"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddStreamAuth(request *AddStreamAuthRequest) (response *AddStreamAuthResponse, err error) {
    return c.AddStreamAuthWithContext(context.Background(), request)
}

// AddStreamAuth
// 用于设置推拉流鉴权配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPULLSTATE = "InvalidParameterValue.InvalidPullState"
//  INVALIDPARAMETERVALUE_INVALIDPUSHSTATE = "InvalidParameterValue.InvalidPushState"
//  INVALIDPARAMETERVALUE_INVALIDSECRET = "InvalidParameterValue.InvalidSecret"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddStreamAuthWithContext(ctx context.Context, request *AddStreamAuthRequest) (response *AddStreamAuthResponse, err error) {
    if request == nil {
        request = NewAddStreamAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddStreamAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddStreamAuthResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserDeviceRequest() (request *AddUserDeviceRequest) {
    request = &AddUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddUserDevice")
    
    
    return
}

func NewAddUserDeviceResponse() (response *AddUserDeviceResponse) {
    response = &AddUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddUserDevice
// 用于新增单个设备。添加设备之后，可根据返回结果到设备上进行配置，配置后等待设备注册/推流。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUNDCLUSTER = "ResourceNotFound.NotFoundCluster"
func (c *Client) AddUserDevice(request *AddUserDeviceRequest) (response *AddUserDeviceResponse, err error) {
    return c.AddUserDeviceWithContext(context.Background(), request)
}

// AddUserDevice
// 用于新增单个设备。添加设备之后，可根据返回结果到设备上进行配置，配置后等待设备注册/推流。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUNDCLUSTER = "ResourceNotFound.NotFoundCluster"
func (c *Client) AddUserDeviceWithContext(ctx context.Context, request *AddUserDeviceRequest) (response *AddUserDeviceResponse, err error) {
    if request == nil {
        request = NewAddUserDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewBatchOperateDeviceRequest() (request *BatchOperateDeviceRequest) {
    request = &BatchOperateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchOperateDevice")
    
    
    return
}

func NewBatchOperateDeviceResponse() (response *BatchOperateDeviceResponse) {
    response = &BatchOperateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchOperateDevice
// 用于批量操作（启用，禁用，删除）设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) BatchOperateDevice(request *BatchOperateDeviceRequest) (response *BatchOperateDeviceResponse, err error) {
    return c.BatchOperateDeviceWithContext(context.Background(), request)
}

// BatchOperateDevice
// 用于批量操作（启用，禁用，删除）设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) BatchOperateDeviceWithContext(ctx context.Context, request *BatchOperateDeviceRequest) (response *BatchOperateDeviceResponse, err error) {
    if request == nil {
        request = NewBatchOperateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchOperateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchOperateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDomainRequest() (request *CheckDomainRequest) {
    request = &CheckDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CheckDomain")
    
    
    return
}

func NewCheckDomainResponse() (response *CheckDomainResponse) {
    response = &CheckDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckDomain
// 用于检测域名是否备案。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOMATCHEDCNAME = "FailedOperation.NoMatchedCname"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOMAIN = "InvalidParameterValue.InvalidDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINNOTRECORD = "ResourceUnavailable.DomainNotRecord"
func (c *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    return c.CheckDomainWithContext(context.Background(), request)
}

// CheckDomain
// 用于检测域名是否备案。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOMATCHEDCNAME = "FailedOperation.NoMatchedCname"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOMAIN = "InvalidParameterValue.InvalidDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINNOTRECORD = "ResourceUnavailable.DomainNotRecord"
func (c *Client) CheckDomainWithContext(ctx context.Context, request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    if request == nil {
        request = NewCheckDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDomainResponse()
    err = c.Send(request, response)
    return
}

func NewControlDevicePTZRequest() (request *ControlDevicePTZRequest) {
    request = &ControlDevicePTZRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDevicePTZ")
    
    
    return
}

func NewControlDevicePTZResponse() (response *ControlDevicePTZResponse) {
    response = &ControlDevicePTZResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlDevicePTZ
// 用于设备通道云台控制，包括转动、变倍、变焦、光圈等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPTZCMD = "InvalidParameterValue.UnSupportedPTZCMD"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePTZ(request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    return c.ControlDevicePTZWithContext(context.Background(), request)
}

// ControlDevicePTZ
// 用于设备通道云台控制，包括转动、变倍、变焦、光圈等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPTZCMD = "InvalidParameterValue.UnSupportedPTZCMD"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePTZWithContext(ctx context.Context, request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    if request == nil {
        request = NewControlDevicePTZRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDevicePTZ require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDevicePTZResponse()
    err = c.Send(request, response)
    return
}

func NewControlDevicePresetRequest() (request *ControlDevicePresetRequest) {
    request = &ControlDevicePresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDevicePreset")
    
    
    return
}

func NewControlDevicePresetResponse() (response *ControlDevicePresetResponse) {
    response = &ControlDevicePresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlDevicePreset
// 用于操作设备预置位，包括设置、删除、调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePreset(request *ControlDevicePresetRequest) (response *ControlDevicePresetResponse, err error) {
    return c.ControlDevicePresetWithContext(context.Background(), request)
}

// ControlDevicePreset
// 用于操作设备预置位，包括设置、删除、调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePresetWithContext(ctx context.Context, request *ControlDevicePresetRequest) (response *ControlDevicePresetResponse, err error) {
    if request == nil {
        request = NewControlDevicePresetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDevicePreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDevicePresetResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceStreamRequest() (request *ControlDeviceStreamRequest) {
    request = &ControlDeviceStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDeviceStream")
    
    
    return
}

func NewControlDeviceStreamResponse() (response *ControlDeviceStreamResponse) {
    response = &ControlDeviceStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlDeviceStream
// 用于获取设备的实时开流地址。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEMISMATCHCHANNEL = "InvalidParameterValue.DeviceMismatchChannel"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMTYPE = "InvalidParameterValue.InvalidStreamType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) ControlDeviceStream(request *ControlDeviceStreamRequest) (response *ControlDeviceStreamResponse, err error) {
    return c.ControlDeviceStreamWithContext(context.Background(), request)
}

// ControlDeviceStream
// 用于获取设备的实时开流地址。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEMISMATCHCHANNEL = "InvalidParameterValue.DeviceMismatchChannel"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMTYPE = "InvalidParameterValue.InvalidStreamType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) ControlDeviceStreamWithContext(ctx context.Context, request *ControlDeviceStreamRequest) (response *ControlDeviceStreamResponse, err error) {
    if request == nil {
        request = NewControlDeviceStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDeviceStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDeviceStreamResponse()
    err = c.Send(request, response)
    return
}

func NewControlRecordRequest() (request *ControlRecordRequest) {
    request = &ControlRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlRecord")
    
    
    return
}

func NewControlRecordResponse() (response *ControlRecordResponse) {
    response = &ControlRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlRecord
// 用于录像回放过程中的倍速、跳转、播放/暂停/停止等控制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_STREAMNOTEXISTORCLOSE = "ResourceNotFound.StreamNotExistOrClose"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_SCALEANDPOSBOTHEXIST = "UnsupportedOperation.ScaleAndPosBothExist"
func (c *Client) ControlRecord(request *ControlRecordRequest) (response *ControlRecordResponse, err error) {
    return c.ControlRecordWithContext(context.Background(), request)
}

// ControlRecord
// 用于录像回放过程中的倍速、跳转、播放/暂停/停止等控制。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_STREAMNOTEXISTORCLOSE = "ResourceNotFound.StreamNotExistOrClose"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_SCALEANDPOSBOTHEXIST = "UnsupportedOperation.ScaleAndPosBothExist"
func (c *Client) ControlRecordWithContext(ctx context.Context, request *ControlRecordRequest) (response *ControlRecordResponse, err error) {
    if request == nil {
        request = NewControlRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlRecordResponse()
    err = c.Send(request, response)
    return
}

func NewControlRecordTimelineRequest() (request *ControlRecordTimelineRequest) {
    request = &ControlRecordTimelineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlRecordTimeline")
    
    
    return
}

func NewControlRecordTimelineResponse() (response *ControlRecordTimelineResponse) {
    response = &ControlRecordTimelineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ControlRecordTimeline
// 用于查询设备本地录像时间轴信息，为NVR/IPC本地存储的录像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlRecordTimeline(request *ControlRecordTimelineRequest) (response *ControlRecordTimelineResponse, err error) {
    return c.ControlRecordTimelineWithContext(context.Background(), request)
}

// ControlRecordTimeline
// 用于查询设备本地录像时间轴信息，为NVR/IPC本地存储的录像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlRecordTimelineWithContext(ctx context.Context, request *ControlRecordTimelineRequest) (response *ControlRecordTimelineResponse, err error) {
    if request == nil {
        request = NewControlRecordTimelineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlRecordTimeline require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlRecordTimelineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAITaskRequest() (request *DeleteAITaskRequest) {
    request = &DeleteAITaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteAITask")
    
    
    return
}

func NewDeleteAITaskResponse() (response *DeleteAITaskResponse) {
    response = &DeleteAITaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAITask
// 删除AI任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DeleteAITask(request *DeleteAITaskRequest) (response *DeleteAITaskResponse, err error) {
    return c.DeleteAITaskWithContext(context.Background(), request)
}

// DeleteAITask
// 删除AI任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DeleteAITaskWithContext(ctx context.Context, request *DeleteAITaskRequest) (response *DeleteAITaskResponse, err error) {
    if request == nil {
        request = NewDeleteAITaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAITask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAITaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteDomain")
    
    
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomain
// 用于删除域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    return c.DeleteDomainWithContext(context.Background(), request)
}

// DeleteDomain
// 用于删除域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatewayRequest() (request *DeleteGatewayRequest) {
    request = &DeleteGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteGateway")
    
    
    return
}

func NewDeleteGatewayResponse() (response *DeleteGatewayResponse) {
    response = &DeleteGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGateway
// 用于删除网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGateway(request *DeleteGatewayRequest) (response *DeleteGatewayResponse, err error) {
    return c.DeleteGatewayWithContext(context.Background(), request)
}

// DeleteGateway
// 用于删除网关。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteGatewayWithContext(ctx context.Context, request *DeleteGatewayRequest) (response *DeleteGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
    request = &DeleteOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteOrganization")
    
    
    return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
    response = &DeleteOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOrganization
// 用于删除组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ORGLINKDEV = "UnsupportedOperation.OrgLinkDev"
//  UNSUPPORTEDOPERATION_ORGLINKORG = "UnsupportedOperation.OrgLinkOrg"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    return c.DeleteOrganizationWithContext(context.Background(), request)
}

// DeleteOrganization
// 用于删除组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ORGLINKDEV = "UnsupportedOperation.OrgLinkDev"
//  UNSUPPORTEDOPERATION_ORGLINKORG = "UnsupportedOperation.OrgLinkOrg"
func (c *Client) DeleteOrganizationWithContext(ctx context.Context, request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordBackupPlanRequest() (request *DeleteRecordBackupPlanRequest) {
    request = &DeleteRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordBackupPlan")
    
    
    return
}

func NewDeleteRecordBackupPlanResponse() (response *DeleteRecordBackupPlanResponse) {
    response = &DeleteRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordBackupPlan
// 用于删除录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupPlan(request *DeleteRecordBackupPlanRequest) (response *DeleteRecordBackupPlanResponse, err error) {
    return c.DeleteRecordBackupPlanWithContext(context.Background(), request)
}

// DeleteRecordBackupPlan
// 用于删除录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupPlanWithContext(ctx context.Context, request *DeleteRecordBackupPlanRequest) (response *DeleteRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordBackupPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordBackupTemplateRequest() (request *DeleteRecordBackupTemplateRequest) {
    request = &DeleteRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordBackupTemplate")
    
    
    return
}

func NewDeleteRecordBackupTemplateResponse() (response *DeleteRecordBackupTemplateResponse) {
    response = &DeleteRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordBackupTemplate
// 用于删除录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupTemplate(request *DeleteRecordBackupTemplateRequest) (response *DeleteRecordBackupTemplateResponse, err error) {
    return c.DeleteRecordBackupTemplateWithContext(context.Background(), request)
}

// DeleteRecordBackupTemplate
// 用于删除录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupTemplateWithContext(ctx context.Context, request *DeleteRecordBackupTemplateRequest) (response *DeleteRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRecordBackupTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordPlanRequest() (request *DeleteRecordPlanRequest) {
    request = &DeleteRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordPlan")
    
    
    return
}

func NewDeleteRecordPlanResponse() (response *DeleteRecordPlanResponse) {
    response = &DeleteRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordPlan
// 用于删除实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    return c.DeleteRecordPlanWithContext(context.Background(), request)
}

// DeleteRecordPlan
// 用于删除实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlanWithContext(ctx context.Context, request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRetrieveTaskRequest() (request *DeleteRecordRetrieveTaskRequest) {
    request = &DeleteRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordRetrieveTask")
    
    
    return
}

func NewDeleteRecordRetrieveTaskResponse() (response *DeleteRecordRetrieveTaskResponse) {
    response = &DeleteRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordRetrieveTask
// 用于删除取回任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCEINUSE_RETRIEVETASKEXECUTING = "ResourceInUse.RetrieveTaskExecuting"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DeleteRecordRetrieveTask(request *DeleteRecordRetrieveTaskRequest) (response *DeleteRecordRetrieveTaskResponse, err error) {
    return c.DeleteRecordRetrieveTaskWithContext(context.Background(), request)
}

// DeleteRecordRetrieveTask
// 用于删除取回任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCEINUSE_RETRIEVETASKEXECUTING = "ResourceInUse.RetrieveTaskExecuting"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DeleteRecordRetrieveTaskWithContext(ctx context.Context, request *DeleteRecordRetrieveTaskRequest) (response *DeleteRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRetrieveTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordTemplateRequest() (request *DeleteRecordTemplateRequest) {
    request = &DeleteRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordTemplate")
    
    
    return
}

func NewDeleteRecordTemplateResponse() (response *DeleteRecordTemplateResponse) {
    response = &DeleteRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordTemplate
// 用于删除实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRecordTemplate(request *DeleteRecordTemplateRequest) (response *DeleteRecordTemplateResponse, err error) {
    return c.DeleteRecordTemplateWithContext(context.Background(), request)
}

// DeleteRecordTemplate
// 用于删除实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRecordTemplateWithContext(ctx context.Context, request *DeleteRecordTemplateRequest) (response *DeleteRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserDeviceRequest() (request *DeleteUserDeviceRequest) {
    request = &DeleteUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteUserDevice")
    
    
    return
}

func NewDeleteUserDeviceResponse() (response *DeleteUserDeviceResponse) {
    response = &DeleteUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserDevice
// 用于删除已添加的设备。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserDevice(request *DeleteUserDeviceRequest) (response *DeleteUserDeviceResponse, err error) {
    return c.DeleteUserDeviceWithContext(context.Background(), request)
}

// DeleteUserDevice
// 用于删除已添加的设备。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserDeviceWithContext(ctx context.Context, request *DeleteUserDeviceRequest) (response *DeleteUserDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteUserDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAITaskRequest() (request *DescribeAITaskRequest) {
    request = &DescribeAITaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeAITask")
    
    
    return
}

func NewDescribeAITaskResponse() (response *DescribeAITaskResponse) {
    response = &DescribeAITaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAITask
// 获取AI任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DescribeAITask(request *DescribeAITaskRequest) (response *DescribeAITaskResponse, err error) {
    return c.DescribeAITaskWithContext(context.Background(), request)
}

// DescribeAITask
// 获取AI任务详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DescribeAITaskWithContext(ctx context.Context, request *DescribeAITaskRequest) (response *DescribeAITaskResponse, err error) {
    if request == nil {
        request = NewDescribeAITaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAITask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAITaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAITaskResultRequest() (request *DescribeAITaskResultRequest) {
    request = &DescribeAITaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeAITaskResult")
    
    
    return
}

func NewDescribeAITaskResultResponse() (response *DescribeAITaskResultResponse) {
    response = &DescribeAITaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAITaskResult
// 获取AI任务识别结果
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CHANNELIDMUSTBENOTEMPTY = "InvalidParameterValue.ChannelIdMustBeNotEmpty"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDBEGINANDENDTIME = "InvalidParameterValue.InvalidBeginAndEndTime"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_OBJECTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ObjectContainIllegalCharacter"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DescribeAITaskResult(request *DescribeAITaskResultRequest) (response *DescribeAITaskResultResponse, err error) {
    return c.DescribeAITaskResultWithContext(context.Background(), request)
}

// DescribeAITaskResult
// 获取AI任务识别结果
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CHANNELIDMUSTBENOTEMPTY = "InvalidParameterValue.ChannelIdMustBeNotEmpty"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDBEGINANDENDTIME = "InvalidParameterValue.InvalidBeginAndEndTime"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_OBJECTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ObjectContainIllegalCharacter"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) DescribeAITaskResultWithContext(ctx context.Context, request *DescribeAITaskResultRequest) (response *DescribeAITaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeAITaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAITaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAITaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCNAMERequest() (request *DescribeCNAMERequest) {
    request = &DescribeCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeCNAME")
    
    
    return
}

func NewDescribeCNAMEResponse() (response *DescribeCNAMEResponse) {
    response = &DescribeCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCNAME
// 用于根据服务节点获取 CNAME 值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCNAME(request *DescribeCNAMERequest) (response *DescribeCNAMEResponse, err error) {
    return c.DescribeCNAMEWithContext(context.Background(), request)
}

// DescribeCNAME
// 用于根据服务节点获取 CNAME 值。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCNAMEWithContext(ctx context.Context, request *DescribeCNAMERequest) (response *DescribeCNAMEResponse, err error) {
    if request == nil {
        request = NewDescribeCNAMERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceChannelRequest() (request *DescribeDeviceChannelRequest) {
    request = &DescribeDeviceChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDeviceChannel")
    
    
    return
}

func NewDescribeDeviceChannelResponse() (response *DescribeDeviceChannelResponse) {
    response = &DescribeDeviceChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceChannel
// 用于查询设备的通道。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceChannel(request *DescribeDeviceChannelRequest) (response *DescribeDeviceChannelResponse, err error) {
    return c.DescribeDeviceChannelWithContext(context.Background(), request)
}

// DescribeDeviceChannel
// 用于查询设备的通道。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceChannelWithContext(ctx context.Context, request *DescribeDeviceChannelRequest) (response *DescribeDeviceChannelResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePresetRequest() (request *DescribeDevicePresetRequest) {
    request = &DescribeDevicePresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDevicePreset")
    
    
    return
}

func NewDescribeDevicePresetResponse() (response *DescribeDevicePresetResponse) {
    response = &DescribeDevicePresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePreset
// 用于查询设备通道预置位信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) DescribeDevicePreset(request *DescribeDevicePresetRequest) (response *DescribeDevicePresetResponse, err error) {
    return c.DescribeDevicePresetWithContext(context.Background(), request)
}

// DescribeDevicePreset
// 用于查询设备通道预置位信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) DescribeDevicePresetWithContext(ctx context.Context, request *DescribeDevicePresetRequest) (response *DescribeDevicePresetResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePresetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePresetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRegionRequest() (request *DescribeDeviceRegionRequest) {
    request = &DescribeDeviceRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDeviceRegion")
    
    
    return
}

func NewDescribeDeviceRegionResponse() (response *DescribeDeviceRegionResponse) {
    response = &DescribeDeviceRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceRegion
// 用于添加设备时，查询设备可以使用的服务节点，查询结果为已经绑定了域名的服务节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDeviceRegion(request *DescribeDeviceRegionRequest) (response *DescribeDeviceRegionResponse, err error) {
    return c.DescribeDeviceRegionWithContext(context.Background(), request)
}

// DescribeDeviceRegion
// 用于添加设备时，查询设备可以使用的服务节点，查询结果为已经绑定了域名的服务节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDeviceRegionWithContext(ctx context.Context, request *DescribeDeviceRegionRequest) (response *DescribeDeviceRegionResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDomain")
    
    
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomain
// 用于查询添加的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    return c.DescribeDomainWithContext(context.Background(), request)
}

// DescribeDomain
// 用于查询添加的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainWithContext(ctx context.Context, request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRegionRequest() (request *DescribeDomainRegionRequest) {
    request = &DescribeDomainRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDomainRegion")
    
    
    return
}

func NewDescribeDomainRegionResponse() (response *DescribeDomainRegionResponse) {
    response = &DescribeDomainRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainRegion
// 用于用户添加域名时，查询可以绑定的服务节点，结果为平台支持的所有服务节点。（注意：每个服务节点只能绑定一个域名）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainRegion(request *DescribeDomainRegionRequest) (response *DescribeDomainRegionResponse, err error) {
    return c.DescribeDomainRegionWithContext(context.Background(), request)
}

// DescribeDomainRegion
// 用于用户添加域名时，查询可以绑定的服务节点，结果为平台支持的所有服务节点。（注意：每个服务节点只能绑定一个域名）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainRegionWithContext(ctx context.Context, request *DescribeDomainRegionRequest) (response *DescribeDomainRegionResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayRequest() (request *DescribeGatewayRequest) {
    request = &DescribeGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeGateway")
    
    
    return
}

func NewDescribeGatewayResponse() (response *DescribeGatewayResponse) {
    response = &DescribeGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGateway
// 用于获取网关详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGateway(request *DescribeGatewayRequest) (response *DescribeGatewayResponse, err error) {
    return c.DescribeGatewayWithContext(context.Background(), request)
}

// DescribeGateway
// 用于获取网关详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGatewayWithContext(ctx context.Context, request *DescribeGatewayRequest) (response *DescribeGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayMonitorRequest() (request *DescribeGatewayMonitorRequest) {
    request = &DescribeGatewayMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeGatewayMonitor")
    
    
    return
}

func NewDescribeGatewayMonitorResponse() (response *DescribeGatewayMonitorResponse) {
    response = &DescribeGatewayMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayMonitor
// 用于获取网关的数据及流量监控信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGatewayMonitor(request *DescribeGatewayMonitorRequest) (response *DescribeGatewayMonitorResponse, err error) {
    return c.DescribeGatewayMonitorWithContext(context.Background(), request)
}

// DescribeGatewayMonitor
// 用于获取网关的数据及流量监控信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGatewayMonitorWithContext(ctx context.Context, request *DescribeGatewayMonitorRequest) (response *DescribeGatewayMonitorResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayProtocolRequest() (request *DescribeGatewayProtocolRequest) {
    request = &DescribeGatewayProtocolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeGatewayProtocol")
    
    
    return
}

func NewDescribeGatewayProtocolResponse() (response *DescribeGatewayProtocolResponse) {
    response = &DescribeGatewayProtocolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayProtocol
// 用于查询网关接入协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGatewayProtocol(request *DescribeGatewayProtocolRequest) (response *DescribeGatewayProtocolResponse, err error) {
    return c.DescribeGatewayProtocolWithContext(context.Background(), request)
}

// DescribeGatewayProtocol
// 用于查询网关接入协议。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGatewayProtocolWithContext(ctx context.Context, request *DescribeGatewayProtocolRequest) (response *DescribeGatewayProtocolResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayProtocolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayProtocol require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayProtocolResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayVersionRequest() (request *DescribeGatewayVersionRequest) {
    request = &DescribeGatewayVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeGatewayVersion")
    
    
    return
}

func NewDescribeGatewayVersionResponse() (response *DescribeGatewayVersionResponse) {
    response = &DescribeGatewayVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayVersion
// 查询网关服务版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGatewayVersion(request *DescribeGatewayVersionRequest) (response *DescribeGatewayVersionResponse, err error) {
    return c.DescribeGatewayVersionWithContext(context.Background(), request)
}

// DescribeGatewayVersion
// 查询网关服务版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeGatewayVersionWithContext(ctx context.Context, request *DescribeGatewayVersionRequest) (response *DescribeGatewayVersionResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
    request = &DescribeOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeOrganization")
    
    
    return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
    response = &DescribeOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganization
// 用于查询组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    return c.DescribeOrganizationWithContext(context.Background(), request)
}

// DescribeOrganization
// 用于查询组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordBackupPlanRequest() (request *DescribeRecordBackupPlanRequest) {
    request = &DescribeRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordBackupPlan")
    
    
    return
}

func NewDescribeRecordBackupPlanResponse() (response *DescribeRecordBackupPlanResponse) {
    response = &DescribeRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordBackupPlan
// 用于查询录像上云计划详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupPlan(request *DescribeRecordBackupPlanRequest) (response *DescribeRecordBackupPlanResponse, err error) {
    return c.DescribeRecordBackupPlanWithContext(context.Background(), request)
}

// DescribeRecordBackupPlan
// 用于查询录像上云计划详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupPlanWithContext(ctx context.Context, request *DescribeRecordBackupPlanRequest) (response *DescribeRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewDescribeRecordBackupPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordBackupTemplateRequest() (request *DescribeRecordBackupTemplateRequest) {
    request = &DescribeRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordBackupTemplate")
    
    
    return
}

func NewDescribeRecordBackupTemplateResponse() (response *DescribeRecordBackupTemplateResponse) {
    response = &DescribeRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordBackupTemplate
// 用于查询录像上云模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupTemplate(request *DescribeRecordBackupTemplateRequest) (response *DescribeRecordBackupTemplateResponse, err error) {
    return c.DescribeRecordBackupTemplateWithContext(context.Background(), request)
}

// DescribeRecordBackupTemplate
// 用于查询录像上云模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupTemplateWithContext(ctx context.Context, request *DescribeRecordBackupTemplateRequest) (response *DescribeRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRecordBackupTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordFileRequest() (request *DescribeRecordFileRequest) {
    request = &DescribeRecordFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordFile")
    
    
    return
}

func NewDescribeRecordFileResponse() (response *DescribeRecordFileResponse) {
    response = &DescribeRecordFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordFile
// 用于查询设备云端录像时间轴信息（即为视频上云后设置录像计划后云存储的录像）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordFile(request *DescribeRecordFileRequest) (response *DescribeRecordFileResponse, err error) {
    return c.DescribeRecordFileWithContext(context.Background(), request)
}

// DescribeRecordFile
// 用于查询设备云端录像时间轴信息（即为视频上云后设置录像计划后云存储的录像）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordFileWithContext(ctx context.Context, request *DescribeRecordFileRequest) (response *DescribeRecordFileResponse, err error) {
    if request == nil {
        request = NewDescribeRecordFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordPlanRequest() (request *DescribeRecordPlanRequest) {
    request = &DescribeRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordPlan")
    
    
    return
}

func NewDescribeRecordPlanResponse() (response *DescribeRecordPlanResponse) {
    response = &DescribeRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordPlan
// 用于查询实时上云计划详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DescribeRecordPlan(request *DescribeRecordPlanRequest) (response *DescribeRecordPlanResponse, err error) {
    return c.DescribeRecordPlanWithContext(context.Background(), request)
}

// DescribeRecordPlan
// 用于查询实时上云计划详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DescribeRecordPlanWithContext(ctx context.Context, request *DescribeRecordPlanRequest) (response *DescribeRecordPlanResponse, err error) {
    if request == nil {
        request = NewDescribeRecordPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordPlaybackUrlRequest() (request *DescribeRecordPlaybackUrlRequest) {
    request = &DescribeRecordPlaybackUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordPlaybackUrl")
    
    
    return
}

func NewDescribeRecordPlaybackUrlResponse() (response *DescribeRecordPlaybackUrlResponse) {
    response = &DescribeRecordPlaybackUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordPlaybackUrl
// 用于获取云端录像回放url地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordPlaybackUrl(request *DescribeRecordPlaybackUrlRequest) (response *DescribeRecordPlaybackUrlResponse, err error) {
    return c.DescribeRecordPlaybackUrlWithContext(context.Background(), request)
}

// DescribeRecordPlaybackUrl
// 用于获取云端录像回放url地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordPlaybackUrlWithContext(ctx context.Context, request *DescribeRecordPlaybackUrlRequest) (response *DescribeRecordPlaybackUrlResponse, err error) {
    if request == nil {
        request = NewDescribeRecordPlaybackUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordPlaybackUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordPlaybackUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRetrieveTaskRequest() (request *DescribeRecordRetrieveTaskRequest) {
    request = &DescribeRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordRetrieveTask")
    
    
    return
}

func NewDescribeRecordRetrieveTaskResponse() (response *DescribeRecordRetrieveTaskResponse) {
    response = &DescribeRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordRetrieveTask
// 用于查询云录像取回任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DescribeRecordRetrieveTask(request *DescribeRecordRetrieveTaskRequest) (response *DescribeRecordRetrieveTaskResponse, err error) {
    return c.DescribeRecordRetrieveTaskWithContext(context.Background(), request)
}

// DescribeRecordRetrieveTask
// 用于查询云录像取回任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DescribeRecordRetrieveTaskWithContext(ctx context.Context, request *DescribeRecordRetrieveTaskRequest) (response *DescribeRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRetrieveTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordSliceRequest() (request *DescribeRecordSliceRequest) {
    request = &DescribeRecordSliceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordSlice")
    
    
    return
}

func NewDescribeRecordSliceResponse() (response *DescribeRecordSliceResponse) {
    response = &DescribeRecordSliceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordSlice
// 平台支持将数据以TS切片的形式存入客户自有COS桶，该接口用于支持客户快捷查询切片信息列表
//
// （注意：只支持标准存储类型的查询）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
func (c *Client) DescribeRecordSlice(request *DescribeRecordSliceRequest) (response *DescribeRecordSliceResponse, err error) {
    return c.DescribeRecordSliceWithContext(context.Background(), request)
}

// DescribeRecordSlice
// 平台支持将数据以TS切片的形式存入客户自有COS桶，该接口用于支持客户快捷查询切片信息列表
//
// （注意：只支持标准存储类型的查询）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
func (c *Client) DescribeRecordSliceWithContext(ctx context.Context, request *DescribeRecordSliceRequest) (response *DescribeRecordSliceResponse, err error) {
    if request == nil {
        request = NewDescribeRecordSliceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordSlice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordSliceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTemplateRequest() (request *DescribeRecordTemplateRequest) {
    request = &DescribeRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordTemplate")
    
    
    return
}

func NewDescribeRecordTemplateResponse() (response *DescribeRecordTemplateResponse) {
    response = &DescribeRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordTemplate
// 用于查询实时上云模板详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeRecordTemplate(request *DescribeRecordTemplateRequest) (response *DescribeRecordTemplateResponse, err error) {
    return c.DescribeRecordTemplateWithContext(context.Background(), request)
}

// DescribeRecordTemplate
// 用于查询实时上云模板详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeRecordTemplateWithContext(ctx context.Context, request *DescribeRecordTemplateRequest) (response *DescribeRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamAuthRequest() (request *DescribeStreamAuthRequest) {
    request = &DescribeStreamAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeStreamAuth")
    
    
    return
}

func NewDescribeStreamAuthResponse() (response *DescribeStreamAuthResponse) {
    response = &DescribeStreamAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamAuth
// 用于查询推拉流鉴权配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStreamAuth(request *DescribeStreamAuthRequest) (response *DescribeStreamAuthResponse, err error) {
    return c.DescribeStreamAuthWithContext(context.Background(), request)
}

// DescribeStreamAuth
// 用于查询推拉流鉴权配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStreamAuthWithContext(ctx context.Context, request *DescribeStreamAuthRequest) (response *DescribeStreamAuthResponse, err error) {
    if request == nil {
        request = NewDescribeStreamAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamAuthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 用于查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// 用于查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDeviceRequest() (request *DescribeUserDeviceRequest) {
    request = &DescribeUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeUserDevice")
    
    
    return
}

func NewDescribeUserDeviceResponse() (response *DescribeUserDeviceResponse) {
    response = &DescribeUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserDevice
// 用于查询设备的详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDevice(request *DescribeUserDeviceRequest) (response *DescribeUserDeviceResponse, err error) {
    return c.DescribeUserDeviceWithContext(context.Background(), request)
}

// DescribeUserDevice
// 用于查询设备的详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDeviceWithContext(ctx context.Context, request *DescribeUserDeviceRequest) (response *DescribeUserDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeUserDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoDownloadUrlRequest() (request *DescribeVideoDownloadUrlRequest) {
    request = &DescribeVideoDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVideoDownloadUrl")
    
    
    return
}

func NewDescribeVideoDownloadUrlResponse() (response *DescribeVideoDownloadUrlResponse) {
    response = &DescribeVideoDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoDownloadUrl
// 用于获取云录像下载 url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOWNLOADURLERROR = "InvalidParameter.DownloadUrlError"
//  INVALIDPARAMETER_DOWNLOADURLHASEXPIRED = "InvalidParameter.DownloadUrlHasExpired"
//  INVALIDPARAMETERVALUE_INVALIDBEGINTIME = "InvalidParameterValue.InvalidBeginTime"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENDTIME = "InvalidParameterValue.InvalidEndTime"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_OUTOFTIMERANGE = "InvalidParameterValue.OutOfTimeRange"
//  OPERATIONDENIED_BANDWIDTHLIMITZERO = "OperationDenied.BandwidthLimitZero"
//  OPERATIONDENIED_CONCURRENTDOWNLOADSOVERLIMIT = "OperationDenied.ConcurrentDownloadsOverLimit"
//  OPERATIONDENIED_CONNECTSLIMITZERO = "OperationDenied.ConnectsLimitZero"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VIDEONOTFOUND = "ResourceNotFound.VideoNotFound"
//  RESOURCEUNAVAILABLE_VIDEOARCHIVED = "ResourceUnavailable.VideoArchived"
func (c *Client) DescribeVideoDownloadUrl(request *DescribeVideoDownloadUrlRequest) (response *DescribeVideoDownloadUrlResponse, err error) {
    return c.DescribeVideoDownloadUrlWithContext(context.Background(), request)
}

// DescribeVideoDownloadUrl
// 用于获取云录像下载 url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOWNLOADURLERROR = "InvalidParameter.DownloadUrlError"
//  INVALIDPARAMETER_DOWNLOADURLHASEXPIRED = "InvalidParameter.DownloadUrlHasExpired"
//  INVALIDPARAMETERVALUE_INVALIDBEGINTIME = "InvalidParameterValue.InvalidBeginTime"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENDTIME = "InvalidParameterValue.InvalidEndTime"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_OUTOFTIMERANGE = "InvalidParameterValue.OutOfTimeRange"
//  OPERATIONDENIED_BANDWIDTHLIMITZERO = "OperationDenied.BandwidthLimitZero"
//  OPERATIONDENIED_CONCURRENTDOWNLOADSOVERLIMIT = "OperationDenied.ConcurrentDownloadsOverLimit"
//  OPERATIONDENIED_CONNECTSLIMITZERO = "OperationDenied.ConnectsLimitZero"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VIDEONOTFOUND = "ResourceNotFound.VideoNotFound"
//  RESOURCEUNAVAILABLE_VIDEOARCHIVED = "ResourceUnavailable.VideoArchived"
func (c *Client) DescribeVideoDownloadUrlWithContext(ctx context.Context, request *DescribeVideoDownloadUrlRequest) (response *DescribeVideoDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeVideoDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewListAITasksRequest() (request *ListAITasksRequest) {
    request = &ListAITasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListAITasks")
    
    
    return
}

func NewListAITasksResponse() (response *ListAITasksResponse) {
    response = &ListAITasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAITasks
// 获取AI任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListAITasks(request *ListAITasksRequest) (response *ListAITasksResponse, err error) {
    return c.ListAITasksWithContext(context.Background(), request)
}

// ListAITasks
// 获取AI任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListAITasksWithContext(ctx context.Context, request *ListAITasksRequest) (response *ListAITasksResponse, err error) {
    if request == nil {
        request = NewListAITasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAITasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAITasksResponse()
    err = c.Send(request, response)
    return
}

func NewListDevicesRequest() (request *ListDevicesRequest) {
    request = &ListDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListDevices")
    
    
    return
}

func NewListDevicesResponse() (response *ListDevicesResponse) {
    response = &ListDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListDevices
// 用于获取对应组织下的设备列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDACCESSPROTOCOL = "InvalidParameterValue.InvalidAccessProtocol"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_INVALIDDEVICETYPE = "InvalidParameterValue.InvalidDeviceType"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListDevices(request *ListDevicesRequest) (response *ListDevicesResponse, err error) {
    return c.ListDevicesWithContext(context.Background(), request)
}

// ListDevices
// 用于获取对应组织下的设备列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDACCESSPROTOCOL = "InvalidParameterValue.InvalidAccessProtocol"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_INVALIDDEVICETYPE = "InvalidParameterValue.InvalidDeviceType"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListDevicesWithContext(ctx context.Context, request *ListDevicesRequest) (response *ListDevicesResponse, err error) {
    if request == nil {
        request = NewListDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListGatewayDevicesRequest() (request *ListGatewayDevicesRequest) {
    request = &ListGatewayDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListGatewayDevices")
    
    
    return
}

func NewListGatewayDevicesResponse() (response *ListGatewayDevicesResponse) {
    response = &ListGatewayDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListGatewayDevices
// 用于查询网关下挂载的设备列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListGatewayDevices(request *ListGatewayDevicesRequest) (response *ListGatewayDevicesResponse, err error) {
    return c.ListGatewayDevicesWithContext(context.Background(), request)
}

// ListGatewayDevices
// 用于查询网关下挂载的设备列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListGatewayDevicesWithContext(ctx context.Context, request *ListGatewayDevicesRequest) (response *ListGatewayDevicesResponse, err error) {
    if request == nil {
        request = NewListGatewayDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGatewayDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGatewayDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListGatewaysRequest() (request *ListGatewaysRequest) {
    request = &ListGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListGateways")
    
    
    return
}

func NewListGatewaysResponse() (response *ListGatewaysResponse) {
    response = &ListGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListGateways
// 用于获取网关列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
func (c *Client) ListGateways(request *ListGatewaysRequest) (response *ListGatewaysResponse, err error) {
    return c.ListGatewaysWithContext(context.Background(), request)
}

// ListGateways
// 用于获取网关列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
func (c *Client) ListGatewaysWithContext(ctx context.Context, request *ListGatewaysRequest) (response *ListGatewaysResponse, err error) {
    if request == nil {
        request = NewListGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationChannelNumbersRequest() (request *ListOrganizationChannelNumbersRequest) {
    request = &ListOrganizationChannelNumbersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListOrganizationChannelNumbers")
    
    
    return
}

func NewListOrganizationChannelNumbersResponse() (response *ListOrganizationChannelNumbersResponse) {
    response = &ListOrganizationChannelNumbersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListOrganizationChannelNumbers
// 用于查询组织目录下的未添加到实时上云计划中的通道数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannelNumbers(request *ListOrganizationChannelNumbersRequest) (response *ListOrganizationChannelNumbersResponse, err error) {
    return c.ListOrganizationChannelNumbersWithContext(context.Background(), request)
}

// ListOrganizationChannelNumbers
// 用于查询组织目录下的未添加到实时上云计划中的通道数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannelNumbersWithContext(ctx context.Context, request *ListOrganizationChannelNumbersRequest) (response *ListOrganizationChannelNumbersResponse, err error) {
    if request == nil {
        request = NewListOrganizationChannelNumbersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationChannelNumbers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationChannelNumbersResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationChannelsRequest() (request *ListOrganizationChannelsRequest) {
    request = &ListOrganizationChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListOrganizationChannels")
    
    
    return
}

func NewListOrganizationChannelsResponse() (response *ListOrganizationChannelsResponse) {
    response = &ListOrganizationChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListOrganizationChannels
// 用于查询组织目录下的通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannels(request *ListOrganizationChannelsRequest) (response *ListOrganizationChannelsResponse, err error) {
    return c.ListOrganizationChannelsWithContext(context.Background(), request)
}

// ListOrganizationChannels
// 用于查询组织目录下的通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannelsWithContext(ctx context.Context, request *ListOrganizationChannelsRequest) (response *ListOrganizationChannelsResponse, err error) {
    if request == nil {
        request = NewListOrganizationChannelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupPlanDevicesRequest() (request *ListRecordBackupPlanDevicesRequest) {
    request = &ListRecordBackupPlanDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupPlanDevices")
    
    
    return
}

func NewListRecordBackupPlanDevicesResponse() (response *ListRecordBackupPlanDevicesResponse) {
    response = &ListRecordBackupPlanDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordBackupPlanDevices
// 用于查询录像上云计划下的设备通道列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELNAME = "InvalidParameterValue.InvalidChannelName"
//  INVALIDPARAMETERVALUE_INVALIDDEVICENAME = "InvalidParameterValue.InvalidDeviceName"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
func (c *Client) ListRecordBackupPlanDevices(request *ListRecordBackupPlanDevicesRequest) (response *ListRecordBackupPlanDevicesResponse, err error) {
    return c.ListRecordBackupPlanDevicesWithContext(context.Background(), request)
}

// ListRecordBackupPlanDevices
// 用于查询录像上云计划下的设备通道列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELNAME = "InvalidParameterValue.InvalidChannelName"
//  INVALIDPARAMETERVALUE_INVALIDDEVICENAME = "InvalidParameterValue.InvalidDeviceName"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
func (c *Client) ListRecordBackupPlanDevicesWithContext(ctx context.Context, request *ListRecordBackupPlanDevicesRequest) (response *ListRecordBackupPlanDevicesResponse, err error) {
    if request == nil {
        request = NewListRecordBackupPlanDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupPlanDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupPlanDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupPlansRequest() (request *ListRecordBackupPlansRequest) {
    request = &ListRecordBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupPlans")
    
    
    return
}

func NewListRecordBackupPlansResponse() (response *ListRecordBackupPlansResponse) {
    response = &ListRecordBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordBackupPlans
// 用于查询录像上云计划列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupPlans(request *ListRecordBackupPlansRequest) (response *ListRecordBackupPlansResponse, err error) {
    return c.ListRecordBackupPlansWithContext(context.Background(), request)
}

// ListRecordBackupPlans
// 用于查询录像上云计划列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupPlansWithContext(ctx context.Context, request *ListRecordBackupPlansRequest) (response *ListRecordBackupPlansResponse, err error) {
    if request == nil {
        request = NewListRecordBackupPlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupTemplatesRequest() (request *ListRecordBackupTemplatesRequest) {
    request = &ListRecordBackupTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupTemplates")
    
    
    return
}

func NewListRecordBackupTemplatesResponse() (response *ListRecordBackupTemplatesResponse) {
    response = &ListRecordBackupTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordBackupTemplates
// 用于查询录像上云模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupTemplates(request *ListRecordBackupTemplatesRequest) (response *ListRecordBackupTemplatesResponse, err error) {
    return c.ListRecordBackupTemplatesWithContext(context.Background(), request)
}

// ListRecordBackupTemplates
// 用于查询录像上云模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupTemplatesWithContext(ctx context.Context, request *ListRecordBackupTemplatesRequest) (response *ListRecordBackupTemplatesResponse, err error) {
    if request == nil {
        request = NewListRecordBackupTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordPlanChannelsRequest() (request *ListRecordPlanChannelsRequest) {
    request = &ListRecordPlanChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordPlanChannels")
    
    
    return
}

func NewListRecordPlanChannelsResponse() (response *ListRecordPlanChannelsResponse) {
    response = &ListRecordPlanChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordPlanChannels
// 用于查询用户下所有实时上云计划中的通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlanChannels(request *ListRecordPlanChannelsRequest) (response *ListRecordPlanChannelsResponse, err error) {
    return c.ListRecordPlanChannelsWithContext(context.Background(), request)
}

// ListRecordPlanChannels
// 用于查询用户下所有实时上云计划中的通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlanChannelsWithContext(ctx context.Context, request *ListRecordPlanChannelsRequest) (response *ListRecordPlanChannelsResponse, err error) {
    if request == nil {
        request = NewListRecordPlanChannelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordPlanChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordPlanChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordPlanDevicesRequest() (request *ListRecordPlanDevicesRequest) {
    request = &ListRecordPlanDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordPlanDevices")
    
    
    return
}

func NewListRecordPlanDevicesResponse() (response *ListRecordPlanDevicesResponse) {
    response = &ListRecordPlanDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordPlanDevices
// 用于查询实时上云计划下的设备通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) ListRecordPlanDevices(request *ListRecordPlanDevicesRequest) (response *ListRecordPlanDevicesResponse, err error) {
    return c.ListRecordPlanDevicesWithContext(context.Background(), request)
}

// ListRecordPlanDevices
// 用于查询实时上云计划下的设备通道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) ListRecordPlanDevicesWithContext(ctx context.Context, request *ListRecordPlanDevicesRequest) (response *ListRecordPlanDevicesResponse, err error) {
    if request == nil {
        request = NewListRecordPlanDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordPlanDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordPlanDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordPlansRequest() (request *ListRecordPlansRequest) {
    request = &ListRecordPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordPlans")
    
    
    return
}

func NewListRecordPlansResponse() (response *ListRecordPlansResponse) {
    response = &ListRecordPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordPlans
// 用于查询实时上云计划列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlans(request *ListRecordPlansRequest) (response *ListRecordPlansResponse, err error) {
    return c.ListRecordPlansWithContext(context.Background(), request)
}

// ListRecordPlans
// 用于查询实时上云计划列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlansWithContext(ctx context.Context, request *ListRecordPlansRequest) (response *ListRecordPlansResponse, err error) {
    if request == nil {
        request = NewListRecordPlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordPlansResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordRetrieveTasksRequest() (request *ListRecordRetrieveTasksRequest) {
    request = &ListRecordRetrieveTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordRetrieveTasks")
    
    
    return
}

func NewListRecordRetrieveTasksResponse() (response *ListRecordRetrieveTasksResponse) {
    response = &ListRecordRetrieveTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordRetrieveTasks
// 用于查询取回任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordRetrieveTasks(request *ListRecordRetrieveTasksRequest) (response *ListRecordRetrieveTasksResponse, err error) {
    return c.ListRecordRetrieveTasksWithContext(context.Background(), request)
}

// ListRecordRetrieveTasks
// 用于查询取回任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordRetrieveTasksWithContext(ctx context.Context, request *ListRecordRetrieveTasksRequest) (response *ListRecordRetrieveTasksResponse, err error) {
    if request == nil {
        request = NewListRecordRetrieveTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordRetrieveTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordRetrieveTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordTemplatesRequest() (request *ListRecordTemplatesRequest) {
    request = &ListRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordTemplates")
    
    
    return
}

func NewListRecordTemplatesResponse() (response *ListRecordTemplatesResponse) {
    response = &ListRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRecordTemplates
// 用于查询实时上云模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordTemplates(request *ListRecordTemplatesRequest) (response *ListRecordTemplatesResponse, err error) {
    return c.ListRecordTemplatesWithContext(context.Background(), request)
}

// ListRecordTemplates
// 用于查询实时上云模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordTemplatesWithContext(ctx context.Context, request *ListRecordTemplatesRequest) (response *ListRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewListRecordTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListSubTasksRequest() (request *ListSubTasksRequest) {
    request = &ListSubTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListSubTasks")
    
    
    return
}

func NewListSubTasksResponse() (response *ListSubTasksResponse) {
    response = &ListSubTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSubTasks
// 用于查询任务的子任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListSubTasks(request *ListSubTasksRequest) (response *ListSubTasksResponse, err error) {
    return c.ListSubTasksWithContext(context.Background(), request)
}

// ListSubTasks
// 用于查询任务的子任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListSubTasksWithContext(ctx context.Context, request *ListSubTasksRequest) (response *ListSubTasksResponse, err error) {
    if request == nil {
        request = NewListSubTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSubTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSubTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTasks
// 用于查询批量任务和简单任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// 用于查询批量任务和简单任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPlayRecordRequest() (request *PlayRecordRequest) {
    request = &PlayRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "PlayRecord")
    
    
    return
}

func NewPlayRecordResponse() (response *PlayRecordResponse) {
    response = &PlayRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PlayRecord
// 用于获取设备本地录像 URL 地址。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) PlayRecord(request *PlayRecordRequest) (response *PlayRecordResponse, err error) {
    return c.PlayRecordWithContext(context.Background(), request)
}

// PlayRecord
// 用于获取设备本地录像 URL 地址。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) PlayRecordWithContext(ctx context.Context, request *PlayRecordRequest) (response *PlayRecordResponse, err error) {
    if request == nil {
        request = NewPlayRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PlayRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPlayRecordResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshDeviceChannelRequest() (request *RefreshDeviceChannelRequest) {
    request = &RefreshDeviceChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "RefreshDeviceChannel")
    
    
    return
}

func NewRefreshDeviceChannelResponse() (response *RefreshDeviceChannelResponse) {
    response = &RefreshDeviceChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefreshDeviceChannel
// 用于刷新国标设备的通道（接口调用后，触发向设备请求通道列表，新增的通道入库，设备上已删除的通道需自行删除、后台不自动删除）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) RefreshDeviceChannel(request *RefreshDeviceChannelRequest) (response *RefreshDeviceChannelResponse, err error) {
    return c.RefreshDeviceChannelWithContext(context.Background(), request)
}

// RefreshDeviceChannel
// 用于刷新国标设备的通道（接口调用后，触发向设备请求通道列表，新增的通道入库，设备上已删除的通道需自行删除、后台不自动删除）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) RefreshDeviceChannelWithContext(ctx context.Context, request *RefreshDeviceChannelRequest) (response *RefreshDeviceChannelResponse, err error) {
    if request == nil {
        request = NewRefreshDeviceChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshDeviceChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshDeviceChannelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAITaskRequest() (request *UpdateAITaskRequest) {
    request = &UpdateAITaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateAITask")
    
    
    return
}

func NewUpdateAITaskResponse() (response *UpdateAITaskResponse) {
    response = &UpdateAITaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAITask
// 更新AI任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_CALLBACKURLCONTAINILLEGALCHARACTER = "InvalidParameterValue.CallbackURLContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELIDALREADYEXISTSINOTHERAITASKS = "InvalidParameterValue.ChannelIdAlreadyExistsInOtherAITasks"
//  INVALIDPARAMETERVALUE_CHANNELLISTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ChannelListContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELNUMBERMUSTBELESSTHANONETHOUSAND = "InvalidParameterValue.ChannelNumberMustBeLessThanOneThousand"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_HASDUPLICATEOPERTIMESLOT = "InvalidParameterValue.HasDuplicateOperTimeSlot"
//  INVALIDPARAMETERVALUE_INVALIDAITASKDESC = "InvalidParameterValue.InvalidAITaskDesc"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDAITASKNAME = "InvalidParameterValue.InvalidAITaskName"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATETAG = "InvalidParameterValue.InvalidTemplateTag"
//  INVALIDPARAMETERVALUE_INVALIDTIMEINTERVAL = "InvalidParameterValue.InvalidTimeInterval"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTCONTAINILLEGALCHARACTER = "InvalidParameterValue.OperTimeSlotContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTNUMBERMUSTBELESSTHANFIVE = "InvalidParameterValue.OperTimeSlotNumberMustBeLessThanFive"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTSTARTMUSTBELESSTHANEND = "InvalidParameterValue.OperTimeSlotStartMustBeLessThanEnd"
//  INVALIDPARAMETERVALUE_TEMPLATETAGMUSTBECONSISTENT = "InvalidParameterValue.TemplateTagMustBeConsistent"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) UpdateAITask(request *UpdateAITaskRequest) (response *UpdateAITaskResponse, err error) {
    return c.UpdateAITaskWithContext(context.Background(), request)
}

// UpdateAITask
// 更新AI任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_CALLBACKURLCONTAINILLEGALCHARACTER = "InvalidParameterValue.CallbackURLContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELIDALREADYEXISTSINOTHERAITASKS = "InvalidParameterValue.ChannelIdAlreadyExistsInOtherAITasks"
//  INVALIDPARAMETERVALUE_CHANNELLISTCONTAINILLEGALCHARACTER = "InvalidParameterValue.ChannelListContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_CHANNELNUMBERMUSTBELESSTHANONETHOUSAND = "InvalidParameterValue.ChannelNumberMustBeLessThanOneThousand"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_HASDUPLICATEOPERTIMESLOT = "InvalidParameterValue.HasDuplicateOperTimeSlot"
//  INVALIDPARAMETERVALUE_INVALIDAITASKDESC = "InvalidParameterValue.InvalidAITaskDesc"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDAITASKNAME = "InvalidParameterValue.InvalidAITaskName"
//  INVALIDPARAMETERVALUE_INVALIDDETECTTYPE = "InvalidParameterValue.InvalidDetectType"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATETAG = "InvalidParameterValue.InvalidTemplateTag"
//  INVALIDPARAMETERVALUE_INVALIDTIMEINTERVAL = "InvalidParameterValue.InvalidTimeInterval"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTCONTAINILLEGALCHARACTER = "InvalidParameterValue.OperTimeSlotContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTNUMBERMUSTBELESSTHANFIVE = "InvalidParameterValue.OperTimeSlotNumberMustBeLessThanFive"
//  INVALIDPARAMETERVALUE_OPERTIMESLOTSTARTMUSTBELESSTHANEND = "InvalidParameterValue.OperTimeSlotStartMustBeLessThanEnd"
//  INVALIDPARAMETERVALUE_TEMPLATETAGMUSTBECONSISTENT = "InvalidParameterValue.TemplateTagMustBeConsistent"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) UpdateAITaskWithContext(ctx context.Context, request *UpdateAITaskRequest) (response *UpdateAITaskResponse, err error) {
    if request == nil {
        request = NewUpdateAITaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAITask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAITaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAITaskStatusRequest() (request *UpdateAITaskStatusRequest) {
    request = &UpdateAITaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateAITaskStatus")
    
    
    return
}

func NewUpdateAITaskStatusResponse() (response *UpdateAITaskStatusResponse) {
    response = &UpdateAITaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAITaskStatus
// 更新 AI 任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISOFF = "FailedOperation.AITaskStatusIsOff"
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDAITASKSTATUS = "InvalidParameterValue.InvalidAITaskStatus"
//  INVALIDPARAMETERVALUE_STATUSMUSTBENOTEMPTY = "InvalidParameterValue.StatusMustBeNotEmpty"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) UpdateAITaskStatus(request *UpdateAITaskStatusRequest) (response *UpdateAITaskStatusResponse, err error) {
    return c.UpdateAITaskStatusWithContext(context.Background(), request)
}

// UpdateAITaskStatus
// 更新 AI 任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITASKSTATUSISOFF = "FailedOperation.AITaskStatusIsOff"
//  FAILEDOPERATION_AITASKSTATUSISON = "FailedOperation.AITaskStatusIsOn"
//  INVALIDPARAMETERVALUE_INVALIDAITASKID = "InvalidParameterValue.InvalidAITaskID"
//  INVALIDPARAMETERVALUE_INVALIDAITASKSTATUS = "InvalidParameterValue.InvalidAITaskStatus"
//  INVALIDPARAMETERVALUE_STATUSMUSTBENOTEMPTY = "InvalidParameterValue.StatusMustBeNotEmpty"
//  RESOURCENOTFOUND_AITASKNOTEXISTED = "ResourceNotFound.AITaskNotExisted"
func (c *Client) UpdateAITaskStatusWithContext(ctx context.Context, request *UpdateAITaskStatusRequest) (response *UpdateAITaskStatusResponse, err error) {
    if request == nil {
        request = NewUpdateAITaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAITaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAITaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceOrganizationRequest() (request *UpdateDeviceOrganizationRequest) {
    request = &UpdateDeviceOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateDeviceOrganization")
    
    
    return
}

func NewUpdateDeviceOrganizationResponse() (response *UpdateDeviceOrganizationResponse) {
    response = &UpdateDeviceOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceOrganization
// 用于批量更改设备的组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceOrganization(request *UpdateDeviceOrganizationRequest) (response *UpdateDeviceOrganizationResponse, err error) {
    return c.UpdateDeviceOrganizationWithContext(context.Background(), request)
}

// UpdateDeviceOrganization
// 用于批量更改设备的组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceOrganizationWithContext(ctx context.Context, request *UpdateDeviceOrganizationRequest) (response *UpdateDeviceOrganizationResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceStatusRequest() (request *UpdateDeviceStatusRequest) {
    request = &UpdateDeviceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateDeviceStatus")
    
    
    return
}

func NewUpdateDeviceStatusResponse() (response *UpdateDeviceStatusResponse) {
    response = &UpdateDeviceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceStatus
// 用于启用/禁用设备，禁用后拒绝设备注册。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceStatus(request *UpdateDeviceStatusRequest) (response *UpdateDeviceStatusResponse, err error) {
    return c.UpdateDeviceStatusWithContext(context.Background(), request)
}

// UpdateDeviceStatus
// 用于启用/禁用设备，禁用后拒绝设备注册。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceStatusWithContext(ctx context.Context, request *UpdateDeviceStatusRequest) (response *UpdateDeviceStatusResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayRequest() (request *UpdateGatewayRequest) {
    request = &UpdateGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateGateway")
    
    
    return
}

func NewUpdateGatewayResponse() (response *UpdateGatewayResponse) {
    response = &UpdateGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGateway
// 用于修改网关信息（支持对网关名称和描述的修改）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateGateway(request *UpdateGatewayRequest) (response *UpdateGatewayResponse, err error) {
    return c.UpdateGatewayWithContext(context.Background(), request)
}

// UpdateGateway
// 用于修改网关信息（支持对网关名称和描述的修改）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateGatewayWithContext(ctx context.Context, request *UpdateGatewayRequest) (response *UpdateGatewayResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationRequest() (request *UpdateOrganizationRequest) {
    request = &UpdateOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateOrganization")
    
    
    return
}

func NewUpdateOrganizationResponse() (response *UpdateOrganizationResponse) {
    response = &UpdateOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateOrganization
// 用于修改组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateOrganization(request *UpdateOrganizationRequest) (response *UpdateOrganizationResponse, err error) {
    return c.UpdateOrganizationWithContext(context.Background(), request)
}

// UpdateOrganization
// 用于修改组织。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateOrganizationWithContext(ctx context.Context, request *UpdateOrganizationRequest) (response *UpdateOrganizationResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordBackupPlanRequest() (request *UpdateRecordBackupPlanRequest) {
    request = &UpdateRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordBackupPlan")
    
    
    return
}

func NewUpdateRecordBackupPlanResponse() (response *UpdateRecordBackupPlanResponse) {
    response = &UpdateRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecordBackupPlan
// 用于修改录像上云计划。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupPlan(request *UpdateRecordBackupPlanRequest) (response *UpdateRecordBackupPlanResponse, err error) {
    return c.UpdateRecordBackupPlanWithContext(context.Background(), request)
}

// UpdateRecordBackupPlan
// 用于修改录像上云计划。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupPlanWithContext(ctx context.Context, request *UpdateRecordBackupPlanRequest) (response *UpdateRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordBackupPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordBackupTemplateRequest() (request *UpdateRecordBackupTemplateRequest) {
    request = &UpdateRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordBackupTemplate")
    
    
    return
}

func NewUpdateRecordBackupTemplateResponse() (response *UpdateRecordBackupTemplateResponse) {
    response = &UpdateRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecordBackupTemplate
// 用于修改录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupTemplate(request *UpdateRecordBackupTemplateRequest) (response *UpdateRecordBackupTemplateResponse, err error) {
    return c.UpdateRecordBackupTemplateWithContext(context.Background(), request)
}

// UpdateRecordBackupTemplate
// 用于修改录像上云模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupTemplateWithContext(ctx context.Context, request *UpdateRecordBackupTemplateRequest) (response *UpdateRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateRecordBackupTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordPlanRequest() (request *UpdateRecordPlanRequest) {
    request = &UpdateRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordPlan")
    
    
    return
}

func NewUpdateRecordPlanResponse() (response *UpdateRecordPlanResponse) {
    response = &UpdateRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecordPlan
// 用于修改实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATIONRANGE = "InvalidParameterValue.InvalidExpirationRange"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTRANSITIONRANGE = "InvalidParameterValue.InvalidTransitionRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  MISSINGPARAMETER_MISSINGLIFERULEPARAM = "MissingParameter.MissingLifeRuleParam"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_CHANNELREPEATADD = "ResourceInUse.ChannelRepeatAdd"
//  RESOURCEINUSE_PLANDELETING = "ResourceInUse.PlanDeleting"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    return c.UpdateRecordPlanWithContext(context.Background(), request)
}

// UpdateRecordPlan
// 用于修改实时上云计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATIONRANGE = "InvalidParameterValue.InvalidExpirationRange"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTRANSITIONRANGE = "InvalidParameterValue.InvalidTransitionRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  MISSINGPARAMETER_MISSINGLIFERULEPARAM = "MissingParameter.MissingLifeRuleParam"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_CHANNELREPEATADD = "ResourceInUse.ChannelRepeatAdd"
//  RESOURCEINUSE_PLANDELETING = "ResourceInUse.PlanDeleting"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordPlanWithContext(ctx context.Context, request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordTemplateRequest() (request *UpdateRecordTemplateRequest) {
    request = &UpdateRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordTemplate")
    
    
    return
}

func NewUpdateRecordTemplateResponse() (response *UpdateRecordTemplateResponse) {
    response = &UpdateRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecordTemplate
// 用于修改实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordTemplate(request *UpdateRecordTemplateRequest) (response *UpdateRecordTemplateResponse, err error) {
    return c.UpdateRecordTemplateWithContext(context.Background(), request)
}

// UpdateRecordTemplate
// 用于修改实时上云模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordTemplateWithContext(ctx context.Context, request *UpdateRecordTemplateRequest) (response *UpdateRecordTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserDeviceRequest() (request *UpdateUserDeviceRequest) {
    request = &UpdateUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateUserDevice")
    
    
    return
}

func NewUpdateUserDeviceResponse() (response *UpdateUserDeviceResponse) {
    response = &UpdateUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserDevice
// 用于修改设备的配置信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTREAMPROTOCOL = "InvalidParameterValue.UnsupportedStreamProtocol"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateUserDevice(request *UpdateUserDeviceRequest) (response *UpdateUserDeviceResponse, err error) {
    return c.UpdateUserDeviceWithContext(context.Background(), request)
}

// UpdateUserDevice
// 用于修改设备的配置信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTREAMPROTOCOL = "InvalidParameterValue.UnsupportedStreamProtocol"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateUserDeviceWithContext(ctx context.Context, request *UpdateUserDeviceRequest) (response *UpdateUserDeviceResponse, err error) {
    if request == nil {
        request = NewUpdateUserDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGatewayRequest() (request *UpgradeGatewayRequest) {
    request = &UpgradeGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpgradeGateway")
    
    
    return
}

func NewUpgradeGatewayResponse() (response *UpgradeGatewayResponse) {
    response = &UpgradeGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeGateway
// 用于网关升级（支持对所有待更新的服务一键升级）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGateway(request *UpgradeGatewayRequest) (response *UpgradeGatewayResponse, err error) {
    return c.UpgradeGatewayWithContext(context.Background(), request)
}

// UpgradeGateway
// 用于网关升级（支持对所有待更新的服务一键升级）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYID = "InvalidParameterValue.InvalidGatewayId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGatewayWithContext(ctx context.Context, request *UpgradeGatewayRequest) (response *UpgradeGatewayResponse, err error) {
    if request == nil {
        request = NewUpgradeGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGatewayResponse()
    err = c.Send(request, response)
    return
}
