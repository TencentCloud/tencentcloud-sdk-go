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

package v20190719

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
    request = &AllocateAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AllocateAddresses")
    
    
    return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
    response = &AllocateAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateAddresses
// 申请一个或多个弹性公网IP（简称 EIP）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSBINDED = "FailedOperation.PrivateIpAddressBinded"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDADDRESSCOUNT = "InvalidParameterValue.InvalidAddressCount"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  MISSINGPARAMETER_MISSINGPRIVATEIPADDRESS = "MissingParameter.MissingPrivateIpAddress"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    
    response = NewAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

// AllocateAddresses
// 申请一个或多个弹性公网IP（简称 EIP）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSBINDED = "FailedOperation.PrivateIpAddressBinded"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDADDRESSCOUNT = "InvalidParameterValue.InvalidAddressCount"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  MISSINGPARAMETER_MISSINGPRIVATEIPADDRESS = "MissingParameter.MissingPrivateIpAddress"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
func (c *Client) AllocateAddressesWithContext(ctx context.Context, request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssignIpv6AddressesRequest() (request *AssignIpv6AddressesRequest) {
    request = &AssignIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AssignIpv6Addresses")
    
    
    return
}

func NewAssignIpv6AddressesResponse() (response *AssignIpv6AddressesResponse) {
    response = &AssignIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignIpv6Addresses
// 本接口（AssignIpv6Addresses）用于弹性网卡申请IPv6地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewAssignIpv6AddressesRequest()
    }
    
    response = NewAssignIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

// AssignIpv6Addresses
// 本接口（AssignIpv6Addresses）用于弹性网卡申请IPv6地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) AssignIpv6AddressesWithContext(ctx context.Context, request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewAssignIpv6AddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssignIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssignPrivateIpAddressesRequest() (request *AssignPrivateIpAddressesRequest) {
    request = &AssignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AssignPrivateIpAddresses")
    
    
    return
}

func NewAssignPrivateIpAddressesResponse() (response *AssignPrivateIpAddressesResponse) {
    response = &AssignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignPrivateIpAddresses
// 弹性网卡申请内网 IP
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    
    response = NewAssignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

// AssignPrivateIpAddresses
// 弹性网卡申请内网 IP
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) AssignPrivateIpAddressesWithContext(ctx context.Context, request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
    request = &AssociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AssociateAddress")
    
    
    return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
    response = &AssociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateAddress
// 将弹性公网IP（简称 EIP）绑定到实例或弹性网卡的指定内网 IP 上。
//
// 将 EIP 绑定到实例（ECM）上，其本质是将 EIP 绑定到实例上主网卡的主内网 IP 上。
//
// 将 EIP 绑定到指定网卡的内网 IP上，内网IP已经绑定了EIP或普通公网IP，则反馈失败。必须先解绑该 EIP，才能再绑定新的。
//
// 只有状态为 UNBIND 的 EIP 才能够绑定内网IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGASSOCIATEENTITY = "MissingParameter.MissingAssociateEntity"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_ALREADYBINDEIP = "UnsupportedOperation.AlreadyBindEip"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDNETWORKINTERFACEIDNOTFOUND = "UnsupportedOperation.InvalidNetworkInterfaceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    
    response = NewAssociateAddressResponse()
    err = c.Send(request, response)
    return
}

// AssociateAddress
// 将弹性公网IP（简称 EIP）绑定到实例或弹性网卡的指定内网 IP 上。
//
// 将 EIP 绑定到实例（ECM）上，其本质是将 EIP 绑定到实例上主网卡的主内网 IP 上。
//
// 将 EIP 绑定到指定网卡的内网 IP上，内网IP已经绑定了EIP或普通公网IP，则反馈失败。必须先解绑该 EIP，才能再绑定新的。
//
// 只有状态为 UNBIND 的 EIP 才能够绑定内网IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGASSOCIATEENTITY = "MissingParameter.MissingAssociateEntity"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_ALREADYBINDEIP = "UnsupportedOperation.AlreadyBindEip"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDNETWORKINTERFACEIDNOTFOUND = "UnsupportedOperation.InvalidNetworkInterfaceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) AssociateAddressWithContext(ctx context.Context, request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// 绑定安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// AssociateSecurityGroups
// 绑定安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachDisks
// 本接口（AttachDisks）用于挂载云硬盘。
//
//  
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘中存在不允许挂载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISKATTACHED = "FailedOperation.DiskAttached"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  INVALIDPARAMETERVALUE_INVALIDDISKID = "InvalidParameterValue.InvalidDiskId"
//  INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_ATTACHEDDISKLIMITEXCEEDED = "LimitExceeded.AttachedDiskLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    if request == nil {
        request = NewAttachDisksRequest()
    }
    
    response = NewAttachDisksResponse()
    err = c.Send(request, response)
    return
}

// AttachDisks
// 本接口（AttachDisks）用于挂载云硬盘。
//
//  
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘中存在不允许挂载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISKATTACHED = "FailedOperation.DiskAttached"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  INVALIDPARAMETERVALUE_INVALIDDISKID = "InvalidParameterValue.InvalidDiskId"
//  INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_ATTACHEDDISKLIMITEXCEEDED = "LimitExceeded.AttachedDiskLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
func (c *Client) AttachDisksWithContext(ctx context.Context, request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    if request == nil {
        request = NewAttachDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewAttachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
    request = &AttachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "AttachNetworkInterface")
    
    
    return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
    response = &AttachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachNetworkInterface
// 弹性网卡绑定云主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    
    response = NewAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

// AttachNetworkInterface
// 弹性网卡绑定云主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeregisterTargetsRequest() (request *BatchDeregisterTargetsRequest) {
    request = &BatchDeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "BatchDeregisterTargets")
    
    
    return
}

func NewBatchDeregisterTargetsResponse() (response *BatchDeregisterTargetsResponse) {
    response = &BatchDeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeregisterTargets
// 批量解绑后端服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchDeregisterTargets(request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchDeregisterTargetsRequest()
    }
    
    response = NewBatchDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

// BatchDeregisterTargets
// 批量解绑后端服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchDeregisterTargetsWithContext(ctx context.Context, request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchDeregisterTargetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewBatchDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyTargetWeightRequest() (request *BatchModifyTargetWeightRequest) {
    request = &BatchModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "BatchModifyTargetWeight")
    
    
    return
}

func NewBatchModifyTargetWeightResponse() (response *BatchModifyTargetWeightResponse) {
    response = &BatchModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyTargetWeight
// 批量修改监听器绑定的后端机器的转发权重。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchModifyTargetWeight(request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewBatchModifyTargetWeightRequest()
    }
    
    response = NewBatchModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

// BatchModifyTargetWeight
// 批量修改监听器绑定的后端机器的转发权重。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchModifyTargetWeightWithContext(ctx context.Context, request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewBatchModifyTargetWeightRequest()
    }
    request.SetContext(ctx)
    
    response = NewBatchModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRegisterTargetsRequest() (request *BatchRegisterTargetsRequest) {
    request = &BatchRegisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "BatchRegisterTargets")
    
    
    return
}

func NewBatchRegisterTargetsResponse() (response *BatchRegisterTargetsResponse) {
    response = &BatchRegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchRegisterTargets
// 批量绑定后端目标。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    
    response = NewBatchRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

// BatchRegisterTargets
// 批量绑定后端目标。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchRegisterTargetsWithContext(ctx context.Context, request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewBatchRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
    request = &CreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateDisks")
    
    
    return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
    response = &CreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDisks
// 本接口（CreateDisks）用于创建云硬盘。
//
// 
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
//
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
//
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    if request == nil {
        request = NewCreateDisksRequest()
    }
    
    response = NewCreateDisksResponse()
    err = c.Send(request, response)
    return
}

// CreateDisks
// 本接口（CreateDisks）用于创建云硬盘。
//
// 
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
//
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
//
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) CreateDisksWithContext(ctx context.Context, request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    if request == nil {
        request = NewCreateDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHaVipRequest() (request *CreateHaVipRequest) {
    request = &CreateHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateHaVip")
    
    
    return
}

func NewCreateHaVipResponse() (response *CreateHaVipResponse) {
    response = &CreateHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHaVip
// 本接口（CreateHaVip）用于创建高可用虚拟IP（HAVIP）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTNOTCURRENTSUBNET = "InvalidParameterValue.ObjectNotCurrentSubnet"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    if request == nil {
        request = NewCreateHaVipRequest()
    }
    
    response = NewCreateHaVipResponse()
    err = c.Send(request, response)
    return
}

// CreateHaVip
// 本接口（CreateHaVip）用于创建高可用虚拟IP（HAVIP）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTNOTCURRENTSUBNET = "InvalidParameterValue.ObjectNotCurrentSubnet"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVipWithContext(ctx context.Context, request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    if request == nil {
        request = NewCreateHaVipRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateHaVipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateImage")
    
    
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImage
// 本接口(CreateImage)用于将实例的系统盘制作为新镜像，创建后的镜像可以用于创建实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

// CreateImage
// 本接口(CreateImage)用于将实例的系统盘制作为新镜像，创建后的镜像可以用于创建实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateKeyPair")
    
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKeyPair
// 用于创建一个 OpenSSH RSA 密钥对，可以用于登录 Linux 实例。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAME = "InvalidParameterValue.InvalidKeyPairName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

// CreateKeyPair
// 用于创建一个 OpenSSH RSA 密钥对，可以用于登录 Linux 实例。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAME = "InvalidParameterValue.InvalidKeyPairName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateListener
// 创建负载均衡监听器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

// CreateListener
// 创建负载均衡监听器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoadBalancer
// 购买负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERNUM = "InvalidParameterValue.InvalidLoadBalancerNum"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERTYPE = "InvalidParameterValue.InvalidLoadBalancerType"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNUMOUTOFRANGE = "InvalidParameterValue.TagNumOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBLIMITEXCEEDED = "LimitExceeded.LBLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT_LOADBALANCERSOLDOUT = "ResourcesSoldOut.LoadBalancerSoldOut"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

// CreateLoadBalancer
// 购买负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERNUM = "InvalidParameterValue.InvalidLoadBalancerNum"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERTYPE = "InvalidParameterValue.InvalidLoadBalancerType"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNUMOUTOFRANGE = "InvalidParameterValue.TagNumOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBLIMITEXCEEDED = "LimitExceeded.LBLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT_LOADBALANCERSOLDOUT = "ResourcesSoldOut.LoadBalancerSoldOut"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModuleRequest() (request *CreateModuleRequest) {
    request = &CreateModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateModule")
    
    
    return
}

func NewCreateModuleResponse() (response *CreateModuleResponse) {
    response = &CreateModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateModule
// 创建模块
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODULENUM = "InvalidParameterValue.InvaildModuleNum"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) CreateModule(request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

// CreateModule
// 创建模块
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODULENUM = "InvalidParameterValue.InvaildModuleNum"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) CreateModuleWithContext(ctx context.Context, request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
    request = &CreateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateNetworkInterface")
    
    
    return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
    response = &CreateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetworkInterface
// 创建弹性网卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    
    response = NewCreateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

// CreateNetworkInterface
// 创建弹性网卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateNetworkInterfaceWithContext(ctx context.Context, request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteTableRequest() (request *CreateRouteTableRequest) {
    request = &CreateRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateRouteTable")
    
    
    return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
    response = &CreateRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRouteTable
// 创建了VPC后，系统会创建一个默认路由表，所有新建的子网都会关联到默认路由表。默认情况下您可以直接使用默认路由表来管理您的路由策略。当您的路由策略较多时，您可以调用创建路由表接口创建更多路由表管理您的路由策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    
    response = NewCreateRouteTableResponse()
    err = c.Send(request, response)
    return
}

// CreateRouteTable
// 创建了VPC后，系统会创建一个默认路由表，所有新建的子网都会关联到默认路由表。默认情况下您可以直接使用默认路由表来管理您的路由策略。当您的路由策略较多时，您可以调用创建路由表接口创建更多路由表管理您的路由策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRouteTableWithContext(ctx context.Context, request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutesRequest() (request *CreateRoutesRequest) {
    request = &CreateRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateRoutes")
    
    
    return
}

func NewCreateRoutesResponse() (response *CreateRoutesResponse) {
    response = &CreateRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoutes
// 创建路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) CreateRoutes(request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    if request == nil {
        request = NewCreateRoutesRequest()
    }
    
    response = NewCreateRoutesResponse()
    err = c.Send(request, response)
    return
}

// CreateRoutes
// 创建路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) CreateRoutesWithContext(ctx context.Context, request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    if request == nil {
        request = NewCreateRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSecurityGroup")
    
    
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroup
// 创建安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    
    response = NewCreateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateSecurityGroup
// 创建安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupPoliciesRequest() (request *CreateSecurityGroupPoliciesRequest) {
    request = &CreateSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSecurityGroupPolicies")
    
    
    return
}

func NewCreateSecurityGroupPoliciesResponse() (response *CreateSecurityGroupPoliciesResponse) {
    response = &CreateSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroupPolicies
// <p>本接口（CreateSecurityGroupPolicies）用于创建安全组规则（SecurityGroupPolicy）。</p>
//
// <p>在 SecurityGroupPolicySet 参数中：</p>
//
// <ul>
//
// <li>Version 安全组规则版本号，用户每次更新安全规则版本会自动加1，防止您更新的路由规则已过期，不填不考虑冲突。</li>
//
// <li>在创建出站和入站规则（Egress 和 Ingress）时：<ul>
//
// <li>Protocol 字段支持输入TCP, UDP, ICMP, GRE, ALL。</li>
//
// <li>CidrBlock 字段允许输入符合cidr格式标准的任意字符串。在基础网络中，如果 CidrBlock 包含您的账户内的云服务器之外的设备在腾讯云的内网 IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。</li>
//
// <li>SecurityGroupId 字段允许输入与待修改的安全组位于相同项目中的安全组 ID，包括这个安全组 ID 本身，代表安全组下所有云服务器的内网 IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个 ID 所关联的云服务器变化而变化，不需要重新修改。</li>
//
// <li>Port 字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当 Protocol 字段是 TCP 或 UDP 时，Port 字段才被接受，即 Protocol 字段不是 TCP 或 UDP 时，Protocol 和 Port 排他关系，不允许同时输入，否则会接口报错。</li>
//
// <li>Action 字段只允许输入 ACCEPT 或 DROP。</li>
//
// <li>CidrBlock, SecurityGroupId, AddressTemplate 是排他关系，不允许同时输入，Protocol + Port 和 ServiceTemplate 二者是排他关系，不允许同时输入。</li>
//
// <li>一次请求中只能创建单个方向的规则, 如果需要指定索引（PolicyIndex）参数, 多条规则的索引必须一致。</li>
//
// </ul></li></ul>
//
// <p>默认接口请求频率限制：20次/秒。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    
    response = NewCreateSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

// CreateSecurityGroupPolicies
// <p>本接口（CreateSecurityGroupPolicies）用于创建安全组规则（SecurityGroupPolicy）。</p>
//
// <p>在 SecurityGroupPolicySet 参数中：</p>
//
// <ul>
//
// <li>Version 安全组规则版本号，用户每次更新安全规则版本会自动加1，防止您更新的路由规则已过期，不填不考虑冲突。</li>
//
// <li>在创建出站和入站规则（Egress 和 Ingress）时：<ul>
//
// <li>Protocol 字段支持输入TCP, UDP, ICMP, GRE, ALL。</li>
//
// <li>CidrBlock 字段允许输入符合cidr格式标准的任意字符串。在基础网络中，如果 CidrBlock 包含您的账户内的云服务器之外的设备在腾讯云的内网 IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。</li>
//
// <li>SecurityGroupId 字段允许输入与待修改的安全组位于相同项目中的安全组 ID，包括这个安全组 ID 本身，代表安全组下所有云服务器的内网 IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个 ID 所关联的云服务器变化而变化，不需要重新修改。</li>
//
// <li>Port 字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当 Protocol 字段是 TCP 或 UDP 时，Port 字段才被接受，即 Protocol 字段不是 TCP 或 UDP 时，Protocol 和 Port 排他关系，不允许同时输入，否则会接口报错。</li>
//
// <li>Action 字段只允许输入 ACCEPT 或 DROP。</li>
//
// <li>CidrBlock, SecurityGroupId, AddressTemplate 是排他关系，不允许同时输入，Protocol + Port 和 ServiceTemplate 二者是排他关系，不允许同时输入。</li>
//
// <li>一次请求中只能创建单个方向的规则, 如果需要指定索引（PolicyIndex）参数, 多条规则的索引必须一致。</li>
//
// </ul></li></ul>
//
// <p>默认接口请求频率限制：20次/秒。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPoliciesWithContext(ctx context.Context, request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
    request = &CreateSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSubnet")
    
    
    return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
    response = &CreateSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubnet
// 创建子网，若创建成功，则此子网会成为此可用区的默认子网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    
    response = NewCreateSubnetResponse()
    err = c.Send(request, response)
    return
}

// CreateSubnet
// 创建子网，若创建成功，则此子网会成为此可用区的默认子网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
    request = &CreateVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateVpc")
    
    
    return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
    response = &CreateVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpc
// 创建私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    
    response = NewCreateVpcResponse()
    err = c.Send(request, response)
    return
}

// CreateVpc
// 创建私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateVpcWithContext(ctx context.Context, request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHaVipRequest() (request *DeleteHaVipRequest) {
    request = &DeleteHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteHaVip")
    
    
    return
}

func NewDeleteHaVipResponse() (response *DeleteHaVipResponse) {
    response = &DeleteHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteHaVip
// 用于删除高可用虚拟IP（HAVIP）
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    if request == nil {
        request = NewDeleteHaVipRequest()
    }
    
    response = NewDeleteHaVipResponse()
    err = c.Send(request, response)
    return
}

// DeleteHaVip
// 用于删除高可用虚拟IP（HAVIP）
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVipWithContext(ctx context.Context, request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    if request == nil {
        request = NewDeleteHaVipRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteHaVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteImage")
    
    
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImage
// 删除镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

// DeleteImage
// 删除镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteListener
// 删除负载均衡监听器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

// DeleteListener
// 删除负载均衡监听器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
    request = &DeleteLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteLoadBalancer")
    
    
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancer
// 删除负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

// DeleteLoadBalancer
// 删除负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancerWithContext(ctx context.Context, request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerListenersRequest() (request *DeleteLoadBalancerListenersRequest) {
    request = &DeleteLoadBalancerListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteLoadBalancerListeners")
    
    
    return
}

func NewDeleteLoadBalancerListenersResponse() (response *DeleteLoadBalancerListenersResponse) {
    response = &DeleteLoadBalancerListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancerListeners
// 删除负载均衡多个监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancerListeners(request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerListenersRequest()
    }
    
    response = NewDeleteLoadBalancerListenersResponse()
    err = c.Send(request, response)
    return
}

// DeleteLoadBalancerListeners
// 删除负载均衡多个监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancerListenersWithContext(ctx context.Context, request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerListenersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModuleRequest() (request *DeleteModuleRequest) {
    request = &DeleteModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteModule")
    
    
    return
}

func NewDeleteModuleResponse() (response *DeleteModuleResponse) {
    response = &DeleteModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModule
// 删除业务模块
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEINMODULE = "FailedOperation.InstanceInModule"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteModule(request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    if request == nil {
        request = NewDeleteModuleRequest()
    }
    
    response = NewDeleteModuleResponse()
    err = c.Send(request, response)
    return
}

// DeleteModule
// 删除业务模块
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEINMODULE = "FailedOperation.InstanceInModule"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteModuleWithContext(ctx context.Context, request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    if request == nil {
        request = NewDeleteModuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
    request = &DeleteNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteNetworkInterface")
    
    
    return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
    response = &DeleteNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNetworkInterface
// 删除弹性网卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    
    response = NewDeleteNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

// DeleteNetworkInterface
// 删除弹性网卡
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNetworkInterfaceWithContext(ctx context.Context, request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
    request = &DeleteRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteRouteTable")
    
    
    return
}

func NewDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
    response = &DeleteRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRouteTable
// 删除路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    
    response = NewDeleteRouteTableResponse()
    err = c.Send(request, response)
    return
}

// DeleteRouteTable
// 删除路由表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
func (c *Client) DeleteRouteTableWithContext(ctx context.Context, request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
    request = &DeleteRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteRoutes")
    
    
    return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
    response = &DeleteRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoutes
// 对某个路由表批量删除路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutesRequest()
    }
    
    response = NewDeleteRoutesResponse()
    err = c.Send(request, response)
    return
}

// DeleteRoutes
// 对某个路由表批量删除路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutesWithContext(ctx context.Context, request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
    request = &DeleteSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSecurityGroup")
    
    
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroup
// 只有当前账号下的安全组允许被删除。
//
// 安全组实例ID如果在其他安全组的规则中被引用，则无法直接删除。这种情况下，需要先进行规则修改，再删除安全组。
//
// 删除的安全组无法再找回，请谨慎调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    
    response = NewDeleteSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteSecurityGroup
// 只有当前账号下的安全组允许被删除。
//
// 安全组实例ID如果在其他安全组的规则中被引用，则无法直接删除。这种情况下，需要先进行规则修改，再删除安全组。
//
// 删除的安全组无法再找回，请谨慎调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupPoliciesRequest() (request *DeleteSecurityGroupPoliciesRequest) {
    request = &DeleteSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSecurityGroupPolicies")
    
    
    return
}

func NewDeleteSecurityGroupPoliciesResponse() (response *DeleteSecurityGroupPoliciesResponse) {
    response = &DeleteSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroupPolicies
// SecurityGroupPolicySet.Version 用于指定要操作的安全组的版本。传入 Version 版本号若不等于当前安全组的最新版本，将返回失败；若不传 Version 则直接删除指定PolicyIndex的规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) DeleteSecurityGroupPolicies(request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    
    response = NewDeleteSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

// DeleteSecurityGroupPolicies
// SecurityGroupPolicySet.Version 用于指定要操作的安全组的版本。传入 Version 版本号若不等于当前安全组的最新版本，将返回失败；若不传 Version 则直接删除指定PolicyIndex的规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) DeleteSecurityGroupPoliciesWithContext(ctx context.Context, request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshots
// 本接口（DeleteSnapshots）用于删除快照。
//
// 
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以返回特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOT = "InvalidParameterValue.InvalidSnapshot"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTID = "InvalidParameterValue.InvalidSnapshotId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

// DeleteSnapshots
// 本接口（DeleteSnapshots）用于删除快照。
//
// 
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以返回特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOT = "InvalidParameterValue.InvalidSnapshot"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTID = "InvalidParameterValue.InvalidSnapshotId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshotsWithContext(ctx context.Context, request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
    request = &DeleteSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSubnet")
    
    
    return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
    response = &DeleteSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSubnet
// 删除子网，若子网为可用区下的默认子网，则默认子网会回退到系统自动创建的默认子网，非用户最新创建的子网。若默认子网不满足需求，可调用设置默认子网接口设置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

// DeleteSubnet
// 删除子网，若子网为可用区下的默认子网，则默认子网会回退到系统自动创建的默认子网，非用户最新创建的子网。若默认子网不满足需求，可调用设置默认子网接口设置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
    request = &DeleteVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteVpc")
    
    
    return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
    response = &DeleteVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpc
// 删除私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    
    response = NewDeleteVpcResponse()
    err = c.Send(request, response)
    return
}

// DeleteVpc
// 删除私有网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteVpcWithContext(ctx context.Context, request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
    request = &DescribeAddressQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeAddressQuota")
    
    
    return
}

func NewDescribeAddressQuotaResponse() (response *DescribeAddressQuotaResponse) {
    response = &DescribeAddressQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddressQuota
// 查询您账户的弹性公网IP（简称 EIP）在当前地域的配额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAddressQuotaRequest()
    }
    
    response = NewDescribeAddressQuotaResponse()
    err = c.Send(request, response)
    return
}

// DescribeAddressQuota
// 查询您账户的弹性公网IP（简称 EIP）在当前地域的配额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeAddressQuotaWithContext(ctx context.Context, request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAddressQuotaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAddressQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
    request = &DescribeAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeAddresses")
    
    
    return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
    response = &DescribeAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddresses
// 查询弹性公网IP列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    
    response = NewDescribeAddressesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAddresses
// 查询弹性公网IP列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
func (c *Client) DescribeAddressesWithContext(ctx context.Context, request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseOverviewRequest() (request *DescribeBaseOverviewRequest) {
    request = &DescribeBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeBaseOverview")
    
    
    return
}

func NewDescribeBaseOverviewResponse() (response *DescribeBaseOverviewResponse) {
    response = &DescribeBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaseOverview
// 获取概览页统计的基本数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeBaseOverview(request *DescribeBaseOverviewRequest) (response *DescribeBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBaseOverviewRequest()
    }
    
    response = NewDescribeBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

// DescribeBaseOverview
// 获取概览页统计的基本数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeBaseOverviewWithContext(ctx context.Context, request *DescribeBaseOverviewRequest) (response *DescribeBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBaseOverviewRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeConfig")
    
    
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfig
// 获取带宽硬盘等数据的限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGBASECONFIGPARAMETER = "MissingParameter.MissingBaseConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfig
// 获取带宽硬盘等数据的限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGBASECONFIGPARAMETER = "MissingParameter.MissingBaseConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeConfigWithContext(ctx context.Context, request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImageTaskRequest() (request *DescribeCustomImageTaskRequest) {
    request = &DescribeCustomImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeCustomImageTask")
    
    
    return
}

func NewDescribeCustomImageTaskResponse() (response *DescribeCustomImageTaskResponse) {
    response = &DescribeCustomImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomImageTask
// 查询导入镜像任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomImageTask(request *DescribeCustomImageTaskRequest) (response *DescribeCustomImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImageTaskRequest()
    }
    
    response = NewDescribeCustomImageTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeCustomImageTask
// 查询导入镜像任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomImageTaskWithContext(ctx context.Context, request *DescribeCustomImageTaskRequest) (response *DescribeCustomImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImageTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCustomImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultSubnetRequest() (request *DescribeDefaultSubnetRequest) {
    request = &DescribeDefaultSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeDefaultSubnet")
    
    
    return
}

func NewDescribeDefaultSubnetResponse() (response *DescribeDefaultSubnetResponse) {
    response = &DescribeDefaultSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultSubnet
// 查询可用区的默认子网
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeDefaultSubnet(request *DescribeDefaultSubnetRequest) (response *DescribeDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultSubnetRequest()
    }
    
    response = NewDescribeDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

// DescribeDefaultSubnet
// 查询可用区的默认子网
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeDefaultSubnetWithContext(ctx context.Context, request *DescribeDefaultSubnetRequest) (response *DescribeDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultSubnetRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// 
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// 
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDisksWithContext(ctx context.Context, request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
    request = &DescribeHaVipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeHaVips")
    
    
    return
}

func NewDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
    response = &DescribeHaVipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHaVips
// 用于查询高可用虚拟IP（HAVIP）列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    if request == nil {
        request = NewDescribeHaVipsRequest()
    }
    
    response = NewDescribeHaVipsResponse()
    err = c.Send(request, response)
    return
}

// DescribeHaVips
// 用于查询高可用虚拟IP（HAVIP）列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeHaVipsWithContext(ctx context.Context, request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    if request == nil {
        request = NewDescribeHaVipsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeHaVipsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRequest() (request *DescribeImageRequest) {
    request = &DescribeImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeImage")
    
    
    return
}

func NewDescribeImageResponse() (response *DescribeImageResponse) {
    response = &DescribeImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImage
// 展示镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeImage(request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    if request == nil {
        request = NewDescribeImageRequest()
    }
    
    response = NewDescribeImageResponse()
    err = c.Send(request, response)
    return
}

// DescribeImage
// 展示镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeImageWithContext(ctx context.Context, request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    if request == nil {
        request = NewDescribeImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImportImageOsRequest() (request *DescribeImportImageOsRequest) {
    request = &DescribeImportImageOsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeImportImageOs")
    
    
    return
}

func NewDescribeImportImageOsResponse() (response *DescribeImportImageOsResponse) {
    response = &DescribeImportImageOsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImportImageOs
// 查询外部导入镜像支持的OS列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
func (c *Client) DescribeImportImageOs(request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

// DescribeImportImageOs
// 查询外部导入镜像支持的OS列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
func (c *Client) DescribeImportImageOsWithContext(ctx context.Context, request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigRequest() (request *DescribeInstanceTypeConfigRequest) {
    request = &DescribeInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstanceTypeConfig")
    
    
    return
}

func NewDescribeInstanceTypeConfigResponse() (response *DescribeInstanceTypeConfigResponse) {
    response = &DescribeInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceTypeConfig
// 获取机型配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCETYPECONFIGPARAMETER = "MissingParameter.MissingInstanceTypeConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstanceTypeConfig(request *DescribeInstanceTypeConfigRequest) (response *DescribeInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigRequest()
    }
    
    response = NewDescribeInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstanceTypeConfig
// 获取机型配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCETYPECONFIGPARAMETER = "MissingParameter.MissingInstanceTypeConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstanceTypeConfigWithContext(ctx context.Context, request *DescribeInstanceTypeConfigRequest) (response *DescribeInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstanceVncUrl")
    
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceVncUrl
// 查询实例管理终端地址
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstanceVncUrl
// 查询实例管理终端地址
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 获取实例的相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDDESCRIBEINSTANCE = "InvalidParameterValue.InvaildDescribeInstance"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDORDERBYFIELD = "InvalidParameterValue.InvalidOrderByField"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstances
// 获取实例的相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDDESCRIBEINSTANCE = "InvalidParameterValue.InvaildDescribeInstance"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDORDERBYFIELD = "InvalidParameterValue.InvalidOrderByField"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstancesDeniedActions")
    
    
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDeniedActions
// 通过实例id获取当前禁止的操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstancesDeniedActions
// 通过实例id获取当前禁止的操作
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesDeniedActionsWithContext(ctx context.Context, request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListeners
// 查询负载均衡的监听器列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

// DescribeListeners
// 查询负载均衡的监听器列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalanceTaskStatusRequest() (request *DescribeLoadBalanceTaskStatusRequest) {
    request = &DescribeLoadBalanceTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeLoadBalanceTaskStatus")
    
    
    return
}

func NewDescribeLoadBalanceTaskStatusResponse() (response *DescribeLoadBalanceTaskStatusResponse) {
    response = &DescribeLoadBalanceTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalanceTaskStatus
// 查询负载均衡相关的任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalanceTaskStatus(request *DescribeLoadBalanceTaskStatusRequest) (response *DescribeLoadBalanceTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalanceTaskStatusRequest()
    }
    
    response = NewDescribeLoadBalanceTaskStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeLoadBalanceTaskStatus
// 查询负载均衡相关的任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalanceTaskStatusWithContext(ctx context.Context, request *DescribeLoadBalanceTaskStatusRequest) (response *DescribeLoadBalanceTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalanceTaskStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLoadBalanceTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeLoadBalancers")
    
    
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancers
// 查询负载均衡实例列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

// DescribeLoadBalancers
// 查询负载均衡实例列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleRequest() (request *DescribeModuleRequest) {
    request = &DescribeModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModule")
    
    
    return
}

func NewDescribeModuleResponse() (response *DescribeModuleResponse) {
    response = &DescribeModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModule
// 获取模块列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModule(request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    if request == nil {
        request = NewDescribeModuleRequest()
    }
    
    response = NewDescribeModuleResponse()
    err = c.Send(request, response)
    return
}

// DescribeModule
// 获取模块列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleWithContext(ctx context.Context, request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    if request == nil {
        request = NewDescribeModuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleDetailRequest() (request *DescribeModuleDetailRequest) {
    request = &DescribeModuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModuleDetail")
    
    
    return
}

func NewDescribeModuleDetailResponse() (response *DescribeModuleDetailResponse) {
    response = &DescribeModuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModuleDetail
// 展示模块详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleDetail(request *DescribeModuleDetailRequest) (response *DescribeModuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeModuleDetailRequest()
    }
    
    response = NewDescribeModuleDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeModuleDetail
// 展示模块详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleDetailWithContext(ctx context.Context, request *DescribeModuleDetailRequest) (response *DescribeModuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeModuleDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeModuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonthPeakNetworkRequest() (request *DescribeMonthPeakNetworkRequest) {
    request = &DescribeMonthPeakNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeMonthPeakNetwork")
    
    
    return
}

func NewDescribeMonthPeakNetworkResponse() (response *DescribeMonthPeakNetworkResponse) {
    response = &DescribeMonthPeakNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMonthPeakNetwork
// 获取客户节点上的出入带宽月峰和计费带宽信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonthPeakNetwork(request *DescribeMonthPeakNetworkRequest) (response *DescribeMonthPeakNetworkResponse, err error) {
    if request == nil {
        request = NewDescribeMonthPeakNetworkRequest()
    }
    
    response = NewDescribeMonthPeakNetworkResponse()
    err = c.Send(request, response)
    return
}

// DescribeMonthPeakNetwork
// 获取客户节点上的出入带宽月峰和计费带宽信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonthPeakNetworkWithContext(ctx context.Context, request *DescribeMonthPeakNetworkRequest) (response *DescribeMonthPeakNetworkResponse, err error) {
    if request == nil {
        request = NewDescribeMonthPeakNetworkRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMonthPeakNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
    request = &DescribeNetworkInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeNetworkInterfaces")
    
    
    return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
    response = &DescribeNetworkInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetworkInterfaces
// 查询弹性网卡列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    
    response = NewDescribeNetworkInterfacesResponse()
    err = c.Send(request, response)
    return
}

// DescribeNetworkInterfaces
// 查询弹性网卡列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeNetworkInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeRequest() (request *DescribeNodeRequest) {
    request = &DescribeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeNode")
    
    
    return
}

func NewDescribeNodeResponse() (response *DescribeNodeResponse) {
    response = &DescribeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNode
// 获取节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGNODEPARAMETER = "MissingParameter.MissingNodeParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeNode(request *DescribeNodeRequest) (response *DescribeNodeResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRequest()
    }
    
    response = NewDescribeNodeResponse()
    err = c.Send(request, response)
    return
}

// DescribeNode
// 获取节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGNODEPARAMETER = "MissingParameter.MissingNodeParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeNodeWithContext(ctx context.Context, request *DescribeNodeRequest) (response *DescribeNodeResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackingQuotaGroupRequest() (request *DescribePackingQuotaGroupRequest) {
    request = &DescribePackingQuotaGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePackingQuotaGroup")
    
    
    return
}

func NewDescribePackingQuotaGroupResponse() (response *DescribePackingQuotaGroupResponse) {
    response = &DescribePackingQuotaGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePackingQuotaGroup
// 使用本接口获取某种机型在某些区域的装箱配额（当使用虚拟机型时，返回的是一组相互关联的装箱配额）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackingQuotaGroup(request *DescribePackingQuotaGroupRequest) (response *DescribePackingQuotaGroupResponse, err error) {
    if request == nil {
        request = NewDescribePackingQuotaGroupRequest()
    }
    
    response = NewDescribePackingQuotaGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribePackingQuotaGroup
// 使用本接口获取某种机型在某些区域的装箱配额（当使用虚拟机型时，返回的是一组相互关联的装箱配额）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackingQuotaGroupWithContext(ctx context.Context, request *DescribePackingQuotaGroupRequest) (response *DescribePackingQuotaGroupResponse, err error) {
    if request == nil {
        request = NewDescribePackingQuotaGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePackingQuotaGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakBaseOverviewRequest() (request *DescribePeakBaseOverviewRequest) {
    request = &DescribePeakBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakBaseOverview")
    
    
    return
}

func NewDescribePeakBaseOverviewResponse() (response *DescribePeakBaseOverviewResponse) {
    response = &DescribePeakBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePeakBaseOverview
// CPU 内存 硬盘等基础信息峰值数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakBaseOverview(request *DescribePeakBaseOverviewRequest) (response *DescribePeakBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakBaseOverviewRequest()
    }
    
    response = NewDescribePeakBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

// DescribePeakBaseOverview
// CPU 内存 硬盘等基础信息峰值数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakBaseOverviewWithContext(ctx context.Context, request *DescribePeakBaseOverviewRequest) (response *DescribePeakBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakBaseOverviewRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePeakBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakNetworkOverviewRequest() (request *DescribePeakNetworkOverviewRequest) {
    request = &DescribePeakNetworkOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakNetworkOverview")
    
    
    return
}

func NewDescribePeakNetworkOverviewResponse() (response *DescribePeakNetworkOverviewResponse) {
    response = &DescribePeakNetworkOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePeakNetworkOverview
// 获取网络峰值数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakNetworkOverview(request *DescribePeakNetworkOverviewRequest) (response *DescribePeakNetworkOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakNetworkOverviewRequest()
    }
    
    response = NewDescribePeakNetworkOverviewResponse()
    err = c.Send(request, response)
    return
}

// DescribePeakNetworkOverview
// 获取网络峰值数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakNetworkOverviewWithContext(ctx context.Context, request *DescribePeakNetworkOverviewRequest) (response *DescribePeakNetworkOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakNetworkOverviewRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePeakNetworkOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePriceRunInstanceRequest() (request *DescribePriceRunInstanceRequest) {
    request = &DescribePriceRunInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePriceRunInstance")
    
    
    return
}

func NewDescribePriceRunInstanceResponse() (response *DescribePriceRunInstanceResponse) {
    response = &DescribePriceRunInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePriceRunInstance
// 查询实例价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePriceRunInstance(request *DescribePriceRunInstanceRequest) (response *DescribePriceRunInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePriceRunInstanceRequest()
    }
    
    response = NewDescribePriceRunInstanceResponse()
    err = c.Send(request, response)
    return
}

// DescribePriceRunInstance
// 查询实例价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePriceRunInstanceWithContext(ctx context.Context, request *DescribePriceRunInstanceRequest) (response *DescribePriceRunInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePriceRunInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePriceRunInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteConflictsRequest() (request *DescribeRouteConflictsRequest) {
    request = &DescribeRouteConflictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeRouteConflicts")
    
    
    return
}

func NewDescribeRouteConflictsResponse() (response *DescribeRouteConflictsResponse) {
    response = &DescribeRouteConflictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRouteConflicts
// 查询自定义路由策略与云联网路由策略冲突列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteConflicts(request *DescribeRouteConflictsRequest) (response *DescribeRouteConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteConflictsRequest()
    }
    
    response = NewDescribeRouteConflictsResponse()
    err = c.Send(request, response)
    return
}

// DescribeRouteConflicts
// 查询自定义路由策略与云联网路由策略冲突列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteConflictsWithContext(ctx context.Context, request *DescribeRouteConflictsRequest) (response *DescribeRouteConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteConflictsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRouteConflictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
    request = &DescribeRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeRouteTables")
    
    
    return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
    response = &DescribeRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRouteTables
// 查询路由表对象列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    
    response = NewDescribeRouteTablesResponse()
    err = c.Send(request, response)
    return
}

// DescribeRouteTables
// 查询路由表对象列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupAssociationStatisticsRequest() (request *DescribeSecurityGroupAssociationStatisticsRequest) {
    request = &DescribeSecurityGroupAssociationStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupAssociationStatistics")
    
    
    return
}

func NewDescribeSecurityGroupAssociationStatisticsResponse() (response *DescribeSecurityGroupAssociationStatisticsResponse) {
    response = &DescribeSecurityGroupAssociationStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupAssociationStatistics
// 查询安全组关联实例统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupAssociationStatistics(request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupAssociationStatisticsRequest()
    }
    
    response = NewDescribeSecurityGroupAssociationStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSecurityGroupAssociationStatistics
// 查询安全组关联实例统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupAssociationStatisticsWithContext(ctx context.Context, request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupAssociationStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupAssociationStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupLimitsRequest() (request *DescribeSecurityGroupLimitsRequest) {
    request = &DescribeSecurityGroupLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupLimits")
    
    
    return
}

func NewDescribeSecurityGroupLimitsResponse() (response *DescribeSecurityGroupLimitsResponse) {
    response = &DescribeSecurityGroupLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupLimits
// 查询用户安全组配额
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroupLimits(request *DescribeSecurityGroupLimitsRequest) (response *DescribeSecurityGroupLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupLimitsRequest()
    }
    
    response = NewDescribeSecurityGroupLimitsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSecurityGroupLimits
// 查询用户安全组配额
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroupLimitsWithContext(ctx context.Context, request *DescribeSecurityGroupLimitsRequest) (response *DescribeSecurityGroupLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupLimitsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupPoliciesRequest() (request *DescribeSecurityGroupPoliciesRequest) {
    request = &DescribeSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupPolicies")
    
    
    return
}

func NewDescribeSecurityGroupPoliciesResponse() (response *DescribeSecurityGroupPoliciesResponse) {
    response = &DescribeSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupPolicies
// 查询安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    
    response = NewDescribeSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

// DescribeSecurityGroupPolicies
// 查询安全组规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupPoliciesWithContext(ctx context.Context, request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroups")
    
    
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroups
// 查看安全组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    
    response = NewDescribeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSecurityGroups
// 查看安全组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshots
// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// 
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSnapshots
// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// 
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
    request = &DescribeSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSubnets")
    
    
    return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
    response = &DescribeSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnets
// 查询子网列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    
    response = NewDescribeSubnetsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSubnets
// 查询子网列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSubnetsWithContext(ctx context.Context, request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetHealthRequest() (request *DescribeTargetHealthRequest) {
    request = &DescribeTargetHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTargetHealth")
    
    
    return
}

func NewDescribeTargetHealthResponse() (response *DescribeTargetHealthResponse) {
    response = &DescribeTargetHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTargetHealth
// 获取负载均衡后端服务的健康检查状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    if request == nil {
        request = NewDescribeTargetHealthRequest()
    }
    
    response = NewDescribeTargetHealthResponse()
    err = c.Send(request, response)
    return
}

// DescribeTargetHealth
// 获取负载均衡后端服务的健康检查状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargetHealthWithContext(ctx context.Context, request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    if request == nil {
        request = NewDescribeTargetHealthRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTargetHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetsRequest() (request *DescribeTargetsRequest) {
    request = &DescribeTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTargets")
    
    
    return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
    response = &DescribeTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTargets
// 查询负载均衡绑定的后端服务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetsRequest()
    }
    
    response = NewDescribeTargetsResponse()
    err = c.Send(request, response)
    return
}

// DescribeTargets
// 查询负载均衡绑定的后端服务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargetsWithContext(ctx context.Context, request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTaskResult")
    
    
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskResult
// 查询EIP异步任务执行结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskResult
// 查询EIP异步任务执行结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStatus
// 本接口(DescribeTaskStatus)用于获取异步任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskStatus
// 本接口(DescribeTaskStatus)用于获取异步任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
    request = &DescribeVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeVpcs")
    
    
    return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
    response = &DescribeVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcs
// 查询私有网络列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    
    response = NewDescribeVpcsResponse()
    err = c.Send(request, response)
    return
}

// DescribeVpcs
// 查询私有网络列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
    request = &DetachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DetachDisks")
    
    
    return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
    response = &DetachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachDisks
// 本接口（DetachDisks）用于卸载云硬盘。
//
// 
//
// * 支持批量操作，卸载挂载在同一主机上的多块云盘。如果多块云盘中存在不允许卸载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机卸载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为卸载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  INVALIDPARAMETERVALUE_INVALIDDISKID = "InvalidParameterValue.InvalidDiskId"
//  INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    if request == nil {
        request = NewDetachDisksRequest()
    }
    
    response = NewDetachDisksResponse()
    err = c.Send(request, response)
    return
}

// DetachDisks
// 本接口（DetachDisks）用于卸载云硬盘。
//
// 
//
// * 支持批量操作，卸载挂载在同一主机上的多块云盘。如果多块云盘中存在不允许卸载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机卸载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为卸载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  INVALIDPARAMETERVALUE_INVALIDDISKID = "InvalidParameterValue.InvalidDiskId"
//  INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
func (c *Client) DetachDisksWithContext(ctx context.Context, request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    if request == nil {
        request = NewDetachDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
    request = &DetachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DetachNetworkInterface")
    
    
    return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
    response = &DetachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachNetworkInterface
// 弹性网卡解绑云主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    
    response = NewDetachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

// DetachNetworkInterface
// 弹性网卡解绑云主机
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRoutesRequest() (request *DisableRoutesRequest) {
    request = &DisableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DisableRoutes")
    
    
    return
}

func NewDisableRoutesResponse() (response *DisableRoutesResponse) {
    response = &DisableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableRoutes
// 禁用已启用的子网路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutes(request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    if request == nil {
        request = NewDisableRoutesRequest()
    }
    
    response = NewDisableRoutesResponse()
    err = c.Send(request, response)
    return
}

// DisableRoutes
// 禁用已启用的子网路由
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutesWithContext(ctx context.Context, request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    if request == nil {
        request = NewDisableRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
    request = &DisassociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateAddress")
    
    
    return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
    response = &DisassociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateAddress
// 解绑弹性公网IP（简称 EIP）
//
// 只有状态为 BIND 和 BIND_ENI 的 EIP 才能进行解绑定操作。
//
// EIP 如果被封堵，则不能进行解绑定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDADDRESSID = "InvalidParameterValue.InvaildAddressId"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_QUOTALIMITEXCEEDED = "UnsupportedOperation.QuotaLimitExceeded"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    
    response = NewDisassociateAddressResponse()
    err = c.Send(request, response)
    return
}

// DisassociateAddress
// 解绑弹性公网IP（简称 EIP）
//
// 只有状态为 BIND 和 BIND_ENI 的 EIP 才能进行解绑定操作。
//
// EIP 如果被封堵，则不能进行解绑定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDADDRESSID = "InvalidParameterValue.InvaildAddressId"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_QUOTALIMITEXCEEDED = "UnsupportedOperation.QuotaLimitExceeded"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) DisassociateAddressWithContext(ctx context.Context, request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateInstancesKeyPairs")
    
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateInstancesKeyPairs
// 用于解除实例的密钥绑定关系。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTSUPPORTED = "InvalidParameterValue.InstanceIdNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRID = "InvalidParameterValue.InvalidKeyPairId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateInstancesKeyPairs
// 用于解除实例的密钥绑定关系。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTSUPPORTED = "InvalidParameterValue.InstanceIdNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRID = "InvalidParameterValue.InvalidKeyPairId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateInstancesKeyPairsWithContext(ctx context.Context, request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateSecurityGroups
// 解绑安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateSecurityGroups
// 解绑安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRoutesRequest() (request *EnableRoutesRequest) {
    request = &EnableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "EnableRoutes")
    
    
    return
}

func NewEnableRoutesResponse() (response *EnableRoutesResponse) {
    response = &EnableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableRoutes
// 启用已禁用的子网路由。
//
// 本接口会校验启用后，是否与已有路由冲突，如果冲突，则无法启用，失败处理。路由冲突时，需要先禁用与之冲突的路由，才能启用该路由。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutes(request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    if request == nil {
        request = NewEnableRoutesRequest()
    }
    
    response = NewEnableRoutesResponse()
    err = c.Send(request, response)
    return
}

// EnableRoutes
// 启用已禁用的子网路由。
//
// 本接口会校验启用后，是否与已有路由冲突，如果冲突，则无法启用，失败处理。路由冲突时，需要先禁用与之冲突的路由，才能启用该路由。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutesWithContext(ctx context.Context, request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    if request == nil {
        request = NewEnableRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewImportCustomImageRequest() (request *ImportCustomImageRequest) {
    request = &ImportCustomImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ImportCustomImage")
    
    
    return
}

func NewImportCustomImageResponse() (response *ImportCustomImageResponse) {
    response = &ImportCustomImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportCustomImage
// 导入自定义镜像，支持 RAW、VHD、QCOW2、VMDK 镜像格式
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNAUTHORIZEDOPERATION_WINDOWSIMAGE = "UnauthorizedOperation.WindowsImage"
func (c *Client) ImportCustomImage(request *ImportCustomImageRequest) (response *ImportCustomImageResponse, err error) {
    if request == nil {
        request = NewImportCustomImageRequest()
    }
    
    response = NewImportCustomImageResponse()
    err = c.Send(request, response)
    return
}

// ImportCustomImage
// 导入自定义镜像，支持 RAW、VHD、QCOW2、VMDK 镜像格式
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNAUTHORIZEDOPERATION_WINDOWSIMAGE = "UnauthorizedOperation.WindowsImage"
func (c *Client) ImportCustomImageWithContext(ctx context.Context, request *ImportCustomImageRequest) (response *ImportCustomImageResponse, err error) {
    if request == nil {
        request = NewImportCustomImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewImportCustomImageResponse()
    err = c.Send(request, response)
    return
}

func NewImportImageRequest() (request *ImportImageRequest) {
    request = &ImportImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ImportImage")
    
    
    return
}

func NewImportImageResponse() (response *ImportImageResponse) {
    response = &ImportImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportImage
// 从CVM产品导入镜像到ECM
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGEDUPLICATE = "InvalidParameterValue.ImageDuplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

// ImportImage
// 从CVM产品导入镜像到ECM
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGEDUPLICATE = "InvalidParameterValue.ImageDuplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateNetworkInterfaceRequest() (request *MigrateNetworkInterfaceRequest) {
    request = &MigrateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "MigrateNetworkInterface")
    
    
    return
}

func NewMigrateNetworkInterfaceResponse() (response *MigrateNetworkInterfaceResponse) {
    response = &MigrateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigrateNetworkInterface
// 弹性网卡迁移
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    
    response = NewMigrateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

// MigrateNetworkInterface
// 弹性网卡迁移
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) MigrateNetworkInterfaceWithContext(ctx context.Context, request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewMigrateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewMigratePrivateIpAddressRequest() (request *MigratePrivateIpAddressRequest) {
    request = &MigratePrivateIpAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "MigratePrivateIpAddress")
    
    
    return
}

func NewMigratePrivateIpAddressResponse() (response *MigratePrivateIpAddressResponse) {
    response = &MigratePrivateIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigratePrivateIpAddress
// 弹性网卡内网IP迁移。
//
// 该接口用于将一个内网IP从一个弹性网卡上迁移到另外一个弹性网卡，主IP地址不支持迁移。
//
// 迁移前后的弹性网卡必须在同一个子网内。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    if request == nil {
        request = NewMigratePrivateIpAddressRequest()
    }
    
    response = NewMigratePrivateIpAddressResponse()
    err = c.Send(request, response)
    return
}

// MigratePrivateIpAddress
// 弹性网卡内网IP迁移。
//
// 该接口用于将一个内网IP从一个弹性网卡上迁移到另外一个弹性网卡，主IP地址不支持迁移。
//
// 迁移前后的弹性网卡必须在同一个子网内。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MigratePrivateIpAddressWithContext(ctx context.Context, request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    if request == nil {
        request = NewMigratePrivateIpAddressRequest()
    }
    request.SetContext(ctx)
    
    response = NewMigratePrivateIpAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressAttributeRequest() (request *ModifyAddressAttributeRequest) {
    request = &ModifyAddressAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyAddressAttribute")
    
    
    return
}

func NewModifyAddressAttributeResponse() (response *ModifyAddressAttributeResponse) {
    response = &ModifyAddressAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressAttribute
// 修改弹性公网IP属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    
    response = NewModifyAddressAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyAddressAttribute
// 修改弹性公网IP属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressAttributeWithContext(ctx context.Context, request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAddressAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressesBandwidthRequest() (request *ModifyAddressesBandwidthRequest) {
    request = &ModifyAddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyAddressesBandwidth")
    
    
    return
}

func NewModifyAddressesBandwidthResponse() (response *ModifyAddressesBandwidthResponse) {
    response = &ModifyAddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressesBandwidth
// 调整弹性公网IP带宽
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyAddressesBandwidthRequest()
    }
    
    response = NewModifyAddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

// ModifyAddressesBandwidth
// 调整弹性公网IP带宽
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressesBandwidthWithContext(ctx context.Context, request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyAddressesBandwidthRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultSubnetRequest() (request *ModifyDefaultSubnetRequest) {
    request = &ModifyDefaultSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyDefaultSubnet")
    
    
    return
}

func NewModifyDefaultSubnetResponse() (response *ModifyDefaultSubnetResponse) {
    response = &ModifyDefaultSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDefaultSubnet
// 修改在一个可用区下创建实例时使用的默认子网（创建实例时，未填写VPC参数时使用的sunbetId）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyDefaultSubnet(request *ModifyDefaultSubnetRequest) (response *ModifyDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewModifyDefaultSubnetRequest()
    }
    
    response = NewModifyDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

// ModifyDefaultSubnet
// 修改在一个可用区下创建实例时使用的默认子网（创建实例时，未填写VPC参数时使用的sunbetId）
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyDefaultSubnetWithContext(ctx context.Context, request *ModifyDefaultSubnetRequest) (response *ModifyDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewModifyDefaultSubnetRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHaVipAttributeRequest() (request *ModifyHaVipAttributeRequest) {
    request = &ModifyHaVipAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyHaVipAttribute")
    
    
    return
}

func NewModifyHaVipAttributeResponse() (response *ModifyHaVipAttributeResponse) {
    response = &ModifyHaVipAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHaVipAttribute
// 用于修改高可用虚拟IP（HAVIP）属性
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyHaVipAttribute(request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHaVipAttributeRequest()
    }
    
    response = NewModifyHaVipAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyHaVipAttribute
// 用于修改高可用虚拟IP（HAVIP）属性
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyHaVipAttributeWithContext(ctx context.Context, request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHaVipAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyHaVipAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
    request = &ModifyImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyImageAttribute")
    
    
    return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
    response = &ModifyImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyImageAttribute
// 本接口（ModifyImageAttribute）用于修改镜像属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyImageAttribute
// 本接口（ModifyImageAttribute）用于修改镜像属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyInstancesAttribute")
    
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesAttribute
// 修改实例的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyInstancesAttribute
// 修改实例的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyInstancesAttributeWithContext(ctx context.Context, request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIpv6AddressesAttributeRequest() (request *ModifyIpv6AddressesAttributeRequest) {
    request = &ModifyIpv6AddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyIpv6AddressesAttribute")
    
    
    return
}

func NewModifyIpv6AddressesAttributeResponse() (response *ModifyIpv6AddressesAttributeResponse) {
    response = &ModifyIpv6AddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIpv6AddressesAttribute
// 本接口（ModifyIpv6AddressesAttribute）用于修改弹性网卡IPv6地址属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) ModifyIpv6AddressesAttribute(request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyIpv6AddressesAttributeRequest()
    }
    
    response = NewModifyIpv6AddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyIpv6AddressesAttribute
// 本接口（ModifyIpv6AddressesAttribute）用于修改弹性网卡IPv6地址属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) ModifyIpv6AddressesAttributeWithContext(ctx context.Context, request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyIpv6AddressesAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyIpv6AddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyListener")
    
    
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyListener
// 修改负载均衡监听器属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}

// ModifyListener
// 修改负载均衡监听器属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyListenerWithContext(ctx context.Context, request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
    request = &ModifyLoadBalancerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyLoadBalancerAttributes")
    
    
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancerAttributes
// 修改负载均衡实例的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

// ModifyLoadBalancerAttributes
// 修改负载均衡实例的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyLoadBalancerAttributesWithContext(ctx context.Context, request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleConfigRequest() (request *ModifyModuleConfigRequest) {
    request = &ModifyModuleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleConfig")
    
    
    return
}

func NewModifyModuleConfigResponse() (response *ModifyModuleConfigResponse) {
    response = &ModifyModuleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleConfig
// 修改模块配置，已关联实例的模块不支持调整配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MODULENOTALLOWCHANGE = "InvalidParameterValue.ModuleNotAllowChange"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleConfig(request *ModifyModuleConfigRequest) (response *ModifyModuleConfigResponse, err error) {
    if request == nil {
        request = NewModifyModuleConfigRequest()
    }
    
    response = NewModifyModuleConfigResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleConfig
// 修改模块配置，已关联实例的模块不支持调整配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MODULENOTALLOWCHANGE = "InvalidParameterValue.ModuleNotAllowChange"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleConfigWithContext(ctx context.Context, request *ModifyModuleConfigRequest) (response *ModifyModuleConfigResponse, err error) {
    if request == nil {
        request = NewModifyModuleConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleDisableWanIpRequest() (request *ModifyModuleDisableWanIpRequest) {
    request = &ModifyModuleDisableWanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleDisableWanIp")
    
    
    return
}

func NewModifyModuleDisableWanIpResponse() (response *ModifyModuleDisableWanIpResponse) {
    response = &ModifyModuleDisableWanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleDisableWanIp
// 修改模块是否禁止分配外网ip的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleDisableWanIp(request *ModifyModuleDisableWanIpRequest) (response *ModifyModuleDisableWanIpResponse, err error) {
    if request == nil {
        request = NewModifyModuleDisableWanIpRequest()
    }
    
    response = NewModifyModuleDisableWanIpResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleDisableWanIp
// 修改模块是否禁止分配外网ip的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleDisableWanIpWithContext(ctx context.Context, request *ModifyModuleDisableWanIpRequest) (response *ModifyModuleDisableWanIpResponse, err error) {
    if request == nil {
        request = NewModifyModuleDisableWanIpRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleDisableWanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleImageRequest() (request *ModifyModuleImageRequest) {
    request = &ModifyModuleImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleImage")
    
    
    return
}

func NewModifyModuleImageResponse() (response *ModifyModuleImageResponse) {
    response = &ModifyModuleImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleImage
// 修改模块的默认镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) ModifyModuleImage(request *ModifyModuleImageRequest) (response *ModifyModuleImageResponse, err error) {
    if request == nil {
        request = NewModifyModuleImageRequest()
    }
    
    response = NewModifyModuleImageResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleImage
// 修改模块的默认镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) ModifyModuleImageWithContext(ctx context.Context, request *ModifyModuleImageRequest) (response *ModifyModuleImageResponse, err error) {
    if request == nil {
        request = NewModifyModuleImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleIpDirectRequest() (request *ModifyModuleIpDirectRequest) {
    request = &ModifyModuleIpDirectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleIpDirect")
    
    
    return
}

func NewModifyModuleIpDirectResponse() (response *ModifyModuleIpDirectResponse) {
    response = &ModifyModuleIpDirectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleIpDirect
// 修改模块IP直通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleIpDirect(request *ModifyModuleIpDirectRequest) (response *ModifyModuleIpDirectResponse, err error) {
    if request == nil {
        request = NewModifyModuleIpDirectRequest()
    }
    
    response = NewModifyModuleIpDirectResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleIpDirect
// 修改模块IP直通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleIpDirectWithContext(ctx context.Context, request *ModifyModuleIpDirectRequest) (response *ModifyModuleIpDirectResponse, err error) {
    if request == nil {
        request = NewModifyModuleIpDirectRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleIpDirectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNameRequest() (request *ModifyModuleNameRequest) {
    request = &ModifyModuleNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleName")
    
    
    return
}

func NewModifyModuleNameResponse() (response *ModifyModuleNameResponse) {
    response = &ModifyModuleNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleName
// 修改模块名称
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleName(request *ModifyModuleNameRequest) (response *ModifyModuleNameResponse, err error) {
    if request == nil {
        request = NewModifyModuleNameRequest()
    }
    
    response = NewModifyModuleNameResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleName
// 修改模块名称
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleNameWithContext(ctx context.Context, request *ModifyModuleNameRequest) (response *ModifyModuleNameResponse, err error) {
    if request == nil {
        request = NewModifyModuleNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNetworkRequest() (request *ModifyModuleNetworkRequest) {
    request = &ModifyModuleNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleNetwork")
    
    
    return
}

func NewModifyModuleNetworkResponse() (response *ModifyModuleNetworkResponse) {
    response = &ModifyModuleNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleNetwork
// 修改模块默认带宽上限
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleNetwork(request *ModifyModuleNetworkRequest) (response *ModifyModuleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyModuleNetworkRequest()
    }
    
    response = NewModifyModuleNetworkResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleNetwork
// 修改模块默认带宽上限
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleNetworkWithContext(ctx context.Context, request *ModifyModuleNetworkRequest) (response *ModifyModuleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyModuleNetworkRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleSecurityGroupsRequest() (request *ModifyModuleSecurityGroupsRequest) {
    request = &ModifyModuleSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleSecurityGroups")
    
    
    return
}

func NewModifyModuleSecurityGroupsResponse() (response *ModifyModuleSecurityGroupsResponse) {
    response = &ModifyModuleSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleSecurityGroups
// 修改模块默认安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModuleSecurityGroups(request *ModifyModuleSecurityGroupsRequest) (response *ModifyModuleSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyModuleSecurityGroupsRequest()
    }
    
    response = NewModifyModuleSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// ModifyModuleSecurityGroups
// 修改模块默认安全组
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModuleSecurityGroupsWithContext(ctx context.Context, request *ModifyModuleSecurityGroupsRequest) (response *ModifyModuleSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyModuleSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyModuleSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateIpAddressesAttributeRequest() (request *ModifyPrivateIpAddressesAttributeRequest) {
    request = &ModifyPrivateIpAddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyPrivateIpAddressesAttribute")
    
    
    return
}

func NewModifyPrivateIpAddressesAttributeResponse() (response *ModifyPrivateIpAddressesAttributeResponse) {
    response = &ModifyPrivateIpAddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrivateIpAddressesAttribute
// 用于修改弹性网卡内网IP属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    
    response = NewModifyPrivateIpAddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyPrivateIpAddressesAttribute
// 用于修改弹性网卡内网IP属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttributeWithContext(ctx context.Context, request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPrivateIpAddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteTableAttributeRequest() (request *ModifyRouteTableAttributeRequest) {
    request = &ModifyRouteTableAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyRouteTableAttribute")
    
    
    return
}

func NewModifyRouteTableAttributeResponse() (response *ModifyRouteTableAttributeResponse) {
    response = &ModifyRouteTableAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRouteTableAttribute
// 修改路由表属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    
    response = NewModifyRouteTableAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyRouteTableAttribute
// 修改路由表属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttributeWithContext(ctx context.Context, request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRouteTableAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupAttributeRequest() (request *ModifySecurityGroupAttributeRequest) {
    request = &ModifySecurityGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySecurityGroupAttribute")
    
    
    return
}

func NewModifySecurityGroupAttributeResponse() (response *ModifySecurityGroupAttributeResponse) {
    response = &ModifySecurityGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupAttribute
// 修改安全组属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAttributeRequest()
    }
    
    response = NewModifySecurityGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifySecurityGroupAttribute
// 修改安全组属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupAttributeWithContext(ctx context.Context, request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifySecurityGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupPoliciesRequest() (request *ModifySecurityGroupPoliciesRequest) {
    request = &ModifySecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySecurityGroupPolicies")
    
    
    return
}

func NewModifySecurityGroupPoliciesResponse() (response *ModifySecurityGroupPoliciesResponse) {
    response = &ModifySecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupPolicies
// 修改安全组出站和入站规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupPolicies(request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    
    response = NewModifySecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

// ModifySecurityGroupPolicies
// 修改安全组出站和入站规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupPoliciesWithContext(ctx context.Context, request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifySecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
    request = &ModifySubnetAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySubnetAttribute")
    
    
    return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
    response = &ModifySubnetAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubnetAttribute
// 修改子网属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifySubnetAttribute
// 修改子网属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifySubnetAttributeWithContext(ctx context.Context, request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
    request = &ModifyTargetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyTargetPort")
    
    
    return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
    response = &ModifyTargetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetPort
// 修改监听器绑定的后端机器的端口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetPortRequest()
    }
    
    response = NewModifyTargetPortResponse()
    err = c.Send(request, response)
    return
}

// ModifyTargetPort
// 修改监听器绑定的后端机器的端口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetPortWithContext(ctx context.Context, request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetPortRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyTargetPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetWeightRequest() (request *ModifyTargetWeightRequest) {
    request = &ModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyTargetWeight")
    
    
    return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
    response = &ModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTargetWeight
// 修改监听器绑定的后端机器的转发权重。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetWeightRequest()
    }
    
    response = NewModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

// ModifyTargetWeight
// 修改监听器绑定的后端机器的转发权重。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetWeightWithContext(ctx context.Context, request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetWeightRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
    request = &ModifyVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyVpcAttribute")
    
    
    return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
    response = &ModifyVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcAttribute
// 修改私有网络（VPC）的相关属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    
    response = NewModifyVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyVpcAttribute
// 修改私有网络（VPC）的相关属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyVpcAttributeWithContext(ctx context.Context, request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "RebootInstances")
    
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebootInstances
// 只有状态为RUNNING的实例才可以进行此操作；接口调用成功时，实例会进入REBOOTING状态；重启实例成功时，实例会进入RUNNING状态；支持强制重启，强制重启的效果等同于关闭物理计算机的电源开关再重新启动。强制重启可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常重启时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

// RebootInstances
// 只有状态为RUNNING的实例才可以进行此操作；接口调用成功时，实例会进入REBOOTING状态；重启实例成功时，实例会进入RUNNING状态；支持强制重启，强制重启的效果等同于关闭物理计算机的电源开关再重新启动。强制重启可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常重启时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
    request = &ReleaseAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ReleaseAddresses")
    
    
    return
}

func NewReleaseAddressesResponse() (response *ReleaseAddressesResponse) {
    response = &ReleaseAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseAddresses
// 释放一个或多个弹性公网IP（简称 EIP）。
//
// 该操作不可逆，释放后 EIP 关联的 IP 地址将不再属于您的名下。
//
// 只有状态为 UNBIND 的 EIP 才能进行释放操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    
    response = NewReleaseAddressesResponse()
    err = c.Send(request, response)
    return
}

// ReleaseAddresses
// 释放一个或多个弹性公网IP（简称 EIP）。
//
// 该操作不可逆，释放后 EIP 关联的 IP 地址将不再属于您的名下。
//
// 只有状态为 UNBIND 的 EIP 才能进行释放操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ReleaseAddressesWithContext(ctx context.Context, request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleaseAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIpv6AddressesRequest() (request *ReleaseIpv6AddressesRequest) {
    request = &ReleaseIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ReleaseIpv6Addresses")
    
    
    return
}

func NewReleaseIpv6AddressesResponse() (response *ReleaseIpv6AddressesResponse) {
    response = &ReleaseIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseIpv6Addresses
// 本接口（UnassignIpv6Addresses）用于释放弹性网卡IPv6地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReleaseIpv6Addresses(request *ReleaseIpv6AddressesRequest) (response *ReleaseIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewReleaseIpv6AddressesRequest()
    }
    
    response = NewReleaseIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

// ReleaseIpv6Addresses
// 本接口（UnassignIpv6Addresses）用于释放弹性网卡IPv6地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReleaseIpv6AddressesWithContext(ctx context.Context, request *ReleaseIpv6AddressesRequest) (response *ReleaseIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewReleaseIpv6AddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleaseIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePrivateIpAddressesRequest() (request *RemovePrivateIpAddressesRequest) {
    request = &RemovePrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "RemovePrivateIpAddresses")
    
    
    return
}

func NewRemovePrivateIpAddressesResponse() (response *RemovePrivateIpAddressesResponse) {
    response = &RemovePrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemovePrivateIpAddresses
// 弹性网卡退还内网 IP。
//
// 退还弹性网卡上的辅助内网IP，接口自动解关联弹性公网 IP。不能退还弹性网卡的主内网IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) RemovePrivateIpAddresses(request *RemovePrivateIpAddressesRequest) (response *RemovePrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewRemovePrivateIpAddressesRequest()
    }
    
    response = NewRemovePrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

// RemovePrivateIpAddresses
// 弹性网卡退还内网 IP。
//
// 退还弹性网卡上的辅助内网IP，接口自动解关联弹性公网 IP。不能退还弹性网卡的主内网IP。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) RemovePrivateIpAddressesWithContext(ctx context.Context, request *RemovePrivateIpAddressesRequest) (response *RemovePrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewRemovePrivateIpAddressesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRemovePrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRouteTableAssociationRequest() (request *ReplaceRouteTableAssociationRequest) {
    request = &ReplaceRouteTableAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceRouteTableAssociation")
    
    
    return
}

func NewReplaceRouteTableAssociationResponse() (response *ReplaceRouteTableAssociationResponse) {
    response = &ReplaceRouteTableAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceRouteTableAssociation
// 修改子网关联的路由表，一个子网只能关联一个路由表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTVPCNOTCURRENTVPC = "InvalidParameterValue.ObjectVpcNotCurrentVpc"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ReplaceRouteTableAssociation(request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    if request == nil {
        request = NewReplaceRouteTableAssociationRequest()
    }
    
    response = NewReplaceRouteTableAssociationResponse()
    err = c.Send(request, response)
    return
}

// ReplaceRouteTableAssociation
// 修改子网关联的路由表，一个子网只能关联一个路由表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTVPCNOTCURRENTVPC = "InvalidParameterValue.ObjectVpcNotCurrentVpc"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ReplaceRouteTableAssociationWithContext(ctx context.Context, request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    if request == nil {
        request = NewReplaceRouteTableAssociationRequest()
    }
    request.SetContext(ctx)
    
    response = NewReplaceRouteTableAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRoutesRequest() (request *ReplaceRoutesRequest) {
    request = &ReplaceRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceRoutes")
    
    
    return
}

func NewReplaceRoutesResponse() (response *ReplaceRoutesResponse) {
    response = &ReplaceRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceRoutes
// 替换路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ReplaceRoutes(request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    
    response = NewReplaceRoutesResponse()
    err = c.Send(request, response)
    return
}

// ReplaceRoutes
// 替换路由策略
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ReplaceRoutesWithContext(ctx context.Context, request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewReplaceRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceSecurityGroupPolicyRequest() (request *ReplaceSecurityGroupPolicyRequest) {
    request = &ReplaceSecurityGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceSecurityGroupPolicy")
    
    
    return
}

func NewReplaceSecurityGroupPolicyResponse() (response *ReplaceSecurityGroupPolicyResponse) {
    response = &ReplaceSecurityGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceSecurityGroupPolicy
// 替换单条安全组路由规则, 单个请求中只能替换单个方向的一条规则, 必须要指定索引（PolicyIndex）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    
    response = NewReplaceSecurityGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

// ReplaceSecurityGroupPolicy
// 替换单条安全组路由规则, 单个请求中只能替换单个方向的一条规则, 必须要指定索引（PolicyIndex）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicyWithContext(ctx context.Context, request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewReplaceSecurityGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesRequest() (request *ResetInstancesRequest) {
    request = &ResetInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstances")
    
    
    return
}

func NewResetInstancesResponse() (response *ResetInstancesResponse) {
    response = &ResetInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstances
// 重装实例，若指定了ImageId参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装；若未指定密码，则密码通过站内信形式随后发送。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstances(request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    if request == nil {
        request = NewResetInstancesRequest()
    }
    
    response = NewResetInstancesResponse()
    err = c.Send(request, response)
    return
}

// ResetInstances
// 重装实例，若指定了ImageId参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装；若未指定密码，则密码通过站内信形式随后发送。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesWithContext(ctx context.Context, request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    if request == nil {
        request = NewResetInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesMaxBandwidthRequest() (request *ResetInstancesMaxBandwidthRequest) {
    request = &ResetInstancesMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstancesMaxBandwidth")
    
    
    return
}

func NewResetInstancesMaxBandwidthResponse() (response *ResetInstancesMaxBandwidthResponse) {
    response = &ResetInstancesMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesMaxBandwidth
// 重置实例的最大带宽上限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesMaxBandwidth(request *ResetInstancesMaxBandwidthRequest) (response *ResetInstancesMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesMaxBandwidthRequest()
    }
    
    response = NewResetInstancesMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

// ResetInstancesMaxBandwidth
// 重置实例的最大带宽上限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesMaxBandwidthWithContext(ctx context.Context, request *ResetInstancesMaxBandwidthRequest) (response *ResetInstancesMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesMaxBandwidthRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesPassword
// 重置处于运行中状态的实例的密码，需要显式指定强制关机参数ForceStop。如果没有显式指定强制关机参数，则只有处于关机状态的实例才允许执行重置密码操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

// ResetInstancesPassword
// 重置处于运行中状态的实例的密码，需要显式指定强制关机参数ForceStop。如果没有显式指定强制关机参数，则只有处于关机状态的实例才允许执行重置密码操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutesRequest() (request *ResetRoutesRequest) {
    request = &ResetRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetRoutes")
    
    
    return
}

func NewResetRoutesResponse() (response *ResetRoutesResponse) {
    response = &ResetRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetRoutes
// 对某个路由表名称和所有路由策略（Route）进行重新设置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutes(request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    
    response = NewResetRoutesResponse()
    err = c.Send(request, response)
    return
}

// ResetRoutes
// 对某个路由表名称和所有路由策略（Route）进行重新设置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutesWithContext(ctx context.Context, request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "RunInstances")
    
    
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunInstances
// 创建ECM实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BLOCKBALANCE = "FailedOperation.BlockBalance"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_IMAGESIZELARGETHANSYSDISKSIZE = "InvalidParameterValue.ImageSizeLargeThanSysDiskSize"
//  INVALIDPARAMETERVALUE_INSTANCECONFIGNOTMATCH = "InvalidParameterValue.InstanceConfigNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTMATCHPID = "InvalidParameterValue.InstanceTypeNotMatchPid"
//  INVALIDPARAMETERVALUE_INVAILDHOSTNAME = "InvalidParameterValue.InvaildHostName"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDBILLINGTYPE = "InvalidParameterValue.InvalidBillingType"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKNUM = "InvalidParameterValue.InvalidDataDiskNum"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECHARGETYPE = "InvalidParameterValue.InvalidInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPEID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSECURITYGROUPID = "InvalidParameterValue.InvalidSecurityGroupID"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCETYPE = "InvalidParameterValue.InvalidZoneInstanceType"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UNMATCHEDBILLINGTYPE = "InvalidParameterValue.UnmatchedBillingType"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_INSTANCESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.InstanceSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_NICORIPLIMITEXCEEDED = "LimitExceeded.NicOrIPLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  LIMITEXCEEDED_VCPULIMITEXCEEDED = "LimitExceeded.VcpuLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCEINSUFFICIENT_INSTANCEQUOTANOTENOUGH = "ResourceInsufficient.InstanceQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

// RunInstances
// 创建ECM实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BLOCKBALANCE = "FailedOperation.BlockBalance"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_IMAGESIZELARGETHANSYSDISKSIZE = "InvalidParameterValue.ImageSizeLargeThanSysDiskSize"
//  INVALIDPARAMETERVALUE_INSTANCECONFIGNOTMATCH = "InvalidParameterValue.InstanceConfigNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTMATCHPID = "InvalidParameterValue.InstanceTypeNotMatchPid"
//  INVALIDPARAMETERVALUE_INVAILDHOSTNAME = "InvalidParameterValue.InvaildHostName"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDBILLINGTYPE = "InvalidParameterValue.InvalidBillingType"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKNUM = "InvalidParameterValue.InvalidDataDiskNum"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECHARGETYPE = "InvalidParameterValue.InvalidInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPEID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSECURITYGROUPID = "InvalidParameterValue.InvalidSecurityGroupID"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCETYPE = "InvalidParameterValue.InvalidZoneInstanceType"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UNMATCHEDBILLINGTYPE = "InvalidParameterValue.UnmatchedBillingType"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_INSTANCESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.InstanceSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_NICORIPLIMITEXCEEDED = "LimitExceeded.NicOrIPLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  LIMITEXCEEDED_VCPULIMITEXCEEDED = "LimitExceeded.VcpuLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCEINSUFFICIENT_INSTANCEQUOTANOTENOUGH = "ResourceInsufficient.InstanceQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "SetLoadBalancerSecurityGroups")
    
    
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetLoadBalancerSecurityGroups
// 设置负载均衡实例的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// SetLoadBalancerSecurityGroups
// 设置负载均衡实例的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetLoadBalancerSecurityGroupsWithContext(ctx context.Context, request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewSetSecurityGroupForLoadbalancersRequest() (request *SetSecurityGroupForLoadbalancersRequest) {
    request = &SetSecurityGroupForLoadbalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "SetSecurityGroupForLoadbalancers")
    
    
    return
}

func NewSetSecurityGroupForLoadbalancersResponse() (response *SetSecurityGroupForLoadbalancersResponse) {
    response = &SetSecurityGroupForLoadbalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetSecurityGroupForLoadbalancers
// 绑定或解绑一个安全组到多个负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}

// SetSecurityGroupForLoadbalancers
// 绑定或解绑一个安全组到多个负载均衡实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetSecurityGroupForLoadbalancersWithContext(ctx context.Context, request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    request.SetContext(ctx)
    
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstances
// 只有状态为STOPPED的实例才可以进行此操作；接口调用成功时，实例会进入STARTING状态；启动实例成功时，实例会进入RUNNING状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

// StartInstances
// 只有状态为STOPPED的实例才可以进行此操作；接口调用成功时，实例会进入STARTING状态；启动实例成功时，实例会进入RUNNING状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstances
// 只有处于"RUNNING"状态的实例才能够进行关机操作；
//
// 调用成功时，实例会进入STOPPING状态；关闭实例成功时，实例会进入STOPPED状态；
//
// 支持强制关闭，强制关机的效果等同于关闭物理计算机的电源开关，强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

// StopInstances
// 只有处于"RUNNING"状态的实例才能够进行关机操作；
//
// 调用成功时，实例会进入STOPPING状态；关闭实例成功时，实例会进入STOPPED状态；
//
// 支持强制关闭，强制关机的效果等同于关闭物理计算机的电源开关，强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
    request = &TerminateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "TerminateDisks")
    
    
    return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
    response = &TerminateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDisks
// 本接口（TerminateDisks）用于退还云硬盘。
//
// 
//
// * 不再使用的云盘，可通过本接口主动退还。
//
// * 本接口支持退还预付费云盘和按小时后付费云盘。按小时后付费云盘可直接退还，预付费云盘需符合退还规则。
//
// * 支持批量操作，每次请求批量云硬盘的上限为50。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSUFFICIENTREFUNDQUOTA = "InvalidParameterValue.InsufficientRefundQuota"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    if request == nil {
        request = NewTerminateDisksRequest()
    }
    
    response = NewTerminateDisksResponse()
    err = c.Send(request, response)
    return
}

// TerminateDisks
// 本接口（TerminateDisks）用于退还云硬盘。
//
// 
//
// * 不再使用的云盘，可通过本接口主动退还。
//
// * 本接口支持退还预付费云盘和按小时后付费云盘。按小时后付费云盘可直接退还，预付费云盘需符合退还规则。
//
// * 支持批量操作，每次请求批量云硬盘的上限为50。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSUFFICIENTREFUNDQUOTA = "InvalidParameterValue.InsufficientRefundQuota"
//  INVALIDPARAMETERVALUE_INVALIDDISK = "InvalidParameterValue.InvalidDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
func (c *Client) TerminateDisksWithContext(ctx context.Context, request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    if request == nil {
        request = NewTerminateDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewTerminateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateInstances
// 销毁实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TERMINATETIMESMALLER = "InvalidParameterValue.TerminateTimeSmaller"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}

// TerminateInstances
// 销毁实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TERMINATETIMESMALLER = "InvalidParameterValue.TerminateTimeSmaller"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
