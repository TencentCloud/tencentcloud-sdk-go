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

package v20201028

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-28"

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


func NewAddSpecifyPrivateZoneVpcRequest() (request *AddSpecifyPrivateZoneVpcRequest) {
    request = &AddSpecifyPrivateZoneVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "AddSpecifyPrivateZoneVpc")
    
    
    return
}

func NewAddSpecifyPrivateZoneVpcResponse() (response *AddSpecifyPrivateZoneVpcResponse) {
    response = &AddSpecifyPrivateZoneVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSpecifyPrivateZoneVpc
// 追加与私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSpecifyPrivateZoneVpc(request *AddSpecifyPrivateZoneVpcRequest) (response *AddSpecifyPrivateZoneVpcResponse, err error) {
    return c.AddSpecifyPrivateZoneVpcWithContext(context.Background(), request)
}

// AddSpecifyPrivateZoneVpc
// 追加与私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddSpecifyPrivateZoneVpcWithContext(ctx context.Context, request *AddSpecifyPrivateZoneVpcRequest) (response *AddSpecifyPrivateZoneVpcResponse, err error) {
    if request == nil {
        request = NewAddSpecifyPrivateZoneVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSpecifyPrivateZoneVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSpecifyPrivateZoneVpcResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndPointRequest() (request *CreateEndPointRequest) {
    request = &CreateEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreateEndPoint")
    
    
    return
}

func NewCreateEndPointResponse() (response *CreateEndPointResponse) {
    response = &CreateEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEndPoint
// 创建终端节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEVPCENDPOINTERROR = "FailedOperation.CreateVpcEndPointError"
//  FAILEDOPERATION_CREATEVPCENDPOINTFAILED = "FailedOperation.CreateVpcEndPointFailed"
//  FAILEDOPERATION_GETTMPCREDFAILED = "FailedOperation.GetTmpCredFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTSERVICENOTEXIST = "InvalidParameter.EndPointServiceNotExist"
//  INVALIDPARAMETERVALUE_UINNOTINWHITELIST = "InvalidParameterValue.UinNotInWhiteList"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateEndPoint(request *CreateEndPointRequest) (response *CreateEndPointResponse, err error) {
    return c.CreateEndPointWithContext(context.Background(), request)
}

// CreateEndPoint
// 创建终端节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEVPCENDPOINTERROR = "FailedOperation.CreateVpcEndPointError"
//  FAILEDOPERATION_CREATEVPCENDPOINTFAILED = "FailedOperation.CreateVpcEndPointFailed"
//  FAILEDOPERATION_GETTMPCREDFAILED = "FailedOperation.GetTmpCredFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTSERVICENOTEXIST = "InvalidParameter.EndPointServiceNotExist"
//  INVALIDPARAMETERVALUE_UINNOTINWHITELIST = "InvalidParameterValue.UinNotInWhiteList"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateEndPointWithContext(ctx context.Context, request *CreateEndPointRequest) (response *CreateEndPointResponse, err error) {
    if request == nil {
        request = NewCreateEndPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEndPointAndEndPointServiceRequest() (request *CreateEndPointAndEndPointServiceRequest) {
    request = &CreateEndPointAndEndPointServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreateEndPointAndEndPointService")
    
    
    return
}

func NewCreateEndPointAndEndPointServiceResponse() (response *CreateEndPointAndEndPointServiceResponse) {
    response = &CreateEndPointAndEndPointServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEndPointAndEndPointService
// 同时创建终端节点和终端节点服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEVPCENDPOINTERROR = "FailedOperation.CreateVpcEndPointError"
//  FAILEDOPERATION_CREATEVPCENDPOINTFAILED = "FailedOperation.CreateVpcEndPointFailed"
//  FAILEDOPERATION_ENDPOINTSERVICECREATEFAILED = "FailedOperation.EndPointServiceCreateFailed"
//  FAILEDOPERATION_ENDPOINTSERVICEDELETEFAILED = "FailedOperation.EndPointServiceDeleteFailed"
//  FAILEDOPERATION_ENDPOINTSERVICEWHITELISTFAILED = "FailedOperation.EndPointServiceWhiteListFailed"
//  FAILEDOPERATION_GETTMPCREDFAILED = "FailedOperation.GetTmpCredFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTSERVICENOTEXIST = "InvalidParameter.EndPointServiceNotExist"
//  INVALIDPARAMETERVALUE_UINNOTINWHITELIST = "InvalidParameterValue.UinNotInWhiteList"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateEndPointAndEndPointService(request *CreateEndPointAndEndPointServiceRequest) (response *CreateEndPointAndEndPointServiceResponse, err error) {
    return c.CreateEndPointAndEndPointServiceWithContext(context.Background(), request)
}

// CreateEndPointAndEndPointService
// 同时创建终端节点和终端节点服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEVPCENDPOINTERROR = "FailedOperation.CreateVpcEndPointError"
//  FAILEDOPERATION_CREATEVPCENDPOINTFAILED = "FailedOperation.CreateVpcEndPointFailed"
//  FAILEDOPERATION_ENDPOINTSERVICECREATEFAILED = "FailedOperation.EndPointServiceCreateFailed"
//  FAILEDOPERATION_ENDPOINTSERVICEDELETEFAILED = "FailedOperation.EndPointServiceDeleteFailed"
//  FAILEDOPERATION_ENDPOINTSERVICEWHITELISTFAILED = "FailedOperation.EndPointServiceWhiteListFailed"
//  FAILEDOPERATION_GETTMPCREDFAILED = "FailedOperation.GetTmpCredFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTSERVICENOTEXIST = "InvalidParameter.EndPointServiceNotExist"
//  INVALIDPARAMETERVALUE_UINNOTINWHITELIST = "InvalidParameterValue.UinNotInWhiteList"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateEndPointAndEndPointServiceWithContext(ctx context.Context, request *CreateEndPointAndEndPointServiceRequest) (response *CreateEndPointAndEndPointServiceResponse, err error) {
    if request == nil {
        request = NewCreateEndPointAndEndPointServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEndPointAndEndPointService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEndPointAndEndPointServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateForwardRuleRequest() (request *CreateForwardRuleRequest) {
    request = &CreateForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreateForwardRule")
    
    
    return
}

func NewCreateForwardRuleResponse() (response *CreateForwardRuleResponse) {
    response = &CreateForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateForwardRule
// 创建自定义转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETER_FORWARDRULEZONEREPEATBIND = "InvalidParameter.ForwardRuleZoneRepeatBind"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) CreateForwardRule(request *CreateForwardRuleRequest) (response *CreateForwardRuleResponse, err error) {
    return c.CreateForwardRuleWithContext(context.Background(), request)
}

// CreateForwardRule
// 创建自定义转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETER_FORWARDRULEZONEREPEATBIND = "InvalidParameter.ForwardRuleZoneRepeatBind"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) CreateForwardRuleWithContext(ctx context.Context, request *CreateForwardRuleRequest) (response *CreateForwardRuleResponse, err error) {
    if request == nil {
        request = NewCreateForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateForwardRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateDNSAccountRequest() (request *CreatePrivateDNSAccountRequest) {
    request = &CreatePrivateDNSAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateDNSAccount")
    
    
    return
}

func NewCreatePrivateDNSAccountResponse() (response *CreatePrivateDNSAccountResponse) {
    response = &CreatePrivateDNSAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivateDNSAccount
// 创建私有域解析账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTEXIST = "InvalidParameter.AccountExist"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateDNSAccount(request *CreatePrivateDNSAccountRequest) (response *CreatePrivateDNSAccountResponse, err error) {
    return c.CreatePrivateDNSAccountWithContext(context.Background(), request)
}

// CreatePrivateDNSAccount
// 创建私有域解析账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTEXIST = "InvalidParameter.AccountExist"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateDNSAccountWithContext(ctx context.Context, request *CreatePrivateDNSAccountRequest) (response *CreatePrivateDNSAccountResponse, err error) {
    if request == nil {
        request = NewCreatePrivateDNSAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateDNSAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateDNSAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateZoneRequest() (request *CreatePrivateZoneRequest) {
    request = &CreatePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZone")
    
    
    return
}

func NewCreatePrivateZoneResponse() (response *CreatePrivateZoneResponse) {
    response = &CreatePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivateZone
// 创建私有域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEZONEFAILED = "FailedOperation.CreateZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALDOMAIN = "InvalidParameter.IllegalDomain"
//  INVALIDPARAMETER_ILLEGALDOMAINTLD = "InvalidParameter.IllegalDomainTld"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_VPCPTRZONEBINDEXCEED = "InvalidParameter.VpcPtrZoneBindExceed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESERVEDDOMAIN = "InvalidParameterValue.ReservedDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) CreatePrivateZone(request *CreatePrivateZoneRequest) (response *CreatePrivateZoneResponse, err error) {
    return c.CreatePrivateZoneWithContext(context.Background(), request)
}

// CreatePrivateZone
// 创建私有域
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEZONEFAILED = "FailedOperation.CreateZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALDOMAIN = "InvalidParameter.IllegalDomain"
//  INVALIDPARAMETER_ILLEGALDOMAINTLD = "InvalidParameter.IllegalDomainTld"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_VPCPTRZONEBINDEXCEED = "InvalidParameter.VpcPtrZoneBindExceed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESERVEDDOMAIN = "InvalidParameterValue.ReservedDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) CreatePrivateZoneWithContext(ctx context.Context, request *CreatePrivateZoneRequest) (response *CreatePrivateZoneResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateZoneRecordRequest() (request *CreatePrivateZoneRecordRequest) {
    request = &CreatePrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZoneRecord")
    
    
    return
}

func NewCreatePrivateZoneRecordResponse() (response *CreatePrivateZoneRecordResponse) {
    response = &CreatePrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivateZoneRecord
// 添加私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATERECORDFAILED = "FailedOperation.CreateRecordFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_MXNOTSUPPORTED = "InvalidParameter.MXNotSupported"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALTTLVALUE = "InvalidParameterValue.IllegalTTLValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateZoneRecord(request *CreatePrivateZoneRecordRequest) (response *CreatePrivateZoneRecordResponse, err error) {
    return c.CreatePrivateZoneRecordWithContext(context.Background(), request)
}

// CreatePrivateZoneRecord
// 添加私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATERECORDFAILED = "FailedOperation.CreateRecordFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_MXNOTSUPPORTED = "InvalidParameter.MXNotSupported"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALTTLVALUE = "InvalidParameterValue.IllegalTTLValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateZoneRecordWithContext(ctx context.Context, request *CreatePrivateZoneRecordRequest) (response *CreatePrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateZoneRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndPointRequest() (request *DeleteEndPointRequest) {
    request = &DeleteEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeleteEndPoint")
    
    
    return
}

func NewDeleteEndPointResponse() (response *DeleteEndPointResponse) {
    response = &DeleteEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEndPoint
// 删除终端节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEVPCENDPOINTFAILED = "FailedOperation.DeleteVpcEndPointFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTBINDFORWARDRULE = "InvalidParameter.EndPointBindForwardRule"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeleteEndPoint(request *DeleteEndPointRequest) (response *DeleteEndPointResponse, err error) {
    return c.DeleteEndPointWithContext(context.Background(), request)
}

// DeleteEndPoint
// 删除终端节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEVPCENDPOINTFAILED = "FailedOperation.DeleteVpcEndPointFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTBINDFORWARDRULE = "InvalidParameter.EndPointBindForwardRule"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeleteEndPointWithContext(ctx context.Context, request *DeleteEndPointRequest) (response *DeleteEndPointResponse, err error) {
    if request == nil {
        request = NewDeleteEndPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteForwardRuleRequest() (request *DeleteForwardRuleRequest) {
    request = &DeleteForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeleteForwardRule")
    
    
    return
}

func NewDeleteForwardRuleResponse() (response *DeleteForwardRuleResponse) {
    response = &DeleteForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteForwardRule
// 删除转发规则并停止转发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeleteForwardRule(request *DeleteForwardRuleRequest) (response *DeleteForwardRuleResponse, err error) {
    return c.DeleteForwardRuleWithContext(context.Background(), request)
}

// DeleteForwardRule
// 删除转发规则并停止转发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeleteForwardRuleWithContext(ctx context.Context, request *DeleteForwardRuleRequest) (response *DeleteForwardRuleResponse, err error) {
    if request == nil {
        request = NewDeleteForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteForwardRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateDNSAccountRequest() (request *DeletePrivateDNSAccountRequest) {
    request = &DeletePrivateDNSAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeletePrivateDNSAccount")
    
    
    return
}

func NewDeletePrivateDNSAccountResponse() (response *DeletePrivateDNSAccountResponse) {
    response = &DeletePrivateDNSAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrivateDNSAccount
// 删除私有域解析账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_EXISTBOUNDVPC = "UnsupportedOperation.ExistBoundVpc"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateDNSAccount(request *DeletePrivateDNSAccountRequest) (response *DeletePrivateDNSAccountResponse, err error) {
    return c.DeletePrivateDNSAccountWithContext(context.Background(), request)
}

// DeletePrivateDNSAccount
// 删除私有域解析账号
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_EXISTBOUNDVPC = "UnsupportedOperation.ExistBoundVpc"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateDNSAccountWithContext(ctx context.Context, request *DeletePrivateDNSAccountRequest) (response *DeletePrivateDNSAccountResponse, err error) {
    if request == nil {
        request = NewDeletePrivateDNSAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateDNSAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateDNSAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateZoneRequest() (request *DeletePrivateZoneRequest) {
    request = &DeletePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeletePrivateZone")
    
    
    return
}

func NewDeletePrivateZoneResponse() (response *DeletePrivateZoneResponse) {
    response = &DeletePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrivateZone
// 删除私有域并停止解析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEZONEFAILED = "FailedOperation.DeleteZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ZONEBINDFORWARDRULE = "ResourceInUse.ZoneBindForwardRule"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateZone(request *DeletePrivateZoneRequest) (response *DeletePrivateZoneResponse, err error) {
    return c.DeletePrivateZoneWithContext(context.Background(), request)
}

// DeletePrivateZone
// 删除私有域并停止解析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEZONEFAILED = "FailedOperation.DeleteZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ZONEBINDFORWARDRULE = "ResourceInUse.ZoneBindForwardRule"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateZoneWithContext(ctx context.Context, request *DeletePrivateZoneRequest) (response *DeletePrivateZoneResponse, err error) {
    if request == nil {
        request = NewDeletePrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateZoneRecordRequest() (request *DeletePrivateZoneRecordRequest) {
    request = &DeletePrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeletePrivateZoneRecord")
    
    
    return
}

func NewDeletePrivateZoneRecordResponse() (response *DeletePrivateZoneRecordResponse) {
    response = &DeletePrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrivateZoneRecord
// 删除私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETELASTBINDVPCRECORDFAILED = "FailedOperation.DeleteLastBindVpcRecordFailed"
//  FAILEDOPERATION_DELETERECORDFAILED = "FailedOperation.DeleteRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateZoneRecord(request *DeletePrivateZoneRecordRequest) (response *DeletePrivateZoneRecordResponse, err error) {
    return c.DeletePrivateZoneRecordWithContext(context.Background(), request)
}

// DeletePrivateZoneRecord
// 删除私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETELASTBINDVPCRECORDFAILED = "FailedOperation.DeleteLastBindVpcRecordFailed"
//  FAILEDOPERATION_DELETERECORDFAILED = "FailedOperation.DeleteRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DeletePrivateZoneRecordWithContext(ctx context.Context, request *DeletePrivateZoneRecordRequest) (response *DeletePrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewDeletePrivateZoneRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateZoneRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSpecifyPrivateZoneVpcRequest() (request *DeleteSpecifyPrivateZoneVpcRequest) {
    request = &DeleteSpecifyPrivateZoneVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DeleteSpecifyPrivateZoneVpc")
    
    
    return
}

func NewDeleteSpecifyPrivateZoneVpcResponse() (response *DeleteSpecifyPrivateZoneVpcResponse) {
    response = &DeleteSpecifyPrivateZoneVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSpecifyPrivateZoneVpc
// 删除与私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSpecifyPrivateZoneVpc(request *DeleteSpecifyPrivateZoneVpcRequest) (response *DeleteSpecifyPrivateZoneVpcResponse, err error) {
    return c.DeleteSpecifyPrivateZoneVpcWithContext(context.Background(), request)
}

// DeleteSpecifyPrivateZoneVpc
// 删除与私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSpecifyPrivateZoneVpcWithContext(ctx context.Context, request *DeleteSpecifyPrivateZoneVpcRequest) (response *DeleteSpecifyPrivateZoneVpcResponse, err error) {
    if request == nil {
        request = NewDeleteSpecifyPrivateZoneVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSpecifyPrivateZoneVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSpecifyPrivateZoneVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountVpcListRequest() (request *DescribeAccountVpcListRequest) {
    request = &DescribeAccountVpcListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeAccountVpcList")
    
    
    return
}

func NewDescribeAccountVpcListResponse() (response *DescribeAccountVpcListResponse) {
    response = &DescribeAccountVpcListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountVpcList
// 获取私有域解析账号的VPC列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
func (c *Client) DescribeAccountVpcList(request *DescribeAccountVpcListRequest) (response *DescribeAccountVpcListResponse, err error) {
    return c.DescribeAccountVpcListWithContext(context.Background(), request)
}

// DescribeAccountVpcList
// 获取私有域解析账号的VPC列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
func (c *Client) DescribeAccountVpcListWithContext(ctx context.Context, request *DescribeAccountVpcListRequest) (response *DescribeAccountVpcListResponse, err error) {
    if request == nil {
        request = NewDescribeAccountVpcListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountVpcList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountVpcListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogRequest() (request *DescribeAuditLogRequest) {
    request = &DescribeAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeAuditLog")
    
    
    return
}

func NewDescribeAuditLogResponse() (response *DescribeAuditLogResponse) {
    response = &DescribeAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLog
// 获取操作日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLog(request *DescribeAuditLogRequest) (response *DescribeAuditLogResponse, err error) {
    return c.DescribeAuditLogWithContext(context.Background(), request)
}

// DescribeAuditLog
// 获取操作日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLogWithContext(ctx context.Context, request *DescribeAuditLogRequest) (response *DescribeAuditLogResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
    request = &DescribeDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeDashboard")
    
    
    return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
    response = &DescribeDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDashboard
// 获取私有域解析概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    return c.DescribeDashboardWithContext(context.Background(), request)
}

// DescribeDashboard
// 获取私有域解析概览
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardWithContext(ctx context.Context, request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndPointListRequest() (request *DescribeEndPointListRequest) {
    request = &DescribeEndPointListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeEndPointList")
    
    
    return
}

func NewDescribeEndPointListResponse() (response *DescribeEndPointListResponse) {
    response = &DescribeEndPointListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEndPointList
// 获取终端节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEndPointList(request *DescribeEndPointListRequest) (response *DescribeEndPointListResponse, err error) {
    return c.DescribeEndPointListWithContext(context.Background(), request)
}

// DescribeEndPointList
// 获取终端节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEndPointListWithContext(ctx context.Context, request *DescribeEndPointListRequest) (response *DescribeEndPointListResponse, err error) {
    if request == nil {
        request = NewDescribeEndPointListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEndPointList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEndPointListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndPointRegionRequest() (request *DescribeEndPointRegionRequest) {
    request = &DescribeEndPointRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeEndPointRegion")
    
    
    return
}

func NewDescribeEndPointRegionResponse() (response *DescribeEndPointRegionResponse) {
    response = &DescribeEndPointRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEndPointRegion
// 查询终端节点已经启用了的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DescribeEndPointRegion(request *DescribeEndPointRegionRequest) (response *DescribeEndPointRegionResponse, err error) {
    return c.DescribeEndPointRegionWithContext(context.Background(), request)
}

// DescribeEndPointRegion
// 查询终端节点已经启用了的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DescribeEndPointRegionWithContext(ctx context.Context, request *DescribeEndPointRegionRequest) (response *DescribeEndPointRegionResponse, err error) {
    if request == nil {
        request = NewDescribeEndPointRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEndPointRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEndPointRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardRuleRequest() (request *DescribeForwardRuleRequest) {
    request = &DescribeForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeForwardRule")
    
    
    return
}

func NewDescribeForwardRuleResponse() (response *DescribeForwardRuleResponse) {
    response = &DescribeForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForwardRule
// 查询转发规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTBINDFORWARDRULE = "InvalidParameter.EndPointBindForwardRule"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DescribeForwardRule(request *DescribeForwardRuleRequest) (response *DescribeForwardRuleResponse, err error) {
    return c.DescribeForwardRuleWithContext(context.Background(), request)
}

// DescribeForwardRule
// 查询转发规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTBINDFORWARDRULE = "InvalidParameter.EndPointBindForwardRule"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_FREQUENCYLIMIT = "UnsupportedOperation.FrequencyLimit"
func (c *Client) DescribeForwardRuleWithContext(ctx context.Context, request *DescribeForwardRuleRequest) (response *DescribeForwardRuleResponse, err error) {
    if request == nil {
        request = NewDescribeForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForwardRuleListRequest() (request *DescribeForwardRuleListRequest) {
    request = &DescribeForwardRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeForwardRuleList")
    
    
    return
}

func NewDescribeForwardRuleListResponse() (response *DescribeForwardRuleListResponse) {
    response = &DescribeForwardRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForwardRuleList
// 查询转发规则列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeForwardRuleList(request *DescribeForwardRuleListRequest) (response *DescribeForwardRuleListResponse, err error) {
    return c.DescribeForwardRuleListWithContext(context.Background(), request)
}

// DescribeForwardRuleList
// 查询转发规则列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeForwardRuleListWithContext(ctx context.Context, request *DescribeForwardRuleListRequest) (response *DescribeForwardRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeForwardRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForwardRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForwardRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateDNSAccountListRequest() (request *DescribePrivateDNSAccountListRequest) {
    request = &DescribePrivateDNSAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateDNSAccountList")
    
    
    return
}

func NewDescribePrivateDNSAccountListResponse() (response *DescribePrivateDNSAccountListResponse) {
    response = &DescribePrivateDNSAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateDNSAccountList
// 获取私有域解析账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateDNSAccountList(request *DescribePrivateDNSAccountListRequest) (response *DescribePrivateDNSAccountListResponse, err error) {
    return c.DescribePrivateDNSAccountListWithContext(context.Background(), request)
}

// DescribePrivateDNSAccountList
// 获取私有域解析账号列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateDNSAccountListWithContext(ctx context.Context, request *DescribePrivateDNSAccountListRequest) (response *DescribePrivateDNSAccountListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateDNSAccountListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateDNSAccountList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateDNSAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneRequest() (request *DescribePrivateZoneRequest) {
    request = &DescribePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZone")
    
    
    return
}

func NewDescribePrivateZoneResponse() (response *DescribePrivateZoneResponse) {
    response = &DescribePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateZone
// 获取私有域信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZone(request *DescribePrivateZoneRequest) (response *DescribePrivateZoneResponse, err error) {
    return c.DescribePrivateZoneWithContext(context.Background(), request)
}

// DescribePrivateZone
// 获取私有域信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneWithContext(ctx context.Context, request *DescribePrivateZoneRequest) (response *DescribePrivateZoneResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneListRequest() (request *DescribePrivateZoneListRequest) {
    request = &DescribePrivateZoneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneList")
    
    
    return
}

func NewDescribePrivateZoneListResponse() (response *DescribePrivateZoneListResponse) {
    response = &DescribePrivateZoneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateZoneList
// 获取私有域列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneList(request *DescribePrivateZoneListRequest) (response *DescribePrivateZoneListResponse, err error) {
    return c.DescribePrivateZoneListWithContext(context.Background(), request)
}

// DescribePrivateZoneList
// 获取私有域列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneListWithContext(ctx context.Context, request *DescribePrivateZoneListRequest) (response *DescribePrivateZoneListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateZoneList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateZoneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneRecordListRequest() (request *DescribePrivateZoneRecordListRequest) {
    request = &DescribePrivateZoneRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneRecordList")
    
    
    return
}

func NewDescribePrivateZoneRecordListResponse() (response *DescribePrivateZoneRecordListResponse) {
    response = &DescribePrivateZoneRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateZoneRecordList
// 获取私有域记录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneRecordList(request *DescribePrivateZoneRecordListRequest) (response *DescribePrivateZoneRecordListResponse, err error) {
    return c.DescribePrivateZoneRecordListWithContext(context.Background(), request)
}

// DescribePrivateZoneRecordList
// 获取私有域记录列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneRecordListWithContext(ctx context.Context, request *DescribePrivateZoneRecordListRequest) (response *DescribePrivateZoneRecordListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneRecordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateZoneRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateZoneRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneServiceRequest() (request *DescribePrivateZoneServiceRequest) {
    request = &DescribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneService")
    
    
    return
}

func NewDescribePrivateZoneServiceResponse() (response *DescribePrivateZoneServiceResponse) {
    response = &DescribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateZoneService
// 查询私有域解析开通状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneService(request *DescribePrivateZoneServiceRequest) (response *DescribePrivateZoneServiceResponse, err error) {
    return c.DescribePrivateZoneServiceWithContext(context.Background(), request)
}

// DescribePrivateZoneService
// 查询私有域解析开通状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneServiceWithContext(ctx context.Context, request *DescribePrivateZoneServiceRequest) (response *DescribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateZoneService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaUsageRequest() (request *DescribeQuotaUsageRequest) {
    request = &DescribeQuotaUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeQuotaUsage")
    
    
    return
}

func NewDescribeQuotaUsageResponse() (response *DescribeQuotaUsageResponse) {
    response = &DescribeQuotaUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuotaUsage
// 查询额度使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaUsage(request *DescribeQuotaUsageRequest) (response *DescribeQuotaUsageResponse, err error) {
    return c.DescribeQuotaUsageWithContext(context.Background(), request)
}

// DescribeQuotaUsage
// 查询额度使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaUsageWithContext(ctx context.Context, request *DescribeQuotaUsageRequest) (response *DescribeQuotaUsageResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotaUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRequestDataRequest() (request *DescribeRequestDataRequest) {
    request = &DescribeRequestDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeRequestData")
    
    
    return
}

func NewDescribeRequestDataResponse() (response *DescribeRequestDataResponse) {
    response = &DescribeRequestDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRequestData
// 获取私有域解析请求量
//
// 可能返回的错误码:
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
func (c *Client) DescribeRequestData(request *DescribeRequestDataRequest) (response *DescribeRequestDataResponse, err error) {
    return c.DescribeRequestDataWithContext(context.Background(), request)
}

// DescribeRequestData
// 获取私有域解析请求量
//
// 可能返回的错误码:
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
func (c *Client) DescribeRequestDataWithContext(ctx context.Context, request *DescribeRequestDataRequest) (response *DescribeRequestDataResponse, err error) {
    if request == nil {
        request = NewDescribeRequestDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRequestData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRequestDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyForwardRuleRequest() (request *ModifyForwardRuleRequest) {
    request = &ModifyForwardRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyForwardRule")
    
    
    return
}

func NewModifyForwardRuleResponse() (response *ModifyForwardRuleResponse) {
    response = &ModifyForwardRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyForwardRule
// 修改转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyForwardRule(request *ModifyForwardRuleRequest) (response *ModifyForwardRuleResponse, err error) {
    return c.ModifyForwardRuleWithContext(context.Background(), request)
}

// ModifyForwardRule
// 修改转发规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINTNOTEXISTS = "InvalidParameter.EndPointNotExists"
//  INVALIDPARAMETER_FORWARDRULENOTEXIST = "InvalidParameter.ForwardRuleNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyForwardRuleWithContext(ctx context.Context, request *ModifyForwardRuleRequest) (response *ModifyForwardRuleResponse, err error) {
    if request == nil {
        request = NewModifyForwardRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyForwardRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyForwardRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRequest() (request *ModifyPrivateZoneRequest) {
    request = &ModifyPrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZone")
    
    
    return
}

func NewModifyPrivateZoneResponse() (response *ModifyPrivateZoneResponse) {
    response = &ModifyPrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateZone
// 修改私有域信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYZONEFAILED = "FailedOperation.ModifyZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) ModifyPrivateZone(request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
    return c.ModifyPrivateZoneWithContext(context.Background(), request)
}

// ModifyPrivateZone
// 修改私有域信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYZONEFAILED = "FailedOperation.ModifyZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) ModifyPrivateZoneWithContext(ctx context.Context, request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRecordRequest() (request *ModifyPrivateZoneRecordRequest) {
    request = &ModifyPrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneRecord")
    
    
    return
}

func NewModifyPrivateZoneRecordResponse() (response *ModifyPrivateZoneRecordResponse) {
    response = &ModifyPrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateZoneRecord
// 修改私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALTTLVALUE = "InvalidParameterValue.IllegalTTLValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneRecord(request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
    return c.ModifyPrivateZoneRecordWithContext(context.Background(), request)
}

// ModifyPrivateZoneRecord
// 修改私有域解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALTTLVALUE = "InvalidParameterValue.IllegalTTLValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneRecordWithContext(ctx context.Context, request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZoneRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneVpcRequest() (request *ModifyPrivateZoneVpcRequest) {
    request = &ModifyPrivateZoneVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneVpc")
    
    
    return
}

func NewModifyPrivateZoneVpcResponse() (response *ModifyPrivateZoneVpcResponse) {
    response = &ModifyPrivateZoneVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateZoneVpc
// 修改私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneVpc(request *ModifyPrivateZoneVpcRequest) (response *ModifyPrivateZoneVpcResponse, err error) {
    return c.ModifyPrivateZoneVpcWithContext(context.Background(), request)
}

// ModifyPrivateZoneVpc
// 修改私有域关联的VPC
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneVpcWithContext(ctx context.Context, request *ModifyPrivateZoneVpcRequest) (response *ModifyPrivateZoneVpcResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZoneVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneVpcResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordsStatusRequest() (request *ModifyRecordsStatusRequest) {
    request = &ModifyRecordsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyRecordsStatus")
    
    
    return
}

func NewModifyRecordsStatusResponse() (response *ModifyRecordsStatusResponse) {
    response = &ModifyRecordsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordsStatus
// 修改解析记录状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  FAILEDOPERATION_UPDATERECORDFAILED = "FailedOperation.UpdateRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
func (c *Client) ModifyRecordsStatus(request *ModifyRecordsStatusRequest) (response *ModifyRecordsStatusResponse, err error) {
    return c.ModifyRecordsStatusWithContext(context.Background(), request)
}

// ModifyRecordsStatus
// 修改解析记录状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  FAILEDOPERATION_UPDATERECORDFAILED = "FailedOperation.UpdateRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
func (c *Client) ModifyRecordsStatusWithContext(ctx context.Context, request *ModifyRecordsStatusRequest) (response *ModifyRecordsStatusResponse, err error) {
    if request == nil {
        request = NewModifyRecordsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAsyncBindVpcStatusRequest() (request *QueryAsyncBindVpcStatusRequest) {
    request = &QueryAsyncBindVpcStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "QueryAsyncBindVpcStatus")
    
    
    return
}

func NewQueryAsyncBindVpcStatusResponse() (response *QueryAsyncBindVpcStatusResponse) {
    response = &QueryAsyncBindVpcStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryAsyncBindVpcStatus
// 查询异步绑定vpc操作状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_UNIQUEIDNOTEXIST = "InvalidParameter.UniqueIdNotExist"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryAsyncBindVpcStatus(request *QueryAsyncBindVpcStatusRequest) (response *QueryAsyncBindVpcStatusResponse, err error) {
    return c.QueryAsyncBindVpcStatusWithContext(context.Background(), request)
}

// QueryAsyncBindVpcStatus
// 查询异步绑定vpc操作状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_UNIQUEIDNOTEXIST = "InvalidParameter.UniqueIdNotExist"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryAsyncBindVpcStatusWithContext(ctx context.Context, request *QueryAsyncBindVpcStatusRequest) (response *QueryAsyncBindVpcStatusResponse, err error) {
    if request == nil {
        request = NewQueryAsyncBindVpcStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAsyncBindVpcStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryAsyncBindVpcStatusResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribePrivateZoneServiceRequest() (request *SubscribePrivateZoneServiceRequest) {
    request = &SubscribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("privatedns", APIVersion, "SubscribePrivateZoneService")
    
    
    return
}

func NewSubscribePrivateZoneServiceResponse() (response *SubscribePrivateZoneServiceResponse) {
    response = &SubscribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubscribePrivateZoneService
// 开通私有域解析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubscribePrivateZoneService(request *SubscribePrivateZoneServiceRequest) (response *SubscribePrivateZoneServiceResponse, err error) {
    return c.SubscribePrivateZoneServiceWithContext(context.Background(), request)
}

// SubscribePrivateZoneService
// 开通私有域解析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubscribePrivateZoneServiceWithContext(ctx context.Context, request *SubscribePrivateZoneServiceRequest) (response *SubscribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewSubscribePrivateZoneServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubscribePrivateZoneService require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubscribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}
