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

package v20230321

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-03-21"

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
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
func (c *Client) AddClusterStorageOption(request *AddClusterStorageOptionRequest) (response *AddClusterStorageOptionResponse, err error) {
    return c.AddClusterStorageOptionWithContext(context.Background(), request)
}

// AddClusterStorageOption
// 本接口（AddClusterStorageOption）用于添加集群存储选项信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
func (c *Client) AddClusterStorageOptionWithContext(ctx context.Context, request *AddClusterStorageOptionRequest) (response *AddClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewAddClusterStorageOptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "AddClusterStorageOption")
    
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
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
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
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) AddNodesWithContext(ctx context.Context, request *AddNodesRequest) (response *AddNodesResponse, err error) {
    if request == nil {
        request = NewAddNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "AddNodes")
    
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
//  INTERNALERROR_AGENTRUNSCRIPTFAIL = "InternalError.AgentRunScriptFail"
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
//  INTERNALERROR_AGENTRUNSCRIPTFAIL = "InternalError.AgentRunScriptFail"
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "AddQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddQueueResponse()
    err = c.Send(request, response)
    return
}

func NewAttachNodesRequest() (request *AttachNodesRequest) {
    request = &AttachNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "AttachNodes")
    
    
    return
}

func NewAttachNodesResponse() (response *AttachNodesResponse) {
    response = &AttachNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachNodes
// 本接口 (AttachNodes) 用于绑定一个或者多个计算节点指定资源到指定集群中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_AGENTRUNSCRIPTFAIL = "InternalError.AgentRunScriptFail"
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
func (c *Client) AttachNodes(request *AttachNodesRequest) (response *AttachNodesResponse, err error) {
    return c.AttachNodesWithContext(context.Background(), request)
}

// AttachNodes
// 本接口 (AttachNodes) 用于绑定一个或者多个计算节点指定资源到指定集群中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_AGENTRUNSCRIPTFAIL = "InternalError.AgentRunScriptFail"
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
func (c *Client) AttachNodesWithContext(ctx context.Context, request *AttachNodesRequest) (response *AttachNodesResponse, err error) {
    if request == nil {
        request = NewAttachNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "AttachNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachNodesResponse()
    err = c.Send(request, response)
    return
}

func NewBindClusterVpcRequest() (request *BindClusterVpcRequest) {
    request = &BindClusterVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "BindClusterVpc")
    
    
    return
}

func NewBindClusterVpcResponse() (response *BindClusterVpcResponse) {
    response = &BindClusterVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindClusterVpc
// 本接口 (BindClusterVpc) 用于为IDC集群绑定VPC和子网。
//
// 
//
// * 绑定VPC后，集群可在该VPC内开启专线/VPN代理。
//
// * VpcId和SubnetId为必填参数，且子网必须属于指定的VPC。
//
// * 若集群已开通代理，需先关闭代理（DisableClusterDedicatedProxy）再变更VPC绑定。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_VPCRESOURCE = "ResourceNotFound.VpcResource"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_HASMANAGEDNODES = "UnsupportedOperation.HasManagedNodes"
//  UNSUPPORTEDOPERATION_PROXYENABLED = "UnsupportedOperation.ProxyEnabled"
func (c *Client) BindClusterVpc(request *BindClusterVpcRequest) (response *BindClusterVpcResponse, err error) {
    return c.BindClusterVpcWithContext(context.Background(), request)
}

// BindClusterVpc
// 本接口 (BindClusterVpc) 用于为IDC集群绑定VPC和子网。
//
// 
//
// * 绑定VPC后，集群可在该VPC内开启专线/VPN代理。
//
// * VpcId和SubnetId为必填参数，且子网必须属于指定的VPC。
//
// * 若集群已开通代理，需先关闭代理（DisableClusterDedicatedProxy）再变更VPC绑定。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_VPCRESOURCE = "ResourceNotFound.VpcResource"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_HASMANAGEDNODES = "UnsupportedOperation.HasManagedNodes"
//  UNSUPPORTEDOPERATION_PROXYENABLED = "UnsupportedOperation.ProxyEnabled"
func (c *Client) BindClusterVpcWithContext(ctx context.Context, request *BindClusterVpcRequest) (response *BindClusterVpcResponse, err error) {
    if request == nil {
        request = NewBindClusterVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "BindClusterVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindClusterVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindClusterVpcResponse()
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
// 
//
// * 本接口为异步接口， 当创建集群请求下发成功后会返回一个集群`ID`和一个`RequestId`，此时创建集群操作并未立即完成。在此期间集群的状态将会处于“PENDING”或者“INITING”，集群创建结果可以通过调用 [DescribeClusters](https://cloud.tencent.com/document/product/1527/72100)  接口查询，如果集群状态(ClusterStatus)变为“RUNNING(运行中)”，则代表集群创建成功，“ INIT_FAILED”代表集群创建失败。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCAM = "InternalError.CallCAM"
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
// 
//
// * 本接口为异步接口， 当创建集群请求下发成功后会返回一个集群`ID`和一个`RequestId`，此时创建集群操作并未立即完成。在此期间集群的状态将会处于“PENDING”或者“INITING”，集群创建结果可以通过调用 [DescribeClusters](https://cloud.tencent.com/document/product/1527/72100)  接口查询，如果集群状态(ClusterStatus)变为“RUNNING(运行中)”，则代表集群创建成功，“ INIT_FAILED”代表集群创建失败。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCAM = "InternalError.CallCAM"
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "CreateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduledActionRequest() (request *CreateScheduledActionRequest) {
    request = &CreateScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "CreateScheduledAction")
    
    
    return
}

func NewCreateScheduledActionResponse() (response *CreateScheduledActionResponse) {
    response = &CreateScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScheduledAction
// 为指定集群队列创建定时伸缩任务，按计划时间自动调整队列的节点数量。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCAM = "InternalError.CallCAM"
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
func (c *Client) CreateScheduledAction(request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    return c.CreateScheduledActionWithContext(context.Background(), request)
}

// CreateScheduledAction
// 为指定集群队列创建定时伸缩任务，按计划时间自动调整队列的节点数量。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCAM = "InternalError.CallCAM"
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
func (c *Client) CreateScheduledActionWithContext(ctx context.Context, request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    if request == nil {
        request = NewCreateScheduledActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "CreateScheduledAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkspacesRequest() (request *CreateWorkspacesRequest) {
    request = &CreateWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "CreateWorkspaces")
    
    
    return
}

func NewCreateWorkspacesResponse() (response *CreateWorkspacesResponse) {
    response = &CreateWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkspaces
// 本接口 (CreateWorkspaces) 用于创建工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
func (c *Client) CreateWorkspaces(request *CreateWorkspacesRequest) (response *CreateWorkspacesResponse, err error) {
    return c.CreateWorkspacesWithContext(context.Background(), request)
}

// CreateWorkspaces
// 本接口 (CreateWorkspaces) 用于创建工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
func (c *Client) CreateWorkspacesWithContext(ctx context.Context, request *CreateWorkspacesRequest) (response *CreateWorkspacesResponse, err error) {
    if request == nil {
        request = NewCreateWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "CreateWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkspacesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteCluster")
    
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_LOCALPATH = "ResourceNotFound.LocalPath"
func (c *Client) DeleteClusterStorageOption(request *DeleteClusterStorageOptionRequest) (response *DeleteClusterStorageOptionResponse, err error) {
    return c.DeleteClusterStorageOptionWithContext(context.Background(), request)
}

// DeleteClusterStorageOption
// 本接口 (DeleteClusterStorageOption) 用于删除集群存储选项信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_LOCALPATH = "ResourceNotFound.LocalPath"
func (c *Client) DeleteClusterStorageOptionWithContext(ctx context.Context, request *DeleteClusterStorageOptionRequest) (response *DeleteClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewDeleteClusterStorageOptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteClusterStorageOption")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterStorageOption require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterStorageOptionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteJob")
    
    
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJob
// 本接口 (DeleteJob) 用于删除一个作业任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    return c.DeleteJobWithContext(context.Background(), request)
}

// DeleteJob
// 本接口 (DeleteJob) 用于删除一个作业任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteNodes")
    
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"
func (c *Client) DeleteQueueWithContext(ctx context.Context, request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
    if request == nil {
        request = NewDeleteQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduledActionRequest() (request *DeleteScheduledActionRequest) {
    request = &DeleteScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DeleteScheduledAction")
    
    
    return
}

func NewDeleteScheduledActionResponse() (response *DeleteScheduledActionResponse) {
    response = &DeleteScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScheduledAction
// 删除指定的定时伸缩任务。
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"
func (c *Client) DeleteScheduledAction(request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    return c.DeleteScheduledActionWithContext(context.Background(), request)
}

// DeleteScheduledAction
// 删除指定的定时伸缩任务。
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"
func (c *Client) DeleteScheduledActionWithContext(ctx context.Context, request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DeleteScheduledAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduledActionResponse()
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeAutoScalingConfigurationWithContext(ctx context.Context, request *DescribeAutoScalingConfigurationRequest) (response *DescribeAutoScalingConfigurationResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeAutoScalingConfiguration")
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterActivities(request *DescribeClusterActivitiesRequest) (response *DescribeClusterActivitiesResponse, err error) {
    return c.DescribeClusterActivitiesWithContext(context.Background(), request)
}

// DescribeClusterActivities
// 本接口（DescribeClusterActivities）用于查询集群活动历史记录列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterActivitiesWithContext(ctx context.Context, request *DescribeClusterActivitiesRequest) (response *DescribeClusterActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterActivitiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeClusterActivities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDedicatedProxyRequest() (request *DescribeClusterDedicatedProxyRequest) {
    request = &DescribeClusterDedicatedProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeClusterDedicatedProxy")
    
    
    return
}

func NewDescribeClusterDedicatedProxyResponse() (response *DescribeClusterDedicatedProxyResponse) {
    response = &DescribeClusterDedicatedProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDedicatedProxy
// 本接口 (DescribeClusterDedicatedProxy) 用于查询IDC集群专线/VPN代理的状态。
//
// 
//
// * 返回终端节点（EndPoint）的当前状态，包括是否就绪、VIP地址等信息。
//
// * 若代理未开通，EndPointReady返回false，EndPointStatus为UNKNOWN。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDedicatedProxy(request *DescribeClusterDedicatedProxyRequest) (response *DescribeClusterDedicatedProxyResponse, err error) {
    return c.DescribeClusterDedicatedProxyWithContext(context.Background(), request)
}

// DescribeClusterDedicatedProxy
// 本接口 (DescribeClusterDedicatedProxy) 用于查询IDC集群专线/VPN代理的状态。
//
// 
//
// * 返回终端节点（EndPoint）的当前状态，包括是否就绪、VIP地址等信息。
//
// * 若代理未开通，EndPointReady返回false，EndPointStatus为UNKNOWN。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDedicatedProxyWithContext(ctx context.Context, request *DescribeClusterDedicatedProxyRequest) (response *DescribeClusterDedicatedProxyResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDedicatedProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeClusterDedicatedProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDedicatedProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDedicatedProxyResponse()
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeClusterStorageOption(request *DescribeClusterStorageOptionRequest) (response *DescribeClusterStorageOptionResponse, err error) {
    return c.DescribeClusterStorageOptionWithContext(context.Background(), request)
}

// DescribeClusterStorageOption
// 本接口 (DescribeClusterStorageOption) 用于查询集群存储选项信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeClusterStorageOptionWithContext(ctx context.Context, request *DescribeClusterStorageOptionRequest) (response *DescribeClusterStorageOptionResponse, err error) {
    if request == nil {
        request = NewDescribeClusterStorageOptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeClusterStorageOption")
    
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInitNodeScriptsRequest() (request *DescribeInitNodeScriptsRequest) {
    request = &DescribeInitNodeScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeInitNodeScripts")
    
    
    return
}

func NewDescribeInitNodeScriptsResponse() (response *DescribeInitNodeScriptsResponse) {
    response = &DescribeInitNodeScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInitNodeScripts
// 本接口 (DescribeInitNodeScripts) 用于查询节点初始化脚本列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeInitNodeScripts(request *DescribeInitNodeScriptsRequest) (response *DescribeInitNodeScriptsResponse, err error) {
    return c.DescribeInitNodeScriptsWithContext(context.Background(), request)
}

// DescribeInitNodeScripts
// 本接口 (DescribeInitNodeScripts) 用于查询节点初始化脚本列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeInitNodeScriptsWithContext(ctx context.Context, request *DescribeInitNodeScriptsRequest) (response *DescribeInitNodeScriptsResponse, err error) {
    if request == nil {
        request = NewDescribeInitNodeScriptsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeInitNodeScripts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInitNodeScripts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInitNodeScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceFamiliesRequest() (request *DescribeInstanceFamiliesRequest) {
    request = &DescribeInstanceFamiliesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeInstanceFamilies")
    
    
    return
}

func NewDescribeInstanceFamiliesResponse() (response *DescribeInstanceFamiliesResponse) {
    response = &DescribeInstanceFamiliesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceFamilies
// 查询指定集群可用的机型族列表，用于弹性伸缩配置时选择机型族。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeInstanceFamilies(request *DescribeInstanceFamiliesRequest) (response *DescribeInstanceFamiliesResponse, err error) {
    return c.DescribeInstanceFamiliesWithContext(context.Background(), request)
}

// DescribeInstanceFamilies
// 查询指定集群可用的机型族列表，用于弹性伸缩配置时选择机型族。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) DescribeInstanceFamiliesWithContext(ctx context.Context, request *DescribeInstanceFamiliesRequest) (response *DescribeInstanceFamiliesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFamiliesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeInstanceFamilies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceFamilies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceFamiliesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobSubmitInfoRequest() (request *DescribeJobSubmitInfoRequest) {
    request = &DescribeJobSubmitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeJobSubmitInfo")
    
    
    return
}

func NewDescribeJobSubmitInfoResponse() (response *DescribeJobSubmitInfoResponse) {
    response = &DescribeJobSubmitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobSubmitInfo
// 本接口用于查询作业的提交信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_JOBSUBMITINFO = "ResourceNotFound.JobSubmitInfo"
func (c *Client) DescribeJobSubmitInfo(request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    return c.DescribeJobSubmitInfoWithContext(context.Background(), request)
}

// DescribeJobSubmitInfo
// 本接口用于查询作业的提交信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_JOBSUBMITINFO = "ResourceNotFound.JobSubmitInfo"
func (c *Client) DescribeJobSubmitInfoWithContext(ctx context.Context, request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJobSubmitInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeJobSubmitInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobSubmitInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobSubmitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// 本接口 (DescribeJobs) 用于查询作业任务列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBATCH = "InternalError.CallBatch"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// 本接口 (DescribeJobs) 用于查询作业任务列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBATCH = "InternalError.CallBatch"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsOverviewRequest() (request *DescribeJobsOverviewRequest) {
    request = &DescribeJobsOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeJobsOverview")
    
    
    return
}

func NewDescribeJobsOverviewResponse() (response *DescribeJobsOverviewResponse) {
    response = &DescribeJobsOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobsOverview
// 本接口 (DescribeJobs) 用于查询作业任务列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeJobsOverview(request *DescribeJobsOverviewRequest) (response *DescribeJobsOverviewResponse, err error) {
    return c.DescribeJobsOverviewWithContext(context.Background(), request)
}

// DescribeJobsOverview
// 本接口 (DescribeJobs) 用于查询作业任务列表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) DescribeJobsOverviewWithContext(ctx context.Context, request *DescribeJobsOverviewRequest) (response *DescribeJobsOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeJobsOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeJobsOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobsOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsOverviewResponse()
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQueueAutoScalingRequest() (request *DescribeQueueAutoScalingRequest) {
    request = &DescribeQueueAutoScalingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeQueueAutoScaling")
    
    
    return
}

func NewDescribeQueueAutoScalingResponse() (response *DescribeQueueAutoScalingResponse) {
    response = &DescribeQueueAutoScalingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQueueAutoScaling
// 查询指定集群的队列弹性伸缩配置信息。
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueueAutoScaling(request *DescribeQueueAutoScalingRequest) (response *DescribeQueueAutoScalingResponse, err error) {
    return c.DescribeQueueAutoScalingWithContext(context.Background(), request)
}

// DescribeQueueAutoScaling
// 查询指定集群的队列弹性伸缩配置信息。
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueueAutoScalingWithContext(ctx context.Context, request *DescribeQueueAutoScalingRequest) (response *DescribeQueueAutoScalingResponse, err error) {
    if request == nil {
        request = NewDescribeQueueAutoScalingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeQueueAutoScaling")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQueueAutoScaling require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQueueAutoScalingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQueueAutoScalingOverviewRequest() (request *DescribeQueueAutoScalingOverviewRequest) {
    request = &DescribeQueueAutoScalingOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeQueueAutoScalingOverview")
    
    
    return
}

func NewDescribeQueueAutoScalingOverviewResponse() (response *DescribeQueueAutoScalingOverviewResponse) {
    response = &DescribeQueueAutoScalingOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQueueAutoScalingOverview
// 查询指定集群的队列弹性伸缩概览信息，包括期望容量、当前容量、当前动态节点数、有效定时任务数等。
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueueAutoScalingOverview(request *DescribeQueueAutoScalingOverviewRequest) (response *DescribeQueueAutoScalingOverviewResponse, err error) {
    return c.DescribeQueueAutoScalingOverviewWithContext(context.Background(), request)
}

// DescribeQueueAutoScalingOverview
// 查询指定集群的队列弹性伸缩概览信息，包括期望容量、当前容量、当前动态节点数、有效定时任务数等。
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
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"
func (c *Client) DescribeQueueAutoScalingOverviewWithContext(ctx context.Context, request *DescribeQueueAutoScalingOverviewRequest) (response *DescribeQueueAutoScalingOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeQueueAutoScalingOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeQueueAutoScalingOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQueueAutoScalingOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQueueAutoScalingOverviewResponse()
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
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
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
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
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeQueues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduledActionsRequest() (request *DescribeScheduledActionsRequest) {
    request = &DescribeScheduledActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeScheduledActions")
    
    
    return
}

func NewDescribeScheduledActionsResponse() (response *DescribeScheduledActionsResponse) {
    response = &DescribeScheduledActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScheduledActions
// 查询指定集群队列的定时伸缩任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
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
func (c *Client) DescribeScheduledActions(request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    return c.DescribeScheduledActionsWithContext(context.Background(), request)
}

// DescribeScheduledActions
// 查询指定集群队列的定时伸缩任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
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
func (c *Client) DescribeScheduledActionsWithContext(ctx context.Context, request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeScheduledActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduledActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduledActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspacesRequest() (request *DescribeWorkspacesRequest) {
    request = &DescribeWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DescribeWorkspaces")
    
    
    return
}

func NewDescribeWorkspacesResponse() (response *DescribeWorkspacesResponse) {
    response = &DescribeWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkspaces
// 本接口（DescribeWorkspaces）用于查询工作空间列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDCVMINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidCvmInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
func (c *Client) DescribeWorkspaces(request *DescribeWorkspacesRequest) (response *DescribeWorkspacesResponse, err error) {
    return c.DescribeWorkspacesWithContext(context.Background(), request)
}

// DescribeWorkspaces
// 本接口（DescribeWorkspaces）用于查询工作空间列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDCVMINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidCvmInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
func (c *Client) DescribeWorkspacesWithContext(ctx context.Context, request *DescribeWorkspacesRequest) (response *DescribeWorkspacesResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DescribeWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspacesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachNodesRequest() (request *DetachNodesRequest) {
    request = &DetachNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DetachNodes")
    
    
    return
}

func NewDetachNodesResponse() (response *DetachNodesResponse) {
    response = &DetachNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachNodes
// 本接口 (DetachNodes) 用于将一个或者多个计算节点从集群中移除，但是不销毁指定计算资源。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDCVMINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidCvmInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
func (c *Client) DetachNodes(request *DetachNodesRequest) (response *DetachNodesResponse, err error) {
    return c.DetachNodesWithContext(context.Background(), request)
}

// DetachNodes
// 本接口 (DetachNodes) 用于将一个或者多个计算节点从集群中移除，但是不销毁指定计算资源。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDCVMINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidCvmInstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
func (c *Client) DetachNodesWithContext(ctx context.Context, request *DetachNodesRequest) (response *DetachNodesResponse, err error) {
    if request == nil {
        request = NewDetachNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DetachNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClusterDedicatedProxyRequest() (request *DisableClusterDedicatedProxyRequest) {
    request = &DisableClusterDedicatedProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "DisableClusterDedicatedProxy")
    
    
    return
}

func NewDisableClusterDedicatedProxyResponse() (response *DisableClusterDedicatedProxyResponse) {
    response = &DisableClusterDedicatedProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableClusterDedicatedProxy
// 本接口 (DisableClusterDedicatedProxy) 用于关闭IDC集群的专线/VPN代理。
//
// 
//
// * 关闭后，系统将删除VPC终端节点（EndPoint），断开IDC集群与云上VPC的网络连接。
//
// * 若代理未开通，调用将返回ProxyNotEnabled错误。
//
// * 操作不可逆，关闭后需重新调用EnableClusterDedicatedProxy开启。
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEENDPOINTFAILED = "InternalError.DeleteEndpointFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_HASMANAGEDNODES = "UnsupportedOperation.HasManagedNodes"
//  UNSUPPORTEDOPERATION_PROXYNOTENABLED = "UnsupportedOperation.ProxyNotEnabled"
func (c *Client) DisableClusterDedicatedProxy(request *DisableClusterDedicatedProxyRequest) (response *DisableClusterDedicatedProxyResponse, err error) {
    return c.DisableClusterDedicatedProxyWithContext(context.Background(), request)
}

// DisableClusterDedicatedProxy
// 本接口 (DisableClusterDedicatedProxy) 用于关闭IDC集群的专线/VPN代理。
//
// 
//
// * 关闭后，系统将删除VPC终端节点（EndPoint），断开IDC集群与云上VPC的网络连接。
//
// * 若代理未开通，调用将返回ProxyNotEnabled错误。
//
// * 操作不可逆，关闭后需重新调用EnableClusterDedicatedProxy开启。
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEENDPOINTFAILED = "InternalError.DeleteEndpointFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_HASMANAGEDNODES = "UnsupportedOperation.HasManagedNodes"
//  UNSUPPORTEDOPERATION_PROXYNOTENABLED = "UnsupportedOperation.ProxyNotEnabled"
func (c *Client) DisableClusterDedicatedProxyWithContext(ctx context.Context, request *DisableClusterDedicatedProxyRequest) (response *DisableClusterDedicatedProxyResponse, err error) {
    if request == nil {
        request = NewDisableClusterDedicatedProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "DisableClusterDedicatedProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClusterDedicatedProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClusterDedicatedProxyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClusterDedicatedProxyRequest() (request *EnableClusterDedicatedProxyRequest) {
    request = &EnableClusterDedicatedProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "EnableClusterDedicatedProxy")
    
    
    return
}

func NewEnableClusterDedicatedProxyResponse() (response *EnableClusterDedicatedProxyResponse) {
    response = &EnableClusterDedicatedProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableClusterDedicatedProxy
// 本接口 (EnableClusterDedicatedProxy) 用于开启IDC集群的专线/VPN代理。
//
// 
//
// * 开启后，系统将自动创建VPC终端节点（EndPoint），实现IDC集群与云上VPC的网络互通。
//
// * 若代理已开通，重复调用将幂等返回已有EndPoint信息。
//
// * SubnetId与VpcId需同时指定或同时不指定。若不指定，则使用集群已绑定的VPC和子网。
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEENDPOINTFAILED = "InternalError.CreateEndpointFailed"
//  INTERNALERROR_ENDPOINTSERVICEWHITELISTFAILED = "InternalError.EndpointServiceWhitelistFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_REGIONDEDICATEDPROXYDISABLED = "UnsupportedOperation.RegionDedicatedProxyDisabled"
//  UNSUPPORTEDOPERATION_REGIONNOTSUPPORTDEDICATEDPROXY = "UnsupportedOperation.RegionNotSupportDedicatedProxy"
//  UNSUPPORTEDOPERATION_VPCALREADYBOUND = "UnsupportedOperation.VpcAlreadyBound"
func (c *Client) EnableClusterDedicatedProxy(request *EnableClusterDedicatedProxyRequest) (response *EnableClusterDedicatedProxyResponse, err error) {
    return c.EnableClusterDedicatedProxyWithContext(context.Background(), request)
}

// EnableClusterDedicatedProxy
// 本接口 (EnableClusterDedicatedProxy) 用于开启IDC集群的专线/VPN代理。
//
// 
//
// * 开启后，系统将自动创建VPC终端节点（EndPoint），实现IDC集群与云上VPC的网络互通。
//
// * 若代理已开通，重复调用将幂等返回已有EndPoint信息。
//
// * SubnetId与VpcId需同时指定或同时不指定。若不指定，则使用集群已绑定的VPC和子网。
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEENDPOINTFAILED = "InternalError.CreateEndpointFailed"
//  INTERNALERROR_ENDPOINTSERVICEWHITELISTFAILED = "InternalError.EndpointServiceWhitelistFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_REGIONDEDICATEDPROXYDISABLED = "UnsupportedOperation.RegionDedicatedProxyDisabled"
//  UNSUPPORTEDOPERATION_REGIONNOTSUPPORTDEDICATEDPROXY = "UnsupportedOperation.RegionNotSupportDedicatedProxy"
//  UNSUPPORTEDOPERATION_VPCALREADYBOUND = "UnsupportedOperation.VpcAlreadyBound"
func (c *Client) EnableClusterDedicatedProxyWithContext(ctx context.Context, request *EnableClusterDedicatedProxyRequest) (response *EnableClusterDedicatedProxyResponse, err error) {
    if request == nil {
        request = NewEnableClusterDedicatedProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "EnableClusterDedicatedProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClusterDedicatedProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClusterDedicatedProxyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateRegisterCodeRequest() (request *GenerateRegisterCodeRequest) {
    request = &GenerateRegisterCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "GenerateRegisterCode")
    
    
    return
}

func NewGenerateRegisterCodeResponse() (response *GenerateRegisterCodeResponse) {
    response = &GenerateRegisterCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateRegisterCode
// 本接口(GenerateRegisterCode)用于为队列创建一个注册码，注册码用于IDC机器的注册纳管。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) GenerateRegisterCode(request *GenerateRegisterCodeRequest) (response *GenerateRegisterCodeResponse, err error) {
    return c.GenerateRegisterCodeWithContext(context.Background(), request)
}

// GenerateRegisterCode
// 本接口(GenerateRegisterCode)用于为队列创建一个注册码，注册码用于IDC机器的注册纳管。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) GenerateRegisterCodeWithContext(ctx context.Context, request *GenerateRegisterCodeRequest) (response *GenerateRegisterCodeResponse, err error) {
    if request == nil {
        request = NewGenerateRegisterCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "GenerateRegisterCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateRegisterCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateRegisterCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateRegisterCommandRequest() (request *GenerateRegisterCommandRequest) {
    request = &GenerateRegisterCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "GenerateRegisterCommand")
    
    
    return
}

func NewGenerateRegisterCommandResponse() (response *GenerateRegisterCommandResponse) {
    response = &GenerateRegisterCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateRegisterCommand
// 本接口 (GenerateRegisterCommand) 用于生成IDC集群的节点注册命令。
//
// 
//
// * 返回的注册命令可直接在IDC机器上以root身份执行，将该机器纳管进指定的IDC集群。
//
// * 当<code>Proxy=true</code>时，系统会先确保集群专线代理就绪（自动开启终端节点并轮询至ACTIVE），再签发注册码并渲染带代理VIP的注册命令；若在超时窗口内代理仍未就绪，将返回<code>FailedOperation.ProxyNotReady</code>。
//
// * 当<code>Proxy=false</code>时，IDC机器需可直连集群，直接签发注册码并渲染注册命令。
//
// * VpcId与SubnetId需同时指定或同时不指定；仅当<code>Proxy=true</code>且集群未绑定VPC时二者必填。当<code>Proxy=false</code>时二者不生效，若仍传入将返回<code>InvalidParameterValue.ParametersNotSupported</code>。
//
// * 若集群此前已开启专线代理并绑定了VPC/子网，本次传入的VpcId/SubnetId与已绑定值不一致时，将返回<code>UnsupportedOperation.VpcAlreadyBound</code>（不支持改绑）。
//
// * 仅支持IDC类型集群，对非IDC集群调用将返回<code>InvalidParameterValue.ParametersNotSupported</code>。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PROXYNOTREADY = "FailedOperation.ProxyNotReady"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_VPCALREADYBOUND = "UnsupportedOperation.VpcAlreadyBound"
func (c *Client) GenerateRegisterCommand(request *GenerateRegisterCommandRequest) (response *GenerateRegisterCommandResponse, err error) {
    return c.GenerateRegisterCommandWithContext(context.Background(), request)
}

// GenerateRegisterCommand
// 本接口 (GenerateRegisterCommand) 用于生成IDC集群的节点注册命令。
//
// 
//
// * 返回的注册命令可直接在IDC机器上以root身份执行，将该机器纳管进指定的IDC集群。
//
// * 当<code>Proxy=true</code>时，系统会先确保集群专线代理就绪（自动开启终端节点并轮询至ACTIVE），再签发注册码并渲染带代理VIP的注册命令；若在超时窗口内代理仍未就绪，将返回<code>FailedOperation.ProxyNotReady</code>。
//
// * 当<code>Proxy=false</code>时，IDC机器需可直连集群，直接签发注册码并渲染注册命令。
//
// * VpcId与SubnetId需同时指定或同时不指定；仅当<code>Proxy=true</code>且集群未绑定VPC时二者必填。当<code>Proxy=false</code>时二者不生效，若仍传入将返回<code>InvalidParameterValue.ParametersNotSupported</code>。
//
// * 若集群此前已开启专线代理并绑定了VPC/子网，本次传入的VpcId/SubnetId与已绑定值不一致时，将返回<code>UnsupportedOperation.VpcAlreadyBound</code>（不支持改绑）。
//
// * 仅支持IDC类型集群，对非IDC集群调用将返回<code>InvalidParameterValue.ParametersNotSupported</code>。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PROXYNOTREADY = "FailedOperation.ProxyNotReady"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  UNSUPPORTEDOPERATION_VPCALREADYBOUND = "UnsupportedOperation.VpcAlreadyBound"
func (c *Client) GenerateRegisterCommandWithContext(ctx context.Context, request *GenerateRegisterCommandRequest) (response *GenerateRegisterCommandResponse, err error) {
    if request == nil {
        request = NewGenerateRegisterCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "GenerateRegisterCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateRegisterCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateRegisterCommandResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateWorkspacesRequest() (request *InquirePriceCreateWorkspacesRequest) {
    request = &InquirePriceCreateWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "InquirePriceCreateWorkspaces")
    
    
    return
}

func NewInquirePriceCreateWorkspacesResponse() (response *InquirePriceCreateWorkspacesResponse) {
    response = &InquirePriceCreateWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateWorkspaces
// 本接口(InquirePriceCreateWorkspaces)用于创建实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) InquirePriceCreateWorkspaces(request *InquirePriceCreateWorkspacesRequest) (response *InquirePriceCreateWorkspacesResponse, err error) {
    return c.InquirePriceCreateWorkspacesWithContext(context.Background(), request)
}

// InquirePriceCreateWorkspaces
// 本接口(InquirePriceCreateWorkspaces)用于创建实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) InquirePriceCreateWorkspacesWithContext(ctx context.Context, request *InquirePriceCreateWorkspacesRequest) (response *InquirePriceCreateWorkspacesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "InquirePriceCreateWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateWorkspacesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyWorkspacesChargeTypeRequest() (request *InquirePriceModifyWorkspacesChargeTypeRequest) {
    request = &InquirePriceModifyWorkspacesChargeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "InquirePriceModifyWorkspacesChargeType")
    
    
    return
}

func NewInquirePriceModifyWorkspacesChargeTypeResponse() (response *InquirePriceModifyWorkspacesChargeTypeResponse) {
    response = &InquirePriceModifyWorkspacesChargeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceModifyWorkspacesChargeType
// 查询按量计费工作空间转换为包年包月的价格。不会创建订单或变更资源。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) InquirePriceModifyWorkspacesChargeType(request *InquirePriceModifyWorkspacesChargeTypeRequest) (response *InquirePriceModifyWorkspacesChargeTypeResponse, err error) {
    return c.InquirePriceModifyWorkspacesChargeTypeWithContext(context.Background(), request)
}

// InquirePriceModifyWorkspacesChargeType
// 查询按量计费工作空间转换为包年包月的价格。不会创建订单或变更资源。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) InquirePriceModifyWorkspacesChargeTypeWithContext(ctx context.Context, request *InquirePriceModifyWorkspacesChargeTypeRequest) (response *InquirePriceModifyWorkspacesChargeTypeResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyWorkspacesChargeTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "InquirePriceModifyWorkspacesChargeType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyWorkspacesChargeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyWorkspacesChargeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterDeletionProtectionRequest() (request *ModifyClusterDeletionProtectionRequest) {
    request = &ModifyClusterDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyClusterDeletionProtection")
    
    
    return
}

func NewModifyClusterDeletionProtectionResponse() (response *ModifyClusterDeletionProtectionResponse) {
    response = &ModifyClusterDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterDeletionProtection
// 修改集群删除保护状态
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) ModifyClusterDeletionProtection(request *ModifyClusterDeletionProtectionRequest) (response *ModifyClusterDeletionProtectionResponse, err error) {
    return c.ModifyClusterDeletionProtectionWithContext(context.Background(), request)
}

// ModifyClusterDeletionProtection
// 修改集群删除保护状态
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCLOUDSDKEXCEPTION = "InternalError.CallCloudSdkException"
//  INVALIDPARAMETER_COMBINATION = "InvalidParameter.Combination"
//  INVALIDPARAMETERVALUE_INVALIDUNDERWRITEPERIOD = "InvalidParameterValue.InvalidUnderwritePeriod"
//  INVALIDPARAMETERVALUE_SPACETYPENOTAVAILABLE = "InvalidParameterValue.SpaceTypeNotAvailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
func (c *Client) ModifyClusterDeletionProtectionWithContext(ctx context.Context, request *ModifyClusterDeletionProtectionRequest) (response *ModifyClusterDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewModifyClusterDeletionProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyClusterDeletionProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInitNodeScriptsRequest() (request *ModifyInitNodeScriptsRequest) {
    request = &ModifyInitNodeScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyInitNodeScripts")
    
    
    return
}

func NewModifyInitNodeScriptsResponse() (response *ModifyInitNodeScriptsResponse) {
    response = &ModifyInitNodeScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInitNodeScripts
// 本接口 (ModifyInitNodeScripts) 用于修改节点初始化脚本。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
func (c *Client) ModifyInitNodeScripts(request *ModifyInitNodeScriptsRequest) (response *ModifyInitNodeScriptsResponse, err error) {
    return c.ModifyInitNodeScriptsWithContext(context.Background(), request)
}

// ModifyInitNodeScripts
// 本接口 (ModifyInitNodeScripts) 用于修改节点初始化脚本。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
func (c *Client) ModifyInitNodeScriptsWithContext(ctx context.Context, request *ModifyInitNodeScriptsRequest) (response *ModifyInitNodeScriptsResponse, err error) {
    if request == nil {
        request = NewModifyInitNodeScriptsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyInitNodeScripts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInitNodeScripts require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInitNodeScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNodeAttributeRequest() (request *ModifyNodeAttributeRequest) {
    request = &ModifyNodeAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyNodeAttribute")
    
    
    return
}

func NewModifyNodeAttributeResponse() (response *ModifyNodeAttributeResponse) {
    response = &ModifyNodeAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNodeAttribute
// 本接口用于修改节点属性
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
func (c *Client) ModifyNodeAttribute(request *ModifyNodeAttributeRequest) (response *ModifyNodeAttributeResponse, err error) {
    return c.ModifyNodeAttributeWithContext(context.Background(), request)
}

// ModifyNodeAttribute
// 本接口用于修改节点属性
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
func (c *Client) ModifyNodeAttributeWithContext(ctx context.Context, request *ModifyNodeAttributeRequest) (response *ModifyNodeAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNodeAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyNodeAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodeAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodeAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduledActionRequest() (request *ModifyScheduledActionRequest) {
    request = &ModifyScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyScheduledAction")
    
    
    return
}

func NewModifyScheduledActionResponse() (response *ModifyScheduledActionResponse) {
    response = &ModifyScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyScheduledAction
// 修改指定的定时伸缩任务配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
func (c *Client) ModifyScheduledAction(request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    return c.ModifyScheduledActionWithContext(context.Background(), request)
}

// ModifyScheduledAction
// 修改指定的定时伸缩任务配置。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"
func (c *Client) ModifyScheduledActionWithContext(ctx context.Context, request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    if request == nil {
        request = NewModifyScheduledActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyScheduledAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyScheduledAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkspacesAttributeRequest() (request *ModifyWorkspacesAttributeRequest) {
    request = &ModifyWorkspacesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyWorkspacesAttribute")
    
    
    return
}

func NewModifyWorkspacesAttributeResponse() (response *ModifyWorkspacesAttributeResponse) {
    response = &ModifyWorkspacesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkspacesAttribute
// 本接口 (ModifyWorkspacesAttribute) 用于修改工作空间的属性（目前只支持修改工作空间的名称）。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) ModifyWorkspacesAttribute(request *ModifyWorkspacesAttributeRequest) (response *ModifyWorkspacesAttributeResponse, err error) {
    return c.ModifyWorkspacesAttributeWithContext(context.Background(), request)
}

// ModifyWorkspacesAttribute
// 本接口 (ModifyWorkspacesAttribute) 用于修改工作空间的属性（目前只支持修改工作空间的名称）。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) ModifyWorkspacesAttributeWithContext(ctx context.Context, request *ModifyWorkspacesAttributeRequest) (response *ModifyWorkspacesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyWorkspacesAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyWorkspacesAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkspacesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkspacesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkspacesChargeTypeRequest() (request *ModifyWorkspacesChargeTypeRequest) {
    request = &ModifyWorkspacesChargeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyWorkspacesChargeType")
    
    
    return
}

func NewModifyWorkspacesChargeTypeResponse() (response *ModifyWorkspacesChargeTypeResponse) {
    response = &ModifyWorkspacesChargeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkspacesChargeType
// 正式提交按量计费工作空间转包年包月订单。仅支持 ONLINE 且计费模式为 POSTPAID_BY_HOUR 的工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) ModifyWorkspacesChargeType(request *ModifyWorkspacesChargeTypeRequest) (response *ModifyWorkspacesChargeTypeResponse, err error) {
    return c.ModifyWorkspacesChargeTypeWithContext(context.Background(), request)
}

// ModifyWorkspacesChargeType
// 正式提交按量计费工作空间转包年包月订单。仅支持 ONLINE 且计费模式为 POSTPAID_BY_HOUR 的工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
func (c *Client) ModifyWorkspacesChargeTypeWithContext(ctx context.Context, request *ModifyWorkspacesChargeTypeRequest) (response *ModifyWorkspacesChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyWorkspacesChargeTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyWorkspacesChargeType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkspacesChargeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkspacesChargeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkspacesRenewFlagRequest() (request *ModifyWorkspacesRenewFlagRequest) {
    request = &ModifyWorkspacesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "ModifyWorkspacesRenewFlag")
    
    
    return
}

func NewModifyWorkspacesRenewFlagResponse() (response *ModifyWorkspacesRenewFlagResponse) {
    response = &ModifyWorkspacesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkspacesRenewFlag
// 本接口 (ModifyWorkspacesAttribute) 用于修改工作空间的属性（目前只支持修改工作空间的名称）。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
//  INVALIDPARAMETERVALUE_SPACEIDNOTFOUND = "InvalidParameterValue.SpaceIdNotFound"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_SPACECHARGETYPE = "UnsupportedOperation.SpaceChargeType"
//  UNSUPPORTEDOPERATION_WORKSPACESTATEARREARS = "UnsupportedOperation.WorkspaceStateArrears"
func (c *Client) ModifyWorkspacesRenewFlag(request *ModifyWorkspacesRenewFlagRequest) (response *ModifyWorkspacesRenewFlagResponse, err error) {
    return c.ModifyWorkspacesRenewFlagWithContext(context.Background(), request)
}

// ModifyWorkspacesRenewFlag
// 本接口 (ModifyWorkspacesAttribute) 用于修改工作空间的属性（目前只支持修改工作空间的名称）。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDSPACEIDMALFORMED = "InvalidParameterValue.InvalidSpaceIdMalformed"
//  INVALIDPARAMETERVALUE_SPACEIDNOTFOUND = "InvalidParameterValue.SpaceIdNotFound"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  UNSUPPORTEDOPERATION_SPACECHARGETYPE = "UnsupportedOperation.SpaceChargeType"
//  UNSUPPORTEDOPERATION_WORKSPACESTATEARREARS = "UnsupportedOperation.WorkspaceStateArrears"
func (c *Client) ModifyWorkspacesRenewFlagWithContext(ctx context.Context, request *ModifyWorkspacesRenewFlagRequest) (response *ModifyWorkspacesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyWorkspacesRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "ModifyWorkspacesRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkspacesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkspacesRenewFlagResponse()
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
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
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
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) SetAutoScalingConfigurationWithContext(ctx context.Context, request *SetAutoScalingConfigurationRequest) (response *SetAutoScalingConfigurationResponse, err error) {
    if request == nil {
        request = NewSetAutoScalingConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "SetAutoScalingConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAutoScalingConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAutoScalingConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewSetQueueAutoScalingRequest() (request *SetQueueAutoScalingRequest) {
    request = &SetQueueAutoScalingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "SetQueueAutoScaling")
    
    
    return
}

func NewSetQueueAutoScalingResponse() (response *SetQueueAutoScalingResponse) {
    response = &SetQueueAutoScalingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetQueueAutoScaling
// 为指定集群的队列配置弹性伸缩策略，包括伸缩容量、扩容方式等。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) SetQueueAutoScaling(request *SetQueueAutoScalingRequest) (response *SetQueueAutoScalingResponse, err error) {
    return c.SetQueueAutoScalingWithContext(context.Background(), request)
}

// SetQueueAutoScaling
// 为指定集群的队列配置弹性伸缩策略，包括伸缩容量、扩容方式等。
//
// 可能返回的错误码:
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"
//  RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"
//  UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"
//  UNSUPPORTEDOPERATION_CLUSTERACCEPTOTHERREQUEST = "UnsupportedOperation.ClusterAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
//  UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"
//  UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
func (c *Client) SetQueueAutoScalingWithContext(ctx context.Context, request *SetQueueAutoScalingRequest) (response *SetQueueAutoScalingResponse, err error) {
    if request == nil {
        request = NewSetQueueAutoScalingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "SetQueueAutoScaling")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetQueueAutoScaling require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetQueueAutoScalingResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitJobRequest() (request *SubmitJobRequest) {
    request = &SubmitJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "SubmitJob")
    
    
    return
}

func NewSubmitJobResponse() (response *SubmitJobResponse) {
    response = &SubmitJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitJob
// 本接口 (SubmitJob) 用于提交一个作业任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBDESCRIPTIONTOOLONG = "InvalidParameterValue.JobDescriptionTooLong"
//  INVALIDPARAMETERVALUE_JOBNAMETOOLONG = "InvalidParameterValue.JobNameTooLong"
//  INVALIDPARAMETERVALUE_TASKDEPENDENCIESUNFEASIBLE = "InvalidParameterValue.TaskDependenciesUnfeasible"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED_TASKINSTANCENUMLIMIT = "LimitExceeded.TaskInstanceNumLimit"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) SubmitJob(request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
    return c.SubmitJobWithContext(context.Background(), request)
}

// SubmitJob
// 本接口 (SubmitJob) 用于提交一个作业任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBDESCRIPTIONTOOLONG = "InvalidParameterValue.JobDescriptionTooLong"
//  INVALIDPARAMETERVALUE_JOBNAMETOOLONG = "InvalidParameterValue.JobNameTooLong"
//  INVALIDPARAMETERVALUE_TASKDEPENDENCIESUNFEASIBLE = "InvalidParameterValue.TaskDependenciesUnfeasible"
//  INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"
//  LIMITEXCEEDED_TASKINSTANCENUMLIMIT = "LimitExceeded.TaskInstanceNumLimit"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"
func (c *Client) SubmitJobWithContext(ctx context.Context, request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
    if request == nil {
        request = NewSubmitJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "SubmitJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
    request = &TerminateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "TerminateJob")
    
    
    return
}

func NewTerminateJobResponse() (response *TerminateJobResponse) {
    response = &TerminateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateJob
// 本接口 (TerminateJob) 用于终止一个作业任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) TerminateJob(request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    return c.TerminateJobWithContext(context.Background(), request)
}

// TerminateJob
// 本接口 (TerminateJob) 用于终止一个作业任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) TerminateJobWithContext(ctx context.Context, request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    if request == nil {
        request = NewTerminateJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "TerminateJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateWorkspacesRequest() (request *TerminateWorkspacesRequest) {
    request = &TerminateWorkspacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("thpc", APIVersion, "TerminateWorkspaces")
    
    
    return
}

func NewTerminateWorkspacesResponse() (response *TerminateWorkspacesResponse) {
    response = &TerminateWorkspacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateWorkspaces
// 本接口 (TerminateWorkspaces) 用于主动退还工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) TerminateWorkspaces(request *TerminateWorkspacesRequest) (response *TerminateWorkspacesResponse, err error) {
    return c.TerminateWorkspacesWithContext(context.Background(), request)
}

// TerminateWorkspaces
// 本接口 (TerminateWorkspaces) 用于主动退还工作空间。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"
//  INVALIDPARAMETERVALUE_JOBNOTFOUND = "InvalidParameterValue.JobNotFound"
//  UNSUPPORTEDOPERATION_JOBSTATENOTSUPPORT = "UnsupportedOperation.JobStateNotSupport"
func (c *Client) TerminateWorkspacesWithContext(ctx context.Context, request *TerminateWorkspacesRequest) (response *TerminateWorkspacesResponse, err error) {
    if request == nil {
        request = NewTerminateWorkspacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "thpc", APIVersion, "TerminateWorkspaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateWorkspaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateWorkspacesResponse()
    err = c.Send(request, response)
    return
}
