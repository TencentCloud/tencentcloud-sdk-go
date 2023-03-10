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

package v20210125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-25"

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


func NewAddDMSPartitionsRequest() (request *AddDMSPartitionsRequest) {
    request = &AddDMSPartitionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AddDMSPartitions")
    
    
    return
}

func NewAddDMSPartitionsResponse() (response *AddDMSPartitionsResponse) {
    response = &AddDMSPartitionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDMSPartitions
// DMS元数据新增分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddDMSPartitions(request *AddDMSPartitionsRequest) (response *AddDMSPartitionsResponse, err error) {
    return c.AddDMSPartitionsWithContext(context.Background(), request)
}

// AddDMSPartitions
// DMS元数据新增分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AddDMSPartitionsWithContext(ctx context.Context, request *AddDMSPartitionsRequest) (response *AddDMSPartitionsResponse, err error) {
    if request == nil {
        request = NewAddDMSPartitionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDMSPartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDMSPartitionsResponse()
    err = c.Send(request, response)
    return
}

func NewAddUsersToWorkGroupRequest() (request *AddUsersToWorkGroupRequest) {
    request = &AddUsersToWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AddUsersToWorkGroup")
    
    
    return
}

func NewAddUsersToWorkGroupResponse() (response *AddUsersToWorkGroupResponse) {
    response = &AddUsersToWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddUsersToWorkGroup
// 添加用户到工作组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"
func (c *Client) AddUsersToWorkGroup(request *AddUsersToWorkGroupRequest) (response *AddUsersToWorkGroupResponse, err error) {
    return c.AddUsersToWorkGroupWithContext(context.Background(), request)
}

// AddUsersToWorkGroup
// 添加用户到工作组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"
func (c *Client) AddUsersToWorkGroupWithContext(ctx context.Context, request *AddUsersToWorkGroupRequest) (response *AddUsersToWorkGroupResponse, err error) {
    if request == nil {
        request = NewAddUsersToWorkGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUsersToWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUsersToWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAlterDMSDatabaseRequest() (request *AlterDMSDatabaseRequest) {
    request = &AlterDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AlterDMSDatabase")
    
    
    return
}

func NewAlterDMSDatabaseResponse() (response *AlterDMSDatabaseResponse) {
    response = &AlterDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AlterDMSDatabase
// DMS元数据更新库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSDatabase(request *AlterDMSDatabaseRequest) (response *AlterDMSDatabaseResponse, err error) {
    return c.AlterDMSDatabaseWithContext(context.Background(), request)
}

// AlterDMSDatabase
// DMS元数据更新库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSDatabaseWithContext(ctx context.Context, request *AlterDMSDatabaseRequest) (response *AlterDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewAlterDMSDatabaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AlterDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewAlterDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewAlterDMSPartitionRequest() (request *AlterDMSPartitionRequest) {
    request = &AlterDMSPartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AlterDMSPartition")
    
    
    return
}

func NewAlterDMSPartitionResponse() (response *AlterDMSPartitionResponse) {
    response = &AlterDMSPartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AlterDMSPartition
// DMS元数据更新分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSPartition(request *AlterDMSPartitionRequest) (response *AlterDMSPartitionResponse, err error) {
    return c.AlterDMSPartitionWithContext(context.Background(), request)
}

// AlterDMSPartition
// DMS元数据更新分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSPartitionWithContext(ctx context.Context, request *AlterDMSPartitionRequest) (response *AlterDMSPartitionResponse, err error) {
    if request == nil {
        request = NewAlterDMSPartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AlterDMSPartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewAlterDMSPartitionResponse()
    err = c.Send(request, response)
    return
}

func NewAlterDMSTableRequest() (request *AlterDMSTableRequest) {
    request = &AlterDMSTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AlterDMSTable")
    
    
    return
}

func NewAlterDMSTableResponse() (response *AlterDMSTableResponse) {
    response = &AlterDMSTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AlterDMSTable
// DMS元数据更新表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSTable(request *AlterDMSTableRequest) (response *AlterDMSTableResponse, err error) {
    return c.AlterDMSTableWithContext(context.Background(), request)
}

// AlterDMSTable
// DMS元数据更新表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSTableWithContext(ctx context.Context, request *AlterDMSTableRequest) (response *AlterDMSTableResponse, err error) {
    if request == nil {
        request = NewAlterDMSTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AlterDMSTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewAlterDMSTableResponse()
    err = c.Send(request, response)
    return
}

func NewAttachUserPolicyRequest() (request *AttachUserPolicyRequest) {
    request = &AttachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AttachUserPolicy")
    
    
    return
}

func NewAttachUserPolicyResponse() (response *AttachUserPolicyResponse) {
    response = &AttachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachUserPolicy
// 绑定鉴权策略到用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    return c.AttachUserPolicyWithContext(context.Background(), request)
}

// AttachUserPolicy
// 绑定鉴权策略到用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) AttachUserPolicyWithContext(ctx context.Context, request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachWorkGroupPolicyRequest() (request *AttachWorkGroupPolicyRequest) {
    request = &AttachWorkGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AttachWorkGroupPolicy")
    
    
    return
}

func NewAttachWorkGroupPolicyResponse() (response *AttachWorkGroupPolicyResponse) {
    response = &AttachWorkGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachWorkGroupPolicy
// 绑定鉴权策略到工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AttachWorkGroupPolicy(request *AttachWorkGroupPolicyRequest) (response *AttachWorkGroupPolicyResponse, err error) {
    return c.AttachWorkGroupPolicyWithContext(context.Background(), request)
}

// AttachWorkGroupPolicy
// 绑定鉴权策略到工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AttachWorkGroupPolicyWithContext(ctx context.Context, request *AttachWorkGroupPolicyRequest) (response *AttachWorkGroupPolicyResponse, err error) {
    if request == nil {
        request = NewAttachWorkGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachWorkGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachWorkGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewBindWorkGroupsToUserRequest() (request *BindWorkGroupsToUserRequest) {
    request = &BindWorkGroupsToUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "BindWorkGroupsToUser")
    
    
    return
}

func NewBindWorkGroupsToUserResponse() (response *BindWorkGroupsToUserResponse) {
    response = &BindWorkGroupsToUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindWorkGroupsToUser
// 绑定工作组到用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) BindWorkGroupsToUser(request *BindWorkGroupsToUserRequest) (response *BindWorkGroupsToUserResponse, err error) {
    return c.BindWorkGroupsToUserWithContext(context.Background(), request)
}

// BindWorkGroupsToUser
// 绑定工作组到用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) BindWorkGroupsToUserWithContext(ctx context.Context, request *BindWorkGroupsToUserRequest) (response *BindWorkGroupsToUserResponse, err error) {
    if request == nil {
        request = NewBindWorkGroupsToUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindWorkGroupsToUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindWorkGroupsToUserResponse()
    err = c.Send(request, response)
    return
}

func NewCancelNotebookSessionStatementRequest() (request *CancelNotebookSessionStatementRequest) {
    request = &CancelNotebookSessionStatementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CancelNotebookSessionStatement")
    
    
    return
}

func NewCancelNotebookSessionStatementResponse() (response *CancelNotebookSessionStatementResponse) {
    response = &CancelNotebookSessionStatementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelNotebookSessionStatement
// 本接口（CancelNotebookSessionStatement）用于取消session statement
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelNotebookSessionStatement(request *CancelNotebookSessionStatementRequest) (response *CancelNotebookSessionStatementResponse, err error) {
    return c.CancelNotebookSessionStatementWithContext(context.Background(), request)
}

// CancelNotebookSessionStatement
// 本接口（CancelNotebookSessionStatement）用于取消session statement
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelNotebookSessionStatementWithContext(ctx context.Context, request *CancelNotebookSessionStatementRequest) (response *CancelNotebookSessionStatementResponse, err error) {
    if request == nil {
        request = NewCancelNotebookSessionStatementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelNotebookSessionStatement require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelNotebookSessionStatementResponse()
    err = c.Send(request, response)
    return
}

func NewCancelNotebookSessionStatementBatchRequest() (request *CancelNotebookSessionStatementBatchRequest) {
    request = &CancelNotebookSessionStatementBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CancelNotebookSessionStatementBatch")
    
    
    return
}

func NewCancelNotebookSessionStatementBatchResponse() (response *CancelNotebookSessionStatementBatchResponse) {
    response = &CancelNotebookSessionStatementBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelNotebookSessionStatementBatch
// 本接口（CancelNotebookSessionStatementBatch）用于按批取消Session statement。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelNotebookSessionStatementBatch(request *CancelNotebookSessionStatementBatchRequest) (response *CancelNotebookSessionStatementBatchResponse, err error) {
    return c.CancelNotebookSessionStatementBatchWithContext(context.Background(), request)
}

// CancelNotebookSessionStatementBatch
// 本接口（CancelNotebookSessionStatementBatch）用于按批取消Session statement。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelNotebookSessionStatementBatchWithContext(ctx context.Context, request *CancelNotebookSessionStatementBatchRequest) (response *CancelNotebookSessionStatementBatchResponse, err error) {
    if request == nil {
        request = NewCancelNotebookSessionStatementBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelNotebookSessionStatementBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelNotebookSessionStatementBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CancelTask")
    
    
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelTask
// 本接口（CancelTask），用于取消任务执行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    return c.CancelTaskWithContext(context.Background(), request)
}

// CancelTask
// 本接口（CancelTask），用于取消任务执行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCheckLockMetaDataRequest() (request *CheckLockMetaDataRequest) {
    request = &CheckLockMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CheckLockMetaData")
    
    
    return
}

func NewCheckLockMetaDataResponse() (response *CheckLockMetaDataResponse) {
    response = &CheckLockMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckLockMetaData
// 元数据锁检查
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckLockMetaData(request *CheckLockMetaDataRequest) (response *CheckLockMetaDataResponse, err error) {
    return c.CheckLockMetaDataWithContext(context.Background(), request)
}

// CheckLockMetaData
// 元数据锁检查
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckLockMetaDataWithContext(ctx context.Context, request *CheckLockMetaDataRequest) (response *CheckLockMetaDataResponse, err error) {
    if request == nil {
        request = NewCheckLockMetaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckLockMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckLockMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDMSDatabaseRequest() (request *CreateDMSDatabaseRequest) {
    request = &CreateDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDMSDatabase")
    
    
    return
}

func NewCreateDMSDatabaseResponse() (response *CreateDMSDatabaseResponse) {
    response = &CreateDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDMSDatabase
// DMS元数据创建库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSDatabase(request *CreateDMSDatabaseRequest) (response *CreateDMSDatabaseResponse, err error) {
    return c.CreateDMSDatabaseWithContext(context.Background(), request)
}

// CreateDMSDatabase
// DMS元数据创建库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSDatabaseWithContext(ctx context.Context, request *CreateDMSDatabaseRequest) (response *CreateDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateDMSDatabaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDMSTableRequest() (request *CreateDMSTableRequest) {
    request = &CreateDMSTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDMSTable")
    
    
    return
}

func NewCreateDMSTableResponse() (response *CreateDMSTableResponse) {
    response = &CreateDMSTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDMSTable
// DMS元数据创建表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSTable(request *CreateDMSTableRequest) (response *CreateDMSTableResponse, err error) {
    return c.CreateDMSTableWithContext(context.Background(), request)
}

// CreateDMSTable
// DMS元数据创建表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSTableWithContext(ctx context.Context, request *CreateDMSTableRequest) (response *CreateDMSTableResponse, err error) {
    if request == nil {
        request = NewCreateDMSTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDMSTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDMSTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataEngineRequest() (request *CreateDataEngineRequest) {
    request = &CreateDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDataEngine")
    
    
    return
}

func NewCreateDataEngineResponse() (response *CreateDataEngineResponse) {
    response = &CreateDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataEngine
// 为用户创建数据引擎
//
// 可能返回的错误码:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERCREATEPROCESSRUNNING = "FailedOperation.AnotherCreateProcessRunning"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_BINDTOOMANYTAGS = "FailedOperation.BindTooManyTags"
//  FAILEDOPERATION_CREATEDATAENGINEFAILED = "FailedOperation.CreateDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_DUPLICATETAGKEY = "FailedOperation.DuplicateTagKey"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_ILLEGALRESOURCE = "FailedOperation.IllegalResource"
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TAGALREADYATTACHED = "FailedOperation.TagAlreadyAttached"
//  FAILEDOPERATION_TAGKEYTOOLONG = "FailedOperation.TagKeyTooLong"
//  FAILEDOPERATION_TAGNOTEXIST = "FailedOperation.TagNotExist"
//  FAILEDOPERATION_TAGVALUETOOLONG = "FailedOperation.TagValueTooLong"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  FAILEDOPERATION_TOOMANYTAGS = "FailedOperation.TooManyTags"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) CreateDataEngine(request *CreateDataEngineRequest) (response *CreateDataEngineResponse, err error) {
    return c.CreateDataEngineWithContext(context.Background(), request)
}

// CreateDataEngine
// 为用户创建数据引擎
//
// 可能返回的错误码:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERCREATEPROCESSRUNNING = "FailedOperation.AnotherCreateProcessRunning"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_BINDTOOMANYTAGS = "FailedOperation.BindTooManyTags"
//  FAILEDOPERATION_CREATEDATAENGINEFAILED = "FailedOperation.CreateDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_DUPLICATETAGKEY = "FailedOperation.DuplicateTagKey"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_ILLEGALRESOURCE = "FailedOperation.IllegalResource"
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TAGALREADYATTACHED = "FailedOperation.TagAlreadyAttached"
//  FAILEDOPERATION_TAGKEYTOOLONG = "FailedOperation.TagKeyTooLong"
//  FAILEDOPERATION_TAGNOTEXIST = "FailedOperation.TagNotExist"
//  FAILEDOPERATION_TAGVALUETOOLONG = "FailedOperation.TagValueTooLong"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  FAILEDOPERATION_TOOMANYTAGS = "FailedOperation.TooManyTags"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) CreateDataEngineWithContext(ctx context.Context, request *CreateDataEngineRequest) (response *CreateDataEngineResponse, err error) {
    if request == nil {
        request = NewCreateDataEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatabaseRequest() (request *CreateDatabaseRequest) {
    request = &CreateDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDatabase")
    
    
    return
}

func NewCreateDatabaseResponse() (response *CreateDatabaseResponse) {
    response = &CreateDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDatabase
// 本接口（CreateDatabase）用于生成建库SQL语句。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDatabase(request *CreateDatabaseRequest) (response *CreateDatabaseResponse, err error) {
    return c.CreateDatabaseWithContext(context.Background(), request)
}

// CreateDatabase
// 本接口（CreateDatabase）用于生成建库SQL语句。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDatabaseWithContext(ctx context.Context, request *CreateDatabaseRequest) (response *CreateDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateDatabaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExportTaskRequest() (request *CreateExportTaskRequest) {
    request = &CreateExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateExportTask")
    
    
    return
}

func NewCreateExportTaskResponse() (response *CreateExportTaskResponse) {
    response = &CreateExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateExportTask
// 该接口（CreateExportTask）用于创建导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateExportTask(request *CreateExportTaskRequest) (response *CreateExportTaskResponse, err error) {
    return c.CreateExportTaskWithContext(context.Background(), request)
}

// CreateExportTask
// 该接口（CreateExportTask）用于创建导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateExportTaskWithContext(ctx context.Context, request *CreateExportTaskRequest) (response *CreateExportTaskResponse, err error) {
    if request == nil {
        request = NewCreateExportTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImportTaskRequest() (request *CreateImportTaskRequest) {
    request = &CreateImportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateImportTask")
    
    
    return
}

func NewCreateImportTaskResponse() (response *CreateImportTaskResponse) {
    response = &CreateImportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImportTask
// 该接口（CreateImportTask）用于创建导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateImportTask(request *CreateImportTaskRequest) (response *CreateImportTaskResponse, err error) {
    return c.CreateImportTaskWithContext(context.Background(), request)
}

// CreateImportTask
// 该接口（CreateImportTask）用于创建导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateImportTaskWithContext(ctx context.Context, request *CreateImportTaskRequest) (response *CreateImportTaskResponse, err error) {
    if request == nil {
        request = NewCreateImportTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInternalTableRequest() (request *CreateInternalTableRequest) {
    request = &CreateInternalTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateInternalTable")
    
    
    return
}

func NewCreateInternalTableResponse() (response *CreateInternalTableResponse) {
    response = &CreateInternalTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInternalTable
// 创建托管存储内表（该接口已废弃）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInternalTable(request *CreateInternalTableRequest) (response *CreateInternalTableResponse, err error) {
    return c.CreateInternalTableWithContext(context.Background(), request)
}

// CreateInternalTable
// 创建托管存储内表（该接口已废弃）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInternalTableWithContext(ctx context.Context, request *CreateInternalTableRequest) (response *CreateInternalTableResponse, err error) {
    if request == nil {
        request = NewCreateInternalTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInternalTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInternalTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookSessionRequest() (request *CreateNotebookSessionRequest) {
    request = &CreateNotebookSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateNotebookSession")
    
    
    return
}

func NewCreateNotebookSessionResponse() (response *CreateNotebookSessionResponse) {
    response = &CreateNotebookSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookSession
// 本接口（CreateNotebookSession）用于创建notebook livy session
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDCODE_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.ResourceNotFoundCode_SessionInsufficientResources"
func (c *Client) CreateNotebookSession(request *CreateNotebookSessionRequest) (response *CreateNotebookSessionResponse, err error) {
    return c.CreateNotebookSessionWithContext(context.Background(), request)
}

// CreateNotebookSession
// 本接口（CreateNotebookSession）用于创建notebook livy session
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDCODE_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.ResourceNotFoundCode_SessionInsufficientResources"
func (c *Client) CreateNotebookSessionWithContext(ctx context.Context, request *CreateNotebookSessionRequest) (response *CreateNotebookSessionResponse, err error) {
    if request == nil {
        request = NewCreateNotebookSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookSessionStatementRequest() (request *CreateNotebookSessionStatementRequest) {
    request = &CreateNotebookSessionStatementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateNotebookSessionStatement")
    
    
    return
}

func NewCreateNotebookSessionStatementResponse() (response *CreateNotebookSessionStatementResponse) {
    response = &CreateNotebookSessionStatementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookSessionStatement
// 本接口（CreateNotebookSessionStatement）用于创建session statement
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateNotebookSessionStatement(request *CreateNotebookSessionStatementRequest) (response *CreateNotebookSessionStatementResponse, err error) {
    return c.CreateNotebookSessionStatementWithContext(context.Background(), request)
}

// CreateNotebookSessionStatement
// 本接口（CreateNotebookSessionStatement）用于创建session statement
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateNotebookSessionStatementWithContext(ctx context.Context, request *CreateNotebookSessionStatementRequest) (response *CreateNotebookSessionStatementResponse, err error) {
    if request == nil {
        request = NewCreateNotebookSessionStatementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookSessionStatement require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookSessionStatementResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookSessionStatementSupportBatchSQLRequest() (request *CreateNotebookSessionStatementSupportBatchSQLRequest) {
    request = &CreateNotebookSessionStatementSupportBatchSQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateNotebookSessionStatementSupportBatchSQL")
    
    
    return
}

func NewCreateNotebookSessionStatementSupportBatchSQLResponse() (response *CreateNotebookSessionStatementSupportBatchSQLResponse) {
    response = &CreateNotebookSessionStatementSupportBatchSQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookSessionStatementSupportBatchSQL
// 本接口（CreateNotebookSessionStatementSupportBatchSQL）用于创建Statement批量运行SQL任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateNotebookSessionStatementSupportBatchSQL(request *CreateNotebookSessionStatementSupportBatchSQLRequest) (response *CreateNotebookSessionStatementSupportBatchSQLResponse, err error) {
    return c.CreateNotebookSessionStatementSupportBatchSQLWithContext(context.Background(), request)
}

// CreateNotebookSessionStatementSupportBatchSQL
// 本接口（CreateNotebookSessionStatementSupportBatchSQL）用于创建Statement批量运行SQL任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateNotebookSessionStatementSupportBatchSQLWithContext(ctx context.Context, request *CreateNotebookSessionStatementSupportBatchSQLRequest) (response *CreateNotebookSessionStatementSupportBatchSQLResponse, err error) {
    if request == nil {
        request = NewCreateNotebookSessionStatementSupportBatchSQLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookSessionStatementSupportBatchSQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookSessionStatementSupportBatchSQLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResultDownloadRequest() (request *CreateResultDownloadRequest) {
    request = &CreateResultDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateResultDownload")
    
    
    return
}

func NewCreateResultDownloadResponse() (response *CreateResultDownloadResponse) {
    response = &CreateResultDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResultDownload
// 创建查询结果下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
func (c *Client) CreateResultDownload(request *CreateResultDownloadRequest) (response *CreateResultDownloadResponse, err error) {
    return c.CreateResultDownloadWithContext(context.Background(), request)
}

// CreateResultDownload
// 创建查询结果下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
func (c *Client) CreateResultDownloadWithContext(ctx context.Context, request *CreateResultDownloadRequest) (response *CreateResultDownloadResponse, err error) {
    if request == nil {
        request = NewCreateResultDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResultDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResultDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScriptRequest() (request *CreateScriptRequest) {
    request = &CreateScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateScript")
    
    
    return
}

func NewCreateScriptResponse() (response *CreateScriptResponse) {
    response = &CreateScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScript
// 该接口（CreateScript）用于创建sql脚本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateScript(request *CreateScriptRequest) (response *CreateScriptResponse, err error) {
    return c.CreateScriptWithContext(context.Background(), request)
}

// CreateScript
// 该接口（CreateScript）用于创建sql脚本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateScriptWithContext(ctx context.Context, request *CreateScriptRequest) (response *CreateScriptResponse, err error) {
    if request == nil {
        request = NewCreateScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScriptResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSparkAppRequest() (request *CreateSparkAppRequest) {
    request = &CreateSparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateSparkApp")
    
    
    return
}

func NewCreateSparkAppResponse() (response *CreateSparkAppResponse) {
    response = &CreateSparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSparkApp
// 创建spark应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
func (c *Client) CreateSparkApp(request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    return c.CreateSparkAppWithContext(context.Background(), request)
}

// CreateSparkApp
// 创建spark应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
func (c *Client) CreateSparkAppWithContext(ctx context.Context, request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSparkAppTaskRequest() (request *CreateSparkAppTaskRequest) {
    request = &CreateSparkAppTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateSparkAppTask")
    
    
    return
}

func NewCreateSparkAppTaskResponse() (response *CreateSparkAppTaskResponse) {
    response = &CreateSparkAppTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSparkAppTask
// 创建spark任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTask(request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    return c.CreateSparkAppTaskWithContext(context.Background(), request)
}

// CreateSparkAppTask
// 创建spark任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTaskWithContext(ctx context.Context, request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkAppTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkAppTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStoreLocationRequest() (request *CreateStoreLocationRequest) {
    request = &CreateStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateStoreLocation")
    
    
    return
}

func NewCreateStoreLocationResponse() (response *CreateStoreLocationResponse) {
    response = &CreateStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStoreLocation
// 该接口（CreateStoreLocation）新增或覆盖计算结果存储位置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateStoreLocation(request *CreateStoreLocationRequest) (response *CreateStoreLocationResponse, err error) {
    return c.CreateStoreLocationWithContext(context.Background(), request)
}

// CreateStoreLocation
// 该接口（CreateStoreLocation）新增或覆盖计算结果存储位置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateStoreLocationWithContext(ctx context.Context, request *CreateStoreLocationRequest) (response *CreateStoreLocationResponse, err error) {
    if request == nil {
        request = NewCreateStoreLocationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableRequest() (request *CreateTableRequest) {
    request = &CreateTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTable")
    
    
    return
}

func NewCreateTableResponse() (response *CreateTableResponse) {
    response = &CreateTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTable
// 本接口（CreateTable）用于生成建表SQL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTable(request *CreateTableRequest) (response *CreateTableResponse, err error) {
    return c.CreateTableWithContext(context.Background(), request)
}

// CreateTable
// 本接口（CreateTable）用于生成建表SQL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTableWithContext(ctx context.Context, request *CreateTableRequest) (response *CreateTableResponse, err error) {
    if request == nil {
        request = NewCreateTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// 本接口（CreateTask）用于创建sql查询任务。（推荐使用CreateTasks接口）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// 本接口（CreateTask）用于创建sql查询任务。（推荐使用CreateTasks接口）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTasksRequest() (request *CreateTasksRequest) {
    request = &CreateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTasks")
    
    
    return
}

func NewCreateTasksResponse() (response *CreateTasksResponse) {
    response = &CreateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTasks
// 本接口（CreateTasks），用于批量创建任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTasks(request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    return c.CreateTasksWithContext(context.Background(), request)
}

// CreateTasks
// 本接口（CreateTasks），用于批量创建任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTasksWithContext(ctx context.Context, request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    if request == nil {
        request = NewCreateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTasksInOrderRequest() (request *CreateTasksInOrderRequest) {
    request = &CreateTasksInOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTasksInOrder")
    
    
    return
}

func NewCreateTasksInOrderResponse() (response *CreateTasksInOrderResponse) {
    response = &CreateTasksInOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTasksInOrder
// 按顺序创建任务（已经废弃，后期不再维护，请使用接口CreateTasks）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateTasksInOrder(request *CreateTasksInOrderRequest) (response *CreateTasksInOrderResponse, err error) {
    return c.CreateTasksInOrderWithContext(context.Background(), request)
}

// CreateTasksInOrder
// 按顺序创建任务（已经废弃，后期不再维护，请使用接口CreateTasks）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateTasksInOrderWithContext(ctx context.Context, request *CreateTasksInOrderRequest) (response *CreateTasksInOrderResponse, err error) {
    if request == nil {
        request = NewCreateTasksInOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTasksInOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTasksInOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkGroupRequest() (request *CreateWorkGroupRequest) {
    request = &CreateWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateWorkGroup")
    
    
    return
}

func NewCreateWorkGroupResponse() (response *CreateWorkGroupResponse) {
    response = &CreateWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWorkGroup
// 创建工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateWorkGroup(request *CreateWorkGroupRequest) (response *CreateWorkGroupResponse, err error) {
    return c.CreateWorkGroupWithContext(context.Background(), request)
}

// CreateWorkGroup
// 创建工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateWorkGroupWithContext(ctx context.Context, request *CreateWorkGroupRequest) (response *CreateWorkGroupResponse, err error) {
    if request == nil {
        request = NewCreateWorkGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookSessionRequest() (request *DeleteNotebookSessionRequest) {
    request = &DeleteNotebookSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteNotebookSession")
    
    
    return
}

func NewDeleteNotebookSessionResponse() (response *DeleteNotebookSessionResponse) {
    response = &DeleteNotebookSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotebookSession
// 本接口（DeleteNotebookSession）用于删除notebook livy session
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DeleteNotebookSession(request *DeleteNotebookSessionRequest) (response *DeleteNotebookSessionResponse, err error) {
    return c.DeleteNotebookSessionWithContext(context.Background(), request)
}

// DeleteNotebookSession
// 本接口（DeleteNotebookSession）用于删除notebook livy session
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DeleteNotebookSessionWithContext(ctx context.Context, request *DeleteNotebookSessionRequest) (response *DeleteNotebookSessionResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebookSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScriptRequest() (request *DeleteScriptRequest) {
    request = &DeleteScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteScript")
    
    
    return
}

func NewDeleteScriptResponse() (response *DeleteScriptResponse) {
    response = &DeleteScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScript
// 该接口（DeleteScript）用于删除sql脚本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteScript(request *DeleteScriptRequest) (response *DeleteScriptResponse, err error) {
    return c.DeleteScriptWithContext(context.Background(), request)
}

// DeleteScript
// 该接口（DeleteScript）用于删除sql脚本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteScriptWithContext(ctx context.Context, request *DeleteScriptRequest) (response *DeleteScriptResponse, err error) {
    if request == nil {
        request = NewDeleteScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSparkAppRequest() (request *DeleteSparkAppRequest) {
    request = &DeleteSparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteSparkApp")
    
    
    return
}

func NewDeleteSparkAppResponse() (response *DeleteSparkAppResponse) {
    response = &DeleteSparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSparkApp
// 删除spark应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSparkApp(request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    return c.DeleteSparkAppWithContext(context.Background(), request)
}

// DeleteSparkApp
// 删除spark应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSparkAppWithContext(ctx context.Context, request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    if request == nil {
        request = NewDeleteSparkAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"
//  UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"
//  UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersFromWorkGroupRequest() (request *DeleteUsersFromWorkGroupRequest) {
    request = &DeleteUsersFromWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteUsersFromWorkGroup")
    
    
    return
}

func NewDeleteUsersFromWorkGroupResponse() (response *DeleteUsersFromWorkGroupResponse) {
    response = &DeleteUsersFromWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsersFromWorkGroup
// 从工作组中删除用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"
func (c *Client) DeleteUsersFromWorkGroup(request *DeleteUsersFromWorkGroupRequest) (response *DeleteUsersFromWorkGroupResponse, err error) {
    return c.DeleteUsersFromWorkGroupWithContext(context.Background(), request)
}

// DeleteUsersFromWorkGroup
// 从工作组中删除用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"
func (c *Client) DeleteUsersFromWorkGroupWithContext(ctx context.Context, request *DeleteUsersFromWorkGroupRequest) (response *DeleteUsersFromWorkGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUsersFromWorkGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsersFromWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersFromWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkGroupRequest() (request *DeleteWorkGroupRequest) {
    request = &DeleteWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteWorkGroup")
    
    
    return
}

func NewDeleteWorkGroupResponse() (response *DeleteWorkGroupResponse) {
    response = &DeleteWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWorkGroup
// 删除工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DeleteWorkGroup(request *DeleteWorkGroupRequest) (response *DeleteWorkGroupResponse, err error) {
    return c.DeleteWorkGroupWithContext(context.Background(), request)
}

// DeleteWorkGroup
// 删除工作组
//
// 可能返回的错误码:
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DeleteWorkGroupWithContext(ctx context.Context, request *DeleteWorkGroupRequest) (response *DeleteWorkGroupResponse, err error) {
    if request == nil {
        request = NewDeleteWorkGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSDatabaseRequest() (request *DescribeDMSDatabaseRequest) {
    request = &DescribeDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSDatabase")
    
    
    return
}

func NewDescribeDMSDatabaseResponse() (response *DescribeDMSDatabaseResponse) {
    response = &DescribeDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDMSDatabase
// DMS元数据获取库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabase(request *DescribeDMSDatabaseRequest) (response *DescribeDMSDatabaseResponse, err error) {
    return c.DescribeDMSDatabaseWithContext(context.Background(), request)
}

// DescribeDMSDatabase
// DMS元数据获取库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabaseWithContext(ctx context.Context, request *DescribeDMSDatabaseRequest) (response *DescribeDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewDescribeDMSDatabaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSPartitionsRequest() (request *DescribeDMSPartitionsRequest) {
    request = &DescribeDMSPartitionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSPartitions")
    
    
    return
}

func NewDescribeDMSPartitionsResponse() (response *DescribeDMSPartitionsResponse) {
    response = &DescribeDMSPartitionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDMSPartitions
// DMS元数据获取分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSPartitions(request *DescribeDMSPartitionsRequest) (response *DescribeDMSPartitionsResponse, err error) {
    return c.DescribeDMSPartitionsWithContext(context.Background(), request)
}

// DescribeDMSPartitions
// DMS元数据获取分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSPartitionsWithContext(ctx context.Context, request *DescribeDMSPartitionsRequest) (response *DescribeDMSPartitionsResponse, err error) {
    if request == nil {
        request = NewDescribeDMSPartitionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSPartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSPartitionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSTableRequest() (request *DescribeDMSTableRequest) {
    request = &DescribeDMSTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSTable")
    
    
    return
}

func NewDescribeDMSTableResponse() (response *DescribeDMSTableResponse) {
    response = &DescribeDMSTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDMSTable
// DMS元数据获取表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSTable(request *DescribeDMSTableRequest) (response *DescribeDMSTableResponse, err error) {
    return c.DescribeDMSTableWithContext(context.Background(), request)
}

// DescribeDMSTable
// DMS元数据获取表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSTableWithContext(ctx context.Context, request *DescribeDMSTableRequest) (response *DescribeDMSTableResponse, err error) {
    if request == nil {
        request = NewDescribeDMSTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSTablesRequest() (request *DescribeDMSTablesRequest) {
    request = &DescribeDMSTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSTables")
    
    
    return
}

func NewDescribeDMSTablesResponse() (response *DescribeDMSTablesResponse) {
    response = &DescribeDMSTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDMSTables
// DMS元数据获取表列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSTables(request *DescribeDMSTablesRequest) (response *DescribeDMSTablesResponse, err error) {
    return c.DescribeDMSTablesWithContext(context.Background(), request)
}

// DescribeDMSTables
// DMS元数据获取表列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSTablesWithContext(ctx context.Context, request *DescribeDMSTablesRequest) (response *DescribeDMSTablesResponse, err error) {
    if request == nil {
        request = NewDescribeDMSTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEnginesRequest() (request *DescribeDataEnginesRequest) {
    request = &DescribeDataEnginesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDataEngines")
    
    
    return
}

func NewDescribeDataEnginesResponse() (response *DescribeDataEnginesResponse) {
    response = &DescribeDataEnginesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataEngines
// 本接口（DescribeDataEngines）用于获取DataEngines信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataEngines(request *DescribeDataEnginesRequest) (response *DescribeDataEnginesResponse, err error) {
    return c.DescribeDataEnginesWithContext(context.Background(), request)
}

// DescribeDataEngines
// 本接口（DescribeDataEngines）用于获取DataEngines信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataEnginesWithContext(ctx context.Context, request *DescribeDataEnginesRequest) (response *DescribeDataEnginesResponse, err error) {
    if request == nil {
        request = NewDescribeDataEnginesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEngines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEnginesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabases
// 本接口（DescribeDatabases）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// 本接口（DescribeDatabases）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEngineUsageInfoRequest() (request *DescribeEngineUsageInfoRequest) {
    request = &DescribeEngineUsageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeEngineUsageInfo")
    
    
    return
}

func NewDescribeEngineUsageInfoResponse() (response *DescribeEngineUsageInfoResponse) {
    response = &DescribeEngineUsageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEngineUsageInfo
// 本接口根据引擎ID获取数据引擎资源使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEngineUsageInfo(request *DescribeEngineUsageInfoRequest) (response *DescribeEngineUsageInfoResponse, err error) {
    return c.DescribeEngineUsageInfoWithContext(context.Background(), request)
}

// DescribeEngineUsageInfo
// 本接口根据引擎ID获取数据引擎资源使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEngineUsageInfoWithContext(ctx context.Context, request *DescribeEngineUsageInfoRequest) (response *DescribeEngineUsageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEngineUsageInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEngineUsageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEngineUsageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForbiddenTableProRequest() (request *DescribeForbiddenTableProRequest) {
    request = &DescribeForbiddenTableProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeForbiddenTablePro")
    
    
    return
}

func NewDescribeForbiddenTableProResponse() (response *DescribeForbiddenTableProResponse) {
    response = &DescribeForbiddenTableProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeForbiddenTablePro
// 本接口（DescribeForbiddenTablePro）用于获取被禁用的表属性列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForbiddenTablePro(request *DescribeForbiddenTableProRequest) (response *DescribeForbiddenTableProResponse, err error) {
    return c.DescribeForbiddenTableProWithContext(context.Background(), request)
}

// DescribeForbiddenTablePro
// 本接口（DescribeForbiddenTablePro）用于获取被禁用的表属性列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForbiddenTableProWithContext(ctx context.Context, request *DescribeForbiddenTableProRequest) (response *DescribeForbiddenTableProResponse, err error) {
    if request == nil {
        request = NewDescribeForbiddenTableProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForbiddenTablePro require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForbiddenTableProResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLakeFsDirSummaryRequest() (request *DescribeLakeFsDirSummaryRequest) {
    request = &DescribeLakeFsDirSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeLakeFsDirSummary")
    
    
    return
}

func NewDescribeLakeFsDirSummaryResponse() (response *DescribeLakeFsDirSummaryResponse) {
    response = &DescribeLakeFsDirSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLakeFsDirSummary
// 查询托管存储指定目录的Summary
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsDirSummary(request *DescribeLakeFsDirSummaryRequest) (response *DescribeLakeFsDirSummaryResponse, err error) {
    return c.DescribeLakeFsDirSummaryWithContext(context.Background(), request)
}

// DescribeLakeFsDirSummary
// 查询托管存储指定目录的Summary
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsDirSummaryWithContext(ctx context.Context, request *DescribeLakeFsDirSummaryRequest) (response *DescribeLakeFsDirSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeLakeFsDirSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLakeFsDirSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLakeFsDirSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLakeFsInfoRequest() (request *DescribeLakeFsInfoRequest) {
    request = &DescribeLakeFsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeLakeFsInfo")
    
    
    return
}

func NewDescribeLakeFsInfoResponse() (response *DescribeLakeFsInfoResponse) {
    response = &DescribeLakeFsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLakeFsInfo
// 查询用户的托管存储信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsInfo(request *DescribeLakeFsInfoRequest) (response *DescribeLakeFsInfoResponse, err error) {
    return c.DescribeLakeFsInfoWithContext(context.Background(), request)
}

// DescribeLakeFsInfo
// 查询用户的托管存储信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsInfoWithContext(ctx context.Context, request *DescribeLakeFsInfoRequest) (response *DescribeLakeFsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLakeFsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLakeFsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLakeFsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionRequest() (request *DescribeNotebookSessionRequest) {
    request = &DescribeNotebookSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSession")
    
    
    return
}

func NewDescribeNotebookSessionResponse() (response *DescribeNotebookSessionResponse) {
    response = &DescribeNotebookSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSession
// 本接口（DescribeNotebookSession）用于获取notebook livy session详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSession(request *DescribeNotebookSessionRequest) (response *DescribeNotebookSessionResponse, err error) {
    return c.DescribeNotebookSessionWithContext(context.Background(), request)
}

// DescribeNotebookSession
// 本接口（DescribeNotebookSession）用于获取notebook livy session详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionWithContext(ctx context.Context, request *DescribeNotebookSessionRequest) (response *DescribeNotebookSessionResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionLogRequest() (request *DescribeNotebookSessionLogRequest) {
    request = &DescribeNotebookSessionLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSessionLog")
    
    
    return
}

func NewDescribeNotebookSessionLogResponse() (response *DescribeNotebookSessionLogResponse) {
    response = &DescribeNotebookSessionLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSessionLog
// 本接口（DescribeNotebookSessionLog）用于获取notebook livy session日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionLog(request *DescribeNotebookSessionLogRequest) (response *DescribeNotebookSessionLogResponse, err error) {
    return c.DescribeNotebookSessionLogWithContext(context.Background(), request)
}

// DescribeNotebookSessionLog
// 本接口（DescribeNotebookSessionLog）用于获取notebook livy session日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionLogWithContext(ctx context.Context, request *DescribeNotebookSessionLogRequest) (response *DescribeNotebookSessionLogResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSessionLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionStatementRequest() (request *DescribeNotebookSessionStatementRequest) {
    request = &DescribeNotebookSessionStatementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSessionStatement")
    
    
    return
}

func NewDescribeNotebookSessionStatementResponse() (response *DescribeNotebookSessionStatementResponse) {
    response = &DescribeNotebookSessionStatementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSessionStatement
// 本接口（DescribeNotebookSessionStatement）用于获取session statement信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionStatement(request *DescribeNotebookSessionStatementRequest) (response *DescribeNotebookSessionStatementResponse, err error) {
    return c.DescribeNotebookSessionStatementWithContext(context.Background(), request)
}

// DescribeNotebookSessionStatement
// 本接口（DescribeNotebookSessionStatement）用于获取session statement信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionStatementWithContext(ctx context.Context, request *DescribeNotebookSessionStatementRequest) (response *DescribeNotebookSessionStatementResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionStatementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSessionStatement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionStatementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionStatementSqlResultRequest() (request *DescribeNotebookSessionStatementSqlResultRequest) {
    request = &DescribeNotebookSessionStatementSqlResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSessionStatementSqlResult")
    
    
    return
}

func NewDescribeNotebookSessionStatementSqlResultResponse() (response *DescribeNotebookSessionStatementSqlResultResponse) {
    response = &DescribeNotebookSessionStatementSqlResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSessionStatementSqlResult
// 本接口（DescribeNotebookSessionStatementSqlResult）用于获取statement运行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookSessionStatementSqlResult(request *DescribeNotebookSessionStatementSqlResultRequest) (response *DescribeNotebookSessionStatementSqlResultResponse, err error) {
    return c.DescribeNotebookSessionStatementSqlResultWithContext(context.Background(), request)
}

// DescribeNotebookSessionStatementSqlResult
// 本接口（DescribeNotebookSessionStatementSqlResult）用于获取statement运行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookSessionStatementSqlResultWithContext(ctx context.Context, request *DescribeNotebookSessionStatementSqlResultRequest) (response *DescribeNotebookSessionStatementSqlResultResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionStatementSqlResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSessionStatementSqlResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionStatementSqlResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionStatementsRequest() (request *DescribeNotebookSessionStatementsRequest) {
    request = &DescribeNotebookSessionStatementsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSessionStatements")
    
    
    return
}

func NewDescribeNotebookSessionStatementsResponse() (response *DescribeNotebookSessionStatementsResponse) {
    response = &DescribeNotebookSessionStatementsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSessionStatements
// 本接口（DescribeNotebookSessionStatements）用于获取Session Statement列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionStatements(request *DescribeNotebookSessionStatementsRequest) (response *DescribeNotebookSessionStatementsResponse, err error) {
    return c.DescribeNotebookSessionStatementsWithContext(context.Background(), request)
}

// DescribeNotebookSessionStatements
// 本接口（DescribeNotebookSessionStatements）用于获取Session Statement列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) DescribeNotebookSessionStatementsWithContext(ctx context.Context, request *DescribeNotebookSessionStatementsRequest) (response *DescribeNotebookSessionStatementsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionStatementsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSessionStatements require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionStatementsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSessionsRequest() (request *DescribeNotebookSessionsRequest) {
    request = &DescribeNotebookSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeNotebookSessions")
    
    
    return
}

func NewDescribeNotebookSessionsResponse() (response *DescribeNotebookSessionsResponse) {
    response = &DescribeNotebookSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSessions
// 本接口（DescribeNotebookSessions）用于获取notebook livy session列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookSessions(request *DescribeNotebookSessionsRequest) (response *DescribeNotebookSessionsResponse, err error) {
    return c.DescribeNotebookSessionsWithContext(context.Background(), request)
}

// DescribeNotebookSessions
// 本接口（DescribeNotebookSessions）用于获取notebook livy session列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookSessionsWithContext(ctx context.Context, request *DescribeNotebookSessionsRequest) (response *DescribeNotebookSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSessionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSessions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResultDownloadRequest() (request *DescribeResultDownloadRequest) {
    request = &DescribeResultDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeResultDownload")
    
    
    return
}

func NewDescribeResultDownloadResponse() (response *DescribeResultDownloadResponse) {
    response = &DescribeResultDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResultDownload
// 查询结果下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeResultDownload(request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    return c.DescribeResultDownloadWithContext(context.Background(), request)
}

// DescribeResultDownload
// 查询结果下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeResultDownloadWithContext(ctx context.Context, request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeResultDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResultDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResultDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScriptsRequest() (request *DescribeScriptsRequest) {
    request = &DescribeScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeScripts")
    
    
    return
}

func NewDescribeScriptsResponse() (response *DescribeScriptsResponse) {
    response = &DescribeScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScripts
// 该接口（DescribeScripts）用于获取所有SQL查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeScripts(request *DescribeScriptsRequest) (response *DescribeScriptsResponse, err error) {
    return c.DescribeScriptsWithContext(context.Background(), request)
}

// DescribeScripts
// 该接口（DescribeScripts）用于获取所有SQL查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeScriptsWithContext(ctx context.Context, request *DescribeScriptsRequest) (response *DescribeScriptsResponse, err error) {
    if request == nil {
        request = NewDescribeScriptsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScripts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppJobRequest() (request *DescribeSparkAppJobRequest) {
    request = &DescribeSparkAppJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppJob")
    
    
    return
}

func NewDescribeSparkAppJobResponse() (response *DescribeSparkAppJobResponse) {
    response = &DescribeSparkAppJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSparkAppJob
// 查询具体的spark应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
func (c *Client) DescribeSparkAppJob(request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    return c.DescribeSparkAppJobWithContext(context.Background(), request)
}

// DescribeSparkAppJob
// 查询具体的spark应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
func (c *Client) DescribeSparkAppJobWithContext(ctx context.Context, request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppJobsRequest() (request *DescribeSparkAppJobsRequest) {
    request = &DescribeSparkAppJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppJobs")
    
    
    return
}

func NewDescribeSparkAppJobsResponse() (response *DescribeSparkAppJobsResponse) {
    response = &DescribeSparkAppJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSparkAppJobs
// 获取spark应用列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkAppJobs(request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    return c.DescribeSparkAppJobsWithContext(context.Background(), request)
}

// DescribeSparkAppJobs
// 获取spark应用列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkAppJobsWithContext(ctx context.Context, request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppTasksRequest() (request *DescribeSparkAppTasksRequest) {
    request = &DescribeSparkAppTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppTasks")
    
    
    return
}

func NewDescribeSparkAppTasksResponse() (response *DescribeSparkAppTasksResponse) {
    response = &DescribeSparkAppTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSparkAppTasks
// 查询spark应用的运行任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSparkAppTasks(request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    return c.DescribeSparkAppTasksWithContext(context.Background(), request)
}

// DescribeSparkAppTasks
// 查询spark应用的运行任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSparkAppTasksWithContext(ctx context.Context, request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStoreLocationRequest() (request *DescribeStoreLocationRequest) {
    request = &DescribeStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeStoreLocation")
    
    
    return
}

func NewDescribeStoreLocationResponse() (response *DescribeStoreLocationResponse) {
    response = &DescribeStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStoreLocation
// 查询计算结果存储位置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStoreLocation(request *DescribeStoreLocationRequest) (response *DescribeStoreLocationResponse, err error) {
    return c.DescribeStoreLocationWithContext(context.Background(), request)
}

// DescribeStoreLocation
// 查询计算结果存储位置。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStoreLocationWithContext(ctx context.Context, request *DescribeStoreLocationRequest) (response *DescribeStoreLocationResponse, err error) {
    if request == nil {
        request = NewDescribeStoreLocationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableRequest() (request *DescribeTableRequest) {
    request = &DescribeTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTable")
    
    
    return
}

func NewDescribeTableResponse() (response *DescribeTableResponse) {
    response = &DescribeTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTable
// 本接口（DescribeTable），用于查询单个表的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTable(request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    return c.DescribeTableWithContext(context.Background(), request)
}

// DescribeTable
// 本接口（DescribeTable），用于查询单个表的详细信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTableWithContext(ctx context.Context, request *DescribeTableRequest) (response *DescribeTableResponse, err error) {
    if request == nil {
        request = NewDescribeTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTables
// 本接口（DescribeTables）用于查询数据表列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// 本接口（DescribeTables）用于查询数据表列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTaskResult")
    
    
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskResult
// 查询任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    return c.DescribeTaskResultWithContext(context.Background(), request)
}

// DescribeTaskResult
// 查询任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 该接口（DescribleTasks）用于查询任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// 该接口（DescribleTasks）用于查询任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsers
// 获取用户列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    return c.DescribeUsersWithContext(context.Background(), request)
}

// DescribeUsers
// 获取用户列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeUsersWithContext(ctx context.Context, request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewsRequest() (request *DescribeViewsRequest) {
    request = &DescribeViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeViews")
    
    
    return
}

func NewDescribeViewsResponse() (response *DescribeViewsResponse) {
    response = &DescribeViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeViews
// 本接口（DescribeViews）用于查询数据视图列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    return c.DescribeViewsWithContext(context.Background(), request)
}

// DescribeViews
// 本接口（DescribeViews）用于查询数据视图列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeViewsWithContext(ctx context.Context, request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    if request == nil {
        request = NewDescribeViewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeViews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeViewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkGroupsRequest() (request *DescribeWorkGroupsRequest) {
    request = &DescribeWorkGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeWorkGroups")
    
    
    return
}

func NewDescribeWorkGroupsResponse() (response *DescribeWorkGroupsResponse) {
    response = &DescribeWorkGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWorkGroups
// 获取工作组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeWorkGroups(request *DescribeWorkGroupsRequest) (response *DescribeWorkGroupsResponse, err error) {
    return c.DescribeWorkGroupsWithContext(context.Background(), request)
}

// DescribeWorkGroups
// 获取工作组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeWorkGroupsWithContext(ctx context.Context, request *DescribeWorkGroupsRequest) (response *DescribeWorkGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachUserPolicyRequest() (request *DetachUserPolicyRequest) {
    request = &DetachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DetachUserPolicy")
    
    
    return
}

func NewDetachUserPolicyResponse() (response *DetachUserPolicyResponse) {
    response = &DetachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachUserPolicy
// 解绑用户鉴权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DetachUserPolicy(request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    return c.DetachUserPolicyWithContext(context.Background(), request)
}

// DetachUserPolicy
// 解绑用户鉴权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DetachUserPolicyWithContext(ctx context.Context, request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    if request == nil {
        request = NewDetachUserPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachWorkGroupPolicyRequest() (request *DetachWorkGroupPolicyRequest) {
    request = &DetachWorkGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DetachWorkGroupPolicy")
    
    
    return
}

func NewDetachWorkGroupPolicyResponse() (response *DetachWorkGroupPolicyResponse) {
    response = &DetachWorkGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachWorkGroupPolicy
// 解绑工作组鉴权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DetachWorkGroupPolicy(request *DetachWorkGroupPolicyRequest) (response *DetachWorkGroupPolicyResponse, err error) {
    return c.DetachWorkGroupPolicyWithContext(context.Background(), request)
}

// DetachWorkGroupPolicy
// 解绑工作组鉴权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DetachWorkGroupPolicyWithContext(ctx context.Context, request *DetachWorkGroupPolicyRequest) (response *DetachWorkGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDetachWorkGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachWorkGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachWorkGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDropDMSDatabaseRequest() (request *DropDMSDatabaseRequest) {
    request = &DropDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDMSDatabase")
    
    
    return
}

func NewDropDMSDatabaseResponse() (response *DropDMSDatabaseResponse) {
    response = &DropDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DropDMSDatabase
// DMS元数据删除库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSDatabase(request *DropDMSDatabaseRequest) (response *DropDMSDatabaseResponse, err error) {
    return c.DropDMSDatabaseWithContext(context.Background(), request)
}

// DropDMSDatabase
// DMS元数据删除库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSDatabaseWithContext(ctx context.Context, request *DropDMSDatabaseRequest) (response *DropDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewDropDMSDatabaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDropDMSPartitionsRequest() (request *DropDMSPartitionsRequest) {
    request = &DropDMSPartitionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDMSPartitions")
    
    
    return
}

func NewDropDMSPartitionsResponse() (response *DropDMSPartitionsResponse) {
    response = &DropDMSPartitionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DropDMSPartitions
// DMS元数据删除分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSPartitions(request *DropDMSPartitionsRequest) (response *DropDMSPartitionsResponse, err error) {
    return c.DropDMSPartitionsWithContext(context.Background(), request)
}

// DropDMSPartitions
// DMS元数据删除分区
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSPartitionsWithContext(ctx context.Context, request *DropDMSPartitionsRequest) (response *DropDMSPartitionsResponse, err error) {
    if request == nil {
        request = NewDropDMSPartitionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDMSPartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDMSPartitionsResponse()
    err = c.Send(request, response)
    return
}

func NewDropDMSTableRequest() (request *DropDMSTableRequest) {
    request = &DropDMSTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDMSTable")
    
    
    return
}

func NewDropDMSTableResponse() (response *DropDMSTableResponse) {
    response = &DropDMSTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DropDMSTable
// DMS元数据删除表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSTable(request *DropDMSTableRequest) (response *DropDMSTableResponse, err error) {
    return c.DropDMSTableWithContext(context.Background(), request)
}

// DropDMSTable
// DMS元数据删除表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSTableWithContext(ctx context.Context, request *DropDMSTableRequest) (response *DropDMSTableResponse, err error) {
    if request == nil {
        request = NewDropDMSTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDMSTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDMSTableResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateCreateMangedTableSqlRequest() (request *GenerateCreateMangedTableSqlRequest) {
    request = &GenerateCreateMangedTableSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "GenerateCreateMangedTableSql")
    
    
    return
}

func NewGenerateCreateMangedTableSqlResponse() (response *GenerateCreateMangedTableSqlResponse) {
    response = &GenerateCreateMangedTableSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateCreateMangedTableSql
// 生成创建托管表语句
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCreateMangedTableSql(request *GenerateCreateMangedTableSqlRequest) (response *GenerateCreateMangedTableSqlResponse, err error) {
    return c.GenerateCreateMangedTableSqlWithContext(context.Background(), request)
}

// GenerateCreateMangedTableSql
// 生成创建托管表语句
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCreateMangedTableSqlWithContext(ctx context.Context, request *GenerateCreateMangedTableSqlRequest) (response *GenerateCreateMangedTableSqlResponse, err error) {
    if request == nil {
        request = NewGenerateCreateMangedTableSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateCreateMangedTableSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateCreateMangedTableSqlResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskJobLogDetailRequest() (request *ListTaskJobLogDetailRequest) {
    request = &ListTaskJobLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ListTaskJobLogDetail")
    
    
    return
}

func NewListTaskJobLogDetailResponse() (response *ListTaskJobLogDetailResponse) {
    response = &ListTaskJobLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTaskJobLogDetail
// 本接口（ListTaskJobLogDetail）用于获取spark-jar日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
func (c *Client) ListTaskJobLogDetail(request *ListTaskJobLogDetailRequest) (response *ListTaskJobLogDetailResponse, err error) {
    return c.ListTaskJobLogDetailWithContext(context.Background(), request)
}

// ListTaskJobLogDetail
// 本接口（ListTaskJobLogDetail）用于获取spark-jar日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
func (c *Client) ListTaskJobLogDetailWithContext(ctx context.Context, request *ListTaskJobLogDetailRequest) (response *ListTaskJobLogDetailResponse, err error) {
    if request == nil {
        request = NewListTaskJobLogDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskJobLogDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskJobLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewLockMetaDataRequest() (request *LockMetaDataRequest) {
    request = &LockMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "LockMetaData")
    
    
    return
}

func NewLockMetaDataResponse() (response *LockMetaDataResponse) {
    response = &LockMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LockMetaData
// 元数据锁
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) LockMetaData(request *LockMetaDataRequest) (response *LockMetaDataResponse, err error) {
    return c.LockMetaDataWithContext(context.Background(), request)
}

// LockMetaData
// 元数据锁
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) LockMetaDataWithContext(ctx context.Context, request *LockMetaDataRequest) (response *LockMetaDataResponse, err error) {
    if request == nil {
        request = NewLockMetaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LockMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewLockMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernEventRuleRequest() (request *ModifyGovernEventRuleRequest) {
    request = &ModifyGovernEventRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyGovernEventRule")
    
    
    return
}

func NewModifyGovernEventRuleResponse() (response *ModifyGovernEventRuleResponse) {
    response = &ModifyGovernEventRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGovernEventRule
// 修改数据治理事件阈值
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyGovernEventRule(request *ModifyGovernEventRuleRequest) (response *ModifyGovernEventRuleResponse, err error) {
    return c.ModifyGovernEventRuleWithContext(context.Background(), request)
}

// ModifyGovernEventRule
// 修改数据治理事件阈值
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyGovernEventRuleWithContext(ctx context.Context, request *ModifyGovernEventRuleRequest) (response *ModifyGovernEventRuleResponse, err error) {
    if request == nil {
        request = NewModifyGovernEventRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernEventRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernEventRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySparkAppRequest() (request *ModifySparkAppRequest) {
    request = &ModifySparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifySparkApp")
    
    
    return
}

func NewModifySparkAppResponse() (response *ModifySparkAppResponse) {
    response = &ModifySparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySparkApp
// 更新spark应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
func (c *Client) ModifySparkApp(request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    return c.ModifySparkAppWithContext(context.Background(), request)
}

// ModifySparkApp
// 更新spark应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
func (c *Client) ModifySparkAppWithContext(ctx context.Context, request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    if request == nil {
        request = NewModifySparkAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUser
// 修改用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// 修改用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkGroupRequest() (request *ModifyWorkGroupRequest) {
    request = &ModifyWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyWorkGroup")
    
    
    return
}

func NewModifyWorkGroupResponse() (response *ModifyWorkGroupResponse) {
    response = &ModifyWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWorkGroup
// 修改工作组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"
func (c *Client) ModifyWorkGroup(request *ModifyWorkGroupRequest) (response *ModifyWorkGroupResponse, err error) {
    return c.ModifyWorkGroupWithContext(context.Background(), request)
}

// ModifyWorkGroup
// 修改工作组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"
func (c *Client) ModifyWorkGroupWithContext(ctx context.Context, request *ModifyWorkGroupRequest) (response *ModifyWorkGroupResponse, err error) {
    if request == nil {
        request = NewModifyWorkGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReportHeartbeatMetaDataRequest() (request *ReportHeartbeatMetaDataRequest) {
    request = &ReportHeartbeatMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ReportHeartbeatMetaData")
    
    
    return
}

func NewReportHeartbeatMetaDataResponse() (response *ReportHeartbeatMetaDataResponse) {
    response = &ReportHeartbeatMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportHeartbeatMetaData
// 上报元数据心跳
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReportHeartbeatMetaData(request *ReportHeartbeatMetaDataRequest) (response *ReportHeartbeatMetaDataResponse, err error) {
    return c.ReportHeartbeatMetaDataWithContext(context.Background(), request)
}

// ReportHeartbeatMetaData
// 上报元数据心跳
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ReportHeartbeatMetaDataWithContext(ctx context.Context, request *ReportHeartbeatMetaDataRequest) (response *ReportHeartbeatMetaDataResponse, err error) {
    if request == nil {
        request = NewReportHeartbeatMetaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportHeartbeatMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportHeartbeatMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewSuspendResumeDataEngineRequest() (request *SuspendResumeDataEngineRequest) {
    request = &SuspendResumeDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "SuspendResumeDataEngine")
    
    
    return
}

func NewSuspendResumeDataEngineResponse() (response *SuspendResumeDataEngineResponse) {
    response = &SuspendResumeDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SuspendResumeDataEngine
// 本接口用于控制暂停或恢复数据引擎
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngine(request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    return c.SuspendResumeDataEngineWithContext(context.Background(), request)
}

// SuspendResumeDataEngine
// 本接口用于控制暂停或恢复数据引擎
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngineWithContext(ctx context.Context, request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    if request == nil {
        request = NewSuspendResumeDataEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SuspendResumeDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewSuspendResumeDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDataEngineRequest() (request *SwitchDataEngineRequest) {
    request = &SwitchDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "SwitchDataEngine")
    
    
    return
}

func NewSwitchDataEngineResponse() (response *SwitchDataEngineResponse) {
    response = &SwitchDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDataEngine
// 切换主备集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SwitchDataEngine(request *SwitchDataEngineRequest) (response *SwitchDataEngineResponse, err error) {
    return c.SwitchDataEngineWithContext(context.Background(), request)
}

// SwitchDataEngine
// 切换主备集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SwitchDataEngineWithContext(ctx context.Context, request *SwitchDataEngineRequest) (response *SwitchDataEngineResponse, err error) {
    if request == nil {
        request = NewSwitchDataEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindWorkGroupsFromUserRequest() (request *UnbindWorkGroupsFromUserRequest) {
    request = &UnbindWorkGroupsFromUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UnbindWorkGroupsFromUser")
    
    
    return
}

func NewUnbindWorkGroupsFromUserResponse() (response *UnbindWorkGroupsFromUserResponse) {
    response = &UnbindWorkGroupsFromUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindWorkGroupsFromUser
// 解绑用户上的用户组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"
func (c *Client) UnbindWorkGroupsFromUser(request *UnbindWorkGroupsFromUserRequest) (response *UnbindWorkGroupsFromUserResponse, err error) {
    return c.UnbindWorkGroupsFromUserWithContext(context.Background(), request)
}

// UnbindWorkGroupsFromUser
// 解绑用户上的用户组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"
func (c *Client) UnbindWorkGroupsFromUserWithContext(ctx context.Context, request *UnbindWorkGroupsFromUserRequest) (response *UnbindWorkGroupsFromUserResponse, err error) {
    if request == nil {
        request = NewUnbindWorkGroupsFromUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindWorkGroupsFromUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindWorkGroupsFromUserResponse()
    err = c.Send(request, response)
    return
}

func NewUnlockMetaDataRequest() (request *UnlockMetaDataRequest) {
    request = &UnlockMetaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UnlockMetaData")
    
    
    return
}

func NewUnlockMetaDataResponse() (response *UnlockMetaDataResponse) {
    response = &UnlockMetaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnlockMetaData
// 元数据解锁
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnlockMetaData(request *UnlockMetaDataRequest) (response *UnlockMetaDataResponse, err error) {
    return c.UnlockMetaDataWithContext(context.Background(), request)
}

// UnlockMetaData
// 元数据解锁
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnlockMetaDataWithContext(ctx context.Context, request *UnlockMetaDataRequest) (response *UnlockMetaDataResponse, err error) {
    if request == nil {
        request = NewUnlockMetaDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlockMetaData require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlockMetaDataResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRowFilterRequest() (request *UpdateRowFilterRequest) {
    request = &UpdateRowFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpdateRowFilter")
    
    
    return
}

func NewUpdateRowFilterResponse() (response *UpdateRowFilterResponse) {
    response = &UpdateRowFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRowFilter
// 此接口用于更新行过滤规则。注意只能更新过滤规则，不能更新规格对象catalog，database和table。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateRowFilter(request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    return c.UpdateRowFilterWithContext(context.Background(), request)
}

// UpdateRowFilter
// 此接口用于更新行过滤规则。注意只能更新过滤规则，不能更新规格对象catalog，database和table。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateRowFilterWithContext(ctx context.Context, request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    if request == nil {
        request = NewUpdateRowFilterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRowFilter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRowFilterResponse()
    err = c.Send(request, response)
    return
}
