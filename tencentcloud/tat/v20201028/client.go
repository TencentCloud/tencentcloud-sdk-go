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


func NewCancelInvocationRequest() (request *CancelInvocationRequest) {
    request = &CancelInvocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "CancelInvocation")
    
    
    return
}

func NewCancelInvocationResponse() (response *CancelInvocationResponse) {
    response = &CancelInvocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelInvocation
// 取消一台或多台CVM实例执行的命令
//
// 
//
// * 如果命令还未下发到agent，任务状态处于处于PENDING、DELIVERING、DELIVER_DELAYED，取消后任务状态是CANCELLED
//
// * 如果命令已下发到agent，任务状态处于RUNNING， 取消后任务状态是TERMINATED
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEISNOTRELATEDTOINVOCATION = "InvalidParameterValue.InstanceIsNotRelatedToInvocation"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVOCATIONNOTFOUND = "ResourceNotFound.InvocationNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CancelInvocation(request *CancelInvocationRequest) (response *CancelInvocationResponse, err error) {
    return c.CancelInvocationWithContext(context.Background(), request)
}

// CancelInvocation
// 取消一台或多台CVM实例执行的命令
//
// 
//
// * 如果命令还未下发到agent，任务状态处于处于PENDING、DELIVERING、DELIVER_DELAYED，取消后任务状态是CANCELLED
//
// * 如果命令已下发到agent，任务状态处于RUNNING， 取消后任务状态是TERMINATED
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEISNOTRELATEDTOINVOCATION = "InvalidParameterValue.InstanceIsNotRelatedToInvocation"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVOCATIONNOTFOUND = "ResourceNotFound.InvocationNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CancelInvocationWithContext(ctx context.Context, request *CancelInvocationRequest) (response *CancelInvocationResponse, err error) {
    if request == nil {
        request = NewCancelInvocationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelInvocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelInvocationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommandRequest() (request *CreateCommandRequest) {
    request = &CreateCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "CreateCommand")
    
    
    return
}

func NewCreateCommandResponse() (response *CreateCommandResponse) {
    response = &CreateCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCommand
// 此接口用于创建命令。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) CreateCommand(request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
    return c.CreateCommandWithContext(context.Background(), request)
}

// CreateCommand
// 此接口用于创建命令。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) CreateCommandWithContext(ctx context.Context, request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
    if request == nil {
        request = NewCreateCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCommandResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvokerRequest() (request *CreateInvokerRequest) {
    request = &CreateInvokerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "CreateInvoker")
    
    
    return
}

func NewCreateInvokerResponse() (response *CreateInvokerResponse) {
    response = &CreateInvokerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInvoker
// 此接口用于创建执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDCRONEXPRESSION = "InvalidParameterValue.InvalidCronExpression"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVOKETIMEEXPIRED = "InvalidParameterValue.InvokeTimeExpired"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
func (c *Client) CreateInvoker(request *CreateInvokerRequest) (response *CreateInvokerResponse, err error) {
    return c.CreateInvokerWithContext(context.Background(), request)
}

// CreateInvoker
// 此接口用于创建执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDCRONEXPRESSION = "InvalidParameterValue.InvalidCronExpression"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVOKETIMEEXPIRED = "InvalidParameterValue.InvokeTimeExpired"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
func (c *Client) CreateInvokerWithContext(ctx context.Context, request *CreateInvokerRequest) (response *CreateInvokerResponse, err error) {
    if request == nil {
        request = NewCreateInvokerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInvoker require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInvokerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCommandRequest() (request *DeleteCommandRequest) {
    request = &DeleteCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DeleteCommand")
    
    
    return
}

func NewDeleteCommandResponse() (response *DeleteCommandResponse) {
    response = &DeleteCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCommand
// 此接口用于删除命令。
//
// 如果命令与执行器关联，则无法被删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCEUNAVAILABLE_COMMANDINEXECUTING = "ResourceUnavailable.CommandInExecuting"
//  RESOURCEUNAVAILABLE_COMMANDININVOKER = "ResourceUnavailable.CommandInInvoker"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DeleteCommand(request *DeleteCommandRequest) (response *DeleteCommandResponse, err error) {
    return c.DeleteCommandWithContext(context.Background(), request)
}

// DeleteCommand
// 此接口用于删除命令。
//
// 如果命令与执行器关联，则无法被删除。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCEUNAVAILABLE_COMMANDINEXECUTING = "ResourceUnavailable.CommandInExecuting"
//  RESOURCEUNAVAILABLE_COMMANDININVOKER = "ResourceUnavailable.CommandInInvoker"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DeleteCommandWithContext(ctx context.Context, request *DeleteCommandRequest) (response *DeleteCommandResponse, err error) {
    if request == nil {
        request = NewDeleteCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInvokerRequest() (request *DeleteInvokerRequest) {
    request = &DeleteInvokerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DeleteInvoker")
    
    
    return
}

func NewDeleteInvokerResponse() (response *DeleteInvokerResponse) {
    response = &DeleteInvokerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInvoker
// 此接口用于删除执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteInvoker(request *DeleteInvokerRequest) (response *DeleteInvokerResponse, err error) {
    return c.DeleteInvokerWithContext(context.Background(), request)
}

// DeleteInvoker
// 此接口用于删除执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteInvokerWithContext(ctx context.Context, request *DeleteInvokerRequest) (response *DeleteInvokerResponse, err error) {
    if request == nil {
        request = NewDeleteInvokerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInvoker require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInvokerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutomationAgentStatusRequest() (request *DescribeAutomationAgentStatusRequest) {
    request = &DescribeAutomationAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeAutomationAgentStatus")
    
    
    return
}

func NewDescribeAutomationAgentStatusResponse() (response *DescribeAutomationAgentStatusResponse) {
    response = &DescribeAutomationAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutomationAgentStatus
// 此接口用于查询自动化助手客户端的状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeAutomationAgentStatus(request *DescribeAutomationAgentStatusRequest) (response *DescribeAutomationAgentStatusResponse, err error) {
    return c.DescribeAutomationAgentStatusWithContext(context.Background(), request)
}

// DescribeAutomationAgentStatus
// 此接口用于查询自动化助手客户端的状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeAutomationAgentStatusWithContext(ctx context.Context, request *DescribeAutomationAgentStatusRequest) (response *DescribeAutomationAgentStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAutomationAgentStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutomationAgentStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutomationAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommandsRequest() (request *DescribeCommandsRequest) {
    request = &DescribeCommandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeCommands")
    
    
    return
}

func NewDescribeCommandsResponse() (response *DescribeCommandsResponse) {
    response = &DescribeCommandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCommands
// 此接口用于查询命令详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeCommands(request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
    return c.DescribeCommandsWithContext(context.Background(), request)
}

// DescribeCommands
// 此接口用于查询命令详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeCommandsWithContext(ctx context.Context, request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
    if request == nil {
        request = NewDescribeCommandsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommands require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommandsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationTasksRequest() (request *DescribeInvocationTasksRequest) {
    request = &DescribeInvocationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvocationTasks")
    
    
    return
}

func NewDescribeInvocationTasksResponse() (response *DescribeInvocationTasksResponse) {
    response = &DescribeInvocationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInvocationTasks
// 此接口用于查询执行任务详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONTASKID = "InvalidParameterValue.InvalidInvocationTaskId"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocationTasks(request *DescribeInvocationTasksRequest) (response *DescribeInvocationTasksResponse, err error) {
    return c.DescribeInvocationTasksWithContext(context.Background(), request)
}

// DescribeInvocationTasks
// 此接口用于查询执行任务详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONTASKID = "InvalidParameterValue.InvalidInvocationTaskId"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocationTasksWithContext(ctx context.Context, request *DescribeInvocationTasksRequest) (response *DescribeInvocationTasksResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
    request = &DescribeInvocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvocations")
    
    
    return
}

func NewDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
    response = &DescribeInvocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInvocations
// 此接口用于查询执行活动详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FILTERVALUEEXCEEDED = "LimitExceeded.FilterValueExceeded"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    return c.DescribeInvocationsWithContext(context.Background(), request)
}

// DescribeInvocations
// 此接口用于查询执行活动详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FILTERVALUEEXCEEDED = "LimitExceeded.FilterValueExceeded"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocationsWithContext(ctx context.Context, request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvokerRecordsRequest() (request *DescribeInvokerRecordsRequest) {
    request = &DescribeInvokerRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvokerRecords")
    
    
    return
}

func NewDescribeInvokerRecordsResponse() (response *DescribeInvokerRecordsResponse) {
    response = &DescribeInvokerRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInvokerRecords
// 此接口用于查询执行器的执行记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
func (c *Client) DescribeInvokerRecords(request *DescribeInvokerRecordsRequest) (response *DescribeInvokerRecordsResponse, err error) {
    return c.DescribeInvokerRecordsWithContext(context.Background(), request)
}

// DescribeInvokerRecords
// 此接口用于查询执行器的执行记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
func (c *Client) DescribeInvokerRecordsWithContext(ctx context.Context, request *DescribeInvokerRecordsRequest) (response *DescribeInvokerRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInvokerRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvokerRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvokerRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvokersRequest() (request *DescribeInvokersRequest) {
    request = &DescribeInvokersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeInvokers")
    
    
    return
}

func NewDescribeInvokersResponse() (response *DescribeInvokersResponse) {
    response = &DescribeInvokersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInvokers
// 此接口用于查询执行器信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
func (c *Client) DescribeInvokers(request *DescribeInvokersRequest) (response *DescribeInvokersResponse, err error) {
    return c.DescribeInvokersWithContext(context.Background(), request)
}

// DescribeInvokers
// 此接口用于查询执行器信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
func (c *Client) DescribeInvokersWithContext(ctx context.Context, request *DescribeInvokersRequest) (response *DescribeInvokersResponse, err error) {
    if request == nil {
        request = NewDescribeInvokersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvokers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvokersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 此接口用于查询 TAT 产品后台地域列表。
//
// RegionState 为 AVAILABLE，代表该地域的 TAT 后台服务已经可用；未返回，代表该地域的 TAT 后台服务尚不可用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 此接口用于查询 TAT 产品后台地域列表。
//
// RegionState 为 AVAILABLE，代表该地域的 TAT 后台服务已经可用；未返回，代表该地域的 TAT 后台服务尚不可用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableInvokerRequest() (request *DisableInvokerRequest) {
    request = &DisableInvokerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "DisableInvoker")
    
    
    return
}

func NewDisableInvokerResponse() (response *DisableInvokerResponse) {
    response = &DisableInvokerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableInvoker
// 此接口用于停止执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableInvoker(request *DisableInvokerRequest) (response *DisableInvokerResponse, err error) {
    return c.DisableInvokerWithContext(context.Background(), request)
}

// DisableInvoker
// 此接口用于停止执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableInvokerWithContext(ctx context.Context, request *DisableInvokerRequest) (response *DisableInvokerResponse, err error) {
    if request == nil {
        request = NewDisableInvokerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableInvoker require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableInvokerResponse()
    err = c.Send(request, response)
    return
}

func NewEnableInvokerRequest() (request *EnableInvokerRequest) {
    request = &EnableInvokerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "EnableInvoker")
    
    
    return
}

func NewEnableInvokerResponse() (response *EnableInvokerResponse) {
    response = &EnableInvokerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableInvoker
// 此接口用于启用执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableInvoker(request *EnableInvokerRequest) (response *EnableInvokerResponse, err error) {
    return c.EnableInvokerWithContext(context.Background(), request)
}

// EnableInvoker
// 此接口用于启用执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableInvokerWithContext(ctx context.Context, request *EnableInvokerRequest) (response *EnableInvokerResponse, err error) {
    if request == nil {
        request = NewEnableInvokerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableInvoker require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableInvokerResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeCommandRequest() (request *InvokeCommandRequest) {
    request = &InvokeCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "InvokeCommand")
    
    
    return
}

func NewInvokeCommandResponse() (response *InvokeCommandResponse) {
    response = &InvokeCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InvokeCommand
// 在指定的实例上触发命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
//
// 
//
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
//
// * 如果命令类型与 agent 运行环境不符，返回失败
//
// * 指定的实例需要处于 VPC 网络
//
// * 指定的实例需要处于 RUNNING 状态
//
// * 不可同时指定 CVM 和 Lighthouse
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMERROR = "FailedOperation.CVMError"
//  FAILEDOPERATION_LIGHTHOUSEERROR = "FailedOperation.LighthouseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUsername"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  RESOURCEUNAVAILABLE_LIGHTHOUSEUNSUPPORTEDREGION = "ResourceUnavailable.LighthouseUnsupportedRegion"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) InvokeCommand(request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
    return c.InvokeCommandWithContext(context.Background(), request)
}

// InvokeCommand
// 在指定的实例上触发命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
//
// 
//
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
//
// * 如果命令类型与 agent 运行环境不符，返回失败
//
// * 指定的实例需要处于 VPC 网络
//
// * 指定的实例需要处于 RUNNING 状态
//
// * 不可同时指定 CVM 和 Lighthouse
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMERROR = "FailedOperation.CVMError"
//  FAILEDOPERATION_LIGHTHOUSEERROR = "FailedOperation.LighthouseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUsername"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  RESOURCEUNAVAILABLE_LIGHTHOUSEUNSUPPORTEDREGION = "ResourceUnavailable.LighthouseUnsupportedRegion"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) InvokeCommandWithContext(ctx context.Context, request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
    if request == nil {
        request = NewInvokeCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeCommandResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCommandRequest() (request *ModifyCommandRequest) {
    request = &ModifyCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "ModifyCommand")
    
    
    return
}

func NewModifyCommandResponse() (response *ModifyCommandResponse) {
    response = &ModifyCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCommand
// 此接口用于修改命令。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) ModifyCommand(request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
    return c.ModifyCommandWithContext(context.Background(), request)
}

// ModifyCommand
// 此接口用于修改命令。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) ModifyCommandWithContext(ctx context.Context, request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
    if request == nil {
        request = NewModifyCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCommandResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInvokerRequest() (request *ModifyInvokerRequest) {
    request = &ModifyInvokerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "ModifyInvoker")
    
    
    return
}

func NewModifyInvokerResponse() (response *ModifyInvokerResponse) {
    response = &ModifyInvokerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInvoker
// 此接口用于修改执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
func (c *Client) ModifyInvoker(request *ModifyInvokerRequest) (response *ModifyInvokerResponse, err error) {
    return c.ModifyInvokerWithContext(context.Background(), request)
}

// ModifyInvoker
// 此接口用于修改执行器。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDINVOKERID = "InvalidParameterValue.InvalidInvokerId"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
func (c *Client) ModifyInvokerWithContext(ctx context.Context, request *ModifyInvokerRequest) (response *ModifyInvokerResponse, err error) {
    if request == nil {
        request = NewModifyInvokerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInvoker require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInvokerResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewReplacedCommandContentRequest() (request *PreviewReplacedCommandContentRequest) {
    request = &PreviewReplacedCommandContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "PreviewReplacedCommandContent")
    
    
    return
}

func NewPreviewReplacedCommandContentResponse() (response *PreviewReplacedCommandContentResponse) {
    response = &PreviewReplacedCommandContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PreviewReplacedCommandContent
// 此接口用于预览自定义参数替换后的命令内容。不会触发真实执行。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETER_PARAMETERNAMEDUPLICATED = "InvalidParameter.ParameterNameDuplicated"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERS = "InvalidParameterValue.LackOfParameters"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) PreviewReplacedCommandContent(request *PreviewReplacedCommandContentRequest) (response *PreviewReplacedCommandContentResponse, err error) {
    return c.PreviewReplacedCommandContentWithContext(context.Background(), request)
}

// PreviewReplacedCommandContent
// 此接口用于预览自定义参数替换后的命令内容。不会触发真实执行。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CONFLICTPARAMETER = "InvalidParameter.ConflictParameter"
//  INVALIDPARAMETER_PARAMETERNAMEDUPLICATED = "InvalidParameter.ParameterNameDuplicated"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERS = "InvalidParameterValue.LackOfParameters"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) PreviewReplacedCommandContentWithContext(ctx context.Context, request *PreviewReplacedCommandContentRequest) (response *PreviewReplacedCommandContentResponse, err error) {
    if request == nil {
        request = NewPreviewReplacedCommandContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewReplacedCommandContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewReplacedCommandContentResponse()
    err = c.Send(request, response)
    return
}

func NewRunCommandRequest() (request *RunCommandRequest) {
    request = &RunCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tat", APIVersion, "RunCommand")
    
    
    return
}

func NewRunCommandResponse() (response *RunCommandResponse) {
    response = &RunCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunCommand
// 执行命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
//
// 
//
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
//
// * 如果命令类型与 agent 运行环境不符，返回失败
//
// * 指定的实例需要处于 VPC 网络
//
// * 指定的实例需要处于 `RUNNING` 状态
//
// * 不可同时指定 CVM 和 Lighthouse
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMERROR = "FailedOperation.CVMError"
//  FAILEDOPERATION_LIGHTHOUSEERROR = "FailedOperation.LighthouseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUsername"
//  INVALIDPARAMETER_PARAMETERNAMEDUPLICATED = "InvalidParameter.ParameterNameDuplicated"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  RESOURCEUNAVAILABLE_LIGHTHOUSEUNSUPPORTEDREGION = "ResourceUnavailable.LighthouseUnsupportedRegion"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
    return c.RunCommandWithContext(context.Background(), request)
}

// RunCommand
// 执行命令，调用成功返回执行活动ID（inv-xxxxxxxx），每个执行活动内部有一个或多个执行任务（invt-xxxxxxxx），每个执行任务代表命令在一台 CVM 或一台 Lighthouse 上的执行记录。
//
// 
//
// * 如果指定实例未安装 agent，或 agent 不在线，返回失败
//
// * 如果命令类型与 agent 运行环境不符，返回失败
//
// * 指定的实例需要处于 VPC 网络
//
// * 指定的实例需要处于 `RUNNING` 状态
//
// * 不可同时指定 CVM 和 Lighthouse
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMERROR = "FailedOperation.CVMError"
//  FAILEDOPERATION_LIGHTHOUSEERROR = "FailedOperation.LighthouseError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUsername"
//  INVALIDPARAMETER_PARAMETERNAMEDUPLICATED = "InvalidParameter.ParameterNameDuplicated"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSBUCKETURL = "InvalidParameterValue.InvalidOutputCOSBucketUrl"
//  INVALIDPARAMETERVALUE_INVALIDOUTPUTCOSKEYPREFIX = "InvalidParameterValue.InvalidOutputCOSKeyPrefix"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  INVALIDPARAMETERVALUE_PARAMETERKEYCONTAINSINVALIDCHAR = "InvalidParameterValue.ParameterKeyContainsInvalidChar"
//  INVALIDPARAMETERVALUE_PARAMETERKEYDUPLICATED = "InvalidParameterValue.ParameterKeyDuplicated"
//  INVALIDPARAMETERVALUE_PARAMETERKEYLENEXCEEDED = "InvalidParameterValue.ParameterKeyLenExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNUMBEREXCEEDED = "InvalidParameterValue.ParameterNumberExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUENOTSTRING = "InvalidParameterValue.ParameterValueNotString"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SUPPORTPARAMETERSONLYIFENABLEPARAMETER = "InvalidParameterValue.SupportParametersOnlyIfEnableParameter"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  RESOURCEUNAVAILABLE_LIGHTHOUSEUNSUPPORTEDREGION = "ResourceUnavailable.LighthouseUnsupportedRegion"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) RunCommandWithContext(ctx context.Context, request *RunCommandRequest) (response *RunCommandResponse, err error) {
    if request == nil {
        request = NewRunCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunCommandResponse()
    err = c.Send(request, response)
    return
}
