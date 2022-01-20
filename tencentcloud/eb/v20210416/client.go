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

package v20210416

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-16"

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


func NewCheckRuleRequest() (request *CheckRuleRequest) {
    request = &CheckRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CheckRule")
    
    
    return
}

func NewCheckRuleResponse() (response *CheckRuleResponse) {
    response = &CheckRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckRule
// 规则检验
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckRule(request *CheckRuleRequest) (response *CheckRuleResponse, err error) {
    if request == nil {
        request = NewCheckRuleRequest()
    }
    
    response = NewCheckRuleResponse()
    err = c.Send(request, response)
    return
}

// CheckRule
// 规则检验
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckRuleWithContext(ctx context.Context, request *CheckRuleRequest) (response *CheckRuleResponse, err error) {
    if request == nil {
        request = NewCheckRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCheckRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCheckTransformationRequest() (request *CheckTransformationRequest) {
    request = &CheckTransformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CheckTransformation")
    
    
    return
}

func NewCheckTransformationResponse() (response *CheckTransformationResponse) {
    response = &CheckTransformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckTransformation
// 用于在ETL配置页面, 测试规则和数据.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckTransformation(request *CheckTransformationRequest) (response *CheckTransformationResponse, err error) {
    if request == nil {
        request = NewCheckTransformationRequest()
    }
    
    response = NewCheckTransformationResponse()
    err = c.Send(request, response)
    return
}

// CheckTransformation
// 用于在ETL配置页面, 测试规则和数据.
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckTransformationWithContext(ctx context.Context, request *CheckTransformationRequest) (response *CheckTransformationResponse, err error) {
    if request == nil {
        request = NewCheckTransformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewCheckTransformationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConnectionRequest() (request *CreateConnectionRequest) {
    request = &CreateConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CreateConnection")
    
    
    return
}

func NewCreateConnectionResponse() (response *CreateConnectionResponse) {
    response = &CreateConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConnection
// 创建事件连接器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_CONNECTIONDESCRIPTION = "InvalidParameterValue.ConnectionDescription"
//  INVALIDPARAMETERVALUE_CONNECTIONNAME = "InvalidParameterValue.ConnectionName"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_CONNECTION = "LimitExceeded.Connection"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCEUNAVAILABLE_CONNECTION = "ResourceUnavailable.Connection"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) CreateConnection(request *CreateConnectionRequest) (response *CreateConnectionResponse, err error) {
    if request == nil {
        request = NewCreateConnectionRequest()
    }
    
    response = NewCreateConnectionResponse()
    err = c.Send(request, response)
    return
}

// CreateConnection
// 创建事件连接器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_CONNECTIONDESCRIPTION = "InvalidParameterValue.ConnectionDescription"
//  INVALIDPARAMETERVALUE_CONNECTIONNAME = "InvalidParameterValue.ConnectionName"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_CONNECTION = "LimitExceeded.Connection"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCEUNAVAILABLE_CONNECTION = "ResourceUnavailable.Connection"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) CreateConnectionWithContext(ctx context.Context, request *CreateConnectionRequest) (response *CreateConnectionResponse, err error) {
    if request == nil {
        request = NewCreateConnectionRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEventBusRequest() (request *CreateEventBusRequest) {
    request = &CreateEventBusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CreateEventBus")
    
    
    return
}

func NewCreateEventBusResponse() (response *CreateEventBusResponse) {
    response = &CreateEventBusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEventBus
// 用于创建事件集
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_EVENTBUS = "LimitExceeded.EventBus"
//  RESOURCEINUSE_DEFAULTEVENTBUS = "ResourceInUse.DefaultEventBus"
//  RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"
func (c *Client) CreateEventBus(request *CreateEventBusRequest) (response *CreateEventBusResponse, err error) {
    if request == nil {
        request = NewCreateEventBusRequest()
    }
    
    response = NewCreateEventBusResponse()
    err = c.Send(request, response)
    return
}

// CreateEventBus
// 用于创建事件集
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_EVENTBUS = "LimitExceeded.EventBus"
//  RESOURCEINUSE_DEFAULTEVENTBUS = "ResourceInUse.DefaultEventBus"
//  RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"
func (c *Client) CreateEventBusWithContext(ctx context.Context, request *CreateEventBusRequest) (response *CreateEventBusResponse, err error) {
    if request == nil {
        request = NewCreateEventBusRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateEventBusResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// 创建事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"
//  INVALIDPARAMETERVALUE_RULENAME = "InvalidParameterValue.RuleName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_RULE = "LimitExceeded.Rule"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

// CreateRule
// 创建事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"
//  INVALIDPARAMETERVALUE_RULENAME = "InvalidParameterValue.RuleName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_RULE = "LimitExceeded.Rule"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetRequest() (request *CreateTargetRequest) {
    request = &CreateTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CreateTarget")
    
    
    return
}

func NewCreateTargetResponse() (response *CreateTargetResponse) {
    response = &CreateTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTarget
// 创建事件目标
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"
//  INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKTYPE = "InvalidParameterValue.CallbackType"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERCHANNEL = "InvalidParameterValue.NoticeReceiverChannel"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERTIMEWINDOW = "InvalidParameterValue.NoticeReceiverTimeWindow"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERTYPE = "InvalidParameterValue.NoticeReceiverUserType"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TARGET = "LimitExceeded.Target"
//  LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  RESOURCEUNAVAILABLE_TARGET = "ResourceUnavailable.Target"
func (c *Client) CreateTarget(request *CreateTargetRequest) (response *CreateTargetResponse, err error) {
    if request == nil {
        request = NewCreateTargetRequest()
    }
    
    response = NewCreateTargetResponse()
    err = c.Send(request, response)
    return
}

// CreateTarget
// 创建事件目标
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"
//  INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKTYPE = "InvalidParameterValue.CallbackType"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERCHANNEL = "InvalidParameterValue.NoticeReceiverChannel"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERTIMEWINDOW = "InvalidParameterValue.NoticeReceiverTimeWindow"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERTYPE = "InvalidParameterValue.NoticeReceiverUserType"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TARGET = "LimitExceeded.Target"
//  LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  RESOURCEUNAVAILABLE_TARGET = "ResourceUnavailable.Target"
func (c *Client) CreateTargetWithContext(ctx context.Context, request *CreateTargetRequest) (response *CreateTargetResponse, err error) {
    if request == nil {
        request = NewCreateTargetRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTargetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTransformationRequest() (request *CreateTransformationRequest) {
    request = &CreateTransformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "CreateTransformation")
    
    
    return
}

func NewCreateTransformationResponse() (response *CreateTransformationResponse) {
    response = &CreateTransformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTransformation
// 用于创建转换器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"
//  INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKTYPE = "InvalidParameterValue.CallbackType"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERCHANNEL = "InvalidParameterValue.NoticeReceiverChannel"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERTIMEWINDOW = "InvalidParameterValue.NoticeReceiverTimeWindow"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERTYPE = "InvalidParameterValue.NoticeReceiverUserType"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TARGET = "LimitExceeded.Target"
//  LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  RESOURCEUNAVAILABLE_TARGET = "ResourceUnavailable.Target"
func (c *Client) CreateTransformation(request *CreateTransformationRequest) (response *CreateTransformationResponse, err error) {
    if request == nil {
        request = NewCreateTransformationRequest()
    }
    
    response = NewCreateTransformationResponse()
    err = c.Send(request, response)
    return
}

// CreateTransformation
// 用于创建转换器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"
//  INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"
//  INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKTYPE = "InvalidParameterValue.CallbackType"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERCHANNEL = "InvalidParameterValue.NoticeReceiverChannel"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERTIMEWINDOW = "InvalidParameterValue.NoticeReceiverTimeWindow"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERTYPE = "InvalidParameterValue.NoticeReceiverUserType"
//  INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TARGET = "LimitExceeded.Target"
//  LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  RESOURCEUNAVAILABLE_TARGET = "ResourceUnavailable.Target"
func (c *Client) CreateTransformationWithContext(ctx context.Context, request *CreateTransformationRequest) (response *CreateTransformationResponse, err error) {
    if request == nil {
        request = NewCreateTransformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTransformationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConnectionRequest() (request *DeleteConnectionRequest) {
    request = &DeleteConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "DeleteConnection")
    
    
    return
}

func NewDeleteConnectionResponse() (response *DeleteConnectionResponse) {
    response = &DeleteConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteConnection
// 删除事件连接器
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETECONNECTION = "FailedOperation.DeleteConnection"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_CONNECTIONDESCRIPTION = "InvalidParameterValue.ConnectionDescription"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DeleteConnection(request *DeleteConnectionRequest) (response *DeleteConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteConnectionRequest()
    }
    
    response = NewDeleteConnectionResponse()
    err = c.Send(request, response)
    return
}

// DeleteConnection
// 删除事件连接器
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETECONNECTION = "FailedOperation.DeleteConnection"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_CONNECTIONDESCRIPTION = "InvalidParameterValue.ConnectionDescription"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DeleteConnectionWithContext(ctx context.Context, request *DeleteConnectionRequest) (response *DeleteConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteConnectionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEventBusRequest() (request *DeleteEventBusRequest) {
    request = &DeleteEventBusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "DeleteEventBus")
    
    
    return
}

func NewDeleteEventBusResponse() (response *DeleteEventBusResponse) {
    response = &DeleteEventBusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEventBus
// 删除事件集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) DeleteEventBus(request *DeleteEventBusRequest) (response *DeleteEventBusResponse, err error) {
    if request == nil {
        request = NewDeleteEventBusRequest()
    }
    
    response = NewDeleteEventBusResponse()
    err = c.Send(request, response)
    return
}

// DeleteEventBus
// 删除事件集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) DeleteEventBusWithContext(ctx context.Context, request *DeleteEventBusRequest) (response *DeleteEventBusResponse, err error) {
    if request == nil {
        request = NewDeleteEventBusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEventBusResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "DeleteRule")
    
    
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRule
// 删除事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  RESOURCEINUSE_RULE = "ResourceInUse.Rule"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

// DeleteRule
// 删除事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  RESOURCEINUSE_RULE = "ResourceInUse.Rule"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetRequest() (request *DeleteTargetRequest) {
    request = &DeleteTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "DeleteTarget")
    
    
    return
}

func NewDeleteTargetResponse() (response *DeleteTargetResponse) {
    response = &DeleteTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTarget
// 删除事件目标
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) DeleteTarget(request *DeleteTargetRequest) (response *DeleteTargetResponse, err error) {
    if request == nil {
        request = NewDeleteTargetRequest()
    }
    
    response = NewDeleteTargetResponse()
    err = c.Send(request, response)
    return
}

// DeleteTarget
// 删除事件目标
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) DeleteTargetWithContext(ctx context.Context, request *DeleteTargetRequest) (response *DeleteTargetResponse, err error) {
    if request == nil {
        request = NewDeleteTargetRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteTargetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTransformationRequest() (request *DeleteTransformationRequest) {
    request = &DeleteTransformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "DeleteTransformation")
    
    
    return
}

func NewDeleteTransformationResponse() (response *DeleteTransformationResponse) {
    response = &DeleteTransformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTransformation
// 用于删除转换器
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) DeleteTransformation(request *DeleteTransformationRequest) (response *DeleteTransformationResponse, err error) {
    if request == nil {
        request = NewDeleteTransformationRequest()
    }
    
    response = NewDeleteTransformationResponse()
    err = c.Send(request, response)
    return
}

// DeleteTransformation
// 用于删除转换器
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) DeleteTransformationWithContext(ctx context.Context, request *DeleteTransformationRequest) (response *DeleteTransformationResponse, err error) {
    if request == nil {
        request = NewDeleteTransformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteTransformationResponse()
    err = c.Send(request, response)
    return
}

func NewGetEventBusRequest() (request *GetEventBusRequest) {
    request = &GetEventBusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "GetEventBus")
    
    
    return
}

func NewGetEventBusResponse() (response *GetEventBusResponse) {
    response = &GetEventBusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEventBus
// 获取事件集详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) GetEventBus(request *GetEventBusRequest) (response *GetEventBusResponse, err error) {
    if request == nil {
        request = NewGetEventBusRequest()
    }
    
    response = NewGetEventBusResponse()
    err = c.Send(request, response)
    return
}

// GetEventBus
// 获取事件集详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) GetEventBusWithContext(ctx context.Context, request *GetEventBusRequest) (response *GetEventBusResponse, err error) {
    if request == nil {
        request = NewGetEventBusRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetEventBusResponse()
    err = c.Send(request, response)
    return
}

func NewGetRuleRequest() (request *GetRuleRequest) {
    request = &GetRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "GetRule")
    
    
    return
}

func NewGetRuleResponse() (response *GetRuleResponse) {
    response = &GetRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRule
// 获取事件规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) GetRule(request *GetRuleRequest) (response *GetRuleResponse, err error) {
    if request == nil {
        request = NewGetRuleRequest()
    }
    
    response = NewGetRuleResponse()
    err = c.Send(request, response)
    return
}

// GetRule
// 获取事件规则详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) GetRuleWithContext(ctx context.Context, request *GetRuleRequest) (response *GetRuleResponse, err error) {
    if request == nil {
        request = NewGetRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransformationRequest() (request *GetTransformationRequest) {
    request = &GetTransformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "GetTransformation")
    
    
    return
}

func NewGetTransformationResponse() (response *GetTransformationResponse) {
    response = &GetTransformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTransformation
// 用于获取转换器详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
func (c *Client) GetTransformation(request *GetTransformationRequest) (response *GetTransformationResponse, err error) {
    if request == nil {
        request = NewGetTransformationRequest()
    }
    
    response = NewGetTransformationResponse()
    err = c.Send(request, response)
    return
}

// GetTransformation
// 用于获取转换器详情
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
func (c *Client) GetTransformationWithContext(ctx context.Context, request *GetTransformationRequest) (response *GetTransformationResponse, err error) {
    if request == nil {
        request = NewGetTransformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetTransformationResponse()
    err = c.Send(request, response)
    return
}

func NewListConnectionsRequest() (request *ListConnectionsRequest) {
    request = &ListConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "ListConnections")
    
    
    return
}

func NewListConnectionsResponse() (response *ListConnectionsResponse) {
    response = &ListConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListConnections
// 获取事件连接器列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) ListConnections(request *ListConnectionsRequest) (response *ListConnectionsResponse, err error) {
    if request == nil {
        request = NewListConnectionsRequest()
    }
    
    response = NewListConnectionsResponse()
    err = c.Send(request, response)
    return
}

// ListConnections
// 获取事件连接器列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) ListConnectionsWithContext(ctx context.Context, request *ListConnectionsRequest) (response *ListConnectionsResponse, err error) {
    if request == nil {
        request = NewListConnectionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewListConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewListEventBusesRequest() (request *ListEventBusesRequest) {
    request = &ListEventBusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "ListEventBuses")
    
    
    return
}

func NewListEventBusesResponse() (response *ListEventBusesResponse) {
    response = &ListEventBusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEventBuses
// 获取事件集列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListEventBuses(request *ListEventBusesRequest) (response *ListEventBusesResponse, err error) {
    if request == nil {
        request = NewListEventBusesRequest()
    }
    
    response = NewListEventBusesResponse()
    err = c.Send(request, response)
    return
}

// ListEventBuses
// 获取事件集列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"
func (c *Client) ListEventBusesWithContext(ctx context.Context, request *ListEventBusesRequest) (response *ListEventBusesResponse, err error) {
    if request == nil {
        request = NewListEventBusesRequest()
    }
    request.SetContext(ctx)
    
    response = NewListEventBusesResponse()
    err = c.Send(request, response)
    return
}

func NewListRulesRequest() (request *ListRulesRequest) {
    request = &ListRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "ListRules")
    
    
    return
}

func NewListRulesResponse() (response *ListRulesResponse) {
    response = &ListRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListRules
// 获取事件规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) ListRules(request *ListRulesRequest) (response *ListRulesResponse, err error) {
    if request == nil {
        request = NewListRulesRequest()
    }
    
    response = NewListRulesResponse()
    err = c.Send(request, response)
    return
}

// ListRules
// 获取事件规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) ListRulesWithContext(ctx context.Context, request *ListRulesRequest) (response *ListRulesResponse, err error) {
    if request == nil {
        request = NewListRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewListRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListTargetsRequest() (request *ListTargetsRequest) {
    request = &ListTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "ListTargets")
    
    
    return
}

func NewListTargetsResponse() (response *ListTargetsResponse) {
    response = &ListTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTargets
// 获取事件目标列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) ListTargets(request *ListTargetsRequest) (response *ListTargetsResponse, err error) {
    if request == nil {
        request = NewListTargetsRequest()
    }
    
    response = NewListTargetsResponse()
    err = c.Send(request, response)
    return
}

// ListTargets
// 获取事件目标列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"
//  INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"
//  INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
func (c *Client) ListTargetsWithContext(ctx context.Context, request *ListTargetsRequest) (response *ListTargetsResponse, err error) {
    if request == nil {
        request = NewListTargetsRequest()
    }
    request.SetContext(ctx)
    
    response = NewListTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewPublishEventRequest() (request *PublishEventRequest) {
    request = &PublishEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "PublishEvent")
    
    
    return
}

func NewPublishEventResponse() (response *PublishEventResponse) {
    response = &PublishEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishEvent
// （已废弃）用于Event事件投递
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) PublishEvent(request *PublishEventRequest) (response *PublishEventResponse, err error) {
    if request == nil {
        request = NewPublishEventRequest()
    }
    
    response = NewPublishEventResponse()
    err = c.Send(request, response)
    return
}

// PublishEvent
// （已废弃）用于Event事件投递
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) PublishEventWithContext(ctx context.Context, request *PublishEventRequest) (response *PublishEventResponse, err error) {
    if request == nil {
        request = NewPublishEventRequest()
    }
    request.SetContext(ctx)
    
    response = NewPublishEventResponse()
    err = c.Send(request, response)
    return
}

func NewPutEventsRequest() (request *PutEventsRequest) {
    request = &PutEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "PutEvents")
    
    
    return
}

func NewPutEventsResponse() (response *PutEventsResponse) {
    response = &PutEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutEvents
// 用于Event事件投递
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) PutEvents(request *PutEventsRequest) (response *PutEventsResponse, err error) {
    if request == nil {
        request = NewPutEventsRequest()
    }
    
    response = NewPutEventsResponse()
    err = c.Send(request, response)
    return
}

// PutEvents
// 用于Event事件投递
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) PutEventsWithContext(ctx context.Context, request *PutEventsRequest) (response *PutEventsResponse, err error) {
    if request == nil {
        request = NewPutEventsRequest()
    }
    request.SetContext(ctx)
    
    response = NewPutEventsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConnectionRequest() (request *UpdateConnectionRequest) {
    request = &UpdateConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "UpdateConnection")
    
    
    return
}

func NewUpdateConnectionResponse() (response *UpdateConnectionResponse) {
    response = &UpdateConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateConnection
// 更新事件连接器
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONNECTIONID = "InvalidParameterValue.ConnectionId"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"
func (c *Client) UpdateConnection(request *UpdateConnectionRequest) (response *UpdateConnectionResponse, err error) {
    if request == nil {
        request = NewUpdateConnectionRequest()
    }
    
    response = NewUpdateConnectionResponse()
    err = c.Send(request, response)
    return
}

// UpdateConnection
// 更新事件连接器
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONNECTIONID = "InvalidParameterValue.ConnectionId"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"
func (c *Client) UpdateConnectionWithContext(ctx context.Context, request *UpdateConnectionRequest) (response *UpdateConnectionResponse, err error) {
    if request == nil {
        request = NewUpdateConnectionRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEventBusRequest() (request *UpdateEventBusRequest) {
    request = &UpdateEventBusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "UpdateEventBus")
    
    
    return
}

func NewUpdateEventBusResponse() (response *UpdateEventBusResponse) {
    response = &UpdateEventBusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateEventBus
// 更新事件集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) UpdateEventBus(request *UpdateEventBusRequest) (response *UpdateEventBusResponse, err error) {
    if request == nil {
        request = NewUpdateEventBusRequest()
    }
    
    response = NewUpdateEventBusResponse()
    err = c.Send(request, response)
    return
}

// UpdateEventBus
// 更新事件集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"
//  OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
func (c *Client) UpdateEventBusWithContext(ctx context.Context, request *UpdateEventBusRequest) (response *UpdateEventBusResponse, err error) {
    if request == nil {
        request = NewUpdateEventBusRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateEventBusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRuleRequest() (request *UpdateRuleRequest) {
    request = &UpdateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "UpdateRule")
    
    
    return
}

func NewUpdateRuleResponse() (response *UpdateRuleResponse) {
    response = &UpdateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRule
// 更新事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_RULENAME = "InvalidParameterValue.RuleName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) UpdateRule(request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRuleRequest()
    }
    
    response = NewUpdateRuleResponse()
    err = c.Send(request, response)
    return
}

// UpdateRule
// 更新事件规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"
//  INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_RULENAME = "InvalidParameterValue.RuleName"
//  INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"
func (c *Client) UpdateRuleWithContext(ctx context.Context, request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTargetRequest() (request *UpdateTargetRequest) {
    request = &UpdateTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "UpdateTarget")
    
    
    return
}

func NewUpdateTargetResponse() (response *UpdateTargetResponse) {
    response = &UpdateTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateTarget
// 更新事件目标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) UpdateTarget(request *UpdateTargetRequest) (response *UpdateTargetResponse, err error) {
    if request == nil {
        request = NewUpdateTargetRequest()
    }
    
    response = NewUpdateTargetResponse()
    err = c.Send(request, response)
    return
}

// UpdateTarget
// 更新事件目标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) UpdateTargetWithContext(ctx context.Context, request *UpdateTargetRequest) (response *UpdateTargetResponse, err error) {
    if request == nil {
        request = NewUpdateTargetRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateTargetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTransformationRequest() (request *UpdateTransformationRequest) {
    request = &UpdateTransformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("eb", APIVersion, "UpdateTransformation")
    
    
    return
}

func NewUpdateTransformationResponse() (response *UpdateTransformationResponse) {
    response = &UpdateTransformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateTransformation
// 用于更新转换器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) UpdateTransformation(request *UpdateTransformationRequest) (response *UpdateTransformationResponse, err error) {
    if request == nil {
        request = NewUpdateTransformationRequest()
    }
    
    response = NewUpdateTransformationResponse()
    err = c.Send(request, response)
    return
}

// UpdateTransformation
// 用于更新转换器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"
//  INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"
//  INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"
//  INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"
//  INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"
//  INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"
//  RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"
func (c *Client) UpdateTransformationWithContext(ctx context.Context, request *UpdateTransformationRequest) (response *UpdateTransformationResponse, err error) {
    if request == nil {
        request = NewUpdateTransformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateTransformationResponse()
    err = c.Send(request, response)
    return
}
