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

package v20200210

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

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


func NewBindNumberCallOutSkillGroupRequest() (request *BindNumberCallOutSkillGroupRequest) {
    request = &BindNumberCallOutSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "BindNumberCallOutSkillGroup")
    
    
    return
}

func NewBindNumberCallOutSkillGroupResponse() (response *BindNumberCallOutSkillGroupResponse) {
    response = &BindNumberCallOutSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindNumberCallOutSkillGroup
// 绑定号码外呼技能组
func (c *Client) BindNumberCallOutSkillGroup(request *BindNumberCallOutSkillGroupRequest) (response *BindNumberCallOutSkillGroupResponse, err error) {
    return c.BindNumberCallOutSkillGroupWithContext(context.Background(), request)
}

// BindNumberCallOutSkillGroup
// 绑定号码外呼技能组
func (c *Client) BindNumberCallOutSkillGroupWithContext(ctx context.Context, request *BindNumberCallOutSkillGroupRequest) (response *BindNumberCallOutSkillGroupResponse, err error) {
    if request == nil {
        request = NewBindNumberCallOutSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindNumberCallOutSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindNumberCallOutSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindStaffSkillGroupListRequest() (request *BindStaffSkillGroupListRequest) {
    request = &BindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "BindStaffSkillGroupList")
    
    
    return
}

func NewBindStaffSkillGroupListResponse() (response *BindStaffSkillGroupListResponse) {
    response = &BindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindStaffSkillGroupList
// 绑定坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupList(request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    return c.BindStaffSkillGroupListWithContext(context.Background(), request)
}

// BindStaffSkillGroupList
// 绑定坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupListWithContext(ctx context.Context, request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewBindStaffSkillGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindStaffSkillGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoCalloutTaskRequest() (request *CreateAutoCalloutTaskRequest) {
    request = &CreateAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateAutoCalloutTask")
    
    
    return
}

func NewCreateAutoCalloutTaskResponse() (response *CreateAutoCalloutTaskResponse) {
    response = &CreateAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoCalloutTask
// 创建自动外呼任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTask(request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    return c.CreateAutoCalloutTaskWithContext(context.Background(), request)
}

// CreateAutoCalloutTask
// 创建自动外呼任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTaskWithContext(ctx context.Context, request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewCreateAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCCSkillGroupRequest() (request *CreateCCCSkillGroupRequest) {
    request = &CreateCCCSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCCCSkillGroup")
    
    
    return
}

func NewCreateCCCSkillGroupResponse() (response *CreateCCCSkillGroupResponse) {
    response = &CreateCCCSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCCCSkillGroup
// 创建技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) CreateCCCSkillGroup(request *CreateCCCSkillGroupRequest) (response *CreateCCCSkillGroupResponse, err error) {
    return c.CreateCCCSkillGroupWithContext(context.Background(), request)
}

// CreateCCCSkillGroup
// 创建技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) CreateCCCSkillGroupWithContext(ctx context.Context, request *CreateCCCSkillGroupRequest) (response *CreateCCCSkillGroupResponse, err error) {
    if request == nil {
        request = NewCreateCCCSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCCSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCCCSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCallOutSessionRequest() (request *CreateCallOutSessionRequest) {
    request = &CreateCallOutSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCallOutSession")
    
    
    return
}

func NewCreateCallOutSessionResponse() (response *CreateCallOutSessionResponse) {
    response = &CreateCallOutSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCallOutSession
// 创建外呼会话，当前仅支持双呼，即先使用平台号码呼出到坐席手机上，坐席接听后，然后再外呼用户，而且由于运营商频率限制，坐席手机号必须先加白名单，避免频控导致外呼失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCallOutSession(request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    return c.CreateCallOutSessionWithContext(context.Background(), request)
}

// CreateCallOutSession
// 创建外呼会话，当前仅支持双呼，即先使用平台号码呼出到坐席手机上，坐席接听后，然后再外呼用户，而且由于运营商频率限制，坐席手机号必须先加白名单，避免频控导致外呼失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCallOutSessionWithContext(ctx context.Context, request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    if request == nil {
        request = NewCreateCallOutSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCallOutSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCallOutSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCarrierPrivilegeNumberApplicantRequest() (request *CreateCarrierPrivilegeNumberApplicantRequest) {
    request = &CreateCarrierPrivilegeNumberApplicantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCarrierPrivilegeNumberApplicant")
    
    
    return
}

func NewCreateCarrierPrivilegeNumberApplicantResponse() (response *CreateCarrierPrivilegeNumberApplicantResponse) {
    response = &CreateCarrierPrivilegeNumberApplicantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCarrierPrivilegeNumberApplicant
// 用于无限频率地呼叫坐席手机
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCarrierPrivilegeNumberApplicant(request *CreateCarrierPrivilegeNumberApplicantRequest) (response *CreateCarrierPrivilegeNumberApplicantResponse, err error) {
    return c.CreateCarrierPrivilegeNumberApplicantWithContext(context.Background(), request)
}

// CreateCarrierPrivilegeNumberApplicant
// 用于无限频率地呼叫坐席手机
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCarrierPrivilegeNumberApplicantWithContext(ctx context.Context, request *CreateCarrierPrivilegeNumberApplicantRequest) (response *CreateCarrierPrivilegeNumberApplicantResponse, err error) {
    if request == nil {
        request = NewCreateCarrierPrivilegeNumberApplicantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCarrierPrivilegeNumberApplicant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCarrierPrivilegeNumberApplicantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExtensionRequest() (request *CreateExtensionRequest) {
    request = &CreateExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateExtension")
    
    
    return
}

func NewCreateExtensionResponse() (response *CreateExtensionResponse) {
    response = &CreateExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateExtension
// 创建话机账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateExtension(request *CreateExtensionRequest) (response *CreateExtensionResponse, err error) {
    return c.CreateExtensionWithContext(context.Background(), request)
}

// CreateExtension
// 创建话机账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateExtensionWithContext(ctx context.Context, request *CreateExtensionRequest) (response *CreateExtensionResponse, err error) {
    if request == nil {
        request = NewCreateExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSDKLoginTokenRequest() (request *CreateSDKLoginTokenRequest) {
    request = &CreateSDKLoginTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateSDKLoginToken")
    
    
    return
}

func NewCreateSDKLoginTokenResponse() (response *CreateSDKLoginTokenResponse) {
    response = &CreateSDKLoginTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSDKLoginToken
// 创建 SDK 登录 Token。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateSDKLoginToken(request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    return c.CreateSDKLoginTokenWithContext(context.Background(), request)
}

// CreateSDKLoginToken
// 创建 SDK 登录 Token。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateSDKLoginTokenWithContext(ctx context.Context, request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    if request == nil {
        request = NewCreateSDKLoginTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSDKLoginToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSDKLoginTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaffRequest() (request *CreateStaffRequest) {
    request = &CreateStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateStaff")
    
    
    return
}

func NewCreateStaffResponse() (response *CreateStaffResponse) {
    response = &CreateStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStaff
// 创建客服账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaff(request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    return c.CreateStaffWithContext(context.Background(), request)
}

// CreateStaff
// 创建客服账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaffWithContext(ctx context.Context, request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    if request == nil {
        request = NewCreateStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStaffResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserSigRequest() (request *CreateUserSigRequest) {
    request = &CreateUserSigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateUserSig")
    
    
    return
}

func NewCreateUserSigResponse() (response *CreateUserSigResponse) {
    response = &CreateUserSigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserSig
// 创建用户数据签名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateUserSig(request *CreateUserSigRequest) (response *CreateUserSigResponse, err error) {
    return c.CreateUserSigWithContext(context.Background(), request)
}

// CreateUserSig
// 创建用户数据签名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateUserSigWithContext(ctx context.Context, request *CreateUserSigRequest) (response *CreateUserSigResponse, err error) {
    if request == nil {
        request = NewCreateUserSigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserSig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserSigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExtensionRequest() (request *DeleteExtensionRequest) {
    request = &DeleteExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DeleteExtension")
    
    
    return
}

func NewDeleteExtensionResponse() (response *DeleteExtensionResponse) {
    response = &DeleteExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteExtension
// 删除话机账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteExtension(request *DeleteExtensionRequest) (response *DeleteExtensionResponse, err error) {
    return c.DeleteExtensionWithContext(context.Background(), request)
}

// DeleteExtension
// 删除话机账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteExtensionWithContext(ctx context.Context, request *DeleteExtensionRequest) (response *DeleteExtensionResponse, err error) {
    if request == nil {
        request = NewDeleteExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStaffRequest() (request *DeleteStaffRequest) {
    request = &DeleteStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DeleteStaff")
    
    
    return
}

func NewDeleteStaffResponse() (response *DeleteStaffResponse) {
    response = &DeleteStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStaff
// 删除坐席信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteStaff(request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    return c.DeleteStaffWithContext(context.Background(), request)
}

// DeleteStaff
// 删除坐席信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteStaffWithContext(ctx context.Context, request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    if request == nil {
        request = NewDeleteStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStaffResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActiveCarrierPrivilegeNumberRequest() (request *DescribeActiveCarrierPrivilegeNumberRequest) {
    request = &DescribeActiveCarrierPrivilegeNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeActiveCarrierPrivilegeNumber")
    
    
    return
}

func NewDescribeActiveCarrierPrivilegeNumberResponse() (response *DescribeActiveCarrierPrivilegeNumberResponse) {
    response = &DescribeActiveCarrierPrivilegeNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeActiveCarrierPrivilegeNumber
// 查询生效运营商白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeActiveCarrierPrivilegeNumber(request *DescribeActiveCarrierPrivilegeNumberRequest) (response *DescribeActiveCarrierPrivilegeNumberResponse, err error) {
    return c.DescribeActiveCarrierPrivilegeNumberWithContext(context.Background(), request)
}

// DescribeActiveCarrierPrivilegeNumber
// 查询生效运营商白名单规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeActiveCarrierPrivilegeNumberWithContext(ctx context.Context, request *DescribeActiveCarrierPrivilegeNumberRequest) (response *DescribeActiveCarrierPrivilegeNumberResponse, err error) {
    if request == nil {
        request = NewDescribeActiveCarrierPrivilegeNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActiveCarrierPrivilegeNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActiveCarrierPrivilegeNumberResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTaskRequest() (request *DescribeAutoCalloutTaskRequest) {
    request = &DescribeAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTask")
    
    
    return
}

func NewDescribeAutoCalloutTaskResponse() (response *DescribeAutoCalloutTaskResponse) {
    response = &DescribeAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoCalloutTask
// 查询自动外呼任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTask(request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    return c.DescribeAutoCalloutTaskWithContext(context.Background(), request)
}

// DescribeAutoCalloutTask
// 查询自动外呼任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTaskWithContext(ctx context.Context, request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTasksRequest() (request *DescribeAutoCalloutTasksRequest) {
    request = &DescribeAutoCalloutTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTasks")
    
    
    return
}

func NewDescribeAutoCalloutTasksResponse() (response *DescribeAutoCalloutTasksResponse) {
    response = &DescribeAutoCalloutTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoCalloutTasks
// 批量查询自动任务外呼
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasks(request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    return c.DescribeAutoCalloutTasksWithContext(context.Background(), request)
}

// DescribeAutoCalloutTasks
// 批量查询自动任务外呼
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasksWithContext(ctx context.Context, request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoCalloutTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCCBuyInfoListRequest() (request *DescribeCCCBuyInfoListRequest) {
    request = &DescribeCCCBuyInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCCCBuyInfoList")
    
    
    return
}

func NewDescribeCCCBuyInfoListResponse() (response *DescribeCCCBuyInfoListResponse) {
    response = &DescribeCCCBuyInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCCBuyInfoList
// 获取用户购买信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoList(request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    return c.DescribeCCCBuyInfoListWithContext(context.Background(), request)
}

// DescribeCCCBuyInfoList
// 获取用户购买信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoListWithContext(ctx context.Context, request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeCCCBuyInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCCBuyInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCCBuyInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallInMetricsRequest() (request *DescribeCallInMetricsRequest) {
    request = &DescribeCallInMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCallInMetrics")
    
    
    return
}

func NewDescribeCallInMetricsResponse() (response *DescribeCallInMetricsResponse) {
    response = &DescribeCallInMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCallInMetrics
// 获取呼入实时数据统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetrics(request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    return c.DescribeCallInMetricsWithContext(context.Background(), request)
}

// DescribeCallInMetrics
// 获取呼入实时数据统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetricsWithContext(ctx context.Context, request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeCallInMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallInMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallInMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCarrierPrivilegeNumberApplicantsRequest() (request *DescribeCarrierPrivilegeNumberApplicantsRequest) {
    request = &DescribeCarrierPrivilegeNumberApplicantsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCarrierPrivilegeNumberApplicants")
    
    
    return
}

func NewDescribeCarrierPrivilegeNumberApplicantsResponse() (response *DescribeCarrierPrivilegeNumberApplicantsResponse) {
    response = &DescribeCarrierPrivilegeNumberApplicantsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCarrierPrivilegeNumberApplicants
// 查询单状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCarrierPrivilegeNumberApplicants(request *DescribeCarrierPrivilegeNumberApplicantsRequest) (response *DescribeCarrierPrivilegeNumberApplicantsResponse, err error) {
    return c.DescribeCarrierPrivilegeNumberApplicantsWithContext(context.Background(), request)
}

// DescribeCarrierPrivilegeNumberApplicants
// 查询单状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCarrierPrivilegeNumberApplicantsWithContext(ctx context.Context, request *DescribeCarrierPrivilegeNumberApplicantsRequest) (response *DescribeCarrierPrivilegeNumberApplicantsResponse, err error) {
    if request == nil {
        request = NewDescribeCarrierPrivilegeNumberApplicantsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCarrierPrivilegeNumberApplicants require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCarrierPrivilegeNumberApplicantsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChatMessagesRequest() (request *DescribeChatMessagesRequest) {
    request = &DescribeChatMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeChatMessages")
    
    
    return
}

func NewDescribeChatMessagesResponse() (response *DescribeChatMessagesResponse) {
    response = &DescribeChatMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChatMessages
// 获取指定服务记录文本聊天内容，需要先使用查询在线客服记录（DescribeIMCdrs） API 获取服务记录 SessionId。
//
// 
//
// 文本聊天记录只保存了 1 年内的，1 年之前会自动清理。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
func (c *Client) DescribeChatMessages(request *DescribeChatMessagesRequest) (response *DescribeChatMessagesResponse, err error) {
    return c.DescribeChatMessagesWithContext(context.Background(), request)
}

// DescribeChatMessages
// 获取指定服务记录文本聊天内容，需要先使用查询在线客服记录（DescribeIMCdrs） API 获取服务记录 SessionId。
//
// 
//
// 文本聊天记录只保存了 1 年内的，1 年之前会自动清理。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
func (c *Client) DescribeChatMessagesWithContext(ctx context.Context, request *DescribeChatMessagesRequest) (response *DescribeChatMessagesResponse, err error) {
    if request == nil {
        request = NewDescribeChatMessagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChatMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChatMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtensionRequest() (request *DescribeExtensionRequest) {
    request = &DescribeExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeExtension")
    
    
    return
}

func NewDescribeExtensionResponse() (response *DescribeExtensionResponse) {
    response = &DescribeExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExtension
// 获取话机信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DescribeExtension(request *DescribeExtensionRequest) (response *DescribeExtensionResponse, err error) {
    return c.DescribeExtensionWithContext(context.Background(), request)
}

// DescribeExtension
// 获取话机信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DescribeExtensionWithContext(ctx context.Context, request *DescribeExtensionRequest) (response *DescribeExtensionResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtensionsRequest() (request *DescribeExtensionsRequest) {
    request = &DescribeExtensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeExtensions")
    
    
    return
}

func NewDescribeExtensionsResponse() (response *DescribeExtensionsResponse) {
    response = &DescribeExtensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExtensions
// 查询话机列表信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeExtensions(request *DescribeExtensionsRequest) (response *DescribeExtensionsResponse, err error) {
    return c.DescribeExtensionsWithContext(context.Background(), request)
}

// DescribeExtensions
// 查询话机列表信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeExtensionsWithContext(ctx context.Context, request *DescribeExtensionsRequest) (response *DescribeExtensionsResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtensions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtensionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIMCdrsRequest() (request *DescribeIMCdrsRequest) {
    request = &DescribeIMCdrsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeIMCdrs")
    
    
    return
}

func NewDescribeIMCdrsResponse() (response *DescribeIMCdrsResponse) {
    response = &DescribeIMCdrsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIMCdrs
// 获取包括全媒体和文本会话两种类型的服务记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeIMCdrs(request *DescribeIMCdrsRequest) (response *DescribeIMCdrsResponse, err error) {
    return c.DescribeIMCdrsWithContext(context.Background(), request)
}

// DescribeIMCdrs
// 获取包括全媒体和文本会话两种类型的服务记录。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeIMCdrsWithContext(ctx context.Context, request *DescribeIMCdrsRequest) (response *DescribeIMCdrsResponse, err error) {
    if request == nil {
        request = NewDescribeIMCdrsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIMCdrs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIMCdrsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNumbersRequest() (request *DescribeNumbersRequest) {
    request = &DescribeNumbersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeNumbers")
    
    
    return
}

func NewDescribeNumbersResponse() (response *DescribeNumbersResponse) {
    response = &DescribeNumbersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNumbers
// 查询号码列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeNumbers(request *DescribeNumbersRequest) (response *DescribeNumbersResponse, err error) {
    return c.DescribeNumbersWithContext(context.Background(), request)
}

// DescribeNumbers
// 查询号码列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeNumbersWithContext(ctx context.Context, request *DescribeNumbersRequest) (response *DescribeNumbersResponse, err error) {
    if request == nil {
        request = NewDescribeNumbersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNumbers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNumbersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePSTNActiveSessionListRequest() (request *DescribePSTNActiveSessionListRequest) {
    request = &DescribePSTNActiveSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePSTNActiveSessionList")
    
    
    return
}

func NewDescribePSTNActiveSessionListResponse() (response *DescribePSTNActiveSessionListResponse) {
    response = &DescribePSTNActiveSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePSTNActiveSessionList
// 获取当前正在通话的会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionList(request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    return c.DescribePSTNActiveSessionListWithContext(context.Background(), request)
}

// DescribePSTNActiveSessionList
// 获取当前正在通话的会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionListWithContext(ctx context.Context, request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    if request == nil {
        request = NewDescribePSTNActiveSessionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePSTNActiveSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePSTNActiveSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectedTelCdrRequest() (request *DescribeProtectedTelCdrRequest) {
    request = &DescribeProtectedTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeProtectedTelCdr")
    
    
    return
}

func NewDescribeProtectedTelCdrResponse() (response *DescribeProtectedTelCdrResponse) {
    response = &DescribeProtectedTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProtectedTelCdr
// 获取主被叫受保护的电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeProtectedTelCdr(request *DescribeProtectedTelCdrRequest) (response *DescribeProtectedTelCdrResponse, err error) {
    return c.DescribeProtectedTelCdrWithContext(context.Background(), request)
}

// DescribeProtectedTelCdr
// 获取主被叫受保护的电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeProtectedTelCdrWithContext(ctx context.Context, request *DescribeProtectedTelCdrRequest) (response *DescribeProtectedTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeProtectedTelCdrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectedTelCdr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectedTelCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillGroupInfoListRequest() (request *DescribeSkillGroupInfoListRequest) {
    request = &DescribeSkillGroupInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeSkillGroupInfoList")
    
    
    return
}

func NewDescribeSkillGroupInfoListResponse() (response *DescribeSkillGroupInfoListResponse) {
    response = &DescribeSkillGroupInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSkillGroupInfoList
// 获取技能组信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoList(request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    return c.DescribeSkillGroupInfoListWithContext(context.Background(), request)
}

// DescribeSkillGroupInfoList
// 获取技能组信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoListWithContext(ctx context.Context, request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillGroupInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillGroupInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillGroupInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffInfoListRequest() (request *DescribeStaffInfoListRequest) {
    request = &DescribeStaffInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffInfoList")
    
    
    return
}

func NewDescribeStaffInfoListResponse() (response *DescribeStaffInfoListResponse) {
    response = &DescribeStaffInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStaffInfoList
// 获取坐席信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoList(request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    return c.DescribeStaffInfoListWithContext(context.Background(), request)
}

// DescribeStaffInfoList
// 获取坐席信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoListWithContext(ctx context.Context, request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStaffInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaffInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaffInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffStatusMetricsRequest() (request *DescribeStaffStatusMetricsRequest) {
    request = &DescribeStaffStatusMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffStatusMetrics")
    
    
    return
}

func NewDescribeStaffStatusMetricsResponse() (response *DescribeStaffStatusMetricsResponse) {
    response = &DescribeStaffStatusMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStaffStatusMetrics
// 获取坐席实时状态统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetrics(request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    return c.DescribeStaffStatusMetricsWithContext(context.Background(), request)
}

// DescribeStaffStatusMetrics
// 获取坐席实时状态统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetricsWithContext(ctx context.Context, request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeStaffStatusMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaffStatusMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaffStatusMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCallInfoRequest() (request *DescribeTelCallInfoRequest) {
    request = &DescribeTelCallInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCallInfo")
    
    
    return
}

func NewDescribeTelCallInfoResponse() (response *DescribeTelCallInfoResponse) {
    response = &DescribeTelCallInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelCallInfo
// 按实例获取电话消耗统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfo(request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    return c.DescribeTelCallInfoWithContext(context.Background(), request)
}

// DescribeTelCallInfo
// 按实例获取电话消耗统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfoWithContext(ctx context.Context, request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTelCallInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelCallInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelCallInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCdrRequest() (request *DescribeTelCdrRequest) {
    request = &DescribeTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCdr")
    
    
    return
}

func NewDescribeTelCdrResponse() (response *DescribeTelCdrResponse) {
    response = &DescribeTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelCdr
// 获取电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdr(request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    return c.DescribeTelCdrWithContext(context.Background(), request)
}

// DescribeTelCdr
// 获取电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdrWithContext(ctx context.Context, request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeTelCdrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelCdr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelSessionRequest() (request *DescribeTelSessionRequest) {
    request = &DescribeTelSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelSession")
    
    
    return
}

func NewDescribeTelSessionResponse() (response *DescribeTelSessionResponse) {
    response = &DescribeTelSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelSession
// 获取 PSTN 会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSession(request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    return c.DescribeTelSessionWithContext(context.Background(), request)
}

// DescribeTelSession
// 获取 PSTN 会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSessionWithContext(ctx context.Context, request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    if request == nil {
        request = NewDescribeTelSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCCCPhoneNumberRequest() (request *DisableCCCPhoneNumberRequest) {
    request = &DisableCCCPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DisableCCCPhoneNumber")
    
    
    return
}

func NewDisableCCCPhoneNumberResponse() (response *DisableCCCPhoneNumberResponse) {
    response = &DisableCCCPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableCCCPhoneNumber
// 停用号码
//
// 可能返回的错误码:
//  FAILEDOPERATION_CURSTATENOTALLOWMODIFY = "FailedOperation.CurStateNotAllowModify"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEADDRESS = "InvalidParameter.DuplicateAddress"
//  INVALIDPARAMETER_DUPLICATEPHONENUMBER = "InvalidParameter.DuplicatePhoneNumber"
//  INVALIDPARAMETER_DUPLICATESIPACCOUNT = "InvalidParameter.DuplicateSipAccount"
//  INVALIDPARAMETER_ILLEGALADDRESS = "InvalidParameter.IllegalAddress"
//  INVALIDPARAMETER_ILLEGALPHONENUMBER = "InvalidParameter.IllegalPhoneNumber"
//  INVALIDPARAMETER_INVALIDADDRESS = "InvalidParameter.InvalidAddress"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPHONENUMBER = "InvalidParameter.InvalidPhoneNumber"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_SIPACCOUNTPASSWORDFORMAT = "InvalidParameter.SipAccountPasswordFormat"
//  INVALIDPARAMETER_SIPACCOUNTUSERFORMAT = "InvalidParameter.SipAccountUserFormat"
//  INVALIDPARAMETER_SIPTRUNKINUSED = "InvalidParameter.SipTrunkInUsed"
//  INVALIDPARAMETER_SIPTRUNKNOTFOUND = "InvalidParameter.SipTrunkNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DisableCCCPhoneNumber(request *DisableCCCPhoneNumberRequest) (response *DisableCCCPhoneNumberResponse, err error) {
    return c.DisableCCCPhoneNumberWithContext(context.Background(), request)
}

// DisableCCCPhoneNumber
// 停用号码
//
// 可能返回的错误码:
//  FAILEDOPERATION_CURSTATENOTALLOWMODIFY = "FailedOperation.CurStateNotAllowModify"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEADDRESS = "InvalidParameter.DuplicateAddress"
//  INVALIDPARAMETER_DUPLICATEPHONENUMBER = "InvalidParameter.DuplicatePhoneNumber"
//  INVALIDPARAMETER_DUPLICATESIPACCOUNT = "InvalidParameter.DuplicateSipAccount"
//  INVALIDPARAMETER_ILLEGALADDRESS = "InvalidParameter.IllegalAddress"
//  INVALIDPARAMETER_ILLEGALPHONENUMBER = "InvalidParameter.IllegalPhoneNumber"
//  INVALIDPARAMETER_INVALIDADDRESS = "InvalidParameter.InvalidAddress"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPHONENUMBER = "InvalidParameter.InvalidPhoneNumber"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_SIPACCOUNTPASSWORDFORMAT = "InvalidParameter.SipAccountPasswordFormat"
//  INVALIDPARAMETER_SIPACCOUNTUSERFORMAT = "InvalidParameter.SipAccountUserFormat"
//  INVALIDPARAMETER_SIPTRUNKINUSED = "InvalidParameter.SipTrunkInUsed"
//  INVALIDPARAMETER_SIPTRUNKNOTFOUND = "InvalidParameter.SipTrunkNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DisableCCCPhoneNumberWithContext(ctx context.Context, request *DisableCCCPhoneNumberRequest) (response *DisableCCCPhoneNumberResponse, err error) {
    if request == nil {
        request = NewDisableCCCPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableCCCPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableCCCPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewHangUpCallRequest() (request *HangUpCallRequest) {
    request = &HangUpCallRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "HangUpCall")
    
    
    return
}

func NewHangUpCallResponse() (response *HangUpCallResponse) {
    response = &HangUpCallResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HangUpCall
// 挂断电话
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) HangUpCall(request *HangUpCallRequest) (response *HangUpCallResponse, err error) {
    return c.HangUpCallWithContext(context.Background(), request)
}

// HangUpCall
// 挂断电话
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) HangUpCallWithContext(ctx context.Context, request *HangUpCallRequest) (response *HangUpCallResponse, err error) {
    if request == nil {
        request = NewHangUpCallRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HangUpCall require credential")
    }

    request.SetContext(ctx)
    
    response = NewHangUpCallResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExtensionRequest() (request *ModifyExtensionRequest) {
    request = &ModifyExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ModifyExtension")
    
    
    return
}

func NewModifyExtensionResponse() (response *ModifyExtensionResponse) {
    response = &ModifyExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyExtension
// 修改话机账号(绑定技能组、绑定坐席账号)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) ModifyExtension(request *ModifyExtensionRequest) (response *ModifyExtensionResponse, err error) {
    return c.ModifyExtensionWithContext(context.Background(), request)
}

// ModifyExtension
// 修改话机账号(绑定技能组、绑定坐席账号)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) ModifyExtensionWithContext(ctx context.Context, request *ModifyExtensionRequest) (response *ModifyExtensionResponse, err error) {
    if request == nil {
        request = NewModifyExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStaffRequest() (request *ModifyStaffRequest) {
    request = &ModifyStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ModifyStaff")
    
    
    return
}

func NewModifyStaffResponse() (response *ModifyStaffResponse) {
    response = &ModifyStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStaff
// 修改客服账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
func (c *Client) ModifyStaff(request *ModifyStaffRequest) (response *ModifyStaffResponse, err error) {
    return c.ModifyStaffWithContext(context.Background(), request)
}

// ModifyStaff
// 修改客服账号
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
func (c *Client) ModifyStaffWithContext(ctx context.Context, request *ModifyStaffRequest) (response *ModifyStaffResponse, err error) {
    if request == nil {
        request = NewModifyStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStaffResponse()
    err = c.Send(request, response)
    return
}

func NewResetExtensionPasswordRequest() (request *ResetExtensionPasswordRequest) {
    request = &ResetExtensionPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ResetExtensionPassword")
    
    
    return
}

func NewResetExtensionPasswordResponse() (response *ResetExtensionPasswordResponse) {
    response = &ResetExtensionPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetExtensionPassword
// 重置话机注册密码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
func (c *Client) ResetExtensionPassword(request *ResetExtensionPasswordRequest) (response *ResetExtensionPasswordResponse, err error) {
    return c.ResetExtensionPasswordWithContext(context.Background(), request)
}

// ResetExtensionPassword
// 重置话机注册密码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
func (c *Client) ResetExtensionPasswordWithContext(ctx context.Context, request *ResetExtensionPasswordRequest) (response *ResetExtensionPasswordResponse, err error) {
    if request == nil {
        request = NewResetExtensionPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetExtensionPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetExtensionPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStopAutoCalloutTaskRequest() (request *StopAutoCalloutTaskRequest) {
    request = &StopAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "StopAutoCalloutTask")
    
    
    return
}

func NewStopAutoCalloutTaskResponse() (response *StopAutoCalloutTaskResponse) {
    response = &StopAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopAutoCalloutTask
// 停止自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTask(request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    return c.StopAutoCalloutTaskWithContext(context.Background(), request)
}

// StopAutoCalloutTask
// 停止自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTaskWithContext(ctx context.Context, request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewStopAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindNumberCallOutSkillGroupRequest() (request *UnbindNumberCallOutSkillGroupRequest) {
    request = &UnbindNumberCallOutSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UnbindNumberCallOutSkillGroup")
    
    
    return
}

func NewUnbindNumberCallOutSkillGroupResponse() (response *UnbindNumberCallOutSkillGroupResponse) {
    response = &UnbindNumberCallOutSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindNumberCallOutSkillGroup
// 解绑号码外呼技能组
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindNumberCallOutSkillGroup(request *UnbindNumberCallOutSkillGroupRequest) (response *UnbindNumberCallOutSkillGroupResponse, err error) {
    return c.UnbindNumberCallOutSkillGroupWithContext(context.Background(), request)
}

// UnbindNumberCallOutSkillGroup
// 解绑号码外呼技能组
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindNumberCallOutSkillGroupWithContext(ctx context.Context, request *UnbindNumberCallOutSkillGroupRequest) (response *UnbindNumberCallOutSkillGroupResponse, err error) {
    if request == nil {
        request = NewUnbindNumberCallOutSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindNumberCallOutSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindNumberCallOutSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindStaffSkillGroupListRequest() (request *UnbindStaffSkillGroupListRequest) {
    request = &UnbindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UnbindStaffSkillGroupList")
    
    
    return
}

func NewUnbindStaffSkillGroupListResponse() (response *UnbindStaffSkillGroupListResponse) {
    response = &UnbindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindStaffSkillGroupList
// 解绑坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupList(request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    return c.UnbindStaffSkillGroupListWithContext(context.Background(), request)
}

// UnbindStaffSkillGroupList
// 解绑坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupListWithContext(ctx context.Context, request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewUnbindStaffSkillGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindStaffSkillGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}
