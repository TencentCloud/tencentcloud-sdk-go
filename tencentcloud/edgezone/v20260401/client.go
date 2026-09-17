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

package v20260401

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-04-01"

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


func NewApplyPublicIpsRequest() (request *ApplyPublicIpsRequest) {
    request = &ApplyPublicIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ApplyPublicIps")
    
    
    return
}

func NewApplyPublicIpsResponse() (response *ApplyPublicIpsResponse) {
    response = &ApplyPublicIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyPublicIps
// 从静态 IP 池为指定公网实例批量申请多个 Ip 地址（随机分配）。申请前需检查用户配额。
//
// 此接口仅适用于 `RouteMode=static` 的公网实例。BGP/OSPF 实例调用此接口将返回错误。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IPV6QUOTAEXCEEDED = "LimitExceeded.Ipv6QuotaExceeded"
//  LIMITEXCEEDED_IPV6QUOTANOTCONFIGURED = "LimitExceeded.Ipv6QuotaNotConfigured"
//  LIMITEXCEEDED_QUOTAEXCEEDED = "LimitExceeded.QuotaExceeded"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_PUBLICIPINSUFFICIENT = "ResourceInsufficient.PublicIpInsufficient"
//  RESOURCEINSUFFICIENT_PUBLICIPV6INSUFFICIENT = "ResourceInsufficient.PublicIpv6Insufficient"
func (c *Client) ApplyPublicIps(request *ApplyPublicIpsRequest) (response *ApplyPublicIpsResponse, err error) {
    return c.ApplyPublicIpsWithContext(context.Background(), request)
}

// ApplyPublicIps
// 从静态 IP 池为指定公网实例批量申请多个 Ip 地址（随机分配）。申请前需检查用户配额。
//
// 此接口仅适用于 `RouteMode=static` 的公网实例。BGP/OSPF 实例调用此接口将返回错误。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IPV6QUOTAEXCEEDED = "LimitExceeded.Ipv6QuotaExceeded"
//  LIMITEXCEEDED_IPV6QUOTANOTCONFIGURED = "LimitExceeded.Ipv6QuotaNotConfigured"
//  LIMITEXCEEDED_QUOTAEXCEEDED = "LimitExceeded.QuotaExceeded"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_PUBLICIPINSUFFICIENT = "ResourceInsufficient.PublicIpInsufficient"
//  RESOURCEINSUFFICIENT_PUBLICIPV6INSUFFICIENT = "ResourceInsufficient.PublicIpv6Insufficient"
func (c *Client) ApplyPublicIpsWithContext(ctx context.Context, request *ApplyPublicIpsRequest) (response *ApplyPublicIpsResponse, err error) {
    if request == nil {
        request = NewApplyPublicIpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ApplyPublicIps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyPublicIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyPublicIpsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeNodeServiceRequest() (request *CreateEdgeNodeServiceRequest) {
    request = &CreateEdgeNodeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreateEdgeNodeService")
    
    
    return
}

func NewCreateEdgeNodeServiceResponse() (response *CreateEdgeNodeServiceResponse) {
    response = &CreateEdgeNodeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeNodeService
// 开通边缘节点计费服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ZONE = "ResourceNotFound.Zone"
func (c *Client) CreateEdgeNodeService(request *CreateEdgeNodeServiceRequest) (response *CreateEdgeNodeServiceResponse, err error) {
    return c.CreateEdgeNodeServiceWithContext(context.Background(), request)
}

// CreateEdgeNodeService
// 开通边缘节点计费服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ZONE = "ResourceNotFound.Zone"
func (c *Client) CreateEdgeNodeServiceWithContext(ctx context.Context, request *CreateEdgeNodeServiceRequest) (response *CreateEdgeNodeServiceResponse, err error) {
    if request == nil {
        request = NewCreateEdgeNodeServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreateEdgeNodeService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeNodeService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeNodeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstances
// 创建物理机实例，系统自动分配物理机资源并完成装机。如果用户未在当前可用区开通计费，系统自动开通。支持并发分配物理机资源，异步执行网络分配和装机任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 创建物理机实例，系统自动分配物理机资源并完成装机。如果用户未在当前可用区开通计费，系统自动开通。支持并发分配物理机资源，异步执行网络分配和装机任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateNetworkInstanceRequest() (request *CreatePrivateNetworkInstanceRequest) {
    request = &CreatePrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreatePrivateNetworkInstance")
    
    
    return
}

func NewCreatePrivateNetworkInstanceResponse() (response *CreatePrivateNetworkInstanceResponse) {
    response = &CreatePrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivateNetworkInstance
// 创建私网实例，一个用户在一个可用区仅支持创建一个私网实例，网络地址由 Network（网络号）和 Mask（掩码位数）两个参数共同决定子网范围。Network 必须是三个 RFC 1918 私有地址段之一的合法网络地址：10.0.0.0/8、172.16.0.0/12 或 192.168.0.0/16，且 host 位必须全为 0（即Network 与 Mask 组合后不能有主机位被置位，例如 10.0.0.1/24 是非法的，应填 10.0.0.0/24）。Mask 的上限统一为 28，下限由所属地址段决定：10.x.x.x 段允许 8～28，172.16.x.x 段允许 12～28，192.168.x.x 段允许 16～28。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRIVATEINSTANCEDUPLICATE = "FailedOperation.PrivateInstanceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDNETWORK = "InvalidParameterValue.InvalidNetwork"
func (c *Client) CreatePrivateNetworkInstance(request *CreatePrivateNetworkInstanceRequest) (response *CreatePrivateNetworkInstanceResponse, err error) {
    return c.CreatePrivateNetworkInstanceWithContext(context.Background(), request)
}

// CreatePrivateNetworkInstance
// 创建私网实例，一个用户在一个可用区仅支持创建一个私网实例，网络地址由 Network（网络号）和 Mask（掩码位数）两个参数共同决定子网范围。Network 必须是三个 RFC 1918 私有地址段之一的合法网络地址：10.0.0.0/8、172.16.0.0/12 或 192.168.0.0/16，且 host 位必须全为 0（即Network 与 Mask 组合后不能有主机位被置位，例如 10.0.0.1/24 是非法的，应填 10.0.0.0/24）。Mask 的上限统一为 28，下限由所属地址段决定：10.x.x.x 段允许 8～28，172.16.x.x 段允许 12～28，192.168.x.x 段允许 16～28。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRIVATEINSTANCEDUPLICATE = "FailedOperation.PrivateInstanceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDNETWORK = "InvalidParameterValue.InvalidNetwork"
func (c *Client) CreatePrivateNetworkInstanceWithContext(ctx context.Context, request *CreatePrivateNetworkInstanceRequest) (response *CreatePrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreatePrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicNetworkInstanceRequest() (request *CreatePublicNetworkInstanceRequest) {
    request = &CreatePublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreatePublicNetworkInstance")
    
    
    return
}

func NewCreatePublicNetworkInstanceResponse() (response *CreatePublicNetworkInstanceResponse) {
    response = &CreatePublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePublicNetworkInstance
// 用户输入可用区ID、公网实例名称、网络线路、路由模式以创建公网实例，一个用户在一个可用区仅支持创建一个公网实例
//
// 路由模式为 **静态** 的公网实例需要用户主动申请和释放公网IP
//
// 路由模式为 **OSPF、BGP** 的公网实例在创建时自动分配公网IP段，销毁时自动释放公网IP段
//
// 可能返回的错误码:
//  FAILEDOPERATION_PUBLICINSTANCEDUPLICATE = "FailedOperation.PublicInstanceDuplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUOTA = "InvalidParameter.InvalidQuota"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_NOAVAILABLECIDR = "ResourceInsufficient.NoAvailableCidr"
//  RESOURCENOTFOUND_ZONENOTFOUND = "ResourceNotFound.ZoneNotFound"
func (c *Client) CreatePublicNetworkInstance(request *CreatePublicNetworkInstanceRequest) (response *CreatePublicNetworkInstanceResponse, err error) {
    return c.CreatePublicNetworkInstanceWithContext(context.Background(), request)
}

// CreatePublicNetworkInstance
// 用户输入可用区ID、公网实例名称、网络线路、路由模式以创建公网实例，一个用户在一个可用区仅支持创建一个公网实例
//
// 路由模式为 **静态** 的公网实例需要用户主动申请和释放公网IP
//
// 路由模式为 **OSPF、BGP** 的公网实例在创建时自动分配公网IP段，销毁时自动释放公网IP段
//
// 可能返回的错误码:
//  FAILEDOPERATION_PUBLICINSTANCEDUPLICATE = "FailedOperation.PublicInstanceDuplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUOTA = "InvalidParameter.InvalidQuota"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_NOAVAILABLECIDR = "ResourceInsufficient.NoAvailableCidr"
//  RESOURCENOTFOUND_ZONENOTFOUND = "ResourceNotFound.ZoneNotFound"
func (c *Client) CreatePublicNetworkInstanceWithContext(ctx context.Context, request *CreatePublicNetworkInstanceRequest) (response *CreatePublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreatePublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateNetworkInstanceRequest() (request *DeletePrivateNetworkInstanceRequest) {
    request = &DeletePrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DeletePrivateNetworkInstance")
    
    
    return
}

func NewDeletePrivateNetworkInstanceResponse() (response *DeletePrivateNetworkInstanceResponse) {
    response = &DeletePrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrivateNetworkInstance
// 删除私网实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PRIVATEINSTANCEINUSE = "ResourceInUse.PrivateInstanceInUse"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) DeletePrivateNetworkInstance(request *DeletePrivateNetworkInstanceRequest) (response *DeletePrivateNetworkInstanceResponse, err error) {
    return c.DeletePrivateNetworkInstanceWithContext(context.Background(), request)
}

// DeletePrivateNetworkInstance
// 删除私网实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PRIVATEINSTANCEINUSE = "ResourceInUse.PrivateInstanceInUse"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) DeletePrivateNetworkInstanceWithContext(ctx context.Context, request *DeletePrivateNetworkInstanceRequest) (response *DeletePrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewDeletePrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DeletePrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicNetworkInstanceRequest() (request *DeletePublicNetworkInstanceRequest) {
    request = &DeletePublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DeletePublicNetworkInstance")
    
    
    return
}

func NewDeletePublicNetworkInstanceResponse() (response *DeletePublicNetworkInstanceResponse) {
    response = &DeletePublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePublicNetworkInstance
// 修改公网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DeletePublicNetworkInstance(request *DeletePublicNetworkInstanceRequest) (response *DeletePublicNetworkInstanceResponse, err error) {
    return c.DeletePublicNetworkInstanceWithContext(context.Background(), request)
}

// DeletePublicNetworkInstance
// 修改公网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DeletePublicNetworkInstanceWithContext(ctx context.Context, request *DeletePublicNetworkInstanceRequest) (response *DeletePublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewDeletePublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DeletePublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
    request = &DescribeInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeInstanceTypes")
    
    
    return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
    response = &DescribeInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceTypes
// 根据 AppId 查询账号下可用区维度的机型配额列表；若传入 Zone，则仅返回指定可用区下的机型配额；若不传，则返回账号下所有可用区的机型配额。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    return c.DescribeInstanceTypesWithContext(context.Background(), request)
}

// DescribeInstanceTypes
// 根据 AppId 查询账号下可用区维度的机型配额列表；若传入 Zone，则仅返回指定可用区下的机型配额；若不传，则返回账号下所有可用区的机型配额。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstanceTypesWithContext(ctx context.Context, request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeInstanceTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 查询物理机实例列表，支持按实例ID、实例名称、可用区、实例状态等条件筛选，并支持分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询物理机实例列表，支持按实例ID、实例名称、可用区、实例状态等条件筛选，并支持分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateNetworkInstancesRequest() (request *DescribePrivateNetworkInstancesRequest) {
    request = &DescribePrivateNetworkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePrivateNetworkInstances")
    
    
    return
}

func NewDescribePrivateNetworkInstancesResponse() (response *DescribePrivateNetworkInstancesResponse) {
    response = &DescribePrivateNetworkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateNetworkInstances
// 查询私网实例，支持通过私网实例ID、私网实例名称、可用区ID等参数进行查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrivateNetworkInstances(request *DescribePrivateNetworkInstancesRequest) (response *DescribePrivateNetworkInstancesResponse, err error) {
    return c.DescribePrivateNetworkInstancesWithContext(context.Background(), request)
}

// DescribePrivateNetworkInstances
// 查询私网实例，支持通过私网实例ID、私网实例名称、可用区ID等参数进行查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrivateNetworkInstancesWithContext(ctx context.Context, request *DescribePrivateNetworkInstancesRequest) (response *DescribePrivateNetworkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrivateNetworkInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePrivateNetworkInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateNetworkInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateNetworkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicIpsRequest() (request *DescribePublicIpsRequest) {
    request = &DescribePublicIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePublicIps")
    
    
    return
}

func NewDescribePublicIpsResponse() (response *DescribePublicIpsResponse) {
    response = &DescribePublicIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicIps
// 查询用户的公网Ip信息，对于路由模式为Static的公网实例，会返回所有已申请的公网Ip信息，对于路由模式为Ospf和Bgp的公网实例，会直接返回网段信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIps(request *DescribePublicIpsRequest) (response *DescribePublicIpsResponse, err error) {
    return c.DescribePublicIpsWithContext(context.Background(), request)
}

// DescribePublicIps
// 查询用户的公网Ip信息，对于路由模式为Static的公网实例，会返回所有已申请的公网Ip信息，对于路由模式为Ospf和Bgp的公网实例，会直接返回网段信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIpsWithContext(ctx context.Context, request *DescribePublicIpsRequest) (response *DescribePublicIpsResponse, err error) {
    if request == nil {
        request = NewDescribePublicIpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePublicIps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicNetworkInstancesRequest() (request *DescribePublicNetworkInstancesRequest) {
    request = &DescribePublicNetworkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePublicNetworkInstances")
    
    
    return
}

func NewDescribePublicNetworkInstancesResponse() (response *DescribePublicNetworkInstancesResponse) {
    response = &DescribePublicNetworkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicNetworkInstances
// 查询公网实例列表，支持按实例ID、实例名称、可用区等条件筛选，并支持分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribePublicNetworkInstances(request *DescribePublicNetworkInstancesRequest) (response *DescribePublicNetworkInstancesResponse, err error) {
    return c.DescribePublicNetworkInstancesWithContext(context.Background(), request)
}

// DescribePublicNetworkInstances
// 查询公网实例列表，支持按实例ID、实例名称、可用区等条件筛选，并支持分页查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribePublicNetworkInstancesWithContext(ctx context.Context, request *DescribePublicNetworkInstancesRequest) (response *DescribePublicNetworkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePublicNetworkInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePublicNetworkInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicNetworkInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicNetworkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneDataRequest() (request *DescribeZoneDataRequest) {
    request = &DescribeZoneDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeZoneData")
    
    
    return
}

func NewDescribeZoneDataResponse() (response *DescribeZoneDataResponse) {
    response = &DescribeZoneDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneData
// 按指标名，查询统计数据。数据按1分钟间隔统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZoneData(request *DescribeZoneDataRequest) (response *DescribeZoneDataResponse, err error) {
    return c.DescribeZoneDataWithContext(context.Background(), request)
}

// DescribeZoneData
// 按指标名，查询统计数据。数据按1分钟间隔统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZoneDataWithContext(ctx context.Context, request *DescribeZoneDataRequest) (response *DescribeZoneDataResponse, err error) {
    if request == nil {
        request = NewDescribeZoneDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeZoneData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// 跨地域聚合查询所有已配置 region 下的可用区列表。支持通过 FilterByAppId 参数控制是否按账号过滤：默认仅返回账号关联的可用区，设为 False 时返回所有可用区。本地域直查数据库，远程地域并发 HTTP 请求后合并返回。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 跨地域聚合查询所有已配置 region 下的可用区列表。支持通过 FilterByAppId 参数控制是否按账号过滤：默认仅返回账号关联的可用区，设为 False 时返回所有可用区。本地域直查数据库，远程地域并发 HTTP 请求后合并返回。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
    request = &ModifyInstanceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyInstanceAttribute")
    
    
    return
}

func NewModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
    response = &ModifyInstanceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAttribute
// 修改物理机实例的属性，支持修改实例名称、变更公网IP（IPv4/IPv6）。InstanceName 和 NewPublicIp 至少传入一个。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
    return c.ModifyInstanceAttributeWithContext(context.Background(), request)
}

// ModifyInstanceAttribute
// 修改物理机实例的属性，支持修改实例名称、变更公网IP（IPv4/IPv6）。InstanceName 和 NewPublicIp 至少传入一个。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyInstanceAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateNetworkInstanceRequest() (request *ModifyPrivateNetworkInstanceRequest) {
    request = &ModifyPrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyPrivateNetworkInstance")
    
    
    return
}

func NewModifyPrivateNetworkInstanceResponse() (response *ModifyPrivateNetworkInstanceResponse) {
    response = &ModifyPrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateNetworkInstance
// 修改私网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) ModifyPrivateNetworkInstance(request *ModifyPrivateNetworkInstanceRequest) (response *ModifyPrivateNetworkInstanceResponse, err error) {
    return c.ModifyPrivateNetworkInstanceWithContext(context.Background(), request)
}

// ModifyPrivateNetworkInstance
// 修改私网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) ModifyPrivateNetworkInstanceWithContext(ctx context.Context, request *ModifyPrivateNetworkInstanceRequest) (response *ModifyPrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewModifyPrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyPrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublicNetworkInstanceRequest() (request *ModifyPublicNetworkInstanceRequest) {
    request = &ModifyPublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyPublicNetworkInstance")
    
    
    return
}

func NewModifyPublicNetworkInstanceResponse() (response *ModifyPublicNetworkInstanceResponse) {
    response = &ModifyPublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublicNetworkInstance
// 修改公网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) ModifyPublicNetworkInstance(request *ModifyPublicNetworkInstanceRequest) (response *ModifyPublicNetworkInstanceResponse, err error) {
    return c.ModifyPublicNetworkInstanceWithContext(context.Background(), request)
}

// ModifyPublicNetworkInstance
// 修改公网实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) ModifyPublicNetworkInstanceWithContext(ctx context.Context, request *ModifyPublicNetworkInstanceRequest) (response *ModifyPublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewModifyPublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyPublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicIpRequest() (request *ReleasePublicIpRequest) {
    request = &ReleasePublicIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ReleasePublicIp")
    
    
    return
}

func NewReleasePublicIpResponse() (response *ReleasePublicIpResponse) {
    response = &ReleasePublicIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleasePublicIp
// 批量释放已分配给 STATIC 公网实例但**未绑定物理服务器**的 IPv4 地址
//
// 此接口仅适用于 STATIC 模式实例。BGP/OSPF 实例的 CIDR 在实例删除时自动归还，无需手动释放单个 IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IPSTILLBOUNDTOSERVER = "FailedOperation.IpStillBoundToServer"
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PUBLICIPNOTAVAILABLE = "InvalidParameterValue.PublicIpNotAvailable"
//  INVALIDPARAMETERVALUE_PUBLICIPV6NOTAVAILABLE = "InvalidParameterValue.PublicIpv6NotAvailable"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReleasePublicIp(request *ReleasePublicIpRequest) (response *ReleasePublicIpResponse, err error) {
    return c.ReleasePublicIpWithContext(context.Background(), request)
}

// ReleasePublicIp
// 批量释放已分配给 STATIC 公网实例但**未绑定物理服务器**的 IPv4 地址
//
// 此接口仅适用于 STATIC 模式实例。BGP/OSPF 实例的 CIDR 在实例删除时自动归还，无需手动释放单个 IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IPSTILLBOUNDTOSERVER = "FailedOperation.IpStillBoundToServer"
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PUBLICIPNOTAVAILABLE = "InvalidParameterValue.PublicIpNotAvailable"
//  INVALIDPARAMETERVALUE_PUBLICIPV6NOTAVAILABLE = "InvalidParameterValue.PublicIpv6NotAvailable"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReleasePublicIpWithContext(ctx context.Context, request *ReleasePublicIpRequest) (response *ReleasePublicIpResponse, err error) {
    if request == nil {
        request = NewReleasePublicIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ReleasePublicIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleasePublicIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleasePublicIpResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// 销毁物理机实例，释放资源。接口同步释放网络资源（IP回收）并更新状态为 terminating，后台异步执行磁盘清理。支持部分成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LEGACYCOMPATBATCHMIXED = "UnsupportedOperation.LegacyCompatBatchMixed"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// 销毁物理机实例，释放资源。接口同步释放网络资源（IP回收）并更新状态为 terminating，后台异步执行磁盘清理。支持部分成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LEGACYCOMPATBATCHMIXED = "UnsupportedOperation.LegacyCompatBatchMixed"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "TerminateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
