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

package v20180317

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-17"

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
    request.Init().WithApiInfo("clb", APIVersion, "AssociateTargetGroups")
    return
}

func NewAssociateTargetGroupsResponse() (response *AssociateTargetGroupsResponse) {
    response = &AssociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateTargetGroups
// 本接口(AssociateTargetGroups)用来将目标组绑定到负载均衡的监听器（四层协议）或转发规则（七层协议）上。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateTargetGroups(request *AssociateTargetGroupsRequest) (response *AssociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateTargetGroupsRequest()
    }
    response = NewAssociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAutoRewriteRequest() (request *AutoRewriteRequest) {
    request = &AutoRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "AutoRewrite")
    return
}

func NewAutoRewriteResponse() (response *AutoRewriteResponse) {
    response = &AutoRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AutoRewrite
// 用户需要先创建出一个HTTPS:443监听器，并在其下创建转发规则。通过调用本接口，系统会自动创建出一个HTTP:80监听器（如果之前不存在），并在其下创建转发规则，与HTTPS:443监听器下的Domains（在入参中指定）对应。创建成功后可以通过HTTP:80地址自动跳转为HTTPS:443地址进行访问。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AutoRewrite(request *AutoRewriteRequest) (response *AutoRewriteResponse, err error) {
    if request == nil {
        request = NewAutoRewriteRequest()
    }
    response = NewAutoRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeregisterTargetsRequest() (request *BatchDeregisterTargetsRequest) {
    request = &BatchDeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "BatchDeregisterTargets")
    return
}

func NewBatchDeregisterTargetsResponse() (response *BatchDeregisterTargetsResponse) {
    response = &BatchDeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeregisterTargets
// 批量解绑四七层后端服务。批量解绑的资源数量上限为500。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchDeregisterTargets(request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchDeregisterTargetsRequest()
    }
    response = NewBatchDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyTargetWeightRequest() (request *BatchModifyTargetWeightRequest) {
    request = &BatchModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "BatchModifyTargetWeight")
    return
}

func NewBatchModifyTargetWeightResponse() (response *BatchModifyTargetWeightResponse) {
    response = &BatchModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyTargetWeight
// BatchModifyTargetWeight 接口用于批量修改负载均衡监听器绑定的后端机器的转发权重。批量修改的资源数量上限为500。本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。<br/>负载均衡的4层和7层监听器支持此接口，传统型负载均衡不支持。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchModifyTargetWeight(request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewBatchModifyTargetWeightRequest()
    }
    response = NewBatchModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRegisterTargetsRequest() (request *BatchRegisterTargetsRequest) {
    request = &BatchRegisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "BatchRegisterTargets")
    return
}

func NewBatchRegisterTargetsResponse() (response *BatchRegisterTargetsResponse) {
    response = &BatchRegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchRegisterTargets
// 批量绑定虚拟主机或弹性网卡，支持跨域绑定，支持四层、七层（TCP、UDP、HTTP、HTTPS）协议绑定。批量绑定的资源数量上限为500。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    response = NewBatchRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClsLogSetRequest() (request *CreateClsLogSetRequest) {
    request = &CreateClsLogSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateClsLogSet")
    return
}

func NewCreateClsLogSetResponse() (response *CreateClsLogSetResponse) {
    response = &CreateClsLogSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClsLogSet
// 创建CLB专有日志集，此日志集用于存储CLB的日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClsLogSet(request *CreateClsLogSetRequest) (response *CreateClsLogSetResponse, err error) {
    if request == nil {
        request = NewCreateClsLogSetRequest()
    }
    response = NewCreateClsLogSetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateListener")
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateListener
// 在一个负载均衡实例下创建监听器。
//
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateLoadBalancer")
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoadBalancer
// 本接口(CreateLoadBalancer)用来创建负载均衡实例（本接口只支持购买按量计费的负载均衡，包年包月的负载均衡请通过控制台购买）。为了使用负载均衡服务，您必须购买一个或多个负载均衡实例。成功调用该接口后，会返回负载均衡实例的唯一 ID。负载均衡实例的类型分为：公网、内网。详情可参考产品说明中的产品类型。
//
// 注意：(1)指定可用区申请负载均衡、跨zone容灾(仅香港支持)【如果您需要体验该功能，请通过 [工单申请](https://console.cloud.tencent.com/workorder/category)】；(2)目前只有北京、上海、广州支持IPv6；(3)一个账号在每个地域的默认购买配额为：公网100个，内网100个。
//
// 本接口为异步接口，接口成功返回后，可使用 DescribeLoadBalancers 接口查询负载均衡实例的状态（如创建中、正常），以确定是否创建成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerSnatIpsRequest() (request *CreateLoadBalancerSnatIpsRequest) {
    request = &CreateLoadBalancerSnatIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateLoadBalancerSnatIps")
    return
}

func NewCreateLoadBalancerSnatIpsResponse() (response *CreateLoadBalancerSnatIpsResponse) {
    response = &CreateLoadBalancerSnatIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoadBalancerSnatIps
// 针对SnatPro负载均衡，这个接口用于添加SnatIp，如果负载均衡没有开启SnatPro，添加SnatIp后会自动开启。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancerSnatIps(request *CreateLoadBalancerSnatIpsRequest) (response *CreateLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerSnatIpsRequest()
    }
    response = NewCreateLoadBalancerSnatIpsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateRule")
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// CreateRule 接口用于在一个已存在的负载均衡七层监听器下创建转发规则，七层监听器中，后端服务必须绑定到规则上而非监听器上。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetGroupRequest() (request *CreateTargetGroupRequest) {
    request = &CreateTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateTargetGroup")
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    response = NewCreateTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "CreateTopic")
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopic
// 创建主题，默认开启全文索引和键值索引。如果不存在CLB专有日志集，则创建失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteListener")
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteListener
// 本接口用来删除负载均衡实例下的监听器（四层和七层）。
//
// 本接口为异步接口，接口返回成功后，需以得到的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
    request = &DeleteLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancer")
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancer
// DeleteLoadBalancer 接口用以删除指定的一个或多个负载均衡实例。成功删除后，会把负载均衡实例下的监听器、转发规则一起删除，并把后端服务解绑。
//
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerListenersRequest() (request *DeleteLoadBalancerListenersRequest) {
    request = &DeleteLoadBalancerListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancerListeners")
    return
}

func NewDeleteLoadBalancerListenersResponse() (response *DeleteLoadBalancerListenersResponse) {
    response = &DeleteLoadBalancerListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancerListeners
// 该接口支持删除负载均衡的多个监听器。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancerListeners(request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerListenersRequest()
    }
    response = NewDeleteLoadBalancerListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerSnatIpsRequest() (request *DeleteLoadBalancerSnatIpsRequest) {
    request = &DeleteLoadBalancerSnatIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancerSnatIps")
    return
}

func NewDeleteLoadBalancerSnatIpsResponse() (response *DeleteLoadBalancerSnatIpsResponse) {
    response = &DeleteLoadBalancerSnatIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancerSnatIps
// 这个接口用于删除SnatPro的负载均衡的SnatIp。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteLoadBalancerSnatIps(request *DeleteLoadBalancerSnatIpsRequest) (response *DeleteLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerSnatIpsRequest()
    }
    response = NewDeleteLoadBalancerSnatIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRewriteRequest() (request *DeleteRewriteRequest) {
    request = &DeleteRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteRewrite")
    return
}

func NewDeleteRewriteResponse() (response *DeleteRewriteResponse) {
    response = &DeleteRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRewrite
// DeleteRewrite 接口支持删除指定转发规则之间的重定向关系。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRewrite(request *DeleteRewriteRequest) (response *DeleteRewriteResponse, err error) {
    if request == nil {
        request = NewDeleteRewriteRequest()
    }
    response = NewDeleteRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRule
// DeleteRule 接口用来删除负载均衡实例七层监听器下的转发规则。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetGroupsRequest() (request *DeleteTargetGroupsRequest) {
    request = &DeleteTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeleteTargetGroups")
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
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    response = NewDeleteTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetGroupInstancesRequest() (request *DeregisterTargetGroupInstancesRequest) {
    request = &DeregisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargetGroupInstances")
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetGroupInstances(request *DeregisterTargetGroupInstancesRequest) (response *DeregisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetGroupInstancesRequest()
    }
    response = NewDeregisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetsRequest() (request *DeregisterTargetsRequest) {
    request = &DeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargets")
    return
}

func NewDeregisterTargetsResponse() (response *DeregisterTargetsResponse) {
    response = &DeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeregisterTargets
// DeregisterTargets 接口用来将一台或多台后端服务从负载均衡的监听器或转发规则上解绑，对于四层监听器，只需指定监听器ID即可，对于七层监听器，还需通过LocationId或Domain+Url指定转发规则。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargets(request *DeregisterTargetsRequest) (response *DeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetsRequest()
    }
    response = NewDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetsFromClassicalLBRequest() (request *DeregisterTargetsFromClassicalLBRequest) {
    request = &DeregisterTargetsFromClassicalLBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargetsFromClassicalLB")
    return
}

func NewDeregisterTargetsFromClassicalLBResponse() (response *DeregisterTargetsFromClassicalLBResponse) {
    response = &DeregisterTargetsFromClassicalLBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeregisterTargetsFromClassicalLB
// DeregisterTargetsFromClassicalLB 接口用于解绑负载均衡后端服务。本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetsFromClassicalLB(request *DeregisterTargetsFromClassicalLBRequest) (response *DeregisterTargetsFromClassicalLBResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetsFromClassicalLBRequest()
    }
    response = NewDeregisterTargetsFromClassicalLBResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIPListRequest() (request *DescribeBlockIPListRequest) {
    request = &DescribeBlockIPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPList")
    return
}

func NewDescribeBlockIPListResponse() (response *DescribeBlockIPListResponse) {
    response = &DescribeBlockIPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlockIPList
// 查询一个负载均衡所封禁的IP列表（黑名单）。（接口灰度中，如需使用请提工单）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBlockIPList(request *DescribeBlockIPListRequest) (response *DescribeBlockIPListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIPListRequest()
    }
    response = NewDescribeBlockIPListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIPTaskRequest() (request *DescribeBlockIPTaskRequest) {
    request = &DescribeBlockIPTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPTask")
    return
}

func NewDescribeBlockIPTaskResponse() (response *DescribeBlockIPTaskResponse) {
    response = &DescribeBlockIPTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlockIPTask
// 根据 ModifyBlockIPList 接口返回的异步任务的ID，查询封禁IP（黑名单）异步任务的执行状态。（接口灰度中，如需使用请提工单）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBlockIPTask(request *DescribeBlockIPTaskRequest) (response *DescribeBlockIPTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIPTaskRequest()
    }
    response = NewDescribeBlockIPTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBByInstanceIdRequest() (request *DescribeClassicalLBByInstanceIdRequest) {
    request = &DescribeClassicalLBByInstanceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBByInstanceId")
    return
}

func NewDescribeClassicalLBByInstanceIdResponse() (response *DescribeClassicalLBByInstanceIdResponse) {
    response = &DescribeClassicalLBByInstanceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClassicalLBByInstanceId
// DescribeClassicalLBByInstanceId用于通过后端实例ID获取传统型负载均衡ID列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBByInstanceId(request *DescribeClassicalLBByInstanceIdRequest) (response *DescribeClassicalLBByInstanceIdResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBByInstanceIdRequest()
    }
    response = NewDescribeClassicalLBByInstanceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBHealthStatusRequest() (request *DescribeClassicalLBHealthStatusRequest) {
    request = &DescribeClassicalLBHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBHealthStatus")
    return
}

func NewDescribeClassicalLBHealthStatusResponse() (response *DescribeClassicalLBHealthStatusResponse) {
    response = &DescribeClassicalLBHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClassicalLBHealthStatus
// DescribeClassicalLBHealthStatus用于获取传统型负载均衡后端的健康状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBHealthStatus(request *DescribeClassicalLBHealthStatusRequest) (response *DescribeClassicalLBHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBHealthStatusRequest()
    }
    response = NewDescribeClassicalLBHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBListenersRequest() (request *DescribeClassicalLBListenersRequest) {
    request = &DescribeClassicalLBListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBListeners")
    return
}

func NewDescribeClassicalLBListenersResponse() (response *DescribeClassicalLBListenersResponse) {
    response = &DescribeClassicalLBListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClassicalLBListeners
// DescribeClassicalLBListeners 接口用于获取传统型负载均衡的监听器信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBListeners(request *DescribeClassicalLBListenersRequest) (response *DescribeClassicalLBListenersResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBListenersRequest()
    }
    response = NewDescribeClassicalLBListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBTargetsRequest() (request *DescribeClassicalLBTargetsRequest) {
    request = &DescribeClassicalLBTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBTargets")
    return
}

func NewDescribeClassicalLBTargetsResponse() (response *DescribeClassicalLBTargetsResponse) {
    response = &DescribeClassicalLBTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClassicalLBTargets
// DescribeClassicalLBTargets用于获取传统型负载均衡绑定的后端服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBTargets(request *DescribeClassicalLBTargetsRequest) (response *DescribeClassicalLBTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBTargetsRequest()
    }
    response = NewDescribeClassicalLBTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClsLogSetRequest() (request *DescribeClsLogSetRequest) {
    request = &DescribeClsLogSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClsLogSet")
    return
}

func NewDescribeClsLogSetResponse() (response *DescribeClsLogSetResponse) {
    response = &DescribeClsLogSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClsLogSet
// 获取用户的CLB专有日志集。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClsLogSet(request *DescribeClsLogSetRequest) (response *DescribeClsLogSetResponse, err error) {
    if request == nil {
        request = NewDescribeClsLogSetRequest()
    }
    response = NewDescribeClsLogSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterResourcesRequest() (request *DescribeClusterResourcesRequest) {
    request = &DescribeClusterResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClusterResources")
    return
}

func NewDescribeClusterResourcesResponse() (response *DescribeClusterResourcesResponse) {
    response = &DescribeClusterResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterResources
// 查询独占集群中的资源列表，支持按集群ID、VIP、负载均衡ID、是否闲置为过滤条件检索。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeClusterResources(request *DescribeClusterResourcesRequest) (response *DescribeClusterResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterResourcesRequest()
    }
    response = NewDescribeClusterResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExclusiveClustersRequest() (request *DescribeExclusiveClustersRequest) {
    request = &DescribeExclusiveClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeExclusiveClusters")
    return
}

func NewDescribeExclusiveClustersResponse() (response *DescribeExclusiveClustersResponse) {
    response = &DescribeExclusiveClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExclusiveClusters
// 查询集群信息列表，支持以集群类型、集群唯一ID、集群名字、集群标签、集群内vip、集群内负载均衡唯一id、集群网络类型、可用区等条件进行检索
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeExclusiveClusters(request *DescribeExclusiveClustersRequest) (response *DescribeExclusiveClustersResponse, err error) {
    if request == nil {
        request = NewDescribeExclusiveClustersRequest()
    }
    response = NewDescribeExclusiveClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeListeners")
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListeners
// DescribeListeners 接口可根据负载均衡器 ID、监听器的协议或端口作为过滤条件获取监听器列表。如果不指定任何过滤条件，则返回该负载均衡实例下的所有监听器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerListByCertIdRequest() (request *DescribeLoadBalancerListByCertIdRequest) {
    request = &DescribeLoadBalancerListByCertIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerListByCertId")
    return
}

func NewDescribeLoadBalancerListByCertIdResponse() (response *DescribeLoadBalancerListByCertIdResponse) {
    response = &DescribeLoadBalancerListByCertIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancerListByCertId
// 根据证书ID查询其在一个地域中所关联到负载均衡实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancerListByCertId(request *DescribeLoadBalancerListByCertIdRequest) (response *DescribeLoadBalancerListByCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerListByCertIdRequest()
    }
    response = NewDescribeLoadBalancerListByCertIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerTrafficRequest() (request *DescribeLoadBalancerTrafficRequest) {
    request = &DescribeLoadBalancerTrafficRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerTraffic")
    return
}

func NewDescribeLoadBalancerTrafficResponse() (response *DescribeLoadBalancerTrafficResponse) {
    response = &DescribeLoadBalancerTrafficResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancerTraffic
// 查询账号下的高流量负载均衡，返回前10个负载均衡。如果是子账号登录，只返回子账号有权限的负载均衡。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeLoadBalancerTraffic(request *DescribeLoadBalancerTrafficRequest) (response *DescribeLoadBalancerTrafficResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerTrafficRequest()
    }
    response = NewDescribeLoadBalancerTrafficResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancers")
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancers
// 查询一个地域的负载均衡实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersDetailRequest() (request *DescribeLoadBalancersDetailRequest) {
    request = &DescribeLoadBalancersDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancersDetail")
    return
}

func NewDescribeLoadBalancersDetailResponse() (response *DescribeLoadBalancersDetailResponse) {
    response = &DescribeLoadBalancersDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancersDetail
// 查询负载均衡的详细信息，包括监听器，规则及后端目标。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancersDetail(request *DescribeLoadBalancersDetailRequest) (response *DescribeLoadBalancersDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersDetailRequest()
    }
    response = NewDescribeLoadBalancersDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeQuota")
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQuota
// 查询用户当前地域下的各项配额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRewriteRequest() (request *DescribeRewriteRequest) {
    request = &DescribeRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeRewrite")
    return
}

func NewDescribeRewriteResponse() (response *DescribeRewriteResponse) {
    response = &DescribeRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRewrite
// DescribeRewrite 接口可根据负载均衡实例ID，查询一个负载均衡实例下转发规则的重定向关系。如果不指定监听器ID或转发规则ID，则返回该负载均衡实例下的所有重定向关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRewrite(request *DescribeRewriteRequest) (response *DescribeRewriteResponse, err error) {
    if request == nil {
        request = NewDescribeRewriteRequest()
    }
    response = NewDescribeRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupInstancesRequest() (request *DescribeTargetGroupInstancesRequest) {
    request = &DescribeTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroupInstances")
    return
}

func NewDescribeTargetGroupInstancesResponse() (response *DescribeTargetGroupInstancesResponse) {
    response = &DescribeTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTargetGroupInstances
// 获取目标组绑定的服务器信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupInstances(request *DescribeTargetGroupInstancesRequest) (response *DescribeTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupInstancesRequest()
    }
    response = NewDescribeTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupListRequest() (request *DescribeTargetGroupListRequest) {
    request = &DescribeTargetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroupList")
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
    if request == nil {
        request = NewDescribeTargetGroupListRequest()
    }
    response = NewDescribeTargetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
    request = &DescribeTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroups")
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
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    response = NewDescribeTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetHealthRequest() (request *DescribeTargetHealthRequest) {
    request = &DescribeTargetHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetHealth")
    return
}

func NewDescribeTargetHealthResponse() (response *DescribeTargetHealthResponse) {
    response = &DescribeTargetHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTargetHealth
// DescribeTargetHealth 接口用来获取负载均衡后端服务的健康检查结果，不支持传统型负载均衡。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    if request == nil {
        request = NewDescribeTargetHealthRequest()
    }
    response = NewDescribeTargetHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetsRequest() (request *DescribeTargetsRequest) {
    request = &DescribeTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargets")
    return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
    response = &DescribeTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTargets
// DescribeTargets 接口用来查询负载均衡实例的某些监听器绑定的后端服务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetsRequest()
    }
    response = NewDescribeTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTaskStatus")
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStatus
// 本接口用于查询异步任务的执行状态，对于非查询类的接口（创建/删除负载均衡实例、监听器、规则以及绑定或解绑后端服务等），在接口调用成功后，都需要使用本接口查询任务最终是否执行成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateTargetGroupsRequest() (request *DisassociateTargetGroupsRequest) {
    request = &DisassociateTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "DisassociateTargetGroups")
    return
}

func NewDisassociateTargetGroupsResponse() (response *DisassociateTargetGroupsResponse) {
    response = &DisassociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateTargetGroups
// 解除规则的目标组关联关系。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateTargetGroups(request *DisassociateTargetGroupsRequest) (response *DisassociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateTargetGroupsRequest()
    }
    response = NewDisassociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewManualRewriteRequest() (request *ManualRewriteRequest) {
    request = &ManualRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ManualRewrite")
    return
}

func NewManualRewriteResponse() (response *ManualRewriteResponse) {
    response = &ManualRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManualRewrite
// 用户手动配置原访问地址和重定向地址，系统自动将原访问地址的请求重定向至对应路径的目的地址。同一域名下可以配置多条路径作为重定向策略，实现http/https之间请求的自动跳转。设置重定向时，需满足如下约束条件：若A已经重定向至B，则A不能再重定向至C（除非先删除老的重定向关系，再建立新的重定向关系），B不能重定向至任何其它地址。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ManualRewrite(request *ManualRewriteRequest) (response *ManualRewriteResponse, err error) {
    if request == nil {
        request = NewManualRewriteRequest()
    }
    response = NewManualRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIPListRequest() (request *ModifyBlockIPListRequest) {
    request = &ModifyBlockIPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyBlockIPList")
    return
}

func NewModifyBlockIPListResponse() (response *ModifyBlockIPListResponse) {
    response = &ModifyBlockIPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlockIPList
// 修改负载均衡的IP（client IP）封禁黑名单列表，一个转发规则最多支持封禁 2000000 个IP，及黑名单容量为 2000000。
//
// （接口灰度中，如需使用请提工单）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIPList(request *ModifyBlockIPListRequest) (response *ModifyBlockIPListResponse, err error) {
    if request == nil {
        request = NewModifyBlockIPListRequest()
    }
    response = NewModifyBlockIPListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
    request = &ModifyDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyDomain")
    return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
    response = &ModifyDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomain
// ModifyDomain接口用来修改负载均衡七层监听器下的域名。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    if request == nil {
        request = NewModifyDomainRequest()
    }
    response = NewModifyDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainAttributesRequest() (request *ModifyDomainAttributesRequest) {
    request = &ModifyDomainAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyDomainAttributes")
    return
}

func NewModifyDomainAttributesResponse() (response *ModifyDomainAttributesResponse) {
    response = &ModifyDomainAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainAttributes
// ModifyDomainAttributes接口用于修改负载均衡7层监听器转发规则的域名级别属性，如修改域名、修改DefaultServer、开启/关闭Http2、修改证书。
//
// 本接口为异步接口，本接口返回成功后，需以返回的RequestId为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainAttributes(request *ModifyDomainAttributesRequest) (response *ModifyDomainAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDomainAttributesRequest()
    }
    response = NewModifyDomainAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyListener")
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyListener
// ModifyListener接口用来修改负载均衡监听器的属性，包括监听器名称、健康检查参数、证书信息、转发策略等。本接口不支持传统型负载均衡。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
    request = &ModifyLoadBalancerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyLoadBalancerAttributes")
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancerAttributes
// 修改负载均衡实例的属性。支持修改负载均衡实例的名称、设置负载均衡的跨域属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyRule")
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRule
// ModifyRule 接口用来修改负载均衡七层监听器下的转发规则的各项属性，包括转发路径、健康检查属性、转发策略等。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupAttributeRequest() (request *ModifyTargetGroupAttributeRequest) {
    request = &ModifyTargetGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupAttribute")
    return
}

func NewModifyTargetGroupAttributeResponse() (response *ModifyTargetGroupAttributeResponse) {
    response = &ModifyTargetGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetGroupAttribute
// 修改目标组的名称或者默认端口属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupAttribute(request *ModifyTargetGroupAttributeRequest) (response *ModifyTargetGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupAttributeRequest()
    }
    response = NewModifyTargetGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupInstancesPortRequest() (request *ModifyTargetGroupInstancesPortRequest) {
    request = &ModifyTargetGroupInstancesPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupInstancesPort")
    return
}

func NewModifyTargetGroupInstancesPortResponse() (response *ModifyTargetGroupInstancesPortResponse) {
    response = &ModifyTargetGroupInstancesPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetGroupInstancesPort
// 批量修改目标组服务器端口。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesPort(request *ModifyTargetGroupInstancesPortRequest) (response *ModifyTargetGroupInstancesPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupInstancesPortRequest()
    }
    response = NewModifyTargetGroupInstancesPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupInstancesWeightRequest() (request *ModifyTargetGroupInstancesWeightRequest) {
    request = &ModifyTargetGroupInstancesWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupInstancesWeight")
    return
}

func NewModifyTargetGroupInstancesWeightResponse() (response *ModifyTargetGroupInstancesWeightResponse) {
    response = &ModifyTargetGroupInstancesWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetGroupInstancesWeight
// 批量修改目标组的服务器权重。
//
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesWeight(request *ModifyTargetGroupInstancesWeightRequest) (response *ModifyTargetGroupInstancesWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupInstancesWeightRequest()
    }
    response = NewModifyTargetGroupInstancesWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
    request = &ModifyTargetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetPort")
    return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
    response = &ModifyTargetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetPort
// ModifyTargetPort接口用于修改监听器绑定的后端服务的端口。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetPortRequest()
    }
    response = NewModifyTargetPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetWeightRequest() (request *ModifyTargetWeightRequest) {
    request = &ModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetWeight")
    return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
    response = &ModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetWeight
// ModifyTargetWeight 接口用于修改负载均衡绑定的后端服务的转发权重。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetWeightRequest()
    }
    response = NewModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetGroupInstancesRequest() (request *RegisterTargetGroupInstancesRequest) {
    request = &RegisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargetGroupInstances")
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetGroupInstances(request *RegisterTargetGroupInstancesRequest) (response *RegisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewRegisterTargetGroupInstancesRequest()
    }
    response = NewRegisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetsRequest() (request *RegisterTargetsRequest) {
    request = &RegisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargets")
    return
}

func NewRegisterTargetsResponse() (response *RegisterTargetsResponse) {
    response = &RegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterTargets
// RegisterTargets 接口用来将一台或多台后端服务绑定到负载均衡的监听器（或7层转发规则），在此之前您需要先行创建相关的4层监听器或7层转发规则。对于四层监听器（TCP、UDP），只需指定监听器ID即可，对于七层监听器（HTTP、HTTPS），还需通过LocationId或者Domain+Url指定转发规则。
//
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargets(request *RegisterTargetsRequest) (response *RegisterTargetsResponse, err error) {
    if request == nil {
        request = NewRegisterTargetsRequest()
    }
    response = NewRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetsWithClassicalLBRequest() (request *RegisterTargetsWithClassicalLBRequest) {
    request = &RegisterTargetsWithClassicalLBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargetsWithClassicalLB")
    return
}

func NewRegisterTargetsWithClassicalLBResponse() (response *RegisterTargetsWithClassicalLBResponse) {
    response = &RegisterTargetsWithClassicalLBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterTargetsWithClassicalLB
// RegisterTargetsWithClassicalLB 接口用于绑定后端服务到传统型负载均衡。本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetsWithClassicalLB(request *RegisterTargetsWithClassicalLBRequest) (response *RegisterTargetsWithClassicalLBResponse, err error) {
    if request == nil {
        request = NewRegisterTargetsWithClassicalLBRequest()
    }
    response = NewRegisterTargetsWithClassicalLBResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertForLoadBalancersRequest() (request *ReplaceCertForLoadBalancersRequest) {
    request = &ReplaceCertForLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "ReplaceCertForLoadBalancers")
    return
}

func NewReplaceCertForLoadBalancersResponse() (response *ReplaceCertForLoadBalancersResponse) {
    response = &ReplaceCertForLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceCertForLoadBalancers
// ReplaceCertForLoadBalancers 接口用以替换负载均衡实例所关联的证书，对于各个地域的负载均衡，如果指定的老的证书ID与其有关联关系，则会先解除关联，再建立新证书与该负载均衡的关联关系。
//
// 此接口支持替换服务端证书或客户端证书。
//
// 需要使用的新证书，可以通过传入证书ID来指定，如果不指定证书ID，则必须传入证书内容等相关信息，用以新建证书并绑定至负载均衡。
//
// 注：本接口仅可从广州地域调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) ReplaceCertForLoadBalancers(request *ReplaceCertForLoadBalancersRequest) (response *ReplaceCertForLoadBalancersResponse, err error) {
    if request == nil {
        request = NewReplaceCertForLoadBalancersRequest()
    }
    response = NewReplaceCertForLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerClsLogRequest() (request *SetLoadBalancerClsLogRequest) {
    request = &SetLoadBalancerClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerClsLog")
    return
}

func NewSetLoadBalancerClsLogResponse() (response *SetLoadBalancerClsLogResponse) {
    response = &SetLoadBalancerClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetLoadBalancerClsLog
// 增加、删除、更新负载均衡的日志服务(CLS)主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerClsLog(request *SetLoadBalancerClsLogRequest) (response *SetLoadBalancerClsLogResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerClsLogRequest()
    }
    response = NewSetLoadBalancerClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerSecurityGroups")
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetLoadBalancerSecurityGroups
// SetLoadBalancerSecurityGroups 接口支持对一个公网负载均衡实例执行设置（绑定、解绑）安全组操作。查询一个负载均衡实例目前已绑定的安全组，可使用 DescribeLoadBalancers 接口。本接口是set语义，
//
// 绑定操作时，入参需要传入负载均衡实例要绑定的所有安全组（已绑定的+新增绑定的）。
//
// 解绑操作时，入参需要传入负载均衡实例执行解绑后所绑定的所有安全组；如果要解绑所有安全组，可不传此参数，或传入空数组。注意：内网负载均衡不支持绑定安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewSetSecurityGroupForLoadbalancersRequest() (request *SetSecurityGroupForLoadbalancersRequest) {
    request = &SetSecurityGroupForLoadbalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("clb", APIVersion, "SetSecurityGroupForLoadbalancers")
    return
}

func NewSetSecurityGroupForLoadbalancersResponse() (response *SetSecurityGroupForLoadbalancersResponse) {
    response = &SetSecurityGroupForLoadbalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetSecurityGroupForLoadbalancers
// 绑定或解绑一个安全组到多个公网负载均衡实例。注意：内网负载均衡不支持绑定安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}
