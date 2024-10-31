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

package v20220601

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-06-01"

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


func NewCreateDeviceVirtualGroupRequest() (request *CreateDeviceVirtualGroupRequest) {
    request = &CreateDeviceVirtualGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "CreateDeviceVirtualGroup")
    
    
    return
}

func NewCreateDeviceVirtualGroupResponse() (response *CreateDeviceVirtualGroupResponse) {
    response = &CreateDeviceVirtualGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceVirtualGroup
// 创建终端自定义分组，私有化调用path为：/capi/Assets/Device/CreateDeviceVirtualGroup
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_AUTORULEPARAMETERERROR = "InvalidParameter.AutoRuleParameterError"
//  INVALIDPARAMETER_DUPLICATEDEVICEVIRTUALGROUPNAME = "InvalidParameter.DuplicateDeviceVirtualGroupName"
//  INVALIDPARAMETER_DUPLICATEIDINBLACKWHITELIST = "InvalidParameter.DuplicateIdInBlackWhiteList"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) CreateDeviceVirtualGroup(request *CreateDeviceVirtualGroupRequest) (response *CreateDeviceVirtualGroupResponse, err error) {
    return c.CreateDeviceVirtualGroupWithContext(context.Background(), request)
}

// CreateDeviceVirtualGroup
// 创建终端自定义分组，私有化调用path为：/capi/Assets/Device/CreateDeviceVirtualGroup
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_AUTORULEPARAMETERERROR = "InvalidParameter.AutoRuleParameterError"
//  INVALIDPARAMETER_DUPLICATEDEVICEVIRTUALGROUPNAME = "InvalidParameter.DuplicateDeviceVirtualGroupName"
//  INVALIDPARAMETER_DUPLICATEIDINBLACKWHITELIST = "InvalidParameter.DuplicateIdInBlackWhiteList"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) CreateDeviceVirtualGroupWithContext(ctx context.Context, request *CreateDeviceVirtualGroupRequest) (response *CreateDeviceVirtualGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeviceVirtualGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceVirtualGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceVirtualGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountGroupsRequest() (request *DescribeAccountGroupsRequest) {
    request = &DescribeAccountGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeAccountGroups")
    
    
    return
}

func NewDescribeAccountGroupsResponse() (response *DescribeAccountGroupsResponse) {
    response = &DescribeAccountGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountGroups
// 以分页的方式查询账号分组列表，私有化调用path为：/capi/Assets/DescribeAccountGroups
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountGroups(request *DescribeAccountGroupsRequest) (response *DescribeAccountGroupsResponse, err error) {
    return c.DescribeAccountGroupsWithContext(context.Background(), request)
}

// DescribeAccountGroups
// 以分页的方式查询账号分组列表，私有化调用path为：/capi/Assets/DescribeAccountGroups
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountGroupsWithContext(ctx context.Context, request *DescribeAccountGroupsRequest) (response *DescribeAccountGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDevices")
    
    
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDevices
// 查询满足条件的终端数据详情，私有化调用path为：/capi/Assets/Device/DescribeDevices
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 查询满足条件的终端数据详情，私有化调用path为：/capi/Assets/Device/DescribeDevices
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLocalAccountsRequest() (request *DescribeLocalAccountsRequest) {
    request = &DescribeLocalAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeLocalAccounts")
    
    
    return
}

func NewDescribeLocalAccountsResponse() (response *DescribeLocalAccountsResponse) {
    response = &DescribeLocalAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLocalAccounts
// 获取账号列表，支持分页，模糊搜索，私有化调用path为：/capi/Assets/Account/DescribeLocalAccounts
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
func (c *Client) DescribeLocalAccounts(request *DescribeLocalAccountsRequest) (response *DescribeLocalAccountsResponse, err error) {
    return c.DescribeLocalAccountsWithContext(context.Background(), request)
}

// DescribeLocalAccounts
// 获取账号列表，支持分页，模糊搜索，私有化调用path为：/capi/Assets/Account/DescribeLocalAccounts
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
func (c *Client) DescribeLocalAccountsWithContext(ctx context.Context, request *DescribeLocalAccountsRequest) (response *DescribeLocalAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeLocalAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLocalAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLocalAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRootAccountGroupRequest() (request *DescribeRootAccountGroupRequest) {
    request = &DescribeRootAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeRootAccountGroup")
    
    
    return
}

func NewDescribeRootAccountGroupResponse() (response *DescribeRootAccountGroupResponse) {
    response = &DescribeRootAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRootAccountGroup
// 查询账号根分组详情。对应“用户与授权管理”里内置不可见的全网根账号组，所有新建的目录，都挂在该全网根账号组下。
//
// 私有化调用path为：capi/Assets/DescribeRootAccountGroup
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRootAccountGroup(request *DescribeRootAccountGroupRequest) (response *DescribeRootAccountGroupResponse, err error) {
    return c.DescribeRootAccountGroupWithContext(context.Background(), request)
}

// DescribeRootAccountGroup
// 查询账号根分组详情。对应“用户与授权管理”里内置不可见的全网根账号组，所有新建的目录，都挂在该全网根账号组下。
//
// 私有化调用path为：capi/Assets/DescribeRootAccountGroup
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRootAccountGroupWithContext(ctx context.Context, request *DescribeRootAccountGroupRequest) (response *DescribeRootAccountGroupResponse, err error) {
    if request == nil {
        request = NewDescribeRootAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRootAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRootAccountGroupResponse()
    err = c.Send(request, response)
    return
}
