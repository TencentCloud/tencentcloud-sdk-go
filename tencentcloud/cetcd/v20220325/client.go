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

package v20220325

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-03-25"

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


func NewCreateEtcdInstanceRequest() (request *CreateEtcdInstanceRequest) {
    request = &CreateEtcdInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "CreateEtcdInstance")
    
    
    return
}

func NewCreateEtcdInstanceResponse() (response *CreateEtcdInstanceResponse) {
    response = &CreateEtcdInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEtcdInstance
// 创建etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED_QUOTAETCDLIMIT = "LimitExceeded.QuotaEtcdLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEtcdInstance(request *CreateEtcdInstanceRequest) (response *CreateEtcdInstanceResponse, err error) {
    return c.CreateEtcdInstanceWithContext(context.Background(), request)
}

// CreateEtcdInstance
// 创建etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED_QUOTAETCDLIMIT = "LimitExceeded.QuotaEtcdLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateEtcdInstanceWithContext(ctx context.Context, request *CreateEtcdInstanceRequest) (response *CreateEtcdInstanceResponse, err error) {
    if request == nil {
        request = NewCreateEtcdInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "CreateEtcdInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEtcdInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEtcdInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEtcdSnapshotRequest() (request *CreateEtcdSnapshotRequest) {
    request = &CreateEtcdSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "CreateEtcdSnapshot")
    
    
    return
}

func NewCreateEtcdSnapshotResponse() (response *CreateEtcdSnapshotResponse) {
    response = &CreateEtcdSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEtcdSnapshot
// 创建etcd快照
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEEXISTALREADY = "FailedOperation.ResourceExistAlready"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateEtcdSnapshot(request *CreateEtcdSnapshotRequest) (response *CreateEtcdSnapshotResponse, err error) {
    return c.CreateEtcdSnapshotWithContext(context.Background(), request)
}

// CreateEtcdSnapshot
// 创建etcd快照
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEEXISTALREADY = "FailedOperation.ResourceExistAlready"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateEtcdSnapshotWithContext(ctx context.Context, request *CreateEtcdSnapshotRequest) (response *CreateEtcdSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateEtcdSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "CreateEtcdSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEtcdSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEtcdSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEtcdSnapshotPolicyRequest() (request *CreateEtcdSnapshotPolicyRequest) {
    request = &CreateEtcdSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "CreateEtcdSnapshotPolicy")
    
    
    return
}

func NewCreateEtcdSnapshotPolicyResponse() (response *CreateEtcdSnapshotPolicyResponse) {
    response = &CreateEtcdSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEtcdSnapshotPolicy
// 创建etcd快照策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEEXISTALREADY = "FailedOperation.ResourceExistAlready"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEtcdSnapshotPolicy(request *CreateEtcdSnapshotPolicyRequest) (response *CreateEtcdSnapshotPolicyResponse, err error) {
    return c.CreateEtcdSnapshotPolicyWithContext(context.Background(), request)
}

// CreateEtcdSnapshotPolicy
// 创建etcd快照策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCEEXISTALREADY = "FailedOperation.ResourceExistAlready"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateEtcdSnapshotPolicyWithContext(ctx context.Context, request *CreateEtcdSnapshotPolicyRequest) (response *CreateEtcdSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewCreateEtcdSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "CreateEtcdSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEtcdSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEtcdSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdAvailableVersionsRequest() (request *DescribeEtcdAvailableVersionsRequest) {
    request = &DescribeEtcdAvailableVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdAvailableVersions")
    
    
    return
}

func NewDescribeEtcdAvailableVersionsResponse() (response *DescribeEtcdAvailableVersionsResponse) {
    response = &DescribeEtcdAvailableVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdAvailableVersions
// 查看etcd可用版本
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdAvailableVersions(request *DescribeEtcdAvailableVersionsRequest) (response *DescribeEtcdAvailableVersionsResponse, err error) {
    return c.DescribeEtcdAvailableVersionsWithContext(context.Background(), request)
}

// DescribeEtcdAvailableVersions
// 查看etcd可用版本
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdAvailableVersionsWithContext(ctx context.Context, request *DescribeEtcdAvailableVersionsRequest) (response *DescribeEtcdAvailableVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdAvailableVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdAvailableVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdAvailableVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdAvailableVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdCreatingProgressRequest() (request *DescribeEtcdCreatingProgressRequest) {
    request = &DescribeEtcdCreatingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdCreatingProgress")
    
    
    return
}

func NewDescribeEtcdCreatingProgressResponse() (response *DescribeEtcdCreatingProgressResponse) {
    response = &DescribeEtcdCreatingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdCreatingProgress
// 查看etcd集群创建进度
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdCreatingProgress(request *DescribeEtcdCreatingProgressRequest) (response *DescribeEtcdCreatingProgressResponse, err error) {
    return c.DescribeEtcdCreatingProgressWithContext(context.Background(), request)
}

// DescribeEtcdCreatingProgress
// 查看etcd集群创建进度
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdCreatingProgressWithContext(ctx context.Context, request *DescribeEtcdCreatingProgressRequest) (response *DescribeEtcdCreatingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdCreatingProgressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdCreatingProgress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdCreatingProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdCreatingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdCredentialsRequest() (request *DescribeEtcdCredentialsRequest) {
    request = &DescribeEtcdCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdCredentials")
    
    
    return
}

func NewDescribeEtcdCredentialsResponse() (response *DescribeEtcdCredentialsResponse) {
    response = &DescribeEtcdCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdCredentials
// 查询etcd访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdCredentials(request *DescribeEtcdCredentialsRequest) (response *DescribeEtcdCredentialsResponse, err error) {
    return c.DescribeEtcdCredentialsWithContext(context.Background(), request)
}

// DescribeEtcdCredentials
// 查询etcd访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdCredentialsWithContext(ctx context.Context, request *DescribeEtcdCredentialsRequest) (response *DescribeEtcdCredentialsResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdCredentialsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdCredentials")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdInstancesRequest() (request *DescribeEtcdInstancesRequest) {
    request = &DescribeEtcdInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdInstances")
    
    
    return
}

func NewDescribeEtcdInstancesResponse() (response *DescribeEtcdInstancesResponse) {
    response = &DescribeEtcdInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdInstances
// 查询etcd实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeEtcdInstances(request *DescribeEtcdInstancesRequest) (response *DescribeEtcdInstancesResponse, err error) {
    return c.DescribeEtcdInstancesWithContext(context.Background(), request)
}

// DescribeEtcdInstances
// 查询etcd实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeEtcdInstancesWithContext(ctx context.Context, request *DescribeEtcdInstancesRequest) (response *DescribeEtcdInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdQuotaRequest() (request *DescribeEtcdQuotaRequest) {
    request = &DescribeEtcdQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdQuota")
    
    
    return
}

func NewDescribeEtcdQuotaResponse() (response *DescribeEtcdQuotaResponse) {
    response = &DescribeEtcdQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdQuota
// 查看etcd集群配额
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdQuota(request *DescribeEtcdQuotaRequest) (response *DescribeEtcdQuotaResponse, err error) {
    return c.DescribeEtcdQuotaWithContext(context.Background(), request)
}

// DescribeEtcdQuota
// 查看etcd集群配额
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdQuotaWithContext(ctx context.Context, request *DescribeEtcdQuotaRequest) (response *DescribeEtcdQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdRegionsRequest() (request *DescribeEtcdRegionsRequest) {
    request = &DescribeEtcdRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdRegions")
    
    
    return
}

func NewDescribeEtcdRegionsResponse() (response *DescribeEtcdRegionsResponse) {
    response = &DescribeEtcdRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdRegions
// 查看etcd支持地域
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdRegions(request *DescribeEtcdRegionsRequest) (response *DescribeEtcdRegionsResponse, err error) {
    return c.DescribeEtcdRegionsWithContext(context.Background(), request)
}

// DescribeEtcdRegions
// 查看etcd支持地域
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) DescribeEtcdRegionsWithContext(ctx context.Context, request *DescribeEtcdRegionsRequest) (response *DescribeEtcdRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdSnapshotPoliciesRequest() (request *DescribeEtcdSnapshotPoliciesRequest) {
    request = &DescribeEtcdSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdSnapshotPolicies")
    
    
    return
}

func NewDescribeEtcdSnapshotPoliciesResponse() (response *DescribeEtcdSnapshotPoliciesResponse) {
    response = &DescribeEtcdSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdSnapshotPolicies
// 查看etcd快照策略
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdSnapshotPolicies(request *DescribeEtcdSnapshotPoliciesRequest) (response *DescribeEtcdSnapshotPoliciesResponse, err error) {
    return c.DescribeEtcdSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeEtcdSnapshotPolicies
// 查看etcd快照策略
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdSnapshotPoliciesWithContext(ctx context.Context, request *DescribeEtcdSnapshotPoliciesRequest) (response *DescribeEtcdSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdSnapshotsRequest() (request *DescribeEtcdSnapshotsRequest) {
    request = &DescribeEtcdSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdSnapshots")
    
    
    return
}

func NewDescribeEtcdSnapshotsResponse() (response *DescribeEtcdSnapshotsResponse) {
    response = &DescribeEtcdSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdSnapshots
// 查看etcd快照列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdSnapshots(request *DescribeEtcdSnapshotsRequest) (response *DescribeEtcdSnapshotsResponse, err error) {
    return c.DescribeEtcdSnapshotsWithContext(context.Background(), request)
}

// DescribeEtcdSnapshots
// 查看etcd快照列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEtcdSnapshotsWithContext(ctx context.Context, request *DescribeEtcdSnapshotsRequest) (response *DescribeEtcdSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdSnapshotsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdSnapshots")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEtcdTasksRequest() (request *DescribeEtcdTasksRequest) {
    request = &DescribeEtcdTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeEtcdTasks")
    
    
    return
}

func NewDescribeEtcdTasksResponse() (response *DescribeEtcdTasksResponse) {
    response = &DescribeEtcdTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEtcdTasks
// 查看etcd相关tasks
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEtcdTasks(request *DescribeEtcdTasksRequest) (response *DescribeEtcdTasksResponse, err error) {
    return c.DescribeEtcdTasksWithContext(context.Background(), request)
}

// DescribeEtcdTasks
// 查看etcd相关tasks
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEtcdTasksWithContext(ctx context.Context, request *DescribeEtcdTasksRequest) (response *DescribeEtcdTasksResponse, err error) {
    if request == nil {
        request = NewDescribeEtcdTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeEtcdTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEtcdTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEtcdTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRPCMethodListRequest() (request *DescribeRPCMethodListRequest) {
    request = &DescribeRPCMethodListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DescribeRPCMethodList")
    
    
    return
}

func NewDescribeRPCMethodListResponse() (response *DescribeRPCMethodListResponse) {
    response = &DescribeRPCMethodListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRPCMethodList
// 获取etcd集群的gRPC方法列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRPCMethodList(request *DescribeRPCMethodListRequest) (response *DescribeRPCMethodListResponse, err error) {
    return c.DescribeRPCMethodListWithContext(context.Background(), request)
}

// DescribeRPCMethodList
// 获取etcd集群的gRPC方法列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRPCMethodListWithContext(ctx context.Context, request *DescribeRPCMethodListRequest) (response *DescribeRPCMethodListResponse, err error) {
    if request == nil {
        request = NewDescribeRPCMethodListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DescribeRPCMethodList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRPCMethodList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRPCMethodListResponse()
    err = c.Send(request, response)
    return
}

func NewDisableEtcdInstanceDeletionProtectionRequest() (request *DisableEtcdInstanceDeletionProtectionRequest) {
    request = &DisableEtcdInstanceDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "DisableEtcdInstanceDeletionProtection")
    
    
    return
}

func NewDisableEtcdInstanceDeletionProtectionResponse() (response *DisableEtcdInstanceDeletionProtectionResponse) {
    response = &DisableEtcdInstanceDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableEtcdInstanceDeletionProtection
// 关闭etcd实例删除保护
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableEtcdInstanceDeletionProtection(request *DisableEtcdInstanceDeletionProtectionRequest) (response *DisableEtcdInstanceDeletionProtectionResponse, err error) {
    return c.DisableEtcdInstanceDeletionProtectionWithContext(context.Background(), request)
}

// DisableEtcdInstanceDeletionProtection
// 关闭etcd实例删除保护
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableEtcdInstanceDeletionProtectionWithContext(ctx context.Context, request *DisableEtcdInstanceDeletionProtectionRequest) (response *DisableEtcdInstanceDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewDisableEtcdInstanceDeletionProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "DisableEtcdInstanceDeletionProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableEtcdInstanceDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableEtcdInstanceDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewEnableEtcdInstanceDeletionProtectionRequest() (request *EnableEtcdInstanceDeletionProtectionRequest) {
    request = &EnableEtcdInstanceDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "EnableEtcdInstanceDeletionProtection")
    
    
    return
}

func NewEnableEtcdInstanceDeletionProtectionResponse() (response *EnableEtcdInstanceDeletionProtectionResponse) {
    response = &EnableEtcdInstanceDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableEtcdInstanceDeletionProtection
// 启用etcd实例删除保护
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableEtcdInstanceDeletionProtection(request *EnableEtcdInstanceDeletionProtectionRequest) (response *EnableEtcdInstanceDeletionProtectionResponse, err error) {
    return c.EnableEtcdInstanceDeletionProtectionWithContext(context.Background(), request)
}

// EnableEtcdInstanceDeletionProtection
// 启用etcd实例删除保护
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableEtcdInstanceDeletionProtectionWithContext(ctx context.Context, request *EnableEtcdInstanceDeletionProtectionRequest) (response *EnableEtcdInstanceDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewEnableEtcdInstanceDeletionProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "EnableEtcdInstanceDeletionProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableEtcdInstanceDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableEtcdInstanceDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEtcdAttributeRequest() (request *ModifyEtcdAttributeRequest) {
    request = &ModifyEtcdAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "ModifyEtcdAttribute")
    
    
    return
}

func NewModifyEtcdAttributeResponse() (response *ModifyEtcdAttributeResponse) {
    response = &ModifyEtcdAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEtcdAttribute
// 修改etcd实例属性
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEtcdAttribute(request *ModifyEtcdAttributeRequest) (response *ModifyEtcdAttributeResponse, err error) {
    return c.ModifyEtcdAttributeWithContext(context.Background(), request)
}

// ModifyEtcdAttribute
// 修改etcd实例属性
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyEtcdAttributeWithContext(ctx context.Context, request *ModifyEtcdAttributeRequest) (response *ModifyEtcdAttributeResponse, err error) {
    if request == nil {
        request = NewModifyEtcdAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "ModifyEtcdAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEtcdAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEtcdAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEtcdConfigurationRequest() (request *ModifyEtcdConfigurationRequest) {
    request = &ModifyEtcdConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "ModifyEtcdConfiguration")
    
    
    return
}

func NewModifyEtcdConfigurationResponse() (response *ModifyEtcdConfigurationResponse) {
    response = &ModifyEtcdConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEtcdConfiguration
// 修改etcd实例配置
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_PROMETHEUSSTATUSERROR = "ResourceUnavailable.PrometheusStatusError"
func (c *Client) ModifyEtcdConfiguration(request *ModifyEtcdConfigurationRequest) (response *ModifyEtcdConfigurationResponse, err error) {
    return c.ModifyEtcdConfigurationWithContext(context.Background(), request)
}

// ModifyEtcdConfiguration
// 修改etcd实例配置
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_PROMETHEUSSTATUSERROR = "ResourceUnavailable.PrometheusStatusError"
func (c *Client) ModifyEtcdConfigurationWithContext(ctx context.Context, request *ModifyEtcdConfigurationRequest) (response *ModifyEtcdConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyEtcdConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "ModifyEtcdConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEtcdConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEtcdConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEtcdSnapshotPolicyRequest() (request *ModifyEtcdSnapshotPolicyRequest) {
    request = &ModifyEtcdSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "ModifyEtcdSnapshotPolicy")
    
    
    return
}

func NewModifyEtcdSnapshotPolicyResponse() (response *ModifyEtcdSnapshotPolicyResponse) {
    response = &ModifyEtcdSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEtcdSnapshotPolicy
// 修改etcd快照策略
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) ModifyEtcdSnapshotPolicy(request *ModifyEtcdSnapshotPolicyRequest) (response *ModifyEtcdSnapshotPolicyResponse, err error) {
    return c.ModifyEtcdSnapshotPolicyWithContext(context.Background(), request)
}

// ModifyEtcdSnapshotPolicy
// 修改etcd快照策略
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) ModifyEtcdSnapshotPolicyWithContext(ctx context.Context, request *ModifyEtcdSnapshotPolicyRequest) (response *ModifyEtcdSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewModifyEtcdSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "ModifyEtcdSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEtcdSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEtcdSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResizeEtcdInstanceRequest() (request *ResizeEtcdInstanceRequest) {
    request = &ResizeEtcdInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "ResizeEtcdInstance")
    
    
    return
}

func NewResizeEtcdInstanceResponse() (response *ResizeEtcdInstanceResponse) {
    response = &ResizeEtcdInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeEtcdInstance
// 扩容etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) ResizeEtcdInstance(request *ResizeEtcdInstanceRequest) (response *ResizeEtcdInstanceResponse, err error) {
    return c.ResizeEtcdInstanceWithContext(context.Background(), request)
}

// ResizeEtcdInstance
// 扩容etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
func (c *Client) ResizeEtcdInstanceWithContext(ctx context.Context, request *ResizeEtcdInstanceRequest) (response *ResizeEtcdInstanceResponse, err error) {
    if request == nil {
        request = NewResizeEtcdInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "ResizeEtcdInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeEtcdInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeEtcdInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeEtcdInstanceRequest() (request *UpgradeEtcdInstanceRequest) {
    request = &UpgradeEtcdInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cetcd", APIVersion, "UpgradeEtcdInstance")
    
    
    return
}

func NewUpgradeEtcdInstanceResponse() (response *UpgradeEtcdInstanceResponse) {
    response = &UpgradeEtcdInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeEtcdInstance
// 升级etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpgradeEtcdInstance(request *UpgradeEtcdInstanceRequest) (response *UpgradeEtcdInstanceResponse, err error) {
    return c.UpgradeEtcdInstanceWithContext(context.Background(), request)
}

// UpgradeEtcdInstance
// 升级etcd实例
//
// 可能返回的错误码:
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpgradeEtcdInstanceWithContext(ctx context.Context, request *UpgradeEtcdInstanceRequest) (response *UpgradeEtcdInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeEtcdInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cetcd", APIVersion, "UpgradeEtcdInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeEtcdInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeEtcdInstanceResponse()
    err = c.Send(request, response)
    return
}
