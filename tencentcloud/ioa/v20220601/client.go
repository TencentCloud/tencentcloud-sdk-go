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


func NewCreateDLPFileDetectTaskRequest() (request *CreateDLPFileDetectTaskRequest) {
    request = &CreateDLPFileDetectTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "CreateDLPFileDetectTask")
    
    
    return
}

func NewCreateDLPFileDetectTaskResponse() (response *CreateDLPFileDetectTaskResponse) {
    response = &CreateDLPFileDetectTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDLPFileDetectTask
// 创建文件鉴定任务，私有化调用path为：capi/DlpOpenApi/CreateDLPFileDetectTask
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_IDENTICALNAME = "InvalidParameter.IdenticalName"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) CreateDLPFileDetectTask(request *CreateDLPFileDetectTaskRequest) (response *CreateDLPFileDetectTaskResponse, err error) {
    return c.CreateDLPFileDetectTaskWithContext(context.Background(), request)
}

// CreateDLPFileDetectTask
// 创建文件鉴定任务，私有化调用path为：capi/DlpOpenApi/CreateDLPFileDetectTask
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_IDENTICALNAME = "InvalidParameter.IdenticalName"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) CreateDLPFileDetectTaskWithContext(ctx context.Context, request *CreateDLPFileDetectTaskRequest) (response *CreateDLPFileDetectTaskResponse, err error) {
    if request == nil {
        request = NewCreateDLPFileDetectTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "CreateDLPFileDetectTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDLPFileDetectTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDLPFileDetectTaskResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "CreateDLPFileDetectionTask")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "CreateDeviceTask")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "CreateDeviceVirtualGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "CreatePrivilegeCode")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeAccountGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggrSoftCategorySoftListRequest() (request *DescribeAggrSoftCategorySoftListRequest) {
    request = &DescribeAggrSoftCategorySoftListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeAggrSoftCategorySoftList")
    
    
    return
}

func NewDescribeAggrSoftCategorySoftListResponse() (response *DescribeAggrSoftCategorySoftListResponse) {
    response = &DescribeAggrSoftCategorySoftListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggrSoftCategorySoftList
// 聚合的分类软件列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftCategorySoftList(request *DescribeAggrSoftCategorySoftListRequest) (response *DescribeAggrSoftCategorySoftListResponse, err error) {
    return c.DescribeAggrSoftCategorySoftListWithContext(context.Background(), request)
}

// DescribeAggrSoftCategorySoftList
// 聚合的分类软件列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftCategorySoftListWithContext(ctx context.Context, request *DescribeAggrSoftCategorySoftListRequest) (response *DescribeAggrSoftCategorySoftListResponse, err error) {
    if request == nil {
        request = NewDescribeAggrSoftCategorySoftListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeAggrSoftCategorySoftList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggrSoftCategorySoftList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggrSoftCategorySoftListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggrSoftDetailRequest() (request *DescribeAggrSoftDetailRequest) {
    request = &DescribeAggrSoftDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeAggrSoftDetail")
    
    
    return
}

func NewDescribeAggrSoftDetailResponse() (response *DescribeAggrSoftDetailResponse) {
    response = &DescribeAggrSoftDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggrSoftDetail
// 聚合的软件详情
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftDetail(request *DescribeAggrSoftDetailRequest) (response *DescribeAggrSoftDetailResponse, err error) {
    return c.DescribeAggrSoftDetailWithContext(context.Background(), request)
}

// DescribeAggrSoftDetail
// 聚合的软件详情
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftDetailWithContext(ctx context.Context, request *DescribeAggrSoftDetailRequest) (response *DescribeAggrSoftDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAggrSoftDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeAggrSoftDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggrSoftDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggrSoftDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAggrSoftDeviceListRequest() (request *DescribeAggrSoftDeviceListRequest) {
    request = &DescribeAggrSoftDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeAggrSoftDeviceList")
    
    
    return
}

func NewDescribeAggrSoftDeviceListResponse() (response *DescribeAggrSoftDeviceListResponse) {
    response = &DescribeAggrSoftDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAggrSoftDeviceList
// 聚合软件的已安装终端列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftDeviceList(request *DescribeAggrSoftDeviceListRequest) (response *DescribeAggrSoftDeviceListResponse, err error) {
    return c.DescribeAggrSoftDeviceListWithContext(context.Background(), request)
}

// DescribeAggrSoftDeviceList
// 聚合软件的已安装终端列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeAggrSoftDeviceListWithContext(ctx context.Context, request *DescribeAggrSoftDeviceListRequest) (response *DescribeAggrSoftDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeAggrSoftDeviceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeAggrSoftDeviceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAggrSoftDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAggrSoftDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLPEdgeNodeGroupsRequest() (request *DescribeDLPEdgeNodeGroupsRequest) {
    request = &DescribeDLPEdgeNodeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDLPEdgeNodeGroups")
    
    
    return
}

func NewDescribeDLPEdgeNodeGroupsResponse() (response *DescribeDLPEdgeNodeGroupsResponse) {
    response = &DescribeDLPEdgeNodeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLPEdgeNodeGroups
// 查询边缘节点分组，私有化调用path为：capi/Connectors/DescribeDLPEdgeNodeGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeDLPEdgeNodeGroups(request *DescribeDLPEdgeNodeGroupsRequest) (response *DescribeDLPEdgeNodeGroupsResponse, err error) {
    return c.DescribeDLPEdgeNodeGroupsWithContext(context.Background(), request)
}

// DescribeDLPEdgeNodeGroups
// 查询边缘节点分组，私有化调用path为：capi/Connectors/DescribeDLPEdgeNodeGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeDLPEdgeNodeGroupsWithContext(ctx context.Context, request *DescribeDLPEdgeNodeGroupsRequest) (response *DescribeDLPEdgeNodeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDLPEdgeNodeGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDLPEdgeNodeGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLPEdgeNodeGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLPEdgeNodeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLPEdgeNodesRequest() (request *DescribeDLPEdgeNodesRequest) {
    request = &DescribeDLPEdgeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDLPEdgeNodes")
    
    
    return
}

func NewDescribeDLPEdgeNodesResponse() (response *DescribeDLPEdgeNodesResponse) {
    response = &DescribeDLPEdgeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLPEdgeNodes
// 查询边缘节点列表，私有化调用path为：capi/DlpOpenApi/DescribeDLPEdgeNodes
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeDLPEdgeNodes(request *DescribeDLPEdgeNodesRequest) (response *DescribeDLPEdgeNodesResponse, err error) {
    return c.DescribeDLPEdgeNodesWithContext(context.Background(), request)
}

// DescribeDLPEdgeNodes
// 查询边缘节点列表，私有化调用path为：capi/DlpOpenApi/DescribeDLPEdgeNodes
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DescribeDLPEdgeNodesWithContext(ctx context.Context, request *DescribeDLPEdgeNodesRequest) (response *DescribeDLPEdgeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeDLPEdgeNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDLPEdgeNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLPEdgeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLPEdgeNodesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDLPFileDetectResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLPFileDetectResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLPFileDetectResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLPFileDetectTaskResultRequest() (request *DescribeDLPFileDetectTaskResultRequest) {
    request = &DescribeDLPFileDetectTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDLPFileDetectTaskResult")
    
    
    return
}

func NewDescribeDLPFileDetectTaskResultResponse() (response *DescribeDLPFileDetectTaskResultResponse) {
    response = &DescribeDLPFileDetectTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLPFileDetectTaskResult
// 查询文件鉴定任务结果
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDLPFileDetectTaskResult(request *DescribeDLPFileDetectTaskResultRequest) (response *DescribeDLPFileDetectTaskResultResponse, err error) {
    return c.DescribeDLPFileDetectTaskResultWithContext(context.Background(), request)
}

// DescribeDLPFileDetectTaskResult
// 查询文件鉴定任务结果
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDLPFileDetectTaskResultWithContext(ctx context.Context, request *DescribeDLPFileDetectTaskResultRequest) (response *DescribeDLPFileDetectTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeDLPFileDetectTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDLPFileDetectTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLPFileDetectTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLPFileDetectTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceChildGroupsRequest() (request *DescribeDeviceChildGroupsRequest) {
    request = &DescribeDeviceChildGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDeviceChildGroups")
    
    
    return
}

func NewDescribeDeviceChildGroupsResponse() (response *DescribeDeviceChildGroupsResponse) {
    response = &DescribeDeviceChildGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceChildGroups
// 查询设备组子分组详情，私有化调用path为：capi/Assets/Device/DescribeDeviceChildGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceChildGroups(request *DescribeDeviceChildGroupsRequest) (response *DescribeDeviceChildGroupsResponse, err error) {
    return c.DescribeDeviceChildGroupsWithContext(context.Background(), request)
}

// DescribeDeviceChildGroups
// 查询设备组子分组详情，私有化调用path为：capi/Assets/Device/DescribeDeviceChildGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDeviceChildGroupsWithContext(ctx context.Context, request *DescribeDeviceChildGroupsRequest) (response *DescribeDeviceChildGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceChildGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDeviceChildGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceChildGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceChildGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceDetailListRequest() (request *DescribeDeviceDetailListRequest) {
    request = &DescribeDeviceDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDeviceDetailList")
    
    
    return
}

func NewDescribeDeviceDetailListResponse() (response *DescribeDeviceDetailListResponse) {
    response = &DescribeDeviceDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceDetailList
// 基于软件查看终端详情列表,私有化调用path为：capi/Software/DescribeDeviceDetailList
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeDeviceDetailList(request *DescribeDeviceDetailListRequest) (response *DescribeDeviceDetailListResponse, err error) {
    return c.DescribeDeviceDetailListWithContext(context.Background(), request)
}

// DescribeDeviceDetailList
// 基于软件查看终端详情列表,私有化调用path为：capi/Software/DescribeDeviceDetailList
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) DescribeDeviceDetailListWithContext(ctx context.Context, request *DescribeDeviceDetailListRequest) (response *DescribeDeviceDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDetailListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDeviceDetailList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceDetailListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDeviceHardwareInfoList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDeviceInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceVirtualGroupsRequest() (request *DescribeDeviceVirtualGroupsRequest) {
    request = &DescribeDeviceVirtualGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "DescribeDeviceVirtualGroups")
    
    
    return
}

func NewDescribeDeviceVirtualGroupsResponse() (response *DescribeDeviceVirtualGroupsResponse) {
    response = &DescribeDeviceVirtualGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceVirtualGroups
// 查询终端自定义分组列表，私有化调用path为：/capi/Assets/Device/DescribeDeviceVirtualGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) DescribeDeviceVirtualGroups(request *DescribeDeviceVirtualGroupsRequest) (response *DescribeDeviceVirtualGroupsResponse, err error) {
    return c.DescribeDeviceVirtualGroupsWithContext(context.Background(), request)
}

// DescribeDeviceVirtualGroups
// 查询终端自定义分组列表，私有化调用path为：/capi/Assets/Device/DescribeDeviceVirtualGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"
//  MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"
func (c *Client) DescribeDeviceVirtualGroupsWithContext(ctx context.Context, request *DescribeDeviceVirtualGroupsRequest) (response *DescribeDeviceVirtualGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceVirtualGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDeviceVirtualGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceVirtualGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceVirtualGroupsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeDevices")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeLocalAccounts")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeRootAccountGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeSoftCensusListByDevice")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeSoftwareInformation")
    
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "DescribeVirtualDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVirtualDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVirtualDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewExportSoftwareInformationListRequest() (request *ExportSoftwareInformationListRequest) {
    request = &ExportSoftwareInformationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ioa", APIVersion, "ExportSoftwareInformationList")
    
    
    return
}

func NewExportSoftwareInformationListResponse() (response *ExportSoftwareInformationListResponse) {
    response = &ExportSoftwareInformationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportSoftwareInformationList
// 导出基于指定终端查看软件信息详情列表查询,私有化调用path为：capi/Software/ExportSoftwareInformationList
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ExportSoftwareInformationList(request *ExportSoftwareInformationListRequest) (response *ExportSoftwareInformationListResponse, err error) {
    return c.ExportSoftwareInformationListWithContext(context.Background(), request)
}

// ExportSoftwareInformationList
// 导出基于指定终端查看软件信息详情列表查询,私有化调用path为：capi/Software/ExportSoftwareInformationList
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
func (c *Client) ExportSoftwareInformationListWithContext(ctx context.Context, request *ExportSoftwareInformationListRequest) (response *ExportSoftwareInformationListResponse, err error) {
    if request == nil {
        request = NewExportSoftwareInformationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "ExportSoftwareInformationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportSoftwareInformationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportSoftwareInformationListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "ioa", APIVersion, "ModifyVirtualDeviceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVirtualDeviceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVirtualDeviceGroupsResponse()
    err = c.Send(request, response)
    return
}
