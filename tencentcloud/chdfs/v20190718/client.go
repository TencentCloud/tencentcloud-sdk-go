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

package v20190718

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-18"

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


func NewCreateAccessGroupRequest() (request *CreateAccessGroupRequest) {
    request = &CreateAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessGroup")
    
    
    return
}

func NewCreateAccessGroupResponse() (response *CreateAccessGroupResponse) {
    response = &CreateAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
    return c.CreateAccessGroupWithContext(context.Background(), request)
}

// CreateAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessGroupWithContext(ctx context.Context, request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccessGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessRulesRequest() (request *CreateAccessRulesRequest) {
    request = &CreateAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessRules")
    
    
    return
}

func NewCreateAccessRulesResponse() (response *CreateAccessRulesResponse) {
    response = &CreateAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建权限规则，权限规则ID和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessRules(request *CreateAccessRulesRequest) (response *CreateAccessRulesResponse, err error) {
    return c.CreateAccessRulesWithContext(context.Background(), request)
}

// CreateAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建权限规则，权限规则ID和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessRulesWithContext(ctx context.Context, request *CreateAccessRulesRequest) (response *CreateAccessRulesResponse, err error) {
    if request == nil {
        request = NewCreateAccessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
    request = &CreateFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateFileSystem")
    
    
    return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
    response = &CreateFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建文件系统（异步）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    return c.CreateFileSystemWithContext(context.Background(), request)
}

// CreateFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建文件系统（异步）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifeCycleRulesRequest() (request *CreateLifeCycleRulesRequest) {
    request = &CreateLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateLifeCycleRules")
    
    
    return
}

func NewCreateLifeCycleRulesResponse() (response *CreateLifeCycleRulesResponse) {
    response = &CreateLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建生命周期规则，生命周期规则ID和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLifeCycleRules(request *CreateLifeCycleRulesRequest) (response *CreateLifeCycleRulesResponse, err error) {
    return c.CreateLifeCycleRulesWithContext(context.Background(), request)
}

// CreateLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建生命周期规则，生命周期规则ID和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLifeCycleRulesWithContext(ctx context.Context, request *CreateLifeCycleRulesRequest) (response *CreateLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewCreateLifeCycleRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLifeCycleRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMountPointRequest() (request *CreateMountPointRequest) {
    request = &CreateMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateMountPoint")
    
    
    return
}

func NewCreateMountPointResponse() (response *CreateMountPointResponse) {
    response = &CreateMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建文件系统挂载点，仅限于创建成功的文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCENOTFOUND_VPCNOTEXISTS = "ResourceNotFound.VpcNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMountPoint(request *CreateMountPointRequest) (response *CreateMountPointResponse, err error) {
    return c.CreateMountPointWithContext(context.Background(), request)
}

// CreateMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 创建文件系统挂载点，仅限于创建成功的文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCENOTFOUND_VPCNOTEXISTS = "ResourceNotFound.VpcNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMountPointWithContext(ctx context.Context, request *CreateMountPointRequest) (response *CreateMountPointResponse, err error) {
    if request == nil {
        request = NewCreateMountPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMountPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRestoreTasksRequest() (request *CreateRestoreTasksRequest) {
    request = &CreateRestoreTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateRestoreTasks")
    
    
    return
}

func NewCreateRestoreTasksResponse() (response *CreateRestoreTasksResponse) {
    response = &CreateRestoreTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRestoreTasks
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建回热任务，回热任务ID、状态和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRestoreTasks(request *CreateRestoreTasksRequest) (response *CreateRestoreTasksResponse, err error) {
    return c.CreateRestoreTasksWithContext(context.Background(), request)
}

// CreateRestoreTasks
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量创建回热任务，回热任务ID、状态和创建时间无需填写。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRestoreTasksWithContext(ctx context.Context, request *CreateRestoreTasksRequest) (response *CreateRestoreTasksResponse, err error) {
    if request == nil {
        request = NewCreateRestoreTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRestoreTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessGroupRequest() (request *DeleteAccessGroupRequest) {
    request = &DeleteAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessGroup")
    
    
    return
}

func NewDeleteAccessGroupResponse() (response *DeleteAccessGroupResponse) {
    response = &DeleteAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (response *DeleteAccessGroupResponse, err error) {
    return c.DeleteAccessGroupWithContext(context.Background(), request)
}

// DeleteAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessGroupWithContext(ctx context.Context, request *DeleteAccessGroupRequest) (response *DeleteAccessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccessGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessRulesRequest() (request *DeleteAccessRulesRequest) {
    request = &DeleteAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessRules")
    
    
    return
}

func NewDeleteAccessRulesResponse() (response *DeleteAccessRulesResponse) {
    response = &DeleteAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量删除权限规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessRules(request *DeleteAccessRulesRequest) (response *DeleteAccessRulesResponse, err error) {
    return c.DeleteAccessRulesWithContext(context.Background(), request)
}

// DeleteAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量删除权限规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessRulesWithContext(ctx context.Context, request *DeleteAccessRulesRequest) (response *DeleteAccessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
    request = &DeleteFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteFileSystem")
    
    
    return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
    response = &DeleteFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除文件系统，不允许删除非空文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILESYSTEMNOTEMPTY = "FailedOperation.FileSystemNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    return c.DeleteFileSystemWithContext(context.Background(), request)
}

// DeleteFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除文件系统，不允许删除非空文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILESYSTEMNOTEMPTY = "FailedOperation.FileSystemNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLifeCycleRulesRequest() (request *DeleteLifeCycleRulesRequest) {
    request = &DeleteLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteLifeCycleRules")
    
    
    return
}

func NewDeleteLifeCycleRulesResponse() (response *DeleteLifeCycleRulesResponse) {
    response = &DeleteLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量删除生命周期规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLifeCycleRules(request *DeleteLifeCycleRulesRequest) (response *DeleteLifeCycleRulesResponse, err error) {
    return c.DeleteLifeCycleRulesWithContext(context.Background(), request)
}

// DeleteLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量删除生命周期规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLifeCycleRulesWithContext(ctx context.Context, request *DeleteLifeCycleRulesRequest) (response *DeleteLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDeleteLifeCycleRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLifeCycleRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMountPointRequest() (request *DeleteMountPointRequest) {
    request = &DeleteMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteMountPoint")
    
    
    return
}

func NewDeleteMountPointResponse() (response *DeleteMountPointResponse) {
    response = &DeleteMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除挂载点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMountPoint(request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
    return c.DeleteMountPointWithContext(context.Background(), request)
}

// DeleteMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 删除挂载点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMountPointWithContext(ctx context.Context, request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
    if request == nil {
        request = NewDeleteMountPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMountPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessGroupsRequest() (request *DescribeAccessGroupsRequest) {
    request = &DescribeAccessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessGroups")
    
    
    return
}

func NewDescribeAccessGroupsResponse() (response *DescribeAccessGroupsResponse) {
    response = &DescribeAccessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessGroups
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看权限组列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
    return c.DescribeAccessGroupsWithContext(context.Background(), request)
}

// DescribeAccessGroups
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看权限组列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroupsWithContext(ctx context.Context, request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRulesRequest() (request *DescribeAccessRulesRequest) {
    request = &DescribeAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessRules")
    
    
    return
}

func NewDescribeAccessRulesResponse() (response *DescribeAccessRulesResponse) {
    response = &DescribeAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过权限组ID查看权限规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
    return c.DescribeAccessRulesWithContext(context.Background(), request)
}

// DescribeAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过权限组ID查看权限规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessRulesWithContext(ctx context.Context, request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemRequest() (request *DescribeFileSystemRequest) {
    request = &DescribeFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystem")
    
    
    return
}

func NewDescribeFileSystemResponse() (response *DescribeFileSystemResponse) {
    response = &DescribeFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看文件系统详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystem(request *DescribeFileSystemRequest) (response *DescribeFileSystemResponse, err error) {
    return c.DescribeFileSystemWithContext(context.Background(), request)
}

// DescribeFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看文件系统详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystemWithContext(ctx context.Context, request *DescribeFileSystemRequest) (response *DescribeFileSystemResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
    request = &DescribeFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystems")
    
    
    return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
    response = &DescribeFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileSystems
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看文件系统列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    return c.DescribeFileSystemsWithContext(context.Background(), request)
}

// DescribeFileSystems
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看文件系统列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileSystems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifeCycleRulesRequest() (request *DescribeLifeCycleRulesRequest) {
    request = &DescribeLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeLifeCycleRules")
    
    
    return
}

func NewDescribeLifeCycleRulesResponse() (response *DescribeLifeCycleRulesResponse) {
    response = &DescribeLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看生命周期规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLifeCycleRules(request *DescribeLifeCycleRulesRequest) (response *DescribeLifeCycleRulesResponse, err error) {
    return c.DescribeLifeCycleRulesWithContext(context.Background(), request)
}

// DescribeLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看生命周期规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLifeCycleRulesWithContext(ctx context.Context, request *DescribeLifeCycleRulesRequest) (response *DescribeLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLifeCycleRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLifeCycleRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointRequest() (request *DescribeMountPointRequest) {
    request = &DescribeMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoint")
    
    
    return
}

func NewDescribeMountPointResponse() (response *DescribeMountPointResponse) {
    response = &DescribeMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看挂载点详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPoint(request *DescribeMountPointRequest) (response *DescribeMountPointResponse, err error) {
    return c.DescribeMountPointWithContext(context.Background(), request)
}

// DescribeMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 查看挂载点详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPointWithContext(ctx context.Context, request *DescribeMountPointRequest) (response *DescribeMountPointResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMountPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointsRequest() (request *DescribeMountPointsRequest) {
    request = &DescribeMountPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoints")
    
    
    return
}

func NewDescribeMountPointsResponse() (response *DescribeMountPointsResponse) {
    response = &DescribeMountPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountPoints
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID或者权限组ID查看挂载点列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPoints(request *DescribeMountPointsRequest) (response *DescribeMountPointsResponse, err error) {
    return c.DescribeMountPointsWithContext(context.Background(), request)
}

// DescribeMountPoints
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID或者权限组ID查看挂载点列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPointsWithContext(ctx context.Context, request *DescribeMountPointsRequest) (response *DescribeMountPointsResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMountPoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMountPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsRequest() (request *DescribeResourceTagsRequest) {
    request = &DescribeResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeResourceTags")
    
    
    return
}

func NewDescribeResourceTagsResponse() (response *DescribeResourceTagsResponse) {
    response = &DescribeResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceTags
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看资源标签列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceTags(request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    return c.DescribeResourceTagsWithContext(context.Background(), request)
}

// DescribeResourceTags
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看资源标签列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceTagsWithContext(ctx context.Context, request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRestoreTasksRequest() (request *DescribeRestoreTasksRequest) {
    request = &DescribeRestoreTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeRestoreTasks")
    
    
    return
}

func NewDescribeRestoreTasksResponse() (response *DescribeRestoreTasksResponse) {
    response = &DescribeRestoreTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRestoreTasks
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看回热任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRestoreTasks(request *DescribeRestoreTasksRequest) (response *DescribeRestoreTasksResponse, err error) {
    return c.DescribeRestoreTasksWithContext(context.Background(), request)
}

// DescribeRestoreTasks
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 通过文件系统ID查看回热任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRestoreTasksWithContext(ctx context.Context, request *DescribeRestoreTasksRequest) (response *DescribeRestoreTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRestoreTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessGroupRequest() (request *ModifyAccessGroupRequest) {
    request = &ModifyAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessGroup")
    
    
    return
}

func NewModifyAccessGroupResponse() (response *ModifyAccessGroupResponse) {
    response = &ModifyAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改权限组属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
    return c.ModifyAccessGroupWithContext(context.Background(), request)
}

// ModifyAccessGroup
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改权限组属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessGroupWithContext(ctx context.Context, request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccessGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessRulesRequest() (request *ModifyAccessRulesRequest) {
    request = &ModifyAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessRules")
    
    
    return
}

func NewModifyAccessRulesResponse() (response *ModifyAccessRulesResponse) {
    response = &ModifyAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量修改权限规则属性，需要指定权限规则ID，支持修改权限规则地址、访问模式和优先级。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessRules(request *ModifyAccessRulesRequest) (response *ModifyAccessRulesResponse, err error) {
    return c.ModifyAccessRulesWithContext(context.Background(), request)
}

// ModifyAccessRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量修改权限规则属性，需要指定权限规则ID，支持修改权限规则地址、访问模式和优先级。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessRulesWithContext(ctx context.Context, request *ModifyAccessRulesRequest) (response *ModifyAccessRulesResponse, err error) {
    if request == nil {
        request = NewModifyAccessRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
    request = &ModifyFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyFileSystem")
    
    
    return
}

func NewModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
    response = &ModifyFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改文件系统属性，仅限于创建成功的文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUOTALESSTHANCURRENTUSED = "FailedOperation.QuotaLessThanCurrentUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
    return c.ModifyFileSystemWithContext(context.Background(), request)
}

// ModifyFileSystem
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改文件系统属性，仅限于创建成功的文件系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUOTALESSTHANCURRENTUSED = "FailedOperation.QuotaLessThanCurrentUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
    if request == nil {
        request = NewModifyFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLifeCycleRulesRequest() (request *ModifyLifeCycleRulesRequest) {
    request = &ModifyLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyLifeCycleRules")
    
    
    return
}

func NewModifyLifeCycleRulesResponse() (response *ModifyLifeCycleRulesResponse) {
    response = &ModifyLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量修改生命周期规则属性，需要指定生命周期规则ID，支持修改生命周期规则名称、路径、转换列表和状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLifeCycleRules(request *ModifyLifeCycleRulesRequest) (response *ModifyLifeCycleRulesResponse, err error) {
    return c.ModifyLifeCycleRulesWithContext(context.Background(), request)
}

// ModifyLifeCycleRules
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 批量修改生命周期规则属性，需要指定生命周期规则ID，支持修改生命周期规则名称、路径、转换列表和状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLifeCycleRulesWithContext(ctx context.Context, request *ModifyLifeCycleRulesRequest) (response *ModifyLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewModifyLifeCycleRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLifeCycleRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMountPointRequest() (request *ModifyMountPointRequest) {
    request = &ModifyMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyMountPoint")
    
    
    return
}

func NewModifyMountPointResponse() (response *ModifyMountPointResponse) {
    response = &ModifyMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改挂载点属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMountPoint(request *ModifyMountPointRequest) (response *ModifyMountPointResponse, err error) {
    return c.ModifyMountPointWithContext(context.Background(), request)
}

// ModifyMountPoint
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改挂载点属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMountPointWithContext(ctx context.Context, request *ModifyMountPointRequest) (response *ModifyMountPointResponse, err error) {
    if request == nil {
        request = NewModifyMountPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMountPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceTagsRequest() (request *ModifyResourceTagsRequest) {
    request = &ModifyResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyResourceTags")
    
    
    return
}

func NewModifyResourceTagsResponse() (response *ModifyResourceTagsResponse) {
    response = &ModifyResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyResourceTags
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改资源标签列表，全量覆盖。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyResourceTags(request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    return c.ModifyResourceTagsWithContext(context.Background(), request)
}

// ModifyResourceTags
// 云API旧版本2019-07-18预下线，所有功能由新版本2020-11-12替代，目前云API主要用作控制台使用。
//
// 
//
// 修改资源标签列表，全量覆盖。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyResourceTagsWithContext(ctx context.Context, request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourceTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceTagsResponse()
    err = c.Send(request, response)
    return
}
