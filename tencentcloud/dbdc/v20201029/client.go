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

package v20201029

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-29"

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


func NewAddNodesToDBCustomClusterRequest() (request *AddNodesToDBCustomClusterRequest) {
    request = &AddNodesToDBCustomClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "AddNodesToDBCustomCluster")
    
    
    return
}

func NewAddNodesToDBCustomClusterResponse() (response *AddNodesToDBCustomClusterResponse) {
    response = &AddNodesToDBCustomClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNodesToDBCustomCluster
// 该接口（AddNodesToDBCustomCluster）用于为 DB Custom 集群添加已存在的节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddNodesToDBCustomCluster(request *AddNodesToDBCustomClusterRequest) (response *AddNodesToDBCustomClusterResponse, err error) {
    return c.AddNodesToDBCustomClusterWithContext(context.Background(), request)
}

// AddNodesToDBCustomCluster
// 该接口（AddNodesToDBCustomCluster）用于为 DB Custom 集群添加已存在的节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddNodesToDBCustomClusterWithContext(ctx context.Context, request *AddNodesToDBCustomClusterRequest) (response *AddNodesToDBCustomClusterResponse, err error) {
    if request == nil {
        request = NewAddNodesToDBCustomClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "AddNodesToDBCustomCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNodesToDBCustomCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNodesToDBCustomClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCheckRoleAuthorizedRequest() (request *CheckRoleAuthorizedRequest) {
    request = &CheckRoleAuthorizedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "CheckRoleAuthorized")
    
    
    return
}

func NewCheckRoleAuthorizedResponse() (response *CheckRoleAuthorizedResponse) {
    response = &CheckRoleAuthorizedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckRoleAuthorized
// 检查服务相关角色是否已创建
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) CheckRoleAuthorized(request *CheckRoleAuthorizedRequest) (response *CheckRoleAuthorizedResponse, err error) {
    return c.CheckRoleAuthorizedWithContext(context.Background(), request)
}

// CheckRoleAuthorized
// 检查服务相关角色是否已创建
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) CheckRoleAuthorizedWithContext(ctx context.Context, request *CheckRoleAuthorizedRequest) (response *CheckRoleAuthorizedResponse, err error) {
    if request == nil {
        request = NewCheckRoleAuthorizedRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "CheckRoleAuthorized")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRoleAuthorized require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRoleAuthorizedResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBCustomClusterRequest() (request *CreateDBCustomClusterRequest) {
    request = &CreateDBCustomClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "CreateDBCustomCluster")
    
    
    return
}

func NewCreateDBCustomClusterResponse() (response *CreateDBCustomClusterResponse) {
    response = &CreateDBCustomClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBCustomCluster
// 该接口（CreateDBCustomCluster）用于创建 DB Custom 集群。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomCluster(request *CreateDBCustomClusterRequest) (response *CreateDBCustomClusterResponse, err error) {
    return c.CreateDBCustomClusterWithContext(context.Background(), request)
}

// CreateDBCustomCluster
// 该接口（CreateDBCustomCluster）用于创建 DB Custom 集群。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomClusterWithContext(ctx context.Context, request *CreateDBCustomClusterRequest) (response *CreateDBCustomClusterResponse, err error) {
    if request == nil {
        request = NewCreateDBCustomClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "CreateDBCustomCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBCustomCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBCustomClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBCustomDisasterRecoverGroupRequest() (request *CreateDBCustomDisasterRecoverGroupRequest) {
    request = &CreateDBCustomDisasterRecoverGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "CreateDBCustomDisasterRecoverGroup")
    
    
    return
}

func NewCreateDBCustomDisasterRecoverGroupResponse() (response *CreateDBCustomDisasterRecoverGroupResponse) {
    response = &CreateDBCustomDisasterRecoverGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBCustomDisasterRecoverGroup
// 该接口（CreateDBCustomDisasterRecoverGroup）用于创建 DB Custom 置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomDisasterRecoverGroup(request *CreateDBCustomDisasterRecoverGroupRequest) (response *CreateDBCustomDisasterRecoverGroupResponse, err error) {
    return c.CreateDBCustomDisasterRecoverGroupWithContext(context.Background(), request)
}

// CreateDBCustomDisasterRecoverGroup
// 该接口（CreateDBCustomDisasterRecoverGroup）用于创建 DB Custom 置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomDisasterRecoverGroupWithContext(ctx context.Context, request *CreateDBCustomDisasterRecoverGroupRequest) (response *CreateDBCustomDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewCreateDBCustomDisasterRecoverGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "CreateDBCustomDisasterRecoverGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBCustomDisasterRecoverGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBCustomDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBCustomNodesRequest() (request *CreateDBCustomNodesRequest) {
    request = &CreateDBCustomNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "CreateDBCustomNodes")
    
    
    return
}

func NewCreateDBCustomNodesResponse() (response *CreateDBCustomNodesResponse) {
    response = &CreateDBCustomNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBCustomNodes
// 该接口（CreateDBCustomNodes）用于创建 DB Custom 节点(需支付)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomNodes(request *CreateDBCustomNodesRequest) (response *CreateDBCustomNodesResponse, err error) {
    return c.CreateDBCustomNodesWithContext(context.Background(), request)
}

// CreateDBCustomNodes
// 该接口（CreateDBCustomNodes）用于创建 DB Custom 节点(需支付)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateDBCustomNodesWithContext(ctx context.Context, request *CreateDBCustomNodesRequest) (response *CreateDBCustomNodesResponse, err error) {
    if request == nil {
        request = NewCreateDBCustomNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "CreateDBCustomNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBCustomNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBCustomNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBCustomDisasterRecoverGroupsRequest() (request *DeleteDBCustomDisasterRecoverGroupsRequest) {
    request = &DeleteDBCustomDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DeleteDBCustomDisasterRecoverGroups")
    
    
    return
}

func NewDeleteDBCustomDisasterRecoverGroupsResponse() (response *DeleteDBCustomDisasterRecoverGroupsResponse) {
    response = &DeleteDBCustomDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBCustomDisasterRecoverGroups
// 该接口（DeleteDBCustomDisasterRecoverGroups）用于删除 DB Custom 置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDBCustomDisasterRecoverGroups(request *DeleteDBCustomDisasterRecoverGroupsRequest) (response *DeleteDBCustomDisasterRecoverGroupsResponse, err error) {
    return c.DeleteDBCustomDisasterRecoverGroupsWithContext(context.Background(), request)
}

// DeleteDBCustomDisasterRecoverGroups
// 该接口（DeleteDBCustomDisasterRecoverGroups）用于删除 DB Custom 置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDBCustomDisasterRecoverGroupsWithContext(ctx context.Context, request *DeleteDBCustomDisasterRecoverGroupsRequest) (response *DeleteDBCustomDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDBCustomDisasterRecoverGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DeleteDBCustomDisasterRecoverGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBCustomDisasterRecoverGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBCustomDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBCustomNodesDisasterRecoverGroupRequest() (request *DeleteDBCustomNodesDisasterRecoverGroupRequest) {
    request = &DeleteDBCustomNodesDisasterRecoverGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DeleteDBCustomNodesDisasterRecoverGroup")
    
    
    return
}

func NewDeleteDBCustomNodesDisasterRecoverGroupResponse() (response *DeleteDBCustomNodesDisasterRecoverGroupResponse) {
    response = &DeleteDBCustomNodesDisasterRecoverGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBCustomNodesDisasterRecoverGroup
// 该接口（DeleteDBCustomNodesDisasterRecoverGroup）用于移除 DB Custom 节点的置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDBCustomNodesDisasterRecoverGroup(request *DeleteDBCustomNodesDisasterRecoverGroupRequest) (response *DeleteDBCustomNodesDisasterRecoverGroupResponse, err error) {
    return c.DeleteDBCustomNodesDisasterRecoverGroupWithContext(context.Background(), request)
}

// DeleteDBCustomNodesDisasterRecoverGroup
// 该接口（DeleteDBCustomNodesDisasterRecoverGroup）用于移除 DB Custom 节点的置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDBCustomNodesDisasterRecoverGroupWithContext(ctx context.Context, request *DeleteDBCustomNodesDisasterRecoverGroupRequest) (response *DeleteDBCustomNodesDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewDeleteDBCustomNodesDisasterRecoverGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DeleteDBCustomNodesDisasterRecoverGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBCustomNodesDisasterRecoverGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBCustomNodesDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterDetailRequest() (request *DescribeDBCustomClusterDetailRequest) {
    request = &DescribeDBCustomClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterDetail")
    
    
    return
}

func NewDescribeDBCustomClusterDetailResponse() (response *DescribeDBCustomClusterDetailResponse) {
    response = &DescribeDBCustomClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterDetail
// 该接口(DescribeDBCustomClusterDetail) 用于查询 DB Custom 集群的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterDetail(request *DescribeDBCustomClusterDetailRequest) (response *DescribeDBCustomClusterDetailResponse, err error) {
    return c.DescribeDBCustomClusterDetailWithContext(context.Background(), request)
}

// DescribeDBCustomClusterDetail
// 该接口(DescribeDBCustomClusterDetail) 用于查询 DB Custom 集群的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterDetailWithContext(ctx context.Context, request *DescribeDBCustomClusterDetailRequest) (response *DescribeDBCustomClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterKubeconfigRequest() (request *DescribeDBCustomClusterKubeconfigRequest) {
    request = &DescribeDBCustomClusterKubeconfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterKubeconfig")
    
    
    return
}

func NewDescribeDBCustomClusterKubeconfigResponse() (response *DescribeDBCustomClusterKubeconfigResponse) {
    response = &DescribeDBCustomClusterKubeconfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterKubeconfig
// 该接口（DescribeDBCustomClusterKubeconfig）用于查询 DB Custom 集群 Kubeconfig。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterKubeconfig(request *DescribeDBCustomClusterKubeconfigRequest) (response *DescribeDBCustomClusterKubeconfigResponse, err error) {
    return c.DescribeDBCustomClusterKubeconfigWithContext(context.Background(), request)
}

// DescribeDBCustomClusterKubeconfig
// 该接口（DescribeDBCustomClusterKubeconfig）用于查询 DB Custom 集群 Kubeconfig。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterKubeconfigWithContext(ctx context.Context, request *DescribeDBCustomClusterKubeconfigRequest) (response *DescribeDBCustomClusterKubeconfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterKubeconfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterKubeconfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterKubeconfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterKubeconfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterNodeConfigRequest() (request *DescribeDBCustomClusterNodeConfigRequest) {
    request = &DescribeDBCustomClusterNodeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterNodeConfig")
    
    
    return
}

func NewDescribeDBCustomClusterNodeConfigResponse() (response *DescribeDBCustomClusterNodeConfigResponse) {
    response = &DescribeDBCustomClusterNodeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterNodeConfig
// 该接口（DescribeDBCustomClusterNodeConfig）用于查询 DB Custom 集群内节点的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodeConfig(request *DescribeDBCustomClusterNodeConfigRequest) (response *DescribeDBCustomClusterNodeConfigResponse, err error) {
    return c.DescribeDBCustomClusterNodeConfigWithContext(context.Background(), request)
}

// DescribeDBCustomClusterNodeConfig
// 该接口（DescribeDBCustomClusterNodeConfig）用于查询 DB Custom 集群内节点的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodeConfigWithContext(ctx context.Context, request *DescribeDBCustomClusterNodeConfigRequest) (response *DescribeDBCustomClusterNodeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterNodeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterNodeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterNodeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterNodeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterNodeResourcesRequest() (request *DescribeDBCustomClusterNodeResourcesRequest) {
    request = &DescribeDBCustomClusterNodeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterNodeResources")
    
    
    return
}

func NewDescribeDBCustomClusterNodeResourcesResponse() (response *DescribeDBCustomClusterNodeResourcesResponse) {
    response = &DescribeDBCustomClusterNodeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterNodeResources
// 该接口（DescribeDBCustomClusterNodeResources）用于查询 DB Custom 集群内节点的资源信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodeResources(request *DescribeDBCustomClusterNodeResourcesRequest) (response *DescribeDBCustomClusterNodeResourcesResponse, err error) {
    return c.DescribeDBCustomClusterNodeResourcesWithContext(context.Background(), request)
}

// DescribeDBCustomClusterNodeResources
// 该接口（DescribeDBCustomClusterNodeResources）用于查询 DB Custom 集群内节点的资源信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodeResourcesWithContext(ctx context.Context, request *DescribeDBCustomClusterNodeResourcesRequest) (response *DescribeDBCustomClusterNodeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterNodeResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterNodeResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterNodeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterNodeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterNodesRequest() (request *DescribeDBCustomClusterNodesRequest) {
    request = &DescribeDBCustomClusterNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterNodes")
    
    
    return
}

func NewDescribeDBCustomClusterNodesResponse() (response *DescribeDBCustomClusterNodesResponse) {
    response = &DescribeDBCustomClusterNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterNodes
// 该接口（DescribeDBCustomClusterNodes）用于查询 DB Custom 集群中的节点列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodes(request *DescribeDBCustomClusterNodesRequest) (response *DescribeDBCustomClusterNodesResponse, err error) {
    return c.DescribeDBCustomClusterNodesWithContext(context.Background(), request)
}

// DescribeDBCustomClusterNodes
// 该接口（DescribeDBCustomClusterNodes）用于查询 DB Custom 集群中的节点列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterNodesWithContext(ctx context.Context, request *DescribeDBCustomClusterNodesRequest) (response *DescribeDBCustomClusterNodesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClusterResourcesRequest() (request *DescribeDBCustomClusterResourcesRequest) {
    request = &DescribeDBCustomClusterResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusterResources")
    
    
    return
}

func NewDescribeDBCustomClusterResourcesResponse() (response *DescribeDBCustomClusterResourcesResponse) {
    response = &DescribeDBCustomClusterResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusterResources
// 该接口（DescribeDBCustomClusterResources）用于查询 DB Custom 集群的资源信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterResources(request *DescribeDBCustomClusterResourcesRequest) (response *DescribeDBCustomClusterResourcesResponse, err error) {
    return c.DescribeDBCustomClusterResourcesWithContext(context.Background(), request)
}

// DescribeDBCustomClusterResources
// 该接口（DescribeDBCustomClusterResources）用于查询 DB Custom 集群的资源信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusterResourcesWithContext(ctx context.Context, request *DescribeDBCustomClusterResourcesRequest) (response *DescribeDBCustomClusterResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClusterResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusterResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusterResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClusterResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomClustersRequest() (request *DescribeDBCustomClustersRequest) {
    request = &DescribeDBCustomClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomClusters")
    
    
    return
}

func NewDescribeDBCustomClustersResponse() (response *DescribeDBCustomClustersResponse) {
    response = &DescribeDBCustomClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomClusters
// 该接口（DescribeDBCustomClusters）为 DB Custom 集群列表查询接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClusters(request *DescribeDBCustomClustersRequest) (response *DescribeDBCustomClustersResponse, err error) {
    return c.DescribeDBCustomClustersWithContext(context.Background(), request)
}

// DescribeDBCustomClusters
// 该接口（DescribeDBCustomClusters）为 DB Custom 集群列表查询接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomClustersWithContext(ctx context.Context, request *DescribeDBCustomClustersRequest) (response *DescribeDBCustomClustersResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomDisasterRecoverGroupQuotaRequest() (request *DescribeDBCustomDisasterRecoverGroupQuotaRequest) {
    request = &DescribeDBCustomDisasterRecoverGroupQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomDisasterRecoverGroupQuota")
    
    
    return
}

func NewDescribeDBCustomDisasterRecoverGroupQuotaResponse() (response *DescribeDBCustomDisasterRecoverGroupQuotaResponse) {
    response = &DescribeDBCustomDisasterRecoverGroupQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomDisasterRecoverGroupQuota
// 该接口（DescribeDBCustomDisasterRecoverGroupQuota）用于查询 DB Custom 置放群组配额。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomDisasterRecoverGroupQuota(request *DescribeDBCustomDisasterRecoverGroupQuotaRequest) (response *DescribeDBCustomDisasterRecoverGroupQuotaResponse, err error) {
    return c.DescribeDBCustomDisasterRecoverGroupQuotaWithContext(context.Background(), request)
}

// DescribeDBCustomDisasterRecoverGroupQuota
// 该接口（DescribeDBCustomDisasterRecoverGroupQuota）用于查询 DB Custom 置放群组配额。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomDisasterRecoverGroupQuotaWithContext(ctx context.Context, request *DescribeDBCustomDisasterRecoverGroupQuotaRequest) (response *DescribeDBCustomDisasterRecoverGroupQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomDisasterRecoverGroupQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomDisasterRecoverGroupQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomDisasterRecoverGroupQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomDisasterRecoverGroupQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomDisasterRecoverGroupsRequest() (request *DescribeDBCustomDisasterRecoverGroupsRequest) {
    request = &DescribeDBCustomDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomDisasterRecoverGroups")
    
    
    return
}

func NewDescribeDBCustomDisasterRecoverGroupsResponse() (response *DescribeDBCustomDisasterRecoverGroupsResponse) {
    response = &DescribeDBCustomDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomDisasterRecoverGroups
// 该接口（DescribeDBCustomDisasterRecoverGroups）用于查询 DB Custom 置放群组列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomDisasterRecoverGroups(request *DescribeDBCustomDisasterRecoverGroupsRequest) (response *DescribeDBCustomDisasterRecoverGroupsResponse, err error) {
    return c.DescribeDBCustomDisasterRecoverGroupsWithContext(context.Background(), request)
}

// DescribeDBCustomDisasterRecoverGroups
// 该接口（DescribeDBCustomDisasterRecoverGroups）用于查询 DB Custom 置放群组列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomDisasterRecoverGroupsWithContext(ctx context.Context, request *DescribeDBCustomDisasterRecoverGroupsRequest) (response *DescribeDBCustomDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomDisasterRecoverGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomDisasterRecoverGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomDisasterRecoverGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomImagesRequest() (request *DescribeDBCustomImagesRequest) {
    request = &DescribeDBCustomImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomImages")
    
    
    return
}

func NewDescribeDBCustomImagesResponse() (response *DescribeDBCustomImagesResponse) {
    response = &DescribeDBCustomImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomImages
// 该接口（DescribeDBCustomImages）用于查询 DB Custom 可用的操作系统镜像列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomImages(request *DescribeDBCustomImagesRequest) (response *DescribeDBCustomImagesResponse, err error) {
    return c.DescribeDBCustomImagesWithContext(context.Background(), request)
}

// DescribeDBCustomImages
// 该接口（DescribeDBCustomImages）用于查询 DB Custom 可用的操作系统镜像列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomImagesWithContext(ctx context.Context, request *DescribeDBCustomImagesRequest) (response *DescribeDBCustomImagesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomNodeSecurityGroupsRequest() (request *DescribeDBCustomNodeSecurityGroupsRequest) {
    request = &DescribeDBCustomNodeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomNodeSecurityGroups")
    
    
    return
}

func NewDescribeDBCustomNodeSecurityGroupsResponse() (response *DescribeDBCustomNodeSecurityGroupsResponse) {
    response = &DescribeDBCustomNodeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomNodeSecurityGroups
// 该接口（DescribeDBCustomNodeSecurityGroups）用于查询 DB Custom 节点安全组信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodeSecurityGroups(request *DescribeDBCustomNodeSecurityGroupsRequest) (response *DescribeDBCustomNodeSecurityGroupsResponse, err error) {
    return c.DescribeDBCustomNodeSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBCustomNodeSecurityGroups
// 该接口（DescribeDBCustomNodeSecurityGroups）用于查询 DB Custom 节点安全组信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodeSecurityGroupsWithContext(ctx context.Context, request *DescribeDBCustomNodeSecurityGroupsRequest) (response *DescribeDBCustomNodeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomNodeSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomNodeSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomNodeSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomNodeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomNodeTypesRequest() (request *DescribeDBCustomNodeTypesRequest) {
    request = &DescribeDBCustomNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomNodeTypes")
    
    
    return
}

func NewDescribeDBCustomNodeTypesResponse() (response *DescribeDBCustomNodeTypesResponse) {
    response = &DescribeDBCustomNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomNodeTypes
// 该接口(DescribeDBCustomNodeTypes) 用于查询 DB Custom 节点支持的机型信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodeTypes(request *DescribeDBCustomNodeTypesRequest) (response *DescribeDBCustomNodeTypesResponse, err error) {
    return c.DescribeDBCustomNodeTypesWithContext(context.Background(), request)
}

// DescribeDBCustomNodeTypes
// 该接口(DescribeDBCustomNodeTypes) 用于查询 DB Custom 节点支持的机型信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodeTypesWithContext(ctx context.Context, request *DescribeDBCustomNodeTypesRequest) (response *DescribeDBCustomNodeTypesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomNodeTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomNodeTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomNodesRequest() (request *DescribeDBCustomNodesRequest) {
    request = &DescribeDBCustomNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomNodes")
    
    
    return
}

func NewDescribeDBCustomNodesResponse() (response *DescribeDBCustomNodesResponse) {
    response = &DescribeDBCustomNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomNodes
// 该接口（DescribeDBCustomNodes）用于查询 DB Custom 节点列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodes(request *DescribeDBCustomNodesRequest) (response *DescribeDBCustomNodesResponse, err error) {
    return c.DescribeDBCustomNodesWithContext(context.Background(), request)
}

// DescribeDBCustomNodes
// 该接口（DescribeDBCustomNodes）用于查询 DB Custom 节点列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomNodesWithContext(ctx context.Context, request *DescribeDBCustomNodesRequest) (response *DescribeDBCustomNodesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomRegionsRequest() (request *DescribeDBCustomRegionsRequest) {
    request = &DescribeDBCustomRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomRegions")
    
    
    return
}

func NewDescribeDBCustomRegionsResponse() (response *DescribeDBCustomRegionsResponse) {
    response = &DescribeDBCustomRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomRegions
// 该接口(DescribeDBCustomRegions) 用于查询 DB Custom 支持的地域列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomRegions(request *DescribeDBCustomRegionsRequest) (response *DescribeDBCustomRegionsResponse, err error) {
    return c.DescribeDBCustomRegionsWithContext(context.Background(), request)
}

// DescribeDBCustomRegions
// 该接口(DescribeDBCustomRegions) 用于查询 DB Custom 支持的地域列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomRegionsWithContext(ctx context.Context, request *DescribeDBCustomRegionsRequest) (response *DescribeDBCustomRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomTaskStatusRequest() (request *DescribeDBCustomTaskStatusRequest) {
    request = &DescribeDBCustomTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomTaskStatus")
    
    
    return
}

func NewDescribeDBCustomTaskStatusResponse() (response *DescribeDBCustomTaskStatusResponse) {
    response = &DescribeDBCustomTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomTaskStatus
// 该接口（DescribeDBCustomTaskStatus）用于查询 DB Custom 任务的状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomTaskStatus(request *DescribeDBCustomTaskStatusRequest) (response *DescribeDBCustomTaskStatusResponse, err error) {
    return c.DescribeDBCustomTaskStatusWithContext(context.Background(), request)
}

// DescribeDBCustomTaskStatus
// 该接口（DescribeDBCustomTaskStatus）用于查询 DB Custom 任务的状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomTaskStatusWithContext(ctx context.Context, request *DescribeDBCustomTaskStatusRequest) (response *DescribeDBCustomTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCustomZonesRequest() (request *DescribeDBCustomZonesRequest) {
    request = &DescribeDBCustomZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBCustomZones")
    
    
    return
}

func NewDescribeDBCustomZonesResponse() (response *DescribeDBCustomZonesResponse) {
    response = &DescribeDBCustomZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBCustomZones
// 该接口(DescribeDBCustomZones) 用于查询指定地域的 DB Custom 可用区列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomZones(request *DescribeDBCustomZonesRequest) (response *DescribeDBCustomZonesResponse, err error) {
    return c.DescribeDBCustomZonesWithContext(context.Background(), request)
}

// DescribeDBCustomZones
// 该接口(DescribeDBCustomZones) 用于查询指定地域的 DB Custom 可用区列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDBCustomZonesWithContext(ctx context.Context, request *DescribeDBCustomZonesRequest) (response *DescribeDBCustomZonesResponse, err error) {
    if request == nil {
        request = NewDescribeDBCustomZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBCustomZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCustomZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCustomZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstances
// 本接口用于查询独享集群内的DB实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口用于查询独享集群内的DB实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostListRequest() (request *DescribeHostListRequest) {
    request = &DescribeHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeHostList")
    
    
    return
}

func NewDescribeHostListResponse() (response *DescribeHostListResponse) {
    response = &DescribeHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostList
// 本接口用于查询主机列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCEUNAVAILABLE_RESOURCESTATUSABNORMALERROR = "ResourceUnavailable.ResourceStatusAbnormalError"
func (c *Client) DescribeHostList(request *DescribeHostListRequest) (response *DescribeHostListResponse, err error) {
    return c.DescribeHostListWithContext(context.Background(), request)
}

// DescribeHostList
// 本接口用于查询主机列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCEUNAVAILABLE_RESOURCESTATUSABNORMALERROR = "ResourceUnavailable.ResourceStatusAbnormalError"
func (c *Client) DescribeHostListWithContext(ctx context.Context, request *DescribeHostListRequest) (response *DescribeHostListResponse, err error) {
    if request == nil {
        request = NewDescribeHostListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeHostList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
    request = &DescribeInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeInstanceDetail")
    
    
    return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
    response = &DescribeInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDetail
// 本接口用于查询独享集群详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    return c.DescribeInstanceDetailWithContext(context.Background(), request)
}

// DescribeInstanceDetail
// 本接口用于查询独享集群详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeInstanceDetailWithContext(ctx context.Context, request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// 本接口用于查询独享集群实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// 本接口用于查询独享集群实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  FAILEDOPERATION_QUERYSPECINFOERROR = "FailedOperation.QuerySpecInfoError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 根据不同地域不同用户，获取集群列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_RESOURCEPARAMETERERROR = "InvalidParameterValue.ResourceParameterError"
//  RESOURCENOTFOUND_FETCHRESOURCEERROR = "ResourceNotFound.FetchResourceError"
//  RESOURCENOTFOUND_FETCHRESOURCELISTERROR = "ResourceNotFound.FetchResourceListError"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 根据不同地域不同用户，获取集群列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_RESOURCEPARAMETERERROR = "InvalidParameterValue.ResourceParameterError"
//  RESOURCENOTFOUND_FETCHRESOURCEERROR = "ResourceNotFound.FetchResourceError"
//  RESOURCENOTFOUND_FETCHRESOURCELISTERROR = "ResourceNotFound.FetchResourceListError"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDBCustomClusterRequest() (request *DestroyDBCustomClusterRequest) {
    request = &DestroyDBCustomClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DestroyDBCustomCluster")
    
    
    return
}

func NewDestroyDBCustomClusterResponse() (response *DestroyDBCustomClusterResponse) {
    response = &DestroyDBCustomClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyDBCustomCluster
// 该接口（DestroyDBCustomCluster）用于销毁 DB Custom 集群。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyDBCustomCluster(request *DestroyDBCustomClusterRequest) (response *DestroyDBCustomClusterResponse, err error) {
    return c.DestroyDBCustomClusterWithContext(context.Background(), request)
}

// DestroyDBCustomCluster
// 该接口（DestroyDBCustomCluster）用于销毁 DB Custom 集群。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyDBCustomClusterWithContext(ctx context.Context, request *DestroyDBCustomClusterRequest) (response *DestroyDBCustomClusterResponse, err error) {
    if request == nil {
        request = NewDestroyDBCustomClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DestroyDBCustomCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyDBCustomCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyDBCustomClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDBCustomNodeRequest() (request *DestroyDBCustomNodeRequest) {
    request = &DestroyDBCustomNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "DestroyDBCustomNode")
    
    
    return
}

func NewDestroyDBCustomNodeResponse() (response *DestroyDBCustomNodeResponse) {
    response = &DestroyDBCustomNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyDBCustomNode
// 该接口（DestroyDBCustomNode）用于销毁 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyDBCustomNode(request *DestroyDBCustomNodeRequest) (response *DestroyDBCustomNodeResponse, err error) {
    return c.DestroyDBCustomNodeWithContext(context.Background(), request)
}

// DestroyDBCustomNode
// 该接口（DestroyDBCustomNode）用于销毁 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyDBCustomNodeWithContext(ctx context.Context, request *DestroyDBCustomNodeRequest) (response *DestroyDBCustomNodeResponse, err error) {
    if request == nil {
        request = NewDestroyDBCustomNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "DestroyDBCustomNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyDBCustomNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyDBCustomNodeResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBCustomNodeRequest() (request *IsolateDBCustomNodeRequest) {
    request = &IsolateDBCustomNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "IsolateDBCustomNode")
    
    
    return
}

func NewIsolateDBCustomNodeResponse() (response *IsolateDBCustomNodeResponse) {
    response = &IsolateDBCustomNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBCustomNode
// 该接口 (IsolateDBCustomNode) 用于隔离 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) IsolateDBCustomNode(request *IsolateDBCustomNodeRequest) (response *IsolateDBCustomNodeResponse, err error) {
    return c.IsolateDBCustomNodeWithContext(context.Background(), request)
}

// IsolateDBCustomNode
// 该接口 (IsolateDBCustomNode) 用于隔离 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) IsolateDBCustomNodeWithContext(ctx context.Context, request *IsolateDBCustomNodeRequest) (response *IsolateDBCustomNodeResponse, err error) {
    if request == nil {
        request = NewIsolateDBCustomNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "IsolateDBCustomNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBCustomNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBCustomNodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomClusterAttributesRequest() (request *ModifyDBCustomClusterAttributesRequest) {
    request = &ModifyDBCustomClusterAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomClusterAttributes")
    
    
    return
}

func NewModifyDBCustomClusterAttributesResponse() (response *ModifyDBCustomClusterAttributesResponse) {
    response = &ModifyDBCustomClusterAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomClusterAttributes
// 该接口（ModifyDBCustomClusterAttributes）用于修改 DB Custom 集群的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterAttributes(request *ModifyDBCustomClusterAttributesRequest) (response *ModifyDBCustomClusterAttributesResponse, err error) {
    return c.ModifyDBCustomClusterAttributesWithContext(context.Background(), request)
}

// ModifyDBCustomClusterAttributes
// 该接口（ModifyDBCustomClusterAttributes）用于修改 DB Custom 集群的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterAttributesWithContext(ctx context.Context, request *ModifyDBCustomClusterAttributesRequest) (response *ModifyDBCustomClusterAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomClusterAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomClusterAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomClusterAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomClusterAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomClusterNodeConfigRequest() (request *ModifyDBCustomClusterNodeConfigRequest) {
    request = &ModifyDBCustomClusterNodeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomClusterNodeConfig")
    
    
    return
}

func NewModifyDBCustomClusterNodeConfigResponse() (response *ModifyDBCustomClusterNodeConfigResponse) {
    response = &ModifyDBCustomClusterNodeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomClusterNodeConfig
// 该接口（ModifyDBCustomClusterNodeConfig）用于修改 DB Custom 集群中节点的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterNodeConfig(request *ModifyDBCustomClusterNodeConfigRequest) (response *ModifyDBCustomClusterNodeConfigResponse, err error) {
    return c.ModifyDBCustomClusterNodeConfigWithContext(context.Background(), request)
}

// ModifyDBCustomClusterNodeConfig
// 该接口（ModifyDBCustomClusterNodeConfig）用于修改 DB Custom 集群中节点的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterNodeConfigWithContext(ctx context.Context, request *ModifyDBCustomClusterNodeConfigRequest) (response *ModifyDBCustomClusterNodeConfigResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomClusterNodeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomClusterNodeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomClusterNodeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomClusterNodeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomClusterTagsRequest() (request *ModifyDBCustomClusterTagsRequest) {
    request = &ModifyDBCustomClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomClusterTags")
    
    
    return
}

func NewModifyDBCustomClusterTagsResponse() (response *ModifyDBCustomClusterTagsResponse) {
    response = &ModifyDBCustomClusterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomClusterTags
// 该接口（ModifyDBCustomClusterTags）用于修改 DB Custom 集群绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterTags(request *ModifyDBCustomClusterTagsRequest) (response *ModifyDBCustomClusterTagsResponse, err error) {
    return c.ModifyDBCustomClusterTagsWithContext(context.Background(), request)
}

// ModifyDBCustomClusterTags
// 该接口（ModifyDBCustomClusterTags）用于修改 DB Custom 集群绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomClusterTagsWithContext(ctx context.Context, request *ModifyDBCustomClusterTagsRequest) (response *ModifyDBCustomClusterTagsResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomClusterTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomClusterTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomClusterTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomClusterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomDisasterRecoverGroupAttributeRequest() (request *ModifyDBCustomDisasterRecoverGroupAttributeRequest) {
    request = &ModifyDBCustomDisasterRecoverGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomDisasterRecoverGroupAttribute")
    
    
    return
}

func NewModifyDBCustomDisasterRecoverGroupAttributeResponse() (response *ModifyDBCustomDisasterRecoverGroupAttributeResponse) {
    response = &ModifyDBCustomDisasterRecoverGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomDisasterRecoverGroupAttribute
// 该接口（ModifyDBCustomDisasterRecoverGroupAttribute）用于修改 DB Custom 置放群组的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomDisasterRecoverGroupAttribute(request *ModifyDBCustomDisasterRecoverGroupAttributeRequest) (response *ModifyDBCustomDisasterRecoverGroupAttributeResponse, err error) {
    return c.ModifyDBCustomDisasterRecoverGroupAttributeWithContext(context.Background(), request)
}

// ModifyDBCustomDisasterRecoverGroupAttribute
// 该接口（ModifyDBCustomDisasterRecoverGroupAttribute）用于修改 DB Custom 置放群组的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomDisasterRecoverGroupAttributeWithContext(ctx context.Context, request *ModifyDBCustomDisasterRecoverGroupAttributeRequest) (response *ModifyDBCustomDisasterRecoverGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomDisasterRecoverGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomDisasterRecoverGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomDisasterRecoverGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomDisasterRecoverGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomDisasterRecoverGroupTagsRequest() (request *ModifyDBCustomDisasterRecoverGroupTagsRequest) {
    request = &ModifyDBCustomDisasterRecoverGroupTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomDisasterRecoverGroupTags")
    
    
    return
}

func NewModifyDBCustomDisasterRecoverGroupTagsResponse() (response *ModifyDBCustomDisasterRecoverGroupTagsResponse) {
    response = &ModifyDBCustomDisasterRecoverGroupTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomDisasterRecoverGroupTags
// 该接口（ModifyDBCustomDisasterRecoverGroupTags）用于修改 DB Custom 置放群组绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomDisasterRecoverGroupTags(request *ModifyDBCustomDisasterRecoverGroupTagsRequest) (response *ModifyDBCustomDisasterRecoverGroupTagsResponse, err error) {
    return c.ModifyDBCustomDisasterRecoverGroupTagsWithContext(context.Background(), request)
}

// ModifyDBCustomDisasterRecoverGroupTags
// 该接口（ModifyDBCustomDisasterRecoverGroupTags）用于修改 DB Custom 置放群组绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomDisasterRecoverGroupTagsWithContext(ctx context.Context, request *ModifyDBCustomDisasterRecoverGroupTagsRequest) (response *ModifyDBCustomDisasterRecoverGroupTagsResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomDisasterRecoverGroupTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomDisasterRecoverGroupTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomDisasterRecoverGroupTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomDisasterRecoverGroupTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomNodeAttributesRequest() (request *ModifyDBCustomNodeAttributesRequest) {
    request = &ModifyDBCustomNodeAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomNodeAttributes")
    
    
    return
}

func NewModifyDBCustomNodeAttributesResponse() (response *ModifyDBCustomNodeAttributesResponse) {
    response = &ModifyDBCustomNodeAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomNodeAttributes
// 该接口（ModifyDBCustomNodeAttributes）用于修改 DB Custom 节点的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeAttributes(request *ModifyDBCustomNodeAttributesRequest) (response *ModifyDBCustomNodeAttributesResponse, err error) {
    return c.ModifyDBCustomNodeAttributesWithContext(context.Background(), request)
}

// ModifyDBCustomNodeAttributes
// 该接口（ModifyDBCustomNodeAttributes）用于修改 DB Custom 节点的属性。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeAttributesWithContext(ctx context.Context, request *ModifyDBCustomNodeAttributesRequest) (response *ModifyDBCustomNodeAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomNodeAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomNodeAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomNodeAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomNodeAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomNodeSecurityGroupsRequest() (request *ModifyDBCustomNodeSecurityGroupsRequest) {
    request = &ModifyDBCustomNodeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomNodeSecurityGroups")
    
    
    return
}

func NewModifyDBCustomNodeSecurityGroupsResponse() (response *ModifyDBCustomNodeSecurityGroupsResponse) {
    response = &ModifyDBCustomNodeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomNodeSecurityGroups
// 该接口（ModifyDBCustomNodeSecurityGroups）用于修改 DB Custom 节点安全组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeSecurityGroups(request *ModifyDBCustomNodeSecurityGroupsRequest) (response *ModifyDBCustomNodeSecurityGroupsResponse, err error) {
    return c.ModifyDBCustomNodeSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBCustomNodeSecurityGroups
// 该接口（ModifyDBCustomNodeSecurityGroups）用于修改 DB Custom 节点安全组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeSecurityGroupsWithContext(ctx context.Context, request *ModifyDBCustomNodeSecurityGroupsRequest) (response *ModifyDBCustomNodeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomNodeSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomNodeSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomNodeSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomNodeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomNodeTagsRequest() (request *ModifyDBCustomNodeTagsRequest) {
    request = &ModifyDBCustomNodeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomNodeTags")
    
    
    return
}

func NewModifyDBCustomNodeTagsResponse() (response *ModifyDBCustomNodeTagsResponse) {
    response = &ModifyDBCustomNodeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomNodeTags
// 该接口（ModifyDBCustomNodeTags）用于修改 DB Custom 节点绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeTags(request *ModifyDBCustomNodeTagsRequest) (response *ModifyDBCustomNodeTagsResponse, err error) {
    return c.ModifyDBCustomNodeTagsWithContext(context.Background(), request)
}

// ModifyDBCustomNodeTags
// 该接口（ModifyDBCustomNodeTags）用于修改 DB Custom 节点绑定的标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodeTagsWithContext(ctx context.Context, request *ModifyDBCustomNodeTagsRequest) (response *ModifyDBCustomNodeTagsResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomNodeTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomNodeTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomNodeTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomNodeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomNodeTypeRequest() (request *ModifyDBCustomNodeTypeRequest) {
    request = &ModifyDBCustomNodeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomNodeType")
    
    
    return
}

func NewModifyDBCustomNodeTypeResponse() (response *ModifyDBCustomNodeTypeResponse) {
    response = &ModifyDBCustomNodeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomNodeType
// 该接口（ModifyDBCustomNodeType）用于修改 DB Custom 节点配置。
//
// 修改节点配置过程中会重启节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) ModifyDBCustomNodeType(request *ModifyDBCustomNodeTypeRequest) (response *ModifyDBCustomNodeTypeResponse, err error) {
    return c.ModifyDBCustomNodeTypeWithContext(context.Background(), request)
}

// ModifyDBCustomNodeType
// 该接口（ModifyDBCustomNodeType）用于修改 DB Custom 节点配置。
//
// 修改节点配置过程中会重启节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) ModifyDBCustomNodeTypeWithContext(ctx context.Context, request *ModifyDBCustomNodeTypeRequest) (response *ModifyDBCustomNodeTypeResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomNodeTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomNodeType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomNodeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomNodeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBCustomNodesDisasterRecoverGroupRequest() (request *ModifyDBCustomNodesDisasterRecoverGroupRequest) {
    request = &ModifyDBCustomNodesDisasterRecoverGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyDBCustomNodesDisasterRecoverGroup")
    
    
    return
}

func NewModifyDBCustomNodesDisasterRecoverGroupResponse() (response *ModifyDBCustomNodesDisasterRecoverGroupResponse) {
    response = &ModifyDBCustomNodesDisasterRecoverGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBCustomNodesDisasterRecoverGroup
// 该接口（ModifyDBCustomNodesDisasterRecoverGroup）用于修改 DB Custom 节点的置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodesDisasterRecoverGroup(request *ModifyDBCustomNodesDisasterRecoverGroupRequest) (response *ModifyDBCustomNodesDisasterRecoverGroupResponse, err error) {
    return c.ModifyDBCustomNodesDisasterRecoverGroupWithContext(context.Background(), request)
}

// ModifyDBCustomNodesDisasterRecoverGroup
// 该接口（ModifyDBCustomNodesDisasterRecoverGroup）用于修改 DB Custom 节点的置放群组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyDBCustomNodesDisasterRecoverGroupWithContext(ctx context.Context, request *ModifyDBCustomNodesDisasterRecoverGroupRequest) (response *ModifyDBCustomNodesDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBCustomNodesDisasterRecoverGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyDBCustomNodesDisasterRecoverGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBCustomNodesDisasterRecoverGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBCustomNodesDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// 本接口用于修改集群名称
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_MODIFYRESOURCEINFOERROR = "FailedOperation.ModifyResourceInfoError"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// 本接口用于修改集群名称
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_MODIFYRESOURCEINFOERROR = "FailedOperation.ModifyResourceInfoError"
//  FAILEDOPERATION_QUERYRESOURCEERROR = "FailedOperation.QueryResourceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ModifyInstanceName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewRebootDBCustomNodesRequest() (request *RebootDBCustomNodesRequest) {
    request = &RebootDBCustomNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "RebootDBCustomNodes")
    
    
    return
}

func NewRebootDBCustomNodesResponse() (response *RebootDBCustomNodesResponse) {
    response = &RebootDBCustomNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootDBCustomNodes
// 该接口（RebootDBCustomNodes）用于重启 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RebootDBCustomNodes(request *RebootDBCustomNodesRequest) (response *RebootDBCustomNodesResponse, err error) {
    return c.RebootDBCustomNodesWithContext(context.Background(), request)
}

// RebootDBCustomNodes
// 该接口（RebootDBCustomNodes）用于重启 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RebootDBCustomNodesWithContext(ctx context.Context, request *RebootDBCustomNodesRequest) (response *RebootDBCustomNodesResponse, err error) {
    if request == nil {
        request = NewRebootDBCustomNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "RebootDBCustomNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootDBCustomNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootDBCustomNodesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveNodesFromDBCustomClusterRequest() (request *RemoveNodesFromDBCustomClusterRequest) {
    request = &RemoveNodesFromDBCustomClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "RemoveNodesFromDBCustomCluster")
    
    
    return
}

func NewRemoveNodesFromDBCustomClusterResponse() (response *RemoveNodesFromDBCustomClusterResponse) {
    response = &RemoveNodesFromDBCustomClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveNodesFromDBCustomCluster
// 该接口（RemoveNodesFromDBCustomCluster）用于从 DB Custom 集群移出节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemoveNodesFromDBCustomCluster(request *RemoveNodesFromDBCustomClusterRequest) (response *RemoveNodesFromDBCustomClusterResponse, err error) {
    return c.RemoveNodesFromDBCustomClusterWithContext(context.Background(), request)
}

// RemoveNodesFromDBCustomCluster
// 该接口（RemoveNodesFromDBCustomCluster）用于从 DB Custom 集群移出节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RemoveNodesFromDBCustomClusterWithContext(ctx context.Context, request *RemoveNodesFromDBCustomClusterRequest) (response *RemoveNodesFromDBCustomClusterResponse, err error) {
    if request == nil {
        request = NewRemoveNodesFromDBCustomClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "RemoveNodesFromDBCustomCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveNodesFromDBCustomCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveNodesFromDBCustomClusterResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBCustomNodeRequest() (request *RenewDBCustomNodeRequest) {
    request = &RenewDBCustomNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "RenewDBCustomNode")
    
    
    return
}

func NewRenewDBCustomNodeResponse() (response *RenewDBCustomNodeResponse) {
    response = &RenewDBCustomNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewDBCustomNode
// 该接口（RenewDBCustomNode）用于给 DB Custom 节点续费，或者给已经隔离的实例解除隔离。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RenewDBCustomNode(request *RenewDBCustomNodeRequest) (response *RenewDBCustomNodeResponse, err error) {
    return c.RenewDBCustomNodeWithContext(context.Background(), request)
}

// RenewDBCustomNode
// 该接口（RenewDBCustomNode）用于给 DB Custom 节点续费，或者给已经隔离的实例解除隔离。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) RenewDBCustomNodeWithContext(ctx context.Context, request *RenewDBCustomNodeRequest) (response *RenewDBCustomNodeResponse, err error) {
    if request == nil {
        request = NewRenewDBCustomNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "RenewDBCustomNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDBCustomNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDBCustomNodeResponse()
    err = c.Send(request, response)
    return
}

func NewResetDBCustomNodeRequest() (request *ResetDBCustomNodeRequest) {
    request = &ResetDBCustomNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ResetDBCustomNode")
    
    
    return
}

func NewResetDBCustomNodeResponse() (response *ResetDBCustomNodeResponse) {
    response = &ResetDBCustomNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDBCustomNode
// 该接口（ResetDBCustomNode）用于重装 DB Custom 节点的操作系统。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetDBCustomNode(request *ResetDBCustomNodeRequest) (response *ResetDBCustomNodeResponse, err error) {
    return c.ResetDBCustomNodeWithContext(context.Background(), request)
}

// ResetDBCustomNode
// 该接口（ResetDBCustomNode）用于重装 DB Custom 节点的操作系统。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetDBCustomNodeWithContext(ctx context.Context, request *ResetDBCustomNodeRequest) (response *ResetDBCustomNodeResponse, err error) {
    if request == nil {
        request = NewResetDBCustomNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ResetDBCustomNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDBCustomNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDBCustomNodeResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDBCustomNodeDiskRequest() (request *ResizeDBCustomNodeDiskRequest) {
    request = &ResizeDBCustomNodeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "ResizeDBCustomNodeDisk")
    
    
    return
}

func NewResizeDBCustomNodeDiskResponse() (response *ResizeDBCustomNodeDiskResponse) {
    response = &ResizeDBCustomNodeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeDBCustomNodeDisk
// 该接口（ResizeDBCustomNodeDisk）用于调整 DB Custom 节点磁盘容量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) ResizeDBCustomNodeDisk(request *ResizeDBCustomNodeDiskRequest) (response *ResizeDBCustomNodeDiskResponse, err error) {
    return c.ResizeDBCustomNodeDiskWithContext(context.Background(), request)
}

// ResizeDBCustomNodeDisk
// 该接口（ResizeDBCustomNodeDisk）用于调整 DB Custom 节点磁盘容量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PARAMETERERROR = "InvalidParameterValue.ParameterError"
func (c *Client) ResizeDBCustomNodeDiskWithContext(ctx context.Context, request *ResizeDBCustomNodeDiskRequest) (response *ResizeDBCustomNodeDiskResponse, err error) {
    if request == nil {
        request = NewResizeDBCustomNodeDiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "ResizeDBCustomNodeDisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDBCustomNodeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDBCustomNodeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewStartDBCustomNodesRequest() (request *StartDBCustomNodesRequest) {
    request = &StartDBCustomNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "StartDBCustomNodes")
    
    
    return
}

func NewStartDBCustomNodesResponse() (response *StartDBCustomNodesResponse) {
    response = &StartDBCustomNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartDBCustomNodes
// 该接口（StartDBCustomNodes）用于开机 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StartDBCustomNodes(request *StartDBCustomNodesRequest) (response *StartDBCustomNodesResponse, err error) {
    return c.StartDBCustomNodesWithContext(context.Background(), request)
}

// StartDBCustomNodes
// 该接口（StartDBCustomNodes）用于开机 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StartDBCustomNodesWithContext(ctx context.Context, request *StartDBCustomNodesRequest) (response *StartDBCustomNodesResponse, err error) {
    if request == nil {
        request = NewStartDBCustomNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "StartDBCustomNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartDBCustomNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartDBCustomNodesResponse()
    err = c.Send(request, response)
    return
}

func NewStopDBCustomNodesRequest() (request *StopDBCustomNodesRequest) {
    request = &StopDBCustomNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbdc", APIVersion, "StopDBCustomNodes")
    
    
    return
}

func NewStopDBCustomNodesResponse() (response *StopDBCustomNodesResponse) {
    response = &StopDBCustomNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopDBCustomNodes
// 该接口（StopDBCustomNodes）用于关机 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopDBCustomNodes(request *StopDBCustomNodesRequest) (response *StopDBCustomNodesResponse, err error) {
    return c.StopDBCustomNodesWithContext(context.Background(), request)
}

// StopDBCustomNodes
// 该接口（StopDBCustomNodes）用于关机 DB Custom 节点。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) StopDBCustomNodesWithContext(ctx context.Context, request *StopDBCustomNodesRequest) (response *StopDBCustomNodesResponse, err error) {
    if request == nil {
        request = NewStopDBCustomNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbdc", APIVersion, "StopDBCustomNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDBCustomNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDBCustomNodesResponse()
    err = c.Send(request, response)
    return
}
