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

package v20220817

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-17"

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


func NewAddGroupMemberRequest() (request *AddGroupMemberRequest) {
    request = &AddGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "AddGroupMember")
    
    
    return
}

func NewAddGroupMemberResponse() (response *AddGroupMemberResponse) {
    response = &AddGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddGroupMember
// 此接口用于添加成员列表到指定群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddGroupMember(request *AddGroupMemberRequest) (response *AddGroupMemberResponse, err error) {
    return c.AddGroupMemberWithContext(context.Background(), request)
}

// AddGroupMember
// 此接口用于添加成员列表到指定群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddGroupMemberWithContext(ctx context.Context, request *AddGroupMemberRequest) (response *AddGroupMemberResponse, err error) {
    if request == nil {
        request = NewAddGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchAddGroupMemberRequest() (request *BatchAddGroupMemberRequest) {
    request = &BatchAddGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchAddGroupMember")
    
    
    return
}

func NewBatchAddGroupMemberResponse() (response *BatchAddGroupMemberResponse) {
    response = &BatchAddGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchAddGroupMember
// 此接口用于批量添加成员列表到指定群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchAddGroupMember(request *BatchAddGroupMemberRequest) (response *BatchAddGroupMemberResponse, err error) {
    return c.BatchAddGroupMemberWithContext(context.Background(), request)
}

// BatchAddGroupMember
// 此接口用于批量添加成员列表到指定群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchAddGroupMemberWithContext(ctx context.Context, request *BatchAddGroupMemberRequest) (response *BatchAddGroupMemberResponse, err error) {
    if request == nil {
        request = NewBatchAddGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchAddGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchAddGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateGroupWithMembersRequest() (request *BatchCreateGroupWithMembersRequest) {
    request = &BatchCreateGroupWithMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchCreateGroupWithMembers")
    
    
    return
}

func NewBatchCreateGroupWithMembersResponse() (response *BatchCreateGroupWithMembersResponse) {
    response = &BatchCreateGroupWithMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchCreateGroupWithMembers
// 此接口用于批量创建群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateGroupWithMembers(request *BatchCreateGroupWithMembersRequest) (response *BatchCreateGroupWithMembersResponse, err error) {
    return c.BatchCreateGroupWithMembersWithContext(context.Background(), request)
}

// BatchCreateGroupWithMembers
// 此接口用于批量创建群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateGroupWithMembersWithContext(ctx context.Context, request *BatchCreateGroupWithMembersRequest) (response *BatchCreateGroupWithMembersResponse, err error) {
    if request == nil {
        request = NewBatchCreateGroupWithMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateGroupWithMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateGroupWithMembersResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateRoomRequest() (request *BatchCreateRoomRequest) {
    request = &BatchCreateRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchCreateRoom")
    
    
    return
}

func NewBatchCreateRoomResponse() (response *BatchCreateRoomResponse) {
    response = &BatchCreateRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchCreateRoom
// 批量创建房间接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateRoom(request *BatchCreateRoomRequest) (response *BatchCreateRoomResponse, err error) {
    return c.BatchCreateRoomWithContext(context.Background(), request)
}

// BatchCreateRoom
// 批量创建房间接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateRoomWithContext(ctx context.Context, request *BatchCreateRoomRequest) (response *BatchCreateRoomResponse, err error) {
    if request == nil {
        request = NewBatchCreateRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateRoomResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteGroupMemberRequest() (request *BatchDeleteGroupMemberRequest) {
    request = &BatchDeleteGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDeleteGroupMember")
    
    
    return
}

func NewBatchDeleteGroupMemberResponse() (response *BatchDeleteGroupMemberResponse) {
    response = &BatchDeleteGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeleteGroupMember
// 此接口用于批量删除成员列表到指定群组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteGroupMember(request *BatchDeleteGroupMemberRequest) (response *BatchDeleteGroupMemberResponse, err error) {
    return c.BatchDeleteGroupMemberWithContext(context.Background(), request)
}

// BatchDeleteGroupMember
// 此接口用于批量删除成员列表到指定群组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteGroupMemberWithContext(ctx context.Context, request *BatchDeleteGroupMemberRequest) (response *BatchDeleteGroupMemberResponse, err error) {
    if request == nil {
        request = NewBatchDeleteGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteRecordRequest() (request *BatchDeleteRecordRequest) {
    request = &BatchDeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDeleteRecord")
    
    
    return
}

func NewBatchDeleteRecordResponse() (response *BatchDeleteRecordResponse) {
    response = &BatchDeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeleteRecord
// 批量删除多个房间的录制文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteRecord(request *BatchDeleteRecordRequest) (response *BatchDeleteRecordResponse, err error) {
    return c.BatchDeleteRecordWithContext(context.Background(), request)
}

// BatchDeleteRecord
// 批量删除多个房间的录制文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteRecordWithContext(ctx context.Context, request *BatchDeleteRecordRequest) (response *BatchDeleteRecordResponse, err error) {
    if request == nil {
        request = NewBatchDeleteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRegisterRequest() (request *BatchRegisterRequest) {
    request = &BatchRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchRegister")
    
    
    return
}

func NewBatchRegisterResponse() (response *BatchRegisterResponse) {
    response = &BatchRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchRegister
// 如果批量注册的用户已存在，则会被覆盖。一次最多注册1000个用户。默认请求频率限制：10次/秒
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegister(request *BatchRegisterRequest) (response *BatchRegisterResponse, err error) {
    return c.BatchRegisterWithContext(context.Background(), request)
}

// BatchRegister
// 如果批量注册的用户已存在，则会被覆盖。一次最多注册1000个用户。默认请求频率限制：10次/秒
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegisterWithContext(ctx context.Context, request *BatchRegisterRequest) (response *BatchRegisterResponse, err error) {
    if request == nil {
        request = NewBatchRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewBindDocumentToRoomRequest() (request *BindDocumentToRoomRequest) {
    request = &BindDocumentToRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BindDocumentToRoom")
    
    
    return
}

func NewBindDocumentToRoomResponse() (response *BindDocumentToRoomResponse) {
    response = &BindDocumentToRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindDocumentToRoom
// 绑定文档到房间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoom(request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    return c.BindDocumentToRoomWithContext(context.Background(), request)
}

// BindDocumentToRoom
// 绑定文档到房间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoomWithContext(ctx context.Context, request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    if request == nil {
        request = NewBindDocumentToRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDocumentToRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDocumentToRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocumentRequest() (request *CreateDocumentRequest) {
    request = &CreateDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateDocument")
    
    
    return
}

func NewCreateDocumentResponse() (response *CreateDocumentResponse) {
    response = &CreateDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDocument
// 创建房间内可以使用的文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// 创建房间内可以使用的文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupWithMembersRequest() (request *CreateGroupWithMembersRequest) {
    request = &CreateGroupWithMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateGroupWithMembers")
    
    
    return
}

func NewCreateGroupWithMembersResponse() (response *CreateGroupWithMembersResponse) {
    response = &CreateGroupWithMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroupWithMembers
// 此接口根据成员列表创建群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithMembers(request *CreateGroupWithMembersRequest) (response *CreateGroupWithMembersResponse, err error) {
    return c.CreateGroupWithMembersWithContext(context.Background(), request)
}

// CreateGroupWithMembers
// 此接口根据成员列表创建群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithMembersWithContext(ctx context.Context, request *CreateGroupWithMembersRequest) (response *CreateGroupWithMembersResponse, err error) {
    if request == nil {
        request = NewCreateGroupWithMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupWithMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupWithMembersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupWithSubGroupRequest() (request *CreateGroupWithSubGroupRequest) {
    request = &CreateGroupWithSubGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateGroupWithSubGroup")
    
    
    return
}

func NewCreateGroupWithSubGroupResponse() (response *CreateGroupWithSubGroupResponse) {
    response = &CreateGroupWithSubGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroupWithSubGroup
// 此接口会聚合子群组创建联合群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithSubGroup(request *CreateGroupWithSubGroupRequest) (response *CreateGroupWithSubGroupResponse, err error) {
    return c.CreateGroupWithSubGroupWithContext(context.Background(), request)
}

// CreateGroupWithSubGroup
// 此接口会聚合子群组创建联合群组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithSubGroupWithContext(ctx context.Context, request *CreateGroupWithSubGroupRequest) (response *CreateGroupWithSubGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupWithSubGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupWithSubGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupWithSubGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoomRequest() (request *CreateRoomRequest) {
    request = &CreateRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateRoom")
    
    
    return
}

func NewCreateRoomResponse() (response *CreateRoomResponse) {
    response = &CreateRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoom
// 创建房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoom(request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    return c.CreateRoomWithContext(context.Background(), request)
}

// CreateRoom
// 创建房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoomWithContext(ctx context.Context, request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    if request == nil {
        request = NewCreateRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSupervisorRequest() (request *CreateSupervisorRequest) {
    request = &CreateSupervisorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateSupervisor")
    
    
    return
}

func NewCreateSupervisorResponse() (response *CreateSupervisorResponse) {
    response = &CreateSupervisorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSupervisor
// 创建巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisor(request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    return c.CreateSupervisorWithContext(context.Background(), request)
}

// CreateSupervisor
// 创建巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisorWithContext(ctx context.Context, request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    if request == nil {
        request = NewCreateSupervisorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSupervisor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSupervisorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppCustomContentRequest() (request *DeleteAppCustomContentRequest) {
    request = &DeleteAppCustomContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteAppCustomContent")
    
    
    return
}

func NewDeleteAppCustomContentResponse() (response *DeleteAppCustomContentResponse) {
    response = &DeleteAppCustomContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAppCustomContent
// 删除设置自定义元素。如果参数scenes为空则删除所有自定义元素，否则删除指定的scene自定义元素。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DeleteAppCustomContent(request *DeleteAppCustomContentRequest) (response *DeleteAppCustomContentResponse, err error) {
    return c.DeleteAppCustomContentWithContext(context.Background(), request)
}

// DeleteAppCustomContent
// 删除设置自定义元素。如果参数scenes为空则删除所有自定义元素，否则删除指定的scene自定义元素。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DeleteAppCustomContentWithContext(ctx context.Context, request *DeleteAppCustomContentRequest) (response *DeleteAppCustomContentResponse, err error) {
    if request == nil {
        request = NewDeleteAppCustomContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocumentRequest() (request *DeleteDocumentRequest) {
    request = &DeleteDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteDocument")
    
    
    return
}

func NewDeleteDocumentResponse() (response *DeleteDocumentResponse) {
    response = &DeleteDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDocument
// 删除文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DeleteDocument(request *DeleteDocumentRequest) (response *DeleteDocumentResponse, err error) {
    return c.DeleteDocumentWithContext(context.Background(), request)
}

// DeleteDocument
// 删除文档
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DeleteDocumentWithContext(ctx context.Context, request *DeleteDocumentRequest) (response *DeleteDocumentResponse, err error) {
    if request == nil {
        request = NewDeleteDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// 此接口用于删除指定群组，支持批量操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// 此接口用于删除指定群组，支持批量操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupMemberRequest() (request *DeleteGroupMemberRequest) {
    request = &DeleteGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteGroupMember")
    
    
    return
}

func NewDeleteGroupMemberResponse() (response *DeleteGroupMemberResponse) {
    response = &DeleteGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroupMember
// 此接口用于删除群组中指定成员
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteGroupMember(request *DeleteGroupMemberRequest) (response *DeleteGroupMemberResponse, err error) {
    return c.DeleteGroupMemberWithContext(context.Background(), request)
}

// DeleteGroupMember
// 此接口用于删除群组中指定成员
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteGroupMemberWithContext(ctx context.Context, request *DeleteGroupMemberRequest) (response *DeleteGroupMemberResponse, err error) {
    if request == nil {
        request = NewDeleteGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
    request = &DeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteRecord")
    
    
    return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
    response = &DeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecord
// 删除指定房间的录制文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecord(request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    return c.DeleteRecordWithContext(context.Background(), request)
}

// DeleteRecord
// 删除指定房间的录制文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoomRequest() (request *DeleteRoomRequest) {
    request = &DeleteRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteRoom")
    
    
    return
}

func NewDeleteRoomResponse() (response *DeleteRoomResponse) {
    response = &DeleteRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoom
// 删除房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoom(request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    return c.DeleteRoomWithContext(context.Background(), request)
}

// DeleteRoom
// 删除房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoomWithContext(ctx context.Context, request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    if request == nil {
        request = NewDeleteRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppDetailRequest() (request *DescribeAppDetailRequest) {
    request = &DescribeAppDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeAppDetail")
    
    
    return
}

func NewDescribeAppDetailResponse() (response *DescribeAppDetailResponse) {
    response = &DescribeAppDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppDetail
// 获取应用详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAppDetail(request *DescribeAppDetailRequest) (response *DescribeAppDetailResponse, err error) {
    return c.DescribeAppDetailWithContext(context.Background(), request)
}

// DescribeAppDetail
// 获取应用详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAppDetailWithContext(ctx context.Context, request *DescribeAppDetailRequest) (response *DescribeAppDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAppDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentMemberListRequest() (request *DescribeCurrentMemberListRequest) {
    request = &DescribeCurrentMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeCurrentMemberList")
    
    
    return
}

func NewDescribeCurrentMemberListResponse() (response *DescribeCurrentMemberListResponse) {
    response = &DescribeCurrentMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCurrentMemberList
// 获取当前房间的成员列表，房间结束或过期后无法使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeCurrentMemberList(request *DescribeCurrentMemberListRequest) (response *DescribeCurrentMemberListResponse, err error) {
    return c.DescribeCurrentMemberListWithContext(context.Background(), request)
}

// DescribeCurrentMemberList
// 获取当前房间的成员列表，房间结束或过期后无法使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeCurrentMemberListWithContext(ctx context.Context, request *DescribeCurrentMemberListRequest) (response *DescribeCurrentMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentMemberListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentRequest() (request *DescribeDocumentRequest) {
    request = &DescribeDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocument")
    
    
    return
}

func NewDescribeDocumentResponse() (response *DescribeDocumentResponse) {
    response = &DescribeDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDocument
// 获取文档信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DescribeDocument(request *DescribeDocumentRequest) (response *DescribeDocumentResponse, err error) {
    return c.DescribeDocumentWithContext(context.Background(), request)
}

// DescribeDocument
// 获取文档信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DescribeDocumentWithContext(ctx context.Context, request *DescribeDocumentRequest) (response *DescribeDocumentResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentsByRoomRequest() (request *DescribeDocumentsByRoomRequest) {
    request = &DescribeDocumentsByRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocumentsByRoom")
    
    
    return
}

func NewDescribeDocumentsByRoomResponse() (response *DescribeDocumentsByRoomResponse) {
    response = &DescribeDocumentsByRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDocumentsByRoom
// 此接口获取指定房间下课件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDocumentsByRoom(request *DescribeDocumentsByRoomRequest) (response *DescribeDocumentsByRoomResponse, err error) {
    return c.DescribeDocumentsByRoomWithContext(context.Background(), request)
}

// DescribeDocumentsByRoom
// 此接口获取指定房间下课件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDocumentsByRoomWithContext(ctx context.Context, request *DescribeDocumentsByRoomRequest) (response *DescribeDocumentsByRoomResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentsByRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocumentsByRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentsByRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroup
// 此接口用于获取群组详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    return c.DescribeGroupWithContext(context.Background(), request)
}

// DescribeGroup
// 此接口用于获取群组详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupListRequest() (request *DescribeGroupListRequest) {
    request = &DescribeGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroupList")
    
    
    return
}

func NewDescribeGroupListResponse() (response *DescribeGroupListResponse) {
    response = &DescribeGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupList
// 获取群组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupList(request *DescribeGroupListRequest) (response *DescribeGroupListResponse, err error) {
    return c.DescribeGroupListWithContext(context.Background(), request)
}

// DescribeGroupList
// 获取群组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupListWithContext(ctx context.Context, request *DescribeGroupListRequest) (response *DescribeGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupMemberListRequest() (request *DescribeGroupMemberListRequest) {
    request = &DescribeGroupMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroupMemberList")
    
    
    return
}

func NewDescribeGroupMemberListResponse() (response *DescribeGroupMemberListResponse) {
    response = &DescribeGroupMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupMemberList
// 此接口用于获取群组成员列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupMemberList(request *DescribeGroupMemberListRequest) (response *DescribeGroupMemberListResponse, err error) {
    return c.DescribeGroupMemberListWithContext(context.Background(), request)
}

// DescribeGroupMemberList
// 此接口用于获取群组成员列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupMemberListWithContext(ctx context.Context, request *DescribeGroupMemberListRequest) (response *DescribeGroupMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupMemberListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomRequest() (request *DescribeRoomRequest) {
    request = &DescribeRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoom")
    
    
    return
}

func NewDescribeRoomResponse() (response *DescribeRoomResponse) {
    response = &DescribeRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoom
// 获取房间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoom(request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    return c.DescribeRoomWithContext(context.Background(), request)
}

// DescribeRoom
// 获取房间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoomWithContext(ctx context.Context, request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    if request == nil {
        request = NewDescribeRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomStatisticsRequest() (request *DescribeRoomStatisticsRequest) {
    request = &DescribeRoomStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoomStatistics")
    
    
    return
}

func NewDescribeRoomStatisticsResponse() (response *DescribeRoomStatisticsResponse) {
    response = &DescribeRoomStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoomStatistics
// 获取房间统计信息，仅可在房间结束后调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatistics(request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    return c.DescribeRoomStatisticsWithContext(context.Background(), request)
}

// DescribeRoomStatistics
// 获取房间统计信息，仅可在房间结束后调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatisticsWithContext(ctx context.Context, request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRoomStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSdkAppIdUsersRequest() (request *DescribeSdkAppIdUsersRequest) {
    request = &DescribeSdkAppIdUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeSdkAppIdUsers")
    
    
    return
}

func NewDescribeSdkAppIdUsersResponse() (response *DescribeSdkAppIdUsersResponse) {
    response = &DescribeSdkAppIdUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSdkAppIdUsers
// 此接口用于获取指定应用ID下用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSdkAppIdUsers(request *DescribeSdkAppIdUsersRequest) (response *DescribeSdkAppIdUsersResponse, err error) {
    return c.DescribeSdkAppIdUsersWithContext(context.Background(), request)
}

// DescribeSdkAppIdUsers
// 此接口用于获取指定应用ID下用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSdkAppIdUsersWithContext(ctx context.Context, request *DescribeSdkAppIdUsersRequest) (response *DescribeSdkAppIdUsersResponse, err error) {
    if request == nil {
        request = NewDescribeSdkAppIdUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSdkAppIdUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSdkAppIdUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomMessageRequest() (request *GetRoomMessageRequest) {
    request = &GetRoomMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRoomMessage")
    
    
    return
}

func NewGetRoomMessageResponse() (response *GetRoomMessageResponse) {
    response = &GetRoomMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRoomMessage
// 获取房间历史消息(房间历史消息保存7天)
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomMessage(request *GetRoomMessageRequest) (response *GetRoomMessageResponse, err error) {
    return c.GetRoomMessageWithContext(context.Background(), request)
}

// GetRoomMessage
// 获取房间历史消息(房间历史消息保存7天)
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomMessageWithContext(ctx context.Context, request *GetRoomMessageRequest) (response *GetRoomMessageResponse, err error) {
    if request == nil {
        request = NewGetRoomMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoomMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetWatermarkRequest() (request *GetWatermarkRequest) {
    request = &GetWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetWatermark")
    
    
    return
}

func NewGetWatermarkResponse() (response *GetWatermarkResponse) {
    response = &GetWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetWatermark
// 获取水印设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetWatermark(request *GetWatermarkRequest) (response *GetWatermarkResponse, err error) {
    return c.GetWatermarkWithContext(context.Background(), request)
}

// GetWatermark
// 获取水印设置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetWatermarkWithContext(ctx context.Context, request *GetWatermarkRequest) (response *GetWatermarkResponse, err error) {
    if request == nil {
        request = NewGetWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewLoginOriginIdRequest() (request *LoginOriginIdRequest) {
    request = &LoginOriginIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginOriginId")
    
    
    return
}

func NewLoginOriginIdResponse() (response *LoginOriginIdResponse) {
    response = &LoginOriginIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LoginOriginId
// 使用源账号登录，源账号为注册时填入的originId
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginId(request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    return c.LoginOriginIdWithContext(context.Background(), request)
}

// LoginOriginId
// 使用源账号登录，源账号为注册时填入的originId
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginIdWithContext(ctx context.Context, request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    if request == nil {
        request = NewLoginOriginIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginOriginId require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginOriginIdResponse()
    err = c.Send(request, response)
    return
}

func NewLoginUserRequest() (request *LoginUserRequest) {
    request = &LoginUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginUser")
    
    
    return
}

func NewLoginUserResponse() (response *LoginUserResponse) {
    response = &LoginUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LoginUser
// 登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUser(request *LoginUserRequest) (response *LoginUserResponse, err error) {
    return c.LoginUserWithContext(context.Background(), request)
}

// LoginUser
// 登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUserWithContext(ctx context.Context, request *LoginUserRequest) (response *LoginUserResponse, err error) {
    if request == nil {
        request = NewLoginUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyGroup")
    
    
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroup
// 此接口修改群组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    return c.ModifyGroupWithContext(context.Background(), request)
}

// ModifyGroup
// 此接口修改群组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyGroupWithContext(ctx context.Context, request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoomRequest() (request *ModifyRoomRequest) {
    request = &ModifyRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyRoom")
    
    
    return
}

func NewModifyRoomResponse() (response *ModifyRoomResponse) {
    response = &ModifyRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRoom
// 修改房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyRoom(request *ModifyRoomRequest) (response *ModifyRoomResponse, err error) {
    return c.ModifyRoomWithContext(context.Background(), request)
}

// ModifyRoom
// 修改房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyRoomWithContext(ctx context.Context, request *ModifyRoomRequest) (response *ModifyRoomResponse, err error) {
    if request == nil {
        request = NewModifyRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoomResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserProfileRequest() (request *ModifyUserProfileRequest) {
    request = &ModifyUserProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyUserProfile")
    
    
    return
}

func NewModifyUserProfileResponse() (response *ModifyUserProfileResponse) {
    response = &ModifyUserProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserProfile
// 此接口用于修改用户配置，如头像，昵称/用户名等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserProfile(request *ModifyUserProfileRequest) (response *ModifyUserProfileResponse, err error) {
    return c.ModifyUserProfileWithContext(context.Background(), request)
}

// ModifyUserProfile
// 此接口用于修改用户配置，如头像，昵称/用户名等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserProfileWithContext(ctx context.Context, request *ModifyUserProfileRequest) (response *ModifyUserProfileResponse, err error) {
    if request == nil {
        request = NewModifyUserProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserProfileResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterUserRequest() (request *RegisterUserRequest) {
    request = &RegisterUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "RegisterUser")
    
    
    return
}

func NewRegisterUserResponse() (response *RegisterUserResponse) {
    response = &RegisterUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterUser
// 注册用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUser(request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    return c.RegisterUserWithContext(context.Background(), request)
}

// RegisterUser
// 注册用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUserWithContext(ctx context.Context, request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    if request == nil {
        request = NewRegisterUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterUserResponse()
    err = c.Send(request, response)
    return
}

func NewSetAppCustomContentRequest() (request *SetAppCustomContentRequest) {
    request = &SetAppCustomContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetAppCustomContent")
    
    
    return
}

func NewSetAppCustomContentResponse() (response *SetAppCustomContentResponse) {
    response = &SetAppCustomContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetAppCustomContent
// 设置应用的自定义内容，包括应用图标，自定义的代码等。如果已存在，则为更新。更新js、css内容后，要生效也需要调用该接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContent(request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    return c.SetAppCustomContentWithContext(context.Background(), request)
}

// SetAppCustomContent
// 设置应用的自定义内容，包括应用图标，自定义的代码等。如果已存在，则为更新。更新js、css内容后，要生效也需要调用该接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContentWithContext(ctx context.Context, request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    if request == nil {
        request = NewSetAppCustomContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewSetWatermarkRequest() (request *SetWatermarkRequest) {
    request = &SetWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetWatermark")
    
    
    return
}

func NewSetWatermarkResponse() (response *SetWatermarkResponse) {
    response = &SetWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetWatermark
// 设置水印
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetWatermark(request *SetWatermarkRequest) (response *SetWatermarkResponse, err error) {
    return c.SetWatermarkWithContext(context.Background(), request)
}

// SetWatermark
// 设置水印
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetWatermarkWithContext(ctx context.Context, request *SetWatermarkRequest) (response *SetWatermarkResponse, err error) {
    if request == nil {
        request = NewSetWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDocumentFromRoomRequest() (request *UnbindDocumentFromRoomRequest) {
    request = &UnbindDocumentFromRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "UnbindDocumentFromRoom")
    
    
    return
}

func NewUnbindDocumentFromRoomResponse() (response *UnbindDocumentFromRoomResponse) {
    response = &UnbindDocumentFromRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindDocumentFromRoom
// 文档从房间解绑
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoom(request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    return c.UnbindDocumentFromRoomWithContext(context.Background(), request)
}

// UnbindDocumentFromRoom
// 文档从房间解绑
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoomWithContext(ctx context.Context, request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    if request == nil {
        request = NewUnbindDocumentFromRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDocumentFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindDocumentFromRoomResponse()
    err = c.Send(request, response)
    return
}
