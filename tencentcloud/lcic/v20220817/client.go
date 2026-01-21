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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "AddGroupMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchAddGroupMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchCreateGroupWithMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchCreateRoom")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchDeleteGroupMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchDeleteRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDescribeDocumentRequest() (request *BatchDescribeDocumentRequest) {
    request = &BatchDescribeDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDescribeDocument")
    
    
    return
}

func NewBatchDescribeDocumentResponse() (response *BatchDescribeDocumentResponse) {
    response = &BatchDescribeDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDescribeDocument
// 批量获取文档详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) BatchDescribeDocument(request *BatchDescribeDocumentRequest) (response *BatchDescribeDocumentResponse, err error) {
    return c.BatchDescribeDocumentWithContext(context.Background(), request)
}

// BatchDescribeDocument
// 批量获取文档详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) BatchDescribeDocumentWithContext(ctx context.Context, request *BatchDescribeDocumentRequest) (response *BatchDescribeDocumentResponse, err error) {
    if request == nil {
        request = NewBatchDescribeDocumentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchDescribeDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDescribeDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDescribeDocumentResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BatchRegister")
    
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
// 功能概述
//
// 本接口提供教学场景下的课程文档预绑定能力，支持将课件课堂ID进行关联映射，实现课件的自动化预加载。
//
// 
//
// 应用场景建议
//
// 
//
// 绑定时机
//
// 推荐在创建课堂预约阶段同步完成文档绑定操作，确保课件资源在课堂开始前完成上传。
//
// 
//
// 接口限制与频控策略
//
// 
//
// 频控维度
//
// 开发者账号维度限频
//
// 默认频控阈值为20 QPS（每秒请求次数）
//
// 
//
// 最佳实践建议
//
// 
//
// 重试策略
//
// 当触发限频错误时，建议采用以下策略：
//
// 启用退避重试机制（建议使用指数退避算法）
//
// 初始重试间隔不低于500ms
//
// 最大重试次数不超过3次
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoom(request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    return c.BindDocumentToRoomWithContext(context.Background(), request)
}

// BindDocumentToRoom
// 功能概述
//
// 本接口提供教学场景下的课程文档预绑定能力，支持将课件课堂ID进行关联映射，实现课件的自动化预加载。
//
// 
//
// 应用场景建议
//
// 
//
// 绑定时机
//
// 推荐在创建课堂预约阶段同步完成文档绑定操作，确保课件资源在课堂开始前完成上传。
//
// 
//
// 接口限制与频控策略
//
// 
//
// 频控维度
//
// 开发者账号维度限频
//
// 默认频控阈值为20 QPS（每秒请求次数）
//
// 
//
// 最佳实践建议
//
// 
//
// 重试策略
//
// 当触发限频错误时，建议采用以下策略：
//
// 启用退避重试机制（建议使用指数退避算法）
//
// 初始重试间隔不低于500ms
//
// 最大重试次数不超过3次
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoomWithContext(ctx context.Context, request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    if request == nil {
        request = NewBindDocumentToRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "BindDocumentToRoom")
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupLiveCodesRequest() (request *CreateGroupLiveCodesRequest) {
    request = &CreateGroupLiveCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateGroupLiveCodes")
    
    
    return
}

func NewCreateGroupLiveCodesResponse() (response *CreateGroupLiveCodesResponse) {
    response = &CreateGroupLiveCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroupLiveCodes
// 创建分组直播参加码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupLiveCodes(request *CreateGroupLiveCodesRequest) (response *CreateGroupLiveCodesResponse, err error) {
    return c.CreateGroupLiveCodesWithContext(context.Background(), request)
}

// CreateGroupLiveCodes
// 创建分组直播参加码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupLiveCodesWithContext(ctx context.Context, request *CreateGroupLiveCodesRequest) (response *CreateGroupLiveCodesResponse, err error) {
    if request == nil {
        request = NewCreateGroupLiveCodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateGroupLiveCodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupLiveCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupLiveCodesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateGroupWithMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateGroupWithSubGroup")
    
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
// 创建课堂
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_DISABLERECORD = "InvalidParameter.DisableRecord"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_LIVETYPE = "InvalidParameter.LiveType"
//  INVALIDPARAMETER_MAXMICNUMBER = "InvalidParameter.MaxMicNumber"
//  INVALIDPARAMETER_RECORDLAYOUT = "InvalidParameter.RecordLayout"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNSUPPORTEDOPERATION_VERSION = "UnsupportedOperation.Version"
func (c *Client) CreateRoom(request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    return c.CreateRoomWithContext(context.Background(), request)
}

// CreateRoom
// 创建课堂
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_DISABLERECORD = "InvalidParameter.DisableRecord"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_LIVETYPE = "InvalidParameter.LiveType"
//  INVALIDPARAMETER_MAXMICNUMBER = "InvalidParameter.MaxMicNumber"
//  INVALIDPARAMETER_RECORDLAYOUT = "InvalidParameter.RecordLayout"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNSUPPORTEDOPERATION_VERSION = "UnsupportedOperation.Version"
func (c *Client) CreateRoomWithContext(ctx context.Context, request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    if request == nil {
        request = NewCreateRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateRoom")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "CreateSupervisor")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteAppCustomContent")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteDocument")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteGroupMember")
    
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
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
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
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteRecord")
    
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
// 删除课堂前，请先删除该课堂下的各类资源（包括录制文件、板书等），并解绑相关课件。
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
// 删除课堂前，请先删除该课堂下的各类资源（包括录制文件、板书等），并解绑相关课件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoomWithContext(ctx context.Context, request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    if request == nil {
        request = NewDeleteRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSupervisorRequest() (request *DeleteSupervisorRequest) {
    request = &DeleteSupervisorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteSupervisor")
    
    
    return
}

func NewDeleteSupervisorResponse() (response *DeleteSupervisorResponse) {
    response = &DeleteSupervisorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSupervisor
// 删除巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DeleteSupervisor(request *DeleteSupervisorRequest) (response *DeleteSupervisorResponse, err error) {
    return c.DeleteSupervisorWithContext(context.Background(), request)
}

// DeleteSupervisor
// 删除巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DeleteSupervisorWithContext(ctx context.Context, request *DeleteSupervisorRequest) (response *DeleteSupervisorResponse, err error) {
    if request == nil {
        request = NewDeleteSupervisorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteSupervisor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSupervisor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSupervisorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// 删除已注册用户。注：如果该成员已被添加到群组，请先在群组中删除该成员。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  FAILEDOPERATION_USERISALREADYINGROUP = "FailedOperation.UserIsAlreadyInGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 删除已注册用户。注：如果该成员已被添加到群组，请先在群组中删除该成员。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  FAILEDOPERATION_USERISALREADYINGROUP = "FailedOperation.UserIsAlreadyInGroup"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhiteBoardSnapshotRequest() (request *DeleteWhiteBoardSnapshotRequest) {
    request = &DeleteWhiteBoardSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteWhiteBoardSnapshot")
    
    
    return
}

func NewDeleteWhiteBoardSnapshotResponse() (response *DeleteWhiteBoardSnapshotResponse) {
    response = &DeleteWhiteBoardSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWhiteBoardSnapshot
// 删除白板板书截图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteWhiteBoardSnapshot(request *DeleteWhiteBoardSnapshotRequest) (response *DeleteWhiteBoardSnapshotResponse, err error) {
    return c.DeleteWhiteBoardSnapshotWithContext(context.Background(), request)
}

// DeleteWhiteBoardSnapshot
// 删除白板板书截图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteWhiteBoardSnapshotWithContext(ctx context.Context, request *DeleteWhiteBoardSnapshotRequest) (response *DeleteWhiteBoardSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteWhiteBoardSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DeleteWhiteBoardSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWhiteBoardSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWhiteBoardSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnswerListRequest() (request *DescribeAnswerListRequest) {
    request = &DescribeAnswerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeAnswerList")
    
    
    return
}

func NewDescribeAnswerListResponse() (response *DescribeAnswerListResponse) {
    response = &DescribeAnswerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAnswerList
// 获取房间答题详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAnswerList(request *DescribeAnswerListRequest) (response *DescribeAnswerListResponse, err error) {
    return c.DescribeAnswerListWithContext(context.Background(), request)
}

// DescribeAnswerList
// 获取房间答题详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAnswerListWithContext(ctx context.Context, request *DescribeAnswerListRequest) (response *DescribeAnswerListResponse, err error) {
    if request == nil {
        request = NewDescribeAnswerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeAnswerList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnswerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnswerListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeAppDetail")
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeCurrentMemberListWithContext(ctx context.Context, request *DescribeCurrentMemberListRequest) (response *DescribeCurrentMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentMemberListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeCurrentMemberList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeveloperRequest() (request *DescribeDeveloperRequest) {
    request = &DescribeDeveloperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDeveloper")
    
    
    return
}

func NewDescribeDeveloperResponse() (response *DescribeDeveloperResponse) {
    response = &DescribeDeveloperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeveloper
// 开发商信息获取
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeveloper(request *DescribeDeveloperRequest) (response *DescribeDeveloperResponse, err error) {
    return c.DescribeDeveloperWithContext(context.Background(), request)
}

// DescribeDeveloper
// 开发商信息获取
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeveloperWithContext(ctx context.Context, request *DescribeDeveloperRequest) (response *DescribeDeveloperResponse, err error) {
    if request == nil {
        request = NewDescribeDeveloperRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeDeveloper")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeveloper require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeveloperResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentsRequest() (request *DescribeDocumentsRequest) {
    request = &DescribeDocumentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocuments")
    
    
    return
}

func NewDescribeDocumentsResponse() (response *DescribeDocumentsResponse) {
    response = &DescribeDocumentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDocuments
// 有新接口替换
//
// 
//
// 批量获取文档信息（已废弃，替代接口BatchDescribeDocument）
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeDocuments(request *DescribeDocumentsRequest) (response *DescribeDocumentsResponse, err error) {
    return c.DescribeDocumentsWithContext(context.Background(), request)
}

// DescribeDocuments
// 有新接口替换
//
// 
//
// 批量获取文档信息（已废弃，替代接口BatchDescribeDocument）
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeDocumentsWithContext(ctx context.Context, request *DescribeDocumentsRequest) (response *DescribeDocumentsResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeDocuments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocuments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeDocumentsByRoom")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeGroup")
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupLiveCodesRequest() (request *DescribeGroupLiveCodesRequest) {
    request = &DescribeGroupLiveCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroupLiveCodes")
    
    
    return
}

func NewDescribeGroupLiveCodesResponse() (response *DescribeGroupLiveCodesResponse) {
    response = &DescribeGroupLiveCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupLiveCodes
// 获取分组直播参加码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VERSION = "UnsupportedOperation.Version"
func (c *Client) DescribeGroupLiveCodes(request *DescribeGroupLiveCodesRequest) (response *DescribeGroupLiveCodesResponse, err error) {
    return c.DescribeGroupLiveCodesWithContext(context.Background(), request)
}

// DescribeGroupLiveCodes
// 获取分组直播参加码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VERSION = "UnsupportedOperation.Version"
func (c *Client) DescribeGroupLiveCodesWithContext(ctx context.Context, request *DescribeGroupLiveCodesRequest) (response *DescribeGroupLiveCodesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupLiveCodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeGroupLiveCodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupLiveCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupLiveCodesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeGroupMemberList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMarqueeRequest() (request *DescribeMarqueeRequest) {
    request = &DescribeMarqueeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeMarquee")
    
    
    return
}

func NewDescribeMarqueeResponse() (response *DescribeMarqueeResponse) {
    response = &DescribeMarqueeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMarquee
// 查询跑马灯配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMarquee(request *DescribeMarqueeRequest) (response *DescribeMarqueeResponse, err error) {
    return c.DescribeMarqueeWithContext(context.Background(), request)
}

// DescribeMarquee
// 查询跑马灯配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMarqueeWithContext(ctx context.Context, request *DescribeMarqueeRequest) (response *DescribeMarqueeResponse, err error) {
    if request == nil {
        request = NewDescribeMarqueeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeMarquee")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMarquee require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMarqueeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuestionListRequest() (request *DescribeQuestionListRequest) {
    request = &DescribeQuestionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeQuestionList")
    
    
    return
}

func NewDescribeQuestionListResponse() (response *DescribeQuestionListResponse) {
    response = &DescribeQuestionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuestionList
// 获取房间提问列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuestionList(request *DescribeQuestionListRequest) (response *DescribeQuestionListResponse, err error) {
    return c.DescribeQuestionListWithContext(context.Background(), request)
}

// DescribeQuestionList
// 获取房间提问列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuestionListWithContext(ctx context.Context, request *DescribeQuestionListRequest) (response *DescribeQuestionListResponse, err error) {
    if request == nil {
        request = NewDescribeQuestionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeQuestionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuestionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuestionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
    request = &DescribeRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRecord")
    
    
    return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
    response = &DescribeRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecord
// 查询录制信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecord(request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    return c.DescribeRecordWithContext(context.Background(), request)
}

// DescribeRecord
// 查询录制信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecordWithContext(ctx context.Context, request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordStreamRequest() (request *DescribeRecordStreamRequest) {
    request = &DescribeRecordStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRecordStream")
    
    
    return
}

func NewDescribeRecordStreamResponse() (response *DescribeRecordStreamResponse) {
    response = &DescribeRecordStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordStream
// 录制流查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRecordStream(request *DescribeRecordStreamRequest) (response *DescribeRecordStreamResponse, err error) {
    return c.DescribeRecordStreamWithContext(context.Background(), request)
}

// DescribeRecordStream
// 录制流查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRecordStreamWithContext(ctx context.Context, request *DescribeRecordStreamRequest) (response *DescribeRecordStreamResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRecordStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTaskRequest() (request *DescribeRecordTaskRequest) {
    request = &DescribeRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRecordTask")
    
    
    return
}

func NewDescribeRecordTaskResponse() (response *DescribeRecordTaskResponse) {
    response = &DescribeRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordTask
// 查询录制任务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTask(request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    return c.DescribeRecordTaskWithContext(context.Background(), request)
}

// DescribeRecordTask
// 查询录制任务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTaskWithContext(ctx context.Context, request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRecordTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTaskResponse()
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
// 获取房间配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoom(request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    return c.DescribeRoomWithContext(context.Background(), request)
}

// DescribeRoom
// 获取房间配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoomWithContext(ctx context.Context, request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    if request == nil {
        request = NewDescribeRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomForbiddenUserRequest() (request *DescribeRoomForbiddenUserRequest) {
    request = &DescribeRoomForbiddenUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoomForbiddenUser")
    
    
    return
}

func NewDescribeRoomForbiddenUserResponse() (response *DescribeRoomForbiddenUserResponse) {
    response = &DescribeRoomForbiddenUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomForbiddenUser
// 根据房间ID获取群组中被禁言的用户列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRoomForbiddenUser(request *DescribeRoomForbiddenUserRequest) (response *DescribeRoomForbiddenUserResponse, err error) {
    return c.DescribeRoomForbiddenUserWithContext(context.Background(), request)
}

// DescribeRoomForbiddenUser
// 根据房间ID获取群组中被禁言的用户列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRoomForbiddenUserWithContext(ctx context.Context, request *DescribeRoomForbiddenUserRequest) (response *DescribeRoomForbiddenUserResponse, err error) {
    if request == nil {
        request = NewDescribeRoomForbiddenUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRoomForbiddenUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomForbiddenUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomForbiddenUserResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatisticsWithContext(ctx context.Context, request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRoomStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeRoomStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScoreListRequest() (request *DescribeScoreListRequest) {
    request = &DescribeScoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeScoreList")
    
    
    return
}

func NewDescribeScoreListResponse() (response *DescribeScoreListResponse) {
    response = &DescribeScoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScoreList
// 获取课堂评分列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScoreList(request *DescribeScoreListRequest) (response *DescribeScoreListResponse, err error) {
    return c.DescribeScoreListWithContext(context.Background(), request)
}

// DescribeScoreList
// 获取课堂评分列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScoreListWithContext(ctx context.Context, request *DescribeScoreListRequest) (response *DescribeScoreListResponse, err error) {
    if request == nil {
        request = NewDescribeScoreListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeScoreList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScoreListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeSdkAppIdUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSdkAppIdUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSdkAppIdUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupervisorsRequest() (request *DescribeSupervisorsRequest) {
    request = &DescribeSupervisorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeSupervisors")
    
    
    return
}

func NewDescribeSupervisorsResponse() (response *DescribeSupervisorsResponse) {
    response = &DescribeSupervisorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupervisors
// 获取巡课列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupervisors(request *DescribeSupervisorsRequest) (response *DescribeSupervisorsResponse, err error) {
    return c.DescribeSupervisorsWithContext(context.Background(), request)
}

// DescribeSupervisors
// 获取巡课列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupervisorsWithContext(ctx context.Context, request *DescribeSupervisorsRequest) (response *DescribeSupervisorsResponse, err error) {
    if request == nil {
        request = NewDescribeSupervisorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeSupervisors")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupervisors require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupervisorsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDetailRequest() (request *DescribeUserDetailRequest) {
    request = &DescribeUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeUserDetail")
    
    
    return
}

func NewDescribeUserDetailResponse() (response *DescribeUserDetailResponse) {
    response = &DescribeUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDetail
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserDetail(request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    return c.DescribeUserDetailWithContext(context.Background(), request)
}

// DescribeUserDetail
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserDetailWithContext(ctx context.Context, request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUserDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeUserDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoardSnapshotRequest() (request *DescribeWhiteBoardSnapshotRequest) {
    request = &DescribeWhiteBoardSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeWhiteBoardSnapshot")
    
    
    return
}

func NewDescribeWhiteBoardSnapshotResponse() (response *DescribeWhiteBoardSnapshotResponse) {
    response = &DescribeWhiteBoardSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteBoardSnapshot
// 查询白板板书截图
//
// 课程结束后，可以查询和以图片的形式导出这些内容，方便后续查看、整理与分享。
//
// 注意：不支持屏幕共享中的板书导出。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWhiteBoardSnapshot(request *DescribeWhiteBoardSnapshotRequest) (response *DescribeWhiteBoardSnapshotResponse, err error) {
    return c.DescribeWhiteBoardSnapshotWithContext(context.Background(), request)
}

// DescribeWhiteBoardSnapshot
// 查询白板板书截图
//
// 课程结束后，可以查询和以图片的形式导出这些内容，方便后续查看、整理与分享。
//
// 注意：不支持屏幕共享中的板书导出。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWhiteBoardSnapshotWithContext(ctx context.Context, request *DescribeWhiteBoardSnapshotRequest) (response *DescribeWhiteBoardSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoardSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "DescribeWhiteBoardSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoardSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteBoardSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewEndRoomRequest() (request *EndRoomRequest) {
    request = &EndRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "EndRoom")
    
    
    return
}

func NewEndRoomResponse() (response *EndRoomResponse) {
    response = &EndRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EndRoom
// 结束房间的直播
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EndRoom(request *EndRoomRequest) (response *EndRoomResponse, err error) {
    return c.EndRoomWithContext(context.Background(), request)
}

// EndRoom
// 结束房间的直播
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EndRoomWithContext(ctx context.Context, request *EndRoomRequest) (response *EndRoomResponse, err error) {
    if request == nil {
        request = NewEndRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "EndRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EndRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewEndRoomResponse()
    err = c.Send(request, response)
    return
}

func NewForbidSendMsgRequest() (request *ForbidSendMsgRequest) {
    request = &ForbidSendMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ForbidSendMsg")
    
    
    return
}

func NewForbidSendMsgResponse() (response *ForbidSendMsgResponse) {
    response = &ForbidSendMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForbidSendMsg
// 禁止指定房间中某些用户在一段时间内发言。
//
// 取消对某些用户的禁言。
//
// 被禁言用户退出房间之后再进入同一房间，禁言仍然有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ForbidSendMsg(request *ForbidSendMsgRequest) (response *ForbidSendMsgResponse, err error) {
    return c.ForbidSendMsgWithContext(context.Background(), request)
}

// ForbidSendMsg
// 禁止指定房间中某些用户在一段时间内发言。
//
// 取消对某些用户的禁言。
//
// 被禁言用户退出房间之后再进入同一房间，禁言仍然有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ForbidSendMsgWithContext(ctx context.Context, request *ForbidSendMsgRequest) (response *ForbidSendMsgResponse, err error) {
    if request == nil {
        request = NewForbidSendMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "ForbidSendMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidSendMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewForbidSendMsgResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomEventRequest() (request *GetRoomEventRequest) {
    request = &GetRoomEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRoomEvent")
    
    
    return
}

func NewGetRoomEventResponse() (response *GetRoomEventResponse) {
    response = &GetRoomEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRoomEvent
// 获取房间事件,仅在课堂结束1小时内有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) GetRoomEvent(request *GetRoomEventRequest) (response *GetRoomEventResponse, err error) {
    return c.GetRoomEventWithContext(context.Background(), request)
}

// GetRoomEvent
// 获取房间事件,仅在课堂结束1小时内有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) GetRoomEventWithContext(ctx context.Context, request *GetRoomEventRequest) (response *GetRoomEventResponse, err error) {
    if request == nil {
        request = NewGetRoomEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "GetRoomEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoomEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomEventResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "GetRoomMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoomMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomsRequest() (request *GetRoomsRequest) {
    request = &GetRoomsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRooms")
    
    
    return
}

func NewGetRoomsResponse() (response *GetRoomsResponse) {
    response = &GetRoomsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRooms
// 获取房间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRooms(request *GetRoomsRequest) (response *GetRoomsResponse, err error) {
    return c.GetRoomsWithContext(context.Background(), request)
}

// GetRooms
// 获取房间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomsWithContext(ctx context.Context, request *GetRoomsRequest) (response *GetRoomsResponse, err error) {
    if request == nil {
        request = NewGetRoomsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "GetRooms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRooms require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "GetWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewKickUserFromRoomRequest() (request *KickUserFromRoomRequest) {
    request = &KickUserFromRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "KickUserFromRoom")
    
    
    return
}

func NewKickUserFromRoomResponse() (response *KickUserFromRoomResponse) {
    response = &KickUserFromRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KickUserFromRoom
// 从房间里面踢出用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) KickUserFromRoom(request *KickUserFromRoomRequest) (response *KickUserFromRoomResponse, err error) {
    return c.KickUserFromRoomWithContext(context.Background(), request)
}

// KickUserFromRoom
// 从房间里面踢出用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) KickUserFromRoomWithContext(ctx context.Context, request *KickUserFromRoomRequest) (response *KickUserFromRoomResponse, err error) {
    if request == nil {
        request = NewKickUserFromRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "KickUserFromRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KickUserFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewKickUserFromRoomResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "LoginOriginId")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "LoginUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "ModifyApp")
    
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "ModifyGroup")
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "ModifyRoom")
    
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
// 此接口用于修改用户信息，例如头像、昵称（用户名）等。注意，课中的用户信息不会立即同步修改，需待下次上课时，修改后的信息才会更新显示。
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
// 此接口用于修改用户信息，例如头像、昵称（用户名）等。注意，课中的用户信息不会立即同步修改，需待下次上课时，修改后的信息才会更新显示。
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "ModifyUserProfile")
    
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUserWithContext(ctx context.Context, request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    if request == nil {
        request = NewRegisterUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "RegisterUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterUserResponse()
    err = c.Send(request, response)
    return
}

func NewSendRoomNormalMessageRequest() (request *SendRoomNormalMessageRequest) {
    request = &SendRoomNormalMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SendRoomNormalMessage")
    
    
    return
}

func NewSendRoomNormalMessageResponse() (response *SendRoomNormalMessageResponse) {
    response = &SendRoomNormalMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendRoomNormalMessage
// 1、按照指定身份发送消息，目前支持表情消息、图片消息、文本消息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendRoomNormalMessage(request *SendRoomNormalMessageRequest) (response *SendRoomNormalMessageResponse, err error) {
    return c.SendRoomNormalMessageWithContext(context.Background(), request)
}

// SendRoomNormalMessage
// 1、按照指定身份发送消息，目前支持表情消息、图片消息、文本消息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendRoomNormalMessageWithContext(ctx context.Context, request *SendRoomNormalMessageRequest) (response *SendRoomNormalMessageResponse, err error) {
    if request == nil {
        request = NewSendRoomNormalMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "SendRoomNormalMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendRoomNormalMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendRoomNormalMessageResponse()
    err = c.Send(request, response)
    return
}

func NewSendRoomNotificationMessageRequest() (request *SendRoomNotificationMessageRequest) {
    request = &SendRoomNotificationMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SendRoomNotificationMessage")
    
    
    return
}

func NewSendRoomNotificationMessageResponse() (response *SendRoomNotificationMessageResponse) {
    response = &SendRoomNotificationMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendRoomNotificationMessage
// App 管理员可以通过该接口在群组中发送通知、公告等。目前仅支持文本消息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendRoomNotificationMessage(request *SendRoomNotificationMessageRequest) (response *SendRoomNotificationMessageResponse, err error) {
    return c.SendRoomNotificationMessageWithContext(context.Background(), request)
}

// SendRoomNotificationMessage
// App 管理员可以通过该接口在群组中发送通知、公告等。目前仅支持文本消息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendRoomNotificationMessageWithContext(ctx context.Context, request *SendRoomNotificationMessageRequest) (response *SendRoomNotificationMessageResponse, err error) {
    if request == nil {
        request = NewSendRoomNotificationMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "SendRoomNotificationMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendRoomNotificationMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendRoomNotificationMessageResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContent(request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    return c.SetAppCustomContentWithContext(context.Background(), request)
}

// SetAppCustomContent
// 设置应用的自定义内容，包括应用图标，自定义的代码等。如果已存在，则为更新。更新js、css内容后，要生效也需要调用该接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContentWithContext(ctx context.Context, request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    if request == nil {
        request = NewSetAppCustomContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "SetAppCustomContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewSetMarqueeRequest() (request *SetMarqueeRequest) {
    request = &SetMarqueeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetMarquee")
    
    
    return
}

func NewSetMarqueeResponse() (response *SetMarqueeResponse) {
    response = &SetMarqueeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetMarquee
// 设置跑马灯参数设置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) SetMarquee(request *SetMarqueeRequest) (response *SetMarqueeResponse, err error) {
    return c.SetMarqueeWithContext(context.Background(), request)
}

// SetMarquee
// 设置跑马灯参数设置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) SetMarqueeWithContext(ctx context.Context, request *SetMarqueeRequest) (response *SetMarqueeResponse, err error) {
    if request == nil {
        request = NewSetMarqueeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "SetMarquee")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetMarquee require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetMarqueeResponse()
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
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
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
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetWatermarkWithContext(ctx context.Context, request *SetWatermarkRequest) (response *SetWatermarkResponse, err error) {
    if request == nil {
        request = NewSetWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "SetWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewStartRecordRequest() (request *StartRecordRequest) {
    request = &StartRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "StartRecord")
    
    
    return
}

func NewStartRecordResponse() (response *StartRecordResponse) {
    response = &StartRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartRecord
// 开始录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) StartRecord(request *StartRecordRequest) (response *StartRecordResponse, err error) {
    return c.StartRecordWithContext(context.Background(), request)
}

// StartRecord
// 开始录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) StartRecordWithContext(ctx context.Context, request *StartRecordRequest) (response *StartRecordResponse, err error) {
    if request == nil {
        request = NewStartRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "StartRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStartRoomRequest() (request *StartRoomRequest) {
    request = &StartRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "StartRoom")
    
    
    return
}

func NewStartRoomResponse() (response *StartRoomResponse) {
    response = &StartRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartRoom
// 开始房间的直播。 说明：开始房间接口调用之前需要有用户进入课堂初始化课堂信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartRoom(request *StartRoomRequest) (response *StartRoomResponse, err error) {
    return c.StartRoomWithContext(context.Background(), request)
}

// StartRoom
// 开始房间的直播。 说明：开始房间接口调用之前需要有用户进入课堂初始化课堂信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartRoomWithContext(ctx context.Context, request *StartRoomRequest) (response *StartRoomResponse, err error) {
    if request == nil {
        request = NewStartRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "StartRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRoomResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordRequest() (request *StopRecordRequest) {
    request = &StopRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "StopRecord")
    
    
    return
}

func NewStopRecordResponse() (response *StopRecordResponse) {
    response = &StopRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopRecord
// 停止录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopRecord(request *StopRecordRequest) (response *StopRecordResponse, err error) {
    return c.StopRecordWithContext(context.Background(), request)
}

// StopRecord
// 停止录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) StopRecordWithContext(ctx context.Context, request *StopRecordRequest) (response *StopRecordResponse, err error) {
    if request == nil {
        request = NewStopRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "StopRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRecordResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "UnbindDocumentFromRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDocumentFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindDocumentFromRoomResponse()
    err = c.Send(request, response)
    return
}

func NewUnblockKickedUserRequest() (request *UnblockKickedUserRequest) {
    request = &UnblockKickedUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "UnblockKickedUser")
    
    
    return
}

func NewUnblockKickedUserResponse() (response *UnblockKickedUserResponse) {
    response = &UnblockKickedUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnblockKickedUser
// 解禁从房间里面踢出的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) UnblockKickedUser(request *UnblockKickedUserRequest) (response *UnblockKickedUserResponse, err error) {
    return c.UnblockKickedUserWithContext(context.Background(), request)
}

// UnblockKickedUser
// 解禁从房间里面踢出的用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) UnblockKickedUserWithContext(ctx context.Context, request *UnblockKickedUserRequest) (response *UnblockKickedUserResponse, err error) {
    if request == nil {
        request = NewUnblockKickedUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lcic", APIVersion, "UnblockKickedUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnblockKickedUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnblockKickedUserResponse()
    err = c.Send(request, response)
    return
}
