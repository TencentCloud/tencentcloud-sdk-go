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

package v20220401

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-04-01"

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


func NewAddNodesRequest() (request *AddNodesRequest) {
    request = &AddNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "AddNodes")
    
    
    return
}

func NewAddNodesResponse() (response *AddNodesResponse) {
    response = &AddNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddNodes
// 本接口(AddNodes)用于添加一个或者多个计算节点或者登录节点到指定集群。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) AddNodes(request *AddNodesRequest) (response *AddNodesResponse, err error) {
    return c.AddNodesWithContext(context.Background(), request)
}

// AddNodes
// 本接口(AddNodes)用于添加一个或者多个计算节点或者登录节点到指定集群。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) AddNodesWithContext(ctx context.Context, request *AddNodesRequest) (response *AddNodesResponse, err error) {
    if request == nil {
        request = NewAddNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNodesResponse()
    err = c.Send(request, response)
    return
}

func NewBindAutoScalingGroupRequest() (request *BindAutoScalingGroupRequest) {
    request = &BindAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "BindAutoScalingGroup")
    
    
    return
}

func NewBindAutoScalingGroupResponse() (response *BindAutoScalingGroupResponse) {
    response = &BindAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAutoScalingGroup
// 本接口(BindAutoScalingGroup)用于为集群队列绑定弹性伸缩组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONID = "ResourceNotFound.LaunchConfigurationId"
//  UNSUPPORTEDOPERATION_AUTOSCALINGGROUPALREADYBINDED = "UnsupportedOperation.AutoScalingGroupAlreadyBinded"
func (c *Client) BindAutoScalingGroup(request *BindAutoScalingGroupRequest) (response *BindAutoScalingGroupResponse, err error) {
    return c.BindAutoScalingGroupWithContext(context.Background(), request)
}

// BindAutoScalingGroup
// 本接口(BindAutoScalingGroup)用于为集群队列绑定弹性伸缩组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONID = "ResourceNotFound.LaunchConfigurationId"
//  UNSUPPORTEDOPERATION_AUTOSCALINGGROUPALREADYBINDED = "UnsupportedOperation.AutoScalingGroupAlreadyBinded"
func (c *Client) BindAutoScalingGroupWithContext(ctx context.Context, request *BindAutoScalingGroupRequest) (response *BindAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewBindAutoScalingGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoScalingGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCluster
// 本接口 (CreateCluster) 用于创建并启动集群。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 本接口 (CreateCluster) 用于创建并启动集群。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCluster
// 本接口（DeleteCluster）用于删除一个指定的集群。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// 本接口（DeleteCluster）用于删除一个指定的集群。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNodesRequest() (request *DeleteNodesRequest) {
    request = &DeleteNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteNodes")
    
    
    return
}

func NewDeleteNodesResponse() (response *DeleteNodesResponse) {
    response = &DeleteNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNodes
// 本接口(DeleteNodes)用于删除指定集群中一个或者多个计算节点或者登录节点。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteNodes(request *DeleteNodesRequest) (response *DeleteNodesResponse, err error) {
    return c.DeleteNodesWithContext(context.Background(), request)
}

// DeleteNodes
// 本接口(DeleteNodes)用于删除指定集群中一个或者多个计算节点或者登录节点。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteNodesWithContext(ctx context.Context, request *DeleteNodesRequest) (response *DeleteNodesResponse, err error) {
    if request == nil {
        request = NewDeleteNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusters
// 本接口（DescribeClusters）用于查询集群列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 本接口（DescribeClusters）用于查询集群列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}
