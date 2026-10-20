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


func NewAddConsoleUsersRequest() (request *AddConsoleUsersRequest) {
    request = &AddConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddConsoleUsers")
    
    
    return
}

func NewAddConsoleUsersResponse() (response *AddConsoleUsersResponse) {
    response = &AddConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddConsoleUsers
// 添加控制台用户，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
func (c *Client) AddConsoleUsers(request *AddConsoleUsersRequest) (response *AddConsoleUsersResponse, err error) {
    return c.AddConsoleUsersWithContext(context.Background(), request)
}

// AddConsoleUsers
// 添加控制台用户，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
func (c *Client) AddConsoleUsersWithContext(ctx context.Context, request *AddConsoleUsersRequest) (response *AddConsoleUsersResponse, err error) {
    if request == nil {
        request = NewAddConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewAddWorkspaceGroupsRequest() (request *AddWorkspaceGroupsRequest) {
    request = &AddWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddWorkspaceGroups")
    
    
    return
}

func NewAddWorkspaceGroupsResponse() (response *AddWorkspaceGroupsResponse) {
    response = &AddWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddWorkspaceGroups
// 添加工作空间用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEGROUPMEMBERQUOTA = "LimitExceeded.WorkspaceGroupMemberQuota"
//  LIMITEXCEEDED_WORKSPACEGROUPQUOTA = "LimitExceeded.WorkspaceGroupQuota"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceGroups(request *AddWorkspaceGroupsRequest) (response *AddWorkspaceGroupsResponse, err error) {
    return c.AddWorkspaceGroupsWithContext(context.Background(), request)
}

// AddWorkspaceGroups
// 添加工作空间用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEGROUPMEMBERQUOTA = "LimitExceeded.WorkspaceGroupMemberQuota"
//  LIMITEXCEEDED_WORKSPACEGROUPQUOTA = "LimitExceeded.WorkspaceGroupQuota"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceGroupsWithContext(ctx context.Context, request *AddWorkspaceGroupsRequest) (response *AddWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewAddWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddWorkspaceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAddWorkspaceUsersRequest() (request *AddWorkspaceUsersRequest) {
    request = &AddWorkspaceUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "AddWorkspaceUsers")
    
    
    return
}

func NewAddWorkspaceUsersResponse() (response *AddWorkspaceUsersResponse) {
    response = &AddWorkspaceUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddWorkspaceUsers
// 创建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceUsers(request *AddWorkspaceUsersRequest) (response *AddWorkspaceUsersResponse, err error) {
    return c.AddWorkspaceUsersWithContext(context.Background(), request)
}

// AddWorkspaceUsers
// 创建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_TENCENTCLOUDSDKERROR = "FailedOperation.TencentCloudSDKError"
//  INTERNALERROR_INTERFACECALLERROR = "InternalError.InterfaceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_JWTAUTHFAILURE = "InvalidParameter.JWTAuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_WORKSPACEUSERQUOTA = "LimitExceeded.WorkspaceUserQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddWorkspaceUsersWithContext(ctx context.Context, request *AddWorkspaceUsersRequest) (response *AddWorkspaceUsersResponse, err error) {
    if request == nil {
        request = NewAddWorkspaceUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "AddWorkspaceUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddWorkspaceUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddWorkspaceUsersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCatalogRequest() (request *CreateCatalogRequest) {
    request = &CreateCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateCatalog")
    
    
    return
}

func NewCreateCatalogResponse() (response *CreateCatalogResponse) {
    response = &CreateCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCatalog
// 创建数据目录接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateCatalog(request *CreateCatalogRequest) (response *CreateCatalogResponse, err error) {
    return c.CreateCatalogWithContext(context.Background(), request)
}

// CreateCatalog
// 创建数据目录接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateCatalogWithContext(ctx context.Context, request *CreateCatalogRequest) (response *CreateCatalogResponse, err error) {
    if request == nil {
        request = NewCreateCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleGroupRequest() (request *CreateConsoleGroupRequest) {
    request = &CreateConsoleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateConsoleGroup")
    
    
    return
}

func NewCreateConsoleGroupResponse() (response *CreateConsoleGroupResponse) {
    response = &CreateConsoleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsoleGroup
// 创建控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateConsoleGroup(request *CreateConsoleGroupRequest) (response *CreateConsoleGroupResponse, err error) {
    return c.CreateConsoleGroupWithContext(context.Background(), request)
}

// CreateConsoleGroup
// 创建控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateConsoleGroupWithContext(ctx context.Context, request *CreateConsoleGroupRequest) (response *CreateConsoleGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsoleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateConsoleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleGroupResponse()
    err = c.Send(request, response)
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
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
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
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
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

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateFolder")
    
    
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

// CreateFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEANACATALOG = "FailedOperation.DuplicateAnaCatalog"
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  LIMITEXCEEDED_SETATSCATALOGLIMITEXCEEDED = "LimitExceeded.SetatsCatalogLimitExceeded"
//  RESOURCEINUSE_CATALOGALREADYEXISTS = "ResourceInUse.CatalogAlreadyExists"
func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchemaRequest() (request *CreateSchemaRequest) {
    request = &CreateSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateSchema")
    
    
    return
}

func NewCreateSchemaResponse() (response *CreateSchemaResponse) {
    response = &CreateSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchema
// 创建schema
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SCHEMAALREADYEXISTS = "ResourceInUse.SchemaAlreadyExists"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) CreateSchema(request *CreateSchemaRequest) (response *CreateSchemaResponse, err error) {
    return c.CreateSchemaWithContext(context.Background(), request)
}

// CreateSchema
// 创建schema
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SCHEMAALREADYEXISTS = "ResourceInUse.SchemaAlreadyExists"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) CreateSchemaWithContext(ctx context.Context, request *CreateSchemaRequest) (response *CreateSchemaResponse, err error) {
    if request == nil {
        request = NewCreateSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchemaResponse()
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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

func NewCreateWorkspaceRequest() (request *CreateWorkspaceRequest) {
    request = &CreateWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkspace")
    
    
    return
}

func NewCreateWorkspaceResponse() (response *CreateWorkspaceResponse) {
    response = &CreateWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspace
// 创建工作空间，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkspace(request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    return c.CreateWorkspaceWithContext(context.Background(), request)
}

// CreateWorkspace
// 创建工作空间，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
func (c *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspaceRoleRequest() (request *CreateWorkspaceRoleRequest) {
    request = &CreateWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "CreateWorkspaceRole")
    
    
    return
}

func NewCreateWorkspaceRoleResponse() (response *CreateWorkspaceRoleResponse) {
    response = &CreateWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaceRole
// 创建工作空间角色
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CUSTOMROLEQUOTA = "LimitExceeded.CustomRoleQuota"
func (c *Client) CreateWorkspaceRole(request *CreateWorkspaceRoleRequest) (response *CreateWorkspaceRoleResponse, err error) {
    return c.CreateWorkspaceRoleWithContext(context.Background(), request)
}

// CreateWorkspaceRole
// 创建工作空间角色
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CUSTOMROLEQUOTA = "LimitExceeded.CustomRoleQuota"
func (c *Client) CreateWorkspaceRoleWithContext(ctx context.Context, request *CreateWorkspaceRoleRequest) (response *CreateWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewCreateWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "CreateWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCatalogRequest() (request *DeleteCatalogRequest) {
    request = &DeleteCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteCatalog")
    
    
    return
}

func NewDeleteCatalogResponse() (response *DeleteCatalogResponse) {
    response = &DeleteCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCatalog
// 删除catalog
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteCatalog(request *DeleteCatalogRequest) (response *DeleteCatalogResponse, err error) {
    return c.DeleteCatalogWithContext(context.Background(), request)
}

// DeleteCatalog
// 删除catalog
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteCatalogWithContext(ctx context.Context, request *DeleteCatalogRequest) (response *DeleteCatalogResponse, err error) {
    if request == nil {
        request = NewDeleteCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsoleGroupsRequest() (request *DeleteConsoleGroupsRequest) {
    request = &DeleteConsoleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteConsoleGroups")
    
    
    return
}

func NewDeleteConsoleGroupsResponse() (response *DeleteConsoleGroupsResponse) {
    response = &DeleteConsoleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsoleGroups
// 删除控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteConsoleGroups(request *DeleteConsoleGroupsRequest) (response *DeleteConsoleGroupsResponse, err error) {
    return c.DeleteConsoleGroupsWithContext(context.Background(), request)
}

// DeleteConsoleGroups
// 删除控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteConsoleGroupsWithContext(ctx context.Context, request *DeleteConsoleGroupsRequest) (response *DeleteConsoleGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteConsoleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteConsoleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsoleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsoleGroupsResponse()
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
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
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
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
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

func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
    request = &DeleteFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteFolder")
    
    
    return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
    response = &DeleteFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFolder
// 删除文件夹
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    return c.DeleteFolderWithContext(context.Background(), request)
}

// DeleteFolder
// 删除文件夹
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    if request == nil {
        request = NewDeleteFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSchemaRequest() (request *DeleteSchemaRequest) {
    request = &DeleteSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteSchema")
    
    
    return
}

func NewDeleteSchemaResponse() (response *DeleteSchemaResponse) {
    response = &DeleteSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSchema
// 删除schema
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) DeleteSchema(request *DeleteSchemaRequest) (response *DeleteSchemaResponse, err error) {
    return c.DeleteSchemaWithContext(context.Background(), request)
}

// DeleteSchema
// 删除schema
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) DeleteSchemaWithContext(ctx context.Context, request *DeleteSchemaRequest) (response *DeleteSchemaResponse, err error) {
    if request == nil {
        request = NewDeleteSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableRequest() (request *DeleteTableRequest) {
    request = &DeleteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteTable")
    
    
    return
}

func NewDeleteTableResponse() (response *DeleteTableResponse) {
    response = &DeleteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTable
// 删除table
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) DeleteTable(request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    return c.DeleteTableWithContext(context.Background(), request)
}

// DeleteTable
// 删除table
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) DeleteTableWithContext(ctx context.Context, request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
    if request == nil {
        request = NewDeleteTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableResponse()
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

func NewDeleteWorkspaceRequest() (request *DeleteWorkspaceRequest) {
    request = &DeleteWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkspace")
    
    
    return
}

func NewDeleteWorkspaceResponse() (response *DeleteWorkspaceResponse) {
    response = &DeleteWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkspace
// 删除工作空间
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
func (c *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (response *DeleteWorkspaceResponse, err error) {
    return c.DeleteWorkspaceWithContext(context.Background(), request)
}

// DeleteWorkspace
// 删除工作空间
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
func (c *Client) DeleteWorkspaceWithContext(ctx context.Context, request *DeleteWorkspaceRequest) (response *DeleteWorkspaceResponse, err error) {
    if request == nil {
        request = NewDeleteWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkspaceRoleRequest() (request *DeleteWorkspaceRoleRequest) {
    request = &DeleteWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "DeleteWorkspaceRole")
    
    
    return
}

func NewDeleteWorkspaceRoleResponse() (response *DeleteWorkspaceRoleResponse) {
    response = &DeleteWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkspaceRole
// 删除工作空间角色
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
func (c *Client) DeleteWorkspaceRole(request *DeleteWorkspaceRoleRequest) (response *DeleteWorkspaceRoleResponse, err error) {
    return c.DeleteWorkspaceRoleWithContext(context.Background(), request)
}

// DeleteWorkspaceRole
// 删除工作空间角色
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
func (c *Client) DeleteWorkspaceRoleWithContext(ctx context.Context, request *DeleteWorkspaceRoleRequest) (response *DeleteWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewDeleteWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "DeleteWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewGetCatalogRequest() (request *GetCatalogRequest) {
    request = &GetCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetCatalog")
    
    
    return
}

func NewGetCatalogResponse() (response *GetCatalogResponse) {
    response = &GetCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCatalog
// 查询catalog详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetCatalog(request *GetCatalogRequest) (response *GetCatalogResponse, err error) {
    return c.GetCatalogWithContext(context.Background(), request)
}

// GetCatalog
// 查询catalog详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetCatalogWithContext(ctx context.Context, request *GetCatalogRequest) (response *GetCatalogResponse, err error) {
    if request == nil {
        request = NewGetCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCatalogResponse()
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
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
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
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
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

func NewGetFolderRequest() (request *GetFolderRequest) {
    request = &GetFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetFolder")
    
    
    return
}

func NewGetFolderResponse() (response *GetFolderResponse) {
    response = &GetFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFolder
// 获取文件夹详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFolder(request *GetFolderRequest) (response *GetFolderResponse, err error) {
    return c.GetFolderWithContext(context.Background(), request)
}

// GetFolder
// 获取文件夹详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) GetFolderWithContext(ctx context.Context, request *GetFolderRequest) (response *GetFolderResponse, err error) {
    if request == nil {
        request = NewGetFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetSchemaRequest() (request *GetSchemaRequest) {
    request = &GetSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetSchema")
    
    
    return
}

func NewGetSchemaResponse() (response *GetSchemaResponse) {
    response = &GetSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSchema
// 查询schema详情
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) GetSchema(request *GetSchemaRequest) (response *GetSchemaResponse, err error) {
    return c.GetSchemaWithContext(context.Background(), request)
}

// GetSchema
// 查询schema详情
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) GetSchemaWithContext(ctx context.Context, request *GetSchemaRequest) (response *GetSchemaResponse, err error) {
    if request == nil {
        request = NewGetSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewGetTableRequest() (request *GetTableRequest) {
    request = &GetTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetTable")
    
    
    return
}

func NewGetTableResponse() (response *GetTableResponse) {
    response = &GetTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTable
// 查询table详情
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) GetTable(request *GetTableRequest) (response *GetTableResponse, err error) {
    return c.GetTableWithContext(context.Background(), request)
}

// GetTable
// 查询table详情
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) GetTableWithContext(ctx context.Context, request *GetTableRequest) (response *GetTableResponse, err error) {
    if request == nil {
        request = NewGetTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTableResponse()
    err = c.Send(request, response)
    return
}

func NewGetTablesRequest() (request *GetTablesRequest) {
    request = &GetTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetTables")
    
    
    return
}

func NewGetTablesResponse() (response *GetTablesResponse) {
    response = &GetTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTables
// 批量获取表的元数据。使用场景：根据表的全名列表（catalogName.schemaName.tableName）批量获取表的详细元数据信息。前置条件：需要有对应workspace的读权限，仅返回当前workspace有权限的catalog下的表。权限校验：通过getAuthedCatalogNames获取有权限的catalogName进行过滤，不在权限范围内的表会被静默跳过。限制：FullNames最大支持20个表全名
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) GetTables(request *GetTablesRequest) (response *GetTablesResponse, err error) {
    return c.GetTablesWithContext(context.Background(), request)
}

// GetTables
// 批量获取表的元数据。使用场景：根据表的全名列表（catalogName.schemaName.tableName）批量获取表的详细元数据信息。前置条件：需要有对应workspace的读权限，仅返回当前workspace有权限的catalog下的表。权限校验：通过getAuthedCatalogNames获取有权限的catalogName进行过滤，不在权限范围内的表会被静默跳过。限制：FullNames最大支持20个表全名
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_TABLENOTFOUND = "ResourceNotFound.TableNotFound"
func (c *Client) GetTablesWithContext(ctx context.Context, request *GetTablesRequest) (response *GetTablesResponse, err error) {
    if request == nil {
        request = NewGetTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTablesResponse()
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
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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

func NewGetWorkspaceRequest() (request *GetWorkspaceRequest) {
    request = &GetWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "GetWorkspace")
    
    
    return
}

func NewGetWorkspaceResponse() (response *GetWorkspaceResponse) {
    response = &GetWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkspace
// 查询工作空间详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkspace(request *GetWorkspaceRequest) (response *GetWorkspaceResponse, err error) {
    return c.GetWorkspaceWithContext(context.Background(), request)
}

// GetWorkspace
// 查询工作空间详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TASKEXECUTIONNOTEXIST = "ResourceNotFound.TaskExecutionNotExist"
func (c *Client) GetWorkspaceWithContext(ctx context.Context, request *GetWorkspaceRequest) (response *GetWorkspaceResponse, err error) {
    if request == nil {
        request = NewGetWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "GetWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkspaceResponse()
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
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEENDELETED = "FailedOperation.WorkflowExecutionHasBeenDeleted"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMGROUPMISSINGERROR = "InvalidParameterValue.ParamGroupMissingError"
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
//  FAILEDOPERATION_WORKFLOWEXECUTIONHASBEENDELETED = "FailedOperation.WorkflowExecutionHasBeenDeleted"
//  FAILEDOPERATION_WORKFLOWNOPERMISSION = "FailedOperation.WorkflowNoPermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMBLANKERROR = "InvalidParameterValue.ParamBlankError"
//  INVALIDPARAMETERVALUE_PARAMGROUPMISSINGERROR = "InvalidParameterValue.ParamGroupMissingError"
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

func NewListCatalogsRequest() (request *ListCatalogsRequest) {
    request = &ListCatalogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListCatalogs")
    
    
    return
}

func NewListCatalogsResponse() (response *ListCatalogsResponse) {
    response = &ListCatalogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCatalogs
// 分页获取数据目录列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListCatalogs(request *ListCatalogsRequest) (response *ListCatalogsResponse, err error) {
    return c.ListCatalogsWithContext(context.Background(), request)
}

// ListCatalogs
// 分页获取数据目录列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListCatalogsWithContext(ctx context.Context, request *ListCatalogsRequest) (response *ListCatalogsResponse, err error) {
    if request == nil {
        request = NewListCatalogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListCatalogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCatalogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCatalogsResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleGroupUsersRequest() (request *ListConsoleGroupUsersRequest) {
    request = &ListConsoleGroupUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleGroupUsers")
    
    
    return
}

func NewListConsoleGroupUsersResponse() (response *ListConsoleGroupUsersResponse) {
    response = &ListConsoleGroupUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleGroupUsers
// 查询控制台用户组成员列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupUsers(request *ListConsoleGroupUsersRequest) (response *ListConsoleGroupUsersResponse, err error) {
    return c.ListConsoleGroupUsersWithContext(context.Background(), request)
}

// ListConsoleGroupUsers
// 查询控制台用户组成员列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupUsersWithContext(ctx context.Context, request *ListConsoleGroupUsersRequest) (response *ListConsoleGroupUsersResponse, err error) {
    if request == nil {
        request = NewListConsoleGroupUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleGroupUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleGroupUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleGroupUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleGroupsRequest() (request *ListConsoleGroupsRequest) {
    request = &ListConsoleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleGroups")
    
    
    return
}

func NewListConsoleGroupsResponse() (response *ListConsoleGroupsResponse) {
    response = &ListConsoleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleGroups
// 查询控制台用户组列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroups(request *ListConsoleGroupsRequest) (response *ListConsoleGroupsResponse, err error) {
    return c.ListConsoleGroupsWithContext(context.Background(), request)
}

// ListConsoleGroups
// 查询控制台用户组列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleGroupsWithContext(ctx context.Context, request *ListConsoleGroupsRequest) (response *ListConsoleGroupsResponse, err error) {
    if request == nil {
        request = NewListConsoleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleRolesRequest() (request *ListConsoleRolesRequest) {
    request = &ListConsoleRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleRoles")
    
    
    return
}

func NewListConsoleRolesResponse() (response *ListConsoleRolesResponse) {
    response = &ListConsoleRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleRoles
// 查询控制台角色列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleRoles(request *ListConsoleRolesRequest) (response *ListConsoleRolesResponse, err error) {
    return c.ListConsoleRolesWithContext(context.Background(), request)
}

// ListConsoleRoles
// 查询控制台角色列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleRolesWithContext(ctx context.Context, request *ListConsoleRolesRequest) (response *ListConsoleRolesResponse, err error) {
    if request == nil {
        request = NewListConsoleRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleRolesResponse()
    err = c.Send(request, response)
    return
}

func NewListConsoleUsersRequest() (request *ListConsoleUsersRequest) {
    request = &ListConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListConsoleUsers")
    
    
    return
}

func NewListConsoleUsersResponse() (response *ListConsoleUsersResponse) {
    response = &ListConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConsoleUsers
// 查询控制台用户列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleUsers(request *ListConsoleUsersRequest) (response *ListConsoleUsersResponse, err error) {
    return c.ListConsoleUsersWithContext(context.Background(), request)
}

// ListConsoleUsers
// 查询控制台用户列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListConsoleUsersWithContext(ctx context.Context, request *ListConsoleUsersRequest) (response *ListConsoleUsersResponse, err error) {
    if request == nil {
        request = NewListConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListFilesRequest() (request *ListFilesRequest) {
    request = &ListFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListFiles")
    
    
    return
}

func NewListFilesResponse() (response *ListFilesResponse) {
    response = &ListFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListFiles
// 获取文件夹和文件列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListFiles(request *ListFilesRequest) (response *ListFilesResponse, err error) {
    return c.ListFilesWithContext(context.Background(), request)
}

// ListFiles
// 获取文件夹和文件列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) ListFilesWithContext(ctx context.Context, request *ListFilesRequest) (response *ListFilesResponse, err error) {
    if request == nil {
        request = NewListFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListFilesResponse()
    err = c.Send(request, response)
    return
}

func NewListSchemasRequest() (request *ListSchemasRequest) {
    request = &ListSchemasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListSchemas")
    
    
    return
}

func NewListSchemasResponse() (response *ListSchemasResponse) {
    response = &ListSchemasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSchemas
// 获取schema列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) ListSchemas(request *ListSchemasRequest) (response *ListSchemasResponse, err error) {
    return c.ListSchemasWithContext(context.Background(), request)
}

// ListSchemas
// 获取schema列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
func (c *Client) ListSchemasWithContext(ctx context.Context, request *ListSchemasRequest) (response *ListSchemasResponse, err error) {
    if request == nil {
        request = NewListSchemasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListSchemas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSchemas require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSchemasResponse()
    err = c.Send(request, response)
    return
}

func NewListTablesRequest() (request *ListTablesRequest) {
    request = &ListTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListTables")
    
    
    return
}

func NewListTablesResponse() (response *ListTablesResponse) {
    response = &ListTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTables
// 获取table列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) ListTables(request *ListTablesRequest) (response *ListTablesResponse, err error) {
    return c.ListTablesWithContext(context.Background(), request)
}

// ListTables
// 获取table列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_CATALOGNOTFOUND = "ResourceNotFound.CatalogNotFound"
//  RESOURCENOTFOUND_SCHEMANOTFOUND = "ResourceNotFound.SchemaNotFound"
func (c *Client) ListTablesWithContext(ctx context.Context, request *ListTablesRequest) (response *ListTablesResponse, err error) {
    if request == nil {
        request = NewListTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTablesResponse()
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
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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

func NewListWorkspaceGroupUsersRequest() (request *ListWorkspaceGroupUsersRequest) {
    request = &ListWorkspaceGroupUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaceGroupUsers")
    
    
    return
}

func NewListWorkspaceGroupUsersResponse() (response *ListWorkspaceGroupUsersResponse) {
    response = &ListWorkspaceGroupUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaceGroupUsers
// 查询工作空间用户组用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceGroupUsers(request *ListWorkspaceGroupUsersRequest) (response *ListWorkspaceGroupUsersResponse, err error) {
    return c.ListWorkspaceGroupUsersWithContext(context.Background(), request)
}

// ListWorkspaceGroupUsers
// 查询工作空间用户组用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceGroupUsersWithContext(ctx context.Context, request *ListWorkspaceGroupUsersRequest) (response *ListWorkspaceGroupUsersResponse, err error) {
    if request == nil {
        request = NewListWorkspaceGroupUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaceGroupUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaceGroupUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspaceGroupUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspaceGroupsRequest() (request *ListWorkspaceGroupsRequest) {
    request = &ListWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaceGroups")
    
    
    return
}

func NewListWorkspaceGroupsResponse() (response *ListWorkspaceGroupsResponse) {
    response = &ListWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaceGroups
// 查询工作空间用户组列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceGroups(request *ListWorkspaceGroupsRequest) (response *ListWorkspaceGroupsResponse, err error) {
    return c.ListWorkspaceGroupsWithContext(context.Background(), request)
}

// ListWorkspaceGroups
// 查询工作空间用户组列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceGroupsWithContext(ctx context.Context, request *ListWorkspaceGroupsRequest) (response *ListWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewListWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspaceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspaceUsersRequest() (request *ListWorkspaceUsersRequest) {
    request = &ListWorkspaceUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaceUsers")
    
    
    return
}

func NewListWorkspaceUsersResponse() (response *ListWorkspaceUsersResponse) {
    response = &ListWorkspaceUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaceUsers
// 查询工作空间用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceUsers(request *ListWorkspaceUsersRequest) (response *ListWorkspaceUsersResponse, err error) {
    return c.ListWorkspaceUsersWithContext(context.Background(), request)
}

// ListWorkspaceUsers
// 查询工作空间用户列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaceUsersWithContext(ctx context.Context, request *ListWorkspaceUsersRequest) (response *ListWorkspaceUsersResponse, err error) {
    if request == nil {
        request = NewListWorkspaceUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaceUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaceUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspaceUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkspacesRequest() (request *ListWorkspacesRequest) {
    request = &ListWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "ListWorkspaces")
    
    
    return
}

func NewListWorkspacesResponse() (response *ListWorkspacesResponse) {
    response = &ListWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkspaces
// 查询工作空间列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspaces(request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
    return c.ListWorkspacesWithContext(context.Background(), request)
}

// ListWorkspaces
// 查询工作空间列表，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
    if request == nil {
        request = NewListWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "ListWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkspacesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveConsoleUsersRequest() (request *RemoveConsoleUsersRequest) {
    request = &RemoveConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RemoveConsoleUsers")
    
    
    return
}

func NewRemoveConsoleUsersResponse() (response *RemoveConsoleUsersResponse) {
    response = &RemoveConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveConsoleUsers
// <p>批量移除控制台用户（单次最多10个；前置校验任一不满足整体拒绝；执行阶段单个失败不中断后续删除，成败以 SuccessUins/FailItems 为准），该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。</p>
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) RemoveConsoleUsers(request *RemoveConsoleUsersRequest) (response *RemoveConsoleUsersResponse, err error) {
    return c.RemoveConsoleUsersWithContext(context.Background(), request)
}

// RemoveConsoleUsers
// <p>批量移除控制台用户（单次最多10个；前置校验任一不满足整体拒绝；执行阶段单个失败不中断后续删除，成败以 SuccessUins/FailItems 为准），该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。</p>
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) RemoveConsoleUsersWithContext(ctx context.Context, request *RemoveConsoleUsersRequest) (response *RemoveConsoleUsersResponse, err error) {
    if request == nil {
        request = NewRemoveConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RemoveConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveConsoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWorkspaceGroupsRequest() (request *RemoveWorkspaceGroupsRequest) {
    request = &RemoveWorkspaceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "RemoveWorkspaceGroups")
    
    
    return
}

func NewRemoveWorkspaceGroupsResponse() (response *RemoveWorkspaceGroupsResponse) {
    response = &RemoveWorkspaceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveWorkspaceGroups
// 移除工作空间用户组
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) RemoveWorkspaceGroups(request *RemoveWorkspaceGroupsRequest) (response *RemoveWorkspaceGroupsResponse, err error) {
    return c.RemoveWorkspaceGroupsWithContext(context.Background(), request)
}

// RemoveWorkspaceGroups
// 移除工作空间用户组
//
// 可能返回的错误码:
//  INVALIDPARAMETER_UGRUSERNOTEXIST = "InvalidParameter.UGRUserNotExist"
//  INVALIDPARAMETERVALUE_GROUPNOTEXIST = "InvalidParameterValue.GroupNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGROUPID = "InvalidParameterValue.InvalidGroupId"
//  UNSUPPORTEDOPERATION_DELETESYSTEMGROUPFAIL = "UnsupportedOperation.DeleteSystemGroupFail"
//  UNSUPPORTEDOPERATION_UPDATESYSTEMGROUPFAIL = "UnsupportedOperation.UpdateSystemGroupFail"
func (c *Client) RemoveWorkspaceGroupsWithContext(ctx context.Context, request *RemoveWorkspaceGroupsRequest) (response *RemoveWorkspaceGroupsResponse, err error) {
    if request == nil {
        request = NewRemoveWorkspaceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "RemoveWorkspaceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWorkspaceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWorkspaceGroupsResponse()
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

func NewStartComputeRequest() (request *StartComputeRequest) {
    request = &StartComputeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "StartCompute")
    
    
    return
}

func NewStartComputeResponse() (response *StartComputeResponse) {
    response = &StartComputeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCompute
// 启动计算资源
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StartCompute(request *StartComputeRequest) (response *StartComputeResponse, err error) {
    return c.StartComputeWithContext(context.Background(), request)
}

// StartCompute
// 启动计算资源
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StartComputeWithContext(ctx context.Context, request *StartComputeRequest) (response *StartComputeResponse, err error) {
    if request == nil {
        request = NewStartComputeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "StartCompute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCompute require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartComputeResponse()
    err = c.Send(request, response)
    return
}

func NewStopComputeRequest() (request *StopComputeRequest) {
    request = &StopComputeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "StopCompute")
    
    
    return
}

func NewStopComputeResponse() (response *StopComputeResponse) {
    response = &StopComputeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCompute
// 停止计算资源
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StopCompute(request *StopComputeRequest) (response *StopComputeResponse, err error) {
    return c.StopComputeWithContext(context.Background(), request)
}

// StopCompute
// 停止计算资源
//
// 可能返回的错误码:
//  INTERNALERROR_BIZEXCEPTION = "InternalError.BizException"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  INTERNALERROR_INVALIDPARAMETER = "InternalError.InvalidParameter"
//  INTERNALERROR_PERMISSIONDENIED = "InternalError.PermissionDenied"
//  INTERNALERROR_RESOURCENOTFOUND = "InternalError.ResourceNotFound"
func (c *Client) StopComputeWithContext(ctx context.Context, request *StopComputeRequest) (response *StopComputeResponse, err error) {
    if request == nil {
        request = NewStopComputeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "StopCompute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCompute require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopComputeResponse()
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

func NewUpdateConsoleGroupRequest() (request *UpdateConsoleGroupRequest) {
    request = &UpdateConsoleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateConsoleGroup")
    
    
    return
}

func NewUpdateConsoleGroupResponse() (response *UpdateConsoleGroupResponse) {
    response = &UpdateConsoleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConsoleGroup
// 修改控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
func (c *Client) UpdateConsoleGroup(request *UpdateConsoleGroupRequest) (response *UpdateConsoleGroupResponse, err error) {
    return c.UpdateConsoleGroupWithContext(context.Background(), request)
}

// UpdateConsoleGroup
// 修改控制台用户组，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
func (c *Client) UpdateConsoleGroupWithContext(ctx context.Context, request *UpdateConsoleGroupRequest) (response *UpdateConsoleGroupResponse, err error) {
    if request == nil {
        request = NewUpdateConsoleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateConsoleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConsoleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConsoleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConsoleUsersRequest() (request *UpdateConsoleUsersRequest) {
    request = &UpdateConsoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateConsoleUsers")
    
    
    return
}

func NewUpdateConsoleUsersResponse() (response *UpdateConsoleUsersResponse) {
    response = &UpdateConsoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConsoleUsers
// 修改控制台用户角色，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
func (c *Client) UpdateConsoleUsers(request *UpdateConsoleUsersRequest) (response *UpdateConsoleUsersResponse, err error) {
    return c.UpdateConsoleUsersWithContext(context.Background(), request)
}

// UpdateConsoleUsers
// 修改控制台用户角色，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
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
func (c *Client) UpdateConsoleUsersWithContext(ctx context.Context, request *UpdateConsoleUsersRequest) (response *UpdateConsoleUsersResponse, err error) {
    if request == nil {
        request = NewUpdateConsoleUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateConsoleUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConsoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConsoleUsersResponse()
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

func NewUpdateFolderRequest() (request *UpdateFolderRequest) {
    request = &UpdateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateFolder")
    
    
    return
}

func NewUpdateFolderResponse() (response *UpdateFolderResponse) {
    response = &UpdateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFolder
// 更新文件夹（支持重命名+移动）
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
func (c *Client) UpdateFolder(request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    return c.UpdateFolderWithContext(context.Background(), request)
}

// UpdateFolder
// 更新文件夹（支持重命名+移动）
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
func (c *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest) (response *UpdateFolderResponse, err error) {
    if request == nil {
        request = NewUpdateFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFolderResponse()
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
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
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

func NewUpdateWorkspaceRequest() (request *UpdateWorkspaceRequest) {
    request = &UpdateWorkspaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspace")
    
    
    return
}

func NewUpdateWorkspaceResponse() (response *UpdateWorkspaceResponse) {
    response = &UpdateWorkspaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspace
// 修改工作空间，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (response *UpdateWorkspaceResponse, err error) {
    return c.UpdateWorkspaceWithContext(context.Background(), request)
}

// UpdateWorkspace
// 修改工作空间，该接口为控制台级接口，仅支持在中心地域调用：国内站请传入 ap-guangzhou，国际站请传入 ap-singapore；其他地域调用将返回 UnsupportedRegion。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceWithContext(ctx context.Context, request *UpdateWorkspaceRequest) (response *UpdateWorkspaceResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspace require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceGroupsRolesRequest() (request *UpdateWorkspaceGroupsRolesRequest) {
    request = &UpdateWorkspaceGroupsRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspaceGroupsRoles")
    
    
    return
}

func NewUpdateWorkspaceGroupsRolesResponse() (response *UpdateWorkspaceGroupsRolesResponse) {
    response = &UpdateWorkspaceGroupsRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspaceGroupsRoles
// 更新工作空间用户组角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceGroupsRoles(request *UpdateWorkspaceGroupsRolesRequest) (response *UpdateWorkspaceGroupsRolesResponse, err error) {
    return c.UpdateWorkspaceGroupsRolesWithContext(context.Background(), request)
}

// UpdateWorkspaceGroupsRoles
// 更新工作空间用户组角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceGroupsRolesWithContext(ctx context.Context, request *UpdateWorkspaceGroupsRolesRequest) (response *UpdateWorkspaceGroupsRolesResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceGroupsRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspaceGroupsRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspaceGroupsRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceGroupsRolesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceRoleRequest() (request *UpdateWorkspaceRoleRequest) {
    request = &UpdateWorkspaceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspaceRole")
    
    
    return
}

func NewUpdateWorkspaceRoleResponse() (response *UpdateWorkspaceRoleResponse) {
    response = &UpdateWorkspaceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspaceRole
// 更新工作空间角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceRole(request *UpdateWorkspaceRoleRequest) (response *UpdateWorkspaceRoleResponse, err error) {
    return c.UpdateWorkspaceRoleWithContext(context.Background(), request)
}

// UpdateWorkspaceRole
// 更新工作空间角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceRoleWithContext(ctx context.Context, request *UpdateWorkspaceRoleRequest) (response *UpdateWorkspaceRoleResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspaceRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspaceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceUsersRolesRequest() (request *UpdateWorkspaceUsersRolesRequest) {
    request = &UpdateWorkspaceUsersRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("databuddy", APIVersion, "UpdateWorkspaceUsersRoles")
    
    
    return
}

func NewUpdateWorkspaceUsersRolesResponse() (response *UpdateWorkspaceUsersRolesResponse) {
    response = &UpdateWorkspaceUsersRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspaceUsersRoles
// 修改工作空间用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceUsersRoles(request *UpdateWorkspaceUsersRolesRequest) (response *UpdateWorkspaceUsersRolesResponse, err error) {
    return c.UpdateWorkspaceUsersRolesWithContext(context.Background(), request)
}

// UpdateWorkspaceUsersRoles
// 修改工作空间用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTHIRDPARTAPIERROR = "FailedOperation.CallThirdPartApiError"
//  FAILEDOPERATION_LABELCOUNTLIMIT = "FailedOperation.LabelCountLimit"
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
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTRESOURCEGROUP = "InvalidParameterValue.TaskTypeNotSupportResourceGroup"
//  INVALIDPARAMETERVALUE_TASKTYPEPROPERTYKEYVALUEREQUESTREQUIREDERROR = "InvalidParameterValue.TaskTypePropertyKeyValueRequestRequiredError"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEEXISTS = "InvalidParameterValue.WorkflowNameExists"
//  INVALIDPARAMETERVALUE_WORKFLOWNAMEINVALID = "InvalidParameterValue.WorkflowNameInvalid"
//  INVALIDPARAMETERVALUE_WORKFLOWSTARTTIMEAFTERENDTIMEERROR = "InvalidParameterValue.WorkflowStartTimeAfterEndTimeError"
//  INVALIDPARAMETERVALUE_WORKFLOWTRIGGERADVANCEDCONFIGERROR = "InvalidParameterValue.WorkflowTriggerAdvancedConfigError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKFLOWNOTEXIST = "ResourceNotFound.WorkflowNotExist"
//  RESOURCENOTFOUND_WORKFLOWNOTFOUND = "ResourceNotFound.WorkflowNotFound"
//  RESOURCENOTFOUND_WORKFLOWTRIGGERNOTFOUND = "ResourceNotFound.WorkflowTriggerNotFound"
func (c *Client) UpdateWorkspaceUsersRolesWithContext(ctx context.Context, request *UpdateWorkspaceUsersRolesRequest) (response *UpdateWorkspaceUsersRolesResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceUsersRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "databuddy", APIVersion, "UpdateWorkspaceUsersRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspaceUsersRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceUsersRolesResponse()
    err = c.Send(request, response)
    return
}
