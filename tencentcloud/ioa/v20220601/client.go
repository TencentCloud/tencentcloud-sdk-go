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


func NewCreateDLPFileDetectionTaskRequest() (request *CreateDLPFileDetectionTaskRequest) {
    request = &CreateDLPFileDetectionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "CreateDLPFileDetectionTask")
    
    
    return
}

func NewCreateDLPFileDetectionTaskResponse() (response *CreateDLPFileDetectionTaskResponse) {
    response = &CreateDLPFileDetectionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDLPFileDetectionTask
// 提交送检任务
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDLPFileDetectionTask(request *CreateDLPFileDetectionTaskRequest) (response *CreateDLPFileDetectionTaskResponse, err error) {
    return c.CreateDLPFileDetectionTaskWithContext(context.Background(), request)
}

// CreateDLPFileDetectionTask
// 提交送检任务
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDLPFileDetectionTaskWithContext(ctx context.Context, request *CreateDLPFileDetectionTaskRequest) (response *CreateDLPFileDetectionTaskResponse, err error) {
    if request == nil {
        request = NewCreateDLPFileDetectionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDLPFileDetectionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDLPFileDetectionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceTaskRequest() (request *CreateDeviceTaskRequest) {
    request = &CreateDeviceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "CreateDeviceTask")
    
    
    return
}

func NewCreateDeviceTaskResponse() (response *CreateDeviceTaskResponse) {
    response = &CreateDeviceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceTask
// 创建获取终端进程网络服务信息任务，私有化调用path为：capi/Assets/Device/DescribeDeviceInfo
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  FAILEDOPERATION_RPCSERVICECALLFAILED = "FailedOperation.RPCServiceCallFailed"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDeviceTask(request *CreateDeviceTaskRequest) (response *CreateDeviceTaskResponse, err error) {
    return c.CreateDeviceTaskWithContext(context.Background(), request)
}

// CreateDeviceTask
// 创建获取终端进程网络服务信息任务，私有化调用path为：capi/Assets/Device/DescribeDeviceInfo
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  FAILEDOPERATION_RPCSERVICECALLFAILED = "FailedOperation.RPCServiceCallFailed"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDeviceTaskWithContext(ctx context.Context, request *CreateDeviceTaskRequest) (response *CreateDeviceTaskResponse, err error) {
    if request == nil {
        request = NewCreateDeviceTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceTaskResponse()
    err = c.Send(request, response)
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

func NewCreatePrivilegeCodeRequest() (request *CreatePrivilegeCodeRequest) {
    request = &CreatePrivilegeCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "CreatePrivilegeCode")
    
    
    return
}

func NewCreatePrivilegeCodeResponse() (response *CreatePrivilegeCodeResponse) {
    response = &CreatePrivilegeCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivilegeCode
// 生成特权码，私有化调用path为：capi/Assets/Device/CreatePrivilegeCode，生成的特权码、卸载码，仅对该设备当天有效
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePrivilegeCode(request *CreatePrivilegeCodeRequest) (response *CreatePrivilegeCodeResponse, err error) {
    return c.CreatePrivilegeCodeWithContext(context.Background(), request)
}

// CreatePrivilegeCode
// 生成特权码，私有化调用path为：capi/Assets/Device/CreatePrivilegeCode，生成的特权码、卸载码，仅对该设备当天有效
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePrivilegeCodeWithContext(ctx context.Context, request *CreatePrivilegeCodeRequest) (response *CreatePrivilegeCodeResponse, err error) {
    if request == nil {
        request = NewCreatePrivilegeCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivilegeCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivilegeCodeResponse()
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

func NewDescribeDLPFileDetectResultRequest() (request *DescribeDLPFileDetectResultRequest) {
    request = &DescribeDLPFileDetectResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDLPFileDetectResult")
    
    
    return
}

func NewDescribeDLPFileDetectResultResponse() (response *DescribeDLPFileDetectResultResponse) {
    response = &DescribeDLPFileDetectResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLPFileDetectResult
// webservice查询文件检测结果
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDLPFileDetectResult(request *DescribeDLPFileDetectResultRequest) (response *DescribeDLPFileDetectResultResponse, err error) {
    return c.DescribeDLPFileDetectResultWithContext(context.Background(), request)
}

// DescribeDLPFileDetectResult
// webservice查询文件检测结果
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDLPFileDetectResultWithContext(ctx context.Context, request *DescribeDLPFileDetectResultRequest) (response *DescribeDLPFileDetectResultResponse, err error) {
    if request == nil {
        request = NewDescribeDLPFileDetectResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLPFileDetectResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLPFileDetectResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceHardwareInfoListRequest() (request *DescribeDeviceHardwareInfoListRequest) {
    request = &DescribeDeviceHardwareInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDeviceHardwareInfoList")
    
    
    return
}

func NewDescribeDeviceHardwareInfoListResponse() (response *DescribeDeviceHardwareInfoListResponse) {
    response = &DescribeDeviceHardwareInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceHardwareInfoList
// 查询满足条件的查询终端硬件信息列表，私有化调用path为：/capi/Assets/Device/DescribeDeviceHardwareInfoList
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceHardwareInfoList(request *DescribeDeviceHardwareInfoListRequest) (response *DescribeDeviceHardwareInfoListResponse, err error) {
    return c.DescribeDeviceHardwareInfoListWithContext(context.Background(), request)
}

// DescribeDeviceHardwareInfoList
// 查询满足条件的查询终端硬件信息列表，私有化调用path为：/capi/Assets/Device/DescribeDeviceHardwareInfoList
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceHardwareInfoListWithContext(ctx context.Context, request *DescribeDeviceHardwareInfoListRequest) (response *DescribeDeviceHardwareInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceHardwareInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceHardwareInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceHardwareInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceInfoRequest() (request *DescribeDeviceInfoRequest) {
    request = &DescribeDeviceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDeviceInfo")
    
    
    return
}

func NewDescribeDeviceInfoResponse() (response *DescribeDeviceInfoResponse) {
    response = &DescribeDeviceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceInfo
// 获取终端进程网络服务信息，私有化调用path为：capi/Assets/Device/DescribeDeviceInfo
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (response *DescribeDeviceInfoResponse, err error) {
    return c.DescribeDeviceInfoWithContext(context.Background(), request)
}

// DescribeDeviceInfo
// 获取终端进程网络服务信息，私有化调用path为：capi/Assets/Device/DescribeDeviceInfo
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceInfoWithContext(ctx context.Context, request *DescribeDeviceInfoRequest) (response *DescribeDeviceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceInfoResponse()
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

func NewDescribeSoftCensusListByDeviceRequest() (request *DescribeSoftCensusListByDeviceRequest) {
    request = &DescribeSoftCensusListByDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeSoftCensusListByDevice")
    
    
    return
}

func NewDescribeSoftCensusListByDeviceResponse() (response *DescribeSoftCensusListByDeviceResponse) {
    response = &DescribeSoftCensusListByDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSoftCensusListByDevice
// 查看终端树下的软件列表,私有化调用path为：capi/Software/DescribeSoftCensusListByDevice
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeSoftCensusListByDevice(request *DescribeSoftCensusListByDeviceRequest) (response *DescribeSoftCensusListByDeviceResponse, err error) {
    return c.DescribeSoftCensusListByDeviceWithContext(context.Background(), request)
}

// DescribeSoftCensusListByDevice
// 查看终端树下的软件列表,私有化调用path为：capi/Software/DescribeSoftCensusListByDevice
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeSoftCensusListByDeviceWithContext(ctx context.Context, request *DescribeSoftCensusListByDeviceRequest) (response *DescribeSoftCensusListByDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeSoftCensusListByDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSoftCensusListByDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSoftCensusListByDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSoftwareInformationRequest() (request *DescribeSoftwareInformationRequest) {
    request = &DescribeSoftwareInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeSoftwareInformation")
    
    
    return
}

func NewDescribeSoftwareInformationResponse() (response *DescribeSoftwareInformationResponse) {
    response = &DescribeSoftwareInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSoftwareInformation
// 查看指定终端的软件详情列表,私有化调用path为：capi/Software/DescribeSoftwareInformation
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeSoftwareInformation(request *DescribeSoftwareInformationRequest) (response *DescribeSoftwareInformationResponse, err error) {
    return c.DescribeSoftwareInformationWithContext(context.Background(), request)
}

// DescribeSoftwareInformation
// 查看指定终端的软件详情列表,私有化调用path为：capi/Software/DescribeSoftwareInformation
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeSoftwareInformationWithContext(ctx context.Context, request *DescribeSoftwareInformationRequest) (response *DescribeSoftwareInformationResponse, err error) {
    if request == nil {
        request = NewDescribeSoftwareInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSoftwareInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSoftwareInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVirtualDevicesRequest() (request *DescribeVirtualDevicesRequest) {
    request = &DescribeVirtualDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeVirtualDevices")
    
    
    return
}

func NewDescribeVirtualDevicesResponse() (response *DescribeVirtualDevicesResponse) {
    response = &DescribeVirtualDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVirtualDevices
// 展示自定义分组终端列表，私有化调用path为：/capi/Assets/DescribeVirtualDevices
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeVirtualDevices(request *DescribeVirtualDevicesRequest) (response *DescribeVirtualDevicesResponse, err error) {
    return c.DescribeVirtualDevicesWithContext(context.Background(), request)
}

// DescribeVirtualDevices
// 展示自定义分组终端列表，私有化调用path为：/capi/Assets/DescribeVirtualDevices
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeVirtualDevicesWithContext(ctx context.Context, request *DescribeVirtualDevicesRequest) (response *DescribeVirtualDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeVirtualDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirtualDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirtualDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVirtualDeviceGroupsRequest() (request *ModifyVirtualDeviceGroupsRequest) {
    request = &ModifyVirtualDeviceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "ModifyVirtualDeviceGroups")
    
    
    return
}

func NewModifyVirtualDeviceGroupsResponse() (response *ModifyVirtualDeviceGroupsResponse) {
    response = &ModifyVirtualDeviceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVirtualDeviceGroups
// 终端自定义分组增减终端，私有化调用path为：/capi/Assets/Device/ModifyVirtualDeviceGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_DATABASEQUERYFAILED = "InternalError.DatabaseQueryFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  INVALIDPARAMETERVALUE_VIRTUALDEVICEGROUPNOTFOUND = "InvalidParameterValue.VirtualDeviceGroupNotFound"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  REQUESTLIMITEXCEEDED_WAITFORTHELASTOPERATIONTOCOMPLETE = "RequestLimitExceeded.WaitForTheLastOperationToComplete"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) ModifyVirtualDeviceGroups(request *ModifyVirtualDeviceGroupsRequest) (response *ModifyVirtualDeviceGroupsResponse, err error) {
    return c.ModifyVirtualDeviceGroupsWithContext(context.Background(), request)
}

// ModifyVirtualDeviceGroups
// 终端自定义分组增减终端，私有化调用path为：/capi/Assets/Device/ModifyVirtualDeviceGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_DATABASEQUERYFAILED = "InternalError.DatabaseQueryFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  INVALIDPARAMETERVALUE_VIRTUALDEVICEGROUPNOTFOUND = "InvalidParameterValue.VirtualDeviceGroupNotFound"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  REQUESTLIMITEXCEEDED_WAITFORTHELASTOPERATIONTOCOMPLETE = "RequestLimitExceeded.WaitForTheLastOperationToComplete"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) ModifyVirtualDeviceGroupsWithContext(ctx context.Context, request *ModifyVirtualDeviceGroupsRequest) (response *ModifyVirtualDeviceGroupsResponse, err error) {
    if request == nil {
        request = NewModifyVirtualDeviceGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirtualDeviceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirtualDeviceGroupsResponse()
    err = c.Send(request, response)
    return
}
