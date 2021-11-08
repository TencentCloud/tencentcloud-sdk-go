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
    if request == nil {
        request = NewCreateCommandRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) CreateInvoker(request *CreateInvokerRequest) (response *CreateInvokerResponse, err error) {
    if request == nil {
        request = NewCreateInvokerRequest()
    }
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
    if request == nil {
        request = NewDeleteCommandRequest()
    }
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteInvoker(request *DeleteInvokerRequest) (response *DeleteInvokerResponse, err error) {
    if request == nil {
        request = NewDeleteInvokerRequest()
    }
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
    if request == nil {
        request = NewDescribeAutomationAgentStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeCommandsRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONID = "InvalidParameterValue.InvalidInvocationId"
//  INVALIDPARAMETERVALUE_INVALIDINVOCATIONTASKID = "InvalidParameterValue.InvalidInvocationTaskId"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocationTasks(request *DescribeInvocationTasksRequest) (response *DescribeInvocationTasksResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationTasksRequest()
    }
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
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationsRequest()
    }
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
func (c *Client) DescribeInvokerRecords(request *DescribeInvokerRecordsRequest) (response *DescribeInvokerRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInvokerRecordsRequest()
    }
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
func (c *Client) DescribeInvokers(request *DescribeInvokersRequest) (response *DescribeInvokersResponse, err error) {
    if request == nil {
        request = NewDescribeInvokersRequest()
    }
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
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableInvoker(request *DisableInvokerRequest) (response *DisableInvokerResponse, err error) {
    if request == nil {
        request = NewDisableInvokerRequest()
    }
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
    if request == nil {
        request = NewEnableInvokerRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTUNSUPPORTEDCOMMANDTYPE = "InvalidParameterValue.AgentUnsupportedCommandType"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDID = "InvalidParameterValue.InvalidCommandId"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
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
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) InvokeCommand(request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
    if request == nil {
        request = NewInvokeCommandRequest()
    }
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
func (c *Client) ModifyCommand(request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
    if request == nil {
        request = NewModifyCommandRequest()
    }
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
//  INVALIDPARAMETERVALUE_PARAMETERDISABLED = "InvalidParameterValue.ParameterDisabled"
//  INVALIDPARAMETERVALUE_PARAMETERINVALIDJSONFORMAT = "InvalidParameterValue.ParameterInvalidJsonFormat"
//  RESOURCENOTFOUND_COMMANDNOTFOUND = "ResourceNotFound.CommandNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
func (c *Client) ModifyInvoker(request *ModifyInvokerRequest) (response *ModifyInvokerResponse, err error) {
    if request == nil {
        request = NewModifyInvokerRequest()
    }
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
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
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
    if request == nil {
        request = NewPreviewReplacedCommandContentRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_COMMANDCONTENTINVALID = "InvalidParameterValue.CommandContentInvalid"
//  INVALIDPARAMETERVALUE_COMMANDNAMEDUPLICATED = "InvalidParameterValue.CommandNameDuplicated"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCE = "InvalidParameterValue.InconsistentInstance"
//  INVALIDPARAMETERVALUE_INVALIDCOMMANDNAME = "InvalidParameterValue.InvalidCommandName"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  INVALIDPARAMETERVALUE_INVALIDWORKINGDIRECTORY = "InvalidParameterValue.InvalidWorkingDirectory"
//  INVALIDPARAMETERVALUE_LACKOFPARAMETERINFO = "InvalidParameterValue.LackOfParameterInfo"
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
//  RESOURCEUNAVAILABLE_AGENTNOTINSTALLED = "ResourceUnavailable.AgentNotInstalled"
//  RESOURCEUNAVAILABLE_AGENTSTATUSNOTONLINE = "ResourceUnavailable.AgentStatusNotOnline"
//  RESOURCEUNAVAILABLE_INSTANCESTATENOTRUNNING = "ResourceUnavailable.InstanceStateNotRunning"
//  UNAUTHORIZEDOPERATION_CAMAUTHFAILED = "UnauthorizedOperation.CamAuthFailed"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
func (c *Client) RunCommand(request *RunCommandRequest) (response *RunCommandResponse, err error) {
    if request == nil {
        request = NewRunCommandRequest()
    }
    response = NewRunCommandResponse()
    err = c.Send(request, response)
    return
}
