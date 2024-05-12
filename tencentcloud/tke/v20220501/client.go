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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNodePoolResponse()
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
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
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
//  INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodePools require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodePoolsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodePoolResponse()
    err = c.Send(request, response)
    return
}
