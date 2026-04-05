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

package v20250115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-01-15"

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


func NewCreateAccelerateAreasRequest() (request *CreateAccelerateAreasRequest) {
    request = &CreateAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateAccelerateAreas")
    
    
    return
}

func NewCreateAccelerateAreasResponse() (response *CreateAccelerateAreasResponse) {
    response = &CreateAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccelerateAreas
// 创建加速地域
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_ACCELERATEREGIONREPEAT = "UnsupportedOperation.AccelerateRegionRepeat"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateAccelerateAreas(request *CreateAccelerateAreasRequest) (response *CreateAccelerateAreasResponse, err error) {
    return c.CreateAccelerateAreasWithContext(context.Background(), request)
}

// CreateAccelerateAreas
// 创建加速地域
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_ACCELERATEREGIONREPEAT = "UnsupportedOperation.AccelerateRegionRepeat"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateAccelerateAreasWithContext(ctx context.Context, request *CreateAccelerateAreasRequest) (response *CreateAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewCreateAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndpointGroupRequest() (request *CreateEndpointGroupRequest) {
    request = &CreateEndpointGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateEndpointGroup")
    
    
    return
}

func NewCreateEndpointGroupResponse() (response *CreateEndpointGroupResponse) {
    response = &CreateEndpointGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEndpointGroup
// 创建终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.EnableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_UDPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.UdpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPLICATIONLAYERENDPOINTGROUPPARAMETER = "MissingParameter.ApplicationLayerEndpointGroupParameter"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
//  UNSUPPORTEDOPERATION_THREENETWORKSENDPOINTGROUP = "UnsupportedOperation.ThreeNetworksEndpointGroup"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_VIRTUALENDPOINTGROUPUNSUPPORTEDTCPANDUDP = "UnsupportedOperation.VirtualEndpointGroupUnsupportedTcpAndUdp"
func (c *Client) CreateEndpointGroup(request *CreateEndpointGroupRequest) (response *CreateEndpointGroupResponse, err error) {
    return c.CreateEndpointGroupWithContext(context.Background(), request)
}

// CreateEndpointGroup
// 创建终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DISABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.DisableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENABLEHEALTHCHECKNOTCARRYPARAMETERS = "InvalidParameter.EnableHealthCheckNotCarryParameters"
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETER_UDPENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.UdpEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_NOTMAPPINGMANYPORT = "InvalidParameterValue.NotMappingManyPort"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPLICATIONLAYERENDPOINTGROUPPARAMETER = "MissingParameter.ApplicationLayerEndpointGroupParameter"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  MISSINGPARAMETER_ENABLEHEALTHCHECKPARAMETER = "MissingParameter.EnableHealthCheckParameter"
//  UNSUPPORTEDOPERATION_EXISTACCEPTACLPOLICY = "UnsupportedOperation.ExistAcceptAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTDROPACCEPTACLRULE = "UnsupportedOperation.ExistDropAcceptAclRule"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_NOTSETCROSSBORDERTYPE = "UnsupportedOperation.NotSetCrossBorderType"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
//  UNSUPPORTEDOPERATION_THREENETWORKSENDPOINTGROUP = "UnsupportedOperation.ThreeNetworksEndpointGroup"
//  UNSUPPORTEDOPERATION_UNABLECREATETHIRDPARTYNODES = "UnsupportedOperation.UnableCreateThirdPartyNodes"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_VIRTUALENDPOINTGROUPUNSUPPORTEDTCPANDUDP = "UnsupportedOperation.VirtualEndpointGroupUnsupportedTcpAndUdp"
func (c *Client) CreateEndpointGroupWithContext(ctx context.Context, request *CreateEndpointGroupRequest) (response *CreateEndpointGroupResponse, err error) {
    if request == nil {
        request = NewCreateEndpointGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateEndpointGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEndpointGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEndpointGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateForwardingRuleRequest() (request *CreateForwardingRuleRequest) {
    request = &CreateForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateForwardingRule")
    
    
    return
}

func NewCreateForwardingRuleResponse() (response *CreateForwardingRuleResponse) {
    response = &CreateForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateForwardingRule
// 创建七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
func (c *Client) CreateForwardingRule(request *CreateForwardingRuleRequest) (response *CreateForwardingRuleResponse, err error) {
    return c.CreateForwardingRuleWithContext(context.Background(), request)
}

// CreateForwardingRule
// 创建七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FORWARDGROUPFORWARDINGRULEUNSUPPORTEDDEFAULTENDPOINTGROUP = "UnsupportedOperation.ForwardGroupForwardingRuleUnsupportedDefaultEndpointGroup"
func (c *Client) CreateForwardingRuleWithContext(ctx context.Context, request *CreateForwardingRuleRequest) (response *CreateForwardingRuleResponse, err error) {
    if request == nil {
        request = NewCreateForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalAcceleratorRequest() (request *CreateGlobalAcceleratorRequest) {
    request = &CreateGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateGlobalAccelerator")
    
    
    return
}

func NewCreateGlobalAcceleratorResponse() (response *CreateGlobalAcceleratorResponse) {
    response = &CreateGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalAccelerator
// 创建全球加速实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEGLOBALACCELERATORFAILED = "FailedOperation.CreateGlobalAcceleratorFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"
//  UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateGlobalAccelerator(request *CreateGlobalAcceleratorRequest) (response *CreateGlobalAcceleratorResponse, err error) {
    return c.CreateGlobalAcceleratorWithContext(context.Background(), request)
}

// CreateGlobalAccelerator
// 创建全球加速实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEGLOBALACCELERATORFAILED = "FailedOperation.CreateGlobalAcceleratorFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACCOUNTFROZEN = "UnsupportedOperation.AccountFrozen"
//  UNSUPPORTEDOPERATION_BILLINGFAILED = "UnsupportedOperation.BillingFailed"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_CURRENTACCOUNTNOTALLOWED = "UnsupportedOperation.CurrentAccountNotAllowed"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateGlobalAcceleratorWithContext(ctx context.Context, request *CreateGlobalAcceleratorRequest) (response *CreateGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewCreateGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// 创建监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpsListenerCannotCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATEINVALIDSTATUS = "UnsupportedOperation.CertificateInvalidStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// 创建监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpsListenerCannotCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATEINVALIDSTATUS = "UnsupportedOperation.CertificateInvalidStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "CreateListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccelerateAreasRequest() (request *DeleteAccelerateAreasRequest) {
    request = &DeleteAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteAccelerateAreas")
    
    
    return
}

func NewDeleteAccelerateAreasResponse() (response *DeleteAccelerateAreasResponse) {
    response = &DeleteAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccelerateAreas
// 删除加速地域
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpsListenerCannotCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATEINVALIDSTATUS = "UnsupportedOperation.CertificateInvalidStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
func (c *Client) DeleteAccelerateAreas(request *DeleteAccelerateAreasRequest) (response *DeleteAccelerateAreasResponse, err error) {
    return c.DeleteAccelerateAreasWithContext(context.Background(), request)
}

// DeleteAccelerateAreas
// 删除加速地域
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSUFFICIENTRESOURCES = "FailedOperation.InsufficientResources"
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_HTTPSLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpsListenerCannotCarryParameters"
//  INVALIDPARAMETER_SINGLEPORT = "InvalidParameter.SinglePort"
//  INVALIDPARAMETER_TCPENDPOINTGROUPCARRYPARAMETERS = "InvalidParameter.TcpEndpointGroupCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SEGMENTPORTRANGE = "InvalidParameterValue.SegmentPortRange"
//  INVALIDPARAMETERVALUE_TCPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.TcpSeriesListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTEQUAL = "InvalidParameterValue.ThirdAreaListenerPortEqual"
//  INVALIDPARAMETERVALUE_THIRDAREALISTENERPORTSEGMENT = "InvalidParameterValue.ThirdAreaListenerPortSegment"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UDPSERIESLISTENERPORTEQUAL = "InvalidParameterValue.UdpSeriesListenerPortEqual"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_CERTIFICATEINVALIDSTATUS = "UnsupportedOperation.CertificateInvalidStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
//  UNSUPPORTEDOPERATION_LISTENERPORTSEGMENT = "UnsupportedOperation.ListenerPortSegment"
func (c *Client) DeleteAccelerateAreasWithContext(ctx context.Context, request *DeleteAccelerateAreasRequest) (response *DeleteAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewDeleteAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndpointGroupsRequest() (request *DeleteEndpointGroupsRequest) {
    request = &DeleteEndpointGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteEndpointGroups")
    
    
    return
}

func NewDeleteEndpointGroupsResponse() (response *DeleteEndpointGroupsResponse) {
    response = &DeleteEndpointGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEndpointGroups
// 删除终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_APPLICATIONLAYERENDPOINTGROUPNOTDELETE = "UnsupportedOperation.ApplicationLayerEndpointGroupNotDelete"
//  UNSUPPORTEDOPERATION_EXISTFORWARDINGRULE = "UnsupportedOperation.ExistForwardingRule"
func (c *Client) DeleteEndpointGroups(request *DeleteEndpointGroupsRequest) (response *DeleteEndpointGroupsResponse, err error) {
    return c.DeleteEndpointGroupsWithContext(context.Background(), request)
}

// DeleteEndpointGroups
// 删除终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_APPLICATIONLAYERENDPOINTGROUPNOTDELETE = "UnsupportedOperation.ApplicationLayerEndpointGroupNotDelete"
//  UNSUPPORTEDOPERATION_EXISTFORWARDINGRULE = "UnsupportedOperation.ExistForwardingRule"
func (c *Client) DeleteEndpointGroupsWithContext(ctx context.Context, request *DeleteEndpointGroupsRequest) (response *DeleteEndpointGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteEndpointGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteEndpointGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEndpointGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEndpointGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteForwardingRuleRequest() (request *DeleteForwardingRuleRequest) {
    request = &DeleteForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteForwardingRule")
    
    
    return
}

func NewDeleteForwardingRuleResponse() (response *DeleteForwardingRuleResponse) {
    response = &DeleteForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteForwardingRule
// 删除七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingRule(request *DeleteForwardingRuleRequest) (response *DeleteForwardingRuleResponse, err error) {
    return c.DeleteForwardingRuleWithContext(context.Background(), request)
}

// DeleteForwardingRule
// 删除七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) DeleteForwardingRuleWithContext(ctx context.Context, request *DeleteForwardingRuleRequest) (response *DeleteForwardingRuleResponse, err error) {
    if request == nil {
        request = NewDeleteForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalAcceleratorRequest() (request *DeleteGlobalAcceleratorRequest) {
    request = &DeleteGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteGlobalAccelerator")
    
    
    return
}

func NewDeleteGlobalAcceleratorResponse() (response *DeleteGlobalAcceleratorResponse) {
    response = &DeleteGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalAccelerator
// 删除全球加速实例
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTGLOBALACCELERATORACLPOLICY = "UnsupportedOperation.ExistGlobalAcceleratorAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAccelerator(request *DeleteGlobalAcceleratorRequest) (response *DeleteGlobalAcceleratorResponse, err error) {
    return c.DeleteGlobalAcceleratorWithContext(context.Background(), request)
}

// DeleteGlobalAccelerator
// 删除全球加速实例
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTACCELERATORAREA = "UnsupportedOperation.ExistAcceleratorArea"
//  UNSUPPORTEDOPERATION_EXISTGLOBALACCELERATORACLPOLICY = "UnsupportedOperation.ExistGlobalAcceleratorAclPolicy"
//  UNSUPPORTEDOPERATION_EXISTLISTENER = "UnsupportedOperation.ExistListener"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) DeleteGlobalAcceleratorWithContext(ctx context.Context, request *DeleteGlobalAcceleratorRequest) (response *DeleteGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// 删除监听器
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// 删除监听器
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DeleteListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerateAreasRequest() (request *DescribeAccelerateAreasRequest) {
    request = &DescribeAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeAccelerateAreas")
    
    
    return
}

func NewDescribeAccelerateAreasResponse() (response *DescribeAccelerateAreasResponse) {
    response = &DescribeAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerateAreas
// 查询加速地域
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateAreas(request *DescribeAccelerateAreasRequest) (response *DescribeAccelerateAreasResponse, err error) {
    return c.DescribeAccelerateAreasWithContext(context.Background(), request)
}

// DescribeAccelerateAreas
// 查询加速地域
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateAreasWithContext(ctx context.Context, request *DescribeAccelerateAreasRequest) (response *DescribeAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerateRegionsRequest() (request *DescribeAccelerateRegionsRequest) {
    request = &DescribeAccelerateRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeAccelerateRegions")
    
    
    return
}

func NewDescribeAccelerateRegionsResponse() (response *DescribeAccelerateRegionsResponse) {
    response = &DescribeAccelerateRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerateRegions
// 查询可选加速区域
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateRegions(request *DescribeAccelerateRegionsRequest) (response *DescribeAccelerateRegionsResponse, err error) {
    return c.DescribeAccelerateRegionsWithContext(context.Background(), request)
}

// DescribeAccelerateRegions
// 查询可选加速区域
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_EXISTENDPOINTGROUP = "UnsupportedOperation.ExistEndpointGroup"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) DescribeAccelerateRegionsWithContext(ctx context.Context, request *DescribeAccelerateRegionsRequest) (response *DescribeAccelerateRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerateRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeAccelerateRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerateRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerateRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossBorderSettlementRequest() (request *DescribeCrossBorderSettlementRequest) {
    request = &DescribeCrossBorderSettlementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    
    return
}

func NewDescribeCrossBorderSettlementResponse() (response *DescribeCrossBorderSettlementResponse) {
    response = &DescribeCrossBorderSettlementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossBorderSettlement
// 查询跨境账单
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlement(request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    return c.DescribeCrossBorderSettlementWithContext(context.Background(), request)
}

// DescribeCrossBorderSettlement
// 查询跨境账单
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlementWithContext(ctx context.Context, request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBorderSettlementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossBorderSettlement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossBorderSettlementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndpointGroupsRequest() (request *DescribeEndpointGroupsRequest) {
    request = &DescribeEndpointGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeEndpointGroups")
    
    
    return
}

func NewDescribeEndpointGroupsResponse() (response *DescribeEndpointGroupsResponse) {
    response = &DescribeEndpointGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEndpointGroups
// 查询终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeEndpointGroups(request *DescribeEndpointGroupsRequest) (response *DescribeEndpointGroupsResponse, err error) {
    return c.DescribeEndpointGroupsWithContext(context.Background(), request)
}

// DescribeEndpointGroups
// 查询终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeEndpointGroupsWithContext(ctx context.Context, request *DescribeEndpointGroupsRequest) (response *DescribeEndpointGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeEndpointGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeEndpointGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEndpointGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEndpointGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardingRuleRequest() (request *DescribeForwardingRuleRequest) {
    request = &DescribeForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeForwardingRule")
    
    
    return
}

func NewDescribeForwardingRuleResponse() (response *DescribeForwardingRuleResponse) {
    response = &DescribeForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForwardingRule
// 查看七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeForwardingRule(request *DescribeForwardingRuleRequest) (response *DescribeForwardingRuleResponse, err error) {
    return c.DescribeForwardingRuleWithContext(context.Background(), request)
}

// DescribeForwardingRule
// 查看七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeForwardingRuleWithContext(ctx context.Context, request *DescribeForwardingRuleRequest) (response *DescribeForwardingRuleResponse, err error) {
    if request == nil {
        request = NewDescribeForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalAcceleratorsRequest() (request *DescribeGlobalAcceleratorsRequest) {
    request = &DescribeGlobalAcceleratorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeGlobalAccelerators")
    
    
    return
}

func NewDescribeGlobalAcceleratorsResponse() (response *DescribeGlobalAcceleratorsResponse) {
    response = &DescribeGlobalAcceleratorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalAccelerators
// 修改全球加速实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeGlobalAccelerators(request *DescribeGlobalAcceleratorsRequest) (response *DescribeGlobalAcceleratorsResponse, err error) {
    return c.DescribeGlobalAcceleratorsWithContext(context.Background(), request)
}

// DescribeGlobalAccelerators
// 修改全球加速实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeGlobalAcceleratorsWithContext(ctx context.Context, request *DescribeGlobalAcceleratorsRequest) (response *DescribeGlobalAcceleratorsResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalAcceleratorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeGlobalAccelerators")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalAccelerators require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalAcceleratorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// 查询监听器
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// 查询监听器
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeListeners")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerateAreasRequest() (request *ModifyAccelerateAreasRequest) {
    request = &ModifyAccelerateAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyAccelerateAreas")
    
    
    return
}

func NewModifyAccelerateAreasResponse() (response *ModifyAccelerateAreasResponse) {
    response = &ModifyAccelerateAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerateAreas
// 修改加速地域
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyAccelerateAreas(request *ModifyAccelerateAreasRequest) (response *ModifyAccelerateAreasResponse, err error) {
    return c.ModifyAccelerateAreasWithContext(context.Background(), request)
}

// ModifyAccelerateAreas
// 修改加速地域
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTDUPLICATEWITHEXISTINGDATA = "InvalidParameter.InputDuplicateWithExistingData"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyAccelerateAreasWithContext(ctx context.Context, request *ModifyAccelerateAreasRequest) (response *ModifyAccelerateAreasResponse, err error) {
    if request == nil {
        request = NewModifyAccelerateAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyAccelerateAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerateAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerateAreasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEndpointGroupRequest() (request *ModifyEndpointGroupRequest) {
    request = &ModifyEndpointGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyEndpointGroup")
    
    
    return
}

func NewModifyEndpointGroupResponse() (response *ModifyEndpointGroupResponse) {
    response = &ModifyEndpointGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEndpointGroup
// 修改终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  UNSUPPORTEDOPERATION_DEFAULTENDPOINTGROUPMODIFY = "UnsupportedOperation.DefaultEndpointGroupModify"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
func (c *Client) ModifyEndpointGroup(request *ModifyEndpointGroupRequest) (response *ModifyEndpointGroupResponse, err error) {
    return c.ModifyEndpointGroupWithContext(context.Background(), request)
}

// ModifyEndpointGroup
// 修改终端节点组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.EndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_HTTPSORIGINSERVERPROTOCOLSUPPORTSPARAMETERS = "InvalidParameter.HttpsOriginServerProtocolSupportsParameters"
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETER_INPUTREPEAT = "InvalidParameter.InputRepeat"
//  INVALIDPARAMETER_ONLYENDPOINTGROUPCHECKTYPEHTTPCARRYPARAMETERS = "InvalidParameter.OnlyEndpointGroupCheckTypeHttpCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERENDPOINTGROUPCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerEndpointGroupCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_LISTENERPORTNOTEQUAL = "InvalidParameterValue.ListenerPortNotEqual"
//  INVALIDPARAMETERVALUE_NOTWITHINPUBLICNETWORK = "InvalidParameterValue.NotWithinPublicNetwork"
//  INVALIDPARAMETERVALUE_SPECIFICCHARACTERS = "InvalidParameterValue.SpecificCharacters"
//  MISSINGPARAMETER_CUSTOMCHECKTYPEPARAMETER = "MissingParameter.CustomCheckTypeParameter"
//  UNSUPPORTEDOPERATION_DEFAULTENDPOINTGROUPMODIFY = "UnsupportedOperation.DefaultEndpointGroupModify"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORT = "UnsupportedOperation.Ipv6NotSupport"
//  UNSUPPORTEDOPERATION_SAMEDOMAIN = "UnsupportedOperation.SameDomain"
//  UNSUPPORTEDOPERATION_SAMEPUBLICIP = "UnsupportedOperation.SamePublicIp"
func (c *Client) ModifyEndpointGroupWithContext(ctx context.Context, request *ModifyEndpointGroupRequest) (response *ModifyEndpointGroupResponse, err error) {
    if request == nil {
        request = NewModifyEndpointGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyEndpointGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEndpointGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEndpointGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyForwardingRuleRequest() (request *ModifyForwardingRuleRequest) {
    request = &ModifyForwardingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyForwardingRule")
    
    
    return
}

func NewModifyForwardingRuleResponse() (response *ModifyForwardingRuleResponse) {
    response = &ModifyForwardingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyForwardingRule
// 修改七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingRule(request *ModifyForwardingRuleRequest) (response *ModifyForwardingRuleResponse, err error) {
    return c.ModifyForwardingRuleWithContext(context.Background(), request)
}

// ModifyForwardingRule
// 修改七层转发规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_ASCIICHARACTERS = "InvalidParameterValue.AsciiCharacters"
//  INVALIDPARAMETERVALUE_FORWARDINGRULEPATHCONFLICT = "InvalidParameterValue.ForwardingRulePathConflict"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_SPECIALCHARACTERS = "InvalidParameterValue.SpecialCharacters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DEFAULTFORWARDINGPOLICYOPERATE = "UnsupportedOperation.DefaultForwardingPolicyOperate"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
//  UNSUPPORTEDOPERATION_TRANSPORTLAYERUNSUPPORTEDOPERATEFORWARDINGPOLICY = "UnsupportedOperation.TransportLayerUnsupportedOperateForwardingPolicy"
func (c *Client) ModifyForwardingRuleWithContext(ctx context.Context, request *ModifyForwardingRuleRequest) (response *ModifyForwardingRuleResponse, err error) {
    if request == nil {
        request = NewModifyForwardingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyForwardingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyForwardingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyForwardingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalAcceleratorRequest() (request *ModifyGlobalAcceleratorRequest) {
    request = &ModifyGlobalAcceleratorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyGlobalAccelerator")
    
    
    return
}

func NewModifyGlobalAcceleratorResponse() (response *ModifyGlobalAcceleratorResponse) {
    response = &ModifyGlobalAcceleratorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalAccelerator
// 修改全球加速实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  UNSUPPORTEDOPERATION_ALREADYENABLECROSSBORDER = "UnsupportedOperation.AlreadyEnableCrossBorder"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAccelerator(request *ModifyGlobalAcceleratorRequest) (response *ModifyGlobalAcceleratorResponse, err error) {
    return c.ModifyGlobalAcceleratorWithContext(context.Background(), request)
}

// ModifyGlobalAccelerator
// 修改全球加速实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTMUSTEXISTONE = "InvalidParameter.InputMustExistOne"
//  UNSUPPORTEDOPERATION_ALREADYENABLECROSSBORDER = "UnsupportedOperation.AlreadyEnableCrossBorder"
//  UNSUPPORTEDOPERATION_CROSSBORDERPROMISENOTSIGNED = "UnsupportedOperation.CrossBorderPromiseNotSigned"
//  UNSUPPORTEDOPERATION_ENABLECROSSBORDER = "UnsupportedOperation.EnableCrossBorder"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTALLOWEDOPERATE = "UnsupportedOperation.InstanceStateNotAllowedOperate"
func (c *Client) ModifyGlobalAcceleratorWithContext(ctx context.Context, request *ModifyGlobalAcceleratorRequest) (response *ModifyGlobalAcceleratorResponse, err error) {
    if request == nil {
        request = NewModifyGlobalAcceleratorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyGlobalAccelerator")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalAccelerator require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalAcceleratorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "ModifyListener")
    
    
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListener
// 修改监听器
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_DUPLICATEINSTANCESTATUS = "UnsupportedOperation.DuplicateInstanceStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    return c.ModifyListenerWithContext(context.Background(), request)
}

// ModifyListener
// 修改监听器
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APPLICATIONLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.ApplicationLayerListenerCannotCarryParameters"
//  INVALIDPARAMETER_CLIENTAFFINITYCLOSE = "InvalidParameter.ClientAffinityCLose"
//  INVALIDPARAMETER_HTTPLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.HttpListenerCannotCarryParameters"
//  INVALIDPARAMETER_TRANSPORTLAYERLISTENERCANNOTCARRYPARAMETERS = "InvalidParameter.TransportLayerListenerCannotCarryParameters"
//  INVALIDPARAMETERVALUE_INSTANCEMISMATCH = "InvalidParameterValue.InstanceMismatch"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_DUPLICATEINSTANCESTATUS = "UnsupportedOperation.DuplicateInstanceStatus"
//  UNSUPPORTEDOPERATION_INSTANCENOTRUNNING = "UnsupportedOperation.InstanceNotRunning"
func (c *Client) ModifyListenerWithContext(ctx context.Context, request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "ModifyListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}
