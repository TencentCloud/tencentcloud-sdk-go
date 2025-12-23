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

package v20220217

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-02-17"

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


func NewCreateCloudRunEnvRequest() (request *CreateCloudRunEnvRequest) {
    request = &CreateCloudRunEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "CreateCloudRunEnv")
    
    
    return
}

func NewCreateCloudRunEnvResponse() (response *CreateCloudRunEnvResponse) {
    response = &CreateCloudRunEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudRunEnv
// 创建云托管环境，并开通资源。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateCloudRunEnv(request *CreateCloudRunEnvRequest) (response *CreateCloudRunEnvResponse, err error) {
    return c.CreateCloudRunEnvWithContext(context.Background(), request)
}

// CreateCloudRunEnv
// 创建云托管环境，并开通资源。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateCloudRunEnvWithContext(ctx context.Context, request *CreateCloudRunEnvRequest) (response *CreateCloudRunEnvResponse, err error) {
    if request == nil {
        request = NewCreateCloudRunEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "CreateCloudRunEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRunEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRunEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudRunServerRequest() (request *CreateCloudRunServerRequest) {
    request = &CreateCloudRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "CreateCloudRunServer")
    
    
    return
}

func NewCreateCloudRunServerResponse() (response *CreateCloudRunServerResponse) {
    response = &CreateCloudRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudRunServer
// 创建云托管服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) CreateCloudRunServer(request *CreateCloudRunServerRequest) (response *CreateCloudRunServerResponse, err error) {
    return c.CreateCloudRunServerWithContext(context.Background(), request)
}

// CreateCloudRunServer
// 创建云托管服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) CreateCloudRunServerWithContext(ctx context.Context, request *CreateCloudRunServerRequest) (response *CreateCloudRunServerResponse, err error) {
    if request == nil {
        request = NewCreateCloudRunServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "CreateCloudRunServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRunServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRunServerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudRunServerRequest() (request *DeleteCloudRunServerRequest) {
    request = &DeleteCloudRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DeleteCloudRunServer")
    
    
    return
}

func NewDeleteCloudRunServerResponse() (response *DeleteCloudRunServerResponse) {
    response = &DeleteCloudRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudRunServer
// 删除云托管服务：包括服务下的版本，镜像，流水线
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteCloudRunServer(request *DeleteCloudRunServerRequest) (response *DeleteCloudRunServerResponse, err error) {
    return c.DeleteCloudRunServerWithContext(context.Background(), request)
}

// DeleteCloudRunServer
// 删除云托管服务：包括服务下的版本，镜像，流水线
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteCloudRunServerWithContext(ctx context.Context, request *DeleteCloudRunServerRequest) (response *DeleteCloudRunServerResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRunServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DeleteCloudRunServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRunServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRunServerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudRunVersionsRequest() (request *DeleteCloudRunVersionsRequest) {
    request = &DeleteCloudRunVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DeleteCloudRunVersions")
    
    
    return
}

func NewDeleteCloudRunVersionsResponse() (response *DeleteCloudRunVersionsResponse) {
    response = &DeleteCloudRunVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudRunVersions
// 批量删除版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCloudRunVersions(request *DeleteCloudRunVersionsRequest) (response *DeleteCloudRunVersionsResponse, err error) {
    return c.DeleteCloudRunVersionsWithContext(context.Background(), request)
}

// DeleteCloudRunVersions
// 批量删除版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCloudRunVersionsWithContext(ctx context.Context, request *DeleteCloudRunVersionsRequest) (response *DeleteCloudRunVersionsResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRunVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DeleteCloudRunVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRunVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRunVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunDeployRecordRequest() (request *DescribeCloudRunDeployRecordRequest) {
    request = &DescribeCloudRunDeployRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunDeployRecord")
    
    
    return
}

func NewDescribeCloudRunDeployRecordResponse() (response *DescribeCloudRunDeployRecordResponse) {
    response = &DescribeCloudRunDeployRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunDeployRecord
// 查询云托管部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudRunDeployRecord(request *DescribeCloudRunDeployRecordRequest) (response *DescribeCloudRunDeployRecordResponse, err error) {
    return c.DescribeCloudRunDeployRecordWithContext(context.Background(), request)
}

// DescribeCloudRunDeployRecord
// 查询云托管部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudRunDeployRecordWithContext(ctx context.Context, request *DescribeCloudRunDeployRecordRequest) (response *DescribeCloudRunDeployRecordResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunDeployRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunDeployRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunDeployRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunDeployRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunEnvsRequest() (request *DescribeCloudRunEnvsRequest) {
    request = &DescribeCloudRunEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunEnvs")
    
    
    return
}

func NewDescribeCloudRunEnvsResponse() (response *DescribeCloudRunEnvsResponse) {
    response = &DescribeCloudRunEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCloudRunEnvs(request *DescribeCloudRunEnvsRequest) (response *DescribeCloudRunEnvsResponse, err error) {
    return c.DescribeCloudRunEnvsWithContext(context.Background(), request)
}

// DescribeCloudRunEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCloudRunEnvsWithContext(ctx context.Context, request *DescribeCloudRunEnvsRequest) (response *DescribeCloudRunEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunPodListRequest() (request *DescribeCloudRunPodListRequest) {
    request = &DescribeCloudRunPodListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunPodList")
    
    
    return
}

func NewDescribeCloudRunPodListResponse() (response *DescribeCloudRunPodListResponse) {
    response = &DescribeCloudRunPodListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunPodList
// 查询云托管Pod实例列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudRunPodList(request *DescribeCloudRunPodListRequest) (response *DescribeCloudRunPodListResponse, err error) {
    return c.DescribeCloudRunPodListWithContext(context.Background(), request)
}

// DescribeCloudRunPodList
// 查询云托管Pod实例列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudRunPodListWithContext(ctx context.Context, request *DescribeCloudRunPodListRequest) (response *DescribeCloudRunPodListResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunPodListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunPodList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunPodList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunPodListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunProcessLogRequest() (request *DescribeCloudRunProcessLogRequest) {
    request = &DescribeCloudRunProcessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunProcessLog")
    
    
    return
}

func NewDescribeCloudRunProcessLogResponse() (response *DescribeCloudRunProcessLogResponse) {
    response = &DescribeCloudRunProcessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunProcessLog
// 查询运行日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudRunProcessLog(request *DescribeCloudRunProcessLogRequest) (response *DescribeCloudRunProcessLogResponse, err error) {
    return c.DescribeCloudRunProcessLogWithContext(context.Background(), request)
}

// DescribeCloudRunProcessLog
// 查询运行日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCloudRunProcessLogWithContext(ctx context.Context, request *DescribeCloudRunProcessLogRequest) (response *DescribeCloudRunProcessLogResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunProcessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunProcessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunProcessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunProcessLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunServerDetailRequest() (request *DescribeCloudRunServerDetailRequest) {
    request = &DescribeCloudRunServerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunServerDetail")
    
    
    return
}

func NewDescribeCloudRunServerDetailResponse() (response *DescribeCloudRunServerDetailResponse) {
    response = &DescribeCloudRunServerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunServerDetail
// 查询云托管服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudRunServerDetail(request *DescribeCloudRunServerDetailRequest) (response *DescribeCloudRunServerDetailResponse, err error) {
    return c.DescribeCloudRunServerDetailWithContext(context.Background(), request)
}

// DescribeCloudRunServerDetail
// 查询云托管服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVERNOTFOUND = "ResourceNotFound.ServerNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_RESOURCEBANNED = "ResourceUnavailable.ResourceBanned"
//  RESOURCEUNAVAILABLE_RESOURCEFROZEN = "ResourceUnavailable.ResourceFrozen"
//  RESOURCEUNAVAILABLE_RESOURCEISOLATED = "ResourceUnavailable.ResourceIsolated"
func (c *Client) DescribeCloudRunServerDetailWithContext(ctx context.Context, request *DescribeCloudRunServerDetailRequest) (response *DescribeCloudRunServerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunServerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunServerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunServerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunServerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunServersRequest() (request *DescribeCloudRunServersRequest) {
    request = &DescribeCloudRunServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunServers")
    
    
    return
}

func NewDescribeCloudRunServersResponse() (response *DescribeCloudRunServersResponse) {
    response = &DescribeCloudRunServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRunServers
// 查询云托管服务列表接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServers(request *DescribeCloudRunServersRequest) (response *DescribeCloudRunServersResponse, err error) {
    return c.DescribeCloudRunServersWithContext(context.Background(), request)
}

// DescribeCloudRunServers
// 查询云托管服务列表接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServersWithContext(ctx context.Context, request *DescribeCloudRunServersRequest) (response *DescribeCloudRunServersResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunServersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeCloudRunServers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvBaseInfoRequest() (request *DescribeEnvBaseInfoRequest) {
    request = &DescribeEnvBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeEnvBaseInfo")
    
    
    return
}

func NewDescribeEnvBaseInfoResponse() (response *DescribeEnvBaseInfoResponse) {
    response = &DescribeEnvBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvBaseInfo
// 查询环境基础信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvBaseInfo(request *DescribeEnvBaseInfoRequest) (response *DescribeEnvBaseInfoResponse, err error) {
    return c.DescribeEnvBaseInfoWithContext(context.Background(), request)
}

// DescribeEnvBaseInfo
// 查询环境基础信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvBaseInfoWithContext(ctx context.Context, request *DescribeEnvBaseInfoRequest) (response *DescribeEnvBaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEnvBaseInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeEnvBaseInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvBaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseOrderRequest() (request *DescribeReleaseOrderRequest) {
    request = &DescribeReleaseOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeReleaseOrder")
    
    
    return
}

func NewDescribeReleaseOrderResponse() (response *DescribeReleaseOrderResponse) {
    response = &DescribeReleaseOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseOrder
// 查询发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReleaseOrder(request *DescribeReleaseOrderRequest) (response *DescribeReleaseOrderResponse, err error) {
    return c.DescribeReleaseOrderWithContext(context.Background(), request)
}

// DescribeReleaseOrder
// 查询发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReleaseOrderWithContext(ctx context.Context, request *DescribeReleaseOrderRequest) (response *DescribeReleaseOrderResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeReleaseOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerManageTaskRequest() (request *DescribeServerManageTaskRequest) {
    request = &DescribeServerManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeServerManageTask")
    
    
    return
}

func NewDescribeServerManageTaskResponse() (response *DescribeServerManageTaskResponse) {
    response = &DescribeServerManageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerManageTask
// 查询服务管理任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerManageTask(request *DescribeServerManageTaskRequest) (response *DescribeServerManageTaskResponse, err error) {
    return c.DescribeServerManageTaskWithContext(context.Background(), request)
}

// DescribeServerManageTask
// 查询服务管理任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerManageTaskWithContext(ctx context.Context, request *DescribeServerManageTaskRequest) (response *DescribeServerManageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeServerManageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeServerManageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerManageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerManageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVersionDetailRequest() (request *DescribeVersionDetailRequest) {
    request = &DescribeVersionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeVersionDetail")
    
    
    return
}

func NewDescribeVersionDetailResponse() (response *DescribeVersionDetailResponse) {
    response = &DescribeVersionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVersionDetail
// 查询版本详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVersionDetail(request *DescribeVersionDetailRequest) (response *DescribeVersionDetailResponse, err error) {
    return c.DescribeVersionDetailWithContext(context.Background(), request)
}

// DescribeVersionDetail
// 查询版本详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVersionDetailWithContext(ctx context.Context, request *DescribeVersionDetailRequest) (response *DescribeVersionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVersionDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "DescribeVersionDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVersionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVersionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewOperateServerManageRequest() (request *OperateServerManageRequest) {
    request = &OperateServerManageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "OperateServerManage")
    
    
    return
}

func NewOperateServerManageResponse() (response *OperateServerManageResponse) {
    response = &OperateServerManageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OperateServerManage
// 操作发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) OperateServerManage(request *OperateServerManageRequest) (response *OperateServerManageResponse, err error) {
    return c.OperateServerManageWithContext(context.Background(), request)
}

// OperateServerManage
// 操作发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) OperateServerManageWithContext(ctx context.Context, request *OperateServerManageRequest) (response *OperateServerManageResponse, err error) {
    if request == nil {
        request = NewOperateServerManageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "OperateServerManage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateServerManage require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateServerManageResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseGrayRequest() (request *ReleaseGrayRequest) {
    request = &ReleaseGrayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "ReleaseGray")
    
    
    return
}

func NewReleaseGrayResponse() (response *ReleaseGrayResponse) {
    response = &ReleaseGrayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseGray
// 灰度发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ReleaseGray(request *ReleaseGrayRequest) (response *ReleaseGrayResponse, err error) {
    return c.ReleaseGrayWithContext(context.Background(), request)
}

// ReleaseGray
// 灰度发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ReleaseGrayWithContext(ctx context.Context, request *ReleaseGrayRequest) (response *ReleaseGrayResponse, err error) {
    if request == nil {
        request = NewReleaseGrayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "ReleaseGray")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseGray require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseGrayResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClsLog
// 查询日志信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    return c.SearchClsLogWithContext(context.Background(), request)
}

// SearchClsLog
// 查询日志信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "SearchClsLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitServerRollbackRequest() (request *SubmitServerRollbackRequest) {
    request = &SubmitServerRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "SubmitServerRollback")
    
    
    return
}

func NewSubmitServerRollbackResponse() (response *SubmitServerRollbackResponse) {
    response = &SubmitServerRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitServerRollback
// 回滚版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SubmitServerRollback(request *SubmitServerRollbackRequest) (response *SubmitServerRollbackResponse, err error) {
    return c.SubmitServerRollbackWithContext(context.Background(), request)
}

// SubmitServerRollback
// 回滚版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SubmitServerRollbackWithContext(ctx context.Context, request *SubmitServerRollbackRequest) (response *SubmitServerRollbackResponse, err error) {
    if request == nil {
        request = NewSubmitServerRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "SubmitServerRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitServerRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitServerRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudRunServerRequest() (request *UpdateCloudRunServerRequest) {
    request = &UpdateCloudRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "UpdateCloudRunServer")
    
    
    return
}

func NewUpdateCloudRunServerResponse() (response *UpdateCloudRunServerResponse) {
    response = &UpdateCloudRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCloudRunServer
// 更新云托管服务
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateCloudRunServer(request *UpdateCloudRunServerRequest) (response *UpdateCloudRunServerResponse, err error) {
    return c.UpdateCloudRunServerWithContext(context.Background(), request)
}

// UpdateCloudRunServer
// 更新云托管服务
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateCloudRunServerWithContext(ctx context.Context, request *UpdateCloudRunServerRequest) (response *UpdateCloudRunServerResponse, err error) {
    if request == nil {
        request = NewUpdateCloudRunServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcbr", APIVersion, "UpdateCloudRunServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudRunServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudRunServerResponse()
    err = c.Send(request, response)
    return
}
