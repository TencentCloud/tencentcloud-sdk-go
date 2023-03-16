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


func NewAddClusterStorageOptionRequest() (request *AddClusterStorageOptionRequest) {
    request = &AddClusterStorageOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "AddClusterStorageOption")
    
    
    return
}

func NewAddClusterStorageOptionResponse() (response *AddClusterStorageOptionResponse) {
    response = &AddClusterStorageOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddClusterStorageOption
// 本接口（AddClusterStorageOption）用于添加集群存储选项信息。
func (c *Client) AddClusterStorageOption(request *AddClusterStorageOptionRequest) (response *AddClusterStorageOptionResponse, err error) {
    return c.AddClusterStorageOptionWithContext(context.Background(), request)
}

// AddClusterStorageOption
// 本接口（AddClusterStorageOption）用于添加集群存储选项信息。
func (c *Client) AddClusterStorageOptionWithContext(ctx context.Context, request *AddClusterStorageOptionRequest) (response *AddClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewAddClusterStorageOptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClusterStorageOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClusterStorageOptionResponse()
    err = c.Send(request, response)
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
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
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
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
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

func NewAddQueueRequest() (request *AddQueueRequest) {
    request = &AddQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "AddQueue")
    
    
    return
}

func NewAddQueueResponse() (response *AddQueueResponse) {
    response = &AddQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddQueue
// 本接口(AddQueue)用于添加队列到指定集群。
//
// * 本接口为目前只支持SchedulerType为SLURM的集群。
//
// * 单个集群中队列数量上限为10个。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUEUENUMLIMIT = "LimitExceeded.QueueNumLimit"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) AddQueue(request *AddQueueRequest) (response *AddQueueResponse, err error) {
    return c.AddQueueWithContext(context.Background(), request)
}

// AddQueue
// 本接口(AddQueue)用于添加队列到指定集群。
//
// * 本接口为目前只支持SchedulerType为SLURM的集群。
//
// * 单个集群中队列数量上限为10个。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUEUENUMLIMIT = "LimitExceeded.QueueNumLimit"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) AddQueueWithContext(ctx context.Context, request *AddQueueRequest) (response *AddQueueResponse, err error) {
    if request == nil {
        request = NewAddQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddQueueResponse()
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPID = "ResourceNotFound.AutoScalingGroupId"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONID = "ResourceNotFound.LaunchConfigurationId"
//  UNSUPPORTEDOPERATION_AUTOSCALINGGROUPALREADYBINDED = "UnsupportedOperation.AutoScalingGroupAlreadyBinded"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) BindAutoScalingGroup(request *BindAutoScalingGroupRequest) (response *BindAutoScalingGroupResponse, err error) {
    return c.BindAutoScalingGroupWithContext(context.Background(), request)
}

// BindAutoScalingGroup
// 本接口(BindAutoScalingGroup)用于为集群队列绑定弹性伸缩组
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_AUTOSCALINGGROUPID = "ResourceNotFound.AutoScalingGroupId"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_LAUNCHCONFIGURATIONID = "ResourceNotFound.LaunchConfigurationId"
//  UNSUPPORTEDOPERATION_AUTOSCALINGGROUPALREADYBINDED = "UnsupportedOperation.AutoScalingGroupAlreadyBinded"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 本接口 (CreateCluster) 用于创建并启动集群。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
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

func NewDeleteClusterStorageOptionRequest() (request *DeleteClusterStorageOptionRequest) {
    request = &DeleteClusterStorageOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteClusterStorageOption")
    
    
    return
}

func NewDeleteClusterStorageOptionResponse() (response *DeleteClusterStorageOptionResponse) {
    response = &DeleteClusterStorageOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterStorageOption
// 本接口 (DeleteClusterStorageOption) 用于删除集群存储选项信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteClusterStorageOption(request *DeleteClusterStorageOptionRequest) (response *DeleteClusterStorageOptionResponse, err error) {
    return c.DeleteClusterStorageOptionWithContext(context.Background(), request)
}

// DeleteClusterStorageOption
// 本接口 (DeleteClusterStorageOption) 用于删除集群存储选项信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DeleteClusterStorageOptionWithContext(ctx context.Context, request *DeleteClusterStorageOptionRequest) (response *DeleteClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewDeleteClusterStorageOptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterStorageOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterStorageOptionResponse()
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_NODEID = "ResourceNotFound.NodeId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_INVALIDNODEROLE = "UnsupportedOperation.InvalidNodeRole"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
func (c *Client) DeleteNodes(request *DeleteNodesRequest) (response *DeleteNodesResponse, err error) {
    return c.DeleteNodesWithContext(context.Background(), request)
}

// DeleteNodes
// 本接口(DeleteNodes)用于删除指定集群中一个或者多个计算节点或者登录节点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_NODEID = "ResourceNotFound.NodeId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_INVALIDNODEROLE = "UnsupportedOperation.InvalidNodeRole"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
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

func NewDeleteQueueRequest() (request *DeleteQueueRequest) {
    request = &DeleteQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteQueue")
    
    
    return
}

func NewDeleteQueueResponse() (response *DeleteQueueResponse) {
    response = &DeleteQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteQueue
// 本接口(DeleteQueue)用于从指定集群删除队列。
//
// * 本接口为目前只支持SchedulerType为SLURM的集群。
//
// 
//
// * 删除队列时，需要保证队列内不存在节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"
func (c *Client) DeleteQueue(request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    return c.DeleteQueueWithContext(context.Background(), request)
}

// DeleteQueue
// 本接口(DeleteQueue)用于从指定集群删除队列。
//
// * 本接口为目前只支持SchedulerType为SLURM的集群。
//
// 
//
// * 删除队列时，需要保证队列内不存在节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"
func (c *Client) DeleteQueueWithContext(ctx context.Context, request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    if request == nil {
        request = NewDeleteQueueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingConfigurationRequest() (request *DescribeAutoScalingConfigurationRequest) {
    request = &DescribeAutoScalingConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeAutoScalingConfiguration")
    
    
    return
}

func NewDescribeAutoScalingConfigurationResponse() (response *DescribeAutoScalingConfigurationResponse) {
    response = &DescribeAutoScalingConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoScalingConfiguration
// 本接口(DescribeAutoScalingConfiguration)用于查询集群弹性伸缩配置信息。本接口仅适用于弹性伸缩类型为THPC_AS的集群。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeAutoScalingConfiguration(request *DescribeAutoScalingConfigurationRequest) (response *DescribeAutoScalingConfigurationResponse, err error) {
    return c.DescribeAutoScalingConfigurationWithContext(context.Background(), request)
}

// DescribeAutoScalingConfiguration
// 本接口(DescribeAutoScalingConfiguration)用于查询集群弹性伸缩配置信息。本接口仅适用于弹性伸缩类型为THPC_AS的集群。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeAutoScalingConfigurationWithContext(ctx context.Context, request *DescribeAutoScalingConfigurationRequest) (response *DescribeAutoScalingConfigurationResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoScalingConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoScalingConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterActivitiesRequest() (request *DescribeClusterActivitiesRequest) {
    request = &DescribeClusterActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeClusterActivities")
    
    
    return
}

func NewDescribeClusterActivitiesResponse() (response *DescribeClusterActivitiesResponse) {
    response = &DescribeClusterActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterActivities
// 本接口（DescribeClusterActivities）用于查询集群活动历史记录列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeClusterActivities(request *DescribeClusterActivitiesRequest) (response *DescribeClusterActivitiesResponse, err error) {
    return c.DescribeClusterActivitiesWithContext(context.Background(), request)
}

// DescribeClusterActivities
// 本接口（DescribeClusterActivities）用于查询集群活动历史记录列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeClusterActivitiesWithContext(ctx context.Context, request *DescribeClusterActivitiesRequest) (response *DescribeClusterActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterStorageOptionRequest() (request *DescribeClusterStorageOptionRequest) {
    request = &DescribeClusterStorageOptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeClusterStorageOption")
    
    
    return
}

func NewDescribeClusterStorageOptionResponse() (response *DescribeClusterStorageOptionResponse) {
    response = &DescribeClusterStorageOptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterStorageOption
// 本接口 (DescribeClusterStorageOption) 用于查询集群存储选项信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeClusterStorageOption(request *DescribeClusterStorageOptionRequest) (response *DescribeClusterStorageOptionResponse, err error) {
    return c.DescribeClusterStorageOptionWithContext(context.Background(), request)
}

// DescribeClusterStorageOption
// 本接口 (DescribeClusterStorageOption) 用于查询集群存储选项信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeClusterStorageOptionWithContext(ctx context.Context, request *DescribeClusterStorageOptionRequest) (response *DescribeClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewDescribeClusterStorageOptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterStorageOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterStorageOptionResponse()
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

func NewDescribeNodesRequest() (request *DescribeNodesRequest) {
    request = &DescribeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeNodes")
    
    
    return
}

func NewDescribeNodesResponse() (response *DescribeNodesResponse) {
    response = &DescribeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNodes
// 本接口 (DescribeNodes) 用于查询指定集群节点概览信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeNodes(request *DescribeNodesRequest) (response *DescribeNodesResponse, err error) {
    return c.DescribeNodesWithContext(context.Background(), request)
}

// DescribeNodes
// 本接口 (DescribeNodes) 用于查询指定集群节点概览信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeNodesWithContext(ctx context.Context, request *DescribeNodesRequest) (response *DescribeNodesResponse, err error) {
    if request == nil {
        request = NewDescribeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQueuesRequest() (request *DescribeQueuesRequest) {
    request = &DescribeQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeQueues")
    
    
    return
}

func NewDescribeQueuesResponse() (response *DescribeQueuesResponse) {
    response = &DescribeQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQueues
// 本接口(DescribeQueues)用于查询指定集群队列概览信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueues(request *DescribeQueuesRequest) (response *DescribeQueuesResponse, err error) {
    return c.DescribeQueuesWithContext(context.Background(), request)
}

// DescribeQueues
// 本接口(DescribeQueues)用于查询指定集群队列概览信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueuesWithContext(ctx context.Context, request *DescribeQueuesRequest) (response *DescribeQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeQueuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewSetAutoScalingConfigurationRequest() (request *SetAutoScalingConfigurationRequest) {
    request = &SetAutoScalingConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "SetAutoScalingConfiguration")
    
    
    return
}

func NewSetAutoScalingConfigurationResponse() (response *SetAutoScalingConfigurationResponse) {
    response = &SetAutoScalingConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetAutoScalingConfiguration
// 本接口(SetAutoScalingConfiguration)用于为集群设置集群弹性伸缩配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) SetAutoScalingConfiguration(request *SetAutoScalingConfigurationRequest) (response *SetAutoScalingConfigurationResponse, err error) {
    return c.SetAutoScalingConfigurationWithContext(context.Background(), request)
}

// SetAutoScalingConfiguration
// 本接口(SetAutoScalingConfiguration)用于为集群设置集群弹性伸缩配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) SetAutoScalingConfigurationWithContext(ctx context.Context, request *SetAutoScalingConfigurationRequest) (response *SetAutoScalingConfigurationResponse, err error) {
    if request == nil {
        request = NewSetAutoScalingConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAutoScalingConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAutoScalingConfigurationResponse()
    err = c.Send(request, response)
    return
}
