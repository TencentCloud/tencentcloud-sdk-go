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

package v20240906

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-09-06"

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


func NewAssociateTargetGroupsRequest() (request *AssociateTargetGroupsRequest) {
    request = &AssociateTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "AssociateTargetGroups")
    
    
    return
}

func NewAssociateTargetGroupsResponse() (response *AssociateTargetGroupsResponse) {
    response = &AssociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateTargetGroups
// 本接口(AssociateTargetGroups)用来将目标组绑定到负载均衡。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateTargetGroups(request *AssociateTargetGroupsRequest) (response *AssociateTargetGroupsResponse, err error) {
    return c.AssociateTargetGroupsWithContext(context.Background(), request)
}

// AssociateTargetGroups
// 本接口(AssociateTargetGroups)用来将目标组绑定到负载均衡。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateTargetGroupsWithContext(ctx context.Context, request *AssociateTargetGroupsRequest) (response *AssociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayLoadBalancerRequest() (request *CreateGatewayLoadBalancerRequest) {
    request = &CreateGatewayLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "CreateGatewayLoadBalancer")
    
    
    return
}

func NewCreateGatewayLoadBalancerResponse() (response *CreateGatewayLoadBalancerResponse) {
    response = &CreateGatewayLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGatewayLoadBalancer
// 本接口(CreateGatewayLoadBalancer)用来创建网关负载均衡实例。为了使用网关负载均衡服务，您必须购买一个或多个网关负载均衡实例。成功调用该接口后，会返回网关负载均衡实例的唯一 ID。
//
// 注意：单个账号在每个地域的默认购买配额为：10个。
//
// 本接口为异步接口，接口成功返回后，可使用 DescribeGatewayLoadBalancers 接口查询负载均衡实例的状态（如创建中、正常），以确定是否创建成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) CreateGatewayLoadBalancer(request *CreateGatewayLoadBalancerRequest) (response *CreateGatewayLoadBalancerResponse, err error) {
    return c.CreateGatewayLoadBalancerWithContext(context.Background(), request)
}

// CreateGatewayLoadBalancer
// 本接口(CreateGatewayLoadBalancer)用来创建网关负载均衡实例。为了使用网关负载均衡服务，您必须购买一个或多个网关负载均衡实例。成功调用该接口后，会返回网关负载均衡实例的唯一 ID。
//
// 注意：单个账号在每个地域的默认购买配额为：10个。
//
// 本接口为异步接口，接口成功返回后，可使用 DescribeGatewayLoadBalancers 接口查询负载均衡实例的状态（如创建中、正常），以确定是否创建成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) CreateGatewayLoadBalancerWithContext(ctx context.Context, request *CreateGatewayLoadBalancerRequest) (response *CreateGatewayLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateGatewayLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGatewayLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGatewayLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetGroupRequest() (request *CreateTargetGroupRequest) {
    request = &CreateTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "CreateTargetGroup")
    
    
    return
}

func NewCreateTargetGroupResponse() (response *CreateTargetGroupResponse) {
    response = &CreateTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTargetGroup
// 创建目标组。该功能正在内测中，如需使用，请通过[工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20LB&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    return c.CreateTargetGroupWithContext(context.Background(), request)
}

// CreateTargetGroup
// 创建目标组。该功能正在内测中，如需使用，请通过[工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20LB&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTargetGroupWithContext(ctx context.Context, request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatewayLoadBalancerRequest() (request *DeleteGatewayLoadBalancerRequest) {
    request = &DeleteGatewayLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DeleteGatewayLoadBalancer")
    
    
    return
}

func NewDeleteGatewayLoadBalancerResponse() (response *DeleteGatewayLoadBalancerResponse) {
    response = &DeleteGatewayLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGatewayLoadBalancer
// DeleteGatewayLoadBalancer 接口用以删除指定的一个或多个网关负载均衡实例。成功删除后，会把网关负载均衡实例与后端服务解绑。
//
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DeleteGatewayLoadBalancer(request *DeleteGatewayLoadBalancerRequest) (response *DeleteGatewayLoadBalancerResponse, err error) {
    return c.DeleteGatewayLoadBalancerWithContext(context.Background(), request)
}

// DeleteGatewayLoadBalancer
// DeleteGatewayLoadBalancer 接口用以删除指定的一个或多个网关负载均衡实例。成功删除后，会把网关负载均衡实例与后端服务解绑。
//
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DeleteGatewayLoadBalancerWithContext(ctx context.Context, request *DeleteGatewayLoadBalancerRequest) (response *DeleteGatewayLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteGatewayLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGatewayLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGatewayLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetGroupsRequest() (request *DeleteTargetGroupsRequest) {
    request = &DeleteTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DeleteTargetGroups")
    
    
    return
}

func NewDeleteTargetGroupsResponse() (response *DeleteTargetGroupsResponse) {
    response = &DeleteTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTargetGroups
// 删除目标组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTargetGroups(request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    return c.DeleteTargetGroupsWithContext(context.Background(), request)
}

// DeleteTargetGroups
// 删除目标组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTargetGroupsWithContext(ctx context.Context, request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetGroupInstancesRequest() (request *DeregisterTargetGroupInstancesRequest) {
    request = &DeregisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DeregisterTargetGroupInstances")
    
    
    return
}

func NewDeregisterTargetGroupInstancesResponse() (response *DeregisterTargetGroupInstancesResponse) {
    response = &DeregisterTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeregisterTargetGroupInstances
// 从目标组中解绑服务器。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetGroupInstances(request *DeregisterTargetGroupInstancesRequest) (response *DeregisterTargetGroupInstancesResponse, err error) {
    return c.DeregisterTargetGroupInstancesWithContext(context.Background(), request)
}

// DeregisterTargetGroupInstances
// 从目标组中解绑服务器。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetGroupInstancesWithContext(ctx context.Context, request *DeregisterTargetGroupInstancesRequest) (response *DeregisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayLoadBalancersRequest() (request *DescribeGatewayLoadBalancersRequest) {
    request = &DescribeGatewayLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeGatewayLoadBalancers")
    
    
    return
}

func NewDescribeGatewayLoadBalancersResponse() (response *DescribeGatewayLoadBalancersResponse) {
    response = &DescribeGatewayLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayLoadBalancers
// 查询一个地域的网关负载均衡实例列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) DescribeGatewayLoadBalancers(request *DescribeGatewayLoadBalancersRequest) (response *DescribeGatewayLoadBalancersResponse, err error) {
    return c.DescribeGatewayLoadBalancersWithContext(context.Background(), request)
}

// DescribeGatewayLoadBalancers
// 查询一个地域的网关负载均衡实例列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) DescribeGatewayLoadBalancersWithContext(ctx context.Context, request *DescribeGatewayLoadBalancersRequest) (response *DescribeGatewayLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupInstanceStatusRequest() (request *DescribeTargetGroupInstanceStatusRequest) {
    request = &DescribeTargetGroupInstanceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeTargetGroupInstanceStatus")
    
    
    return
}

func NewDescribeTargetGroupInstanceStatusResponse() (response *DescribeTargetGroupInstanceStatusResponse) {
    response = &DescribeTargetGroupInstanceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupInstanceStatus
// 查询目标组后端服务状态。目前仅支持网关负载均衡类型的目标组支持查询后端服务状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) DescribeTargetGroupInstanceStatus(request *DescribeTargetGroupInstanceStatusRequest) (response *DescribeTargetGroupInstanceStatusResponse, err error) {
    return c.DescribeTargetGroupInstanceStatusWithContext(context.Background(), request)
}

// DescribeTargetGroupInstanceStatus
// 查询目标组后端服务状态。目前仅支持网关负载均衡类型的目标组支持查询后端服务状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) DescribeTargetGroupInstanceStatusWithContext(ctx context.Context, request *DescribeTargetGroupInstanceStatusRequest) (response *DescribeTargetGroupInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupInstanceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupInstanceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupInstanceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupInstancesRequest() (request *DescribeTargetGroupInstancesRequest) {
    request = &DescribeTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeTargetGroupInstances")
    
    
    return
}

func NewDescribeTargetGroupInstancesResponse() (response *DescribeTargetGroupInstancesResponse) {
    response = &DescribeTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupInstances
// 获取目标组绑定的服务器信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupInstances(request *DescribeTargetGroupInstancesRequest) (response *DescribeTargetGroupInstancesResponse, err error) {
    return c.DescribeTargetGroupInstancesWithContext(context.Background(), request)
}

// DescribeTargetGroupInstances
// 获取目标组绑定的服务器信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupInstancesWithContext(ctx context.Context, request *DescribeTargetGroupInstancesRequest) (response *DescribeTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupListRequest() (request *DescribeTargetGroupListRequest) {
    request = &DescribeTargetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeTargetGroupList")
    
    
    return
}

func NewDescribeTargetGroupListResponse() (response *DescribeTargetGroupListResponse) {
    response = &DescribeTargetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupList
// 获取目标组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupList(request *DescribeTargetGroupListRequest) (response *DescribeTargetGroupListResponse, err error) {
    return c.DescribeTargetGroupListWithContext(context.Background(), request)
}

// DescribeTargetGroupList
// 获取目标组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupListWithContext(ctx context.Context, request *DescribeTargetGroupListRequest) (response *DescribeTargetGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
    request = &DescribeTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeTargetGroups")
    
    
    return
}

func NewDescribeTargetGroupsResponse() (response *DescribeTargetGroupsResponse) {
    response = &DescribeTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroups
// 查询目标组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroups(request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    return c.DescribeTargetGroupsWithContext(context.Background(), request)
}

// DescribeTargetGroups
// 查询目标组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupsWithContext(ctx context.Context, request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskStatus
// 本接口用于查询异步任务的执行状态，对于非查询类的接口（创建/删除负载均衡实例等），在接口调用成功后，都需要使用本接口查询任务最终是否执行成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// 本接口用于查询异步任务的执行状态，对于非查询类的接口（创建/删除负载均衡实例等），在接口调用成功后，都需要使用本接口查询任务最终是否执行成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateTargetGroupsRequest() (request *DisassociateTargetGroupsRequest) {
    request = &DisassociateTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "DisassociateTargetGroups")
    
    
    return
}

func NewDisassociateTargetGroupsResponse() (response *DisassociateTargetGroupsResponse) {
    response = &DisassociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateTargetGroups
// 解除负载均衡和目标组的关联关系。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateTargetGroups(request *DisassociateTargetGroupsRequest) (response *DisassociateTargetGroupsResponse, err error) {
    return c.DisassociateTargetGroupsWithContext(context.Background(), request)
}

// DisassociateTargetGroups
// 解除负载均衡和目标组的关联关系。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateTargetGroupsWithContext(ctx context.Context, request *DisassociateTargetGroupsRequest) (response *DisassociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateGatewayLoadBalancerRequest() (request *InquirePriceCreateGatewayLoadBalancerRequest) {
    request = &InquirePriceCreateGatewayLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "InquirePriceCreateGatewayLoadBalancer")
    
    
    return
}

func NewInquirePriceCreateGatewayLoadBalancerResponse() (response *InquirePriceCreateGatewayLoadBalancerResponse) {
    response = &InquirePriceCreateGatewayLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateGatewayLoadBalancer
// InquirePriceCreateGatewayLoadBalancer接口查询创建网关负载均衡的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquirePriceCreateGatewayLoadBalancer(request *InquirePriceCreateGatewayLoadBalancerRequest) (response *InquirePriceCreateGatewayLoadBalancerResponse, err error) {
    return c.InquirePriceCreateGatewayLoadBalancerWithContext(context.Background(), request)
}

// InquirePriceCreateGatewayLoadBalancer
// InquirePriceCreateGatewayLoadBalancer接口查询创建网关负载均衡的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquirePriceCreateGatewayLoadBalancerWithContext(ctx context.Context, request *InquirePriceCreateGatewayLoadBalancerRequest) (response *InquirePriceCreateGatewayLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateGatewayLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateGatewayLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateGatewayLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGatewayLoadBalancerAttributeRequest() (request *ModifyGatewayLoadBalancerAttributeRequest) {
    request = &ModifyGatewayLoadBalancerAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "ModifyGatewayLoadBalancerAttribute")
    
    
    return
}

func NewModifyGatewayLoadBalancerAttributeResponse() (response *ModifyGatewayLoadBalancerAttributeResponse) {
    response = &ModifyGatewayLoadBalancerAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGatewayLoadBalancerAttribute
// ModifyGatewayLoadBalancerAttribute 接口用于修改负载均衡实例的属性。支持修改负载均衡实例的名称、带宽上限。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
func (c *Client) ModifyGatewayLoadBalancerAttribute(request *ModifyGatewayLoadBalancerAttributeRequest) (response *ModifyGatewayLoadBalancerAttributeResponse, err error) {
    return c.ModifyGatewayLoadBalancerAttributeWithContext(context.Background(), request)
}

// ModifyGatewayLoadBalancerAttribute
// ModifyGatewayLoadBalancerAttribute 接口用于修改负载均衡实例的属性。支持修改负载均衡实例的名称、带宽上限。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
func (c *Client) ModifyGatewayLoadBalancerAttributeWithContext(ctx context.Context, request *ModifyGatewayLoadBalancerAttributeRequest) (response *ModifyGatewayLoadBalancerAttributeResponse, err error) {
    if request == nil {
        request = NewModifyGatewayLoadBalancerAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGatewayLoadBalancerAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGatewayLoadBalancerAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupAttributeRequest() (request *ModifyTargetGroupAttributeRequest) {
    request = &ModifyTargetGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "ModifyTargetGroupAttribute")
    
    
    return
}

func NewModifyTargetGroupAttributeResponse() (response *ModifyTargetGroupAttributeResponse) {
    response = &ModifyTargetGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupAttribute
// 修改目标组的名称、健康探测等属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupAttribute(request *ModifyTargetGroupAttributeRequest) (response *ModifyTargetGroupAttributeResponse, err error) {
    return c.ModifyTargetGroupAttributeWithContext(context.Background(), request)
}

// ModifyTargetGroupAttribute
// 修改目标组的名称、健康探测等属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupAttributeWithContext(ctx context.Context, request *ModifyTargetGroupAttributeRequest) (response *ModifyTargetGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupInstancesWeightRequest() (request *ModifyTargetGroupInstancesWeightRequest) {
    request = &ModifyTargetGroupInstancesWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "ModifyTargetGroupInstancesWeight")
    
    
    return
}

func NewModifyTargetGroupInstancesWeightResponse() (response *ModifyTargetGroupInstancesWeightResponse) {
    response = &ModifyTargetGroupInstancesWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupInstancesWeight
// 修改目标组的服务器权重。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesWeight(request *ModifyTargetGroupInstancesWeightRequest) (response *ModifyTargetGroupInstancesWeightResponse, err error) {
    return c.ModifyTargetGroupInstancesWeightWithContext(context.Background(), request)
}

// ModifyTargetGroupInstancesWeight
// 修改目标组的服务器权重。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesWeightWithContext(ctx context.Context, request *ModifyTargetGroupInstancesWeightRequest) (response *ModifyTargetGroupInstancesWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupInstancesWeightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupInstancesWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupInstancesWeightResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetGroupInstancesRequest() (request *RegisterTargetGroupInstancesRequest) {
    request = &RegisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gwlb", APIVersion, "RegisterTargetGroupInstances")
    
    
    return
}

func NewRegisterTargetGroupInstancesResponse() (response *RegisterTargetGroupInstancesResponse) {
    response = &RegisterTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterTargetGroupInstances
// 注册服务器到目标组。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetGroupInstances(request *RegisterTargetGroupInstancesRequest) (response *RegisterTargetGroupInstancesResponse, err error) {
    return c.RegisterTargetGroupInstancesWithContext(context.Background(), request)
}

// RegisterTargetGroupInstances
// 注册服务器到目标组。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetGroupInstancesWithContext(ctx context.Context, request *RegisterTargetGroupInstancesRequest) (response *RegisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewRegisterTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}
