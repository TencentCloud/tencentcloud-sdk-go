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

package v20220501

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-05-01"

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


func NewCreateHealthCheckPolicyRequest() (request *CreateHealthCheckPolicyRequest) {
    request = &CreateHealthCheckPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateHealthCheckPolicy")
    
    
    return
}

func NewCreateHealthCheckPolicyResponse() (response *CreateHealthCheckPolicyResponse) {
    response = &CreateHealthCheckPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHealthCheckPolicy
// 创建健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreateHealthCheckPolicy(request *CreateHealthCheckPolicyRequest) (response *CreateHealthCheckPolicyResponse, err error) {
    return c.CreateHealthCheckPolicyWithContext(context.Background(), request)
}

// CreateHealthCheckPolicy
// 创建健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreateHealthCheckPolicyWithContext(ctx context.Context, request *CreateHealthCheckPolicyRequest) (response *CreateHealthCheckPolicyResponse, err error) {
    if request == nil {
        request = NewCreateHealthCheckPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "CreateHealthCheckPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHealthCheckPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHealthCheckPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNodePoolRequest() (request *CreateNodePoolRequest) {
    request = &CreateNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "CreateNodePool")
    
    
    return
}

func NewCreateNodePoolResponse() (response *CreateNodePoolResponse) {
    response = &CreateNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNodePool
// 创建 TKE 节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateNodePool(request *CreateNodePoolRequest) (response *CreateNodePoolResponse, err error) {
    return c.CreateNodePoolWithContext(context.Background(), request)
}

// CreateNodePool
// 创建 TKE 节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateNodePoolWithContext(ctx context.Context, request *CreateNodePoolRequest) (response *CreateNodePoolResponse, err error) {
    if request == nil {
        request = NewCreateNodePoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "CreateNodePool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterMachinesRequest() (request *DeleteClusterMachinesRequest) {
    request = &DeleteClusterMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteClusterMachines")
    
    
    return
}

func NewDeleteClusterMachinesResponse() (response *DeleteClusterMachinesResponse) {
    response = &DeleteClusterMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterMachines
// 删除原生节点池节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterMachines(request *DeleteClusterMachinesRequest) (response *DeleteClusterMachinesResponse, err error) {
    return c.DeleteClusterMachinesWithContext(context.Background(), request)
}

// DeleteClusterMachines
// 删除原生节点池节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteClusterMachinesWithContext(ctx context.Context, request *DeleteClusterMachinesRequest) (response *DeleteClusterMachinesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DeleteClusterMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHealthCheckPolicyRequest() (request *DeleteHealthCheckPolicyRequest) {
    request = &DeleteHealthCheckPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteHealthCheckPolicy")
    
    
    return
}

func NewDeleteHealthCheckPolicyResponse() (response *DeleteHealthCheckPolicyResponse) {
    response = &DeleteHealthCheckPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHealthCheckPolicy
// 删除健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DeleteHealthCheckPolicy(request *DeleteHealthCheckPolicyRequest) (response *DeleteHealthCheckPolicyResponse, err error) {
    return c.DeleteHealthCheckPolicyWithContext(context.Background(), request)
}

// DeleteHealthCheckPolicy
// 删除健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DeleteHealthCheckPolicyWithContext(ctx context.Context, request *DeleteHealthCheckPolicyRequest) (response *DeleteHealthCheckPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteHealthCheckPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DeleteHealthCheckPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHealthCheckPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHealthCheckPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNodePoolRequest() (request *DeleteNodePoolRequest) {
    request = &DeleteNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DeleteNodePool")
    
    
    return
}

func NewDeleteNodePoolResponse() (response *DeleteNodePoolResponse) {
    response = &DeleteNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNodePool
// 删除 TKE 节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNodePool(request *DeleteNodePoolRequest) (response *DeleteNodePoolResponse, err error) {
    return c.DeleteNodePoolWithContext(context.Background(), request)
}

// DeleteNodePool
// 删除 TKE 节点池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNodePoolWithContext(ctx context.Context, request *DeleteNodePoolRequest) (response *DeleteNodePoolResponse, err error) {
    if request == nil {
        request = NewDeleteNodePoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DeleteNodePool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstances
// 查询集群下节点实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    return c.DescribeClusterInstancesWithContext(context.Background(), request)
}

// DescribeClusterInstances
// 查询集群下节点实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeClusterInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterMachinesRequest() (request *DescribeClusterMachinesRequest) {
    request = &DescribeClusterMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusterMachines")
    
    
    return
}

func NewDescribeClusterMachinesResponse() (response *DescribeClusterMachinesResponse) {
    response = &DescribeClusterMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterMachines
// 查询托原生点列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterMachines(request *DescribeClusterMachinesRequest) (response *DescribeClusterMachinesResponse, err error) {
    return c.DescribeClusterMachinesWithContext(context.Background(), request)
}

// DescribeClusterMachines
// 查询托原生点列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
func (c *Client) DescribeClusterMachinesWithContext(ctx context.Context, request *DescribeClusterMachinesRequest) (response *DescribeClusterMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeClusterMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// 查询集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 查询集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMNOAUTH = "InternalError.CamNoAuth"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
//  INTERNALERROR_QUOTAMAXCLSLIMIT = "InternalError.QuotaMaxClsLimit"
//  INTERNALERROR_QUOTAMAXNODLIMIT = "InternalError.QuotaMaxNodLimit"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthCheckPoliciesRequest() (request *DescribeHealthCheckPoliciesRequest) {
    request = &DescribeHealthCheckPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckPolicies")
    
    
    return
}

func NewDescribeHealthCheckPoliciesResponse() (response *DescribeHealthCheckPoliciesResponse) {
    response = &DescribeHealthCheckPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthCheckPolicies
// 查询健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeHealthCheckPolicies(request *DescribeHealthCheckPoliciesRequest) (response *DescribeHealthCheckPoliciesResponse, err error) {
    return c.DescribeHealthCheckPoliciesWithContext(context.Background(), request)
}

// DescribeHealthCheckPolicies
// 查询健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeHealthCheckPoliciesWithContext(ctx context.Context, request *DescribeHealthCheckPoliciesRequest) (response *DescribeHealthCheckPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeHealthCheckPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeHealthCheckPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthCheckPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthCheckPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthCheckPolicyBindingsRequest() (request *DescribeHealthCheckPolicyBindingsRequest) {
    request = &DescribeHealthCheckPolicyBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckPolicyBindings")
    
    
    return
}

func NewDescribeHealthCheckPolicyBindingsResponse() (response *DescribeHealthCheckPolicyBindingsResponse) {
    response = &DescribeHealthCheckPolicyBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthCheckPolicyBindings
// 查询健康检测策略绑定关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeHealthCheckPolicyBindings(request *DescribeHealthCheckPolicyBindingsRequest) (response *DescribeHealthCheckPolicyBindingsResponse, err error) {
    return c.DescribeHealthCheckPolicyBindingsWithContext(context.Background(), request)
}

// DescribeHealthCheckPolicyBindings
// 查询健康检测策略绑定关系
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeHealthCheckPolicyBindingsWithContext(ctx context.Context, request *DescribeHealthCheckPolicyBindingsRequest) (response *DescribeHealthCheckPolicyBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeHealthCheckPolicyBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeHealthCheckPolicyBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthCheckPolicyBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthCheckPolicyBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthCheckTemplateRequest() (request *DescribeHealthCheckTemplateRequest) {
    request = &DescribeHealthCheckTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeHealthCheckTemplate")
    
    
    return
}

func NewDescribeHealthCheckTemplateResponse() (response *DescribeHealthCheckTemplateResponse) {
    response = &DescribeHealthCheckTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthCheckTemplate
// 查询健康检测策略模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeHealthCheckTemplate(request *DescribeHealthCheckTemplateRequest) (response *DescribeHealthCheckTemplateResponse, err error) {
    return c.DescribeHealthCheckTemplateWithContext(context.Background(), request)
}

// DescribeHealthCheckTemplate
// 查询健康检测策略模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeHealthCheckTemplateWithContext(ctx context.Context, request *DescribeHealthCheckTemplateRequest) (response *DescribeHealthCheckTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeHealthCheckTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeHealthCheckTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthCheckTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthCheckTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodePoolsRequest() (request *DescribeNodePoolsRequest) {
    request = &DescribeNodePoolsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "DescribeNodePools")
    
    
    return
}

func NewDescribeNodePoolsResponse() (response *DescribeNodePoolsResponse) {
    response = &DescribeNodePoolsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNodePools
// 查询 TKE 节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeNodePools(request *DescribeNodePoolsRequest) (response *DescribeNodePoolsResponse, err error) {
    return c.DescribeNodePoolsWithContext(context.Background(), request)
}

// DescribeNodePools
// 查询 TKE 节点池列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"
//  RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
func (c *Client) DescribeNodePoolsWithContext(ctx context.Context, request *DescribeNodePoolsRequest) (response *DescribeNodePoolsResponse, err error) {
    if request == nil {
        request = NewDescribeNodePoolsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "DescribeNodePools")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodePoolsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterMachineRequest() (request *ModifyClusterMachineRequest) {
    request = &ModifyClusterMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyClusterMachine")
    
    
    return
}

func NewModifyClusterMachineResponse() (response *ModifyClusterMachineResponse) {
    response = &ModifyClusterMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterMachine
// 修改原生节点
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyClusterMachine(request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
    return c.ModifyClusterMachineWithContext(context.Background(), request)
}

// ModifyClusterMachine
// 修改原生节点
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyClusterMachineWithContext(ctx context.Context, request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
    if request == nil {
        request = NewModifyClusterMachineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "ModifyClusterMachine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterMachine require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterMachineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHealthCheckPolicyRequest() (request *ModifyHealthCheckPolicyRequest) {
    request = &ModifyHealthCheckPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyHealthCheckPolicy")
    
    
    return
}

func NewModifyHealthCheckPolicyResponse() (response *ModifyHealthCheckPolicyResponse) {
    response = &ModifyHealthCheckPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHealthCheckPolicy
// 修改健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyHealthCheckPolicy(request *ModifyHealthCheckPolicyRequest) (response *ModifyHealthCheckPolicyResponse, err error) {
    return c.ModifyHealthCheckPolicyWithContext(context.Background(), request)
}

// ModifyHealthCheckPolicy
// 修改健康检测策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyHealthCheckPolicyWithContext(ctx context.Context, request *ModifyHealthCheckPolicyRequest) (response *ModifyHealthCheckPolicyResponse, err error) {
    if request == nil {
        request = NewModifyHealthCheckPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "ModifyHealthCheckPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHealthCheckPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHealthCheckPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodePoolRequest() (request *ModifyNodePoolRequest) {
    request = &ModifyNodePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "ModifyNodePool")
    
    
    return
}

func NewModifyNodePoolResponse() (response *ModifyNodePoolResponse) {
    response = &ModifyNodePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNodePool
// 更新 TKE 节点池
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNodePool(request *ModifyNodePoolRequest) (response *ModifyNodePoolResponse, err error) {
    return c.ModifyNodePoolWithContext(context.Background(), request)
}

// ModifyNodePool
// 更新 TKE 节点池
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNodePoolWithContext(ctx context.Context, request *ModifyNodePoolRequest) (response *ModifyNodePoolResponse, err error) {
    if request == nil {
        request = NewModifyNodePoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "ModifyNodePool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodePoolResponse()
    err = c.Send(request, response)
    return
}

func NewRebootMachinesRequest() (request *RebootMachinesRequest) {
    request = &RebootMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "RebootMachines")
    
    
    return
}

func NewRebootMachinesResponse() (response *RebootMachinesResponse) {
    response = &RebootMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootMachines
// 重启原生节点实例
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RebootMachines(request *RebootMachinesRequest) (response *RebootMachinesResponse, err error) {
    return c.RebootMachinesWithContext(context.Background(), request)
}

// RebootMachines
// 重启原生节点实例
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RebootMachinesWithContext(ctx context.Context, request *RebootMachinesRequest) (response *RebootMachinesResponse, err error) {
    if request == nil {
        request = NewRebootMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "RebootMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewSetMachineLoginRequest() (request *SetMachineLoginRequest) {
    request = &SetMachineLoginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "SetMachineLogin")
    
    
    return
}

func NewSetMachineLoginResponse() (response *SetMachineLoginResponse) {
    response = &SetMachineLoginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetMachineLogin
// 设置是否开启节点登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) SetMachineLogin(request *SetMachineLoginRequest) (response *SetMachineLoginResponse, err error) {
    return c.SetMachineLoginWithContext(context.Background(), request)
}

// SetMachineLogin
// 设置是否开启节点登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) SetMachineLoginWithContext(ctx context.Context, request *SetMachineLoginRequest) (response *SetMachineLoginResponse, err error) {
    if request == nil {
        request = NewSetMachineLoginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "SetMachineLogin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetMachineLogin require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetMachineLoginResponse()
    err = c.Send(request, response)
    return
}

func NewStartMachinesRequest() (request *StartMachinesRequest) {
    request = &StartMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "StartMachines")
    
    
    return
}

func NewStartMachinesResponse() (response *StartMachinesResponse) {
    response = &StartMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartMachines
// 本接口 (StartMachines) 用于启动一个或多个原生节点实例。
//
// 
//
// 只有状态为 Stopped 的实例才可以进行此操作。
//
// 接口调用成功后，等待一分钟左右，实例会进入 Running 状态。
//
// 支持批量操作。每次请求批量实例的上限为100。
//
// 本接口为同步接口，启动实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为 Running，则代表启动实例操作成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartMachines(request *StartMachinesRequest) (response *StartMachinesResponse, err error) {
    return c.StartMachinesWithContext(context.Background(), request)
}

// StartMachines
// 本接口 (StartMachines) 用于启动一个或多个原生节点实例。
//
// 
//
// 只有状态为 Stopped 的实例才可以进行此操作。
//
// 接口调用成功后，等待一分钟左右，实例会进入 Running 状态。
//
// 支持批量操作。每次请求批量实例的上限为100。
//
// 本接口为同步接口，启动实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为 Running，则代表启动实例操作成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartMachinesWithContext(ctx context.Context, request *StartMachinesRequest) (response *StartMachinesResponse, err error) {
    if request == nil {
        request = NewStartMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "StartMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewStopMachinesRequest() (request *StopMachinesRequest) {
    request = &StopMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tke", APIVersion, "StopMachines")
    
    
    return
}

func NewStopMachinesResponse() (response *StopMachinesResponse) {
    response = &StopMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopMachines
// 本接口 (StopMachines) 用于关闭一个或多个原生节点实例。
//
// 
//
// 只有状态为 Running 的实例才可以进行此操作。
//
// 接口调用成功时，实例会进入 Stopping 状态；关闭实例成功时，实例会进入 Stopped 状态。
//
// 支持强制关闭。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
//
// 支持批量操作。每次请求批量实例的上限为 100。
//
// 本接口为同步接口，关闭实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为stopped_with_charging，则代表关闭实例操作成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopMachines(request *StopMachinesRequest) (response *StopMachinesResponse, err error) {
    return c.StopMachinesWithContext(context.Background(), request)
}

// StopMachines
// 本接口 (StopMachines) 用于关闭一个或多个原生节点实例。
//
// 
//
// 只有状态为 Running 的实例才可以进行此操作。
//
// 接口调用成功时，实例会进入 Stopping 状态；关闭实例成功时，实例会进入 Stopped 状态。
//
// 支持强制关闭。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
//
// 支持批量操作。每次请求批量实例的上限为 100。
//
// 本接口为同步接口，关闭实例请求发送成功后会返回一个RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeClusterInstances 接口查询，如果实例的状态为stopped_with_charging，则代表关闭实例操作成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopMachinesWithContext(ctx context.Context, request *StopMachinesRequest) (response *StopMachinesResponse, err error) {
    if request == nil {
        request = NewStopMachinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tke", APIVersion, "StopMachines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMachinesResponse()
    err = c.Send(request, response)
    return
}
