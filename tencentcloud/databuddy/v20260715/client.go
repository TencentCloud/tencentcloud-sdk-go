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

package v20260715

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-07-15"

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


func NewCreateFileRequest() (request *CreateFileRequest) {
    request = &CreateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateFile")
    
    
    return
}

func NewCreateFileResponse() (response *CreateFileResponse) {
    response = &CreateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFile
// 在Studio（统一开发 IDE）的工作空间文件树中新建一个文件（Notebook/SQL/Python等），创建成功后返回文件的完整元信息。
//
// 
//
// **前置条件**
//
// 1. WorkspaceId 对应工作空间存在，且调用方为该工作空间成员；
//
// 2. ParentFolderPath 对应的父文件夹必须存在，且调用方对其有写权限（根目录传 `/`）；
//
// 3. FileName 在同一父文件夹下不能重名（含后缀比较）；
//
// 4. FileName 后缀必须与 FileType 匹配（`.ipynb`↔`NOTEBOOK_FILE`、`.sql`↔`SQL_FILE`）；
//
// 5. 需带文件内容创建时通过 Storage 传入（大文件走 COS 中转，小文件放 Storage.Content）。
func (c *Client) CreateFile(request *CreateFileRequest) (response *CreateFileResponse, err error) {
    return c.CreateFileWithContext(context.Background(), request)
}

// CreateFile
// 在Studio（统一开发 IDE）的工作空间文件树中新建一个文件（Notebook/SQL/Python等），创建成功后返回文件的完整元信息。
//
// 
//
// **前置条件**
//
// 1. WorkspaceId 对应工作空间存在，且调用方为该工作空间成员；
//
// 2. ParentFolderPath 对应的父文件夹必须存在，且调用方对其有写权限（根目录传 `/`）；
//
// 3. FileName 在同一父文件夹下不能重名（含后缀比较）；
//
// 4. FileName 后缀必须与 FileType 匹配（`.ipynb`↔`NOTEBOOK_FILE`、`.sql`↔`SQL_FILE`）；
//
// 5. 需带文件内容创建时通过 Storage 传入（大文件走 COS 中转，小文件放 Storage.Content）。
func (c *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest) (response *CreateFileResponse, err error) {
    if request == nil {
        request = NewCreateFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileRequest() (request *DeleteFileRequest) {
    request = &DeleteFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteFile")
    
    
    return
}

func NewDeleteFileResponse() (response *DeleteFileResponse) {
    response = &DeleteFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFile
// 将文件移入回收站（软删除），同时清理该文件的版本记录与执行结果快照。
//
// 
//
// **前置条件**
//
// 1. FileId 对应文件必须存在且为活跃状态；
//
// 2. 调用方对该文件有删除权限；
//
// 3. 文件未被工作流任务引用。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | 缺少 FileId | 请传入 FileId  |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId |
//
// | `ResourceInUse.FileReferencedByTask` | 1030204 | 文件被工作流任务引用，不允许删除 | 请先解除任务引用后再删除 |
//
// | `UnauthorizedOperation.FileDeleteDenied` | 1030303 | 对该文件无删除权限 | 请联系文件负责人或空间管理员授权 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) DeleteFile(request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    return c.DeleteFileWithContext(context.Background(), request)
}

// DeleteFile
// 将文件移入回收站（软删除），同时清理该文件的版本记录与执行结果快照。
//
// 
//
// **前置条件**
//
// 1. FileId 对应文件必须存在且为活跃状态；
//
// 2. 调用方对该文件有删除权限；
//
// 3. 文件未被工作流任务引用。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | 缺少 FileId | 请传入 FileId  |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId |
//
// | `ResourceInUse.FileReferencedByTask` | 1030204 | 文件被工作流任务引用，不允许删除 | 请先解除任务引用后再删除 |
//
// | `UnauthorizedOperation.FileDeleteDenied` | 1030303 | 对该文件无删除权限 | 请联系文件负责人或空间管理员授权 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_CREATEWORKFLOWFAILED = "FailedOperation.CreateWorkflowFailed"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
//  FAILEDOPERATION_WORKFLOWCOUNTLIMIT = "FailedOperation.WorkflowCountLimit"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) DeleteFileWithContext(ctx context.Context, request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    if request == nil {
        request = NewDeleteFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkflow")
    
    
    return
}

func NewDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
    response = &DeleteWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetFileRequest() (request *GetFileRequest) {
    request = &GetFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetFile")
    
    
    return
}

func NewGetFileResponse() (response *GetFileResponse) {
    response = &GetFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFile
// 获取文件的元信息，可选包含文件内容，支持按版本读取历史快照。
//
// 
//
// **前置条件**
//
// 1. FileId 与 FilePath 二选一，至少传一个；同时传时以 FileId 为准；
//
// 2. 对应文件必须存在，且调用方对该文件有读权限；
//
// 3. 传 VersionId 时该版本必须存在。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | FileId 与 FilePath 同时为空 | FileId 与 FilePath 二选一，至少传一个 |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId 或 FilePath |
//
// | `ResourceNotFound.FileVersionNotFound` | 1030205 | 指定的文件版本不存在 | 请确认 VersionId，或调用 ListFileVersions 获取 |
//
// | `UnauthorizedOperation.FileReadDenied` | 1030304 | 对该文件无读权限 | 请联系文件负责人或空间管理员授权 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) GetFile(request *GetFileRequest) (response *GetFileResponse, err error) {
    return c.GetFileWithContext(context.Background(), request)
}

// GetFile
// 获取文件的元信息，可选包含文件内容，支持按版本读取历史快照。
//
// 
//
// **前置条件**
//
// 1. FileId 与 FilePath 二选一，至少传一个；同时传时以 FileId 为准；
//
// 2. 对应文件必须存在，且调用方对该文件有读权限；
//
// 3. 传 VersionId 时该版本必须存在。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | FileId 与 FilePath 同时为空 | FileId 与 FilePath 二选一，至少传一个 |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId 或 FilePath |
//
// | `ResourceNotFound.FileVersionNotFound` | 1030205 | 指定的文件版本不存在 | 请确认 VersionId，或调用 ListFileVersions 获取 |
//
// | `UnauthorizedOperation.FileReadDenied` | 1030304 | 对该文件无读权限 | 请联系文件负责人或空间管理员授权 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ONEFLOWRESOURCENOEXISTERROR = "ResourceNotFound.OneFlowResourceNoExistError"
func (c *Client) GetFileWithContext(ctx context.Context, request *GetFileRequest) (response *GetFileResponse, err error) {
    if request == nil {
        request = NewGetFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRequest() (request *GetWorkflowRequest) {
    request = &GetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflow")
    
    
    return
}

func NewGetWorkflowResponse() (response *GetWorkflowResponse) {
    response = &GetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflow
// 获取工作流详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) GetWorkflow(request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    return c.GetWorkflowWithContext(context.Background(), request)
}

// GetWorkflow
// 获取工作流详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRunRequest() (request *GetWorkflowRunRequest) {
    request = &GetWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflowRun")
    
    
    return
}

func NewGetWorkflowRunResponse() (response *GetWorkflowRunResponse) {
    response = &GetWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflowRun
// 查询工作流运行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
func (c *Client) GetWorkflowRun(request *GetWorkflowRunRequest) (response *GetWorkflowRunResponse, err error) {
    return c.GetWorkflowRunWithContext(context.Background(), request)
}

// GetWorkflowRun
// 查询工作流运行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
func (c *Client) GetWorkflowRunWithContext(ctx context.Context, request *GetWorkflowRunRequest) (response *GetWorkflowRunResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowTaskRunRequest() (request *GetWorkflowTaskRunRequest) {
    request = &GetWorkflowTaskRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkflowTaskRun")
    
    
    return
}

func NewGetWorkflowTaskRunResponse() (response *GetWorkflowTaskRunResponse) {
    response = &GetWorkflowTaskRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflowTaskRun
// 查询任务运行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkflowTaskRun(request *GetWorkflowTaskRunRequest) (response *GetWorkflowTaskRunResponse, err error) {
    return c.GetWorkflowTaskRunWithContext(context.Background(), request)
}

// GetWorkflowTaskRun
// 查询任务运行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkflowTaskRunWithContext(ctx context.Context, request *GetWorkflowTaskRunRequest) (response *GetWorkflowTaskRunResponse, err error) {
    if request == nil {
        request = NewGetWorkflowTaskRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkflowTaskRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflowTaskRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowTaskRunResponse()
    err = c.Send(request, response)
    return
}

func NewKillWorkflowRunRequest() (request *KillWorkflowRunRequest) {
    request = &KillWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "KillWorkflowRun")
    
    
    return
}

func NewKillWorkflowRunResponse() (response *KillWorkflowRunResponse) {
    response = &KillWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillWorkflowRun
// 终止工作流的运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_NOWORKFLOWEXECUTIONNEEDOPERATE = "FailedOperation.NoWorkflowExecutionNeedOperate"
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEDELETE = "FailedOperation.WorkflowExecutionHasBeDelete"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_WORKFLOWEXECUTIONHASREACHEDFINALSTATECANNOTBESTOPPED = "UnsupportedOperation.WorkflowExecutionHasReachedFinalStateCannotBeStopped"
func (c *Client) KillWorkflowRun(request *KillWorkflowRunRequest) (response *KillWorkflowRunResponse, err error) {
    return c.KillWorkflowRunWithContext(context.Background(), request)
}

// KillWorkflowRun
// 终止工作流的运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_NOWORKFLOWEXECUTIONNEEDOPERATE = "FailedOperation.NoWorkflowExecutionNeedOperate"
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEDELETE = "FailedOperation.WorkflowExecutionHasBeDelete"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_WORKFLOWEXECUTIONHASREACHEDFINALSTATECANNOTBESTOPPED = "UnsupportedOperation.WorkflowExecutionHasReachedFinalStateCannotBeStopped"
func (c *Client) KillWorkflowRunWithContext(ctx context.Context, request *KillWorkflowRunRequest) (response *KillWorkflowRunResponse, err error) {
    if request == nil {
        request = NewKillWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "KillWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowRunsRequest() (request *ListWorkflowRunsRequest) {
    request = &ListWorkflowRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflowRuns")
    
    
    return
}

func NewListWorkflowRunsResponse() (response *ListWorkflowRunsResponse) {
    response = &ListWorkflowRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowRuns
// 工作流运行列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTWORKFLOWEXECUTIONS = "FailedOperation.ExistWorkflowExecutions"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  INVALIDPARAMETERVALUE_WORKFLOWENDTIMELESSSTARTTIME = "InvalidParameterValue.WorkflowEndTimeLessStartTime"
//  INVALIDPARAMETERVALUE_WORKFLOWQUERYENDTIMEANDSTARTTIMEEXCEED = "InvalidParameterValue.WorkflowQueryEndTimeAndStartTimeExceed"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
func (c *Client) ListWorkflowRuns(request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    return c.ListWorkflowRunsWithContext(context.Background(), request)
}

// ListWorkflowRuns
// 工作流运行列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTWORKFLOWEXECUTIONS = "FailedOperation.ExistWorkflowExecutions"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  INVALIDPARAMETERVALUE_WORKFLOWENDTIMELESSSTARTTIME = "InvalidParameterValue.WorkflowEndTimeLessStartTime"
//  INVALIDPARAMETERVALUE_WORKFLOWQUERYENDTIMEANDSTARTTIMEEXCEED = "InvalidParameterValue.WorkflowQueryEndTimeAndStartTimeExceed"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
func (c *Client) ListWorkflowRunsWithContext(ctx context.Context, request *ListWorkflowRunsRequest) (response *ListWorkflowRunsResponse, err error) {
    if request == nil {
        request = NewListWorkflowRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflowRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowTaskRunsRequest() (request *ListWorkflowTaskRunsRequest) {
    request = &ListWorkflowTaskRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflowTaskRuns")
    
    
    return
}

func NewListWorkflowTaskRunsResponse() (response *ListWorkflowTaskRunsResponse) {
    response = &ListWorkflowTaskRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowTaskRuns
// 查询工作流任务历史运行列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowTaskRuns(request *ListWorkflowTaskRunsRequest) (response *ListWorkflowTaskRunsResponse, err error) {
    return c.ListWorkflowTaskRunsWithContext(context.Background(), request)
}

// ListWorkflowTaskRuns
// 查询工作流任务历史运行列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowTaskRunsWithContext(ctx context.Context, request *ListWorkflowTaskRunsRequest) (response *ListWorkflowTaskRunsResponse, err error) {
    if request == nil {
        request = NewListWorkflowTaskRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflowTaskRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowTaskRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowTaskRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowsRequest() (request *ListWorkflowsRequest) {
    request = &ListWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkflows")
    
    
    return
}

func NewListWorkflowsResponse() (response *ListWorkflowsResponse) {
    response = &ListWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflows(request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    return c.ListWorkflowsWithContext(context.Background(), request)
}

// ListWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LISTWORKFLOWFILTERPARAMERROR = "InvalidParameterValue.ListWorkflowFilterParamError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowsWithContext(ctx context.Context, request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    if request == nil {
        request = NewListWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewRerunWorkflowRunRequest() (request *RerunWorkflowRunRequest) {
    request = &RerunWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RerunWorkflowRun")
    
    
    return
}

func NewRerunWorkflowRunResponse() (response *RerunWorkflowRunResponse) {
    response = &RerunWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunWorkflowRun
// 重跑工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RERUNWORKFLOWFAIL = "FailedOperation.RerunWorkflowFail"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMSERROR = "InvalidParameter.InvalidParamsError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RerunWorkflowRun(request *RerunWorkflowRunRequest) (response *RerunWorkflowRunResponse, err error) {
    return c.RerunWorkflowRunWithContext(context.Background(), request)
}

// RerunWorkflowRun
// 重跑工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RERUNWORKFLOWFAIL = "FailedOperation.RerunWorkflowFail"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMSERROR = "InvalidParameter.InvalidParamsError"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  RESOURCENOTFOUND_WORKFLOWEXECUTIONNOTEXIST = "ResourceNotFound.WorkflowExecutionNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RerunWorkflowRunWithContext(ctx context.Context, request *RerunWorkflowRunRequest) (response *RerunWorkflowRunResponse, err error) {
    if request == nil {
        request = NewRerunWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RerunWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewRunWorkflowRequest() (request *RunWorkflowRequest) {
    request = &RunWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RunWorkflow")
    
    
    return
}

func NewRunWorkflowResponse() (response *RunWorkflowResponse) {
    response = &RunWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunWorkflow
// 运行工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RUNWORKFLOWFAIL = "FailedOperation.RunWorkflowFail"
//  FAILEDOPERATION_RUNWORKFLOWFAILEXECUTIONIDEMPTY = "FailedOperation.RunWorkflowFailExecutionIdEmpty"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RunWorkflow(request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    return c.RunWorkflowWithContext(context.Background(), request)
}

// RunWorkflow
// 运行工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_RUNWORKFLOWFAIL = "FailedOperation.RunWorkflowFail"
//  FAILEDOPERATION_RUNWORKFLOWFAILEXECUTIONIDEMPTY = "FailedOperation.RunWorkflowFailExecutionIdEmpty"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_PARAMNULLERROR = "InvalidParameterValue.ParamNullError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWTASKNOTEXIST = "ResourceNotFound.WorkflowTaskNotExist"
func (c *Client) RunWorkflowWithContext(ctx context.Context, request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    if request == nil {
        request = NewRunWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RunWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindWorkflowBundleRequest() (request *UnbindWorkflowBundleRequest) {
    request = &UnbindWorkflowBundleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UnbindWorkflowBundle")
    
    
    return
}

func NewUnbindWorkflowBundleResponse() (response *UnbindWorkflowBundleResponse) {
    response = &UnbindWorkflowBundleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindWorkflowBundle
// 解绑工作流Bundle信息
//
// 说明：本接口语义等同于规范动词清单中的 Detach，因兼容既有产品形态保留 Unbind 命名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UnbindWorkflowBundle(request *UnbindWorkflowBundleRequest) (response *UnbindWorkflowBundleResponse, err error) {
    return c.UnbindWorkflowBundleWithContext(context.Background(), request)
}

// UnbindWorkflowBundle
// 解绑工作流Bundle信息
//
// 说明：本接口语义等同于规范动词清单中的 Detach，因兼容既有产品形态保留 Unbind 命名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UnbindWorkflowBundleWithContext(ctx context.Context, request *UnbindWorkflowBundleRequest) (response *UnbindWorkflowBundleResponse, err error) {
    if request == nil {
        request = NewUnbindWorkflowBundleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UnbindWorkflowBundle")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindWorkflowBundle require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindWorkflowBundleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFileRequest() (request *UpdateFileRequest) {
    request = &UpdateFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateFile")
    
    
    return
}

func NewUpdateFileResponse() (response *UpdateFileResponse) {
    response = &UpdateFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFile
// 更新文件内容与运行配置（计算资源、默认 catalog/schema、参数等），返回更新后的文件元信息。
//
// 
//
// **前置条件**
//
// 1. FileId 对应文件必须存在且为活跃状态；
//
// 2. 调用方对该文件有写权限；
//
// 3. 仅更新配置时不传 Storage；仅更新内容时不传 FileConfig；
//
// 4. FileConfig.ResourceId 非空时会校验资源类型与文件类型的匹配性。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | 缺少 FileId | 请传入 FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `InvalidParameterValue.ResourceId` | 1030104 | 计算资源类型与文件类型不匹配 | Python/Notebook 选数据计算资源，SQL 选数据分析资源 |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId，或调用 GetFile 校验文件状态 |
//
// | `UnauthorizedOperation.FileWriteDenied` | 1030302 | 对该文件无写权限 | 请联系文件负责人或空间管理员授权 |
//
// | `FailedOperation.FileStorageUpdateFailed` | 1030401 | 文件内容写入存储失败 | 请稍后重试，持续失败请携带 RequestId 联系支持 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFile(request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
    return c.UpdateFileWithContext(context.Background(), request)
}

// UpdateFile
// 更新文件内容与运行配置（计算资源、默认 catalog/schema、参数等），返回更新后的文件元信息。
//
// 
//
// **前置条件**
//
// 1. FileId 对应文件必须存在且为活跃状态；
//
// 2. 调用方对该文件有写权限；
//
// 3. 仅更新配置时不传 Storage；仅更新内容时不传 FileConfig；
//
// 4. FileConfig.ResourceId 非空时会校验资源类型与文件类型的匹配性。
//
// 
//
// **错误码（Module 均为 `Studio`）**
//
// 
//
// | 错误码（Code） | InnerCode | 描述 | 处理建议 |
//
// | --- | --- | --- | --- |
//
// | `MissingParameter.WorkspaceId` | 1030001 | 缺少 WorkspaceId | 请传入 WorkspaceId |
//
// | `MissingParameter.FileId` | 1030003 | 缺少 FileId | 请传入 FileId |
//
// | `InvalidParameterValue.FileType` | 1030102 | FileType 取值不支持 | FileType 取 FILE/NOTEBOOK_FILE/SQL_FILE |
//
// | `InvalidParameterValue.ResourceId` | 1030104 | 计算资源类型与文件类型不匹配 | Python/Notebook 选数据计算资源，SQL 选数据分析资源 |
//
// | `ResourceNotFound.FileNotFound` | 1030203 | 文件不存在或已删除 | 请确认 FileId，或调用 GetFile 校验文件状态 |
//
// | `UnauthorizedOperation.FileWriteDenied` | 1030302 | 对该文件无写权限 | 请联系文件负责人或空间管理员授权 |
//
// | `FailedOperation.FileStorageUpdateFailed` | 1030401 | 文件内容写入存储失败 | 请稍后重试，持续失败请携带 RequestId 联系支持 |
//
// | `InternalError` | 1030900 | 服务内部异常 | 请携带 RequestId 联系支持 |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) UpdateFileWithContext(ctx context.Context, request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
    if request == nil {
        request = NewUpdateFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowRequest() (request *UpdateWorkflowRequest) {
    request = &UpdateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkflow")
    
    
    return
}

func NewUpdateWorkflowResponse() (response *UpdateWorkflowResponse) {
    response = &UpdateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflow
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    return c.UpdateWorkflowWithContext(context.Background(), request)
}

// UpdateWorkflow
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_UPDATEWORKFLOWFAILED = "FailedOperation.UpdateWorkflowFailed"
//  FAILEDOPERATION_WORKFLOWBUNDLENOPERMISSION = "FailedOperation.WorkflowBundleNoPermission"
//  FAILEDOPERATION_WORKFLOWCREATELOCKACQUIREFAILED = "FailedOperation.WorkflowCreateLockAcquireFailed"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATETASKNAMEERROR = "InvalidParameterValue.DuplicateTaskNameError"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYELEMENTCOUNTLIMIT = "InvalidParameterValue.LoopDataArrayElementCountLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYJSONINVALID = "InvalidParameterValue.LoopDataArrayJsonInvalid"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAY = "InvalidParameterValue.LoopDataArrayNotJsonArray"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYNOTJSONARRAYLITERAL = "InvalidParameterValue.LoopDataArrayNotJsonArrayLiteral"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUEBLANK = "InvalidParameterValue.LoopDataArrayValueBlank"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVALUELENGTHLIMIT = "InvalidParameterValue.LoopDataArrayValueLengthLimit"
//  INVALIDPARAMETERVALUE_LOOPDATAARRAYVARIABLEEXPRESSIONINVALID = "InvalidParameterValue.LoopDataArrayVariableExpressionInvalid"
//  INVALIDPARAMETERVALUE_PARAMILLEGALERROR = "InvalidParameterValue.ParamIllegalError"
//  INVALIDPARAMETERVALUE_TASKHOOKVALIDATIONFAILED = "InvalidParameterValue.TaskHookValidationFailed"
//  INVALIDPARAMETERVALUE_TASKNAMECONTAINSILLEGALCHARACTERSERROR = "InvalidParameterValue.TaskNameContainsIllegalCharactersError"
//  INVALIDPARAMETERVALUE_TASKNAMEEXCEEDSLIMITERROR = "InvalidParameterValue.TaskNameExceedsLimitError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkflowWithContext(ctx context.Context, request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowResponse()
    err = c.Send(request, response)
    return
}
